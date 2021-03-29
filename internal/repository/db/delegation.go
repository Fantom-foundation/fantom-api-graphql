// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	colDelegations          = "delegations"
	fiDelegationPk          = "_id"
	fiDelegationAddress     = "addr"
	fiDelegationToValidator = "to"
	fiDelegationCreated     = "createdTime"
)

// initDelegationCollection initializes the delegation collection with
// indexes and additional parameters needed by the app.
func (db *MongoDbBridge) initDelegationCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index ordinal key along with the primary key
	unique := true
	ix = append(ix, mongo.IndexModel{
		Keys: bson.D{{fiDelegationAddress, 1}, {fiDelegationToValidator, 1}},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})

	// index delegator and receiving validator
	// ix = append(ix, mongo.IndexModel{Keys: bson.D{{fiDelegationAddress, 1}}})
	// ix = append(ix, mongo.IndexModel{Keys: bson.D{{fiDelegationToValidator, 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{fiDelegationCreated, -1}}})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for delegation collection; %s", err.Error())
	}

	// log we done that
	db.log.Debugf("delegation collection initialized")
}

// AddDelegation stores a delegation in the database if it doesn't exist.
func (db *MongoDbBridge) AddDelegation(dl *types.Delegation) error {
	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(colDelegations)

	// if the transaction already exists, we don't need to add it
	// just make sure the transaction accounts were processed
	if db.isDelegationKnown(col, dl) {
		return db.UpdateDelegation(dl)
	}

	// try to do the insert
	if _, err := col.InsertOne(context.Background(), dl); err != nil {
		db.log.Critical(err)
		return err
	}

	// make sure contracts collection is initialized
	if db.initDelegation != nil {
		db.initDelegation.Do(func() { db.initDelegationCollection(col); db.initDelegation = nil })
	}
	return nil
}

// UpdateDelegation updates the given delegation in database.
func (db *MongoDbBridge) UpdateDelegation(dl *types.Delegation) error {
	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(colDelegations)

	// try to update a delegation by replacing it in the database
	er, err := col.ReplaceOne(context.Background(), bson.D{{fiDelegationPk, dl.Uid()}}, dl)
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// do we actually have the document
	if 0 == er.MatchedCount {
		return fmt.Errorf("can not update, the delegation does not exist in database")
	}
	return nil
}

// isDelegationKnown checks if the given delegation exists in the database.
func (db *MongoDbBridge) isDelegationKnown(col *mongo.Collection, dl *types.Delegation) bool {
	// try to find the contract in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{
		{fiDelegationPk, dl.Uid()},
	}, options.FindOne().SetProjection(bson.D{
		{fiContractPk, true},
	}))

	// error on lookup?
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			return false
		}

		// inform that we can not get the PK; should not happen
		db.log.Errorf("can not get existing delegation pk; %s", sr.Err().Error())
		return false
	}

	return true
}

// DelegationsCount calculates total number of delegations in the database.
func (db *MongoDbBridge) DelegationsCount() (uint64, error) {
	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(colDelegations)

	// do the counting
	val, err := col.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		db.log.Errorf("can not count documents in delegations collection; %s", err.Error())
		return 0, err
	}

	return uint64(val), nil
}

// dlgListInit initializes list of delegations based on provided cursor, count, and filter.
func (db *MongoDbBridge) dlgListInit(col *mongo.Collection, cursor *string, count int32, filter *bson.D) (*types.DelegationList, error) {
	// make sure some filter is used
	if nil == filter {
		filter = &bson.D{}
	}

	// find how many transactions do we have in the database
	total, err := col.CountDocuments(context.Background(), *filter)
	if err != nil {
		db.log.Errorf("can not count delegations")
		return nil, err
	}

	// make the list and notify the size of it
	db.log.Debugf("found %d filtered delegations", total)
	list := types.DelegationList{
		Collection: make([]*types.Delegation, 0),
		Total:      uint64(total),
		First:      "",
		Last:       "",
		IsStart:    total == 0,
		IsEnd:      total == 0,
		Filter:     *filter,
	}

	// is the list non-empty? return the list with properly calculated range marks
	if 0 < total {
		return db.dlgListCollectRangeMarks(col, &list, cursor, count)
	}

	// this is an empty list
	db.log.Debug("empty delegations list created")
	return &list, nil
}

// trxListWithRangeMarks returns the transaction list with proper First/Last marks of the transaction range.
func (db *MongoDbBridge) dlgListCollectRangeMarks(col *mongo.Collection, list *types.DelegationList, cursor *string, count int32) (*types.DelegationList, error) {
	var err error

	// find out the cursor ordinal index
	if cursor == nil && count > 0 {
		// get the highest available pk
		list.First, err = db.dlgListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{fiDelegationCreated, -1}, {fiDelegationPk, -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available pk
		list.First, err = db.dlgListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{fiDelegationCreated, 1}, {fiDelegationPk, 1}}))
		list.IsEnd = true

	} else if cursor != nil {
		// the cursor itself is the starting point
		list.First, err = db.dlgListBorderPk(col,
			bson.D{{fiDelegationPk, *cursor}},
			options.FindOne())
	}

	// check the error
	if err != nil {
		db.log.Errorf("can not find the initial delegation")
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("delegation list initialized with PK %s", list.First)
	return list, nil
}

// dlgListBorderPk finds the top PK of the delegations collection based on given filter and options.
func (db *MongoDbBridge) dlgListBorderPk(col *mongo.Collection, filter bson.D, opt *options.FindOneOptions) (string, error) {
	// prep container
	var row struct {
		Value string `bson:"_id"`
	}

	// make sure we pull only what we need
	opt.SetProjection(bson.D{{fiDelegationPk, true}})
	sr := col.FindOne(context.Background(), filter, opt)

	// try to decode
	err := sr.Decode(&row)
	if err != nil {
		return "", err
	}

	return row.Value, nil
}

// dlgListFilter creates a filter for delegations list loading.
func (db *MongoDbBridge) dlgListFilter(cursor *string, count int32, list *types.DelegationList) *bson.D {
	// build an extended filter for the query; add PK (decoded cursor) to the original filter
	if cursor == nil {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: fiDelegationPk, Value: bson.D{{"$gte", list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: fiDelegationPk, Value: bson.D{{"$lte", list.First}}})
		}
	} else {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: fiDelegationPk, Value: bson.D{{"$gt", list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: fiDelegationPk, Value: bson.D{{"$lt", list.First}}})
		}
	}

	// return the new filter
	return &list.Filter
}

// dlgListOptions creates a filter options set for delegations list search.
func (db *MongoDbBridge) dlgListOptions(count int32) *options.FindOptions {
	// prep options
	opt := options.Find()

	// how to sort results in the collection
	// from high (new) to low (old) by default; reversed if loading from bottom
	sd := -1
	if count < 0 {
		sd = 1
	}

	// sort with the direction we want
	opt.SetSort(bson.D{{fiDelegationCreated, sd}, {fiDelegationPk, sd}})

	// prep the loading limit
	var limit = int64(count)
	if limit < 0 {
		limit = -limit
	}

	// apply the limit, try to get one more record so we can detect list end
	opt.SetLimit(limit + 1)
	return opt
}

// dlgListLoad load the initialized list of delegations from database.
func (db *MongoDbBridge) dlgListLoad(col *mongo.Collection, cursor *string, count int32, list *types.DelegationList) (err error) {
	// get the context for loader
	ctx := context.Background()

	// load the data
	ld, err := col.Find(ctx, db.dlgListFilter(cursor, count, list), db.dlgListOptions(count))
	if err != nil {
		db.log.Errorf("error loading delegations list; %s", err.Error())
		return err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			db.log.Errorf("error closing delegations list cursor; %s", err.Error())
		}
	}()

	// loop and load the list; we may not store the last value
	var dlg *types.Delegation
	for ld.Next(ctx) {
		// append a previous value to the list, if we have one
		if dlg != nil {
			list.Collection = append(list.Collection, dlg)
		}

		// try to decode the next row
		var row types.Delegation
		if err = ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode the delegation list row; %s", err.Error())
			return err
		}

		// use this row as the next item
		dlg = &row
	}

	// we should have all the items already; we may just need to check if a boundary was reached
	list.IsEnd = (cursor == nil && count < 0) || (count > 0 && int32(len(list.Collection)) < count)
	list.IsStart = (cursor == nil && count > 0) || (count < 0 && int32(len(list.Collection)) < -count)

	// add the last item as well if we hit the boundary
	if (list.IsStart || list.IsEnd) && dlg != nil {
		list.Collection = append(list.Collection, dlg)
	}
	return nil
}

// Delegations pulls list of transaction hashes starting on the specified cursor.
func (db *MongoDbBridge) Delegations(cursor *string, count int32, filter *bson.D) (*types.DelegationList, error) {
	// nothing to load?
	if count == 0 {
		return nil, fmt.Errorf("nothing to do, zero delegations requested")
	}

	// get the collection and context
	col := db.client.Database(db.dbName).Collection(colDelegations)

	// init the list
	list, err := db.dlgListInit(col, cursor, count, filter)
	if err != nil {
		db.log.Errorf("can not build delegation list; %s", err.Error())
		return nil, err
	}

	// load data if there are any
	if list.Total > 0 {
		err = db.dlgListLoad(col, cursor, count, list)
		if err != nil {
			db.log.Errorf("can not load delegation list from database; %s", err.Error())
			return nil, err
		}

		// reverse on negative so new-er delegations will be on top
		if count < 0 {
			list.Reverse()
		}
	}

	return list, nil
}

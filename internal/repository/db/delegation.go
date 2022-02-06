// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
	"time"
)

// colDelegations represents the name of the delegations collection
const colDelegations = "delegations"

// ErrUnknownDelegation represents an error given on an unknown delegation update attempt.
var ErrUnknownDelegation = fmt.Errorf("unknown delegation")

// initDelegationCollection initializes the delegation collection with
// indexes and additional parameters needed by the app.
func (db *MongoDbBridge) initDelegationCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index delegation address and the validator; this is how we find a specific unique delegation
	unique := true
	ix = append(ix, mongo.IndexModel{
		Keys: bson.D{{Key: types.FiDelegationAddress, Value: 1}, {Key: types.FiDelegationToValidator, Value: 1}},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})

	// index delegator, receiving validator, and creation time stamp
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiDelegationAddress, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiDelegationToValidator, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiDelegationOrdinal, Value: -1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiDelegationStamp, Value: -1}}})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for delegation collection; %s", err.Error())
	}

	// log we're done that
	db.log.Debugf("delegation collection initialized")
}

// Delegation returns details of a delegation from an address to a validator ID.
func (db *MongoDbBridge) Delegation(addr *common.Address, valID *hexutil.Big) (*types.Delegation, error) {
	// get the collection for delegations
	col := db.client.Database(db.dbName).Collection(colDelegations)

	// try to find the delegation in the database
	sr := col.FindOne(context.Background(), bson.D{
		{Key: types.FiDelegationAddress, Value: addr.String()},
		{Key: types.FiDelegationToValidator, Value: valID.String()},
	})

	// do we have the data?
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			db.log.Errorf("delegation %s to #%d not found", addr.String(), valID.ToInt().Uint64())
			return nil, ErrUnknownDelegation
		}
		return nil, sr.Err()
	}

	// try to decode
	var dlg types.Delegation
	if err := sr.Decode(&dlg); err != nil {
		return nil, err
	}
	return &dlg, nil
}

// AddDelegation stores a delegation in the database if it doesn't exist.
func (db *MongoDbBridge) AddDelegation(dl *types.Delegation) error {
	// get the collection for delegations
	col := db.client.Database(db.dbName).Collection(colDelegations)

	// if the delegation already exists, update it with the new data
	if db.isDelegationKnown(col, dl) {
		return db.UpdateDelegation(dl)
	}

	// try to do the insert
	if _, err := col.InsertOne(context.Background(), dl); err != nil {
		db.log.Criticalf("can not add delegation %s to %d; %s", dl.Address.String(), dl.ToStakerId.ToInt().Uint64(), err.Error())
		return err
	}

	// make sure delegation collection is initialized
	if db.initDelegations != nil {
		db.initDelegations.Do(func() { db.initDelegationCollection(col); db.initDelegations = nil })
	}
	return nil
}

// UpdateDelegation updates the given delegation in database.
func (db *MongoDbBridge) UpdateDelegation(dl *types.Delegation) error {
	// get the collection for delegations
	col := db.client.Database(db.dbName).Collection(colDelegations)

	// calculate the value to 9 digits (and 18 billions remain available)
	val := new(big.Int).Div(dl.AmountDelegated.ToInt(), types.DelegationDecimalsCorrection).Uint64()

	// notify
	db.log.Debugf("updating delegation %s to #%d value to %d",
		dl.Address.String(), dl.ToStakerId.ToInt().Uint64(), val)

	// try to update a delegation by replacing it in the database
	// we use address and validator ID to identify unique delegation
	er, err := col.UpdateOne(context.Background(), bson.D{
		{Key: types.FiDelegationAddress, Value: dl.Address.String()},
		{Key: types.FiDelegationToValidator, Value: dl.ToStakerId.String()},
	}, bson.D{{Key: "$set", Value: bson.D{
		{Key: types.FiDelegationOrdinal, Value: dl.OrdinalIndex()},
		{Key: types.FiDelegationStamp, Value: time.Unix(int64(dl.CreatedTime), 0)},
		{Key: types.FiDelegationTransaction, Value: dl.Transaction.String()},
		{Key: types.FiDelegationToValidatorAddress, Value: dl.ToStakerAddress.String()},
		{Key: types.FiDelegationAmountActive, Value: dl.AmountDelegated.String()},
		{Key: types.FiDelegationValue, Value: val},
	}}}, new(options.UpdateOptions).SetUpsert(true))
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// do we actually have the document
	if 0 == er.MatchedCount {
		db.log.Errorf("delegation %s to %d not found", dl.Address.String(), dl.ToStakerId.ToInt().Uint64())
	}

	// make sure delegation collection is initialized
	if db.initDelegations != nil {
		db.initDelegations.Do(func() { db.initDelegationCollection(col); db.initDelegations = nil })
	}
	return nil
}

// UpdateDelegationBalance updates the given delegation active balance in database to the given amount.
func (db *MongoDbBridge) UpdateDelegationBalance(addr *common.Address, valID *hexutil.Big, amo *hexutil.Big) error {
	// get the collection for delegations
	col := db.client.Database(db.dbName).Collection(colDelegations)
	val := new(big.Int).Div(amo.ToInt(), types.DelegationDecimalsCorrection).Uint64()

	// notify
	db.log.Debugf("%s delegation to #%d value changed to %d", addr.String(), valID.ToInt().Uint64(), val)

	// update the transaction details
	ur, err := col.UpdateOne(context.Background(),
		bson.D{
			{Key: types.FiDelegationAddress, Value: addr.String()},
			{Key: types.FiDelegationToValidator, Value: valID.String()},
		},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: types.FiDelegationAmountActive, Value: amo.String()},
			{Key: types.FiDelegationValue, Value: val},
		}}})
	if err != nil {
		// log the issue
		db.log.Criticalf("delegation balance can not be updated; %s", err.Error())
		return err
	}

	// any match?
	if ur.MatchedCount == 0 {
		db.log.Errorf("delegation %s to %d not found", addr.String(), valID.ToInt().Uint64())
		return ErrUnknownDelegation
	}
	return nil
}

// isDelegationKnown checks if the given delegation exists in the database.
func (db *MongoDbBridge) isDelegationKnown(col *mongo.Collection, dl *types.Delegation) bool {
	// try to find the delegation in the database
	sr := col.FindOne(context.Background(), bson.D{
		{Key: types.FiDelegationAddress, Value: dl.Address.String()},
		{Key: types.FiDelegationToValidator, Value: dl.ToStakerId.String()},
	}, options.FindOne().SetProjection(bson.D{
		{Key: types.FiDelegationPk, Value: true},
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

// DelegationsCountFiltered calculates total number of delegations in the database for the given filter.
func (db *MongoDbBridge) DelegationsCountFiltered(filter *bson.D) (uint64, error) {
	return db.CountFiltered(db.client.Database(db.dbName).Collection(colDelegations), filter)
}

// DelegationsCount calculates total number of delegations in the database.
func (db *MongoDbBridge) DelegationsCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(colDelegations))
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
		First:      0,
		Last:       0,
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

// trxListWithRangeMarks returns a list of delegations with proper First/Last marks.
func (db *MongoDbBridge) dlgListCollectRangeMarks(col *mongo.Collection, list *types.DelegationList, cursor *string, count int32) (*types.DelegationList, error) {
	var err error

	// find out the cursor ordinal index
	if cursor == nil && count > 0 {
		// get the highest available pk
		list.First, err = db.dlgListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: types.FiDelegationOrdinal, Value: -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available pk
		list.First, err = db.dlgListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: types.FiDelegationOrdinal, Value: 1}}))
		list.IsEnd = true

	} else if cursor != nil {
		// decode the cursor
		id, err := primitive.ObjectIDFromHex(*cursor)
		if err != nil {
			db.log.Errorf("invalid delegation cursor ID; %s", err.Error())
			return nil, err
		}

		// look for the first ordinal to make sure it's there
		list.First, err = db.dlgListBorderPk(col,
			append(list.Filter, bson.E{Key: types.FiDelegationPk, Value: id}),
			options.FindOne())
	}

	// check the error
	if err != nil {
		db.log.Errorf("can not find the initial delegation; %s", err.Error())
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("delegation list starts from #%d", list.First)
	return list, nil
}

// dlgListBorderPk finds the top PK of the delegations collection based on given filter and options.
func (db *MongoDbBridge) dlgListBorderPk(col *mongo.Collection, filter bson.D, opt *options.FindOneOptions) (uint64, error) {
	// prep container
	var row struct {
		Value uint64 `bson:"orx"`
	}

	// make sure we pull only what we need
	opt.SetProjection(bson.D{{Key: types.FiDelegationOrdinal, Value: true}})
	sr := col.FindOne(context.Background(), filter, opt)

	// try to decode
	err := sr.Decode(&row)
	if err != nil {
		return 0, err
	}

	return row.Value, nil
}

// dlgListFilter creates a filter for delegations list loading.
func (db *MongoDbBridge) dlgListFilter(cursor *string, count int32, list *types.DelegationList) *bson.D {
	// build an extended filter for the query; add PK (decoded cursor) to the original filter
	if cursor == nil {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: types.FiDelegationOrdinal, Value: bson.D{{Key: "$lte", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: types.FiDelegationOrdinal, Value: bson.D{{Key: "$gte", Value: list.First}}})
		}
	} else {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: types.FiDelegationOrdinal, Value: bson.D{{Key: "$lt", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: types.FiDelegationOrdinal, Value: bson.D{{Key: "$gt", Value: list.First}}})
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
		count = -count
	}

	// sort with the direction we want
	opt.SetSort(bson.D{{Key: types.FiDelegationOrdinal, Value: sd}})

	// apply the limit, try to get one more record so we can detect list end
	opt.SetLimit(int64(count) + 1)
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

// Delegations pulls list of delegations starting at the specified cursor.
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
			count = -count
		}

		// cut the end?
		if len(list.Collection) > int(count) {
			list.Collection = list.Collection[:len(list.Collection)-1]
		}
	}

	return list, nil
}

// DelegationsAll pulls list of delegations for the given filter un-paged.
func (db *MongoDbBridge) DelegationsAll(filter *bson.D) ([]*types.Delegation, error) {
	// get the collection and context
	col := db.client.Database(db.dbName).Collection(colDelegations)
	list := make([]*types.Delegation, 0)
	ctx := context.Background()

	// load the data
	ld, err := col.Find(ctx, filter, options.Find().SetSort(bson.D{{Key: types.FiDelegationStamp, Value: -1}}))
	if err != nil {
		db.log.Errorf("error loading full delegations list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer db.closeCursor(ld)

	for ld.Next(ctx) {
		// try to decode the next row
		var row types.Delegation
		if err = ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode the full delegation list row; %s", err.Error())
			return nil, err
		}

		// use this row as the next item
		list = append(list, &row)
	}
	return list, nil
}

// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// colEpochs represents the name of the epochs collection in database.
	colEpochs = "epochs"

	// fiEpochPk is the name of the primary key of the collection.
	fiEpochPk = "_id"

	// fiEpochEndTime is the name of the epoch end field in the collection.
	fiEpochEndTime = "end"
)

// initEpochsCollection initializes the epochs collection with
// indexes and additional parameters needed by the app.
func (db *MongoDbBridge) initEpochsCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index ordinal key sorted from high to low since this is the way we usually list
	ix = append(ix, mongo.IndexModel{
		Keys:    bson.D{{Key: fiEpochEndTime, Value: -1}},
		Options: new(options.IndexOptions).SetUnique(true),
	})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for epoch collection; %s", err.Error())
	}
	db.log.Debugf("epochs collection initialized")
}

// AddEpoch stores an epoch reference in connected persistent storage.
func (db *MongoDbBridge) AddEpoch(e *types.Epoch) error {
	// do we have all needed data? we reject epochs without any stake
	if e == nil || e.EndTime == 0 || e.StakeTotalAmount.ToInt().Cmp(intZero) <= 0 {
		return fmt.Errorf("empty epoch received")
	}

	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(colEpochs)

	// if the transaction already exists, we don't need to add it
	// just make sure the transaction accounts were processed
	if db.isEpochKnown(col, e) {
		return nil
	}

	// try to do the insert
	if _, err := col.InsertOne(context.Background(), e); err != nil {
		db.log.Critical(err)
		return err
	}

	// make sure epochs collection is initialized
	if db.initEpochs != nil {
		db.initEpochs.Do(func() { db.initEpochsCollection(col); db.initEpochs = nil })
	}

	// log what we did
	db.log.Debugf("epoch #%d added to database", e.Id)
	return nil
}

// isEpochKnown checks if the given epoch has already been added to the database
func (db *MongoDbBridge) isEpochKnown(col *mongo.Collection, e *types.Epoch) bool {
	// try to find the epoch in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{
		{Key: fiEpochPk, Value: int64(e.Id)},
	}, options.FindOne().SetProjection(bson.D{
		{Key: fiEpochPk, Value: true},
	}))

	// error on lookup?
	if sr.Err() == nil {
		return true
	}

	db.log.Debugf("epoch #%d not found in database; %s", e.Id, sr.Err().Error())
	return false
}

// LastKnownEpoch provides the number of the newest epoch stored in the database.
func (db *MongoDbBridge) LastKnownEpoch() (uint64, error) {
	return db.epochListBorderPk(db.client.Database(db.dbName).Collection(colEpochs), options.FindOne().SetSort(bson.D{{Key: fiEpochEndTime, Value: -1}}))
}

// EpochsCount calculates total number of epochs in the database.
func (db *MongoDbBridge) EpochsCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(colEpochs))
}

// epochListInit initializes list of epochs based on provided cursor, count.
func (db *MongoDbBridge) epochListInit(col *mongo.Collection, cursor *string, count int32) (*types.EpochList, error) {
	// find how many transactions do we have in the database
	total, err := db.EpochsCount()
	if err != nil {
		db.log.Errorf("can not count epochs")
		return nil, err
	}

	// make the list and notify the size of it
	list := types.EpochList{
		Collection: make([]*types.Epoch, 0),
		Total:      total,
		First:      0,
		Last:       0,
		IsStart:    total == 0,
		IsEnd:      total == 0,
	}

	// is the list non-empty? return the list with properly calculated range marks
	if 0 < total {
		return db.epochListCollectRangeMarks(col, &list, cursor, count)
	}

	// this is an empty list
	db.log.Debug("empty epoch list created")
	return &list, nil
}

// epochListCollectRangeMarks returns a list of epochs with proper First/Last marks.
func (db *MongoDbBridge) epochListCollectRangeMarks(col *mongo.Collection, list *types.EpochList, cursor *string, count int32) (*types.EpochList, error) {
	var err error

	// find out the cursor ordinal index
	if cursor == nil && count > 0 {
		// get the highest available pk
		list.First, err = db.epochListBorderPk(col, options.FindOne().SetSort(bson.D{{Key: fiEpochEndTime, Value: -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available pk
		list.First, err = db.epochListBorderPk(col, options.FindOne().SetSort(bson.D{{Key: fiEpochEndTime, Value: 1}}))
		list.IsEnd = true

	} else if cursor != nil {
		// the cursor itself is the starting point
		list.First, err = hexutil.DecodeUint64(*cursor)
	}

	// check the error
	if err != nil {
		db.log.Errorf("can not find the initial epoch")
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("epoch list initialized with epoch #%d", list.First)
	return list, nil
}

// rewListBorderPk finds the top PK of the reward claims collection based on given filter and options.
func (db *MongoDbBridge) epochListBorderPk(col *mongo.Collection, opt *options.FindOneOptions) (uint64, error) {
	// prep container
	var row struct {
		Value uint64 `bson:"_id"`
	}

	// make sure we pull only what we need
	opt.SetProjection(bson.D{{Key: fiEpochPk, Value: true}})

	// try to decode
	sr := col.FindOne(context.Background(), bson.D{}, opt)
	err := sr.Decode(&row)
	if err != nil {
		return 0, err
	}
	return row.Value, nil
}

// epochListFilter creates a filter for epoch list loading.
func (db *MongoDbBridge) epochListFilter(cursor *string, count int32, list *types.EpochList) *bson.D {
	// build an extended filter for the query; add PK (decoded cursor) to the original filter
	if cursor == nil {
		if count > 0 {
			return &bson.D{{Key: fiEpochPk, Value: bson.D{{Key: "$lte", Value: list.First}}}}
		}
		return &bson.D{{Key: fiEpochPk, Value: bson.D{{Key: "$gte", Value: list.First}}}}
	}

	// with cursor provided we need to skip the identified line
	if count > 0 {
		return &bson.D{{Key: fiEpochPk, Value: bson.D{{Key: "$lt", Value: list.First}}}}
	}
	return &bson.D{{Key: fiEpochPk, Value: bson.D{{Key: "$gt", Value: list.First}}}}
}

// epochListOptions creates a filter options set for epochs list search.
func (db *MongoDbBridge) epochListOptions(count int32) *options.FindOptions {
	// prep options
	opt := options.Find()

	// how to sort results in the collection
	// from high (new) to low (old) by default; reversed if loading from bottom
	sd := -1
	if count < 0 {
		sd = 1
	}

	// sort with the direction we want
	opt.SetSort(bson.D{{Key: fiEpochEndTime, Value: sd}})

	// prep the loading limit
	var limit = int64(count)
	if limit < 0 {
		limit = -limit
	}

	// apply the limit, try to get one more record so we can detect list end
	opt.SetLimit(limit + 1)
	return opt
}

// epochListLoad loads the initialized list of epochs from database.
func (db *MongoDbBridge) epochListLoad(col *mongo.Collection, cursor *string, count int32, list *types.EpochList) (err error) {
	// get the context for loader
	ctx := context.Background()

	// load the data
	ld, err := col.Find(ctx, db.epochListFilter(cursor, count, list), db.epochListOptions(count))
	if err != nil {
		db.log.Errorf("error loading epochs list; %s", err.Error())
		return err
	}

	defer db.closeCursor(ld)

	// loop and load the list; we may not store the last value
	var e *types.Epoch
	for ld.Next(ctx) {
		// append a previous value to the list, if we have one
		if e != nil {
			list.Collection = append(list.Collection, e)
		}

		// try to decode the next row
		var row types.Epoch
		if err = ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode epoch list row; %s", err.Error())
			return err
		}

		// use this row as the next item
		e = &row
	}

	// we should have all the items already; we may just need to check if a boundary was reached
	list.IsEnd = (cursor == nil && count < 0) || (count > 0 && int32(len(list.Collection)) < count)
	list.IsStart = (cursor == nil && count > 0) || (count < 0 && int32(len(list.Collection)) < -count)

	// add the last item as well if we hit the boundary
	if (list.IsStart || list.IsEnd) && e != nil {
		list.Collection = append(list.Collection, e)
	}
	return nil
}

// Epochs pulls list of epochs starting at the specified cursor.
func (db *MongoDbBridge) Epochs(cursor *string, count int32) (*types.EpochList, error) {
	// nothing to load?
	if count == 0 {
		return nil, fmt.Errorf("nothing to do, zero epochs requested")
	}

	// get the collection and context
	col := db.client.Database(db.dbName).Collection(colEpochs)

	// init the list
	list, err := db.epochListInit(col, cursor, count)
	if err != nil {
		db.log.Errorf("can not build epoch list; %s", err.Error())
		return nil, err
	}

	// load data if there are any
	if list.Total > 0 {
		err = db.epochListLoad(col, cursor, count, list)
		if err != nil {
			db.log.Errorf("can not load epoch list; %s", err.Error())
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

// TreasuryTotal aggregates the total amount of treasury fee across all epochs.
func (db *MongoDbBridge) TreasuryTotal() (int64, error) {
	col := db.client.Database(db.dbName).Collection(colEpochs)

	// aggregate the total amount of burned native tokens
	cr, err := col.Aggregate(context.Background(), mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "amount", Value: bson.D{{Key: "$sum", Value: "$treasured"}}},
		}}},
	})
	if err != nil {
		db.log.Errorf("can not collect total treasure fee; %s", err.Error())
		return 0, err
	}

	defer db.closeCursor(cr)
	if !cr.Next(context.Background()) {
		return 0, fmt.Errorf("treasure fee aggregation failed")
	}

	var row struct {
		Amount int64 `bson:"amount"`
	}
	if err := cr.Decode(&row); err != nil {
		db.log.Errorf("can not decode treasure fee aggregation cursor; %s", err.Error())
		return 0, err
	}
	return row.Amount, nil
}

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
	"time"
)

// colBurns represents the name of the native FTM burns collection in database.
const colBurns = "burns"
const colBurnsAggregate = "burns_sum"

// burnBaseAggregateDate represents the timestamp of the materialised burn aggregate
var burnBaseAggregateDate = time.Unix(0, 0)

/*
Initialize aggregate table:
	db.burns.aggregate([
		{$group: {
			"_id": null,
			"amount": { $sum: "$amount"}
		}},
		{$project: {
			"_id": new ISODate("1970-01-01T00:00:00Z"),
			"amount": true
		}},
		{$merge: {
			"into": "burns_sum",
			"on": "_id",
			"whenMatched": "replace",
			"whenNotMatched": "insert"
		}}
	])
*/

// initBurnsCollection initializes the burn collection indexes.
func (db *MongoDbBridge) initBurnsCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index delegator + validator
	unique := true
	ix = append(ix, mongo.IndexModel{
		Keys: bson.D{
			{Key: "block", Value: 1},
		},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for withdrawals collection; %s", err.Error())
	}

	// log we are done that
	db.log.Debugf("burns collection initialized")
}

// BurnByBlock pulls a burn information for the given block number, if available.
func (db *MongoDbBridge) BurnByBlock(bn hexutil.Uint64) (*types.FtmBurn, error) {
	col := db.client.Database(db.dbName).Collection(colBurns)

	// try to find existing burn
	sr := col.FindOne(context.Background(), bson.D{{Key: "block", Value: bn}})
	if sr.Err() != nil {
		// if the burn has not been found, add this as a new one
		if sr.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}

		db.log.Errorf("could not load FTM burn at #%d; %s", bn, sr.Err())
		return nil, sr.Err()
	}

	// decode existing burn and update
	var ex types.FtmBurn
	if err := sr.Decode(&ex); err != nil {
		db.log.Errorf("could not decode FTM burn at #%d; %s", bn, sr.Err())
		return nil, sr.Err()
	}

	return &ex, nil
}

// StoreBurn stores the given native FTM burn record.
func (db *MongoDbBridge) StoreBurn(burn *types.FtmBurn) error {
	if burn == nil {
		return nil
	}

	col := db.client.Database(db.dbName).Collection(colBurns)

	// make sure burns collection is initialized
	if db.initBurns != nil {
		db.initBurns.Do(func() { db.initBurnsCollection(col); db.initBurns = nil })
	}

	// insert/update the burn data
	re, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: "block", Value: burn.BlockNumber}},
		bson.D{{Key: "$set", Value: burn}},
		options.Update().SetUpsert(true))
	if err != nil {
		db.log.Criticalf("could not update burn #%d; %s", burn.BlockNumber, err.Error())
		return err
	}

	if re.UpsertedCount > 0 {
		db.burnAddBurnValue(burn.Value())
	}
	return nil
}

// burnAddBurnValue adds the given value to the total burned amount.
func (db *MongoDbBridge) burnAddBurnValue(v int64) {
	col := db.client.Database(db.dbName).Collection(colBurnsAggregate)

	_, err := col.UpdateByID(context.Background(), burnBaseAggregateDate, bson.D{
		{Key: "$inc", Value: bson.D{{Key: "amount", Value: v}}},
	})
	if err != nil {
		db.log.Criticalf("could not update burned total; %s", err.Error())
	}
}

// burnAddTotalAggregate creates a container for the total aggregate.
func (db *MongoDbBridge) burnAddTotalAggregate() error {
	col := db.client.Database(db.dbName).Collection(colBurnsAggregate)
	_, err := col.UpdateByID(context.Background(), burnBaseAggregateDate, bson.D{
		{Key: "$setOnInsert", Value: bson.D{
			{Key: "_id", Value: burnBaseAggregateDate},
			{Key: "amount", Value: int64(0)},
		}},
	}, options.Update().SetUpsert(true))
	return err
}

// BurnCount estimates the number of burn records in the database.
func (db *MongoDbBridge) BurnCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(colBurns))
}

// BurnTotal aggregates the total amount of burned fee across all blocks.
func (db *MongoDbBridge) BurnTotal() (int64, error) {
	col := db.client.Database(db.dbName).Collection(colBurnsAggregate)

	sr := col.FindOne(context.Background(), bson.D{{Key: "_id", Value: burnBaseAggregateDate}})
	if sr.Err() != nil {
		db.log.Criticalf("could not get burned total; %s", sr.Err().Error())
		if sr.Err() == mongo.ErrNoDocuments {
			return 0, db.burnAddTotalAggregate()
		}

		return 0, sr.Err()
	}

	var out struct {
		Amount int64 `bson:"amount"`
	}
	if err := sr.Decode(&out); err != nil {
		return 0, err
	}
	return out.Amount, nil
}

// BurnTotalSlow aggregates the total amount of burned fee across all blocks.
func (db *MongoDbBridge) BurnTotalSlow() (int64, error) {
	col := db.client.Database(db.dbName).Collection(colBurns)

	// aggregate the total amount of burned native tokens
	cr, err := col.Aggregate(context.Background(), mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "amount", Value: bson.D{{Key: "$sum", Value: "$amount"}}},
		}}},
	})
	if err != nil {
		db.log.Errorf("can not collect total burned fee; %s", err.Error())
		return 0, err
	}

	defer db.closeCursor(cr)
	if !cr.Next(context.Background()) {
		return 0, fmt.Errorf("burned fee aggregation failed")
	}

	var row struct {
		Amount int64 `bson:"amount"`
	}
	if err := cr.Decode(&row); err != nil {
		db.log.Errorf("can not decode burned fee aggregation cursor; %s", err.Error())
		return 0, err
	}
	return row.Amount, nil
}

// BurnList provides list of native FTM burns per blocks stored in the persistent database.
func (db *MongoDbBridge) BurnList(count int64) ([]types.FtmBurn, error) {
	col := db.client.Database(db.dbName).Collection(colBurns)

	cr, err := col.Find(context.Background(), bson.D{}, options.Find().SetSort(bson.D{{Key: "block", Value: -1}}).SetLimit(count))
	if err != nil {
		db.log.Errorf("failed to load burns; %s", err.Error())
		return nil, err
	}
	defer db.closeCursor(cr)

	ctx := context.Background()
	list := make([]types.FtmBurn, 0, count)

	for cr.Next(ctx) {
		var row types.FtmBurn
		if err := cr.Decode(&row); err != nil {
			db.log.Errorf("failed to decode burn; %s", err.Error())
			continue
		}
		list = append(list, row)
	}

	return list, nil
}

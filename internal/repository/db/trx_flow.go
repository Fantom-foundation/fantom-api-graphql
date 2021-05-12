// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	// db.trx_volume.createIndex({"stamp": 1}, {unique: true})
	// coTransactionVolume represents the name of the trx flow collection.
	coTransactionVolume = "trx_volume"

	// fiTrxVolumePk name of the primary key of the transaction volume row.
	fiTrxVolumePk = "_id"

	// fiTrxVolumeStamp name of the field of the trx volume time stamp.
	fiTrxVolumeStamp = "stamp"
)

// TrxDailyFlowList loads a range of daily trx volumes from the database.
func (db *MongoDbBridge) TrxDailyFlowList(from *time.Time, to *time.Time) ([]*types.DailyTrxVolume, error) {
	// log what we do
	db.log.Debugf("loading trx flow between %s and %s", from.String(), to.String())

	// get the collection and context
	ctx := context.Background()
	col := db.client.Database(db.dbName).Collection(coTransactionVolume)

	// pull the data; make sure there is a limit to the range
	ld, err := col.Find(ctx, trxDailyFlowListFilter(from, to), options.Find().SetSort(bson.D{{fiTrxVolumePk, 1}}).SetLimit(365))
	if err != nil {
		db.log.Errorf("can not load daily flow; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err := ld.Close(ctx)
		if err != nil {
			db.log.Errorf("error closing daily flow list cursor; %s", err.Error())
		}
	}()

	// load the list
	return loadTrxDailyFlowList(ld)
}

// TrxRecentTrxSpeed provides the number of transaction per second on the defined range in seconds.
func (db *MongoDbBridge) TrxRecentTrxSpeed(sec int32) (float64, error) {
	// make sure the request makes sense and calculate the left boundary
	if sec < 60 {
		sec = 60
	}
	from := time.Now().UTC().Add(time.Duration(-sec) * time.Second)
	col := db.client.Database(db.dbName).Collection(coTransactions)

	// find how many transactions do we have in the database
	total, err := col.CountDocuments(context.Background(), bson.D{
		{fiTransactionTimeStamp, bson.D{
			{"$gte", from},
		}},
	})
	if err != nil {
		db.log.Errorf("can not count recent transactions")
		return 0, err
	}

	// any transactions at all?
	if total == 0 {
		return 0, nil
	}
	return float64(total) / float64(sec), nil
}

// trxDailyFlowListFilter creates a filter for loading trx flow data based on provided
// range dates.
func trxDailyFlowListFilter(from *time.Time, to *time.Time) *bson.D {
	// prep the filter
	filter := bson.D{}

	// add start filter
	if from != nil {
		filter = append(filter, bson.E{Key: fiTrxVolumeStamp, Value: bson.D{{Key: "$gte", Value: *from}}})
	}

	// add end filter
	if to != nil {
		filter = append(filter, bson.E{Key: fiTrxVolumeStamp, Value: bson.D{{Key: "$lte", Value: *to}}})
	}

	return &filter
}

// loadTrxDailyFlowList load the trx flow list from provided DB cursor.
func loadTrxDailyFlowList(ld *mongo.Cursor) ([]*types.DailyTrxVolume, error) {
	// prep the result list
	ctx := context.Background()
	list := make([]*types.DailyTrxVolume, 0)

	// loop and load
	for ld.Next(ctx) {
		// try to decode the next row
		var row types.DailyTrxVolume
		if err := ld.Decode(&row); err != nil {
			return nil, err
		}

		// we have one
		list = append(list, &row)
	}
	return list, nil
}

// TrxDailyFlowUpdate performs an update on the daily trx flow data
// for the given date range directly.
func (db *MongoDbBridge) TrxDailyFlowUpdate(from time.Time) error {
	// log what we do
	db.log.Noticef("updating trx flow after %s", from)

	// we aggregate transactions
	col := db.client.Database(db.dbName).Collection(coTransactions)

	// get the collection
	cr, err := col.Aggregate(context.Background(), mongo.Pipeline{
		{{"$match", bson.D{
			{"stamp", bson.D{{"$gte", from}}},
		}}},
		{{"$group", bson.D{
			{"_id", bson.D{
				{"$dateToString", bson.D{
					{"format", "%Y-%m-%d"},
					{"date", "$stamp"},
				}},
			}},
			{"volume", bson.D{{"$sum", "$amo"}}},
			{"value", bson.D{{"$sum", 1}}},
		}}},
		{{"$project", bson.D{
			{"stamp", bson.D{{"$toDate", "$_id"}}},
			{"volume", 1},
			{"value", 1},
		}}},
		{{"$merge", bson.D{
			{"into", "trx_volume"},
			{"on", "_id"},
			{"whenMatched", "replace"},
			{"whenNotMatched", "insert"},
		}}},
	})
	if err != nil {
		db.log.Errorf("can not update trx flow; %s", err.Error())
		return err
	}

	// close the cursor, we don't really need the data
	if err := cr.Close(context.Background()); err != nil {
		db.log.Errorf("can not close aggregate cursor; %s", err.Error())
	}
	return nil
}

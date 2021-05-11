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

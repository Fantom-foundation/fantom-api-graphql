// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	// colEpochs represents the name of the epochs' collection in database.
	colGasPrice = "gas_price"
)

// gasPriceCollectionIndexes provides a list of indexes expected to exist on the prices' collection.
func gasPriceCollectionIndexes() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	ixGasPriceFrom := "ix_gas_price_from"
	ix[0] = mongo.IndexModel{
		Keys:    bson.D{{Key: types.FiGasPriceTimeFrom, Value: 1}},
		Options: &options.IndexOptions{Name: &ixGasPriceFrom},
	}

	ixGasPriceTo := "ix_gas_price_to"
	ix[1] = mongo.IndexModel{
		Keys:    bson.D{{Key: types.FiGasPriceTimeTo, Value: 1}},
		Options: &options.IndexOptions{Name: &ixGasPriceTo},
	}

	return ix
}

// AddGasPricePeriod stores a new record for the gas price evaluation
// into the persistent collection.
func (db *MongoDbBridge) AddGasPricePeriod(gp *types.GasPricePeriod) error {
	// do we have anything to store at all?
	if gp == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(colGasPrice)

	// try to do the insert
	if _, err := col.InsertOne(context.Background(), gp); err != nil {
		db.log.Errorf("can not store gas price value; %s", err)
		return err
	}

	return nil
}

// GasPricePeriodCount calculates total number of gas price period records in the database.
func (db *MongoDbBridge) GasPricePeriodCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(colGasPrice))
}

// GasPriceTicks provides a list of gas price ticks for the given time period.
func (db *MongoDbBridge) GasPriceTicks(from *time.Time, to *time.Time) ([]types.GasPricePeriod, error) {
	// get the collection
	col := db.client.Database(db.dbName).Collection(colGasPrice)

	// find ticks inside the date/time range
	cursor, err := col.Find(context.Background(), bson.D{
		{Key: "from", Value: bson.D{{Key: "$gte", Value: from}}},
		{Key: "to", Value: bson.D{{Key: "$lte", Value: to}}},
	}, options.Find().SetSort(bson.D{{Key: "from", Value: 1}}))
	if err != nil {
		db.log.Errorf("can not pull gas price ticks; %s", err.Error())
		return nil, err
	}

	// make sure to close the cursor
	defer db.closeCursor(cursor)

	// load all the data from the database
	list := make([]types.GasPricePeriod, 0)
	for cursor.Next(context.Background()) {
		var row types.GasPricePeriod

		if err := cursor.Decode(&row); err != nil {
			db.log.Errorf("could not decode gas price tick; %s", err.Error())
			return nil, err
		}

		list = append(list, row)
	}

	return list, nil
}

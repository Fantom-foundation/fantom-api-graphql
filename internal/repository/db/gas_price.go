// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// colEpochs represents the name of the gas price collection in database.
	colGasPrice = "gas_price"
)

// initGasPriceCollection initializes the gas price period collection with
// indexes and additional parameters needed by the app.
func (db *MongoDbBridge) initGasPriceCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index sender and recipient
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiGasPriceTimeFrom, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiGasPriceTimeTo, Value: 1}}})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for gas price collection; %s", err.Error())
	}

	// log we are done that
	db.log.Debugf("gas price collection initialized")
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

	// make sure gas price collection is initialized
	if db.initGasPrice != nil {
		db.initGasPrice.Do(func() { db.initGasPriceCollection(col); db.initGasPrice = nil })
	}
	return nil
}

// GasPricePeriodCount calculates total number of gas price period records in the database.
func (db *MongoDbBridge) GasPricePeriodCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(colGasPrice))
}

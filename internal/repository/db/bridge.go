// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDbBridge represents Mongo DB abstraction layer.
type MongoDbBridge struct {
	client *mongo.Client
	log    logger.Logger
	dbName string
}

// New creates a new Mongo Db connection bridge.
func New(cfg *config.Config, log logger.Logger) (*MongoDbBridge, error) {
	// log what we do
	log.Debugf("connecting database at %s/%s", cfg.Db.Url, cfg.Db.DbName)

	// open the database connection
	con, err := connectDb(&cfg.Db)
	if err != nil {
		log.Criticalf("can not contact the database; %s", err.Error())
		return nil, err
	}

	// log the event
	log.Notice("database connection established")

	// return the bridge
	return &MongoDbBridge{
		client: con,
		log:    log,
		dbName: cfg.Db.DbName,
	}, nil
}

// connectDb opens Mongo database connection
func connectDb(cfg *config.Database) (*mongo.Client, error) {
	// get empty unrestricted context
	ctx := context.Background()

	// create new Mongo client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Url))
	if err != nil {
		return nil, err
	}

	// validate the connection was indeed established
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Close will terminate or finish all operations and close the connection to Mongo database.
func (db *MongoDbBridge) Close() {
	// do we have a client?
	if db.client != nil {
		// try to disconnect
		err := db.client.Disconnect(context.Background())
		if err != nil {
			db.log.Errorf("error on closing database connection; %s", err.Error())
		}

		// inform
		db.log.Info("database connection is closed")
	}
}

// getAggregateValue extract single aggregate value for a given collection and aggregation pipeline.
func (db *MongoDbBridge) getAggregateValue(col *mongo.Collection, pipeline *bson.A) (uint64, error) {
	// work with context
	ctx := context.Background()

	// use aggregate pipeline to get the result set, should be just one row
	res, err := col.Aggregate(ctx, *pipeline)
	if err != nil {
		db.log.Errorf("can not get aggregate value; %s", err.Error())
		return 0, err
	}

	// don't forget to close the result cursor
	defer func() {
		// close the cursor
		err = res.Close(ctx)
		if err != nil {
			db.log.Errorf("closing aggregation cursor failed; %s", err.Error())
		}
	}()

	// get the value
	if !res.Next(ctx) {
		db.log.Error("aggregate document not found")
		return 0, err
	}

	// prep container; we are interested in just one value
	var row struct {
		Id    string `bson:"_id"`
		Value int64  `bson:"value"`
	}

	// try to decode the response
	err = res.Decode(&row)
	if err != nil {
		db.log.Errorf("can not parse aggregate value; %s", err.Error())
		return 0, err
	}

	// not a valid aggregate value
	if row.Value < 0 {
		db.log.Error("aggregate value not found")
		return 0, fmt.Errorf("item not found")
	}

	return uint64(row.Value), nil
}

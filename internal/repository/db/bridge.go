// Db package implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// offChainDatabaseName specifies the name of the database being used to store off-chain data.
const offChainDatabaseName = "fantom"

// Bridge represents Mongo DB abstraction layer.
type MongoDbBridge struct {
	client *mongo.Client
	log    logger.Logger

	// fnBalance is a function for retrieving specified account balance.
	fnBalance func(*types.Account) (*hexutil.Big, error)
}

// New creates a new Mongo Db connection bridge.
func New(cfg *config.Config, log logger.Logger) (*MongoDbBridge, error) {
	// get empty unrestricted context
	ctx := context.Background()

	// create new Mongo client
	client, err := mongo.Connect(ctx, clientOptions(cfg))
	if err != nil {
		log.Critical(err)
		return nil, err
	}

	// validate the connection was indeed established
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Critical(err)
		return nil, err
	}

	// log the event
	log.Notice("database backend connection established")

	// make a new Bridge
	return &MongoDbBridge{
		client: client,
		log:    log,
	}, nil
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

// clientOptions creates a new Mongo configuration for connecting the database backend.
func clientOptions(cfg *config.Config) *options.ClientOptions {
	return options.Client().ApplyURI(cfg.MongoUrl)
}

// SetBalance sets the account balance retrieval callback.
func (db *MongoDbBridge) SetBalance(fn func(*types.Account) (*hexutil.Big, error)) {
	db.fnBalance = fn
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

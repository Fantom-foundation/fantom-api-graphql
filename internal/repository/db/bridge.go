// Db package implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// offChainDatabaseName specifies the name of the database being used to store off-chain data.
const offChainDatabaseName = "fantom"

// Bridge represents Mongo DB abstraction layer.
type MongoDbBridge struct {
	client *mongo.Client
	log    logger.Logger
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

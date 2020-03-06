// Db package implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Bridge represents Mongo DB abstraction layer.
type Bridge struct {
	db  *mongo.Client
	log logger.Logger
}

// New creates a new Mongo Db connection bridge.
func New(cfg *config.Config, log logger.Logger) (*Bridge, error) {
	// create new Mongo client
	client, err := mongo.Connect(context.Background(), clientOptions(cfg))
	if err != nil {
		log.Critical(err)
		return nil, err
	}

	// validate the connection was indeed established
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Critical(err)
		return nil, err
	}

	// log the event
	log.Noticef("database backend connection established at [%]", cfg.MongoUrl)

	// make a new Bridge
	return &Bridge{
		db:  client,
		log: log,
	}, nil
}

// clientOptions creates a new Mongo configuration for connecting the database backend.
func clientOptions(cfg *config.Config) *options.ClientOptions {
	return options.Client().ApplyURI(cfg.MongoUrl)
}

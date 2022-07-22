// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository/db/registry"
	"fmt"
	"math/big"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDbBridge represents Mongo DB abstraction layer.
type MongoDbBridge struct {
	client *mongo.Client
	log    logger.Logger
	dbName string

	// sync DB related processes
	wg  sync.WaitGroup
	sig []chan bool

	// init state marks
	initAccounts     *sync.Once
	initTransactions *sync.Once
	initContracts    *sync.Once
	initSwaps        *sync.Once
	initDelegations  *sync.Once
	initWithdrawals  *sync.Once
	initRewards      *sync.Once
	initErc20Trx     *sync.Once
	initFMintTrx     *sync.Once
	initEpochs       *sync.Once
	initGasPrice     *sync.Once
	initBurns        *sync.Once
}

// docListCountAggregationTimeout represents a max duration of DB query executed to calculate
// exact document count in filtered collection. If this duration is exceeded, the query fails
// ad we fall back to full collection documents count estimation.
const docListCountAggregationTimeout = 500 * time.Millisecond

// intZero represents an empty big value.
var intZero = new(big.Int)

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
	db := &MongoDbBridge{
		client: con,
		log:    log,
		dbName: cfg.Db.DbName,
	}

	// check the state
	db.updateDatabaseIndexes()
	db.CheckDatabaseInitState()
	return db, nil
}

// connectDb opens Mongo database connection
func connectDb(cfg *config.Database) (*mongo.Client, error) {
	// get empty unrestricted context
	ctx := context.Background()

	// create new Mongo client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Url).SetRegistry(registry.DefaultRegistry()))
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
		// prep context
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		// try to disconnect
		err := db.client.Disconnect(ctx)
		if err != nil {
			db.log.Errorf("error on closing database connection; %s", err.Error())
		}

		// inform
		db.log.Info("database connection is closed")
		cancel()
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

// CheckDatabaseInitState verifies if database collections have been
// already initialized and marks the empty collections so they can be properly
// configured when created.
func (db *MongoDbBridge) CheckDatabaseInitState() {
	// log what we do
	db.log.Debugf("checking database init state")

	db.collectionNeedInit("accounts", db.AccountCount, &db.initAccounts)
	db.collectionNeedInit("transactions", db.TransactionsCount, &db.initTransactions)
	db.collectionNeedInit("contracts", db.ContractCount, &db.initContracts)
	db.collectionNeedInit("swaps", db.SwapCount, &db.initSwaps)
	db.collectionNeedInit("delegations", db.DelegationsCount, &db.initDelegations)
	db.collectionNeedInit("withdrawals", db.WithdrawalsCount, &db.initWithdrawals)
	db.collectionNeedInit("rewards", db.RewardsCount, &db.initRewards)
	db.collectionNeedInit("erc20 transactions", db.ErcTransactionCount, &db.initErc20Trx)
	db.collectionNeedInit("fmint transactions", db.FMintTransactionCount, &db.initFMintTrx)
	db.collectionNeedInit("epochs", db.EpochsCount, &db.initEpochs)
	db.collectionNeedInit("gas price periods", db.GasPricePeriodCount, &db.initGasPrice)
	db.collectionNeedInit("burned fees", db.BurnCount, &db.initBurns)
}

// checkAccountCollectionState checks the Accounts' collection state.
func (db *MongoDbBridge) collectionNeedInit(name string, counter func() (uint64, error), init **sync.Once) {
	// use the counter to get the collection size
	count, err := counter()
	if err != nil {
		db.log.Errorf("can not check %s count; %s", name, err.Error())
		return
	}

	// collection not empty,
	if 0 != count {
		db.log.Debugf("found %d %s", count, name)
		return
	}

	// collection init needed, create the init control
	db.log.Noticef("%s collection empty", name)
	var once sync.Once
	*init = &once
}

// CountFiltered calculates total number of documents in the given collection for the given filter.
func (db *MongoDbBridge) CountFiltered(col *mongo.Collection, filter *bson.D) (uint64, error) {
	// make sure some filter is used
	if nil == filter {
		filter = &bson.D{}
	}

	// do the counting
	val, err := col.CountDocuments(context.Background(), *filter, options.Count().SetMaxTime(5*time.Second))
	if err != nil {
		db.log.Errorf("can not count documents in rewards collection; %s", err.Error())
		return db.EstimateCount(col)
	}
	return uint64(val), nil
}

// EstimateCount calculates an estimated number of documents in the given collection.
func (db *MongoDbBridge) EstimateCount(col *mongo.Collection) (uint64, error) {
	// do the counting
	val, err := col.EstimatedDocumentCount(context.Background())
	if err != nil {
		db.log.Errorf("can not count documents in rewards collection; %s", err.Error())
		return 0, err
	}
	return uint64(val), nil
}

// listDocumentsCount tries to calculate precise documents count and if it's not counted in limited
// time, use general estimation to speed up the loader.
func (db *MongoDbBridge) listDocumentsCount(col *mongo.Collection, filter *bson.D) (int64, error) {
	// try to count the proper way
	total, err := col.CountDocuments(context.Background(), filter, options.Count().SetMaxTime(docListCountAggregationTimeout))
	if err == nil {
		return total, nil
	}

	// it failed in the limited time we gave it
	db.log.Errorf("can not count documents properly; %s", err.Error())

	// just estimate the whole collection size
	total, err = col.EstimatedDocumentCount(context.Background())
	if err != nil {
		db.log.Errorf("can not count documents")
		return 0, err
	}
	return total, nil
}

// closeCursor closes the given query cursor and reports possible issue if it fails.
func (db *MongoDbBridge) closeCursor(c *mongo.Cursor) {
	if err := c.Close(context.Background()); err != nil {
		db.log.Errorf("failed to close query cursor; %s", err.Error())
	}
}

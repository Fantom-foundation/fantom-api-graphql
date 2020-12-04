// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"crypto/sha256"
	"fantom-api-graphql/internal/types"
	"fmt"
	"math/big"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (

	// coUniswap is the name of the off-chain database collection storing transaction details.
	coUniswap = "uniswap"

	// fiSwapPk is the name of the primary key field of the swap collection.
	fiSwapPk         = "_id"
	fiSwapBlock      = "blk"
	fiSwapTxHash     = "tx"
	fiSwapPair       = "pair"
	fiSwapTimestamp  = "timestamp"
	fiSwapSender     = "sender"
	fiSwapTo         = "to"
	fiSwapAmount0in  = "am0in"
	fiSwapAmount0out = "am0out"
	fiSwapAmount1in  = "am1in"
	fiSwapAmount1out = "am1out"
)

// getHash generates hash for swap from transaction hash and pair address
func getHash(swap *types.Swap) *types.Hash {
	hashBytes := swap.Hash.Big().Bytes()
	pairBytes := swap.Pair.Bytes()
	sum := sha256.Sum256(append(hashBytes, pairBytes...))
	swapHash := types.BytesToHash(sum[:])
	return &swapHash
}

// initUniswapCollection initializes the swap collection with
// indexes and additional parameters needed by the app.
func (db *MongoDbBridge) initUniswapCollection(col *mongo.Collection) {
	if !db.initSwaps {
		return
	}

	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index for primary key
	ix = append(ix, mongo.IndexModel{
		Keys: bson.D{{fiSwapPk, 1}},
	})

	// index sender and recipient
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{fiSwapTimestamp, 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{fiSwapSender, 1}}})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for swap collection; %s", err.Error())
	}

	// log we done that
	db.initSwaps = false
	db.log.Debugf("swap collection initialized")
}

// shouldAddSwap validates if the swap should be added to the persistent storage.
func (db *MongoDbBridge) shouldAddSwap(col *mongo.Collection, swap *types.Swap) bool {
	// check if swap already exists
	swapHash := getHash(swap)
	exists, err := db.IsSwapKnown(col, swapHash)
	if err != nil {
		db.log.Critical(err)
		return false
	}

	// if the transaction already exists, we don't need to do anything here
	return !exists
}

// UniswapAdd stores a swap reference in connected persistent storage.
func (db *MongoDbBridge) UniswapAdd(swap *types.Swap) error {
	// do we have all needed data?
	if swap == nil {
		return fmt.Errorf("can not add empty swap")
	}

	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(coUniswap)

	// if the swap already exists, we don't need to add it
	// just make sure the transaction accounts were processed
	if !db.shouldAddSwap(col, swap) {
		return nil
	}

	// calculate swap hash to use it as a pk
	swapHash := getHash(swap)

	// try to do the insert
	if _, err := col.InsertOne(context.Background(),
		/*
			&bson.D{
				{fiSwapPk, swapHash.String()},
				{fiSwapBlock, uint64(*swap.BlockNumber)},
				{fiSwapTimestamp, uint64(*swap.TimeStamp)},
			}); err != nil {
		*/

		swapData(&bson.D{
			{fiSwapPk, swapHash.String()},
			{fiSwapBlock, uint64(*swap.BlockNumber)},
			{fiSwapTimestamp, uint64(*swap.TimeStamp)},
		}, swap)); err != nil {

		db.log.Critical(err)
		return err
	}

	// add transaction to the db
	db.log.Debugf("swap %s added to database", swapHash.String())

	// check init state
	db.initUniswapCollection(col)
	return nil
}

// smallerNumber for a big.Int of the 1e18 wei
func smallerNumber(nr1 *big.Int, nr2 *big.Int) uint64 {
	return nr1.Div(nr1, nr2).Uint64()
}

// swapData collects the data for the given swap.
func swapData(base *bson.D, swap *types.Swap) bson.D {
	// make a new instance if needed
	if base == nil {
		base = &bson.D{}
	}

	// making amount numbrs smaller to be able to call agregate functions in database
	var divNr, _ = new(big.Int).SetString("1000000", 0)

	// add the extended data
	*base = append(*base,
		bson.E{Key: fiSwapTxHash, Value: swap.Hash.String()},
		bson.E{Key: fiSwapPair, Value: swap.Pair.String()},
		bson.E{Key: fiSwapSender, Value: swap.Sender.String()},
		bson.E{Key: fiSwapTo, Value: swap.To.String()},
		bson.E{Key: fiSwapAmount0in, Value: smallerNumber(swap.Amount0In, divNr)},
		bson.E{Key: fiSwapAmount0out, Value: smallerNumber(swap.Amount0Out, divNr)},
		bson.E{Key: fiSwapAmount1in, Value: smallerNumber(swap.Amount1In, divNr)},
		bson.E{Key: fiSwapAmount1out, Value: smallerNumber(swap.Amount1Out, divNr)},
	)
	return *base
}

// IsSwapKnown checks if swap document already exists in the database.
func (db *MongoDbBridge) IsSwapKnown(col *mongo.Collection, hash *types.Hash) (bool, error) {
	// try to find swap in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{
		{fiSwapPk, hash.String()},
	}, options.FindOne().SetProjection(bson.D{
		{fiSwapPk, true},
	}))

	// error on lookup?
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			// add swap to the db
			db.log.Debugf("swap %s not found in database", hash.String())
			return false, nil
		}

		// log the error of the lookup
		db.log.Error("can not get existing swap pk")
		return false, sr.Err()
	}

	// add swap to the db
	return true, nil
}

// SwapCount returns the number of swaps stored in the database.
func (db *MongoDbBridge) SwapCount() (uint64, error) {
	// get the collection and context
	col := db.client.Database(db.dbName).Collection(coUniswap)

	// find how many swaps do we have in the database
	total, err := col.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		db.log.Errorf("can not count swaps")
		return 0, err
	}

	// inform what we are about to do
	db.log.Debugf("found %d swaps in off-chain database", total)
	return uint64(total), nil
}

// LastKnownSwapBlock returns number of the last known block stored in the database.
func (db *MongoDbBridge) LastKnownSwapBlock() (uint64, error) {
	// prep search options
	opt := options.FindOne()
	opt.SetSort(bson.D{{fiSwapBlock, -1}})
	opt.SetProjection(bson.D{{fiSwapBlock, true}})

	// get the swaps collection
	col := db.client.Database(db.dbName).Collection(coUniswap)
	res := col.FindOne(context.Background(), bson.D{}, opt)
	if res.Err() != nil {
		// may be no block at all
		if res.Err() == mongo.ErrNoDocuments {
			db.log.Info("no blocks found in database")
			return 0, nil
		}

		// log issue
		db.log.Error("can not get the top block")
		return 0, res.Err()
	}

	// get the actual value
	var swap struct {
		Block uint64 `bson:"blk"`
	}

	// get the data
	err := res.Decode(&swap)
	if err != nil {
		db.log.Error("can not decode the top block")
		return 0, res.Err()
	}

	return swap.Block, nil
}

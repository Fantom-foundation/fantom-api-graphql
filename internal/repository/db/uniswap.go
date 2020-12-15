// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"crypto/sha256"
	"fantom-api-graphql/internal/types"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	fiSwapDate       = "date"
	fiSwapSender     = "sender"
	fiSwapTo         = "to"
	fiSwapAmount0in  = "am0in"
	fiSwapAmount0out = "am0out"
	fiSwapAmount1in  = "am1in"
	fiSwapAmount1out = "am1out"
)

// decChange holds information, how many decimals will be added/removed
var decChange = new(big.Int).SetUint64(1000000000)

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
		Keys: bson.D{{Key: fiSwapPk, Value: 1}},
	})

	// index sender and recipient
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: fiSwapDate, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: fiSwapSender, Value: 1}}})

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
		//primitive.Timestamp{T:uint32(time.Now().Unix())
		swapData(&bson.D{
			{Key: fiSwapPk, Value: swapHash.String()},
			{Key: fiSwapBlock, Value: uint64(*swap.BlockNumber)},
			{Key: fiSwapDate, Value: primitive.NewDateTimeFromTime(time.Unix((int64)(*swap.TimeStamp), 0))},
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

// removeDecimals for a big.Int of the 1e18 wei
func removeDecimals(nr1 *big.Int) uint64 {
	// making amount numbrs smaller to be able to call agregate functions in database
	return nr1.Div(nr1, decChange).Uint64()
}

// return Decimals for a big.Int of the 1e18 wei
func returnDecimals(nr1 *big.Int) *big.Int {
	// making amount numbrs bigger again
	return nr1.Mul(nr1, decChange)
}

// swapData collects the data for the given swap.
func swapData(base *bson.D, swap *types.Swap) bson.D {
	// make a new instance if needed
	if base == nil {
		base = &bson.D{}
	}

	// add the extended data
	*base = append(*base,
		bson.E{Key: fiSwapTxHash, Value: swap.Hash.String()},
		bson.E{Key: fiSwapPair, Value: swap.Pair.String()},
		bson.E{Key: fiSwapSender, Value: swap.Sender.String()},
		bson.E{Key: fiSwapTo, Value: swap.To.String()},
		bson.E{Key: fiSwapAmount0in, Value: removeDecimals(swap.Amount0In)},
		bson.E{Key: fiSwapAmount0out, Value: removeDecimals(swap.Amount0Out)},
		bson.E{Key: fiSwapAmount1in, Value: removeDecimals(swap.Amount1In)},
		bson.E{Key: fiSwapAmount1out, Value: removeDecimals(swap.Amount1Out)},
	)
	return *base
}

// IsSwapKnown checks if swap document already exists in the database.
func (db *MongoDbBridge) IsSwapKnown(col *mongo.Collection, hash *types.Hash) (bool, error) {
	// try to find swap in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{
		{Key: fiSwapPk, Value: hash.String()},
	}, options.FindOne().SetProjection(bson.D{
		{Key: fiSwapPk, Value: true},
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
	opt.SetSort(bson.D{{Key: fiSwapBlock, Value: -1}})
	opt.SetProjection(bson.D{{Key: fiSwapBlock, Value: true}})

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

// Volume represents one single sum of volumes for specified pair
type Volume struct {
	ID    string `bson:"_id"`
	Total int64  `bson:"total"`
}

// UniswapVolume resolves volume of swap trades for specified pair and date interval.
// If toTime is 0, then it calculates volumes till now
func (db *MongoDbBridge) UniswapVolume(pairAddress *common.Address, fromTime int64, toTime int64) (types.DefiSwapVolume, error) {

	// translate unix time into mongo primitive date
	fTime := primitive.NewDateTimeFromTime(time.Unix(fromTime, 0))

	var dt bson.D

	// construct date condition
	if toTime != 0 {
		tTime := primitive.NewDateTimeFromTime(time.Unix(toTime, 0))
		dt = bson.D{{Key: "$gte", Value: fTime}, {Key: "$lte", Value: tTime}}
	} else {
		dt = bson.D{{Key: "$gte", Value: fTime}}
	}

	// create command pipeline
	pipe := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "date", Value: dt},
			{Key: "pair", Value: pairAddress.String()}}}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$pair"},
			{Key: "total", Value: bson.M{"$sum": bson.D{
				{Key: "$add", Value: bson.A{"$am0in", "$am0out"}}}}},
		}}},
	}

	// query collection
	col := db.client.Database(db.dbName).Collection(coUniswap)
	cursor, err := col.Aggregate(context.Background(), pipe)
	def := types.DefiSwapVolume{
		PairAddress: pairAddress,
		Volume:      big.NewInt(0)}

	if err != nil {
		db.log.Errorf("Can not get swap volumes: %s", err.Error())
		return def, err
	} else {

		defer cursor.Close(context.Background())
		// get result and fill return data
		for cursor.Next(context.Background()) {
			var val Volume
			err := cursor.Decode(&val)
			if err != nil {
				fmt.Println(err.Error())
			}

			v := returnDecimals(big.NewInt(val.Total))
			def.Volume = v
		}
	}

	return def, nil
}

// UniswapTimeVolumes resolves volumes of swap trades for specified pair grouped by date interval.
// If toTime is 0, then it calculates volumes till now
func (db *MongoDbBridge) UniswapTimeVolumes(pairAddress *common.Address, fromTime int64, toTime int64) ([]types.DefiSwapVolume, error) {

	fTime := primitive.NewDateTimeFromTime(time.Unix(fromTime, 0))

	var dt bson.D

	if toTime != 0 {
		tTime := primitive.NewDateTimeFromTime(time.Unix(toTime, 0))
		dt = bson.D{{Key: "$gte", Value: fTime}, {Key: "$lte", Value: tTime}}
	} else {
		dt = bson.D{{Key: "$gte", Value: fTime}}
	}

	// create query pipeline
	pipe := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "date", Value: dt},
			{Key: "pair", Value: pairAddress.String()}}}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{
				{Key: "year", Value: bson.D{
					{Key: "$year", Value: "$date"}}},
				{Key: "month", Value: bson.D{
					{Key: "$month", Value: "$date"}}},
				{Key: "day", Value: bson.D{
					{Key: "$dayOfMonth", Value: "$date"}}}}},
			{Key: "total", Value: bson.M{"$sum": bson.D{
				{Key: "$add", Value: bson.A{"$am0in", "$am0out"}}}}},
		}}},
		{{Key: "$sort", Value: bson.D{
			{Key: "_id.year", Value: -1},
			{Key: "_id.month", Value: -1},
			{Key: "_id.day", Value: -1},
		}}},
	}

	list := make([]types.DefiSwapVolume, 0)

	// execute query
	col := db.client.Database(db.dbName).Collection(coUniswap)
	cursor, err := col.Aggregate(context.Background(), pipe)

	if err != nil {
		fmt.Println(err.Error())
		return list, nil
	} else {

		defer cursor.Close(context.Background())

		// structure for maping data from db
		type myDate struct {
			Year  int `bson:"year"`
			Month int `bson:"month"`
			Day   int `bson:"day"`
		}
		// represents one single sum of volumes for specified pair
		type volumeData struct {
			ID    myDate `bson:"_id"`
			Total int64  `bson:"total"`
		}

		// iterate thru results and construct data
		for cursor.Next(context.Background()) {
			var val volumeData
			err := cursor.Decode(&val)
			if err != nil {
				fmt.Println(err.Error())
			}
			// adding first 0 for numbers 0..9 to have it in the string
			day := strconv.Itoa(val.ID.Day)
			if val.ID.Day < 10 {
				day = "0" + day
			}
			mnt := strconv.Itoa(val.ID.Month)
			if val.ID.Month < 10 {
				mnt = "0" + mnt
			}
			def := types.DefiSwapVolume{
				PairAddress: pairAddress,
				Volume:      returnDecimals(big.NewInt(val.Total)),
				DateString:  strconv.Itoa(val.ID.Year) + "-" + mnt + "-" + day,
			}
			list = append(list, def)
		}
	}

	return list, nil
}

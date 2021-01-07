// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"crypto/sha256"
	"fantom-api-graphql/internal/types"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (

	// coUniswap is the name of the off-chain database collection storing Uniswap swap details.
	coUniswap = "uniswap"

	// fiSwapPk is the name of the primary key field of the swap collection.
	fiSwapPk         = "_id"
	fiSwapType       = "type"
	fiSwapBlock      = "blk"
	fiSwapTxHash     = "tx"
	fiSwapPair       = "pair"
	fiSwapDate       = "date"
	fiSwapSender     = "sender"
	fiSwapAmount0in  = "am0in"
	fiSwapAmount0out = "am0out"
	fiSwapAmount1in  = "am1in"
	fiSwapAmount1out = "am1out"
	fiSwapReserve0   = "reserve0"
	fiSwapReserve1   = "reserve1"
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

	// index date and sender
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
	exists, err := db.IsSwapKnown(col, swapHash, swap)
	if err != nil {
		db.log.Critical(err)
		return false
	}

	// if the transaction already exists, we don't need to do anything here
	return !exists
}

// isZeroSwap checks if amounts are not zero to avoid divide by 0 during calculations in db
func isZeroSwap(swap *types.Swap) bool {
	if swap.Type == types.SwapSync {
		return false
	}
	am := big.Int{}
	am0small := removeDecimals(am.Add(swap.Amount0In, swap.Amount0Out))
	am1small := removeDecimals(am.Add(swap.Amount1In, swap.Amount1Out))

	if am0small == 0 || am1small == 0 {
		return true
	}
	return false
}

// UniswapAdd stores a swap reference in connected persistent storage.
func (db *MongoDbBridge) UniswapAdd(swap *types.Swap) error {
	// do we have all needed data?
	if swap == nil {
		return fmt.Errorf("can not add empty swap")
	}

	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(coUniswap)

	// check for zero amounts in the swap, because of future div by 0 during agregation in db
	if isZeroSwap(swap) {
		db.log.Debugf("Swap from block %v will not be added, because swap amount is 0 ater removing decimals", uint64(*swap.BlockNumber))
		return nil
	}

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
			{Key: fiSwapDate, Value: primitive.NewDateTimeFromTime(time.Unix((int64)(*swap.TimeStamp), 0).UTC())},
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
		bson.E{Key: fiSwapType, Value: swap.Type},
		bson.E{Key: fiSwapTxHash, Value: swap.Hash.String()},
		bson.E{Key: fiSwapPair, Value: swap.Pair.String()},
		bson.E{Key: fiSwapSender, Value: swap.Sender.String()},
		bson.E{Key: fiSwapAmount0in, Value: removeDecimals(swap.Amount0In)},
		bson.E{Key: fiSwapAmount0out, Value: removeDecimals(swap.Amount0Out)},
		bson.E{Key: fiSwapAmount1in, Value: removeDecimals(swap.Amount1In)},
		bson.E{Key: fiSwapAmount1out, Value: removeDecimals(swap.Amount1Out)},
		bson.E{Key: fiSwapReserve0, Value: removeDecimals(swap.Reserve0)},
		bson.E{Key: fiSwapReserve1, Value: removeDecimals(swap.Reserve1)},
	)
	return *base
}

// IsSwapKnown checks if swap document already exists in the database.
func (db *MongoDbBridge) IsSwapKnown(col *mongo.Collection, hash *types.Hash, swap *types.Swap) (bool, error) {
	// try to find swap in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{
		{Key: fiSwapPk, Value: hash.String()}})

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

	// swap is known, jus log and return true
	db.log.Debugf("Swap %s is already in database.", hash.String())

	// if swap is sync type, then update reserves
	if swap.Type == types.SwapSync {
		db.log.Debugf("Updating reserves for Swap %s", hash.String())
		_, err := col.UpdateOne(context.Background(),
			bson.M{fiSwapPk: hash.String()},
			bson.D{
				{Key: "$set", Value: bson.M{fiSwapReserve0: removeDecimals(swap.Reserve0)}},
				{Key: "$set", Value: bson.M{fiSwapReserve1: removeDecimals(swap.Reserve1)}}})
		if err != nil {
			db.log.Errorf("Unable to update reserves for Swap %s", hash.String())
		}
	} else {

		// in case the sync event was recorded first, update reserves into actual swap
		// and delete sync record.
		type Values struct {
			Type     int   `bson:"type"`
			Reserve0 int64 `bson:"reserve0"`
			Reserve1 int64 `bson:"reserve1"`
		}
		var values Values
		sr.Decode(&values)

		if types.SwapSync == values.Type {
			db.log.Debugf("Updating reserve for swap: %s, reserve0: %v, reserve1: %v", hash.String(), values.Reserve0, values.Reserve1)
			col.DeleteOne(context.Background(), bson.D{{Key: fiSwapPk, Value: hash.String()}})
			swap.Reserve0 = returnDecimals(big.NewInt(values.Reserve0))
			swap.Reserve1 = returnDecimals(big.NewInt(values.Reserve1))
			return false, nil
		}
	}
	// no need to change data
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

	// -1 is for confuiguration document with last correct swap block number
	if total > 1 {
		total--
	}

	// inform what we are about to do
	db.log.Debugf("found %d swaps in off-chain database", total)
	return uint64(total), nil
}

// LastKnownSwapBlock returns number of the last known block stored in the database.
func (db *MongoDbBridge) LastKnownSwapBlock() (uint64, error) {

	// search for document with last swap block number
	query := bson.D{
		{Key: "lastSwapSyncBlk", Value: bson.D{
			{Key: "$exists", Value: "true"}}},
	}

	// get the swaps collection
	col := db.client.Database(db.dbName).Collection(coUniswap)
	res := col.FindOne(context.Background(), query)
	if res.Err() != nil {
		// may be no block at all
		if res.Err() == mongo.ErrNoDocuments {
			db.log.Info("No document with last swap block number in database starting from 0.")
			return 0, nil
		}

		// log issue
		db.log.Error("Can not get the last correct swap block number, starting from 0.")
		return 0, res.Err()
	}

	// get the actual value
	var swap struct {
		Block uint64 `bson:"lastSwapSyncBlk"`
	}

	// get the data
	err := res.Decode(&swap)
	if err != nil {
		db.log.Error("Can not resolve id of the last correct swap block in db. Starting from 0.")
		return 0, res.Err()
	}

	return swap.Block, nil
}

// UniswapUpdateLastKnownSwapBlock stores a last correctly saved swap block number into persistent storage.
func (db *MongoDbBridge) UniswapUpdateLastKnownSwapBlock(blkNumber uint64) error {

	// is valid block number
	if blkNumber == 0 {
		return fmt.Errorf("No need to store zero value, will start from 0 next time")
	}

	// document for update with last swap block number
	query := bson.D{
		{Key: "lastSwapSyncBlk", Value: bson.D{
			{Key: "$exists", Value: "true"}}},
	}

	data := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "lastSwapSyncBlk", Value: blkNumber}}},
	}

	// get the collection for transactions and insert data
	col := db.client.Database(db.dbName).Collection(coUniswap)
	if _, err := col.UpdateOne(context.Background(),
		query, data, options.Update().SetUpsert(true)); err != nil {

		db.log.Critical(err)
		return err
	}

	// log
	db.log.Debugf("Block %d was set as a last correct uniswap block into database", blkNumber)
	return nil
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
func (db *MongoDbBridge) UniswapTimeVolumes(pairAddress *common.Address, resolution string, fromTime int64, toTime int64) ([]types.DefiSwapVolume, error) {

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
			{Key: "_id", Value: getGroupBsonD(resolution)},
			{Key: "total", Value: bson.M{"$sum": bson.D{
				{Key: "$add", Value: bson.A{"$am0in", "$am0out"}}}}},
		}}},
		{{Key: "$sort", Value: bson.D{
			{Key: "_id", Value: 1},
		}}},
	}

	list := make([]types.DefiSwapVolume, 0)

	// execute query
	col := db.client.Database(db.dbName).Collection(coUniswap)
	cursor, err := col.Aggregate(context.Background(), pipe)

	if err != nil {
		db.log.Errorf(err.Error())
		return list, nil
	} else {
		defer cursor.Close(context.Background())

		// iterate thru results and construct data
		for cursor.Next(context.Background()) {
			var val Volume
			err := cursor.Decode(&val)
			if err != nil {
				db.log.Errorf(err.Error())
			}
			def := types.DefiSwapVolume{
				PairAddress: pairAddress,
				Volume:      returnDecimals(big.NewInt(val.Total)),
				DateString:  val.ID,
			}
			list = append(list, def)
		}
	}

	return list, nil
}

// getGroupBsonD is a helper function for constructing group db request
func getGroupBsonD(resolution string) bson.D {

	// initial value set to 1 minute
	mul := 60 * 1000

	switch resolution {
	case "month":
		mul *= 30 * 24 * 60
	case "day":
		mul *= 24 * 60
	case "4h":
		mul *= 4 * 60
	case "1h":
		mul *= 60
	case "30m":
		mul *= 30
	case "15m":
		mul *= 15
	case "5m":
		mul *= 5
	case "1m":
		mul *= 1
	default:
		mul *= 24 * 60
	}

	format := "%Y-%m-%dT%H:%M:%S.000Z"
	groupBson := bson.D{
		{Key: "$dateToString", Value: bson.D{
			{Key: "format", Value: format},
			{Key: "date", Value: bson.D{
				{Key: "$add", Value: bson.A{bson.D{
					{Key: "$subtract", Value: bson.A{
						bson.D{{Key: "$subtract", Value: bson.A{"$date", 0}}},
						bson.D{{Key: "$mod", Value: bson.A{bson.D{
							{Key: "$toLong", Value: bson.D{
								{Key: "$subtract", Value: bson.A{"$date", 0}}}}},
							mul}},
						},
					}}},
					0}},
			},
			}}},
	}
	return groupBson
}

// getDateBsonD is a helper function for constructing date db request
func getDateBsonD(fromTime int64, toTime int64) bson.D {
	fTime := primitive.NewDateTimeFromTime(time.Unix(fromTime, 0))

	var dt bson.D

	if toTime != 0 {
		tTime := primitive.NewDateTimeFromTime(time.Unix(toTime, 0))
		dt = bson.D{{Key: "$gte", Value: fTime}, {Key: "$lte", Value: tTime}}
	} else {
		dt = bson.D{{Key: "$gte", Value: fTime}}
	}

	return dt
}

// UniswapTimePrices resolves price of swap trades for specified pair grouped by date interval.
// If toTime is 0, then it calculates prices till now
func (db *MongoDbBridge) UniswapTimePrices(pairAddress *common.Address, resolution string, fromTime int64, toTime int64, direction int32) ([]types.DefiTimePrice, error) {

	tokenASum := bson.D{{Key: "$add", Value: bson.A{"$am0in", "$am0out"}}}
	tokenBSum := bson.D{{Key: "$add", Value: bson.A{"$am1in", "$am1out"}}}

	// creating priceBsonD bson request object
	var priceBsonD bson.D
	if direction == 0 {
		priceBsonD = bson.D{{Key: "$divide", Value: bson.A{tokenASum, tokenBSum}}}
	} else {
		priceBsonD = bson.D{{Key: "$divide", Value: bson.A{tokenBSum, tokenASum}}}
	}

	// create query pipeline
	pipe := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "date", Value: getDateBsonD(fromTime, toTime)},
			{Key: "pair", Value: pairAddress.String()}}}},
		{{Key: "$sort", Value: bson.D{
			{Key: "date", Value: 1},
		}}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: getGroupBsonD(resolution)},
			{Key: "low", Value: bson.D{
				{Key: "$min", Value: priceBsonD}}},
			{Key: "high", Value: bson.D{
				{Key: "$max", Value: priceBsonD}}},
			{Key: "open", Value: bson.D{
				{Key: "$first", Value: priceBsonD}}},
			{Key: "close", Value: bson.D{
				{Key: "$last", Value: priceBsonD}}},
			{Key: "avg", Value: bson.D{
				{Key: "$avg", Value: priceBsonD}}},
		}}},
		{{Key: "$sort", Value: bson.D{
			{Key: "_id", Value: 1},
		}}},
	}

	list := make([]types.DefiTimePrice, 0)

	// execute query
	col := db.client.Database(db.dbName).Collection(coUniswap)
	cursor, err := col.Aggregate(context.Background(), pipe)
	if err != nil {
		db.log.Errorf(err.Error())
		return list, nil
	} else {
		defer cursor.Close(context.Background())

		// iterate thru results and construct data
		for cursor.Next(context.Background()) {
			var priceVal types.DefiTimePrice
			err := cursor.Decode(&priceVal)
			if err != nil {
				db.log.Errorf(err.Error())
			}
			priceVal.PairAddress = *pairAddress
			list = append(list, priceVal)
		}
	}

	return list, nil
}

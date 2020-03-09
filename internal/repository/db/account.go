// Db package implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
)

const (
	// coAccountTransactions is the name of the off-chain database collection storing reference between
	// accounts and transactions.
	coAccountTransactions = "acc_trx"

	// fiAccountTransactionPk is the name of the primary key field of the account to transaction collection.
	fiAccountTransactionPk = "_id"

	// fiAccountHash is the name of the account address field in the collection.
	fiAccountAddress = "acc"

	// fiTransactionHash is the name of the transaction hash field in the collection.
	fiTransactionHash = "trx"

	// hexBlockMask is a bit mask we use to validate that a block number is within available PK range (49 bits).
	blockNumberMask uint64 = 0x1FFFFFFFFFFFF

	// blockNumberShift represents bits to shift the block number left to make room for trx index and direction.
	blockNumberShift uint = 15

	// hexBlockMask is a bit mask we use to validate that a transaction index in block is within PK range (14 bits).
	transactionIndexMask uint64 = 0x3FFF

	// transactionIndexShift represents bits to shift the trx index left to make room for trx direction.
	transactionIndexShift uint = 1
)

// refAccountTransaction represents a single account hash to transaction hash reference.
type refAccountTransaction struct {
	// _id is unique trx identifier constructed from trx direction, a block number, and trx index inside the block.
	_id     uint64
	account types.Hash
	trx     types.Hash
}

// AddTransaction stores a transaction reference in connected persistent storage.
//
// Please note that contract creation transactions will not have a recipient and so they will appear only once.
// Regular transactions will be stored twice, for a sending address and also for a receiving address.
func (db *MongoDbBridge) AddTransaction(block *big.Int, trxIndex *uint64, trx *types.Transaction) error {
	// get empty and unrestricted context
	ctx := context.Background()

	// get the collection for account transactions
	col := db.client.Database(offChainDatabaseName).Collection(coAccountTransactions)

	// make the base PK if possible
	pk, err := calcAccountTransactionPk(block, trxIndex)
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// store outgoing transaction ref
	err = db.setAccountTransaction(&ctx, col, &pk, &trx.From, &trx.Hash)
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// do we have any receiving part?
	if trx.To != nil {
		// set the key to receiving
		pk = pk | 1

		// store incoming transaction ref
		err = db.setAccountTransaction(&ctx, col, &pk, trx.To, &trx.Hash)
		if err != nil {
			db.log.Critical(err)
			return err
		}
	}

	return nil
}

// calcAccountTransactionPk calculates base transaction primary key index from the block number and trx index in the block
// such as:
//     The lowest bit is reserved for direction; 0 = outgoing and 1 = incoming trx for the given account.
//     Next 14 bits hold trx index inside the block; this gives us 16383 potential transactions in a block.
//     Next 49 bits hold the block number; this gives us 562.949.953.421.311 potential blocks.
//
//		It means that, on rate of 100kBlocks/s, we will run out of PK indexes for blocks in 178 years.
//		Thankfuly, the Y2198 breakdown can be easily fixed with uint256 arithmetics we plan
//		to support as soon as possible.
func calcAccountTransactionPk(block *big.Int, trxIndex *uint64) (uint64, error) {
	var pk uint64

	// is it possible to convert the block number?
	if !block.IsUint64() {
		return pk, fmt.Errorf("block number is way too high to be stored")
	}

	// convert and validate that we are safely within the 49 bits reserved for the block number
	bx := block.Uint64()
	if bx != (bx & blockNumberMask) {
		return pk, fmt.Errorf("block number bit count is higer than available primary key space")
	}

	// validate transaction index
	if *trxIndex != (*trxIndex & transactionIndexMask) {
		return pk, fmt.Errorf("transaction index within the block is too high to be stored")
	}

	// make the PK index
	return pk | (bx << blockNumberShift) | (*trxIndex << transactionIndexShift), nil
}

// setAccountTransaction stores the given transaction hash in relationship with the address under the given PK.
func (db *MongoDbBridge) setAccountTransaction(ctx *context.Context, col *mongo.Collection, pk *uint64, addr *common.Address, trx *types.Hash) error {
	// prep options
	opt := options.Update().SetUpsert(true)

	// do the update based on given PK; we don't need to pull the document updated
	_, err := col.UpdateOne(*ctx, bson.D{{fiAccountTransactionPk, pk}}, bson.D{
		{fiAccountTransactionPk, *pk},
		{fiAccountAddress, addr.String()},
		{fiTransactionHash, *trx},
	}, opt)

	return err
}

// AccountTransactions returns capped list of transaction hashes related to an account from off-chain database.
//
// Transactions are loaded based on following option combinations:
//    - nil anchor, positive count => start from the most recent trx and return at most <count> older trx hashes
//    - nil anchor, negative count => start from the oldest trx and return at most <count> newer trx hashes
//    - NOT nil anchor, positive count => start after the anchored trx and return at most <count> older trx hashes
//    - NOT nil anchor, negative count => start before the anchored trx and return at most <count> newer trx hashes
//
// Transaction hashes are always returned ordered by their relative age in ascending order, i.e. from new to old.
// The relative age of transaction is derived from the relative age of the block they belong to (i.e. the block number)
// and the index of the transaction inside the block.
func (db *MongoDbBridge) AccountTransactions(acc *types.Account, anchor *string, count int) ([]*types.Hash, error) {
	// zero count? technically correct request for an empty set, but we should not allow it
	if count == 0 {
		return nil, fmt.Errorf("zero cap on transaction count is not allowed")
	}

	// get empty & unrestricted context
	ctx := context.Background()

	// get the collection for account transactions
	col := db.client.Database(offChainDatabaseName).Collection(coAccountTransactions)

	// start looking for the data
	cur, err := col.Find(ctx, db.accTrxFindFilter(col, acc.Address, anchor, count), db.accTrxFindOptions(count))
	if err != nil {
		db.log.Error("can not get account to transaction list from the off-chain database")
		return nil, err
	}

	// always close the cursor before we leave; log errors on closing to be clean
	defer func() {
		if err := cur.Close(ctx); err != nil {
			db.log.Error(err)
		}
	}()

	// loop and decode hashes to build the result set
	result := make([]*types.Hash, 0)
	for cur.Next(ctx) {
		var e refAccountTransaction

		err := cur.Decode(&e)
		if err != nil {
			db.log.Critical(err)
			return nil, err
		}

		result = append(result, &e.trx)
	}

	// did we get an error in scan?
	if err := cur.Err(); err != nil {
		db.log.Critical(err)
		return nil, err
	}

	// if we sorted from bottom up, we need to reverse the result set so newer transactions are on top
	if count < 0 {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
	}

	return result, nil
}

// accTrxFindFilter builds filter for account to hash query in relationship to an optional anchored trx hash.
func (db *MongoDbBridge) accTrxFindFilter(col *mongo.Collection, addr *common.Address, anchor *string, count int) interface{} {
	// do we need to check the anchor?
	if nil != anchor {
		// get PK of the anchored transaction for the account address
		pk, err := db.accTrxPkByAnchor(col, addr, anchor)
		if err == nil {
			// operator to be used for find changes for negative count
			var op = "$gt"
			if count < 0 {
				op = "$lt"
			}

			// prep the filter
			return bson.D{
				{fiAccountAddress, addr.String()},
				{fiAccountTransactionPk, bson.D{
					{op, pk},
				}},
			}
		}

		// anchored trx not found? log and fallback to simple search; maybe the anchor was wrong
		db.log.Error("can not find anchor for account to transaction find")
	}

	// simply filter records by address
	return bson.D{
		{fiAccountAddress, addr.String()},
	}
}

// accTrxPkByAnchor finds primary key of a transaction anchor for account address.
func (db *MongoDbBridge) accTrxPkByAnchor(col *mongo.Collection, addr *common.Address, anchor *string) (uint64, error) {
	// get empty & unrestricted context
	ctx := context.Background()

	var pk uint64
	err := col.FindOne(ctx, bson.D{
		{fiAccountAddress, addr.String()},
		{fiTransactionHash, *anchor}}).Decode(&pk)

	// do we have the result?
	if err != nil {
		return 0, err
	}

	return pk, nil
}

// accTrxFindOptions builds query options based on count value of the request.
func (db *MongoDbBridge) accTrxFindOptions(count int) *options.FindOptions {
	// what direction do we sort depends on the value of count
	var sd = 1
	if count < 0 {
		sd = -1
		count = -count
	}

	// prep limit option
	opt := options.Find()
	opt.SetLimit(int64(count))
	opt.SetSort(bson.D{{fiAccountTransactionPk, sd}})

	return opt
}

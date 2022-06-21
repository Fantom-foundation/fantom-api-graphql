// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// trxLargeInputWall represents the largest transaction input block we store in the off-chain database.
	// Larger inputs (like contract deployments) need to be loaded from the blockchain directly if needed.
	trxLargeInputWall = 32 * 8

	// colTransactions is the name of the off-chain database collection storing transaction details.
	colTransactions = "transaction"

	// fiTransactionBlockHash is the block hash field of the transaction.
	fiTransactionBlockHash = "blk_h"

	// fiTransactionBlockNumber is the block number field of the transaction.
	fiTransactionBlockNumber = "blk"

	// fiTransactionTimeStamp is the name of the field of the transaction time stamp.
	fiTransactionTimeStamp = "stamp"

	// fiTransactionFrom is the name of the address field of the sender account.
	// db.transaction.createIndex({from:1}).
	fiTransactionFrom = "from"

	// fiTransactionGas is the transaction gas field of the transaction.
	fiTransactionGas = "gas_lim"

	// fiTransactionGasUsed is the transaction gas used field of the transaction.
	fiTransactionGasUsed = "gas_use"

	// fiTransactionCumulativeGasUsed is the transaction cumulative gas used field of the transaction.
	fiTransactionCumulativeGasUsed = "gas_cum"

	// fiTransactionGasPrice is the transaction gas price field of the transaction.
	fiTransactionGasPrice = "gas_pri"

	// fiTransactionHash is the transaction hash field of the transaction.
	fiTransactionHash = "_id"

	// fiTransactionNonce is the transaction nonce field of the transaction.
	fiTransactionNonce = "nonce"

	// fiTransactionTo is the name of the address field of the recipient account.
	// null for contract creation.
	// db.transaction.createIndex({to:1}).
	fiTransactionTo = "to"

	// fiTransactionContractAddress is the transaction contract address field of the transaction.
	fiTransactionContractAddress = "contr"

	// fiTransactionValue is the name of the field of the transaction value.
	fiTransactionValue = "value"

	// fiTransactionInputData is the transaction input data field of the transaction.
	fiTransactionInputData = "input"

	// fiTransactionLargeInput is the transaction large input field of the transaction.
	fiTransactionLargeInput = "large"

	// fiTransactionIndex is the transaction index field of the transaction.
	fiTransactionIndex = "bix"

	// fiTransactionStatus is the transaction status field of the transaction.
	fiTransactionStatus = "stat"

	// fiTransactionLogs is the transaction logs field of the transaction.
	fiTransactionLogs = "logs"

	// fiTransactionOrdinalIndex is the name of the transaction ordinal index in the blockchain field.
	// db.transaction.createIndex({_id:1,orx:-1},{unique:true})
	fiTransactionOrdinalIndex = "orx"

	// fiTransactionAmount is the transaction amount field of the transaction.
	fiTransactionAmount = "amount"

	// fiTransactionGasGWei is the transaction gas gWei field of the transaction.
	fiTransactionGasGWei = "gwx100"
)

// transactionCollectionIndexes provides a list of indexes expected to exist on the transactions' collection.
func transactionCollectionIndexes() []mongo.IndexModel {
	// prepare index models
	ix := make([]mongo.IndexModel, 6)

	// index ordinal key sorted from high to low since this is the way we usually list
	unique := true
	ixTxOrdinal := "ix_tx_orx"
	ix[0] = mongo.IndexModel{
		Keys: bson.D{{Key: fiTransactionOrdinalIndex, Value: -1}},
		Options: &options.IndexOptions{
			Unique: &unique,
			Name:   &ixTxOrdinal,
		},
	}

	ixTxFrom := "ix_tx_from"
	ix[1] = mongo.IndexModel{
		Keys:    bson.D{{Key: fiTransactionFrom, Value: 1}},
		Options: &options.IndexOptions{Name: &ixTxFrom},
	}

	ixTxTo := "ix_tx_to"
	ix[2] = mongo.IndexModel{
		Keys:    bson.D{{Key: fiTransactionTo, Value: 1}},
		Options: &options.IndexOptions{Name: &ixTxTo},
	}

	ixTxStamp := "ix_tx_stamp"
	ix[3] = mongo.IndexModel{
		Keys:    bson.D{{Key: fiTransactionTimeStamp, Value: 1}},
		Options: &options.IndexOptions{Name: &ixTxStamp},
	}

	// sender + ordinal index
	ixTxSenderOrdinal := "ix_tx_from_orx"
	ix[4] = mongo.IndexModel{
		Keys: bson.D{{Key: fiTransactionFrom, Value: 1}, {Key: fiTransactionOrdinalIndex, Value: -1}},
		Options: &options.IndexOptions{
			Name:   &ixTxSenderOrdinal,
			Unique: &unique,
		},
	}

	// recipient + ordinal index
	ixTxRecipientOrdinal := "ix_tx_to_orx"
	ix[5] = mongo.IndexModel{
		Keys: bson.D{{Key: fiTransactionTo, Value: 1}, {Key: fiTransactionOrdinalIndex, Value: -1}},
		Options: &options.IndexOptions{
			Name:   &ixTxRecipientOrdinal,
			Unique: &unique,
		},
	}

	return ix
}

// shouldAddTransaction validates if the transaction should be added to the persistent storage.
func (db *MongoDbBridge) shouldAddTransaction(col *mongo.Collection, trx *types.Transaction) bool {
	// check if the transaction already exists
	exists, err := db.IsTransactionKnown(col, &trx.Hash)
	if err != nil {
		db.log.Critical(err)
		return false
	}

	// if the transaction already exists, we don't need to do anything here
	return !exists
}

// AddTransaction stores a transaction reference in connected persistent storage.
func (db *MongoDbBridge) AddTransaction(block *types.Block, trx *types.Transaction) error {
	// do we have all needed data?
	if block == nil || trx == nil {
		return fmt.Errorf("can not add empty transaction")
	}

	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(colTransactions)

	// if the transaction already exists, we don't need to add it
	// just make sure the transaction accounts were processed
	if !db.shouldAddTransaction(col, trx) {
		return db.UpdateTransaction(col, trx)
	}

	trx.LargeInput = len(trx.InputData) > trxLargeInputWall
	if trx.LargeInput {
		trx.InputData = hexutil.Bytes{}
	}

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: fiTransactionHash, Value: trx.Hash}},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: fiTransactionBlockHash, Value: trx.BlockHash},
				{Key: fiTransactionBlockNumber, Value: trx.BlockNumber},
				{Key: fiTransactionTimeStamp, Value: trx.TimeStamp},
				{Key: fiTransactionFrom, Value: trx.From},
				{Key: fiTransactionGas, Value: trx.Gas},
				{Key: fiTransactionGasUsed, Value: trx.GasUsed},
				{Key: fiTransactionCumulativeGasUsed, Value: trx.CumulativeGasUsed},
				{Key: fiTransactionGasPrice, Value: trx.GasPrice},
				{Key: fiTransactionNonce, Value: trx.Nonce},
				{Key: fiTransactionTo, Value: trx.To},
				{Key: fiTransactionContractAddress, Value: trx.ContractAddress},
				{Key: fiTransactionValue, Value: trx.Value},
				{Key: fiTransactionInputData, Value: trx.InputData},
				{Key: fiTransactionLargeInput, Value: trx.LargeInput},
				{Key: fiTransactionIndex, Value: trx.Index},
				{Key: fiTransactionStatus, Value: trx.Status},
				{Key: fiTransactionLogs, Value: trx.Logs},
				{Key: fiTransactionOrdinalIndex, Value: trx.ComputedOrdinalIndex()},
				{Key: fiTransactionAmount, Value: trx.ComputedAmount()},
				{Key: fiTransactionGasGWei, Value: trx.ComputedGWei()},
			}},
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fiTransactionHash, Value: trx.Hash},
			}},
		},
		options.Update().SetUpsert(true),
	); err != nil {
		db.log.Critical(err)
		return err
	}

	// add transaction to the db
	db.log.Debugf("transaction %s added to database", trx.Hash.String())

	return nil
}

// UpdateTransaction updates transaction data in the database collection.
func (db *MongoDbBridge) UpdateTransaction(col *mongo.Collection, trx *types.Transaction) error {
	// notify
	db.log.Debugf("updating transaction %s", trx.Hash.String())

	// try to update a delegation by replacing it in the database
	// we use address and validator ID to identify unique delegation
	er, err := col.UpdateOne(context.Background(), bson.D{
		{Key: fiTransactionHash, Value: trx.Hash},
	}, bson.D{{Key: "$set", Value: bson.D{
		{Key: fiTransactionOrdinalIndex, Value: trx.ComputedOrdinalIndex()},
		{Key: fiTransactionFrom, Value: trx.From},
		{Key: fiTransactionValue, Value: trx.Value},
		{Key: fiTransactionTimeStamp, Value: trx.TimeStamp},
	}}}, new(options.UpdateOptions).SetUpsert(false))
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// do we actually have the document
	if 0 == er.MatchedCount {
		return fmt.Errorf("can not update, the transaction not found in database")
	}
	return nil
}

// IsTransactionKnown checks if a transaction document already exists in the database.
func (db *MongoDbBridge) IsTransactionKnown(col *mongo.Collection, hash *common.Hash) (bool, error) {
	// try to find the transaction in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{
		{Key: fiTransactionHash, Value: hash},
	}, options.FindOne().SetProjection(bson.D{
		{Key: fiTransactionHash, Value: true},
	}))

	// error on lookup?
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			// add transaction to the db
			db.log.Debugf("transaction %s not found in database", hash.String())
			return false, nil
		}

		// log the error of the lookup
		db.log.Error("can not get existing transaction pk")
		return false, sr.Err()
	}

	// add transaction to the db
	return true, nil
}

// initTrxList initializes list of transactions based on provided cursor and count.
func (db *MongoDbBridge) initTrxList(col *mongo.Collection, cursor *string, count int32, filter *bson.D) (*types.TransactionList, error) {
	// make sure some filter is used
	if nil == filter {
		filter = &bson.D{}
	}

	// find how many transactions do we have in the database
	total, err := db.listDocumentsCount(col, filter)
	if err != nil {
		db.log.Errorf("can not count transactions")
		return nil, err
	}

	// make the list and notify the size of it
	db.log.Debugf("found %d filtered transactions", total)
	list := types.TransactionList{
		Collection: make([]*types.Transaction, 0),
		Total:      uint64(total),
		First:      0,
		Last:       0,
		IsStart:    total == 0,
		IsEnd:      total == 0,
		Filter:     *filter,
	}

	// is the list non-empty? return the list with properly calculated range marks
	if 0 < total {
		return db.trxListWithRangeMarks(col, &list, cursor, count, filter)
	}

	// this is an empty list
	db.log.Debug("empty transaction list created")
	return &list, nil
}

// trxListWithRangeMarks returns the transaction list with proper First/Last marks of the transaction range.
func (db *MongoDbBridge) trxListWithRangeMarks(
	col *mongo.Collection,
	list *types.TransactionList,
	cursor *string,
	count int32,
	filter *bson.D,
) (*types.TransactionList, error) {
	var err error

	// find out the cursor ordinal index
	if cursor == nil && count > 0 {
		// get the highest available ordinal index (top transaction)
		list.First, err = db.findBorderOrdinalIndex(col,
			*filter,
			options.FindOne().SetSort(bson.D{{Key: fiTransactionOrdinalIndex, Value: -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available ordinal index (top transaction)
		list.First, err = db.findBorderOrdinalIndex(col,
			*filter,
			options.FindOne().SetSort(bson.D{{Key: fiTransactionOrdinalIndex, Value: 1}}))
		list.IsEnd = true

	} else if cursor != nil {
		// get the highest available ordinal index (top transaction)
		list.First, err = db.findBorderOrdinalIndex(col,
			bson.D{{Key: fiTransactionHash, Value: *cursor}},
			options.FindOne())
	}

	// check the error
	if err != nil {
		db.log.Errorf("can not find the initial transactions")
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("transaction list initialized with ordinal index %d", list.First)
	return list, nil
}

// findBorderOrdinalIndex finds the highest, or lowest ordinal index in the collection.
// For negative sort it will return highest and for positive sort it will return lowest available value.
func (db *MongoDbBridge) findBorderOrdinalIndex(col *mongo.Collection, filter bson.D, opt *options.FindOneOptions) (uint64, error) {
	// prep container
	var row struct {
		Value uint64 `bson:"orx"`
	}

	// make sure we pull only what we need
	opt.SetProjection(bson.D{{Key: "orx", Value: true}})
	sr := col.FindOne(context.Background(), filter, opt)

	// try to decode
	err := sr.Decode(&row)
	if err != nil {
		return 0, err
	}

	return row.Value, nil
}

// txListFilter creates a filter for transaction list search.
func (db *MongoDbBridge) txListFilter(cursor *string, count int32, list *types.TransactionList) *bson.D {
	// inform what we are about to do
	db.log.Debugf("transaction filter starts from index %d", list.First)

	// build the filter query
	if cursor == nil {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: fiTransactionOrdinalIndex, Value: bson.D{{Key: "$lte", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: fiTransactionOrdinalIndex, Value: bson.D{{Key: "$gte", Value: list.First}}})
		}
	} else {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: fiTransactionOrdinalIndex, Value: bson.D{{Key: "$lt", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: fiTransactionOrdinalIndex, Value: bson.D{{Key: "$gt", Value: list.First}}})
		}
	}

	// log the filter
	return &list.Filter
}

// txListOptions creates a filter options set for transactions list search.
func (db *MongoDbBridge) txListOptions(count int32) *options.FindOptions {
	// prep options
	opt := options.Find()

	// how to sort results in the collection
	if count > 0 {
		// from high (new) to low (old)
		opt.SetSort(bson.D{{Key: fiTransactionOrdinalIndex, Value: -1}})
	} else {
		// from low (old) to high (new)
		opt.SetSort(bson.D{{Key: fiTransactionOrdinalIndex, Value: 1}})
	}

	// prep the loading limit
	var limit = int64(count)
	if limit < 0 {
		limit = -limit
	}

	// apply the limit, try to get one more transaction
	// so we can detect list end
	opt.SetLimit(limit + 1)
	return opt
}

// txListLoad load the initialized list from database
func (db *MongoDbBridge) txListLoad(col *mongo.Collection, cursor *string, count int32, list *types.TransactionList) error {
	// get the context for loader
	ctx := context.Background()

	// load the data
	ld, err := col.Find(ctx, db.txListFilter(cursor, count, list), db.txListOptions(count))
	if err != nil {
		db.log.Errorf("error loading transactions list; %s", err.Error())
		return err
	}

	defer db.closeCursor(ld)

	// loop and load
	var trx *types.Transaction
	for ld.Next(ctx) {
		// process the last found hash
		if trx != nil {
			list.Collection = append(list.Collection, trx)
		}

		// try to decode the next row
		var row types.Transaction
		if err := ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode the list row; %s", err.Error())
			return err
		}

		// we have one
		trx = &row
	}

	// we should have all the items already; we may just need to check if a boundary was reached
	list.IsEnd = (cursor == nil && count < 0) || (count > 0 && int32(len(list.Collection)) < count)
	list.IsStart = (cursor == nil && count > 0) || (count < 0 && int32(len(list.Collection)) < -count)

	// add the last item as well if we hit the boundary
	if (list.IsStart || list.IsEnd) && trx != nil {
		list.Collection = append(list.Collection, trx)
	}
	return nil
}

// TransactionsCount returns the number of transactions stored in the database.
func (db *MongoDbBridge) TransactionsCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(colTransactions))
}

// Transactions pulls list of transaction hashes starting on the specified cursor.
func (db *MongoDbBridge) Transactions(cursor *string, count int32, filter *bson.D) (*types.TransactionList, error) {
	// nothing to load?
	if count == 0 {
		return nil, fmt.Errorf("nothing to do, zero transactions requested")
	}

	// get the collection and context
	col := db.client.Database(db.dbName).Collection(colTransactions)

	// init the list
	list, err := db.initTrxList(col, cursor, count, filter)
	if err != nil {
		db.log.Errorf("can not build transactions list; %s", err.Error())
		return nil, err
	}

	// load data if there are any
	if list.Total > 0 {
		err = db.txListLoad(col, cursor, count, list)
		if err != nil {
			db.log.Errorf("can not load transactions list from database; %s", err.Error())
			return nil, err
		}

		// reverse on negative so new-er transaction will be on top
		if count < 0 {
			list.Reverse()
			count = -count
		}

		// cut the end?
		if len(list.Collection) > int(count) {
			list.Collection = list.Collection[:len(list.Collection)-1]
		}
	}

	return list, nil
}

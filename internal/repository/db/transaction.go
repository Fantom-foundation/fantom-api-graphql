// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// coTransaction is the name of the off-chain database collection storing transaction details.
	coTransactions = "transaction"

	// fiTransactionPk is the name of the primary key field of the transaction collection.
	fiTransactionPk = "_id"

	// fiTransactionOrdinalIndex is the name of the transaction ordinal index in the blockchain field.
	// db.transaction.createIndex({_id:1,orx:-1},{unique:true})
	fiTransactionOrdinalIndex = "orx"

	// fiTransactionBlock is the name of the block number field of the transaction.
	fiTransactionBlock = "blk"

	// fiTransactionSender is the name of the address field of the sender account.
	// db.transaction.createIndex({from:1}).
	fiTransactionSender = "from"

	// fiTransactionRecipient is the name of the address field of the recipient account.
	// null for contract creation.
	// db.transaction.createIndex({to:1}).
	fiTransactionRecipient = "to"

	// fiTransactionContract is the name of the address field of the smart contract created.
	// null if not contract creation.
	fiTransactionContract = "sc"

	// fiTransactionValue is the name of the value transferred in WEI field.
	fiTransactionValue = "val"

	// fiTransactionTimestamp is the name of the transaction time stamp field.
	fiTransactionTimestamp = "ts"

	// fiTransactionPropagated is the name of the field indicating processed transaction.
	fiTransactionPropagated = "isok"

	// fiTransactionIsERC20Call is the name of the field indicating an ERC20 call.
	fiTransactionIsERC20Call = "iserc"

	// fiTransactionTargetContract is the name of the field of target smart contract type
	// this transaction addressed. If undetected, or not a contract address, the field is empty.
	fiTransactionTargetContract = "tc"

	// fiTransactionTargetCall is the name of the field of target smart contract function name.
	// If the function has not been detected yet, the field is empty.
	fiTransactionTargetCall = "call"
)

// initTransactionsCollection initializes the transaction collection with
// indexes and additional parameters needed by the app.
func (db *MongoDbBridge) initTransactionsCollection(col *mongo.Collection) {
	if !db.initTransactions {
		return
	}

	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index ordinal key along with the primary key
	unique := true
	ix = append(ix, mongo.IndexModel{
		Keys: bson.D{{fiTransactionPk, 1}, {fiTransactionOrdinalIndex, -1}},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})

	// index sender and recipient
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{fiTransactionSender, 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{fiTransactionRecipient, 1}}})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for transaction collection; %s", err.Error())
	}

	// log we done that
	db.initTransactions = false
	db.log.Debugf("transactions collection initialized")
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

// decodeOptAddresses decodes an options address for transaction saving process.
func decodeOptAddresses(adr *common.Address) *string {
	var result *string
	if adr != nil {
		val := adr.String()
		result = &val
	}
	return result
}

// TransactionDetails extends the transaction detail with additional information
// collected from the off-chain database if available.
func (db *MongoDbBridge) TransactionDetails(trx *types.Transaction) error {
	// missing the actual transaction base
	if trx == nil {
		return fmt.Errorf("can not extend empty transaction")
	}

	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(coTransactions)

	// try to find the result
	sr := col.FindOne(context.Background(),
		bson.D{{fiTransactionPk, trx.Hash.String()}},
		options.FindOne().SetProjection(bson.D{
			{fiTransactionOrdinalIndex, true},
			{fiTransactionTargetCall, true},
			{fiTransactionTargetContract, true},
			{fiTransactionPropagated, true},
			{fiTransactionIsERC20Call, true},
		}))
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			db.log.Debugf("transaction %s not stored in db yet", trx.Hash.String())
			return nil
		}

		// log issue here
		db.log.Error("can not get transaction data from DB; %s", sr.Err().Error())
		return sr.Err()
	}

	// try to decode the extended data
	db.log.Debugf("decoding extended information for %s", trx.Hash.String())
	if err := sr.Decode(trx); err != nil {
		// log issue here
		db.log.Error("can not decode transaction data from DB; %s", sr.Err().Error())
		return sr.Err()
	}

	// reset the target call
	if trx.TargetContractType != nil && *trx.TargetContractType == "" {
		trx.TargetContractType = nil
		trx.TargetFunctionCall = nil
	}

	// log the ordinal index of the found transaction
	db.log.Debugf("found transaction %s as #%d", trx.Hash.String(), trx.OrdinalIndex)
	return nil
}

// AddTransaction stores a transaction reference in connected persistent storage.
func (db *MongoDbBridge) AddTransaction(block *types.Block, trx *types.Transaction) error {
	// do we have all needed data?
	if block == nil || trx == nil {
		return fmt.Errorf("can not add empty transaction")
	}

	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(coTransactions)

	// if the transaction already exists, we don't need to add it
	// just make sure the transaction accounts were processed
	if !db.shouldAddTransaction(col, trx) {
		return nil
	}

	// calculate the ordinal index of the transaction
	trx.OrdinalIndex = types.TransactionIndex(block, trx)

	// try to do the insert
	if _, err := col.InsertOne(context.Background(),
		transactionData(&bson.D{
			{fiTransactionPk, trx.Hash.String()},
			{fiTransactionOrdinalIndex, trx.OrdinalIndex},
			{fiTransactionBlock, uint64(block.Number)},
			{fiTransactionTimestamp, uint64(block.TimeStamp)},
		}, trx)); err != nil {
		db.log.Critical(err)
		return err
	}

	// add transaction to the db
	db.log.Infof("transaction %s added to database", trx.Hash.String())

	// check init state
	db.initTransactionsCollection(col)
	return nil
}

// UpdateTransaction modifies the transaction data in the DB to match
// the values of the record.
func (db *MongoDbBridge) UpdateTransaction(trx *types.Transaction) error {
	// any transaction given?
	if trx == nil {
		return fmt.Errorf("no transaction given")
	}

	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(coTransactions)

	// update the transaction details
	if _, err := col.UpdateOne(context.Background(),
		bson.D{{fiTransactionPk, trx.Hash.String()}},
		bson.D{{"$set", transactionData(nil, trx)}}); err != nil {
		// log the issue
		db.log.Criticalf("transaction %s can not be updated; %s", trx.Hash.String(), err.Error())
		return err
	}

	return nil
}

// transactionData collects the data for the given transaction.
func transactionData(base *bson.D, trx *types.Transaction) bson.D {
	// make a new instance if needed
	if base == nil {
		base = &bson.D{}
	}

	// add the extended data
	*base = append(*base,
		bson.E{Key: fiTransactionSender, Value: trx.From.String()},
		bson.E{Key: fiTransactionRecipient, Value: decodeOptAddresses(trx.To)},
		bson.E{Key: fiTransactionContract, Value: decodeOptAddresses(trx.ContractAddress)},
		bson.E{Key: fiTransactionValue, Value: trx.Value.String()},
		bson.E{Key: fiTransactionTargetContract, Value: trx.TargetContractType},
		bson.E{Key: fiTransactionTargetCall, Value: trx.TargetFunctionCall},
		bson.E{Key: fiTransactionPropagated, Value: trx.IsProcessed},
		bson.E{Key: fiTransactionIsERC20Call, Value: trx.IsErc20Call},
	)
	return *base
}

// isTransactionKnown checks if a transaction document already exists in the database.
func (db *MongoDbBridge) IsTransactionKnown(col *mongo.Collection, hash *types.Hash) (bool, error) {
	// try to find the transaction in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{
		{fiTransactionPk, hash.String()},
	}, options.FindOne().SetProjection(bson.D{
		{fiTransactionPk, true},
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

// TransactionMarkPropagated marks given transaction as processed
// and ready to be served full to API users.
// AccountQueue processor call this function as a callback.
func (db *MongoDbBridge) TransactionMarkPropagated(trx *types.Transaction) error {
	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(coTransactions)

	// update the transaction details
	if _, err := col.UpdateOne(context.Background(),
		bson.D{{fiTransactionPk, trx.Hash.String()}},
		bson.D{{"$set", bson.D{{fiTransactionPropagated, true}}}}); err != nil {
		// log the issue
		db.log.Criticalf("transaction %s can not be updated; %s", trx.Hash.String(), err.Error())
		return err
	}

	// update local state and log we are done with the trx
	trx.IsProcessed = true
	db.log.Debugf("transaction %s state changed", trx.Hash.String())
	return nil
}

// LastKnownBlock returns number of the last known block stored in the database.
func (db *MongoDbBridge) LastKnownBlock() (uint64, error) {
	// prep search options
	opt := options.FindOne()
	opt.SetSort(bson.D{{fiTransactionBlock, -1}})
	opt.SetProjection(bson.D{{fiTransactionBlock, true}})

	// get the collection for account transactions
	col := db.client.Database(db.dbName).Collection(coTransactions)
	res := col.FindOne(context.Background(), bson.D{{fiTransactionPropagated, true}}, opt)
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
	var tx struct {
		Block uint64 `bson:"blk"`
	}

	// get the data
	err := res.Decode(&tx)
	if err != nil {
		db.log.Error("can not decode the top block")
		return 0, res.Err()
	}

	return tx.Block, nil
}

// initTrxList initializes list of transactions based on provided cursor and count.
func (db *MongoDbBridge) initTrxList(col *mongo.Collection, cursor *string, count int32, filter *bson.D) (*types.TransactionHashList, error) {
	// make sure some filter is used
	if nil == filter {
		filter = &bson.D{}
	}

	// find how many transactions do we have in the database
	total, err := col.CountDocuments(context.Background(), *filter)
	if err != nil {
		db.log.Errorf("can not count transactions")
		return nil, err
	}

	// make the list and notify the size of it
	db.log.Debugf("found %d filtered transactions", total)
	list := types.TransactionHashList{
		Collection: make([]*types.Hash, 0),
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
	list *types.TransactionHashList,
	cursor *string,
	count int32,
	filter *bson.D,
) (*types.TransactionHashList, error) {
	var err error

	// find out the cursor ordinal index
	if cursor == nil && count > 0 {
		// get the highest available ordinal index (top transaction)
		list.First, err = db.findBorderOrdinalIndex(col,
			*filter,
			options.FindOne().SetSort(bson.D{{fiTransactionOrdinalIndex, -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available ordinal index (top transaction)
		list.First, err = db.findBorderOrdinalIndex(col,
			*filter,
			options.FindOne().SetSort(bson.D{{fiTransactionOrdinalIndex, 1}}))
		list.IsEnd = true

	} else if cursor != nil {
		// get the highest available ordinal index (top transaction)
		list.First, err = db.findBorderOrdinalIndex(col,
			bson.D{{fiTransactionPk, *cursor}},
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
	opt.SetProjection(bson.D{{"orx", true}})
	sr := col.FindOne(context.Background(), filter, opt)

	// try to decode
	err := sr.Decode(&row)
	if err != nil {
		return 0, err
	}

	return row.Value, nil
}

// txListFilter creates a filter for transaction list search.
func (db *MongoDbBridge) txListFilter(cursor *string, count int32, list *types.TransactionHashList) *bson.D {
	// inform what we are about to do
	db.log.Debugf("transaction filter starts from index %d", list.First)

	// build the filter query
	if cursor == nil {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: fiTransactionOrdinalIndex, Value: bson.D{{"$lte", list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: fiTransactionOrdinalIndex, Value: bson.D{{"$gte", list.First}}})
		}
	} else {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: fiTransactionOrdinalIndex, Value: bson.D{{"$lt", list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: fiTransactionOrdinalIndex, Value: bson.D{{"$gt", list.First}}})
		}
	}

	// log the filter
	return &list.Filter
}

// txListOptions creates a filter options set for transactions list search.
func (db *MongoDbBridge) txListOptions(count int32) *options.FindOptions {
	// prep options
	opt := options.Find()
	opt.SetProjection(bson.D{{fiTransactionPk, true}, {fiTransactionOrdinalIndex, true}})

	// how to sort results in the collection
	if count > 0 {
		// from high (new) to low (old)
		opt.SetSort(bson.D{{fiTransactionOrdinalIndex, -1}})
	} else {
		// from low (old) to high (new)
		opt.SetSort(bson.D{{fiTransactionOrdinalIndex, 1}})
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
func (db *MongoDbBridge) txListLoad(col *mongo.Collection, cursor *string, count int32, list *types.TransactionHashList) error {
	// get the context for loader
	ctx := context.Background()

	// load the data
	ld, err := col.Find(ctx, db.txListFilter(cursor, count, list), db.txListOptions(count))
	if err != nil {
		db.log.Errorf("error loading transactions list; %s", err.Error())
		return err
	}

	// close the cursor as we leave
	defer func() {
		err := ld.Close(ctx)
		if err != nil {
			db.log.Errorf("error closing transactions list cursor; %s", err.Error())
		}
	}()

	// loop and load
	var hash *types.Hash
	for ld.Next(ctx) {
		// process the last found hash
		if hash != nil {
			list.Collection = append(list.Collection, hash)
		}

		// get the next hash
		var row struct {
			Id  string `bson:"_id"`
			Orx uint64 `bson:"orx"`
		}

		// try to decode the next row
		if err := ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode the list row; %s", err.Error())
			return err
		}

		// decode the value
		h := types.HexToHash(row.Id)
		hash = &h
	}

	// we should have all the items already; we may just need to check if a boundary was reached
	list.IsEnd = (cursor == nil && count < 0) || (count > 0 && int32(len(list.Collection)) < count)
	list.IsStart = (cursor == nil && count > 0) || (count < 0 && int32(len(list.Collection)) < -count)

	// add the last item as well if we hit the boundary
	if (list.IsStart || list.IsEnd) && hash != nil {
		list.Collection = append(list.Collection, hash)
	}
	return nil
}

// TransactionsCount returns the number of transactions stored in the database.
func (db *MongoDbBridge) TransactionsCount() (uint64, error) {
	// get the collection and context
	col := db.client.Database(db.dbName).Collection(coTransactions)

	// find how many transactions do we have in the database
	total, err := col.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		db.log.Errorf("can not count transactions")
		return 0, err
	}

	// inform what we are about to do
	db.log.Debugf("found %d transactions in off-chain database", total)
	return uint64(total), nil
}

// Transactions pulls list of transaction hashes starting on the specified cursor.
func (db *MongoDbBridge) Transactions(cursor *string, count int32, filter *bson.D) (*types.TransactionHashList, error) {
	// nothing to load?
	if count == 0 {
		return nil, fmt.Errorf("nothing to do, zero transactions requested")
	}

	// get the collection and context
	col := db.client.Database(db.dbName).Collection(coTransactions)

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
		}
	}

	return list, nil
}

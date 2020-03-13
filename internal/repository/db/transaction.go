package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
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
	fiTransactionOrdinalIndex = "orx"

	// fiTransactionBlock is the name of the block number field of the transaction.
	fiTransactionBlock = "blk"

	// fiTransactionSender is the name of the address field of the sender's account.
	fiTransactionSender = "from"

	// fiTransactionRecipient is the name of the address field of the recipients's account. null for contract creation.
	fiTransactionRecipient = "to"

	// fiTransactionValue is the name of the value transferred in WEI field.
	fiTransactionValue = "val"

	// fiTransactionTimestamp is the name of the transaction time stamp field.
	fiTransactionTimestamp = "ts"
)

/*
// tblTransaction represents a single transaction record in the database.
type tblTransaction struct {
	Id        common.Hash    `bson:"_id"`
	Block     uint64         `bson:"blk"`
	From      common.Address `bson:"from"`
	To        common.Address `bson:"to"`
	Value     hexutil.Big    `bson:"val"`
	Timestamp hexutil.Uint64 `bson:"tx"`
}
*/

// AddTransaction stores a transaction reference in connected persistent storage.
func (db *MongoDbBridge) AddTransaction(block *types.Block, trx *types.Transaction) error {
	// do we have all needed data?
	if block == nil || trx == nil {
		return fmt.Errorf("can not add empty transaction")
	}

	// check if the transaction already exists
	exists, err := db.isTransactionKnown(&trx.Hash)
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// if the transaction already exists, we don't need to do anything here
	if exists {
		return nil
	}

	// get the collection for transactions
	col := db.client.Database(offChainDatabaseName).Collection(coTransactions)

	// recipient address may not be defined so we need to do a bit more parsing
	var rcAddress *string = nil
	if trx.To != nil {
		rcp := trx.To.String()
		rcAddress = &rcp
	}

	// try to do the insert
	_, err = col.InsertOne(context.Background(), bson.D{
		{fiTransactionPk, trx.Hash.String()},
		{fiTransactionOrdinalIndex, trxOrdinalIndex(trx)},
		{fiTransactionBlock, uint64(block.Number)},
		{fiTransactionSender, trx.From.String()},
		{fiTransactionRecipient, rcAddress},
		{fiTransactionValue, trx.Value.String()},
		{fiTransactionTimestamp, uint64(block.TimeStamp)},
	})
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// add the transaction to the sender's address list
	return db.propagateTrxToAccounts(block, trx)
}

// getTrxOrdinalIndex calculates ordinal index in the whole blockchain.
func trxOrdinalIndex(trx *types.Transaction) uint64 {
	return (uint64(*trx.BlockNumber) << 14) | uint64(*trx.TrxIndex)
}

// propagateTrxToAccounts push given transaction to sender's account and also to recipient's account, if exists.
func (db *MongoDbBridge) propagateTrxToAccounts(block *types.Block, trx *types.Transaction) error {
	// propagate to sender
	sender := types.Account{Address: trx.From}
	err := db.AddAccountTransaction(&sender, block, trx)
	if err != nil {
		db.log.Error("can not push the transaction to sender account")
		return err
	}

	// do we have a receiving account? may not be present for contract creating transactions
	if trx.To != nil {
		recipient := types.Account{Address: *trx.To}
		err := db.AddAccountTransaction(&recipient, block, trx)
		if err != nil {
			db.log.Error("can not push the transaction to sender account")
			return err
		}
	}

	return nil
}

// isTransactionKnown checks if a transaction document already exists in the database.
func (db *MongoDbBridge) isTransactionKnown(hash *types.Hash) (bool, error) {
	// get the collection for account transactions
	col := db.client.Database(offChainDatabaseName).Collection(coTransactions)

	// try to find the account in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{
		{fiTransactionPk, hash.String()},
	}, options.FindOne().SetProjection(bson.D{{fiTransactionPk, true}}))

	// error on lookup?
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			return false, nil
		}

		db.log.Error("can not get existing transaction pk")
		return false, sr.Err()
	}

	return true, nil
}

// LastKnownBlock returns number of the last known block stored in the database.
func (db *MongoDbBridge) LastKnownBlock() (uint64, error) {
	// prep search options
	opt := options.FindOne()
	opt.SetSort(bson.D{{fiTransactionBlock, -1}})
	opt.SetProjection(bson.D{{fiTransactionBlock, true}})

	// get the collection for account transactions
	col := db.client.Database(offChainDatabaseName).Collection(coTransactions)
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

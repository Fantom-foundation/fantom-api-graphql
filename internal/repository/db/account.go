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
	"math/big"
	"strings"
)

const (
	// coAccount is the name of the off-chain database collection storing account details.
	coAccounts = "account"

	// fiAccountTransactionPk is the name of the primary key field of the account to transaction collection.
	fiAccountPk = "_id"

	// fiAccountActivity is the name of the account last activity timestamp field.
	fiAccountActivity = "act"

	// fiAccountBalance is the name of the current account balance field in special units (FTM * 1000).
	fiAccountBalance = "bal"

	// fiAccountTxList is the name of the list of account transactions sub-document structure.
	fiAccountTxList = "trx"

	// fiAccountTxHash is the name of the account transaction hash field.
	fiAccountTxHash = "hash"

	// fiAccountTxHashPath is the full path of the account transaction hash field.
	fiAccountTxHashPath = "trx.hash"

	// fiAccountTxDirection is the name of the account transaction direction field.
	// Direction is 0 for outgoing and 1 for incoming transactions.
	fiAccountTxDirection = "dir"

	// fiAccountTxValue is the name of the account transaction value field.
	fiAccountTxValue = "val"

	// fiAccountTxTimeStamp is the name of the account transaction time stamp field.
	fiAccountTxTimeStamp = "ts"
)

// how many WEIs are in out account balance units (FTM*100)
var weiInBalanceUnits = new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil)

// tblAccountTransaction represents a transaction sub-document under the account master document.
type tblAccountTransaction struct {
	Hash      string `bson:"hash" json:"hash"`
	Direction int8   `bson:"dir" json:"dir"`
	Value     string `bson:"val" json:"val"`
	TimeStamp uint64 `bson:"ts" json:"tx"`
}

// AddAccount stores an account in the blockchain if not exists.
func (db *MongoDbBridge) AddAccount(acc *types.Account) error {
	// do we have account data?
	if acc == nil {
		return fmt.Errorf("can not add empty account")
	}

	// check the account existence in the database
	exists, err := db.isAccountKnown(&acc.Address)
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// if the account already exists, we don't need to do anything here
	if exists {
		return nil
	}

	// get the collection for account transactions
	col := db.client.Database(offChainDatabaseName).Collection(coAccounts)

	// do the update based on given PK; we don't need to pull the document updated
	_, err = col.InsertOne(context.Background(), bson.D{
		{fiAccountPk, acc.Address.String()},
		{fiAccountBalance, "0x0"},
		{fiAccountActivity, nil},
		{fiAccountTxList, bson.A{}},
	})

	// error on lookup?
	if err != nil {
		db.log.Error("can not insert new account")
		return err
	}

	return nil
}

// AddAccountTransaction add a given transaction to the given account address.
func (db *MongoDbBridge) AddAccountTransaction(acc *types.Account, block *types.Block, trx *types.Transaction) error {
	// do we have all needed data?
	if block == nil || trx == nil {
		return fmt.Errorf("can not add empty transaction")
	}

	// make sure the account exists
	if err := db.AddAccount(acc); err != nil {
		return err
	}

	// get the collection for account transactions
	col := db.client.Database(offChainDatabaseName).Collection(coAccounts)

	// what is the direction
	var dir = 0
	if trx.To != nil && acc.Address.String() == trx.To.String() {
		dir = 1
	}

	// get account balance
	bal, err := db.fnBalance(acc)
	if err != nil {
		return fmt.Errorf("can not get account balance")
	}

	// calculate balance in FTM*100 units (so we keep some decimals without worrying about them)
	ftm := new(big.Int).Div(bal.ToInt(), weiInBalanceUnits)

	// add the transaction to the account set
	_, err = col.UpdateOne(context.Background(),
		bson.D{{fiAccountPk, acc.Address.String()}},
		bson.D{
			{"$set", bson.D{
				{fiAccountActivity, uint64(block.TimeStamp)},
				{fiAccountBalance, ftm.Uint64()},
			}},
			{"$addToSet", bson.D{
				{fiAccountTxList, bson.D{
					{fiAccountTxHash, trx.Hash.String()},
					{fiAccountTxDirection, dir},
					{fiAccountTxValue, trx.Value.String()},
					{fiAccountTxTimeStamp, uint64(block.TimeStamp)},
				}},
			}},
		})

	// error on update?
	if err != nil {
		db.log.Errorf("can not add transaction to account; %s", err.Error())
		return err
	}

	return nil
}

// isAccountTransactionKnown verifies if the transaction is already listed for the account address given.
func (db *MongoDbBridge) isAccountTransactionKnown(addr *common.Address, hash *types.Hash) (bool, error) {
	// get the collection for account transactions
	col := db.client.Database(offChainDatabaseName).Collection(coAccounts)

	// try to find the account in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{
		{fiAccountPk, addr.String()},
		{fiAccountTxHashPath, hash.String()},
	}, options.FindOne().SetProjection(bson.D{{fiAccountPk, true}}))

	// error on lookup?
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			return false, nil
		}

		db.log.Error("can not get existing account transaction hash record")
		return false, sr.Err()
	}

	return true, nil
}

// isAccountKnown checks if an account document already exists in the database.
func (db *MongoDbBridge) isAccountKnown(addr *common.Address) (bool, error) {
	// get the collection for account transactions
	col := db.client.Database(offChainDatabaseName).Collection(coAccounts)

	// try to find the account in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{
		{fiAccountPk, addr.String()},
	}, options.FindOne().SetProjection(bson.D{{fiAccountPk, true}}))

	// error on lookup?
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			return false, nil
		}

		db.log.Error("can not get existing account pk")
		return false, sr.Err()
	}

	return true, nil
}

// initAccountTrxList initializes a list of transaction for given account address.
func (db *MongoDbBridge) initAccountTrxList(col *mongo.Collection, addr *common.Address, cursor *string, count int32) (*types.TransactionHashList, error) {
	// inform what we are about to do
	db.log.Debugf("initializing transactions list for account [%s]", addr.String())

	// get the total number of transaction for the account
	total, err := db.AccountTrxCount(col, addr)
	if err != nil {
		db.log.Errorf("can not list transaction for the account")
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("found %d transactions on account [%s]", total, addr.String())

	// transaction records are stored as an array so we work with indexes
	// old transactions are on top and new on the bottom of the list since we $push to it
	var start int64
	var end int64

	// decide where do we start and where do we end the tx slice
	if cursor == nil && count > 0 {
		// no cursor and positive count => we need to go from bottom up (new to old)
		start = int64(total - uint64(count))
		end = int64(total) - 1

	} else if cursor == nil && count < 0 {
		// no cursor and negative count => we need to go from top to bottom (old to new)
		end = int64(-count)

	} else if cursor != nil {
		// get the cursor id
		cursorIdx, err := db.accountTrxIndex(col, addr, cursor)
		if err != nil {
			db.log.Errorf("invalid cursor value; %s", err.Error())
			return nil, err
		}

		// log the activity
		db.log.Debugf("trx [%s] index is %d", *cursor, cursorIdx)

		// positive count => new to old => button up
		if count > 0 {
			start = int64(cursorIdx) - int64(count)
			end = int64(cursorIdx) - 1
		} else {
			start = int64(cursorIdx) + 1
			end = start - int64(count) - 1
		}
	}

	// make sure we are not beyond the start
	if start < 0 {
		start = 0
	}

	// make sure we are not beyond the end
	if (uint64(end) + 1) >= total {
		end = int64(total) - 1
	}

	// make sure end is not above start
	if end < start {
		end = start
	}

	// inform what we are about to do
	db.log.Debugf("new account transactions list for %d items [%d:%d]", count, start, end)

	// prep an empty list marking already clear boundaries
	// please note we have newer trx-es at the bottom and the oldest is on top, so the boundary calc may seem odd
	list := types.TransactionHashList{
		Collection: make([]*types.Hash, 0),
		Total:      total,
		First:      uint64(start),
		Last:       uint64(end),
		IsStart:    (uint64(end) + 1) == total,
		IsEnd:      start == 0,
	}

	return &list, nil
}

// getAccountTrxPipeline creates an aggregation pipeline for account transactions list based on given criteria.
func (db *MongoDbBridge) accountTrxPipeline(addr *common.Address, first uint64, last uint64) bson.A {
	// build name of the aggregated array element
	var sb strings.Builder
	sb.WriteString("$")
	sb.WriteString(fiAccountTxList)

	// what is the size of the slice we need?
	pull := last - first + 1

	// use aggregation pipeline to pull the defined slice of the TX list
	pipeline := bson.A{
		/* match account */
		bson.D{{"$match",
			bson.D{{fiAccountPk, addr.String()}},
		}},

		/* use projection to get specified slice of the embedded documents */
		bson.D{{"$project",
			bson.D{{"trx",
				bson.D{{"$slice",
					bson.A{sb.String(), first, pull},
				}},
			}},
		}},
	}

	return pipeline
}

// getAccountTrxList loads the list of transaction records from the Mongo database.
func (db *MongoDbBridge) accountTrxList(col *mongo.Collection, addr *common.Address, first uint64, last uint64) ([]tblAccountTransaction, error) {
	// get context
	ctx := context.Background()

	// call for the data
	res, err := col.Aggregate(ctx, db.accountTrxPipeline(addr, first, last))
	if err != nil {
		db.log.Errorf("can not get aggregate value; %s", err.Error())
		return nil, err
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
		db.log.Error("account and/or transaction not found")
		return nil, err
	}

	// prep container; we are interested in just one value
	row := struct {
		Id  string                  `bson:"_id"`
		Trx []tblAccountTransaction `bson:"trx"`
	}{}

	// try to decode the response
	err = res.Decode(&row)
	if err != nil {
		db.log.Errorf("can not parse account record; %s", err.Error())
		return nil, err
	}

	return row.Trx, nil
}

// AccountTransactions loads list of transaction hashes of an account.
func (db *MongoDbBridge) AccountTransactions(acc *types.Account, cursor *string, count int32) (*types.TransactionHashList, error) {
	// nothing to load?
	if count == 0 {
		return nil, fmt.Errorf("nothing to do, zero blocks requested")
	}

	// no account given?
	if acc == nil {
		return nil, fmt.Errorf("can not list transactions of empty account")
	}

	// get the collection
	col := db.client.Database(offChainDatabaseName).Collection(coAccounts)

	// init the list
	list, err := db.initAccountTrxList(col, &acc.Address, cursor, count)
	if err != nil {
		db.log.Errorf("account transactions list init failed; %s", err.Error())
		return nil, err
	}

	// get the data
	trx, err := db.accountTrxList(col, &acc.Address, list.First, list.Last)
	if err != nil {
		db.log.Errorf("can not build account transactions list; %s", err.Error())
		return nil, err
	}

	// extract data
	for _, tx := range trx {
		hash := types.HexToHash(tx.Hash)
		list.Collection = append(list.Collection, &hash)
	}

	// reverse the collection since we always have newer transactions on top
	list.Reverse()
	return list, nil
}

// AccountTrxCount calculates total number of transaction associated with an account.
func (db *MongoDbBridge) AccountTrxCount(col *mongo.Collection, addr *common.Address) (uint64, error) {
	// no address?
	if addr == nil {
		return 0, fmt.Errorf("can not get number of transaction of empty account")
	}

	// build name of the aggregated array element
	var sb strings.Builder
	sb.WriteString("$")
	sb.WriteString(fiAccountTxList)

	// prep aggregation pipeline
	pipeline := bson.A{
		/* match account */
		bson.D{{"$match",
			bson.D{{fiAccountPk, addr.String()}},
		}},

		/* get the size of the embedded array */
		bson.D{{"$project",
			bson.D{{"value",
				bson.D{{"$size", sb.String()}},
			}},
		}},
	}

	return db.getAggregateValue(col, &pipeline)
}

// AccountCount calculates total number of accounts in the database.
func (db *MongoDbBridge) AccountCount() (hexutil.Uint64, error) {
	// get the collection for transactions
	col := db.client.Database(offChainDatabaseName).Collection(coAccounts)

	// do the counting
	val, err := col.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		db.log.Errorf("can not cound documents in accounts collection; %s", err.Error())
		return 0, err
	}

	return hexutil.Uint64(uint64(val)), nil
}

// accountTrxIndex finds index of a transaction in account transactions array.
func (db *MongoDbBridge) accountTrxIndex(col *mongo.Collection, addr *common.Address, hash *string) (uint64, error) {
	// no address or no hash?
	if addr == nil || hash == nil {
		return 0, fmt.Errorf("can not get transaction index of empty account or transaction hash")
	}

	// build name of the aggregated array element path
	var sb strings.Builder
	sb.WriteString("$")
	sb.WriteString(fiAccountTxHashPath)

	// prep aggregation pipeline for index lookup
	pipeline := bson.A{
		/* match account */
		bson.D{{"$match",
			bson.D{{fiAccountPk, addr.String()}},
		}},

		/* get index of the given hash in embedded array */
		bson.D{{"$project",
			bson.D{{"value",
				bson.D{{"$indexOfArray", bson.A{sb.String(), *hash}}},
			}},
		}},
	}

	return db.getAggregateValue(col, &pipeline)
}

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
	// coAccount is the name of the off-chain database collection storing account details.
	coAccounts = "account"

	// fiAccountTransactionPk is the name of the primary key field
	// of the account to transaction collection.
	fiAccountPk = "_id"

	// fiAccountType is the name of the field of the account contract type.
	fiAccountType = "type"

	// fiScCreationTx is the name of the field of the transaction hash
	// which created the contract, if the account is a contract.
	fiScCreationTx = "sc"
)

// the account base row
type AccountRow struct {
	Type   string      `bson:"type"`
	Sc     *string     `bson:"sc"`
	ScHash *types.Hash `bson:"-"`
}

// Account tries to load an account identified by the address given from
// the off-chain database.
func (db *MongoDbBridge) Account(addr *common.Address) (*types.Account, error) {
	// get the collection for account transactions
	col := db.client.Database(db.dbName).Collection(coAccounts)

	// try to find the account
	sr := col.FindOne(context.Background(), bson.D{{fiAccountPk, addr.String()}}, options.FindOne())

	// error on lookup?
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}

		db.log.Error("can not get existing account %s; %s", addr.String(), sr.Err().Error())
		return nil, sr.Err()
	}

	// try to decode the row
	var row AccountRow
	err := sr.Decode(&row)
	if err != nil {
		db.log.Error("can not decode account %s; %s", addr.String(), err.Error())
		return nil, err
	}

	// any hash?
	if row.Sc != nil {
		h := types.HexToHash(*row.Sc)
		row.ScHash = &h
	}

	return &types.Account{
		Address:    *addr,
		ContractTx: row.ScHash,
		Type:       row.Type,
	}, nil
}

// initAccountsCollection initializes the account collection with
// indexes and additional parameters needed by the app.
func (db *MongoDbBridge) initAccountsCollection() {
	if !db.initAccounts {
		return
	}

	// log we are done here
	db.initAccounts = false
	db.log.Debugf("accounts collection initialized")
}

// AddAccount stores an account in the blockchain if not exists.
func (db *MongoDbBridge) AddAccount(acc *types.Account) error {
	// do we have account data?
	if acc == nil {
		return fmt.Errorf("can not add empty account")
	}

	// get the collection for account transactions
	col := db.client.Database(db.dbName).Collection(coAccounts)

	// extract contract creation transaction if available
	var conTx *string
	if acc.ContractTx != nil {
		cx := acc.ContractTx.String()
		conTx = &cx
	}

	// do the update based on given PK; we don't need to pull the document updated
	_, err := col.InsertOne(context.Background(), bson.D{
		{fiAccountPk, acc.Address.String()},
		{fiScCreationTx, conTx},
		{fiAccountType, acc.Type},
	})

	// error on lookup?
	if err != nil {
		db.log.Error("can not insert new account")
		return err
	}

	// check init state
	db.initAccountsCollection()
	return nil
}

// isAccountKnown checks if an account document already exists in the database.
func (db *MongoDbBridge) IsAccountKnown(addr *common.Address) (bool, error) {
	// get the collection for account transactions
	col := db.client.Database(db.dbName).Collection(coAccounts)

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

// AccountCount calculates total number of accounts in the database.
func (db *MongoDbBridge) AccountCount() (hexutil.Uint64, error) {
	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(coAccounts)

	// do the counting
	val, err := col.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		db.log.Errorf("can not count documents in accounts collection; %s", err.Error())
		return 0, err
	}

	return hexutil.Uint64(uint64(val)), nil
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

	// log what we do here
	db.log.Debugf("loading transactions of %s", acc.Address.String())

	// make the filter for [(from = Account) OR (to = Account)]
	filter := bson.D{{"$or", bson.D{{"from", acc.Address.String()}, {"to", acc.Address.String()}}}}

	// return list of transactions filtered by the account
	return db.Transactions(cursor, count, &filter)
}

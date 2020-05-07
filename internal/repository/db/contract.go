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
	// coContract is the name of the off-chain database collection storing smart contract details.
	coContract = "contract"

	// fiContractPk is the name of the primary key field of the contract collection.
	fiContractPk = "_id"

	// fiContractOrdinalIndex is the name of the contract ordinal index in the blockchain.
	fiContractOrdinalIndex = "orx"

	// fiContractAddress is the name of the address field of the contract.
	fiContractAddress = "adr"

	// fiContractTransaction is the name of the contract creation transaction
	// field of the sender's account.
	fiContractTransaction = "tx"

	// fiContractTimestamp is the name of the contract time stamp field.
	fiContractTimestamp = "ts"

	// fiContractName is the name of the contract name field.
	fiContractName = "name"

	// fiContractVersion is the name of the contract version id field.
	fiContractVersion = "ver"

	// fiContractCompiler is the name of the contract compiler id field.
	fiContractCompiler = "cv"

	// fiContractSource is the name of the contract source code field.
	fiContractSource = "sol"

	// fiContractAbi is the name of the contract ABI field.
	fiContractAbi = "abi"

	// fiContractSourceValidated is the name of the contract source code
	// validation timestamp field.
	fiContractSourceValidated = "ok"
)

// AddContract stores a smart contract reference in connected persistent storage.
func (db *MongoDbBridge) AddContract(block *types.Block, trx *types.Transaction) error {
	// do we have all needed data?
	if block == nil || trx == nil {
		return fmt.Errorf("can not add empty contract transaction")
	}

	// make sure this is indeed a contract transaction
	if trx.ContractAddress == nil {
		return fmt.Errorf("not a contract creation transaction")
	}

	// get the collection for transactions
	col := db.client.Database(offChainDatabaseName).Collection(coContract)

	// check if the transaction already exists
	exists, err := db.isContractKnown(col, trx.ContractAddress)
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// if the transaction already exists, we don't need to do anything here
	if exists {
		return nil
	}

	// try to do the insert
	_, err = col.InsertOne(context.Background(), bson.D{
		{fiContractPk, trx.ContractAddress.String()},
		{fiContractOrdinalIndex, db.transactionIndex(block, trx)},
		{fiContractAddress, trx.ContractAddress.String()},
		{fiContractTransaction, trx.Hash.String()},
		{fiContractTimestamp, uint64(block.TimeStamp)},
		{fiContractName, nil},
		{fiContractVersion, nil},
		{fiContractCompiler, nil},
		{fiContractSource, nil},
		{fiContractAbi, nil},
		{fiContractSourceValidated, nil},
	})
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// inform and quit
	db.log.Debugf("added smart contract reference [%s]", trx.ContractAddress.String())
	return nil
}

// isContractKnown checks if a smart contract document already exists in the database.
func (db *MongoDbBridge) isContractKnown(col *mongo.Collection, addr *common.Address) (bool, error) {
	// try to find the contract in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{
		{fiContractPk, addr.String()},
	}, options.FindOne().SetProjection(bson.D{
		{fiContractPk, true},
	}))

	// error on lookup?
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			return false, nil
		}

		// inform that we can not get the PK; should not happen
		db.log.Error("can not get existing contract pk")
		return false, sr.Err()
	}

	return true, nil
}

// ContractTransaction returns contract creation transaction hash if available.
func (db *MongoDbBridge) ContractTransaction(addr *common.Address) (*types.Hash, error) {
	// get the contract details from database
	c, err := db.Contract(addr)
	if err != nil {
		db.log.Errorf("can not get the contract transaction for [%s]; %s", addr.String(), err.Error())
		return nil, err
	}

	// contract not found
	if c == nil {
		return nil, nil
	}

	// return the hash
	return &c.TransactionHash, nil
}

// Contract returns details of a smart contract stored in the Mongo database
// if available, or nil if contract does not exist.
func (db *MongoDbBridge) Contract(addr *common.Address) (*types.Contract, error) {
	// get the collection for transactions
	col := db.client.Database(offChainDatabaseName).Collection(coContract)

	// try to find the contract in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{{fiContractPk, addr.String()}})

	// error on lookup?
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}

		// inform that we can not get the PK; should not happen
		db.log.Errorf("can not get contract details; %s", sr.Err().Error())
		return nil, sr.Err()
	}

	// prep container
	var row struct {
		Id          string  `bson:"_id"`
		Orx         uint64  `bson:"orx"`
		Address     string  `bson:"adr"`
		Transaction string  `bson:"tx"`
		TimeStamp   uint64  `bson:"ts"`
		Name        *string `bson:"name"`
		Version     *string `bson:"ver"`
		Compiler    *string `bson:"cv"`
		SourceCode  *string `bson:"sol"`
		Abi         *string `bson:"abi"`
		Validated   *uint64 `bson:"ok"`
	}

	// try to decode
	err := sr.Decode(&row)
	if err != nil {
		db.log.Errorf("can not get contract details; %s", err.Error())
		return nil, err
	}

	// prep the contract
	con := types.Contract{
		Address:         common.HexToAddress(row.Address),
		TransactionHash: types.HexToHash(row.Transaction),
		TimeStamp:       (hexutil.Uint64)(row.TimeStamp),
	}

	// do we have the source code?
	if row.SourceCode != nil {
		con.SourceCode = *row.SourceCode
	}

	// do we have the ABI definition?
	if row.Abi != nil {
		con.Abi = *row.Abi
	}

	// do we have the contract name?
	if row.Name != nil {
		con.Name = *row.Name
	}

	// do we have the validation time stamp?
	if row.Validated != nil {
		val := *row.Validated
		con.Validated = (*hexutil.Uint64)(&val)
	}

	return &con, nil
}

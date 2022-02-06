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
	"strconv"
)

const (
	// coContract is the name of the off-chain database collection storing smart contract details.
	coContract = "contract"

	// fiContractPk is the name of the primary key field of the contract collection.
	fiContractPk = "_id"

	// fiContractOrdinalIndex is the name of the contract ordinal index in the blockchain.
	// db.contract.createIndex({_id:1,orx:-1},{unique:true})
	fiContractOrdinalIndex = "orx"

	// fiContractSourceValidated is the name of the contract source code
	// validation timestamp field.
	fiContractSourceValidated = "val"
)

// initContractsCollection initializes the contracts collection with
// indexes and additional parameters needed by the app.
func (db *MongoDbBridge) initContractsCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index ordinal key along with the primary key
	unique := true
	ix = append(ix, mongo.IndexModel{
		Keys: bson.D{{Key: fiContractPk, Value: 1}, {Key: fiContractOrdinalIndex, Value: -1}},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for contracts collection; %s", err.Error())
	}

	// log we done that
	db.log.Debugf("contracts collection initialized")
}

// AddContract stores a smart contract reference in connected persistent storage.
func (db *MongoDbBridge) AddContract(sc *types.Contract) error {
	// do we have all needed data?
	if sc == nil {
		return fmt.Errorf("can not add empty contract")
	}

	// get the collection for contracts
	col := db.client.Database(db.dbName).Collection(coContract)

	// check if the contract already exists
	exists, err := db.isContractKnown(col, &sc.Address)
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// if the contract already exists, we update it to match the new content
	if exists {
		db.log.Debugf("contract %s known, updating", sc.Address.String())
		return db.UpdateContract(sc)
	}

	// try to do the insert
	if _, err = col.InsertOne(context.Background(), sc); err != nil {
		db.log.Critical(err)
		return err
	}

	// make sure contracts collection is initialized
	if db.initContracts != nil {
		db.initContracts.Do(func() { db.initContractsCollection(col); db.initContracts = nil })
	}

	db.log.Debugf("added smart contract at %s", sc.Address.String())
	return nil
}

// UpdateContract updates smart contract information in database to reflect
// new validation or similar changes passed from repository.
func (db *MongoDbBridge) UpdateContract(sc *types.Contract) error {
	// complain about missing contract data
	if sc == nil {
		db.log.Criticalf("can not update empty contract")
		return fmt.Errorf("no contract given to update")
	}

	// get the collection for contracts
	col := db.client.Database(db.dbName).Collection(coContract)

	// update the contract details
	if _, err := col.UpdateOne(context.Background(),
		bson.D{{Key: fiContractPk, Value: sc.Address.String()}},
		bson.D{{Key: "$set", Value: sc}}); err != nil {
		// log the issue
		db.log.Errorf("can not update contract details at %s; %s", sc.Address.String(), err.Error())
		return err
	}

	return nil
}

// IsContractKnown checks if a smart contract document already exists in the database.
func (db *MongoDbBridge) IsContractKnown(addr *common.Address) bool {
	// check the contract existence in the database
	known, err := db.isContractKnown(db.client.Database(db.dbName).Collection(coContract), addr)
	if err != nil {
		return false
	}

	return known
}

// isContractKnown checks if a smart contract document already exists in the database.
func (db *MongoDbBridge) isContractKnown(col *mongo.Collection, addr *common.Address) (bool, error) {
	// try to find the contract in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{
		{Key: fiContractPk, Value: addr.String()},
	}, options.FindOne().SetProjection(bson.D{
		{Key: fiContractPk, Value: true},
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
func (db *MongoDbBridge) ContractTransaction(addr *common.Address) (*common.Hash, error) {
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
	col := db.client.Database(db.dbName).Collection(coContract)

	// try to find the contract in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{{Key: fiContractPk, Value: addr.String()}})

	// error on lookup?
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}

		// inform that we can not get the PK; should not happen
		db.log.Errorf("can not get contract %s details; %s", addr.String(), sr.Err().Error())
		return nil, sr.Err()
	}

	// try to decode the contract data
	var con types.Contract
	err := sr.Decode(&con)
	if err != nil {
		db.log.Errorf("can not decode contract %s details; %s", addr.String(), err.Error())
		return nil, err
	}

	// inform
	db.log.Debugf("loaded contract %s", addr.String())
	return &con, nil
}

// ContractCount calculates total number of contracts in the database.
func (db *MongoDbBridge) ContractCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(coContract))
}

// contractListTotal find the total amount of contracts for the criteria and populates the list
func (db *MongoDbBridge) contractListTotal(col *mongo.Collection, validatedOnly bool, list *types.ContractList) error {
	// validation filter
	filter := bson.D{}
	if validatedOnly {
		filter = bson.D{{Key: fiContractSourceValidated, Value: bson.D{{Key: "$ne", Value: nil}}}}
	}

	// find how many contracts do we have in the database
	total, err := col.CountDocuments(context.Background(), filter)
	if err != nil {
		db.log.Errorf("can not count contracts")
		return err
	}

	// apply the total count
	list.Total = uint64(total)
	return nil
}

// contractListTopFilter constructs a filter for finding the top item of the list.
// Consider creating DB index db.contract.createIndex({_id:1,orx:-1},{unique:true}).
func contractListTopFilter(validatedOnly bool, cursor *string) (*bson.D, error) {
	// what is the requested ordinal index from cursor, if any
	var ix uint64
	if cursor != nil {
		var err error

		// get the ordinal index based on cursor
		ix, err = strconv.ParseUint(*cursor, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid cursor value; %s", err.Error())
		}
	}

	// with cursor and any validation status
	filter := bson.D{}
	if cursor != nil && !validatedOnly {
		filter = bson.D{{Key: fiContractOrdinalIndex, Value: ix}}
	}

	// no cursor, but validation status filter
	if cursor == nil && validatedOnly {
		filter = bson.D{{Key: fiContractSourceValidated, Value: bson.D{{Key: "$ne", Value: nil}}}}
	}

	// with cursor and also the validation filter
	if cursor != nil && validatedOnly {
		filter = bson.D{{Key: fiContractSourceValidated, Value: bson.D{{Key: "$ne", Value: nil}}}, {Key: fiContractOrdinalIndex, Value: ix}}
	}
	return &filter, nil
}

// contractListTop find the first contract of the list based on provided criteria and populates the list.
func (db *MongoDbBridge) contractListTop(col *mongo.Collection, validatedOnly bool, cursor *string, count int32, list *types.ContractList) error {
	// get the filter
	filter, err := contractListTopFilter(validatedOnly, cursor)
	if err != nil {
		db.log.Errorf("can not find top contract for the list; %s", err.Error())
		return err
	}

	// find out the cursor ordinal index
	if cursor == nil && count > 0 {
		// get the highest available ordinal index (top smart contract)
		list.First, err = db.findBorderOrdinalIndex(col,
			*filter,
			options.FindOne().SetSort(bson.D{{Key: fiContractOrdinalIndex, Value: -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available ordinal index (bottom smart contract)
		list.First, err = db.findBorderOrdinalIndex(col,
			*filter,
			options.FindOne().SetSort(bson.D{{Key: fiContractOrdinalIndex, Value: 1}}))
		list.IsEnd = true

	} else if cursor != nil {
		// get the highest available ordinal index (top smart contract)
		list.First, err = db.findBorderOrdinalIndex(col,
			*filter,
			options.FindOne())
	}

	// check the error
	if err != nil {
		db.log.Errorf("can not find the initial contract")
		return err
	}
	return nil
}

// contractListInit initializes list of contracts based on provided cursor and count.
func (db *MongoDbBridge) contractListInit(col *mongo.Collection, validatedOnly bool, cursor *string, count int32) (*types.ContractList, error) {
	// make the list
	list := types.ContractList{
		Collection: make([]*types.Contract, 0),
		Total:      0,
		First:      0,
		Last:       0,
		IsStart:    false,
		IsEnd:      false,
	}

	// calculate the total number of contracts in the list
	if err := db.contractListTotal(col, validatedOnly, &list); err != nil {
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("found %d contracts in off-chain database", list.Total)

	// find the top contract of the list
	if err := db.contractListTop(col, validatedOnly, cursor, count, &list); err != nil {
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("contract list initialized with ordinal index %d", list.First)
	return &list, nil
}

// contractListFilter creates a filter for contract list search.
func (db *MongoDbBridge) contractListFilter(validatedOnly bool, cursor *string, count int32, list *types.ContractList) *bson.D {
	// inform what we are about to do
	db.log.Debugf("contract filter starts from index %d", list.First)
	ordinalOp := "$lte"

	// no cursor and bottom up list
	if cursor == nil && count < 0 {
		ordinalOp = "$gte"
	}

	// we have the cursor and we scan from top
	if cursor != nil && count > 0 {
		ordinalOp = "$lt"
	}

	// we have the cursor and we scan from bottom
	if cursor != nil && count < 0 {
		ordinalOp = "$gt"
	}

	// build the filter query
	var filter bson.D
	if validatedOnly {
		// filter validated only contracts
		filter = bson.D{
			{Key: fiContractOrdinalIndex, Value: bson.D{{Key: ordinalOp, Value: list.First}}},
			{Key: fiContractSourceValidated, Value: bson.D{{Key: "$ne", Value: nil}}},
		}
	} else {
		// filter all contracts
		filter = bson.D{{Key: fiContractOrdinalIndex, Value: bson.D{{Key: ordinalOp, Value: list.First}}}}
	}
	return &filter
}

// contractListOptions creates a filter options set for contract list search.
func (db *MongoDbBridge) contractListOptions(count int32) *options.FindOptions {
	// prep options
	opt := options.Find()

	// how to sort results in the collection
	if count > 0 {
		// from high (new) to low (old)
		opt.SetSort(bson.D{{Key: fiContractOrdinalIndex, Value: -1}})
	} else {
		// from low (old) to high (new)
		opt.SetSort(bson.D{{Key: fiContractOrdinalIndex, Value: 1}})
	}

	// prep the loading limit
	var limit = int64(count)
	if limit < 0 {
		limit = -limit
	}

	// try to get one more
	limit++

	// apply the limit
	opt.SetLimit(limit)
	return opt
}

// contractListLoad loads the initialized contract list from persistent database.
func (db *MongoDbBridge) contractListLoad(col *mongo.Collection, validatedOnly bool, cursor *string, count int32, list *types.ContractList) error {
	// get the context for loader
	ctx := context.Background()

	// load the data
	ld, err := col.Find(ctx, db.contractListFilter(validatedOnly, cursor, count, list), db.contractListOptions(count))
	if err != nil {
		db.log.Errorf("error loading contract list; %s", err.Error())
		return err
	}

	defer db.closeCursor(ld)

	// loop and load
	var contract *types.Contract
	for ld.Next(ctx) {
		// process the last found hash
		if contract != nil {
			list.Collection = append(list.Collection, contract)
			list.Last = contract.Uid()
		}

		// try to decode the next row
		var con types.Contract
		if err := ld.Decode(&con); err != nil {
			db.log.Errorf("can not decode contract the list row; %s", err.Error())
			return err
		}

		// keep this one
		contract = &con
	}

	// we should have all the items already; we may just need to check if a boundary was reached
	if contract != nil {
		list.IsEnd = count > 0 && int32(len(list.Collection)) < count
		list.IsStart = count < 0 && int32(len(list.Collection)) < -count

		// add the last item as well
		if list.IsStart || list.IsEnd {
			list.Collection = append(list.Collection, contract)
			list.Last = contract.Uid()
		}
	}

	return nil
}

// Contracts provides list of smart contracts stored in the persistent storage.
func (db *MongoDbBridge) Contracts(validatedOnly bool, cursor *string, count int32) (*types.ContractList, error) {
	// nothing to load?
	if count == 0 {
		return nil, fmt.Errorf("nothing to do, zero contracts requested")
	}

	// get the collection and context
	col := db.client.Database(db.dbName).Collection(coContract)

	// init the list
	list, err := db.contractListInit(col, validatedOnly, cursor, count)
	if err != nil {
		db.log.Errorf("can not build contract list; %s", err.Error())
		return nil, err
	}

	// load data
	err = db.contractListLoad(col, validatedOnly, cursor, count, list)
	if err != nil {
		db.log.Errorf("can not load contracts list from database; %s", err.Error())
		return nil, err
	}

	// shift the first item on cursor
	if cursor != nil {
		list.First = list.Collection[0].Uid()
	}

	// reverse on negative so new-er contracts will be on top
	if count < 0 {
		list.Reverse()
		count = -count
	}

	// cut the end?
	if len(list.Collection) > int(count) {
		list.Collection = list.Collection[:len(list.Collection)-1]
	}
	return list, nil
}

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

	// fiContractAddress is the name of the address field of the contract.
	fiContractAddress = "adr"

	// fiContractTransaction is the name of the contract creation transaction
	// field of the sender's account.
	fiContractTransaction = "tx"

	// fiContractTimestamp is the name of the contract time stamp field.
	fiContractTimestamp = "ts"

	// fiContractName is the name of the contract name field.
	fiContractName = "name"

	// fiContractSupport is the name of the contract support contact field.
	fiContractSupport = "sup"

	// fiContractVersion is the name of the contract version id field.
	fiContractVersion = "ver"

	// fiContractCompiler is the name of the contract compiler id field.
	fiContractCompiler = "cv"

	// fiContractLicense is the name of the contract open source license field.
	fiContractLicense = "lic"

	// fiContractIsOptimized is the name of the contract compiler optimizer enabled field.
	fiContractIsOptimized = "isOpt"

	// fiContractOptimizationRuns is the name of the contract compiler optimization runs field.
	fiContractOptimizationRuns = "runs"

	// fiContractSource is the name of the contract source code field.
	fiContractSource = "sol"

	// fiContractSourceHash is the name of the contract source code hash field.
	fiContractSourceHash = "soh"

	// fiContractAbi is the name of the contract ABI field.
	fiContractAbi = "abi"

	// fiContractSourceValidated is the name of the contract source code
	// validation timestamp field.
	fiContractSourceValidated = "ok"
)

// ContractRow defines a row in the Contract collection.
type contractRow struct {
	Id             string  `bson:"_id"`
	Orx            uint64  `bson:"orx"`
	Address        string  `bson:"adr"`
	Transaction    string  `bson:"tx"`
	TimeStamp      uint64  `bson:"ts"`
	Name           *string `bson:"name"`
	Support        *string `bson:"sup"`
	Version        *string `bson:"ver"`
	Compiler       *string `bson:"cv"`
	License        *string `bson:"lic"`
	IsOptimized    bool    `bson:"isOpt"`
	OptimizerRuns  uint64  `bson:"runs"`
	SourceCode     *string `bson:"sol"`
	SourceCodeHash *string `bson:"soh"`
	Abi            *string `bson:"abi"`
	Validated      *uint64 `bson:"ok"`
}

// initContractsCollection initializes the contracts collection with
// indexes and additional parameters needed by the app.
func (db *MongoDbBridge) initContractsCollection(col *mongo.Collection) {
	if !db.initContracts {
		return
	}

	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index ordinal key along with the primary key
	unique := true
	ix = append(ix, mongo.IndexModel{
		Keys: bson.D{{fiContractPk, 1}, {fiContractOrdinalIndex, -1}},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for contracts collection; %s", err.Error())
	}

	// log we done that
	db.initContracts = false
	db.log.Debugf("contracts collection initialized")
}

// AddContract stores a smart contract reference in connected persistent storage.
func (db *MongoDbBridge) AddContract(con *types.Contract) error {
	// do we have all needed data?
	if con == nil {
		return fmt.Errorf("can not add empty contract")
	}

	// get the collection for contracts
	col := db.client.Database(db.dbName).Collection(coContract)

	// check if the contract already exists
	exists, err := db.isContractKnown(col, &con.Address)
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// if the contract already exists, we update it to match the new content
	if exists {
		db.log.Debugf("contract %s known, updating", con.Address.String())
		return db.UpdateContract(con)
	}

	// get base contract data
	base := bson.D{
		{fiContractPk, con.Address.String()},
		{fiContractOrdinalIndex, con.OrdinalIndex},
		{fiContractAddress, con.Address.String()},
		{fiContractTransaction, con.TransactionHash.String()},
		{fiContractTimestamp, uint64(con.TimeStamp)},
	}

	// try to do the insert
	if _, err = col.InsertOne(context.Background(), contractData(con, base)); err != nil {
		db.log.Critical(err)
		return err
	}

	// inform and quit
	db.log.Debugf("added smart contract reference [%s]", con.Address.String())

	// init the collection
	db.initContractsCollection(col)
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

	// check if the contract already exists
	exists, err := db.isContractKnown(col, &sc.Address)
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// if contract can not be found, we can update
	if !exists {
		db.log.Errorf("can not find contract %s for update", sc.Address.String())
		return fmt.Errorf("contract not found")
	}

	// update the contract details
	if _, err = col.UpdateOne(context.Background(),
		bson.D{{fiContractPk, sc.Address.String()}},
		bson.D{{"$set", contractData(sc, nil)}}); err != nil {
		// log the issue
		db.log.Errorf("can not update contract details; %s", err.Error())
		return err
	}

	return nil
}

// contractData collects the contract data into the db structure we use
// for insert/update operations.
func contractData(sc *types.Contract, data bson.D) bson.D {
	// make sure we have the container
	if data == nil {
		data = bson.D{}
	}

	// append common data
	data = append(data,
		bson.E{Key: fiContractName, Value: sc.Name},
		bson.E{Key: fiContractSupport, Value: sc.SupportContact},
		bson.E{Key: fiContractVersion, Value: sc.Version},
		bson.E{Key: fiContractCompiler, Value: sc.Compiler},
		bson.E{Key: fiContractLicense, Value: sc.License},
		bson.E{Key: fiContractIsOptimized, Value: sc.IsOptimized},
		bson.E{Key: fiContractOptimizationRuns, Value: sc.OptimizeRuns},
		bson.E{Key: fiContractSource, Value: sc.SourceCode},
		bson.E{Key: fiContractAbi, Value: sc.Abi},
	)

	// do we have the source code hash?
	if sc.SourceCodeHash != nil {
		data = append(data, bson.E{Key: fiContractSourceHash, Value: sc.SourceCodeHash.String()})
	} else {
		data = append(data, bson.E{Key: fiContractSourceHash, Value: nil})
	}

	// do we have the validation mark?
	if sc.Validated != nil {
		data = append(data, bson.E{Key: fiContractSourceValidated, Value: uint64(*sc.Validated)})
	} else {
		data = append(data, bson.E{Key: fiContractSourceValidated, Value: nil})
	}

	return data
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

// newContract creates a new contract structure from provided DB row record
func newContract(row *contractRow) *types.Contract {
	// prep the contract
	con := types.Contract{
		OrdinalIndex:    row.Orx,
		Address:         common.HexToAddress(row.Address),
		TransactionHash: types.HexToHash(row.Transaction),
		TimeStamp:       (hexutil.Uint64)(row.TimeStamp),
		IsOptimized:     row.IsOptimized,
		OptimizeRuns:    int32(row.OptimizerRuns),
	}

	// do we have the source code?
	if row.SourceCode != nil {
		con.SourceCode = *row.SourceCode
	}

	// do we have the source code hash?
	if row.SourceCodeHash != nil {
		hash := types.HexToHash(*row.SourceCodeHash)
		con.SourceCodeHash = &hash
	}

	// do we have the ABI definition?
	if row.Abi != nil {
		con.Abi = *row.Abi
	}

	// do we have the contract name?
	if row.Name != nil {
		con.Name = *row.Name
	}

	// do we have the support contact name?
	if row.Support != nil {
		con.SupportContact = *row.Support
	}

	// do we have the contract version?
	if row.Version != nil {
		con.Version = *row.Version
	}

	// do we have the contract License?
	if row.License != nil {
		con.License = *row.License
	}

	// do we have the contract compiler?
	if row.Compiler != nil {
		con.Compiler = *row.Compiler
	}

	// do we have the validation time stamp?
	if row.Validated != nil {
		val := *row.Validated
		con.Validated = (*hexutil.Uint64)(&val)
	}

	return &con
}

// Contract returns details of a smart contract stored in the Mongo database
// if available, or nil if contract does not exist.
func (db *MongoDbBridge) Contract(addr *common.Address) (*types.Contract, error) {
	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(coContract)

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
	var row contractRow

	// try to decode
	err := sr.Decode(&row)
	if err != nil {
		db.log.Errorf("can not get contract details; %s", err.Error())
		return nil, err
	}

	return newContract(&row), nil
}

// ContractCount calculates total number of contracts in the database.
func (db *MongoDbBridge) ContractCount() (uint64, error) {
	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(coContract)

	// do the counting
	val, err := col.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		db.log.Errorf("can not count documents in contracts collection; %s", err.Error())
		return 0, err
	}

	return uint64(val), nil
}

// contractListTotal find the total amount of contracts for the criteria and populates the list
func (db *MongoDbBridge) contractListTotal(col *mongo.Collection, validatedOnly bool, list *types.ContractList) error {
	// prep the empty filter
	filter := bson.D{}

	// validation filter
	if validatedOnly {
		filter = bson.D{{db.dbName, bson.D{{"$ne", nil}}}}
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

	// prep the empty filter (no cursor and any validation status)
	filter := bson.D{}

	// with cursor and any validation status
	if cursor != nil && !validatedOnly {
		filter = bson.D{{fiContractOrdinalIndex, ix}}
	}

	// no cursor, but validation status filter
	if cursor == nil && validatedOnly {
		filter = bson.D{{fiContractSourceValidated, bson.D{{"$ne", nil}}}}
	}

	// with cursor and also the validation filter
	if cursor != nil && validatedOnly {
		filter = bson.D{{fiContractSourceValidated, bson.D{{"$ne", nil}}}, {fiContractOrdinalIndex, ix}}
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
			options.FindOne().SetSort(bson.D{{fiContractOrdinalIndex, -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available ordinal index (bottom smart contract)
		list.First, err = db.findBorderOrdinalIndex(col,
			*filter,
			options.FindOne().SetSort(bson.D{{fiContractOrdinalIndex, 1}}))
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

	// prep base operator
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
			{fiContractOrdinalIndex, bson.D{{ordinalOp, list.First}}},
			{fiContractSourceValidated, bson.D{{"$ne", nil}}},
		}
	} else {
		// filter all contracts
		filter = bson.D{{fiContractOrdinalIndex, bson.D{{ordinalOp, list.First}}}}
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
		opt.SetSort(bson.D{{fiContractOrdinalIndex, -1}})
	} else {
		// from low (old) to high (new)
		opt.SetSort(bson.D{{fiContractOrdinalIndex, 1}})
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

	// close the cursor as we leave
	defer func() {
		err := ld.Close(ctx)
		if err != nil {
			db.log.Errorf("error closing contract list cursor; %s", err.Error())
		}
	}()

	// loop and load
	var contract *types.Contract
	for ld.Next(ctx) {
		// process the last found hash
		if contract != nil {
			list.Collection = append(list.Collection, contract)
			list.Last = contract.OrdinalIndex
		}

		// get the next hash
		var row contractRow

		// try to decode the next row
		if err := ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode contract the list row; %s", err.Error())
			return err
		}

		// decode the value
		contract = newContract(&row)
	}

	// we should have all the items already; we may just need to check if a boundary was reached
	if cursor != nil {
		list.IsEnd = count > 0 && int32(len(list.Collection)) < count
		list.IsStart = count < 0 && int32(len(list.Collection)) < -count

		// add the last item as well
		if (list.IsStart || list.IsEnd) && contract != nil {
			list.Collection = append(list.Collection, contract)
			list.Last = contract.OrdinalIndex
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
		list.First = list.Collection[0].OrdinalIndex
	}

	// reverse on negative so new-er contracts will be on top
	if count < 0 {
		list.Reverse()
	}

	return list, nil
}

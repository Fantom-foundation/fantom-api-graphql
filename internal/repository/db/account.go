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
	// coAccount is the name of the off-chain database collection storing account details.
	colAccounts = "account"

	// fiAccountAddress represents account address column inside the collection.
	fiAccountAddress = "_id"

	// fiAccountName represents account name column inside the collection.
	fiAccountName = "name"

	// fiAccountType represents account type column inside the collection.
	fiAccountType = "act"

	// fiAccountIsContract represents is contract column inside the collection.
	fiAccountIsContract = "is_code"

	// fiAccountFirstAppearance represents first appearance column inside the collection.
	fiAccountFirstAppearance = "since"

	// fiAccountDeployedBy represents deployer column inside the collection.
	fiAccountDeployedBy = "deployer"

	// fiAccountDeploymentTrx represents deployment transaction column inside the collection.
	fiAccountDeploymentTrx = "dtx"

	// fiAccountContractUid represents contract uid column inside the collection.
	fiAccountContractUid = "cuid"

	// defaultTokenListLength is the number of ERC20 tokens pulled by default on negative count
	defaultTokenListLength = 25
)

// Account tries to load an account identified by the address given from
// the off-chain database.
func (db *MongoDbBridge) Account(addr *common.Address) (*types.Account, error) {
	col := db.client.Database(db.dbName).Collection(colAccounts)

	// try to find the account
	sr := col.FindOne(context.Background(), bson.D{{Key: defaultPK, Value: addr}}, options.FindOne())
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}

		db.log.Error("can not get existing account %s; %s", addr.String(), sr.Err().Error())
		return nil, sr.Err()
	}

	// try to decode the row
	var row types.Account
	if err := sr.Decode(&row); err != nil {
		db.log.Error("can not decode account %s; %s", addr.String(), err.Error())
		return nil, err
	}

	return &row, nil
}

// StoreAccount stores a unique account record in the database.
func (db *MongoDbBridge) StoreAccount(acc *types.Account) error {
	// do we have account data?
	if acc == nil {
		return fmt.Errorf("can not add empty account")
	}

	// insert the account if not exists
	col := db.client.Database(db.dbName).Collection(colAccounts)
	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{
			{Key: fiAccountAddress, Value: acc.Address},
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: fiAccountName, Value: acc.Name},
				{Key: fiAccountType, Value: acc.AccountType},
				{Key: fiAccountIsContract, Value: acc.IsContract},
				{Key: fiAccountFirstAppearance, Value: acc.FirstAppearance},
				{Key: fiAccountDeployedBy, Value: acc.DeployedBy},
				{Key: fiAccountDeploymentTrx, Value: acc.DeploymentTrx},
				{Key: fiAccountContractUid, Value: acc.ComputedContractUid()},
			}},
			{Key: "$setOnInsert", Value: bson.D{
				{Key: fiAccountAddress, Value: acc.Address},
			}},
		},
		options.Update().SetUpsert(true),
	); err != nil {
		db.log.Errorf("can not add account %s; %s", acc.Address.String(), err)
		return err
	}

	db.log.Debugf("added account at %s", acc.Address.String())
	return nil
}

// IsAccountKnown checks if an account document already exists in the database.
func (db *MongoDbBridge) IsAccountKnown(addr *common.Address) (bool, error) {
	col := db.client.Database(db.dbName).Collection(colAccounts)

	sr := col.FindOne(context.Background(), bson.D{{Key: defaultPK, Value: addr}}, options.FindOne().SetProjection(bson.D{{Key: defaultPK, Value: true}}))
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			return false, nil
		}

		db.log.Error("can not get existing account pk")
		return false, sr.Err()
	}
	return true, nil
}

// AccountCount calculates total number of accounts in the database.
func (db *MongoDbBridge) AccountCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(colAccounts))
}

// AccountTransactions loads list of transaction hashes of an account.
func (db *MongoDbBridge) AccountTransactions(addr *common.Address, rec *common.Address, cursor *string, count int32) (*types.TransactionList, error) {
	if addr == nil {
		return nil, fmt.Errorf("can not list transactions of empty account")
	}

	db.log.Debugf("loading transactions of %s", addr.String())
	if rec == nil {
		// return list of transactions
		filter := bson.D{
			{Key: "$or", Value: bson.A{
				bson.D{{Key: "from", Value: addr.String()}},
				bson.D{{Key: "to", Value: addr.String()}},
			}}}
		return db.Transactions(cursor, count, &filter)
	}

	// return list of transactions filtered by the account and recipient
	filter := bson.D{{Key: "from", Value: addr.String()}, {Key: "to", Value: rec.String()}}
	return db.Transactions(cursor, count, &filter)
}

// ErcTokensList returns a list of known ERC tokens.
func (db *MongoDbBridge) ErcTokensList(count int32, accountType int) ([]common.Address, error) {
	// make sure the count is positive; use default size if not
	if count <= 0 {
		count = defaultTokenListLength
	}

	// log what we do
	db.log.Debugf("loading %d ERC token accounts", count)

	// get the collection for contracts
	col := db.client.Database(db.dbName).Collection(colAccounts)

	// make the filter for ERC tokens only and pull them ordered by activity
	filter := bson.D{{Key: fiAccountType, Value: accountType}}
	opt := options.Find().SetLimit(int64(count))

	// load the data
	cursor, err := col.Find(context.Background(), filter, opt)
	if err != nil {
		db.log.Errorf("error loading ERC tokens list; %s", err.Error())
		return nil, err
	}

	return db.loadAccountAddressList(cursor)
}

// Contracts provides list of smart contracts stored in the persistent storage.
func (db *MongoDbBridge) Contracts(cursor *string, count int32) (*types.ContractList, error) {
	// nothing to load?
	if count == 0 {
		return nil, fmt.Errorf("nothing to do, zero contracts requested")
	}

	// get the collection and context
	col := db.client.Database(db.dbName).Collection(colAccounts)

	filter := bson.D{bson.E{Key: fiAccountType, Value: bson.D{{Key: "$ne", Value: types.AccountTypeWallet}}}}

	// init the list
	list, err := db.contractListInit(col, cursor, count, &filter)
	if err != nil {
		db.log.Errorf("can not build contract list; %s", err.Error())
		return nil, err
	}

	// load data
	err = db.contractListLoad(col, cursor, count, list)
	if err != nil {
		db.log.Errorf("can not load contracts list from database; %s", err.Error())
		return nil, err
	}

	// check collection has data
	if len(list.Collection) == 0 {
		return list, nil
	}

	// shift the first item on cursor
	if cursor != nil {
		list.First = uint64(*list.Collection[0].ContractUid)
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

// loadAccountAddressList loads the list of addresses from given cursor.
func (db *MongoDbBridge) loadAccountAddressList(cursor *mongo.Cursor) ([]common.Address, error) {
	// close the cursor as we leave
	defer db.closeCursor(cursor)

	// loop and load
	list := make([]common.Address, 0)
	for cursor.Next(context.Background()) {
		var row types.Account
		if err := cursor.Decode(&row); err != nil {
			db.log.Errorf("can not decode ERC contracts list row; %s", err.Error())
			return nil, err
		}

		list = append(list, row.Address)
	}

	return list, nil
}

// contractListInit initializes list of contracts based on provided cursor and count.
func (db *MongoDbBridge) contractListInit(col *mongo.Collection, cursor *string, count int32, filter *bson.D) (*types.ContractList, error) {
	total, err := col.CountDocuments(context.Background(), *filter)
	if err != nil {
		db.log.Errorf("can not count contracts")
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("found %d contracts in off-chain database", total)
	list := types.ContractList{
		Collection: make([]*types.Account, 0),
		Total:      uint64(total),
		First:      0,
		Last:       0,
		IsStart:    total == 0,
		IsEnd:      total == 0,
		Filter:     *filter,
	}

	// is the list non-empty? return the list with properly calculated range marks
	if 0 < total {
		return db.contractListCollectRangeMarks(col, &list, cursor, count)
	}
	// this is an empty list
	db.log.Debug("empty transaction list created")
	return &list, nil
}

// contractListCollectRangeMarks returns a list of contracts with proper First/Last marks.
func (db *MongoDbBridge) contractListCollectRangeMarks(col *mongo.Collection, list *types.ContractList, cursor *string, count int32) (*types.ContractList, error) {
	var err error

	// find out the cursor ordinal index
	if cursor == nil && count > 0 {
		// get the highest available pk
		list.First, err = db.contractListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: fiAccountContractUid, Value: -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available pk
		list.First, err = db.contractListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: fiAccountContractUid, Value: 1}}))
		list.IsEnd = true

	} else if cursor != nil {
		// the cursor itself is the starting point
		list.First, err = db.contractListBorderPk(col,
			bson.D{{Key: defaultPK, Value: *cursor}},
			options.FindOne())
	}

	// check the error
	if err != nil {
		db.log.Errorf("can not find the initial contract")
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("Contract list initialized with ordinal %d", list.First)
	return list, nil
}

// contractListBorderPk finds the top PK of the contracts collection based on given filter and options.
func (db *MongoDbBridge) contractListBorderPk(col *mongo.Collection, filter bson.D, opt *options.FindOneOptions) (uint64, error) {
	// prep container
	var row struct {
		Value int64 `bson:"cuid"`
	}

	// make sure we pull only what we need
	opt.SetProjection(bson.D{{Key: fiAccountContractUid, Value: true}})

	// try to decode
	sr := col.FindOne(context.Background(), filter, opt)
	err := sr.Decode(&row)
	if err != nil {
		return 0, err
	}
	return uint64(row.Value), nil
}

// contractListLoad load the initialized list of contracts from database.
func (db *MongoDbBridge) contractListLoad(col *mongo.Collection, cursor *string, count int32, list *types.ContractList) (err error) {
	// get the context for loader
	ctx := context.Background()

	// load the data
	ld, err := col.Find(ctx, db.contractListFilter(cursor, count, list), db.contractListOptions(count))
	if err != nil {
		db.log.Errorf("error loading contracts list; %s", err.Error())
		return err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			db.log.Errorf("error closing contracts list cursor; %s", err.Error())
		}
	}()

	// loop and load the list; we may not store the last value
	var acc *types.Account
	for ld.Next(ctx) {
		// append a previous value to the list, if we have one
		if acc != nil {
			list.Collection = append(list.Collection, acc)
		}

		// try to decode the next row
		var row types.Account
		if err = ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode the contract list row; %s", err.Error())
			return err
		}

		// use this row as the next item
		acc = &row
	}

	// we should have all the items already; we may just need to check if a boundary was reached
	list.IsEnd = (cursor == nil && count < 0) || (count > 0 && int32(len(list.Collection)) < count)
	list.IsStart = (cursor == nil && count > 0) || (count < 0 && int32(len(list.Collection)) < -count)

	// add the last item as well if we hit the boundary
	if ((count < 0 && list.IsStart) || (count > 0 && list.IsEnd)) && acc != nil {
		list.Collection = append(list.Collection, acc)
	}
	return nil
}

// contractListFilter creates a filter for contract list loading.
func (db *MongoDbBridge) contractListFilter(cursor *string, count int32, list *types.ContractList) *bson.D {
	// build an extended filter for the query; add PK (decoded cursor) to the original filter
	if cursor == nil {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: fiAccountContractUid, Value: bson.D{{Key: "$lte", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: fiAccountContractUid, Value: bson.D{{Key: "$gte", Value: list.First}}})
		}
	} else {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: fiAccountContractUid, Value: bson.D{{Key: "$lt", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: fiAccountContractUid, Value: bson.D{{Key: "$gt", Value: list.First}}})
		}
	}
	// return the new filter
	return &list.Filter
}

// contractListOptions creates a filter options set for contract list search.
func (db *MongoDbBridge) contractListOptions(count int32) *options.FindOptions {
	// prep options
	opt := options.Find()

	// how to sort results in the collection
	// from high (new) to low (old) by default; reversed if loading from bottom
	sd := -1
	if count < 0 {
		sd = 1
	}

	// sort with the direction we want
	opt.SetSort(bson.D{{Key: fiAccountContractUid, Value: sd}})

	// prep the loading limit
	var limit = int64(count)
	if limit < 0 {
		limit = -limit
	}

	// apply the limit, try to get one more record so we can detect list end
	opt.SetLimit(limit + 1)
	return opt
}

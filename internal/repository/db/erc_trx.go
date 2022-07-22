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
	"time"
)

// colErcTransactions represents the name of the ERC20 transaction collection in database.
const colErcTransactions = "erc20trx"

// initErc20TrxCollection initializes the ERC20 transaction list collection with
// indexes and additional parameters needed by the app.
func (db *MongoDbBridge) initErc20TrxCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index specific elements
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiTokenTransactionToken, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiTokenTransactionSender, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiTokenTransactionRecipient, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiTokenTransactionOrdinal, Value: -1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiTokenTransactionCallHash, Value: 1}}})

	// sender + ordinal index
	tox := "to_tok"
	ix = append(ix, mongo.IndexModel{
		Keys: bson.D{{Key: types.FiTokenTransactionRecipient, Value: 1}, {Key: types.FiTokenTransactionToken, Value: 1}},
		Options: &options.IndexOptions{
			Name: &tox,
		},
	})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for ERC20 trx collection; %s", err.Error())
	}

	// log we are done that
	db.log.Debugf("ERC20 trx collection initialized")
}

// AddERC20Transaction stores an ERC20 transaction in the database if it doesn't exist.
func (db *MongoDbBridge) AddERC20Transaction(trx *types.TokenTransaction) error {
	// get the collection for delegations
	col := db.client.Database(db.dbName).Collection(colErcTransactions)

	// is it a new one?
	if db.isErcTransactionKnown(col, trx) {
		return nil
	}

	// try to do the insert
	if _, err := col.InsertOne(context.Background(), trx); err != nil {
		db.log.Critical(err)
		return err
	}

	// make sure delegation collection is initialized
	if db.initErc20Trx != nil {
		db.initErc20Trx.Do(func() { db.initErc20TrxCollection(col); db.initErc20Trx = nil })
	}
	return nil
}

// isErcTransactionKnown checks if the given delegation exists in the database.
func (db *MongoDbBridge) isErcTransactionKnown(col *mongo.Collection, trx *types.TokenTransaction) bool {
	// try to find the delegation in the database
	sr := col.FindOne(context.Background(), bson.D{
		{Key: types.FiTokenTransactionPk, Value: trx.Pk()},
	}, options.FindOne().SetProjection(bson.D{
		{Key: types.FiTokenTransactionPk, Value: true},
	}))

	// error on lookup?
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			return false
		}
		// inform that we can not get the PK; should not happen
		db.log.Errorf("can not get existing ERC transaction pk; %s", sr.Err().Error())
		return false
	}
	return true
}

// ErcTransactionCountFiltered calculates total number of ERC20 transactions
// in the database for the given filter.
func (db *MongoDbBridge) ErcTransactionCountFiltered(filter *bson.D) (uint64, error) {
	return db.CountFiltered(db.client.Database(db.dbName).Collection(colErcTransactions), filter)
}

// ErcTransactionCount calculates total number of ERC20 transactions in the database.
func (db *MongoDbBridge) ErcTransactionCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(colErcTransactions))
}

// ercTrxListInit initializes list of ERC20 transactions based on provided cursor, count, and filter.
func (db *MongoDbBridge) ercTrxListInit(col *mongo.Collection, cursor *string, count int32, filter *bson.D) (*types.TokenTransactionList, error) {
	// make sure some filter is used
	if nil == filter {
		filter = &bson.D{}
	}

	// find how many transactions do we have in the database
	total, err := db.ErcTransactionCountFiltered(filter)
	if err != nil {
		db.log.Errorf("can not count ERC20 transactions")
		return nil, err
	}

	// make the list and notify the size of it
	db.log.Debugf("found %d filtered ERC20 transactions", total)
	list := types.TokenTransactionList{
		Collection: make([]*types.TokenTransaction, 0),
		Total:      total,
		First:      0,
		Last:       0,
		IsStart:    total == 0,
		IsEnd:      total == 0,
		Filter:     *filter,
	}

	// is the list non-empty? return the list with properly calculated range marks
	if 0 < total {
		return db.ercTrxListCollectRangeMarks(col, &list, cursor, count)
	}
	// this is an empty list
	db.log.Debug("empty erc trx list created")
	return &list, nil
}

// ercTrxListCollectRangeMarks returns a list of ERC20 transactions with proper First/Last marks.
func (db *MongoDbBridge) ercTrxListCollectRangeMarks(col *mongo.Collection, list *types.TokenTransactionList, cursor *string, count int32) (*types.TokenTransactionList, error) {
	var err error

	// find out the cursor ordinal index
	if cursor == nil && count > 0 {
		// get the highest available pk
		list.First, err = db.ercTrxListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: types.FiTokenTransactionOrdinal, Value: -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available pk
		list.First, err = db.ercTrxListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: types.FiTokenTransactionOrdinal, Value: 1}}))
		list.IsEnd = true

	} else if cursor != nil {
		// the cursor itself is the starting point
		list.First, err = db.ercTrxListBorderPk(col,
			bson.D{{Key: types.FiTokenTransactionPk, Value: *cursor}},
			options.FindOne())
	}

	// check the error
	if err != nil {
		db.log.Errorf("can not find the initial ERC20 trx")
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("ERC20 transaction list initialized with ordinal %d", list.First)
	return list, nil
}

// ercTrxListBorderPk finds the top PK of the ERC20 transactions collection based on given filter and options.
func (db *MongoDbBridge) ercTrxListBorderPk(col *mongo.Collection, filter bson.D, opt *options.FindOneOptions) (uint64, error) {
	// prep container
	var row struct {
		Value uint64 `bson:"orx"`
	}

	// make sure we pull only what we need
	opt.SetProjection(bson.D{{Key: types.FiTokenTransactionOrdinal, Value: true}})

	// try to decode
	sr := col.FindOne(context.Background(), filter, opt)
	err := sr.Decode(&row)
	if err != nil {
		return 0, err
	}
	return row.Value, nil
}

// ercTrxListFilter creates a filter for ERC20 transaction list loading.
func (db *MongoDbBridge) ercTrxListFilter(cursor *string, count int32, list *types.TokenTransactionList) *bson.D {
	// build an extended filter for the query; add PK (decoded cursor) to the original filter
	if cursor == nil {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: types.FiTokenTransactionOrdinal, Value: bson.D{{Key: "$lte", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: types.FiTokenTransactionOrdinal, Value: bson.D{{Key: "$gte", Value: list.First}}})
		}
	} else {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: types.FiTokenTransactionOrdinal, Value: bson.D{{Key: "$lt", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: types.FiTokenTransactionOrdinal, Value: bson.D{{Key: "$gt", Value: list.First}}})
		}
	}
	// return the new filter
	return &list.Filter
}

// ercTrxListOptions creates a filter options set for ERC20 transactions list search.
func (db *MongoDbBridge) ercTrxListOptions(count int32) *options.FindOptions {
	// prep options
	opt := options.Find()

	// how to sort results in the collection
	// from high (new) to low (old) by default; reversed if loading from bottom
	sd := -1
	if count < 0 {
		sd = 1
	}

	// sort with the direction we want
	opt.SetSort(bson.D{{Key: types.FiTokenTransactionOrdinal, Value: sd}})

	// prep the loading limit
	var limit = int64(count)
	if limit < 0 {
		limit = -limit
	}

	// apply the limit, try to get one more record so we can detect list end
	opt.SetLimit(limit + 1)
	return opt
}

// ercTrxListLoad load the initialized list of ERC20 transactions from database.
func (db *MongoDbBridge) ercTrxListLoad(col *mongo.Collection, cursor *string, count int32, list *types.TokenTransactionList) (err error) {
	// get the context for loader
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// load the data
	ld, err := col.Find(ctx, db.ercTrxListFilter(cursor, count, list), db.ercTrxListOptions(count))
	if err != nil {
		db.log.Errorf("error loading ERC20 transactions list; %s", err.Error())
		return err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			db.log.Errorf("error closing ERC20 transactions list cursor; %s", err.Error())
		}
	}()

	// loop and load the list; we may not store the last value
	var trx *types.TokenTransaction
	for ld.Next(context.Background()) {
		// append a previous value to the list, if we have one
		if trx != nil {
			list.Collection = append(list.Collection, trx)
		}

		// try to decode the next row
		var row types.TokenTransaction
		if err = ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode the ERC20 transaction list row; %s", err.Error())
			return err
		}

		// use this row as the next item
		trx = &row
	}

	// we should have all the items already; we may just need to check if a boundary was reached
	list.IsEnd = (cursor == nil && count < 0) || (count > 0 && int32(len(list.Collection)) < count)
	list.IsStart = (cursor == nil && count > 0) || (count < 0 && int32(len(list.Collection)) < -count)

	// add the last item as well if we hit the boundary
	if ((count < 0 && list.IsStart) || (count > 0 && list.IsEnd)) && trx != nil {
		list.Collection = append(list.Collection, trx)
	}
	return nil
}

// Erc20Transactions pulls list of ERC20 transactions starting at the specified cursor.
func (db *MongoDbBridge) Erc20Transactions(cursor *string, count int32, filter *bson.D) (*types.TokenTransactionList, error) {
	// nothing to load?
	if count == 0 {
		return nil, fmt.Errorf("nothing to do, zero erc transactions requested")
	}

	// get the collection and context
	col := db.client.Database(db.dbName).Collection(colErcTransactions)

	// init the list
	list, err := db.ercTrxListInit(col, cursor, count, filter)
	if err != nil {
		db.log.Errorf("can not build erc transaction list; %s", err.Error())
		return nil, err
	}

	// load data if there are any
	if list.Total > 0 {
		err = db.ercTrxListLoad(col, cursor, count, list)
		if err != nil {
			db.log.Errorf("can not load erc transaction list from database; %s", err.Error())
			return nil, err
		}

		// reverse on negative so new-er trx will be on top
		if count < 0 {
			list.Reverse()
		}
	}
	return list, nil
}

// Erc20Assets provides list of unique token addresses linked by transactions to the given owner address.
func (db *MongoDbBridge) Erc20Assets(owner common.Address, count int32) ([]common.Address, error) {
	// nothing to load?
	if count <= 1 {
		return nil, fmt.Errorf("nothing to do, zero erc assets requested")
	}

	// get the collection and context
	col := db.client.Database(db.dbName).Collection(colErcTransactions)
	refs, err := col.Distinct(context.Background(), types.FiTokenTransactionToken,
		bson.D{{Key: "to", Value: owner.String()}},
	)
	if err != nil {
		db.log.Errorf("can not pull assets for %s; %s", owner.String(), err.Error())
		return nil, err
	}

	// prep the output array
	res := make([]common.Address, len(refs))
	for i, a := range refs {
		res[i] = common.HexToAddress(a.(string))
	}
	return res, nil
}

// TokenTransactionsByCall provides list of token transactions for the given blockchain transaction call.
func (db *MongoDbBridge) TokenTransactionsByCall(trxHash *common.Hash) ([]*types.TokenTransaction, error) {
	col := db.client.Database(db.dbName).Collection(colErcTransactions)

	// search for values
	ld, err := col.Find(
		context.Background(),
		bson.D{{Key: types.FiTokenTransactionCallHash, Value: trxHash.String()}},
		options.Find().SetSort(bson.D{{Key: types.FiTokenTransactionOrdinal, Value: -1}}),
	)

	defer db.closeCursor(ld)

	// loop and load the list; we may not store the last value
	list := make([]*types.TokenTransaction, 0)
	for ld.Next(context.Background()) {
		var row types.TokenTransaction
		if err = ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode the token transaction; %s", err.Error())
			return nil, err
		}

		// use this row as the next item
		list = append(list, &row)
	}
	return list, nil
}

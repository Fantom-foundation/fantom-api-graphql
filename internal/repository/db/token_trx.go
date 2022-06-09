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
	// colTokenTransactions represents the name of the token (ERC20, ERC721, ERC1155) transaction collection in database.
	colTokenTransactions = "token_trx"

	// filTokenTrxTokenAddress is the name of the token address column in token transactions collection
	filTokenTrxTokenAddress = "addr"

	// filTokenTrxTokenAddress is the name of the token sender address column in token transactions collection
	filTokenTrxSender = "from"

	// filTokenTrxTokenAddress is the name of the token recipient address column in token transactions collection
	filTokenTrxRecipient = "to"

	// filTokenTrxTransaction is the name of the transaction type column in token transactions collection
	filTokenTrxType = "tx_type"

	// filTokenTrxTransaction is the name of the transaction hash column in token transactions collection
	filTokenTrxTransaction = "trx"

	// filTokenTrxOrdinalIndex is the name of the ordinal index column in token transactions collection
	filTokenTrxOrdinalIndex = "ordinal"
)

// tokenTrxCollectionIndexes provides a list of indexes expected to exist on the tokens' transactions collection.
func tokenTrxCollectionIndexes() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 5)

	ixTrxSender := "ix_trx_sender"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "from", Value: 1}}, Options: &options.IndexOptions{Name: &ixTrxSender}}

	ixTrxRecipient := "ix_trx_recipient"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "to", Value: 1}}, Options: &options.IndexOptions{Name: &ixTrxRecipient}}

	ixTrxOrdinal := "ix_trx_ordinal"
	ix[2] = mongo.IndexModel{Keys: bson.D{{Key: "ordinal", Value: -1}}, Options: &options.IndexOptions{Name: &ixTrxOrdinal}}

	ixTrxHash := "ix_trx_hash"
	ix[3] = mongo.IndexModel{Keys: bson.D{{Key: "trx", Value: 1}}, Options: &options.IndexOptions{Name: &ixTrxHash}}

	ixOrdinalRecipient := "ix_ord_to"
	ix[4] = mongo.IndexModel{
		Keys: bson.D{{Key: "to", Value: 1}, {Key: "ordinal", Value: 1}},
		Options: &options.IndexOptions{
			Name: &ixOrdinalRecipient,
		},
	}

	return ix
}

// StoreTokenTransaction stores a token transaction in the database.
func (db *MongoDbBridge) StoreTokenTransaction(trx *types.TokenTransaction) error {
	if trx == nil {
		return fmt.Errorf("can not add empty transaction")
	}

	col := db.client.Database(db.dbName).Collection(colTokenTransactions)
	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: defaultPK, Value: trx.Pk()}},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: filTokenTrxOrdinalIndex, Value: trx.Index()},
			}},
			{Key: "$setOnInsert", Value: trx.WithValue(8)}, // TODO: precision?
		},
		options.Update().SetUpsert(true),
	); err != nil {
		db.log.Errorf("can not add token transaction; %s", err)
		return err
	}

	db.log.Debugf("added token trx at %s, log #%d", trx.Transaction.String(), trx.LogIndex)
	return nil
}

// ErcTransactionCountFiltered calculates total number of ERC20 transactions
// in the database for the given filter.
func (db *MongoDbBridge) ErcTransactionCountFiltered(filter *bson.D) (uint64, error) {
	return db.CountFiltered(db.client.Database(db.dbName).Collection(colTokenTransactions), filter)
}

// ErcTransactionCount calculates total number of ERC20 transactions in the database.
func (db *MongoDbBridge) ErcTransactionCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(colTokenTransactions))
}

// Erc20Transactions pulls list of ERC20 transactions starting at the specified cursor.
func (db *MongoDbBridge) Erc20Transactions(token *common.Address, acc *common.Address, txType []int32, cursor *string, count int32) (*types.TokenTransactionList, error) {
	// nothing to load?
	if count == 0 {
		return nil, fmt.Errorf("nothing to do, zero erc transactions requested")
	}

	// get the collection and context
	col := db.client.Database(db.dbName).Collection(colTokenTransactions)

	// prep the filter
	filter := bson.D{}

	// filter specific token
	if token != nil {
		filter = append(filter, bson.E{
			Key:   filTokenTrxTokenAddress,
			Value: token.String(),
		})
	}

	// common address (sender or recipient)
	if acc != nil {
		filter = append(filter, bson.E{
			Key: "$or",
			Value: bson.A{bson.D{{
				Key:   filTokenTrxSender,
				Value: acc.String(),
			}}, bson.D{{
				Key:   filTokenTrxRecipient,
				Value: acc.String(),
			}}},
		})
	}

	// type of the transaction
	if txType != nil {
		filter = append(filter, bson.E{
			Key: filTokenTrxType,
			Value: bson.D{{
				Key:   "$in",
				Value: txType,
			}},
		})
	}

	// init the list
	list, err := db.ercTrxListInit(col, cursor, count, &filter)
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
	col := db.client.Database(db.dbName).Collection(colTokenTransactions)
	refs, err := col.Distinct(context.Background(), filTokenTrxTokenAddress,
		bson.D{{Key: filTokenTrxRecipient, Value: owner.String()}},
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
	col := db.client.Database(db.dbName).Collection(colTokenTransactions)

	// search for values
	ld, err := col.Find(
		context.Background(),
		bson.D{{Key: filTokenTrxTransaction, Value: trxHash.String()}},
		options.Find().SetSort(bson.D{{Key: filTokenTrxOrdinalIndex, Value: -1}}),
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

// isErcTransactionKnown checks if the given delegation exists in the database.
func (db *MongoDbBridge) isErcTransactionKnown(col *mongo.Collection, trx *types.TokenTransaction) bool {
	// try to find the delegation in the database
	sr := col.FindOne(context.Background(), bson.D{
		{Key: defaultPK, Value: trx.Pk()},
	}, options.FindOne().SetProjection(bson.D{
		{Key: defaultPK, Value: true},
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

// ercTrxListInit initializes list of ERC20 transactions based on provided cursor, count, and filter.
func (db *MongoDbBridge) ercTrxListInit(col *mongo.Collection, cursor *string, count int32, filter *bson.D) (*types.TokenTransactionList, error) {
	// make sure some filter is used
	if nil == filter {
		filter = &bson.D{}
	}

	// find how many transactions do we have in the database
	total, err := col.CountDocuments(context.Background(), *filter)
	if err != nil {
		db.log.Errorf("can not count ERC20 transactions")
		return nil, err
	}

	// make the list and notify the size of it
	db.log.Debugf("found %d filtered ERC20 transactions", total)
	list := types.TokenTransactionList{
		Collection: make([]*types.TokenTransaction, 0),
		Total:      uint64(total),
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
			options.FindOne().SetSort(bson.D{{Key: filTokenTrxOrdinalIndex, Value: -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available pk
		list.First, err = db.ercTrxListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: filTokenTrxOrdinalIndex, Value: 1}}))
		list.IsEnd = true

	} else if cursor != nil {
		// the cursor itself is the starting point
		list.First, err = db.ercTrxListBorderPk(col,
			bson.D{{Key: defaultPK, Value: *cursor}},
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
	opt.SetProjection(bson.D{{Key: filTokenTrxOrdinalIndex, Value: true}})

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
			list.Filter = append(list.Filter, bson.E{Key: filTokenTrxOrdinalIndex, Value: bson.D{{Key: "$lte", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: filTokenTrxOrdinalIndex, Value: bson.D{{Key: "$gte", Value: list.First}}})
		}
	} else {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: filTokenTrxOrdinalIndex, Value: bson.D{{Key: "$lt", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: filTokenTrxOrdinalIndex, Value: bson.D{{Key: "$gt", Value: list.First}}})
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
	opt.SetSort(bson.D{{Key: filTokenTrxOrdinalIndex, Value: sd}})

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
	ctx := context.Background()

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
	for ld.Next(ctx) {
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

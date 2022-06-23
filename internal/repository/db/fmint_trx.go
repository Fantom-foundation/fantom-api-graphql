// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// colFMintTransactions represents the name of the fMint transaction collection in database.
const colFMintTransactions = "fmint_trx"

// fMintTrxCollectionIndexes provides a list of indexes expected to exist on the fmint transactions' collection.
func fMintTrxCollectionIndexes() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 4)

	ixFMintTxToken := "ix_fmint_tx_tok"
	ix[0] = mongo.IndexModel{
		Keys:    bson.D{{Key: types.FiFMintTransactionTokenAddress, Value: 1}},
		Options: &options.IndexOptions{Name: &ixFMintTxToken},
	}

	ixFMintTxUser := "ix_fmint_tx_usr"
	ix[1] = mongo.IndexModel{
		Keys:    bson.D{{Key: types.FiFMintTransactionUser, Value: 1}},
		Options: &options.IndexOptions{Name: &ixFMintTxUser},
	}

	ixFMintTxStamp := "ix_fmint_tx_stamp"
	ix[2] = mongo.IndexModel{
		Keys:    bson.D{{Key: types.FiFMintTransactionTimestamp, Value: -1}},
		Options: &options.IndexOptions{Name: &ixFMintTxStamp},
	}

	ixFMintTxOrdinal := "ix_fmin_tx_orx"
	ix[3] = mongo.IndexModel{
		Keys:    bson.D{{Key: types.FiFMintTransactionOrdinal, Value: -1}},
		Options: &options.IndexOptions{Name: &ixFMintTxOrdinal},
	}

	return ix
}

// AddFMintTransaction stores an fMint transaction in the database if it doesn't exist.
func (db *MongoDbBridge) AddFMintTransaction(trx *types.FMintTransaction) error {
	// get the collection for delegations
	col := db.client.Database(db.dbName).Collection(colFMintTransactions)

	// is it a new one?
	if db.isFMintTransactionKnown(col, trx) {
		return nil
	}

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: defaultPK, Value: trx.ComputedPk()}},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: types.FiFMintTransactionUser, Value: trx.UserAddress},
				{Key: types.FiFMintTransactionTokenAddress, Value: trx.TokenAddress},
				{Key: types.FiFmintTransactionType, Value: trx.Type},
				{Key: types.FiFmintTransactionAmount, Value: trx.Amount},
				{Key: types.FiFmintTransactionFee, Value: trx.Fee},
				{Key: types.FiFmintTransactionFeeValue, Value: trx.ComputedFeeValue()},
				{Key: types.FiFmintTransactionTrxHash, Value: trx.TrxHash},
				{Key: types.FiFmintTransactionTrxIndex, Value: trx.TrxIndex},
				{Key: types.FiFMintTransactionTimestamp, Value: trx.TimeStamp},
				{Key: types.FiFMintTransactionValue, Value: trx.ComputedValue()},
				{Key: types.FiFMintTransactionOrdinal, Value: trx.ComputedOrdinalIndex()},
			}},
			{Key: "$setOnInsert", Value: bson.D{
				{Key: defaultPK, Value: trx.ComputedPk()},
			}},
		},
		options.Update().SetUpsert(true),
	); err != nil {
		db.log.Critical(err)
		return err
	}

	return nil
}

// isFMintTransactionKnown checks if the given delegation exists in the database.
func (db *MongoDbBridge) isFMintTransactionKnown(col *mongo.Collection, trx *types.FMintTransaction) bool {
	// try to find the delegation in the database
	sr := col.FindOne(context.Background(), bson.D{
		{Key: defaultPK, Value: trx.ComputedPk()},
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
		db.log.Errorf("can not get existing fMint transaction pk; %s", sr.Err().Error())
		return false
	}
	return true
}

// FMintTransactionCount calculates total number of fMint transactions in the database.
func (db *MongoDbBridge) FMintTransactionCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(colFMintTransactions))
}

// FMintTransactionCountFiltered calculates total number of sMint transactions
// in the database for the given filter.
func (db *MongoDbBridge) FMintTransactionCountFiltered(filter *bson.D) (uint64, error) {
	return db.CountFiltered(db.client.Database(db.dbName).Collection(colFMintTransactions), filter)
}

// FMintTransactions pulls list of fMint transactions starting at the specified cursor.
func (db *MongoDbBridge) FMintTransactions(cursor *string, count int32, filter *bson.D) (*types.FMintTransactionList, error) {
	// nothing to load?
	if count == 0 {
		return nil, fmt.Errorf("nothing to do, zero fMint transactions requested")
	}

	// get the collection and context
	col := db.client.Database(db.dbName).Collection(colFMintTransactions)

	// init the list
	list, err := db.fMintTrxListInit(col, cursor, count, filter)
	if err != nil {
		db.log.Errorf("can not build fMint transaction list; %s", err.Error())
		return nil, err
	}

	// load data if there are any
	if list.Total > 0 {
		err = db.fMintTrxListLoad(col, cursor, count, list)
		if err != nil {
			db.log.Errorf("can not load fMint transaction list from database; %s", err.Error())
			return nil, err
		}

		// reverse on negative so new-er trx will be on top
		if count < 0 {
			list.Reverse()
		}
	}
	return list, nil
}

// fMintTrxListInit initializes list of fMint transactions based on provided cursor, count, and filter.
func (db *MongoDbBridge) fMintTrxListInit(col *mongo.Collection, cursor *string, count int32, filter *bson.D) (*types.FMintTransactionList, error) {
	// make sure some filter is used
	if nil == filter {
		filter = &bson.D{}
	}

	// find how many transactions do we have in the database
	total, err := col.CountDocuments(context.Background(), *filter)
	if err != nil {
		db.log.Errorf("can not count fMint transactions")
		return nil, err
	}

	// make the list and notify the size of it
	db.log.Debugf("found %d filtered fmint transactions", total)
	list := types.FMintTransactionList{
		Collection: make([]*types.FMintTransaction, 0),
		Total:      uint64(total),
		First:      0,
		Last:       0,
		IsStart:    total == 0,
		IsEnd:      total == 0,
		Filter:     *filter,
	}

	// is the list non-empty? return the list with properly calculated range marks
	if 0 < total {
		return db.fMintTrxListCollectRangeMarks(col, &list, cursor, count)
	}
	// this is an empty list
	db.log.Debug("empty fMint trx list created")
	return &list, nil
}

// fMintTrxListCollectRangeMarks finds range marks of a list of fMint transactions with proper First/Last marks.
func (db *MongoDbBridge) fMintTrxListCollectRangeMarks(col *mongo.Collection, list *types.FMintTransactionList, cursor *string, count int32) (*types.FMintTransactionList, error) {
	var err error

	// find out the cursor ordinal index
	if cursor == nil && count > 0 {
		// get the highest available pk
		list.First, err = db.fMintTrxListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: types.FiFMintTransactionOrdinal, Value: -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available pk
		list.First, err = db.fMintTrxListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: types.FiFMintTransactionOrdinal, Value: 1}}))
		list.IsEnd = true

	} else if cursor != nil {
		// the cursor itself is the starting point
		list.First, err = db.fMintTrxListBorderPk(col,
			bson.D{{Key: defaultPK, Value: *cursor}},
			options.FindOne())
	}

	// check the error
	if err != nil {
		db.log.Errorf("can not find the initial fMint trx")
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("fMint transaction list initialized with ordinal %s", list.First)
	return list, nil
}

// fMintTrxListBorderPk finds the top PK of the ERC20 transactions collection based on given filter and options.
func (db *MongoDbBridge) fMintTrxListBorderPk(col *mongo.Collection, filter bson.D, opt *options.FindOneOptions) (uint64, error) {
	// prep container
	var row struct {
		Value uint64 `bson:"orx"`
	}

	// make sure we pull only what we need
	opt.SetProjection(bson.D{{Key: types.FiFMintTransactionOrdinal, Value: true}})

	// try to decode
	sr := col.FindOne(context.Background(), filter, opt)
	err := sr.Decode(&row)
	if err != nil {
		return 0, err
	}
	return row.Value, nil
}

// fMintTrxListFilter creates a filter for fMint transaction list loading.
func (db *MongoDbBridge) fMintTrxListFilter(cursor *string, count int32, list *types.FMintTransactionList) *bson.D {
	// build an extended filter for the query; add PK (decoded cursor) to the original filter
	if cursor == nil {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: types.FiFMintTransactionOrdinal, Value: bson.D{{Key: "$lte", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: types.FiFMintTransactionOrdinal, Value: bson.D{{Key: "$gte", Value: list.First}}})
		}
	} else {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: types.FiFMintTransactionOrdinal, Value: bson.D{{Key: "$lt", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: types.FiFMintTransactionOrdinal, Value: bson.D{{Key: "$gt", Value: list.First}}})
		}
	}
	// return the new filter
	return &list.Filter
}

// fMintTrxListOptions creates a filter options set for fMint transactions list search.
func (db *MongoDbBridge) fMintTrxListOptions(count int32) *options.FindOptions {
	// prep options
	opt := options.Find()

	// how to sort results in the collection
	// from high (new) to low (old) by default; reversed if loading from bottom
	sd := -1
	if count < 0 {
		sd = 1
	}

	// sort with the direction we want
	opt.SetSort(bson.D{{Key: types.FiFMintTransactionOrdinal, Value: sd}})

	// prep the loading limit
	var limit = int64(count)
	if limit < 0 {
		limit = -limit
	}

	// apply the limit, try to get one more record, so we can detect list end
	opt.SetLimit(limit + 1)
	return opt
}

// fMintTrxListLoad load the initialized list of fMint transactions from database.
func (db *MongoDbBridge) fMintTrxListLoad(col *mongo.Collection, cursor *string, count int32, list *types.FMintTransactionList) (err error) {
	ctx := context.Background()

	// load the data
	ld, err := col.Find(ctx, db.fMintTrxListFilter(cursor, count, list), db.fMintTrxListOptions(count))
	if err != nil {
		db.log.Errorf("error loading fMint transactions list; %s", err.Error())
		return err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			db.log.Errorf("error closing fMint transactions list cursor; %s", err.Error())
		}
	}()

	// loop and load the list; we may not store the last value
	var trx *types.FMintTransaction
	for ld.Next(ctx) {
		// append a previous value to the list, if we have one
		if trx != nil {
			list.Collection = append(list.Collection, trx)
		}

		// try to decode the next row
		var row types.FMintTransaction
		if err = ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode the fMint transaction list row; %s", err.Error())
			return err
		}

		// use this row as the next item
		trx = &row
	}

	// we should have all the items already; we may just need to check if a boundary was reached
	list.IsEnd = (cursor == nil && count < 0) || (count > 0 && int32(len(list.Collection)) < count)
	list.IsStart = (cursor == nil && count > 0) || (count < 0 && int32(len(list.Collection)) < -count)

	// add the last item as well if we hit the boundary
	if (list.IsStart || list.IsEnd) && trx != nil {
		list.Collection = append(list.Collection, trx)
	}
	return nil
}

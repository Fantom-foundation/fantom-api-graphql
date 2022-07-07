// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// colTokenOwnerships is the name of the collection keeping token nft ownership.
	colNftOwnerships = "nft_ownerships"

	// fiNftOwnershipContract is the name of the DB column of the NFT contract address.
	fiNftOwnershipContract = "contract"

	// fiNftOwnershipTokenId is the name of the DB column of the token ID.
	fiNftOwnershipTokenId = "token_id"

	// fiNftOwnershipOwner is the name of the DB column of the token owner address.
	fiNftOwnershipOwner = "owner"

	// fiNftOwnershipAmount is the name of the DB column of the NFT amount.
	fiNftOwnershipAmount = "amount"

	// fiNftOwnershipObtained is the name of the DB column of the obtained date.
	fiNftOwnershipObtained = "obtained"

	// fiNftOwnershipTrx is the name of the DB column of the transaction hash.
	fiNftOwnershipTrx = "trx"

	// fiNftOwnershipTokenName is the name of the DB column of the NFT name.
	fiNftOwnershipTokenName = "token_name"

	// fiNftOwnershipOrdinalIndex is the name of the DB column of the ordinal index.
	fiNftOwnershipOrdinalIndex = "orx"
)

// StoreNftOwnership updates NFT ownership in the persistent storage. If the Amount dropped to zero,
// the record is deleted.
func (db *MongoDbBridge) StoreNftOwnership(no *types.NftOwnership) error {
	if no == nil {
		return fmt.Errorf("no nft ownership value to store")
	}

	// remove record with zero Amount
	if no.Amount.ToInt().Uint64() == 0 {
		return db.deleteNftOwnership(no)
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(colNftOwnerships)

	// try to do the upsert
	id := no.ComputedPk()
	rs, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: defaultPK, Value: id}},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: fiNftOwnershipContract, Value: no.Contract},
				{Key: fiNftOwnershipTokenId, Value: no.TokenId},
				{Key: fiNftOwnershipOwner, Value: no.Owner},
				{Key: fiNftOwnershipAmount, Value: no.Amount},
				{Key: fiNftOwnershipObtained, Value: no.Obtained},
				{Key: fiNftOwnershipTrx, Value: no.Trx},
				{Key: fiNftOwnershipTokenName, Value: no.TokenName},
				{Key: fiNftOwnershipOrdinalIndex, Value: no.ComputedOrdinalIndex()},
			}},
			{Key: "$setOnInsert", Value: bson.D{{Key: defaultPK, Value: id}}},
		},
		options.Update().SetUpsert(true),
	)

	if err != nil {
		db.log.Errorf("can not store nft ownership %s of %s / #%s; %s",
			no.Owner.String(), no.Contract.String(), no.TokenId.String(), err.Error())
		return err
	}

	if rs.UpsertedCount > 0 {
		db.log.Debugf("token %s / #%s nft ownership updated to %s",
			no.Contract.String(), no.TokenId.String(), no.Owner.String())
	}

	return nil
}

// ListNftOwnerships resolves list of nft ownerships based on input data.
func (db *MongoDbBridge) ListNftOwnerships(contract *common.Address, tokenId *hexutil.Big, owner *common.Address, cursor *string, count int32) (out *types.NftOwnershipList, err error) {
	filter := bson.D{}
	if contract != nil {
		filter = append(filter, primitive.E{Key: fiNftOwnershipContract, Value: contract.String()})
	}
	if tokenId != nil {
		filter = append(filter, primitive.E{Key: fiNftOwnershipTokenId, Value: tokenId.String()})
	}
	if owner != nil {
		filter = append(filter, primitive.E{Key: fiNftOwnershipOwner, Value: owner.String()})
	}
	return db.listOwnerships(&filter, cursor, count)
}

// ListOwnerships resolves list of nft ownerships based on input filter.
func (db *MongoDbBridge) listOwnerships(filter *bson.D, cursor *string, count int32) (out *types.NftOwnershipList, err error) {
	// make sure some filter is used
	if filter == nil {
		filter = &bson.D{}
	}

	col := db.client.Database(db.dbName).Collection(colNftOwnerships)

	// init the list
	list, err := db.nftOwnershipListInit(col, cursor, count, filter)
	if err != nil {
		db.log.Errorf("can not build nft ownership list; %s", err.Error())
		return nil, err
	}

	if count == 0 {
		return list, nil // interested in Total only
	}

	if list.Total > 0 {
		err = db.nftOwnershipListLoad(col, cursor, count, list)
		if err != nil {
			db.log.Errorf("can not load transactions list from database; %s", err.Error())
			return nil, err
		}

		// reverse on negative so new-er nft ownerships will be on top
		if count < 0 {
			list.Reverse()
			count = -count
		}

		// cut the end?
		if len(list.Collection) > int(count) {
			list.Collection = list.Collection[:len(list.Collection)-1]
		}
	}

	return list, nil
}

// nftOwnershipListInit initializes list of nft ownerships based on provided cursor and count.
func (db *MongoDbBridge) nftOwnershipListInit(col *mongo.Collection, cursor *string, count int32, filter *bson.D) (*types.NftOwnershipList, error) {
	total, err := col.CountDocuments(context.Background(), *filter)
	if err != nil {
		db.log.Errorf("can not count nft ownerships")
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("found %d nft ownerships in off-chain database", total)
	list := types.NftOwnershipList{
		Collection: make([]*types.NftOwnership, 0),
		Total:      uint64(total),
		First:      0,
		Last:       0,
		IsStart:    total == 0,
		IsEnd:      total == 0,
		Filter:     *filter,
	}

	// is the list non-empty? return the list with properly calculated range marks
	if 0 < total {
		return db.nftOwnershipListCollectRangeMarks(col, &list, cursor, count)
	}
	// this is an empty list
	db.log.Debug("empty nft ownership list created")
	return &list, nil
}

// nftOwnershipListCollectRangeMarks returns a list of nft ownerships with proper First/Last marks.
func (db *MongoDbBridge) nftOwnershipListCollectRangeMarks(col *mongo.Collection, list *types.NftOwnershipList, cursor *string, count int32) (*types.NftOwnershipList, error) {
	var err error

	// find out the cursor ordinal index
	if cursor == nil && count > 0 {
		// get the highest available pk
		list.First, err = db.nftOwnershipListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: fiNftOwnershipOrdinalIndex, Value: -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available pk
		list.First, err = db.nftOwnershipListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: fiNftOwnershipOrdinalIndex, Value: 1}}))
		list.IsEnd = true

	} else if cursor != nil {
		var oid primitive.ObjectID
		oid, err = primitive.ObjectIDFromHex(*cursor)

		if err != nil {
			return nil, fmt.Errorf("invalid cursor value; %s", err.Error())
		}

		// the cursor itself is the starting point
		list.First, err = db.nftOwnershipListBorderPk(col,
			bson.D{{Key: defaultPK, Value: oid}},
			options.FindOne())
	}

	// check the error
	if err != nil {
		db.log.Errorf("can not find the initial nft ownership")
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("nft ownership list initialized with ordinal %d", list.First)
	return list, nil
}

// nftOwnershipListBorderPk finds the top PK of the nft ownership collection based on given filter and options.
func (db *MongoDbBridge) nftOwnershipListBorderPk(col *mongo.Collection, filter bson.D, opt *options.FindOneOptions) (uint64, error) {
	// prep container
	var row struct {
		Value uint64 `bson:"orx"`
	}

	// make sure we pull only what we need
	opt.SetProjection(bson.D{{Key: fiNftOwnershipOrdinalIndex, Value: true}})

	// try to decode
	sr := col.FindOne(context.Background(), filter, opt)

	err := sr.Decode(&row)
	if err != nil {
		return 0, err
	}

	return row.Value, nil
}

// nftOwnershipListLoad load the initialized list of nft ownerships from database.
func (db *MongoDbBridge) nftOwnershipListLoad(col *mongo.Collection, cursor *string, count int32, list *types.NftOwnershipList) (err error) {
	// get the context for loader
	ctx := context.Background()

	// load the data
	ld, err := col.Find(ctx, db.nftOwnershipListFilter(cursor, count, list), db.nftOwnershipListOptions(count))
	if err != nil {
		db.log.Errorf("error loading nft ownership list; %s", err.Error())
		return err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			db.log.Errorf("error closing nft ownership list cursor; %s", err.Error())
		}
	}()

	// loop and load the list; we may not store the last value
	var acc *types.NftOwnership
	for ld.Next(ctx) {
		// append a previous value to the list, if we have one
		if acc != nil {
			list.Collection = append(list.Collection, acc)
		}

		// try to decode the next row
		var row types.NftOwnership
		if err = ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode the nft ownership list row; %s", err.Error())
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

// nftOwnershipListFilter creates a filter for nft ownership list loading.
func (db *MongoDbBridge) nftOwnershipListFilter(cursor *string, count int32, list *types.NftOwnershipList) *bson.D {
	// build an extended filter for the query; add PK (decoded cursor) to the original filter
	if cursor == nil {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: fiNftOwnershipOrdinalIndex, Value: bson.D{{Key: "$lte", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: fiNftOwnershipOrdinalIndex, Value: bson.D{{Key: "$gte", Value: list.First}}})
		}
	} else {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: fiNftOwnershipOrdinalIndex, Value: bson.D{{Key: "$lt", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: fiNftOwnershipOrdinalIndex, Value: bson.D{{Key: "$gt", Value: list.First}}})
		}
	}
	// return the new filter
	return &list.Filter
}

// nftOwnershipListOptions creates a filter options set for contract list search.
func (db *MongoDbBridge) nftOwnershipListOptions(count int32) *options.FindOptions {
	// prep options
	opt := options.Find()

	// how to sort results in the collection
	// from high (new) to low (old) by default; reversed if loading from bottom
	sd := -1
	if count < 0 {
		sd = 1
	}

	// sort with the direction we want
	opt.SetSort(bson.D{{Key: fiNftOwnershipOrdinalIndex, Value: sd}})

	// prep the loading limit
	var limit = int64(count)
	if limit < 0 {
		limit = -limit
	}

	// apply the limit, try to get one more record so we can detect list end
	opt.SetLimit(limit + 1)
	return opt
}

// nftOwnershipCollectionIndexes provides a list of indexes expected to exist on nft ownership records.
func nftOwnershipCollectionIndexes() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 3)

	ixContractToken := "ix_no_contract_token"
	ix[0] = mongo.IndexModel{
		Keys:    bson.D{{Key: fiNftOwnershipContract, Value: 1}, {Key: fiNftOwnershipTokenId, Value: 1}},
		Options: &options.IndexOptions{Name: &ixContractToken},
	}

	ixOwner := "ix_no_owner"
	ix[1] = mongo.IndexModel{
		Keys:    bson.D{{Key: fiNftOwnershipOwner, Value: 1}},
		Options: &options.IndexOptions{Name: &ixOwner},
	}

	ixOrx := "ix_no_orx"
	ix[2] = mongo.IndexModel{
		Keys:    bson.D{{Key: fiNftOwnershipOrdinalIndex, Value: 1}},
		Options: &options.IndexOptions{Name: &ixOrx},
	}

	return ix
}

// deleteNftOwnership removes NFT ownership from the persistent storage.
// We do this if the token qty drops to zero on the owner address.
func (db *MongoDbBridge) deleteNftOwnership(no *types.NftOwnership) error {
	if no == nil {
		return fmt.Errorf("no nft ownership value to delete")
	}

	col := db.client.Database(db.dbName).Collection(colNftOwnerships)

	dr, err := col.DeleteOne(context.Background(), bson.D{{Key: defaultPK, Value: no.ComputedPk()}})
	if err != nil {
		db.log.Errorf("can not delete nft ownership %s of %s / #%s; %s",
			no.Owner.String(), no.Contract.String(), no.TokenId.String(), err.Error())
		return err
	}

	if dr.DeletedCount > 0 {
		db.log.Debugf("token %s / #%s ownership by %s deleted",
			no.Contract.String(), no.TokenId.String(), no.Owner.String())
	}
	return nil
}

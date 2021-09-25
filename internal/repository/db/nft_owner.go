// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
)

const (
	// colNFTOwner represents the name of the NFT owners collection in database.
	colNFTOwner = "nft_owner"
)

// initNFTOwnerCollection initializes the NFT owners collection with
// indexes and additional parameters needed by the app.
func (db *MongoDbBridge) initNFTOwnerCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)
	unique := true

	// index sender and recipient
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiNFTOwnerContract, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiNFTOwnerAddress, Value: 1}}})

	// unique specific token owner
	ix = append(ix, mongo.IndexModel{
		Keys: bson.D{
			{Key: types.FiNFTOwnerContract, Value: 1},
			{Key: types.FiNFTTokenID, Value: 1},
			{Key: types.FiNFTOwnerAddress, Value: 1},
		},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for NFT owners collection; %s", err.Error())
	}

	// log we are done that
	db.log.Debugf("NFT owners collection initialized")
}

// AddNFTOwner adds a new owner to the collection.
// We need to support both ERC-721 and ERC-1155; if the owner already exists, we may
// simply need to update their ownership quantity; if it doesn't, we will insert a new one
// if this is a transfer, we also need to update the original ownership
func (db *MongoDbBridge) AddNFTOwner(nfo *types.NFTOwner) error {
	if nfo == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(colNFTOwner)

	// check for existing record; skip insert if we know it already
	if db.isNFTOwnerKnown(col, nfo) {
		return nil
	}

	// try to do the insert
	if _, err := col.InsertOne(context.Background(), nfo); err != nil {
		db.log.Criticalf("can not add NFT owner %s, contract %s, token %s; %s",
			nfo.OwnerAddress.String(),
			nfo.ContractAddress.String(),
			nfo.TokenId.String(), err.Error())
		return err
	}

	// make sure delegation collection is initialized
	if db.initNFTOwner != nil {
		db.initNFTOwner.Do(func() { db.initNFTOwnerCollection(col); db.initNFTOwner = nil })
	}
	return nil
}

// UpdateNFTOwner changes specific NFT token owner record.
func (db *MongoDbBridge) UpdateNFTOwner(nfo *types.NFTOwner) error {
	if nfo == nil {
		return fmt.Errorf("no value to store")
	}

	// get the collection
	col := db.client.Database(db.dbName).Collection(colNFTOwner)

	// try to execute the update
	up, err := col.UpdateOne(context.Background(), bson.D{
		{Key: types.FiNFTOwnerContract, Value: nfo.ContractAddress.String()},
		{Key: types.FiNFTTokenID, Value: nfo.TokenId.String()},
		{Key: types.FiNFTOwnerAddress, Value: nfo.OwnerAddress.String()},
	}, bson.D{{Key: "$set", Value: bson.D{
		{Key: types.FiNFTTokenOwnerQuantity, Value: nfo.Quantity.String()},
	}}})
	if err != nil {
		db.log.Criticalf("can not update NFT owner; %s", err.Error())
		return err
	}

	// we should have updated the NFT owner; if not, the owner wasn't registered, which is an error!
	if 0 == up.MatchedCount {
		db.log.Criticalf("NFT token %s owner not known on %s", nfo.TokenId.String(), nfo.ContractAddress.String())
		return db.AddNFTOwner(nfo)
	}

	log.Debug("NFT token %s on %s owned by %s updated", nfo.TokenId.String(), nfo.ContractAddress.String(), nfo.OwnerAddress.String())
	return nil
}

// DeleteNFTOwner removes the given NFT owner record from the database.
func (db *MongoDbBridge) DeleteNFTOwner(nft *common.Address, tokenID *big.Int, owner *common.Address) error {
	// get the collection
	col := db.client.Database(db.dbName).Collection(colNFTOwner)

	dl, err := col.DeleteOne(context.Background(), bson.D{
		{Key: types.FiNFTOwnerContract, Value: nft.String()},
		{Key: types.FiNFTTokenID, Value: ((*hexutil.Big)(tokenID)).String()},
		{Key: types.FiNFTOwnerAddress, Value: owner.String()},
	})
	if err != nil {
		db.log.Criticalf("%s NFT token %s of owner %s not deleted; %s",
			nft.String(), ((*hexutil.Big)(tokenID)).String(), owner.String(), err.Error())
		return err
	}

	// any documents deleted?
	if 0 == dl.DeletedCount {
		db.log.Errorf("%s NFT token %s of owner %s not known",
			nft.String(), ((*hexutil.Big)(tokenID)).String(), owner.String())
	}
	return nil
}

// NFTOwnerQty returns the quantity of an NFT registered for the given owner.
func (db *MongoDbBridge) NFTOwnerQty(nft *common.Address, tokenID *big.Int, owner *common.Address) (*big.Int, error) {
	// get the collection
	col := db.client.Database(db.dbName).Collection(colNFTOwner)

	// try to find the owner in the database
	sr := col.FindOne(context.Background(), bson.D{
		{Key: types.FiNFTOwnerContract, Value: nft.String()},
		{Key: types.FiNFTTokenID, Value: ((*hexutil.Big)(tokenID)).String()},
		{Key: types.FiNFTOwnerAddress, Value: owner.String()},
	}, options.FindOne().SetProjection(bson.D{
		{Key: types.FiNFTTokenOwnerQuantity, Value: true},
	}))

	// error on lookup?
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		db.log.Errorf("can not get existing NFT owner pk; %s", sr.Err().Error())
		return nil, sr.Err()
	}

	var row struct {
		Qty string `bson:"qty"`
	}
	if err := sr.Decode(row); err != nil {
		db.log.Errorf("can not get decode NFT owner pk; %s", err.Error())
		return nil, err
	}

	return hexutil.DecodeBig(row.Qty)
}

// isNFTOwnerKnown checks if the given NFT owner exists in the database.
func (db *MongoDbBridge) isNFTOwnerKnown(col *mongo.Collection, nfo *types.NFTOwner) bool {
	// try to find the owner in the database
	sr := col.FindOne(context.Background(), bson.D{
		{Key: types.FiNFTOwnerContract, Value: nfo.ContractAddress.String()},
		{Key: types.FiNFTTokenID, Value: nfo.TokenId.String()},
		{Key: types.FiNFTOwnerAddress, Value: nfo.ContractAddress.String()},
	}, options.FindOne().SetProjection(bson.D{
		{Key: "_id", Value: true},
	}))

	// error on lookup?
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			return false
		}
		db.log.Errorf("can not get existing NFT owner pk; %s", sr.Err().Error())
		return false
	}
	return true
}

// NFTOwnerCountFiltered calculates total number of delegations in the database for the given filter.
func (db *MongoDbBridge) NFTOwnerCountFiltered(filter *bson.D) (uint64, error) {
	return db.CountFiltered(db.client.Database(db.dbName).Collection(colNFTOwner), filter)
}

// NFTOwnerCount calculates estimated total number of NFT owners in the database.
func (db *MongoDbBridge) NFTOwnerCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(colNFTOwner))
}

// nfoListInit initializes list of NFT owners based on provided cursor, count, and filter.
func (db *MongoDbBridge) nfoListInit(col *mongo.Collection, cursor *string, count int32, filter *bson.D) (*types.NFTOwnerList, error) {
	// make sure some filter is used
	if nil == filter {
		filter = &bson.D{}
	}

	// find how many owners do we have in the database
	total, err := col.CountDocuments(context.Background(), *filter)
	if err != nil {
		db.log.Errorf("can not count NFT owners")
		return nil, err
	}

	// make the list and notify the size of it
	db.log.Debugf("found %d filtered NFT owners", total)
	list := types.NFTOwnerList{
		Collection: make([]*types.NFTOwner, 0),
		Total:      uint64(total),
		First:      0,
		Last:       0,
		IsStart:    total == 0,
		IsEnd:      total == 0,
		Filter:     *filter,
	}

	// if the list is not empty, return it with properly calculated range marks
	if 0 < total {
		return db.nfoRangeMarks(col, &list, cursor, count)
	}

	db.log.Debug("empty NFT owners list created")
	return &list, nil
}

// nfoRangeMarks updates the list of NFT owners with proper First/Last marks.
func (db *MongoDbBridge) nfoRangeMarks(col *mongo.Collection, list *types.NFTOwnerList, cursor *string, count int32) (*types.NFTOwnerList, error) {
	var err error

	// find out the cursor ordinal index
	if cursor == nil && count > 0 {
		// get the highest available pk
		list.First, err = db.nfoBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: types.FiNFTOwnerPK, Value: -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available pk
		list.First, err = db.nfoBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: types.FiNFTOwnerPK, Value: 1}}))
		list.IsEnd = true

	} else if cursor != nil {
		id, err := hexutil.DecodeUint64(*cursor)
		if err != nil {
			db.log.Errorf("invalid NFT owners cursor; %s", err.Error())
			return nil, err
		}

		list.First, err = db.nfoBorderPk(col,
			bson.D{{Key: "_id", Value: int64(id)}},
			options.FindOne())
	}

	// check the error
	if err != nil {
		db.log.Errorf("can not find the initial withdraw request")
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("withdraw requests list initialized with PK %s", list.First)
	return list, nil
}

// wrListBorderPk finds the top PK of the NFT owners collection based on given filter and options.
func (db *MongoDbBridge) nfoBorderPk(col *mongo.Collection, filter bson.D, opt *options.FindOneOptions) (uint64, error) {
	// prep container
	var row struct {
		Value uint64 `bson:"_id"`
	}

	// make sure we pull only what we need
	opt.SetProjection(bson.D{{Key: types.FiNFTOwnerPK, Value: true}})
	sr := col.FindOne(context.Background(), filter, opt)

	// try to decode
	if err := sr.Decode(&row); err != nil {
		return 0, err
	}
	return row.Value, nil
}

// nfoListFilter creates a filter for NFT owners list loading.
func (db *MongoDbBridge) nfoListFilter(cursor *string, count int32, list *types.NFTOwnerList) *bson.D {
	if cursor == nil {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: types.FiNFTOwnerPK, Value: bson.D{{Key: "$lte", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: types.FiNFTOwnerPK, Value: bson.D{{Key: "$gte", Value: list.First}}})
		}
	} else {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: types.FiNFTOwnerPK, Value: bson.D{{Key: "$lt", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: types.FiNFTOwnerPK, Value: bson.D{{Key: "$gt", Value: list.First}}})
		}
	}
	return &list.Filter
}

// nfoListOptions creates a filter options set for NFT owners list search.
func (db *MongoDbBridge) nfoListOptions(count int32) *options.FindOptions {
	opt := options.Find()

	// how to sort results in the collection
	// from high (new) to low (old) by default; reversed if loading from bottom
	sd := -1
	if count < 0 {
		sd = 1
	}

	// sort with the direction we want
	opt.SetSort(bson.D{{Key: types.FiNFTOwnerPK, Value: sd}})

	// prep the loading limit
	var limit = int64(count)
	if limit < 0 {
		limit = -limit
	}

	// apply the limit, try to get one more record, so we can detect list end
	opt.SetLimit(limit + 1)
	return opt
}

// nfoListLoad loads the initialized list of NFT owners from database.
func (db *MongoDbBridge) nfoListLoad(col *mongo.Collection, cursor *string, count int32, list *types.NFTOwnerList) (err error) {
	// get the context for loader
	ctx := context.Background()

	// load the data
	ld, err := col.Find(ctx, db.nfoListFilter(cursor, count, list), db.nfoListOptions(count))
	if err != nil {
		db.log.Errorf("error loading with requests list; %s", err.Error())
		return err
	}

	// close the cursor as we leave
	defer func() {
		err = ld.Close(ctx)
		if err != nil {
			db.log.Errorf("error closing NFT owners list cursor; %s", err.Error())
		}
	}()

	// loop and load the list; we may not store the last value
	var nfo *types.NFTOwner
	for ld.Next(ctx) {
		// append a previous value to the list, if we have one
		if nfo != nil {
			list.Collection = append(list.Collection, nfo)
		}

		// try to decode the next row
		var row types.NFTOwner
		if err = ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode the NFT owner list row; %s", err.Error())
			return err
		}

		// use this row as the next item
		nfo = &row
	}

	// we should have all the items already; we may just need to check if a boundary was reached
	list.IsEnd = (cursor == nil && count < 0) || (count > 0 && int32(len(list.Collection)) < count)
	list.IsStart = (cursor == nil && count > 0) || (count < 0 && int32(len(list.Collection)) < -count)

	// add the last item as well if we hit the boundary
	if (list.IsStart || list.IsEnd) && nfo != nil {
		list.Collection = append(list.Collection, nfo)
	}
	return nil
}

// NFTOwners pulls list of NFT owners starting at the specified cursor.
func (db *MongoDbBridge) NFTOwners(cursor *string, count int32, filter *bson.D) (*types.NFTOwnerList, error) {
	// nothing to load?
	if count == 0 {
		return nil, fmt.Errorf("nothing to do, zero NFT owners requested")
	}

	// get the collection and context
	col := db.client.Database(db.dbName).Collection(colNFTOwner)

	// init the list
	list, err := db.nfoListInit(col, cursor, count, filter)
	if err != nil {
		db.log.Errorf("can not build NFT owners list; %s", err.Error())
		return nil, err
	}

	// load data if there are any
	if list.Total > 0 {
		err = db.nfoListLoad(col, cursor, count, list)
		if err != nil {
			db.log.Errorf("can not load NFT owners list from database; %s", err.Error())
			return nil, err
		}

		// reverse on negative so new-er NFT owners will be on top
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

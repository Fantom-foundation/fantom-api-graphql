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
	"time"
)

const (
	// coBallot is the name of the off-chain database collection storing ballot smart contract details.
	coBallot = "ballot"

	// fiBallotPk is the name of the primary key field of the ballot contract collection.
	fiBallotPk = "_id"

	// fiBallotOrdinalIndex is the name of the ballot ordinal index field.
	fiBallotOrdinalIndex = "orx"

	// fiBallotAddress is the name of the address field of the ballot contract.
	fiBallotAddress = "addr"

	// fiBallotName is the name of the ballot contract name field.
	fiBallotName = "name"

	// fiBallotDetailsUrl is the name of the ballot contract details URL field.
	fiBallotDetailsUrl = "url"

	// fiBallotStart is the name of the ballot starting timestamp field.
	fiBallotStart = "start"

	// fiBallotEnd is the name of the ballot ending timestamp field.
	fiBallotEnd = "end"

	// fiBallotProposals is the name of the list of ballot proposals sub-document structure.
	fiBallotProposalsList = "opt"
)

// AddBallot stores a ballot smart contract reference in connected persistent storage.
func (db *MongoDbBridge) AddBallot(ballot *types.Ballot) error {
	// do we have all needed data?
	if ballot == nil {
		return fmt.Errorf("can not add empty ballot contract")
	}

	// get the collection for contracts
	col := db.client.Database(db.dbName).Collection(coBallot)

	// check if the ballot already exists in the database
	exists, err := db.isBallotKnown(col, &ballot.Address)
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// if the ballot already exists, we don't need to do anything here
	if exists {
		return nil
	}

	// try to do the insert
	_, err = col.InsertOne(context.Background(), bson.D{
		{fiBallotPk, ballot.Address.String()},
		{fiBallotOrdinalIndex, ballot.OrdinalIndex},
		{fiBallotAddress, ballot.Address.String()},
		{fiBallotName, ballot.Name},
		{fiBallotDetailsUrl, ballot.DetailsUrl},
		{fiBallotStart, uint64(ballot.Start)},
		{fiBallotEnd, uint64(ballot.End)},
		{fiBallotProposalsList, ballot.Proposals},
	})

	// make sure we are ok
	if err != nil {
		db.log.Critical(err)
		return err
	}

	// inform and quit
	db.log.Debugf("added ballot reference [%s]", ballot.Address.String())
	return nil
}

// isBallotKnown checks if the ballot contract document already exists in the database.
func (db *MongoDbBridge) isBallotKnown(col *mongo.Collection, addr *common.Address) (bool, error) {
	// try to find the ballot contract in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{
		{fiBallotPk, addr.String()},
	}, options.FindOne().SetProjection(bson.D{
		{fiBallotPk, true},
	}))

	// error on lookup?
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			return false, nil
		}

		// inform that we can not get the PK; should not happen
		db.log.Error("can not get existing ballot pk")
		return false, sr.Err()
	}

	return true, nil
}

// Ballot returns details of a ballot smart contract stored in the Mongo database
// if available, or nil if the ballot does not exist.
func (db *MongoDbBridge) Ballot(addr *common.Address) (*types.Ballot, error) {
	// get the collection for transactions
	col := db.client.Database(db.dbName).Collection(coBallot)

	// try to find the contract in the database (it may already exist)
	sr := col.FindOne(context.Background(), bson.D{{fiContractPk, addr.String()}})

	// error on lookup?
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}

		// inform that we can not get the PK; should not happen
		db.log.Errorf("can not get ballot details; %s", sr.Err().Error())
		return nil, sr.Err()
	}

	// prep container
	var row types.Ballot

	// try to decode
	err := sr.Decode(&row)
	if err != nil {
		db.log.Errorf("can not get ballot details; %s", err.Error())
		return nil, err
	}

	return &row, nil
}

// ballotListTotal find the total amount of ballots for the criteria and populates the list
func (db *MongoDbBridge) ballotListTotal(col *mongo.Collection, list *types.BallotList) error {
	// find how many ballots do we have in the database
	total, err := col.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		db.log.Errorf("can not count contracts")
		return err
	}

	// apply the total count
	list.Total = uint64(total)
	return nil
}

// ballotListTopFilter constructs a filter for finding the top item of the list.
// Consider creating DB index db.ballot.createIndex({_id:1,orx:-1},{unique:true}).
func ballotListTopFilter(cursor *string) (*bson.D, error) {
	// no filter by default
	filter := bson.D{}

	// do we have the cursor? we need a filter for that
	if cursor != nil {
		// parse the cursor into an ordinal index
		ix, err := strconv.ParseUint(*cursor, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid ballot cursor value; %s", err.Error())
		}

		// update the filter to reflect the ordinal index
		filter = bson.D{{fiBallotOrdinalIndex, ix}}
	}

	return &filter, nil
}

// ballotListTop find the first ballot of the list based on provided criteria and populates the list.
func (db *MongoDbBridge) ballotListTop(col *mongo.Collection, cursor *string, count int32, list *types.BallotList) error {
	// get the filter
	filter, err := ballotListTopFilter(cursor)
	if err != nil {
		db.log.Errorf("can not find top ballot filter for the list; %s", err.Error())
		return err
	}

	// find out the cursor ordinal index
	if cursor == nil && count > 0 {
		// get the highest available ordinal index (top smart contract)
		list.First, err = db.findBorderOrdinalIndex(col,
			*filter,
			options.FindOne().SetSort(bson.D{{fiBallotOrdinalIndex, -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available ordinal index (bottom smart contract)
		list.First, err = db.findBorderOrdinalIndex(col,
			*filter,
			options.FindOne().SetSort(bson.D{{fiBallotOrdinalIndex, 1}}))
		list.IsEnd = true

	} else if cursor != nil {
		// get the highest available ordinal index (top ballot contract)
		list.First, err = db.findBorderOrdinalIndex(col,
			*filter,
			options.FindOne())
	}

	// check the error
	if err != nil {
		db.log.Errorf("can not find the initial ballot")
		return err
	}

	return nil
}

// ballotListInit initializes list of ballots based on provided cursor and count.
func (db *MongoDbBridge) ballotListInit(col *mongo.Collection, cursor *string, count int32) (*types.BallotList, error) {
	// make the list
	list := types.BallotList{
		Collection: make([]*types.Ballot, 0),
		Total:      0,
		First:      0,
		Last:       0,
		IsStart:    false,
		IsEnd:      false,
	}

	// calculate the total number of contracts in the list
	if err := db.ballotListTotal(col, &list); err != nil {
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("found %d ballots in off-chain database", list.Total)

	// any ballots in the list?
	if list.Total > 0 {
		// find the top contract of the list
		if err := db.ballotListTop(col, cursor, count, &list); err != nil {
			return nil, err
		}

		// inform what we are about to do
		db.log.Debugf("ballots list initialized with ordinal index %d", list.First)
	}

	return &list, nil
}

// ballotListFilter creates a filter for ballot contracts list search.
func (db *MongoDbBridge) ballotListFilter(cursor *string, count int32, list *types.BallotList) *bson.D {
	// inform what we are about to do
	db.log.Debugf("ballots filter starts from index %d", list.First)

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

	// filter all contracts
	filter := bson.D{{fiBallotOrdinalIndex, bson.D{{ordinalOp, list.First}}}}
	return &filter
}

// ballotListOptions creates a filter options set for contract list search.
func (db *MongoDbBridge) ballotListOptions(count int32) *options.FindOptions {
	// prep options
	opt := options.Find()

	// how to sort results in the collection
	if count > 0 {
		// from high (new) to low (old)
		opt.SetSort(bson.D{{fiBallotOrdinalIndex, -1}})
	} else {
		// from low (old) to high (new)
		opt.SetSort(bson.D{{fiBallotOrdinalIndex, 1}})
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

// ballotListLoad loads the initialized list of ballots from persistent database.
func (db *MongoDbBridge) ballotListLoad(col *mongo.Collection, cursor *string, count int32, list *types.BallotList) error {
	// get the context for loader
	ctx := context.Background()

	// load the data
	ld, err := col.Find(ctx, db.ballotListFilter(cursor, count, list), db.ballotListOptions(count))
	if err != nil {
		db.log.Errorf("error loading ballots list; %s", err.Error())
		return err
	}

	// close the cursor as we leave
	defer func() {
		err := ld.Close(ctx)
		if err != nil {
			db.log.Errorf("error closing ballots list cursor; %s", err.Error())
		}
	}()

	// loop and load
	var ballot *types.Ballot
	for ld.Next(ctx) {
		// process the last found hash
		if ballot != nil {
			list.Collection = append(list.Collection, ballot)
			list.Last = ballot.OrdinalIndex
		}

		var row types.Ballot

		// try to decode the next row
		if err := ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode ballot list row; %s", err.Error())
			return err
		}

		// decode the value
		row.Address = common.HexToAddress(row.AddressString)
		ballot = &row
	}

	// if the start is not decided yet; calculate from the results
	if !list.IsStart {
		list.IsStart = count < 0 && int32(len(list.Collection)) < -count
	}

	// if the end is not decided yet; calculate from the results
	if !list.IsEnd {
		list.IsEnd = count > 0 && int32(len(list.Collection)) < count
	}

	// add the last item as well
	if (list.IsStart || list.IsEnd) && ballot != nil {
		list.Collection = append(list.Collection, ballot)
		list.Last = ballot.OrdinalIndex
	}

	return nil
}

// Ballots provides list of ballots stored in the persistent storage.
func (db *MongoDbBridge) Ballots(cursor *string, count int32) (*types.BallotList, error) {
	// nothing to load?
	if count == 0 {
		return nil, fmt.Errorf("nothing to do, zero ballots requested")
	}

	// get the collection and context
	col := db.client.Database(db.dbName).Collection(coBallot)

	// init the list
	list, err := db.ballotListInit(col, cursor, count)
	if err != nil {
		db.log.Errorf("can not build ballots list; %s", err.Error())
		return nil, err
	}

	// any items in the list at all?
	if list.Total > 0 {
		// load data
		err = db.ballotListLoad(col, cursor, count, list)
		if err != nil {
			db.log.Errorf("can not load ballots list data from database; %s", err.Error())
			return nil, err
		}

		// shift the first item on cursor
		if cursor != nil {
			list.First = list.Collection[0].OrdinalIndex
		}

		// reverse on negative so new-er ballots will be on top
		if count < 0 {
			list.Reverse()
		}
	}

	return list, nil
}

// BallotsClosed returns a list of <count> recently closed Ballots.
// We can not decide if the ballot has been already resolved here since we don't
// keep that information in the database, so this has to be resolved later if needed.
// NOTE: Consider creating DB index db.ballot.createIndex({end:-1},{unique:false}).
func (db *MongoDbBridge) BallotsClosed(count uint32) ([]types.Ballot, error) {
	// get the collection and context
	col := db.client.Database(db.dbName).Collection(coBallot)

	// prepare search options and filter
	filter := bson.D{{fiBallotEnd, bson.D{{"$lt", time.Now().UTC().Unix()}}}}
	opt := options.Find().SetSort(bson.D{{fiBallotOrdinalIndex, -1}}).SetLimit(int64(count))

	// load the data
	ld, err := col.Find(context.Background(), filter, opt)
	if err != nil {
		db.log.Errorf("error loading closed ballots list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer func() {
		err := ld.Close(context.Background())
		if err != nil {
			db.log.Errorf("error closing closed ballots list cursor; %s", err.Error())
		}
	}()

	// loop and load
	list := make([]types.Ballot, 0)
	for ld.Next(context.Background()) {
		// make the ballot record
		var row types.Ballot

		// try to decode the next row
		if err := ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode closed ballot list row; %s", err.Error())
			return nil, err
		}

		// decode the value
		row.Address = common.HexToAddress(row.AddressString)
		list = append(list, row)
	}

	return list, nil
}

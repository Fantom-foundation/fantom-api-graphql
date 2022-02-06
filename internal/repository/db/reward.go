// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
)

// colRewards represents the name of the reward claim collection in database.
const colRewards = "rewards"

// initRewardsCollection initializes the reward claims collection with
// indexes and additional parameters needed by the app.
func (db *MongoDbBridge) initRewardsCollection(col *mongo.Collection) {
	// prepare index models
	ix := make([]mongo.IndexModel, 0)

	// index delegator, receiving validator, and creation time stamp
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiRewardClaimAddress, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiRewardClaimToValidator, Value: 1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiRewardClaimOrdinal, Value: -1}}})
	ix = append(ix, mongo.IndexModel{Keys: bson.D{{Key: types.FiRewardClaimedTimeStamp, Value: -1}}})

	// create indexes
	if _, err := col.Indexes().CreateMany(context.Background(), ix); err != nil {
		db.log.Panicf("can not create indexes for reward claims collection; %s", err.Error())
	}

	// log we done that
	db.log.Debugf("reward claims collection initialized")
}

// AddRewardClaim stores a reward claim in the database if it doesn't exist.
func (db *MongoDbBridge) AddRewardClaim(rc *types.RewardClaim) error {
	// get the collection for delegations
	col := db.client.Database(db.dbName).Collection(colRewards)

	// if the delegation already exists, update it with the new one
	if db.isRewardClaimKnown(col, rc) {
		return nil
	}

	// try to do the insert
	if _, err := col.InsertOne(context.Background(), rc); err != nil {
		db.log.Critical(err)
		return err
	}

	// make sure delegation collection is initialized
	if db.initRewards != nil {
		db.initRewards.Do(func() { db.initRewardsCollection(col); db.initRewards = nil })
	}
	return nil
}

// isDelegationKnown checks if the given delegation exists in the database.
func (db *MongoDbBridge) isRewardClaimKnown(col *mongo.Collection, rc *types.RewardClaim) bool {
	// try to find the delegation in the database
	sr := col.FindOne(context.Background(), bson.D{
		{Key: types.FiRewardClaimPk, Value: rc.Pk()},
	}, options.FindOne().SetProjection(bson.D{
		{Key: types.FiRewardClaimPk, Value: true},
	}))

	// error on lookup?
	if sr.Err() != nil {
		// may be ErrNoDocuments, which we seek
		if sr.Err() == mongo.ErrNoDocuments {
			return false
		}
		// inform that we can not get the PK; should not happen
		db.log.Errorf("can not get existing reward claim pk; %s", sr.Err().Error())
		return false
	}

	return true
}

// RewardsCountFiltered calculates total number of reward claims in the database for the given filter.
func (db *MongoDbBridge) RewardsCountFiltered(filter *bson.D) (uint64, error) {
	return db.CountFiltered(db.client.Database(db.dbName).Collection(colRewards), filter)
}

// RewardsCount calculates total number of reward claims in the database.
func (db *MongoDbBridge) RewardsCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(colRewards))
}

// rewListInit initializes list of delegations based on provided cursor, count, and filter.
func (db *MongoDbBridge) rewListInit(col *mongo.Collection, cursor *string, count int32, filter *bson.D) (*types.RewardClaimsList, error) {
	// make sure some filter is used
	if nil == filter {
		filter = &bson.D{}
	}

	// find how many transactions do we have in the database
	total, err := col.CountDocuments(context.Background(), *filter)
	if err != nil {
		db.log.Errorf("can not count reward claims")
		return nil, err
	}

	// make the list and notify the size of it
	db.log.Debugf("found %d filtered reward claims", total)
	list := types.RewardClaimsList{
		Collection: make([]*types.RewardClaim, 0),
		Total:      uint64(total),
		First:      0,
		Last:       0,
		IsStart:    total == 0,
		IsEnd:      total == 0,
		Filter:     *filter,
	}

	// is the list non-empty? return the list with properly calculated range marks
	if 0 < total {
		return db.rewListCollectRangeMarks(col, &list, cursor, count)
	}

	// this is an empty list
	db.log.Debug("empty reward claims list created")
	return &list, nil
}

// rewListCollectRangeMarks returns a list of reward claims with proper First/Last marks.
func (db *MongoDbBridge) rewListCollectRangeMarks(col *mongo.Collection, list *types.RewardClaimsList, cursor *string, count int32) (*types.RewardClaimsList, error) {
	var err error

	// find out the cursor ordinal index
	if cursor == nil && count > 0 {
		// get the highest available pk
		list.First, err = db.rewListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: types.FiRewardClaimOrdinal, Value: -1}}))
		list.IsStart = true

	} else if cursor == nil && count < 0 {
		// get the lowest available pk
		list.First, err = db.rewListBorderPk(col,
			list.Filter,
			options.FindOne().SetSort(bson.D{{Key: types.FiRewardClaimOrdinal, Value: 1}}))
		list.IsEnd = true

	} else if cursor != nil {
		// the cursor itself is the starting point
		list.First, err = db.rewListBorderPk(col,
			bson.D{{Key: types.FiRewardClaimPk, Value: *cursor}},
			options.FindOne())
	}

	// check the error
	if err != nil {
		db.log.Errorf("can not find the initial reward claim")
		return nil, err
	}

	// inform what we are about to do
	db.log.Debugf("reward claim list initialized with ordinal %d", list.First)
	return list, nil
}

// rewListBorderPk finds the top PK of the reward claims collection based on given filter and options.
func (db *MongoDbBridge) rewListBorderPk(col *mongo.Collection, filter bson.D, opt *options.FindOneOptions) (uint64, error) {
	// prep container
	var row struct {
		Value uint64 `bson:"orx"`
	}

	// make sure we pull only what we need
	opt.SetProjection(bson.D{{Key: types.FiRewardClaimOrdinal, Value: true}})

	// try to decode
	sr := col.FindOne(context.Background(), filter, opt)
	err := sr.Decode(&row)
	if err != nil {
		return 0, err
	}
	return row.Value, nil
}

// rewListFilter creates a filter for reward claims list loading.
func (db *MongoDbBridge) rewListFilter(cursor *string, count int32, list *types.RewardClaimsList) *bson.D {
	// build an extended filter for the query; add PK (decoded cursor) to the original filter
	if cursor == nil {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: types.FiRewardClaimOrdinal, Value: bson.D{{Key: "$lte", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: types.FiRewardClaimOrdinal, Value: bson.D{{Key: "$gte", Value: list.First}}})
		}
	} else {
		if count > 0 {
			list.Filter = append(list.Filter, bson.E{Key: types.FiRewardClaimOrdinal, Value: bson.D{{Key: "$lt", Value: list.First}}})
		} else {
			list.Filter = append(list.Filter, bson.E{Key: types.FiRewardClaimOrdinal, Value: bson.D{{Key: "$gt", Value: list.First}}})
		}
	}
	// return the new filter
	return &list.Filter
}

// rewListOptions creates a filter options set for reward claims list search.
func (db *MongoDbBridge) rewListOptions(count int32) *options.FindOptions {
	// prep options
	opt := options.Find()

	// how to sort results in the collection
	// from high (new) to low (old) by default; reversed if loading from bottom
	sd := -1
	if count < 0 {
		sd = 1
	}

	// sort with the direction we want
	opt.SetSort(bson.D{{Key: types.FiRewardClaimOrdinal, Value: sd}})

	// prep the loading limit
	var limit = int64(count)
	if limit < 0 {
		limit = -limit
	}

	// apply the limit, try to get one more record so we can detect list end
	opt.SetLimit(limit + 1)
	return opt
}

// rewListLoad load the initialized list of reward claims from database.
func (db *MongoDbBridge) rewListLoad(col *mongo.Collection, cursor *string, count int32, list *types.RewardClaimsList) (err error) {
	// get the context for loader
	ctx := context.Background()

	// load the data
	ld, err := col.Find(ctx, db.rewListFilter(cursor, count, list), db.rewListOptions(count))
	if err != nil {
		db.log.Errorf("error loading reward claims list; %s", err.Error())
		return err
	}

	// close the cursor as we leave
	defer db.closeCursor(ld)

	// loop and load the list; we may not store the last value
	var rwc *types.RewardClaim
	for ld.Next(ctx) {
		// append a previous value to the list, if we have one
		if rwc != nil {
			list.Collection = append(list.Collection, rwc)
		}

		// try to decode the next row
		var row types.RewardClaim
		if err = ld.Decode(&row); err != nil {
			db.log.Errorf("can not decode the reward claim list row; %s", err.Error())
			return err
		}

		// use this row as the next item
		rwc = &row
	}

	// we should have all the items already; we may just need to check if a boundary was reached
	list.IsEnd = (cursor == nil && count < 0) || (count > 0 && int32(len(list.Collection)) < count)
	list.IsStart = (cursor == nil && count > 0) || (count < 0 && int32(len(list.Collection)) < -count)

	// add the last item as well if we hit the boundary
	if (list.IsStart || list.IsEnd) && rwc != nil {
		list.Collection = append(list.Collection, rwc)
	}
	return nil
}

// RewardClaims pulls list of reward claims starting at the specified cursor.
func (db *MongoDbBridge) RewardClaims(cursor *string, count int32, filter *bson.D) (*types.RewardClaimsList, error) {
	// nothing to load?
	if count == 0 {
		return nil, fmt.Errorf("nothing to do, zero reward claims requested")
	}

	// get the collection and context
	col := db.client.Database(db.dbName).Collection(colRewards)

	// init the list
	list, err := db.rewListInit(col, cursor, count, filter)
	if err != nil {
		db.log.Errorf("can not build reward claims list; %s", err.Error())
		return nil, err
	}

	// load data if there are any
	if list.Total > 0 {
		err = db.rewListLoad(col, cursor, count, list)
		if err != nil {
			db.log.Errorf("can not load reward claims list from database; %s", err.Error())
			return nil, err
		}

		// reverse on negative so new-er delegations will be on top
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

// RewardsSumValue calculates sum of values for all the reward claims by a filter.
func (db *MongoDbBridge) RewardsSumValue(filter *bson.D) (*big.Int, error) {
	return db.sumFieldValue(
		db.client.Database(db.dbName).Collection(colRewards),
		types.FiRewardClaimedValue,
		filter,
		types.RewardDecimalsCorrection)
}

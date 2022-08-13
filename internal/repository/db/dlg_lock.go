// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// colLockedDelegations represents the name of the delegations lock collection
const colLockedDelegations = "locked_dlg"

// lockedDelegationsIndexes provides a list of indexes expected to exist on the locked delegations' collection.
func lockedDelegationsIndexes() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 1)

	unique := true
	ixDelegation := "ix_delegation"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "from", Value: 1}, {Key: "to", Value: 1}}, Options: &options.IndexOptions{
		Name:   &ixDelegation,
		Unique: &unique,
	}}

	return ix
}

// StoreLockedDelegation stores the given locked delegation into the database.
func (db *MongoDbBridge) StoreLockedDelegation(dl *types.LockedDelegation) error {
	col := db.client.Database(db.dbName).Collection(colLockedDelegations)

	// delete any outdated lock first
	_, err := col.DeleteOne(context.Background(), bson.D{
		{Key: "from", Value: dl.Delegator},
		{Key: "to", Value: dl.ValidatorId},
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "created", Value: dl.Locked}},
			bson.D{{Key: "expires", Value: bson.D{{Key: "$lte", Value: dl.Locked}}}},
		}},
	})
	if err != nil {
		db.log.Errorf("could not remove outdated locked delegation; %s", err.Error())
		return err
	}

	// attempt to update
	_, err = col.UpdateOne(
		context.Background(),
		bson.D{
			{Key: "from", Value: dl.Delegator},
			{Key: "to", Value: dl.ValidatorId},
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "locked", Value: dl.Locked},
				{Key: "duration", Value: dl.Duration},
				{Key: "expires", Value: dl.LockedUntil},
			}},
			{Key: "$inc", Value: bson.D{
				{Key: "value", Value: dl.Value},
			}},
			{Key: "$setOnInsert", Value: bson.D{
				{Key: "from", Value: dl.Delegator},
				{Key: "to", Value: dl.ValidatorId},
				{Key: "created", Value: dl.Locked},
			}},
		},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		db.log.Errorf("could not store locked delegation; %s", err.Error())
		return err
	}

	return nil
}

// AdjustLockedDelegation updates value of the given locked delegation by the give amount in the database.
func (db *MongoDbBridge) AdjustLockedDelegation(dlg common.Address, validatorID int64, delta int64) error {
	col := db.client.Database(db.dbName).Collection(colLockedDelegations)

	// attempt to update
	_, err := col.UpdateOne(
		context.Background(),
		bson.D{
			{Key: "from", Value: dlg},
			{Key: "to", Value: validatorID},
		},
		bson.D{
			{Key: "$inc", Value: bson.D{
				{Key: "value", Value: delta},
			}},
		},
		options.Update().SetUpsert(false),
	)
	if err != nil {
		db.log.Errorf("could not store locked delegation; %s", err.Error())
		return err
	}

	return nil
}

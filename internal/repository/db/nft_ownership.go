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

const (
	// colTokenOwnerships is the name of the collection keeping token nft ownership.
	colNftOwnerships = "nft_ownerships"

	// fiOwnershipContract is the name of the DB column of the NFT contract address.
	fiNftOwnershipContract = "contract"

	// fiOwnershipTokenId is the name of the DB column of the token ID.
	fiNftOwnershipTokenId = "token_id"

	// fiOwnershipOwner is the name of the DB column of the token owner address.
	fiNftOwnershipOwner = "owner"

	// fiNftOwnershipAmount is the name of the DB column of the token amount.
	fiNftOwnershipAmount = "amount"

	// fiNftOwnershipObtained is the name of the DB column of the obtained date.
	fiNftOwnershipObtained = "obtained"

	// fiNftOwnershipTrx is the name of the DB column of the transaction hash.
	fiNftOwnershipTrx = "trx"
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
			{Key: "$set", Value: no},
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

// nftOwnershipCollectionIndexes provides a list of indexes expected to exist on nft ownership records.
func nftOwnershipCollectionIndexes() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

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

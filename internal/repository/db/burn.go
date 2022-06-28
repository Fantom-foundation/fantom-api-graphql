// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"bytes"
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
)

// colBurns represents the name of the native FTM burns collection in database.
const colBurns = "burns"

// burnCollectionIndexes provides a list of indexes expected to exist on the native FTM burns' collection.
func burnCollectionIndexes() []mongo.IndexModel {
	// prepare index models
	ix := make([]mongo.IndexModel, 1)

	ixBurnBlock := "ix_burn_block"
	unique := true
	ix[0] = mongo.IndexModel{
		Keys: bson.D{
			{Key: "block", Value: 1},
		},
		Options: &options.IndexOptions{
			Name:   &ixBurnBlock,
			Unique: &unique,
		},
	}

	return ix
}

// BurnByBlock pulls a burn information for the given block number, if available.
func (db *MongoDbBridge) BurnByBlock(bn hexutil.Uint64) (*types.FtmBurn, error) {
	col := db.client.Database(db.dbName).Collection(colBurns)

	// try to find existing burn
	sr := col.FindOne(context.Background(), bson.D{{Key: "block", Value: bn}})
	if sr.Err() != nil {
		// if the burn has not been found, add this as a new one
		if sr.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}

		db.log.Errorf("could not load FTM burn at #%d; %s", bn, sr.Err())
		return nil, sr.Err()
	}

	// decode existing burn and update
	var ex types.FtmBurn
	if err := sr.Decode(&ex); err != nil {
		db.log.Errorf("could not decode FTM burn at #%d; %s", bn, sr.Err())
		return nil, sr.Err()
	}

	return &ex, nil
}

// StoreBurn stores the given native FTM burn record.
func (db *MongoDbBridge) StoreBurn(burn *types.FtmBurn) error {
	if burn == nil {
		return nil
	}

	col := db.client.Database(db.dbName).Collection(colBurns)

	// pull the burn
	ex, err := db.BurnByBlock(burn.BlockNumber)
	if err != nil {
		return err
	}

	// is this a new burn record?
	if ex == nil {
		_, err = col.InsertOne(context.Background(), burn)
		return err
	}

	// check validity for saving
	if !db.isBurnValidForSave(burn, ex) {
		return nil
	}

	// add the new value to the existing one
	val := new(big.Int).Add((*big.Int)(&ex.Amount), (*big.Int)(&burn.Amount))
	ex.Amount = (hexutil.Big)(*val)

	// update the list of included transactions
	if burn.TxList != nil && len(burn.TxList) > 0 {
		if ex.TxList == nil {
			ex.TxList = make([]common.Hash, 0)
		}

		for _, v := range burn.TxList {
			ex.TxList = append(ex.TxList, v)
		}
	}

	// update the record
	_, err = col.UpdateOne(context.Background(), bson.D{{Key: "block", Value: ex.BlockNumber}}, bson.D{{Key: "$set", Value: ex}})
	return err
}

// isBurnValidForSave checks if the new burn should be stored within the database.
func (db *MongoDbBridge) isBurnValidForSave(burn *types.FtmBurn, ex *types.FtmBurn) bool {
	if burn == nil || ex == nil || ex.TxList == nil || burn.TxList == nil {
		db.log.Criticalf("#%d invalid burn check %t; %t", uint64(burn.BlockNumber), burn.TxList, ex.TxList)
		return false
	}

	// all the transactions can already be included
	var found int

	for _, in := range burn.TxList {
		for _, e := range ex.TxList {
			if bytes.Compare(in.Bytes(), e.Bytes()) == 0 {
				found++
				break
			}
		}
	}

	// do we have them all? if so, we have nothing to do here
	if found == len(burn.TxList) {
		return false
	}

	// we can not handle partial update (some transactions are already included, but not all)
	if found > 0 && found < len(burn.TxList) {
		db.log.Criticalf("invalid partial burn received at #%d", burn.BlockNumber)
		return false
	}

	return true
}

// BurnCount estimates the number of burn records in the database.
func (db *MongoDbBridge) BurnCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(colBurns))
}

// BurnTotal aggregates the total amount of burned fee across all blocks.
func (db *MongoDbBridge) BurnTotal() (int64, error) {
	col := db.client.Database(db.dbName).Collection(colBurns)

	// aggregate the total amount of burned native tokens
	cr, err := col.Aggregate(context.Background(), mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "amount", Value: bson.D{{Key: "$sum", Value: "$amount"}}},
		}}},
	})
	if err != nil {
		db.log.Errorf("can not collect total burned fee; %s", err.Error())
		return 0, err
	}

	defer db.closeCursor(cr)
	if !cr.Next(context.Background()) {
		return 0, fmt.Errorf("burned fee aggregation failed")
	}

	var row struct {
		Amount int64 `bson:"amount"`
	}
	if err := cr.Decode(&row); err != nil {
		db.log.Errorf("can not decode burned fee aggregation cursor; %s", err.Error())
		return 0, err
	}
	return row.Amount, nil
}

// BurnList provides list of native FTM burns per blocks stored in the persistent database.
func (db *MongoDbBridge) BurnList(count int64) ([]types.FtmBurn, error) {
	col := db.client.Database(db.dbName).Collection(colBurns)

	cr, err := col.Find(context.Background(), bson.D{}, options.Find().SetSort(bson.D{{Key: "block", Value: -1}}).SetLimit(count))
	if err != nil {
		db.log.Errorf("failed to load burns; %s", err.Error())
		return nil, err
	}
	defer db.closeCursor(cr)

	ctx := context.Background()
	list := make([]types.FtmBurn, 0, count)

	for cr.Next(ctx) {
		var row types.FtmBurn
		if err := cr.Decode(&row); err != nil {
			db.log.Errorf("failed to decode burn; %s", err.Error())
			continue
		}
		list = append(list, row)
	}

	return list, nil
}

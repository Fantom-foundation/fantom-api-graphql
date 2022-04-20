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
	"math/big"
)

// colBurns represents the name of the native FTM burns collection in database.
const colBurns = "burns"

// StoreBurn stores the given native FTM burn record.
func (db *MongoDbBridge) StoreBurn(burn *types.FtmBurn) error {
	if burn == nil {
		return nil
	}

	col := db.client.Database(db.dbName).Collection(colBurns)

	// try to find existing burn
	sr := col.FindOne(context.Background(), bson.D{{Key: "block", Value: burn.Block}})
	if sr.Err() != nil {
		// if the burn has not been found, add this as a new one
		if sr.Err() == mongo.ErrNoDocuments {
			_, err := col.InsertOne(context.Background(), burn)
			return err
		}

		db.log.Errorf("could not load FTM burn at #%d; %s", burn.Block, sr.Err())
		return sr.Err()
	}

	// decode existing burn and update
	var ex types.FtmBurn
	if err := sr.Decode(&sr); err != nil {
		db.log.Errorf("could not decode FTM burn at #%d; %s", burn.Block, sr.Err())
		return sr.Err()
	}

	// all the transactions can already be included
	if ex.TxList != nil && burn.TxList != nil {
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
			return nil
		}

		// we can not handle partial update (some transactions are already included, but not all)
		if found > 0 && found < len(burn.TxList) {
			db.log.Criticalf("invalid partial burn received at #%d", burn.Block)
			return fmt.Errorf("partial burn update rejected at #%d", burn.Block)
		}
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
	_, err := col.UpdateOne(context.Background(), bson.D{{Key: "block", Value: ex.Block}}, bson.D{{Key: "$set", Value: ex}})
	return err
}

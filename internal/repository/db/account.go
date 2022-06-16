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
	// coAccount is the name of the off-chain database collection storing account details.
	colAccounts = "account"

	// fiAccountType represents account type column inside the collection.
	fiAccountType = "act"

	// defaultTokenListLength is the number of ERC20 tokens pulled by default on negative count
	defaultTokenListLength = 25
)

// Account tries to load an account identified by the address given from
// the off-chain database.
func (db *MongoDbBridge) Account(addr *common.Address) (*types.Account, error) {
	col := db.client.Database(db.dbName).Collection(colAccounts)

	// try to find the account
	sr := col.FindOne(context.Background(), bson.D{{Key: defaultPK, Value: addr}}, options.FindOne())
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}

		db.log.Error("can not get existing account %s; %s", addr.String(), sr.Err().Error())
		return nil, sr.Err()
	}

	// try to decode the row
	var row types.Account
	if err := sr.Decode(&row); err != nil {
		db.log.Error("can not decode account %s; %s", addr.String(), err.Error())
		return nil, err
	}

	return &row, nil
}

// StoreAccount stores a unique account record in the database.
func (db *MongoDbBridge) StoreAccount(acc *types.Account) error {
	// do we have account data?
	if acc == nil {
		return fmt.Errorf("can not add empty account")
	}

	// insert the account if not exists
	col := db.client.Database(db.dbName).Collection(colAccounts)
	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{{Key: defaultPK, Value: acc.Address}},
		bson.D{{Key: "$set", Value: acc}},
		options.Update().SetUpsert(true),
	); err != nil {
		db.log.Errorf("can not add account %s; %s", acc.Address.String(), err)
		return err
	}

	db.log.Debugf("added account at %s", acc.Address.String())
	return nil
}

// IsAccountKnown checks if an account document already exists in the database.
func (db *MongoDbBridge) IsAccountKnown(addr *common.Address) (bool, error) {
	col := db.client.Database(db.dbName).Collection(colAccounts)

	sr := col.FindOne(context.Background(), bson.D{{Key: defaultPK, Value: addr}}, options.FindOne().SetProjection(bson.D{{Key: defaultPK, Value: true}}))
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			return false, nil
		}

		db.log.Error("can not get existing account pk")
		return false, sr.Err()
	}
	return true, nil
}

// AccountCount calculates total number of accounts in the database.
func (db *MongoDbBridge) AccountCount() (uint64, error) {
	return db.EstimateCount(db.client.Database(db.dbName).Collection(colAccounts))
}

// AccountTransactions loads list of transaction hashes of an account.
func (db *MongoDbBridge) AccountTransactions(addr *common.Address, rec *common.Address, cursor *string, count int32) (*types.TransactionList, error) {
	if addr == nil {
		return nil, fmt.Errorf("can not list transactions of empty account")
	}

	db.log.Debugf("loading transactions of %s", addr.String())
	if rec == nil {
		// return list of transactions
		filter := bson.D{
			{Key: "$or", Value: bson.A{
				bson.D{{Key: "from", Value: addr.String()}},
				bson.D{{Key: "to", Value: addr.String()}},
			}}}
		return db.Transactions(cursor, count, &filter)
	}

	// return list of transactions filtered by the account and recipient
	filter := bson.D{{Key: "from", Value: addr.String()}, {Key: "to", Value: rec.String()}}
	return db.Transactions(cursor, count, &filter)
}

// ErcTokensList returns a list of known ERC tokens.
func (db *MongoDbBridge) ErcTokensList(count int32, accountType int) ([]common.Address, error) {
	// make sure the count is positive; use default size if not
	if count <= 0 {
		count = defaultTokenListLength
	}

	// log what we do
	db.log.Debugf("loading %d ERC token accounts", count)

	// get the collection for contracts
	col := db.client.Database(db.dbName).Collection(colAccounts)

	// make the filter for ERC tokens only and pull them ordered by activity
	filter := bson.D{{Key: fiAccountType, Value: accountType}}
	opt := options.Find().SetLimit(int64(count))

	// load the data
	cursor, err := col.Find(context.Background(), filter, opt)
	if err != nil {
		db.log.Errorf("error loading ERC tokens list; %s", err.Error())
		return nil, err
	}

	return db.loadAccountAddressList(cursor)
}

// loadAccountAddressList loads the list of addresses from given cursor.
func (db *MongoDbBridge) loadAccountAddressList(cursor *mongo.Cursor) ([]common.Address, error) {
	// close the cursor as we leave
	defer db.closeCursor(cursor)

	// loop and load
	list := make([]common.Address, 0)
	for cursor.Next(context.Background()) {
		var row types.Account
		if err := cursor.Decode(&row); err != nil {
			db.log.Errorf("can not decode ERC contracts list row; %s", err.Error())
			return nil, err
		}

		list = append(list, row.Address)
	}

	return list, nil
}

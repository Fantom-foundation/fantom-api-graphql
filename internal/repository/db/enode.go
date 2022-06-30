// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// colNetworkNodes represents the name of the detected network nodes collection
const colNetworkNodes = "discovery"

// ErrUnknownNetworkNode error is returned on update attempt of an unknown node record.
var ErrUnknownNetworkNode = fmt.Errorf("unknown network node")

// operaNodeCollectionIndexes provides a list of indexes expected to exist on the opera nodes' collection.
func operaNodeCollectionIndexes() []mongo.IndexModel {
	ix := make([]mongo.IndexModel, 2)

	unique := true
	ixNodeAddress := "ix_node_address"
	ix[0] = mongo.IndexModel{Keys: bson.D{{Key: "node", Value: 1}}, Options: &options.IndexOptions{
		Name:   &ixNodeAddress,
		Unique: &unique,
	}}

	ixNodeSeen := "ix_node_seen"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "seen_last", Value: -1}}, Options: &options.IndexOptions{Name: &ixNodeSeen}}

	ixNodeChecked := "ix_node_checked"
	ix[1] = mongo.IndexModel{Keys: bson.D{{Key: "checked", Value: 1}}, Options: &options.IndexOptions{Name: &ixNodeChecked}}

	return ix
}

// NetworkNode gets a stored record of Opera network node by its network identifier.
func (db *MongoDbBridge) NetworkNode(nid enode.ID) (*types.OperaNode, error) {
	col := db.client.Database(db.dbName).Collection(colNetworkNodes)

	// try to find the delegation in the database
	sr := col.FindOne(context.Background(), bson.D{
		{Key: defaultPK, Value: nid},
	})

	// do we have the data?
	if sr.Err() != nil {
		if sr.Err() == mongo.ErrNoDocuments {
			db.log.Errorf("node ID %s not found", nid.String())
			return nil, nil
		}
		return nil, sr.Err()
	}

	// try to decode
	var node types.OperaNode
	if err := sr.Decode(&node); err != nil {
		return nil, err
	}
	return &node, nil
}

// NetworkNodeEvict removes record of a lost Opera network node from the database by its network identifier.
func (db *MongoDbBridge) NetworkNodeEvict(nid enode.ID) error {
	col := db.client.Database(db.dbName).Collection(colNetworkNodes)
	_, err := col.DeleteOne(context.Background(), bson.D{{Key: defaultPK, Value: nid}})
	return err
}

// StoreNetworkNode stores the given Opera node record in the persistent database.
func (db *MongoDbBridge) StoreNetworkNode(node *types.OperaNode) error {
	col := db.client.Database(db.dbName).Collection(colNetworkNodes)

	// update or insert the node into the database
	nu, err := col.UpdateByID(context.Background(), node.Node.ID(), bson.D{
		{Key: "$set", Value: node},
	}, options.Update().SetUpsert(true))

	if err != nil {
		db.log.Errorf("can not store node %s; %s", node.Node.ID(), err.Error())
		return err
	}

	if nu.MatchedCount == 0 {
		db.log.Debugf("new network node %s stored", node.Node.ID())
	}
	return nil
}

// IsNetworkNodeKnown checks if the given network node is already registered in the persistent database.
func (db *MongoDbBridge) IsNetworkNodeKnown(id enode.ID) bool {
	col := db.client.Database(db.dbName).Collection(colNetworkNodes)

	sr := col.FindOne(
		context.Background(),
		bson.D{{Key: defaultPK, Value: id}},
		options.FindOne().SetProjection(bson.D{{Key: defaultPK, Value: true}}),
	)
	if sr.Err() != nil && sr.Err() != mongo.ErrNoDocuments {
		db.log.Errorf("could not check network node %s; %s", id.String(), sr.Err().Error())
	}

	return sr.Err() == nil
}

// NetworkNodeConfirmCheck confirms successful check of the given Opera network node.
func (db *MongoDbBridge) NetworkNodeConfirmCheck(id enode.ID) error {
	col := db.client.Database(db.dbName).Collection(colNetworkNodes)

	now := time.Now().UTC()
	ur, err := col.UpdateByID(context.Background(), id, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "fails", Value: int64(0)},
			{Key: "seen_last", Value: now},
			{Key: "checked", Value: now},
		}},
		{Key: "$inc", Value: bson.D{
			{Key: "score", Value: int64(1)},
		}},
	})

	if err != nil {
		db.log.Errorf("could not confirm node %s; %s", id.String(), err.Error())
		return err
	}

	if ur.ModifiedCount == 0 {
		return ErrUnknownNetworkNode
	}
	return nil
}

// NetworkNodeFailCheck registers failed check of the given Opera network node.
func (db *MongoDbBridge) NetworkNodeFailCheck(id enode.ID) (time.Time, int64, int64, error) {
	col := db.client.Database(db.dbName).Collection(colNetworkNodes)

	// we need a pipeline here to be able to aggregate
	now := time.Now().UTC()
	cu, err := col.Aggregate(context.Background(), mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: defaultPK, Value: id},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "seen_last", Value: true},
			{Key: "score", Value: bson.D{
				{Key: "$floor", Value: bson.D{
					{Key: "$divide", Value: bson.A{"$score", 2}},
				}},
			}},
			{Key: "fails", Value: bson.D{
				{Key: "$add", Value: bson.A{
					"$fails",
					1,
				}},
			}},
		}}},
		{{Key: "$addFields", Value: bson.D{
			{Key: "checked", Value: now},
		}}},
		{{Key: "$merge", Value: bson.D{
			{Key: "into", Value: colNetworkNodes},
			{Key: "on", Value: defaultPK},
			{Key: "whenMatched", Value: "merge"},
			{Key: "whenNotMatched", Value: "discard"},
		}}},
	})

	if err != nil {
		db.log.Errorf("could not register node %s check failure; %s", id.String(), err.Error())
		return time.Time{}, 0, 0, err
	}

	defer db.closeCursor(cu)
	if !cu.Next(context.Background()) {
		return time.Time{}, 0, 0, ErrUnknownNetworkNode
	}

	var row struct {
		LastSeen time.Time `bson:"seen_last"`
		Score    int64     `bson:"score"`
		Fails    int64     `bson:"fails"`
	}
	if err := cu.Decode(&row); err != nil {
		db.log.Errorf("can not decode failure record; %s", err.Error())
		return time.Time{}, 0, 0, nil
	}

	return row.LastSeen, row.Score, row.Fails, nil
}

// NetworkNodeUpdateBatch provides a list of Opera network node addresses most suitable for status update
// based on the registered time of the latest check.
func (db *MongoDbBridge) NetworkNodeUpdateBatch() ([]*enode.Node, error) {
	// calculate the update batch window
	// we try to avoid rapid update attempts, but the updater must finish a previous batch before pulling the new one
	window := time.Now().Add(-30 * time.Minute)

	col := db.client.Database(db.dbName).Collection(colNetworkNodes)
	cu, err := col.Find(
		context.Background(),
		bson.D{{Key: "checked", Value: bson.D{{Key: "$lt", Value: window}}}},
		options.Find().SetSort(bson.D{{Key: "checked", Value: 1}}).SetLimit(100),
	)

	if err != nil {
		db.log.Errorf("invalid network node batch call; %s", err.Error())
		return nil, err
	}

	defer db.closeCursor(cu)
	return db.NetworkNodeAddressList(cu, 100)
}

// NetworkNodeBootstrapSet provides a set of known nodes to be co-used to bootstrap new search.
func (db *MongoDbBridge) NetworkNodeBootstrapSet() []*enode.Node {
	col := db.client.Database(db.dbName).Collection(colNetworkNodes)

	// sample random set of nodes without failed checks, sorted down from the most recently seen
	cu, err := col.Aggregate(context.Background(), mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "fails", Value: 0},
		}}},
		{{Key: "$sort", Value: bson.D{
			{Key: "seen_last", Value: -1},
		}}},
		{{Key: "$sample", Value: bson.D{
			{Key: "size", Value: 25},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "node", Value: true},
		}}},
	})

	defer db.closeCursor(cu)

	// load the list of addresses; the slice is returned even if an error happened
	l, err := db.NetworkNodeAddressList(cu, 25)
	if err != nil {
		db.log.Errorf("could not load bootstrap batch; %s", err.Error())
	}
	return l
}

// NetworkNodeAddressList provides a list of Opera network node addresses list for the given filter, sorting,
// and expected batch size.
func (db *MongoDbBridge) NetworkNodeAddressList(cur *mongo.Cursor, size uint) ([]*enode.Node, error) {
	list := make([]*enode.Node, 0, size)

	for cur.Next(context.Background()) {
		var n struct {
			Node enode.Node `bson:"node"`
		}
		if err := cur.Decode(&n); err != nil {
			db.log.Errorf("could not decode network node; %s", err.Error())
			continue
		}
		list = append(list, &n.Node)
	}

	return list, nil
}

// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/p2p/enode"
	"time"
)

// OperaNode represents a node on Fantom Opera network.
type OperaNode struct {
	// Node address in native form
	Node enode.Node `bson:"node"`

	// The score tracks how many liveliness checks were performed. It is incremented by one
	// every time the node passes a check, and halved every time it doesn't.
	Score int64 `bson:"score"`

	// Fails tracks the number of failed checks performed on the node.
	Fails int64 `bson:"fails"`

	// These two track the time of last successful contact.
	Found    time.Time `bson:"seen_first"`
	LastSeen time.Time `bson:"seen_last"`

	// This one tracks the time of our last attempt to contact the node.
	LastCheck time.Time `bson:"checked"`

	// Server geolocation
	LocationContinent string  `bson:"geo_con"`
	LocationCountry   string  `bson:"geo_cry"`
	LocationCity      string  `bson:"geo_city"`
	LocationRegion    string  `bson:"geo_sub"`
	LocationTimeZone  string  `bson:"geo_tz"`
	LocationLatitude  float64 `bson:"geo_lt"`
	LocationLongitude float64 `bson:"geo_lg"`
	LocationAccuracy  int16   `bson:"geo_acu"`
}

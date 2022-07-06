// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/p2p/enode"
	"time"
)

// GeoLocation represent geographic location.
type GeoLocation struct {
	Continent string  `bson:"continent"`
	Country   string  `bson:"country"`
	Region    string  `bson:"region"`
	City      string  `bson:"city"`
	TimeZone  string  `bson:"tz"`
	Latitude  float64 `bson:"lat"`
	Longitude float64 `bson:"lon"`
	Accuracy  uint16  `bson:"accuracy"`
}

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

	// Geographic location information
	Location GeoLocation `bson:"location"`
}

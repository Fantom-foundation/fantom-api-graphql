// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/p2p/enode"
	"time"
)

const (
	OperaNodeGeoAggregationLevelContinent = iota
	OperaNodeGeoAggregationLevelCountry
	OperaNodeGeoAggregationLevelState
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

// OperaNodeLocationAggregate represents an aggregated summary of Opera network nodes
// based on their geographic location.
type OperaNodeLocationAggregate struct {
	// Region represents the name of the location of the aggregation group
	// based on selected detail level.
	Region string `bson:"_id"`

	// TopRegion represents the name of the top level location of the aggregation group.
	TopRegion string `bson:"top_region" json:"topRegion"`

	// Latitude represents average geographic coordinates of the aggregate.
	Latitude  float64 `bson:"lat" json:"latitude"`
	Longitude float64 `bson:"lon" json:"longitude"`

	// Count represents the number of nodes in the aggregation group.
	Count int32 `bson:"total"`
}

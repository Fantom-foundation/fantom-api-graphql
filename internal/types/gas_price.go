// Package types implements different core types of the API.
package types

import (
	"time"
)

const (
	// GasPricePeriodTypeSuggestion represents the type of gas price period data from the Opera node suggestion call.
	GasPricePeriodTypeSuggestion = iota
)

// GasPrice represents an extended gas price estimator.
type GasPrice struct {
	Fast    float64 `json:"fast"`
	Fastest float64 `json:"fastest"`
	SafeLow float64 `json:"safeLow"`
	Average float64 `json:"average"`
}

// GasPricePeriod represents a data set of interval of gas price
// estimation provided by the Opera node.
type GasPricePeriod struct {
	Type  int8      `json:"type" bson:"type"`
	Open  int64     `json:"open" bson:"open"`
	Close int64     `json:"close" bson:"close"`
	Min   int64     `json:"min" bson:"min"`
	Max   int64     `json:"max" bson:"max"`
	Avg   int64     `json:"avg" bson:"avg"`
	From  time.Time `json:"from" bson:"from"`
	To    time.Time `json:"to" bson:"to"`
	Tick  int64     `json:"tick" bson:"tick"`
}

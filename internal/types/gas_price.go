// Package types implements different core types of the API.
package types

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	// GasPricePeriodTypeSuggestion represents the type of gas price period data from the Opera node suggestion call.
	GasPricePeriodTypeSuggestion = iota
)

const (
	// FiGasPriceTimeFrom is the name of the starting time stamp column in the collection.
	FiGasPriceTimeFrom = "from"

	// FiGasPriceTimeTo is the name of the ending time stamp column in the collection.
	FiGasPriceTimeTo = "to"

	/*
		// FiGasPriceAmountOpen is the name of the adjusted opening gas price amount column in the collection.
		FiGasPriceAmountOpen = "open"

		// FiGasPriceAmountMax is the name of the adjusted maximum gas price amount column in the collection.
		FiGasPriceAmountMax = "max"

		// FiGasPriceAmountMin is the name of the adjusted minimum gas price amount column in the collection.
		FiGasPriceAmountMin = "min"

		// FiGasPriceAmountAvg is the name of the adjusted average gas price amount column in the collection.
		FiGasPriceAmountAvg = "avg"

		// FiGasPriceAmountClose is the name of the adjusted closing gas price amount column in the collection.
		FiGasPriceAmountClose = "close"

		// FiGasPriceTimeTick is the name of the tick speed column in the collection.
		FiGasPriceTimeTick = "tick"
	*/
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

// MarshalBSON creates a BSON representation of the gas price estimation record.
func (gpp *GasPricePeriod) MarshalBSON() ([]byte, error) {
	return bson.Marshal(*gpp)
}

// UnmarshalBSON updates the value from BSON source.
func (gpp *GasPricePeriod) UnmarshalBSON(data []byte) (err error) {
	return bson.Unmarshal(data, gpp)
}

// Package types implements different core types of the API.
package types

// GasPrice represents an extended gas price estimator.
type GasPrice struct {
	Fast    int32 `json:"fast"`
	Fastest int32 `json:"fastest"`
	SafeLow int32 `json:"safeLow"`
	Average int32 `json:"average"`
}

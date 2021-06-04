// Package types implements different core types of the API.
package types

// GasPrice represents an extended gas price estimator.
type GasPrice struct {
	Fast    float64 `json:"fast"`
	Fastest float64 `json:"fastest"`
	SafeLow float64 `json:"safeLow"`
	Average float64 `json:"average"`
}

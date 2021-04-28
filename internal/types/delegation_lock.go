// Package types implements different core types of the API.
package types

import "github.com/ethereum/go-ethereum/common/hexutil"

// DelegationLock represents a lock related to a delegation
type DelegationLock struct {
	LockedAmount    hexutil.Big    `json:"lockedStake"`
	LockedFromEpoch hexutil.Uint64 `json:"fromEpoch"`
	LockedUntil     hexutil.Uint64 `json:"endTime"`
	Duration        hexutil.Uint64 `json:"duration"`
}

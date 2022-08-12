// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// DelegationLock represents a lock related to a delegation
type DelegationLock struct {
	LockedAmount    hexutil.Big    `json:"lockedStake"`
	LockedFromEpoch hexutil.Uint64 `json:"fromEpoch"`
	LockedUntil     hexutil.Uint64 `json:"endTime"`
	Duration        hexutil.Uint64 `json:"duration"`
}

// DelegationUnlock represents an unlock request on a locked delegation.
type DelegationUnlock struct {
	When    time.Time   `bson:"when"`
	Amount  hexutil.Big `bson:"amount"`
	Penalty hexutil.Big `bson:"penalty"`
	Value   int64       `bson:"value"`
	Fine    int64       `bson:"fine"`
}

// LockedDelegation represents an information about locked delegation.
type LockedDelegation struct {
	Delegator   common.Address `bson:"from"`
	ValidatorId int64          `bson:"to"`
	Locked      time.Time      `bson:"locked"`
	Duration    int64          `bson:"duration"`
	LockedUntil time.Time      `bson:"expires"`
	Amount      hexutil.Big    `bson:"amount,omitempty"`
	Value       int64          `bson:"value"`
}

// LockedDelegationValue converts the given delegation lock amount to the value as stored.
func LockedDelegationValue(amo *big.Int) int64 {
	if amo == nil {
		return 0
	}
	return new(big.Int).Div(amo, DelegationDecimalsCorrection).Int64()
}

// SetAmount updates the amount locked and also the value of the lock.
func (ld *LockedDelegation) SetAmount(amo *hexutil.Big) {
	if amo != nil {
		ld.Amount = *amo
	}
	ld.Value = LockedDelegationValue(amo.ToInt())
}

// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Delegation represents a delegator in Opera blockchain.
type Delegation struct {
	Address          common.Address  `json:"address"`
	ToStakerId       hexutil.Uint64  `json:"toStakerID"`
	CreatedEpoch     hexutil.Uint64  `json:"createdEpoch"`
	CreatedTime      hexutil.Uint64  `json:"createdTime"`
	DeactivatedEpoch *hexutil.Uint64 `json:"deactivatedEpoch"`
	DeactivatedTime  *hexutil.Uint64 `json:"deactivatedTime"`
	AmountDelegated  *hexutil.Big    `json:"amount"`
	ClaimedReward    *hexutil.Big    `json:"claimedRewards"`
}

// DelegationLock represents a lock related to a delegation
type DelegationLock struct {
	LockedFromEpoch  hexutil.Uint64  `json:"lockedFrom"`
	LockedUntil      hexutil.Uint64  `json:"lockedUntil"`
}
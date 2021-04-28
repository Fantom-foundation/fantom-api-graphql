// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// PendingRewards represents a rewards waiting to be claimed structure.
type PendingRewards struct {
	Address common.Address
	Staker  hexutil.Big
	Amount  hexutil.Big
}

// FromEpoch returns the first unpaid epoch.
// This is deprecated in SFCv3.
func (pr PendingRewards) FromEpoch() hexutil.Uint64 {
	return 0
}

// ToEpoch returns the last unpaid epoch.
// This is deprecated in SFCv3.
func (pr PendingRewards) ToEpoch() hexutil.Uint64 {
	return 0
}

// IsOverRange returns if the pending reward is over safe range.
// This is deprecated in SFCv3.
func (pr PendingRewards) IsOverRange() bool {
	return false
}

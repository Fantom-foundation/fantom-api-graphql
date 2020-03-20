// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Delegator represents a delegator in Opera blockchain.
type Delegator struct {
	Address          common.Address  `json:"address"`
	ToStakerId       hexutil.Uint64  `json:"toStakerID"`
	CreatedEpoch     hexutil.Uint64  `json:"createdEpoch"`
	CreatedTime      hexutil.Uint64  `json:"createdTime"`
	DeactivatedEpoch *hexutil.Uint64 `json:"deactivatedEpoch"`
	DeactivatedTime  *hexutil.Uint64 `json:"deactivatedTime"`
	Amount           *hexutil.Big    `json:"amount"`
	ClaimedReward    *hexutil.Big    `json:"claimedRewards"`
}

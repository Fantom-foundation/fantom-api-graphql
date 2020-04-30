// Package types implements different core types of the API.
package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Epoch represents epoch detail.
type Epoch struct {
	Id                     hexutil.Uint64
	EndTime                hexutil.Big
	Duration               hexutil.Big
	EpochFee               hexutil.Big
	TotalBaseRewardWeight  hexutil.Big
	TotalTxRewardWeight    hexutil.Big
	BaseRewardPerSecond    hexutil.Big
	StakeTotalAmount       hexutil.Big
	DelegationsTotalAmount hexutil.Big
	TotalSupply            hexutil.Big
}

// UnmarshalEpoch parses the JSON-encoded Epoch data.
func UnmarshalEpoch(data []byte) (*Epoch, error) {
	var ep Epoch
	err := json.Unmarshal(data, &ep)
	return &ep, err
}

// Marshal returns the JSON encoding of Epoch.
func (e *Epoch) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

// Package types implements different core types of the API.
package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Epoch represents epoch detail.
type Epoch struct {
	Id                    hexutil.Uint64 `json:"id"`
	EndTime               hexutil.Uint64 `json:"end"`
	EpochFee              hexutil.Big    `json:"fee"`
	TotalBaseRewardWeight hexutil.Big    `json:"trw"`
	TotalTxRewardWeight   hexutil.Big    `json:"txw"`
	BaseRewardPerSecond   hexutil.Big    `json:"brw"`
	StakeTotalAmount      hexutil.Big    `json:"stk"`
	TotalSupply           hexutil.Big    `json:"sup"`
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

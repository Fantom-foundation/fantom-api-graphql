// Package types implements different core types of the API.
package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Epoch represents epoch detail.
type Epoch struct {
	Id                    hexutil.Uint64 `json:"id"  bson:"_id"`
	EndTime               uint64         `json:"end" bson:"end"`
	EpochFee              hexutil.Big    `json:"fee" bson:"fee"`
	TotalBaseRewardWeight hexutil.Big    `json:"trw" bson:"brw"`
	TotalTxRewardWeight   hexutil.Big    `json:"txw" bson:"trw"`
	BaseRewardPerSecond   hexutil.Big    `json:"brw" bson:"rew"`
	StakeTotalAmount      hexutil.Big    `json:"stk" bson:"stake"`
	TotalSupply           hexutil.Big    `json:"sup" bson:"supply"`
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

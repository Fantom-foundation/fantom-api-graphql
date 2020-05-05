// Package types implements different core types of the API.
package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
)

// Account represents an Opera account at the blockchain.
type Account struct {
	Address  common.Address `json:"address"`
	Contract *Hash          `json:"contract"`
}

// UnmarshalAccount parses the JSON-encoded account data.
func UnmarshalAccount(data []byte) (*Account, error) {
	var acc Account
	err := json.Unmarshal(data, &acc)
	return &acc, err
}

// Marshal returns the JSON encoding of account.
func (acc *Account) Marshal() ([]byte, error) {
	return json.Marshal(acc)
}

// Package types implements different core types of the API.
package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
)

const (
	// AccountTypeWallet identifies accounts of the type regular wallet
	AccountTypeWallet = "wallet"

	// AccountTypeContract identifies a generic contract type
	AccountTypeContract = "contract"

	// AccountTypeERC20Token identifies a contract of type ERC20 token
	AccountTypeERC20Token = "ERC20"
)

// Account represents an Opera account at the blockchain.
type Account struct {
	Address    common.Address `json:"address"`
	ContractTx *Hash          `json:"contract"`
	Type       string         `json:"type"`
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

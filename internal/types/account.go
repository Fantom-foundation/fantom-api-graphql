// Package types implements different core types of the API.
package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	// AccountTypeWallet identifies accounts of the type regular wallet
	AccountTypeWallet = "wallet"

	// AccountTypeContract identifies a generic contract type
	AccountTypeContract = "contract"

	// AccountTypeSFC identifies the Special Function Contract
	AccountTypeSFC = "SFC"

	// AccountTypeERC20Token identifies a contract of type ERC20 token
	AccountTypeERC20Token = "ERC20"

	// AccountTypeERC721Contract identifies a contract of type ERC721 non-fungible token
	AccountTypeERC721Contract = "ERC721"

	// AccountTypeERC1155Contract identifies a multi-token contract of type ERC1155
	AccountTypeERC1155Contract = "ERC1155"
)

// Account represents an Opera account at the blockchain.
type Account struct {
	Address      common.Address `json:"address"`
	ContractTx   *common.Hash   `json:"contract"`
	Type         string         `json:"type"`
	LastActivity hexutil.Uint64 `json:"ats"`
	TrxCounter   hexutil.Uint64 `json:"trc"`
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

// Package types implements different core types of the API.
package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
)

// Erc20Token represents an ERC20 token contract
// on the blockchain.
type Erc20Token struct {
	// Address represents the address of the ERC20 contract on chain.
	Address common.Address `json:"address"`

	// Name represents an extended name of the token.
	Name string `json:"name"`

	// Symbol represents an abbreviation for the token.
	Symbol string `json:"symbol"`

	// Decimals is the number of decimals the token supports.
	// The most common value is 18 to mimic the ETH to WEI relationship.
	// USD pairs on ChainLink (we use for price oracles) use 8 digits.
	Decimals int32 `json:"decimals"`
}

// UnmarshalErc20Token parses the JSON-encoded account data.
func UnmarshalErc20Token(data []byte) (*Erc20Token, error) {
	var erc20 Erc20Token
	err := json.Unmarshal(data, &erc20)
	return &erc20, err
}

// Marshal returns the JSON encoding of account.
func (erc20 *Erc20Token) Marshal() ([]byte, error) {
	return json.Marshal(erc20)
}

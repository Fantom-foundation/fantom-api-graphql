// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// Erc721Token represents an ERC721 token contract
// on the blockchain.
type Erc721Token struct {
	// Address represents the address of the ERC20 contract on chain.
	Address common.Address `json:"address"`

	// Name represents an extended name of the token.
	Name string `json:"name"`

	// Symbol represents an abbreviation for the token.
	Symbol string `json:"symbol"`
}

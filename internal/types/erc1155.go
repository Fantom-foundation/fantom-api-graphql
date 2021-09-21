// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// Erc1155Contract represents an ERC1155 token contract
// on the blockchain.
type Erc1155Contract struct {
	// Address represents the address of the ERC20 contract on chain.
	Address common.Address `json:"address"`
}

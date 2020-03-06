// Package types implements different core types of the API.
package types

import (
	"math/big"
)

// Block represents a basic information provided by the API about block inside Opera blockchain.
type Block struct {
	Hash      Hash
	Number    big.Int
	TimeStamp uint64
	Txs       []*Hash
}

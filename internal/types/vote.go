// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Vote represents a vote of a given Voter in a given Ballot.
type Vote struct {
	// Voter represents an address participating in a Ballot.
	Voter common.Address

	// Ballot represents the address of the voting contract.
	Ballot common.Address

	// Vote represents the selected proposal the voter chose.
	// The vote is nil if the voter didn't participate in the ballot.
	Vote *hexutil.Uint64
}

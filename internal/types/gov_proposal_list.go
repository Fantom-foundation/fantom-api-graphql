// Package types implements different core types of the API.
package types

import "github.com/ethereum/go-ethereum/common"

// GovernanceProposalList represents a list of governance proposals list.
type GovernanceProposalList struct {
	// List keeps the actual Collection.
	Collection []*GovernanceProposal

	// Total indicates total number of contracts in the whole collection.
	Total uint64

	// First is the index of the first item on the list
	First common.Address

	// Last is the index of the last item on the list
	Last common.Address

	// IsStart indicates there are no contracts available above the list currently.
	IsStart bool

	// IsEnd indicates there are no contracts available below the list currently.
	IsEnd bool
}

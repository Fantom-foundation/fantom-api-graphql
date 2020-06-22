// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Ballot represents an Opera smart contract at the blockchain.
type Ballot struct {
	// OrdinalIndex is the ordinal ballot contract index in the database.
	OrdinalIndex uint64 `json:"orx" bson:"orx"`

	// Address represents the address of the ballot contract
	AddressString string         `json:"addr" bson:"addr"`
	Address       common.Address `json:"-" bson:"-"`

	// Name of the ballot smart contract, if available.
	Name string `json:"name" bson:"name"`

	// URL of the ballot detailed information page.
	DetailsUrl string `json:"url" bson:"url"`

	// An approximate timestamp after which the ballot opens for voting.
	Start hexutil.Uint64 `json:"start" bson:"start"`

	// An approximate timestamp after which the ballot
	// is closed and no longer accepts votes.
	End hexutil.Uint64 `json:"end" bson:"end"`

	// List of proposals of the ballot.
	Proposals []string `json:"opt" bson:"opt"`
}

// BallotList represents a list of ballot smart contracts.
type BallotList struct {
	// List keeps the actual collection of ballots.
	Collection []*Ballot

	// Total indicates total number of transaction in the whole collection.
	Total uint64

	// First is the index of the first item on the list
	First uint64

	// Last is the index of the last item on the list
	Last uint64

	// IsStart indicates there are no contracts available above the list currently.
	IsStart bool

	// IsEnd indicates there are no contracts available below the list currently.
	IsEnd bool
}

// Reverse reverses the order of ballots in the list.
func (c *BallotList) Reverse() {
	// anything to swap at all?
	if c.Collection == nil || len(c.Collection) < 2 {
		return
	}

	// swap elements
	for i, j := 0, len(c.Collection)-1; i < j; i, j = i+1, j-1 {
		c.Collection[i], c.Collection[j] = c.Collection[j], c.Collection[i]
	}

	// swap indexes
	c.First, c.Last = c.Last, c.First
}

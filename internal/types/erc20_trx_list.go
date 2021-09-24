// Package types implements different core types of the API.
package types

import "go.mongodb.org/mongo-driver/bson"

// TokenTransactionList represents a list of ERC20/ERC721/ERC1155 transactions.
type TokenTransactionList struct {
	// List keeps the actual Collection.
	Collection []*TokenTransaction

	// Total indicates total number of ERC transactions in the whole collection.
	Total uint64

	// First is the index of the first item on the list
	First uint64

	// Last is the index of the last item on the list
	Last uint64

	// IsStart indicates there are no ERC transactions available above the list currently.
	IsStart bool

	// IsEnd indicates there are no ERC transactions available below the list currently.
	IsEnd bool

	// Filter represents the base filter used for filtering the list
	Filter bson.D
}

// Reverse reverses the order of ERC transactions in the list.
func (c *TokenTransactionList) Reverse() {
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

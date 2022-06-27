// Package types implements different core types of the API.
package types

import "go.mongodb.org/mongo-driver/bson"

// NftOwnershipList represents a list of nft ownership.
type NftOwnershipList struct {
	// List keeps the actual Collection.
	Collection []*NftOwnership

	// Total indicates total number of delegations in the whole collection.
	Total uint64

	// First is the index of the first item on the list
	First string

	// Last is the index of the last item on the list
	Last string

	// IsStart indicates there are no delegations available above the list currently.
	IsStart bool

	// IsEnd indicates there are no delegations available below the list currently.
	IsEnd bool

	// Filter represents the base filter used for filtering the list
	Filter bson.D
}

// Reverse reverses the order of delegations in the list.
func (c *NftOwnershipList) Reverse() {
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

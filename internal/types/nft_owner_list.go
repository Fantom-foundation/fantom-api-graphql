// Package types implements different core types of the API.
package types

import "go.mongodb.org/mongo-driver/bson"

// NFTOwnerList represents a list of NFT owners.
type NFTOwnerList struct {
	// List keeps the actual Collection.
	Collection []*NFTOwner

	// Total indicates total number of NFT owners in the whole collection.
	Total uint64

	// First is the index of the first item on the list
	First uint64

	// Last is the index of the last item on the list
	Last uint64

	// IsStart indicates there are no rows above the list currently.
	IsStart bool

	// IsEnd indicates there are no rows below the list currently.
	IsEnd bool

	// Filter represents the base filter used for filtering the list
	Filter bson.D
}

// Reverse reverses the order of fMint transactions in the list.
func (l *NFTOwnerList) Reverse() {
	// anything to swap at all?
	if l.Collection == nil || len(l.Collection) < 2 {
		return
	}

	// swap elements
	for i, j := 0, len(l.Collection)-1; i < j; i, j = i+1, j-1 {
		l.Collection[i], l.Collection[j] = l.Collection[j], l.Collection[i]
	}

	// swap indexes
	l.First, l.Last = l.Last, l.First
}

// Package types implements different core types of the API.
package types

import "go.mongodb.org/mongo-driver/bson"

// TransactionList represents a list of transactions.
type TransactionList struct {
	// Collection represent list of transactions' hash.
	Collection []*Transaction

	// Total indicates total number of transaction in the whole collection.
	Total uint64

	// First is the index of the first item on the list
	First uint64

	// Last is the index of the last item on the list
	Last uint64

	// IsStart indicates there are no transactions available above the list currently.
	IsStart bool

	// IsEnd indicates there are no transactions available below the list currently.
	IsEnd bool

	// Filter represents the base filter used for filtering the list
	Filter bson.D
}

// Reverse reverses the order of transactions in the list.
func (b *TransactionList) Reverse() {
	// anything to swap at all?
	if b.Collection == nil || len(b.Collection) < 2 {
		return
	}

	// swap elements
	for i, j := 0, len(b.Collection)-1; i < j; i, j = i+1, j-1 {
		b.Collection[i], b.Collection[j] = b.Collection[j], b.Collection[i]
	}

	// swap indexes
	b.First, b.Last = b.Last, b.First
}

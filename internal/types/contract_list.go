// Package types implements different core types of the API.
package types

// ContractList represents a list of smart contracts.
type ContractList struct {
	// List keeps the actual Collection.
	Collection []*Contract

	// Total indicates total number of contracts in the whole collection.
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

// Reverse reverses the order of contracts in the list.
func (c *ContractList) Reverse() {
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

// Package types implements different core types of the API.
package types

// EpochList represents a list of epochs.
type EpochList struct {
	// List keeps the actual Collection.
	Collection []*Epoch

	// Total indicates total number of withdraw requests in the whole collection.
	Total uint64

	// First is the index of the first item on the list
	First uint64

	// Last is the index of the last item on the list
	Last uint64

	// IsStart indicates there are no withdraw requests available above the list currently.
	IsStart bool

	// IsEnd indicates there are no withdraw requests available below the list currently.
	IsEnd bool
}

// Reverse reverses the order of delegations in the list.
func (c *EpochList) Reverse() {
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

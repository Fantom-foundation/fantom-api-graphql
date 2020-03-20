// Package types implements different core types of the API.
package types

// BlockList represents a list of blocks.
type BlockList struct {
	// List keeps the actual Collection.
	Collection []*Block

	// IsStart indicates there are no blocks available above the list currently.
	IsStart bool

	// IsEnd indicates there are no blocks available below the list currently.
	IsEnd bool
}

// Reverse reverses the order of blocks in the list.
func (b *BlockList) Reverse() {
	// anything to swap at all?
	if b.Collection == nil || len(b.Collection) < 2 {
		return
	}

	// swap elements
	for i, j := 0, len(b.Collection)-1; i < j; i, j = i+1, j-1 {
		b.Collection[i], b.Collection[j] = b.Collection[j], b.Collection[i]
	}
}

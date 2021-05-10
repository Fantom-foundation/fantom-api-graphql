// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"fantom-api-graphql/internal/types"
	"unsafe"
)

// AddBlock adds a new block to the in-memory ring for fast load.
func (b *MemBridge) AddBlock(blk *types.Block) {
	if blk != nil {
		b.blkRing.Add((unsafe.Pointer)(blk))
	}
}

// ListBlocks pulls the list of blocks from the block ring.
func (b *MemBridge) ListBlocks(length int) []*types.Block {
	// pull raw data
	l := b.blkRing.List(length)

	// convert to the target type
	out := make([]*types.Block, len(l))
	for i := 0; i < len(l); i++ {
		out[i] = (*types.Block)(l[i])
	}
	return out
}

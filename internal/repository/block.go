/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"errors"
	"fantom-api-graphql/internal/repository/cache"
	"fantom-api-graphql/internal/repository/rpc"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	etc "github.com/ethereum/go-ethereum/core/types"
	eth "github.com/ethereum/go-ethereum/rpc"
)

// ErrBlockNotFound represents an error returned if a block can not be found.
var ErrBlockNotFound = errors.New("requested block can not be found in Opera blockchain")

// ObservedHeaders provides a channel fed with new headers observed
// by the connected blockchain node.
func (p *proxy) ObservedHeaders() chan *etc.Header {
	return p.rpc.ObservedBlockProxy()
}

// BlockHeight returns the current height of the Opera blockchain in blocks.
func (p *proxy) BlockHeight() (*hexutil.Big, error) {
	return p.rpc.BlockHeight()
}

// LastKnownBlock returns number of the last block known to the repository.
func (p *proxy) LastKnownBlock() (uint64, error) {
	return p.db.LastKnownBlock()
}

// UpdateLastKnownBlock update record about last known block.
func (p *proxy) UpdateLastKnownBlock(blockNo *hexutil.Uint64) error {
	return p.db.UpdateLastKnownBlock(blockNo)
}

// CacheBlock puts a block to the internal block cache.
func (p *proxy) CacheBlock(blk *types.Block) {
	p.cache.AddBlock(blk)
}

// BlockByNumber returns a block at Opera blockchain represented by a number. Top block is returned if the number
// is not provided.
// If the block is not found, ErrBlockNotFound error is returned.
func (p *proxy) BlockByNumber(num *hexutil.Uint64) (*types.Block, error) {
	// return the top block if block number is not provided
	if num == nil {
		tag := rpc.BlockTypeLatest
		return p.blockByTag(&tag)
	}
	return p.getBlock(num.String(), p.blockByTag)
}

// BlockByHash returns a block at Opera blockchain represented by a hash. Top block is returned if the hash
// is not provided.
// If the block is not found, ErrBlockNotFound error is returned.
func (p *proxy) BlockByHash(hash *common.Hash) (*types.Block, error) {
	// do we have a hash?
	if hash == nil {
		tag := rpc.BlockTypeLatest
		return p.blockByTag(&tag)
	}
	return p.getBlock(hash.String(), p.rpc.BlockByHash)
}

// getBlock gets a block of given tag from cache, or from a repository pull function.
func (p *proxy) getBlock(tag string, pull func(*string) (*types.Block, error)) (*types.Block, error) {
	// inform what we do
	p.log.Debugf("block [%s] requested", tag)

	// try to use the in-memory cache
	if blk := p.cache.PullBlock(tag); blk != nil {
		// inform what we do
		p.log.Debugf("block [%s] loaded from cache", tag)

		// return the block
		return blk, nil
	}

	// extract the block from the chain
	blk, err := pull(&tag)
	if err != nil {
		// block simply not found?
		if err == eth.ErrNoResult {
			p.log.Warning("block not found in the blockchain")
			return nil, ErrBlockNotFound
		}

		// something went wrong
		return nil, err
	}

	// try to store the block in cache for future use
	err = p.cache.PushBlock(tag, blk)
	if err != nil {
		p.log.Errorf("can not cache; %s", err.Error())
	}

	// inform what we do
	p.log.Debugf("block [%s] loaded by pulling", tag)
	return blk, nil
}

// blockByTag returns a block at Opera blockchain represented by given tag.
// The tag could be an encoded block number, or a predefined string tag for "earliest", "latest" or "pending" block.
func (p *proxy) blockByTag(tag *string) (*types.Block, error) {
	// inform what we do
	p.log.Debugf("loading block [%s]", *tag)

	// extract the block
	block, err := p.rpc.Block(tag)
	if err != nil {
		// block simply not found?
		if err == eth.ErrNoResult {
			p.log.Warning("block not found in the blockchain")
			return nil, ErrBlockNotFound
		}

		// something went wrong
		return nil, err
	}

	return block, nil
}

// initBlockList finds and returns the first block of the list and initializes the list internals accordingly.
func (p *proxy) initBlockList(num *uint64, count int32) (*types.Block, *types.BlockList, error) {
	// start from the latest block by default
	var tag = rpc.BlockTypeLatest

	// we may want to start from bottom, or specific block instead
	if num == nil && count < 0 {
		tag = rpc.BlockTypeEarliest
	} else if num != nil {
		tag = hexutil.EncodeUint64(*num)
	}

	// inform what we are about to do
	p.log.Debugf("initializing a new blocks list using tag [%s]", tag)

	// get the latest block to start from
	fb, err := p.blockByTag(&tag)
	if err != nil {
		p.log.Critical("the starting block not found in the blockchain")
		return nil, nil, err
	}

	// make sure the first block is valid
	if fb == nil {
		p.log.Critical("the starting block is not valid")
		return nil, nil, fmt.Errorf("received invalid first block of the list")
	}

	// inform what we are about to do
	p.log.Debugf("block list starts with block [%s]", fb.Number.String())

	// prep an empty list marking already clear boundaries for missing block cursor/number
	list := types.BlockList{
		Collection: make([]*types.Block, 0),
		IsStart:    num == nil && count > 0,
		IsEnd:      num == nil && count < 0,
	}

	return fb, &list, nil
}

// pullBlocks pulls specified list of blocks from repository and calculates boundary situation.
func (p *proxy) pullBlocks(num *uint64, count int32, toPull int32, current *types.Block, list *types.BlockList) {
	// prep the scan vars
	var next *types.Block
	var tag hexutil.Uint64
	var err error

	// loop to pull all the blocks requested
	for i := int32(0); i < toPull; i++ {
		// do we have any next block waiting to be used?
		if next != nil {
			// update the list with the current pending block only if it's valid for the list
			if num == nil || uint64(current.Number) != *num {
				list.Collection = append(list.Collection, current)
			}

			// move search to next block
			current = next
		}

		// we always have a <current> block; either from successful initBlockList, or from previous scan iteration
		// we assume blocks are always consecutive; if not, the search will stop on the gap gracefully
		if count > 0 {
			tag = hexutil.Uint64(uint64(current.Number) - 1)
		} else {
			tag = hexutil.Uint64(uint64(current.Number) + 1)
		}

		// try to get next block; break the loop on search issue; in that case <next> will be nil
		next, err = p.BlockByNumber(&tag)
		if err != nil {
			break
		}
	}

	// update the list with the last pending block only if it's valid for the list
	if num == nil || uint64(current.Number) != *num {
		list.Collection = append(list.Collection, current)
	}

	// return the result
	list.IsStart, list.IsEnd = checkBlocksListBoundary(count, next, list)
}

// checkListBoundary verifies if the list of blocks is on one of the edges.
func checkBlocksListBoundary(count int32, next *types.Block, list *types.BlockList) (bool, bool) {
	return list.IsStart || (count < 0 && next == nil), list.IsEnd || (count > 0 && next == nil)
}

// Blocks pulls list of blocks starting on the specified block number and going up, or down based on count number.
// If the initial block number is not provided, we start on top, or bottom based on count value.
//
// No-number boundaries are handled as follows:
// 	- For positive count we start from the most recent block and scan to older blocks.
// 	- For negative count we start from the first block and scan to newer blocks.
func (p *proxy) Blocks(num *uint64, count int32) (*types.BlockList, error) {
	// nothing to load?
	if count == 0 {
		return nil, fmt.Errorf("nothing to do, zero blocks requested")
	}

	// fast blocks list from the rings available?
	if num == nil && count > 0 && count < cache.BlockRingCacheSize {
		bl, err := p.RecentBlocks(int(count))
		if err == nil {
			return bl, nil
		}
	}

	// slow block list
	return p.makeBlocksList(num, count)
}

// makeBlocksList creates a block list for defined blocks range.
func (p *proxy) makeBlocksList(num *uint64, count int32) (*types.BlockList, error) {
	// init the list
	current, list, err := p.initBlockList(num, count)
	if err != nil {
		return nil, err
	}

	// how many to pull at most
	toPull := count
	if count < 0 {
		toPull = -count
	}

	// if in the middle we try to pull one extra block to find out if we reached aan end
	if num != nil {
		toPull++
	}

	// get specified list of block from repository
	p.pullBlocks(num, count, toPull, current, list)

	// if we scanned from bottom up, we need to reverse the list so newer blocks are on top
	if count < 0 {
		list.Reverse()
	}
	return list, nil
}

// RecentBlocks pulls a list of the most recent blocks from the ring cache.
func (p *proxy) RecentBlocks(length int) (*types.BlockList, error) {
	// pull the quick list
	bl := p.cache.ListBlocks(length)

	// does it make sense? if so, make the list from it
	if len(bl) > 0 {
		return &types.BlockList{
			Collection: bl,
			IsStart:    true,
			IsEnd:      false,
		}, nil
	}
	return nil, fmt.Errorf("recent blocks list not available")
}

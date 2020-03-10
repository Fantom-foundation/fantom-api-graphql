/*
Repository package implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"errors"
	"fantom-api-graphql/internal/repository/rpc"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/rpc"
)

// ErrBlockNotFound represents an error returned if a block can not be found.
var ErrBlockNotFound = errors.New("requested block can not be found in Opera blockchain")

// BlockByNumber returns a block at Opera blockchain represented by a number. Top block is returned if the number
// is not provided.
// If the block is not found, ErrBlockNotFound error is returned.
func (p *proxy) BlockByNumber(num *hexutil.Uint64) (*types.Block, error) {
	// inform what we do
	p.log.Infof("block requested")

	// return the top block if block number is not provided
	if num == nil {
		return p.blockByTag(rpc.BlockTypeLatest)
	}

	// try to use the in-memory cache
	if blk := p.cache.PullBlock(num); blk != nil {
		// inform what we do
		p.log.Infof("loaded block [%s] from cache", num.String())

		// return the block
		return blk, nil
	}

	// go to the chain
	blk, err := p.blockByTag(num.String())
	if err != nil {
		return nil, err
	}

	// try to store the block in cache for future use
	err = p.cache.PushBlock(blk)
	if err != nil {
		p.log.Error(err)
	}

	return blk, nil
}

// BlockByHash returns a block at Opera blockchain represented by a hash. Top block is returned if the hash
// is not provided.
// If the block is not found, ErrBlockNotFound error is returned.
func (p *proxy) BlockByHash(hash *types.Hash) (*types.Block, error) {
	// inform what we do
	p.log.Infof("block requested")

	// return the top block if hash is not provided
	if hash == nil {
		return p.blockByTag(rpc.BlockTypeLatest)
	}

	// extract the block from the chain
	block, err := p.rpc.BlockByHash(hash)
	if err != nil {
		// block simply not found?
		if err == eth.ErrNoResult {
			p.log.Warning("block not found in the blockchain")
			return nil, ErrBlockNotFound
		}

		// something went wrong
		return nil, err
	}

	// inform what we do
	p.log.Infof("block [%s] loaded from rpc", hash.String())
	return block, nil
}

// blockByTag returns a block at Opera blockchain represented by given tag.
// The tag could be an encoded block number, or a predefined string tag for "earliest", "latest" or "pending" block.
func (p *proxy) blockByTag(tag string) (*types.Block, error) {
	// inform what we do
	p.log.Infof("loading [%s] block", tag)

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

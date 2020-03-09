/*
Rpc package implements bridge to Lachesis full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Opera/Lachesis node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Lachesis RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Lachesis RPC interface with connection limited to specified endpoints.

We strongly discourage opening Lachesis RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"fantom-api-graphql/internal/types"
)

// BlockTypeLatest represents the latest available block in blockchain.
const BlockTypeLatest = "latest"

// Block returns information about a blockchain block by encoded hex number, or by a type tag.
// For tag based loading use predefined BlockType contacts.
func (b *OperaBridge) Block(numTag string) (*types.Block, error) {
	// keep track of the operation
	b.log.Debugf("loading details of block num/tag %s", numTag)

	// call for data
	var block types.Block
	err := b.rpc.Call(&block, "ftm_getBlockByNumber", numTag, false)
	if err != nil {
		b.log.Error("block could not be extracted")
		return nil, err
	}

	// keep track of the operation
	b.log.Debugf("block found for num/tag %s", numTag)
	return &block, nil
}

// BlockByHash returns information about a blockchain block by hash.
func (b *OperaBridge) BlockByHash(hash *types.Hash) (*types.Block, error) {
	// keep track of the operation
	b.log.Debugf("loading details of block %s", hash.String())

	// call for data
	var block types.Block
	err := b.rpc.Call(&block, "ftm_getBlockByHash", hash.String(), false)
	if err != nil {
		b.log.Error("block could not be extracted")
		return nil, err
	}

	// inform and return
	b.log.Debugf("block found for hash %s", hash.String())
	return &block, nil
}

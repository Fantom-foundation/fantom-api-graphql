/*
Package rpc implements bridge to Lachesis full node API interface.

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
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// BlockTypeLatest represents the latest available block in blockchain.
const (
	BlockTypeLatest   = "latest"
	BlockTypeEarliest = "earliest"
)

// BlockHeight returns the current block height of the Opera blockchain.
func (ftm *FtmBridge) BlockHeight() (*hexutil.Big, error) {
	// keep track of the operation
	ftm.log.Debugf("checking current block height")

	// call for data
	var height hexutil.Big
	err := ftm.rpc.Call(&height, "ftm_blockNumber")
	if err != nil {
		ftm.log.Error("block height could not be obtained")
		return nil, err
	}

	// inform and return
	ftm.log.Debugf("current block height is %s", height.String())
	return &height, nil
}

// Block returns information about a blockchain block by encoded hex number, or by a type tag.
// For tag based loading use predefined BlockType contacts.
func (ftm *FtmBridge) Block(numTag *string) (*types.Block, error) {
	// keep track of the operation
	ftm.log.Debugf("loading details of block num/tag %s", *numTag)

	// call for data
	var block types.Block
	err := ftm.rpc.Call(&block, "ftm_getBlockByNumber", numTag, false)
	if err != nil {
		ftm.log.Error("block could not be extracted")
		return nil, err
	}

	// detect block not found situation; block number is zero and the hash is also zero
	if uint64(block.Number) == 0 && block.Hash.Big().Cmp(big.NewInt(0)) == 0 {
		ftm.log.Debugf("block [%s] not found", *numTag)
		return nil, fmt.Errorf("block not found")
	}

	// keep track of the operation
	ftm.log.Debugf("block found for num/tag %s", *numTag)
	return &block, nil
}

// BlockByHash returns information about a blockchain block by hash.
func (ftm *FtmBridge) BlockByHash(hash *string) (*types.Block, error) {
	// keep track of the operation
	ftm.log.Debugf("loading details of block %s", *hash)

	// call for data
	var block types.Block
	err := ftm.rpc.Call(&block, "ftm_getBlockByHash", hash, false)
	if err != nil {
		ftm.log.Error("block could not be extracted")
		return nil, err
	}

	// detect block not found situation
	if uint64(block.Number) == 0 {
		ftm.log.Debugf("block [%s] not found", *hash)
		return nil, fmt.Errorf("block not found")
	}

	// inform and return
	ftm.log.Debugf("block found for hash %s", *hash)
	return &block, nil
}

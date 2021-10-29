// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Block represents resolvable blockchain block structure.
type Block struct {
	types.Block
}

// NewBlock builds new resolvable block structure.
func NewBlock(blk *types.Block) *Block {
	if blk == nil {
		return nil
	}
	return &Block{Block: *blk}
}

// Block resolves blockchain block by number or by hash. If neither is provided, the most recent block is given.
func (rs *rootResolver) Block(args *struct {
	Number *hexutil.Uint64
	Hash   *common.Hash
}) (*Block, error) {
	// do we have the number, or hash is not given?
	if args.Number != nil || args.Hash == nil {
		b, err := repository.R().BlockByNumber(args.Number)
		return NewBlock(b), err
	}

	// simply pull the block by hash
	b, err := repository.R().BlockByHash(args.Hash)
	return NewBlock(b), err
}

// Parent resolves parent block information to the given block.
func (blk *Block) Parent() (*Block, error) {
	// get the parent block by hash
	parent, err := repository.R().BlockByHash(&blk.ParentHash)
	return NewBlock(parent), err
}

// TxHashList resolves list of hashes of transaction bundled in the block.
func (blk *Block) TxHashList() []common.Hash {
	// make the container and fill it with data
	txs := make([]common.Hash, len(blk.Txs))
	for i, hash := range blk.Txs {
		txs[i] = *hash
	}
	return txs
}

// TxList resolves list of transaction details of the transactions bundled in the block.
func (blk *Block) TxList() ([]*Transaction, error) {
	// make the container
	txs := make([]*Transaction, len(blk.Txs))

	// loop the hashes and extract transactions
	for i, hash := range blk.Txs {
		trx, err := repository.R().Transaction(hash)
		if err != nil {
			return nil, err
		}

		// make a resolvable transaction
		txs[i] = NewTransaction(trx)
	}

	return txs, nil
}

// TransactionCount resolves number of transactions in the block.
func (blk *Block) TransactionCount() *int32 {
	count := int32(len(blk.Txs))
	return &count
}

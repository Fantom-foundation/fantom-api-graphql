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
	repo repository.Repository
	types.Block
}

// NewBlock builds new resolvable block structure.
func NewBlock(blk *types.Block, repo repository.Repository) *Block {
	return &Block{
		repo:  repo,
		Block: *blk,
	}
}

// Block resolves blockchain block by number or by hash. If neither is provided, the most recent block is given.
func (rs *rootResolver) Block(args *struct {
	Number *hexutil.Uint64
	Hash   *common.Hash
}) (*Block, error) {
	// do we have the number, or hash is not given?
	if args.Number != nil || args.Hash == nil {
		// get the block by the number, or get the top block
		block, err := rs.repo.BlockByNumber(args.Number)
		if err != nil {
			rs.log.Errorf("could not get the specified block")
			return nil, err
		}

		return NewBlock(block, rs.repo), nil
	}

	// simply pull the block by hash
	block, err := rs.repo.BlockByHash(args.Hash)
	if err != nil {
		rs.log.Errorf("could not get the specified block")
		return nil, err
	}

	return NewBlock(block, rs.repo), nil
}

// Parent resolves parent block information to the given block.
func (blk *Block) Parent() (*Block, error) {
	// get the parent block by hash
	parent, err := blk.repo.BlockByHash(&blk.ParentHash)
	if err != nil {
		return nil, err
	}

	return NewBlock(parent, blk.repo), nil
}

// TxHashList resolves list of hashes of transaction bundled in the block.
func (blk *Block) TxHashList() []common.Hash {
	// make the container
	txs := make([]common.Hash, len(blk.Txs))

	// loop hashes and extract them
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
		trx, err := blk.repo.Transaction(hash)
		if err != nil {
			return txs, err
		}

		// make a resolvable transaction
		txs[i] = NewTransaction(trx, blk.repo)
	}

	return txs, nil
}

// TransactionCount resolves number of transactions in the block.
func (blk *Block) TransactionCount() *int32 {
	count := int32(len(blk.Txs))
	return &count
}

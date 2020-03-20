// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// listMaxEdgesPerRequest maximal number of edges end-client can request in one query.
const listMaxEdgesPerRequest = 100

// BlockList represents resolvable list of blockchain block edges structure.
type BlockList struct {
	repo       repository.Repository
	list       *types.BlockList
	TotalCount hexutil.Big
}

// BlockListEdge represents a single edge of a block list structure.
type BlockListEdge struct {
	Block  *Block
	Cursor Cursor
}

// NewBlockList builds new resolvable list of blocks.
func NewBlockList(blocks *types.BlockList, totalCount *hexutil.Big, repo repository.Repository) *BlockList {
	return &BlockList{
		repo:       repo,
		list:       blocks,
		TotalCount: *totalCount,
	}
}

// Blocks resolves list of blockchain blocks encapsulated in a listable structure.
func (rs *rootResolver) Blocks(args *struct {
	Cursor *Cursor
	Count  int32
}) (*BlockList, error) {
	// find the cursor
	var num *uint64

	// do we have a cursor? try to decode it into an actual block number
	if args.Cursor != nil {
		val, err := hexutil.DecodeUint64(string(*args.Cursor))
		if err != nil {
			rs.log.Errorf("invalid block cursor [%s]; %s", args.Cursor, err.Error())
		}
		num = &val
	}

	// get the first block so we know the total
	bh, err := rs.repo.BlockHeight()
	if err != nil {
		return nil, err
	}

	// limit query size
	if args.Count > listMaxEdgesPerRequest {
		args.Count = listMaxEdgesPerRequest
	}

	// get the block list from repository
	bl, err := rs.repo.Blocks(num, args.Count)
	if err != nil {
		rs.log.Errorf("can not get blocks list; %s", err.Error())
		return nil, err
	}

	return NewBlockList(bl, bh, rs.repo), nil
}

// PageInfo resolves the current page information for the blocks list.
func (bl *BlockList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if bl.list == nil || bl.list.Collection == nil || len(bl.list.Collection) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get the first and last elements
	first := Cursor(bl.list.Collection[0].Number.String())
	last := Cursor(bl.list.Collection[len(bl.list.Collection)-1].Number.String())
	return NewListPageInfo(&first, &last, !bl.list.IsEnd, !bl.list.IsStart)
}

// Edges resolves list of block list edges for the linked block list.
func (bl *BlockList) Edges() []*BlockListEdge {
	// do we have any items? return empty list if not
	if bl.list == nil || bl.list.Collection == nil || len(bl.list.Collection) == 0 {
		return make([]*BlockListEdge, 0)
	}

	// make the list
	edges := make([]*BlockListEdge, len(bl.list.Collection))
	for i, b := range bl.list.Collection {
		// get the resolvable block
		blk := NewBlock(b, bl.repo)

		// make the element
		edge := BlockListEdge{
			Block:  blk,
			Cursor: Cursor(b.Number.String()),
		}

		// add it to the list
		edges[i] = &edge
	}

	return edges
}

// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// BlockList represents resolvable list of blockchain block edges structure.
type BlockList struct {
	list       *types.BlockList
	TotalCount hexutil.Big
}

// BlockListEdge represents a single edge of a block list structure.
type BlockListEdge struct {
	Block  *Block
	Cursor Cursor
}

// NewBlockList builds new resolvable list of blocks.
func NewBlockList(blocks *types.BlockList, totalCount *hexutil.Big) *BlockList {
	return &BlockList{
		list:       blocks,
		TotalCount: *totalCount,
	}
}

// Blocks resolves list of blockchain blocks encapsulated in a listable structure.
func (rs *rootResolver) Blocks(args *struct {
	Cursor *Cursor
	Count  int32
}) (*BlockList, error) {
	// do we have a cursor? try to decode it into an actual block number
	var num *uint64
	if args.Cursor != nil {
		val, err := hexutil.DecodeUint64(string(*args.Cursor))
		if err != nil {
			log.Errorf("invalid block cursor [%s]; %s", args.Cursor, err.Error())
		}
		num = &val
	}

	// get the first block so we know the total
	bh, err := repository.R().BlockHeight()
	if err != nil {
		return nil, err
	}

	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// get the block list from repository
	bl, err := repository.R().Blocks(num, args.Count)
	if err != nil {
		log.Errorf("can not get blocks list; %s", err.Error())
		return nil, err
	}

	return NewBlockList(bl, bh), nil
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
		// make the element
		edge := BlockListEdge{
			Block:  NewBlock(b),
			Cursor: Cursor(b.Number.String()),
		}

		// add it to the list
		edges[i] = &edge
	}
	return edges
}

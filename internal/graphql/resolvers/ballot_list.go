// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
)

// BallotList represents resolvable list of ballot edges structure.
type BallotList struct {
	repo repository.Repository
	list *types.BallotList
}

// BallotListEdge represents a single edge of a ballot list structure.
type BallotListEdge struct {
	Ballot *Ballot
	Cursor Cursor
}

// NewBallotList builds new resolvable list of ballots.
func NewBallotList(bl *types.BallotList, repo repository.Repository) *BallotList {
	return &BallotList{
		repo: repo,
		list: bl,
	}
}

// Ballots resolves list of ballots encapsulated in a listable structure.
func (rs *rootResolver) Ballots(args *struct {
	Cursor *Cursor
	Count  int32
}) (*BallotList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// get the contract list from repository
	bl, err := rs.repo.Ballots((*string)(args.Cursor), args.Count)
	if err != nil {
		rs.log.Errorf("can not get ballots list; %s", err.Error())
		return nil, err
	}

	return NewBallotList(bl, rs.repo), nil
}

// TotalCount resolves the total number of ballots in the list.
func (bl *BallotList) TotalCount() hexutil.Big {
	val := (*hexutil.Big)(big.NewInt(int64(bl.list.Total)))
	return *val
}

// PageInfo resolves the current page information for the ballot list.
func (bl *BallotList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if bl.list == nil || bl.list.Collection == nil || len(bl.list.Collection) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get the first and last elements
	first := Cursor(strconv.FormatUint(bl.list.First, 10))
	last := Cursor(strconv.FormatUint(bl.list.Last, 10))
	return NewListPageInfo(&first, &last, !bl.list.IsEnd, !bl.list.IsStart)
}

// Edges resolves list of edges for the linked ballot list.
func (bl *BallotList) Edges() []*BallotListEdge {
	// do we have any items? return empty list if not
	if bl.list == nil || bl.list.Collection == nil || len(bl.list.Collection) == 0 {
		return make([]*BallotListEdge, 0)
	}

	// make the list
	edges := make([]*BallotListEdge, len(bl.list.Collection))
	for i, b := range bl.list.Collection {
		// make the contract ref
		bb := NewBallot(b, bl.repo)

		// make the element
		edge := BallotListEdge{
			Ballot: bb,
			Cursor: Cursor(strconv.FormatUint(bb.OrdinalIndex, 10)),
		}

		// add it to the list
		edges[i] = &edge
	}

	return edges
}

// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"sort"
)

// DelegationList represents resolvable list of blockchain delegation edges structure.
type DelegationList struct {
	repo   repository.Repository
	cursor uint64
	count  int32
	list   []types.Delegation
}

// DelegationListEdge represents a single edge of a delegations list structure.
type DelegationListEdge struct {
	Delegation Delegation
	Cursor     Cursor
}

// NewDelegationList builds new resolvable list of delegations.
func NewDelegationList(dl []types.Delegation, cursor uint64, count int32, repo repository.Repository) *DelegationList {
	return &DelegationList{
		repo:   repo,
		cursor: cursor,
		count:  count,
		list:   dl,
	}
}

// TotalCount resolves the total number of delegations in the list.
func (dl *DelegationList) TotalCount() hexutil.Big {
	val := (*hexutil.Big)(big.NewInt(int64(len(dl.list))))
	return *val
}

// firstIndex return the index of the first item of requested slice of the delegations list.
func (dl *DelegationList) firstIndex() uint64 {
	// negative count?
	if dl.count < 0 {
		// adjust for the list position
		val := int64(dl.cursor) + int64(dl.count)
		if val < 0 {
			val = 0
		}
		return uint64(val)
	}

	return dl.cursor
}

// lastIndex return the index of the last item of requested slice of the delegations list.
func (dl *DelegationList) lastIndex() uint64 {
	// negative count?
	if dl.count < 0 {
		if dl.cursor > 0 {
			return dl.cursor
		}
		return 0
	}

	// adjust for the list position
	val := dl.cursor + uint64(dl.count)
	if val > uint64(len(dl.list)) {
		val = uint64(len(dl.list))
	}

	return val
}

// PageInfo resolves the current page information for the transaction list.
func (dl *DelegationList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if len(dl.list) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get slice range
	iFirst := dl.firstIndex()
	iLast := dl.lastIndex()

	// get the first and last elements
	first := Cursor(hexutil.Uint64(iFirst).String())
	last := Cursor(hexutil.Uint64(iLast - 1).String())
	return NewListPageInfo(&first, &last, iLast+1 < uint64(len(dl.list)), iFirst > 0)
}

// Edges resolves list of block list edges for the linked block list.
func (dl *DelegationList) Edges() []*DelegationListEdge {
	// do we have any items? return empty list if not
	if len(dl.list) == 0 {
		return make([]*DelegationListEdge, 0)
	}

	// get slice range
	iFirst := dl.firstIndex()
	iLast := dl.lastIndex()

	// any range available?
	if iLast <= iFirst {
		return make([]*DelegationListEdge, 0)
	}

	// make the list
	edges := make([]*DelegationListEdge, iLast-iFirst)
	for i, d := range dl.list[iFirst:iLast] {
		// make delegator
		del := NewDelegation(&d, dl.repo)

		// make the edge
		e := DelegationListEdge{
			Delegation: *del,
			Cursor:     Cursor(hexutil.Uint64(iFirst + uint64(i)).String()),
		}

		// add to the list and advance the edge list index
		edges[i] = &e
	}

	return edges
}

// parseDelegationsCursor decodes incoming cursor for the delegations list.
func parseDelegationsCursor(c *Cursor, count int32, dl []types.Delegation) uint64 {
	// try to parse the cursor if any
	var cr uint64
	if c != nil {
		// try to decode the cursor
		var err error
		cr, err = hexutil.DecodeUint64(string(*c))
		if err != nil {
			return 0
		}

		// adjust position based on the count
		if count > 0 {
			cr = cr + 1
		}
	} else if count < 0 {
		// start from the bottom
		cr = uint64(len(dl))
	}

	// cursor beyond range? reset to top
	if cr > uint64(len(dl)) {
		cr = 0
	}

	return cr
}

// Resolves a list of delegations information of a staker.
func (rs *rootResolver) DelegationsOf(args *struct {
	Staker hexutil.Uint64
	Cursor *Cursor
	Count  int32
}) (*DelegationList, error) {
	// any arguments?
	if args == nil {
		return nil, fmt.Errorf("missing delegations input")
	}

	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// get the list
	dl, err := rs.repo.DelegationsOf(args.Staker)
	if err != nil {
		return nil, err
	}

	// sort by the date of creation
	sort.Sort(DelegationsByAge(dl))
	return NewDelegationList(dl, parseDelegationsCursor(args.Cursor, args.Count, dl), args.Count, rs.repo), nil
}

// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// DelegationList represents resolvable list of blockchain delegation edges structure.
type DelegationList struct {
	types.DelegationList
}

// DelegationListEdge represents a single edge of a delegations list structure.
type DelegationListEdge struct {
	Delegation *Delegation
}

// NewDelegationList builds new resolvable list of delegations.
func NewDelegationList(dl *types.DelegationList) *DelegationList {
	return &DelegationList{*dl}
}

// TotalCount resolves the total number of delegations in the list.
func (dl *DelegationList) TotalCount() hexutil.Uint64 {
	return hexutil.Uint64(dl.Total)
}

// PageInfo resolves the current page information for the delegations list.
func (dl *DelegationList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if dl.Collection == nil || len(dl.Collection) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get the first and last elements
	first := Cursor(dl.Collection[0].ID)
	last := Cursor(dl.Collection[len(dl.Collection)-1].ID)
	return NewListPageInfo(&first, &last, !dl.IsEnd, !dl.IsStart)
}

// Edges resolves list of delegation list edges for the linked block list.
func (dl *DelegationList) Edges() []*DelegationListEdge {
	// do we have any items? return empty list if not
	if dl.Collection == nil || len(dl.Collection) == 0 {
		return make([]*DelegationListEdge, 0)
	}

	// make the list
	edges := make([]*DelegationListEdge, len(dl.Collection))
	for i, d := range dl.Collection {
		edges[i] = &DelegationListEdge{Delegation: NewDelegation(d)}
	}
	return edges
}

// Cursor generates the cursor for the current delegation list edge.
func (dle *DelegationListEdge) Cursor() Cursor {
	return Cursor(dle.Delegation.ID)
}

// DelegationsOf resolves a list of delegations information of a staker.
func (rs *rootResolver) DelegationsOf(args *struct {
	Staker hexutil.Big
	Cursor *Cursor
	Count  int32
}) (*DelegationList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// get the list
	dl, err := repository.R().DelegationsOfValidator(&args.Staker, (*string)(args.Cursor), args.Count)
	if err != nil {
		return nil, err
	}

	// return the list
	return NewDelegationList(dl), nil
}

// DelegationsByAddress resolves a list of own delegations by the account address.
func (rs *rootResolver) DelegationsByAddress(args *struct {
	Address common.Address
	Cursor  *Cursor
	Count   int32
}) (*DelegationList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// get the list of delegations
	dl, err := repository.R().DelegationsByAddress(&args.Address, (*string)(args.Cursor), args.Count)
	if err != nil {
		return nil, err
	}

	// return the resolvable list
	return NewDelegationList(dl), nil
}

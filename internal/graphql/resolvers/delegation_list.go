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
	Cursor     Cursor
}

// NewDelegationList builds new resolvable list of delegations.
func NewDelegationList(dl *types.DelegationList) *DelegationList {
	return &DelegationList{*dl}
}

// TotalCount resolves the total number of delegations in the list.
func (dl *DelegationList) TotalCount() hexutil.Uint64 {
	return hexutil.Uint64(dl.Total)
}

// PageInfo resolves the current page information for the transaction list.
func (dl *DelegationList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if dl.Collection == nil || len(dl.Collection) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get the first and last elements
	first := Cursor(hexutil.Uint64(dl.First).String())
	last := Cursor(hexutil.Uint64(dl.Last).String())

	return NewListPageInfo(&first, &last, !dl.IsEnd, !dl.IsStart)
}

// Edges resolves list of block list edges for the linked block list.
func (dl *DelegationList) Edges() []*DelegationListEdge {
	// do we have any items? return empty list if not
	if dl.Collection == nil || len(dl.Collection) == 0 {
		return make([]*DelegationListEdge, 0)
	}

	// make the list
	edges := make([]*DelegationListEdge, len(dl.Collection))
	for i, d := range dl.Collection {
		edges[i] = &DelegationListEdge{
			Delegation: NewDelegation(d),
			Cursor:     Cursor(hexutil.Uint64(d.Uid()).String()),
		}
	}

	return edges
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

	// decode cursor
	var cr *uint64 = nil
	if args.Cursor != nil {
		cv, err := hexutil.DecodeUint64(string(*args.Cursor))
		if err != nil {
			return nil, err
		}
		cr = &cv
	}

	// get the list
	dl, err := repository.R().DelegationsOfValidator(&args.Staker, (*hexutil.Uint64)(cr), args.Count)
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

	// decode cursor
	var cr *uint64 = nil
	if args.Cursor != nil {
		cv, err := hexutil.DecodeUint64(string(*args.Cursor))
		if err != nil {
			return nil, err
		}
		cr = &cv
	}

	// get the list of delegations
	dl, err := repository.R().DelegationsByAddress(&args.Address, (*hexutil.Uint64)(cr), args.Count)
	if err != nil {
		return nil, err
	}

	// return the resolvable list
	return NewDelegationList(dl), nil
}

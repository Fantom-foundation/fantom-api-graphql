// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// NftOwnershipList represents resolvable list of nft ownership edges structure.
type NftOwnershipList struct {
	types.NftOwnershipList
}

// NftOwnershipListEdge represents a single edge of an nft ownership list structure.
type NftOwnershipListEdge struct {
	NftOwnership *NftOwnership
}

// NewNftOwnershipList builds new resolvable list of nft ownerships.
func NewNftOwnershipList(nol *types.NftOwnershipList) *NftOwnershipList {
	return &NftOwnershipList{NftOwnershipList: *nol}
}

// NftOwnerships resolves a paginated list of nft tokens.
func (rs *rootResolver) NftOwnerships(args struct {
	Cursor     *Cursor
	Count      int32
	Collection *common.Address
	Owner      *common.Address
	TokenId    *hexutil.Big
}) (*NftOwnershipList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, accMaxTransactionsPerRequest)

	// get nft ownership list from repository
	nol, err := repository.R().ListNftOwnerships(
		args.Collection,
		args.TokenId,
		args.Owner,
		(*string)(args.Cursor),
		args.Count,
	)
	if err != nil {
		return nil, err
	}

	return NewNftOwnershipList(nol), nil
}

// TotalCount resolves the total number of ERC20 transactions in the list.
func (nol *NftOwnershipList) TotalCount() hexutil.Big {
	val := (*hexutil.Big)(new(big.Int).SetUint64(nol.Total))
	return *val
}

// PageInfo resolves the current page information for the ERC20 transaction list.
func (nol *NftOwnershipList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if nol.Collection == nil || len(nol.Collection) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get the first and last elements
	first := Cursor(nol.Collection[0].ComputedPk().String())
	last := Cursor(nol.Collection[len(nol.Collection)-1].ComputedPk().String())
	return NewListPageInfo(&first, &last, !nol.IsEnd, !nol.IsStart)
}

// Edges resolves list of edges for the linked smart contract list.
func (nol *NftOwnershipList) Edges() []*NftOwnershipListEdge {
	// do we have any items? return empty list if not
	if nol.Collection == nil || len(nol.Collection) == 0 {
		return make([]*NftOwnershipListEdge, 0)
	}

	// make the list
	edges := make([]*NftOwnershipListEdge, len(nol.Collection))
	for i, c := range nol.Collection {
		// make the element
		edge := NftOwnershipListEdge{
			NftOwnership: NewNftOwnership(c),
		}
		edges[i] = &edge
	}

	return edges
}

// Cursor resolves the nft ownership cursor in the edges list.
func (nole *NftOwnershipListEdge) Cursor() Cursor {
	return Cursor(nole.NftOwnership.ComputedPk().String())
}

package resolvers

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// TokenTransactionList represents resolvable list of ERC20 transaction edges structure.
type TokenTransactionList struct {
	types.TokenTransactionList
}

// TokenTransactionListEdge represents a single edge of an ERC20 transaction list structure.
type TokenTransactionListEdge struct {
	Trx *TokenTransaction
}

// NewERC20TransactionList builds new resolvable list of ERC20 transactions.
func NewERC20TransactionList(tl *types.TokenTransactionList) *TokenTransactionList {
	return &TokenTransactionList{TokenTransactionList: *tl}
}

// TotalCount resolves the total number of ERC20 transactions in the list.
func (txl *TokenTransactionList) TotalCount() hexutil.Big {
	val := (*hexutil.Big)(new(big.Int).SetUint64(txl.Total))
	return *val
}

// PageInfo resolves the current page information for the ERC20 transaction list.
func (txl *TokenTransactionList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if txl.Collection == nil || len(txl.Collection) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get the first and last elements
	first := Cursor(txl.Collection[0].ID)
	last := Cursor(txl.Collection[len(txl.Collection)-1].ID)
	return NewListPageInfo(&first, &last, !txl.IsEnd, !txl.IsStart)
}

// Edges resolves list of edges for the linked smart contract list.
func (txl *TokenTransactionList) Edges() []*TokenTransactionListEdge {
	// do we have any items? return empty list if not
	if txl.Collection == nil || len(txl.Collection) == 0 {
		return make([]*TokenTransactionListEdge, 0)
	}

	// make the list
	edges := make([]*TokenTransactionListEdge, len(txl.Collection))
	for i, c := range txl.Collection {
		// make the element
		edge := TokenTransactionListEdge{
			Trx: NewTokenTransaction(c),
		}
		edges[i] = &edge
	}

	return edges
}

// Cursor resolves the ERC20 transaction cursor in the edges list.
func (tle *TokenTransactionListEdge) Cursor() Cursor {
	return Cursor(tle.Trx.ID)
}

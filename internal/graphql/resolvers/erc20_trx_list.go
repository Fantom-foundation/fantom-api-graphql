package resolvers

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// ERC20TransactionList represents resolvable list of ERC20 transaction edges structure.
type ERC20TransactionList struct {
	types.TokenTransactionList
}

// ERC20TransactionListEdge represents a single edge of an ERC20 transaction list structure.
type ERC20TransactionListEdge struct {
	Trx *ERC20Transaction
}

// NewERC20TransactionList builds new resolvable list of ERC20 transactions.
func NewERC20TransactionList(tl *types.TokenTransactionList) *ERC20TransactionList {
	return &ERC20TransactionList{TokenTransactionList: *tl}
}

// TotalCount resolves the total number of ERC20 transactions in the list.
func (txl *ERC20TransactionList) TotalCount() hexutil.Big {
	val := (*hexutil.Big)(new(big.Int).SetUint64(txl.Total))
	return *val
}

// PageInfo resolves the current page information for the ERC20 transaction list.
func (txl *ERC20TransactionList) PageInfo() (*ListPageInfo, error) {
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
func (txl *ERC20TransactionList) Edges() []*ERC20TransactionListEdge {
	// do we have any items? return empty list if not
	if txl.Collection == nil || len(txl.Collection) == 0 {
		return make([]*ERC20TransactionListEdge, 0)
	}

	// make the list
	edges := make([]*ERC20TransactionListEdge, len(txl.Collection))
	for i, c := range txl.Collection {
		// make the element
		edge := ERC20TransactionListEdge{
			Trx: NewErc20Transaction(c),
		}
		edges[i] = &edge
	}

	return edges
}

// Cursor resolves the ERC20 transaction cursor in the edges list.
func (tle *ERC20TransactionListEdge) Cursor() Cursor {
	return Cursor(tle.Trx.ID)
}

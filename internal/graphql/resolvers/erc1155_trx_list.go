package resolvers

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// ERC1155TransactionList represents resolvable list of ERC1155 transaction edges structure.
type ERC1155TransactionList struct {
	types.TokenTransactionList
}

// ERC1155TransactionListEdge represents a single edge of an ERC1155 transaction list structure.
type ERC1155TransactionListEdge struct {
	Trx *ERC1155Transaction
}

// NewERC1155TransactionList builds new resolvable list of ERC1155 transactions.
func NewERC1155TransactionList(tl *types.TokenTransactionList) *ERC1155TransactionList {
	return &ERC1155TransactionList{TokenTransactionList: *tl}
}

// TotalCount resolves the total number of ERC1155 transactions in the list.
func (txl *ERC1155TransactionList) TotalCount() hexutil.Big {
	val := (*hexutil.Big)(new(big.Int).SetUint64(txl.Total))
	return *val
}

// PageInfo resolves the current page information for the ERC1155 transaction list.
func (txl *ERC1155TransactionList) PageInfo() (*ListPageInfo, error) {
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
func (txl *ERC1155TransactionList) Edges() []*ERC1155TransactionListEdge {
	// do we have any items? return empty list if not
	if txl.Collection == nil || len(txl.Collection) == 0 {
		return make([]*ERC1155TransactionListEdge, 0)
	}

	// make the list
	edges := make([]*ERC1155TransactionListEdge, len(txl.Collection))
	for i, c := range txl.Collection {
		// make the element
		edge := ERC1155TransactionListEdge{
			Trx: NewErc1155Transaction(c),
		}
		edges[i] = &edge
	}

	return edges
}

// Cursor resolves the ERC1155 transaction cursor in the edges list.
func (tle *ERC1155TransactionListEdge) Cursor() Cursor {
	return Cursor(tle.Trx.ID)
}

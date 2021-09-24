package resolvers

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// ERC721TransactionList represents resolvable list of ERC721 transaction edges structure.
type ERC721TransactionList struct {
	types.TokenTransactionList
}

// ERC721TransactionListEdge represents a single edge of an ERC721 transaction list structure.
type ERC721TransactionListEdge struct {
	Trx *ERC721Transaction
}

// NewERC721TransactionList builds new resolvable list of ERC721 transactions.
func NewERC721TransactionList(tl *types.TokenTransactionList) *ERC721TransactionList {
	return &ERC721TransactionList{TokenTransactionList: *tl}
}

// TotalCount resolves the total number of ERC721 transactions in the list.
func (txl *ERC721TransactionList) TotalCount() hexutil.Big {
	val := (*hexutil.Big)(new(big.Int).SetUint64(txl.Total))
	return *val
}

// PageInfo resolves the current page information for the ERC721 transaction list.
func (txl *ERC721TransactionList) PageInfo() (*ListPageInfo, error) {
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
func (txl *ERC721TransactionList) Edges() []*ERC721TransactionListEdge {
	// do we have any items? return empty list if not
	if txl.Collection == nil || len(txl.Collection) == 0 {
		return make([]*ERC721TransactionListEdge, 0)
	}

	// make the list
	edges := make([]*ERC721TransactionListEdge, len(txl.Collection))
	for i, c := range txl.Collection {
		// make the element
		edge := ERC721TransactionListEdge{
			Trx: NewErc721Transaction(c),
		}
		edges[i] = &edge
	}

	return edges
}

// Cursor resolves the ERC721 transaction cursor in the edges list.
func (tle *ERC721TransactionListEdge) Cursor() Cursor {
	return Cursor(tle.Trx.ID)
}

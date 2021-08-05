// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// TransactionList represents resolvable list of blockchain transaction edges structure.
type TransactionList struct {
	types.TransactionList
}

// TransactionListEdge represents a single edge of a transaction list structure.
type TransactionListEdge struct {
	Transaction *Transaction
	Cursor      Cursor
}

// NewTransactionList builds new resolvable list of transactions.
func NewTransactionList(txs *types.TransactionList) *TransactionList {
	return &TransactionList{
		TransactionList: *txs,
	}
}

// Transactions resolves list of blockchain transactions encapsulated in a listable structure.
func (rs *rootResolver) Transactions(args *struct {
	Cursor *Cursor
	Count  int32
}) (*TransactionList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// get the transaction hash list from repository
	txs, err := repository.R().Transactions((*string)(args.Cursor), args.Count)
	if err != nil {
		log.Errorf("can not get transactions list; %s", err.Error())
		return nil, err
	}
	return NewTransactionList(txs), nil
}

// TotalCount resolves the total number of transactions in the list.
func (tl *TransactionList) TotalCount() hexutil.Big {
	val := (*hexutil.Big)(big.NewInt(int64(tl.Total)))
	return *val
}

// PageInfo resolves the current page information for the transaction list.
func (tl *TransactionList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if tl.Collection == nil || len(tl.Collection) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get the first and last elements
	first := Cursor(tl.Collection[0].Hash.String())
	last := Cursor(tl.Collection[len(tl.Collection)-1].Hash.String())
	return NewListPageInfo(&first, &last, !tl.IsEnd, !tl.IsStart)
}

// Edges resolves list of transaction list edges for the linked transaction list.
func (tl *TransactionList) Edges() []*TransactionListEdge {
	// do we have any items? return empty list if not
	if tl.Collection == nil || len(tl.Collection) == 0 {
		return make([]*TransactionListEdge, 0)
	}

	// make the list
	edges := make([]*TransactionListEdge, len(tl.Collection))
	for i, t := range tl.Collection {
		// make the element
		edges[i] = &TransactionListEdge{
			Transaction: NewTransaction(t),
			Cursor:      Cursor(t.Hash.String()),
		}
	}
	return edges
}

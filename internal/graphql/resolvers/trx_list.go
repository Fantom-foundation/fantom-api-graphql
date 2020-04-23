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
	repo repository.Repository
	list *types.TransactionHashList
}

// TransactionListEdge represents a single edge of a transaction list structure.
type TransactionListEdge struct {
	Transaction *Transaction
	Cursor      Cursor
}

// NewTransactionList builds new resolvable list of transactions.
func NewTransactionList(txs *types.TransactionHashList, repo repository.Repository) *TransactionList {
	return &TransactionList{
		repo: repo,
		list: txs,
	}
}

// Transactions resolves list of blockchain transactions encapsulated in a listable structure.
func (rs *rootResolver) Transactions(args *struct {
	Cursor *Cursor
	Count  int32
}) (*TransactionList, error) {
	// limit query size
	if args.Count > listMaxEdgesPerRequest {
		args.Count = listMaxEdgesPerRequest
	}

	// get the transaction hash list from repository
	txs, err := rs.repo.Transactions((*string)(args.Cursor), args.Count)
	if err != nil {
		rs.log.Errorf("can not get transactions list; %s", err.Error())
		return nil, err
	}

	return NewTransactionList(txs, rs.repo), nil
}

// TotalCount resolves the total number of transactions in the list.
func (tl *TransactionList) TotalCount() hexutil.Big {
	val := (*hexutil.Big)(big.NewInt(int64(tl.list.Total)))
	return *val
}

// PageInfo resolves the current page information for the transaction list.
func (tl *TransactionList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if tl.list == nil || tl.list.Collection == nil || len(tl.list.Collection) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get the first and last elements
	first := Cursor(tl.list.Collection[0].String())
	last := Cursor(tl.list.Collection[len(tl.list.Collection)-1].String())
	return NewListPageInfo(&first, &last, !tl.list.IsEnd, !tl.list.IsStart)
}

// Edges resolves list of block list edges for the linked block list.
func (tl *TransactionList) Edges() []*TransactionListEdge {
	// do we have any items? return empty list if not
	if tl.list == nil || tl.list.Collection == nil || len(tl.list.Collection) == 0 {
		return make([]*TransactionListEdge, 0)
	}

	// make the list
	edges := make([]*TransactionListEdge, len(tl.list.Collection))
	for i, t := range tl.list.Collection {
		// get the transaction
		tx, err := tl.repo.Transaction(t)
		if err == nil {
			// get the resolvable transaction
			trx := NewTransaction(tx, tl.repo)

			// make the element
			edge := TransactionListEdge{
				Transaction: trx,
				Cursor:      Cursor(t.String()),
			}

			// add it to the list
			edges[i] = &edge
		}
	}

	return edges
}

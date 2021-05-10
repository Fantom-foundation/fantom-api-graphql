// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"fantom-api-graphql/internal/types"
	"unsafe"
)

// AddTransaction adds a new transaction to the in-memory ring for fast load.
func (b *MemBridge) AddTransaction(trx *types.Transaction) {
	if trx != nil {
		b.trxRing.Add((unsafe.Pointer)(trx))
	}
}

// ListTransactions pulls the list of transactions from the trx ring.
func (b *MemBridge) ListTransactions(length int) []*types.Transaction {
	l := b.trxRing.List(length)

	// convert to the target type
	out := make([]*types.Transaction, len(l))
	for i := 0; i < len(l); i++ {
		out[i] = (*types.Transaction)(l[i])
	}
	return out
}

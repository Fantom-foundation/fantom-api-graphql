package resolvers

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// TokenTransaction represents a resolvable generic token transaction.
type TokenTransaction struct {
	types.TokenTransaction
}

// NewTokenTransaction creates a new instance of resolvable generic token transaction.
func NewTokenTransaction(t *types.TokenTransaction) *TokenTransaction {
	return &TokenTransaction{TokenTransaction: *t}
}

// Hash resolves base transaction call hash.
func (ttx *TokenTransaction) Hash() common.Hash {
	return ttx.Transaction
}

// BlockNumber resolves the number of the transaction block.
func (ttx *TokenTransaction) BlockNumber() hexutil.Uint64 {
	return hexutil.Uint64(ttx.TokenTransaction.BlockNumber)
}

// Type resolves human-readable type of the transaction.
func (ttx *TokenTransaction) Type() string {
	return ercTrxTypeToName(ttx.TokenTransaction.Type)
}

// LogIndex resolves the index of the token transaction inside the block.
func (ttx *TokenTransaction) LogIndex() hexutil.Uint64 {
	return hexutil.Uint64(ttx.TokenTransaction.LogIndex)
}

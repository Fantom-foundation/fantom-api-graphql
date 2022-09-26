package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
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

// TrxIndex resolves index of the transaction in the block.
func (ttx *TokenTransaction) TrxIndex() hexutil.Uint64 {
	return hexutil.Uint64(ttx.TokenTransaction.TrxIndex)
}

// Type resolves human-readable type of the transaction.
func (ttx *TokenTransaction) Type() string {
	return ercTrxTypeToName(ttx.TokenTransaction.TrxType)
}

// TokenName resolves the name of the ERC20 token contract, if available.
func (ttx *TokenTransaction) TokenName() (name string, err error) {
	return repository.R().Erc20Name(&ttx.TokenTransaction.TokenAddress)
}

// TokenSymbol resolves the symbol of the ERC20 token contract, if available.
func (ttx *TokenTransaction) TokenSymbol() (sym string, err error) {
	return repository.R().Erc20Symbol(&ttx.TokenTransaction.TokenAddress)
}

// TokenDecimals resolves the amount of decimals of the ERC20 token contract, if available.
func (ttx *TokenTransaction) TokenDecimals() (decimals int32, err error) {
	return repository.R().Erc20Decimals(&ttx.TokenTransaction.TokenAddress)
}

// Timestamp returns timestamp.
func (ttx *TokenTransaction) Timestamp() hexutil.Uint64 {
	return hexutil.Uint64(ttx.TimeStamp.Unix())
}

// TokenType returns type of token.
func (ttx *TokenTransaction) TokenType() string {
	// always return erc20 to keep compatibility with existing graphql schema
	return "ERC20"
}

// TokenId returns token identifier.
func (ttx *TokenTransaction) TokenId() hexutil.Big {
	// always return 0 to keep compatibility with existing graphql schema
	tokenId := big.NewInt(0)
	return hexutil.Big(*tokenId)
}

package resolvers

import (
	"fantom-api-graphql/internal/repository"
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

// TokenName resolves the name of the ERC token contract, if available.
func (ttx *TokenTransaction) TokenName() (name string, err error) {
	switch ttx.TokenTransaction.TokenType {
	case types.AccountTypeERC20Token:
		name, err = repository.R().Erc20Name(&ttx.TokenTransaction.TokenAddress)
	case types.AccountTypeERC721Contract:
		name, err = repository.R().Erc721Name(&ttx.TokenTransaction.TokenAddress)
	default:
		name, err = "", nil
	}
	return
}

// TokenSymbol resolves the symbol of the ERC token contract, if available.
func (ttx *TokenTransaction) TokenSymbol() (sym string, err error) {
	switch ttx.TokenTransaction.TokenType {
	case types.AccountTypeERC20Token:
		sym, err = repository.R().Erc20Symbol(&ttx.TokenTransaction.TokenAddress)
	case types.AccountTypeERC721Contract:
		sym, err = repository.R().Erc721Symbol(&ttx.TokenTransaction.TokenAddress)
	default:
		sym, err = "", nil
	}
	return
}

// TokenDecimals resolves the amount of decimals of the ERC token contract, if available.
func (ttx *TokenTransaction) TokenDecimals() (decimals int32, err error) {
	switch ttx.TokenTransaction.TokenType {
	case types.AccountTypeERC20Token:
		decimals, err = repository.R().Erc20Decimals(&ttx.TokenTransaction.TokenAddress)
	case types.AccountTypeERC721Contract:
		decimals = 0
	default:
		decimals = 0
	}
	return
}

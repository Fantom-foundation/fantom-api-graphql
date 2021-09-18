package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

// ERC20Transaction represents a resolvable ERC20 token transaction.
type ERC20Transaction struct {
	types.TokenTransaction
}

// NewErc20Transaction creates a new instance of resolvable ERC20 transaction.
func NewErc20Transaction(t *types.TokenTransaction) *ERC20Transaction {
	return &ERC20Transaction{TokenTransaction: *t}
}

// TrxHash resolves the hash of the transaction executing the ERC20 call.
func (trx *ERC20Transaction) TrxHash() common.Hash {
	return trx.TokenTransaction.Transaction
}

// Transaction resolves an instance of the transaction executing the ERC20 call.
func (trx *ERC20Transaction) Transaction() (*Transaction, error) {
	// get the transaction from repo
	tx, err := repository.R().Transaction(&trx.TokenTransaction.Transaction)
	if err != nil {
		return nil, err
	}
	return NewTransaction(tx), nil
}

// Token resolves instance of the ERC20 token involved.
func (trx *ERC20Transaction) Token() *ERC20Token {
	return NewErc20Token(&trx.TokenAddress)
}

// TrxType resolves the type of the ERC20 transaction.
func (trx *ERC20Transaction) TrxType() string {
	return ercTrxTypeToName(trx.Type)
}

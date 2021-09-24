package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

// ERC1155Transaction represents a resolvable ERC1155 token transaction.
type ERC1155Transaction struct {
	types.TokenTransaction
}

// NewErc1155Transaction creates a new instance of resolvable ERC1155 transaction.
func NewErc1155Transaction(t *types.TokenTransaction) *ERC1155Transaction {
	return &ERC1155Transaction{TokenTransaction: *t}
}

// TrxHash resolves the hash of the transaction executing the ERC1155 call.
func (trx *ERC1155Transaction) TrxHash() common.Hash {
	return trx.TokenTransaction.Transaction
}

// Transaction resolves an instance of the transaction executing the ERC1155 call.
func (trx *ERC1155Transaction) Transaction() (*Transaction, error) {
	// get the transaction from repo
	tx, err := repository.R().Transaction(&trx.TokenTransaction.Transaction)
	if err != nil {
		return nil, err
	}
	return NewTransaction(tx), nil
}

// Token resolves instance of the ERC1155 token involved.
func (trx *ERC1155Transaction) Token() *ERC1155Contract {
	return NewErc1155Contract(&trx.TokenAddress)
}

// TrxType resolves the type of the ERC1155 transaction.
func (trx *ERC1155Transaction) TrxType() string {
	return ercTrxTypeToName(trx.Type)
}

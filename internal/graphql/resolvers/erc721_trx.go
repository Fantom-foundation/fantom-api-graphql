package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

// ERC721Transaction represents a resolvable ERC721 token transaction.
type ERC721Transaction struct {
	types.TokenTransaction
}

// NewErc721Transaction creates a new instance of resolvable ERC721 transaction.
func NewErc721Transaction(t *types.TokenTransaction) *ERC721Transaction {
	return &ERC721Transaction{TokenTransaction: *t}
}

// TrxHash resolves the hash of the transaction executing the ERC721 call.
func (trx *ERC721Transaction) TrxHash() common.Hash {
	return trx.TokenTransaction.Transaction
}

// Transaction resolves an instance of the transaction executing the ERC721 call.
func (trx *ERC721Transaction) Transaction() (*Transaction, error) {
	// get the transaction from repo
	tx, err := repository.R().Transaction(&trx.TokenTransaction.Transaction)
	if err != nil {
		return nil, err
	}
	return NewTransaction(tx), nil
}

// Token resolves instance of the ERC721 token involved.
func (trx *ERC721Transaction) Token() *ERC721Contract {
	return NewErc721Contract(&trx.TokenAddress)
}

// TrxType resolves the type of the ERC721 transaction.
func (trx *ERC721Transaction) TrxType() string {
	return ercTrxTypeToName(trx.Type)
}

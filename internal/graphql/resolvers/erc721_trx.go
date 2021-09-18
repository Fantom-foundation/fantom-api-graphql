package resolvers

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ERC721TrxTypeNameTransfer = "TRANSFER"
	ERC721TrxTypeNameMint     = "MINT"
	ERC721TrxTypeNameBurn     = "BURN"
	ERC721TrxTypeNameApproval = "APPROVAL"
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
func (trx *ERC721Transaction) Token() *ERC721Token {
	return NewErc721Token(&trx.TokenAddress)
}

// TrxType resolves the type of the ERC721 transaction.
func (trx *ERC721Transaction) TrxType() string {
	switch trx.Type {
	case types.TokenTrxTypeTransfer:
		// minting
		if config.EmptyAddress == trx.Sender.String() {
			return ERC721TrxTypeNameMint
		}
		// burning
		if config.EmptyAddress == trx.Recipient.String() {
			return ERC721TrxTypeNameBurn
		}
		//regular transfer
		return ERC721TrxTypeNameTransfer
	case types.TokenTrxTypeApproval:
		return ERC721TrxTypeNameApproval
	}
	return "OTHER"
}

// erc721TrxTypeByName returns numeric type of the ERC721 transaction by its name.
// Returns nil if the name is not recognized.
func erc721TrxTypeByName(name string) *int32 {
	// decode the transaction type filter
	var txType *int32
	switch name {
	case ERC721TrxTypeNameMint:
		i := int32(types.TokenTrxTypeTransfer)
		txType = &i
	case ERC721TrxTypeNameBurn:
		i := int32(types.TokenTrxTypeTransfer)
		txType = &i
	case ERC721TrxTypeNameTransfer:
		i := int32(types.TokenTrxTypeTransfer)
		txType = &i
	case ERC721TrxTypeNameApproval:
		i := int32(types.TokenTrxTypeApproval)
		txType = &i
	}
	return txType
}

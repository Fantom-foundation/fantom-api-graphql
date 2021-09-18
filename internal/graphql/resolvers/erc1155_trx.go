package resolvers

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ERC1155TrxTypeNameTransfer = "TRANSFER"
	ERC1155TrxTypeNameMint     = "MINT"
	ERC1155TrxTypeNameBurn     = "BURN"
	ERC1155TrxTypeNameApproval = "APPROVAL"
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
	switch trx.Type {
	case types.TokenTrxTypeTransfer:
		// minting
		if config.EmptyAddress == trx.Sender.String() {
			return ERC1155TrxTypeNameMint
		}
		// burning
		if config.EmptyAddress == trx.Recipient.String() {
			return ERC1155TrxTypeNameBurn
		}
		//regular transfer
		return ERC1155TrxTypeNameTransfer
	case types.TokenTrxTypeApproval:
		return ERC1155TrxTypeNameApproval
	}
	return "OTHER"
}

// erc1155TrxTypeByName returns numeric type of the ERC1155 transaction by its name.
// Returns nil if the name is not recognized.
func erc1155TrxTypeByName(name string) *int32 {
	// decode the transaction type filter
	var txType *int32
	switch name {
	case ERC1155TrxTypeNameMint:
		i := int32(types.TokenTrxTypeTransfer)
		txType = &i
	case ERC1155TrxTypeNameBurn:
		i := int32(types.TokenTrxTypeTransfer)
		txType = &i
	case ERC1155TrxTypeNameTransfer:
		i := int32(types.TokenTrxTypeTransfer)
		txType = &i
	case ERC1155TrxTypeNameApproval:
		i := int32(types.TokenTrxTypeApproval)
		txType = &i
	}
	return txType
}

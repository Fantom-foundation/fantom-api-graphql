package resolvers

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ERC20TrxTypeNameTransfer = "TRANSFER"
	ERC20TrxTypeNameMint     = "MINT"
	ERC20TrxTypeNameBurn     = "BURN"
	ERC20TrxTypeNameApproval = "APPROVAL"
)

// ERC20Transaction represents a resolvable ERC20 token transaction.
type ERC20Transaction struct {
	types.Erc20Transaction
}

// NewErc20Transaction creates a new instance of resolvable ERC20 transaction.
func NewErc20Transaction(t *types.Erc20Transaction) *ERC20Transaction {
	return &ERC20Transaction{Erc20Transaction: *t}
}

// TrxHash resolves the hash of the transaction executing the ERC20 call.
func (trx *ERC20Transaction) TrxHash() common.Hash {
	return trx.Erc20Transaction.Transaction
}

// Transaction resolves an instance of the transaction executing the ERC20 call.
func (trx *ERC20Transaction) Transaction() (*Transaction, error) {
	// get the transaction from repo
	tx, err := repository.R().Transaction(&trx.Erc20Transaction.Transaction)
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
	switch trx.Type {
	case types.ERC20TrxTypeTransfer:
		// minting
		if config.EmptyAddress == trx.Sender.String() {
			return ERC20TrxTypeNameMint
		}
		// burning
		if config.EmptyAddress == trx.Recipient.String() {
			return ERC20TrxTypeNameBurn
		}
		//regular transfer
		return ERC20TrxTypeNameTransfer
	case types.ERC20TrxTypeApproval:
		return ERC20TrxTypeNameApproval
	}
	return "OTHER"
}

// erc20TrxTypeByName returns numeric type of the ERC20 transaction by its name.
// Returns nil if the name is not recognized.
func erc20TrxTypeByName(name string) *int32 {
	// decode the transaction type filter
	var txType *int32
	switch name {
	case ERC20TrxTypeNameMint:
		i := int32(types.ERC20TrxTypeTransfer)
		txType = &i
	case ERC20TrxTypeNameBurn:
		i := int32(types.ERC20TrxTypeTransfer)
		txType = &i
	case ERC20TrxTypeNameTransfer:
		i := int32(types.ERC20TrxTypeTransfer)
		txType = &i
	case ERC20TrxTypeNameApproval:
		i := int32(types.ERC20TrxTypeApproval)
		txType = &i
	}
	return txType
}

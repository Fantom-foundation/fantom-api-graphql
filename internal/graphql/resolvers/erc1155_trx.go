package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
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

// TrxType resolves the type of the ERC721 transaction.
func (trx *ERC1155Transaction) TrxType() string {
	return ercTrxTypeToName(trx.TokenTransaction.TrxType)
}

// TrxIndex resolves the transaction index.
func (trx *ERC1155Transaction) TrxIndex() hexutil.Uint64 {
	return hexutil.Uint64(trx.TokenTransaction.TrxIndex)
}

// TokenType returns type of token.
func (trx *ERC1155Transaction) TokenType() string {
	// always return erc1155 to keep compatibility with existing graphql schema
	return "ERC1155"
}

// Amount returns token amount.
func (trx *ERC1155Transaction) Amount() hexutil.Big {
	// always return 0 to keep compatibility with existing graphql schema
	tokenId := big.NewInt(0)
	return hexutil.Big(*tokenId)
}

// TokenId returns token identifier.
func (trx *ERC1155Transaction) TokenId() hexutil.Big {
	// always return 0 to keep compatibility with existing graphql schema
	tokenId := big.NewInt(0)
	return hexutil.Big(*tokenId)
}

// Timestamp returns timestamp.
func (trx *ERC1155Transaction) Timestamp() hexutil.Uint64 {
	return hexutil.Uint64(trx.TimeStamp.Unix())
}

/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

// StoreTokenTransaction stores ERC20/ERC721/ERC1155 transaction into the repository.
func (p *proxy) StoreTokenTransaction(trx *types.TokenTransaction) error {
	return p.db.StoreTokenTransaction(trx)
}

// TokenTransactionsByCall provides a list of token transaction made inside a specific
// transaction call (blockchain transaction).
func (p *proxy) TokenTransactionsByCall(trxHash *common.Hash) ([]*types.TokenTransaction, error) {
	return p.db.TokenTransactionsByCall(trxHash)
}

// TokenTransactions provides list of ERC20 transactions.
func (p *proxy) TokenTransactions(token *common.Address, acc *common.Address, txType []int32, cursor *string, count int32) (*types.TokenTransactionList, error) {
	return p.db.Erc20Transactions(token, acc, txType, cursor, count)
}

// Erc20Assets provides a list of known assets for the given owner.
func (p *proxy) Erc20Assets(owner common.Address, count int32) ([]common.Address, error) {
	return p.db.Erc20Assets(owner, count)
}

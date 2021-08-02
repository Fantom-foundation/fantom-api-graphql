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
	"go.mongodb.org/mongo-driver/bson"
)

// StoreErc20Transaction stores ERC20 transaction into the repository.
func (p *proxy) StoreErc20Transaction(trx *types.Erc20Transaction) error {
	return p.db.AddERC20Transaction(trx)
}

// Erc20Transactions provides list of ERC20 transactions based on given filters.
func (p *proxy) Erc20Transactions(token *common.Address, acc *common.Address, tt *int32, cursor *string, count int32) (*types.Erc20TransactionList, error) {
	// prep the filter
	fi := bson.D{}

	// filter specific token
	if token != nil {
		fi = append(fi, bson.E{
			Key:   types.FiErc20TransactionToken,
			Value: token.String(),
		})
	}

	// common address (sender or recipient)
	if acc != nil {
		fi = append(fi, bson.E{
			Key: "$or",
			Value: bson.A{bson.D{{
				Key:   types.FiErc20TransactionSender,
				Value: acc.String(),
			}}, bson.D{{
				Key:   types.FiErc20TransactionRecipient,
				Value: acc.String(),
			}}},
		})
	}

	// type of the transaction
	if tt != nil {
		fi = append(fi, bson.E{
			Key:   types.FiErc20TransactionType,
			Value: *tt,
		})
	}

	// do loading
	return p.db.Erc20Transactions(cursor, count, &fi)
}

// Erc20Assets provides a list of known assets for the given owner.
func (p *proxy) Erc20Assets(owner common.Address, count int32) ([]common.Address, error) {
	return p.db.Erc20Assets(owner, count)
}

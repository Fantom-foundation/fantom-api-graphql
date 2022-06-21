/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import "fantom-api-graphql/internal/types"

// AddFMintTransaction adds the specified fMint transaction to persistent storage.
func (p *proxy) AddFMintTransaction(trx *types.FMintTransaction) error {
	return p.db.AddFMintTransaction(trx)
}

// FMintUsers loads the list of fMint users and their associated tokens
// used for a specified transaction type.
func (p *proxy) FMintUsers(tt int32) ([]*types.FMintUserTokens, error) {
	return p.db.FMintUsers(tt)
}

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
)

// StoreFtmBurn stores the given native FTM burn per block record into the persistent storage.
func (p *proxy) StoreFtmBurn(burn *types.FtmBurn) error {
	return p.db.StoreBurn(burn)
}

// FtmBurnTotal provides the total amount of burned native FTM.
func (p *proxy) FtmBurnTotal() (int64, error) {
	return p.db.BurnTotal()
}

// FtmBurnList provides list of per-block burned native FTM tokens.
func (p *proxy) FtmBurnList(count int64) ([]types.FtmBurn, error) {
	return p.db.BurnList(count)
}

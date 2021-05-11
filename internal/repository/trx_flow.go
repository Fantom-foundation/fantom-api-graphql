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
	"time"
)

// TrxFlowVolume resolves the list of daily trx flow aggregations.
func (p *proxy) TrxFlowVolume(from *time.Time, to *time.Time) ([]*types.DailyTrxVolume, error) {
	return p.db.TrxDailyFlowList(from, to)
}

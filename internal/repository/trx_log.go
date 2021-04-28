/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	retypes "github.com/ethereum/go-ethereum/core/types"
	"sync"
)

// QueueTrxLog queues the given transaction log record for processing.
func (p *proxy) QueueTrxLog(log *retypes.Log, wg *sync.WaitGroup) {
	// put the log record to the processing queue
	p.orc.logsQueue <- &eventTrxLog{
		wg:  wg,
		Log: *log,
	}
}

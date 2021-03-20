/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/types"
	"sync"
)

// trxDispatcher implements dispatcher of new transactions in the blockchain.
type trxDispatcher struct {
	service
	buffer chan *evtTransaction
}

// evtTransaction represents a single incoming transaction event to be processed.
type evtTransaction struct {
	block *types.Block
	trx   *types.Transaction
}

// NewTrxDispatcher creates a new transaction dispatcher instance.
func newTrxDispatcher(buffer chan *evtTransaction, repo Repository, log logger.Logger, wg *sync.WaitGroup) *trxDispatcher {
	// create new dispatcher
	return &trxDispatcher{
		service: newService("trx dispatcher", repo, log, wg),
		buffer:  buffer,
	}
}

// run starts the transaction dispatcher job
func (td *trxDispatcher) run() {
	// add self to the wait group and run the dispatch routine
	td.wg.Add(1)
	go td.dispatch()
}

// Dispatch implements the dispatcher reader and router routine.
func (td *trxDispatcher) dispatch() {
	// log action
	td.log.Notice("trx dispatcher is running")

	// don't forget to sign off after we are done
	defer func() {
		// log finish
		td.log.Notice("trx dispatcher is closed")
		td.wg.Done()
	}()

	// wait for transactions and process them
	for {
		// try to read next transaction
		select {
		case toDispatch := <-td.buffer:
			// validate
			if toDispatch.block == nil || toDispatch.trx == nil {
				td.log.Critical("dispatcher received invalid transaction")
				continue
			}

			// dispatch the received
			err := td.repo.TransactionAdd(toDispatch.block, toDispatch.trx)
			if err != nil {
				td.log.Error("could not dispatch transaction")
				td.log.Error(err)
			}
		case <-td.sigStop:
			// stop the routine
			td.log.Info("trx dispatcher received stop signal")
			return
		}
	}
}

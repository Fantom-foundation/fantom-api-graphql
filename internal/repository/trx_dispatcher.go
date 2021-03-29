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

// trxDispatchQueueCapacity is the number of transactions kept in the dispatch buffer.
const trxDispatchQueueCapacity = 1000

// trxDispatcher implements dispatcher of new transactions in the blockchain.
type trxDispatcher struct {
	service
	buffer chan *eventTransaction
}

// eventTransaction represents a single incoming transaction event to be processed.
type eventTransaction struct {
	block *types.Block
	trx   *types.Transaction
}

// NewTrxDispatcher creates a new transaction dispatcher instance.
func newTrxDispatcher(buffer chan *eventTransaction, repo Repository, log logger.Logger, wg *sync.WaitGroup) *trxDispatcher {
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

// dispatch implements the dispatcher reader and router routine.
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
			td.process(toDispatch)
		case <-td.sigStop:
			// stop the routine
			td.log.Info("trx dispatcher received stop signal")
			return
		}
	}
}

// process processes the given transaction event into the required targets.
func (td *trxDispatcher) process(evt *eventTransaction) {
	// make the work group for this trx
	var wg sync.WaitGroup

	// process transaction into the accounts
	td.propagateTrxToAccounts(evt, &wg)

	// process transaction logs
	for _, lg := range evt.trx.Logs {
		wg.Add(1)
		td.repo.QueueTrxLog(&lg, &wg)
	}

	// store the transaction into the database once the processing is done
	// we spawn a lot of go-routines here so we should test the optimal queue length above
	go td.waitAndStore(evt.block, evt.trx, &wg)
}

// waitAndStore waits for the transaction processing to finish and stores the transaction into db.
func (td *trxDispatcher) waitAndStore(blk *types.Block, trx *types.Transaction, wg *sync.WaitGroup) {
	wg.Wait()
	if err := td.repo.StoreTransaction(blk, trx); err != nil {
		td.log.Errorf("can not store transaction %s from block %d", trx.Hash.String(), blk.Number)
	}
}

// propagateTrxToAccounts pushes given transaction to accounts on both sides.
// The function is executed in a separate thread so it doesn't block
func (td *trxDispatcher) propagateTrxToAccounts(evt *eventTransaction, wg *sync.WaitGroup) {
	// log what we do here
	td.log.Debugf("analyzing accounts involved in trx %s", evt.trx.Hash.String())

	// the sender is always present
	wg.Add(1)
	td.repo.QueueAccount(evt.block, evt.trx, &evt.trx.From, nil, wg)

	// do we have a recipient?
	if evt.trx.To != nil {
		wg.Add(1)
		td.repo.QueueAccount(evt.block, evt.trx, evt.trx.To, nil, wg)
	}

	// no contract creation? we are done
	if evt.trx.ContractAddress == nil {
		return
	}

	// queue the new contract to be processed as well
	td.log.Debugf("contract %s found at trx %s", evt.trx.ContractAddress.String(), evt.trx.Hash.String())
	wg.Add(1)
	td.repo.QueueAccount(evt.block, evt.trx, evt.trx.ContractAddress, &evt.trx.Hash, wg)
	return
}

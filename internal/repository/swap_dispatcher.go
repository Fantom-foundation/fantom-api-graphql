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

// swapDispatcher implements dispatcher of swap events in the blockchain.
type swapDispatcher struct {
	service
	buffer chan *evtSwap
}

// evtSwap represents a single incoming swap event to be processed.
type evtSwap struct {
	swp *types.Swap
}

// newSwapDispatcher creates a new swap dispatcher instance.
func newSwapDispatcher(buffer chan *evtSwap, repo Repository, log logger.Logger, wg *sync.WaitGroup) *swapDispatcher {
	// create new dispatcher
	return &swapDispatcher{
		service: newService("swapDispatcher", repo, log, wg),
		buffer:  buffer,
	}
}

// run starts the transaction dispatcher job
func (sd *swapDispatcher) run() {
	// inform about the action
	sd.log.Notice("starting uniswap dispatcher")

	// add self to the wait group and run the dispatch routine
	sd.wg.Add(1)
	go sd.dispatch()
}

// dispatch implements the dispatcher reader and router routine.
func (sd *swapDispatcher) dispatch() {
	// don't forget to sign off after we are done
	defer func() {
		// log finish
		sd.log.Notice("swap dispatcher done")
		sd.wg.Done()
	}()

	// what we dispatch
	var toDispatch *evtSwap

	// wait for swap and process it
	for {
		// try to read next swap
		select {
		case toDispatch = <-sd.buffer:
			// validate
			if toDispatch.swp == nil {
				sd.log.Critical("swap dispatcher received invalid data")
				continue
			}

			// dispatch the received
			err := sd.repo.UniswapAdd(toDispatch.swp)
			if err != nil {
				sd.log.Error("could not dispatch swap")
				sd.log.Error(err)
			}
		case <-sd.sigStop:
			// stop the routine
			return
		}
	}
}

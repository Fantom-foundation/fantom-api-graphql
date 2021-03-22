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
	buffer chan *types.Swap
}

// newSwapDispatcher creates a new swap dispatcher instance.
func newSwapDispatcher(buffer chan *types.Swap, repo Repository, log logger.Logger, wg *sync.WaitGroup) *swapDispatcher {
	// create new dispatcher
	return &swapDispatcher{
		service: newService("swap dispatcher", repo, log, wg),
		buffer:  buffer,
	}
}

// run starts the transaction dispatcher job
func (swd *swapDispatcher) run() {
	// add self to the wait group and run the dispatch routine
	swd.wg.Add(1)
	go swd.dispatch()
}

// dispatch implements the dispatcher reader and router routine.
func (swd *swapDispatcher) dispatch() {
	// log action
	swd.log.Notice("uniswap dispatcher is running")

	// don't forget to sign off after we are done
	defer func() {
		// log finish
		swd.log.Notice("uniswap dispatcher is closed")
		swd.wg.Done()
	}()

	// wait for swap and process it
	for {
		// try to read next swap
		select {
		case toDispatch := <-swd.buffer:
			// validate
			if toDispatch == nil {
				swd.log.Critical("swap dispatcher received invalid data")
				continue
			}

			// dispatch the received
			err := swd.repo.UniswapAdd(toDispatch)
			if err != nil {
				swd.log.Error("could not dispatch swap")
				swd.log.Error(err)
			}
		case <-swd.sigStop:
			// stop the routine
			return
		}
	}
}

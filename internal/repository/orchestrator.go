/*
Repository package implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/logger"
	"sync"
)

// trxDispatchBufferCapacity is the number of transactions kept in the dispatch buffer.
const trxDispatchBufferCapacity = 20000

// Orchestrator implements repository synchronization and monitoring control
type orchestrator struct {
	service

	// orchestrator managed channels
	trxBuffer chan *evtTransaction
	sysDone   chan bool

	// services being orchestrated
	txd *trxDispatcher
	sys *scanner
	mon *blockMonitor
}

// NewOrchestrator creates a new instance of repository orchestrator.
func newOrchestrator(repo Repository, log logger.Logger) *orchestrator {
	// make a wait group for orchestrated services
	var wg sync.WaitGroup

	// create new orchestrator
	or := orchestrator{
		service: newService("orchestrator", repo, log, &wg),
	}

	// init the orchestration
	or.init()

	// start orchestrating
	wg.Add(1)
	go or.orchestrate()
	return &or
}

// Close signals orchestrator to terminate all orchestrated services.
func (or *orchestrator) close() {
	// signal the service to close
	or.service.close()

	// signal scanner
	if or.sys != nil {
		or.sys.close()
	}

	// signal monitor
	if or.mon != nil {
		or.mon.close()
	}

	// signal tx dispatcher
	if or.txd != nil {
		or.txd.close()
	}

	// wait scanners to terminate
	or.log.Debugf("waiting for services to finish")
	or.wg.Wait()

	// close owned channels
	close(or.trxBuffer)
	close(or.sysDone)

	// we are done
	or.log.Notice("orchestrator done")
}

// Start initiates the orchestrator work.
func (or *orchestrator) init() {
	// create a channel for transaction dispatcher
	or.trxBuffer = make(chan *evtTransaction, trxDispatchBufferCapacity)

	// make the transaction dispatcher; it starts dispatching immediately
	or.txd = NewTrxDispatcher(or.trxBuffer, or.repo, or.log, or.wg)

	// create sync scanner; it starts scanning immediately
	or.sysDone = make(chan bool, 1)
	or.sys = newScanner(or.trxBuffer, or.sysDone, or.repo, or.log, or.wg)

	// create monitor; it waits for sync scanner to finish
	or.mon = NewBlockMonitor(or.repo.FtmConnection(), or.trxBuffer, or.repo, or.log, or.wg)
}

// Orchestrate starts the service orchestration.
func (or *orchestrator) orchestrate() {
	// log action
	or.log.Notice("orchestrator is running")

	// don't forget to sign off after we are done
	defer func() {
		// log finish
		or.log.Notice("orchestrator scheduler is closing")

		// signal to wait group we are done
		or.wg.Done()
	}()

	// wait for either stop signal, or scanner to finish
	for {
		select {
		case <-or.sigStop:
			// stop signal received?
			return
		case <-or.sysDone:
			// log action
			or.log.Notice("synchronization finished")

			// scanner is done, start monitoring
			or.mon.run()
		}
	}
}

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
	"time"
)

// trxDispatchBufferCapacity is the number of transactions kept in the dispatch buffer.
const trxDispatchBufferCapacity = 20000

// Orchestrator implements repository synchronization and monitoring control
type orchestrator struct {
	service

	// orchestrator managed channels
	trxBuffer        chan *evtTransaction
	sysDone          chan bool
	reScan           chan bool
	sigKillScheduler chan bool

	// count re-scans
	reScanCounter uint

	// services being orchestrated
	txd *trxDispatcher
	sys *scanner
	mon *blockMonitor
	sti *stiMonitor
}

// NewOrchestrator creates a new instance of repository orchestrator.
func newOrchestrator(repo Repository, log logger.Logger) *orchestrator {
	// make a wait group for orchestrated services
	var wg sync.WaitGroup

	// create new orchestrator
	or := orchestrator{
		service:          newService("orchestrator", repo, log, &wg),
		sigKillScheduler: make(chan bool, 1),
	}

	// init the orchestration
	or.init()

	// start orchestrating
	wg.Add(1)
	go or.orchestrate()
	return &or
}

// close signals orchestrator to terminate all orchestrated services.
func (or *orchestrator) close() {
	// signal the service to close
	or.service.close()

	// signal scanner
	or.sys.close()

	// signal monitor
	or.mon.close()

	// signal tx dispatcher
	or.txd.close()

	// signal sti monitor
	or.sti.close()

	// kill re-scan scheduler
	or.sigKillScheduler <- true

	// wait scanners to terminate
	or.log.Debugf("waiting for services to finish")
	or.wg.Wait()

	// close owned channels
	close(or.trxBuffer)
	close(or.sysDone)
	close(or.sigKillScheduler)

	// we are done
	or.log.Notice("orchestrator done")
}

// setBlockChannel registers a channel for notifying new block events.
func (or *orchestrator) setBlockChannel(ch chan *types.Block) {
	or.mon.onBlock = ch
}

// setTrxChannel registers a channel for notifying new transaction events.
func (or *orchestrator) setTrxChannel(ch chan *types.Transaction) {
	or.mon.onTransaction = ch
}

// init initiates the orchestrator work.
func (or *orchestrator) init() {
	// create a channel for transaction dispatcher
	or.trxBuffer = make(chan *evtTransaction, trxDispatchBufferCapacity)

	// make the transaction dispatcher; it starts dispatching immediately
	or.txd = NewTrxDispatcher(or.trxBuffer, or.repo, or.log, or.wg)

	// create sync scanner; it starts scanning immediately
	or.sysDone = make(chan bool, 1)
	or.sys = newScanner(or.trxBuffer, or.sysDone, or.repo, or.log, or.wg)

	// create block monitor; it waits for sync scanner to finish
	or.reScan = make(chan bool, 1)
	or.mon = NewBlockMonitor(or.repo.FtmConnection(), or.trxBuffer, or.reScan, or.repo, or.log, or.wg)

	// create staker information monitor; it starts right away on slow peace
	or.sti = newStiMonitor(or.repo, or.log, or.wg)
}

// orchestrate starts the service orchestration.
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
		case <-or.reScan:
			// advance counter
			or.reScanCounter++

			// log action
			or.log.Warningf("re-scan #%d requested by terminated monitoring", or.reScanCounter)

			// start re-scan scheduler
			or.wg.Add(1)
			go or.scheduleRescan()
		}
	}
}

// scheduleRescan schedules block chain re-scan on monitoring failure.
func (or *orchestrator) scheduleRescan() {
	// don't forget to sign off after we are done
	defer func() {
		// log finish
		or.log.Notice("orchestrator re-scan scheduler is done")

		// signal to wait group we are done
		or.wg.Done()
	}()

	// calculate delay duration of this re-scan
	// we increase delay between re-scans so we don't consume too much resources
	// if the Lachesis is dropping subscriptions but is still available for RPC calls
	var dur = time.Duration(or.reScanCounter*2) * time.Second
	or.log.Warningf("re-scan scheduled after %d seconds", dur)

	// wait for either stop signal, or scanner to finish
	for {
		select {
		case <-or.sigKillScheduler:
			return
		case <-time.After(dur):
			or.sys.run()
			return
		}
	}
}

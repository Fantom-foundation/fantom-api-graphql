/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/types"
	"sync"
	"time"
)

// trxDispatchBufferCapacity is the number of transactions kept in the dispatch buffer.
const trxDispatchBufferCapacity = 20000
const swapDispatchBufferCapacity = 20000

// Orchestrator implements repository synchronization and monitoring control
type orchestrator struct {
	service

	// orchestrator managed channels
	trxBuffer          chan *evtTransaction
	swapBuffer         chan *evtSwap
	accountQueue       chan *accountQueueRequest
	contractCallsQueue chan *types.Transaction
	sysDone            chan bool
	swsDone            chan bool
	reScan             chan bool
	sigKillScheduler   chan bool

	// count re-scans
	reScanCounter uint

	// services being orchestrated
	txd *trxDispatcher
	sys *scanner
	sws *uniswapScanner
	swd *swapDispatcher
	smo *UniswapMonitor
	mon *blockMonitor
	sti *stiMonitor
	acq *accountQueue
	ccq *contractCallQueue
}

// NewOrchestrator creates a new instance of repository orchestrator.
func newOrchestrator(repo Repository, log logger.Logger, cfg *config.Repository) *orchestrator {
	// make a wait group for orchestrated services
	var wg sync.WaitGroup

	// create new orchestrator
	or := orchestrator{
		service:          newService("orchestrator", repo, log, &wg),
		sigKillScheduler: make(chan bool, 1),
	}

	// init the orchestration
	or.init(cfg)

	// return the orchestrator
	return &or
}

// init initiates the orchestrator work.
func (or *orchestrator) init(cfg *config.Repository) {
	// create a channel for transaction and swap dispatcher
	or.trxBuffer = make(chan *evtTransaction, trxDispatchBufferCapacity)
	or.swapBuffer = make(chan *evtSwap, swapDispatchBufferCapacity)
	or.accountQueue = make(chan *accountQueueRequest, accountQueueLength)
	or.contractCallsQueue = make(chan *types.Transaction, contractCallQueueLength)

	// make the transaction dispatcher; it starts dispatching immediately
	or.txd = newTrxDispatcher(or.trxBuffer, or.repo, or.log, or.wg)

	// make the swap dispatcher; it starts dispatching immediately
	or.swd = newSwapDispatcher(or.swapBuffer, or.repo, or.log, or.wg)

	// make the account queue handler; it starts processing immediately
	or.acq = newAccountQueue(or.accountQueue, or.repo, or.log, or.wg)

	// make the contract call analyzer queue handler; it starts processing immediately
	or.ccq = newContractCallQueue(or.contractCallsQueue, or.repo, or.log, or.wg)

	// create sync scanner; it starts scanning immediately
	or.sysDone = make(chan bool, 1)
	or.sys = newScanner(or.trxBuffer, or.sysDone, or.repo, or.log, or.wg)

	// create swap scanner, which loads uniswaps to local db immediately
	or.swsDone = make(chan bool, 1)
	or.sws = newUniswapScanner(or.swapBuffer, or.swsDone, or.repo, or.log, or.wg)

	// create block monitor; it waits for sync scanner to finish
	or.reScan = make(chan bool, 1)
	or.mon = NewBlockMonitor(or.repo.FtmConnection(), or.trxBuffer, or.reScan, or.repo, or.log, or.wg)

	// create the Uniswap monitor
	or.smo = NewUniswapMonitor(or.repo.FtmConnection(), or.swapBuffer, or.repo, or.log, or.wg)

	// create staker information monitor; it starts right away on slow peace
	if cfg.MonitorStakers {
		or.sti = newStiMonitor(or.repo, or.log, or.wg)
	}
}

// run starts the orchestrator work
func (or *orchestrator) run() {
	// start all the services in the expected order
	// queue dispatchers come first
	or.txd.run()
	or.swd.run()
	or.acq.run()
	or.ccq.run()

	// now the scanners so we sync the off-chain database
	or.sys.run()
	or.sws.run()

	// finally monitors
	or.mon.run()
	or.smo.run()

	// most of them are fired separately, once the sync is done, though
	if or.sti != nil {
		or.sti.run()
	}

	// start orchestrating
	or.wg.Add(1)
	go or.orchestrate()
}

// close signals orchestrator to terminate all orchestrated services.
func (or *orchestrator) close() {
	// close all the services
	or.closeServices()

	// kill re-scan scheduler
	or.sigKillScheduler <- true

	// wait scanners to terminate
	or.log.Debugf("waiting for services to finish")
	or.wg.Wait()

	// close owned channels
	close(or.trxBuffer)
	close(or.swapBuffer)
	close(or.swsDone)
	close(or.sysDone)
	close(or.sigKillScheduler)

	// we are done
	or.log.Notice("orchestrator done")
}

// closeServices signals services of the orchestrator to close
func (or *orchestrator) closeServices() {
	// signal the orchestrator to close
	or.service.close()

	// signal the services to close
	or.sys.close()
	or.sws.close()
	or.swd.close()
	or.smo.close()
	or.mon.close()
	or.txd.close()
	or.acq.close()
	or.ccq.close()

	// signal sti monitor if it exists
	if or.sti != nil {
		or.sti.close()
	}
}

// setBlockChannel registers a channel for notifying new block events.
func (or *orchestrator) setBlockChannel(ch chan *types.Block) {
	or.mon.onBlock = ch
}

// setTrxChannel registers a channel for notifying new transaction events.
func (or *orchestrator) setTrxChannel(ch chan *types.Transaction) {
	or.mon.onTransaction = ch
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
			or.log.Notice("blocks synchronization finished")
			// scanner is done, start monitoring
			or.mon.run()
		case <-or.swsDone:
			or.log.Notice("swaps synchronization finished")
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
			or.sws.run()
			return
		}
	}
}

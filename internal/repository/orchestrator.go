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

// swapDispatchQueueCapacity is the number of Uniswap swap kept in the dispatch buffer.
const swapDispatchQueueCapacity = 20000

// orScannersCount is the number of scanner services running concurrently.
const orScannersCount = 2

// Orchestrator implements repository synchronization and monitoring control
type orchestrator struct {
	service

	// orchestrator managed channels
	trxDispatcherQueue chan *eventTransaction
	accountQueue       chan *accountEvent
	logsQueue          chan *types.LogRecord

	// orchestration related events
	blkScanDone   chan bool
	reScan        chan bool
	reScanSigStop chan bool

	// re-scans counter
	reScanCounter uint

	// dispatchers handle objects put into a corresponding queue
	// and process them into the database, and/or any other way needed
	txd *trxDispatcher
	acd *accountDispatcher
	lod *logsDispatcher

	// scanners synchronize content of the database with the current
	// chain state to catch up after the API service restarted
	bls *blockScanner
	sfs *sfcScanner

	// monitors watch for newly created objects on the block chain
	// and push those object into their processing queues to be sorted out
	blm *blockMonitor
	stm *stiMonitor
	txf *txFlowUpdater
	gps *gasPriceSuggestionMonitor
}

// NewOrchestrator creates a new instance of repository orchestrator.
func newOrchestrator(repo Repository, log logger.Logger, cfg *config.Config) *orchestrator {
	// make a wait group for orchestrated services
	var wg sync.WaitGroup

	// create new orchestrator
	or := orchestrator{
		// make the service
		service: newService("orchestrator", repo, log, &wg),

		// make queues for dispatchers; scanners and monitors will use them to push new objects
		trxDispatcherQueue: make(chan *eventTransaction, trxDispatchQueueCapacity),
		accountQueue:       make(chan *accountEvent, accountQueueLength),
		logsQueue:          make(chan *types.LogRecord, logQueueLength),

		// make sure to accept all the incoming signals
		reScanSigStop: make(chan bool, 1),
	}

	// init the orchestration
	or.init(cfg)
	return &or
}

// init initiates the orchestrator work.
func (or *orchestrator) init(cfg *config.Config) {
	// make all the dispatchers first so they can receive and process objects
	or.gps = newGasPriceSuggestionMonitor(or.repo, or.log, or.wg)
	or.txd = newTrxDispatcher(or.trxDispatcherQueue, or.repo, or.log, or.wg)
	or.acd = newAccountDispatcher(or.accountQueue, or.repo, or.log, or.wg)
	or.lod = newLogsDispatcher(or.logsQueue, or.repo, or.log, or.wg)

	// create sync blockScanner; it starts scanning immediately
	or.blkScanDone = make(chan bool, 1)
	or.bls = newBlockScanner(or.trxDispatcherQueue, or.blkScanDone, or.repo, or.log, or.wg, &cfg.RepoCommand)

	// SFC scanner
	or.sfs = newSFCScanner(or.repo, or.log, or.wg)

	// create block monitor; it waits for sync blockScanner to finish
	or.reScan = make(chan bool, orScannersCount)
	or.blm = NewBlockMonitor(or.repo.FtmConnection(), or.trxDispatcherQueue, or.reScan, or.repo, or.log, or.wg)

	// create staker information monitor; it starts right away on slow peace
	if cfg.Repository.MonitorStakers {
		or.stm = newStiMonitor(or.repo, or.log, or.wg)
	}

	// create trx flow updater
	or.txf = NewTxFlowUpdater(or.repo, or.log, or.wg)
}

// run starts the orchestrator work
func (or *orchestrator) run() {
	// queue dispatchers come first so they cna process new objects
	or.txd.run()
	or.acd.run()
	or.lod.run()

	// now the scanners so we sync the off-chain database
	or.bls.run()
	or.sfs.run()

	// finally monitors
	or.txf.run()
	or.gps.run()

	// stakers info monitor may not be run at all
	if or.stm != nil {
		or.stm.run()
	}

	// start orchestrating
	or.wg.Add(1)
	go or.orchestrate()
}

// close signals orchestrator to terminate all orchestrated services.
func (or *orchestrator) close() {
	// close all the services
	or.closeServices()

	// wait scanners to terminate
	or.log.Debugf("waiting for services to finish")
	or.wg.Wait()

	// close owned channels
	close(or.trxDispatcherQueue)
	close(or.blkScanDone)
	close(or.reScanSigStop)

	// we are done
	or.log.Notice("orchestrator done")
}

// closeServices signals services of the orchestrator to close
func (or *orchestrator) closeServices() {
	// we are done
	or.log.Notice("closing services")

	// signal the orchestrator to close
	or.reScanSigStop <- true
	or.service.close()

	// signal monitors to close
	or.blm.close()
	or.txf.close()
	or.gps.close()

	// signal scanners to close
	or.bls.close()
	or.sfs.close()

	// signal close to sti monitor as well, if it exists
	if or.stm != nil {
		or.stm.close()
	}

	// signal dispatchers to close
	or.txd.close()
	or.acd.close()
	or.lod.close()
}

// setBlockChannel registers a channel for notifying new block events.
func (or *orchestrator) setBlockChannel(ch chan *types.Block) {
	or.blm.onBlock = ch
}

// setTrxChannel registers a channel for notifying new transaction events.
func (or *orchestrator) setTrxChannel(ch chan *types.Transaction) {
	or.blm.onTransaction = ch
}

// orchestrate starts the service orchestration.
func (or *orchestrator) orchestrate() {
	// log action
	or.log.Notice("orchestrator is running")

	// don't forget to sign off after we are done
	defer func() {
		// log finish
		or.log.Notice("orchestrator is closed")

		// signal to wait group we are done
		or.wg.Done()
	}()

	// wait for either stop signal, or blockScanner to finish
	for {
		select {
		case <-or.sigStop:
			return

		case <-or.blkScanDone:
			or.log.Notice("blocks synchronization finished")
			or.blm.run()

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

// scheduleRescan schedules blockchain re-scan on monitoring failure.
func (or *orchestrator) scheduleRescan() {
	// inform
	or.log.Notice("re-scan scheduler is running")
	or.blm.close()

	// don't forget to sign off after we are done
	defer func() {
		// log finish
		or.log.Notice("re-scan scheduler is closed")

		// signal to wait group we are done
		or.wg.Done()
	}()

	// calculate delay duration of this re-scan
	// we increase delay between re-scans so we don't consume too much resources
	// if the Lachesis is dropping subscriptions but is still available for RPC calls
	var dur = time.Duration(or.reScanCounter*10) * time.Second
	or.log.Warningf("re-scan scheduled after %d seconds", dur)

	// wait for either stop signal, or blockScanner to finish
	for {
		select {
		case <-or.reScanSigStop:
			return
		case <-time.After(dur):
			or.bls.run()
			return
		}
	}
}

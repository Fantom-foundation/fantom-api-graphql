// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"sync"
)

// swapDispatchQueueCapacity is the number of Uniswap swap kept in the dispatch buffer.
const swapDispatchQueueCapacity = 20000

// Orchestrator implements service orchestration.
type Orchestrator struct {
	log logger.Logger
	wg  *sync.WaitGroup

	// special services with external dependency
	bld *blockDispatcher
	trd *trxDispatcher
	acd *accDispatcher
	lgd *logDispatcher

	// collection of all the managed services
	svc []Svc
}

// NewOrchestrator creates a new instance of service orchestrator.
func NewOrchestrator(cfg *config.Config) *Orchestrator {
	// create new orchestrator
	or := Orchestrator{
		log: log,
		wg:  new(sync.WaitGroup),
		svc: make([]Svc, 0, 10),
	}

	// init the orchestration
	or.init(cfg)
	return &or
}

// Run starts all the services prepared to be run.
func (or *Orchestrator) Run() {
	// ready services to be started
	or.ready()

	// start services
	for _, s := range or.svc {
		s.run()
	}
}

// Close signals orchestrator to terminate all orchestrated services.
func (or *Orchestrator) Close() {
	or.log.Noticef("orchestrator received a close signal")

	// pass the signal to all the services
	for _, s := range or.svc {
		or.log.Noticef("closing %s", s.name())
		s.close()
	}

	// wait scanners to terminate
	or.log.Notice("waiting for services to finish")
	or.wg.Wait()

	// we are done
	or.log.Notice("orchestrator done")
}

// SetBlockChannel registers a channel for notifying new block events.
func (or *Orchestrator) SetBlockChannel(ch chan *types.Block) {
	or.bld.onBlock = ch
}

// SetTrxChannel registers a channel for notifying new transaction events.
func (or *Orchestrator) SetTrxChannel(ch chan *types.Transaction) {
	or.trd.onTransaction = ch
}

// Init the orchestrator.
func (or *Orchestrator) init(cfg *config.Config) {
	// make the block dispatcher
	or.bld = &blockDispatcher{or: or}
	or.svc = append(or.svc, or.bld)

	// make the transaction dispatcher
	or.trd = &trxDispatcher{or: or}
	or.svc = append(or.svc, or.trd)

	// make account dispatcher
	or.acd = &accDispatcher{or: or}
	or.svc = append(or.svc, or.acd)

	// make log dispatcher
	or.lgd = &logDispatcher{or: or}
	or.svc = append(or.svc, or.lgd)
}

// initServices initializes all the services added to the orchestrator
func (or *Orchestrator) ready() {
	// init all the services
	for _, s := range or.svc {
		s.init()
	}

	// connect services' input channels to their source
	or.trd.inTransaction = or.bld.outTransaction
	or.acd.inAccount = or.trd.outAccount
	or.lgd.inLog = or.trd.outLog
}

// started signals to the orchestrator that the calling service
// has been started and is functioning.
func (or *Orchestrator) started(svc Svc) {
	or.wg.Add(1)
	or.log.Noticef("%s is running", svc.name())
}

// finished signals to the orchestrator that the calling service
// has been terminated and is no longer running.
func (or *Orchestrator) finished(svc Svc) {
	or.wg.Done()
	or.log.Noticef("%s terminated", svc.name())
}

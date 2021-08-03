// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"sync"
)

// Orchestrator implements service orchestration.
type Orchestrator struct {
	wg  *sync.WaitGroup

	// special services with external dependency
	bld *blockDispatcher
	trd *trxDispatcher
	acd *accDispatcher
	lgd *logDispatcher

	// collection of all the managed services
	svc []Svc
}

// repo represents the repository used by services to handle data exchange.
var repo repository.Repository

// log represents the logger used by services to inform about performed actions.
var log logger.Logger

// NewOrchestrator creates a new instance of service orchestrator.
func NewOrchestrator(cfg *config.Config, r repository.Repository) *Orchestrator {
	// create new orchestrator
	or := Orchestrator{
		wg:  new(sync.WaitGroup),
		svc: make([]Svc, 0, 10),
	}

	// update the package references
	repo = r
	log = r.Log()

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
	log.Noticef("orchestrator received a close signal")

	// pass the signal to all the services
	for _, s := range or.svc {
		log.Noticef("closing %s", s.name())
		s.close()
	}

	// wait scanners to terminate
	log.Notice("waiting for services to finish")
	or.wg.Wait()

	// we are done
	log.Notice("orchestrator done")
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
	log.Noticef("%s is running", svc.name())
}

// finished signals to the orchestrator that the calling service
// has been terminated and is no longer running.
func (or *Orchestrator) finished(svc Svc) {
	or.wg.Done()
	log.Noticef("%s terminated", svc.name())
}

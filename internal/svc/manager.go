// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"fmt"
	"sync"
)

// ServiceManager implements service manager.
type ServiceManager struct {
	wg *sync.WaitGroup

	// special services with external dependency
	ora *orchestrator
	bld *blockDispatcher
	trd *trxDispatcher
	acd *accDispatcher
	lgd *logDispatcher
	bls *blkScanner
	bud *burnDispatcher

	// collection of all the managed services
	svc []Svc
}

// newServiceManager creates a new instance of service manager.
func newServiceManager() *ServiceManager {
	// make sure we have what we need
	if log == nil {
		panic(fmt.Errorf("logger not available"))
	}
	if cfg == nil {
		panic(fmt.Errorf("configuration not available"))
	}

	// create new orchestrator
	sm := ServiceManager{
		wg:  new(sync.WaitGroup),
		svc: make([]Svc, 0, 15),
	}

	// init the orchestration
	sm.init()
	return &sm
}

// Run starts all the services prepared to be run.
func (mgr *ServiceManager) Run() {
	// get local copy of the repository
	repo = repository.R()

	// init all the services to the starting state
	for _, s := range mgr.svc {
		s.init()
	}

	// start services
	for _, s := range mgr.svc {
		s.run()
	}
}

// Close signals orchestrator to terminate all orchestrated services.
func (mgr *ServiceManager) Close() {
	log.Noticef("svc manager received a close signal")

	// pass the signal to all the services
	for _, s := range mgr.svc {
		log.Noticef("closing %s", s.name())
		s.close()
	}

	// wait scanners to terminate
	log.Notice("waiting for services to finish")
	mgr.wg.Wait()

	// we are done
	log.Notice("svc manager closed")
}

// SetBlockChannel registers a channel for notifying new block events.
func (mgr *ServiceManager) SetBlockChannel(ch chan *types.Block) {
	mgr.bld.onBlock = ch
}

// SetTrxChannel registers a channel for notifying new transaction events.
func (mgr *ServiceManager) SetTrxChannel(ch chan *types.Transaction) {
	mgr.trd.onTransaction = ch
}

// Init the svc manager.
func (mgr *ServiceManager) init() {
	// make the block dispatcher
	mgr.bld = &blockDispatcher{service: service{mgr: mgr}}
	mgr.svc = append(mgr.svc, mgr.bld)

	// make the transaction dispatcher
	mgr.trd = &trxDispatcher{service: service{mgr: mgr}}
	mgr.svc = append(mgr.svc, mgr.trd)

	// make account dispatcher
	mgr.acd = &accDispatcher{service: service{mgr: mgr}}
	mgr.svc = append(mgr.svc, mgr.acd)

	// make log dispatcher
	mgr.lgd = &logDispatcher{service: service{mgr: mgr}}
	mgr.svc = append(mgr.svc, mgr.lgd)

	// make burn dispatcher
	mgr.bud = &burnDispatcher{service: service{mgr: mgr}}
	mgr.svc = append(mgr.svc, mgr.bud)

	// make block scanner
	mgr.bls = &blkScanner{service: service{mgr: mgr}, cfg: cfg.RepoCommand}
	mgr.svc = append(mgr.svc, mgr.bls)

	// make epoch scanner
	mgr.svc = append(mgr.svc, &epochScanner{service: service{mgr: mgr}})

	// make staker information scanner only if we have the contract address
	if cfg.Staking.StiContract.String() != config.EmptyAddress {
		mgr.svc = append(mgr.svc, &stiScanner{service: service{mgr: mgr}})
	}

	// make gas price suggestion monitor
	mgr.svc = append(mgr.svc, &gpsMonitor{service: service{mgr: mgr}})

	// make transaction flow monitor
	mgr.svc = append(mgr.svc, &trxFlowMonitor{service: service{mgr: mgr}})

	// make the network discovery
	mgr.svc = append(mgr.svc, &netCrawler{service: service{mgr: mgr}})

	// add orchestrator as the last service, so it can safely operate on all the other
	mgr.ora = &orchestrator{service: service{mgr: mgr}}
	mgr.svc = append(mgr.svc, mgr.ora)
}

// started signals to the manager that the calling service
// has been started and is functioning.
func (mgr *ServiceManager) started(svc Svc) {
	mgr.wg.Add(1)
	log.Noticef("%s is running", svc.name())
}

// finished signals to the manager that the calling service
// has been terminated and is no longer running.
func (mgr *ServiceManager) finished(svc Svc) {
	mgr.wg.Done()
	log.Noticef("%s terminated", svc.name())
}

// BlockHeight provides identifier of the top known block.
func (mgr *ServiceManager) BlockHeight() uint64 {
	if mgr.bls == nil {
		return 0
	}
	return mgr.bls.blockHeight()
}

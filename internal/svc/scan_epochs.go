// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"time"
)

// epsScanTickerDuration represents the frequency of the scanner ticker.
const epsScanTickerDuration = 25 * time.Millisecond

// epsObserverTickerDuration represents the frequency of the sealed epoch observer.
const epsObserverTickerDuration = 10 * time.Second

// epsStoreQueueLength represents the capacity of the epoch scanner store queue.
const epsStoreQueueLength = 100

// epochScanner implements blockchain epochs scanner.
// We can not scan anything before epoch #5574 (migration); full data available after #5577
type epochScanner struct {
	service
	observeTick *time.Ticker
	scanTick    *time.Ticker
	current     uint64
	top         *types.Epoch
	queue       chan *types.Epoch
}

// name returns the name of the service used by orchestrator.
func (eps *epochScanner) name() string {
	return "SFC epoch scanner"
}

// init prepares the epoch scanner sig channel and storage queue.
func (eps *epochScanner) init() {
	eps.sigStop = make(chan struct{})
	eps.queue = make(chan *types.Epoch, epsStoreQueueLength)
}

// run feeds initial state and starts the epoch scanner threads.
func (eps *epochScanner) run() {
	// make sure we are orchestrated
	if eps.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", eps.name()))
	}

	// try to find the last known epoch, we want to start where we left off
	var err error
	eps.current, err = repo.LastKnownEpoch()
	if err != nil {
		log.Criticalf("can not get the last known epoch; %s", err.Error())
		eps.current = 1
	}

	// signal orchestrator that we start two threads
	eps.mgr.started(eps)
	go eps.dequeue()
	eps.mgr.started(eps)
	go eps.execute()
}

// close terminates the block dispatcher.
func (eps *epochScanner) close() {
	if eps.observeTick != nil {
		eps.observeTick.Stop()
		eps.scanTick.Stop()
	}

	if eps.sigStop != nil {
		close(eps.sigStop)
	}
}

// execute performs the epoch scan starting from the given epoch ID.
func (eps *epochScanner) execute() {
	// make sure to clean up
	defer func() {
		close(eps.queue)
		eps.mgr.finished(eps)
	}()

	// init tickers
	eps.observeTick = time.NewTicker(epsObserverTickerDuration)
	eps.scanTick = time.NewTicker(epsScanTickerDuration)

	// log the status
	log.Noticef("epoch scan starts at #%d", eps.current)

	for {
		select {
		case <-eps.sigStop:
			return
		case <-eps.scanTick.C:
			eps.next()
		case <-eps.observeTick.C:
			if eps.top != nil && uint64(eps.top.Id) > eps.current {
				log.Noticef("epoch scan at #%d of #%d", eps.current, eps.top.Id)
			}
			eps.observe()
		}
	}
}

// updateState updates observed
func (eps *epochScanner) observe() {
	// what is the top epoch number
	ep, err := repo.CurrentSealedEpoch()
	if err != nil {
		log.Criticalf("can not get sealed epoch; %s", err.Error())
		return
	}

	// still on the same epoch
	if ep == nil || (eps.top != nil && eps.top.Id == ep.Id) {
		return
	}

	// a new epoch found
	log.Noticef("current sealed epoch is #%d; scanner at #%d", ep.Id, eps.current)
	eps.top = ep
}

// next processes epoch data based on the stored current epoch number.
func (eps *epochScanner) next() {
	// do we have a space to grow to
	if eps.top == nil {
		log.Debugf("top sealed epoch not known")
		return
	}

	// already reached top epoch
	if uint64(eps.top.Id) < eps.current {
		log.Debugf("already on top epoch #%d", eps.current)
		return
	}

	// try to get the epoch data
	ep, err := repo.Epoch((*hexutil.Uint64)(&eps.current))
	if err != nil {
		log.Errorf("can not get epoch #%d; %s", eps.current, err.Error())
		return
	}

	// process and move to the next epoch
	log.Debugf("loaded epoch #%d", ep.Id)
	select {
	case eps.queue <- ep:
		eps.current++
	case <-eps.sigStop:
	}
}

// dequeue consumes the queue and sends epochs to the persistent storage.
func (eps *epochScanner) dequeue() {
	for {
		// pull the next epoch from the queue if possible and store it
		ep, ok := <-eps.queue
		if !ok {
			log.Noticef("SFC epoch scanner queue closed")
			eps.mgr.finished(eps)
			return
		}
		eps.store(ep)
	}
}

// store consumes the queue and sends epochs to the persistent storage.
func (eps *epochScanner) store(ep *types.Epoch) {
	// log the epoch
	if ep.EndTime == 0 {
		log.Warningf("epoch #%d details not available", eps.current)
		return
	}

	// log what we do
	log.Debugf("processing epoch #%d", ep.Id)

	// add the epoch to the database
	err := repo.AddEpoch(ep)
	if err != nil {
		log.Errorf("can not store epoch #%d; %s", ep.Id, err.Error())
	}
}

// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/atomic"
	"time"
)

// epsScanTickerDuration represents the frequency of the scanner ticker.
const epsScanTickerDuration = 100 * time.Millisecond

// epsObserverTickerDuration represents the frequency of the sealed epoch observer.
const epsObserverTickerDuration = 10 * time.Second

// epsStoreQueueLength represents the capacity of the epoch scanner store queue.
const epsStoreQueueLength = 100

// epochScanner implements blockchain epochs scanner.
// We can not scan anything before epoch #5574 (migration); full data available after #5577
type epochScanner struct {
	or          *Orchestrator
	sigStop     chan bool
	observeTick *time.Ticker
	scanTick    *time.Ticker
	current     *atomic.Uint64
	top         *types.Epoch
	queue       chan *types.Epoch
}

// name returns the name of the service.
func (eps *epochScanner) name() string {
	return "SFC epoch scanner"
}

// init prepares the epoch scanner to perform its function.
func (eps *epochScanner) init() {
	eps.sigStop = make(chan bool, 1)
	eps.queue = make(chan *types.Epoch, epsStoreQueueLength)
}

// run starts the account queue to life.
func (eps *epochScanner) run() {
	// make sure we are orchestrated
	if eps.or == nil {
		panic(fmt.Errorf("no orchestrator set"))
	}

	// try to find the lats known epoch
	cur, err := repo.LastKnownEpoch()
	if err != nil {
		log.Criticalf("can not get the last known epoch; %s", err.Error())
		cur = 1
	}

	// store the current epoch ID
	eps.current = atomic.NewUint64(cur)

	// signal orchestrator we started and go
	eps.or.started(eps)
	go eps.dequeue()
	go eps.scan()
}

// close terminates the block dispatcher.
func (eps *epochScanner) close() {
	eps.observeTick.Stop()
	eps.scanTick.Stop()

	eps.sigStop <- true
}

// scan performs the epoch scan starting from the given epoch ID.
func (eps *epochScanner) scan() {
	// make sure to clean up
	defer func() {
		close(eps.sigStop)
		close(eps.queue)

		eps.or.finished(eps)
	}()

	// init tickers
	eps.observeTick = time.NewTicker(epsObserverTickerDuration)
	eps.scanTick = time.NewTicker(epsScanTickerDuration)

	// log the status
	log.Noticef("epoch scan starts at #%d", eps.current.Load())

	for {
		select {
		case <-eps.sigStop:
			return
		case <-eps.scanTick.C:
			eps.next()
		case <-eps.observeTick.C:
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
	log.Noticef("current sealed epoch is #%d", ep.Id)
	eps.top = ep
}

// next processes epoch data based on the stored current epoch number.
func (eps *epochScanner) next() {
	// do we have a space to grow to
	if eps.top == nil || uint64(eps.top.Id) < eps.current.Load() {
		return
	}

	// try to get the epoch data
	id := eps.current.Load()
	ep, err := repo.Epoch((*hexutil.Uint64)(&id))
	if err != nil {
		log.Errorf("can not get epoch #%d; %s", id, err.Error())
		return
	}

	// process and move to the next epoch
	eps.queue <- ep
}

// dequeue consumes the queue and sends epochs to the persistent storage.
func (eps *epochScanner) dequeue() {
	for {
		// pull the next epoch from the queue
		ep, ok := <-eps.queue
		if !ok {
			log.Noticef("epoch scanner queue terminated")
			return
		}

		// store the epoch
		eps.store(ep)

		// update current epoch ID
		eps.current.Store(uint64(ep.Id))
	}
}

// store consumes the queue and sends epochs to the persistent storage.
func (eps *epochScanner) store(ep *types.Epoch) {
	// log the epoch
	if ep.EndTime == 0 {
		log.Debugf("epoch #%d details not available", eps.current)
		return
	}

	// log what we do
	log.Noticef("processing epoch #%d", ep.Id)

	// add the epoch to the database
	err := repo.AddEpoch(ep)
	if err != nil {
		log.Errorf("can not store epoch #%d; %s", ep.Id, err.Error())
	}
}

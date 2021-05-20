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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"sync"
	"time"
)

// sfcScanner implements SFC contract scanner used to extract data to off-chain storage.
// we can not scan anything before #5574 (migration); full data available after #5577
type sfcScanner struct {
	service
	epochQueue   chan *types.Epoch
	sigTermQueue chan bool
}

// newSFCScanner creates new instance of the SFC scanner service.
func newSFCScanner(repo Repository, log logger.Logger, wg *sync.WaitGroup) *sfcScanner {
	// create new blockScanner instance
	return &sfcScanner{
		service:      newService("sfc scanner", repo, log, wg),
		sigTermQueue: make(chan bool, 1),
		epochQueue:   make(chan *types.Epoch, 50),
	}
}

// run initializes the sfcScanner and starts scanning
func (sfs *sfcScanner) run() {
	// try to find the lats known epoch
	start, err := sfs.repo.LastKnownEpoch()
	if err != nil {
		sfs.log.Criticalf("can not get the last known epoch; %s", err.Error())
		start = 1
	}

	// start the epoch queue
	sfs.wg.Add(1)
	go sfs.monitor()

	// start blockScanner
	sfs.wg.Add(1)
	go sfs.scan(start)
}

// scan performs the actual SFC scanning operation
func (sfs *sfcScanner) scan(from uint64) {
	var (
		current    = from
		top        *types.Epoch
		topTicker  = time.NewTicker(10 * time.Second)
		scanTicker = time.NewTicker(50 * time.Millisecond)
	)

	// don't forget to sign off after we are done
	defer func() {
		// stop tickers
		topTicker.Stop()
		scanTicker.Stop()

		// report done
		sfs.log.Notice("epoch scanner done")
		sfs.wg.Done()
	}()

	sfs.log.Noticef("SFC scan starts at epoch #%d", current)
	for {
		select {
		case <-sfs.sigStop:
			sfs.sigTermQueue <- true
			return
		case <-topTicker.C:
			if err := sfs.updateTop(&top); err != nil {
				sfs.log.Criticalf("can not get sealed epoch; %s", err.Error())
				return
			}
			if current <= uint64(top.Id) {
				sfs.log.Infof("epoch scanner at #%d, top sealed epoch #%d", current, top.Id)
			}
		case <-scanTicker.C:
			// next epoch available?
			if top == nil || current > uint64(top.Id) {
				continue
			}
			if err := sfs.push(hexutil.Uint64(current)); err != nil {
				return
			}
			current++
		}
	}
}

// push loads the epoch from repository and pushes it for processing.
func (sfs *sfcScanner) push(epochID hexutil.Uint64) error {
	// try to get the epoch data
	ep, err := sfs.repo.Epoch(&epochID)
	if err != nil {
		sfs.log.Errorf("can not get epoch #%d; %s", epochID, err.Error())
		return err
	}

	// push the epoch for processing?
	if ep.EndTime > 0 {
		sfs.epochQueue <- ep
	}
	return nil
}

// updateTopEpoch updates the top sealed epoch from the SFC.
func (sfs *sfcScanner) updateTop(top **types.Epoch) (err error) {
	// what is the top epoch number
	ep, err := sfs.repo.CurrentSealedEpoch()
	if err != nil {
		sfs.log.Criticalf("can not get sealed epoch; %s", err.Error())
		return err
	}

	// still on the same epoch
	if ep == nil || ((*top) != nil && (*top).Id == ep.Id) {
		return nil
	}

	// a new epoch found
	sfs.log.Noticef("current sealed epoch is #%d", ep.Id)
	*top = ep
	return nil
}

// monitor collects epochs for processing and sends tem to the repository
// for storing in the off-chain database.
func (sfs *sfcScanner) monitor() {
	// don't forget to sign off after we are done
	defer func() {
		// report done
		sfs.log.Notice("epoch monitor done")
		sfs.wg.Done()
	}()

	sfs.log.Notice("epoch monitor started")
	for {
		select {
		case <-sfs.sigTermQueue:
			return
		case ep := <-sfs.epochQueue:
			sfs.process(ep)
		}
	}
}

// process processes the identified sealed epoch if possible.
func (sfs *sfcScanner) process(epoch *types.Epoch) {
	// log what we do
	sfs.log.Debugf("processing epoch #%d", epoch.Id)

	// add the epoch to the database
	err := sfs.repo.AddEpoch(epoch)
	if err != nil {
		sfs.log.Errorf("can not store epoch #%d; %s", epoch.Id, err.Error())
		return
	}
}

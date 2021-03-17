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
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"github.com/ethereum/go-ethereum/event"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ftm "github.com/ethereum/go-ethereum/rpc"
)

// uniswapNewPairMonitorQueueSize represents the size of the queue for processing new uniswap pair events.
const uniswapNewPairMonitorQueueSize = 50

// UniswapMonitor represents a processor capturing new swap events on the blockchain
type UniswapMonitor struct {
	service

	con      *ftm.Client
	swapChan chan *evtSwap

	// pairs represents a list of pair monitors we control
	pairs []*UniswapPairMonitor

	// newPairEvtCh represents a channel receiving events for new pairs we may want to monitor
	newPairEvtCh chan *contracts.UniswapFactoryPairCreated
	newPairSub   event.Subscription
}

// NewUniswapMonitor creates a new uniswap monitor instance.
func NewUniswapMonitor(con *ftm.Client, buffer chan *evtSwap, repo Repository, log logger.Logger, wg *sync.WaitGroup) *UniswapMonitor {
	// create new monitor instance
	um := UniswapMonitor{
		service:  newService("uniswap monitor", repo, log, wg),
		swapChan: buffer,
		con:      con,
		pairs:    make([]*UniswapPairMonitor, 0),
	}

	// return monitor instance
	return &um
}

// run starts monitoring for uniswap swap events
func (um *UniswapMonitor) run() {
	// log
	um.log.Notice("initializing uniswap contract monitor")

	// start pairs monitoring
	if err := um.runPairsMonitor(); err != nil {
		um.log.Errorf("uniswap monitor can not be started; %s", err.Error())
		return
	}

	// start pairs monitoring
	if err := um.runFactoryMonitor(); err != nil {
		um.log.Errorf("uniswap monitor can not be started; %s", err.Error())
		return
	}

	// inform about the monitor being started
	um.log.Notice("uniswap swap event monitor started")
}

// initFactoryMonitor initializes the core uniswap factory contract monitoring.
func (um *UniswapMonitor) runFactoryMonitor() error {
	// gets the Uniswap factory instance according to config address of the contract
	uniswapFactory, err := um.repo.UniswapFactoryContract()
	if err != nil {
		um.log.Errorf("uniswap factory contract instance was not created; %s", err.Error())
		return err
	}

	// create channel for capturing event for new pair creation
	um.newPairEvtCh = make(chan *contracts.UniswapFactoryPairCreated, uniswapNewPairMonitorQueueSize)
	um.newPairSub, err = uniswapFactory.WatchPairCreated(&bind.WatchOpts{}, um.newPairEvtCh, nil, nil)
	if err != nil {
		um.log.Errorf("can not subscribe to uniswap factory contract for tracking new pair creation; %s", err.Error())
		return err
	}

	// start uniswap factory monitoring routine
	um.wg.Add(1)
	go um.monitor()

	return nil
}

// runPairsMonitor starts monitoring of all pairs managed by the uniswap contract.
func (um *UniswapMonitor) runPairsMonitor() error {
	//get all pairs in blockchain
	pairs, err := um.repo.UniswapPairs()
	if err != nil {
		um.log.Errorf("uniswap pairs can not be resolved; %s", err.Error())
		return err
	}

	// log the size of the monitoring
	um.log.Noticef("found %d uniswap pairs to watch", len(pairs))

	// start monitoring for all pairs
	for _, pair := range pairs {
		// make the pair monitor and add the pair monitor to the list
		pm := NewUniswapPairMonitor(um.con, um.swapChan, um.repo, um.log, um.wg, &pair)
		um.pairs = append(um.pairs, pm)

		// start the pair monitor
		pm.run()
	}

	return nil
}

// monitor monitors new uniswap pairs event and adds the new pairs to the pair monitor collection.
func (um *UniswapMonitor) monitor() {
	// don't forget to sign off after we are done
	defer func() {
		// unsubscribe from new pairs monitoring
		um.newPairSub.Unsubscribe()

		// log
		um.log.Notice("uniswap factory monitor stopped")

		// signal to wait group we are done
		um.wg.Done()
	}()

	// loop here
	for {
		select {
		case <-um.sigStop:
			// the master monitor is terminating
			// signal all the pair monitors to stop as well
			for _, pair := range um.pairs {
				pair.close()
			}

			return
		case newPair := <-um.newPairEvtCh:
			// make the pair monitor and add the pair monitor to the list
			pm := NewUniswapPairMonitor(um.con, um.swapChan, um.repo, um.log, um.wg, &newPair.Pair)
			um.pairs = append(um.pairs, pm)

			// start the pair monitor
			pm.run()
		}
	}
}

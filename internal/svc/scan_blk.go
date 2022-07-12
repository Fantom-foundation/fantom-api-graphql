// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"sync/atomic"
	"time"
)

// blsObserverTickBaseDuration represents the frequency of the scanner status observer.
const blsObserverTickBaseDuration = 5 * time.Second

// blsScanTickBaseDuration represents the frequency of the scanner default progress.
const blsScanTickBaseDuration = 5 * time.Millisecond

// blsObserverTickIdleDuration represents the frequency of the scanner status observer on idle.
const blsObserverTickIdleDuration = 1 * time.Minute

// blsScanTickIdleDuration represents the frequency of the scanner re-check on idle.
const blsScanTickIdleDuration = 5 * time.Minute

// blsBlockBufferCapacity represents the capacity of the found blocks channel.
// When the channel is full, the push will have to wait for room here and the scanner
// will be slowed down naturally.
const blsBlockBufferCapacity = 1000

// blsReScanHysteresis is the number of blocks we wait from dispatcher until a re-scan kicks in.
const blsReScanHysteresis = 100

// blkScanner implements scanner loading previous/unknown blockchain blocks.
type blkScanner struct {
	service
	cfg            config.RepoCmd
	outBlock       chan *types.Block
	outStateSwitch chan bool
	inDispatched   chan uint64
	observeTick    *time.Ticker
	scanTick       *time.Ticker
	onIdle         bool
	from           uint64
	next           uint64
	to             uint64
	done           uint64
}

// name returns the name of the service used by orchestrator.
func (bls *blkScanner) name() string {
	return "block scanner"
}

// init prepares the block scanner.
func (bls *blkScanner) init() {
	bls.onIdle = false
	bls.sigStop = make(chan struct{})
	bls.outStateSwitch = make(chan bool, 1)
	bls.outBlock = make(chan *types.Block, blsBlockBufferCapacity)
}

// run starts the block dispatcher
func (bls *blkScanner) run() {
	if bls.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", bls.name()))
	}
	if bls.inDispatched == nil {
		panic(fmt.Errorf("no input block number channel"))
	}

	// get the scanner range
	start, err := bls.boundaries()
	if err != nil {
		log.Errorf("scanner can not proceed; %s", err.Error())
		return
	}

	// signal orchestrator we started and go
	log.Noticef("block scan starts at #%d", start)
	bls.from = start
	bls.next = start

	bls.mgr.started(bls)
	go bls.execute()
}

// close terminates the block dispatcher.
func (bls *blkScanner) close() {
	if bls.scanTick != nil {
		bls.scanTick.Stop()
		bls.observeTick.Stop()
	}
	if bls.sigStop != nil {
		close(bls.sigStop)
	}
}

// boundaries provides the block scanner initial range.
func (bls *blkScanner) boundaries() (uint64, error) {
	// get the newest known transaction
	lnb, err := repo.LastKnownBlock()
	if err != nil {
		log.Critical("can not scan blockchain; %s", err.Error())
		return 0, err
	}

	// apply re-scan
	if lnb > bls.cfg.BlockScanReScan {
		log.Debugf("last known block is #%d, re-scanning %d blocks", lnb, bls.cfg.BlockScanReScan)
		lnb = lnb - bls.cfg.BlockScanReScan
	}
	return lnb, nil
}

// execute scans blockchain blocks in the given range and push found blocks
// to the output channel for processing.
func (bls *blkScanner) execute() {
	defer func() {
		close(bls.outBlock)
		close(bls.outStateSwitch)
		bls.mgr.finished(bls)
	}()

	// set initial state and start the tickers for observer and scanner
	bls.observe()
	bls.observeTick = time.NewTicker(blsObserverTickBaseDuration)
	bls.scanTick = time.NewTicker(blsScanTickBaseDuration)

	// do the scan
	for {
		select {
		case <-bls.sigStop:
			return
		case bin, ok := <-bls.inDispatched:
			// ignore block re-scans; do not skip blocks in dispatched # counter
			done := atomic.LoadUint64(&bls.done)
			if ok && (done == 0 || int64(bin)-int64(done) == 1) {
				atomic.StoreUint64(&bls.done, bin)
			}
		case <-bls.observeTick.C:
			bls.updateState(bls.observe())
		case <-bls.scanTick.C:
			bls.shift()
		}
	}
}

// observe updates the scanner final block and logs the progress.
// It returns expected idle state to be used to transition if needed.
func (bls *blkScanner) observe() bool {
	// try to get the block height
	bh, err := repo.BlockHeight()
	if err != nil {
		log.Errorf("can not get current block height; %s", err.Error())
		return false
	}

	// if on idle, wait for the dispatcher to catch up with the blocks
	// we use a hysteresis to delay state flip back to active scan
	// we compare current block height with the latest known dispatched block number
	target := bh.ToInt().Uint64()
	done := atomic.LoadUint64(&bls.done)

	if bls.onIdle && target < done+blsReScanHysteresis {
		bls.next = done
		bls.from = done
		log.Infof("block scanner idling at #%d, head at #%d", bls.next, target)
		return true
	}

	// adjust target block number; log the progress of the scan
	bls.to = target
	log.Infof("block scanner at #%d of <#%d, #%d>, #%d dispatched", bls.next, bls.from, bls.to, done)
	return bls.to < bls.next
}

// updateState change scanner state if needed.
// It resets the internal tickers according to the target state.
func (bls *blkScanner) updateState(target bool) {
	// if the state already match, do nothing
	if target == bls.onIdle {
		return
	}

	// switch the state; advertise the transition
	log.Noticef("block scanner idle state toggled to %t", target)
	bls.onIdle = target

	select {
	case bls.outStateSwitch <- target:
	case <-bls.sigStop:
		return
	}

	// going full speed
	if !target {
		bls.observeTick.Reset(blsObserverTickBaseDuration)
		bls.scanTick.Reset(blsScanTickBaseDuration)
		return
	}

	// going idle
	bls.observeTick.Reset(blsObserverTickIdleDuration)
	bls.scanTick.Reset(blsScanTickIdleDuration)
}

// next pulls the next block if available and pushes it for processing.
func (bls *blkScanner) shift() {
	// we may not need to pull at all, if on updateState
	if bls.onIdle {
		return
	}

	// are we at the end? check the status
	if bls.next > bls.to {
		bls.updateState(bls.observe())
		return
	}

	// pull the current block
	block, err := repo.BlockByNumber((*hexutil.Uint64)(&bls.next))
	if err != nil {
		log.Errorf("block #%d not available; %s", bls.next, err.Error())
		return
	}

	// push the block for processing and advance to the next expected block
	// observe possible stop signal during a wait for the block queue slot
	select {
	case bls.outBlock <- block:
		bls.next++
	case <-bls.sigStop:
	}
}

// blockHeight provides information about processed block height.
func (bls *blkScanner) blockHeight() uint64 {
	return atomic.LoadUint64(&bls.done)
}

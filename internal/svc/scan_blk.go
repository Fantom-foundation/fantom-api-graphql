// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/atomic"
	"time"
)

// blsObserverTickBaseDuration represents the frequency of the scanner status observer.
const blsObserverTickBaseDuration = 5 * time.Second

// blsScanTickBaseDuration represents the frequency of the scanner default progress.
const blsScanTickBaseDuration = 3 * time.Millisecond

// blsObserverTickIdleDuration represents the frequency of the scanner status observer on idle.
const blsObserverTickIdleDuration = 30 * time.Second

// blsScanTickIdleDuration represents the frequency of the scanner re-check on idle.
const blsScanTickIdleDuration = 10 * time.Second

// blsReScanHysteresis is the number of blocks we wait from dispatcher until a re-scan kicks in.
const blsReScanHysteresis = 50

// blsBlockBufferCapacity represents the capacity of the found blocks channel.
const blsBlockBufferCapacity = 20000

// blkScanner implements scanner loading previous/unknown blockchain blocks.
type blkScanner struct {
	service
	cfg          config.RepoCmd
	outBlock     chan *types.Block
	inDispatched chan uint64
	observeTick  *time.Ticker
	scanTick     *time.Ticker
	onIdle       *atomic.Bool
	toIdle       chan bool
	from         uint64
	next         uint64
	to           uint64
	done         uint64
}

// name returns the name of the service used by orchestrator.
func (bls *blkScanner) name() string {
	return "block scanner"
}

// init prepares the block scanner.
func (bls *blkScanner) init() {
	bls.sigStop = make(chan bool, 1)
	bls.toIdle = make(chan bool, 1)
	bls.onIdle = atomic.NewBool(false)
	bls.outBlock = make(chan *types.Block, blsBlockBufferCapacity)
}

// run starts the block dispatcher
func (bls *blkScanner) run() {
	if bls.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", bls.name()))
	}
	if bls.outBlock == nil {
		panic(fmt.Errorf("no output block channel"))
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
		bls.sigStop <- true
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
		close(bls.sigStop)
		close(bls.outBlock)
		close(bls.toIdle)
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
			// ignore block re-scans
			if ok && bin > bls.done {
				bls.done = bin
			}
		case <-bls.observeTick.C:
			bls.idle(bls.observe())
		case <-bls.scanTick.C:
			bls.shift()
		}
	}
}

// observe updates the scanner final block to scan towards and logs the progress.
// Returns expected idle state (true in case the top block has been reached).
func (bls *blkScanner) observe() bool {
	// try to get the block height
	bh, err := repo.BlockHeight()
	if err != nil {
		log.Errorf("can not get scanner target; %s", err.Error())
		return false
	}

	// if on idle, wait for the dispatcher to catch up with the blocks
	// we use a hysteresis boundary before the state is flipped to re-scan
	onIdle := bls.onIdle.Load()
	target := bh.ToInt().Uint64()
	diff := target - bls.done
	if onIdle && diff < blsReScanHysteresis {
		bls.next = bls.done
		bls.from = bls.done
		log.Infof("block scanner idling on #%d, head on #%d", bls.next, target)
		return false
	}

	// log the progress of the scan process
	bls.to = target
	log.Infof("block scanner on #%d of <#%d, #%d>, #%d dispatched", bls.next, bls.from, bls.to, bls.done)

	// should we run on full speed (false), or on idle (true)?
	return bls.to < bls.next
}

// idle change scanner idle state if needed by resetting the internal tickers.
func (bls *blkScanner) idle(target bool) {
	// if the state already match, nothing to do
	if target == bls.onIdle.Load() {
		return
	}

	// switch the idle state
	log.Noticef("block scanner idle toggled to %t", target)
	bls.onIdle.Toggle()

	// going to active?
	if !target {
		bls.observeTick.Reset(blsObserverTickBaseDuration)
		bls.scanTick.Reset(blsScanTickBaseDuration)
		return
	}

	// going to idle, signal to orchestrator
	bls.observeTick.Reset(blsObserverTickIdleDuration)
	bls.scanTick.Reset(blsScanTickIdleDuration)
	bls.toIdle <- true
}

// next pulls the next block if available and pushes it for processing.
func (bls *blkScanner) shift() {
	// are we on the end already
	if bls.next > bls.to {
		return
	}

	// we may not need to pull at all, if on idle
	if bls.onIdle.Load() {
		return
	}

	// pull the current block
	block, err := repo.BlockByNumber((*hexutil.Uint64)(&bls.next))
	if err != nil {
		log.Errorf("block #%d not available; %s", bls.next, err.Error())
		return
	}

	// push the block for processing and advance to the next expected block
	bls.outBlock <- block
	bls.next++
}

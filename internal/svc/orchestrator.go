// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/repository/cache/ring"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	etc "github.com/ethereum/go-ethereum/core/types"
	"unsafe"
)

// orBlockCacheCapacity represents the capacity of the local block cache.
const orBlockCacheCapacity = 50

// orchestrator implements service responsible for moderating connections between other services.
type orchestrator struct {
	service
	blkCache          *ring.Ring
	pushHeads         bool
	inScanStateSwitch chan bool
}

// name returns the name of the service used by manager.
func (or *orchestrator) name() string {
	return "orchestrator"
}

// init sets the initial connection state for the managed services
func (or *orchestrator) init() {
	or.sigStop = make(chan struct{})
	or.blkCache = ring.New(orBlockCacheCapacity)

	// connect services' input channels to their source
	or.mgr.trd.inTransaction = or.mgr.bld.outTransaction
	or.mgr.acd.inAccount = or.mgr.trd.outAccount
	or.mgr.lgd.inLog = or.mgr.trd.outLog
	or.mgr.bld.inBlock = or.mgr.bls.outBlock
	or.mgr.bls.inDispatched = or.mgr.bld.outDispatched
	or.mgr.bud.inTransaction = or.mgr.trd.outTransaction
	or.inScanStateSwitch = or.mgr.bls.outStateSwitch

	// read initial block scanner state
	// no need to worry about race condition, init() is called sequentially and this is the last one
	or.pushHeads = or.mgr.bls.onIdle
}

// run starts the orchestrator service business
func (or *orchestrator) run() {
	// make sure we are orchestrated
	if or.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", or.name()))
	}

	// signal manager we started and go
	or.mgr.started(or)
	go or.execute()
}

// execute performs the active data routing between received heads
// and either local blocks cache, or the block dispatcher,
// based on the block scanner state. The scanner state transition is observed
// by an inbound channel.
func (or *orchestrator) execute() {
	defer func() {
		or.mgr.finished(or)
	}()

	// access the new heads queue
	// it's filled with new heads as the connected node processes blocks from the network
	heads := repo.ObservedHeaders()
	for {
		select {
		case <-or.sigStop:
			return
		case h, ok := <-heads:
			if ok {
				or.handleNewHead(h)
			}
		case idle, ok := <-or.inScanStateSwitch:
			if ok {
				or.pushHeads = idle
				if idle {
					or.unloadCache()
				}
			}
		}
	}
}

// handleNewHead processes incoming header and handles it according
// to the state of the block scanner by either pushing the corresponding block
// to dispatcher queue, or by putting the block to the local ring cache for future use.
func (or *orchestrator) handleNewHead(h *etc.Header) {
	// get the block
	bn := h.Number.Uint64()
	blk, err := repo.BlockByNumber((*hexutil.Uint64)(&bn))
	if err != nil {
		log.Errorf("block #%d not available; %s", bn, err.Error())
		return
	}

	// if the block scanner is on idle, push the block directly to processing queue
	if or.pushHeads {
		or.mgr.bld.inBlock <- blk
		return
	}

	// store the block in the ring cache for now; we build up an in-memory collection of the latest blocks
	// after the block scanner goes idle, we unload these blocks from ring cache
	// into the processing queue to make sure we didn't miss any of them on scan to idle transition
	or.blkCache.Add(unsafe.Pointer(blk))
}

// unloadCache pushes all the blocks currently stored in cache (e.g. blocks of the most recent heads)
// into the block processing queue to make sure they get all processed, and we don't miss any
// on block scanner full speed to idle transition (consistency feature, may not be needed).
func (or *orchestrator) unloadCache() {
	// inform about the cache unloading
	log.Noticef("unloading block cache for processing")

	// pull all available cached blocks and reset the ring to make sure
	// we don't re-send the same blocks again
	l := or.blkCache.List(orBlockCacheCapacity)
	or.blkCache.Reset()

	// push them all to dispatcher for processing
	for _, blk := range l {
		log.Infof("cached block #%d sent for processing", (*types.Block)(blk).Number)
		or.mgr.bld.inBlock <- (*types.Block)(blk)
	}
}

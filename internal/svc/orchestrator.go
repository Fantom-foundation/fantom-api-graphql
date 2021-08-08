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
const orBlockCacheCapacity = 25

// orchestrator implements service responsible for moderating connections between other services.
type orchestrator struct {
	service
	blkCache *ring.Ring
}

// name returns the name of the service used by orchestrator.
func (or *orchestrator) name() string {
	return "orchestrator"
}

// init sets the initial connection state for the managed services
func (or *orchestrator) init() {
	or.sigStop = make(chan bool, 1)
	or.blkCache = ring.New(orBlockCacheCapacity)

	// connect services' input channels to their source
	or.mgr.trd.inTransaction = or.mgr.bld.outTransaction
	or.mgr.acd.inAccount = or.mgr.trd.outAccount
	or.mgr.lgd.inLog = or.mgr.trd.outLog
	or.mgr.bld.inBlock = or.mgr.bls.outBlock
	or.mgr.bls.inDispatched = or.mgr.bld.outDispatched
}

// run starts the block dispatcher
func (or *orchestrator) run() {
	// make sure we are orchestrated
	if or.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", or.name()))
	}

	// signal orchestrator we started and go
	or.mgr.started(or)
	go or.execute()
}

// execute performs the service interaction management
func (or *orchestrator) execute() {
	defer func() {
		close(or.sigStop)
		or.mgr.finished(or)
	}()

	heads := repo.ObservedHeaders()

	// do the scan
	for {
		select {
		case <-or.sigStop:
			return
		case h, ok := <-heads:
			if ok {
				or.handleNewHead(h)
			}
		case _, ok := <-or.mgr.bls.toIdle:
			if ok {
				or.unloadCache()
			}
		}
	}
}

// handleNewHead processes incoming observed header and handles it according
// to the state of the block scanner by either pushing the corresponding block
// to dispatcher queue, or by putting the block to the ring for future use.
func (or *orchestrator) handleNewHead(h *etc.Header) {
	// get the block
	bn := h.Number.Uint64()
	blk, err := repo.BlockByNumber((*hexutil.Uint64)(&bn))
	if err != nil {
		log.Errorf("block #%d not available; %s", bn, err.Error())
		return
	}

	// if the block scanner is on idle, push to process
	if or.mgr.bls.onIdle.Load() {
		or.mgr.bld.inBlock <- blk
		return
	}

	// store the block in the cache for now
	or.blkCache.Add(unsafe.Pointer(blk))
}

// unloadCache pushes all the blocks currently stored in cache
// into the block processing queue to make sure they get all processed.
func (or *orchestrator) unloadCache() {
	// inform about the cache unloading
	log.Noticef("unloading block cache for processing")

	// pull all available cached blocks and send them to dispatch
	l := or.blkCache.List(orBlockCacheCapacity)
	for _, blk := range l {
		or.mgr.bld.inBlock <- (*types.Block)(blk)
	}

	// reset the ring cache depth, so we don't process those blocks again
	or.blkCache.Reset()
}

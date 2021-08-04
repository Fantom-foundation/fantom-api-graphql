// Package svc implements blockchain data processing services.
package svc

import (
	"fmt"
)

// orchestrator implements service responsible for moderating connections between other services.
type orchestrator struct {
	service
}

// name returns the name of the service used by orchestrator.
func (or *orchestrator) name() string {
	return "moderator"
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

// init sets the initial connection state for the managed services
func (or *orchestrator) init() {
	or.sigStop = make(chan bool, 1)

	// connect services' input channels to their source
	or.mgr.trd.inTransaction = or.mgr.bld.outTransaction
	or.mgr.acd.inAccount = or.mgr.trd.outAccount
	or.mgr.lgd.inLog = or.mgr.trd.outLog
	or.mgr.bls.outBlock = or.mgr.bld.inBlock
}

// execute performs the service interaction management
func (or *orchestrator) execute() {
	defer func() {
		close(or.sigStop)
		or.mgr.finished(or)
	}()

	// do the scan
	for {
		// capture stop signal
		select {
		case <-or.sigStop:
			return
		}
	}
}

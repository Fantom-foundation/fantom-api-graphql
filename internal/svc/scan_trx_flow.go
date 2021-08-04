// Package svc implements blockchain data processing services.
package svc

import (
	"fmt"
	"time"
)

const (
	// trxFlowUpdaterPeriod represents the period in which we do trx flow updates.
	trxFlowUpdaterPeriod = 7 * time.Minute

	// trxCountUpdaterPeriod represents the period in which the trx count estimation
	// is updated from the underlying database.
	trxCountUpdaterPeriod = 30 * time.Minute
)

// trxFlowMonitor represents a service for transaction flow monitoring.
type trxFlowMonitor struct {
	service
	flowTicker  *time.Ticker
	countTicker *time.Ticker
}

// name returns a human-readable name of the service used by the manager.
func (tfm *trxFlowMonitor) name() string {
	return "trx flow monitor"
}

// run starts the transaction flow monitoring.
func (tfm *trxFlowMonitor) run() {
	// make sure we are orchestrated
	if tfm.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", tfm.name()))
	}

	// start go routine for processing
	tfm.mgr.started(tfm)
	go tfm.execute()
}

// close terminates the gas price suggestion monitor.
func (tfm *trxFlowMonitor) close() {
	if tfm.flowTicker != nil {
		tfm.flowTicker.Stop()
		tfm.countTicker.Stop()
	}
	if tfm.sigStop != nil {
		tfm.sigStop <- true
	}
}

// execute performs regular ticker based checks on the transaction flow
// and sends the collected data to persistent repository.
func (tfm *trxFlowMonitor) execute() {
	defer func() {
		close(tfm.sigStop)
		tfm.mgr.finished(tfm)
	}()

	// do initial trx count update
	go tfm.updateCount()

	// start to control the monitor
	tfm.flowTicker = time.NewTicker(trxFlowUpdaterPeriod)
	tfm.countTicker = time.NewTicker(trxCountUpdaterPeriod)

	// loop here
	for {
		select {
		case <-tfm.sigStop:
			return
		case <-tfm.flowTicker.C:
			repo.TrxFlowUpdate()
		case <-tfm.countTicker.C:
			go tfm.updateCount()
		}
	}
}

// updateCount updates trx counter estimation.
func (tfm *trxFlowMonitor) updateCount() {
	// pull the value from DB
	val, err := repo.TransactionsCount()
	if err != nil {
		log.Errorf("can not update trx count estimation; %s", err.Error())
		return
	}

	// update the estimate
	repo.UpdateTrxCountEstimate(val)
}

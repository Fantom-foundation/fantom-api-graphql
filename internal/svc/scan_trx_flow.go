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
	trxCountUpdaterPeriod = 31 * time.Minute

	// burnFlowUpdaterPeriod represents the period in which the burn aggregate
	// is updated from the underlying database.
	burnFlowUpdaterPeriod = 23 * time.Minute
)

// trxFlowMonitor represents a service for transaction flow monitoring.
type trxFlowMonitor struct {
	service
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
	if tfm.sigStop != nil {
		close(tfm.sigStop)
	}
}

// execute performs regular ticker based checks on the transaction flow
// and sends the collected data to persistent repository.
func (tfm *trxFlowMonitor) execute() {
	// start to control the monitor
	flowTicker := time.NewTicker(trxFlowUpdaterPeriod)
	countTicker := time.NewTicker(trxCountUpdaterPeriod)
	burnTicker := time.NewTicker(burnFlowUpdaterPeriod)

	defer func() {
		flowTicker.Stop()
		countTicker.Stop()
		burnTicker.Stop()

		tfm.mgr.finished(tfm)
	}()

	// do initial trx count update
	go tfm.updateCount()

	// loop here
	for {
		select {
		case <-tfm.sigStop:
			return
		case <-flowTicker.C:
			repo.TrxFlowUpdate()
		case <-burnTicker.C:
			repo.BurnDailyUpdate()
		case <-countTicker.C:
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

/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/types"
	"time"
)

// trxFlowUpdateRange represents the range for which we do the trx flow update.
const trxFlowUpdateRange = -2 * 24 * time.Hour

// burnFlowUpdateRange represents the range for which we do the aggregated burn flow update.
const burnFlowUpdateRange = -7 * 24 * time.Hour

// TrxFlowVolume resolves the list of daily trx flow aggregations.
func (p *proxy) TrxFlowVolume(from *time.Time, to *time.Time) ([]*types.DailyTrxVolume, error) {
	return p.db.TrxDailyFlowList(from, to)
}

// TrxFlowSpeed provides speed of transaction per second for the last <sec> seconds.
func (p *proxy) TrxFlowSpeed(sec int32) (float64, error) {
	return p.db.TrxRecentTrxSpeed(sec)
}

// TrxGasSpeed provides speed of gas consumption per second by transactions.
func (p *proxy) TrxGasSpeed(from *time.Time, to *time.Time) (float64, error) {
	return p.db.TrxGasSpeed(from, to)
}

// TrxFlowUpdate executes the trx flow update in the database.
func (p *proxy) TrxFlowUpdate() {
	// calculate previous midnight
	now := time.Now().UTC()
	h, m, s := now.Clock()
	from := now.Add(time.Duration(-(h*3600 + m*60 + s)) * time.Second).Add(time.Duration(-now.Nanosecond()) * time.Nanosecond).Add(trxFlowUpdateRange)

	// do the update
	err := p.db.TrxDailyFlowUpdate(from)
	if err != nil {
		p.log.Criticalf("can not update trx flow; %s", err.Error())
	}

	// log success
	p.log.Debugf("trx flow updated")
}

// BurnDailyUpdate provides an aggregated amount of burned FTMs by days.
func (p *proxy) BurnDailyUpdate() {
	// calculate previous midnight
	now := time.Now().UTC()
	h, m, s := now.Clock()
	from := now.Add(time.Duration(-(h*3600 + m*60 + s)) * time.Second).Add(time.Duration(-now.Nanosecond()) * time.Nanosecond).Add(burnFlowUpdateRange)

	// do the update
	err := p.db.FeeFlowAggregateUpdate(from, now)
	if err != nil {
		p.log.Criticalf("can not update burn aggregate flow; %s", err.Error())
	}

	// log success
	p.log.Info("burn flow aggregate updated")
}

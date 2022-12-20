/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/types"
	"math"
	"math/big"
	"time"
)

// BurnTreasuryShare is the share between burned and treasury fee.
type BurnTreasuryShare struct {
	SinceBlockID    uint64
	UntilBlockID    uint64
	Since           int64
	Until           int64
	ToBurn          *big.Int
	ToTreasury      *big.Int
	ToRewards       *big.Int
	DigitCorrection *big.Int
}

var (
	// burnTreasury defines the amount of FTM burned and sent to treasury of different setups based on starting block
	burnTreasury = []BurnTreasuryShare{
		{
			SinceBlockID:    0,
			UntilBlockID:    48552263,
			Since:           0,
			Until:           1665060100,
			ToBurn:          big.NewInt(300),
			ToTreasury:      new(big.Int),
			ToRewards:       big.NewInt(700),
			DigitCorrection: big.NewInt(1000),
		},
		{
			SinceBlockID:    48552264,
			UntilBlockID:    math.MaxUint64,
			Since:           1665060101,
			Until:           math.MaxInt64,
			ToBurn:          big.NewInt(200),
			ToTreasury:      big.NewInt(100),
			ToRewards:       big.NewInt(700),
			DigitCorrection: big.NewInt(1000),
		},
	}
)

// StoreFtmBurn stores the given native FTM burn per block record into the persistent storage.
func (p *proxy) StoreFtmBurn(burn *types.FtmBurn) error {
	p.cache.FtmBurnUpdate(burn, p.db.BurnTotal)
	return p.db.StoreBurn(burn)
}

// FtmBurnTotal provides the total amount of burned native FTM.
func (p *proxy) FtmBurnTotal() (int64, error) {
	return p.cache.FtmBurnTotal(p.db.BurnTotal)
}

// FtmBurnList provides list of per-block burned native FTM tokens.
func (p *proxy) FtmBurnList(count int64) ([]*types.FtmBurn, error) {
	return p.db.BurnList(count)
}

// BurnTreasuryStashShareByBlock finds treasury/burn share for the given block ID.
func (p *proxy) BurnTreasuryStashShareByBlock(blk uint64) *BurnTreasuryShare {
	for _, s := range burnTreasury {
		if blk >= s.SinceBlockID && blk <= s.UntilBlockID {
			return &s
		}
	}
	return nil
}

// BurnTreasuryStashShareByTimeStamp finds treasury/burn share for the given time stamp.
func (p *proxy) BurnTreasuryStashShareByTimeStamp(ts int64) *BurnTreasuryShare {
	for _, s := range burnTreasury {
		if ts >= s.Since && ts <= s.Until {
			return &s
		}
	}
	return nil
}

// FeeFlow provides a list of fee flow aggregates for the given date range.
func (p *proxy) FeeFlow(from, to time.Time) ([]*types.FtmDailyBurn, error) {
	p.log.Infof("loading fee flow from %s to %s", from.String(), to.String())
	return p.db.FeeFlow(from, to)
}

/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"bytes"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// sfcDecimalUnit represents decimal units adjustment used by SFC contract
// on certain values calculation to preserve calculations precision.
var sfcDecimalUnit = new(big.Int).SetUint64(1e18)

// SfcDecimalUnit returns the decimal unit adjustment used by the SFC contract.
func (p *proxy) SfcDecimalUnit() *big.Int {
	return sfcDecimalUnit
}

// SfcVersion returns current version of the SFC contract.
func (p *proxy) SfcVersion() (hexutil.Uint64, error) {
	return p.rpc.SfcVersion()
}

// SfcConfiguration provides SFC contract configuration.
func (p *proxy) SfcConfiguration() (*types.SfcConfig, error) {
	// try cache first
	c := p.cache.PullSfcConfig()
	if c == nil {
		// load the config with all the values filled
		c = &types.SfcConfig{
			MinValidatorStake:      p.pullSfcConfigValue(p.rpc.SfcMinValidatorStake),
			MaxDelegatedRatio:      p.pullSfcConfigValue(p.rpc.SfcMaxDelegatedRatio),
			MinLockupDuration:      p.pullSfcConfigValue(p.rpc.SfcMinLockupDuration),
			MaxLockupDuration:      p.pullSfcConfigValue(p.rpc.SfcMaxLockupDuration),
			WithdrawalPeriodEpochs: p.pullSfcConfigValue(p.rpc.SfcWithdrawalPeriodEpochs),
			WithdrawalPeriodTime:   p.pullSfcConfigValue(p.rpc.SfcWithdrawalPeriodTime),
		}
		// cache for future use
		p.cache.PushSfcConfig(c)
	}
	return c, nil
}

// pullSfcConfigValue pulls SFC config value for the given value loader function.
func (p *proxy) pullSfcConfigValue(f func() (*big.Int, error)) hexutil.Big {
	val, err := f()
	if err != nil {
		p.log.Errorf("can not load SFC config value; %s", err.Error())
		return hexutil.Big{}
	}
	return (hexutil.Big)(*val)
}

// CurrentEpoch returns the id of the current epoch.
func (p *proxy) CurrentEpoch() (hexutil.Uint64, error) {
	return p.rpc.CurrentEpoch()
}

// Epoch returns the structure of the current epoch.
func (p *proxy) Epoch(id *hexutil.Uint64) (*types.Epoch, error) {
	// get the current epoch if the id has not been provided
	if id == nil || *id == 0 {
		// ask for the sealed epoch id
		val, err := p.rpc.CurrentSealedEpoch()
		if err != nil {
			return nil, err
		}
		id = &val
	}

	// try the cache first
	ep := p.cache.PullEpoch(id)
	if ep != nil {
		return ep, nil
	}

	// pull from remote
	ep, err := p.rpc.Epoch(*id)
	if err != nil {
		return nil, err
	}

	// find fee share
	share := p.BurnTreasuryStashShareByTimeStamp(int64(ep.EndTime))
	if share == nil {
		p.log.Criticalf("epoch fee share not found at %d", ep.Id)
		return nil, fmt.Errorf("fee share not found")
	}

	// calculate fee shares
	ep.EpochFeeBurn = hexutil.Big(*new(big.Int).Div(new(big.Int).Mul(ep.EpochFee.ToInt(), share.ToBurn), share.DigitCorrection))
	ep.EpochFeeTreasury = hexutil.Big(*new(big.Int).Div(new(big.Int).Mul(ep.EpochFee.ToInt(), share.ToTreasury), share.DigitCorrection))

	// cache for future use
	p.cache.PushEpoch(ep)
	return ep, nil
}

// CurrentSealedEpoch returns the data of the latest sealed epoch.
// This is used for reward estimation calculation and we don't need
// real time data, but rather faster response time.
// So, we use cache for handling the response.
// It will not be updated in sync with the SFC contract.
// If you need real time response, please use the Epoch(id) function instead.
func (p *proxy) CurrentSealedEpoch() (*types.Epoch, error) {
	// inform what we do
	p.log.Debug("latest sealed epoch requested")

	// we need to go the slow path
	id, err := p.rpc.CurrentSealedEpoch()
	if err != nil {
		p.log.Errorf("can not get the id of the last sealed epoch; %s", err.Error())
		return nil, err
	}
	return p.Epoch(&id)
}

// TotalStaked calculates current total staked amount for all stakers.
func (p *proxy) TotalStaked() (*hexutil.Big, error) {
	// try cache first
	value := p.cache.PullTotalStaked()
	if value != nil {
		p.log.Debugf("total staked amount loaded from memory cache")
		return value, nil
	}

	// get the actual live value
	total, err := p.rpc.TotalStaked()
	if err != nil {
		p.log.Errorf("can not get the total staked amount; %s", err.Error())
		return nil, err
	}

	// store in cache
	if err := p.cache.PushTotalStaked((*hexutil.Big)(total)); err != nil {
		// log issue
		p.log.Errorf("can not store total staked amount in memory; %s", err.Error())
	}

	// return the value
	return (*hexutil.Big)(total), nil
}

// RewardsAllowed returns the reward lock status from SFC.
func (p *proxy) RewardsAllowed() (bool, error) {
	return p.rpc.RewardsAllowed()
}

// LockingAllowed indicates if the stake locking has been enabled in SFC.
func (p *proxy) LockingAllowed() (bool, error) {
	return p.rpc.LockingAllowed()
}

// IsSfcContract returns true if the given address points to the SFC contract.
func (p *proxy) IsSfcContract(addr *common.Address) bool {
	return bytes.Equal(addr.Bytes(), p.cfg.Staking.SFCContract.Bytes())
}

// LastKnownEpoch returns the id of the last known and scanned epoch.
func (p *proxy) LastKnownEpoch() (uint64, error) {
	return p.db.LastKnownEpoch()
}

// AddEpoch stores an epoch reference in connected persistent storage.
func (p *proxy) AddEpoch(e *types.Epoch) error {
	return p.db.AddEpoch(e)
}

// Epochs pulls list of epochs starting at the specified cursor.
func (p *proxy) Epochs(cursor *string, count int32) (*types.EpochList, error) {
	return p.db.Epochs(cursor, count)
}

// FtmTreasuryTotal provides the total amount of native FTM sent into treasury.
func (p *proxy) FtmTreasuryTotal() (int64, error) {
	// pull the database value
	val, err := p.db.TreasuryTotal()
	if err != nil {
		return 0, err
	}

	// add correction for the sponsored amount
	return val, nil
}

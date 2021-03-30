/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"bytes"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// SfcVersion returns current version of the SFC contract.
func (p *proxy) SfcVersion() (hexutil.Uint64, error) {
	return p.rpc.SfcVersion()
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

		// use this id instead of nil
		id = &val
	}

	return p.rpc.Epoch(*id)
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

	// try to use the in-memory cache
	if ep := p.cache.PullLastEpoch(); ep != nil {
		// inform what we do
		p.log.Debug("latest sealed epoch loaded from cache")
		return ep, nil
	}

	// we need to go the slow path
	id, err := p.rpc.CurrentSealedEpoch()
	if err != nil {
		// inform what we do
		p.log.Errorf("can not get the id of the last sealed epoch; %s", err.Error())
		return nil, err
	}

	// get the epoch from SFC
	ep, err := p.rpc.Epoch(id)
	if err != nil {
		p.log.Errorf("can not get data of the last sealed epoch; %s", err.Error())
		return nil, err
	}

	// try to store the epoch in cache for future use
	err = p.cache.PushLastEpoch(ep)
	if err != nil {
		p.log.Error(err)
	}

	// inform what we do
	p.log.Debugf("epoch %s loaded from sfc", id.String())
	return ep, nil
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

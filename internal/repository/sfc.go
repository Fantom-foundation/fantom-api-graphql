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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// CurrentEpoch returns the id of the current epoch.
func (p *proxy) CurrentEpoch() (hexutil.Uint64, error) {
	return p.rpc.CurrentEpoch()
}

// Epoch returns the id of the current epoch.
func (p *proxy) Epoch(id hexutil.Uint64) (types.Epoch, error) {
	return p.rpc.Epoch(id)
}

// LastStakerId returns the last staker id in Opera blockchain.
func (p *proxy) LastStakerId() (hexutil.Uint64, error) {
	return p.rpc.LastStakerId()
}

// StakersNum returns the number of stakers in Opera blockchain.
func (p *proxy) StakersNum() (hexutil.Uint64, error) {
	return p.rpc.StakersNum()
}

// Staker extract a staker information from SFC smart contract.
func (p *proxy) Staker(id hexutil.Uint64) (*types.Staker, error) {
	return p.rpc.Staker(id)
}

// Staker extract a staker information by address.
func (p *proxy) StakerByAddress(addr common.Address) (*types.Staker, error) {
	return p.rpc.StakerByAddress(addr)
}

// DelegationsOf extract a list of delegations for a given staker.
func (p *proxy) DelegationsOf(staker hexutil.Uint64) ([]types.Delegator, error) {
	return p.rpc.DelegationsOf(staker)
}

// Delegation returns a detail of delegation for the given address.
func (p *proxy) Delegation(addr common.Address) (*types.Delegator, error) {
	return p.rpc.Delegation(addr)
}

// Delegation returns a detail of delegation for the given address.
func (p *proxy) DelegationRewards(addr string) (types.PendingRewards, error) {
	p.log.Debugf("processing %s", addr)
	return p.rpc.DelegationRewards(addr)
}

// WithdrawRequests extracts a list of partial withdraw requests
// for the given address.
func (p *proxy) WithdrawRequests(addr *common.Address) ([]*types.WithdrawRequest, error) {
	// log the action
	p.log.Debugf("loading withdraw requests for [%s]", addr.String())

	// proxy the request directly to RPC/SFC
	return p.rpc.WithdrawRequests(addr)
}

// DeactivatedDelegation extracts a list of deactivated delegation requests
// for the given address.
func (p *proxy) DeactivatedDelegation(addr *common.Address) ([]*types.DeactivatedDelegation, error) {
	// log the action
	p.log.Debugf("loading deactivated delegation requests for [%s]", addr.String())

	// proxy the request directly to RPC/SFC
	return p.rpc.DeactivatedDelegation(addr)
}

// delegatedAmount calculates total amount currently delegated
// and amount locked in pending un-delegation.
// Partial Un-delegations are subtracted during the preparation
// phase, but total un-delegations are subtracted only when
// the delegation is closed.
func (p *proxy) DelegatedAmountExtended(dl *types.Delegator) (*big.Int, *big.Int, error) {
	// log the action
	p.log.Debugf("loading extended delegation amounts for [%s]", dl.Address.String())

	// proxy the request directly to RPC/SFC
	return p.rpc.DelegatedAmountExtended(dl)
}

// TotalStaked calculates current total staked amount for all stakers.
func (p *proxy) TotalStaked() (*hexutil.Big, error) {
	// try cache first
	value := p.cache.PullTotalStaked()
	if value != nil {
		p.log.Debugf("total staked amount loaded from memory cache")
		return value, nil
	}

	// get the top staker
	topId, err := p.rpc.LastStakerId()
	if err != nil {
		p.log.Errorf("can not get the last staker; %s", err.Error())
		return nil, err
	}

	// make new accumulator
	total := new(big.Int)

	// go over all the validators
	var id uint64
	for id = 1; id <= uint64(topId); id++ {
		// get this validator
		val, err := p.rpc.Staker(hexutil.Uint64(id))
		if err == nil && val.TotalStake != nil {
			// advance the total sum
			total = new(big.Int).Add(total, val.TotalStake.ToInt())
		}
	}

	// store in cache
	if err := p.cache.PushTotalStaked((*hexutil.Big)(total)); err != nil {
		// log issue
		p.log.Errorf("can not store total staked amount in memory; %s", err.Error())
	}

	// return the value
	return (*hexutil.Big)(total), nil
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

		// return the block
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
		// inform what we do
		p.log.Errorf("can not get data of the last sealed epoch; %s", err.Error())
		return nil, err
	}

	// try to store the block in cache for future use
	err = p.cache.PushLastEpoch(&ep)
	if err != nil {
		p.log.Error(err)
	}

	// inform what we do
	p.log.Debugf("epoch [%s] loaded from sfc", id.String())
	return &ep, nil
}

// RewardsAllowed returns the reward lock status from SFC.
func (p *proxy) RewardsAllowed() (bool, error) {
	return p.rpc.RewardsAllowed()
}

// RewardsStash returns the amount of WEI stashed for the given address.
func (p *proxy) RewardsStash(addr *common.Address) (*big.Int, error) {
	return p.rpc.Stashed(*addr, big.NewInt(0))
}

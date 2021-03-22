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

// SfcVersion returns current version of the SFC contract.
func (p *proxy) SfcVersion() (hexutil.Uint64, error) {
	return p.rpc.SfcVersion()
}

// CurrentEpoch returns the id of the current epoch.
func (p *proxy) CurrentEpoch() (hexutil.Uint64, error) {
	return p.rpc.CurrentEpoch()
}

// Epoch returns the structure of the current epoch.
func (p *proxy) Epoch(id *hexutil.Uint64) (types.Epoch, error) {
	// get the current epoch if the id has not been provided
	if id == nil || *id == 0 {
		// ask for the sealed epoch id
		val, err := p.rpc.CurrentSealedEpoch()
		if err != nil {
			return types.Epoch{}, err
		}

		// use this id instead of nil
		id = &val
	}

	return p.rpc.Epoch(*id)
}

// LastStakerId returns the last staker id in Opera blockchain.
func (p *proxy) LastStakerId() (hexutil.Uint64, error) {
	return p.rpc.LastStakerId()
}

// StakersNum returns the number of stakers in Opera blockchain.
func (p *proxy) StakersNum() (hexutil.Uint64, error) {
	return p.rpc.StakersNum()
}

// IsStaker returns if the given address is an SFC staker.
func (p *proxy) IsStaker(addr *common.Address) (bool, error) {
	return p.rpc.IsStaker(addr)
}

// Staker extract a staker information from SFC smart contract.
func (p *proxy) Staker(id hexutil.Uint64) (*types.Staker, error) {
	return p.rpc.Staker(id)
}

// StakerAddress extract a staker address for the given staker ID.
func (p *proxy) StakerAddress(id hexutil.Uint64) (common.Address, error) {
	return p.rpc.StakerAddress(id)
}

// Staker extract a staker information by address.
func (p *proxy) StakerByAddress(addr common.Address) (*types.Staker, error) {
	return p.rpc.StakerByAddress(addr)
}

// IsDelegating returns if the given address is an SFC delegator.
func (p *proxy) IsDelegating(addr *common.Address) (bool, error) {
	return p.rpc.IsDelegating(addr)
}

// DelegationsOf extract a list of delegations for a given staker.
func (p *proxy) DelegationsOf(staker hexutil.Uint64) ([]types.Delegation, error) {
	return p.rpc.DelegationsOf(staker)
}

// DelegationsByAddress returns a list of all delegations
// of a given delegator address.
func (p *proxy) DelegationsByAddress(addr common.Address) ([]types.Delegation, error) {
	return p.rpc.DelegationsByAddress(addr)
}

// Delegation returns a detail of delegation for the given address.
func (p *proxy) Delegation(addr common.Address, staker hexutil.Uint64) (*types.Delegation, error) {
	return p.rpc.Delegation(addr, staker)
}

// DelegationLock returns delegation lock information using SFC contract binding.
func (p *proxy) DelegationLock(delegation *types.Delegation) (*types.DelegationLock, error) {
	return p.rpc.DelegationLock(delegation)
}

// Delegation returns a detail of delegation for the given address.
func (p *proxy) DelegationRewards(addr string, staker hexutil.Uint64) (types.PendingRewards, error) {
	p.log.Debugf("loading rewards of %s to %d", addr, staker)
	return p.rpc.DelegationRewards(addr, staker)
}

// WithdrawRequests extracts a list of partial withdraw requests
// for the given address.
func (p *proxy) WithdrawRequests(addr *common.Address, stakerId *hexutil.Uint64) ([]*types.WithdrawRequest, error) {
	// log the action
	p.log.Debugf("loading withdraw requests for [%s], staker %d", addr.String(), stakerId)

	// proxy the request directly to RPC/SFC with no staker filter
	if stakerId == nil {
		return p.rpc.WithdrawRequests(addr, nil)
	}

	// proxy the request directly to RPC/SFC, filter specific staker
	return p.rpc.WithdrawRequests(addr, new(big.Int).SetUint64(uint64(*stakerId)))
}

// DeactivatedDelegation extracts a list of deactivated delegation requests
// for the given address.
func (p *proxy) DeactivatedDelegation(addr *common.Address, stakerId *hexutil.Uint64) ([]*types.DeactivatedDelegation, error) {
	// log the action
	p.log.Debugf("loading deactivated delegation requests for [%s], staker %d", addr.String(), stakerId)

	// proxy the request directly to RPC/SFC, no staker
	if stakerId == nil {
		return p.rpc.DeactivatedDelegation(addr, nil)
	}

	// proxy the request directly to RPC/SFC, filter specific staker only
	return p.rpc.DeactivatedDelegation(addr, new(big.Int).SetUint64(uint64(*stakerId)))
}

// delegatedAmount calculates total amount currently delegated
// and amount locked in pending un-delegation.
// Partial Un-delegations are subtracted during the preparation
// phase, but total un-delegations are subtracted only when
// the delegation is closed.
func (p *proxy) DelegatedAmountExtended(dl *types.Delegation) (*big.Int, *big.Int, error) {
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

// LockingAllowed indicates if the stake locking has been enabled in SFC.
func (p *proxy) LockingAllowed() (bool, error) {
	return p.rpc.LockingAllowed()
}

// RewardsStash returns the amount of WEI stashed for the given address.
func (p *proxy) RewardsStash(addr *common.Address) (*big.Int, error) {
	return p.rpc.Stashed(*addr, big.NewInt(0))
}

// DelegationFluidStakingActive signals if the delegation is upgraded to Fluid Staking model.
func (p *proxy) DelegationFluidStakingActive(dl *types.Delegation) (bool, error) {
	return p.rpc.DelegationFluidStakingActive(dl)
}

// DelegationPaidUntilEpoch resolves the id of the last epoch rewards has been paid to."
func (p *proxy) DelegationPaidUntilEpoch(dl *types.Delegation) (hexutil.Uint64, error) {
	return p.rpc.DelegationPaidUntilEpoch(dl)
}

// DelegationOutstandingSFTM returns the amount of sFTM tokens for the delegation
// identified by the delegator address and the stakerId.
func (p *proxy) DelegationOutstandingSFTM(addr *common.Address, toStaker *hexutil.Uint64) (hexutil.Big, error) {
	return p.rpc.DelegationOutstandingSFTM(addr, toStaker)
}

// DelegationTokenizerUnlocked returns the status of SFC Tokenizer lock
// for a delegation identified by the address and staker id.
func (p *proxy) DelegationTokenizerUnlocked(addr *common.Address, toStaker *hexutil.Uint64) (bool, error) {
	return p.rpc.DelegationTokenizerUnlocked(addr, toStaker)
}

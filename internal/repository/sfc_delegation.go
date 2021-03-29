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
	"go.mongodb.org/mongo-driver/bson"
)

// IsDelegating returns if the given address is an SFC delegator.
func (p *proxy) IsDelegating(addr *common.Address) (bool, error) {
	count, err := p.db.DelegationsCountFiltered(&bson.D{{"addr", addr.String()}})
	if err != nil {
		p.log.Errorf("can not check delegation by address; %s", addr.String())
		return false, err
	}
	return 0 < count, nil
}

// DelegationsOfValidator extract a list of delegations for a given validator.
func (p *proxy) DelegationsOfValidator(valID *hexutil.Big, cursor *string, count int32) (*types.DelegationList, error) {
	p.log.Debugf("loading delegations of #%d", valID.ToInt().Uint64())
	return p.db.Delegations(cursor, count, &bson.D{{"to", valID.String()}})
}

// DelegationsByAddress returns a list of all delegations of a given delegator address.
func (p *proxy) DelegationsByAddress(addr *common.Address, cursor *string, count int32) (*types.DelegationList, error) {
	p.log.Debugf("loading delegations of %s", addr.String())
	return p.db.Delegations(cursor, count, &bson.D{{"addr", addr.String()}})
}

// StoreDelegation stores the delegation in persistent database.
func (p *proxy) StoreDelegation(dl *types.Delegation) error {
	return p.db.AddDelegation(dl)
}

// Delegation returns a detail of delegation for the given address.
func (p *proxy) Delegation(addr *common.Address, valID *hexutil.Big) (*types.Delegation, error) {
	p.log.Debugf("loading delegation of %s to #%d", addr.String(), valID.ToInt().Uint64())
	return p.db.Delegation(addr, valID)
}

// DelegationLock returns delegation lock information using SFC contract binding.
func (p *proxy) DelegationLock(dlg *types.Delegation) (*types.DelegationLock, error) {
	p.log.Debugf("loading lock information for %s to #%d", dlg.Address.String(), dlg.ToStakerId.ToInt().Uint64())
	return p.rpc.DelegationLock(dlg)
}

// PendingRewards returns a detail of pending rewards for the given delegation address and validator ID.
func (p *proxy) PendingRewards(addr *common.Address, valID *hexutil.Big) (*types.PendingRewards, error) {
	p.log.Debugf("loading pending rewards of %s to #%d", addr.String(), valID.ToInt().Uint64())
	return p.rpc.PendingRewards(addr, valID.ToInt())
}

// DelegationOutstandingSFTM returns the amount of sFTM tokens for the delegation
// identified by the delegator address and the stakerId.
func (p *proxy) DelegationOutstandingSFTM(addr *common.Address, toStaker *hexutil.Big) (*hexutil.Big, error) {
	val, err := p.rpc.DelegationOutstandingSFTM(addr, toStaker.ToInt())
	if err != nil {
		return nil, err
	}
	return (*hexutil.Big)(val), nil
}

// DelegationTokenizerUnlocked returns the status of SFC Tokenizer lock
// for a delegation identified by the address and staker id.
func (p *proxy) DelegationTokenizerUnlocked(addr *common.Address, toStaker *hexutil.Big) (bool, error) {
	return p.rpc.DelegationTokenizerUnlocked(addr, toStaker.ToInt())
}

// DelegationFluidStakingActive signals if the delegation is upgraded to Fluid Staking model.
func (p *proxy) DelegationFluidStakingActive(dl *types.Delegation) (bool, error) {
	return p.rpc.DelegationFluidStakingActive(dl)
}

/*
Repository package implements repository for handling fast and efficient access to data required
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

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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// LastValidatorId returns the last staker id in Opera blockchain.
func (p *proxy) LastValidatorId() (uint64, error) {
	return p.rpc.LastValidatorId()
}

// ValidatorsCount returns the number of stakers in Opera blockchain.
func (p *proxy) ValidatorsCount() (uint64, error) {
	return p.rpc.ValidatorsCount()
}

// IsValidator returns if the given address is an SFC staker.
func (p *proxy) IsValidator(addr *common.Address) (bool, error) {
	return p.rpc.IsValidator(addr)
}

// ValidatorAddress extract a staker address for the given staker ID.
func (p *proxy) ValidatorAddress(id *hexutil.Big) (*common.Address, error) {
	// try to use cache
	adr := p.cache.PullValidatorAddress(id)
	if nil != adr {
		return adr, nil
	}

	// pull from SFC
	adr, err := p.rpc.ValidatorAddress((*big.Int)(id))
	if err != nil {
		return nil, err
	}

	// store to cache for future use and return the value we got
	p.cache.PushValidatorAddress(id, adr)
	return adr, nil
}

// Validator extract a staker information from SFC smart contract.
func (p *proxy) Validator(id *hexutil.Big) (*types.Validator, error) {
	return p.rpc.Validator((*big.Int)(id))
}

// ValidatorByAddress extract a staker information by address.
func (p *proxy) ValidatorByAddress(addr *common.Address) (*types.Validator, error) {
	return p.rpc.ValidatorByAddress(addr)
}

// SfcMaxDelegatedRatio extracts a ratio between self delegation and received stake.
func (p *proxy) SfcMaxDelegatedRatio() (*big.Int, error) {
	// try cache first
	val := p.cache.PullSfcMaxDelegatedRatio()
	if val != nil {
		return val, nil
	}

	// pull from the SFC contract
	val, err := p.rpc.SfcMaxDelegatedRatio()
	if err != nil {
		return nil, err
	}

	// store for future use
	p.cache.PushSfcMaxDelegatedRatio(val)
	return val, nil
}

// ValidatorDowntime pulls information about validator downtime from the RPC interface.
func (p *proxy) ValidatorDowntime(valID *hexutil.Big) (uint64, uint64, error) {
	return p.rpc.ValidatorDowntime(valID)
}

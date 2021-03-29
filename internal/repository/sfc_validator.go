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

// LastValidatorId returns the last staker id in Opera blockchain.
func (p *proxy) LastValidatorId() (uint64, error) {
	return p.rpc.LastValidatorId()
}

// ValidatorsCount returns the number of stakers in Opera blockchain.
func (p *proxy) ValidatorsCount() (uint64, error) {
	return p.rpc.ValidatorsCount()
}

// IsStaker returns if the given address is an SFC staker.
func (p *proxy) IsValidator(addr *common.Address) (bool, error) {
	return p.rpc.IsValidator(addr)
}

// StakerAddress extract a staker address for the given staker ID.
func (p *proxy) ValidatorAddress(id hexutil.Uint64) (*common.Address, error) {
	return p.rpc.ValidatorAddress(new(big.Int).SetUint64(uint64(id)))
}

// Staker extract a staker information from SFC smart contract.
func (p *proxy) Validator(id hexutil.Uint64) (*types.Validator, error) {
	return p.rpc.Validator(new(big.Int).SetUint64(uint64(id)))
}

// Staker extract a staker information by address.
func (p *proxy) ValidatorByAddress(addr *common.Address) (*types.Validator, error) {
	return p.rpc.ValidatorByAddress(addr)
}

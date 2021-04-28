// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// CurrentEpoch resolves the id of the current epoch of the Opera blockchain.
func (rs *rootResolver) CurrentEpoch() (hexutil.Uint64, error) {
	return repository.R().CurrentEpoch()
}

// LastStakerId resolves the last staker id in Opera blockchain.
func (rs *rootResolver) LastStakerId() (hexutil.Uint64, error) {
	val, err := repository.R().LastValidatorId()
	if err != nil {
		return 0, err
	}
	return hexutil.Uint64(val), nil
}

// StakersNum resolves the number of stakers in Opera blockchain.
func (rs *rootResolver) StakersNum() (hexutil.Uint64, error) {
	val, err := repository.R().ValidatorsCount()
	if err != nil {
		return 0, err
	}
	return hexutil.Uint64(val), nil
}

// Staker resolves a validator information from SFC smart contract.
func (rs *rootResolver) Staker(args struct {
	Id      *hexutil.Big
	Address *common.Address
}) (*Staker, error) {
	// by ID or by address?
	if args.Id != nil {
		st, err := repository.R().Validator(args.Id)
		if err != nil {
			return nil, err
		}
		return NewStaker(st), err
	}

	st, err := repository.R().ValidatorByAddress(args.Address)
	if err != nil {
		return nil, err
	}
	return NewStaker(st), err
}

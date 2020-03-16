// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// CurrentEpoch resolves the id of the current epoch of the Opera blockchain.
func (rs *rootResolver) CurrentEpoch() (hexutil.Uint64, error) {
	return rs.repo.CurrentEpoch()
}

// Epoch resolves information about epoch of the given id.
func (rs *rootResolver) Epoch(args *struct{ Id hexutil.Uint64 }) (types.Epoch, error) {
	return rs.repo.Epoch(args.Id)
}

// Resolves the last staker id in Opera blockchain.
func (rs *rootResolver) LastStakerId() (hexutil.Uint64, error) {
	return rs.repo.LastStakerId()
}

// Resolves the number of stakers in Opera blockchain.
func (rs *rootResolver) StakersNum() (hexutil.Uint64, error) {
	return rs.repo.StakersNum()
}

// Resolves a staker information from SFC smart contract.
func (rs *rootResolver) Staker(args *struct {
	Id      *hexutil.Uint64
	Address *common.Address
}) (*Staker, error) {
	// by id
	if args.Id != nil {
		st, err := rs.repo.Staker(*args.Id)
		if err != nil {
			return nil, err
		}

		return NewStaker(st, rs.repo), nil
	}

	// by address
	if args.Address != nil {
		st, err := rs.repo.StakerByAddress(*args.Address)
		if err != nil {
			return nil, err
		}

		return NewStaker(st, rs.repo), nil
	}

	// return error
	return nil, fmt.Errorf("missing staker id or address")
}

// Resolves a staker information from SFC smart contract.
func (rs *rootResolver) Stakers() ([]Staker, error) {
	// get the number
	num, err := rs.repo.LastStakerId()
	if err != nil {
		rs.log.Errorf("can not get the highest staker id; %s", err.Error())
		return nil, err
	}

	// make the list
	list := make([]Staker, num)

	// loop to get the data
	for i := uint64(0); i < uint64(num); i++ {
		// extract the staker info
		st, err := rs.repo.Staker(hexutil.Uint64(i + 1))
		if err != nil {
			rs.log.Criticalf("can not extract staker information; %s", err.Error())
		} else {
			// store to the list
			list[i] = *NewStaker(st, rs.repo)
		}
	}

	return list, nil
}

// Resolves a list of delegations information of a staker.
func (rs *rootResolver) DelegationsOf(args *struct{ Staker hexutil.Uint64 }) ([]types.Delegator, error) {
	return rs.repo.DelegationsOf(args.Staker)
}

// Delegation resolves details of a delegator by it's address.
func (rs *rootResolver) Delegation(args *struct{ Address common.Address }) (*types.Delegator, error) {
	return rs.repo.Delegation(args.Address)
}

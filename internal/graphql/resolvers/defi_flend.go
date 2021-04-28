// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum/common"
)

// LendingPool represents a resolvable object Lending pool.
type LendingPool struct {
}

// FLendLendingPool resolves lending pool instance
func (rs *rootResolver) FLendLendingPool() (*LendingPool, error) {
	return &LendingPool{}, nil
}

// ReserveData resolves asset reserve data from lending pool
func (lp *LendingPool) ReserveData(args *struct{ Address common.Address }) (*types.ReserveData, error) {
	return repository.R().FLendGetLendingPoolReserveData(&args.Address)
}

// ReserveList resolves list of assets in lending pool
func (lp *LendingPool) ReserveList() ([]common.Address, error) {
	return repository.R().FLendGetReserveList()
}

// ReserveDataList resolves list of assets data in lending pool
func (lp *LendingPool) ReserveDataList() ([]*types.ReserveData, error) {
	// get the list
	rl, err := repository.R().FLendGetReserveList()
	if err != nil {
		return nil, err
	}

	// make the container
	rdl := make([]*types.ReserveData, len(rl))
	for i, adr := range rl {
		rdl[i], err = repository.R().FLendGetLendingPoolReserveData(&adr)
		if err != nil {
			return nil, err
		}
	}
	return rdl, nil
}

// UserAccountData resolves user account data from lending pool
func (lp *LendingPool) UserAccountData(args *struct{ Address common.Address }) (*types.FLendUserAccountData, error) {
	return repository.R().FLendGetUserAccountData(&args.Address)
}

// UserDepositHistory resolves user account deposit history data from lending pool
func (lp *LendingPool) UserDepositHistory(args *struct {
	Address *common.Address
	Asset   *common.Address
}) ([]*types.FLendDeposit, error) {
	return repository.R().FLendGetUserDepositHistory(args.Address, args.Asset)
}

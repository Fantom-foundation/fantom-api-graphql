// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum/common"
)

// LendingPool represents a resolvable object Lending pool.
type LendingPool struct {
	repo repository.Repository
}

// FLendLendingPool resolves lending pool instance
func (rs *rootResolver) FLendLendingPool() (*LendingPool, error) {
	return &LendingPool{rs.repo}, nil
}

// ReserveData resolves asset reserve data from lending pool
func (lp *LendingPool) ReserveData(args *struct{ Address common.Address }) (*types.ReserveData, error) {

	rd, err := lp.repo.FLendGetLendingPoolReserveData(&args.Address)
	if err != nil {
		lp.repo.Log().Errorf("Can not get lending pool reserve data for address %s: %s", args.Address.String(), err.Error())
		return nil, err
	}
	return rd, nil
}

// ReserveList resolves list of assets in lending pool
func (lp *LendingPool) ReserveList() ([]common.Address, error) {

	rl, err := lp.repo.FLendGetReserveList()
	if err != nil {
		lp.repo.Log().Errorf("Can not get lending pool reserves list: %s", err.Error())
		return nil, err
	}
	return rl, nil
}

// ReserveDataList resolves list of assets data in lending pool
func (lp *LendingPool) ReserveDataList() ([]*types.ReserveData, error) {

	rl, err := lp.repo.FLendGetReserveList()
	if err != nil {
		lp.repo.Log().Errorf("Can not get lending pool reserves list: %s", err.Error())
		return nil, err
	}

	rdl := make([]*types.ReserveData, 0)
	for _, reserveData := range rl {
		rd, err := lp.repo.FLendGetLendingPoolReserveData(&reserveData)
		if err != nil {
			lp.repo.Log().Errorf("Can not get lending pool reserve data for address %s: %s", reserveData.String(), err.Error())
			return nil, err
		}
		rdl = append(rdl, rd)
	}
	return rdl, nil
}

// UserAccountData resolves user account data from lending pool
func (lp *LendingPool) UserAccountData(args *struct{ Address common.Address }) (*types.FLendUserAccountData, error) {

	uad, err := lp.repo.FLendGetUserAccountData(&args.Address)
	if err != nil {
		lp.repo.Log().Errorf("Can not get lending pool user account data for address %s: %s", args.Address.String(), err.Error())
		return nil, err
	}
	return uad, nil
}

// UserDepositHistory resolves user account deposit history data from lending pool
func (lp *LendingPool) UserDepositHistory(args *struct {
	Address *common.Address
	Asset   *common.Address
}) ([]*types.FLendDeposit, error) {

	uad, err := lp.repo.FLendGetUserDepositHistory(args.Address, args.Asset)
	if err != nil {
		lp.repo.Log().Errorf("Can not get lending pool deposit history data for user %s, asset %s: %s", args.Address.String(), args.Asset.String(), err.Error())
		return nil, err
	}
	return uad, nil
}

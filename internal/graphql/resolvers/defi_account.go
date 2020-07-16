// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// defiTokenBalanceTypeCollateral is the name of collateral type token balance.
const defiTokenBalanceTypeCollateral = "COLLATERAL"

// defiTokenBalanceTypeDebt is the name of debt type token balance.
const defiTokenBalanceTypeDebt = "DEBT"

// DefiAccount represents resolvable DeFi account information.
type DefiAccount struct {
	repo repository.Repository
	types.DefiAccount
}

// DefiTokenBalance represents a resolvable DeFi token balance information.
type DefiTokenBalance struct {
	repo         repository.Repository
	OwnerAddress common.Address
	TokenAddress common.Address
	Type         string
}

// NewDefiAccount creates new instance of resolvable DeFi account.
func NewDefiAccount(da *types.DefiAccount, repo repository.Repository) *DefiAccount {
	return &DefiAccount{
		DefiAccount: *da,
		repo:        repo,
	}
}

// NewDefiTokenBalance creates a new DeFi token balance, either collateral, or debt.
func NewDefiTokenBalance(owner common.Address, token common.Address, tp string, repo repository.Repository) *DefiTokenBalance {
	return &DefiTokenBalance{
		repo:         repo,
		Type:         tp,
		OwnerAddress: owner,
		TokenAddress: token,
	}
}

// DefiAccount resolves details of a DeFi account by it's address.
func (rs *rootResolver) DefiAccount(args *struct{ Owner common.Address }) (*DefiAccount, error) {
	// get the delegator detail from backend
	da, err := rs.repo.DefiAccount(args.Owner)
	if err != nil {
		return nil, err
	}

	return NewDefiAccount(da, rs.repo), nil
}

// Collateral resolves the list of collateral token balance containers.
func (da *DefiAccount) Collateral() []*DefiTokenBalance {
	// make the container
	list := make([]*DefiTokenBalance, 0)

	// loop all the collateral addresses and make the token balance record
	for _, token := range da.CollateralList {
		list = append(list, NewDefiTokenBalance(da.Address, token, defiTokenBalanceTypeCollateral, da.repo))
	}

	return list
}

// Debt resolves the list of debt token balance containers.
func (da *DefiAccount) Debt() []*DefiTokenBalance {
	// make the container
	list := make([]*DefiTokenBalance, 0)

	// loop all the collateral addresses and make the token balance record
	for _, token := range da.DebtList {
		list = append(list, NewDefiTokenBalance(da.Address, token, defiTokenBalanceTypeDebt, da.repo))
	}

	return list
}

// Token resolves the token information from the related token address.
func (dtb *DefiTokenBalance) Token() (*types.DefiToken, error) {
	return dtb.repo.DefiToken(&dtb.TokenAddress)
}

// Balance resolves the balance of the token for the related token address.
func (dtb *DefiTokenBalance) Balance() (hexutil.Big, error) {
	return dtb.repo.DefiTokenBalance(&dtb.OwnerAddress, &dtb.TokenAddress, dtb.Type)
}

// Value resolves the value of the token for the related token address in fUSD.
func (dtb *DefiTokenBalance) Value() (hexutil.Big, error) {
	return dtb.repo.DefiTokenValue(&dtb.OwnerAddress, &dtb.TokenAddress, dtb.Type)
}

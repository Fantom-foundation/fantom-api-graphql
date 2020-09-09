// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// FMintAccount represents resolvable DeFi account information.
type FMintAccount struct {
	repo repository.Repository
	types.FMintAccount
}

// FMintTokenBalance represents a resolvable DeFi token balance information.
type FMintTokenBalance struct {
	repo         repository.Repository
	OwnerAddress common.Address
	TokenAddress common.Address
	Type         types.DefiTokenType
}

// NewFMintAccount creates new instance of resolvable DeFi account.
func NewFMintAccount(ac *types.FMintAccount, repo repository.Repository) *FMintAccount {
	return &FMintAccount{
		FMintAccount: *ac,
		repo:         repo,
	}
}

// NewFMintTokenBalance creates a new DeFi token balance, either collateral, or debt.
func NewFMintTokenBalance(owner common.Address, token common.Address, tp types.DefiTokenType, repo repository.Repository) *FMintTokenBalance {
	return &FMintTokenBalance{
		repo:         repo,
		Type:         tp,
		OwnerAddress: owner,
		TokenAddress: token,
	}
}

// FMintAccount resolves details of a DeFi account by it's address.
func (rs *rootResolver) FMintAccount(args *struct{ Owner common.Address }) (*FMintAccount, error) {
	// get the delegator detail from backend
	ac, err := rs.repo.FMintAccount(args.Owner)
	if err != nil {
		return nil, err
	}

	return NewFMintAccount(ac, rs.repo), nil
}

// Collateral resolves the list of collateral token balance containers.
func (fac *FMintAccount) Collateral() []*FMintTokenBalance {
	// make the container
	list := make([]*FMintTokenBalance, 0)

	// loop all the collateral addresses and make the token balance record
	for _, token := range fac.CollateralList {
		list = append(list, NewFMintTokenBalance(fac.Address, token, types.DefiTokenTypeCollateral, fac.repo))
	}

	return list
}

// Debt resolves the list of debt token balance containers.
func (fac *FMintAccount) Debt() []*FMintTokenBalance {
	// make the container
	list := make([]*FMintTokenBalance, 0)

	// loop all the collateral addresses and make the token balance record
	for _, token := range fac.DebtList {
		list = append(list, NewFMintTokenBalance(fac.Address, token, types.DefiTokenTypeDebt, fac.repo))
	}

	return list
}

// TotalCollateralRewardsAmount resolves the total amount of rewards
// accumulated on the account for the excessive collateral deposits.
func (fac *FMintAccount) TotalCollateralRewardsAmount() hexutil.Big {
	return hexutil.Big{}
}

// Token resolves the token information from the related token address.
func (mintBalance *FMintTokenBalance) Token() (*DefiToken, error) {
	// get the token backend
	tk, err := mintBalance.repo.DefiToken(&mintBalance.TokenAddress)
	if err != nil {
		return nil, err
	}

	// resolve the token
	return NewDefiToken(tk, mintBalance.repo), nil
}

// Balance resolves the balance of the token for the related token address.
func (mintBalance *FMintTokenBalance) Balance() (hexutil.Big, error) {
	return mintBalance.repo.FMintTokenBalance(&mintBalance.OwnerAddress, &mintBalance.TokenAddress, mintBalance.Type)
}

// Value resolves the value of the token for the related token address in fUSD.
func (mintBalance *FMintTokenBalance) Value() (hexutil.Big, error) {
	return mintBalance.repo.FMintTokenValue(&mintBalance.OwnerAddress, &mintBalance.TokenAddress, mintBalance.Type)
}

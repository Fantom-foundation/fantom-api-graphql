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
	types.FMintAccount
}

// FMintTokenBalance represents a resolvable DeFi token balance information.
type FMintTokenBalance struct {
	OwnerAddress common.Address
	TokenAddress common.Address
	Type         types.DefiTokenType
}

// NewFMintAccount creates new instance of resolvable DeFi account.
func NewFMintAccount(ac *types.FMintAccount) *FMintAccount {
	return &FMintAccount{FMintAccount: *ac}
}

// NewFMintTokenBalance creates a new DeFi token balance, either collateral, or debt.
func NewFMintTokenBalance(owner common.Address, token common.Address, tp types.DefiTokenType) *FMintTokenBalance {
	return &FMintTokenBalance{
		Type:         tp,
		OwnerAddress: owner,
		TokenAddress: token,
	}
}

// FMintAccount resolves details of a DeFi account by its address.
func (rs *rootResolver) FMintAccount(args *struct{ Owner common.Address }) (*FMintAccount, error) {
	// get the delegator detail from backend
	ac, err := repository.R().FMintAccount(args.Owner)
	if err != nil {
		return nil, err
	}

	return NewFMintAccount(ac), nil
}

// Collateral resolves the list of collateral token balance containers.
func (fac *FMintAccount) Collateral() []*FMintTokenBalance {
	// prep container and loop all the collateral addresses
	// to create the token balance record for each
	list := make([]*FMintTokenBalance, len(fac.CollateralList))
	for i, token := range fac.CollateralList {
		list[i] = NewFMintTokenBalance(fac.Address, token, types.DefiTokenTypeCollateral)
	}
	return list
}

// Debt resolves the list of debt token balance containers.
func (fac *FMintAccount) Debt() []*FMintTokenBalance {
	// make the container
	// loop all the collateral addresses and make the token balance record
	list := make([]*FMintTokenBalance, len(fac.DebtList))
	for i, token := range fac.DebtList {
		list[i] = NewFMintTokenBalance(fac.Address, token, types.DefiTokenTypeDebt)
	}
	return list
}

// RewardsEarned resolves the total amount of rewards
// accumulated on the account for the excessive collateral deposits.
func (fac *FMintAccount) RewardsEarned() (hexutil.Big, error) {
	return repository.R().FMintRewardsEarned(&fac.Address)
}

// RewardsStashed resolves the total amount of rewards
// accumulated on the account in the stash.
func (fac *FMintAccount) RewardsStashed() (hexutil.Big, error) {
	return repository.R().FMintRewardsStashed(&fac.Address)
}

// CanClaimRewards resolves the fMint account flag for being allowed
// to claim earned rewards.
func (fac *FMintAccount) CanClaimRewards() (bool, error) {
	return repository.R().FMintCanClaimRewards(&fac.Address)
}

// CanReceiveRewards resolves the fMint account flag for being eligible
// to receive earned rewards. If the collateral to debt ration drop below
// certain value, earned rewards are burned.
func (fac *FMintAccount) CanReceiveRewards() (bool, error) {
	return repository.R().FMintCanReceiveRewards(&fac.Address)
}

// CanPushNewRewards resolves the flag about the new rewards unlocked
// and ready for push.
func (fac *FMintAccount) CanPushNewRewards() (bool, error) {
	return repository.R().FMintCanPushRewards()
}

// Token resolves the token information from the related token address.
func (mb *FMintTokenBalance) Token() (*DefiToken, error) {
	// get the token backend
	tk, err := repository.R().DefiToken(&mb.TokenAddress)
	if err != nil {
		return nil, err
	}

	// resolve the token
	return NewDefiToken(tk), nil
}

// Balance resolves the balance of the token for the related token address.
func (mb *FMintTokenBalance) Balance() (hexutil.Big, error) {
	return repository.R().FMintTokenBalance(&mb.OwnerAddress, &mb.TokenAddress, mb.Type)
}

// Value resolves the value of the token for the related token address in fUSD.
func (mb *FMintTokenBalance) Value() (hexutil.Big, error) {
	return repository.R().FMintTokenValue(&mb.OwnerAddress, &mb.TokenAddress, mb.Type)
}

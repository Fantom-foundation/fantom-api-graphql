// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"math"
)

// define names of the fMint operations recognized by the protocol resolvers.
const (
	FMintPurposeCollateral = "FMINT_COLLATERAL"
	FMintPurposeDebt       = "FMINT_DEBT"
)

// FMintUserToken represents a resolvable pair of user and its fMint token
// used for a specific purpose.
type FMintUserToken struct {
	// Purpose represents the type of usage of the token by the user.
	Purpose string

	// UserAddress represents the address of the fMint user.
	UserAddress common.Address

	// TokenAddress represents the address of the associated token.
	TokenAddress common.Address
}

// fMintPurposeToType converts purpose to transaction type.
func fMintPurposeToType(p string) int32 {
	switch p {
	case FMintPurposeCollateral:
		return types.FMintTrxTypeDeposit
	case FMintPurposeDebt:
		return types.FMintTrxTypeMint
	default:
		log.Criticalf("unknown purpose %s", p)
		return math.MaxInt32
	}
}

// fMintTypeToPurpose converts transaction type to purpose.
func fMintTypeToPurpose(t int32) string {
	switch t {
	case types.FMintTrxTypeDeposit:
		return FMintPurposeCollateral
	case types.FMintTrxTypeWithdraw:
		return FMintPurposeCollateral
	case types.FMintTrxTypeMint:
		return FMintPurposeDebt
	case types.FMintTrxTypeRepay:
		return FMintPurposeDebt
	default:
		log.Criticalf("unknown fMint transaction type #%d", t)
		return FMintPurposeCollateral
	}
}

// FMintUserTokens resolves list of fMint users and associated tokens
// used for specified purpose.
func (rs *rootResolver) FMintUserTokens(args struct{ Purpose string }) ([]*FMintUserToken, error) {
	// get the aggregated list of addresses and their tokens
	list, err := repository.R().FMintUsers(fMintPurposeToType(args.Purpose))
	if err != nil {
		return nil, err
	}

	// collect all the users with their tokens into the output set
	out := make([]*FMintUserToken, 0)
	for _, usr := range list {
		for _, tok := range usr.Tokens {
			out = append(out, &FMintUserToken{
				Purpose:      fMintTypeToPurpose(usr.Purpose),
				UserAddress:  usr.User,
				TokenAddress: tok,
			})
		}
	}
	return out, nil
}

// Account resolves account of the fMint user.
func (fut *FMintUserToken) Account() (*FMintAccount, error) {
	// get the delegator detail from backend
	ac, err := repository.R().FMintAccount(fut.UserAddress)
	if err != nil {
		return nil, err
	}

	return NewFMintAccount(ac), nil
}

// Token resolves the detail of the associated ERC20 token.
func (fut *FMintUserToken) Token() *ERC20Token {
	return NewErc20Token(&fut.TokenAddress)
}

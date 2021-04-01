// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
)

// WithdrawRequest represents resolvable partial withdraw request
// from either stake or delegation structure.
type WithdrawRequest struct {
	types.WithdrawRequest
}

// NewWithdrawRequest builds new resolvable partial withdraw request structure.
func NewWithdrawRequest(wr *types.WithdrawRequest) WithdrawRequest {
	return WithdrawRequest{WithdrawRequest: *wr}
}

// Account resolves the account detail of the partial withdraw request.
func (wr WithdrawRequest) Account() (*Account, error) {
	// get the account detail by address
	acc, err := repository.R().Account(&wr.Address)
	if err != nil {
		return nil, err
	}

	// return the account detail
	return NewAccount(acc), nil
}

// Staker resolves the withdraw request staker detail, if available.
func (wr WithdrawRequest) Staker() (*Staker, error) {
	// get staker detail by the staker id
	st, err := repository.R().Validator(wr.StakerID)
	if err != nil {
		return nil, err
	}

	// return the staker information
	return NewStaker(st), nil
}

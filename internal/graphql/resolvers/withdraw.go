// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

// Id resolves unique internal identifier of the Withdraw request.
func (wr WithdrawRequest) Id() Cursor {
	return Cursor(wr.RequestTrx.String())
}

// WithdrawRequestID resolves the SFC identifier of the request
// unique to the delegation address and the target validator ID.
func (wr WithdrawRequest) WithdrawRequestID() hexutil.Big {
	if wr.WithdrawRequest.WithdrawRequestID == nil {
		return hexutil.Big{}
	}
	return *wr.WithdrawRequest.WithdrawRequestID
}

// StakerID resolves the identifier of the validator ID the request points to.
func (wr WithdrawRequest) StakerID() hexutil.Big {
	if wr.WithdrawRequest.StakerID == nil {
		return hexutil.Big{}
	}
	return *wr.WithdrawRequest.StakerID
}

// Amount resolves the amount of tokens the withdraw request is for.
func (wr WithdrawRequest) Amount() hexutil.Big {
	if wr.WithdrawRequest.Amount == nil {
		return hexutil.Big{}
	}
	return *wr.WithdrawRequest.Amount
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
	st, err := repository.R().Validator(wr.WithdrawRequest.StakerID)
	if err != nil {
		return nil, err
	}

	// return the staker information
	return NewStaker(st), nil
}

// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
)

// WithdrawRequest represents resolvable partial withdraw request
// from either stake or delegation structure.
type WithdrawRequest struct {
	repo repository.Repository
	types.WithdrawRequest
}

// NewWithdrawRequest builds new resolvable partial withdraw request structure.
func NewWithdrawRequest(wr *types.WithdrawRequest, repo repository.Repository) WithdrawRequest {
	return WithdrawRequest{
		repo:            repo,
		WithdrawRequest: *wr,
	}
}

// Account resolves the account detail of the partial withdraw request.
func (wr WithdrawRequest) Account() (*Account, error) {
	// get the account detail by address
	acc, err := wr.repo.Account(&wr.Address)
	if err != nil {
		return nil, err
	}

	// return the account detail
	return NewAccount(acc, wr.repo), nil
}

// Staker resolves the withdraw request staker detail, if available.
func (wr WithdrawRequest) Staker() (*Staker, error) {
	// get staker detail by the staker id
	st, err := wr.repo.Staker(wr.StakerID)
	if err != nil {
		return nil, err
	}

	// return the staker information
	return NewStaker(st, wr.repo), nil
}

// RequestBlock resolves the withdraw request registration block details.
func (wr WithdrawRequest) RequestBlock() (*Block, error) {
	// get the block details by it's identifier
	blk, err := wr.repo.BlockByNumber(&wr.RequestBlockNumber)
	if err != nil {
		return nil, err
	}

	// return the block details
	return NewBlock(blk, wr.repo), nil
}

// WithdrawBlock resolves the withdraw finalization block details.
// This is resolved only if the withdraw request has been already
// finalized and we know the finalization details.
func (wr WithdrawRequest) WithdrawBlock() (*Block, error) {
	// do we have the finalization block details?
	if wr.WithdrawBlockNumber == nil {
		return nil, nil
	}

	// get the block details by it's identifier
	blk, err := wr.repo.BlockByNumber(wr.WithdrawBlockNumber)
	if err != nil {
		return nil, err
	}

	// return the block details
	return NewBlock(blk, wr.repo), nil
}

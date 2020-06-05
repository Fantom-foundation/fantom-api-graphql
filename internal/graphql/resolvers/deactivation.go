// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
)

// DeactivatedDelegation represents resolvable delegation deactivation request
// from delegation structure.
type DeactivatedDelegation struct {
	repo repository.Repository
	types.DeactivatedDelegation
}

// NewDeactivatedDelegation builds new resolvable deactivated delegation request structure.
func NewDeactivatedDelegation(dd *types.DeactivatedDelegation, repo repository.Repository) DeactivatedDelegation {
	return DeactivatedDelegation{
		repo:                  repo,
		DeactivatedDelegation: *dd,
	}
}

// Account resolves the account detail of the deactivated delegation request.
func (dd DeactivatedDelegation) Account() (*Account, error) {
	// get the account detail by address
	acc, err := dd.repo.Account(&dd.Address)
	if err != nil {
		return nil, err
	}

	// return the account detail
	return NewAccount(acc, dd.repo), nil
}

// Staker resolves the deactivated delegation request staker detail, if available.
func (dd DeactivatedDelegation) Staker() (*Staker, error) {
	// get staker detail by the staker id
	st, err := dd.repo.Staker(dd.StakerID)
	if err != nil {
		return nil, err
	}

	// return the staker information
	return NewStaker(st, dd.repo), nil
}

// RequestBlock resolves the deactivated delegation request registration block details.
func (dd DeactivatedDelegation) RequestBlock() (*Block, error) {
	// get the block details by it's identifier
	blk, err := dd.repo.BlockByNumber(&dd.RequestBlockNumber)
	if err != nil {
		return nil, err
	}

	// return the block details
	return NewBlock(blk, dd.repo), nil
}

// WithdrawBlock resolves the withdraw finalization block details.
// This is resolved only if the deactivated delegation request has been already
// finalized and we know the finalization details.
func (dd DeactivatedDelegation) WithdrawBlock() (*Block, error) {
	// do we have the finalization block details?
	if dd.WithdrawBlockNumber == nil {
		return nil, nil
	}

	// get the block details by it's identifier
	blk, err := dd.repo.BlockByNumber(dd.WithdrawBlockNumber)
	if err != nil {
		return nil, err
	}

	// return the block details
	return NewBlock(blk, dd.repo), nil
}

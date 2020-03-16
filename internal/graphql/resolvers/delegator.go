package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
)

// Delegator represents resolvable delegator detail.
type Delegator struct {
	types.Delegator
	repo repository.Repository
}

// NewDelegator creates new instance of resolvable Delegator.
func NewDelegator(d *types.Delegator, repo repository.Repository) *Delegator {
	return &Delegator{
		Delegator: *d,
		repo:      repo,
	}
}

// PendingRewards resolves pending rewards for the delegator account.
func (del Delegator) PendingRewards() (types.PendingRewards, error) {
	// get the rewards
	r, err := del.repo.DelegationRewards(del.Address)
	if err != nil {
		return types.PendingRewards{}, err
	}

	return r, nil
}

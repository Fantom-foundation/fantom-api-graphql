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
	r, err := del.repo.DelegationRewards(del.Address.String())
	if err != nil {
		return types.PendingRewards{}, err
	}

	return r, nil
}

// DelegationsByAge represents a list of delegations sortable by their age of creation.
type DelegationsByAge []Delegator

// Len returns size of the delegation list.
func (d DelegationsByAge) Len() int {
	return len(d)
}

// Less compares two delegations and returns true if the first is lower than the last.
// We use it to sort delegations by time created having newer on top.
func (d DelegationsByAge) Less(i, j int) bool {
	return uint64(d[i].CreatedTime) > uint64(d[j].CreatedTime)
}

// Swap changes position of two delegations in the list.
func (d DelegationsByAge) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

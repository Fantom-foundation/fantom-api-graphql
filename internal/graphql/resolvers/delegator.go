// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
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

// Delegation resolves details of a delegator by it's address.
func (rs *rootResolver) Delegation(args *struct{ Address common.Address }) (*Delegator, error) {
	// get the delegator detail from backend
	d, err := rs.repo.Delegation(args.Address)
	if err != nil {
		return nil, err
	}

	return NewDelegator(d, rs.repo), nil
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
type DelegationsByAge []types.Delegator

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

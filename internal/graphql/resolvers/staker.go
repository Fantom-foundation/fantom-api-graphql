package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"sort"
)

// Staker represents  resolvable staker record.
type Staker struct {
	types.Staker
	repo repository.Repository
}

// NewStaker creates a new resolvable staker structure
func NewStaker(st *types.Staker, repo repository.Repository) *Staker {
	return &Staker{Staker: *st, repo: repo}
}

// Delegations resolves list of delegations associated with the staker.
func (st Staker) Delegations() ([]Delegator, error) {
	// get delegations
	dl, err := st.repo.DelegationsOf(st.Id)
	if err != nil {
		return nil, err
	}

	// make the list
	list := make([]Delegator, len(dl))

	// prep the list
	for i := 0; i < len(dl); i++ {
		list[i] = *NewDelegator(&dl[i], st.repo)
	}

	// sort by the date of creation
	sort.Sort(DelegationsByAge(list))
	return list, nil
}

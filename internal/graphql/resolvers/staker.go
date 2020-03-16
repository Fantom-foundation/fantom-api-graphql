package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
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

// Delegations resolves list of delegators associated with the staker.
func (st Staker) Delegations() ([]types.Delegator, error) {
	return st.repo.DelegationsOf(st.Id)
}

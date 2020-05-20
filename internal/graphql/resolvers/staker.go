// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"fmt"
	"sort"
)

// Staker represents resolvable staker record.
type Staker struct {
	types.Staker
	repo repository.Repository
}

// NewStaker creates a new resolvable staker structure
func NewStaker(st *types.Staker, repo repository.Repository) *Staker {
	return &Staker{Staker: *st, repo: repo}
}

// Delegations resolves list of delegations associated with the staker.
func (st Staker) Delegations(args *struct {
	Cursor *Cursor
	Count  int32
}) (*DelegatorList, error) {
	// any arguments?
	if args == nil {
		return nil, fmt.Errorf("missing delegations input")
	}

	// get delegations
	dl, err := st.repo.DelegationsOf(st.Id)
	if err != nil {
		return nil, err
	}

	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, accMaxTransactionsPerRequest)

	// sort by the date of creation
	sort.Sort(DelegationsByAge(dl))
	return NewDelegatorList(dl, parseDelegationsCursor(args.Cursor, args.Count, dl), args.Count, st.repo), nil
}

// StakerInfo resolves extended staker information if available.
func (st Staker) StakerInfo() *types.StakerInfo {
	return st.repo.RetrieveStakerInfo(st.Id)
}

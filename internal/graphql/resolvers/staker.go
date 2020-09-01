// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"fmt"
	"sort"
	"time"
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
}) (*DelegationList, error) {
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
	return NewDelegationList(dl, parseDelegationsCursor(args.Cursor, args.Count, dl), args.Count, st.repo), nil
}

// StakerInfo resolves extended staker information if available.
func (st Staker) StakerInfo() *types.StakerInfo {
	return st.repo.RetrieveStakerInfo(st.Id)
}

// IsStakeLocked signals if the stake is locked right now.
func (st Staker) IsStakeLocked() bool {
	return st.Staker.LockedFromEpoch > 0 && uint64(st.Staker.LockedUntil) < uint64(time.Now().UTC().Unix())
}

// WithdrawRequests resolves partial withdraw requests of the staker.
// We load withdraw requests of the stake only, not the stake delegators.
func (st Staker) WithdrawRequests() ([]WithdrawRequest, error) {
	// pull the requests list from remote server
	wr, err := st.repo.WithdrawRequests(&st.StakerAddress)
	if err != nil {
		return nil, err
	}

	// create new result set
	list := make([]WithdrawRequest, 0)

	// sort the list by age of the request
	// we want the newest requests on top
	sort.Sort(types.WithdrawRequestsByAge(wr))

	// iterate over the sorted list and populate the output array
	for _, req := range wr {
		list = append(list, NewWithdrawRequest(req, st.repo))
	}

	// return the final resolvable list
	return list, nil
}

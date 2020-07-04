// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

// Ballot represents resolvable ballot record.
type Ballot struct {
	types.Ballot
	isFinalized *bool
	repo        repository.Repository
}

// BallotProposal represents a single resolved ballot proposal.
type BallotProposal struct {
	Id   int32
	Name string
}

// NewBallot creates a new resolvable ballot structure
func NewBallot(bt *types.Ballot, repo repository.Repository) *Ballot {
	return &Ballot{Ballot: *bt, repo: repo}
}

// Ballot resolves details of a voting ballot by it's address.
func (rs *rootResolver) Ballot(args *struct{ Address common.Address }) (*Ballot, error) {
	// get the ballot for the address
	bl, err := rs.repo.BallotByAddress(&args.Address)
	if err != nil {
		return nil, err
	}

	// return the resolved ballot record
	return NewBallot(bl, rs.repo), nil
}

// IsOpen resolves open status of a ballot.
func (bt *Ballot) IsOpen() bool {
	// get the current time
	now := time.Now()
	return int64(bt.Start) < now.Unix() && int64(bt.End) > now.Unix()
}

// IsFinalized resolves finalized status of a ballot.
// If a ballot has been finalized, the winner is known.
func (bt *Ballot) IsFinalized() (bool, error) {
	// do we have the value pre-cached?
	if bt.isFinalized == nil {
		fin, err := bt.repo.BallotIsFinalized(&bt.Address)
		if err != nil {
			return false, err
		}

		// rewrite the value
		bt.isFinalized = &fin
	}

	return *bt.isFinalized, nil
}

// Winner resolves the winning proposal index, if the ballot
// has been finalized already. Returns nil if the ballot is pending.
func (bt *Ballot) Winner() (*int32, error) {
	// make sure the ballot is already finalized
	fin, err := bt.IsFinalized()
	if err != nil {
		return nil, err
	}

	// if the ballot is not resolved, do not dare
	if !fin {
		return nil, nil
	}

	// check the ballot winner
	win, err := bt.repo.BallotWinner(&bt.Address)
	if err != nil {
		return nil, err
	}

	// no winner
	if win == nil {
		return nil, nil
	}

	// convert
	val := int32(*win)
	return &val, nil
}

// Contract resolves managing smart contract of the ballot.
func (bt *Ballot) Contract() (*Contract, error) {
	// get a contract to be validated if any
	sc, err := bt.repo.Contract(&bt.Address)
	if err != nil {
		return nil, err
	}

	return NewContract(sc, bt.repo), nil
}

// Proposals resolves a list of available proposals
// e.g. options to vote for.
func (bt *Ballot) Proposals() ([]BallotProposal, error) {
	// make the container
	props := make([]BallotProposal, 0)

	// loop all proposals available
	for ix, prop := range bt.Ballot.Proposals {
		// hack the smart contract update after 7/1/2020
		if ix > 0 || uint64(bt.Ballot.End) < 1593561600 {
			props = append(props, BallotProposal{
				Id:   int32(ix),
				Name: prop,
			})
		}
	}

	return props, nil
}

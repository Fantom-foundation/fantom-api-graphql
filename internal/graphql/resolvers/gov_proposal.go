// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// GovernanceProposal represents details of a Governance Proposal.
type GovernanceProposal struct {
	// repo represents the reference to the server repository
	repo repository.Repository

	// the actual proposal details
	types.GovernanceProposal
}

// zeroInt represents an empty Big INT value used for comparison.
var zeroInt = new(big.Int)

// OptionState resolves a state of a given Proposal option identified
// by it's id (index position) in the Proposal options list.
func (gp *GovernanceProposal) OptionState(args *struct{ OptionId hexutil.Big }) (*types.GovernanceOptionState, error) {
	return gp.repo.GovernanceOptionState(&gp.GovernanceId, &gp.Id, &args.OptionId)
}

// OptionState resolves a state of a given Proposal option identified
// by it's id (index position) in the Proposal options list.
func (gp *GovernanceProposal) OptionStates() ([]*types.GovernanceOptionState, error) {
	return gp.repo.GovernanceOptionStates(&gp.GovernanceId, &gp.Id)
}

// Vote resolves the vote for the given <from> address linked
// with the <delegatedTo> delegation recipient.
func (gp *GovernanceProposal) Vote(args *struct {
	From        common.Address
	DelegatedTo *common.Address
}) (*types.GovernanceVote, error) {
	return gp.repo.GovernanceVote(&gp.GovernanceId, &gp.Id, &args.From, args.DelegatedTo)
}

// Governance resolves the parent Governance instance.
func (gp *GovernanceProposal) Governance() (*GovernanceContract, error) {
	// get the governance contract by address
	gc, err := gp.repo.GovernanceContractBy(&gp.GovernanceId)
	if err != nil {
		return nil, err
	}

	return &GovernanceContract{
		repo:    gp.repo,
		Name:    gc.Name,
		Type:    gc.Type,
		Address: gp.GovernanceId,
	}, nil
}

// State resolves the state of the Governance Proposal.
func (gp *GovernanceProposal) State() (*types.GovernanceProposalState, error) {
	// get the state
	return gp.repo.GovernanceProposalState(&gp.GovernanceId, &gp.Id)
}

// TotalWeight resolves the total available voting power which can influence
// the proposal outcome.
func (gp *GovernanceProposal) TotalWeight() (hexutil.Big, error) {
	return gp.repo.GovernanceTotalWeight(&gp.GovernanceId)
}

// VotedWeightRatio represents what percentage of the total voting power already
// placed a vote either directly, or though a delegation.
func (gp *GovernanceProposal) VotedWeightRatio() int32 {
	// get the total weight
	total, err := gp.TotalWeight()
	if err != nil {
		gp.repo.Log().Errorf("can not calculate voted ration for %s; %s", gp.GovernanceId.String(), err.Error())
		return 0
	}

	// no weight? simply scratch the test and return zero
	if 0 == total.ToInt().Cmp(zeroInt) {
		return 0
	}

	// get the current proposal state
	state, err := gp.State()
	if err != nil {
		gp.repo.Log().Errorf("can not calculate voted ration for %s; %s", gp.GovernanceId.String(), err.Error())
		return 0
	}

	// no votes?
	if 0 == state.Votes.ToInt().Cmp(zeroInt) {
		return 0
	}

	// calculate the rate
	return int32(new(big.Int).Div(new(big.Int).Mul(big.NewInt(1000), state.Votes.ToInt()), total.ToInt()).Int64())
}

// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/sync/singleflight"
	"math/big"
)

// GovernanceProposal represents details of a Governance Proposal.
type GovernanceProposal struct {
	types.GovernanceProposal
	cg *singleflight.Group
}

// zeroInt represents an empty Big INT value used for comparison.
var zeroInt = new(big.Int)

// NewGovernanceProposal creates a new resolvable GovernanceProposal instance.
func NewGovernanceProposal(gp *types.GovernanceProposal) *GovernanceProposal {
	return &GovernanceProposal{
		GovernanceProposal: *gp,
		cg:                 new(singleflight.Group),
	}
}

// OptionState resolves a state of a given Proposal option identified
// by it's id (index position) in the Proposal options list.
func (gp *GovernanceProposal) OptionState(args *struct{ OptionId hexutil.Big }) (*types.GovernanceOptionState, error) {
	return repository.R().GovernanceOptionState(&gp.GovernanceId, &gp.Id, &args.OptionId)
}

// OptionState resolves a state of a given Proposal option identified
// by it's id (index position) in the Proposal options list.
func (gp *GovernanceProposal) OptionStates() ([]*types.GovernanceOptionState, error) {
	return repository.R().GovernanceOptionStates(&gp.GovernanceId, &gp.Id)
}

// Vote resolves the vote for the given <from> address linked
// with the <delegatedTo> delegation recipient.
func (gp *GovernanceProposal) Vote(args *struct {
	From        common.Address
	DelegatedTo *common.Address
}) (*types.GovernanceVote, error) {
	return repository.R().GovernanceVote(&gp.GovernanceId, &gp.Id, &args.From, args.DelegatedTo)
}

// Governance resolves the parent Governance instance.
func (gp *GovernanceProposal) Governance() (*GovernanceContract, error) {
	// get the governance contract by address
	gc, err := repository.R().GovernanceContractBy(&gp.GovernanceId)
	if err != nil {
		return nil, err
	}

	return &GovernanceContract{
		Name:    gc.Name,
		Type:    gc.Type,
		Address: gp.GovernanceId,
	}, nil
}

// State resolves the state of the Governance Proposal.
func (gp *GovernanceProposal) State() (*types.GovernanceProposalState, error) {
	// make sure to call this only once in parallel processing
	gps, err, _ := gp.cg.Do("state", func() (interface{}, error) {
		return repository.R().GovernanceProposalState(&gp.GovernanceId, &gp.Id)
	})
	if err != nil {
		return nil, err
	}
	return gps.(*types.GovernanceProposalState), nil
}

// TotalWeight resolves the total available voting power which can influence
// the proposal outcome.
func (gp *GovernanceProposal) TotalWeight() (hexutil.Big, error) {
	// make sure to call it only once if in parallel processing
	wt, err, _ := gp.cg.Do("weight", func() (interface{}, error) {
		return repository.R().GovernanceTotalWeight(&gp.GovernanceId)
	})
	if err != nil {
		return hexutil.Big{}, err
	}
	return wt.(hexutil.Big), nil
}

// VotedWeightRatio represents what percentage of the total voting power already
// placed a vote either directly, or though a delegation.
func (gp *GovernanceProposal) VotedWeightRatio() int32 {
	// get the total weight
	total, err := gp.TotalWeight()
	if err != nil || 0 == total.ToInt().Cmp(zeroInt) {
		return 0
	}

	// get the current proposal state
	state, err := gp.State()
	if err != nil || 0 == state.Votes.ToInt().Cmp(zeroInt) {
		return 0
	}

	// calculate the rate
	return int32(new(big.Int).Div(new(big.Int).Mul(big.NewInt(1000), state.Votes.ToInt()), total.ToInt()).Int64())
}

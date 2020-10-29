// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GovernanceProposal represents details of a Governance Proposal.
type GovernanceProposal struct {
	// repo represents the reference to the server repository
	repo repository.Repository

	// the actual proposal details
	types.GovernanceProposal
}

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

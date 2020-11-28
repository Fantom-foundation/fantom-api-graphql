// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GovernanceContract represents details of a Governance contract
// recognized by the API server.
type GovernanceContract struct {
	// repo represents the reference to the server repository
	repo repository.Repository

	// Type represents the internal type of the Governance contract
	Type string

	// Name represents the name of the contract
	Name string

	// Address represents the address of the Governance contract
	Address common.Address
}

// GovContract resolves a governance contract details recognized by the API by address.
func (rs *rootResolver) GovContract(args struct{ Address common.Address }) (*GovernanceContract, error) {
	// get the contract by the address
	gc, err := rs.repo.GovernanceContractBy(&args.Address)
	if err != nil {
		return nil, err
	}

	// return the contract info
	return &GovernanceContract{
		repo:    rs.repo,
		Type:    gc.Type,
		Name:    gc.Name,
		Address: args.Address,
	}, nil
}

// GovContracts resolves list of governance contracts details recognized by the API.
func (rs *rootResolver) GovContracts() ([]*GovernanceContract, error) {
	// do we know any contracts?
	if nil == rs.cfg.Governance.Contracts || 0 == len(rs.cfg.Governance.Contracts) {
		return nil, fmt.Errorf("no governance contracts recognized")
	}

	// make the output array
	res := make([]*GovernanceContract, 0)

	// loop over known governance contracts
	for _, gc := range rs.cfg.Governance.Contracts {
		// add to the structure
		res = append(res, &GovernanceContract{
			repo:    rs.repo,
			Name:    gc.Name,
			Type:    gc.Type,
			Address: gc.Address,
		})
	}

	return res, nil
}

// TotalProposals resolves the number of proposals registered within
// the governance contract.
func (gc *GovernanceContract) TotalProposals() (hexutil.Big, error) {
	return gc.repo.GovernanceProposalsCount(&gc.Address)
}

// Proposal resolves single proposal of the Governance contract specified
// by the proposal id inside the contract.
func (gc *GovernanceContract) Proposal(args *struct{ Id hexutil.Big }) (*GovernanceProposal, error) {
	// get the proposal
	prop, err := gc.repo.GovernanceProposal(&gc.Address, &args.Id)
	if err != nil {
		return nil, err
	}

	// wrap the proposal for resolving
	return &GovernanceProposal{
		repo:               gc.repo,
		GovernanceProposal: *prop,
	}, nil
}

// Proposals resolves list of Governance contract proposals encapsulated in a listable structure.
func (gc *GovernanceContract) Proposals(args *struct {
	Cursor     *Cursor
	Count      int32
	ActiveOnly bool
}) (*GovernanceProposalList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// get the list of all proposals
	list, err := gc.repo.GovernanceProposals([]*common.Address{&gc.Address}, (*string)(args.Cursor), args.Count, args.ActiveOnly)
	if err != nil {
		return nil, err
	}

	return NewGovernanceProposalList(list, gc.repo), nil
}

// DelegationsBy resolves list of delegations an address has in context of the given
// governance contract.
func (gc *GovernanceContract) DelegationsBy(args struct{ From common.Address }) ([]common.Address, error) {
	// decide by the contract type
	switch gc.Type {
	case "sfc":
		return gc.sfcDelegationsBy(args.From)
	}

	// no delegations by default
	return []common.Address{}, nil
}

// CanVote resolves if the given address can post votes in context of the given governance contract.
func (gc *GovernanceContract) CanVote(args struct{ From common.Address }) (bool, error) {
	// decide by the contract type
	switch gc.Type {
	case "sfc":
		return gc.sfcCanVote(args.From)
	}

	// voting disabled by default
	return false, nil
}

// sfcDelegationsBy resolves delegations of the SFC type.
func (gc *GovernanceContract) sfcDelegationsBy(addr common.Address) ([]common.Address, error) {
	// get SFC delegations list
	dl, err := gc.repo.DelegationsByAddress(addr)
	if err != nil {
		return nil, err
	}

	// prep the result container
	res := make([]common.Address, 0)

	// check if the address is a staker
	// if so, it also delegates to itself in the context of the contract
	isStaker, err := gc.repo.IsStaker(&addr)
	if err == nil && isStaker {
		res = append(res, addr)
	}

	// loop delegations to make the list
	for _, d := range dl {
		// get the staker info
		staker, err := gc.repo.StakerAddress(d.ToStakerId)
		if err != nil {
			gc.repo.Log().Errorf("error loading staker %d info; %s", d.ToStakerId, addr.String())
			return nil, err
		}

		// get the
		res = append(res, staker)
	}

	// log delegations found
	gc.repo.Log().Debugf("%d delegations on %s", len(res), addr.String())
	return res, nil
}

// sfcCanVote resolves if a given address can vote in SFC governance context.
func (gc *GovernanceContract) sfcCanVote(addr common.Address) (bool, error) {
	// is the address a staker?
	isStaker, err := gc.repo.IsStaker(&addr)
	if err != nil {
		return false, err
	}

	// if the account is a staker, we are done here
	if isStaker {
		gc.repo.Log().Debugf("%s is validator", addr.String())
		return true, nil
	}

	// check delegating status
	isDelegation, err := gc.repo.IsDelegating(&addr)
	if err != nil {
		return false, err
	}

	return isDelegation, nil
}

// ProposalFee resolves the fee required by the Governance contract to allow
// new proposal to be placed.
func (gc *GovernanceContract) ProposalFee() (hexutil.Big, error) {
	return gc.repo.GovernanceProposalFee(&gc.Address)
}

// totalVotingPower resolves the total available voting power.
func (gc *GovernanceContract) TotalVotingPower() (hexutil.Big, error) {
	return gc.repo.GovernanceTotalWeight(&gc.Address)
}

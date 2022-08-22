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
	gc, err := repository.R().GovernanceContractBy(args.Address)
	if err != nil {
		return nil, err
	}

	// return the contract info
	return &GovernanceContract{
		Type:    gc.Type,
		Name:    gc.Name,
		Address: args.Address,
	}, nil
}

// GovContracts resolves list of governance contracts details recognized by the API.
func (rs *rootResolver) GovContracts() ([]*GovernanceContract, error) {
	// do we know any contracts?
	if nil == cfg.Governance.Contracts || 0 == len(cfg.Governance.Contracts) {
		return nil, fmt.Errorf("no governance contracts recognized")
	}

	// make the output array
	res := make([]*GovernanceContract, len(cfg.Governance.Contracts))
	for i, gc := range cfg.Governance.Contracts {
		// add to the structure
		res[i] = &GovernanceContract{
			Name:    gc.Name,
			Type:    gc.Type,
			Address: gc.Address,
		}
	}
	return res, nil
}

// TotalProposals resolves the number of proposals registered within
// the governance contract.
func (gc *GovernanceContract) TotalProposals() (hexutil.Big, error) {
	return repository.R().GovernanceProposalsCount(&gc.Address)
}

// Proposal resolves single proposal of the Governance contract specified
// by the proposal id inside the contract.
func (gc *GovernanceContract) Proposal(args *struct{ Id hexutil.Big }) (*GovernanceProposal, error) {
	// get the proposal
	prop, err := repository.R().GovernanceProposal(gc.Address, &args.Id)
	if err != nil {
		return nil, err
	}

	// wrap the proposal for resolving
	return NewGovernanceProposal(prop), nil
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
	list, err := repository.R().GovernanceProposals([]common.Address{gc.Address}, (*string)(args.Cursor), args.Count, args.ActiveOnly)
	if err != nil {
		return nil, err
	}

	return NewGovernanceProposalList(list), nil
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
	log.Debugf("unknown governance type of %s", gc.Address.Hex())
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
	log.Debugf("unknown governance type of %s", gc.Address.Hex())
	return false, nil
}

// sfcDelegationsBy resolves delegations of the SFC type.
func (gc *GovernanceContract) sfcDelegationsBy(addr common.Address) ([]common.Address, error) {
	// get SFC delegations list
	dl, err := repository.R().DelegationsByAddressAll(&addr)
	if err != nil {
		return nil, err
	}

	// make the active delegations list
	res := make([]common.Address, 0)
	for _, d := range dl {
		// is the delegation ok for voting?
		if 0 == d.AmountDelegated.ToInt().Uint64() {
			log.Debugf("delegation to %d from address %s is deactivated", d.ToStakerId, addr.String())
			continue
		}
		res = append(res, d.ToStakerAddress)
	}

	// log delegations found
	log.Debugf("%d delegations on %s", len(res), addr.String())
	return res, nil
}

// sfcCanVote resolves if a given address can vote in SFC governance context.
func (gc *GovernanceContract) sfcCanVote(addr common.Address) (bool, error) {
	// even validators are actually delegating to themselves on SFCv3
	return repository.R().IsDelegating(&addr)
}

// ProposalFee resolves the fee required by the Governance contract to allow
// new proposal to be placed.
func (gc *GovernanceContract) ProposalFee() (hexutil.Big, error) {
	return repository.R().GovernanceProposalFee(&gc.Address)
}

// TotalVotingPower resolves the total available voting power.
func (gc *GovernanceContract) TotalVotingPower() (hexutil.Big, error) {
	return repository.R().GovernanceTotalWeight(&gc.Address)
}

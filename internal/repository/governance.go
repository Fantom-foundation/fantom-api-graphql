/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GovernanceProposalsCount provides the total number of proposals
// in a given Governance contract.
func (p *proxy) GovernanceProposalsCount(address *common.Address) (hexutil.Big, error) {
	return p.rpc.GovernanceProposalsCount(address)
}

// GovernanceProposal provides a detail of Proposal of a governance contract
// specified by its id.
func (p *proxy) GovernanceProposal(gov common.Address, id *hexutil.Big) (*types.GovernanceProposal, error) {
	return p.rpc.GovernanceProposal(gov, id)
}

// GovernanceProposalState provides a state of Proposal of a governance contract
// specified by its id.
func (p *proxy) GovernanceProposalState(gov *common.Address, id *hexutil.Big) (*types.GovernanceProposalState, error) {
	return p.rpc.GovernanceProposalState(gov, id)
}

// GovernanceOptionState returns a state of the given option of a proposal.
func (p *proxy) GovernanceOptionState(gov *common.Address, propId *hexutil.Big, optId *hexutil.Big) (*types.GovernanceOptionState, error) {
	return p.rpc.GovernanceOptionState(gov, propId, optId)
}

// GovernanceOptionStates returns a list of states of options of a proposal.
func (p *proxy) GovernanceOptionStates(gov *common.Address, propId *hexutil.Big, optRange int) ([]*types.GovernanceOptionState, error) {
	return p.rpc.GovernanceOptionStates(gov, propId, optRange)
}

// GovernanceVote provides a single vote in the Governance Proposal context.
func (p *proxy) GovernanceVote(
	gov *common.Address,
	propId *hexutil.Big,
	from *common.Address,
	delegatedTo *common.Address) (*types.GovernanceVote, error) {
	return p.rpc.GovernanceVote(gov, propId, from, delegatedTo)
}

// GovernanceContractBy provides governance contract details by its address.
func (p *proxy) GovernanceContractBy(addr common.Address) (config.GovernanceContract, error) {
	// try to pull the config from the map
	if gc, ok := p.govContracts[addr.String()]; ok {
		return gc, nil
	}

	// contract not found
	return config.GovernanceContract{}, fmt.Errorf("governance contract %s not found", addr.String())
}

// GovernanceProposalFee returns the fee payable for a new proposal
// in given Governance contract context.
func (p *proxy) GovernanceProposalFee(gov *common.Address) (hexutil.Big, error) {
	return p.rpc.GovernanceProposalFee(gov)
}

// GovernanceTotalWeight provides the total weight of all available votes
// in the governance contract identified by the address.
func (p *proxy) GovernanceTotalWeight(gov *common.Address) (hexutil.Big, error) {
	// try cache
	we := p.cache.PullGovernanceTotalWeight(gov)
	if we == nil {
		// get the governance config
		cfg, ok := p.govContracts[gov.String()]
		if !ok {
			return hexutil.Big{}, fmt.Errorf("unknown governance %s", gov.String())
		}

		// do it slow way
		var err error
		we, err = p.rpc.GovernanceTotalWeight(&cfg.Governable)
		if err != nil {
			p.log.Errorf("can not pull governance total weight for %s; %s", gov.String(), err.Error())
			return hexutil.Big{}, err
		}

		// store the weight into cache
		if err = p.cache.PushGovernanceTotalWeight(gov, we); err != nil {
			p.log.Errorf("can not cache governance total weight for %s; %s", gov.String(), err.Error())
		}
	}

	return *we, nil
}

// GovernanceProposalOwner returns the owner of the given proposal contract.
func (p *proxy) GovernanceProposalOwner(gp *common.Address) (*common.Address, error) {
	return p.rpc.GovernanceProposalOwner(gp)
}

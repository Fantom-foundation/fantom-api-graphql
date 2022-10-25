/*
Package rpc implements bridge to Opera full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Opera/Opera node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Opera RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Opera RPC interface with connection limited to specified endpoints.

We strongly discourage opening Opera RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"context"
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
)

//go:generate tools/abigen.sh --abi ./contracts/abi/gov_governance.abi --pkg contracts --type Governance --out ./contracts/governance.go
//go:generate tools/abigen.sh --abi ./contracts/abi/gov_iproposal.abi --pkg contracts --type GovernanceProposal --out ./contracts/gov_iproposal.go
//go:generate tools/abigen.sh --abi ./contracts/abi/gov_governable.abi --pkg contracts --type Governable --out ./contracts/gov_governable.go

// proposalExtended represents the extended information
// of a governance proposal.
type govProposalExtended struct {
	Name string
	Desc string
}

// GovernanceProposalsCount provides the total number of proposals
// in a given Governance contract.
func (ftm *FtmBridge) GovernanceProposalsCount(gov *common.Address) (hexutil.Big, error) {
	// get the contract
	gc, err := contracts.NewGovernance(*gov, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not access governance %s; %s", gov.String(), err.Error())
		return hexutil.Big{}, err
	}

	// get the last proposal id
	id, err := gc.LastProposalID(nil)
	if err != nil {
		ftm.log.Errorf("can not count governance %s proposals; %s", gov.String(), err.Error())
		return hexutil.Big{}, err
	}

	return hexutil.Big(*id), nil
}

// GovernanceProposal provides a detail of Proposal of a governance contract
// specified by its id.
func (ftm *FtmBridge) GovernanceProposal(gov common.Address, id *hexutil.Big) (*types.GovernanceProposal, error) {
	// get the contract
	gc, err := contracts.NewGovernance(gov, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not access governance %s; %s", gov.String(), err.Error())
		return nil, err
	}

	// get the detail
	gp, err := ftm.governanceProposalDetail(gc, gov, id.ToInt())
	if err != nil {
		ftm.log.Errorf("proposal %d not available in %s; %s", id.ToInt().Uint64(), gov.String(), err.Error())
		return nil, err
	}

	return gp, nil
}

// GovernanceProposal provides a detail of Proposal of a governance contract
// specified by its id.
func (ftm *FtmBridge) governanceProposalDetail(gc *contracts.Governance, govId common.Address, id *big.Int) (*types.GovernanceProposal, error) {
	// try to get proposal params
	data, err := gc.ProposalParams(nil, id)
	if err != nil {
		return nil, err
	}

	// get some details about the proposal from its contract
	ext, err := ftm.GovernanceProposalDetails(&data.ProposalContract)
	if err != nil {
		return nil, err
	}

	// make and return the proposal detail
	return &types.GovernanceProposal{
		GovernanceId:  govId,
		Id:            hexutil.Big(*id),
		Name:          ext.Name,
		Description:   ext.Desc,
		Contract:      data.ProposalContract,
		ProposalType:  hexutil.Uint64(data.PType.Uint64()),
		IsExecutable:  data.Executable > 0,
		MinVotes:      hexutil.Big(*data.MinVotes),
		MinAgreement:  hexutil.Big(*data.MinAgreement),
		OpinionScales: govConvertScales(data.OpinionScales),
		Options:       govConvertOptions(data.Options),
		VotingStarts:  hexutil.Uint64(data.VotingStartTime.Uint64()),
		VotingMayEnd:  hexutil.Uint64(data.VotingMinEndTime.Uint64()),
		VotingMustEnd: hexutil.Uint64(data.VotingMaxEndTime.Uint64()),
	}, nil
}

// GovernanceProposalState provides a state of Proposal of a governance contract
// specified by its id.
func (ftm *FtmBridge) GovernanceProposalState(gov *common.Address, id *hexutil.Big) (*types.GovernanceProposalState, error) {
	// get the contract
	gc, err := contracts.NewGovernance(*gov, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not access governance %s; %s", gov.String(), err.Error())
		return nil, err
	}

	// get the state
	st, err := gc.ProposalState(nil, id.ToInt())
	if err != nil {
		ftm.log.Errorf("can not access governance %s proposal %d state; %s", gov.String(), id.ToInt().Int64(), err.Error())
		return nil, err
	}

	// does the response make sense?
	if st.Status == nil || st.Votes == nil {
		return &types.GovernanceProposalState{}, nil
	}

	return &types.GovernanceProposalState{
		IsResolved: st.Status.Uint64() == 1,
		WinnerId:   (*hexutil.Big)(st.WinnerOptionID),
		Votes:      hexutil.Big(*st.Votes),
		Status:     hexutil.Big(*st.Status),
	}, nil
}

// govConvertOptions converts the encoded options list into
// an array of strings for processing convenience.
func govConvertOptions(opt [][32]byte) []string {
	// prep result
	res := make([]string, len(opt))

	// loop the options
	for i, v := range opt {
		res[i] = strings.TrimRight(string(v[:]), "\u0000 ")
	}

	return res
}

// govConvertScales converts scales of the governance proposal
// for processing convenience.
func govConvertScales(sc []*big.Int) []hexutil.Uint64 {
	// prep the result
	res := make([]hexutil.Uint64, len(sc))

	// loop incoming scales
	for i, v := range sc {
		res[i] = hexutil.Uint64(v.Uint64())
	}

	return res
}

// GovernanceProposalDetails provides a detail of Proposal of a governance contract
// specified by its id.
func (ftm *FtmBridge) GovernanceProposalDetails(prop *common.Address) (*govProposalExtended, error) {
	// get the proposal contract
	pp, err := contracts.NewGovernanceProposal(*prop, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not access governance proposal %s; %s", prop.String(), err.Error())
		return nil, err
	}

	// prep the container
	ge := govProposalExtended{}

	// load the name
	ge.Name, err = pp.Name(nil)
	if err != nil {
		ftm.log.Errorf("governance proposal %s name not available; %s", prop.String(), err.Error())
		return nil, err
	}

	// load the description
	ge.Desc, err = pp.Description(nil)
	if err != nil {
		ftm.log.Errorf("governance proposal %s description not available; %s", prop.String(), err.Error())
		return nil, err
	}

	return &ge, nil
}

// GovernanceOptionState returns a state of the given option of a proposal.
func (ftm *FtmBridge) GovernanceOptionState(gov *common.Address, propId *hexutil.Big, optId *hexutil.Big) (*types.GovernanceOptionState, error) {
	// get the contract
	gc, err := contracts.NewGovernance(*gov, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not access governance %s; %s", gov.String(), err.Error())
		return nil, err
	}

	// et the state from the contract
	gs, err := ftm.GovernanceOptionStateById(gc, propId, optId)
	if err != nil {
		ftm.log.Errorf("governance %s proposal #%d state #%d not available; %s",
			gov.String(), propId.ToInt().Uint64(), optId.ToInt().Uint64(), err.Error())
		return nil, err
	}

	return gs, nil
}

// GovernanceOptionStates returns a list of states of options of a proposal.
func (ftm *FtmBridge) GovernanceOptionStates(gov *common.Address, propId *hexutil.Big, optRange int) ([]*types.GovernanceOptionState, error) {
	// get the contract
	gc, err := contracts.NewGovernance(*gov, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not access governance %s; %s", gov.String(), err.Error())
		return nil, err
	}

	// make the container and collect the states
	res := make([]*types.GovernanceOptionState, optRange)

	// loop over all possible states and check them one by one
	for i := 0; i < optRange; i++ {
		// get the state of this option
		res[i], err = ftm.GovernanceOptionStateById(gc, propId, (*hexutil.Big)(big.NewInt(int64(i))))
		if err != nil {
			ftm.log.Errorf("unknown option #%d on governance %s; %s", i, gov.String(), err.Error())
			break
		}
	}
	return res, nil
}

// GovernanceOptionStateById returns a state of the given option of a proposal.
func (ftm *FtmBridge) GovernanceOptionStateById(gc *contracts.Governance, propId *hexutil.Big, optId *hexutil.Big) (*types.GovernanceOptionState, error) {
	// get the state
	data, err := gc.ProposalOptionState(nil, propId.ToInt(), optId.ToInt())
	if err != nil {
		return nil, err
	}

	// construct the state and return the values
	return &types.GovernanceOptionState{
		OptionId:       *optId,
		Votes:          hexutil.Big(*data.Votes),
		AgreementRatio: hexutil.Big(*data.AgreementRatio),
		Agreement:      hexutil.Big(*data.Agreement),
	}, nil
}

// GovernanceVote provides a single vote in the Governance Proposal context.
func (ftm *FtmBridge) GovernanceVote(
	gov *common.Address,
	propId *hexutil.Big,
	from *common.Address,
	delegatedTo *common.Address) (*types.GovernanceVote, error) {
	// get the contract
	gc, err := contracts.NewGovernance(*gov, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not access governance %s; %s", gov.String(), err.Error())
		return nil, err
	}

	// if no delegation is in play, use source address as the delegation
	if delegatedTo == nil {
		delegatedTo = from
	}

	// get the vote details
	vote, err := gc.GetVote(nil, *from, *delegatedTo, propId.ToInt())
	if err != nil {
		ftm.log.Errorf("can not access vote of %s on governance %s; %s", from.String(), gov.String(), err.Error())
		return nil, err
	}

	return &types.GovernanceVote{
		GovernanceId: *gov,
		ProposalId:   *propId,
		From:         *from,
		DelegatedTo:  delegatedTo,
		Weight:       hexutil.Big(*vote.Weight),
		Choices:      govConvertScales(vote.Choices),
	}, nil
}

// GovernanceProposalsBy loads list of proposals of the given Governance contract.
func (ftm *FtmBridge) GovernanceProposalsBy(gov common.Address) ([]*types.GovernanceProposal, error) {
	// get the contract
	gc, err := contracts.NewGovernance(gov, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not access governance %s; %s", gov.String(), err.Error())
		return nil, err
	}

	// get the max number of proposals
	maxProposalId, err := gc.LastProposalID(nil)
	if err != nil {
		ftm.log.Errorf("can not count governance %s proposals; %s", gov.String(), err.Error())
		return nil, err
	}

	// log what we do
	ftm.log.Noticef("loading %d proposals of %s", maxProposalId.Uint64(), gov.String())

	// make the array; the maxProposalId starts with 1, so we need array for one less
	result := make([]*types.GovernanceProposal, 0)

	// loop the sys to load proposals
	for i := maxProposalId.Int64(); i > 0; i-- {
		// pull the proposal details
		gp, err := ftm.governanceProposalDetail(gc, gov, big.NewInt(i))
		if err != nil {
			ftm.log.Errorf("can not access governance %s proposal %d; %s", gov.String(), i, err.Error())
			return nil, err
		}

		// keep the proposal in the list
		ftm.log.Debugf("found proposal #%d on %s", gp.Id.ToInt().Uint64(), gov.String())
		result = append(result, gp)
	}

	// log what we do
	ftm.log.Debugf("%d proposals of %s loaded", maxProposalId.Uint64(), gov.String())
	return result, nil
}

// GovernanceProposalFee returns the fee payable for a new proposal
// in given Governance contract context.
func (ftm *FtmBridge) GovernanceProposalFee(gov *common.Address) (hexutil.Big, error) {
	// get the contract
	gc, err := contracts.NewGovernance(*gov, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not access governance %s; %s", gov.String(), err.Error())
		return hexutil.Big{}, err
	}

	// get the fee
	fee, err := gc.ProposalFee(nil)
	if err != nil {
		ftm.log.Errorf("governance %s fee not available; %s", gov.String(), err.Error())
		return hexutil.Big{}, err
	}

	return hexutil.Big(*fee), nil
}

// GovernanceTotalWeight returns the total available voting weight for all proposals
// of a governance contract. The address given must be the Governable contract linked
// to the core Governance.
func (ftm *FtmBridge) GovernanceTotalWeight(ge *common.Address) (*hexutil.Big, error) {
	// get the contract
	goe, err := contracts.NewGovernable(*ge, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not access governable adapter %s; %s", ge.String(), err.Error())
		return nil, err
	}

	// get the total available voting weight
	w, err := goe.GetTotalWeight(ftm.DefaultCallOpts())
	if err != nil {
		ftm.log.Errorf("total weight not available on adapter %s; %s", ge.String(), err.Error())
		return nil, err
	}

	return (*hexutil.Big)(w), nil
}

// GovernanceProposalOwner returns the owner of the given proposal contract.
func (ftm *FtmBridge) GovernanceProposalOwner(gp *common.Address) (*common.Address, error) {
	// make the call
	data, err := ftm.eth.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   gp,
		Data: common.Hex2Bytes("8da5cb5b"), /* owner() public view returns (address) */
	}, nil)
	if err != nil {
		ftm.log.Criticalf("proposal owner not found; %s", err.Error())
		return nil, err
	}

	// make an address off the result
	a := common.Address{}
	a.SetBytes(data)
	return &a, nil
}

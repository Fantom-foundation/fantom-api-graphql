// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GovernanceProposal represents a Governance proposal record.
type GovernanceProposal struct {
	// GovernanceId represents the identifier of the Governance
	// contract this Proposal belongs to.
	GovernanceId common.Address

	// Id represents the identifier of the Proposal in the Governance
	// contract managing the proposal.
	Id hexutil.Big

	// Name represents a name of the Proposal.
	Name string

	// Description represents a textual description of the Proposal.
	Description string

	// Contract represents the contract of the Proposal. Each Proposal
	// is represented by a contract responsible for maintaining the Proposal
	// parameters, options and finalization actions.
	Contract common.Address

	// ProposalType represents the type of the Proposal that corresponds
	// with the Proposal Template.
	ProposalType hexutil.Uint64

	// IsExecutable identifies if the proposal will be finalized
	// by executing a finalizing code.
	IsExecutable bool

	// MinVotes corresponds with the minimal number of votes
	// required by the Proposal to be settled in any way
	// other than REJECTED.
	MinVotes hexutil.Big

	// MinAgreement represents the minimal agreement ratio
	// required to be reached on any of the Proposal options
	// so the Proposal could be settled in any way
	// other than REJECTED.
	MinAgreement hexutil.Big

	// OpinionScales is the scale of opinions on available options.
	// A voter provides a single opinion picked from the scale
	// for each option during the voting for a proposal.
	// I.e.: Scales {0, 2, 3, 4, 5} represent opinions of
	// {strongly disagree, disagree, neutral, agree and strongly agree}.
	OpinionScales []hexutil.Uint64

	// Options is a list of options available on the Proposal.
	// A voter must provide their opinion expressed by a chosen scale
	// for each option on the list. It's generally better to scatter
	// opinions across options instead of having a binary view.
	Options []string

	// VotingStarts is the time stamp of the voting getting opened
	// to receive votes.
	VotingStarts hexutil.Uint64

	// VotingMayEnd is the time stamp when the voting could be closed
	// if enough votes are collected to settle the Proposal (winner option is selectable).
	VotingMayEnd hexutil.Uint64

	// VotingMustEnd is the time stamp when the voting must be closed.
	// If enough votes to settle the Proposal were not collected up until this time
	// the Proposal is rejected and will not be settled in any way (no winner option is selectable).
	VotingMustEnd hexutil.Uint64
}

// GovernanceProposalState represents a state
type GovernanceProposalState struct {
	// IsResolved signals if the Proposal is already resolved.
	IsResolved bool

	// WinnerId is the identifier of the winning option.
	WinnerId *hexutil.Big

	// votes is the number of votes received on the Proposal.
	Votes hexutil.Big

	// status represents the status of the Proposal.
	// 0 = Initial, 1 = Resolved, 2 = Failed, 4 = Canceled, 8 = Execution Expired
	Status hexutil.Big
}

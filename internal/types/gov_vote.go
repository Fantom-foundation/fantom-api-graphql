// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GovernanceVote represents a vote in the Governance Proposal.
type GovernanceVote struct {
	// GovernanceId is the identifier of the Governance contract.
	GovernanceId common.Address

	// ProposalId is the identifier of the Governance Proposal inside the contract.
	ProposalId hexutil.Big

	// From is the address of the voting party
	From common.Address

	// DelegatedTo is the address of the delegation the vote refers to.
	DelegatedTo *common.Address

	// Weight represents the weight of the vote
	Weight hexutil.Big

	// Choices represents the list of opinions on the Proposal options the vote
	// presented.
	Choices []hexutil.Uint64
}

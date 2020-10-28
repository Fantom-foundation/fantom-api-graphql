// Package types implements different core types of the API.
package types

import "github.com/ethereum/go-ethereum/common/hexutil"

// GovernanceOptionState represents information about a state
// of a given Governance Proposal option.
type GovernanceOptionState struct {
	// optionId is the identifier of the option,
	// effectively option index in the options array
	OptionId hexutil.Big

	// Votes is the number of votes received on the option.
	Votes hexutil.Big

	// AgreementRatio represents the ratio of the option agreement across votes.
	AgreementRatio hexutil.Big

	// agreement represents the absolute value of the agreement across votes.
	Agreement hexutil.Big
}

// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"crypto/sha256"
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"html"
	"regexp"
)

const (
	// scValidationMinSourceCodeLength is the minimum length of a validated
	// smart contract source code.
	scMinSourceCodeLength = 32

	// scMaxNameLength is the maximum accepted length of a smart contract name.
	scMaxNameLength = 64

	// scMaxNameLength is the maximum accepted length of a smart contract version.
	scMaxVersionLength = 14

	// scMaxSupportLinkLength is the maximum accepted length of smart contract
	// support link.
	scMaxSupportLinkLength = 64
)

// scVersionSyntaxRegexp represents a regular expression for testing smart contract
// version string syntax. We enforce specific syntax on provided contract versions.
var scVersionSyntaxRegexp = regexp.MustCompile("^\\w?(\\d+\\.)+\\d+$")

// Contract represents resolvable blockchain smart contract structure.
type Contract struct {
	types.Contract
}

// ContractValidationInput represents an input structure used
// to validate contract source code against deployed contract byte code.
type ContractValidationInput struct {
	// Address represents the deployment address of the contract being validated.
	Address common.Address `json:"address"`

	// Name represents an optional name of the contract.
	Name *string `json:"name,omitempty"`

	// Version represents an optional version of the contract.
	// We assume version to be constructed from numbers and dots
	// with optional character at the beginning.
	// I.e. "v1.5.17"
	Version *string `json:"version,omitempty"`

	// SupportContact represents an optional contact information
	// the contract validator wants to publish with the contract
	// details.
	SupportContact *string `json:"supportContact,omitempty"`

	// License represents an optional contact open source license
	// being used.
	License *string `json:"license,omitempty"`

	// IsOptimized signals that the contract byte code was optimized
	// during compilation.
	Optimized bool `json:"optimized"`

	// OptimizeRuns represents number of optimization runs used
	// during the contract compilation.
	OptimizeRuns int32 `json:"optimizeRuns"`

	// SourceCode represents the Solidity source code to be validated.
	SourceCode string `json:"sourceCode"`
}

// NewContract builds new resolvable smart contract structure.
func NewContract(con *types.Contract) *Contract {
	return &Contract{Contract: *con}
}

// DeployedBy resolves the deployment transaction of the contract.
func (con *Contract) DeployedBy() (*Transaction, error) {
	tr, err := repository.R().Transaction(&con.TransactionHash)
	return NewTransaction(tr), err
}

// sanitizeStringOption sanitizes and validates optional string value from the
// smart contract validation check.
func sanitizeStringOption(o *string, length int) (bool, *string) {
	// nil is always ok
	if o == nil {
		return true, nil
	}

	// escape the string
	val := html.EscapeString(*o)
	if len(val) > length {
		return false, nil
	}

	return true, &val
}

// isValidationValid checks the contract validation input and asses
// if it can be processed.
func isValidationValid(in *ContractValidationInput) error {
	// source code must be at least defined number of glyphs long
	if len(in.SourceCode) < scMinSourceCodeLength {
		return fmt.Errorf("contract source code is too short to be valid")
	}

	// collect sanitize result
	var res bool

	// check the name of the contract
	if res, in.Name = sanitizeStringOption(in.Name, scMaxNameLength); !res {
		return fmt.Errorf("contract name is too long to be valid")
	}

	// check the version of the contract
	if res, in.Version = sanitizeStringOption(in.Version, scMaxVersionLength); !res {
		return fmt.Errorf("contract version is too long to be valid")
	}

	// check the contact information of the contract
	if res, in.SupportContact = sanitizeStringOption(in.SupportContact, scMaxSupportLinkLength); !res {
		return fmt.Errorf("contract contact information is too long to be valid")
	}

	// validate the version syntax
	if in.Version != nil && !scVersionSyntaxRegexp.MatchString(*in.Version) {
		return fmt.Errorf("invalid version information provided")
	}

	// validate the version syntax
	if in.OptimizeRuns < 0 {
		return fmt.Errorf("invalid number of optimization runs provided")
	}

	return nil
}

// sourceHash calculates hash of the given source code so we can verify that
// incoming validation source code is not the same one we already know.
func sourceHash(sc string) common.Hash {
	// calculate SHA256 hash of the source code
	sum := sha256.Sum256([]byte(sc))
	return common.BytesToHash(sum[:])
}

// updateContractFromInput update Contract data from provided input structure.
func updateContractFromInput(con *ContractValidationInput, sc *types.Contract) {
	// update the contract detail and pass it to validation
	sc.SourceCode = con.SourceCode
	sc.IsOptimized = con.Optimized
	sc.OptimizeRuns = con.OptimizeRuns

	// pass the intended name
	if con.Name != nil {
		sc.Name = *con.Name
	}

	// pass the intended version
	if con.Version != nil {
		sc.Version = *con.Version
	}

	// pass the intended license
	if con.License != nil {
		sc.License = *con.License
	}

	// pass the intended support contact
	if con.SupportContact != nil {
		sc.SupportContact = *con.SupportContact
	}
}

// ValidateContract resolves smart contract source code vs. deployed byte code and marks
// the contract as validated if the match is found. Peer API points are ringed on success
// to notify them about the change.
func (rs *rootResolver) ValidateContract(args *struct{ Contract ContractValidationInput }) (*Contract, error) {
	// validate the input
	if err := isValidationValid(&args.Contract); err != nil {
		log.Errorf("can not validate contract, validation request is not valid; %s", err.Error())
		return nil, err
	}

	// get a contract to be validated if any
	sc, err := repository.R().Contract(&args.Contract.Address)
	if err != nil {
		log.Errorf("contract [%s] not found", args.Contract.Address.String())
		return nil, err
	}

	// if we already have this source code, no need to do any updates
	hash := sourceHash(args.Contract.SourceCode)
	if sc.SourceCodeHash != nil && hash.String() == sc.SourceCodeHash.String() {
		log.Debugf("contract [%s] source code is already known", sc.Address.String())
		return NewContract(sc), nil
	}

	// copy relevant information from input into the contract struct
	sc.SourceCodeHash = &hash
	updateContractFromInput(&args.Contract, sc)

	// do the validation
	if err := repository.R().ValidateContract(sc); err != nil {
		log.Errorf("contract validation failed; %s", err.Error())
		return nil, err
	}

	// initiate contract syncing in a separated routine
	// we don't really need to wait for it, so let it run
	go rs.syncContract(*sc)

	// return the final updated contract
	return NewContract(sc), nil
}

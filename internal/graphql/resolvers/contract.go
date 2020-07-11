// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"bytes"
	"compress/gzip"
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	repo repository.Repository
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

	// Compiler represents the name of contract compiler to be used for validation.
	Compiler *string `json:"compiler"`

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
func NewContract(con *types.Contract, repo repository.Repository) *Contract {
	return &Contract{
		repo:     repo,
		Contract: *con,
	}
}

// DeployedBy resolves the deployment transaction of the contract.
func (con *Contract) DeployedBy() (*Transaction, error) {
	tr, err := con.repo.Transaction(&con.TransactionHash)
	if err != nil {
		return nil, err
	}

	// make the resolvable transaction object
	trx := NewTransaction(tr, con.repo)
	return trx, nil
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

// updateContractFromInput update Contract data from provided input structure.
func updateContractFromInput(con *ContractValidationInput, sc *types.Contract) {
	// update the contract detail and pass it to validation
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

	// pass the intended support contact
	if con.Compiler != nil {
		sc.Compiler = *con.Compiler
	}
}

// compressContractSourceCode compresses the given contract source code.
// The compressed source code is what we save in the persistent database
// and also what we use for on-chain validation repository.
func compressContractSourceCode(src string) (string, error) {
	// prep the GZIP compressor
	buf := new(bytes.Buffer)
	w, err := gzip.NewWriterLevel(buf, gzip.BestCompression)
	if err != nil {
		return "", err
	}

	// write the source code into the GZIP writer
	if _, err := w.Write([]byte(src)); err != nil {
		return "", err
	}

	// flush the writer so we have all the data in
	// this may not be needed, but we want to make sure
	if err := w.Flush(); err != nil {
		return "", err
	}

	// close the writer
	if err := w.Close(); err != nil {
		return "", err
	}

	// return the compressed data
	return buf.String(), nil
}

// ValidateContract resolves smart contract source code vs. deployed byte code and marks
// the contract as validated if the match is found. Peer API points are ringed on success
// to notify them about the change.
func (rs *rootResolver) ValidateContract(args *struct{ Contract ContractValidationInput }) (*Contract, error) {
	// validate the input
	if err := isValidationValid(&args.Contract); err != nil {
		rs.log.Errorf("can not validate contract, validation request is not valid; %s", err.Error())
		return nil, err
	}

	// get a contract to be validated if any
	sc, err := rs.repo.Contract(&args.Contract.Address)
	if err != nil {
		rs.log.Errorf("contract [%s] not found", args.Contract.Address.String())
		return nil, err
	}

	// compress contract source code
	src, err := compressContractSourceCode(args.Contract.SourceCode)
	if err != nil {
		rs.log.Errorf("source code compression failed; %s", err.Error())
		return nil, err
	}

	// if we already have this source code, no need to do any updates
	hash := crypto.Keccak256Hash([]byte(src))
	if sc.SourceCodeHash != nil && hash.String() == sc.SourceCodeHash.String() {
		rs.log.Debugf("contract [%s] source code is already known", sc.Address.String())
		return NewContract(sc, rs.repo), nil
	}

	// copy relevant information from input into the contract struct
	updateContractFromInput(&args.Contract, sc)
	sc.SourceCodeHash = &hash
	sc.SourceCode = src

	// do the validation
	if err := rs.repo.ValidateContract(sc); err != nil {
		rs.log.Errorf("contract validation failed; %s", err.Error())
		return nil, err
	}

	// return the final updated contract
	return NewContract(sc, rs.repo), nil
}

// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Not functional, placed to "support" existing GraphQL schema only

// Contract represents resolvable blockchain smart contract structure.
type Contract struct {
	Address         common.Address
	TransactionHash common.Hash
	TimeStamp       hexutil.Uint64
	Name            string
	Version         string
	SupportContact  string
	License         string
	Compiler        string
	SourceCode      string
	Abi             string
	Validated       *hexutil.Uint64
}

// NewContract builds new resolvable contract structure.
func NewContract(acc *types.Account) *Contract {
	return &Contract{
		Address:         acc.Address,
		TransactionHash: *acc.DeploymentTrx,
		TimeStamp:       hexutil.Uint64(acc.FirstAppearance.Unix()),
		Name:            acc.Name,
		Version:         "",
		SupportContact:  "",
		License:         "",
		Compiler:        "",
		SourceCode:      "",
		Abi:             "",
		Validated:       nil,
	}
}

// ContractValidationInput represents an input structure used
// to validate contract source code against deployed contract byte code.
type ContractValidationInput struct {
	Address        common.Address `json:"address"`
	Name           *string        `json:"name,omitempty"`
	Version        *string        `json:"version,omitempty"`
	SupportContact *string        `json:"supportContact,omitempty"`
	License        *string        `json:"license,omitempty"`
	Optimized      bool           `json:"optimized"`
	OptimizeRuns   int32          `json:"optimizeRuns"`
	SourceCode     string         `json:"sourceCode"`
}

// DeployedBy resolves the deployment transaction of the contract.
func (con *Contract) DeployedBy() (*Transaction, error) {
	tr, err := repository.R().Transaction(&con.TransactionHash)
	return NewTransaction(tr), err
}

// ValidateContract resolves smart contract source code vs. deployed byte code and marks
// the contract as validated if the match is found. Peer API points are ringed on success
// to notify them about the change.
func (rs *rootResolver) ValidateContract(args *struct{ Contract ContractValidationInput }) (*Contract, error) {
	return nil, fmt.Errorf("method not implemented")
}

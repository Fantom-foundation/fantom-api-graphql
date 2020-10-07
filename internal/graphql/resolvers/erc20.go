// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ERC20Token represents a generic ERC20 token
type ERC20Token struct {
	repo        repository.Repository
	Address     common.Address
	TotalSupply hexutil.Big
}

// NewErc20Token creates a new instance of resolvable ERC20 token, it also validates
// the token existence by loading the total supply of the token
// before making a resolvable instance.
func NewErc20Token(adr *common.Address, repo repository.Repository) *ERC20Token {
	// get the total supply of the token and validate the token existence
	ts, err := repo.Erc20TotalSupply(adr)
	if err != nil {
		return nil
	}

	// make the instance of the token
	return &ERC20Token{
		repo:        repo,
		Address:     *adr,
		TotalSupply: ts,
	}
}

// Erc20Token resolves an instance of ERC20 token if available.
func (rs *rootResolver) Erc20Token(args *struct{ Token common.Address }) *ERC20Token {
	return NewErc20Token(&args.Token, rs.repo)
}

// FMintTokenAllowance resolves the amount of ERC20 tokens unlocked
// by the token owner for DeFi operations.
func (rs *rootResolver) FMintTokenAllowance(args *struct {
	Owner common.Address
	Token common.Address
}) (hexutil.Big, error) {
	return rs.repo.Erc20Allowance(&args.Token, &args.Owner, nil)
}

// ErcTotalSupply resolves the current total supply of the specified token.
func (rs *rootResolver) ErcTotalSupply(args *struct{ Token common.Address }) (hexutil.Big, error) {
	return rs.repo.Erc20TotalSupply(&args.Token)
}

// ErcTokenBalance resolves the current available balance of the specified token
// for the specified owner.
func (rs *rootResolver) ErcTokenBalance(args *struct {
	Owner common.Address
	Token common.Address
}) (hexutil.Big, error) {
	return rs.repo.Erc20BalanceOf(&args.Token, &args.Owner)
}

// ErcTokenAllowance resolves the current amount of ERC20 tokens unlocked
// by the token owner for the spender to be manipulated with.
func (rs *rootResolver) ErcTokenAllowance(args *struct {
	Token   common.Address
	Owner   common.Address
	Spender common.Address
}) (hexutil.Big, error) {
	return rs.repo.Erc20Allowance(&args.Token, &args.Owner, &args.Spender)
}

// Name resolves the name of the given ERC20 token.
func (token *ERC20Token) Name() string {
	// get the token name, if available
	name, err := token.repo.Erc20Name(&token.Address)
	if err != nil {
		return ""
	}

	return name
}

// Symbol resolves the symbol of the given ERC20 token.
func (token *ERC20Token) Symbol() string {
	// get the token symbol, if available
	symbol, err := token.repo.Erc20Symbol(&token.Address)
	if err != nil {
		return ""
	}

	return symbol
}

// Decimals resolves the number of decimals of the given ERC20 token.
func (token *ERC20Token) Decimals() int32 {
	// get the number of decimals, if available
	deci, err := token.repo.Erc20Decimals(&token.Address)
	if err != nil {
		return 0
	}

	return deci
}

// BalanceOf resolves the available balance of the given ERC20 token to a user.
func (token *ERC20Token) BalanceOf(args *struct{ Owner common.Address }) (hexutil.Big, error) {
	return token.repo.Erc20BalanceOf(&token.Address, &args.Owner)
}

// Allowance resolves the unlocked allowance of the given ERC20 token from the owner to spender.
func (token *ERC20Token) Allowance(args *struct {
	Owner   common.Address
	Spender common.Address
}) (hexutil.Big, error) {
	return token.repo.Erc20Allowance(&token.Address, &args.Owner, &args.Spender)
}

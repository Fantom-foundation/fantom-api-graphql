// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ERC20Token represents a generic ERC20 token
type ERC20Token struct {
	types.Erc20Token
}

// NewErc20Token creates a new instance of resolvable ERC20 token, it also validates
// the token existence by loading the total supply of the token
// before making a resolvable instance.
func NewErc20Token(adr *common.Address) *ERC20Token {
	// get the total supply of the token and validate the token existence
	erc20, err := repository.R().Erc20Token(adr)
	if err != nil {
		return nil
	}
	// make the instance of the token
	return &ERC20Token{*erc20}
}

// Erc20Token resolves an instance of ERC20 token if available.
func (rs *rootResolver) Erc20Token(args *struct{ Token common.Address }) *ERC20Token {
	return NewErc20Token(&args.Token)
}

// FMintTokenAllowance resolves the amount of ERC20 tokens unlocked
// by the token owner for DeFi operations.
func (rs *rootResolver) FMintTokenAllowance(args *struct {
	Owner common.Address
	Token common.Address
}) hexutil.Big {
	a, err := repository.R().Erc20Allowance(&args.Token, &args.Owner, nil)
	if err != nil {
		log.Errorf("allowance of %s for %s not known; %s", args.Token.String(), args.Owner.String(), err.Error())
		return hexutil.Big{}
	}
	return a
}

// ErcTotalSupply resolves the current total supply of the specified token.
func (rs *rootResolver) ErcTotalSupply(args *struct{ Token common.Address }) hexutil.Big {
	s, err := repository.R().Erc20TotalSupply(&args.Token)
	if err != nil {
		log.Errorf("total supply of %s not known; %s", args.Token.String(), err.Error())
		return hexutil.Big{}
	}
	return s
}

// ErcTokenBalance resolves the current available balance of the specified token
// for the specified owner.
func (rs *rootResolver) ErcTokenBalance(args *struct {
	Owner common.Address
	Token common.Address
}) hexutil.Big {
	b, err := repository.R().Erc20BalanceOf(&args.Token, &args.Owner)
	if err != nil {
		log.Errorf("balance of %s for %s not known; %s", args.Token.String(), args.Owner.String(), err.Error())
		return hexutil.Big{}
	}
	return b
}

// ErcTokenAllowance resolves the current amount of ERC20 tokens unlocked
// by the token owner for the spender to be manipulated with.
func (rs *rootResolver) ErcTokenAllowance(args *struct {
	Token   common.Address
	Owner   common.Address
	Spender common.Address
}) hexutil.Big {
	a, err := repository.R().Erc20Allowance(&args.Token, &args.Owner, &args.Spender)
	if err != nil {
		log.Errorf("allowance of %s for %s -> %s not known; %s", args.Token.String(), args.Owner.String(), args.Spender.String(), err.Error())
		return hexutil.Big{}
	}
	return a
}

// TotalSupply resolves the total supply of the given ERC20 token.
func (token *ERC20Token) TotalSupply() hexutil.Big {
	s, err := repository.R().Erc20TotalSupply(&token.Address)
	if err != nil {
		log.Errorf("total supply of %s not known; %s", token.Address.String(), err.Error())
		return hexutil.Big{}
	}
	return s
}

// BalanceOf resolves the available balance of the given ERC20 token to a user.
func (token *ERC20Token) BalanceOf(args *struct{ Owner common.Address }) hexutil.Big {
	b, err := repository.R().Erc20BalanceOf(&token.Address, &args.Owner)
	if err != nil {
		log.Errorf("balance of %s for %s not known; %s", token.Address.String(), args.Owner.String(), err.Error())
		return hexutil.Big{}
	}
	return b
}

// Allowance resolves the unlocked allowance of the given ERC20 token from the owner to spender.
func (token *ERC20Token) Allowance(args *struct {
	Owner   common.Address
	Spender common.Address
}) hexutil.Big {
	a, err := repository.R().Erc20Allowance(&token.Address, &args.Owner, &args.Spender)
	if err != nil {
		log.Errorf("allowance of %s for %s -> %s not known; %s", token.Address.String(), args.Owner.String(), args.Spender.String(), err.Error())
		return hexutil.Big{}
	}
	return a
}

// LogoURL resolves an URL of the token logo.
func (token *ERC20Token) LogoURL() string {
	return repository.R().Erc20LogoURL(&token.Address)
}

// TotalDeposit represents the total amount of tokens deposited to fMint as collateral.
func (token *ERC20Token) TotalDeposit() hexutil.Big {
	d, err := repository.R().FMintTokenTotalBalance(&token.Address, types.DefiTokenTypeCollateral)
	if err != nil {
		log.Errorf("unknown deposit of %s; %s", token.Address.String(), err.Error())
		return hexutil.Big{}
	}
	return d
}

// TotalDebt represents the total amount of tokens borrowed/minted on fMint.
func (token *ERC20Token) TotalDebt() hexutil.Big {
	d, err := repository.R().FMintTokenTotalBalance(&token.Address, types.DefiTokenTypeDebt)
	if err != nil {
		log.Errorf("unknown debt of %s; %s", token.Address.String(), err.Error())
		return hexutil.Big{}
	}
	return d
}

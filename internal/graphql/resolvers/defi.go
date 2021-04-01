// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// defiWFTMSymbol is the symbol used for wrapped FTM tokens.
const defiWFTMSymbol = "WFTM"

// DefiToken represents a resolvable DeFi token instance.
type DefiToken struct {
	types.DefiToken
}

// NewDefiToken creates a new instance of resolvable DeFi token.
func NewDefiToken(tk *types.DefiToken) *DefiToken {
	return &DefiToken{DefiToken: *tk}
}

// DefiTokens resolves list of DeFi tokens available for the DeFi functions.
func (rs *rootResolver) DefiTokens() ([]*DefiToken, error) {
	// pass the call to repository
	tkList, err := repository.R().DefiTokens()
	if err != nil {
		return nil, err
	}

	// prep the container
	list := make([]*DefiToken, len(tkList))
	for ix, tk := range tkList {
		list[ix] = NewDefiToken(&tk)
	}
	return list, nil
}

// DefiNativeToken resolves the native FTM wrapper token.
func (rs *rootResolver) DefiNativeToken() *ERC20Token {
	// get the token address
	adr, err := repository.R().NativeTokenAddress()
	if err != nil {
		return nil
	}
	return NewErc20Token(adr)
}

// Price resolves the value of the token in ref. denomination
// using on-chain price oracle.
func (dt *DefiToken) Price() (hexutil.Big, error) {
	return repository.R().DefiTokenPrice(&dt.Address)
}

// AvailableBalance resolves the total amount of ERC20 tokens
// available to the specified token holder.
func (dt *DefiToken) AvailableBalance(args *struct{ Owner common.Address }) (hexutil.Big, error) {
	return repository.R().Erc20BalanceOf(&dt.Address, &args.Owner)
}

// Allowance resolves the total amount of ERC20 tokens unlocked
// by the token holder for DeFi operations.
func (dt *DefiToken) Allowance(args *struct{ Owner common.Address }) (hexutil.Big, error) {
	return repository.R().Erc20Allowance(&dt.Address, &args.Owner, nil)
}

// CanWrapFTM signals if the token can be used to wrap native FTM
// to get some amount of it.
func (dt *DefiToken) CanWrapFTM() bool {
	return dt.Symbol == defiWFTMSymbol
}

// TotalSupply represents the total amount of tokens on supply.
func (dt *DefiToken) TotalSupply() (hexutil.Big, error) {
	return repository.R().Erc20TotalSupply(&dt.Address)
}

// TotalDeposit represents the total amount of tokens deposited to fMint as collateral.
func (dt *DefiToken) TotalDeposit() (hexutil.Big, error) {
	return repository.R().FMintTokenTotalBalance(&dt.Address, types.DefiTokenTypeCollateral)
}

// TotalDebt represents the total amount of tokens borrowed/minted on fMint.
func (dt *DefiToken) TotalDebt() (hexutil.Big, error) {
	return repository.R().FMintTokenTotalBalance(&dt.Address, types.DefiTokenTypeDebt)
}

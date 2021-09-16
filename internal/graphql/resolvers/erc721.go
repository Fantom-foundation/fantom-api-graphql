// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// ERC721Token represents a generic ERC721 token
type ERC721Token struct {
	types.Erc721Token
}

// NewErc721Token creates a new instance of resolvable ERC721 token.
func NewErc721Token(adr *common.Address) *ERC721Token {
	// get the total supply of the token and validate the token existence
	token, err := repository.R().Erc721Token(adr)
	if err != nil {
		return nil
	}
	// make the instance of the token
	return &ERC721Token{*token}
}

// Erc721Token resolves an instance of ERC721 token if available.
func (rs *rootResolver) Erc721Token(args *struct{ Token common.Address }) *ERC721Token {
	return NewErc721Token(&args.Token)
}

// TotalSupply resolves the total supply of the given ERC20 token.
func (token *ERC721Token) TotalSupply() (*hexutil.Big, error) {
	totalSupply, err := repository.R().Erc721TotalSupply(&token.Address)
	if err != nil { // ignore err, return null
		return nil, nil
	} else {
		return &totalSupply, err
	}
}

// BalanceOf resolves the available balance of the given ERC721 token to a user.
func (token *ERC721Token) BalanceOf(args *struct{ Owner common.Address }) (hexutil.Big, error) {
	return repository.R().Erc721BalanceOf(&token.Address, &args.Owner)
}

// TokenURI provides URI of Metadata JSON Schema of the ERC721 token.
func (token *ERC721Token) TokenURI(args *struct{ TokenId hexutil.Big }) (*string, error) {
	tokenId := big.Int(args.TokenId)
	uri, err := repository.R().Erc721TokenURI(&token.Address, &tokenId)
	if err != nil { // ignore err, return null
		return nil, nil
	} else {
		return &uri, err
	}
}

// OwnerOf provides information about NFT token ownership.
func (token *ERC721Token) OwnerOf(args *struct{ TokenId hexutil.Big }) (*common.Address, error) {
	tokenId := big.Int(args.TokenId)
	owner, err := repository.R().Erc721OwnerOf(&token.Address, &tokenId)
	if err != nil { // ignore err, return null
		return nil, nil
	} else {
		return &owner, err
	}
}

// GetApproved provides information about operator approved to manipulate with the NFT token.
func (token *ERC721Token) GetApproved(args *struct{ TokenId hexutil.Big }) (*common.Address, error) {
	tokenId := big.Int(args.TokenId)
	operator, err := repository.R().Erc721GetApproved(&token.Address, &tokenId)
	if err != nil { // ignore err, return null
		return nil, nil
	} else {
		return &operator, err
	}
}

// IsApprovedForAll provides information about operator approved to manipulate with NFT tokens of given owner.
func (token *ERC721Token) IsApprovedForAll(args *struct{ Owner common.Address; Operator common.Address }) (*bool, error) {
	isApproved, err := repository.R().Erc721IsApprovedForAll(&token.Address, &args.Owner, &args.Operator)
	if err != nil { // ignore err, return null
		return nil, nil
	} else {
		return &isApproved, err
	}
}

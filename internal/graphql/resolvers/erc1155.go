package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// ERC1155Contract represents a generic ERC721 token
type ERC1155Contract struct {
	types.Erc1155Contract
}

// NewErc1155Contract creates a new instance of resolvable ERC1155 contract.
func NewErc1155Contract(adr *common.Address) *ERC1155Contract {
	return &ERC1155Contract{ types.Erc1155Contract{Address: *adr} }
}

// Erc1155Contract resolves an instance of ERC1155 contract if available.
func (rs *rootResolver) Erc1155Contract(args *struct{ Address common.Address }) *ERC1155Contract {
	return NewErc1155Contract(&args.Address)
}

// Uri provides URI of Metadata JSON Schema of the token.
func (token *ERC1155Contract) Uri(args *struct{ TokenId hexutil.Big }) (*string, error) {
	tokenId := big.Int(args.TokenId)
	uri, err := repository.R().Erc1155Uri(&token.Address, &tokenId)
	if err != nil { // optional - ignore err, return null
		return nil, nil
	} else {
		return &uri, err
	}
}

// BalanceOf resolves the available balance of the given token for a user.
func (token *ERC1155Contract) BalanceOf(args *struct{ Owner common.Address; TokenId hexutil.Big }) (hexutil.Big, error) {
	tokenId := big.Int(args.TokenId)
	balance,err := repository.R().Erc1155BalanceOf(&token.Address, &args.Owner, &tokenId)
	if err != nil || balance == nil {
		return hexutil.Big{}, err
	} else {
		return hexutil.Big(*balance), err
	}
}

// BalanceOfBatch resolves the available balances of the given tokens and owners.
func (token *ERC1155Contract) BalanceOfBatch(args *struct{ Owners []common.Address; TokenIds []hexutil.Big }) ([]hexutil.Big, error) {
	tokenIds := make([]*big.Int, len(args.TokenIds))
	for i, tokenId := range args.TokenIds {
		value := big.Int(tokenId)
		tokenIds[i] = &value
	}

	balances,err := repository.R().Erc1155BalanceOfBatch(&token.Address, &args.Owners, tokenIds)
	if err != nil || balances == nil {
		return nil, err
	} else {
		hexBalances := make([]hexutil.Big, len(balances))
		for i, balance := range balances {
			hexBalances[i] = hexutil.Big(*balance)
		}
		return hexBalances, nil
	}
}

// IsApprovedForAll provides information about operator approved to manipulate with tokens of given owner.
func (token *ERC1155Contract) IsApprovedForAll(args *struct{ Owner common.Address; Operator common.Address }) (*bool, error) {
	isApproved, err := repository.R().Erc1155IsApprovedForAll(&token.Address, &args.Owner, &args.Operator)
	if err != nil { // ignore err, return null
		return nil, nil
	} else {
		return &isApproved, err
	}
}

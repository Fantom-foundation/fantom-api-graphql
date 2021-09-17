package repository

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Erc1155Contract returns an ERC1155 contract for the given address, if available.
func (p *proxy) Erc1155Contract(addr *common.Address) (*types.Erc1155Contract, error) {
	return &types.Erc1155Contract{Address: *addr}, nil
}

// Erc1155Uri provides URI of Metadata JSON Schema of the token.
func (p *proxy) Erc1155Uri(token *common.Address, tokenId *big.Int) (string, error) {
	return p.rpc.Erc1155Uri(token, tokenId)
}

// Erc1155BalanceOf provides amount of NFT tokens owned by given owner.
func (p *proxy) Erc1155BalanceOf(token *common.Address, owner *common.Address, tokenId *big.Int) (*big.Int, error) {
	return p.rpc.Erc1155BalanceOf(token, owner, tokenId)
}

// Erc1155BalanceOfBatch provides amount of NFT tokens owned by given owner.
func (p *proxy) Erc1155BalanceOfBatch(token *common.Address, owners *[]common.Address, tokenIds []*big.Int) ([]*big.Int, error) {
	return p.rpc.Erc1155BalanceOfBatch(token, owners, tokenIds)
}

// Erc1155IsApprovedForAll provides information about operator approved to manipulate with NFT tokens of given owner.
func (p *proxy) Erc1155IsApprovedForAll(token *common.Address, owner *common.Address, operator *common.Address) (bool, error) {
	return p.rpc.Erc1155IsApprovedForAll(token, owner, operator)
}

// Erc1155ContractsList returns a list of known ERC1155 tokens ordered by their activity.
func (p *proxy) Erc1155ContractsList(count int32) ([]common.Address, error) {
	return p.db.Erc1155ContractsList(count)
}

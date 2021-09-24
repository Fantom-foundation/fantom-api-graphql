package repository

import (
	"fantom-api-graphql/internal/repository/cache"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// Erc721Contract returns an ERC721 token for the given address, if available.
func (p *proxy) Erc721Contract(addr *common.Address) (*types.Erc721Contract, error) {
	// get the token
	val, err, _ := p.apiRequestGroup.Do(cache.ErcTokenId(addr, cache.Erc721CacheIdPrefix), func() (interface{}, error) {
		// try cache first
		var err error
		tok := p.cache.PullErc721Contract(addr)
		if tok != nil {
			p.log.Debugf("found cached token %s", tok.Address.String())
			return tok, nil
		}

		tok, err = p.loadErc721ContractDetails(&types.Erc721Contract{Address: *addr})
		if err != nil {
			p.log.Errorf("can not load ERC721 token at %s; %s", addr.String(), err.Error())
			return nil, err
		}

		p.log.Debugf("found ERC-721 %s at %s", tok.Symbol, tok.Address.String())
		err = p.cache.PushErc721Contract(tok)
		if err != nil {
			p.log.Errorf("can not keep ERC-721 token %s in cache; %s", addr.String(), err.Error())
		}
		return tok, nil
	})

	if err != nil {
		return nil, err
	}
	return val.(*types.Erc721Contract), nil
}

func (p *proxy) loadErc721ContractDetails(token *types.Erc721Contract) (*types.Erc721Contract, error) {
	var err error

	// get the name (ignore fail - name is optional in ERC721)
	token.Name, err = p.rpc.Erc721Name(&token.Address)
	if err != nil {
		p.log.Noticef("ERC721 name failed for %s; %s", token.Address.String(), err.Error())
		return nil, err
	}

	// get symbol (ignore fail - symbol is optional in ERC721)
	token.Symbol, err = p.rpc.Erc721Symbol(&token.Address)
	if err != nil {
		p.log.Noticef("ERC721 symbol failed for %s; %s", token.Address.String(), err.Error())
		return nil, err
	}

	return token, nil
}

// Erc721Name provides information about the name of the ERC721 token.
func (p *proxy) Erc721Name(adr *common.Address) (string, error) {
	tk, err := p.Erc721Contract(adr)
	if err != nil {
		return "", err
	}
	return tk.Name, nil
}

// Erc721Symbol provides information about the symbol of the ERC721 token.
func (p *proxy) Erc721Symbol(adr *common.Address) (string, error) {
	tk, err := p.Erc721Contract(adr)
	if err != nil {
		return "", err
	}
	return tk.Symbol, nil
}

// Erc721BalanceOf provides amount of NFT tokens owned by given owner in given ERC721 contract.
func (p *proxy) Erc721BalanceOf(token *common.Address, owner *common.Address) (hexutil.Big, error) {
	return p.rpc.Erc721BalanceOf(token, owner)
}

// Erc721TotalSupply provides information about all available tokens.
func (p *proxy) Erc721TotalSupply(token *common.Address) (hexutil.Big, error) {
	return p.rpc.Erc721TotalSupply(token)
}

// Erc721TokenURI provides URI of Metadata JSON Schema of the ERC721 token.
func (p *proxy) Erc721TokenURI(token *common.Address, tokenId *big.Int) (string, error) {
	return p.rpc.Erc721TokenURI(token, tokenId)
}

// Erc721OwnerOf provides information about NFT token ownership.
func (p *proxy) Erc721OwnerOf(token *common.Address, tokenId *big.Int) (common.Address, error) {
	return p.rpc.Erc721OwnerOf(token, tokenId)
}

// Erc721GetApproved provides information about operator approved to manipulate with the NFT token.
func (p *proxy) Erc721GetApproved(token *common.Address, tokenId *big.Int) (common.Address, error) {
	return p.rpc.Erc721GetApproved(token, tokenId)
}

// Erc721IsApprovedForAll provides information about operator approved to manipulate with NFT tokens of given owner.
func (p *proxy) Erc721IsApprovedForAll(token *common.Address, owner *common.Address, operator *common.Address) (bool, error) {
	return p.rpc.Erc721IsApprovedForAll(token, owner, operator)
}

// Erc721ContractsList returns a list of known ERC721 tokens ordered by their activity.
func (p *proxy) Erc721ContractsList(count int32) ([]common.Address, error) {
	return p.db.Erc721ContractsList(count)
}

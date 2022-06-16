/*
Package rpc implements bridge to Opera full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Opera/Opera node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Opera RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Opera RPC interface with connection limited to specified endpoints.

We strongly discourage opening Opera RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

//go:generate tools/abigen.sh --abi ./contracts/abi/erc721.abi --pkg contracts --type ERC721 --out ./contracts/erc721_token.go

// Erc721Name provides information about the name of the ERC721 token.
func (ftm *FtmBridge) Erc721Name(token *common.Address) (string, error) {
	// connect the contract
	contract, err := contracts.NewERC721(*token, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact ERC721 contract; %s", err.Error())
		return "", err
	}

	// get the token name
	name, err := contract.Name(nil)
	if err != nil {
		ftm.log.Errorf("ERC721 token %s name not available; %s", token.String(), err.Error())
		return "", err
	}

	return name, nil
}

// Erc721Symbol provides information about the symbol of the ERC721 token.
func (ftm *FtmBridge) Erc721Symbol(token *common.Address) (string, error) {
	// connect the contract
	contract, err := contracts.NewERC721(*token, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact ERC721 contract; %s", err.Error())
		return "", err
	}

	// get the token name
	symbol, err := contract.Symbol(nil)
	if err != nil {
		ftm.log.Errorf("ERC721 token %s symbol not available; %s", token.String(), err.Error())
		return "", err
	}

	return symbol, nil
}

// Erc721BalanceOf provides amount of NFT tokens owned by given owner in given ERC721 contract.
func (ftm *FtmBridge) Erc721BalanceOf(token *common.Address, owner *common.Address) (hexutil.Big, error) {
	// connect the contract
	contract, err := contracts.NewERC721(*token, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact ERC721 contract; %s", err.Error())
		return hexutil.Big{}, err
	}

	// get the balance
	val, err := contract.BalanceOf(nil, *owner)
	if err != nil {
		ftm.log.Errorf("can not ERC721 %s balance for %s; %s", token.String(), owner.String(), err.Error())
		return hexutil.Big{}, err
	}

	// make sur we always have a value; at least zero
	// this should always be the case since the contract should
	// return zero even for unknown owners, but let's be sure here
	if val == nil {
		ftm.log.Errorf("no balance available for ERC721 %s, owner %s", token.String(), owner.String())
		val = new(big.Int)
	}

	// return the account balance
	return hexutil.Big(*val), nil
}

// Erc721TotalSupply provides information about all available tokens
func (ftm *FtmBridge) Erc721TotalSupply(token *common.Address) (hexutil.Big, error) {
	// connect the contract
	contract, err := contracts.NewERC721(*token, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact ERC721 contract; %s", err.Error())
		return hexutil.Big{}, err
	}

	// get the amount of tokens allowed for DeFi
	val, err := contract.TotalSupply(nil)
	if err != nil {
		ftm.log.Errorf("can not get ERC721 %s total supply; %s", token.String(), err.Error())
		return hexutil.Big{}, err
	}

	// make sure we always have a value; at least zero
	if val == nil {
		ftm.log.Errorf("no supply available for ERC721 %s", token.String())
		val = new(big.Int)
	}

	// return the account balance
	return hexutil.Big(*val), nil
}

// Erc721TokenURI provides URI of Metadata JSON Schema of the ERC721 token.
func (ftm *FtmBridge) Erc721TokenURI(token *common.Address, tokenId *big.Int) (string, error) {
	// connect the contract
	contract, err := contracts.NewERC721(*token, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact ERC721 contract; %s", err.Error())
		return "", err
	}

	// get the token name
	uri, err := contract.TokenURI(nil, tokenId)
	if err != nil {
		ftm.log.Errorf("ERC721 token %s/%s URI not available; %s", token.String(), tokenId.String(), err.Error())
		return "", err
	}

	return uri, nil
}

// Erc721OwnerOf provides information about NFT token ownership
func (ftm *FtmBridge) Erc721OwnerOf(token *common.Address, tokenId *big.Int) (common.Address, error) {
	// connect the contract
	contract, err := contracts.NewERC721(*token, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact ERC721 contract; %s", err.Error())
		return common.Address{}, err
	}

	owner, err := contract.OwnerOf(nil, tokenId)
	if err != nil {
		ftm.log.Errorf("can not get ERC721 %s owner of %s; %s", token.String(), tokenId.String(), err.Error())
		return common.Address{}, err
	}

	return owner, nil
}

// Erc721GetApproved provides information about operator approved to manipulate with the NFT token.
func (ftm *FtmBridge) Erc721GetApproved(token *common.Address, tokenId *big.Int) (common.Address, error) {
	// connect the contract
	contract, err := contracts.NewERC721(*token, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact ERC721 contract; %s", err.Error())
		return common.Address{}, err
	}

	owner, err := contract.GetApproved(nil, tokenId)
	if err != nil {
		ftm.log.Errorf("can not get ERC721 %s approved operator of %s; %s", token.String(), tokenId.String(), err.Error())
		return common.Address{}, err
	}

	return owner, nil
}

// Erc721IsApprovedForAll provides information about operator approved to manipulate with NFT tokens of given owner.
func (ftm *FtmBridge) Erc721IsApprovedForAll(token *common.Address, owner *common.Address, operator *common.Address) (bool, error) {
	// connect the contract
	contract, err := contracts.NewERC721(*token, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact ERC721 contract; %s", err.Error())
		return false, err
	}

	isApproved, err := contract.IsApprovedForAll(nil, *owner, *operator)
	if err != nil {
		ftm.log.Errorf("can not get ERC721 %s approved-for-all status for owner %s and operator %s; %s", token.String(), owner.String(), operator.String(), err.Error())
		return false, err
	}

	return isApproved, nil
}

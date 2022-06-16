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
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

//go:generate tools/abigen.sh --abi ./contracts/abi/erc1155.abi --pkg contracts --type ERC1155 --out ./contracts/erc1155_token.go

// Erc1155Uri provides URI of Metadata JSON Schema of the ERC1155 token.
func (ftm *FtmBridge) Erc1155Uri(token *common.Address, tokenId *big.Int) (string, error) {
	// connect the contract
	contract, err := contracts.NewERC1155(*token, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact ERC1155 contract; %s", err.Error())
		return "", err
	}

	// get the token name
	uri, err := contract.Uri(nil, tokenId)
	if err != nil {
		ftm.log.Errorf("ERC1155 token %s/%s URI not available; %s", token.String(), tokenId.String(), err.Error())
		return "", err
	}

	return uri, nil
}

// Erc1155BalanceOf provides amount of tokens owned by given owner in given ERC1155 contract.
func (ftm *FtmBridge) Erc1155BalanceOf(token *common.Address, owner *common.Address, tokenId *big.Int) (*big.Int, error) {
	// connect the contract
	contract, err := contracts.NewERC1155(*token, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact ERC1155 contract; %s", err.Error())
		return nil, err
	}

	balance, err := contract.BalanceOf(nil, *owner, tokenId)
	if err != nil {
		ftm.log.Errorf("can not get ERC1155 %s/%s balance for %s; %s", token.String(), tokenId.String(), owner.String(), err.Error())
		return nil, err
	}

	return balance, nil
}

// Erc1155BalanceOfBatch provides amounts of tokens owned by given owners in given ERC1155 contract.
func (ftm *FtmBridge) Erc1155BalanceOfBatch(token *common.Address, owners *[]common.Address, tokenIds []*big.Int) ([]*big.Int, error) {
	// connect the contract
	contract, err := contracts.NewERC1155(*token, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact ERC1155 contract; %s", err.Error())
		return nil, err
	}

	balances, err := contract.BalanceOfBatch(nil, *owners, tokenIds)
	if err != nil {
		ftm.log.Errorf("can not get ERC1155 batch balance for %s; %s", token.String(), err.Error())
		return nil, err
	}

	return balances, nil
}

// Erc1155IsApprovedForAll provides information about operator approved to manipulate with tokens of given owner.
func (ftm *FtmBridge) Erc1155IsApprovedForAll(token *common.Address, owner *common.Address, operator *common.Address) (bool, error) {
	// connect the contract
	contract, err := contracts.NewERC1155(*token, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact ERC1155 contract; %s", err.Error())
		return false, err
	}

	isApproved, err := contract.IsApprovedForAll(nil, *owner, *operator)
	if err != nil {
		ftm.log.Errorf("can not get ERC1155 %s approved-for-all status for owner %s and operator %s; %s", token.String(), owner.String(), operator.String(), err.Error())
		return false, err
	}

	return isApproved, nil
}

var erc1155contractAbi *abi.ABI // parsed ABI singleton

func Erc1155ParseTransferBatchData(data []byte) (ids []*big.Int, values []*big.Int, err error) {
	if erc1155contractAbi == nil {
		contractAbi, err := abi.JSON(strings.NewReader(contracts.ERC1155MetaData.ABI))
		if err != nil {
			return nil, nil, err
		}
		erc1155contractAbi = &contractAbi
	}

	outs, err := erc1155contractAbi.Unpack("TransferBatch", data)
	if err != nil {
		return nil, nil, err
	}
	ids = (outs[0]).([]*big.Int)
	values = (outs[1]).([]*big.Int)
	return ids, values, err
}

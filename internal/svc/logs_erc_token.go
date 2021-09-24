// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/repository/rpc"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// handleErcTokenApproval handles Approval event on ERC20 or ERC721 token.
// event Approval(address indexed owner, address indexed spender, uint256 value)
func handleErcTokenApproval(lr *types.LogRecord) {
	handleErcTransaction(lr, types.TokenTrxTypeApproval)
}

// handleErcTokenTransfer handles Transfer event on ERC20 or ERC721 token.
// event Transfer(address indexed from, address indexed to, uint256 value)
func handleErcTokenTransfer(lr *types.LogRecord) {
	handleErcTransaction(lr, types.TokenTrxTypeTransfer)
}

// handleErcTransaction handles Approval and/or Transfer event on an ERC20/ERC721 token.
// event Transfer(address indexed from,  address indexed to, uint256 value); // ERC20
// event Transfer(address indexed from,  address indexed to, uint256 indexed _tokenId); // ERC721
// event Approval(address indexed owner, address indexed operator, uint256 value); // ERC20
// event Approval(address indexed owner, address indexed operator, uint256 indexed _tokenId); // ERC721
func handleErcTransaction(lr *types.LogRecord, trxType int32) {

	// ERC20 has 2 indexed params (=> 3 topics) and 1 non-indexed uint256 param (=> 32 bytes)
	if len(lr.Topics) == 3 && len(lr.Data) == 32 {
		from := common.BytesToAddress(lr.Topics[1].Bytes())
		to := common.BytesToAddress(lr.Topics[2].Bytes())
		amount := new(big.Int).SetBytes(lr.Data[:])
		tokenId := big.NewInt(0)
		storeTokenTransaction(lr, types.AccountTypeERC20Token, tokenTrxType(trxType, from, to), from, to, *amount, *tokenId, 0)
		return
	}

	// ERC721 has 3 indexed params (=> 4 topics) and no non-indexed param (=> 0 bytes)
	if len(lr.Topics) == 4 && len(lr.Data) == 0 {
		from := common.BytesToAddress(lr.Topics[1].Bytes())
		to := common.BytesToAddress(lr.Topics[2].Bytes())
		amount := big.NewInt(1)
		tokenId := new(big.Int).SetBytes(lr.Topics[3].Bytes())
		storeTokenTransaction(lr, types.AccountTypeERC721Contract, tokenTrxType(trxType, from, to), from, to, *amount, *tokenId, 0)
		return
	}

	log.Debugf("Unrecognized ERC-20/ERC-721 Transfer/Approval from tx %s (%d data bytes, %d topics)", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
}

// event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 tokenId, uint256 value)
func handleErc1155TransferSingle(lr *types.LogRecord) {
	// 3 indexed params, 2 uint256 params
	if len(lr.Topics) == 4 && len(lr.Data) == 64 {
		from := common.BytesToAddress(lr.Topics[2].Bytes())
		to := common.BytesToAddress(lr.Topics[3].Bytes())
		tokenId := new(big.Int).SetBytes(lr.Data[0:32])
		amount := new(big.Int).SetBytes(lr.Data[32:64])
		storeTokenTransaction(lr, types.AccountTypeERC1155Contract, tokenTrxType(types.TokenTrxTypeTransfer, from, to), from, to, *amount, *tokenId, 0)
		return
	}
	log.Debugf("Unrecognized ERC1155 TransferSingle from tx %s (%d data bytes, %d topics)", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
}

// event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func handleErc1155TransferBatch(lr *types.LogRecord) {
	// 3 indexed params
	if len(lr.Topics) == 4 {
		from := common.BytesToAddress(lr.Topics[2].Bytes())
		to := common.BytesToAddress(lr.Topics[3].Bytes())
		ids, values, err := rpc.Erc1155ParseTransferBatchData(lr.Data)
		if err != nil {
			log.Errorf("failed to parse ERC1155 TransferBatch data - trx %s; %s", lr.TxHash.String(), err.Error())
		}
		if len(ids) != len(values) {
			log.Errorf("ERC1155 TransferBatch ids and values length differs - trx %s", lr.TxHash.String())
		}
		for i := range ids {
			log.Infof("ERC1155 storing TransferBatch - trx %s - len %d", lr.TxHash.String(), len(ids))
			storeTokenTransaction(lr, types.AccountTypeERC1155Contract, types.TokenTrxTypeTransfer, from, to, *values[i], *ids[i], uint16(i))
		}
		return
	}
	log.Debugf("Unrecognized ERC-1155 TransferBatch from tx %s (%d data bytes, %d topics)", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
}

func tokenTrxType(trxType int32, from common.Address, to common.Address) int32 {
	if trxType == types.TokenTrxTypeTransfer && config.EmptyAddress == from.String() {
		return types.TokenTrxTypeMint
	}
	if trxType == types.TokenTrxTypeTransfer && config.EmptyAddress == to.String() {
		return types.TokenTrxTypeBurn
	}
	return trxType
}

// storeTokenTransaction handles general token (ERC20/ERC721/ERC1155) transaction.
func storeTokenTransaction(lr *types.LogRecord, tokenType string, eventType int32, from common.Address, to common.Address, amount big.Int, tokenId big.Int, seq uint16) {
	if err := repo.StoreTokenTransaction(&types.TokenTransaction{
		Transaction:  lr.TxHash,
		TrxIndex:     hexutil.Uint64(uint64(lr.TxIndex)),
		TokenAddress: lr.Address,
		Type:         eventType,
		TokenType:    tokenType,
		Sender:       from,
		Recipient:    to,
		Amount:       hexutil.Big(amount),
		TokenId:      hexutil.Big(tokenId),
		TimeStamp:    lr.Block.TimeStamp,
		LogIndex:     lr.Index,
		BlockNumber:  lr.BlockNumber,
		Seq:          seq, // sequence of erc transactions emitted by one log event - non-zero only for batch transfer events
	}); err != nil {
		log.Errorf("can not store token %s trx for call %s; %s", tokenType, lr.TxHash.String(), err.Error())
	}
}

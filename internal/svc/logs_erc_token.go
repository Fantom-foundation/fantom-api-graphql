// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// handleErcApproval handles Approval event on ERC20/721.
// event Approval(address indexed owner, address indexed spender, uint256 value)
// event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func handleErcApproval(lr *types.LogRecord) {
	if isErc20Transaction(lr) {
		processErc20Transaction(lr, types.TokenTrxTypeApproval)
		return
	}
	if isErc721Transaction(lr) {
		// do nothing on erc721 approval
		return
	}
	log.Debugf("Unrecognized ERC-20/721 Approval from tx %s (%d data bytes, %d topics)", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
}

// handleErcTransfer handles Transfer event on ERC20.
// event Transfer(address indexed from, address indexed to, uint256 value)
// event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func handleErcTransfer(lr *types.LogRecord) {
	if isErc20Transaction(lr) {
		processErc20Transaction(lr, types.TokenTrxTypeTransfer)
		return
	}
	if isErc721Transaction(lr) {
		processErc721Transfer(lr)
		return
	}
	log.Debugf("Unrecognized ERC-20/721 Transfer from tx %s (%d data bytes, %d topics)", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
}

// handleErc1155TransferSingle handles Transfer event on ERC1155
// event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 tokenId, uint256 value)
func handleErc1155TransferSingle(lr *types.LogRecord) {
	// 3 indexed params, 2 uint256 params
	if len(lr.Topics) == 4 && len(lr.Data) == 64 {
		// TODO: Handle transfer ownership
		return
	}
	log.Debugf("Unrecognized ERC1155 TransferSingle from tx %s (%d data bytes, %d topics)", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
}

// handleErc1155TransferBatch handles TransferBatch event on ERC1155
// event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func handleErc1155TransferBatch(lr *types.LogRecord) {
	// 3 indexed params
	if len(lr.Topics) == 4 {
		// TODO: Handle transfer ownership
		return
	}
	log.Debugf("Unrecognized ERC-1155 TransferBatch from tx %s (%d data bytes, %d topics)", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
}

// processErc20Transaction handles Approval and/or Transfer event on an ERC20 token.
func processErc20Transaction(lr *types.LogRecord, trxType int32) {
	from := common.BytesToAddress(lr.Topics[1].Bytes())
	to := common.BytesToAddress(lr.Topics[2].Bytes())
	amount := new(big.Int).SetBytes(lr.Data[:])

	decimals, err := repo.Erc20Decimals(&lr.Address)
	if err != nil {
		log.Errorf("can not get token decimals for token address %s, call %s; %s", lr.Address.String(), lr.TxHash.String(), err.Error())
		return
	}

	if err := repo.StoreTokenTransaction(&types.TokenTransaction{
		Transaction:  lr.TxHash,
		TrxIndex:     uint64(lr.TxIndex),
		TokenAddress: lr.Address,
		TrxType:      tokenTrxType(trxType, from, to),
		Sender:       from,
		Recipient:    to,
		Amount:       hexutil.Big(*amount),
		TimeStamp:    time.Unix(int64(lr.Block.TimeStamp), 0),
		LogIndex:     uint64(lr.Index),
		BlockNumber:  lr.BlockNumber,
	}, uint8(decimals)); err != nil {
		log.Errorf("can not store token transaction for call %s; %s", lr.TxHash.String(), err.Error())
	}
}

// handleErcTransfer handles Transfer event on ERC721.
func processErc721Transfer(lr *types.LogRecord) {
	// TODO: Handle transfer ownership
}

// tokenTrxType detects detailed type of ERC transfer based on common type and addresses involved.
func tokenTrxType(trxType int32, from common.Address, to common.Address) int32 {
	if trxType == types.TokenTrxTypeTransfer {
		if config.EmptyAddress == from.String() {
			return types.TokenTrxTypeMint
		}
		if config.EmptyAddress == to.String() {
			return types.TokenTrxTypeBurn
		}
	}
	return trxType
}

// isErc20Transaction checks is valid ERC20 transaction
func isErc20Transaction(lr *types.LogRecord) bool {
	// ERC20 has 2 indexed params (=> 3 topics) and 1 non-indexed uint256 param (=> 32 bytes)
	return len(lr.Topics) == 3 && len(lr.Data) == 32
}

// isErc721Transaction checks is valid ERC721 transaction
func isErc721Transaction(lr *types.LogRecord) bool {
	// ERC721 has 3 indexed params (=> 4 topics) and no non-indexed param (=> 0 bytes)
	return len(lr.Topics) == 4 && len(lr.Data) == 0
}

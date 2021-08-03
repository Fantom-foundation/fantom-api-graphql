// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// handleErcTokenApproval handles Approval event on an ERC20 token.
// event Approval(address indexed owner, address indexed spender, uint256 value)
func handleErcTokenApproval(lr *types.LogRecord) {
	handleErc20Transaction(lr, types.ERC20TrxTypeApproval)
}

// handleErcTokenTransfer handles Transfer event on an ERC20 token.
// event Transfer(address indexed from, address indexed to, uint256 value)
func handleErcTokenTransfer(lr *types.LogRecord) {
	handleErc20Transaction(lr, types.ERC20TrxTypeTransfer)
}

// handleErc20Transaction handles Approval and/or Transfer event on an ERC20 token.
// event Transfer(address indexed from, address indexed to, uint256 value);
// event Approval(address indexed owner, address indexed spender, uint256 value);
func handleErc20Transaction(lr *types.LogRecord, t int32) {
	// sanity check for data (1x uint256 (value) = 32 bytes)
	if len(lr.Data) != 32 || len(lr.Topics) != 3 {
		log.Errorf("%s log invalid for ERC20; expected 32 bytes data, %d bytes given, %d topics given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		handleErc721Transaction(lr, t)
		return
	}

	// handle it as an ERC20 transaction
	handleTokenTransaction(lr, types.AccountTypeERC20Token, t, new(big.Int).SetBytes(lr.Data[:]))
}

// handleErc721Transaction handles Approval and/or Transfer event of an ERC721 token.
// event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
// event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);
func handleErc721Transaction(lr *types.LogRecord, t int32) {
	// sanity check for data (1x uint256 (value) = 32 bytes)
	if len(lr.Data) != 0 || len(lr.Topics) != 4 {
		log.Errorf("%s log invalid for ERC721; expected 0 bytes data, %d bytes given, %d topics given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	// handle it as an ERC721 transaction
	handleTokenTransaction(lr, types.AccountTypeERC721Token, t, new(big.Int).SetBytes(lr.Topics[3].Bytes()))
}

// handleTokenTransaction handles general token transaction.
// This works on both ERC20 and ERC721 tokens.
// In case of ERC721 the <val> value represents the non-fungible token ID.
func handleTokenTransaction(lr *types.LogRecord, tokenType string, eventType int32, val *big.Int) {
	// get the addresses of participants from the log record
	from := common.BytesToAddress(lr.Topics[1].Bytes())
	to := common.BytesToAddress(lr.Topics[2].Bytes())

	// store the trx
	if err := repo.StoreErc20Transaction(&types.Erc20Transaction{
		Transaction:  lr.TxHash,
		TrxIndex:     hexutil.Uint64(uint64(lr.TxIndex)),
		TokenAddress: lr.Address,
		Type:         eventType,
		TokenType:    tokenType,
		Sender:       from,
		Recipient:    to,
		Amount:       (hexutil.Big)(*val),
		TimeStamp:    lr.Block.TimeStamp,
	}); err != nil {
		log.Errorf("can not store token %s trx for call %s; %s", tokenType, lr.TxHash.String(), err.Error())
	}
}

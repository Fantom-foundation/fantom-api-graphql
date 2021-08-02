/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// handleErcTokenApproval handles Approval event on an ERC20 token.
// event Approval(address indexed owner, address indexed spender, uint256 value)
func handleErcTokenApproval(log *types.LogRecord) {
	handleErc20Transaction(log, types.ERC20TrxTypeApproval)
}

// handleErcTokenTransfer handles Transfer event on an ERC20 token.
// event Transfer(address indexed from, address indexed to, uint256 value)
func handleErcTokenTransfer(log *types.LogRecord) {
	handleErc20Transaction(log, types.ERC20TrxTypeTransfer)
}

// handleErc20Transaction handles Approval and/or Transfer event on an ERC20 token.
// event Transfer(address indexed from, address indexed to, uint256 value);
// event Approval(address indexed owner, address indexed spender, uint256 value);
func handleErc20Transaction(log *types.LogRecord, t int32) {
	// sanity check for data (1x uint256 (value) = 32 bytes)
	if len(log.Data) != 32 || len(log.Topics) != 3 {
		R().Log().Errorf("%s log invalid for ERC20; expected 32 bytes data, %d bytes given, %d topics given", log.TxHash.String(), len(log.Data), len(log.Topics))
		handleErc721Transaction(log, t)
		return
	}

	// handle it as an ERC20 transaction
	handleTokenTransaction(log, types.AccountTypeERC20Token, t, new(big.Int).SetBytes(log.Data[:]))
}

// handleErc721Transaction handles Approval and/or Transfer event of an ERC721 token.
// event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
// event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);
func handleErc721Transaction(log *types.LogRecord, t int32) {
	// sanity check for data (1x uint256 (value) = 32 bytes)
	if len(log.Data) != 0 || len(log.Topics) != 4 {
		R().Log().Errorf("%s log invalid for ERC721; expected 0 bytes data, %d bytes given, %d topics given", log.TxHash.String(), len(log.Data), len(log.Topics))
		return
	}

	// handle it as an ERC721 transaction
	handleTokenTransaction(log, types.AccountTypeERC721Token, t, new(big.Int).SetBytes(log.Topics[3].Bytes()))
}

// handleTokenTransaction handles general token transaction.
// This works on both ERC20 and ERC721 tokens.
// In case of ERC721 the <val> value represents the non-fungible token ID.
func handleTokenTransaction(log *types.LogRecord, tokenType string, eventType int32, val *big.Int) {
	// get the addresses of participants from the log record
	from := common.BytesToAddress(log.Topics[1].Bytes())
	to := common.BytesToAddress(log.Topics[2].Bytes())

	// store the trx
	if err := R().StoreErc20Transaction(&types.Erc20Transaction{
		Transaction:  log.TxHash,
		TrxIndex:     hexutil.Uint64(uint64(log.TxIndex)),
		TokenAddress: log.Address,
		Type:         eventType,
		TokenType:    tokenType,
		Sender:       from,
		Recipient:    to,
		Amount:       (hexutil.Big)(*val),
		TimeStamp:    log.Block.TimeStamp,
	}); err != nil {
		R().Log().Errorf("can not store token %s trx for call %s; %s", tokenType, log.TxHash.String(), err.Error())
	}
}

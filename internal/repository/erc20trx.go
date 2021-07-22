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
	retypes "github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
)

// StoreErc20Transaction stores ERC20 transaction into the repository.
func (p *proxy) StoreErc20Transaction(trx *types.Erc20Transaction) error {
	return p.db.AddERC20Transaction(trx)
}

// Erc20Transactions provides list of ERC20 transactions based on given filters.
func (p *proxy) Erc20Transactions(token *common.Address, acc *common.Address, tt *int32, cursor *string, count int32) (*types.Erc20TransactionList, error) {
	// prep the filter
	fi := bson.D{}

	// filter specific token
	if token != nil {
		fi = append(fi, bson.E{
			Key:   types.FiErc20TransactionToken,
			Value: token.String(),
		})
	}

	// common address (sender or recipient)
	if acc != nil {
		fi = append(fi, bson.E{
			Key: "$or",
			Value: bson.A{bson.D{{
				Key:   types.FiErc20TransactionSender,
				Value: acc.String(),
			}}, bson.D{{
				Key:   types.FiErc20TransactionRecipient,
				Value: acc.String(),
			}}},
		})
	}

	// type of the transaction
	if tt != nil {
		fi = append(fi, bson.E{
			Key:   types.FiErc20TransactionType,
			Value: *tt,
		})
	}

	// do loading
	return p.db.Erc20Transactions(cursor, count, &fi)
}

// Erc20Assets provides a list of known assets for the given owner.
func (p *proxy) Erc20Assets(owner common.Address, count int32) ([]common.Address, error) {
	return p.db.Erc20Assets(owner, count)
}

// handleErc20Approval handles Approval event on an ERC20 token.
// event Approval(address indexed owner, address indexed spender, uint256 value)
func handleErc20Approval(log *retypes.Log, ld *logsDispatcher) {
	handleErc20Transaction(log, ld, types.ERC20TrxTypeApproval)
}

// handleErc20Transfer handles Transfer event on an ERC20 token.
// event Transfer(address indexed from, address indexed to, uint256 value)
func handleErc20Transfer(log *retypes.Log, ld *logsDispatcher) {
	handleErc20Transaction(log, ld, types.ERC20TrxTypeTransfer)
}

// handleErc20Transaction handles Approval and/or Transfer event on an ERC20 token.
// event Transfer(address indexed from, address indexed to, uint256 value);
// event Approval(address indexed owner, address indexed spender, uint256 value);
func handleErc20Transaction(log *retypes.Log, ld *logsDispatcher, t int32) {
	// sanity check for data (1x uint256 (value) = 32 bytes)
	if len(log.Data) != 32 || len(log.Topics) != 3 {
		ld.log.Errorf("%s log invalid for ERC20; expected 32 bytes data, %d bytes given, %d topics given", log.TxHash.String(), len(log.Data), len(log.Topics))
		handleErc721Transaction(log, ld, t)
		return
	}

	// handle it as an ERC20 transaction
	handleTokenTransaction(log, ld, t, types.AccountTypeERC20Token, new(big.Int).SetBytes(log.Data[:]))
}

// handleErc721Transaction handles Approval and/or Transfer event of an ERC721 token.
// event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
// event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);
func handleErc721Transaction(log *retypes.Log, ld *logsDispatcher, t int32) {
	// sanity check for data (1x uint256 (value) = 32 bytes)
	if len(log.Data) != 0 || len(log.Topics) != 4 {
		ld.log.Errorf("%s log invalid for ERC721; expected 0 bytes data, %d bytes given, %d topics given", log.TxHash.String(), len(log.Data), len(log.Topics))
		return
	}

	// handle it as an ERC20 transaction
	handleTokenTransaction(log, ld, t, types.AccountTypeERC721Token, new(big.Int).SetBytes(log.Topics[3].Bytes()))
}

// handleTokenTransaction handles general token transaction.
// This works on both ERC20 and ERC721 tokens.
func handleTokenTransaction(log *retypes.Log, ld *logsDispatcher, t int32, tType string, val *big.Int) {
	// get the block
	blk := hexutil.Uint64(log.BlockNumber)
	block, err := ld.repo.BlockByNumber(&blk)
	if err != nil {
		ld.log.Errorf("can not decode block #%d from log record; %s", blk, err.Error())
		return
	}

	// get the data elements from the log record
	from := common.BytesToAddress(log.Topics[1].Bytes())
	to := common.BytesToAddress(log.Topics[2].Bytes())

	// store the trx
	if err := ld.repo.StoreErc20Transaction(&types.Erc20Transaction{
		Transaction:  log.TxHash,
		TrxIndex:     hexutil.Uint64(uint64(log.TxIndex)),
		TokenAddress: log.Address,
		Type:         t,
		TokenType:    tType,
		Sender:       from,
		Recipient:    to,
		Amount:       (hexutil.Big)(*val),
		TimeStamp:    block.TimeStamp,
	}); err != nil {
		ld.log.Errorf("can not store token trx for call %s; %s", log.TxHash.String(), err.Error())
	}
}

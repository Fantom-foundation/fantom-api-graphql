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
func handleErc20Transaction(log *retypes.Log, ld *logsDispatcher, t int32) {
	// sanity check for data (1x uint256 = 32 bytes)
	if len(log.Data) != 32 {
		ld.log.Criticalf("%s log invalid data length; expected 32 bytes, given %d bytes", log.TxHash.String(), len(log.Data))
		return
	}

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
	amo := new(big.Int).SetBytes(log.Data[:])

	// store the trx
	if err := ld.repo.StoreErc20Transaction(&types.Erc20Transaction{
		Transaction: log.TxHash,
		TrxIndex:    hexutil.Uint64(uint64(log.TxIndex)),
		Token:       log.Address,
		Type:        t,
		Sender:      from,
		Recipient:   to,
		Amount:      (hexutil.Big)(*amo),
		TimeStamp:   block.TimeStamp,
	}); err != nil {
		ld.log.Errorf("can not store ERC20 transaction; %s", err.Error())
	}
}

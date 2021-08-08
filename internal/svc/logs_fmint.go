// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// handleFMintDeposit handles a new deposit on fMint contract.
// event Deposited(address indexed token, address indexed user, uint256 amount)
func handleFMintDeposit(lr *types.LogRecord) {
	// sanity check for data (1 uint256 = 32 bytes); call + token + user = 3 topics
	if len(lr.Data) != 32 || len(lr.Topics) != 3 {
		log.Criticalf("%s invalid event; expected 32 bytes, %d bytes given; expected 3 topics, %d given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	handleNewFMintRecord(
		lr,
		types.FMintTrxTypeDeposit,
		common.BytesToAddress(lr.Topics[2].Bytes()),
		common.BytesToAddress(lr.Topics[1].Bytes()),
		new(big.Int).SetBytes(lr.Data),
		new(big.Int),
	)
}

// handleFMintWithdraw handles a new withdrawal on fMint contract.
// event Withdrawn(address indexed token, address indexed user, uint256 amount)
func handleFMintWithdraw(lr *types.LogRecord) {
	// sanity check for data (1 uint256 = 32 bytes); call + token + user = 3 topics
	if len(lr.Data) != 32 || len(lr.Topics) != 3 {
		log.Criticalf("%s invalid event; expected 32 bytes, %d bytes given; expected 3 topics, %d given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	handleNewFMintRecord(
		lr,
		types.FMintTrxTypeWithdraw,
		common.BytesToAddress(lr.Topics[2].Bytes()),
		common.BytesToAddress(lr.Topics[1].Bytes()),
		new(big.Int).SetBytes(lr.Data),
		new(big.Int),
	)
}

// handleFMintMint handles a new mint (tokens borrow) on fMint contract.
// event Minted(address indexed token, address indexed user, uint256 amount, uint256 fee)
func handleFMintMint(lr *types.LogRecord) {
	// sanity check for data (2 uint256 = 64 bytes); call + token + user = 3 topics
	if len(lr.Data) != 64 || len(lr.Topics) != 3 {
		log.Criticalf("%s invalid event; expected 64 bytes, %d bytes given; expected 3 topics, %d given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	handleNewFMintRecord(
		lr,
		types.FMintTrxTypeMint,
		common.BytesToAddress(lr.Topics[2].Bytes()),
		common.BytesToAddress(lr.Topics[1].Bytes()),
		new(big.Int).SetBytes(lr.Data[:32]),
		new(big.Int).SetBytes(lr.Data[32:]),
	)
}

// handleFMintRepay handles a new repay (debt repay) on fMint contract.
// event Repaid(address indexed token, address indexed user, uint256 amount)
func handleFMintRepay(lr *types.LogRecord) {
	// sanity check for data (1 uint256 = 32 bytes); call + token + user = 3 topics
	if len(lr.Data) != 32 || len(lr.Topics) != 3 {
		log.Criticalf("%s invalid event; expected 32 bytes, %d bytes given; expected 3 topics, %d given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	handleNewFMintRecord(
		lr,
		types.FMintTrxTypeRepay,
		common.BytesToAddress(lr.Topics[2].Bytes()),
		common.BytesToAddress(lr.Topics[1].Bytes()),
		new(big.Int).SetBytes(lr.Data),
		new(big.Int),
	)
}

// handleNewFMintRecord creates an fMint record with the given data
// and pushes it into the persistent storage for future reference.
func handleNewFMintRecord(lr *types.LogRecord, tp int32, user common.Address, token common.Address, amount *big.Int, fee *big.Int) {
	err := repo.AddFMintTransaction(&types.FMintTransaction{
		UserAddress:  user,
		TokenAddress: token,
		Type:         tp,
		Amount:       (hexutil.Big)(*amount),
		Fee:          (hexutil.Big)(*fee),
		TrxHash:      lr.TxHash,
		TrxIndex:     int64(lr.TxIndex)<<8 ^ int64(lr.Index),
		TimeStamp:    lr.Block.TimeStamp,
	})
	if err != nil {
		log.Errorf("can not register fMint trx %s; %s", lr.TxHash.String(), err.Error())
	}
}

// handleFMintReward handles a new reward claim on fMint contract.
// event RewardPaid(address indexed user, uint256 reward)
func handleFMintReward(lr *types.LogRecord) {
	// sanity check for data (1 uint256 = 32 bytes); call + user = 2 topics
	if len(lr.Data) != 32 || len(lr.Topics) != 2 {
		log.Criticalf("%s invalid event; expected 32 bytes, %d bytes given; expected 2 topics, %d given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

}

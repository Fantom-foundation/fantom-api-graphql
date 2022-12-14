// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// burnDispatcher implements dispatcher of new native FTM burns on the blockchain.
type burnDispatcher struct {
	service
	inTransaction chan *eventTrx
}

// name returns the name of the service used by orchestrator.
func (bud *burnDispatcher) name() string {
	return "native burn dispatcher"
}

// init prepares the transaction dispatcher to perform its function.
func (bud *burnDispatcher) init() {
	bud.sigStop = make(chan struct{})
}

// run starts the transaction dispatcher job
func (bud *burnDispatcher) run() {
	// make sure we are orchestrated
	if bud.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", bud.name()))
	}

	// signal orchestrator we started and go
	bud.mgr.started(bud)
	go bud.execute()
}

// execute runs the main loop of the burn dispatcher.
func (bud *burnDispatcher) execute() {
	defer func() {
		bud.mgr.finished(bud)
	}()

	var current *types.FtmBurn
	for {
		select {
		case <-bud.sigStop:
			return
		case tx, ok := <-bud.inTransaction:
			if !ok {
				return
			}
			current = bud.process(tx, current)
		}
	}
}

// process incoming transaction event to extract the burn information.
func (bud *burnDispatcher) process(tx *eventTrx, burn *types.FtmBurn) *types.FtmBurn {
	txFee, txTreasury, txBurn, txReward := bud.burnedFee(tx.trx)

	// no previous burn to check against? make a new record
	if burn == nil {
		return &types.FtmBurn{
			BlockNumber:    tx.blk.Number,
			BlkTimeStamp:   time.Unix(int64(tx.blk.TimeStamp), 0),
			BurnAmount:     hexutil.Big(*txBurn),
			FeeAmount:      hexutil.Big(*txFee),
			TreasuryAmount: hexutil.Big(*txTreasury),
			RewardsAmount:  hexutil.Big(*txReward),
			TxList:         append(make([]common.Hash, 0), tx.trx.Hash),
		}
	}

	// a new block may have been received
	if tx.blk.Number != burn.BlockNumber {
		val := float64(new(big.Int).Div((*big.Int)(&burn.BurnAmount), types.BurnDecimalsCorrection).Int64()) / 1_000_000
		log.Debugf("collected block burn of %.4f FTM at #%d", val, burn.BlockNumber)

		if err := repo.StoreFtmBurn(burn); err != nil {
			log.Warningf("could not store previous burn; %s", err.Error())
		}

		return &types.FtmBurn{
			BlockNumber:    tx.blk.Number,
			BlkTimeStamp:   time.Unix(int64(tx.blk.TimeStamp), 0),
			BurnAmount:     hexutil.Big(*txBurn),
			FeeAmount:      hexutil.Big(*txFee),
			TreasuryAmount: hexutil.Big(*txTreasury),
			RewardsAmount:  hexutil.Big(*txReward),
			TxList:         append(make([]common.Hash, 0), tx.trx.Hash),
		}
	}

	// just add this fee to the existing
	burn.BurnAmount = hexutil.Big(*new(big.Int).Add((*big.Int)(&burn.BurnAmount), txBurn))
	burn.FeeAmount = hexutil.Big(*new(big.Int).Add((*big.Int)(&burn.FeeAmount), txFee))
	burn.TreasuryAmount = hexutil.Big(*new(big.Int).Add((*big.Int)(&burn.TreasuryAmount), txTreasury))
	burn.RewardsAmount = hexutil.Big(*new(big.Int).Add((*big.Int)(&burn.RewardsAmount), txReward))

	// remember the transaction ref
	burn.TxList = append(burn.TxList, tx.trx.Hash)
	return burn
}

// burnedFee calculates the amount of burned FTMs from the transaction fee.
func (bud *burnDispatcher) burnedFee(trx *types.Transaction) (*big.Int, *big.Int, *big.Int, *big.Int) {
	// the transaction is pending, no fees are stashed yet
	if trx.BlockNumber == nil {
		return new(big.Int), new(big.Int), new(big.Int), new(big.Int)
	}

	// calculate the total amount of fee paid for the transaction in consumed gas
	fee := new(big.Int).Mul((*big.Int)(&trx.GasPrice), big.NewInt(int64(*trx.GasUsed)))

	// find the appropriate share for this transaction block
	share := repo.BurnTreasuryStashShareByBlock(uint64(*trx.BlockNumber))
	if share == nil {
		return fee, new(big.Int), new(big.Int), new(big.Int)
	}

	treasury := new(big.Int).Div(new(big.Int).Mul(fee, share.ToTreasury), share.DigitCorrection)
	burn := new(big.Int).Div(new(big.Int).Mul(fee, share.ToBurn), share.DigitCorrection)
	reward := new(big.Int).Div(new(big.Int).Mul(fee, share.ToRewards), share.DigitCorrection)
	return fee, treasury, burn, reward
}

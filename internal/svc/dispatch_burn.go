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

var (
	// feePartToBurn represents the percentage of the transaction fee burned in 4 digits fixed int.
	feePartToBurn = big.NewInt(300)

	// feeBurnDigitCorrection reduces the fixed number of amount digits back to WEI after burn calculation.
	feeBurnDigitCorrection = big.NewInt(1000)
)

// name returns the name of the service used by orchestrator.
func (bud *burnDispatcher) name() string {
	return "native burn dispatcher"
}

// init prepares the transaction dispatcher to perform its function.
func (bud *burnDispatcher) init() {
	bud.sigStop = make(chan bool, 1)
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

// close terminates the block dispatcher.
func (bud *burnDispatcher) close() {
	if bud.sigStop != nil {
		bud.sigStop <- true
	}
}

// execute runs the main loop of the burn dispatcher.
func (bud *burnDispatcher) execute() {
	defer func() {
		close(bud.sigStop)
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
	// no previous burn to check against? make a new record
	if burn == nil {
		return &types.FtmBurn{
			BlockNumber:  tx.blk.Number,
			BlkTimeStamp: time.Unix(int64(tx.blk.TimeStamp), 0),
			Amount:       hexutil.Big(*bud.burnedFee(tx.trx)),
			TxList:       append(make([]common.Hash, 0), tx.trx.Hash),
		}
	}

	// a new block may have been received
	if tx.blk.Number != burn.BlockNumber {
		val := float64(new(big.Int).Div((*big.Int)(&burn.Amount), types.BurnDecimalsCorrection).Int64()) / 1_000_000
		log.Debugf("collected block burn of %.4f FTM at #%d", val, burn.BlockNumber)

		if err := repo.StoreFtmBurn(burn); err != nil {
			log.Warningf("could not store previous burn; %s", err.Error())
		}

		return &types.FtmBurn{
			BlockNumber:  tx.blk.Number,
			BlkTimeStamp: time.Unix(int64(tx.blk.TimeStamp), 0),
			Amount:       hexutil.Big(*bud.burnedFee(tx.trx)),
			TxList:       append(make([]common.Hash, 0), tx.trx.Hash),
		}
	}

	// just add this fee to the existing
	burn.Amount = hexutil.Big(*new(big.Int).Add((*big.Int)(&burn.Amount), bud.burnedFee(tx.trx)))
	burn.TxList = append(burn.TxList, tx.trx.Hash)
	return burn
}

// burnedFee calculates the amount of burned FTMs from the transaction fee.
func (bud *burnDispatcher) burnedFee(trx *types.Transaction) *big.Int {
	// calculate the total amount of fee paid for the transaction in consumed gas
	fee := new(big.Int).Mul((*big.Int)(&trx.GasPrice), big.NewInt(int64(*trx.GasUsed)))

	// now get 30% by multiplying by 300 and dividing by 1000
	return new(big.Int).Div(new(big.Int).Mul(fee, feePartToBurn), feeBurnDigitCorrection)
}

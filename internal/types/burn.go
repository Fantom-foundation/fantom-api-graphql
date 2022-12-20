// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
	"time"
)

var (
	// BurnDecimalsCorrection is used to reduce precision of an amount of burned FTM
	BurnDecimalsCorrection = new(big.Int).SetUint64(1_000_000_000)

	// BurnFTMDecimalsCorrection is used to convert reduced precision burned amount to FTM units.
	BurnFTMDecimalsCorrection = float64(1_000_000_000)
)

// FtmBurn represents deflation of native tokens by burning.
type FtmBurn struct {
	BlockNumber    hexutil.Uint64 `bson:"block"`
	BlkTimeStamp   time.Time      `bson:"ts"`
	BurnAmount     hexutil.Big    `bson:"amount"`
	TreasuryAmount hexutil.Big    `bson:"treasury"`
	FeeAmount      hexutil.Big    `bson:"fee"`
	RewardsAmount  hexutil.Big    `bson:"rewards"`
	TxList         []common.Hash  `bson:"tx_list"`
}

// FtmDailyBurn represents a burn of native tokens by days.
type FtmDailyBurn struct {
	TickDate       time.Time `bson:"_id"`
	BlocksCount    int32     `bson:"blocks_count"`
	BurnedAmount   int64     `bson:"burn_amount"`
	TreasuryAmount int64     `bson:"treasury_amount"`
	RewardsAmount  int64     `bson:"rewards_amount"`
	FeeAmount      int64     `bson:"fee_amount"`
}

// MarshalBSON returns a BSON document for the FTM burn.
func (burn *FtmBurn) MarshalBSON() ([]byte, error) {
	amount := new(big.Int).Div(burn.BurnAmount.ToInt(), BurnDecimalsCorrection)
	fee := new(big.Int).Div(burn.FeeAmount.ToInt(), BurnDecimalsCorrection)
	treasury := new(big.Int).Div(burn.TreasuryAmount.ToInt(), BurnDecimalsCorrection)
	rewards := new(big.Int).Div(burn.RewardsAmount.ToInt(), BurnDecimalsCorrection)

	row := struct {
		Block          int64     `bson:"block"`
		TimeStamp      time.Time `bson:"ts"`
		Value          string    `bson:"value"`
		Amount         int64     `bson:"amount"`
		FeeValue       string    `bson:"fee_value"`
		FeeAmount      int64     `bson:"fee_amount"`
		TreasuryValue  string    `bson:"try_value"`
		TreasuryAmount int64     `bson:"try_amount"`
		RewardsValue   string    `bson:"rew_value"`
		RewardsAmount  int64     `bson:"rew_amount"`
		TxList         []string  `bson:"tx_list"`
	}{
		Block:          int64(burn.BlockNumber),
		TimeStamp:      burn.BlkTimeStamp,
		Value:          burn.BurnAmount.String(),
		Amount:         amount.Int64(),
		FeeValue:       burn.FeeAmount.String(),
		FeeAmount:      fee.Int64(),
		TreasuryValue:  burn.TreasuryAmount.String(),
		TreasuryAmount: treasury.Int64(),
		RewardsValue:   burn.RewardsAmount.String(),
		RewardsAmount:  rewards.Int64(),
		TxList:         make([]string, len(burn.TxList)),
	}
	for i, v := range burn.TxList {
		row.TxList[i] = v.String()
	}
	return bson.Marshal(row)
}

// UnmarshalBSON updates the value from BSON source.
func (burn *FtmBurn) UnmarshalBSON(data []byte) (err error) {
	var row struct {
		Block         int64     `bson:"block"`
		TimeStamp     time.Time `bson:"ts"`
		Value         string    `bson:"value"`
		FeeValue      string    `bson:"fee_value"`
		TreasuryValue string    `bson:"try_value"`
		RewardsValue  string    `bson:"rew_value"`
		TxList        []string  `bson:"tx_list"`
	}

	err = bson.Unmarshal(data, &row)
	if err != nil {
		return err
	}

	burn.BlockNumber = hexutil.Uint64(row.Block)
	burn.BlkTimeStamp = row.TimeStamp
	burn.BurnAmount = (hexutil.Big)(*hexutil.MustDecodeBig(row.Value))
	burn.FeeAmount = (hexutil.Big)(*hexutil.MustDecodeBig(row.FeeValue))
	burn.TreasuryAmount = (hexutil.Big)(*hexutil.MustDecodeBig(row.TreasuryValue))
	burn.RewardsAmount = (hexutil.Big)(*hexutil.MustDecodeBig(row.RewardsValue))

	burn.TxList = make([]common.Hash, len(row.TxList))
	for i, v := range row.TxList {
		burn.TxList[i] = common.HexToHash(v)
	}
	return nil
}

// Timestamp return UNIX stamp of the burn.
func (burn *FtmBurn) Timestamp() hexutil.Uint64 {
	return hexutil.Uint64(burn.BlkTimeStamp.Unix())
}

// FtmValue returns FTM amount of burned tokens.
func (burn *FtmBurn) FtmValue() float64 {
	return float64(new(big.Int).Div(burn.BurnAmount.ToInt(), BurnDecimalsCorrection).Int64()) / BurnFTMDecimalsCorrection
}

// Value returns FTM amount of burned tokens.
func (burn *FtmBurn) Value() int64 {
	return new(big.Int).Div(burn.BurnAmount.ToInt(), BurnDecimalsCorrection).Int64()
}

// Amount returns amount of burned tokens in WEI.
func (burn *FtmBurn) Amount() hexutil.Big {
	return burn.BurnAmount
}

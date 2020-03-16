package types

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Epoch represents epoch detail.
type Epoch struct {
	Id                     hexutil.Uint64
	EndTime                hexutil.Big
	Duration               hexutil.Big
	EpochFee               hexutil.Big
	TotalBaseRewardWeight  hexutil.Big
	TotalTxRewardWeight    hexutil.Big
	BaseRewardPerSecond    hexutil.Big
	StakeTotalAmount       hexutil.Big
	DelegationsTotalAmount hexutil.Big
	TotalSupply            hexutil.Big
}

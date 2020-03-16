package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Staker represents a staker information.
type Staker struct {
	Id                       hexutil.Uint64 `json:"id"`
	StakerAddress            common.Address `json:"address"`
	TotalStake               *hexutil.Big   `json:"totalStake"`
	Stake                    *hexutil.Big   `json:"stake"`
	DelegatedMe              *hexutil.Big   `json:"delegatedMe"`
	IsValidator              bool           `json:"isValidator"`
	IsActive                 bool           `json:"isActive"`
	IsCheater                bool           `json:"isCheater"`
	IsOffline                bool           `json:"isOffline"`
	CreatedEpoch             hexutil.Uint64 `json:"createdEpoch"`
	CreatedTime              hexutil.Uint64 `json:"createdTime"`
	DeactivatedEpoch         hexutil.Uint64 `json:"deactivatedEpoch"`
	DeactivatedTime          hexutil.Uint64 `json:"deactivatedTime"`
	MissedBlocks             hexutil.Uint64 `json:"missedBlocks"`
	Downtime                 hexutil.Uint64 `json:"downtime"`
	Poi                      *hexutil.Big   `json:"poi"`
	BaseRewardWeight         *hexutil.Big   `json:"baseRewardWeight"`
	TxRewardWeight           *hexutil.Big   `json:"txRewardWeight"`
	ValidationScore          *hexutil.Big   `json:"validationScore"`
	OriginationScore         *hexutil.Big   `json:"originationScore"`
	ClaimedRewards           *hexutil.Big   `json:"claimedRewards"`
	DelegationClaimedRewards *hexutil.Big   `json:"delegatorsClaimedRewards"`
}

// Package types implements different core types of the API.
package types

import (
	"encoding/json"
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
	Status                   hexutil.Uint64 `json:"status"`

	/* delegation limits */
	TotalDelegatedLimit hexutil.Big
	DelegatedLimit      hexutil.Big
}

// StakerInfo holds extended staker information.
type StakerInfo struct {
	// Name represents the name of the staker
	Name *string `json:"name"`

	// LogoUrl represents staker logo URL
	LogoUrl *string `json:"logoUrl"`

	// Website represents a link to staker website
	Website *string `json:"website"`

	// Contact represents a link to contact to the staker
	Contact *string `json:"contact"`
}

// UnmarshalStakerInfo parses the JSON-encoded staker information data.
func UnmarshalStakerInfo(data []byte) (*StakerInfo, error) {
	var sti StakerInfo
	err := json.Unmarshal(data, &sti)
	return &sti, err
}

// Marshal returns the JSON encoding of staker information.
func (sti *StakerInfo) Marshal() ([]byte, error) {
	return json.Marshal(sti)
}

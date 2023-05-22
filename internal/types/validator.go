// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"time"
)

// Validator represents a validator information.
type Validator struct {
	Id               hexutil.Big    `json:"id"`
	StakerAddress    common.Address `json:"address"`
	TotalStake       *hexutil.Big   `json:"totalStake"`
	Status           hexutil.Uint64 `json:"status"`
	CreatedEpoch     hexutil.Uint64 `json:"createdEpoch"`
	CreatedTime      hexutil.Uint64 `json:"createdTime"`
	DeactivatedEpoch hexutil.Uint64 `json:"deactivatedEpoch"`
	DeactivatedTime  hexutil.Uint64 `json:"deactivatedTime"`
}

// OfflineValidator represents a validator with current downtime.
type OfflineValidator struct {
	ID         uint64
	Address    common.Address
	DownTime   Downtime
	DownBlocks uint64
}

type Downtime uint64

func (d Downtime) String() string {
	return time.Duration(d).Round(time.Second).String()
}

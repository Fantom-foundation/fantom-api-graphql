// Package types implements different core types of the API.
package types

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// Epoch represents epoch detail.
type Epoch struct {
	Id                    hexutil.Uint64 `json:"id"`
	EndTime               hexutil.Uint64 `json:"end"`
	EpochFee              hexutil.Big    `json:"fee"`
	TotalBaseRewardWeight hexutil.Big    `json:"trw"`
	TotalTxRewardWeight   hexutil.Big    `json:"txw"`
	BaseRewardPerSecond   hexutil.Big    `json:"brw"`
	StakeTotalAmount      hexutil.Big    `json:"stk"`
	TotalSupply           hexutil.Big    `json:"sup"`
}

// BsonEpoch represents the epoch data structure for BSON formatting.
type BsonEpoch struct {
	ID                  int64     `bson:"_id"`
	EndTime             int64     `bson:"et"`
	End                 time.Time `bson:"end"`
	Fee                 string    `bson:"fee"`
	BaseRewardWeight    string    `bson:"brw"`
	TrxRewardWeight     string    `bson:"trw"`
	BaseRewardPerSecond string    `bson:"rew"`
	TotalStake          string    `bson:"stake"`
	TotalSupply         string    `bson:"supply"`
}

// UnmarshalEpoch parses the JSON-encoded Epoch data.
func UnmarshalEpoch(data []byte) (*Epoch, error) {
	var ep Epoch
	err := json.Unmarshal(data, &ep)
	return &ep, err
}

// Marshal returns the JSON encoding of Epoch.
func (e *Epoch) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

// MarshalBSON creates a BSON representation of the Epoch record.
func (e *Epoch) MarshalBSON() ([]byte, error) {
	// prep the structure for saving
	return bson.Marshal(BsonEpoch{
		ID:                  int64(e.Id),
		EndTime:             int64(e.EndTime),
		End:                 time.Unix(int64(e.EndTime), 0),
		Fee:                 e.EpochFee.String(),
		BaseRewardWeight:    e.TotalBaseRewardWeight.String(),
		TrxRewardWeight:     e.TotalTxRewardWeight.String(),
		BaseRewardPerSecond: e.BaseRewardPerSecond.String(),
		TotalStake:          e.StakeTotalAmount.String(),
		TotalSupply:         e.TotalSupply.String(),
	})
}

// UnmarshalBSON updates the value from BSON source.
func (e *Epoch) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode stored epoch")
		}
	}()

	// try to decode BSON data
	var row BsonEpoch
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	// transfer the data points
	e.Id = (hexutil.Uint64)(row.ID)
	e.EndTime = (hexutil.Uint64)(row.EndTime)
	e.EpochFee = (hexutil.Big)(*hexutil.MustDecodeBig(row.Fee))
	e.TotalBaseRewardWeight = (hexutil.Big)(*hexutil.MustDecodeBig(row.BaseRewardWeight))
	e.TotalTxRewardWeight = (hexutil.Big)(*hexutil.MustDecodeBig(row.TrxRewardWeight))
	e.BaseRewardPerSecond = (hexutil.Big)(*hexutil.MustDecodeBig(row.BaseRewardPerSecond))
	e.StakeTotalAmount = (hexutil.Big)(*hexutil.MustDecodeBig(row.TotalStake))
	e.TotalSupply = (hexutil.Big)(*hexutil.MustDecodeBig(row.TotalSupply))
	return nil
}

// Package types implements different core types of the API.
package types

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
)

// Delegation represents a delegator in Opera blockchain.
type Delegation struct {
	Transaction     common.Hash    `json:"trx"`
	Address         common.Address `json:"address"`
	ToStakerId      *hexutil.Big   `json:"toStakerID"`
	AmountDelegated *hexutil.Big   `json:"amount"`
	CreatedTime     hexutil.Uint64 `json:"createdTime"`
}

// Uid returns a unique identifier for the given delegation.
// We construct the UID from the time the delegation was created, a part of the creation transaction hash,
// and part of the target validator index.
func (dl *Delegation) Uid() uint64 {
	return uint64(dl.CreatedTime)<<16 | (dl.ToStakerId.ToInt().Uint64()&0xFF)<<8 | (binary.BigEndian.Uint64(dl.Transaction[:8]) & 0xFF)
}

// MarshalBSON creates a BSON representation of the delegation record.
func (dl *Delegation) MarshalBSON() ([]byte, error) {
	pom := struct {
		Uid    uint64 `bson:"_id"`
		Trx    string `bson:"trx"`
		Addr   string `bson:"addr"`
		To     string `bson:"to"`
		CrTime uint64 `bson:"cr_time"`
		Amount string `bson:"amount"`
	}{
		Uid:    dl.Uid(),
		Trx:    dl.Transaction.String(),
		Addr:   dl.Address.String(),
		To:     dl.ToStakerId.String(),
		CrTime: uint64(dl.CreatedTime),
		Amount: dl.AmountDelegated.String(),
	}
	return bson.Marshal(pom)
}

// UnmarshalBSON updates the value from BSON source.
func (dl *Delegation) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode BIG number in delegation unmarshal")
		}
	}()

	// try to decode the BSON data
	var row struct {
		Uid    uint64 `bson:"_id"`
		Trx    string `bson:"trx"`
		Addr   string `bson:"addr"`
		To     string `bson:"to"`
		CrTime uint64 `bson:"cr_time"`
		Amount string `bson:"amount"`
	}
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	// transfer values
	dl.Transaction = common.HexToHash(row.Trx)
	dl.Address = common.HexToAddress(row.Addr)
	dl.ToStakerId = (*hexutil.Big)(hexutil.MustDecodeBig(row.To))
	dl.CreatedTime = hexutil.Uint64(row.CrTime)
	dl.AmountDelegated = (*hexutil.Big)(hexutil.MustDecodeBig(row.Amount))
	return nil
}

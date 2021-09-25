// Package types implements different core types of the API.
package types

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	FiNFTOwnerPK            = "_id"
	FiNFTOwnerContract      = "contract"
	FiNFTOwnerAddress       = "owner"
	FiNFTTokenID            = "token"
	FiNFTTokenOwnerQuantity = "qty"
)

// NFTOwner represents a non-fungible token owner reference structure.
type NFTOwner struct {
	ContractAddress common.Address `json:"contractAddress"`
	TokenId         hexutil.Big    `json:"tokenId"`
	TokenType       string         `json:"tokenType"`
	OwnerAddress    common.Address `json:"ownerAddress"`
	Quantity        hexutil.Big    `json:"qty"`
	TimeStamp       time.Time      `json:"timeStamp"`
	BlockNumber     hexutil.Uint64 `json:"blockNumber"`
	LogIndex        hexutil.Uint64 `json:"logIndex"`
}

// NFTOwnerRow represents the database row af on NFT owner record.
type NFTOwnerRow struct {
	Ordinal   int64     `bson:"_id"`
	Contract  string    `bson:"contract"`
	TokenID   string    `bson:"token"`
	TokenType string    `bson:"type"`
	Owner     string    `bson:"owner"`
	Qty       string    `bson:"qty"`
	TimeStamp time.Time `bson:"since"`
}

// OrdinalIndex generates ordinal index of the NFT token owner record.
func (nfo *NFTOwner) OrdinalIndex() int64 {
	return (int64(nfo.BlockNumber)&0x7FFFFFFFFFF)<<20 | int64(nfo.LogIndex)
}

// MarshalBSON returns a BSON document for the NFT owner record.
func (nfo *NFTOwner) MarshalBSON() ([]byte, error) {
	row := NFTOwnerRow{
		Ordinal:   nfo.OrdinalIndex(),
		Contract:  nfo.ContractAddress.String(),
		TokenID:   nfo.TokenId.String(),
		TokenType: nfo.TokenType,
		Owner:     nfo.OwnerAddress.String(),
		Qty:       nfo.Quantity.String(),
		TimeStamp: time.Time{},
	}
	return bson.Marshal(row)
}

// UnmarshalBSON updates the NFT owner record from provided BSON source.
func (nfo *NFTOwner) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode BIG number in NFT owner unmarshal")
		}
	}()

	var row NFTOwnerRow
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	nfo.TokenType = row.TokenType
	nfo.ContractAddress = common.HexToAddress(row.Contract)
	nfo.OwnerAddress = common.HexToAddress(row.Owner)
	nfo.TimeStamp = row.TimeStamp
	nfo.BlockNumber = hexutil.Uint64(row.Ordinal >> 20)
	nfo.LogIndex = hexutil.Uint64(row.Ordinal & 0xFFFFF)

	val, err := hexutil.DecodeBig(row.TokenID)
	if err != nil {
		return err
	}
	nfo.TokenId = (hexutil.Big)(*val)

	val, err = hexutil.DecodeBig(row.Qty)
	if err != nil {
		return err
	}
	nfo.Quantity = (hexutil.Big)(*val)

	return nil
}

package types

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// NftOwnership represents a record of NFT owner identification.
type NftOwnership struct {
	Contract  common.Address `bson:"contract"`
	TokenId   hexutil.Big    `bson:"token_id"`
	Owner     common.Address `bson:"owner"`
	Amount    hexutil.Big    `bson:"amount"`
	Obtained  time.Time      `bson:"obtained"`
	Trx       common.Hash    `bson:"trx"`
	TokenName *string        `bson:"token_name"`
}

// ComputedPk generates unique identifier for the NFT owner record.
// Collision approx for p(n)=1e-12: n = sqrt(2 x 2^96 x 2^-39) = 536.870.912 documents
func (no *NftOwnership) ComputedPk() primitive.ObjectID {
	hash := sha256.New()
	hash.Write(no.Contract.Bytes())
	hash.Write(no.TokenId.ToInt().Bytes())
	hash.Write(no.Owner.Bytes())

	var id [12]byte
	copy(id[:], hash.Sum(nil))
	return id
}

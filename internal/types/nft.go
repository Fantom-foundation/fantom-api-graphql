package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"time"
)

// NftOwnership represents a record of NFT owner identification.
type NftOwnership struct {
	Contract common.Address `bson:"contract"`
	TokenId  hexutil.Big    `bson:"token_id"`
	Owner    common.Address `bson:"owner"`
	Amount   hexutil.Big    `bson:"amount"`
	Obtained time.Time      `bson:"obtained"`
	Trx      common.Hash    `bson:"trx"`
}

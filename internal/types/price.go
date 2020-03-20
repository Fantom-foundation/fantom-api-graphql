// Package types implements different core types of the API.
package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Price represents a conversion price between FTM tokens an a target symbol.
type Price struct {
	FromSymbol    string         `json:"FROMSYMBOL"`
	ToSymbol      string         `json:"TOSYMBOL"`
	Price         float64        `json:"PRICE"`
	Open24        float64        `json:"OPEN24HOUR"`
	High24        float64        `json:"HIGH24HOUR"`
	Low24         float64        `json:"LOW24HOUR"`
	Volume24      float64        `json:"VOLUME24HOUR"`
	Change24      float64        `json:"CHANGE24HOUR"`
	ChangePct24   float64        `json:"CHANGEPCT24HOUR"`
	TotalVolume24 float64        `json:"TOTALVOLUME24H"`
	Supply        float64        `json:"SUPPLY"`
	MarketCap     float64        `json:"MKTCAP"`
	LastUpdate    hexutil.Uint64 `json:"LASTUPDATE"`
}

// UnmarshalPrice parses the JSON-encoded price data.
func UnmarshalPrice(data []byte) (Price, error) {
	var pri Price
	err := json.Unmarshal(data, &pri)
	return pri, err
}

// Marshal returns the JSON encoding of price.
func (pri Price) Marshal() ([]byte, error) {
	return json.Marshal(pri)
}

/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"encoding/json"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"time"
)

const (
	ownPriceSymbol          = "FTM"
	priceApiAddress         = "https://min-api.cryptocompare.com/data/pricemultifull?"
	priceApiSourceSymbolVar = "fsyms="
	priceApiTargetSymbolVar = "tsyms="

	// pricePullRequestTimeout is number of seconds we wait for the price information request to finish.
	pricePullRequestTimeout = 5
)

// GasPrice pulls the current amount of WEI for single Gas.
func (p *proxy) GasPrice() (hexutil.Big, error) {
	return p.rpc.GasPrice()
}

// GasPriceExtended provides extended gas price information.
func (p *proxy) GasPriceExtended() (*types.GasPrice, error) {
	// get the current gas price
	gp, err := p.rpc.GasPrice()
	if err != nil {
		return nil, err
	}

	// calculate the gas price in Gwei units
	gWei := math.Round(float64(gp.ToInt().Int64())/float64(10000000)) / 10.0
	return &types.GasPrice{
		Fast:    gWei,
		Fastest: gWei,
		SafeLow: gWei,
		Average: gWei,
	}, nil
}

// GasPriceTicks provides a list of gas price ticks for the given time period.
func (p *proxy) GasPriceTicks(from *time.Time, to *time.Time) ([]types.GasPricePeriod, error) {
	return p.db.GasPriceTicks(from, to)
}

// GasEstimate calculates the estimated amount of Gas required to perform
// transaction described by the input params.
func (p *proxy) GasEstimate(trx *struct {
	From  *common.Address
	To    *common.Address
	Value *hexutil.Big
	Data  *string
}) (*hexutil.Uint64, error) {
	return p.rpc.GasEstimate(trx)
}

// isValidPriceSymbol checks if the requested symbol is a valid price symbol we support
func (p *proxy) isValidPriceSymbol(sym string) bool {
	// check against supported price symbols from configuration
	for _, vs := range p.cfg.DeFi.PriceSymbols {
		if strings.EqualFold(vs, sym) {
			return true
		}
	}
	return false
}

// Price returns a price information for the given target symbol.
func (p *proxy) Price(sym string) (types.Price, error) {
	// check the symbol validity
	if !p.isValidPriceSymbol(sym) {
		return types.Price{}, fmt.Errorf("unknown price symbol requested")
	}

	// inform what we do
	p.log.Infof("loading price info for symbol [%s]", sym)

	// try to use the in-memory cache
	if pri := p.cache.PullPrice(sym); pri != nil {
		// inform what we do
		p.log.Infof("price [%s] loaded from cache", sym)
		return *pri, nil
	}

	// call for the price from an external source
	pri, err := p.requestPrice(sym)
	if err != nil {
		// inform what we do
		p.log.Errorf("price [%s] not available; %s", sym, err.Error())
		return types.Price{}, err
	}

	// inform what we do
	p.log.Infof("price [%s] obtained from an API request", sym)
	return pri, nil
}

// requestPrice requests the price from an external 3rd party API
// inside a request group.
func (p *proxy) requestPrice(sym string) (types.Price, error) {
	// call for the price inside a named request group
	pri, err, _ := p.apiRequestGroup.Do(priceRequestName(sym), func() (interface{}, error) {
		return p.requestRemotePrice(sym)
	})

	// any error on the process?
	if err != nil {
		return types.Price{}, err
	}

	// return the price we have
	return pri.(types.Price), nil
}

// priceRequestName generates a name for the price pull request.
func priceRequestName(sym string) string {
	var sb strings.Builder
	sb.WriteString("price+")
	sb.WriteString(sym)
	return sb.String()
}

// getPriceApiUrl builds REST API endpoint URL for the given target symbol.
func getPriceApiUrl(sym string) string {
	// use the builder
	var sb strings.Builder

	sb.WriteString(priceApiAddress)
	sb.WriteString(priceApiSourceSymbolVar)
	sb.WriteString(ownPriceSymbol)
	sb.WriteString("&")
	sb.WriteString(priceApiTargetSymbolVar)
	sb.WriteString(sym)

	return sb.String()
}

// requestRemotePrice pulls the price for given symbol from an external API
// and ensures the result, if valid, is stored in cache for future use
func (p *proxy) requestRemotePrice(sym string) (types.Price, error) {
	// make the request tpo remote API
	pri, err := p.makePriceRequest(sym)
	if err != nil {
		return types.Price{}, err
	}

	// try to store the price in cache for future use
	err = p.cache.PushPrice(sym, &pri)
	if err != nil {
		p.log.Error(err)
	}

	// inform what we got here
	p.log.Infof("price loaded: %s -> %s = %f", pri.FromSymbol, pri.ToSymbol, pri.Price)
	return pri, nil
}

// makePriceRequest executes a request to remote API to pull the price
// and return the result from the pull.
func (p *proxy) makePriceRequest(sym string) (types.Price, error) {
	// prep the request
	req, err := http.NewRequest("GET", getPriceApiUrl(sym), nil)
	if err != nil {
		return types.Price{}, fmt.Errorf("can not create HTTP request for price API; %s", err.Error())
	}

	// do the request
	client := &http.Client{Timeout: time.Second * pricePullRequestTimeout}
	resp, err := client.Do(req)
	if err != nil {
		return types.Price{}, fmt.Errorf("can not query price API; %s", err.Error())
	}

	// don't forget to close
	defer func() {
		// log the HTTP request
		p.log.Debugf("finished HTTP request to pull [%s] price", sym)

		// close the connection
		err := resp.Body.Close()
		if err != nil {
			p.log.Errorf("error closing price API request; %s", err.Error())
		}

		// is this a panic?
		if r := recover(); r != nil {
			p.log.Errorf("error parsing price API response; %s", r)
		}
	}()

	// read the data
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.Price{}, fmt.Errorf("can not read price API response; %s", err.Error())
	}

	// we need to be able to read the data
	var data map[string]map[string]map[string]map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return types.Price{}, fmt.Errorf("can not decode price API response; %s", err.Error())
	}

	return types.Price{
		FromSymbol:    (data["RAW"][ownPriceSymbol][sym]["FROMSYMBOL"]).(string),
		ToSymbol:      (data["RAW"][ownPriceSymbol][sym]["TOSYMBOL"]).(string),
		Price:         (data["RAW"][ownPriceSymbol][sym]["PRICE"]).(float64),
		Open24:        (data["RAW"][ownPriceSymbol][sym]["OPEN24HOUR"]).(float64),
		High24:        (data["RAW"][ownPriceSymbol][sym]["HIGH24HOUR"]).(float64),
		Low24:         (data["RAW"][ownPriceSymbol][sym]["LOW24HOUR"]).(float64),
		Volume24:      (data["RAW"][ownPriceSymbol][sym]["VOLUME24HOUR"]).(float64),
		Change24:      (data["RAW"][ownPriceSymbol][sym]["CHANGE24HOUR"]).(float64),
		ChangePct24:   (data["RAW"][ownPriceSymbol][sym]["CHANGEPCT24HOUR"]).(float64),
		TotalVolume24: (data["RAW"][ownPriceSymbol][sym]["TOTALVOLUME24H"]).(float64),
		Supply:        (data["RAW"][ownPriceSymbol][sym]["SUPPLY"]).(float64),
		MarketCap:     (data["RAW"][ownPriceSymbol][sym]["MKTCAP"]).(float64),
		LastUpdate:    hexutil.Uint64(uint64((data["RAW"][ownPriceSymbol][sym]["LASTUPDATE"]).(float64))),
	}, nil
}

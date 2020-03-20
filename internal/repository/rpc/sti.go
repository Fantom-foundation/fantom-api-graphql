/*
Package rpc implements bridge to Lachesis full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Opera/Lachesis node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Lachesis RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Lachesis RPC interface with connection limited to specified endpoints.

We strongly discourage opening Lachesis RPC interface for unrestricted Internet access.
*/
package rpc

//go:generate abigen --abi ./contracts/st_info.abi --pkg rpc --type StakerInfoContract --out ./sti_bind.go

import (
	"encoding/json"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"
)

// stiRequestTimeout is number of seconds we wait for the staker information request to finish.
const stiRequestTimeout = 10

// stiContractAddress holds deployment address of the Staker info smart contract.
var stiContractAddress = common.HexToAddress("0x92ffad75b8a942d149621a39502cdd8ad1dd57b4")

// StakerInfo extracts an extended staker information from smart contact by their id.
func (ftm *FtmBridge) StakerInfo(id hexutil.Uint64) (*types.StakerInfo, error) {
	// keep track of the operation
	ftm.log.Debugf("loading staker information for staker #%d", id)

	// instantiate the contract and display its name
	contract, err := NewStakerInfoContract(stiContractAddress, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate STI contract: %v", err)
		return nil, err
	}

	// call for data
	url, err := contract.GetInfo(nil, big.NewInt(int64(id)))
	if err != nil {
		ftm.log.Errorf("failed to get the staker information: %v", err)
		return nil, err
	}

	// var url string
	if len(url) == 0 {
		ftm.log.Debugf("no information for staker #%d", id)
		return nil, nil
	}

	// try to download JSON for the info
	return ftm.downloadStakerInfo(url)
}

// downloadStakerInfo tries to download staker information from the given URL address.
func (ftm *FtmBridge) downloadStakerInfo(url string) (*types.StakerInfo, error) {
	// log what we are about to do
	ftm.log.Debugf("downloading staker info address [%s]", url)

	// make a http client
	cl := http.Client{Timeout: time.Second * stiRequestTimeout}

	// prep request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		ftm.log.Errorf("can not request given staker info; %s", err.Error())
		return nil, err
	}

	// be honest, set agent
	req.Header.Set("User-Agent", "Fantom GraphQL API Server")

	// process the request
	res, err := cl.Do(req)
	if err != nil {
		ftm.log.Errorf("can not download staker info; %s", err.Error())
		return nil, err
	}

	// read the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ftm.log.Errorf("can not read staker info response; %s", err.Error())
		return nil, err
	}

	// try to parse
	var info types.StakerInfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		ftm.log.Errorf("invalid response for staker info; %s", err.Error())
		return nil, err
	}

	// do we have anything?
	if info.Name != nil {
		ftm.log.Debugf("found staker [%s]", *info.Name)
	}

	return &info, nil
}

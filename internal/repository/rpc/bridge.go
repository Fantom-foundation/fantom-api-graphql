/*
Package rpc implements bridge to Opera full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Opera/Opera node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Opera RPC interface for remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Opera RPC interface with connection limited to specified endpoints.

We strongly discourage opening Opera RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"context"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	etc "github.com/ethereum/go-ethereum/core/types"
	eth "github.com/ethereum/go-ethereum/ethclient"
	ftm "github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/sync/singleflight"
	"strings"
	"sync"
)

// rpcHeadProxyChannelCapacity represents the capacity of the new received blocks proxy channel.
const rpcHeadProxyChannelCapacity = 10000

// FtmBridge represents Opera RPC abstraction layer.
type FtmBridge struct {
	rpc *ftm.Client
	eth *eth.Client
	log logger.Logger
	cg  *singleflight.Group

	// fMintCfg represents the configuration of the fMint protocol
	sigConfig     *config.ServerSignature
	sfcConfig     *config.Staking
	uniswapConfig *config.DeFiUniswap

	// extended minter config
	fMintCfg fMintConfig
	fLendCfg fLendConfig

	// common contracts
	sfcAbi      *abi.ABI
	sfcContract *contracts.SfcContract
	sfcShards   sfcShards

	// received blocks proxy
	wg       *sync.WaitGroup
	sigClose chan bool
	headers  chan *etc.Header
}

// New creates new Opera RPC connection bridge.
func New(cfg *config.Config, log logger.Logger) (*FtmBridge, error) {
	cli, con, err := connect(cfg, log)
	if err != nil {
		log.Criticalf("can not open connection; %s", err.Error())
		return nil, err
	}

	// build the bridge structure using the con we have
	br := &FtmBridge{
		rpc: cli,
		eth: con,
		log: log,
		cg:  new(singleflight.Group),

		// special configuration options below this line
		sigConfig:     &cfg.Signature,
		sfcConfig:     &cfg.Staking,
		uniswapConfig: &cfg.DeFi.Uniswap,
		fMintCfg: fMintConfig{
			addressProvider: cfg.DeFi.FMint.AddressProvider,
		},
		fLendCfg: fLendConfig{lendigPoolAddress: cfg.DeFi.FLend.LendingPool},

		// empty shards
		sfcShards: sfcShards{log: log, client: con, sfc: cfg.Staking.SFCContract},

		// configure block observation loop
		wg:       new(sync.WaitGroup),
		sigClose: make(chan bool, 1),
		headers:  make(chan *etc.Header, rpcHeadProxyChannelCapacity),
	}

	// inform about the local address of the API node
	log.Noticef("using signature address %s", br.sigConfig.Address.String())

	// add the bridge ref to the fMintCfg and return the instance
	br.fMintCfg.bridge = br
	br.run()
	return br, nil
}

// connect opens connections we need to communicate with the blockchain node.
func connect(cfg *config.Config, log logger.Logger) (*ftm.Client, *eth.Client, error) {
	// log what we do
	log.Debugf("connecting blockchain node at %s", cfg.Opera.ApiNodeUrl)

	// try to establish a connection
	client, err := ftm.Dial(cfg.Opera.ApiNodeUrl)
	if err != nil {
		log.Critical(err)
		return nil, nil, err
	}

	// try to establish a for smart contract interaction
	con, err := eth.Dial(cfg.Opera.ApiNodeUrl)
	if err != nil {
		log.Critical(err)
		return nil, nil, err
	}

	// log
	log.Notice("node connection open")
	return client, con, nil
}

// run starts the bridge threads required to collect blockchain data.
func (ftm *FtmBridge) run() {
	ftm.wg.Add(1)
	go ftm.observeBlocks()
}

// terminate kills the bridge threads to end the bridge gracefully.
func (ftm *FtmBridge) terminate() {
	ftm.sigClose <- true
	ftm.wg.Wait()
	ftm.log.Noticef("rpc threads terminated")
}

// Close will finish all pending operations and terminate the Opera RPC connection
func (ftm *FtmBridge) Close() {
	// terminate threads before we close connections
	ftm.terminate()

	// do we have a connection?
	if ftm.rpc != nil {
		ftm.rpc.Close()
		ftm.eth.Close()
		ftm.log.Info("blockchain connections are closed")
	}
}

// Connection returns open Opera connection.
func (ftm *FtmBridge) Connection() *ftm.Client {
	return ftm.rpc
}

// DefaultCallOpts creates a default record for call options.
func (ftm *FtmBridge) DefaultCallOpts() *bind.CallOpts {
	// get the default call opts only once if called in parallel
	co, _, _ := ftm.cg.Do("default-call-opts", func() (interface{}, error) {
		return &bind.CallOpts{
			Pending:     false,
			From:        ftm.sigConfig.Address,
			BlockNumber: nil,
			Context:     context.Background(),
		}, nil
	})
	return co.(*bind.CallOpts)
}

// SfcContract returns instance of SFC contract for interaction.
func (ftm *FtmBridge) SfcContract() *contracts.SfcContract {
	// lazy create SFC contract instance
	if nil == ftm.sfcContract {
		// instantiate the contract and display its name
		var err error
		ftm.sfcContract, err = contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
		if err != nil {
			ftm.log.Criticalf("failed to instantiate SFC contract; %s", err.Error())
			panic(err)
		}
	}
	return ftm.sfcContract
}

// SfcAbi returns a parse ABI of the AFC contract.
func (ftm *FtmBridge) SfcAbi() *abi.ABI {
	if nil == ftm.sfcAbi {
		ab, err := abi.JSON(strings.NewReader(contracts.SfcContractABI))
		if err != nil {
			ftm.log.Criticalf("failed to parse SFC contract ABI; %s", err.Error())
			panic(err)
		}
		ftm.sfcAbi = &ab
	}
	return ftm.sfcAbi
}

// ObservedBlockProxy provides a channel fed with new headers observed
// by the connected blockchain node.
func (ftm *FtmBridge) ObservedBlockProxy() chan *etc.Header {
	return ftm.headers
}

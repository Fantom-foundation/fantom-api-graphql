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

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/ethclient"
	ftm "github.com/ethereum/go-ethereum/rpc"
)

// FtmBridge represents Lachesis RPC abstraction layer.
type FtmBridge struct {
	rpc *ftm.Client
	eth *eth.Client
	log logger.Logger

	// fMintCfg represents the configuration of the fMint protocol
	fMintCfg fMintConfig

	// uniswap configuration elements used by the bridge
	uniswapCore   common.Address
	uniswapRouter common.Address
}

// New creates new Lachesis RPC connection bridge.
func New(cfg *config.Config, log logger.Logger) (*FtmBridge, error) {
	// try to establish a connection
	client, err := ftm.Dial(cfg.LachesisUrl)
	if err != nil {
		log.Critical(err)
		return nil, err
	}

	// log
	log.Notice("full node rpc connection established")

	// try to establish a for smart contract interaction
	con, err := eth.Dial(cfg.LachesisUrl)
	if err != nil {
		log.Critical(err)
		return nil, err
	}

	// log
	log.Notice("smart contact connection established")

	// return the Bridge
	br := &FtmBridge{
		rpc: client,
		eth: con,
		log: log,

		// special configuration options below this line
		fMintCfg: fMintConfig{
			addressProvider: common.HexToAddress(cfg.DefiFMintAddressProvider),
		},

		// uniswap config
		uniswapCore:   common.HexToAddress(cfg.DefiUniswapCore),
		uniswapRouter: common.HexToAddress(cfg.DefiUniswapRouter),
	}

	// add the bridge ref to the fMintCfg and return the instance
	br.fMintCfg.bridge = br
	return br, nil
}

// Close will finish all pending operations and terminate the Lachesis RPC connection
func (ftm *FtmBridge) Close() {
	// do we have a connection?
	if ftm.rpc != nil {
		ftm.rpc.Close()
		ftm.eth.Close()
		ftm.log.Info("blockchain connections are closed")
	}
}

// Connection returns open Opera/Lachesis connection.
func (ftm *FtmBridge) Connection() *ftm.Client {
	return ftm.rpc
}

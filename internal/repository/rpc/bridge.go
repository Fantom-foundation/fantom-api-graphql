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
	"context"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"golang.org/x/sync/singleflight"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	eth "github.com/ethereum/go-ethereum/ethclient"
	ftm "github.com/ethereum/go-ethereum/rpc"
)

// FtmBridge represents Lachesis RPC abstraction layer.
type FtmBridge struct {
	rpc *ftm.Client
	eth *eth.Client
	log logger.Logger

	// we need a Group to use single flight to control price pulls
	reqGroup *singleflight.Group

	// fMintCfg represents the configuration of the fMint protocol
	sigConfig     *config.ServerSignature
	sfcConfig     *config.Staking
	uniswapConfig *config.DeFiUniswap

	// extended minter config
	fMintCfg fMintConfig
	fLendCfg fLendConfig
}

// New creates new Lachesis RPC connection bridge.
func New(cfg *config.Config, log logger.Logger) (*FtmBridge, error) {
	// log what we do
	log.Debugf("connecting block chain node at %s", cfg.Lachesis.Url)

	// try to establish a connection
	client, err := ftm.Dial(cfg.Lachesis.Url)
	if err != nil {
		log.Critical(err)
		return nil, err
	}

	// log
	log.Notice("block chain node online")

	// try to establish a for smart contract interaction
	con, err := eth.Dial(cfg.Lachesis.Url)
	if err != nil {
		log.Critical(err)
		return nil, err
	}

	// log
	log.Notice("smart contact connection open")

	// return the Bridge
	br := &FtmBridge{
		rpc: client,
		eth: con,
		log: log,

		// make the group
		reqGroup: new(singleflight.Group),

		// special configuration options below this line
		sigConfig:     &cfg.MySignature,
		sfcConfig:     &cfg.Staking,
		uniswapConfig: &cfg.DeFi.Uniswap,
		fMintCfg: fMintConfig{
			addressProvider: cfg.DeFi.FMint.AddressProvider,
		},
		fLendCfg: fLendConfig{lendigPoolAddress: cfg.DeFi.FLend.LendingPool},
	}

	// inform about the local address of the API node
	log.Noticef("using signature address %s", br.sigConfig.Address.String())

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

// DefaultCallOpts creates a default record for call options.
func (ftm *FtmBridge) DefaultCallOpts() *bind.CallOpts {
	return &bind.CallOpts{
		Pending:     false,
		From:        ftm.sigConfig.Address,
		BlockNumber: ftm.MustBlockHeight(),
		Context:     context.Background(),
	}
}

/*
Rpc package implements bridge to Lachesis full node API interface.

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
	ftm "github.com/ethereum/go-ethereum/rpc"
)

// Bridge represents Lachesis RPC abstraction layer.
type FtmBridge struct {
	rpc *ftm.Client
	log logger.Logger
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
	log.Notice("full node connection established")

	// return the Bridge
	return &FtmBridge{
		rpc: client,
		log: log,
	}, nil
}

// Close will finish all pending operations and terminate the Lachesis RPC connection
func (ftm *FtmBridge) Close() {
	// do we have a connection?
	if ftm.rpc != nil {
		ftm.rpc.Close()
		ftm.log.Info("full node connection is closed")
	}
}

// Client returns open Opera/Lachesis connection.
func (ftm *FtmBridge) Connection() *ftm.Client {
	return ftm.rpc
}

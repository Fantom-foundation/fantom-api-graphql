package main

import (
	"context"
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
)

// connector represents the node gateway connection
type connector struct {
	client   *ethclient.Client
	gov      *contracts.Governance
	key      *bind.TransactOpts
	batchLen uint64
}

// connect opens connection to the Opera RPC interface
func connect(cfg *gvConf) (*ethclient.Client, error) {
	// log what we do and try to dial the node
	fmt.Printf("Opening RPC connection to %s\n", cfg.nodeRpcUrl)

	// open the node connection
	con, err := ethclient.Dial(cfg.nodeRpcUrl)
	if err != nil {
		fmt.Printf("Can not connect the node at %s!\n%s\n", cfg.nodeRpcUrl, err.Error())
		return nil, err
	}

	// get network ID so we can assume we are talking to something we expect
	nid, err := con.NetworkID(context.Background())
	if err != nil {
		fmt.Printf("Can not detect the network at %s!\n%s\n", cfg.nodeRpcUrl, err.Error())
		return nil, err
	}

	// show the result
	fmt.Printf("Network #%d contacted.\n", nid.Uint64())
	return con, nil
}

// governance connects with the Governance contract to allow interactions
func governance(cfg *gvConf, con *ethclient.Client) (*contracts.Governance, error) {
	// connect the contract
	fmt.Printf("Accessing Governance contract %s\n", cfg.govAddress)
	return contracts.NewGovernance(common.HexToAddress(cfg.govAddress), con)
}

// unlockKey tries to unlock the given key file and prepares the transaction
// signature provider so we can interact with the Governance contract through
// signed transactions as needed.
func unlockKey(cfg *gvConf) (*bind.TransactOpts, error) {
	// read the key store
	f, err := os.Open(cfg.sigKeyFile)
	if err != nil {
		fmt.Printf("Can not open key store at %s!\n%s\n", cfg.sigKeyFile, err.Error())
		return nil, err
	}

	// ensure proper cleanup
	defer func() {
		// make sure to close the opened key store file
		if err := f.Close(); err != nil {
			fmt.Printf("Error closing key store file.\n%s\n", err.Error())
		}
	}()

	// create the transactor
	tr, err := bind.NewTransactor(f, cfg.sigPasswd)
	if err != nil {
		fmt.Printf("Unlock password rejected!\n%s\n", err.Error())
		return nil, err
	}

	// force the gas limit
	tr.GasLimit = 4000000

	// log transaction signer and return the opts
	fmt.Printf("Ulocked account %s.\n", tr.From.String())
	return tr, nil
}

// attach opens the node connection and attaches the Governance contract
// so we can process tasks.
func attach(cfg *gvConf) (*connector, error) {
	// get the Node connection
	con, err := connect(cfg)
	if err != nil {
		fmt.Printf("Opera node not available!\n%s\n", err.Error())
		return nil, err
	}

	// get the governance contract interface
	gov, err := governance(cfg, con)
	if err != nil {
		fmt.Printf("Can not access Governance contract!\n%s\n", err.Error())
		return nil, err
	}

	// unlock the key
	key, err := unlockKey(cfg)
	if err != nil {
		fmt.Printf("Can not unlock the key!\n%s\n", err.Error())
		return nil, err
	}

	// return the connector
	return &connector{
		client:   con,
		gov:      gov,
		key:      key,
		batchLen: cfg.batchLength,
	}, nil
}

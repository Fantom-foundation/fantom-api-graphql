package gvclear

import (
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fmt"
	eth "github.com/ethereum/go-ethereum/ethclient"
)

// connect opens connection to the Opera RPC interface
func connect(cfg *gvConf) (*eth.Client, error) {
	// log what we do and try to dial the node
	fmt.Printf("Opening RPC connection to %s", cfg.nodeRpcUrl)
	return eth.Dial(cfg.nodeRpcUrl)
}

// governance connects with the Governance contract to allow interactions
func governance(cfg *gvConf, con *eth.Client) (*contracts.Governance, err) {
	// connect the contract
	fmt.Printf("Accessing Governance contract %s", cfg.govAddress)
	return contracts.NewGovernance(common.HexToAddress(cfg.govAddress), con)
}

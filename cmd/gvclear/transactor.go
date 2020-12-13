package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// defaultCallOpts creates a default record for call options.
func defaultCallOpts(adr *common.Address, con *ethclient.Client) *bind.CallOpts {
	return &bind.CallOpts{
		Pending:     false,
		From:        *adr,
		BlockNumber: mustBlockHeight(con),
		Context:     context.Background(),
	}
}

// mustBlockHeight returns the current block height
// of the block chain. It returns nil if the block height can not be pulled.
func mustBlockHeight(con *ethclient.Client) *big.Int {
	// get the current block number
	blk, err := con.BlockNumber(context.Background())
	if err != nil {
		fmt.Printf("Can not get the current block height!\n%s\n", err.Error())
		return nil
	}

	return new(big.Int).SetUint64(blk)
}

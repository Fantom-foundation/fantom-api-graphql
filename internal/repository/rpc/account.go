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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// AccountBalance reads balance of account from Lachesis node.
func (b *Bridge) AccountBalance(addr *common.Address) (*big.Int, error) {
	// use RPC to make the call
	var balance string
	err := b.rpc.Call(&balance, "ftm_getBalance", addr.Hex(), "latest")
	if err != nil {
		b.log.Errorf("can not get balance of account [%s]", addr.Hex())
		return nil, err
	}

	// decode the response from remote server
	val, err := hexutil.DecodeBig(balance)
	if err != nil {
		b.log.Errorf("can not decode balance of account [%s]", addr.Hex())
		return nil, err
	}

	return val, nil
}

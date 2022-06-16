/*
Package rpc implements bridge to Opera full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Opera/Opera node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Opera RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Opera RPC interface with connection limited to specified endpoints.

We strongly discourage opening Opera RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// AccountBalance reads balance of account from Opera node.
func (ftm *FtmBridge) AccountBalance(addr *common.Address) (*hexutil.Big, error) {
	// use RPC to make the call
	var balance string
	err := ftm.rpc.Call(&balance, "ftm_getBalance", addr.Hex(), "latest")
	if err != nil {
		ftm.log.Errorf("can not get balance of account [%s]", addr.Hex())
		return nil, err
	}

	// decode the response from remote server
	val, err := hexutil.DecodeBig(balance)
	if err != nil {
		ftm.log.Errorf("can not decode balance of account [%s]", addr.Hex())
		return nil, err
	}

	return (*hexutil.Big)(val), nil
}

// AccountNonce returns the total number of transaction of account from Opera node.
func (ftm *FtmBridge) AccountNonce(addr *common.Address) (*hexutil.Uint64, error) {
	var nonce hexutil.Uint64
	err := ftm.rpc.Call(&nonce, "ftm_getTransactionCount", addr.Hex(), "latest")
	if err != nil {
		ftm.log.Errorf("can not get number of transaction of account [%s]", addr.Hex())
		return nil, err
	}
	return &nonce, nil
}

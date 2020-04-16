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

import "github.com/ethereum/go-ethereum/common/hexutil"

// GasPrice resolves the current amount of WEI for single Gas.
func (ftm *FtmBridge) GasPrice() (hexutil.Uint64, error) {
	// keep track of the operation
	ftm.log.Debugf("checking current gas price")

	// call for data
	var price hexutil.Big
	err := ftm.rpc.Call(&price, "ftm_gasPrice")
	if err != nil {
		ftm.log.Error("current gas price could not be obtained")
		return hexutil.Uint64(0), err
	}

	// if the price safely within the range
	if !price.ToInt().IsUint64() {
		ftm.log.Error("current gas price is too high and can not be extracted")
		return hexutil.Uint64(0), err
	}

	// inform and return
	ftm.log.Debugf("current gas price is %d", uint64(price.ToInt().Uint64()))
	return hexutil.Uint64(price.ToInt().Uint64()), nil
}

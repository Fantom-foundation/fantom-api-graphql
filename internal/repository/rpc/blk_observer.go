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
	"github.com/ethereum/go-ethereum"
	"time"
)

// ftmHeadsObserverSubscribeTick represents the time between subscription attempts.
const ftmHeadsObserverSubscribeTick = 30 * time.Second

// observeBlocks collects new blocks from the blockchain network
// and posts them into the proxy channel for processing.
func (ftm *FtmBridge) observeBlocks() {
	var sub ethereum.Subscription
	defer func() {
		if sub != nil {
			sub.Unsubscribe()
		}
		ftm.log.Noticef("block observer done")
		ftm.wg.Done()
	}()

	sub = ftm.blockSubscription()
	for {
		// re-subscribe if the subscription ref is not valid
		if sub == nil {
			tm := time.NewTimer(ftmHeadsObserverSubscribeTick)
			select {
			case <-ftm.sigClose:
				return
			case <-tm.C:
				sub = ftm.blockSubscription()
				continue
			}
		}

		// use the subscriptions
		select {
		case <-ftm.sigClose:
			return
		case err := <-sub.Err():
			ftm.log.Errorf("block subscription failed; %s", err.Error())
			sub = nil
		}
	}
}

// blockSubscription provides a subscription for new blocks received
// by the connected blockchain node.
func (ftm *FtmBridge) blockSubscription() ethereum.Subscription {
	sub, err := ftm.rpc.EthSubscribe(context.Background(), ftm.headers, "newHeads")
	if err != nil {
		ftm.log.Criticalf("can not observe new blocks; %s", err.Error())
		return nil
	}
	return sub
}

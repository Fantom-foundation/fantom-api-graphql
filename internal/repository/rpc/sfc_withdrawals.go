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
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// withdrawnRequest represents a finalized partial withdraw
// request in the blockchain.
// We do not need any of the addresses here since we pair
// corresponding requests just by the request id.  We always use
// the same address filter for both requests list and finalization.
type withdrawnRequest struct {
	RequestId   big.Int
	BlockNumber uint64
	Penalty     big.Int
}

// WithdrawRequests extracts a list of partial withdraw requests
// for the given address.
func (ftm *FtmBridge) WithdrawRequests(addr *common.Address) ([]*types.WithdrawRequest, error) {
	// we need to have the address to continue
	if addr == nil {
		ftm.log.Error("can not pull withdraw requests for empty address")
		return nil, fmt.Errorf("withdraw requests address filter not defined")
	}

	// prepare to interact with the SFC contract
	contract, err := NewSfcContract(sfcContractAddress, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %v", err)
		return nil, err
	}

	// get a list of finalized requests
	fin, err := ftm.withdrawnByRequest(contract, *addr)
	if err != nil {
		ftm.log.Error("can not pull finalized withdraw requests; %s", err.Error())
		return nil, err
	}

	// get a list of requests
	return ftm.createdWithdrawRequests(contract, *addr, fin)
}

// withdrawnByRequest pulls list of finalized withdraw requests
// from the SFC contract events using filter iterator.
func (ftm *FtmBridge) withdrawnByRequest(sfc *SfcContract, addr common.Address) ([]withdrawnRequest, error) {
	// create event iterator for the finalized withdraw requests
	it, err := sfc.FilterPartialWithdrawnByRequest(nil, []common.Address{addr}, nil, nil)
	if err != nil {
		ftm.log.Errorf("failed to get finalized withdraws iterator; %s", err.Error())
		return nil, err
	}

	// make sure to close the iterator on leaving
	defer func() {
		err := it.Close()
		if err != nil {
			ftm.log.Errorf("failed to close finalized withdraws iterator; %s", err.Error())
		}
	}()

	// make the container
	list := make([]withdrawnRequest, 0)

	// loop through the iterator
	for it.Next() {
		// make sure this is a valid record
		if it.Event.WrID == nil || it.Event.Penalty == nil {
			ftm.log.Error("invalid partial withdraw finalization record found")
			continue
		}

		// populate the local struct with data we need
		fr := withdrawnRequest{
			RequestId:   *it.Event.WrID,
			BlockNumber: it.Event.Raw.BlockNumber,
			Penalty:     *it.Event.Penalty,
		}

		// add to the list
		list = append(list, fr)
	}

	// detect possible error in traversing the iterator
	if err := it.Error(); err != nil {
		ftm.log.Errorf("failed to pull from finalized withdraws iterator; %s", err.Error())
		return nil, err
	}

	return list, nil
}

// createdWithdrawRequests pulls list of created withdraw requests
// from the SFC contract events using filter iterator.
func (ftm *FtmBridge) createdWithdrawRequests(sfc *SfcContract, addr common.Address, fin []withdrawnRequest) ([]*types.WithdrawRequest, error) {
	// create event iterator for the created withdraw requests
	it, err := sfc.FilterCreatedWithdrawRequest(nil, []common.Address{addr}, nil, nil)
	if err != nil {
		ftm.log.Errorf("failed to get created withdraw requests iterator; %s", err.Error())
		return nil, err
	}

	// make sure to close the iterator on leaving
	defer func() {
		err := it.Close()
		if err != nil {
			ftm.log.Errorf("failed to close created withdraw requests iterator; %s", err.Error())
		}
	}()

	// make the container
	list := make([]*types.WithdrawRequest, 0)

	// loop through the iterator
	for it.Next() {
		// make sure this is a valid record
		if it.Event.WrID == nil || it.Event.Amount == nil || it.Event.StakerID == nil {
			ftm.log.Error("invalid created withdraw request record found")
			continue
		}

		// make the structure of the withdraw request
		// the finalization details are nil by default
		wr := types.WithdrawRequest{
			Address:            it.Event.Auth,
			Receiver:           it.Event.Receiver,
			StakerID:           hexutil.Uint64(it.Event.StakerID.Uint64()),
			WithdrawRequestID:  hexutil.Big(*it.Event.WrID),
			IsDelegation:       it.Event.Delegation,
			Amount:             hexutil.Big(*it.Event.Amount),
			RequestBlockNumber: hexutil.Uint64(it.Event.Raw.BlockNumber),
		}

		// try to extend the request with finalized details if available
		pairWithdrawRequestWithFinalized(&wr, fin)

		// add to output list
		list = append(list, &wr)
	}

	// detect possible error in traversing the iterator
	if err := it.Error(); err != nil {
		ftm.log.Errorf("failed to pull from created withdraw requests iterator; %s", err.Error())
		return nil, err
	}

	return list, nil
}

// pairWithdrawRequestWithFinalized tries to pair the withdraw request with one
// of the finalized withdrawals to populate finished requests details.
func pairWithdrawRequestWithFinalized(wr *types.WithdrawRequest, fin []withdrawnRequest) {
	// loop all finalized requests available to find
	// the corresponding one if exists
	for _, fr := range fin {
		// compare request id which is the unique request identifier
		// for the given address
		if fr.RequestId.Cmp(wr.WithdrawRequestID.ToInt()) == 0 {
			// what block is this finalized withdrawal registered in?
			blk := hexutil.Uint64(fr.BlockNumber)

			// what penalty has been applied
			pen := hexutil.Big(fr.Penalty)

			// update the request
			wr.WithdrawBlockNumber = &blk
			wr.WithdrawPenalty = &pen

			// we are done here
			break
		}
	}
}

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

// finalizedDeactivatedDelegation represents a finalized deactivated delegation
// request in the blockchain.
type finalizedDeactivatedDelegation struct {
	StakerID    uint64
	BlockNumber uint64
	Penalty     big.Int
}

// DeactivatedDelegation extracts a list of deactivated delegation requests
// for the given address.
func (ftm *FtmBridge) DeactivatedDelegation(addr *common.Address, stakerId hexutil.Uint64) ([]*types.DeactivatedDelegation, error) {
	// get a list of deactivation
	return ftm.deactivatedDelegationList(addr, new(big.Int).SetUint64(uint64(stakerId)))
}

// PendingDeactivation finds pending full withdraw request for the given address and staker id.
// If there is a pending deactivation, all the remaining amount delegated is going to be withdrawn.
func (ftm *FtmBridge) PendingDeactivation(addr *common.Address, staker *big.Int) (*types.DeactivatedDelegation, error) {
	// get withdraw requests list
	list, err := ftm.deactivatedDelegationList(addr, staker)
	if err != nil {
		return nil, err
	}

	// loop over the list of requests and add non-finished
	for _, req := range list {
		if req.WithdrawBlockNumber == nil {
			return req, nil
		}
	}

	return nil, nil
}

// deactivatedDelegationList extracts a list of deactivated delegation requests
// for the given address and staker.
func (ftm *FtmBridge) deactivatedDelegationList(addr *common.Address, staker *big.Int) ([]*types.DeactivatedDelegation, error) {
	// we need to have the address to continue
	if addr == nil {
		ftm.log.Error("can not pull deactivated delegation requests for empty address")
		return nil, fmt.Errorf("deactivated delegation requests address filter not defined")
	}

	// prepare to interact with the SFC contract
	contract, err := NewSfcContract(sfcContractAddress, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %v", err)
		return nil, err
	}

	// get a list of finalized requests
	fin, err := ftm.finalizedDeactivatedDelegation(contract, *addr, staker)
	if err != nil {
		ftm.log.Error("can not pull finalized deactivated delegation requests; %s", err.Error())
		return nil, err
	}

	// get a list of requests
	return ftm.createdDeactivatedDelegation(contract, *addr, staker, fin)
}

// createdDeactivatedDelegation pulls list of created deactivated delegation requests
// from the SFC contract events using filter iterator.
func (ftm *FtmBridge) createdDeactivatedDelegation(sfc *SfcContract, addr common.Address, staker *big.Int, fin []finalizedDeactivatedDelegation) ([]*types.DeactivatedDelegation, error) {
	var it *SfcContractDeactivatedDelegationIterator
	var err error

	// create event iterator for the created deactivated delegation requests
	if staker == nil {
		// we don't care about the staker here
		it, err = sfc.FilterDeactivatedDelegation(nil, []common.Address{addr}, nil)
	} else {
		// we don't care about the staker here
		it, err = sfc.FilterDeactivatedDelegation(nil, []common.Address{addr}, []*big.Int{staker})
	}

	// check for errors before we pull the list
	if err != nil {
		ftm.log.Errorf("failed to get created deactivated delegation requests iterator; %s", err.Error())
		return nil, err
	}

	// make sure to close the iterator on leaving
	defer func() {
		err := it.Close()
		if err != nil {
			ftm.log.Errorf("failed to close created deactivated delegation requests iterator; %s", err.Error())
		}
	}()

	// pull the list
	return ftm.createdDeactivatedDelegationList(it, fin)
}

// createdDeactivatedDelegationList pulls list of created deactivated delegation requests
// from the given iterator.
func (ftm *FtmBridge) createdDeactivatedDelegationList(it *SfcContractDeactivatedDelegationIterator, fin []finalizedDeactivatedDelegation) ([]*types.DeactivatedDelegation, error) {
	// make the container
	list := make([]*types.DeactivatedDelegation, 0)

	// loop through the iterator
	for it.Next() {
		// make sure this is a valid record
		if it.Event.StakerID == nil {
			ftm.log.Error("invalid created deactivated delegation request record found")
			continue
		}

		// make the structure of the withdraw request
		// the finalization details are nil by default
		dd := types.DeactivatedDelegation{
			Address:            it.Event.Delegator,
			StakerID:           hexutil.Uint64(it.Event.StakerID.Uint64()),
			RequestBlockNumber: hexutil.Uint64(it.Event.Raw.BlockNumber),
		}

		// try to extend the request with finalized details if available
		pairDeactivatedDelegationWithFinalized(&dd, fin)

		// add to output list
		list = append(list, &dd)
	}

	// detect possible error in traversing the iterator
	if err := it.Error(); err != nil {
		ftm.log.Errorf("failed to pull from created deactivated delegation requests iterator; %s", err.Error())
		return nil, err
	}

	return list, nil
}

// finalizedDeactivatedDelegation extracts a list of finalized deactivated delegation requests
// for the given address.
func (ftm *FtmBridge) finalizedDeactivatedDelegation(sfc *SfcContract, addr common.Address, staker *big.Int) ([]finalizedDeactivatedDelegation, error) {
	var it *SfcContractWithdrawnDelegationIterator
	var err error

	// create event iterator for the finalized deactivated delegation requests
	if staker == nil {
		// forget about the staker and use just the delegator address
		it, err = sfc.FilterWithdrawnDelegation(nil, []common.Address{addr}, nil)
	} else {
		// make sure to limit the staker along with the delegator address
		it, err = sfc.FilterWithdrawnDelegation(nil, []common.Address{addr}, []*big.Int{staker})
	}

	// check errors before we pull the actual list from the iterator
	if err != nil {
		ftm.log.Errorf("failed to get finalized deactivated delegation iterator; %s", err.Error())
		return nil, err
	}

	// make sure to close the iterator on leaving
	defer func() {
		err := it.Close()
		if err != nil {
			ftm.log.Errorf("failed to close finalized deactivated delegation iterator; %s", err.Error())
		}
	}()

	// pull the list
	return ftm.finalizedDeactivatedDelegationList(it)
}

// finalizedDeactivatedDelegation extracts a list of finalized deactivated delegation requests
// for the given address.
func (ftm *FtmBridge) finalizedDeactivatedDelegationList(it *SfcContractWithdrawnDelegationIterator) ([]finalizedDeactivatedDelegation, error) {
	// make the container
	list := make([]finalizedDeactivatedDelegation, 0)

	// loop through the iterator
	for it.Next() {
		// make sure this is a valid record
		if it.Event.ToStakerID == nil || it.Event.Penalty == nil {
			ftm.log.Error("invalid finalization deactivated delegation record found")
			continue
		}

		// populate the local struct with data we need
		fr := finalizedDeactivatedDelegation{
			StakerID:    it.Event.ToStakerID.Uint64(),
			BlockNumber: it.Event.Raw.BlockNumber,
			Penalty:     *it.Event.Penalty,
		}

		// add to the list
		list = append(list, fr)
	}

	// detect possible error in traversing the iterator
	if err := it.Error(); err != nil {
		ftm.log.Errorf("failed to pull from finalized deactivated delegation iterator; %s", err.Error())
		return nil, err
	}

	return list, nil
}

// pairDeactivatedDelegationWithFinalized tries to pair the deactivated delegation
// request with one of the finalized deactivated delegation
// to populate finished requests details.
func pairDeactivatedDelegationWithFinalized(dd *types.DeactivatedDelegation, fin []finalizedDeactivatedDelegation) {
	// loop all finalized deactivated delegation available to find
	// the corresponding one if exists
	for _, fr := range fin {
		// we check for the same staker id and block number above opening request
		if uint64(dd.StakerID) == fr.StakerID && uint64(dd.RequestBlockNumber) < fr.BlockNumber {
			// what block is this finalized deactivated delegation registered in?
			blk := hexutil.Uint64(fr.BlockNumber)

			// what penalty has been applied
			pen := hexutil.Big(fr.Penalty)

			// update the request
			dd.WithdrawBlockNumber = &blk
			dd.WithdrawPenalty = &pen

			// we are done here
			break
		}
	}
}

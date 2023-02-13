/*
Package rpc implements bridge to Opera full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Opera/Opera node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Opera RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Opera RPC interface with connection limited to specified endpoints.

We strongly discourage opening Opera RPC interface for unrestricted Internet access.
---
curl -H "Content-Type: application/json" -X POST --data '{"method":"eth_call","params":[{"from":"0x0000000000000000000000000000000000000000","to":"0xFC00FACE00000000000000000000000000000000","gas":"0x1ec77","gasPrice":"0x92e59e9300","value":"0x0","data":"0xc5f956af"},"latest"],"id":1,"jsonrpc":"2.0"}' https://rpcapi-tracing.fantom.network/
*/
package rpc

import (
	"context"
	"fantom-api-graphql/internal/logger"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
	"sync"
	"time"
)

// sfcShards holds a cached values for SFC contract shards.
type sfcShards struct {
	log       logger.Logger
	client    *ethclient.Client
	sfc       common.Address
	update    time.Time
	lock      sync.Mutex
	constants common.Address
	treasury  common.Address
}

var (
	// callSigConstAddress represents function signature for constants contract address
	callSigConstAddress = []byte{0xd4, 0x6f, 0xa5, 0x18}

	// callSigTreasuryAddress represents function signature for treasury address (treasuryAddress() 0xc5f956af)
	callSigTreasuryAddress = []byte{0xc5, 0xf9, 0x56, 0xaf}

	// callSigMinSelfStake represents function signature for minSelfStake() returning uint256
	callSigMinSelfStake = []byte{0xc5, 0xf5, 0x30, 0xaf}

	// callSigMaxDelegatedRatio represents function signature for maxDelegatedRatio() returning uint256
	callSigMaxDelegatedRatio = []byte{0x22, 0x65, 0xf2, 0x84}

	// callSigValidatorCommission represents function signature for validatorCommission() returning uint256
	callSigValidatorCommission = []byte{0xa7, 0x78, 0x65, 0x15}

	// callSigBurntFeeShare represents function signature for burntFeeShare() returning uint256
	callSigBurntFeeShare = []byte{0xc7, 0x4d, 0xd6, 0x21}

	// callSigTreasuryFeeShare represents function signature for treasuryFeeShare() returning uint256
	callSigTreasuryFeeShare = []byte{0x94, 0xc3, 0xe9, 0x14}

	// callSigUnlockedRewardRatio represents function signature for unlockedRewardRatio() returning uint256
	callSigUnlockedRewardRatio = []byte{0x5e, 0x23, 0x08, 0xd2}

	// callSigMinLockupDuration represents function signature for minLockupDuration() returning uint256
	callSigMinLockupDuration = []byte{0x0d, 0x7b, 0x26, 0x09}

	// callSigMaxLockupDuration represents function signature for maxLockupDuration() returning uint256
	callSigMaxLockupDuration = []byte{0x0d, 0x49, 0x55, 0xe3}

	// callSigWithdrawalPeriodEpochs represents function signature for withdrawalPeriodEpochs() returning uint256
	callSigWithdrawalPeriodEpochs = []byte{0x65, 0x0a, 0xcd, 0x66}

	// callSigWithdrawalPeriodTime represents function signature for withdrawalPeriodTime() returning uint256
	callSigWithdrawalPeriodTime = []byte{0xb8, 0x2b, 0x84, 0x27}

	// callSigBaseRewardPerSecond represents function signature for baseRewardPerSecond() returning uint256
	callSigBaseRewardPerSecond = []byte{0xd9, 0xa7, 0xc1, 0xf9}

	// callSigTargetGasPowerPerSecond represents function signature for targetGasPowerPerSecond() returning uint256
	callSigTargetGasPowerPerSecond = []byte{0x3a, 0x3e, 0xf6, 0x6c}
)

// sfcLoadShards loads shards addresses from SFC contract.
func (sha *sfcShards) assertShards() {
	// do we have up-to-date shards addresses
	if sha.update.After(time.Now()) {
		return
	}

	// new timeout
	sha.update = time.Now().Add(1 * time.Hour)
	sha.loadConstantsAddress()
}

// callForAddress calls given contract and returns decoded address from the reponse.
func callForAddress(contract common.Address, call []byte, client *ethclient.Client, log logger.Logger) (common.Address, error) {
	data, err := client.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   &contract,
		Data: call,
	}, nil)
	if err != nil {
		log.Criticalf("failed to get address; %s", err.Error())
		return common.Address{}, err
	}

	// the address is at the tail of 32 bytes long ABI response; we skip bytes 0..11 and use the rest
	adr := common.Address{}
	adr.SetBytes(data[12:])
	return adr, nil
}

// sfcLoadConstantsAddress loads SFC constants shard address from SFC contract.
// function constsAddress() external view returns (address)
func (sha *sfcShards) loadConstantsAddress() {
	var err error

	// get constants ref
	sha.constants, err = callForAddress(sha.sfc, callSigConstAddress, sha.client, sha.log)
	if err != nil {
		sha.log.Criticalf("failed to get SFC constants address; %s", err.Error())
		return
	}

	// get treasury vault ref
	sha.treasury, err = callForAddress(sha.sfc, callSigTreasuryAddress, sha.client, sha.log)
	if err != nil {
		sha.log.Criticalf("failed to get SFC treasury address; %s", err.Error())
		return
	}

	sha.log.Noticef("SFC constants address is %s", sha.constants.String())
	sha.log.Noticef("SFC treasury address is %s", sha.treasury.String())
}

// sfcConstants provides address of SFC constants shard.
func (sha *sfcShards) sfcConstants() common.Address {
	// protect config
	sha.lock.Lock()
	defer sha.lock.Unlock()

	// make sure we have shards
	sha.assertShards()
	return sha.constants
}

// constantBySignature pulls a single uint256 value from the SFC ConstantsManager using given call signature.
func (sha *sfcShards) constantBySignature(sig []byte) (*big.Int, error) {
	// get the constants shard address
	cad := sha.sfcConstants()

	// call the contract to get this value
	data, err := sha.client.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   &cad,
		Data: sig,
	}, nil)
	if err != nil {
		sha.log.Errorf("SFC constant %s failed; %s", hexutils.BytesToHex(sig), err.Error())
		return nil, err
	}

	val := new(big.Int).SetBytes(data)
	sha.log.Debugf("SFC constant %s = %s", hexutils.BytesToHex(sig), (*hexutil.Big)(val).String())

	return val, nil
}

// minSelfStake provides minimum amount of stake for a validator in WEI, i.e. 500000 FTM.
func (sha *sfcShards) minSelfStake() (*big.Int, error) {
	return sha.constantBySignature(callSigMinSelfStake)
}

// maxDelegatedRatio provides maximum ratio of delegations a validator can have, i.e. 15 times of self-stake.
func (sha *sfcShards) maxDelegatedRatio() (*big.Int, error) {
	return sha.constantBySignature(callSigMaxDelegatedRatio)
}

// validatorCommission provides the commission fee in percentage a validator will get from a delegation, i.e. 15%.
func (sha *sfcShards) validatorCommission() (*big.Int, error) {
	return sha.constantBySignature(callSigValidatorCommission)
}

// burntFeeShare provides the percentage of fees to burn, i.e. 20%.
func (sha *sfcShards) burntFeeShare() (*big.Int, error) {
	return sha.constantBySignature(callSigBurntFeeShare)
}

// treasuryFeeShare provides the percentage of fees to transfer to treasury address, i.e. 10%.
func (sha *sfcShards) treasuryFeeShare() (*big.Int, error) {
	return sha.constantBySignature(callSigTreasuryFeeShare)
}

// unlockedRewardRatio provides the ratio of the reward rate at base rate (no lock), i.e. 30%.
func (sha *sfcShards) unlockedRewardRatio() (*big.Int, error) {
	return sha.constantBySignature(callSigUnlockedRewardRatio)
}

// minLockupDuration provides the minimum duration of a stake/delegation lockup, i.e. 2 weeks.
func (sha *sfcShards) minLockupDuration() (*big.Int, error) {
	return sha.constantBySignature(callSigMinLockupDuration)
}

// maxLockupDuration provides the maximum duration of a stake/delegation lockup, i.e. 1 year.
func (sha *sfcShards) maxLockupDuration() (*big.Int, error) {
	return sha.constantBySignature(callSigMaxLockupDuration)
}

// withdrawalPeriodEpochs provides the number of epochs that undelegated stake is locked for.
func (sha *sfcShards) withdrawalPeriodEpochs() (*big.Int, error) {
	return sha.constantBySignature(callSigWithdrawalPeriodEpochs)
}

// withdrawalPeriodTime provides the number of seconds that undelegated stake is locked for.
func (sha *sfcShards) withdrawalPeriodTime() (*big.Int, error) {
	return sha.constantBySignature(callSigWithdrawalPeriodTime)
}

// baseRewardPerSecond provides the base value for rewards distributed as a stake reward per second.
func (sha *sfcShards) baseRewardPerSecond() (*big.Int, error) {
	return sha.constantBySignature(callSigBaseRewardPerSecond)
}

// targetGasPowerPerSecond provides the base value for target network gas per second.
func (sha *sfcShards) targetGasPowerPerSecond() (*big.Int, error) {
	return sha.constantBySignature(callSigTargetGasPowerPerSecond)
}

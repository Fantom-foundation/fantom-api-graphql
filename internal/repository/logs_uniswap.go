/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// handleUniswapSwap processes Uniswap Swap event log emitted when a sender trades
// input tokens to gain output tokens, this is the basic type of trade on an Uniswap pair.
// UniswapPair::Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func handleUniswapSwap(log *types.LogRecord) {
	// sanity check for data (4 x uint256 = 4x32 bytes = 128 bytes), (1 x subject topic + 2 x address = 3 topics)
	if len(log.Data) != 128 || len(log.Topics) != 3 {
		R().Log().Criticalf("%s log invalid data length; expected 128 bytes, %d bytes given; expected 3 topics, %d given", log.TxHash.String(), len(log.Data), len(log.Topics))
		return
	}

	// get the reserve values
	in0 := new(big.Int).SetBytes(log.Data[:32])
	in1 := new(big.Int).SetBytes(log.Data[32:64])
	out0 := new(big.Int).SetBytes(log.Data[64:96])
	out1 := new(big.Int).SetBytes(log.Data[96:])
	sender := common.BytesToAddress(log.Topics[1].Bytes())
	z := new(big.Int)

	// log what we got
	R().Log().Debugf("uniswap SWAP on pair %s, block #%d, for %s, in0: %s, in1: %s, out0: %s, out1: %s",
		log.Address.String(),
		log.Block.Number,
		sender.String(),
		((*hexutil.Big)(in0)).String(),
		((*hexutil.Big)(in1)).String(),
		((*hexutil.Big)(out0)).String(),
		((*hexutil.Big)(out1)).String(),
	)

	// store the swap to repository
	err := R().UniswapAdd(&types.Swap{
		OrdIndex:    log.Trx.Uid(),
		BlockNumber: &log.Block.Number,
		Type:        types.SwapMint,
		TimeStamp:   &log.Block.TimeStamp,
		Pair:        log.Address,
		Sender:      log.Trx.From,
		Hash:        log.Trx.Hash,
		Amount0In:   in0,
		Amount0Out:  out0,
		Amount1In:   in1,
		Amount1Out:  out1,
		Reserve0:    z,
		Reserve1:    z,
	})
	if err != nil {
		R().Log().Errorf("could not store uniswap event; %s", err.Error())
	}
}

// handleUniswapMint processes Uniswap Mint event log emitted when a sender adds new liquidity
// to an Uniswap token pair to increase their share on the pair rewards.
// UniswapPair::Mint(address indexed sender, uint256 amount0, uint256 amount1)
func handleUniswapMint(log *types.LogRecord) {
	// sanity check for data (2 x uint256 = 2x32 bytes = 64 bytes), (1 x subject topic + 1 x address = 2 topics)
	if len(log.Data) != 64 || len(log.Topics) != 2 {
		R().Log().Criticalf("%s log invalid data length; expected 64 bytes, %d bytes given; expected 2 topics, %d given", log.TxHash.String(), len(log.Data), len(log.Topics))
		return
	}

	// get the reserve values
	a0 := new(big.Int).SetBytes(log.Data[:32])
	a1 := new(big.Int).SetBytes(log.Data[32:])
	sender := common.BytesToAddress(log.Topics[1].Bytes())
	z := new(big.Int)

	// log what we got
	R().Log().Debugf("uniswap MINT on pair %s, block #%d, for %s, amount0: %s, amount1: %s",
		log.Address.String(), log.Block.Number, sender.String(), ((*hexutil.Big)(a0)).String(), ((*hexutil.Big)(a1)).String())

	// store the swap to repository
	err := R().UniswapAdd(&types.Swap{
		OrdIndex:    log.Trx.Uid(),
		BlockNumber: &log.Block.Number,
		Type:        types.SwapMint,
		TimeStamp:   &log.Block.TimeStamp,
		Pair:        log.Address,
		Sender:      log.Trx.From,
		Hash:        log.Trx.Hash,
		Amount0In:   a0,
		Amount0Out:  z,
		Amount1In:   a1,
		Amount1Out:  z,
		Reserve0:    z,
		Reserve1:    z,
	})
	if err != nil {
		R().Log().Errorf("could not store uniswap event; %s", err.Error())
	}
}

// handleUniswapBurn processes Uniswap Burn event log emitted when a sender claims liquidity
// from an Uniswap token pair to decrease their share on the pair and claim accumulated rewards.
// UniswapPair::Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func handleUniswapBurn(log *types.LogRecord) {
	// sanity check for data (2 x uint256 = 2x32 bytes = 64 bytes), (1 x subject topic + 2 x address = 3 topics)
	if len(log.Data) != 64 || len(log.Topics) != 3 {
		R().Log().Criticalf("%s log invalid data length; expected 64 bytes, %d bytes given; expected 3 topics, %d given", log.TxHash.String(), len(log.Data), len(log.Topics))
		return
	}

	// get the reserve values
	a0 := new(big.Int).SetBytes(log.Data[:32])
	a1 := new(big.Int).SetBytes(log.Data[32:])
	sender := common.BytesToAddress(log.Topics[1].Bytes())
	z := new(big.Int)

	// log what we got
	R().Log().Debugf("uniswap MINT on pair %s, block #%d, for %s, amount0: %s, amount1: %s",
		log.Address.String(),
		log.Block.Number,
		sender.String(),
		((*hexutil.Big)(a0)).String(),
		((*hexutil.Big)(a1)).String(),
	)

	// store the swap to repository
	err := R().UniswapAdd(&types.Swap{
		OrdIndex:    log.Trx.Uid(),
		BlockNumber: &log.Block.Number,
		Type:        types.SwapBurn,
		TimeStamp:   &log.Block.TimeStamp,
		Pair:        log.Address,
		Sender:      log.Trx.From,
		Hash:        log.Trx.Hash,
		Amount0In:   z,
		Amount0Out:  a0,
		Amount1In:   z,
		Amount1Out:  a1,
		Reserve0:    z,
		Reserve1:    z,
	})
	if err != nil {
		R().Log().Errorf("could not store uniswap event; %s", err.Error())
	}
}

// handleUniswapSync processes Uniswap Sync event log.
// UniswapPair::Sync(uint112 reserve0, uint112 reserve1)
func handleUniswapSync(log *types.LogRecord) {
	// sanity check for data (2 x uint112 = 2x32 bytes = 64 bytes)
	if len(log.Data) != 64 {
		R().Log().Criticalf("%s log invalid data length; expected 64 bytes, %d bytes given; expected 1 topic, %d given", log.TxHash.String(), len(log.Data), len(log.Topics))
		return
	}

	// get the reserve values
	r0 := new(big.Int).SetBytes(log.Data[:32])
	r1 := new(big.Int).SetBytes(log.Data[32:])
	z := new(big.Int)

	// log what we got
	R().Log().Debugf("uniswap SYNC on pair %s, block #%d, reserve0: %s, reserve1: %s",
		log.Address.String(),
		log.Block.Number,
		((*hexutil.Big)(r0)).String(),
		((*hexutil.Big)(r1)).String(),
	)

	// store the swap to repository
	err := R().UniswapAdd(&types.Swap{
		OrdIndex:    log.Trx.Uid(),
		BlockNumber: &log.Block.Number,
		Type:        types.SwapSync,
		TimeStamp:   &log.Block.TimeStamp,
		Pair:        log.Address,
		Sender:      log.Trx.From,
		Hash:        log.Trx.Hash,
		Amount0In:   z,
		Amount0Out:  z,
		Amount1In:   z,
		Amount1Out:  z,
		Reserve0:    r0,
		Reserve1:    r1,
	})
	if err != nil {
		R().Log().Errorf("could not store uniswap event; %s", err.Error())
	}
}

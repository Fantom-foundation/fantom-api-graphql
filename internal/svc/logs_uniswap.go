// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// handleUniswapSwap processes Uniswap Swap event lr emitted when a sender trades
// input tokens to gain output tokens, this is the basic type of trade on an Uniswap pair.
// UniswapPair::Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func handleUniswapSwap(lr *types.LogRecord) {
	// sanity check for data (4 x uint256 = 4x32 bytes = 128 bytes), (1 x subject topic + 2 x address = 3 topics)
	if len(lr.Data) != 128 || len(lr.Topics) != 3 {
		log.Criticalf("%s lr invalid data length; expected 128 bytes, %d bytes given; expected 3 topics, %d given",
			lr.TxHash.String(),
			len(lr.Data),
			len(lr.Topics),
		)
		return
	}

	// get the reserve values
	in0 := new(big.Int).SetBytes(lr.Data[:32])
	in1 := new(big.Int).SetBytes(lr.Data[32:64])
	out0 := new(big.Int).SetBytes(lr.Data[64:96])
	out1 := new(big.Int).SetBytes(lr.Data[96:])
	sender := common.BytesToAddress(lr.Topics[1].Bytes())
	z := new(big.Int)

	// lr what we got
	log.Debugf("uniswap SWAP on pair %s, block #%d, for %s, in0: %s, in1: %s, out0: %s, out1: %s",
		lr.Address.String(),
		lr.Block.Number,
		sender.String(),
		((*hexutil.Big)(in0)).String(),
		((*hexutil.Big)(in1)).String(),
		((*hexutil.Big)(out0)).String(),
		((*hexutil.Big)(out1)).String(),
	)

	// store the swap to repository
	err := repo.UniswapAdd(&types.Swap{
		OrdIndex:    lr.Trx.Uid(),
		BlockNumber: &lr.Block.Number,
		Type:        types.SwapMint,
		TimeStamp:   &lr.Block.TimeStamp,
		Pair:        lr.Address,
		Sender:      lr.Trx.From,
		Hash:        lr.Trx.Hash,
		Amount0In:   in0,
		Amount0Out:  out0,
		Amount1In:   in1,
		Amount1Out:  out1,
		Reserve0:    z,
		Reserve1:    z,
	})
	if err != nil {
		log.Errorf("could not store uniswap event; %s", err.Error())
	}
}

// handleUniswapMint processes Uniswap Mint event lr emitted when a sender adds new liquidity
// to an Uniswap token pair to increase their share on the pair rewards.
// UniswapPair::Mint(address indexed sender, uint256 amount0, uint256 amount1)
func handleUniswapMint(lr *types.LogRecord) {
	// sanity check for data (2 x uint256 = 2x32 bytes = 64 bytes), (1 x subject topic + 1 x address = 2 topics)
	if len(lr.Data) != 64 || len(lr.Topics) != 2 {
		log.Criticalf("%s lr invalid data length; expected 64 bytes, %d bytes given; expected 2 topics, %d given",
			lr.TxHash.String(),
			len(lr.Data),
			len(lr.Topics),
		)
		return
	}

	// get the reserve values
	a0 := new(big.Int).SetBytes(lr.Data[:32])
	a1 := new(big.Int).SetBytes(lr.Data[32:])
	sender := common.BytesToAddress(lr.Topics[1].Bytes())
	z := new(big.Int)

	// lr what we got
	log.Debugf("uniswap MINT on pair %s, block #%d, for %s, amount0: %s, amount1: %s",
		lr.Address.String(),
		lr.Block.Number,
		sender.String(),
		((*hexutil.Big)(a0)).String(),
		((*hexutil.Big)(a1)).String(),
	)

	// store the swap to repository
	err := repo.UniswapAdd(&types.Swap{
		OrdIndex:    lr.Trx.Uid(),
		BlockNumber: &lr.Block.Number,
		Type:        types.SwapMint,
		TimeStamp:   &lr.Block.TimeStamp,
		Pair:        lr.Address,
		Sender:      lr.Trx.From,
		Hash:        lr.Trx.Hash,
		Amount0In:   a0,
		Amount0Out:  z,
		Amount1In:   a1,
		Amount1Out:  z,
		Reserve0:    z,
		Reserve1:    z,
	})
	if err != nil {
		log.Errorf("could not store uniswap event; %s", err.Error())
	}
}

// handleUniswapBurn processes Uniswap Burn event lr emitted when a sender claims liquidity
// from an Uniswap token pair to decrease their share on the pair and claim accumulated rewards.
// UniswapPair::Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func handleUniswapBurn(lr *types.LogRecord) {
	// sanity check for data (2 x uint256 = 2x32 bytes = 64 bytes), (1 x subject topic + 2 x address = 3 topics)
	if len(lr.Data) != 64 || len(lr.Topics) != 3 {
		log.Criticalf("%s lr invalid data length; expected 64 bytes, %d bytes given; expected 3 topics, %d given",
			lr.TxHash.String(),
			len(lr.Data),
			len(lr.Topics),
		)
		return
	}

	// get the reserve values
	a0 := new(big.Int).SetBytes(lr.Data[:32])
	a1 := new(big.Int).SetBytes(lr.Data[32:])
	sender := common.BytesToAddress(lr.Topics[1].Bytes())
	z := new(big.Int)

	// lr what we got
	log.Debugf("uniswap MINT on pair %s, block #%d, for %s, amount0: %s, amount1: %s",
		lr.Address.String(),
		lr.Block.Number,
		sender.String(),
		((*hexutil.Big)(a0)).String(),
		((*hexutil.Big)(a1)).String(),
	)

	// store the swap to repository
	err := repo.UniswapAdd(&types.Swap{
		OrdIndex:    lr.Trx.Uid(),
		BlockNumber: &lr.Block.Number,
		Type:        types.SwapBurn,
		TimeStamp:   &lr.Block.TimeStamp,
		Pair:        lr.Address,
		Sender:      lr.Trx.From,
		Hash:        lr.Trx.Hash,
		Amount0In:   z,
		Amount0Out:  a0,
		Amount1In:   z,
		Amount1Out:  a1,
		Reserve0:    z,
		Reserve1:    z,
	})
	if err != nil {
		log.Errorf("could not store uniswap event; %s", err.Error())
	}
}

// handleUniswapSync processes Uniswap Sync event lr.
// UniswapPair::Sync(uint112 reserve0, uint112 reserve1)
func handleUniswapSync(lr *types.LogRecord) {
	// sanity check for data (2 x uint112 = 2x32 bytes = 64 bytes)
	if len(lr.Data) != 64 {
		log.Criticalf("%s lr invalid data length; expected 64 bytes, %d bytes given; expected 1 topic, %d given",
			lr.TxHash.String(),
			len(lr.Data),
			len(lr.Topics),
		)
		return
	}

	// get the reserve values
	r0 := new(big.Int).SetBytes(lr.Data[:32])
	r1 := new(big.Int).SetBytes(lr.Data[32:])
	z := new(big.Int)

	// lr what we got
	log.Debugf("uniswap SYNC on pair %s, block #%d, reserve0: %s, reserve1: %s",
		lr.Address.String(),
		lr.Block.Number,
		((*hexutil.Big)(r0)).String(),
		((*hexutil.Big)(r1)).String(),
	)

	// store the swap to repository
	err := repo.UniswapAdd(&types.Swap{
		OrdIndex:    lr.Trx.Uid(),
		BlockNumber: &lr.Block.Number,
		Type:        types.SwapSync,
		TimeStamp:   &lr.Block.TimeStamp,
		Pair:        lr.Address,
		Sender:      lr.Trx.From,
		Hash:        lr.Trx.Hash,
		Amount0In:   z,
		Amount0Out:  z,
		Amount1In:   z,
		Amount1Out:  z,
		Reserve0:    r0,
		Reserve1:    r1,
	})
	if err != nil {
		log.Errorf("could not store uniswap event; %s", err.Error())
	}
}

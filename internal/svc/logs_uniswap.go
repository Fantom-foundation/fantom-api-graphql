// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// uniswapKnownPairs represents a map of known pairs and their participation in our approved Uniswap instance.
var uniswapKnownPairs = make(map[common.Address]bool, 1000)

// uniswapOrdinalIndex calculates ordinal index of the given Uniswap transaction.
func uniswapOrdinalIndex(lr *types.LogRecord) uint64 {
	return ((uint64(lr.Block.Number) << 14) & 0x7FFFFFFFFFFFFFFF) | ((uint64(lr.TxIndex) << 8) & 0x3fff) | (uint64(lr.Index) & 0xff)
}

// isKnownUniswapPair checks if the given (expected) uniswap pair address
// belongs to our approved Uniswap instance.
func isKnownUniswapPair(pair *common.Address) bool {
	// do we know the address already?
	apr, ok := uniswapKnownPairs[*pair]
	if !ok {
		apr = false

		// get the full list of all known pairs and use it to check this one
		l, err := repo.UniswapPairs()
		if err != nil {
			log.Errorf("unknown uniswap pairs list; %s", err.Error())
			return false
		}
		for _, a := range l {
			if a.String() == pair.String() {
				apr = true
				break
			}
		}

		// store the value locally for future use
		uniswapKnownPairs[*pair] = apr
	}
	return apr
}

// handleUniswapSwap processes Uniswap Swap event lr emitted when a sender trades
// input tokens to gain output tokens, this is the basic type of trade on an Uniswap pair.
// UniswapPair::Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func handleUniswapSwap(lr *types.LogRecord) {
	if !isKnownUniswapPair(&lr.Address) {
		log.Debugf("rejected %s uniswap swap #%d on unknown pair", lr.TxHash.String(), lr.Index)
		return
	}

	// sanity check for data (4 x uint256 = 4x32 bytes = 128 bytes), (1 x subject topic + 2 x address = 3 topics)
	if len(lr.Data) != 128 || len(lr.Topics) != 3 {
		log.Errorf("%s invalid data length; expected 128 bytes, %d bytes given; expected 3 topics, %d given",
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
		OrdIndex:    uniswapOrdinalIndex(lr),
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
		log.Errorf("%s could not store uniswap event #%d; %s", lr.TxHash.String(), lr.Index, err.Error())
	}
}

// handleUniswapMint processes Uniswap Mint event lr emitted when a sender adds new liquidity
// to an Uniswap token pair to increase their share on the pair rewards.
// UniswapPair::Mint(address indexed sender, uint256 amount0, uint256 amount1)
func handleUniswapMint(lr *types.LogRecord) {
	if !isKnownUniswapPair(&lr.Address) {
		log.Debugf("rejected %s uniswap mint #%d on unknown pair", lr.TxHash.String(), lr.Index)
		return
	}

	// sanity check for data (2 x uint256 = 2x32 bytes = 64 bytes), (1 x subject topic + 1 x address = 2 topics)
	if len(lr.Data) != 64 || len(lr.Topics) != 2 {
		log.Errorf("%s invalid data length; expected 64 bytes, %d bytes given; expected 2 topics, %d given",
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
		OrdIndex:    uniswapOrdinalIndex(lr),
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
		log.Errorf("%s could not store uniswap event #%d; %s", lr.TxHash.String(), lr.Index, err.Error())
	}
}

// handleUniswapBurn processes Uniswap Burn event lr emitted when a sender claims liquidity
// from an Uniswap token pair to decrease their share on the pair and claim accumulated rewards.
// UniswapPair::Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func handleUniswapBurn(lr *types.LogRecord) {
	if !isKnownUniswapPair(&lr.Address) {
		log.Debugf("rejected %s uniswap burn #%d on unknown pair", lr.TxHash.String(), lr.Index)
		return
	}

	// sanity check for data (2 x uint256 = 2x32 bytes = 64 bytes), (1 x subject topic + 2 x address = 3 topics)
	if len(lr.Data) != 64 || len(lr.Topics) != 3 {
		log.Errorf("%s invalid data length; expected 64 bytes, %d bytes given; expected 3 topics, %d given",
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
		OrdIndex:    uniswapOrdinalIndex(lr),
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
		log.Errorf("%s could not store uniswap event #%d; %s", lr.TxHash.String(), lr.Index, err.Error())
	}
}

// handleUniswapSync processes Uniswap Sync event lr.
// UniswapPair::Sync(uint112 reserve0, uint112 reserve1)
func handleUniswapSync(lr *types.LogRecord) {
	if !isKnownUniswapPair(&lr.Address) {
		log.Debugf("rejected %s uniswap sync #%d on unknown pair", lr.TxHash.String(), lr.Index)
		return
	}

	// sanity check for data (2 x uint112 = 2x32 bytes = 64 bytes)
	if len(lr.Data) != 64 {
		log.Errorf("%s invalid data length; expected 64 bytes, %d bytes given; expected 1 topic, %d given",
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
		OrdIndex:    uniswapOrdinalIndex(lr),
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
		log.Errorf("%s could not store uniswap event #%d; %s", lr.TxHash.String(), lr.Index, err.Error())
	}
}

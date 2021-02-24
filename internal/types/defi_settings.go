// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// DefiSettings represents a set of current DeFi contract configuration options
// and details used to calculate various DeFi specific operators and fees.
type DefiSettings struct {
	// MintFee4 is the current fee applied to all minting operations in fMint protocol.
	// Value is represented in 4 digits, e.g. value 25 = 0.0025 => 0.25% fee.
	MintFee4 hexutil.Big

	// MinCollateralRatio4 is the minimal allowed ratio between
	// collateral and debt values in ref. denomination (fUSD)
	// on which the borrow trade is allowed.
	// Value is represented in 4 digits,
	// e.g. value 25000 = 3.0x => (debt x 3.0 <= collateral)
	MinCollateralRatio4 hexutil.Big

	// rewardCollateralRatio4 is the minimal ratio between
	// collateral and debt values in ref. denomination (fUSD)
	// on which the account is eligible for rewards distribution.
	// Collateral below this ratio means all the pending rewards
	// will be burnt and lost.
	RewardCollateralRatio4 hexutil.Big

	// Decimals represents the decimals / digits correction
	// applied to the fees and ratios internally to correctly represent
	// fraction numbers. E.g. correction value 4 => ratio/fee x 10000.
	Decimals int32

	// FMintContract represents the address of the DeFi fMint
	// core contract.
	FMintContract common.Address

	// FMintAddressProvider represents the address of the fMint address provider.
	FMintAddressProvider common.Address

	// FMintTokenRegistry represents the address of the DeFi token registry.
	FMintTokenRegistry common.Address

	// FMintRewardDistribution represents the address of the DeFi fMint
	// reward distribution contract.
	FMintRewardDistribution common.Address

	// FMintCollateralPool represents the address of the DeFi fMint
	// collateral pool contract.
	FMintCollateralPool common.Address

	// FMintDebtPool represents the address of the DeFi fMint
	// debt pool contract.
	FMintDebtPool common.Address

	// PriceOracleAggregate is the address of the current price oracle
	// aggregate used by the DeFe to obtain USD price of tokens managed.
	PriceOracleAggregate common.Address

	// FLendingPool is the address of the Lending pool interface for fLend
	FLendingPool common.Address
}

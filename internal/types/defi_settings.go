// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// DefiSettings represents a set of current DeFi contract configuration options
// and details used to calculate various DeFi specific operators and fees.
type DefiSettings struct {
	// TradeFee4 is the current fee applied to all direct trading operations.
	// Value is represented in 4 digits, e.g. value 25 = 0.0025 => 0.25% fee.
	TradeFee4 hexutil.Big

	// LoanFee4 is the current entry fee applied to all lending operations.
	// Value is represented in 4 digits, e.g. value 25 = 0.0025 => 0.25% fee.
	LoanFee4 hexutil.Big

	// MinCollateralRatio4 is the minimal allowed ratio between
	// collateral and debt values in ref. denomination (fUSD)
	// on which the borrow trade is allowed.
	// Value is represented in 4 digits,
	// e.g. value 25000 = 3.0x => (debt x 3.0 <= collateral)
	MinCollateralRatio4 hexutil.Big

	// MinCollateralRatio4 is the minimal allowed ratio between
	// collateral and debt values in ref. denomination (fUSD)
	// on which the borrow trade is allowed.
	// Value is represented in 4 digits,
	// e.g. value 25000 = 2.5x => (debt x 2.25 <= collateral)
	WarningCollateralRatio4 hexutil.Big

	// LiqCollateralRatio4 is the liquidation ratio between
	// collateral and debt values in ref. denomination (fUSD).
	// If the current ratio drops below this value, the position
	// is liquidated.
	// Value is represented in 4 digits,
	// e.g. value 15000 = 1.5x => (debt x 1.5 <= collateral)
	LiqCollateralRatio4 hexutil.Big

	// Decimals represents the decimals / digits correction
	// applied to the fees and ratios internally to correctly represent
	// fraction numbers. E.g. correction value 4 => ratio/fee x 10000.
	Decimals int32

	// PriceOracleAggregate is the address of the current price oracle
	// aggregate used by the DeFe to obtain USD price of tokens managed.
	PriceOracleAggregate common.Address
}

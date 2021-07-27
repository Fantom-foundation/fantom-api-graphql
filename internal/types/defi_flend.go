// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ReserveData represents a Lending pool asset reserve data.
type ReserveData struct {

	// address of the asset
	AssetAddress common.Address

	// number in the reserveList() array
	ID int32

	// bitmask encoded asset reserve configuration data
	Configuration hexutil.Big

	// liquidity index in ray
	LiquidityIndex hexutil.Big

	// variable borrow index in ray
	VariableBorrowIndex hexutil.Big

	// current supply / liquidity / deposit rate in ray
	CurrentLiquidityRate hexutil.Big

	// current variable borrow rate in ray
	CurrentVariableBorrowRate hexutil.Big

	// current stable borrow rate in ray
	CurrentStableBorrowRate hexutil.Big

	// timestamp of when reserve data was last updated
	LastUpdateTimestamp hexutil.Big

	// address of associated aToken (tokenised deposit)
	ATokenAddress common.Address

	// address of associated stable debt token
	StableDebtTokenAddress common.Address

	// address of associated variable debt token
	VariableDebtTokenAddress common.Address

	// address of interest rate strategy
	InterestRateStrategyAddress common.Address
}

// FLendUserAccountData represents a Lending pool user data.
type FLendUserAccountData struct {

	// total collateral in FUSD of the user
	TotalCollateralFUSD hexutil.Big

	// total debt in FUSD of the user
	TotalDebtFUSD hexutil.Big

	// borrowing power left of the user in FUSD
	AvailableBorrowsFUSD hexutil.Big

	// liquidation threshold of the user
	CurrentLiquidationThreshold hexutil.Big

	// Loan To Value of the user
	Ltv hexutil.Big

	// current health factor of the user
	HealthFactor hexutil.Big

	// configuration data
	ConfigurationData hexutil.Big
}

// FLendDeposit represents a Lending pool deposit event data.
type FLendDeposit struct {

	// address of the asset
	AssetAddress common.Address

	// address of the user
	UserAddress common.Address

	// address of the on behalf of
	OnBehalfOfAddress common.Address

	// deposit amount
	Amount hexutil.Big

	// referral code
	ReferralCode int32

	// time of deposit
	Timestamp hexutil.Uint64
}

// FLendBorrow represents a Lending pool borrow event data.
type FLendBorrow struct {

	// address of the asset
	AssetAddress common.Address

	// address of the user
	UserAddress common.Address

	// address of the on behalf of
	OnBehalfOfAddress common.Address

	// deposit amount
	Amount hexutil.Big

	// interest rate mode
	InterestRateMode int32

	// borrow rate
	BorrowRate int32

	// referral code
	ReferralCode int32

	// time of deposit
	Timestamp hexutil.Uint64
}

// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// DefiToken represents a single DeFi managed token information.
type DefiToken struct {
	// Address of the token is used as the token's unique identifier.
	Address common.Address `json:"address"`

	// Index represents the index of the token in the registry
	// starting from 1.
	Index hexutil.Uint64 `json:"index"`

	// Name represents an extended name of the token.
	Name string `json:"name"`

	// Symbol represents an abbreviation for the token.
	Symbol string `json:"symbol"`

	// LogoUrl represents an URL address of the token logo image.
	LogoUrl string `json:"logo"`

	// Decimals is the number of decimals the token supports.
	// The most common value is 18 to mimic the ETH to WEI relationship.
	// USD pairs on ChainLink (we use for price oracles) use 8 digits.
	Decimals int32 `json:"decimals"`

	// PriceDecimals is the number of decimals the price oracle of the token uses.
	// USD pairs of the the ChainLink compatible price oracle we utilize
	// usually have 8 digits.
	PriceDecimals int32 `json:"priceDecimals"`

	// IsActive signals if the token can be used in the DeFi functions at all.
	IsActive bool `json:"isActive"`

	// CanDeposit signals if the token can be used in deposit as a collateral asset.
	CanDeposit bool `json:"canDeposit"`

	// CanMint signals if the token can be used in fMint as a target asset.
	CanMint bool `json:"canMint"`

	// canBorrow signals if the token is available for FLend borrow operations.
	CanBorrow bool `json:"canBorrow"`

	// canTrade signals if the token is available for FTrade direct trading operations.
	CanTrade bool `json:"canTrade"`

	// VolatilityIndex represents an index of volatility of the token used internally.
	VolatilityIndex hexutil.Big `json:"volatilityIndex"`
}

// DefiTokenType represents the type of a token analyzed.
type DefiTokenType string

// types of token recognized by the DeFi protocol
const (
	DefiTokenTypeCollateral DefiTokenType = "COLLATERAL"
	DefiTokenTypeDebt                     = "DEBT"
)

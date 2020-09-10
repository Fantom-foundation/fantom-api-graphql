// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// FMintAccount represents information about specific DeFi/fMint account
// identified by the owner's address.
type FMintAccount struct {
	// address of the DeFi account.
	Address common.Address

	// Collateral represents a list of all collateral assets on this account.
	CollateralList []common.Address

	// Debt represents the list of all the borrowed tokens.
	DebtList []common.Address

	// CollateralValue represents the current collateral value
	// in ref. denomination (fUSD).
	CollateralValue hexutil.Big

	// DebtValue represents the current debt value
	// in ref. denomination (fUSD).
	DebtValue hexutil.Big
}

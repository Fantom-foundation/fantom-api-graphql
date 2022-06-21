// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"time"
)

const (
	// AccountTypeWallet identifies a generic account, probably a wallet.
	AccountTypeWallet = iota

	// AccountTypeContract identifies a contract of unknown detailed type.
	AccountTypeContract

	// AccountTypeSFCContract identifies the Special Function Contract account type.
	AccountTypeSFCContract

	// AccountTypeERC20Contract identifies a contract of type ERC20 token.
	AccountTypeERC20Contract

	// AccountTypeERC721Contract identifies a contract of type ERC721 non-fungible token.
	AccountTypeERC721Contract

	// AccountTypeERC1155Contract identifies a non-fungible token contract of type ERC1155.
	AccountTypeERC1155Contract
)

// Account represents an Opera account at the blockchain.
type Account struct {
	Name            string          `bson:"name"`
	Address         common.Address  `bson:"addr"`
	AccountType     int             `bson:"act"`
	IsContract      bool            `bson:"is_code"`
	FirstAppearance time.Time       `bson:"since"`
	DeployedBy      *common.Address `bson:"deployer"`
	DeploymentTrx   *common.Hash    `bson:"dtx"`
}

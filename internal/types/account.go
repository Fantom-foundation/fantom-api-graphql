// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"time"
)

const (
	// AccountTypeUnknown identifies a generic account, probably a wallet.
	AccountTypeUnknown = "WALLET"

	// AccountTypeContract identifies a contract of unknown detailed type.
	AccountTypeContract = "CONTRACT"

	// AccountTypeSFCContract identifies the Special Function Contract account type.
	AccountTypeSFCContract = "SFC"

	// AccountTypeERC20Contract identifies a contract of type ERC20 token.
	AccountTypeERC20Contract = "ERC20"

	// AccountTypeERC721Contract identifies a contract of type ERC721 non-fungible token.
	AccountTypeERC721Contract = "ERC721"

	// AccountTypeERC1155Contract identifies a non-fungible token contract of type ERC1155.
	AccountTypeERC1155Contract = "ERC1155"
)

// Account represents an Opera account at the blockchain.
type Account struct {
	Name            string          `bson:"name"`
	Address         common.Address  `bson:"addr"`
	AccountType     string          `bson:"act"`
	IsContract      bool            `bson:"is_contract"`
	FirstAppearance time.Time       `bson:"ts"`
	DeployedBy      *common.Address `bson:"deployer"`
	DeploymentTrx   *common.Hash    `bson:"dep_trx"`
	LastActivity    time.Time       `bson:"last_tx"`
	TrxCounter      hexutil.Uint64  `bson:"nonce"`
}

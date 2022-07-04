// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/repository/rpc"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// handleErcTokenApproval handles Approval event on ERC20/721.
// event Approval(address indexed owner, address indexed spender, uint256 value)
// event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func handleErcTokenApproval(lr *types.LogRecord) {
	if isErc20Transaction(lr) {
		processErc20Transaction(lr, types.TokenTrxTypeApproval)
		return
	}
	if !isErc721Transaction(lr) {
		// log all but erc721 and erc20 approvals
		log.Debugf("Unrecognized ERC-20/721 Approval from tx %s (%d data bytes, %d topics)", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
	}
}

// handleErcTokenTransfer handles Transfer event on ERC20.
// event Transfer(address indexed from, address indexed to, uint256 value)
// event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func handleErcTokenTransfer(lr *types.LogRecord) {
	if isErc20Transaction(lr) {
		processErc20Transaction(lr, types.TokenTrxTypeTransfer)
		return
	}
	if isErc721Transaction(lr) {
		processErc721Transfer(lr)
		return
	}
	log.Debugf("Unrecognized ERC-20/721 Transfer from tx %s (%d data bytes, %d topics)", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
}

// handleErc1155TransferSingle handles Transfer event on ERC1155
// event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 tokenId, uint256 value)
func handleErc1155TransferSingle(lr *types.LogRecord) {
	// 3 indexed params, 2 uint256 params
	if len(lr.Topics) != 4 || len(lr.Data) != 64 {
		log.Debugf("Unrecognized ERC1155 TransferSingle from tx %s (%d data bytes, %d topics)", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	// extract data
	from := common.BytesToAddress(lr.Topics[2].Bytes())
	to := common.BytesToAddress(lr.Topics[3].Bytes())
	tokenId := hexutil.Big(*new(big.Int).SetBytes(lr.Data[:32]))
	var tokenName *string

	if uri, _ := repo.Erc1155Uri(&lr.Address, tokenId.ToInt()); uri != "" {
		tokenName = getTokenName(uri)
	}

	// updated recipient balance if method is not burn
	if to.String() != config.EmptyAddress {
		storeOwnership(lr.Address, tokenId, to, erc1155Balance(&lr.Address, &to, tokenId), lr.Block.TimeStamp, lr.TxHash, tokenName)
	}

	// updated sender balance if method is not mint
	if from.String() != config.EmptyAddress {
		storeOwnership(lr.Address, tokenId, from, erc1155Balance(&lr.Address, &from, tokenId), lr.Block.TimeStamp, lr.TxHash, tokenName)
	}
}

// handleErc1155TransferBatch handles TransferBatch event on ERC1155
// event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func handleErc1155TransferBatch(lr *types.LogRecord) {
	// 3 indexed params
	if len(lr.Topics) != 4 {
		log.Debugf("Unrecognized ERC-1155 TransferBatch from tx %s (%d data bytes, %d topics)", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	from := common.BytesToAddress(lr.Topics[2].Bytes())
	to := common.BytesToAddress(lr.Topics[3].Bytes())
	ids, values, err := rpc.Erc1155ParseTransferBatchData(lr.Data)

	if err != nil {
		log.Errorf("failed to parse ERC1155 TransferBatch data - trx %s; %s", lr.TxHash.String(), err.Error())
	}

	if len(ids) != len(values) {
		log.Errorf("ERC1155 TransferBatch ids and values length differs - trx %s", lr.TxHash.String())
	}

	for i := range ids {
		id := hexutil.Big(*ids[i])
		var tokenName *string

		if uri, _ := repo.Erc1155Uri(&lr.Address, id.ToInt()); uri != "" {
			tokenName = getTokenName(uri)
		}

		// updated recipient balance if method is not burn
		if to.String() != config.EmptyAddress {
			storeOwnership(lr.Address, id, to, erc1155Balance(&lr.Address, &to, id), lr.Block.TimeStamp, lr.TxHash, tokenName)
		}

		// updated sender balance if method is not mint
		if from.String() != config.EmptyAddress {
			storeOwnership(lr.Address, id, from, erc1155Balance(&lr.Address, &from, id), lr.Block.TimeStamp, lr.TxHash, tokenName)
		}
	}
}

// processErc20Transaction handles Approval and/or Transfer event on an ERC20 token.
func processErc20Transaction(lr *types.LogRecord, trxType int32) {
	from := common.BytesToAddress(lr.Topics[1].Bytes())
	to := common.BytesToAddress(lr.Topics[2].Bytes())
	amount := new(big.Int).SetBytes(lr.Data[:])

	decimals, err := repo.Erc20Decimals(&lr.Address)
	if err != nil {
		log.Errorf("can not get token decimals for token address %s, call %s; %s", lr.Address.String(), lr.TxHash.String(), err.Error())
		return
	}

	if err := repo.StoreTokenTransaction(&types.TokenTransaction{
		Transaction:  lr.TxHash,
		TrxIndex:     uint64(lr.TxIndex),
		TokenAddress: lr.Address,
		TrxType:      tokenTrxType(trxType, from, to),
		Sender:       from,
		Recipient:    to,
		Amount:       hexutil.Big(*amount),
		TimeStamp:    time.Unix(int64(lr.Block.TimeStamp), 0),
		LogIndex:     uint64(lr.Index),
		BlockNumber:  lr.BlockNumber,
	}, uint8(decimals)); err != nil {
		log.Errorf("can not store token transaction for call %s; %s", lr.TxHash.String(), err.Error())
	}
}

// processErc721Transfer handles Transfer event on ERC721.
func processErc721Transfer(lr *types.LogRecord) {
	// extract details
	from := common.BytesToAddress(lr.Topics[1].Bytes())
	to := common.BytesToAddress(lr.Topics[2].Bytes())
	tokenId := hexutil.Big(*new(big.Int).SetBytes(lr.Topics[3].Bytes()))
	var tokenName *string

	if uri, _ := repo.Erc721TokenURI(&lr.Address, tokenId.ToInt()); uri != "" {
		tokenName = getTokenName(uri)
	}

	// updated recipient balance if method is not burn
	if to.String() != config.EmptyAddress {
		storeOwnership(lr.Address, tokenId, to, erc721Balance(&lr.Address, &to, tokenId), lr.Block.TimeStamp, lr.TxHash, tokenName)
	}

	// updated sender balance if method is not mint
	if from.String() != config.EmptyAddress {
		storeOwnership(lr.Address, tokenId, from, erc721Balance(&lr.Address, &from, tokenId), lr.Block.TimeStamp, lr.TxHash, tokenName)
	}
}

// getTokenName returns token name.
func getTokenName(uri string) *string {
	md, err := repo.GetTokenJsonMetadata(uri)
	if err != nil {
		log.Errorf("could not fetch token json meta data; %s", err.Error())
		return nil
	}
	return md.Name
}

// erc721Balance returns current balance for given contract, owner and token identifier.
func erc721Balance(contract *common.Address, owner *common.Address, tokenId hexutil.Big) hexutil.Big {
	no, err := repo.Erc721OwnerOf(contract, tokenId.ToInt())
	if err != nil {
		log.Errorf("erc721 token balance unknown; %s", err.Error())
		return hexutil.Big{}
	}

	// balance is 1 when owner matches
	if no.String() == owner.String() {
		return *(*hexutil.Big)(new(big.Int).SetUint64(1))
	}

	return hexutil.Big{}
}

// erc1155Balance returns current balance for given contract, owner and token identifier.
func erc1155Balance(contract *common.Address, owner *common.Address, tokenId hexutil.Big) hexutil.Big {
	qty, err := repo.Erc1155BalanceOf(contract, owner, tokenId.ToInt())
	if err != nil {
		log.Criticalf("erc1155 token balance unknown; %s", err.Error())
		return hexutil.Big{}
	}

	return hexutil.Big(*qty)
}

// tokenTrxType detects detailed type of ERC transfer based on common type and addresses involved.
func tokenTrxType(trxType int32, from common.Address, to common.Address) int32 {
	if trxType == types.TokenTrxTypeTransfer {
		if config.EmptyAddress == from.String() {
			return types.TokenTrxTypeMint
		}
		if config.EmptyAddress == to.String() {
			return types.TokenTrxTypeBurn
		}
	}
	return trxType
}

// isErc20Transaction checks is valid ERC20 transaction
func isErc20Transaction(lr *types.LogRecord) bool {
	// ERC20 has 2 indexed params (=> 3 topics) and 1 non-indexed uint256 param (=> 32 bytes)
	return len(lr.Topics) == 3 && len(lr.Data) == 32
}

// isErc721Transaction checks is valid ERC721 transaction
func isErc721Transaction(lr *types.LogRecord) bool {
	// ERC721 has 3 indexed params (=> 4 topics) and no non-indexed param (=> 0 bytes)
	return len(lr.Topics) == 4 && len(lr.Data) == 0
}

// storeOwnership stores nft ownership into persistent storage
func storeOwnership(contract common.Address, tokenId hexutil.Big, owner common.Address, amount hexutil.Big, obtained hexutil.Uint64, trx common.Hash, tokenName *string) {
	if err := repo.StoreNftOwnership(&types.NftOwnership{
		Contract:  contract,
		TokenId:   tokenId,
		Owner:     owner,
		Amount:    amount,
		Obtained:  time.Unix(int64(obtained), 0),
		Trx:       trx,
		TokenName: tokenName,
	}); err != nil {
		log.Errorf("failed to update nft ownership at %s/%d; %s",
			contract.String(), tokenId.ToInt(), err.Error())
	}
}

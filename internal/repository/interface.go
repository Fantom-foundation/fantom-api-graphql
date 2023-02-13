/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/repository/p2p"
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"math/big"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	etc "github.com/ethereum/go-ethereum/core/types"
)

// Repository interface defines functions the underlying implementation provides to API resolvers.
type Repository interface {
	// Account returns account at Opera blockchain for an address, nil if not found.
	Account(*common.Address) (*types.Account, error)

	// AccountBalance returns the current balance of an account at Opera blockchain.
	AccountBalance(*common.Address) (*hexutil.Big, error)

	// AccountNonce returns the current number of sent transactions of an account at Opera blockchain.
	AccountNonce(*common.Address) (*hexutil.Uint64, error)

	// AccountTransactions returns list of transaction hashes for account at Opera blockchain.
	//
	// String cursor represents cursor based on which the list is loaded. If null,
	// it loads either from top, or bottom of the list, based on the value
	// of the integer count. The integer represents the number of transaction loaded at most.
	//
	// For positive number, the list starts right after the cursor
	// (or on top without one) and loads at most defined number of transactions older than that.
	//
	// For negative number, the list starts right before the cursor
	// (or at the bottom without one) and loads at most defined number
	// of transactions newer than that.
	//
	// Transactions are always sorted from newer to older.
	AccountTransactions(*common.Address, *common.Address, *string, int32) (*types.TransactionList, error)

	// AccountsActive total number of accounts known to repository.
	AccountsActive() (hexutil.Uint64, error)

	// AccountIsKnown checks if the account of the given address is known to the API server.
	AccountIsKnown(*common.Address) bool

	// StoreAccount adds specified account detail into the repository.
	StoreAccount(*types.Account) error

	// AccountMarkActivity marks the latest account activity in the repository.
	AccountMarkActivity(*common.Address, uint64) error

	// BlockHeight returns the current height of the Opera blockchain in blocks.
	BlockHeight() (*hexutil.Big, error)

	// LastKnownBlock returns number of the last block known to the repository.
	LastKnownBlock() (uint64, error)

	// UpdateLastKnownBlock update record about last known block.
	UpdateLastKnownBlock(blockNo *hexutil.Uint64) error

	// ObservedHeaders provides a channel fed with new headers observed
	// by the connected blockchain node.
	ObservedHeaders() chan *etc.Header

	// BlockByNumber returns a block at Opera blockchain represented by a number.
	// Top block is returned if the number is not provided.
	// If the block is not found, ErrBlockNotFound error is returned.
	BlockByNumber(*hexutil.Uint64) (*types.Block, error)

	// BlockByHash returns a block at Opera blockchain represented by a hash.
	// Top block is returned if the hash is not provided.
	// If the block is not found, ErrBlockNotFound error is returned.
	BlockByHash(*common.Hash) (*types.Block, error)

	// Blocks pulls list of blocks starting on the specified block number
	// and going up, or down based on count number.
	Blocks(*uint64, int32) (*types.BlockList, error)

	// CacheBlock puts a block to the internal block ring cache.
	CacheBlock(blk *types.Block)

	// Contract extract a smart contract information by address if available.
	Contract(*common.Address) (*types.Contract, error)

	// Contracts returns list of smart contracts at Opera blockchain.
	Contracts(bool, *string, int32) (*types.ContractList, error)

	// ValidateContract tries to validate contract byte code using
	// provided source code. If successful, the contract information
	// is updated the the repository.
	ValidateContract(*types.Contract) error

	// StoreContract updates the contract in repository.
	StoreContract(*types.Contract) error

	// SfcVersion returns current version of the SFC contract.
	SfcVersion() (hexutil.Uint64, error)

	// SfcDecimalUnit returns the decimal unit adjustment used by the SFC contract.
	SfcDecimalUnit() *big.Int

	// CurrentEpoch returns the id of the current epoch.
	CurrentEpoch() (hexutil.Uint64, error)

	// LastKnownEpoch returns the id of the last known and scanned epoch.
	LastKnownEpoch() (uint64, error)

	// AddEpoch stores an epoch reference in connected persistent storage.
	AddEpoch(e *types.Epoch) error

	// Epoch returns the id of the current epoch.
	Epoch(*hexutil.Uint64) (*types.Epoch, error)

	// CurrentSealedEpoch returns the data of the latest sealed epoch.
	CurrentSealedEpoch() (*types.Epoch, error)

	// Epochs pulls list of epochs starting at the specified cursor.
	Epochs(cursor *string, count int32) (*types.EpochList, error)

	// TotalStaked calculates current total staked amount for all stakers.
	TotalStaked() (*hexutil.Big, error)

	// RewardsAllowed returns the reward lock status from SFC.
	RewardsAllowed() (bool, error)

	// LockingAllowed indicates if the stake locking has been enabled in SFC.
	LockingAllowed() (bool, error)

	// IsSfcContract returns true if the given address points to the SFC contract.
	IsSfcContract(*common.Address) bool

	// IsStiContract returns true if the given address points to the STI contract.
	IsStiContract(*common.Address) bool

	// StoreTransaction adds a new incoming transaction from blockchain to the repository.
	StoreTransaction(*types.Block, *types.Transaction) error

	// LoadTransaction returns a transaction at Opera blockchain
	// by a hash loaded directly from the node.
	LoadTransaction(hash *common.Hash) (*types.Transaction, error)

	// Transaction returns a transaction at Opera blockchain by a hash, nil if not found.
	Transaction(*common.Hash) (*types.Transaction, error)

	// Transactions returns list of transaction hashes at Opera blockchain.
	Transactions(*string, int32) (*types.TransactionList, error)

	// TransactionsCount returns total number of transactions in the block chain.
	TransactionsCount() (uint64, error)

	// EstimateTransactionsCount returns an approximate amount of transactions on the network.
	EstimateTransactionsCount() (hexutil.Uint64, error)

	// IncTrxCountEstimate bumps the value of transaction counter estimator.
	IncTrxCountEstimate(diff uint64)

	// UpdateTrxCountEstimate updates the value of transaction counter estimator.
	UpdateTrxCountEstimate(val uint64)

	// CacheTransaction puts a transaction to the internal ring cache.
	CacheTransaction(trx *types.Transaction)

	// SendTransaction sends raw signed and RLP encoded transaction to the block chain.
	SendTransaction(hexutil.Bytes) (*types.Transaction, error)

	// LastValidatorId returns the last validator id in Opera blockchain.
	LastValidatorId() (uint64, error)

	// ValidatorsCount returns the number of stakers in Opera blockchain.
	ValidatorsCount() (uint64, error)

	// IsValidator returns TRUE if the given address is an SFC staker.
	IsValidator(*common.Address) (bool, error)

	// ValidatorAddress extract a staker address for the given staker ID.
	ValidatorAddress(*hexutil.Big) (*common.Address, error)

	// Validator extract a staker information from SFC smart contract.
	Validator(*hexutil.Big) (*types.Validator, error)

	// ValidatorByAddress extract a staker information by address.
	ValidatorByAddress(*common.Address) (*types.Validator, error)

	// ValidatorDowntime pulls information about validator downtime from the RPC interface.
	ValidatorDowntime(*hexutil.Big) (uint64, uint64, error)

	// SfcConfiguration provides SFC contract configuration.
	SfcConfiguration() (*types.SfcConfig, error)

	// SfcMaxDelegatedRatio extracts a ratio between self delegation and received stake.
	SfcMaxDelegatedRatio() (*big.Int, error)

	// PullStakerInfo extracts an extended staker information from smart contact.
	PullStakerInfo(*hexutil.Big) (*types.StakerInfo, error)

	// StoreStakerInfo stores staker information to in-memory cache for future use.
	StoreStakerInfo(*hexutil.Big, *types.StakerInfo) error

	// RetrieveStakerInfo gets staker information from in-memory if available.
	RetrieveStakerInfo(*hexutil.Big) *types.StakerInfo

	// IsDelegating returns if the given address is an SFC delegator.
	IsDelegating(*common.Address) (bool, error)

	// StoreDelegation stores a delegation in the persistent repository.
	StoreDelegation(*types.Delegation) error

	// UpdateDelegationBalance updates active balance of the given delegation.
	UpdateDelegationBalance(*common.Address, *hexutil.Big, func(*big.Int) error) error

	// Delegation returns a detail of delegation for the given address and validator ID.
	Delegation(*common.Address, *hexutil.Big) (*types.Delegation, error)

	// DelegationAmountStaked returns the current amount of staked tokens
	// for the given delegation.
	DelegationAmountStaked(*common.Address, *hexutil.Big) (*big.Int, error)

	// DelegationsByAddress returns a list of all delegations of a given delegator address.
	DelegationsByAddress(*common.Address, *string, int32) (*types.DelegationList, error)

	// DelegationsByAddressAll returns a list of all delegations of the given address un-paged.
	DelegationsByAddressAll(addr *common.Address) ([]*types.Delegation, error)

	// DelegationsOfValidator extracts a list of delegations for a validator by its ID.
	DelegationsOfValidator(*hexutil.Big, *string, int32) (*types.DelegationList, error)

	// DelegationLock returns delegation lock information using SFC contract binding.
	DelegationLock(*common.Address, *hexutil.Big) (*types.DelegationLock, error)

	// DelegationUnlockPenalty returns the amount of penalty applied on given stake unlock.
	DelegationUnlockPenalty(addr *common.Address, valID *big.Int, amount *big.Int) (hexutil.Big, error)

	// DelegationAmountUnlocked returns delegation lock information using SFC contract binding.
	DelegationAmountUnlocked(addr *common.Address, valID *big.Int) (hexutil.Big, error)

	// StoreLockedDelegation stores the given locked delegation into the database.
	StoreLockedDelegation(dl *types.LockedDelegation) error

	// AdjustLockedDelegation reduces the given locked delegation by the give amount in the database.
	AdjustLockedDelegation(common.Address, int64, int64) error

	// PendingRewards returns a detail of pending rewards for the given delegation.
	PendingRewards(*common.Address, *hexutil.Big) (*types.PendingRewards, error)

	// DelegationOutstandingSFTM returns the amount of sFTM tokens for the delegation
	// identified by the delegator address and the staker id.
	DelegationOutstandingSFTM(*common.Address, *hexutil.Big) (*hexutil.Big, error)

	// DelegationTokenizerUnlocked returns the status of SFC Tokenizer lock
	// for a delegation identified by the address and staker id.
	DelegationTokenizerUnlocked(*common.Address, *hexutil.Big) (bool, error)

	// DelegationFluidStakingActive signals if the delegation is upgraded to Fluid Staking model.
	DelegationFluidStakingActive(*common.Address, *hexutil.Big) (bool, error)

	// StoreWithdrawRequest stores the given withdraw request in persistent storage.
	StoreWithdrawRequest(*types.WithdrawRequest) error

	// UpdateWithdrawRequest stores the updated withdraw request in persistent storage.
	UpdateWithdrawRequest(*types.WithdrawRequest) error

	// WithdrawRequest extracts details of a withdraw request specified by the delegator, validator and request ID.
	WithdrawRequest(*common.Address, *hexutil.Big, *hexutil.Big) (*types.WithdrawRequest, error)

	// WithdrawRequests extracts a list of withdraw requests for the given address and validator.
	WithdrawRequests(*common.Address, *hexutil.Big, *string, int32, bool) (*types.WithdrawRequestList, error)

	// WithdrawRequestsPendingTotal is the total value of all pending withdrawal requests
	// for the given delegator and target staker ID.
	WithdrawRequestsPendingTotal(*common.Address, *hexutil.Big) (*big.Int, error)

	// StoreRewardClaim stores reward claim record in the persistent repository.
	StoreRewardClaim(*types.RewardClaim) error

	// RewardsClaimed returns the sum of all the claimed rewards
	// for the given delegator address and validator ID.
	RewardsClaimed(adr *common.Address, valId *big.Int, since *int64, until *int64) (*big.Int, error)

	// RewardClaims provides list of reward claims for the given criteria.
	RewardClaims(*common.Address, *big.Int, *string, int32) (*types.RewardClaimsList, error)

	// Price returns a price information for the given target symbol.
	Price(sym string) (types.Price, error)

	// GasPrice provides the raw suggested value for the gas price.
	GasPrice() (hexutil.Big, error)

	// GasPriceExtended provides extended gas price information.
	GasPriceExtended() (*types.GasPrice, error)

	// StoreGasPricePeriod stores gas price period data into the persistent storage.
	StoreGasPricePeriod(*types.GasPricePeriod) error

	// GasEstimate calculates the estimated amount of Gas required to perform
	// transaction described by the input params.
	GasEstimate(*struct {
		From  *common.Address
		To    *common.Address
		Value *hexutil.Big
		Data  *string
	}) (*hexutil.Uint64, error)

	// DefiConfiguration loads the current DeFi contract settings.
	DefiConfiguration() (*types.DefiSettings, error)

	// DefiTokens resolves list of DeFi tokens available for the DeFi functions.
	DefiTokens() ([]types.DefiToken, error)

	// DefiToken loads details of a single DeFi token by it's address.
	DefiToken(*common.Address) (*types.DefiToken, error)

	// DefiTokenPrice loads the current price of the given token
	// from on-chain price oracle.
	DefiTokenPrice(*common.Address) (hexutil.Big, error)

	// FMintAccount loads details of a DeFi/fMint account identified by the owner address.
	FMintAccount(common.Address) (*types.FMintAccount, error)

	// FMintTokenBalance loads balance of a single DeFi token by it's address.
	FMintTokenBalance(*common.Address, *common.Address, types.DefiTokenType) (hexutil.Big, error)

	// FMintTokenTotalBalance loads total balance of a single DeFi token by it's address.
	FMintTokenTotalBalance(*common.Address, types.DefiTokenType) (hexutil.Big, error)

	// FMintTokenValue loads value of a single DeFi token by it's address in fUSD.
	FMintTokenValue(*common.Address, *common.Address, types.DefiTokenType) (hexutil.Big, error)

	// FMintRewardsEarned resolves the total amount of rewards
	// accumulated on the account for the excessive collateral deposits.
	FMintRewardsEarned(*common.Address) (hexutil.Big, error)

	// FMintRewardsStashed represents the total amount of rewards
	// accumulated on the account in stash.
	FMintRewardsStashed(*common.Address) (hexutil.Big, error)

	// FMintCanClaimRewards resolves the fMint account flag for being allowed
	// to claim earned rewards.
	FMintCanClaimRewards(*common.Address) (bool, error)

	// FMintCanReceiveRewards resolves the fMint account flag for being eligible
	// to receive earned rewards. If the collateral to debt ration drop below
	// certain value, earned rewards are burned.
	FMintCanReceiveRewards(*common.Address) (bool, error)

	// FMintCanPushRewards signals if there are any rewards unlocked
	// on the rewards' distribution contract and can be pushed to account.
	FMintCanPushRewards() (bool, error)

	// FMintUsers loads the list of fMint users and their associated tokens used for a specified purpose.
	FMintUsers(int32) ([]*types.FMintUserTokens, error)

	// AddFMintTransaction adds the specified fMint transaction to persistent storage.
	AddFMintTransaction(*types.FMintTransaction) error

	// UniswapPairs returns list of all token pairs managed by Uniswap core.
	UniswapPairs() ([]common.Address, error)

	// UniswapKnownPairs returns list of all known and whitelisted token pairs managed by Uniswap core.
	UniswapKnownPairs() ([]common.Address, error)

	// UniswapPair returns an address of an Uniswap pair for the given tokens.
	UniswapPair(*common.Address, *common.Address) (*common.Address, error)

	// UniswapAmountsOut resolves a list of output amounts for the given
	// input amount and a list of tokens to be used to make the swap operation.
	UniswapAmountsOut(amountIn hexutil.Big, tokens []common.Address) ([]hexutil.Big, error)

	// UniswapAmountsIn resolves a list of input amounts for the given
	// output amount and a list of tokens to be used to make the swap operation.
	UniswapAmountsIn(amountOut hexutil.Big, tokens []common.Address) ([]hexutil.Big, error)

	// UniswapQuoteInput calculates optimal input on sibling token based on input amount and
	// self reserves of the analyzed token.
	UniswapQuoteInput(amountIn hexutil.Big, reserveMy hexutil.Big, reserveSibling hexutil.Big) (hexutil.Big, error)

	// UniswapTokens returns list of addresses of tokens involved in an Uniswap pair.
	UniswapTokens(*common.Address) ([]common.Address, error)

	// UniswapReserves returns list of token reserve amounts in a Uniswap pair.
	UniswapReserves(*common.Address) ([]hexutil.Big, error)

	// UniswapReservesTimeStamp returns the timestamp of the reserves of a Uniswap pair.
	UniswapReservesTimeStamp(*common.Address) (hexutil.Uint64, error)

	// UniswapCumulativePrices returns list of token cumulative prices of a Uniswap pair.
	UniswapCumulativePrices(*common.Address) ([]hexutil.Big, error)

	// UniswapLastKValue returns the last value of the pool control coefficient.
	UniswapLastKValue(*common.Address) (hexutil.Big, error)

	// UniswapPairContract returns instance of this contract according to given pair address
	UniswapPairContract(*common.Address) (*contracts.UniswapPair, error)

	// UniswapAdd adds a new incoming swap from blockchain to the repository.
	UniswapAdd(*types.Swap) error

	// LastKnownSwapBlock returns number of the last block known to the repository with swap event.
	LastKnownSwapBlock() (uint64, error)

	// UniswapUpdateLastKnownSwapBlock stores a last correctly saved swap block number into persistent storage.
	UniswapUpdateLastKnownSwapBlock(blkNumber uint64) error

	// UniswapFactoryContract returns an instance of an Uniswap factory
	UniswapFactoryContract() (*contracts.UniswapFactory, error)

	// UniswapVolume returns swap volume for specified uniswap pair
	UniswapVolume(*common.Address, int64, int64) (types.DefiSwapVolume, error)

	// UniswapTimeVolumes returns grouped volumes for specified pair, time and resolution
	UniswapTimeVolumes(*common.Address, string, int64, int64) ([]types.DefiSwapVolume, error)

	// UniswapTimePrices returns grouped prices for specified pair, time and resolution
	UniswapTimePrices(*common.Address, string, int64, int64, int32) ([]types.DefiTimePrice, error)

	// UniswapTimeReserves returns grouped reserves for specified pair, time and resolution
	UniswapTimeReserves(*common.Address, string, int64, int64) ([]types.DefiTimeReserve, error)

	// UniswapActions provides list of uniswap actions stored in the persistent db.
	UniswapActions(*common.Address, *string, int32, int32) (*types.UniswapActionList, error)

	// NativeTokenAddress returns address of the native token wrapper, if available.
	NativeTokenAddress() (*common.Address, error)

	// TokenTransactions provides list of ERC20/ERC721/ERC1155 transactions based on given filters.
	TokenTransactions(tokenType string, token *common.Address, tokenId *big.Int, acc *common.Address, txType []int32, cursor *string, count int32) (*types.TokenTransactionList, error)

	// TokenTransactionsByCall provides a list of token transaction made inside a specific
	// transaction call (blockchain transaction).
	TokenTransactionsByCall(*common.Hash) ([]*types.TokenTransaction, error)

	// Erc20Token returns an ERC20 token for the given address, if available.
	Erc20Token(*common.Address) (*types.Erc20Token, error)

	// Erc20TokensList returns a list of known ERC20 tokens ordered by their activity.
	Erc20TokensList(int32) ([]common.Address, error)

	// Erc20Assets provides list of ERC20 tokens involved with the given owner.
	Erc20Assets(common.Address, int32) ([]common.Address, error)

	// Erc20BalanceOf load the current available balance of and ERC20 token identified by the token
	// contract address for an identified owner address.
	Erc20BalanceOf(*common.Address, *common.Address) (hexutil.Big, error)

	// Erc20Allowance loads the current amount of ERC20 tokens unlocked for DeFi
	// contract by the token owner.
	Erc20Allowance(*common.Address, *common.Address, *common.Address) (hexutil.Big, error)

	// Erc20TotalSupply provides information about all available tokens
	Erc20TotalSupply(*common.Address) (hexutil.Big, error)

	// Erc20Name provides information about the name of the ERC20 token.
	Erc20Name(*common.Address) (string, error)

	// Erc20Symbol provides information about the symbol of the ERC20 token.
	Erc20Symbol(*common.Address) (string, error)

	// Erc20Decimals provides information about the decimals of the ERC20 token.
	Erc20Decimals(*common.Address) (int32, error)

	// Erc20LogoURL provides URL address of a logo of the ERC20 token.
	Erc20LogoURL(*common.Address) string

	// StoreTokenTransaction stores ERC20/ERC721/ERC1155 transaction into the repository.
	StoreTokenTransaction(*types.TokenTransaction) error

	// Erc165SupportsInterface provides information about support of the interface by the contract.
	Erc165SupportsInterface(contract *common.Address, interfaceID [4]byte) (bool, error)

	// Erc721Contract returns an ERC721 token for the given address, if available.
	Erc721Contract(*common.Address) (*types.Erc721Contract, error)

	// Erc721ContractsList returns a list of known ERC721 tokens ordered by their activity.
	Erc721ContractsList(int32) ([]common.Address, error)

	// Erc721Name provides information about the name of the ERC721 token.
	Erc721Name(*common.Address) (string, error)

	// Erc721Symbol provides information about the symbol of the ERC721 token.
	Erc721Symbol(*common.Address) (string, error)

	// Erc721TotalSupply provides information about all available tokens.
	Erc721TotalSupply(token *common.Address) (hexutil.Big, error)

	// Erc721BalanceOf provides amount of NFT tokens owned by given owner in given ERC721 contract.
	Erc721BalanceOf(token *common.Address, owner *common.Address) (hexutil.Big, error)

	// Erc721TokenURI provides URI of Metadata JSON Schema of the ERC721 token.
	Erc721TokenURI(token *common.Address, tokenId *big.Int) (string, error)

	// Erc721OwnerOf provides information about NFT token ownership.
	Erc721OwnerOf(token *common.Address, tokenId *big.Int) (common.Address, error)

	// Erc721GetApproved provides information about operator approved to manipulate with the NFT token.
	Erc721GetApproved(token *common.Address, tokenId *big.Int) (common.Address, error)

	// Erc721IsApprovedForAll provides information about operator approved to manipulate with NFT tokens of given owner.
	Erc721IsApprovedForAll(token *common.Address, owner *common.Address, operator *common.Address) (bool, error)

	// Erc1155ContractsList returns a list of known ERC1155 contracts ordered by their activity.
	Erc1155ContractsList(int32) ([]common.Address, error)

	// Erc1155Uri provides URI of Metadata JSON Schema of the token.
	Erc1155Uri(token *common.Address, tokenId *big.Int) (string, error)

	// Erc1155BalanceOf provides amount of NFT tokens owned by given owner.
	Erc1155BalanceOf(token *common.Address, owner *common.Address, tokenId *big.Int) (*big.Int, error)

	// Erc1155BalanceOfBatch provides amount of NFT tokens owned by given owner.
	Erc1155BalanceOfBatch(token *common.Address, owners *[]common.Address, tokenIds []*big.Int) ([]*big.Int, error)

	// Erc1155IsApprovedForAll provides information about operator approved to manipulate with NFT tokens of given owner.
	Erc1155IsApprovedForAll(token *common.Address, owner *common.Address, operator *common.Address) (bool, error)

	// GovernanceContractBy provides governance contract details by its address.
	GovernanceContractBy(common.Address) (config.GovernanceContract, error)

	// GovernanceProposalsCount provides the total number of proposals
	// in a given Governance contract.
	GovernanceProposalsCount(*common.Address) (hexutil.Big, error)

	// GovernanceProposal provides a detail of Proposal of a governance contract
	// specified by its id.
	GovernanceProposal(common.Address, *hexutil.Big) (*types.GovernanceProposal, error)

	// GovernanceProposalState provides a state of Proposal of a governance contract
	// specified by its id.
	GovernanceProposalState(*common.Address, *hexutil.Big) (*types.GovernanceProposalState, error)

	// GovernanceOptionState returns a state of the given option of a proposal.
	GovernanceOptionState(*common.Address, *hexutil.Big, *hexutil.Big) (*types.GovernanceOptionState, error)

	// GovernanceOptionStates returns a list of states of options of a proposal.
	GovernanceOptionStates(*common.Address, *hexutil.Big, int) ([]*types.GovernanceOptionState, error)

	// GovernanceVote provides a single vote in the Governance Proposal context.
	GovernanceVote(*common.Address, *hexutil.Big, *common.Address, *common.Address) (*types.GovernanceVote, error)

	// GovernanceProposals loads list of proposals from given set of Governance contracts.
	GovernanceProposals([]common.Address, *string, int32, bool) (*types.GovernanceProposalList, error)

	// GovernanceProposalFee returns the fee payable for a new proposal
	// in given Governance contract context.
	GovernanceProposalFee(*common.Address) (hexutil.Big, error)

	// GovernanceProposalOwner returns the owner of the given proposal contract.
	GovernanceProposalOwner(*common.Address) (*common.Address, error)

	// GovernanceTotalWeight provides the total weight of all available votes
	// in the governance contract identified by the address.
	GovernanceTotalWeight(*common.Address) (hexutil.Big, error)

	// FLendGetLendingPool resolves lending pool contract instance
	// to be able to get calls and information from this contract
	FLendGetLendingPool() (*contracts.ILendingPool, error)

	// FLendGetLendingPoolReserveData resolves reserve data
	// according to given address
	FLendGetLendingPoolReserveData(*common.Address) (*types.ReserveData, error)

	// FLendGetUserAccountData resolves user account data for
	// specified address
	FLendGetUserAccountData(*common.Address) (*types.FLendUserAccountData, error)

	// FLendGetReserveList resolves list of reserves in lending pool
	FLendGetReserveList() ([]common.Address, error)

	// FLendGetUserDepositHistory resolves deposit history
	// data for specified user and asset address
	FLendGetUserDepositHistory(*common.Address, *common.Address) ([]*types.FLendDeposit, error)

	// TrxFlowVolume resolves the list of daily trx flow aggregations.
	TrxFlowVolume(from *time.Time, to *time.Time) ([]*types.DailyTrxVolume, error)

	// TrxGasSpeed provides speed of gas consumption per second by transactions.
	TrxGasSpeed(from *time.Time, to *time.Time) (float64, error)

	// GasPriceTicks provides a list of gas price ticks for the given time period.
	GasPriceTicks(from *time.Time, to *time.Time) ([]types.GasPricePeriod, error)

	// TrxFlowUpdate executes the trx flow update in the database.
	TrxFlowUpdate()

	// TrxFlowSpeed provides speed of transaction per second for the last <sec> seconds.
	TrxFlowSpeed(sec int32) (float64, error)

	// BurnDailyUpdate update daily aggregate of the burn flow.
	BurnDailyUpdate()

	// StoreFtmBurn stores the given native FTM burn per block record into the persistent storage.
	StoreFtmBurn(burn *types.FtmBurn) error

	// FtmBurnTotal provides the total amount of burned native FTM.
	FtmBurnTotal() (int64, error)

	// FtmBurnList provides list of per-block burned native FTM tokens.
	FtmBurnList(count int64) ([]*types.FtmBurn, error)

	// BurnTreasuryStashShareByBlock finds treasury/burn share for the given block ID.
	BurnTreasuryStashShareByBlock(blk uint64) *BurnTreasuryShare

	// BurnTreasuryStashShareByTimeStamp finds treasury/burn share for the given time stamp.
	BurnTreasuryStashShareByTimeStamp(ts int64) *BurnTreasuryShare

	// FtmTreasuryTotal provides the total amount of native FTM sent into treasury.
	FtmTreasuryTotal() (int64, error)

	// FeeFlow provides a list of fee flow aggregates for the given date range.
	FeeFlow(from, to time.Time) ([]*types.FtmDailyBurn, error)

	// NetworkNode returns instance of Opera network node record by its ID.
	NetworkNode(nid enode.ID) (*types.OperaNode, error)

	// StoreNetworkNode stores the given Opera node record in the persistent database.
	StoreNetworkNode(node *types.OperaNode) error

	// IsNetworkNodeKnown checks if the given network node is already registered in the persistent database.
	IsNetworkNodeKnown(id enode.ID) bool

	// NetworkNodeConfirmCheck confirms successful check of the given Opera network node.
	NetworkNodeConfirmCheck(node *enode.Node, bhp p2p.BlockHeightProvider) (bool, error)

	// NetworkNodeFailCheck registers failed check of the given Opera network node.
	NetworkNodeFailCheck(node *enode.Node) error

	// PeerInformation returns detailed information of the given peer, if it can be obtained.
	PeerInformation(node *enode.Node, bhp p2p.BlockHeightProvider) (*types.OperaNodeInformation, error)

	// NetworkNodeUpdateBatch provides a list of Opera network node addresses most suitable for status update
	// based on the registered time of the latest check.
	NetworkNodeUpdateBatch() ([]*enode.Node, error)

	// NetworkNodeBootstrapSet provides a set of known nodes to be co-used to bootstrap new search.
	NetworkNodeBootstrapSet() []*enode.Node

	// GeoLocation provides geographic location information for the given IP address using GeoIP bridge.
	GeoLocation(net.IP) (types.GeoLocation, error)

	// NetworkNodesGeoAggregated provides a list of aggregated opera nodes based on given location detail level.
	NetworkNodesGeoAggregated(level int) ([]*types.OperaNodeLocationAggregate, error)

	// Close and cleanup the repository.
	Close()
}

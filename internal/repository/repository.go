/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"bytes"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository/cache"
	"fantom-api-graphql/internal/repository/db"
	"fantom-api-graphql/internal/repository/rpc"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ftm "github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

// Repository interface defines functions the underlying implementation provides to API resolvers.
type Repository interface {
	// log provides access to the system wide logger.
	Log() logger.Logger

	// FtmConnection returns open connection to Opera/Lachesis full node.
	FtmConnection() *ftm.Client

	// Account returns account at Opera blockchain for an address, nil if not found.
	Account(*common.Address) (*types.Account, error)

	// AccountBalance returns the current balance of an account at Opera blockchain.
	AccountBalance(*types.Account) (*hexutil.Big, error)

	// AccountNonce returns the current number of sent transactions of an account at Opera blockchain.
	AccountNonce(*types.Account) (*hexutil.Uint64, error)

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
	AccountTransactions(*types.Account, *string, int32) (*types.TransactionHashList, error)

	// Returns total number of accounts known to repository.
	AccountsActive() (hexutil.Uint64, error)

	// AccountIsKnown checks if the account of the given address is known to the API server.
	AccountIsKnown(*common.Address) bool

	// AccountAdd adds specified account detail into the repository.
	AccountAdd(*types.Account) error

	// AccountMarkActivity marks the latest account activity in the repository.
	AccountMarkActivity(*types.Account, uint64) error

	// Block returns a block at Opera blockchain represented by a number.
	// Top block is returned if the number is not provided.
	// If the block is not found, ErrBlockNotFound error is returned.
	BlockByNumber(*hexutil.Uint64) (*types.Block, error)

	// BlockHeight returns the current height of the Opera blockchain in blocks.
	BlockHeight() (*hexutil.Big, error)

	// LastKnownBlock returns number of the last block known to the repository.
	LastKnownBlock() (uint64, error)

	// IsSfcContract returns true if the given address points to the SFC contract.
	IsSfcContract(*common.Address) bool

	// IsStiContract returns true if the given address points to the STI contract.
	IsStiContract(*common.Address) bool

	// CurrentEpoch returns the id of the current epoch.
	CurrentEpoch() (hexutil.Uint64, error)

	// CurrentSealedEpoch returns the data of the latest sealed epoch.
	CurrentSealedEpoch() (*types.Epoch, error)

	// Epoch returns the id of the current epoch.
	Epoch(hexutil.Uint64) (types.Epoch, error)

	// Block returns a block at Opera blockchain represented by a hash.
	// Top block is returned if the hash is not provided.
	// If the block is not found, ErrBlockNotFound error is returned.
	BlockByHash(*types.Hash) (*types.Block, error)

	// AddTransaction adds a new incoming transaction from blockchain to the repository.
	TransactionAdd(*types.Block, *types.Transaction) error

	// TransactionUpdate modifies a transaction record in the repository.
	TransactionUpdate(*types.Transaction) error

	// Transaction returns a transaction at Opera blockchain by a hash, nil if not found.
	Transaction(*types.Hash) (*types.Transaction, error)

	// TransactionsCount returns total number of transactions in the block chain.
	TransactionsCount() (hexutil.Uint64, error)

	// Transactions returns list of transaction hashes at Opera blockchain.
	Transactions(*string, int32) (*types.TransactionHashList, error)

	// Collection pulls list of blocks starting on the specified block number
	// and going up, or down based on count number.
	Blocks(*uint64, int32) (*types.BlockList, error)

	// LastStakerId returns the last staker id in Opera blockchain.
	LastStakerId() (hexutil.Uint64, error)

	// StakersNum returns the number of stakers in Opera blockchain.
	StakersNum() (hexutil.Uint64, error)

	// IsStaker returns if the given address is an SFC staker.
	IsStaker(*common.Address) (bool, error)

	// Staker extract a staker information from SFC smart contract.
	Staker(hexutil.Uint64) (*types.Staker, error)

	// StakerAddress extract a staker address for the given staker ID.
	StakerAddress(hexutil.Uint64) (common.Address, error)

	// Staker extract a staker information by address.
	StakerByAddress(common.Address) (*types.Staker, error)

	// TotalStaked calculates current total staked amount for all stakers.
	TotalStaked() (*hexutil.Big, error)

	// StakerInfo extracts an extended staker information from smart contact.
	PullStakerInfo(hexutil.Uint64) (*types.StakerInfo, error)

	// StoreStakerInfo stores staker information to in-memory cache for future use.
	StoreStakerInfo(hexutil.Uint64, types.StakerInfo) error

	// RetrieveStakerInfo gets staker information from in-memory if available.
	RetrieveStakerInfo(hexutil.Uint64) *types.StakerInfo

	// IsDelegating returns if the given address is an SFC delegator.
	IsDelegating(*common.Address) (bool, error)

	// Delegation returns a detail of delegation for the given address.
	Delegation(common.Address, hexutil.Uint64) (*types.Delegation, error)

	// DelegationsByAddress returns a list of all delegations
	// of a given delegator address.
	DelegationsByAddress(common.Address) ([]types.Delegation, error)

	// DelegationsOf extracts a list of delegations for a given staker.
	DelegationsOf(hexutil.Uint64) ([]types.Delegation, error)

	// DelegationLock returns delegation lock information using SFC contract binding.
	DelegationLock(*types.Delegation) (*types.DelegationLock, error)

	// Delegation returns a detail of delegation for the given address.
	DelegationRewards(string, hexutil.Uint64) (types.PendingRewards, error)

	// DelegationOutstandingSFTM returns the amount of sFTM tokens for the delegation
	// identified by the delegator address and the staker id.
	DelegationOutstandingSFTM(*common.Address, *hexutil.Uint64) (hexutil.Big, error)

	// DelegationTokenizerUnlocked returns the status of SFC Tokenizer lock
	// for a delegation identified by the address and staker id.
	DelegationTokenizerUnlocked(*common.Address, *hexutil.Uint64) (bool, error)

	// WithdrawRequests extracts a list of partial withdraw requests
	// for the given address.
	WithdrawRequests(*common.Address, *hexutil.Uint64) ([]*types.WithdrawRequest, error)

	// DeactivatedDelegation extracts a list of deactivated delegation requests
	// for the given address.
	DeactivatedDelegation(*common.Address, *hexutil.Uint64) ([]*types.DeactivatedDelegation, error)

	// SfcVersion returns current version of the SFC contract.
	SfcVersion() (hexutil.Uint64, error)

	// LockingAllowed indicates if the stake locking has been enabled in SFC.
	LockingAllowed() (bool, error)

	// RewardsAllowed returns the reward lock status from SFC.
	RewardsAllowed() (bool, error)

	// RewardsStash returns the amount of WEI stashed for the given address.
	RewardsStash(*common.Address) (*big.Int, error)

	// delegatedAmount calculates total amount currently delegated
	// and amount locked in pending un-delegation.
	// Partial Un-delegations are subtracted during the preparation
	// phase, but total un-delegations are subtracted only when
	// the delegation is closed.
	DelegatedAmountExtended(*types.Delegation) (*big.Int, *big.Int, error)

	// DelegationFluidStakingActive signals if the delegation is upgraded to Fluid Staking model.
	DelegationFluidStakingActive(*types.Delegation) (bool, error)

	// DelegationPaidUntilEpoch resolves the id of the last epoch rewards has been paid to."
	DelegationPaidUntilEpoch(*types.Delegation) (hexutil.Uint64, error)

	// Price returns a price information for the given target symbol.
	Price(sym string) (types.Price, error)

	// GasPrice resolves the current amount of WEI for single Gas.
	GasPrice() (hexutil.Uint64, error)

	// SendTransaction sends raw signed and RLP encoded transaction to the block chain.
	SendTransaction(hexutil.Bytes) (*types.Transaction, error)

	// SetBlockChannel registers a channel for notifying new block events.
	SetBlockChannel(chan *types.Block)

	// SetTrxChannel registers a channel for notifying new transaction events.
	SetTrxChannel(chan *types.Transaction)

	// Contract extract a smart contract information by address if available.
	Contract(*common.Address) (*types.Contract, error)

	// ContractAdd adds new contract into the repository.
	ContractAdd(*types.Contract) error

	// Contracts returns list of smart contracts at Opera blockchain.
	Contracts(bool, *string, int32) (*types.ContractList, error)

	// ValidateContract tries to validate contract byte code using
	// provided source code. If successful, the contract information
	// is updated the the repository.
	ValidateContract(*types.Contract) error

	// Ballots returns list of ballots at Opera blockchain.
	Ballots(*string, int32) (*types.BallotList, error)

	// BallotsClosed returns a list of <count> recently closed Ballots.
	// If the finalized is set to false, the list contains Ballots waiting
	// to be resolved.
	BallotsClosed(bool, uint32) ([]types.Ballot, error)

	// BallotsActive returns a list of at most <count> currently active Ballots.
	BallotsActive(uint32) ([]types.Ballot, error)

	// BallotByAddress returns a ballot information by the contract address.
	BallotByAddress(*common.Address) (*types.Ballot, error)

	// BallotIsFinalized returns the finalized status of a ballot.
	BallotIsFinalized(*common.Address) (bool, error)

	// BallotWinner returns the winning proposal index, or nil if not decided.
	BallotWinner(*common.Address) (*uint64, error)

	// Votes returns a list of votes for the given votes and list of ballots.
	Votes(common.Address, []common.Address) ([]types.Vote, error)

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
	// on the rewards distribution contract and can be pushed to accounts.
	FMintCanPushRewards() (bool, error)

	// UniswapPairs returns list of all token pairs managed by Uniswap core.
	UniswapPairs() ([]common.Address, error)

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

	// UniswapTokens returns list of addresses of tokens involved in a Uniswap pair.
	UniswapTokens(*common.Address) ([]common.Address, error)

	// UniswapReserves returns list of token reserve amounts in a Uniswap pair.
	UniswapReserves(*common.Address) ([]hexutil.Big, error)

	// UniswapReservesTimeStamp returns the timestamp of the reserves of a Uniswap pair.
	UniswapReservesTimeStamp(*common.Address) (hexutil.Uint64, error)

	// UniswapCumulativePrices returns list of token cumulative prices of a Uniswap pair.
	UniswapCumulativePrices(*common.Address) ([]hexutil.Big, error)

	// UniswapLastKValue returns the last value of the pool control coefficient.
	UniswapLastKValue(*common.Address) (hexutil.Big, error)

	// NativeTokenAddress returns address of the native token wrapper, if available.
	NativeTokenAddress() (*common.Address, error)

	// Erc20Token returns an ERC20 token rfor the given address, if available.
	Erc20Token(*common.Address) (*types.Erc20Token, error)

	// Erc20TokensList returns a list of known ERC20 tokens ordered by their activity.
	Erc20TokensList(int32) ([]common.Address, error)

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

	// GovernanceContractBy provides governance contract details by its address.
	GovernanceContractBy(*common.Address) (*config.GovernanceContract, error)

	// GovernanceProposalsCount provides the total number of prpoposals
	// in a given Governance contract.
	GovernanceProposalsCount(*common.Address) (hexutil.Big, error)

	// GovernanceProposal provides a detail of Proposal of a governance contract
	// specified by its id.
	GovernanceProposal(*common.Address, *hexutil.Big) (*types.GovernanceProposal, error)

	// GovernanceProposalState provides a state of Proposal of a governance contract
	// specified by its id.
	GovernanceProposalState(*common.Address, *hexutil.Big) (*types.GovernanceProposalState, error)

	// GovernanceOptionState returns a state of the given option of a proposal.
	GovernanceOptionState(*common.Address, *hexutil.Big, *hexutil.Big) (*types.GovernanceOptionState, error)

	// GovernanceOptionStates returns a list of states of options of a proposal.
	GovernanceOptionStates(*common.Address, *hexutil.Big) ([]*types.GovernanceOptionState, error)

	// GovernanceVote provides a single vote in the Governance Proposal context.
	GovernanceVote(*common.Address, *hexutil.Big, *common.Address, *common.Address) (*types.GovernanceVote, error)

	// GovernanceProposals loads list of proposals from given set of Governance contracts.
	GovernanceProposals([]*common.Address, *string, int32, bool) (*types.GovernanceProposalList, error)

	// GovernanceProposalFee returns the fee payable for a new proposal
	// in given Governance contract context.
	GovernanceProposalFee(*common.Address) (hexutil.Big, error)

	// GovernanceTotalWeight provides the total weight of all available votes
	// in the governance contract identified by the address.
	GovernanceTotalWeight(*common.Address) (hexutil.Big, error)

	// Close and cleanup the repository.
	Close()
}

// Proxy represents Repository interface implementation and controls access to data
// trough several low level bridges.
type proxy struct {
	cache *cache.MemBridge
	db    *db.MongoDbBridge
	rpc   *rpc.FtmBridge
	log   logger.Logger
	cfg   *config.Config

	// governance contracts reference
	govContracts map[string]*config.GovernanceContract

	// smart contract compilers
	solCompiler string

	// official ballot source addresses
	ballotSources []string

	// service orchestrator reference
	orc *orchestrator
}

// New creates new instance of Repository implementation, namely proxy structure.
func New(cfg *config.Config, log logger.Logger) (Repository, error) {
	// create new in-memory cache bridge
	caBridge, dbBridge, rpcBridge, err := connect(cfg, log)
	if err != nil {
		log.Criticalf("repository init failed")
		return nil, err
	}

	// construct the proxy instance
	p := proxy{
		cache: caBridge,
		db:    dbBridge,
		rpc:   rpcBridge,
		log:   log,
		cfg:   cfg,

		govContracts: governanceContractsMap(&cfg.Governance),

		// keep reference to the SOL compiler
		solCompiler: cfg.Compiler.DefaultSolCompilerPath,

		// keep the ballot sources ref
		ballotSources: cfg.Voting.Sources,
	}

	// make the service orchestrator
	p.orc = newOrchestrator(&p, log, &cfg.Repository)

	// return the proxy
	return &p, nil
}

// governanceContractsMap creates map of governance contracts keyed
// by the contract address.
func governanceContractsMap(cfg *config.Governance) map[string]*config.GovernanceContract {
	// prep the result set
	res := make(map[string]*config.GovernanceContract)

	// collect all the configured governance contracts into the map
	for _, gv := range cfg.Contracts {
		res[gv.Address.String()] = &gv
	}
	return res
}

// connect opens connections to the external sources we need.
func connect(cfg *config.Config, log logger.Logger) (*cache.MemBridge, *db.MongoDbBridge, *rpc.FtmBridge, error) {
	// create new in-memory cache bridge
	caBridge, err := cache.New(cfg, log)
	if err != nil {
		log.Criticalf("can not create in-memory cache bridge, %s", err.Error())
		return nil, nil, nil, err
	}

	// create new database connection bridge
	dbBridge, err := db.New(cfg, log)
	if err != nil {
		log.Criticalf("can not connect backend persistent storage, %s", err.Error())
		return nil, nil, nil, err
	}

	// create new Lachesis RPC bridge
	rpcBridge, err := rpc.New(cfg, log)
	if err != nil {
		log.Criticalf("can not connect Lachesis RPC interface, %s", err.Error())
		return nil, nil, nil, err
	}

	// try to validate the solidity compiler by asking for it's version
	if _, err := compiler.SolidityVersion(cfg.Compiler.DefaultSolCompilerPath); err != nil {
		log.Criticalf("can not invoke the Solidity compiler, %s", err.Error())
		// return nil, nil, nil, err
	}

	return caBridge, dbBridge, rpcBridge, nil
}

// Close with close all connections and clean up the pending work for graceful termination.
func (p *proxy) Close() {
	// inform about actions
	p.log.Notice("repository is closing")

	// initiate orchestrator closing process
	p.orc.close()

	// close connections
	p.db.Close()
	p.rpc.Close()

	// inform about actions
	p.log.Notice("repository done")
}

// FtmClient returns open connection to Opera/Lachesis full node.
func (p *proxy) Log() logger.Logger {
	return p.log
}

// FtmClient returns open connection to Opera/Lachesis full node.
func (p *proxy) FtmConnection() *ftm.Client {
	return p.rpc.Connection()
}

// SetBlockChannel registers a channel for notifying new block events.
func (p *proxy) SetBlockChannel(ch chan *types.Block) {
	p.orc.setBlockChannel(ch)
}

// SetTrxChannel registers a channel for notifying new transactions events.
func (p *proxy) SetTrxChannel(ch chan *types.Transaction) {
	p.orc.setTrxChannel(ch)
}

// IsSfcContract returns true if the given address points to the SFC contract.
func (p *proxy) IsSfcContract(addr *common.Address) bool {
	return bytes.Equal(addr.Bytes(), p.cfg.Staking.SFCContract.Bytes())
}

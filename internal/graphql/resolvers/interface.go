// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"context"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ApiResolver represents the API interface expected to handle API access points
type ApiResolver interface {
	// State resolves current state of the blockchain.
	State() (CurrentState, error)

	// SfcConfig resolves the current SFC configuration.
	SfcConfig() SfcConfig

	// Version resolves current version of the API server.
	Version() string

	// Epochs resolves a list of epochs for the given cursor and count.
	Epochs(args struct {
		Cursor *Cursor
		Count  int32
	}) (*EpochList, error)

	// Account resolves blockchain account by address.
	Account(struct{ Address common.Address }) (*Account, error)

	// Contracts resolves list of blockchain smart contracts encapsulated in a listable structure.
	Contracts(*struct {
		ValidatedOnly bool
		Cursor        *Cursor
		Count         int32
	}) (*ContractList, error)

	// ValidateContract resolves smart contract source code vs. deployed byte code and marks
	// the contract as validated if the match is found. Peer API points are ringed on success
	// to notify them about the change.
	ValidateContract(*struct{ Contract ContractValidationInput }) (*Contract, error)

	// Block resolves blockchain block by number or by hash. If neither is provided, the most recent block is given.
	Block(*struct {
		Number *hexutil.Uint64
		Hash   *common.Hash
	}) (*Block, error)

	// Blocks resolves list of blockchain blocks encapsulated in a listable structure.
	Blocks(*struct {
		Cursor *Cursor
		Count  int32
	}) (*BlockList, error)

	// Transaction resolves blockchain transaction by hash.
	Transaction(*struct{ Hash common.Hash }) (*Transaction, error)

	// Transactions resolves list of blockchain transactions encapsulated in a listable structure.
	Transactions(*struct {
		Cursor *Cursor
		Count  int32
	}) (*TransactionList, error)

	// OnBlock resolves subscription to new blocks' event broadcast.
	OnBlock(ctx context.Context) <-chan *Block

	// OnTransaction resolves subscription to new transactions' event broadcast.
	OnTransaction(ctx context.Context) <-chan *Transaction

	// CurrentEpoch resolves id of the current epoch.
	CurrentEpoch() (hexutil.Uint64, error)

	// Epoch resolves information about epoch of the given id.
	Epoch(*struct{ Id *hexutil.Uint64 }) (Epoch, error)

	// LastStakerId resolves the last staker id in Opera blockchain.
	LastStakerId() (hexutil.Uint64, error)

	// StakersNum resolves the number of stakers in Opera blockchain.
	StakersNum() (hexutil.Uint64, error)

	// Staker resolves a staker information from SFC smart contract.
	Staker(struct {
		Id      *hexutil.Big
		Address *common.Address
	}) (*Staker, error)

	// Stakers resolves a list of staker information from SFC smart contract.
	Stakers() ([]*Staker, error)

	// Delegation resolves details of a delegator by its address.
	Delegation(*struct {
		Address common.Address
		Staker  hexutil.Big
	}) (*Delegation, error)

	// DelegationsOf a list of delegations information of a staker.
	DelegationsOf(*struct {
		Staker hexutil.Big
		Cursor *Cursor
		Count  int32
	}) (*DelegationList, error)

	// DelegationsByAddress a list of own delegations by the account address.
	DelegationsByAddress(*struct {
		Address common.Address
		Cursor  *Cursor
		Count   int32
	}) (*DelegationList, error)

	// Price resolves price details of the Opera blockchain token for the given target symbols.
	Price(*struct{ To string }) (types.Price, error)

	// GasPrice resolves the current amount of WEI for single Gas.
	GasPrice() (hexutil.Uint64, error)

	// EstimateGas resolves the estimated amount of Gas required to perform
	// transaction described by the input params.
	EstimateGas(struct {
		From  *common.Address
		To    *common.Address
		Value *hexutil.Big
		Data  *string
	}) (*hexutil.Uint64, error)

	// EstimateRewards resolves reward estimation for the given address or amount staked.
	EstimateRewards(*struct {
		Address *common.Address
		Amount  *hexutil.Uint64
	}) (EstimatedRewards, error)

	// SfcRewardsCollectedAmount resolves the amount of collected rewards
	// based on provided filtering criteria.
	SfcRewardsCollectedAmount(struct {
		Delegator *common.Address
		Staker    *hexutil.Big
		Since     *hexutil.Uint64
		Until     *hexutil.Uint64
	}) (hexutil.Big, error)

	// SendTransaction sends raw signed and RLP encoded transaction to the blockchain.
	SendTransaction(*struct{ Tx hexutil.Bytes }) (*Transaction, error)

	// DefiConfiguration resolves the current DeFi contract settings.
	DefiConfiguration() (*DefiConfiguration, error)

	// DefiTokens resolves list of DeFi tokens available for the DeFi functions.
	DefiTokens() ([]*DefiToken, error)

	// DefiUniswapPairs resolves a list of all pairs managed by the Uniswap core.
	DefiUniswapPairs() []*UniswapPair

	// DefiUniswapAmountsOut resolves a list of output amounts for the given
	// input amount and a list of tokens to be used to make the swap operation.
	DefiUniswapAmountsOut(*struct {
		AmountIn hexutil.Big
		Tokens   []common.Address
	}) ([]hexutil.Big, error)

	// DefiUniswapAmountsIn resolves a list of input amounts for the given
	// output amount and a list of tokens to be used to make the swap operation.
	DefiUniswapAmountsIn(*struct {
		AmountOut hexutil.Big
		Tokens    []common.Address
	}) ([]hexutil.Big, error)

	// DefiUniswapQuoteLiquidity resolves a list of optimal amounts of tokens
	// to be added to both sides of a pair on addLiquidity call.
	DefiUniswapQuoteLiquidity(*struct {
		Tokens    []common.Address
		AmountsIn []hexutil.Big
	}) ([]hexutil.Big, error)

	// FMintAccount resolves details of a specified DeFi account.
	FMintAccount(*struct{ Owner common.Address }) (*FMintAccount, error)

	// FMintTokenAllowance resolves the amount of ERC20 tokens unlocked
	// by the token owner for DeFi/fMint protocol operations.
	FMintTokenAllowance(args *struct {
		Owner common.Address
		Token common.Address
	}) hexutil.Big

	// Erc20Token resolves an instance of ERC20 token if available.
	Erc20Token(*struct{ Token common.Address }) *ERC20Token

	// Erc20TokenList resolves a list of instances of ERC20 tokens.
	Erc20TokenList(struct{ Count int32 }) ([]*ERC20Token, error)

	// Erc20Assets resolves a list of instances of ERC20 tokens for the given owner.
	Erc20Assets(struct {
		Owner common.Address
		Count int32
	}) ([]*ERC20Token, error)

	// ErcTokenBalance resolves the current available balance of the specified token
	// for the specified owner.
	ErcTokenBalance(args *struct {
		Owner common.Address
		Token common.Address
	}) hexutil.Big

	// ErcTotalSupply resolves the current total supply of the specified token.
	ErcTotalSupply(args *struct{ Token common.Address }) hexutil.Big

	// ErcTokenAllowance resolves the current amount of ERC20 tokens unlocked
	// by the token owner for the spender to be manipulated with.
	ErcTokenAllowance(args *struct {
		Token   common.Address
		Owner   common.Address
		Spender common.Address
	}) hexutil.Big

	// GovContracts resolves list of governance contracts details recognized by the API.
	GovContracts() ([]*GovernanceContract, error)

	// GovContract provides a specific Governance contract information by its address.
	GovContract(struct{ Address common.Address }) (*GovernanceContract, error)

	// GovProposals represents list of joined proposals across all the Governance contracts.
	GovProposals(struct {
		Cursor     *Cursor
		Count      int32
		ActiveOnly bool
	}) (*GovernanceProposalList, error)

	// TrxVolume resolves list of daily aggregations
	// of the network transaction flow.
	TrxVolume(args struct {
		From *string
		To   *string
	}) ([]*DailyTrxVolume, error)

	// TrxSpeed resolves the recent speed of the network in transactions processed per second.
	TrxSpeed(args struct {
		Range int32
	}) (float64, error)

	// TrxGasSpeed resolves the gas consumption speed
	// of the network in transactions processed per second.
	TrxGasSpeed(args struct {
		Range int32
		To    *string
	}) (float64, error)

	// Close terminates resolver broadcast management.
	Close()
}

package gqlschema

// Auto generated GraphQL schema bundle
const schema = `
# DefiToken represents a token available for DeFi operations.
type DefiToken {
    # address of the token is used as the token's unique identifier.
    address: Address!

    # name of the token.
    name: String!

    # symbol used as an abbreviation for the token.
    symbol: String!

    # logoUrl is the URL of the token logo image.
    logoUrl: String!

    # decimals is the number of decimals the token supports.
    # The most common value is 18 to mimic the ETH to WEI relationship.
    decimals: Int!

    # isActive signals if the token can be used
    # in the DeFi functions at all.
    isActive: Boolean!

    # canWrapFTM signals if the token can be used
    # to wrap native FTM tokens for DeFi trading.
    canWrapFTM: Boolean!

    # canDeposit signals if the token can be used
    # in deposit as a collateral asset.
    canDeposit: Boolean!

    # canMint signals if the token can be used
    # in fMint protocol as the target token.
    canMint: Boolean!

    # canBorrow signals if the token is available
    # for FLend borrow operations.
    canBorrow: Boolean!

    # canTrade signals if the token is available
    # for FTrade direct trading operations.
    canTrade: Boolean!

    # price represents the value of the token in ref. denomination.
    # We use fUSD tokens as the synth reference value.
    price: BigInt!

    # priceDecimals is the number of decimals used on the price
    # field to properly handle value calculations without loosing precision.
    priceDecimals: Int!

    # availableBalance represents the total available balance of the token
    # on the account regardless of the DeFi usage of the token.
    # It's effectively the amount available held by the ERC20 token
    # on the account behalf.
    availableBalance(owner: Address!): BigInt!

    # defiAllowance represents the amount of ERC20 tokens unlocked
    # by the owner / token holder to be accessible for DeFi operations.
    # If an operation requires access to certain ERC20 token, the DeFi
    # contract must be allowed to make a transfer of required amount
    # of tokens from the owner to the DeFi Liquidity Poll.
    # If it's not given, the operation will fail.
    allowance(owner: Address!): BigInt!

    # totalSupply represents total amount of tokens across all accounts
    totalSupply: BigInt!
}

# DefiTokenBalanceType represents the type of DeFi token balance record.
enum DefiTokenBalanceType {
    COLLATERAL
    DEBT
}

# Price represents price information of core Opera token
type Price {
    "Source unit symbol."
    fromSymbol: String!

    "Target unit symbol."
    toSymbol: String!

    "Price of the source symbol unit in target symbol unit."
    price: Float!

    "Price change in last 24h."
    change24: Float!

    "Price change in percent in last 24h."
    changePct24: Float!

    "Open 24h price."
    open24: Float!

    "Highest 24h price."
    high24: Float!

    "Lowest 24h price."
    low24: Float!

    "Volume exchanged in last 24h price."
    volume24: Float!

    "Market cap of the source unit."
    marketCap: Float!

    "Timestamp of the last update of this price value."
    lastUpdate: Long!
}

# TransactionList is a list of transaction edges provided by sequential access request.
type TransactionList {
    # Edges contains provided edges of the sequential list.
    edges: [TransactionListEdge!]!

    # TotalCount is the maximum number of transactions available for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page of transaction edges.
    pageInfo: ListPageInfo!
}

# TransactionListEdge is a single edge in a sequential list of transactions.
type TransactionListEdge {
    cursor: Cursor!
    transaction: Transaction!
}


# BlockList is a list of block edges provided by sequential access request.
type BlockList {
    # Edges contains provided edges of the sequential list.
    edges: [BlockListEdge!]!

    # TotalCount is the maximum number of blocks available for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page of block edges.
    pageInfo: ListPageInfo!
}

# BlockListEdge is a single edge in a sequential list of blocks.
type BlockListEdge {
    cursor: Cursor!
    block: Block!
}

# StakerInfo represents extended staker information from smart contract.
type StakerInfo {
    "Name represents the name of the staker."
    name: String

    "LogoUrl represents staker logo URL."
    logoUrl: String

    "Website represents a link to stakers website."
    website: String

    "Contact represents a link to contact to the staker."
    contact: String
}
# ListPageInfo contains information about a sequential access list page.
type ListPageInfo {
    # First is the cursor of the first edge of the edges list. null for empty list.
    first: Cursor

    # Last if the cursor of the last edge of the edges list. null for empty list.
    last: Cursor

    # HasNext specifies if there is another edge after the last one.
    hasNext: Boolean!

    # HasNext specifies if there is another edge before the first one.
    hasPrevious: Boolean!
}
# Transaction is an Opera block chain transaction.
type Transaction {
    # Hash is the unique hash of this transaction.
    hash: Hash!

    # Nonce is the number of transactions sent by the account prior to this transaction.
    nonce: Long!

    # Index is the index of this transaction in the block. This will
    # be null if the transaction is in a pending pool.
    index: Long

    # From is the address of the account that sent this transaction
    from: Address!

    # Sender is the account that sent this transaction
    sender: Account!

    # To is the account the transaction was sent to.
    # This is null for contract creating transactions.
    to: Address

    # contractAddress represents the address of smart contract
    # deployed by this transaction;
    # null if the transaction is not contract creation
    contractAddress: Address

    # Recipient is the account that received this transaction.
    # Null for contract creating transaction.
    recipient: Account

    # Value is the value sent along with this transaction in WEI.
    value: BigInt!

    # GasPrice is the price of gas per unit in WEI.
    gasPrice: BigInt!

    # Gas represents gas provided by the sender.
    gas: Long!

    # GasUsed is the amount of gas that was used on processing this transaction.
    # If the transaction is pending, this field will be null.
    gasUsed: Long

    # InputData is the data supplied to the target of the transaction.
    # Contains smart contract byte code if this is contract creation.
    # Contains encoded contract state mutating function call if recipient
    # is a contract address.
    inputData: Bytes!

    # BlockHash is the hash of the block this transaction was assigned to.
    # Null if the transaction is pending.
    blockHash: Hash

    # BlockHash is the hash of the block this transaction was assigned to.
    # Null if the transaction is pending.
    blockNumber: Long

    # Block is the block this transaction was assigned to. This will be null if
    # the transaction is pending.
    block: Block

    # Status is the return status of the transaction. This will be 1 if the
    # transaction succeeded, or 0 if it failed (due to a revert, or due to
    # running out of gas). If the transaction has not yet been processed, this
    # field will be null.
    status: Long

    # isErc20Call indicates if this is an ERC20 call.
    isErc20Call: Boolean!

    # relayToken represents the reference of the token being referenced
    # by the relay call. It's null if the transaction is actually not a relay
    # call.
    relayToken: ERC20Token
}

# Block is an Opera block chain block.
type Block {
    # Number is the number of this block, starting at 0 for the genesis block.
    number: Long!

    # Hash is the unique block hash of this block.
    hash: Hash!

    # Parent is the parent block of this block.
    parent: Block

    # TransactionCount is the number of transactions in this block.
    transactionCount: Int

    # Timestamp is the unix timestamp at which this block was mined.
    timestamp: Long!

    # GasLimit represents the maximum gas allowed in this block.
    gasLimit: Long!

    # GasUsed represents the actual total used gas by all transactions in this block.
    gasUsed: Long!

    # txHashList is the list of unique hash values of transaction
    # assigned to the block.
    txHashList: [Hash!]!

    # txList is a list of transactions assigned to the block.
    txList: [Transaction!]!
}

# ERC20Token represents a generic ERC20 token.
type ERC20Token {
    # address of the token is used as the token's unique identifier.
    address: Address!

    # name of the token.
    name: String!

    # symbol used as an abbreviation for the token.
    symbol: String!

    # decimals is the number of decimals the token supports.
    # The most common value is 18 to mimic the ETH to WEI relationship.
    decimals: Int!

    # totalSupply represents total amount of tokens across all accounts
    totalSupply: BigInt!

    # logoURL represents a URL address of a logo of the token. It's always
    # provided, but unknown tokens have this set to a generic logo file.
    logoURL: String!

    # balanceOf represents the total available balance of the token
    # on the account regardless of the DeFi usage of the token.
    # It's effectively the amount available held by the ERC20 token
    # on the account behalf.
    balanceOf(owner: Address!): BigInt!

    # allowance represents the amount of ERC20 tokens unlocked
    # by the owner / token holder to be accessible for the given spender.
    allowance(owner: Address!, spender: Address!): BigInt!
}

# DelegationList is a list of delegations edges provided by sequential access request.
type DelegationList {
    "Edges contains provided edges of the sequential list."
    edges: [DelegationListEdge!]!

    """
    TotalCount is the maximum number of delegations
    available for sequential access.
    """
    totalCount: BigInt!

    "PageInfo is an information about the current page of delegation edges."
    pageInfo: ListPageInfo!
}

# DelegationListEdge is a single edge in a sequential list of delegations.
type DelegationListEdge {
    "Cursor defines a scroll key to this edge."
    cursor: Cursor!

    "Delegator represents the delegator provided by this list edge."
    delegation: Delegation!
}

# Delegation represents a delegation on Opera blockchain.
type Delegation {
    "Address of the delegator account."
    address: Address!

    "Identifier of the staker the delegation belongs to."
    toStakerId: Long!

    "Epoch in which the delegation was made."
    createdEpoch: Long!

    "Timestamp of the delegation creation."
    createdTime: Long!

    "Epoch in which the delegation was deactivated."
    deactivatedEpoch: Long

    "Timestamp of the deactivation of the delegation."
    deactivatedTime: Long

    "Amount delegated. It includes all the pending un-delegations."
    amount: BigInt!

    """
    Amount delegated. The current amount still registered
    by SFC contract as delegated amount. It does include pending
    deactivation, but does not include partial un-delegations.
    """
    amountDelegated: BigInt

    "Amount locked in pending un-delegations and withdrawals."
    amountInWithdraw: BigInt!

    "Amount of rewards claimed."
    claimedReward: BigInt

    "Pending rewards for the delegation."
    pendingRewards: PendingRewards!

    "The id of the last epoch rewards has been paid to."
    paidUntilEpoch: Long!

    "List of withdraw requests of the delegation."
    withdrawRequests: [WithdrawRequest!]!

    "List of full delegation deactivation."
    deactivation: [DeactivatedDelegation!]!

    "isFluidStakingActive indicates if the delegation is upgraded to fluid staking."
    isFluidStakingActive: Boolean!

    "isDelegationLocked indicates if the delegation is locked."
    isDelegationLocked: Boolean!

    "lockedFromEpoch represents the id of epoch the lock has been created."
    lockedFromEpoch: Long!

    "lockedUntil represents the timestamp up to which the delegation is locked, zero if not locked."
    lockedUntil: Long!

    # outstandingSFTM represents the amount of sFTM tokens representing the tokenised stake
    # minted and un-repaid on this delegation.
    outstandingSFTM: BigInt!

    # tokenizerAllowedToWithdraw indicates if the stake tokeniser allows the stake
    # to be withdrawn. That means all the sFTM tokens have been repaid and the sFTM
    # debt is effectively zero for the delegation.
    tokenizerAllowedToWithdraw: Boolean!
}

# PendingRewards represents a detail of pending rewards for staking and delegations
type PendingRewards {
    "Staker the pending reward relates to."
    staker: Long!

    "Pending rewards amount."
    amount: BigInt!

    "The first unpaid epoch."
    fromEpoch: Long!

    "The last unpaid epoch."
    toEpoch: Long!
}

# Represents epoch information.
type Epoch {
    "Number the epoch end."
    id: Long!

    "Timestamp of the epoch end."
    endTime: BigInt!

    "Epoch duration in seconds."
    duration: BigInt!

    "Fee at the epoch."
    epochFee: BigInt!

    "Total base reward weight on epoch."
    totalBaseRewardWeight: BigInt!

    "Total transaction reward weight on epoch."
    totalTxRewardWeight: BigInt!

    "Base reward per second of epoch."
    baseRewardPerSecond: BigInt!

    "Total amount staked."
    stakeTotalAmount: BigInt!

    "Total amount delegated."
    delegationsTotalAmount: BigInt!

    "Total supply amount."
    totalSupply: BigInt!
}

# Contract defines block-chain smart contract information container
type Contract {
    "Address represents the contract address."
    address: Address!

    "DeployedBy represents the smart contract deployment transaction reference."
    deployedBy: Transaction!

    "transactionHash represents the smart contract deployment transaction hash."
    transactionHash: Hash!

    "Smart contract name. Empty if not available."
    name: String!

    "Smart contract version identifier. Empty if not available."
    version: String!

    """
    License specifies an open source license the contract was published with.
    Empty if not specified.
    """
    license: String!

    "Smart contract author contact. Empty if not available."
    supportContact: String!

    "Smart contract compiler identifier. Empty if not available."
    compiler: String!

    "Smart contract source code. Empty if not available."
    sourceCode: String!

    "Smart contract ABI definition. Empty if not available."
    abi: String!

    """
    Validated is the unix timestamp at which the source code was validated
    against the deployed byte code. Null if not validated yet.
    """
    validated: Long

    "Timestamp is the unix timestamp at which this smart contract was deployed."
    timestamp: Long!
}

# ContractValidationInput represents a set of data sent from client
# to validate deployed contract with the provided source code.
input ContractValidationInput {
    "Address of the contract being validated."
    address: Address!

    "Optional smart contract name. Maximum allowed length is 64 characters."
    name: String

    "Optional smart contract version identifier. Maximum allowed length is 14 characters."
    version: String

    "Optional smart contract author contact. Maximum allowed length is 64 characters."
    supportContact: String

    """
    License specifies an open source license the contract was published with.
    Empty if not specified.
    """
    license: String

    "Optimized specifies if the compiler was set to optimize the byte code."
    optimized: Boolean = true

    """
    OptimizeRuns specifies number of optimization runs the compiler was set
    to execute during the byte code optimizing.
    """
    optimizeRuns: Int = 200

    "Smart contract source code."
    sourceCode: String!
}

# ContractList is a list of smart contract edges provided by sequential access request.
type ContractList {
    # Edges contains provided edges of the sequential list.
    edges: [ContractListEdge!]!

    # TotalCount is the maximum number of contracts available for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page of contract edges.
    pageInfo: ListPageInfo!
}

# TransactionListEdge is a single edge in a sequential list of transactions.
type ContractListEdge {
    cursor: Cursor!
    contract: Contract!
}

# Hash is a 32 byte binary string, represented by 0x prefixed hexadecimal.
scalar Hash

# Address is a 20 byte Opera address, represented as 0x prefixed hexadecimal number.
scalar Address

# BigInt is a large integer value. Input is accepted as either a JSON number,
# or a hexadecimal string alternatively prefixed with 0x. Output is 0x prefixed hexadecimal.
scalar BigInt

# Long is a 64 bit unsigned integer value.
scalar Long

# Bytes is an arbitrary length binary string, represented as 0x-prefixed hexadecimal.
# An empty byte string is represented as '0x'.
scalar Bytes

# Cursor is a string representing position in a sequential list of edges.
scalar Cursor

# CurrentState represents the current active state
# of the chain information condensed on one place.
type CurrentState {
    # epoch is the last sealed Epoch structure
    sealedEpoch: Epoch!

    # blocks represents number of blocks in the chain.
    blocks: Long!

    # transactions represents number of transactions in the chain.
    transactions: Long!

    # validators represents number of validators in the network.
    validators: Long!

    # accounts represents number of accounts participating on transactions.
    accounts: Long!

    # sfcContractAddress is the address of the SFC contract
    # used for PoS staking control.
    sfcContractAddress: Address!

    # sfcLockingEnabled indicates if the SFC locking feature is enabled.
    sfcLockingEnabled: Boolean!

    # sfcVersion indicates the current version of the SFC contract.
    # The version is encoded into 3 bytes representing ASCII version numbers
    # with the most significant byte first [<8bit major><8bit minor><8bit revision>].
    # I.e. Version 1.0.2 = "102" = 0x313032
    sfcVersion: Long!
}
# Represents staker information.
type Staker {
    "Id number the staker."
    id: Long!

    "Staker address."
    stakerAddress: Address!

    "Amount of total staked tokens in WEI."
    totalStake: BigInt

    "Amount of own staked tokens in WEI."
    stake: BigInt

    "Amount of tokens delegated to the staker in WEI."
    delegatedMe: BigInt

    """
    Maximum total amount of tokens allowed to be delegated
    to the staker in WEI.
    This value depends on the amount of self staked tokens.
    """
    totalDelegatedLimit: BigInt!

    """
    Maximum amount of tokens allowed to be delegated to the staker
    on a new delegation in WEI.
    This value depends on the amount of self staked tokens.
    """
    delegatedLimit: BigInt!

    "Is this a validator record."
    isValidator: Boolean!

    "Is the staker active."
    isActive: Boolean!

    "Is the staker considered to be cheater."
    isCheater: Boolean!

    "Is the staker offline."
    isOffline: Boolean!

    "isStakeLocked signals if the staker locked the stake."
    isStakeLocked: Boolean!

    "Epoch in which the staker was created."
    createdEpoch: Long!

    "Timestamp of the staker creation."
    createdTime: Long!

    "lockedFromEpoch is the identifier of the epoch the stake lock was created."
    lockedFromEpoch: Long!

    "lockedUntil is the timestamp up to which the stake is locked, zero if not locked."
    lockedUntil: Long!

    "Epoch in which the staker was deactivated."
    deactivatedEpoch: Long!

    "Timestamp of the staker deactivation."
    deactivatedTime: Long!

    "How many blocks the staker missed."
    missedBlocks: Long!

    "Number of seconds the staker is offline."
    downtime: Long!

    "Proof of importance score."
    poi: BigInt

    "Base weight for rewards distribution."
    baseRewardWeight: BigInt

    "Weight for transaction rewards distribution."
    txRewardWeight: BigInt

    "Validation score."
    validationScore: BigInt

    "Origination score."
    originationScore: BigInt

    "Amount of rewards claimed in WEI."
    claimedRewards: BigInt

    "Amount of rewards claimed by delegators in WEI."
    delegationClaimedRewards: BigInt

    """
    List of delegations of this staker. Cursor is used to obtain specific slice
    of the staker's delegations. The most recent delegations
    are provided if cursor is omitted.
    """
    delegations(cursor: Cursor, count: Int = 25):DelegationList!

    """
    Status is a binary encoded status of the staker.
    Ok = 0, bin 1 = Fork Detected, bin 256 = Validator Offline
    """
    status: Long!

    "StakerInfo represents extended staker information from smart contract."
    stakerInfo: StakerInfo

    """
    List of withdraw requests of the stake.
    Contains only withdrawal requests of the staking account,
    not the requests of the stake delegators.
    """
    withdrawRequests: [WithdrawRequest!]!
}

# FMintAccount represents an informastion about account details
# in DeFi/fMint protocol.
type FMintAccount {
    # address of the DeFi account.
    address: Address!

    # collateralList represents a list of all collateral tokens
    # linked with the account.
    collateralList: [Address!]!

    # collaterals represents a list of all collateral assets.
    collateral: [FMintTokenBalance!]!

    # collateralValue represents the current collateral value
    # in ref. denomination (fUSD).
    collateralValue: BigInt!

    # debtList represents a list of all debt tokens linked with the account.
    debtList: [Address!]!

    # debts represents the list of all the current borrowed tokens.
    debt: [FMintTokenBalance!]!

    # debtValue represents the current debt value
    # in ref. denomination (fUSD).
    debtValue: BigInt!

    # rewardsEarned represents accumulated rewards
    # earned on the DeFi / fMint account for the excessive
    # collateral value. Please note that the rewards could still
    # be burned, if the account is not eligible to claim the reward.
    rewardsEarned: BigInt!

    # rewardsStashed represents accumulated rewards
    # earned on the DeFi / fMint account for the excessive
    # collateral value and stored into the stash for future
    # claim.
    rewardsStashed: BigInt!

    # canClaimRewards informs if the fMint account collateral
    # to debt is high enough to allow earned rewards claiming.
    canClaimRewards: Boolean!

    # canReceiveRewards informs if the fMint account collateral
    # to debt is high enough to receive earned rewards. If the ratio
    # is below configured one, earned rewards are burned.
    canReceiveRewards: Boolean!

    # canPushNewRewards indicates if new rewards are unlocked
    # inside the reward distribution and can be pushed into
    # the system to distribute them among eligible accounts.
    canPushNewRewards: Boolean!
}

# FMintTokenBalance represents a balance of a specific DeFi token
# on an fMint protocol account.
# The balance is used for both collateral deposits and minting debt.
type FMintTokenBalance {
    # type represents the type of the balance record.
    type: DefiTokenBalanceType!

    # tokenAddress represents unique identifier of the token.
    tokenAddress: Address!

    # token represents the detail of the token
    token: DefiToken!

    # current balance of the token on the account.
    balance: BigInt!

    # value of the current balance of the token on the account
    # in ref. denomination (fUSD).
    value: BigInt!
}

# DefiSettings represents the set of current settings and limits
# applied to DeFi operations.
type DefiSettings {
    # mintFee4 is the current fee applied to all minting operations on fMint protocol.
    # Value is represented in 4 digits, e.g. value 25 = 0.0025 => 0.25% fee.
    mintFee4: BigInt!

    # minCollateralRatio4 is the minimal allowed ratio between
    # collateral and debt values in ref. denomination (fUSD)
    # on which the borrow trade is allowed.
    # Value is represented in 4 digits,
    # e.g. value 25000 = 3.0x => (debt x 3.0 <= collateral)
    minCollateralRatio4: BigInt!

    # rewardCollateralRatio4 is the minimal ratio between
    # collateral and debt values in ref. denomination (fUSD)
    # on which the account is eligible for rewards distribution.
    # Collateral below this ratio means all the pending rewards
    # will be burnt and lost.
    rewardCollateralRatio4: BigInt!

    # decimals represents the decimals / digits correction
    # applied to the fees and ratios internally to correctly represent
    # fraction numbers. E.g. correction value 4 => ratio/fee x 10000.
    decimals: Int!

    # priceOracleAggregate is the address of the current price oracle
    # aggregate used by the DeFi to obtain USD price of tokens managed.
    priceOracleAggregate: Address!

    # StakeTokenizerContract is the address of the Stake Tokenizer contract.
    StakeTokenizerContract: Address!

    # StakeTokenizedERC20Token is the address of the Tokenized Stake ERC20 contract.
    StakeTokenizedERC20Token: Address!

    # fMintAddress is the address of the fMint contract.
    fMintContract: Address!

	# fMintAddressProvider represents the address of the fMint address provider.
	fMintAddressProvider: Address!

    # tokenRegistryContract is the address of the fMint token registry.
    fMintTokenRegistry: Address!

    # fMintRewardDistribution is the address of the DeFi fMint
    # reward distribution contract.
    fMintRewardDistribution: Address!

    # fMintCollateralPool is the address of the fMint collateral pool.
    fMintCollateralPool: Address!

    # fMintDebtPool is the address of the fMint debt pool.
    fMintDebtPool: Address!

    # uniswapCoreFactory is the address of the Uniswap Core Factory contract.
    uniswapCoreFactory: Address!

    # uniswapRouter is the address of the Uniswap Router contract.
    uniswapRouter: Address!
}

# EstimatedRewards represents a calculated rewards estimation for an account or amount staked
type EstimatedRewards {
    "Amount of FTM tokens expected to be staked for the calculation."
    staked: Long!

    """
    dailyReward represents amount of FTM tokens estimated
    to be rewarded for staked amount in average per day.
    """
    dailyReward: BigInt!

    """
    weeklyReward represents amount of FTM tokens estimated
    to be rewarded for staked amount in average per week.
    """
    weeklyReward: BigInt!

    """
    monthlyReward represents amount of FTM tokens estimated
    to be rewarded for staked amount in average per month.
    """
    monthlyReward: BigInt!

    """
    yearlyReward represents amount of FTM tokens estimated
    to be rewarded for staked amount in average per year.
    """
    yearlyReward: BigInt!

    """
    currentRewardYearRate represents average reward rate
    for any staked amount in average per year.
    The value is calculated as linear gross proceeds for staked amount
    of tokens yearly.
    """
    currentRewardRateYearly: Int!

    """
    Total amount of staked FTM tokens used for the calculation in WEI units.
    The estimation uses total staked amount, not the effective amount provided
    by the last epoch. The effective amount does include current
    un-delegations and also skips offline self-staking and flagged staking.
    """
    totalStaked: BigInt!

    """
    Information about the last sealed epoch of the Opera blockchain.
    The epoch provides useful information about total FTM supply,
    total amount staked, rewards rate and weight, fee, etc.
    """
    lastEpoch: Epoch!
}
# WithdrawRequest represents a request for partial stake withdraw.
type WithdrawRequest {
    "Address of the authorized request."
    address: Address!

    "Address of the receiving account."
    receiver: Address!

    "Account receiving the withdraw request."
    account: Account!

    "Staker Id of the staker involved in the withdraw request."
    stakerID: Long!

    "Details of the staker involved in the withdraw request."
    staker: Staker!

    "Unique withdraw request identifier."
    withdrawRequestID: BigInt!

    "Is this a partial delegation withdraw, or staker withdraw?"
    isDelegation: Boolean!

    "Amount of WEI requested to be withdrawn."
    amount: BigInt!

    "Block in which the withdraw request was registered."
    requestBlock: Block!

    """
    Block in which the withdraw request was finalized.
    The value is NULL for pending request.
    """
    withdrawBlock: Block

    """
    Amount of WEI slashed as a penalty for cheating.
    The penalty is applied not only to staker withdraw,
    but also to delegations of a cheating staker.
    The value is NULL for pending requests.
    """
    withdrawPenalty: BigInt
}

# DeactivatedDelegation represents a prepared delegation full withdraw.
# Fully withdrawn delegations must be prepared first and finalized
# only after the lock down period passes.
type DeactivatedDelegation {
    "Address of the delegator."
    address: Address!

    "Staker Id of the staker involved in the withdraw request."
    stakerID: Long!

    "Details of the staker involved in the withdraw request."
    staker: Staker!

    "Block in which the delegation deactivation was registered."
    requestBlock: Block!

    """
    Block in which the delegation was withdrawn.
    The value is NULL for pending request.
    """
    withdrawBlock: Block

    """
    Amount of WEI slashed as a penalty for cheating.
    The penalty is applied to delegators' rewards
    in case their staker is flagged.
    The value is NULL for pending requests.
    """
    withdrawPenalty: BigInt
}

# UniswapPair represents the information about single
# Uniswap pair managed by the Uniswap Core.
type UniswapPair {
    # pairAddress represents the Address of the Pair
    # and also the address of the ERC20 token
    # managing the share of each liquidity participant.
    pairAddress: Address!

    # tokens represent the list of tokens in the pair.
    # The array always contain two tokens, their order
    # is irrelevant from the Uniswap perspective, but
    # we always return them in the same order.
    tokens: [ERC20Token!]!

    # reserves of the tokens of the pair.
    # The reserve index inside the array corresponds
    # with the token position.
    reserves: [BigInt!]!

    # The timestamp of the block
    # in which this reserves state was reached.
    reservesTimeStamp: Long!

    # cumulative prices of the tokens of the pair.
    # The price index inside the array corresponds
    # with the token position.
    cumulativePrices: [BigInt!]!

    # lastKValue represents the last coeficient
    # of reserves multiplied. It's the value Uniswap protocol
    # uses to control reserves growth on both sides of the pool.
    lastKValue: BigInt!

    # totalSupply represents the total amount
    # of the pair tokens in circulation and represents
    # the total share pool of all the participants.
    totalSupply: BigInt!

    # share represents the share of the given user/participant on the pair.
    # To get the share percentage, divide this value by the total supply
    # of the pair.
    shareOf(user: Address!): BigInt!
}


# Account defines block-chain account information container
type Account {
    "Address is the address of the account."
    address: Address!

    "Balance is the current balance of the Account in WEI."
    balance: BigInt!

    """
    TotalValue is the current total value fo the account in WEI.
    It includes available balance,
    delegated amount and pending staking rewards.
    """
    totalValue: BigInt!

    """
    Stashed represents the amount of WEI stashed
    on this account, if any. Stashed amount comes from
    the delegation and validation rewards.
    """
    stashed: BigInt!

    """
    canUnStash informs if there is a stash which can be claimed.
    Please note that stash claiming can be locked inside SFC.
    """
    canUnStash: Boolean!

    "txCount represents number of transaction sent from the account."
    txCount: Long!

    """
    txList represents list of transactions of the account
    in form of TransactionList.
    """
    txList (cursor:Cursor, count:Int!): TransactionList!

    "Details of a staker, if the account is a staker."
    staker: Staker

    "List of delegations of the account, if the account is a delegator."
    delegations(cursor:Cursor, count:Int = 25): DelegationList!

    "Details about smart contract, if the account is a smart contract."
    contract: Contract
}

# GovernanceContract represents basic information
# about a Governance contract deployed on the block chain.
type GovernanceContract {
    # name represents the name of the contract
    name: String!

    # address represents the address of the Governance contract
    address: Address!

    # totalProposals represents the total number of proposals
    # managed by the Governance contract.
    totalProposals: BigInt!

    # proposals represents list of proposals on the contract.
    proposals(cursor:Cursor, count:Int!, activeOnly: Boolean = false):GovernanceProposalList!

    # proposal provides specific Governance Proposal detail identified
    # by its ID inside the Governance contract.
    proposal(id: BigInt!):GovernanceProposal

    # delegationsBy represents list of delegations for the given address.
    # If the address does not delegate, the list is empty.
    # Delegations are handled by the governed contract, so this list may
    # be always empty for certain Governance instances. If the list is empty
    # the source address may still be eligible for voting by itself.
    delegationsBy(from: Address!): [Address!]!

    # canVote checks if the given address can submit votes to Proposals
    # of this Governance conract. The ability to vote is bound
    # to the governed contract logic and can be unavailable
    # to some network participants on certain situation.
    canVote(from: Address!): Boolean!

    # proposalFee represents the fee required by the Governance
    # to accept proposals. The fee is never refunded,
    # even if a Proposal is canceled.
    proposalFee: BigInt!

    # totalVotingPower represents the total voting power available
    # on the Governance contract in the form of votes
    # weight.
    totalVotingPower: BigInt!
}

# GovernanceProposalList is a list of governance proposal edges
# provided by sequential access request.
type GovernanceProposalList {
    # Edges contains provided edges of the sequential list.
    edges: [GovernanceProposalListEdge!]!

    # TotalCount is the maximum number of governance proposals
    # available for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page of governance
    # proposal edges.
    pageInfo: ListPageInfo!
}

# TransactionListEdge is a single edge in a sequential list
# of governance proposals.
type GovernanceProposalListEdge {
    cursor: Cursor!
    proposal: GovernanceProposal!
}

# GovernanceProposal represents the details of a single proposal
# in the governance contract.
type GovernanceProposal {
    # governanceId represents the identifier of the Governance
    # contract this Proposal belongs to.
    governanceId: Address!

    # governance represents the Governance contract reference.
    # Please make sure not to engage in a circular reference too deep.
    governance: GovernanceContract!

    # id identifier of the proposal in the governance contract
    # the proposal is managed by.
    id: BigInt!

    # name represents a name of the Proposal.
    name: String!

    # description represents a textual description of the Proposal.
    description: String!

    # state represents the state of the Proposal.
    state: ProposalState!

    # contract represents the contract of the Proposal. Each Proposal
    # is represented by a contract responsible for maintaining the Proposal
    # parameters, options and finalization actions.
    contract: Address!

    # proposalType represents the type of the Proposal that corresponds
    # with the Proposal Template.
    proposalType: Long!

    # isExecutable identifies if the proposal will be finalized
    # by executing a finalizing code.
    isExecutable: Boolean!

    # minVotes corresponds with the minimal weight of votes
    # required by the Proposal to be settled in any way
    # other than REJECTED.
    minVotes: BigInt!

    # minAgreement represents the minimal agreement weight
    # required to be reached on any of the Proposal options
    # so the Proposal could be settled in any way
    # other than REJECTED.
    minAgreement: BigInt!

    # totalWeight represents the total voting weight
    # of all voters allwed on the proposal. This is effectively
    # the maximum weight an option can gain if all the voters
    # would favor it with the top value of the scale.
    totalWeight: BigInt!

    # votedWeightRatio represents the percentage of the total voting weight
    # already counted towards the proposal options. The ratio increases
    # as more voters place their votes. If no vote was placed the value is zero,
    # if all the voters placed their votes aither directly, or over a delegation,
    # the value is 100.
    votedWeightRatio: Int!

    # opinionScales is the scale of opinions on available options.
    # A voter provides a single opinion picked from the scale
    # for each option during the voting for a proposal.
    # I.e.: Scales {0, 2, 3, 4, 5} represent opinions of
    # {strongly disagree, disagree, neutral, agree and strongly agree}.
    opinionScales: [Long!]!

    # options is a list of options available on the Proposal.
    # A voter must provide their opinion expressed by a chosen scale
    # for each option on the list. It's generally better to scatter
    # opinions across options instead of having a binary view.
    options: [String!]!

    # votingStarts is the time stamp of the voting getting opened
    # to receive votes.
    votingStarts: Long!

    # votingMayEnd is the time stamp when the voting could be closed
    # if enough votes are collected to settle the Proposal (winner option is selectable).
    votingMayEnd: Long!

    # votingMustEnd is the time stamp when the voting must be closed.
    # If enough votes to settle the Proposal were not collected up until this time
    # the Proposal is rejected and will not be settled in any way (no winner option is selectable).
    votingMustEnd: Long!

    # optionStates is the list of states of all the options in the Proposal.
    # Warning: This is an expensive call, use with caution.
    optionStates: [OptionState!]!

    # optionState represents a state of the selected option of the Proposal.
    optionState(optionId:BigInt!):OptionState

    # vote pulls the vote for the given <from> address linked with the <delegatedTo> delegation
    # recipient. If the <from> address is not delegator in the context of the governance
    # subject contract, the <delegatedTo> may be left empty, or set to the same address
    # as the <from> address.
    vote(from: Address!, delegatedTo: Address): GovernanceVote
}

# ProposalState represents the state of the whole proposal.
type ProposalState {
    # isResolved signals if the Proposal is already resolved.
    isResolved: Boolean!

    # winnerId is the identifier of the winning option.
    winnerId: BigInt

    # votes is the number of votes received on the Proposal.
    votes: BigInt!

    # status represents the status of the Proposal.
    # 0 = Initial, 1 = Resolved, 2 = Failed, 4 = Canceled, 8 = Execution Expired
    status: BigInt!
}

# OptionState represents a state in options of a Proposal.
type OptionState {
    # optionId is the identifier of the option,
    # effectively option index in the options array
    optionId: BigInt!

    # votes is the weight of all votes received across all votes;
    # the projection of the votes to this state uses it to calculate
    # actual agreement.
    votes: BigInt!

    # agreement represents the rated weight of all the votes towards this option
    # based on the opinion scale of the proposal and selected opinion scale level of
    # each vote.
    # this effectively reflects the absolute weight of affection of all voters
    # towards this option.
    agreement: BigInt!

    # agreementRatio represents the relative ratio of the option agreement
    # to the total weight of all votes in 18 digits.
    agreementRatio: BigInt!
}

# GovernanceVote is the vote in the context of the given Governance Proposal.
type GovernanceVote {
    # governanceId is the identifier of the Governance contract.
    governanceId: Address!

    # proposalId is the identifier of the proposal of the contract.
    proposalId: BigInt!

    # from is the address of the voting party
    from: Address!

    # delegatedTo is the address of the delegation the vote refers to.
    delegatedTo: Address

    # weight represents the weight of the vote
    weight: BigInt!

    # choices represents the list of opinions on the Proposal options the vote
    # presented.
    choices: [Long!]!
}
# Ballot represents an official deployed ballot contract
# used for Fantom Opera related voting poll.
type Ballot {
    # Address of the ballot, correspond with the smart contract address.
    address: Address!

    # Deployed smart contract handling the ballot voting.
    contract: Contract!

    # Short name of the ballot.
    name: String!

    # URL of the ballot detailed information page.
    detailsUrl: String!

    # An approximate timestamp after which the ballot opens for voting.
    start: Long!

    # An approximate timestamp after which the ballot
    # is closed and no longer accepts votes.
    end: Long!

    # Informs if the ballot is open for voting.
    isOpen: Boolean!

    # Informs if the ballot has already been finalized
    # and the winning proposal is available.
    isFinalized: Boolean!

    # List of proposals of the ballot.
    proposals: [BallotProposal!]!

    # Index of the winning proposal.
    # Is NULL if the ballot has not been finalized yet.
    winner: Int
}

# BallotProposal represents a proposal in the ballot.
type BallotProposal {
    # id is the ballot proposal identifier.
    id: Int!

    # name is the name of the proposal option.
    name: String!
}

# Vote represents a selected vote in a ballot.
type Vote {
    # Address of the ballot the Vote relates to.
    ballot: Address!

    # Address of the voter who placed the vote.
    voter: Address!

    # Account of the voter who placed the vote.
    account: Account!

    # The selected proposal index the voter chose.
    vote: Int
}

# BallotList is a list of ballot edges provided by sequential access request.
type BallotList {
    # Edges contains provided edges of the sequential list.
    edges: [BallotListEdge!]!

    # TotalCount is the maximum number of ballots available
    # for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page
    # of ballot edges.
    pageInfo: ListPageInfo!
}

# BallotListEdge is a single edge in a sequential list of ballots.
type BallotListEdge {
    cursor: Cursor!
    ballot: Ballot!
}

# Root schema definition
schema {
    query: Query
    mutation: Mutation
    subscription: Subscription
}

# Entry points for querying the API
type Query {
    # version represents the API server version responding to your requests.
    version: String!

    # State represents the current state of the blockchain and network.
    state: CurrentState!

    # Total number of accounts active on the Opera blockchain.
    accountsActive:Long!

    # Get an Account information by hash address.
    account(address:Address!):Account!

    # Get list of Contracts with at most <count> edges.
    # If <count> is positive, return edges after the cursor,
    # if negative, return edges before the cursor.
    # For undefined cursor, positive <count> starts the list from top,
    # negative <count> starts the list from bottom.
    # ValidatedOnly specifies if the list should contain all the Contracts,
    # or just contracts with validated byte code and available source/ABI.
    contracts(validatedOnly: Boolean = false, cursor:Cursor, count:Int!):ContractList!

    # Get block information by number or by hash.
    # If neither is provided, the most recent block is given.
    block(number:Long, hash: Hash):Block

    # Get transaction information for given transaction hash.
    transaction(hash:Hash!):Transaction

    # Get list of Blocks with at most <count> edges.
    # If <count> is positive, return edges after the cursor,
    # if negative, return edges before the cursor.
    # For undefined cursor, positive <count> starts the list from top,
    # negative <count> starts the list from bottom.
    blocks(cursor:Cursor, count:Int!):BlockList!

    # Get list of Transactions with at most <count> edges.
    # If <count> is positive, return edges after the cursor,
    # if negative, return edges before the cursor.
    # For undefined cursor, positive <count> starts the list from top,
    # negative <count> starts the list from bottom.
    transactions(cursor:Cursor, count:Int!):TransactionList!

    # Get the id of the current epoch of the Opera blockchain.
    currentEpoch:Long!

    # Get information about specified epoch. Returns current epoch information
    # if id is not provided.
    epoch(id: Long!): Epoch!

    # The last staker id in Opera blockchain.
    lastStakerId: Long!

    # The number of stakers in Opera blockchain.
    stakersNum: Long!

    # Staker information. The staker is loaded either by numeric ID,
    # or by address. null if none is provided.
    staker(id: Long, address: Address): Staker

    # List of staker information from SFC smart contract.
    stakers: [Staker!]!

    # The list of delegations for the given staker ID.
    # Cursor is used to obtain specific slice of the staker's delegations.
    # The most recent delegations are provided if cursor is omitted.
    delegationsOf(staker:Long!, cursor: Cursor, count: Int = 25): DelegationList!

    # Get the details of a specific delegation by it's delegator address
    # and staker the delegation belongs to.
    delegation(address:Address!, staker: Long!): Delegation

    # Get the list of all delegations by it's delegator address.
    delegationsByAddress(address:Address!, cursor: Cursor, count: Int = 25): DelegationList!

    # Returns the current price per gas in WEI units.
    gasPrice: Long!

    # Get price details of the Opera blockchain token for the given target symbols.
    price(to:String!):Price!

    # Get calculated staking rewards for an account or given
    # staking amount in FTM tokens.
    # At least one of the address and amount parameters must be provided.
    # If you provide both, the address takes precedence and the amount is ignored.
    estimateRewards(address:Address, amount:Long):EstimatedRewards!

    # Get official ballot information by its address.
    ballot(address: Address!):Ballot

    # Get list of official Ballots with at most <count> edges.
    # If <count> is positive, return edges after the cursor,
    # if negative, return edges before the cursor.
    # For undefined cursor, positive <count> starts the list from top,
    # negative <count> starts the list from bottom.
    ballots(cursor: Cursor, count: Int!):BallotList!

    # Get list of recently closed official Ballots
    # with at most <count> edges. If the <finalized> is set to false
    # the list contains ballots, which ended, but were not resolved
    # yet.
    ballotsClosed(finalized: Boolean = true, count: Int = 25):[Ballot!]!

    # Get list of currently active Ballots with at most <count> edges.
    ballotsActive(count: Int = 25):[Ballot!]!

    # List of all votes of the given voter identified by the address
    # for the given list of ballots identified by an array of ballot
    # addresses.
    votes(voter:Address!, ballots:[Address!]!):[Vote!]!
    
    # defiConfiguration exposes the current DeFi contract setup.
    defiConfiguration:DefiSettings!

    # defiTokens represents a list of all available DeFi tokens.
    defiTokens:[DefiToken!]!

    # defiNativeToken represents the information about the native token
    # wrapper ERC20 contract. Returns NULL if the native token wraper
    # is not available.
    defiNativeToken: ERC20Token

    # fMintAccount provides DeFi/fMint information about an account on fMint protocol.
    fMintAccount(owner: Address!):FMintAccount!

    # fMintTokenAllowance resolves the amount of ERC20 tokens unlocked
    # by the token owner for DeFi/fMint operations.
    fMintTokenAllowance(owner: Address!, token: Address!):BigInt!

    # defiUniswapPairs represents a list of all pairs managed
    # by the Uniswap Core contract on Opera blockchain.
    defiUniswapPairs: [UniswapPair!]!

    # defiUniswapAmountsOut calculates the expected output amounts
    # required to finalize a swap operation specified by a list of
    # tokens involved in the swap steps and the input amount.
    # At least two addresses of tokens must be given
    # for the calculation to succeed.
    defiUniswapAmountsOut(amountIn: BigInt!, tokens:[Address!]!): [BigInt!]!

    # defiUniswapAmountsIn calculates the expected input amounts
    # required to finalize a swap operation specified by a list of
    # tokens involved in the swap steps and the output amount.
    # At least two addresses of tokens must be given
    # for the calculation to succeed.
    defiUniswapAmountsIn(amountOut: BigInt!, tokens:[Address!]!): [BigInt!]!

    # defiUniswapQuoteLiquidity calculates optimal amount of tokens
    # of an Uniswap pair defined by a pair of tokens for the given amount
    # of both tokens desired to be added to the liquidity pool.
    # The function can be used to calculate minimal amount of tokens expected
    # to be added to the pool on both sides on addLiquidity call.
    # Please note "amountsIn" must be in the same order as are the tokens.
    defiUniswapQuoteLiquidity(tokens:[Address!]!, amountsIn:[BigInt!]!): [BigInt!]!

    # erc20Token provides the information about an ERC20 token specified by it's
    # address, if available. The resolver returns NULL if the token does not exist.
    erc20Token(token: Address!):ERC20Token

    # erc20TokenList provides list of the most active ERC20 tokens
    # deployed on the block chain.
    erc20TokenList(count: Int = 50):[ERC20Token!]!

    # ercTotalSupply provides the current total supply amount of a specified ERC20 token
    # identified by it's ERC20 contract address.
    ercTotalSupply(token: Address!):BigInt!

    # ercTokenBalance provides the current available balance of a specified ERC20 token
    # identified by it's ERC20 contract address.
    ercTokenBalance(owner: Address!, token: Address!):BigInt!

    # ercTokenAllowance provides the current amount of ERC20 tokens unlocked
    # by the token owner for the spender to be manipulated with.
    ercTokenAllowance(token: Address!, owner: Address!, spender: Address!):BigInt!

    # govContracts provides list of governance contracts.
    govContracts:[GovernanceContract!]!

    # govContract provides a specific Governance contract information by its address.
    govContract(address: Address!): GovernanceContract

    # govProposals represents list of joined proposals across all the Governance contracts.
    govProposals(cursor:Cursor, count:Int!, activeOnly: Boolean = false):GovernanceProposalList!
}

# Mutation endpoints for modifying the data
type Mutation {
    # SendTransaction submits a raw signed transaction into the block chain.
    # The tx parameter represents raw signed and RLP encoded transaction data.
    sendTransaction(tx: Bytes!):Transaction

    # Validate a deployed contract byte code with the provided source code
    # so potential users can check the contract source code, access contract ABI
    # to be able to interact with the contract and get the right metadata.
    # Returns updated contract information. If the contract can not be validated,
    # it raises a GraphQL error.
    validateContract(contract: ContractValidationInput!): Contract!
}

# Subscriptions to live events broadcasting
type Subscription {
    # Subscribe to receive information about new blocks in the blockchain.
    onBlock: Block!

    # Subscribe to receive information about new transactions in the blockchain.
    onTransaction: Transaction!
}

`

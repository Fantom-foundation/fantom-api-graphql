package gqlschema

// Auto generated GraphQL schema bundle
const schema = `
# FMintUserToken represents a pair of fMint protocol user
# and a token used by the user for a specific operation
# as reported by fMint users listings.
type FMintUserToken {
    # purpose represents the type of usage of the token by the user.
    purpose: FMintUserTokenPurpose!

    # userAddress represents the address of the user account.
    userAddress: Address!

    # account represents the full record of the fMint account
    account: FMintAccount!

    # tokenAddress represents the address of the associated token.
    tokenAddress: Address!

    # token represents the detail of the token associated.
    token: ERC20Token!
}

# FMintUserTokenPurpose represents the purpose of the fMint user token pair.
enum FMintUserTokenPurpose {
    FMINT_COLLATERAL
    FMINT_DEBT
}
# DailyTrxVolume represents a view of an aggregated flow
# of transactions on the network on specific day.
type DailyTrxVolume {
    # day represents the day of the aggregation in format YYYY-MM-DD
    # i.e. 2021-01-23 for January 23rd, 2021
    day: String!

    # volume represent the number of transactions originated / mined
    # by the network on the day.
    volume: Int!

    # amount represents the total value of native tokens transferred
    # by the network on the day. Please note this includes only direct
    # token transfers.
    amount: BigInt!

    # gas represents the total amount of gas consumed by transactions
    # on the network on the day.
    gas: BigInt!
}

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

    # totalDeposited represents total amount of deposited tokens collateral on fMint.
    totalDeposit: BigInt!

    # totalDebt represents total amount of borrowed/minted tokens on fMint.
    totalDebt: BigInt!
}

# DefiTokenBalanceType represents the type of DeFi token balance record.
enum DefiTokenBalanceType {
    COLLATERAL
    DEBT
}

# TokenTransactionType represents a type of ERC-20/ERC-721/ERC-1155 transaction.
enum TokenTransactionType {
    TRANSFER
    MINT
    BURN
    APPROVAL
    APPROVAL_FOR_ALL
    OTHER
}

# ERC20Transaction represents a transaction on an ERC20 token.
type ERC20Transaction {
    # trxHash represents a hash of the transaction
    # executing the ERC20 call.
    trxHash: Bytes32!

    # transaction represents the transaction
    # executing the ERC20 call.
    transaction: Transaction!

    # trxIndex represents the index
    # of the ERC20 call in the transaction logs.
    trxIndex: Long!

    # tokenAddress represents the address
    # of the ERC20 token contract.
    tokenAddress: Address!

    # token represents the token detail involved.
    token: ERC20Token!

    # trxType is the type of the transaction.
    trxType: TokenTransactionType!

    # sender represents the address of the token owner
    # sending the tokens, e.g. the sender.
    sender: Address!

    # recipient represents the address of the token recipient.
    recipient: Address!

    # amount represents the amount of tokens involved
    # in the transaction; please make sure to interpret the amount
    # with the correct number of decimals from the ERC20 token detail.
    amount: BigInt!

    # timeStamp represents the Unix epoch time stamp
    # of the ERC20 transaction processing.
    timeStamp: Long!
}
# RewardClaimList is a list of reward claims linked to delegations.
type RewardClaimList {
    # Edges contains provided edges of the sequential list.
    edges: [RewardClaimListEdge!]!

    # TotalCount is the maximum number of reward claims
    # available for sequential access.
    totalCount: Long!

    # PageInfo is an information about the current page
    # of reward claim edges.
    pageInfo: ListPageInfo!
}

# RewardClaimListEdge is a single edge in a sequential list
# of reward claims.
type RewardClaimListEdge {
    # Cursor defines a scroll key to this edge.
    cursor: Cursor!

    # claim represents the reward claim detail provided by this list edge.
    claim: RewardClaim!
}

# EpochList is a list of epoch edges provided by sequential access request.
type EpochList {
    # Edges contains provided edges of the sequential list.
    edges: [EpochListEdge!]!

    # TotalCount is the maximum number of epochs
    # available for sequential access.
    totalCount: Long!

    # PageInfo is an information about the current page of epoch list edges.
    pageInfo: ListPageInfo!
}

# EpochListEdge is a single edge in a sequential list of epochs.
type EpochListEdge {
    #Cursor defines a scroll key to this edge.
    cursor: Cursor!

    # epoch represents the Epoch provided by this list edge.
    epoch: Epoch!
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

# ERC1155Contract represents a generic ERC1155 multi-token contract.
type ERC1155Contract {
    # address of the token is used as the token's unique identifier.
    address: Address!

    # uri provides URI of Metadata JSON Schema for given token.
    uri(tokenId: BigInt!): String

    # balanceOf represents amount of tokens on the account.
    balanceOf(owner: Address!, tokenId: BigInt!): BigInt!

    # balanceOf represents amount of tokens on the account.
    balanceOfBatch(owners: [Address!]!, tokenIds: [BigInt!]!): [BigInt!]!

    # isApprovedForAll queries the approval status of an operator for a given owner.
    isApprovedForAll(owner: Address!, operator: Address!): Boolean
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
# GasPriceTick represents a collected gas price tick.
type GasPriceTick {
    # fromTime is the time of the tick measurement start
    fromTime: Time!

    # toTime is the time of the tick measurement end
    toTime: Time!

    # openPrice is the opening gas price in the tick
    openPrice: Long!

    # closePrice is the closing gas price in the tick
    closePrice: Long!

    # minPrice is the lowest reached price in the tick
    minPrice: Long!

    # maxPrice is the highest reached price in the tick
    maxPrice: Long!

    # avgPrice is the average reached price in the tick
    avgPrice: Long!
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
    hash: Bytes32!

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
    blockHash: Bytes32

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

    # tokenTransactions represents a list of generic token transactions executed in the scope
    # of the transaction call; token type and transaction type is provided.
    tokenTransactions: [TokenTransaction!]!

    # erc20Transactions provides list of ERC-20 token transactions executed in the scope
    # of this blockchain transaction call.
    erc20Transactions: [ERC20Transaction!]!

    # erc721Transactions provides list of ERC-721 NFT transactions executed in the scope
    # of this blockchain transaction call.
    erc721Transactions: [ERC721Transaction!]!

    # erc1155Transactions provides list of ERC-1155 NFT transactions executed in the scope
    # of this blockchain transaction call.
    erc1155Transactions: [ERC1155Transaction!]!
}

# NetworkNodeGroupLevel represents the detail of network node count aggregation.
enum NetworkNodeGroupLevel {
    CONTINENT
    COUNTRY
    STATE
}

# NetworkNodeGroup represents an aggregated group of Opera network nodes.
type NetworkNodeGroup {
    # topRegion represents the name of the top level location of the aggregation group.
    topRegion: String!

    # region represents the name of the location of the aggregation group
    # based on selected detail level.
    region: String!

    # count represents the number of nodes in the aggregation group.
    count: Int!

    # latitude represents an average geographic coordinate
    # that specifies the north–south position of the group on the Earth's surface.
    latitude: Float!

    # longitude represents an average geographic coordinate
    # that specifies the east–west position of the group on the Earth's surface.
    longitude: Float!

    # pct represents the percentage share of the aggregation group
    # compared to the number of all known active nodes. The number is provided
    # as fixed point integer with 1 decimal precision (i.e. 258 = 25.8%, 1000 = 100%)
    pct: Int!
}

# NetworkNodeGroupList represents a list of network node groups with a specified detail level.
type NetworkNodeGroupList {
    # level represents the level of detail of the aggregation group list.
    level: NetworkNodeGroupLevel!

    # totalCount represents the total number of nodes in the list.
    totalCount: Int!

    # groups represents an array of groups in the list.
    groups: [NetworkNodeGroup!]!
}

# Block is an Opera block chain block.
type Block {
    # Number is the number of this block, starting at 0 for the genesis block.
    number: Long!

    # Hash is the unique block hash of this block.
    hash: Bytes32!

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
    txHashList: [Bytes32!]!

    # txList is a list of transactions assigned to the block.
    txList: [Transaction!]!
}

# ERC721Contract represents a generic ERC721 non-fungible tokens (NFT) contract.
type ERC721Contract {
    # address of the token is used as the token's unique identifier.
    address: Address!

    # name of the token.
    name: String!

    # symbol used as an abbreviation for the token.
    symbol: String!

    # totalSupply represents total amount of tokens across all accounts
    totalSupply: BigInt

    # balanceOf represents amount of tokens on the account.
    balanceOf(owner: Address!): BigInt!

    # tokenURI provides URI of Metadata JSON Schema of the token.
    tokenURI(tokenId: BigInt!): String

    # ownerOf provides the owner of NFT identified by tokenId
    ownerOf(tokenId: BigInt!): Address

    # getApproved provides the operator approved by owner
    getApproved(tokenId: BigInt!): Address

    # isApprovedForAll queries the approval status of an operator for a given owner.
    isApprovedForAll(owner: Address!, operator: Address!): Boolean
}

# SfcConfig represents the configuration of the SFC contract
# responsible for managing the staking economy of the network.
type SfcConfig {
    # minValidatorStake is the minimal amount of tokens required
    # to register a validator account with the default self stake.
    minValidatorStake: BigInt!

    # maxDelegatedRatio is the maximal ratio between a validator self stake
    # and the sum of all the received stakes of the validator.
    # The value is provided as a multiplier number with 18 decimals.
    maxDelegatedRatio: BigInt!

    # minLockupDuration is the lowest possible number of seconds
    # a delegation can be locked for.
    minLockupDuration: BigInt!

    # maxLockupDuration is the highest possible number of seconds
    # a delegation can be locked for.
    maxLockupDuration: BigInt!

    # withdrawalPeriodEpochs is the minimal number of epochs
    # between an un-delegation and corresponding withdraw request.
    # The delay is enforced on withdraw call.
    withdrawalPeriodEpochs: BigInt!

    # withdrawalPeriodTime is the minimal number of seconds
    # between an un-delegation and corresponding withdraw request.
    # The delay is enforced on withdraw call.
    withdrawalPeriodTime: BigInt!
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

    # totalDeposited represents total amount of deposited tokens collateral on fMint.
    totalDeposit: BigInt!

    # totalDebt represents total amount of borrowed/minted tokens on fMint.
    totalDebt: BigInt!
}

# DelegationList is a list of delegations edges provided by sequential access request.
type DelegationList {
    "Edges contains provided edges of the sequential list."
    edges: [DelegationListEdge!]!

    """
    TotalCount is the maximum number of delegations
    available for sequential access.
    """
    totalCount: Long!

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

# Delegation represents a delegation on Opera block chain.
type Delegation {
    # Address of the delegator account.
    address: Address!

    # Identifier of the staker the delegation belongs to.
    toStakerId: BigInt!

    # Notifies the client that this stake is actually a self-stake
    # of the validator.
    isSelfStake: Boolean!

    # Time stamp of the delegation creation.
    createdTime: Long!

    # Amount delegated in WEI. The value includes all the pending un-delegations.
    amount: BigInt!

    # Current active amount delegated in WEI.
    amountDelegated: BigInt!

    # Amount locked in pending un-delegations in WEI.
    amountInWithdraw: BigInt!

    # Total amount of rewards claimed.
    claimedReward: BigInt!

    # Pending rewards for the delegation in WEI.
    pendingRewards: PendingRewards!

    # List of withdraw requests of the delegation,
    # sorted fro the newest to the oldest requests.
    withdrawRequests(cursor: Cursor, count: Int = 50, activeOnly: Boolean = false): [WithdrawRequest!]!

    # rewardClaims provides a list of reward claims
    # of the delegation as a scrollable list of edges with details of claims.
    rewardClaims(cursor: Cursor, count: Int = 25): RewardClaimList!

    # isFluidStakingActive indicates if the delegation is upgraded to fluid staking.
    isFluidStakingActive: Boolean!

    # isDelegationLocked indicates if the delegation is locked.
    isDelegationLocked: Boolean!

    # lockedFromEpoch represents the id of epoch the lock has been created.
    lockedFromEpoch: Long!

    # lockDuration represents the duration the lock has been placed for.
    lockDuration: Long!

    # lockedUntil represents the time stamp up to which
    # the delegation is locked, zero if not locked.
    lockedUntil: Long!

    # lockedAmount represents the amount of delegation stake locked.
    # The undelegate process must call unlock prior to creating withdraw
    # request if outstanding unlocked amount
    # is lower than demanded amount to undelegate.
    lockedAmount: BigInt!

    # unlockedAmount represents the amount
    # of delegation stake available for undelegate.
    unlockedAmount: BigInt!

    # unlockPenalty provides the mount of penalty applied
    # to the stake amount on premature unlock
    unlockPenalty(amount: BigInt!): BigInt!

    # outstandingSFTM represents the amount of sFTM tokens representing
    # the tokenized stake minted and un-repaid on this delegation.
    outstandingSFTM: BigInt!

    # tokenizerAllowedToWithdraw indicates if the stake tokenizer allows the stake
    # to be withdrawn. That means all the sFTM tokens have been repaid and the sFTM
    # debt is effectively zero for the delegation.
    tokenizerAllowedToWithdraw: Boolean!
}

# PendingRewards represents a detail of pending rewards for staking and delegations
type PendingRewards {
    # address of the delegation the reward belongs to.
    address: Address!

    # Staker the pending reward relates to.
    staker: BigInt!

    # Pending rewards amount.
    amount: BigInt!

    # The first unpaid epoch. Is not used for SFCv3.
    fromEpoch: Long!

    # The last unpaid epoch. Is not used for SFCv3.
    toEpoch: Long!

    # isOverRange signals that the rewards calculation
    # can not be done due to too many unclaimed epochs.
    # Is not used for SFCv3.
    isOverRange: Boolean!
}

# Represents epoch information.
type Epoch {
    # Identifier of the epoch.
    id: Long!

    # Timestamp of the epoch end.
    endTime: Long!

    # Epoch duration in seconds.
    duration: Long!

    # Fee at the epoch.
    epochFee: BigInt!

    # Total base reward weight on epoch.
    totalBaseRewardWeight: BigInt!

    # Total transaction reward weight on epoch.
    totalTxRewardWeight: BigInt!

    # Base reward per second of epoch.
    baseRewardPerSecond: BigInt!

    # Total amount staked. This includes all the staked
    # amount including validators' self stake.
    stakeTotalAmount: BigInt!

    # Total supply amount.
    totalSupply: BigInt!
}

# ERC721TransactionList is a list of ERC721 transaction edges provided by sequential access request.
type ERC721TransactionList {
    # Edges contains provided edges of the sequential list.
    edges: [ERC721TransactionListEdge!]!

    # TotalCount is the maximum number of ERC721 transactions available for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page of ERC721 transaction edges.
    pageInfo: ListPageInfo!
}

# TransactionListEdge is a single edge in a sequential list of ERC721 transactions.
type ERC721TransactionListEdge {
    cursor: Cursor!
    trx: ERC721Transaction!
}

# Contract defines block-chain smart contract information container
type Contract {
    "Address represents the contract address."
    address: Address!

    "DeployedBy represents the smart contract deployment transaction reference."
    deployedBy: Transaction!

    "transactionHash represents the smart contract deployment transaction hash."
    transactionHash: Bytes32!

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

# ERC20TransactionList is a list of ERC20 transaction edges provided by sequential access request.
type ERC20TransactionList {
    # Edges contains provided edges of the sequential list.
    edges: [ERC20TransactionListEdge!]!

    # TotalCount is the maximum number of ERC20 transactions available for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page of ERC20 transaction edges.
    pageInfo: ListPageInfo!
}

# TransactionListEdge is a single edge in a sequential list of ERC20 transactions.
type ERC20TransactionListEdge {
    cursor: Cursor!
    trx: ERC20Transaction!
}

# Bytes32 is a 32 byte binary string, represented by 0x prefixed hexadecimal hash.
scalar Bytes32

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

# Time represents date and time including time zone information in RFC3339 format.
scalar Time

# CurrentState represents the current active state
# of the chain information condensed on one place.
type CurrentState {
    # epoch is the last sealed Epoch structure
    sealedEpoch: Epoch!

    # blocks represents number of blocks in the chain.
    blocks: BigInt!

    # transactions represents number of transactions in the chain.
    transactions: Long!

    # validators represents number of validators in the network.
    validators: Long!

    # accounts represents number of accounts participating on transactions.
    accounts: Long!

    # sfcVersion indicates the current version of the SFC contract.
    # The version is encoded into 3 bytes representing ASCII version numbers
    # with the most significant byte first [<8bit major><8bit minor><8bit revision>].
    # I.e. Version 1.0.2 = "102" = 0x313032
    sfcVersion: Long!

    # sfcContractAddress is the address of the SFC contract
    # used for PoS staking control.
    sfcContractAddress: Address!

    # sfcLockingEnabled indicates if the SFC locking feature is enabled.
    sfcLockingEnabled: Boolean!
}
# UniswapActionList is a list of uniswap action edges provided by sequential access request.
type UniswapActionList {
    # Edges contains provided edges of the sequential list.
    edges: [UniswapActionListEdge!]!

    # TotalCount is the maximum number of uniswap actions available for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page of uniswap action edges.
    pageInfo: ListPageInfo!
}

# UniswapActionListEdge is a single edge in a sequential list of uniswap actions.
type UniswapActionListEdge {
    cursor: Cursor!
    uniswapAction: UniswapAction!
}

# UniswapAction represents a Uniswap action - swap, mint, burn
type UniswapAction {

    # id of the action in the persistent db
    id: Bytes32!

    # UniswapPair represents the information about single
    # Uniswap pair managed by the Uniswap Core.
    uniswapPair: UniswapPair!

    # pairAddress is address of the action's uniswap pair
    pairAddress: Address!

    # transactionHash represents the hash for this acstion transaction
    transactionHash: Bytes32!

    # sender is address of action owner account
    sender: Address!

    # type represents action type:
    # 0 - swap
    # 1 - mint
    # 2 - burn
    type: Int!

    # blockNr is number of the block for this action
    blockNr: Long!

    # Time represents UTC ISO time tag for this reserve value
    time: Long!

    # amount0in is amount of incoming tokens for Token0 in this action
    amount0in: BigInt!

    # amount0out is amount of outgoing tokens for Token0 in this action
    amount0out: BigInt!

    # amount1in is amount of In tokens for Token1 in this action
    amount1in: BigInt!

    # amount1out is amount of outgoing tokens for Token1 in this action
    amount1out: BigInt!
}

# Represents staker information.
type Staker {
    # ID number the staker.
    id: BigInt!

    # Staker address.
    stakerAddress: Address!

    # Amount of total staked tokens in WEI.
    totalStake: BigInt

    # Amount of own staked tokens in WEI.
    stake: BigInt!

    # Amount of tokens delegated to the staker in WEI.
    delegatedMe: BigInt!

    # Maximum total amount of tokens allowed to be delegated
    # to the staker in WEI.
    # This value depends on the amount of self staked tokens.
    totalDelegatedLimit: BigInt!

    # Maximum amount of tokens allowed to be delegated to the staker
    # on a new delegation in WEI.
    # This value depends on the amount of self staked tokens.
    delegatedLimit: BigInt!

    # Is the staker active.
    isActive: Boolean!

    # Is TRUE for validators withdrawing their validation stake.
    isWithdrawn: Boolean!

    # Is the staker considered to be cheater.
    isCheater: Boolean!

    # Is the staker offline.
    isOffline: Boolean!

    # isStakeLocked signals if the staker locked the stake.
    isStakeLocked: Boolean!

    # Epoch in which the staker was created.
    createdEpoch: Long!

    # Timestamp of the staker creation.
    createdTime: Long!

    # lockedFromEpoch is the identifier of the epoch the stake lock was created.
    lockedFromEpoch: Long!

    # lockedUntil is the timestamp up to which the stake is locked, zero if not locked.
    lockedUntil: Long!

    # Epoch in which the staker was deactivated.
    deactivatedEpoch: Long!

    # Timestamp of the staker deactivation.
    deactivatedTime: Long!

    # How many blocks the staker missed.
    missedBlocks: Long!

    # Number of seconds the staker is offline.
    downtime: Long!

    # List of delegations of this staker. Cursor is used to obtain specific slice
    # of the staker delegations. The most recent delegations
    # are provided if cursor is omitted.
    delegations(cursor: Cursor, count: Int = 25):DelegationList!

    # Status is a binary encoded status of the staker.
    # Ok = 0, bin 1 = Fork Detected, bin 256 = Validator Offline
    status: Long!

    # StakerInfo represents extended staker information from smart contract.
    stakerInfo: StakerInfo
}

# StakerFlagFilter represents a filter type for stakers with the given flag.
enum StakerFlagFilter {
    IS_ACTIVE
    IS_WITHDRAWN
    IS_OFFLINE
    IS_CHEATER
}

# ERC1155TransactionList is a list of ERC1155 transaction edges provided by sequential access request.
type ERC1155TransactionList {
    # Edges contains provided edges of the sequential list.
    edges: [ERC1155TransactionListEdge!]!

    # TotalCount is the maximum number of ERC1155 transactions available for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page of ERC1155 transaction edges.
    pageInfo: ListPageInfo!
}

# TransactionListEdge is a single edge in a sequential list of ERC1155 transactions.
type ERC1155TransactionListEdge {
    cursor: Cursor!
    trx: ERC1155Transaction!
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

# FtmBlockBurn represents a native FTM tokens burn record per created block.
type FtmBlockBurn {
    # blockNumber represents the number of the block.
    blockNumber: Long!

    # Timestamp is the unix timestamp at which this block was created.
    timestamp: Long!

    # amount represents the amount of FTM tokens burned in WEI units (18 digits fixed INT encoded as HEX number).
    amount: BigInt!

    # ftmValue represents FTM value of the burned FTM tokens.
    ftmValue: Float!
}
# EstimatedRewards represents a calculated rewards estimation for an account or amount staked
type EstimatedRewards {
    # Amount of FTM tokens expected to be staked for the calculation.
    staked: Long!

    # dailyReward represents amount of FTM tokens estimated
    # to be rewarded for staked amount in average per day.
    dailyReward: BigInt!

    # weeklyReward represents amount of FTM tokens estimated
    # to be rewarded for staked amount in average per week.
    weeklyReward: BigInt!

    # monthlyReward represents amount of FTM tokens estimated
    # to be rewarded for staked amount in average per month.
    monthlyReward: BigInt!

    # yearlyReward represents amount of FTM tokens estimated
    # to be rewarded for staked amount in average per year.
    yearlyReward: BigInt!

    # currentRewardYearRate represents average reward rate
    # for any staked amount in average per year.
    # The value is calculated as linear gross proceeds for staked amount
    # of tokens yearly.
    currentRewardRateYearly: Int!

    # Total amount of staked FTM tokens used for the calculation in WEI units.
    # The estimation uses total staked amount, not the effective amount provided
    # by the last epoch. The effective amount does not include current
    # un-delegations and also skips offline self-staking and flagged staking.
    totalStaked: BigInt!

    # Information about the last sealed epoch of the Opera blockchain.
    # The epoch provides useful information about total FTM supply,
    # total amount staked, rewards rate and weight, fee, etc.
    lastEpoch: Epoch!
}

# WithdrawRequest represents a request for partial stake withdraw.
type WithdrawRequest {
    # Cursor is the internal cursor ID of the withdraw request.
    id: Cursor!

    # Address of the authorized request.
    address: Address!

    # Account of the authorized request.
    account: Account!

    # StakerID represents the identifier of the validator
    # the withdraw request points to.
    stakerID: BigInt!

    # Details of the staker involved in the withdraw request.
    staker: Staker!

    # Unique withdraw request identifier.
    withdrawRequestID: BigInt!

    # Amount of tokens to be withdrawn in WEI.
    amount: BigInt!

    # CreatedTime represents the time stamp of the request creation.
    createdTime: Long!

    # WithdrawTime represents the time stamp of the request finalization.
    # If the request is pending, the withdrawTime will be NULL.
    withdrawTime: Long
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

    # lastKValue represents the last coefficient
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


# DefiUniswapVolume represents a calculated volume for swap pairs in history
type DefiUniswapVolume {

    # UniswapPair represents the information about single
    # Uniswap pair managed by the Uniswap Core.
    uniswapPair: UniswapPair!

    # pairAddress represents the Address of the Pair
    pairAddress: Address!

    # dailyVolume returns swap volume for last 24 hours
    dailyVolume: BigInt!

    # weeklyVolume returns swap volume for last 7 days
    weeklyVolume: BigInt!

    # monthlyVolume returns swap volume for last month
    monthlyVolume: BigInt!

    # YearlyVolume returns swap volume for last year
    yearlyVolume: BigInt!

    # IsInFUSD indicates if TokenA from the pair has a price value to be able
    # to calculate value in fUSD
    isInFUSD: Boolean!

}

# DefiSwaps represents swap volume for given pair and time interval
type DefiTimeVolume {

    # pairAddress represents the Address of the Pair
    pairAddress: Address!

    # time indicates a period for this volume
    time: String!

    # value represents amount of the volume
    value: BigInt!
}

# DefiTimePrice represents a calculated price for swap pairs in history
type DefiTimePrice {

	# pairAddress represents the Address of the Pair
    pairAddress: Address!

    # time indicates a period for this price
    time: String!

    # opening price for this time period
    open: Float!

    # closing price for this time period
	close: Float!

    # lowest price for this time period
	low: Float!

    # highest price for this time period
	high: Float!

    # average price for this time period
    average: Float!
}

# DefiTimeReserve represents a Uniswap pair reserve in history
type DefiTimeReserve {

    # UniswapPair represents the information about single
    # Uniswap pair managed by the Uniswap Core.
    uniswapPair: UniswapPair!

    # Time represents UTC ISO time tag for this reserve values
    time: String!

    # ReserveClose is close reserve for this time period
	# for both tokens. Index inside the array corresponds
    # with the token position.
    reserveClose: [BigInt!]!
}
# LendingPool represents a lendingpool instance.
type LendingPool {

    # Returns all assets reserve addresses
    reserveList: [Address!]!

    # A list of all assets reserves with its data
    reserveDataList: [ReserveData!]!

    # Asset reserve data just for one asset address
	reserveData(address: Address!): ReserveData!

    # User account data for specified user address
    userAccountData(address: Address!): FLendUserData!

    # User account deposit event history data
    userDepositHistory(address: Address, asset: Address): [FLendDeposit!]!
}

# ReserveData represents a lendingpool asset data.
# Unit Ray is 1e27.
type ReserveData {

    # address of the asset
    assetAddress: Address!

    # number in the reserveList() array
    ID: Int!

    # bitmask encoded asset reserve configuration data
    configuration: BigInt!

    # liquidity index in ray
    liquidityIndex: BigInt!

    # variable borrow index in ray
    variableBorrowIndex: BigInt!

    # current supply / liquidity / deposit rate in ray
    currentLiquidityRate: BigInt!

    # current variable borrow rate in ray
    currentVariableBorrowRate: BigInt!

    # current stable borrow rate in ray
    currentStableBorrowRate: BigInt!

    # timestamp of when reserve data was last updated
    lastUpdateTimestamp: BigInt!

    # address of associated aToken (tokenised deposit)
    aTokenAddress: Address!

    # address of associated stable debt token
	stableDebtTokenAddress: Address!

    # address of associated variable debt token
	variableDebtTokenAddress: Address!

    # address of interest rate strategy
    interestRateStrategyAddress: Address!
}


# FLendUserData represents a lendingpool user data.
type FLendUserData {

    # total collateral in FUSD of the user
	totalCollateralFUSD: BigInt!

    # total debt in FUSD of the user
	totalDebtFUSD: BigInt!

    # borrowing power left of the user in FUSD
	availableBorrowsFUSD: BigInt!

    # liquidation threshold of the user
	currentLiquidationThreshold: BigInt!

    # Loan To Value of the user
	ltv: BigInt!

    # current health factor of the user
	healthFactor: BigInt!

    # configuration data
    configurationData: BigInt!
}

# FLendDeposit represents a lendingpool deposit event data.
type FLendDeposit {

    # address of the asset
	assetAddress: Address!

	# address of the user
	userAddress: Address!

    # address of the on behalf of
	onBehalfOfAddress: Address!

	# deposit amount
	amount: BigInt!

	# referral code
	referralCode: Int!

    # time of deposit
    timestamp: Long!
}

# FLendBorrow represents a lending pool borrow event data.
type FLendBorrow {
    # address of the asset
	assetAddress: Address!

	# address of the user
	userAddress: Address!

    # address of the on behalf of
	onBehalfOfAddress: Address!

	# deposit amount
	amount: BigInt!

    # interest rate mode
    interestRateMode: Int!

    # borrow rate
    borrowRate: Int!

	# referral code
	referralCode: Int!

    # time of deposit
    timestamp: Long!
}
# RewardClaim represents
type RewardClaim {
    # address represents the address of the delegator
    address: Address!

    # toStakerId represents the ID of the validator the delegation
    # is placed on
    toStakerId: BigInt!

    # claimed represents the time stamp of the reward claim
    # in Unix Epoch units, e.g. number of seconds from the Unix Epoch start.
    claimed: Long!

    # amount represents the amount of tokens rewarded on the claim.
    amount: BigInt!

    # isRestaked signals if the claim was added to the delegation
    # effectively increasing the staked amount and raising the delegation value.
    isRestaked: Boolean!

    # trxHash is the hash pf the transaction calling for the rewards
    # to be processed and granted.
    trxHash: Bytes32!
}
# ERC1155Transaction represents a transaction on an ERC1155 NFT token.
type ERC1155Transaction {
    # trxHash represents a hash of the transaction
    # executing the ERC1155 call.
    trxHash: Bytes32!

    # transaction represents the transaction
    # executing the ERC1155 call.
    transaction: Transaction!

    # trxIndex represents the index
    # of the ERC1155 call in the transaction logs.
    trxIndex: Long!

    # tokenAddress represents the address
    # of the ERC1155 token contract.
    tokenAddress: Address!

    # token represents the ERC1155 contract detail involved.
    token: ERC1155Contract!

    # tokenId represents the NFT token - one ERC1155 contract can handle multiple NFTs.
    tokenId: BigInt!

    # trxType is the type of the transaction.
    trxType: TokenTransactionType!

    # sender represents the address of the token owner
    # sending the tokens, e.g. the sender.
    sender: Address!

    # recipient represents the address of the token recipient.
    recipient: Address!

    # amount represents the amount of tokens involved in the transaction;
    # please make sure to interpret the amount with the correct number of decimals
    # from the token Metadata JSON Schema.
    amount: BigInt!

    # timeStamp represents the Unix epoch time stamp
    # of the ERC1155 transaction processing.
    timeStamp: Long!
}
# FeeFlowDaily represents daily aggregated flow of the transaction fee distribution in native FTM.
type FeeFlowDaily {
    # date is the signature date of the data point, the time part is set to 00:00:00Z.
    date: Time!

    # blocksCount represents the number of blocks included in the data point.
    blocksCount: Int!

    # fee is the amount of FTM collected from transaction fees;
    # represented as fixed point decimal of FTM with 9 digits.
    fee: Long!

    # feeFTM is the amount of FTM collected from transaction fees;
    # represented as floating point value in FTM units.
    feeFTM: Float!

    # burned is the amount of FTM burned;
    # represented as fixed point decimal with 9 digits.
    burned: Long!

    # burnedFTM is the amount of FTM burned;
    # represented as floating point value in FTM units.
    burnedFTM: Float!

    # treasury is the amount of FTM sent to treasury;
    # represented as fixed point decimal with 9 digits.
    treasury: Long!

    # treasuryFTM is the amount of FTM sent to treasury;
    # represented as floating point value in FTM units.
    treasuryFTM: Float!

    # rewards is the amount of FTM sent to rewards distribution;
    # represented as fixed point decimal with 9 digits.
    # -----------------------------------------------------------------------------------------------------------
    # Please note this is the max amount of rewards available for distribution. The actual amount is scaled
    # down based on locking period of individual stakers and the real amount distributed by the SFC contract
    # will be lower in most cases. The remaining reward tokens after the scaling are effectively also burned
    # and removed from the total supply, but this process is not reflected in this aggregated approximation.
    # Please see the current Fantom SFC contract implementation for the rewards distribution details.
    rewards: Long!

    # rewardsFTM is the amount of FTM sent to rewards distribution;
    # represented as floating point value in FTM units.
    rewardsFTM: Float!
}
# ERC721Transaction represents a transaction on an ERC721 NFT token.
type ERC721Transaction {
    # trxHash represents a hash of the transaction
    # executing the ERC721 call.
    trxHash: Bytes32!

    # transaction represents the transaction
    # executing the ERC721 call.
    transaction: Transaction!

    # trxIndex represents the index
    # of the ERC721 call in the transaction logs.
    trxIndex: Long!

    # tokenAddress represents the address
    # of the ERC721 token contract.
    tokenAddress: Address!

    # token represents the ERC721 contract detail involved.
    token: ERC721Contract!

    # tokenId represents the NFT token - one ERC721 contract can handle multiple NFTs.
    tokenId: BigInt!

    # trxType is the type of the transaction.
    trxType: TokenTransactionType!

    # sender represents the address of the token owner
    # sending the tokens, e.g. the sender.
    sender: Address!

    # recipient represents the address of the token recipient.
    recipient: Address!

    # amount represents the amount of tokens involved
    # in the transaction; please make sure to interpret the amount
    # with the correct number of decimals from the ERC721 token detail.
    amount: BigInt!

    # timeStamp represents the Unix epoch time stamp
    # of the ERC721 transaction processing.
    timeStamp: Long!
}
# TokenTransaction represents a generic token transaction
# of a supported type of token.
type TokenTransaction {
    # Hash is the hash of the executed transaction call.
    hash: Bytes32!

    # trxIndex is the index of the transaction call in a block.
    trxIndex: Long!

    # blockNumber represents the number of the block
    # the transaction was executed in.
    blockNumber: Long!

    # tokenAddress represents the address of the token involved.
    tokenAddress: Address!

    # tokenName represents the name of the token contract.
    # Is empty, if not provided for the given token.
    tokenName: String!

    # tokenSymbol represents the symbol of the token contract.
    # Is empty, if not provided for the given token.
    tokenSymbol: String!

    # tokenType represents the type of the token (i.e. ERC20/ERC721/ERC1155).
    tokenType: String!

    # tokenDecimals is the number of decimals the token supports.
    # The most common value is 18 to mimic the ETH to WEI relationship.
    tokenDecimals: Int!

    # type represents the type of the transaction executed (i.e. Transfer/Mint/Approval).
    type: String!

    # sender of the transaction.
    sender: Address!

    # recipient of the transaction.
    recipient: Address!

    # amount of tokens involved in the transaction.
    amount: BigInt!

    # multi-token contracts (ERC-721/ERC-1155) token ID involved in the transaction.
    tokenId: BigInt!

    # time stamp of the block processing.
    timeStamp: Long!
}

# Account defines block-chain account information container
type Account {
    # Address is the address of the account.
    address: Address!

    # Balance is the current balance of the Account in WEI.
    balance: BigInt!

    # TotalValue is the current total value of the account in WEI.
    # It includes available balance, delegated amount and pending rewards.
    # NOTE: This values is slow to calculate.
    totalValue: BigInt!

    # txCount represents number of transaction sent from the account (Nonce).
    txCount: Long!

    # txList represents list of transactions of the account in form of TransactionList.
    txList(recipient: Address, cursor:Cursor, count:Int!): TransactionList!

    # erc20TxList represents list of ERC20 transactions of the account.
    erc20TxList(cursor:Cursor, count:Int = 25, token: Address, txType: [TokenTransactionType!]): ERC20TransactionList!

    # erc721TxList represents list of ERC721 transactions of the account.
    erc721TxList(cursor:Cursor, count:Int = 25, token: Address, tokenId: BigInt, txType: [TokenTransactionType!]): ERC721TransactionList!

    # erc1155TxList represents list of ERC1155 transactions of the account.
    erc1155TxList(cursor:Cursor, count:Int = 25, token: Address, tokenId: BigInt, txType: [TokenTransactionType!]): ERC1155TransactionList!

    # Details of a staker, if the account is a staker.
    staker: Staker

    # List of delegations of the account, if the account is a delegator.
    delegations(cursor:Cursor, count:Int = 25): DelegationList!

    # Details about smart contract, if the account is a smart contract.
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

    # owner represents the owning wallet of the proposal
    owner: Address!

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
    # of all voters allowed on the proposal. This is effectively
    # the maximum weight an option can gain if all the voters
    # would favor it with the top value of the scale.
    totalWeight: BigInt!

    # votedWeightRatio represents the percentage of the total voting weight
    # already counted towards the proposal options. The ratio increases
    # as more voters place their votes.
    # The value is normalised to 1 digit precision, to get a percentage
    # you need to divide the value by 10.
    # The value is zero if no vote was placed. The value is 1000
    # if all the voters placed their votes either directly,
    # or through a vote delegation mechanism.
    # Please note the value is an estimation. The voting status
    # does not closely reflect changes in the total voting power,
    # especially after the voting is closed.
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

    # sfcConfig provides the current configuration
    # of the SFC contract managing the block chain staking economy.
    sfcConfig: SfcConfig!

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
    block(number:Long, hash: Bytes32):Block

    # Get list of Blocks with at most <count> edges.
    # If <count> is positive, return edges after the cursor,
    # if negative, return edges before the cursor.
    # For undefined cursor, positive <count> starts the list from top,
    # negative <count> starts the list from bottom.
    blocks(cursor:Cursor, count:Int!):BlockList!

    # Get transaction information for given transaction hash.
    transaction(hash:Bytes32!):Transaction

    # Get list of Transactions with at most <count> edges.
    # If <count> is positive, return edges after the cursor,
    # if negative, return edges before the cursor.
    # For undefined cursor, positive <count> starts the list from top,
    # negative <count> starts the list from bottom.
    transactions(cursor:Cursor, count:Int!):TransactionList!

    # Get filtered list of ERC20 Transactions.
    erc20Transactions(cursor:Cursor, count:Int = 25, token: Address, account: Address, txType: [TokenTransactionType!]): ERC20TransactionList!

    # Get filtered list of ERC721 Transactions.
    erc721Transactions(cursor:Cursor, count:Int = 25, token: Address, tokenId: BigInt, account: Address, txType: [TokenTransactionType!]): ERC721TransactionList!

    # Get filtered list of ERC1155 Transactions.
    erc1155Transactions(cursor:Cursor, count:Int = 25, token: Address, tokenId: BigInt, account: Address, txType: [TokenTransactionType!]): ERC1155TransactionList!

    # Get the id of the current epoch of the Opera blockchain.
    currentEpoch:Long!

    # Get information about specified epoch. Returns current epoch information
    # if id is not provided.
    epoch(id: Long): Epoch!

    # Get a scrollable list of epochs sorted from the last one back by default.
    epochs(cursor: Cursor, count: Int = 25): EpochList!

    # The last staker id in Opera blockchain.
    lastStakerId: Long!

    # The number of stakers in Opera blockchain.
    stakersNum: Long!

    # Staker information. The staker is loaded either by numeric ID,
    # or by address. null if none is provided.
    staker(id: BigInt, address: Address): Staker

    # List of staker information from SFC smart contract.
    stakers: [Staker!]!

    # stakersWithFlag provides list of staker information from SFC smart contract
    # for staker with the given flag set to TRUE. This can be used to obtain a subset
    # of stakers in a given state of staking process.
    stakersWithFlag(flag: StakerFlagFilter!): [Staker!]!

    # The list of delegations for the given staker ID.
    # Cursor is used to obtain specific slice of the staker delegations.
    # The most recent delegations are provided if cursor is omitted.
    delegationsOf(staker:BigInt!, cursor: Cursor, count: Int = 25): DelegationList!

    # Get the details of a specific delegation by it's delegator address
    # and staker the delegation belongs to.
    delegation(address:Address!, staker: BigInt!): Delegation

    # Get the list of all delegations by it's delegator address.
    delegationsByAddress(address:Address!, cursor: Cursor, count: Int = 25): DelegationList!

    # Returns the current price per gas in WEI units.
    gasPrice: Long!

    # estimateGas returns the estimated amount of gas required
    # for the transaction described by the parameters of the call.
    estimateGas(from: Address, to: Address, value: BigInt, data: String): Long

    # Get price details of the Opera blockchain token for the given target symbols.
    price(to:String!):Price!

    # Get calculated staking rewards for an account or given
    # staking amount in FTM tokens.
    # At least one of the address and amount parameters must be provided.
    # If you provide both, the address takes precedence and the amount is ignored.
    estimateRewards(address:Address, amount:Long):EstimatedRewards!

    # sfcRewardsCollectedAmount provides an amount of rewards collected based on given
    # filtering options, which are all optional. If no filter option is passed,
    # the total amount of collected rewards is being presented.
    sfcRewardsCollectedAmount(delegator: Address, staker: BigInt, since: Long, until: Long): BigInt!

    # defiConfiguration exposes the current DeFi contract setup.
    defiConfiguration:DefiSettings!

    # defiTokens represents a list of all available DeFi tokens.
    defiTokens:[DefiToken!]!

    # defiNativeToken represents the information about the native token
    # wrapper ERC20 contract. Returns NULL if the native token wrapper
    # is not available.
    defiNativeToken: ERC20Token

    # fMintAccount provides DeFi/fMint information about an account on fMint protocol.
    fMintAccount(owner: Address!):FMintAccount!

    # fMintTokenAllowance resolves the amount of ERC20 tokens unlocked
    # by the token owner for DeFi/fMint operations.
    fMintTokenAllowance(owner: Address!, token: Address!):BigInt!

    # fMintUserTokens resolves a list of pairs of fMint users and their tokens
    # used for a specified purpose.
    fMintUserTokens(purpose:FMintUserTokenPurpose=FMINT_COLLATERAL):[FMintUserToken!]!

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

    # defiUniswapVolumes represents a list of pairs and their historical values
    # of traded volumes
    defiUniswapVolumes:[DefiUniswapVolume!]!

    # defiTimeVolumes returns volumes for specified pair, time resolution and interval.
    # Address is pair address and is mandatory.
    # Resolution can be {month, day, 4h, 1h, 30m 15m, 5m, 1m}, is optional, default is a day.
    # Dates are in unix UTC number and are optional. When not provided
    # then it takes period for last month till now.
    defiTimeVolumes(address:Address!, resolution:String, fromDate:Int, toDate:Int):[DefiTimeVolume!]!

    # defiTimePrices returns prices for specified pair, time resolution and interval.
    # Address is pair address and is mandatory.
    # Resolution can be {month, day, 4h, 1h, 30m 15m, 5m, 1m}, is optional, default is a day.
    # Direction specifies price calculation, default 0 is for TokenA/TokenB otherwise TokenB/TokenA
    # Dates are in unix UTC number and are optional. When not provided
    # then it takes period for last month till now.
    defiTimePrices(address:Address!, resolution:String, fromDate:Int, toDate:Int, direction:Int):[DefiTimePrice!]!

    # defiTimeReserves returns reserves for specified pair, time resolution and interval.
    # Address is pair address and is mandatory.
    # Resolution can be {month, day, 4h, 1h, 30m 15m, 5m, 1m}, is optional, default is a day.
    # Dates are in unix UTC number and are optional. When not provided
    # then it takes period for last month till now.
    defiTimeReserves(address:Address!, resolution:String, fromDate:Int, toDate:Int):[DefiTimeReserve!]!

    # Get list of Uniswap actions with at most <count> edges.
    # If <count> is positive, return edges after the cursor,
    # if negative, return edges before the cursor.
    # For undefined cursor, positive <count> starts the list from top,
    # negative <count> starts the list from bottom.
    # Address can be used for specifying actions for one Uniswap pair.
    # ActionType represents action type:
    # 0 - swap,
    # 1 - mint,
    # 2 - burn,
    defiUniswapActions(pairAddress:Address, cursor:Cursor, count:Int!, actionType:Int):UniswapActionList!

    # erc20Token provides the information about an ERC20 token specified by it's
    # address, if available. The resolver returns NULL if the token does not exist.
    erc20Token(token: Address!):ERC20Token

    # erc20TokenList provides list of the most active ERC20 tokens
    # deployed on the block chain.
    erc20TokenList(count: Int = 50):[ERC20Token!]!

    # erc20Assets provides list of tokens owned by the given
    # account address.
    erc20Assets(owner: Address!, count: Int = 50):[ERC20Token!]!

    # ercTotalSupply provides the current total supply amount of a specified ERC20 token
    # identified by it's ERC20 contract address.
    ercTotalSupply(token: Address!):BigInt!

    # ercTokenBalance provides the current available balance of a specified ERC20 token
    # identified by it's ERC20 contract address.
    ercTokenBalance(owner: Address!, token: Address!):BigInt!

    # ercTokenAllowance provides the current amount of ERC20 tokens unlocked
    # by the token owner for the spender to be manipulated with.
    ercTokenAllowance(token: Address!, owner: Address!, spender: Address!):BigInt!

    # erc721Contract provides the information about ERC721 non-fungible token (NFT) by it's address.
    erc721Contract(token: Address!):ERC721Contract

    # erc721ContractList provides list of the most active ERC721 non-fungible tokens (NFT) on the block chain.
    erc721ContractList(count: Int = 50):[ERC721Contract!]!

    # erc1155Token provides the information about ERC1155 multi-token contract by it's address.
    erc1155Contract(address: Address!):ERC1155Contract

    # erc1155ContractList provides list of the most active ERC1155 multi-token contract on the block chain.
    erc1155ContractList(count: Int = 50):[ERC1155Contract!]!

    # govContracts provides list of governance contracts.
    govContracts:[GovernanceContract!]!

    # govContract provides a specific Governance contract information by its address.
    govContract(address: Address!): GovernanceContract

    # govProposals represents list of joined proposals across all the Governance contracts.
    govProposals(cursor:Cursor, count:Int!, activeOnly: Boolean = false):GovernanceProposalList!

    # fLendLendingPool represents an instance of an fLend Lending pool
    fLendLendingPool: LendingPool!

    # trxVolume provides a list of daily aggregations of the network transaction flow.
    # If boundaries are not defined, last 90 days of aggregated trx flow is provided.
    # Boundaries are defined in format YYYY-MM-DD, i.e. 2021-01-23 for January 23rd, 2021.
    trxVolume(from:String, to:String):[DailyTrxVolume!]!

    # trxSpeed provides the recent speed of the network
    # as number of transactions processed per second
    # calculated for the given range denominated in secods. I.e. range:300 means last 5 minutes.
    # Minimal range is 60 seconds, any range below this value will be adjusted to 60 seconds.
    trxSpeed(range: Int = 1200): Float!

    # trxGasSpeed provides average gas consumed by transactions, either base or cumulative,
    # per second in the given date/time period. Please specify the ending date and time
    # as RFC3339 time stamp, i.e. 2021-05-14T00:00:00.000Z. The current time is used if not defined.
    # The range represents the number of seconds prior the end time stamp
    # we use to calculate the average gas consumption.
    trxGasSpeed(range: Int = 1200, to: String): Float!

    # gasPriceList provides a list of gas price ticks for the given date/time span.
    # If the end time is not specified, the list is provided up to the current date/time.
    # The maximal date/time span of the list is 30 days.
    gasPriceList(from: Time! to: Time): [GasPriceTick!]!

    # ftmBurnedTotal provides the total amount of native FTM tokens burned
    # by the chain from paid transaction fees in WEI units.
    ftmBurnedTotal: BigInt!

    # ftmBurnedTotalAmount provides the total amount of native FTM tokens burned
    # by the chain from paid transaction fees in FTM units.
    ftmBurnedTotalAmount: Float!

    # ftmLatestBlockBurnList provides a list of latest burned native FTM tokens per-block.
    ftmLatestBlockBurnList(count: Int = 25): [FtmBlockBurn!]!

    # dailyFeeFlow provides a list of fee distribution information aggregated by days.
    dailyFeeFlow(from: Time, to: Time): [FeeFlowDaily!]!

    # ftmTreasuryTotal provides the total amount of native FTM tokens sent into treasury
    # by the chain from paid transaction fees in WEI units.
    ftmTreasuryTotal: BigInt!

    # ftmTreasuryTotalAmount provides the total amount of native FTM tokens sent into treasury
    # by the chain from paid transaction fees in FTM units.
    ftmTreasuryTotalAmount: Float!

    # networkNodesAggregated provides an aggregated list of network nodes on the Opera network.
    networkNodesAggregated(level: NetworkNodeGroupLevel = COUNTRY): NetworkNodeGroupList!
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

package gqlschema

// Auto generated GraphQL schema bundle; created 2020-03-16 11:29
const schema = `
# Delegator represents a delegation on Opera blockchain.
type Delegator {
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

    "Amount delegated."
    amount: BigInt

    "Amount of rewards claimed."
    claimedReward: BigInt

    "Pending rewards for the delegation."
    pendingRewards: PendingRewards!
}

# Account defines block-chain account information container
type Account {
    "Address is the address of the account."
    address: Address!

    "Balance is the current balance of the Account in WEI."
    balance: BigInt!

    "txCount represents number of transaction sent from the account."
    txCount: Long!

    "txList represents list of transactions of the account in form of TransactionList."
    txList (cursor:Cursor, count:Int!): TransactionList!

    "Details of a staker, if the account is a staker."
    staker: Staker

    "Details of delegation, if the account is a delegator."
    delegation: Delegator
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

    # To is the account the transaction was sent to. This is null for contract creating transactions.
    to: Address

    # Recipient is the account that received this transaction. Null for contract creating transaction.
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
    inputData: Bytes!

    # BlockHash is the hash of the block this transaction was assigned to. Null if the transaction is pending.
    blockHash: Hash

    # BlockHash is the hash of the block this transaction was assigned to. Null if the transaction is pending.
    blockNumber: Long

    # Block is the block this transaction was assigned to. This will be null if
    # the transaction is pending.
    block: Block

    # Status is the return status of the transaction. This will be 1 if the
    # transaction succeeded, or 0 if it failed (due to a revert, or due to
    # running out of gas). If the transaction has not yet been processed, this
    # field will be null.
    status: Long
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

# Represents epoch information.
type Epoch {
    "Number the epoch end."
    id: Long!

    "Timestamp of the epoch end."
    endTime: BigInt!

    "Epoch duration in secods."
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

    "Amount of tokens delegated to the staker."
    delegatedMe: BigInt

    "Is this a validator record."
    isValidator: Boolean!

    "Is the staker active."
    isActive: Boolean!

    "Is the staker considered to be cheater."
    isCheater: Boolean!

    "Is the staker offline."
    isOffline: Boolean!

    "Epoch in which the staker was created."
    createdEpoch: Long!

    "Timestamp of the staker creation."
    createdTime: Long!

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

    "List of delegations of this staker."
    delegations:[Delegator!]!
}

# PendingRewards represents a detail of pending rewards for staking and delegations
type PendingRewards {
    "Pedning rewards amount."
    amount: BigInt!

    "The first unpaid epoch."
    fromEpoch: Long!

    "The last unpaid epoch."
    toEpoch: Long!
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

    # txHashList is the list of unique hash values of transaction
    # assigned to the block.
    txHashList: [Hash!]!

    # txList is a list of transactions assigned to the block.
    txList: [Transaction!]!
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

# ListPageInfo contains information about a sequential access list page.
type ListPageInfo {
    # First is the cursor of the first edge of the edges list. null for empty list.
    first: Cursor

    # Last if the cusrsor of the last edge of the edges list. null for empty list.
    last: Cursor

    # HasNext specifies if there is another edge after the last one.
    hasNext: Boolean!

    # HasNext specifies if there is another edge before the first one.
    hasPrevious: Boolean!
}
# Root schema definition
schema {
    query: Query
    subscription: Subscription
}

# Entry points for querying the API
type Query {
    "Get an Account information by hash address."
    account(address:Address!):Account!

    "Get block information by number or by hash. If neither is provided, the most recent block is given."
    block(number:Long, hash: Hash):Block

    "Get transaction information for given transaction hash."
    transaction(hash:Hash!):Transaction

    """
    Get list of Blocks with at most <count> edges.
    If <count> is positive, return edges after the cursor, if negative, return edges before the cursor.
    For udefined cursor, positive <count> starts the list from top, negative <count> starts the list from bottom.
    """
    blocks(cursor:Cursor, count:Int!):BlockList!

    """
    Get list of Transactions with at most <count> edges.
    If <count> is positive, return edges after the cursor, if negative, return edges before the cursor.
    For udefined cursor, positive <count> starts the list from top, negative <count> starts the list from bottom.
    """
    transactions(cursor:Cursor, count:Int!):TransactionList!

    "Get the id of the current epoch of the Opera blockchain."
    currentEpoch:Long!

    "Get information about specified epoch. Returns current epoch information if id is not provided."
    epoch(id: Long!): Epoch!

    "The last staker id in Opera blockchain."
    lastStakerId: Long!

    "The number of stakers in Opera blockchain."
    stakersNum: Long!

    "Staker information. The staker is loaded either by numberic ID, or by address. null if none is provided."
    staker(id: Long, address: Address): Staker

    "List of staker information from SFC smart contract."
    stakers: [Staker!]!

    "The list of delegations for the given staker ID."
    delegationsOf(staker:Long!): [Delegator!]!

    "Get the details of a delegator by it's address."
    delegation(address:Address!): Delegator
}

# Subscriptions to live events broadcasting
type Subscription {
    "Subscribe to receive information about new blocks in the blockchain."
    onBlock: Block!

    "Subscribe to receive information about new transactions in the blockchain."
    onTransaction: Transaction!
}

`

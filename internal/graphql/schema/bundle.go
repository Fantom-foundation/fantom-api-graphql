package gqlschema

// Auto generated GraphQL schema bundle; created , 2020-03-03 02:50
const schema = `
# Account defines block-chain account information container
type Account {
    "Address is the address of the account."
    address: Address!

    "Balance is the current balance of the Account in WEI."
    balance: BigInt!

    "Nonce is the number of transactions sent from this account."
    nonce: Long!
}

# Transaction is an Opera block chain transaction.
type Transaction {
    # Hash is the unique hash of this transaction.
    hash: Hash!

    # Nonce is the number of transactions sent by the account prior to this transaction.
    nonce: Long!

    # Index is the index of this transaction in the block. This will
    # be null if the transaction is in a pending pool.
    index: Int

    # From is the account that sent this transaction
    from: Account!

    # To is the account the transaction was sent to. This is null for
    # contract-creating transactions.
    to: Account

    # Value is the value sent along with this transaction in WEI.
    value: BigInt!

    # GasPrice is the price og gas per unit in WEI.
    gasPrice: BigInt!

    # GasLimit is the maximum amount of gas this transaction can consume.
    gasLimit: Long!

    # GasUsed is the amount of gas that was used on processing this transaction.
    # If the transaction is pending, this field will be null.
    gasUsed: Long

    # InputData is the data supplied to the target of the transaction.
    inputData: Bytes!

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
    block: Transaction!
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
    timestamp: BigInt!

    # TransactionHashList is the list of unique hash values of transaction
    # assigned to the block.
    transactionHashList: [Hash!]!

    # Transactions is a list of transactions assigned to the block.
    transactions: [Transaction!]!

    # TransactionAt returns the transaction at the specified index. If
    # the transaction is not available for this block, this field will be null.
    transactionAt(index: Int!): Transaction
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
    # First is the cursor of the first edge of the blocks list.
    first: Cursor!

    # Last if the cusrsor of the last edge of the blocks list.
    last: Cursor!

    # HasNext specifies if there is another edge after the last one.
    hasNext: Boolean

    # HasNext specifies if there is another edge before the first one.
    hasPrevious: Boolean
}
# Root schema definition
schema {
    query:Query
}

# Entry points for querying the API
type Query {
    "Get an Account information by hash address."
    account(address:Address!):Account!

    "Get current Block information."
    currentBlock:Block

    "Get block information for given Block number."
    block(number:Int!):Block

    "Get list of Blocks with at most count edges after the cursor. If cursor is not defined, start from top."
    blocks(cursor:ID, count:Int!):BlockList!

    "Get transaction information for given transaction hash."
    transaction(hash:ID!):Transaction

    "Get list of Transaction with at most count edges after the cursor. If cursor is not defined, start from top."
    transactions(cursor:ID, count:Int!):TransactionList!
}

`

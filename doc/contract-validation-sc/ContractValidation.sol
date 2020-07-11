pragma solidity ^0.5.0;

// @title Fantom deployed smart contract validation service.
contract ContractValidation {
    // ValidatedContract structure represents an information about deployed
    // smart contract, validated through one of our API nodes.
    struct ValidatedContract {
        string name;        // name of the contract
        bytes32 version;    // version information of the contract
        bytes32 license;    // license of the contract
        string contact;     // support contact of the contract
        bytes32 compiler;   // compiler used to compile/validate the contract
        bytes32 clVersion;  // version of the compiler used to compile the contract
        bool isOptimized;   // was the compiler run with "optimized" option
        int32 optRounds;    // optimization rounds used by the compiler to build the contract
        bytes32 srcHash;    // compressed source code hash
        uint validated;     // Unix timestamp of the contract validation
    }

    // admin is the address managing the contract validation container.
    address public admin;

    // validators are the addresses allowed to submit contract validations.
    mapping(address => bool) public validators;

    // contracts maps contract address to the validation record.
    mapping(address => ValidatedContract) public contracts;

    // ValidatorAdded is emitted when a new validator is added to the contract.
    event ValidatorAdded(address indexed addr, uint timestamp);

    // ValidatorDropped is emitted when a validator is removed from the contract.
    event ValidatorDropped(address indexed addr, uint timestamp);

    // ContractValidated signals contract validation for the given contract address.
    event ContractValidated(address indexed addr, bytes32 srcHash, uint timestamp);

    // ContractValidationDropped signals contract validation has been canceled.
    event ContractValidationDropped(address indexed addr, uint timestamp);

    // constructor creates new instance of the validation contract.
    constructor() public {
        // keep the contract creator as the admin
        admin = msg.sender;

        // add the admin as the first validator
        validators[msg.sender] = true;
    }

    // addValidators function adds new validators to the contract.
    function addValidators(address[] calldata list) external {
        // only administrator can do this
        require(msg.sender == admin, "only admin can add validators");

        // add new validators to the list
        for (uint i = 0; i < list.length; i++) {
            validators[list[i]] = true;
            emit ValidatorAdded(list[i], now);
        }
    }

    // dropValidators function removes specified validators from the contract.
    function dropValidators(address[] calldata list) external {
        // only administrator can do this
        require(msg.sender == admin, "only admin can drop validators");

        // delete the validators off the list
        for (uint i = 0; i < list.length; i++) {
            delete (validators[list[i]]);
            emit ValidatorDropped(list[i], now);
        }
    }

    // validate registers new validated deployed contract and stores the validation
    // information in so it can be used later.
    function validated(
        address addr,
        string calldata name,
        bytes32 version,
        bytes32 license,
        string calldata contact,
        bytes32 compiler,
        bytes32 clVersion,
        bool isOptimized,
        int32 optRounds,
        bytes calldata source) external
    {
        // only validators ca do this
        require(validators[msg.sender], "only validators can validate");

        // get the storage for the contract
        // we deliberately allow editing here, the information about previous
        // validation is still available on-chain in the event log.
        ValidatedContract storage inContract = contracts[addr];

        // update the contract details from incoming data
        inContract.name = name;
        inContract.version = version;
        inContract.license = license;
        inContract.contact = contact;
        inContract.compiler = compiler;
        inContract.clVersion = clVersion;
        inContract.isOptimized = isOptimized;
        inContract.optRounds = optRounds;
        inContract.validated = now;

        // calculate the hash for the source code
        // we don't store the source here, but it's still part of the call
        // so a re-check can extract it from the transaction,
        // verify the hash,if needed, and use it for local re-validation.
        inContract.srcHash = keccak256(source);

        // emit the event for new validated contract
        emit ContractValidated(addr, inContract.srcHash, now);
    }

    // dropValidation removes validation information from the contract.
    function dropValidation(address addr) external {
        // only administrator can drop validation
        require(msg.sender == admin, "only admin can drop validation");

        // drop the validation as requested
        delete (contracts[addr]);

        // emit an event about it
        emit ContractValidationDropped(addr, now);
    }
}

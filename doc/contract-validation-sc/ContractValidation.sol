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

    // validators are the addresses of the nodes allowed to submit contract validations.
    mapping(address => bool) public validators;

    // contracts maps contract address to the validation record.
    mapping(address => ValidatedContract) public contracts;

    // ContractValidated signals contract validation for the given contract address.
    event ContractValidated(address indexed addr, bytes32 srcHash);

    // ContractValidationDropped signals contract validation has been canceled.
    event ContractValidationDropped(address indexed addr);

    // constructor create new instance of the validation contract.
    constructor() public {
        // keep the contract creator as the admin
        admin = msg.sender;

        // add the admin as the first validator
        validators[msg.sender] = true;
    }

    // addValidator adds new validator address to the contract.
    function addValidator(address val) public {
        // only administrator can do this
        require(msg.sender == admin, "only admin can add validators");

        // add the new validator to the list
        validators[val] = true;
    }

    // dropValidator disables specified validator from being able to send new validations.
    function dropValidator(address val) public {
        // only administrator can do this
        require(msg.sender == admin, "only admin can drop validators");

        // delete the validator off the list
        delete (validators[val]);
    }

    // validate registers new validated deployed contract and stores the validation
    // information in so it can be used later.
    function validated(
        address addr,
        string memory name,
        bytes32 version,
        bytes32 license,
        string memory contact,
        bytes32 compiler,
        bytes32 clVersion,
        bool isOptimized,
        int32 optRounds,
        bytes memory source) public
    {
        // only validators ca do this
        require(validators[msg.sender], "only validators can validate");

        // get the storage for the contract
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

        // calculate the hash
        inContract.srcHash = keccak256(source);

        // emit the event for new validated contract
        emit ContractValidated(addr, inContract.srcHash);
    }

    // dropValidation removes validation information from the contract.
    function dropValidation(address addr) public {
        // only administrator can drop validation
        require(msg.sender == admin, "only admin can drop validation");

        // drop the validation as requested
        delete (contracts[addr]);

        // emit an event about it
        emit ContractValidationDropped(addr);
    }
}

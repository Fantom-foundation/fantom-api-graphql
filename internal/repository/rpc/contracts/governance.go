// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GovernanceMetaData contains all meta data concerning the Governance contract.
var GovernanceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"ProposalExecutionExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"ProposalRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"ProposalResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"TasksErased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startIdx\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endIdx\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"handled\",\"type\":\"uint256\"}],\"name\":\"TasksHandled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"VoteCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"VoteWeightOverridden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"VoteWeightUnOverridden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"choices\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastProposalID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxExecutionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxOptions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minVotesRatio\",\"type\":\"uint256\"}],\"name\":\"minVotesAbsolute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"overriddenWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposalBurntFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposalFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taskErasingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taskHandlingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governableContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proposalVerifier\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pType\",\"type\":\"uint256\"},{\"internalType\":\"enumProposal.ExecType\",\"name\":\"executable\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgreement\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"opinionScales\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"options\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"proposalContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingMinEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingMaxEndTime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionID\",\"type\":\"uint256\"}],\"name\":\"proposalOptionState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"agreementRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"agreement\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"winnerOptionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"getVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"choices\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tasksCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"assignment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"choices\",\"type\":\"uint256[]\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposalContract\",\"type\":\"address\"}],\"name\":\"createProposal\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"cancelProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"handleTasks\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"tasksCleanup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"calculateVotingTally\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"proposalResolved\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"winnerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"cancelVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"voterAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"recountVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceMetaData.ABI instead.
var GovernanceABI = GovernanceMetaData.ABI

// Governance is an auto generated Go binding around an Ethereum contract.
type Governance struct {
	GovernanceCaller     // Read-only binding to the contract
	GovernanceTransactor // Write-only binding to the contract
	GovernanceFilterer   // Log filterer for contract events
}

// GovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceSession struct {
	Contract     *Governance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceCallerSession struct {
	Contract *GovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceTransactorSession struct {
	Contract     *GovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceRaw struct {
	Contract *Governance // Generic contract binding to access the raw methods on
}

// GovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceCallerRaw struct {
	Contract *GovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceTransactorRaw struct {
	Contract *GovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernance creates a new instance of Governance, bound to a specific deployed contract.
func NewGovernance(address common.Address, backend bind.ContractBackend) (*Governance, error) {
	contract, err := bindGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// NewGovernanceCaller creates a new read-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceCaller(address common.Address, caller bind.ContractCaller) (*GovernanceCaller, error) {
	contract, err := bindGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceCaller{contract: contract}, nil
}

// NewGovernanceTransactor creates a new write-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTransactor, error) {
	contract, err := bindGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTransactor{contract: contract}, nil
}

// NewGovernanceFilterer creates a new log filterer instance of Governance, bound to a specific deployed contract.
func NewGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceFilterer, error) {
	contract, err := bindGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceFilterer{contract: contract}, nil
}

// bindGovernance binds a generic wrapper to an already deployed contract.
func bindGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.GovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transact(opts, method, params...)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Governance *GovernanceCaller) Bytes32ToString(opts *bind.CallOpts, _bytes32 [32]byte) (string, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "bytes32ToString", _bytes32)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Governance *GovernanceSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _Governance.Contract.Bytes32ToString(&_Governance.CallOpts, _bytes32)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Governance *GovernanceCallerSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _Governance.Contract.Bytes32ToString(&_Governance.CallOpts, _bytes32)
}

// CalculateVotingTally is a free data retrieval call binding the contract method 0x14262d79.
//
// Solidity: function calculateVotingTally(uint256 proposalID) view returns(bool proposalResolved, uint256 winnerID, uint256 votes)
func (_Governance *GovernanceCaller) CalculateVotingTally(opts *bind.CallOpts, proposalID *big.Int) (struct {
	ProposalResolved bool
	WinnerID         *big.Int
	Votes            *big.Int
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "calculateVotingTally", proposalID)

	outstruct := new(struct {
		ProposalResolved bool
		WinnerID         *big.Int
		Votes            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProposalResolved = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.WinnerID = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Votes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateVotingTally is a free data retrieval call binding the contract method 0x14262d79.
//
// Solidity: function calculateVotingTally(uint256 proposalID) view returns(bool proposalResolved, uint256 winnerID, uint256 votes)
func (_Governance *GovernanceSession) CalculateVotingTally(proposalID *big.Int) (struct {
	ProposalResolved bool
	WinnerID         *big.Int
	Votes            *big.Int
}, error) {
	return _Governance.Contract.CalculateVotingTally(&_Governance.CallOpts, proposalID)
}

// CalculateVotingTally is a free data retrieval call binding the contract method 0x14262d79.
//
// Solidity: function calculateVotingTally(uint256 proposalID) view returns(bool proposalResolved, uint256 winnerID, uint256 votes)
func (_Governance *GovernanceCallerSession) CalculateVotingTally(proposalID *big.Int) (struct {
	ProposalResolved bool
	WinnerID         *big.Int
	Votes            *big.Int
}, error) {
	return _Governance.Contract.CalculateVotingTally(&_Governance.CallOpts, proposalID)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 i) view returns(bool active, uint256 assignment, uint256 proposalID)
func (_Governance *GovernanceCaller) GetTask(opts *bind.CallOpts, i *big.Int) (struct {
	Active     bool
	Assignment *big.Int
	ProposalID *big.Int
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getTask", i)

	outstruct := new(struct {
		Active     bool
		Assignment *big.Int
		ProposalID *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Assignment = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProposalID = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 i) view returns(bool active, uint256 assignment, uint256 proposalID)
func (_Governance *GovernanceSession) GetTask(i *big.Int) (struct {
	Active     bool
	Assignment *big.Int
	ProposalID *big.Int
}, error) {
	return _Governance.Contract.GetTask(&_Governance.CallOpts, i)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 i) view returns(bool active, uint256 assignment, uint256 proposalID)
func (_Governance *GovernanceCallerSession) GetTask(i *big.Int) (struct {
	Active     bool
	Assignment *big.Int
	ProposalID *big.Int
}, error) {
	return _Governance.Contract.GetTask(&_Governance.CallOpts, i)
}

// GetVote is a free data retrieval call binding the contract method 0xb9e6842b.
//
// Solidity: function getVote(address from, address delegatedTo, uint256 proposalID) view returns(uint256 weight, uint256[] choices)
func (_Governance *GovernanceCaller) GetVote(opts *bind.CallOpts, from common.Address, delegatedTo common.Address, proposalID *big.Int) (struct {
	Weight  *big.Int
	Choices []*big.Int
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getVote", from, delegatedTo, proposalID)

	outstruct := new(struct {
		Weight  *big.Int
		Choices []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Weight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Choices = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetVote is a free data retrieval call binding the contract method 0xb9e6842b.
//
// Solidity: function getVote(address from, address delegatedTo, uint256 proposalID) view returns(uint256 weight, uint256[] choices)
func (_Governance *GovernanceSession) GetVote(from common.Address, delegatedTo common.Address, proposalID *big.Int) (struct {
	Weight  *big.Int
	Choices []*big.Int
}, error) {
	return _Governance.Contract.GetVote(&_Governance.CallOpts, from, delegatedTo, proposalID)
}

// GetVote is a free data retrieval call binding the contract method 0xb9e6842b.
//
// Solidity: function getVote(address from, address delegatedTo, uint256 proposalID) view returns(uint256 weight, uint256[] choices)
func (_Governance *GovernanceCallerSession) GetVote(from common.Address, delegatedTo common.Address, proposalID *big.Int) (struct {
	Weight  *big.Int
	Choices []*big.Int
}, error) {
	return _Governance.Contract.GetVote(&_Governance.CallOpts, from, delegatedTo, proposalID)
}

// LastProposalID is a free data retrieval call binding the contract method 0xa1d373d7.
//
// Solidity: function lastProposalID() view returns(uint256)
func (_Governance *GovernanceCaller) LastProposalID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "lastProposalID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastProposalID is a free data retrieval call binding the contract method 0xa1d373d7.
//
// Solidity: function lastProposalID() view returns(uint256)
func (_Governance *GovernanceSession) LastProposalID() (*big.Int, error) {
	return _Governance.Contract.LastProposalID(&_Governance.CallOpts)
}

// LastProposalID is a free data retrieval call binding the contract method 0xa1d373d7.
//
// Solidity: function lastProposalID() view returns(uint256)
func (_Governance *GovernanceCallerSession) LastProposalID() (*big.Int, error) {
	return _Governance.Contract.LastProposalID(&_Governance.CallOpts)
}

// MaxExecutionPeriod is a free data retrieval call binding the contract method 0xc25c2f26.
//
// Solidity: function maxExecutionPeriod() pure returns(uint256)
func (_Governance *GovernanceCaller) MaxExecutionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "maxExecutionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxExecutionPeriod is a free data retrieval call binding the contract method 0xc25c2f26.
//
// Solidity: function maxExecutionPeriod() pure returns(uint256)
func (_Governance *GovernanceSession) MaxExecutionPeriod() (*big.Int, error) {
	return _Governance.Contract.MaxExecutionPeriod(&_Governance.CallOpts)
}

// MaxExecutionPeriod is a free data retrieval call binding the contract method 0xc25c2f26.
//
// Solidity: function maxExecutionPeriod() pure returns(uint256)
func (_Governance *GovernanceCallerSession) MaxExecutionPeriod() (*big.Int, error) {
	return _Governance.Contract.MaxExecutionPeriod(&_Governance.CallOpts)
}

// MaxOptions is a free data retrieval call binding the contract method 0x5df17077.
//
// Solidity: function maxOptions() pure returns(uint256)
func (_Governance *GovernanceCaller) MaxOptions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "maxOptions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxOptions is a free data retrieval call binding the contract method 0x5df17077.
//
// Solidity: function maxOptions() pure returns(uint256)
func (_Governance *GovernanceSession) MaxOptions() (*big.Int, error) {
	return _Governance.Contract.MaxOptions(&_Governance.CallOpts)
}

// MaxOptions is a free data retrieval call binding the contract method 0x5df17077.
//
// Solidity: function maxOptions() pure returns(uint256)
func (_Governance *GovernanceCallerSession) MaxOptions() (*big.Int, error) {
	return _Governance.Contract.MaxOptions(&_Governance.CallOpts)
}

// MinVotesAbsolute is a free data retrieval call binding the contract method 0x1191e270.
//
// Solidity: function minVotesAbsolute(uint256 totalWeight, uint256 minVotesRatio) pure returns(uint256)
func (_Governance *GovernanceCaller) MinVotesAbsolute(opts *bind.CallOpts, totalWeight *big.Int, minVotesRatio *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "minVotesAbsolute", totalWeight, minVotesRatio)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVotesAbsolute is a free data retrieval call binding the contract method 0x1191e270.
//
// Solidity: function minVotesAbsolute(uint256 totalWeight, uint256 minVotesRatio) pure returns(uint256)
func (_Governance *GovernanceSession) MinVotesAbsolute(totalWeight *big.Int, minVotesRatio *big.Int) (*big.Int, error) {
	return _Governance.Contract.MinVotesAbsolute(&_Governance.CallOpts, totalWeight, minVotesRatio)
}

// MinVotesAbsolute is a free data retrieval call binding the contract method 0x1191e270.
//
// Solidity: function minVotesAbsolute(uint256 totalWeight, uint256 minVotesRatio) pure returns(uint256)
func (_Governance *GovernanceCallerSession) MinVotesAbsolute(totalWeight *big.Int, minVotesRatio *big.Int) (*big.Int, error) {
	return _Governance.Contract.MinVotesAbsolute(&_Governance.CallOpts, totalWeight, minVotesRatio)
}

// OverriddenWeight is a free data retrieval call binding the contract method 0x2177d6fc.
//
// Solidity: function overriddenWeight(address , uint256 ) view returns(uint256)
func (_Governance *GovernanceCaller) OverriddenWeight(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "overriddenWeight", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OverriddenWeight is a free data retrieval call binding the contract method 0x2177d6fc.
//
// Solidity: function overriddenWeight(address , uint256 ) view returns(uint256)
func (_Governance *GovernanceSession) OverriddenWeight(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Governance.Contract.OverriddenWeight(&_Governance.CallOpts, arg0, arg1)
}

// OverriddenWeight is a free data retrieval call binding the contract method 0x2177d6fc.
//
// Solidity: function overriddenWeight(address , uint256 ) view returns(uint256)
func (_Governance *GovernanceCallerSession) OverriddenWeight(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Governance.Contract.OverriddenWeight(&_Governance.CallOpts, arg0, arg1)
}

// ProposalBurntFee is a free data retrieval call binding the contract method 0x8a444bd7.
//
// Solidity: function proposalBurntFee() pure returns(uint256)
func (_Governance *GovernanceCaller) ProposalBurntFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "proposalBurntFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalBurntFee is a free data retrieval call binding the contract method 0x8a444bd7.
//
// Solidity: function proposalBurntFee() pure returns(uint256)
func (_Governance *GovernanceSession) ProposalBurntFee() (*big.Int, error) {
	return _Governance.Contract.ProposalBurntFee(&_Governance.CallOpts)
}

// ProposalBurntFee is a free data retrieval call binding the contract method 0x8a444bd7.
//
// Solidity: function proposalBurntFee() pure returns(uint256)
func (_Governance *GovernanceCallerSession) ProposalBurntFee() (*big.Int, error) {
	return _Governance.Contract.ProposalBurntFee(&_Governance.CallOpts)
}

// ProposalFee is a free data retrieval call binding the contract method 0xc27cabb5.
//
// Solidity: function proposalFee() pure returns(uint256)
func (_Governance *GovernanceCaller) ProposalFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "proposalFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalFee is a free data retrieval call binding the contract method 0xc27cabb5.
//
// Solidity: function proposalFee() pure returns(uint256)
func (_Governance *GovernanceSession) ProposalFee() (*big.Int, error) {
	return _Governance.Contract.ProposalFee(&_Governance.CallOpts)
}

// ProposalFee is a free data retrieval call binding the contract method 0xc27cabb5.
//
// Solidity: function proposalFee() pure returns(uint256)
func (_Governance *GovernanceCallerSession) ProposalFee() (*big.Int, error) {
	return _Governance.Contract.ProposalFee(&_Governance.CallOpts)
}

// ProposalOptionState is a free data retrieval call binding the contract method 0x5f89801e.
//
// Solidity: function proposalOptionState(uint256 proposalID, uint256 optionID) view returns(uint256 votes, uint256 agreementRatio, uint256 agreement)
func (_Governance *GovernanceCaller) ProposalOptionState(opts *bind.CallOpts, proposalID *big.Int, optionID *big.Int) (struct {
	Votes          *big.Int
	AgreementRatio *big.Int
	Agreement      *big.Int
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "proposalOptionState", proposalID, optionID)

	outstruct := new(struct {
		Votes          *big.Int
		AgreementRatio *big.Int
		Agreement      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Votes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AgreementRatio = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Agreement = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalOptionState is a free data retrieval call binding the contract method 0x5f89801e.
//
// Solidity: function proposalOptionState(uint256 proposalID, uint256 optionID) view returns(uint256 votes, uint256 agreementRatio, uint256 agreement)
func (_Governance *GovernanceSession) ProposalOptionState(proposalID *big.Int, optionID *big.Int) (struct {
	Votes          *big.Int
	AgreementRatio *big.Int
	Agreement      *big.Int
}, error) {
	return _Governance.Contract.ProposalOptionState(&_Governance.CallOpts, proposalID, optionID)
}

// ProposalOptionState is a free data retrieval call binding the contract method 0x5f89801e.
//
// Solidity: function proposalOptionState(uint256 proposalID, uint256 optionID) view returns(uint256 votes, uint256 agreementRatio, uint256 agreement)
func (_Governance *GovernanceCallerSession) ProposalOptionState(proposalID *big.Int, optionID *big.Int) (struct {
	Votes          *big.Int
	AgreementRatio *big.Int
	Agreement      *big.Int
}, error) {
	return _Governance.Contract.ProposalOptionState(&_Governance.CallOpts, proposalID, optionID)
}

// ProposalParams is a free data retrieval call binding the contract method 0xcfa1afa3.
//
// Solidity: function proposalParams(uint256 proposalID) view returns(uint256 pType, uint8 executable, uint256 minVotes, uint256 minAgreement, uint256[] opinionScales, bytes32[] options, address proposalContract, uint256 votingStartTime, uint256 votingMinEndTime, uint256 votingMaxEndTime)
func (_Governance *GovernanceCaller) ProposalParams(opts *bind.CallOpts, proposalID *big.Int) (struct {
	PType            *big.Int
	Executable       uint8
	MinVotes         *big.Int
	MinAgreement     *big.Int
	OpinionScales    []*big.Int
	Options          [][32]byte
	ProposalContract common.Address
	VotingStartTime  *big.Int
	VotingMinEndTime *big.Int
	VotingMaxEndTime *big.Int
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "proposalParams", proposalID)

	outstruct := new(struct {
		PType            *big.Int
		Executable       uint8
		MinVotes         *big.Int
		MinAgreement     *big.Int
		OpinionScales    []*big.Int
		Options          [][32]byte
		ProposalContract common.Address
		VotingStartTime  *big.Int
		VotingMinEndTime *big.Int
		VotingMaxEndTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PType = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Executable = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.MinVotes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MinAgreement = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OpinionScales = *abi.ConvertType(out[4], new([]*big.Int)).(*[]*big.Int)
	outstruct.Options = *abi.ConvertType(out[5], new([][32]byte)).(*[][32]byte)
	outstruct.ProposalContract = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.VotingStartTime = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.VotingMinEndTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.VotingMaxEndTime = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalParams is a free data retrieval call binding the contract method 0xcfa1afa3.
//
// Solidity: function proposalParams(uint256 proposalID) view returns(uint256 pType, uint8 executable, uint256 minVotes, uint256 minAgreement, uint256[] opinionScales, bytes32[] options, address proposalContract, uint256 votingStartTime, uint256 votingMinEndTime, uint256 votingMaxEndTime)
func (_Governance *GovernanceSession) ProposalParams(proposalID *big.Int) (struct {
	PType            *big.Int
	Executable       uint8
	MinVotes         *big.Int
	MinAgreement     *big.Int
	OpinionScales    []*big.Int
	Options          [][32]byte
	ProposalContract common.Address
	VotingStartTime  *big.Int
	VotingMinEndTime *big.Int
	VotingMaxEndTime *big.Int
}, error) {
	return _Governance.Contract.ProposalParams(&_Governance.CallOpts, proposalID)
}

// ProposalParams is a free data retrieval call binding the contract method 0xcfa1afa3.
//
// Solidity: function proposalParams(uint256 proposalID) view returns(uint256 pType, uint8 executable, uint256 minVotes, uint256 minAgreement, uint256[] opinionScales, bytes32[] options, address proposalContract, uint256 votingStartTime, uint256 votingMinEndTime, uint256 votingMaxEndTime)
func (_Governance *GovernanceCallerSession) ProposalParams(proposalID *big.Int) (struct {
	PType            *big.Int
	Executable       uint8
	MinVotes         *big.Int
	MinAgreement     *big.Int
	OpinionScales    []*big.Int
	Options          [][32]byte
	ProposalContract common.Address
	VotingStartTime  *big.Int
	VotingMinEndTime *big.Int
	VotingMaxEndTime *big.Int
}, error) {
	return _Governance.Contract.ProposalParams(&_Governance.CallOpts, proposalID)
}

// ProposalState is a free data retrieval call binding the contract method 0xd26331d4.
//
// Solidity: function proposalState(uint256 proposalID) view returns(uint256 winnerOptionID, uint256 votes, uint256 status)
func (_Governance *GovernanceCaller) ProposalState(opts *bind.CallOpts, proposalID *big.Int) (struct {
	WinnerOptionID *big.Int
	Votes          *big.Int
	Status         *big.Int
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "proposalState", proposalID)

	outstruct := new(struct {
		WinnerOptionID *big.Int
		Votes          *big.Int
		Status         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WinnerOptionID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Votes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalState is a free data retrieval call binding the contract method 0xd26331d4.
//
// Solidity: function proposalState(uint256 proposalID) view returns(uint256 winnerOptionID, uint256 votes, uint256 status)
func (_Governance *GovernanceSession) ProposalState(proposalID *big.Int) (struct {
	WinnerOptionID *big.Int
	Votes          *big.Int
	Status         *big.Int
}, error) {
	return _Governance.Contract.ProposalState(&_Governance.CallOpts, proposalID)
}

// ProposalState is a free data retrieval call binding the contract method 0xd26331d4.
//
// Solidity: function proposalState(uint256 proposalID) view returns(uint256 winnerOptionID, uint256 votes, uint256 status)
func (_Governance *GovernanceCallerSession) ProposalState(proposalID *big.Int) (struct {
	WinnerOptionID *big.Int
	Votes          *big.Int
	Status         *big.Int
}, error) {
	return _Governance.Contract.ProposalState(&_Governance.CallOpts, proposalID)
}

// TaskErasingReward is a free data retrieval call binding the contract method 0xaea2ddbd.
//
// Solidity: function taskErasingReward() pure returns(uint256)
func (_Governance *GovernanceCaller) TaskErasingReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "taskErasingReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskErasingReward is a free data retrieval call binding the contract method 0xaea2ddbd.
//
// Solidity: function taskErasingReward() pure returns(uint256)
func (_Governance *GovernanceSession) TaskErasingReward() (*big.Int, error) {
	return _Governance.Contract.TaskErasingReward(&_Governance.CallOpts)
}

// TaskErasingReward is a free data retrieval call binding the contract method 0xaea2ddbd.
//
// Solidity: function taskErasingReward() pure returns(uint256)
func (_Governance *GovernanceCallerSession) TaskErasingReward() (*big.Int, error) {
	return _Governance.Contract.TaskErasingReward(&_Governance.CallOpts)
}

// TaskHandlingReward is a free data retrieval call binding the contract method 0xe3a4d709.
//
// Solidity: function taskHandlingReward() pure returns(uint256)
func (_Governance *GovernanceCaller) TaskHandlingReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "taskHandlingReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskHandlingReward is a free data retrieval call binding the contract method 0xe3a4d709.
//
// Solidity: function taskHandlingReward() pure returns(uint256)
func (_Governance *GovernanceSession) TaskHandlingReward() (*big.Int, error) {
	return _Governance.Contract.TaskHandlingReward(&_Governance.CallOpts)
}

// TaskHandlingReward is a free data retrieval call binding the contract method 0xe3a4d709.
//
// Solidity: function taskHandlingReward() pure returns(uint256)
func (_Governance *GovernanceCallerSession) TaskHandlingReward() (*big.Int, error) {
	return _Governance.Contract.TaskHandlingReward(&_Governance.CallOpts)
}

// TasksCount is a free data retrieval call binding the contract method 0xbb6a0f07.
//
// Solidity: function tasksCount() view returns(uint256)
func (_Governance *GovernanceCaller) TasksCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "tasksCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TasksCount is a free data retrieval call binding the contract method 0xbb6a0f07.
//
// Solidity: function tasksCount() view returns(uint256)
func (_Governance *GovernanceSession) TasksCount() (*big.Int, error) {
	return _Governance.Contract.TasksCount(&_Governance.CallOpts)
}

// TasksCount is a free data retrieval call binding the contract method 0xbb6a0f07.
//
// Solidity: function tasksCount() view returns(uint256)
func (_Governance *GovernanceCallerSession) TasksCount() (*big.Int, error) {
	return _Governance.Contract.TasksCount(&_Governance.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes4)
func (_Governance *GovernanceCaller) Version(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "version")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes4)
func (_Governance *GovernanceSession) Version() ([4]byte, error) {
	return _Governance.Contract.Version(&_Governance.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes4)
func (_Governance *GovernanceCallerSession) Version() ([4]byte, error) {
	return _Governance.Contract.Version(&_Governance.CallOpts)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 proposalID) returns()
func (_Governance *GovernanceTransactor) CancelProposal(opts *bind.TransactOpts, proposalID *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "cancelProposal", proposalID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 proposalID) returns()
func (_Governance *GovernanceSession) CancelProposal(proposalID *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.CancelProposal(&_Governance.TransactOpts, proposalID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 proposalID) returns()
func (_Governance *GovernanceTransactorSession) CancelProposal(proposalID *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.CancelProposal(&_Governance.TransactOpts, proposalID)
}

// CancelVote is a paid mutator transaction binding the contract method 0x21edf2eb.
//
// Solidity: function cancelVote(address delegatedTo, uint256 proposalID) returns()
func (_Governance *GovernanceTransactor) CancelVote(opts *bind.TransactOpts, delegatedTo common.Address, proposalID *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "cancelVote", delegatedTo, proposalID)
}

// CancelVote is a paid mutator transaction binding the contract method 0x21edf2eb.
//
// Solidity: function cancelVote(address delegatedTo, uint256 proposalID) returns()
func (_Governance *GovernanceSession) CancelVote(delegatedTo common.Address, proposalID *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.CancelVote(&_Governance.TransactOpts, delegatedTo, proposalID)
}

// CancelVote is a paid mutator transaction binding the contract method 0x21edf2eb.
//
// Solidity: function cancelVote(address delegatedTo, uint256 proposalID) returns()
func (_Governance *GovernanceTransactorSession) CancelVote(delegatedTo common.Address, proposalID *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.CancelVote(&_Governance.TransactOpts, delegatedTo, proposalID)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x07eecaf5.
//
// Solidity: function createProposal(address proposalContract) payable returns()
func (_Governance *GovernanceTransactor) CreateProposal(opts *bind.TransactOpts, proposalContract common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "createProposal", proposalContract)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x07eecaf5.
//
// Solidity: function createProposal(address proposalContract) payable returns()
func (_Governance *GovernanceSession) CreateProposal(proposalContract common.Address) (*types.Transaction, error) {
	return _Governance.Contract.CreateProposal(&_Governance.TransactOpts, proposalContract)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x07eecaf5.
//
// Solidity: function createProposal(address proposalContract) payable returns()
func (_Governance *GovernanceTransactorSession) CreateProposal(proposalContract common.Address) (*types.Transaction, error) {
	return _Governance.Contract.CreateProposal(&_Governance.TransactOpts, proposalContract)
}

// HandleTasks is a paid mutator transaction binding the contract method 0x6b2ad7d8.
//
// Solidity: function handleTasks(uint256 startIdx, uint256 quantity) returns()
func (_Governance *GovernanceTransactor) HandleTasks(opts *bind.TransactOpts, startIdx *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "handleTasks", startIdx, quantity)
}

// HandleTasks is a paid mutator transaction binding the contract method 0x6b2ad7d8.
//
// Solidity: function handleTasks(uint256 startIdx, uint256 quantity) returns()
func (_Governance *GovernanceSession) HandleTasks(startIdx *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.HandleTasks(&_Governance.TransactOpts, startIdx, quantity)
}

// HandleTasks is a paid mutator transaction binding the contract method 0x6b2ad7d8.
//
// Solidity: function handleTasks(uint256 startIdx, uint256 quantity) returns()
func (_Governance *GovernanceTransactorSession) HandleTasks(startIdx *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.HandleTasks(&_Governance.TransactOpts, startIdx, quantity)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _governableContract, address _proposalVerifier) returns()
func (_Governance *GovernanceTransactor) Initialize(opts *bind.TransactOpts, _governableContract common.Address, _proposalVerifier common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "initialize", _governableContract, _proposalVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _governableContract, address _proposalVerifier) returns()
func (_Governance *GovernanceSession) Initialize(_governableContract common.Address, _proposalVerifier common.Address) (*types.Transaction, error) {
	return _Governance.Contract.Initialize(&_Governance.TransactOpts, _governableContract, _proposalVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _governableContract, address _proposalVerifier) returns()
func (_Governance *GovernanceTransactorSession) Initialize(_governableContract common.Address, _proposalVerifier common.Address) (*types.Transaction, error) {
	return _Governance.Contract.Initialize(&_Governance.TransactOpts, _governableContract, _proposalVerifier)
}

// RecountVote is a paid mutator transaction binding the contract method 0x69921506.
//
// Solidity: function recountVote(address voterAddr, address delegatedTo, uint256 proposalID) returns()
func (_Governance *GovernanceTransactor) RecountVote(opts *bind.TransactOpts, voterAddr common.Address, delegatedTo common.Address, proposalID *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "recountVote", voterAddr, delegatedTo, proposalID)
}

// RecountVote is a paid mutator transaction binding the contract method 0x69921506.
//
// Solidity: function recountVote(address voterAddr, address delegatedTo, uint256 proposalID) returns()
func (_Governance *GovernanceSession) RecountVote(voterAddr common.Address, delegatedTo common.Address, proposalID *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.RecountVote(&_Governance.TransactOpts, voterAddr, delegatedTo, proposalID)
}

// RecountVote is a paid mutator transaction binding the contract method 0x69921506.
//
// Solidity: function recountVote(address voterAddr, address delegatedTo, uint256 proposalID) returns()
func (_Governance *GovernanceTransactorSession) RecountVote(voterAddr common.Address, delegatedTo common.Address, proposalID *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.RecountVote(&_Governance.TransactOpts, voterAddr, delegatedTo, proposalID)
}

// TasksCleanup is a paid mutator transaction binding the contract method 0x491adeee.
//
// Solidity: function tasksCleanup(uint256 quantity) returns()
func (_Governance *GovernanceTransactor) TasksCleanup(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "tasksCleanup", quantity)
}

// TasksCleanup is a paid mutator transaction binding the contract method 0x491adeee.
//
// Solidity: function tasksCleanup(uint256 quantity) returns()
func (_Governance *GovernanceSession) TasksCleanup(quantity *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.TasksCleanup(&_Governance.TransactOpts, quantity)
}

// TasksCleanup is a paid mutator transaction binding the contract method 0x491adeee.
//
// Solidity: function tasksCleanup(uint256 quantity) returns()
func (_Governance *GovernanceTransactorSession) TasksCleanup(quantity *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.TasksCleanup(&_Governance.TransactOpts, quantity)
}

// Vote is a paid mutator transaction binding the contract method 0x172c18b1.
//
// Solidity: function vote(address delegatedTo, uint256 proposalID, uint256[] choices) returns()
func (_Governance *GovernanceTransactor) Vote(opts *bind.TransactOpts, delegatedTo common.Address, proposalID *big.Int, choices []*big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "vote", delegatedTo, proposalID, choices)
}

// Vote is a paid mutator transaction binding the contract method 0x172c18b1.
//
// Solidity: function vote(address delegatedTo, uint256 proposalID, uint256[] choices) returns()
func (_Governance *GovernanceSession) Vote(delegatedTo common.Address, proposalID *big.Int, choices []*big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Vote(&_Governance.TransactOpts, delegatedTo, proposalID, choices)
}

// Vote is a paid mutator transaction binding the contract method 0x172c18b1.
//
// Solidity: function vote(address delegatedTo, uint256 proposalID, uint256[] choices) returns()
func (_Governance *GovernanceTransactorSession) Vote(delegatedTo common.Address, proposalID *big.Int, choices []*big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Vote(&_Governance.TransactOpts, delegatedTo, proposalID, choices)
}

// GovernanceProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the Governance contract.
type GovernanceProposalCanceledIterator struct {
	Event *GovernanceProposalCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalCanceled represents a ProposalCanceled event raised by the Governance contract.
type GovernanceProposalCanceled struct {
	ProposalID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalID)
func (_Governance *GovernanceFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*GovernanceProposalCanceledIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalCanceledIterator{contract: _Governance.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalID)
func (_Governance *GovernanceFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *GovernanceProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalCanceled)
				if err := _Governance.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalID)
func (_Governance *GovernanceFilterer) ParseProposalCanceled(log types.Log) (*GovernanceProposalCanceled, error) {
	event := new(GovernanceProposalCanceled)
	if err := _Governance.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Governance contract.
type GovernanceProposalCreatedIterator struct {
	Event *GovernanceProposalCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalCreated represents a ProposalCreated event raised by the Governance contract.
type GovernanceProposalCreated struct {
	ProposalID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 proposalID)
func (_Governance *GovernanceFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*GovernanceProposalCreatedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalCreatedIterator{contract: _Governance.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 proposalID)
func (_Governance *GovernanceFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *GovernanceProposalCreated) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalCreated)
				if err := _Governance.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCreated is a log parse operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 proposalID)
func (_Governance *GovernanceFilterer) ParseProposalCreated(log types.Log) (*GovernanceProposalCreated, error) {
	event := new(GovernanceProposalCreated)
	if err := _Governance.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalExecutionExpiredIterator is returned from FilterProposalExecutionExpired and is used to iterate over the raw logs and unpacked data for ProposalExecutionExpired events raised by the Governance contract.
type GovernanceProposalExecutionExpiredIterator struct {
	Event *GovernanceProposalExecutionExpired // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalExecutionExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalExecutionExpired)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalExecutionExpired)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalExecutionExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalExecutionExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalExecutionExpired represents a ProposalExecutionExpired event raised by the Governance contract.
type GovernanceProposalExecutionExpired struct {
	ProposalID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecutionExpired is a free log retrieval operation binding the contract event 0xe8365dd25802fb5113a4ebd6fbe5fee885b5ea470b6b1467f3f4df69e490ed87.
//
// Solidity: event ProposalExecutionExpired(uint256 proposalID)
func (_Governance *GovernanceFilterer) FilterProposalExecutionExpired(opts *bind.FilterOpts) (*GovernanceProposalExecutionExpiredIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalExecutionExpired")
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalExecutionExpiredIterator{contract: _Governance.contract, event: "ProposalExecutionExpired", logs: logs, sub: sub}, nil
}

// WatchProposalExecutionExpired is a free log subscription operation binding the contract event 0xe8365dd25802fb5113a4ebd6fbe5fee885b5ea470b6b1467f3f4df69e490ed87.
//
// Solidity: event ProposalExecutionExpired(uint256 proposalID)
func (_Governance *GovernanceFilterer) WatchProposalExecutionExpired(opts *bind.WatchOpts, sink chan<- *GovernanceProposalExecutionExpired) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalExecutionExpired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalExecutionExpired)
				if err := _Governance.contract.UnpackLog(event, "ProposalExecutionExpired", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalExecutionExpired is a log parse operation binding the contract event 0xe8365dd25802fb5113a4ebd6fbe5fee885b5ea470b6b1467f3f4df69e490ed87.
//
// Solidity: event ProposalExecutionExpired(uint256 proposalID)
func (_Governance *GovernanceFilterer) ParseProposalExecutionExpired(log types.Log) (*GovernanceProposalExecutionExpired, error) {
	event := new(GovernanceProposalExecutionExpired)
	if err := _Governance.contract.UnpackLog(event, "ProposalExecutionExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalRejectedIterator is returned from FilterProposalRejected and is used to iterate over the raw logs and unpacked data for ProposalRejected events raised by the Governance contract.
type GovernanceProposalRejectedIterator struct {
	Event *GovernanceProposalRejected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalRejected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalRejected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalRejected represents a ProposalRejected event raised by the Governance contract.
type GovernanceProposalRejected struct {
	ProposalID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalRejected is a free log retrieval operation binding the contract event 0xd92fba445edb3153b571e6df782d7a66fd0ce668519273670820ee3a86da0ef4.
//
// Solidity: event ProposalRejected(uint256 proposalID)
func (_Governance *GovernanceFilterer) FilterProposalRejected(opts *bind.FilterOpts) (*GovernanceProposalRejectedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalRejected")
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalRejectedIterator{contract: _Governance.contract, event: "ProposalRejected", logs: logs, sub: sub}, nil
}

// WatchProposalRejected is a free log subscription operation binding the contract event 0xd92fba445edb3153b571e6df782d7a66fd0ce668519273670820ee3a86da0ef4.
//
// Solidity: event ProposalRejected(uint256 proposalID)
func (_Governance *GovernanceFilterer) WatchProposalRejected(opts *bind.WatchOpts, sink chan<- *GovernanceProposalRejected) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalRejected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalRejected)
				if err := _Governance.contract.UnpackLog(event, "ProposalRejected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalRejected is a log parse operation binding the contract event 0xd92fba445edb3153b571e6df782d7a66fd0ce668519273670820ee3a86da0ef4.
//
// Solidity: event ProposalRejected(uint256 proposalID)
func (_Governance *GovernanceFilterer) ParseProposalRejected(log types.Log) (*GovernanceProposalRejected, error) {
	event := new(GovernanceProposalRejected)
	if err := _Governance.contract.UnpackLog(event, "ProposalRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalResolvedIterator is returned from FilterProposalResolved and is used to iterate over the raw logs and unpacked data for ProposalResolved events raised by the Governance contract.
type GovernanceProposalResolvedIterator struct {
	Event *GovernanceProposalResolved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalResolved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalResolved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalResolved represents a ProposalResolved event raised by the Governance contract.
type GovernanceProposalResolved struct {
	ProposalID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalResolved is a free log retrieval operation binding the contract event 0x663674d96fd5c2a954bf75ad2e6795f9c9701eb687a7a8f3297c7a299467c941.
//
// Solidity: event ProposalResolved(uint256 proposalID)
func (_Governance *GovernanceFilterer) FilterProposalResolved(opts *bind.FilterOpts) (*GovernanceProposalResolvedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalResolved")
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalResolvedIterator{contract: _Governance.contract, event: "ProposalResolved", logs: logs, sub: sub}, nil
}

// WatchProposalResolved is a free log subscription operation binding the contract event 0x663674d96fd5c2a954bf75ad2e6795f9c9701eb687a7a8f3297c7a299467c941.
//
// Solidity: event ProposalResolved(uint256 proposalID)
func (_Governance *GovernanceFilterer) WatchProposalResolved(opts *bind.WatchOpts, sink chan<- *GovernanceProposalResolved) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalResolved)
				if err := _Governance.contract.UnpackLog(event, "ProposalResolved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalResolved is a log parse operation binding the contract event 0x663674d96fd5c2a954bf75ad2e6795f9c9701eb687a7a8f3297c7a299467c941.
//
// Solidity: event ProposalResolved(uint256 proposalID)
func (_Governance *GovernanceFilterer) ParseProposalResolved(log types.Log) (*GovernanceProposalResolved, error) {
	event := new(GovernanceProposalResolved)
	if err := _Governance.contract.UnpackLog(event, "ProposalResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceTasksErasedIterator is returned from FilterTasksErased and is used to iterate over the raw logs and unpacked data for TasksErased events raised by the Governance contract.
type GovernanceTasksErasedIterator struct {
	Event *GovernanceTasksErased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceTasksErasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTasksErased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceTasksErased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceTasksErasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTasksErasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTasksErased represents a TasksErased event raised by the Governance contract.
type GovernanceTasksErased struct {
	Quantity *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTasksErased is a free log retrieval operation binding the contract event 0xdd64b087b4726a6d5006bfae34869d1044e415cd6fc53054b69619ee337b730d.
//
// Solidity: event TasksErased(uint256 quantity)
func (_Governance *GovernanceFilterer) FilterTasksErased(opts *bind.FilterOpts) (*GovernanceTasksErasedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "TasksErased")
	if err != nil {
		return nil, err
	}
	return &GovernanceTasksErasedIterator{contract: _Governance.contract, event: "TasksErased", logs: logs, sub: sub}, nil
}

// WatchTasksErased is a free log subscription operation binding the contract event 0xdd64b087b4726a6d5006bfae34869d1044e415cd6fc53054b69619ee337b730d.
//
// Solidity: event TasksErased(uint256 quantity)
func (_Governance *GovernanceFilterer) WatchTasksErased(opts *bind.WatchOpts, sink chan<- *GovernanceTasksErased) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "TasksErased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTasksErased)
				if err := _Governance.contract.UnpackLog(event, "TasksErased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTasksErased is a log parse operation binding the contract event 0xdd64b087b4726a6d5006bfae34869d1044e415cd6fc53054b69619ee337b730d.
//
// Solidity: event TasksErased(uint256 quantity)
func (_Governance *GovernanceFilterer) ParseTasksErased(log types.Log) (*GovernanceTasksErased, error) {
	event := new(GovernanceTasksErased)
	if err := _Governance.contract.UnpackLog(event, "TasksErased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceTasksHandledIterator is returned from FilterTasksHandled and is used to iterate over the raw logs and unpacked data for TasksHandled events raised by the Governance contract.
type GovernanceTasksHandledIterator struct {
	Event *GovernanceTasksHandled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceTasksHandledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTasksHandled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceTasksHandled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceTasksHandledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTasksHandledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTasksHandled represents a TasksHandled event raised by the Governance contract.
type GovernanceTasksHandled struct {
	StartIdx *big.Int
	EndIdx   *big.Int
	Handled  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTasksHandled is a free log retrieval operation binding the contract event 0x010b51ac92c2d2e8fe9c84de31065ca878ff8274a1e5e6d31799520b7d16bb58.
//
// Solidity: event TasksHandled(uint256 startIdx, uint256 endIdx, uint256 handled)
func (_Governance *GovernanceFilterer) FilterTasksHandled(opts *bind.FilterOpts) (*GovernanceTasksHandledIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "TasksHandled")
	if err != nil {
		return nil, err
	}
	return &GovernanceTasksHandledIterator{contract: _Governance.contract, event: "TasksHandled", logs: logs, sub: sub}, nil
}

// WatchTasksHandled is a free log subscription operation binding the contract event 0x010b51ac92c2d2e8fe9c84de31065ca878ff8274a1e5e6d31799520b7d16bb58.
//
// Solidity: event TasksHandled(uint256 startIdx, uint256 endIdx, uint256 handled)
func (_Governance *GovernanceFilterer) WatchTasksHandled(opts *bind.WatchOpts, sink chan<- *GovernanceTasksHandled) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "TasksHandled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTasksHandled)
				if err := _Governance.contract.UnpackLog(event, "TasksHandled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTasksHandled is a log parse operation binding the contract event 0x010b51ac92c2d2e8fe9c84de31065ca878ff8274a1e5e6d31799520b7d16bb58.
//
// Solidity: event TasksHandled(uint256 startIdx, uint256 endIdx, uint256 handled)
func (_Governance *GovernanceFilterer) ParseTasksHandled(log types.Log) (*GovernanceTasksHandled, error) {
	event := new(GovernanceTasksHandled)
	if err := _Governance.contract.UnpackLog(event, "TasksHandled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceVoteCanceledIterator is returned from FilterVoteCanceled and is used to iterate over the raw logs and unpacked data for VoteCanceled events raised by the Governance contract.
type GovernanceVoteCanceledIterator struct {
	Event *GovernanceVoteCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceVoteCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceVoteCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceVoteCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceVoteCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceVoteCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceVoteCanceled represents a VoteCanceled event raised by the Governance contract.
type GovernanceVoteCanceled struct {
	Voter       common.Address
	DelegatedTo common.Address
	ProposalID  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVoteCanceled is a free log retrieval operation binding the contract event 0x666685d133047310e2a2e8c4f6794b6dccb4e9ad9c6903ac753fb10d8918b649.
//
// Solidity: event VoteCanceled(address voter, address delegatedTo, uint256 proposalID)
func (_Governance *GovernanceFilterer) FilterVoteCanceled(opts *bind.FilterOpts) (*GovernanceVoteCanceledIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "VoteCanceled")
	if err != nil {
		return nil, err
	}
	return &GovernanceVoteCanceledIterator{contract: _Governance.contract, event: "VoteCanceled", logs: logs, sub: sub}, nil
}

// WatchVoteCanceled is a free log subscription operation binding the contract event 0x666685d133047310e2a2e8c4f6794b6dccb4e9ad9c6903ac753fb10d8918b649.
//
// Solidity: event VoteCanceled(address voter, address delegatedTo, uint256 proposalID)
func (_Governance *GovernanceFilterer) WatchVoteCanceled(opts *bind.WatchOpts, sink chan<- *GovernanceVoteCanceled) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "VoteCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceVoteCanceled)
				if err := _Governance.contract.UnpackLog(event, "VoteCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCanceled is a log parse operation binding the contract event 0x666685d133047310e2a2e8c4f6794b6dccb4e9ad9c6903ac753fb10d8918b649.
//
// Solidity: event VoteCanceled(address voter, address delegatedTo, uint256 proposalID)
func (_Governance *GovernanceFilterer) ParseVoteCanceled(log types.Log) (*GovernanceVoteCanceled, error) {
	event := new(GovernanceVoteCanceled)
	if err := _Governance.contract.UnpackLog(event, "VoteCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceVoteWeightOverriddenIterator is returned from FilterVoteWeightOverridden and is used to iterate over the raw logs and unpacked data for VoteWeightOverridden events raised by the Governance contract.
type GovernanceVoteWeightOverriddenIterator struct {
	Event *GovernanceVoteWeightOverridden // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceVoteWeightOverriddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceVoteWeightOverridden)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceVoteWeightOverridden)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceVoteWeightOverriddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceVoteWeightOverriddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceVoteWeightOverridden represents a VoteWeightOverridden event raised by the Governance contract.
type GovernanceVoteWeightOverridden struct {
	Voter common.Address
	Diff  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVoteWeightOverridden is a free log retrieval operation binding the contract event 0x68fe85c5f71a2900fddf574935a27d4f1cb28af34d4fa2742b202684b45d3d14.
//
// Solidity: event VoteWeightOverridden(address voter, uint256 diff)
func (_Governance *GovernanceFilterer) FilterVoteWeightOverridden(opts *bind.FilterOpts) (*GovernanceVoteWeightOverriddenIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "VoteWeightOverridden")
	if err != nil {
		return nil, err
	}
	return &GovernanceVoteWeightOverriddenIterator{contract: _Governance.contract, event: "VoteWeightOverridden", logs: logs, sub: sub}, nil
}

// WatchVoteWeightOverridden is a free log subscription operation binding the contract event 0x68fe85c5f71a2900fddf574935a27d4f1cb28af34d4fa2742b202684b45d3d14.
//
// Solidity: event VoteWeightOverridden(address voter, uint256 diff)
func (_Governance *GovernanceFilterer) WatchVoteWeightOverridden(opts *bind.WatchOpts, sink chan<- *GovernanceVoteWeightOverridden) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "VoteWeightOverridden")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceVoteWeightOverridden)
				if err := _Governance.contract.UnpackLog(event, "VoteWeightOverridden", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteWeightOverridden is a log parse operation binding the contract event 0x68fe85c5f71a2900fddf574935a27d4f1cb28af34d4fa2742b202684b45d3d14.
//
// Solidity: event VoteWeightOverridden(address voter, uint256 diff)
func (_Governance *GovernanceFilterer) ParseVoteWeightOverridden(log types.Log) (*GovernanceVoteWeightOverridden, error) {
	event := new(GovernanceVoteWeightOverridden)
	if err := _Governance.contract.UnpackLog(event, "VoteWeightOverridden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceVoteWeightUnOverriddenIterator is returned from FilterVoteWeightUnOverridden and is used to iterate over the raw logs and unpacked data for VoteWeightUnOverridden events raised by the Governance contract.
type GovernanceVoteWeightUnOverriddenIterator struct {
	Event *GovernanceVoteWeightUnOverridden // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceVoteWeightUnOverriddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceVoteWeightUnOverridden)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceVoteWeightUnOverridden)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceVoteWeightUnOverriddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceVoteWeightUnOverriddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceVoteWeightUnOverridden represents a VoteWeightUnOverridden event raised by the Governance contract.
type GovernanceVoteWeightUnOverridden struct {
	Voter common.Address
	Diff  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVoteWeightUnOverridden is a free log retrieval operation binding the contract event 0x11d7313426c62d856bd4ea20cdef8b93af4b40d2ea5b8f6f962fc705dbdcdbef.
//
// Solidity: event VoteWeightUnOverridden(address voter, uint256 diff)
func (_Governance *GovernanceFilterer) FilterVoteWeightUnOverridden(opts *bind.FilterOpts) (*GovernanceVoteWeightUnOverriddenIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "VoteWeightUnOverridden")
	if err != nil {
		return nil, err
	}
	return &GovernanceVoteWeightUnOverriddenIterator{contract: _Governance.contract, event: "VoteWeightUnOverridden", logs: logs, sub: sub}, nil
}

// WatchVoteWeightUnOverridden is a free log subscription operation binding the contract event 0x11d7313426c62d856bd4ea20cdef8b93af4b40d2ea5b8f6f962fc705dbdcdbef.
//
// Solidity: event VoteWeightUnOverridden(address voter, uint256 diff)
func (_Governance *GovernanceFilterer) WatchVoteWeightUnOverridden(opts *bind.WatchOpts, sink chan<- *GovernanceVoteWeightUnOverridden) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "VoteWeightUnOverridden")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceVoteWeightUnOverridden)
				if err := _Governance.contract.UnpackLog(event, "VoteWeightUnOverridden", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteWeightUnOverridden is a log parse operation binding the contract event 0x11d7313426c62d856bd4ea20cdef8b93af4b40d2ea5b8f6f962fc705dbdcdbef.
//
// Solidity: event VoteWeightUnOverridden(address voter, uint256 diff)
func (_Governance *GovernanceFilterer) ParseVoteWeightUnOverridden(log types.Log) (*GovernanceVoteWeightUnOverridden, error) {
	event := new(GovernanceVoteWeightUnOverridden)
	if err := _Governance.contract.UnpackLog(event, "VoteWeightUnOverridden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Governance contract.
type GovernanceVotedIterator struct {
	Event *GovernanceVoted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceVoted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceVoted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceVoted represents a Voted event raised by the Governance contract.
type GovernanceVoted struct {
	Voter       common.Address
	DelegatedTo common.Address
	ProposalID  *big.Int
	Choices     []*big.Int
	Weight      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x6e5f0f6e0ce2bdcdb0a82952fc6eb90c4c22f0b6228e4619b5dc2118e1166a12.
//
// Solidity: event Voted(address voter, address delegatedTo, uint256 proposalID, uint256[] choices, uint256 weight)
func (_Governance *GovernanceFilterer) FilterVoted(opts *bind.FilterOpts) (*GovernanceVotedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Voted")
	if err != nil {
		return nil, err
	}
	return &GovernanceVotedIterator{contract: _Governance.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x6e5f0f6e0ce2bdcdb0a82952fc6eb90c4c22f0b6228e4619b5dc2118e1166a12.
//
// Solidity: event Voted(address voter, address delegatedTo, uint256 proposalID, uint256[] choices, uint256 weight)
func (_Governance *GovernanceFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *GovernanceVoted) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Voted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceVoted)
				if err := _Governance.contract.UnpackLog(event, "Voted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoted is a log parse operation binding the contract event 0x6e5f0f6e0ce2bdcdb0a82952fc6eb90c4c22f0b6228e4619b5dc2118e1166a12.
//
// Solidity: event Voted(address voter, address delegatedTo, uint256 proposalID, uint256[] choices, uint256 weight)
func (_Governance *GovernanceFilterer) ParseVoted(log types.Log) (*GovernanceVoted, error) {
	event := new(GovernanceVoted)
	if err := _Governance.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

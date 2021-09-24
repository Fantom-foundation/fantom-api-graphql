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

// GovernanceProposalMetaData contains all meta data concerning the GovernanceProposal contract.
var GovernanceProposalMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"pType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"executable\",\"outputs\":[{\"internalType\":\"enumProposal.ExecType\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minAgreement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"opinionScales\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"options\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votingStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votingMinEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votingMaxEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"optionID\",\"type\":\"uint256\"}],\"name\":\"execute_call\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"selfAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionID\",\"type\":\"uint256\"}],\"name\":\"execute_delegatecall\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GovernanceProposalABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceProposalMetaData.ABI instead.
var GovernanceProposalABI = GovernanceProposalMetaData.ABI

// GovernanceProposal is an auto generated Go binding around an Ethereum contract.
type GovernanceProposal struct {
	GovernanceProposalCaller     // Read-only binding to the contract
	GovernanceProposalTransactor // Write-only binding to the contract
	GovernanceProposalFilterer   // Log filterer for contract events
}

// GovernanceProposalCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceProposalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceProposalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceProposalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceProposalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceProposalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceProposalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceProposalSession struct {
	Contract     *GovernanceProposal // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GovernanceProposalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceProposalCallerSession struct {
	Contract *GovernanceProposalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// GovernanceProposalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceProposalTransactorSession struct {
	Contract     *GovernanceProposalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// GovernanceProposalRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceProposalRaw struct {
	Contract *GovernanceProposal // Generic contract binding to access the raw methods on
}

// GovernanceProposalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceProposalCallerRaw struct {
	Contract *GovernanceProposalCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceProposalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceProposalTransactorRaw struct {
	Contract *GovernanceProposalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernanceProposal creates a new instance of GovernanceProposal, bound to a specific deployed contract.
func NewGovernanceProposal(address common.Address, backend bind.ContractBackend) (*GovernanceProposal, error) {
	contract, err := bindGovernanceProposal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposal{GovernanceProposalCaller: GovernanceProposalCaller{contract: contract}, GovernanceProposalTransactor: GovernanceProposalTransactor{contract: contract}, GovernanceProposalFilterer: GovernanceProposalFilterer{contract: contract}}, nil
}

// NewGovernanceProposalCaller creates a new read-only instance of GovernanceProposal, bound to a specific deployed contract.
func NewGovernanceProposalCaller(address common.Address, caller bind.ContractCaller) (*GovernanceProposalCaller, error) {
	contract, err := bindGovernanceProposal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalCaller{contract: contract}, nil
}

// NewGovernanceProposalTransactor creates a new write-only instance of GovernanceProposal, bound to a specific deployed contract.
func NewGovernanceProposalTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceProposalTransactor, error) {
	contract, err := bindGovernanceProposal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalTransactor{contract: contract}, nil
}

// NewGovernanceProposalFilterer creates a new log filterer instance of GovernanceProposal, bound to a specific deployed contract.
func NewGovernanceProposalFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceProposalFilterer, error) {
	contract, err := bindGovernanceProposal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalFilterer{contract: contract}, nil
}

// bindGovernanceProposal binds a generic wrapper to an already deployed contract.
func bindGovernanceProposal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceProposalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceProposal *GovernanceProposalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceProposal.Contract.GovernanceProposalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceProposal *GovernanceProposalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceProposal.Contract.GovernanceProposalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceProposal *GovernanceProposalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceProposal.Contract.GovernanceProposalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceProposal *GovernanceProposalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceProposal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceProposal *GovernanceProposalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceProposal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceProposal *GovernanceProposalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceProposal.Contract.contract.Transact(opts, method, params...)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_GovernanceProposal *GovernanceProposalCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GovernanceProposal.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_GovernanceProposal *GovernanceProposalSession) Description() (string, error) {
	return _GovernanceProposal.Contract.Description(&_GovernanceProposal.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_GovernanceProposal *GovernanceProposalCallerSession) Description() (string, error) {
	return _GovernanceProposal.Contract.Description(&_GovernanceProposal.CallOpts)
}

// Executable is a free data retrieval call binding the contract method 0x16131a18.
//
// Solidity: function executable() view returns(uint8)
func (_GovernanceProposal *GovernanceProposalCaller) Executable(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GovernanceProposal.contract.Call(opts, &out, "executable")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Executable is a free data retrieval call binding the contract method 0x16131a18.
//
// Solidity: function executable() view returns(uint8)
func (_GovernanceProposal *GovernanceProposalSession) Executable() (uint8, error) {
	return _GovernanceProposal.Contract.Executable(&_GovernanceProposal.CallOpts)
}

// Executable is a free data retrieval call binding the contract method 0x16131a18.
//
// Solidity: function executable() view returns(uint8)
func (_GovernanceProposal *GovernanceProposalCallerSession) Executable() (uint8, error) {
	return _GovernanceProposal.Contract.Executable(&_GovernanceProposal.CallOpts)
}

// MinAgreement is a free data retrieval call binding the contract method 0x6eb067a0.
//
// Solidity: function minAgreement() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalCaller) MinAgreement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceProposal.contract.Call(opts, &out, "minAgreement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAgreement is a free data retrieval call binding the contract method 0x6eb067a0.
//
// Solidity: function minAgreement() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalSession) MinAgreement() (*big.Int, error) {
	return _GovernanceProposal.Contract.MinAgreement(&_GovernanceProposal.CallOpts)
}

// MinAgreement is a free data retrieval call binding the contract method 0x6eb067a0.
//
// Solidity: function minAgreement() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalCallerSession) MinAgreement() (*big.Int, error) {
	return _GovernanceProposal.Contract.MinAgreement(&_GovernanceProposal.CallOpts)
}

// MinVotes is a free data retrieval call binding the contract method 0x3cc228fd.
//
// Solidity: function minVotes() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalCaller) MinVotes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceProposal.contract.Call(opts, &out, "minVotes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVotes is a free data retrieval call binding the contract method 0x3cc228fd.
//
// Solidity: function minVotes() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalSession) MinVotes() (*big.Int, error) {
	return _GovernanceProposal.Contract.MinVotes(&_GovernanceProposal.CallOpts)
}

// MinVotes is a free data retrieval call binding the contract method 0x3cc228fd.
//
// Solidity: function minVotes() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalCallerSession) MinVotes() (*big.Int, error) {
	return _GovernanceProposal.Contract.MinVotes(&_GovernanceProposal.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GovernanceProposal *GovernanceProposalCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GovernanceProposal.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GovernanceProposal *GovernanceProposalSession) Name() (string, error) {
	return _GovernanceProposal.Contract.Name(&_GovernanceProposal.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GovernanceProposal *GovernanceProposalCallerSession) Name() (string, error) {
	return _GovernanceProposal.Contract.Name(&_GovernanceProposal.CallOpts)
}

// OpinionScales is a free data retrieval call binding the contract method 0x6f24f8a1.
//
// Solidity: function opinionScales() view returns(uint256[])
func (_GovernanceProposal *GovernanceProposalCaller) OpinionScales(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _GovernanceProposal.contract.Call(opts, &out, "opinionScales")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// OpinionScales is a free data retrieval call binding the contract method 0x6f24f8a1.
//
// Solidity: function opinionScales() view returns(uint256[])
func (_GovernanceProposal *GovernanceProposalSession) OpinionScales() ([]*big.Int, error) {
	return _GovernanceProposal.Contract.OpinionScales(&_GovernanceProposal.CallOpts)
}

// OpinionScales is a free data retrieval call binding the contract method 0x6f24f8a1.
//
// Solidity: function opinionScales() view returns(uint256[])
func (_GovernanceProposal *GovernanceProposalCallerSession) OpinionScales() ([]*big.Int, error) {
	return _GovernanceProposal.Contract.OpinionScales(&_GovernanceProposal.CallOpts)
}

// Options is a free data retrieval call binding the contract method 0x1069143a.
//
// Solidity: function options() view returns(bytes32[])
func (_GovernanceProposal *GovernanceProposalCaller) Options(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _GovernanceProposal.contract.Call(opts, &out, "options")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// Options is a free data retrieval call binding the contract method 0x1069143a.
//
// Solidity: function options() view returns(bytes32[])
func (_GovernanceProposal *GovernanceProposalSession) Options() ([][32]byte, error) {
	return _GovernanceProposal.Contract.Options(&_GovernanceProposal.CallOpts)
}

// Options is a free data retrieval call binding the contract method 0x1069143a.
//
// Solidity: function options() view returns(bytes32[])
func (_GovernanceProposal *GovernanceProposalCallerSession) Options() ([][32]byte, error) {
	return _GovernanceProposal.Contract.Options(&_GovernanceProposal.CallOpts)
}

// PType is a free data retrieval call binding the contract method 0x216b9116.
//
// Solidity: function pType() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalCaller) PType(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceProposal.contract.Call(opts, &out, "pType")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PType is a free data retrieval call binding the contract method 0x216b9116.
//
// Solidity: function pType() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalSession) PType() (*big.Int, error) {
	return _GovernanceProposal.Contract.PType(&_GovernanceProposal.CallOpts)
}

// PType is a free data retrieval call binding the contract method 0x216b9116.
//
// Solidity: function pType() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalCallerSession) PType() (*big.Int, error) {
	return _GovernanceProposal.Contract.PType(&_GovernanceProposal.CallOpts)
}

// VotingMaxEndTime is a free data retrieval call binding the contract method 0xa61e393a.
//
// Solidity: function votingMaxEndTime() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalCaller) VotingMaxEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceProposal.contract.Call(opts, &out, "votingMaxEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingMaxEndTime is a free data retrieval call binding the contract method 0xa61e393a.
//
// Solidity: function votingMaxEndTime() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalSession) VotingMaxEndTime() (*big.Int, error) {
	return _GovernanceProposal.Contract.VotingMaxEndTime(&_GovernanceProposal.CallOpts)
}

// VotingMaxEndTime is a free data retrieval call binding the contract method 0xa61e393a.
//
// Solidity: function votingMaxEndTime() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalCallerSession) VotingMaxEndTime() (*big.Int, error) {
	return _GovernanceProposal.Contract.VotingMaxEndTime(&_GovernanceProposal.CallOpts)
}

// VotingMinEndTime is a free data retrieval call binding the contract method 0x9b3a8a07.
//
// Solidity: function votingMinEndTime() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalCaller) VotingMinEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceProposal.contract.Call(opts, &out, "votingMinEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingMinEndTime is a free data retrieval call binding the contract method 0x9b3a8a07.
//
// Solidity: function votingMinEndTime() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalSession) VotingMinEndTime() (*big.Int, error) {
	return _GovernanceProposal.Contract.VotingMinEndTime(&_GovernanceProposal.CallOpts)
}

// VotingMinEndTime is a free data retrieval call binding the contract method 0x9b3a8a07.
//
// Solidity: function votingMinEndTime() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalCallerSession) VotingMinEndTime() (*big.Int, error) {
	return _GovernanceProposal.Contract.VotingMinEndTime(&_GovernanceProposal.CallOpts)
}

// VotingStartTime is a free data retrieval call binding the contract method 0x2019a608.
//
// Solidity: function votingStartTime() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalCaller) VotingStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceProposal.contract.Call(opts, &out, "votingStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingStartTime is a free data retrieval call binding the contract method 0x2019a608.
//
// Solidity: function votingStartTime() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalSession) VotingStartTime() (*big.Int, error) {
	return _GovernanceProposal.Contract.VotingStartTime(&_GovernanceProposal.CallOpts)
}

// VotingStartTime is a free data retrieval call binding the contract method 0x2019a608.
//
// Solidity: function votingStartTime() view returns(uint256)
func (_GovernanceProposal *GovernanceProposalCallerSession) VotingStartTime() (*big.Int, error) {
	return _GovernanceProposal.Contract.VotingStartTime(&_GovernanceProposal.CallOpts)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x63012d01.
//
// Solidity: function execute_call(uint256 optionID) returns()
func (_GovernanceProposal *GovernanceProposalTransactor) ExecuteCall(opts *bind.TransactOpts, optionID *big.Int) (*types.Transaction, error) {
	return _GovernanceProposal.contract.Transact(opts, "execute_call", optionID)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x63012d01.
//
// Solidity: function execute_call(uint256 optionID) returns()
func (_GovernanceProposal *GovernanceProposalSession) ExecuteCall(optionID *big.Int) (*types.Transaction, error) {
	return _GovernanceProposal.Contract.ExecuteCall(&_GovernanceProposal.TransactOpts, optionID)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x63012d01.
//
// Solidity: function execute_call(uint256 optionID) returns()
func (_GovernanceProposal *GovernanceProposalTransactorSession) ExecuteCall(optionID *big.Int) (*types.Transaction, error) {
	return _GovernanceProposal.Contract.ExecuteCall(&_GovernanceProposal.TransactOpts, optionID)
}

// ExecuteDelegatecall is a paid mutator transaction binding the contract method 0xaa3daba9.
//
// Solidity: function execute_delegatecall(address selfAddress, uint256 optionID) returns()
func (_GovernanceProposal *GovernanceProposalTransactor) ExecuteDelegatecall(opts *bind.TransactOpts, selfAddress common.Address, optionID *big.Int) (*types.Transaction, error) {
	return _GovernanceProposal.contract.Transact(opts, "execute_delegatecall", selfAddress, optionID)
}

// ExecuteDelegatecall is a paid mutator transaction binding the contract method 0xaa3daba9.
//
// Solidity: function execute_delegatecall(address selfAddress, uint256 optionID) returns()
func (_GovernanceProposal *GovernanceProposalSession) ExecuteDelegatecall(selfAddress common.Address, optionID *big.Int) (*types.Transaction, error) {
	return _GovernanceProposal.Contract.ExecuteDelegatecall(&_GovernanceProposal.TransactOpts, selfAddress, optionID)
}

// ExecuteDelegatecall is a paid mutator transaction binding the contract method 0xaa3daba9.
//
// Solidity: function execute_delegatecall(address selfAddress, uint256 optionID) returns()
func (_GovernanceProposal *GovernanceProposalTransactorSession) ExecuteDelegatecall(selfAddress common.Address, optionID *big.Int) (*types.Transaction, error) {
	return _GovernanceProposal.Contract.ExecuteDelegatecall(&_GovernanceProposal.TransactOpts, selfAddress, optionID)
}

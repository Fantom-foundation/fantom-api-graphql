// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rpc

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BallotContractABI is the input ABI used to generate the binding from.
const BallotContractABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proposalNames\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ballot\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winner\",\"type\":\"uint256\"}],\"name\":\"Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ballot\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ballot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeds\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chairperson\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"totals\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stamps\",\"type\":\"uint256[]\"}],\"name\":\"feedWeights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposalsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposal\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightStamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BallotContract is an auto generated Go binding around an Ethereum contract.
type BallotContract struct {
	BallotContractCaller     // Read-only binding to the contract
	BallotContractTransactor // Write-only binding to the contract
	BallotContractFilterer   // Log filterer for contract events
}

// BallotContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BallotContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BallotContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BallotContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BallotContractSession struct {
	Contract     *BallotContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BallotContractCallerSession struct {
	Contract *BallotContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BallotContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BallotContractTransactorSession struct {
	Contract     *BallotContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BallotContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BallotContractRaw struct {
	Contract *BallotContract // Generic contract binding to access the raw methods on
}

// BallotContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BallotContractCallerRaw struct {
	Contract *BallotContractCaller // Generic read-only contract binding to access the raw methods on
}

// BallotContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BallotContractTransactorRaw struct {
	Contract *BallotContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBallotContract creates a new instance of BallotContract, bound to a specific deployed contract.
func NewBallotContract(address common.Address, backend bind.ContractBackend) (*BallotContract, error) {
	contract, err := bindBallotContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BallotContract{BallotContractCaller: BallotContractCaller{contract: contract}, BallotContractTransactor: BallotContractTransactor{contract: contract}, BallotContractFilterer: BallotContractFilterer{contract: contract}}, nil
}

// NewBallotContractCaller creates a new read-only instance of BallotContract, bound to a specific deployed contract.
func NewBallotContractCaller(address common.Address, caller bind.ContractCaller) (*BallotContractCaller, error) {
	contract, err := bindBallotContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BallotContractCaller{contract: contract}, nil
}

// NewBallotContractTransactor creates a new write-only instance of BallotContract, bound to a specific deployed contract.
func NewBallotContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BallotContractTransactor, error) {
	contract, err := bindBallotContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BallotContractTransactor{contract: contract}, nil
}

// NewBallotContractFilterer creates a new log filterer instance of BallotContract, bound to a specific deployed contract.
func NewBallotContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BallotContractFilterer, error) {
	contract, err := bindBallotContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BallotContractFilterer{contract: contract}, nil
}

// bindBallotContract binds a generic wrapper to an already deployed contract.
func bindBallotContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BallotContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BallotContract *BallotContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BallotContract.Contract.BallotContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BallotContract *BallotContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BallotContract.Contract.BallotContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BallotContract *BallotContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BallotContract.Contract.BallotContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BallotContract *BallotContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BallotContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BallotContract *BallotContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BallotContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BallotContract *BallotContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BallotContract.Contract.contract.Transact(opts, method, params...)
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(bytes32 name, string url, uint256 start, uint256 end, bool finalized, uint256 votes, uint256 feeds)
func (_BallotContract *BallotContractCaller) Ballot(opts *bind.CallOpts) (struct {
	Name      [32]byte
	Url       string
	Start     *big.Int
	End       *big.Int
	Finalized bool
	Votes     *big.Int
	Feeds     *big.Int
}, error) {
	ret := new(struct {
		Name      [32]byte
		Url       string
		Start     *big.Int
		End       *big.Int
		Finalized bool
		Votes     *big.Int
		Feeds     *big.Int
	})
	out := ret
	err := _BallotContract.contract.Call(opts, out, "ballot")
	return *ret, err
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(bytes32 name, string url, uint256 start, uint256 end, bool finalized, uint256 votes, uint256 feeds)
func (_BallotContract *BallotContractSession) Ballot() (struct {
	Name      [32]byte
	Url       string
	Start     *big.Int
	End       *big.Int
	Finalized bool
	Votes     *big.Int
	Feeds     *big.Int
}, error) {
	return _BallotContract.Contract.Ballot(&_BallotContract.CallOpts)
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(bytes32 name, string url, uint256 start, uint256 end, bool finalized, uint256 votes, uint256 feeds)
func (_BallotContract *BallotContractCallerSession) Ballot() (struct {
	Name      [32]byte
	Url       string
	Start     *big.Int
	End       *big.Int
	Finalized bool
	Votes     *big.Int
	Feeds     *big.Int
}, error) {
	return _BallotContract.Contract.Ballot(&_BallotContract.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_BallotContract *BallotContractCaller) Chairperson(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BallotContract.contract.Call(opts, out, "chairperson")
	return *ret0, err
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_BallotContract *BallotContractSession) Chairperson() (common.Address, error) {
	return _BallotContract.Contract.Chairperson(&_BallotContract.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_BallotContract *BallotContractCallerSession) Chairperson() (common.Address, error) {
	return _BallotContract.Contract.Chairperson(&_BallotContract.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 weight, uint256 votes)
func (_BallotContract *BallotContractCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name   [32]byte
	Weight *big.Int
	Votes  *big.Int
}, error) {
	ret := new(struct {
		Name   [32]byte
		Weight *big.Int
		Votes  *big.Int
	})
	out := ret
	err := _BallotContract.contract.Call(opts, out, "proposals", arg0)
	return *ret, err
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 weight, uint256 votes)
func (_BallotContract *BallotContractSession) Proposals(arg0 *big.Int) (struct {
	Name   [32]byte
	Weight *big.Int
	Votes  *big.Int
}, error) {
	return _BallotContract.Contract.Proposals(&_BallotContract.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 weight, uint256 votes)
func (_BallotContract *BallotContractCallerSession) Proposals(arg0 *big.Int) (struct {
	Name   [32]byte
	Weight *big.Int
	Votes  *big.Int
}, error) {
	return _BallotContract.Contract.Proposals(&_BallotContract.CallOpts, arg0)
}

// ProposalsCount is a free data retrieval call binding the contract method 0x0a9f46ad.
//
// Solidity: function proposalsCount() view returns(uint256)
func (_BallotContract *BallotContractCaller) ProposalsCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BallotContract.contract.Call(opts, out, "proposalsCount")
	return *ret0, err
}

// ProposalsCount is a free data retrieval call binding the contract method 0x0a9f46ad.
//
// Solidity: function proposalsCount() view returns(uint256)
func (_BallotContract *BallotContractSession) ProposalsCount() (*big.Int, error) {
	return _BallotContract.Contract.ProposalsCount(&_BallotContract.CallOpts)
}

// ProposalsCount is a free data retrieval call binding the contract method 0x0a9f46ad.
//
// Solidity: function proposalsCount() view returns(uint256)
func (_BallotContract *BallotContractCallerSession) ProposalsCount() (*big.Int, error) {
	return _BallotContract.Contract.ProposalsCount(&_BallotContract.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint256 vote, uint256 voted, uint256 weight, uint256 weightStamp)
func (_BallotContract *BallotContractCaller) Votes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Vote        *big.Int
	Voted       *big.Int
	Weight      *big.Int
	WeightStamp *big.Int
}, error) {
	ret := new(struct {
		Vote        *big.Int
		Voted       *big.Int
		Weight      *big.Int
		WeightStamp *big.Int
	})
	out := ret
	err := _BallotContract.contract.Call(opts, out, "votes", arg0)
	return *ret, err
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint256 vote, uint256 voted, uint256 weight, uint256 weightStamp)
func (_BallotContract *BallotContractSession) Votes(arg0 common.Address) (struct {
	Vote        *big.Int
	Voted       *big.Int
	Weight      *big.Int
	WeightStamp *big.Int
}, error) {
	return _BallotContract.Contract.Votes(&_BallotContract.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint256 vote, uint256 voted, uint256 weight, uint256 weightStamp)
func (_BallotContract *BallotContractCallerSession) Votes(arg0 common.Address) (struct {
	Vote        *big.Int
	Voted       *big.Int
	Weight      *big.Int
	WeightStamp *big.Int
}, error) {
	return _BallotContract.Contract.Votes(&_BallotContract.CallOpts, arg0)
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(uint256, uint256, bytes32)
func (_BallotContract *BallotContractCaller) Winner(opts *bind.CallOpts) (*big.Int, *big.Int, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _BallotContract.contract.Call(opts, out, "winner")
	return *ret0, *ret1, *ret2, err
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(uint256, uint256, bytes32)
func (_BallotContract *BallotContractSession) Winner() (*big.Int, *big.Int, [32]byte, error) {
	return _BallotContract.Contract.Winner(&_BallotContract.CallOpts)
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(uint256, uint256, bytes32)
func (_BallotContract *BallotContractCallerSession) Winner() (*big.Int, *big.Int, [32]byte, error) {
	return _BallotContract.Contract.Winner(&_BallotContract.CallOpts)
}

// FeedWeights is a paid mutator transaction binding the contract method 0xee2ef228.
//
// Solidity: function feedWeights(address[] voters, uint256[] totals, uint256[] stamps) returns()
func (_BallotContract *BallotContractTransactor) FeedWeights(opts *bind.TransactOpts, voters []common.Address, totals []*big.Int, stamps []*big.Int) (*types.Transaction, error) {
	return _BallotContract.contract.Transact(opts, "feedWeights", voters, totals, stamps)
}

// FeedWeights is a paid mutator transaction binding the contract method 0xee2ef228.
//
// Solidity: function feedWeights(address[] voters, uint256[] totals, uint256[] stamps) returns()
func (_BallotContract *BallotContractSession) FeedWeights(voters []common.Address, totals []*big.Int, stamps []*big.Int) (*types.Transaction, error) {
	return _BallotContract.Contract.FeedWeights(&_BallotContract.TransactOpts, voters, totals, stamps)
}

// FeedWeights is a paid mutator transaction binding the contract method 0xee2ef228.
//
// Solidity: function feedWeights(address[] voters, uint256[] totals, uint256[] stamps) returns()
func (_BallotContract *BallotContractTransactorSession) FeedWeights(voters []common.Address, totals []*big.Int, stamps []*big.Int) (*types.Transaction, error) {
	return _BallotContract.Contract.FeedWeights(&_BallotContract.TransactOpts, voters, totals, stamps)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BallotContract *BallotContractTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BallotContract.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BallotContract *BallotContractSession) Finalize() (*types.Transaction, error) {
	return _BallotContract.Contract.Finalize(&_BallotContract.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BallotContract *BallotContractTransactorSession) Finalize() (*types.Transaction, error) {
	return _BallotContract.Contract.Finalize(&_BallotContract.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_BallotContract *BallotContractTransactor) Vote(opts *bind.TransactOpts, proposal *big.Int) (*types.Transaction, error) {
	return _BallotContract.contract.Transact(opts, "vote", proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_BallotContract *BallotContractSession) Vote(proposal *big.Int) (*types.Transaction, error) {
	return _BallotContract.Contract.Vote(&_BallotContract.TransactOpts, proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_BallotContract *BallotContractTransactorSession) Vote(proposal *big.Int) (*types.Transaction, error) {
	return _BallotContract.Contract.Vote(&_BallotContract.TransactOpts, proposal)
}

// BallotContractFinalizedIterator is returned from FilterFinalized and is used to iterate over the raw logs and unpacked data for Finalized events raised by the BallotContract contract.
type BallotContractFinalizedIterator struct {
	Event *BallotContractFinalized // Event containing the contract specifics and raw log

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
func (it *BallotContractFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotContractFinalized)
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
		it.Event = new(BallotContractFinalized)
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
func (it *BallotContractFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotContractFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotContractFinalized represents a Finalized event raised by the BallotContract contract.
type BallotContractFinalized struct {
	Ballot common.Address
	Winner *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalized is a free log retrieval operation binding the contract event 0x66b6851664a82efe6b871e434faba2b11421d2dad65eb71a344ae76cca8a2b86.
//
// Solidity: event Finalized(address indexed ballot, uint256 winner)
func (_BallotContract *BallotContractFilterer) FilterFinalized(opts *bind.FilterOpts, ballot []common.Address) (*BallotContractFinalizedIterator, error) {

	var ballotRule []interface{}
	for _, ballotItem := range ballot {
		ballotRule = append(ballotRule, ballotItem)
	}

	logs, sub, err := _BallotContract.contract.FilterLogs(opts, "Finalized", ballotRule)
	if err != nil {
		return nil, err
	}
	return &BallotContractFinalizedIterator{contract: _BallotContract.contract, event: "Finalized", logs: logs, sub: sub}, nil
}

// WatchFinalized is a free log subscription operation binding the contract event 0x66b6851664a82efe6b871e434faba2b11421d2dad65eb71a344ae76cca8a2b86.
//
// Solidity: event Finalized(address indexed ballot, uint256 winner)
func (_BallotContract *BallotContractFilterer) WatchFinalized(opts *bind.WatchOpts, sink chan<- *BallotContractFinalized, ballot []common.Address) (event.Subscription, error) {

	var ballotRule []interface{}
	for _, ballotItem := range ballot {
		ballotRule = append(ballotRule, ballotItem)
	}

	logs, sub, err := _BallotContract.contract.WatchLogs(opts, "Finalized", ballotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotContractFinalized)
				if err := _BallotContract.contract.UnpackLog(event, "Finalized", log); err != nil {
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

// ParseFinalized is a log parse operation binding the contract event 0x66b6851664a82efe6b871e434faba2b11421d2dad65eb71a344ae76cca8a2b86.
//
// Solidity: event Finalized(address indexed ballot, uint256 winner)
func (_BallotContract *BallotContractFilterer) ParseFinalized(log types.Log) (*BallotContractFinalized, error) {
	event := new(BallotContractFinalized)
	if err := _BallotContract.contract.UnpackLog(event, "Finalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BallotContractVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the BallotContract contract.
type BallotContractVotedIterator struct {
	Event *BallotContractVoted // Event containing the contract specifics and raw log

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
func (it *BallotContractVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotContractVoted)
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
		it.Event = new(BallotContractVoted)
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
func (it *BallotContractVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotContractVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotContractVoted represents a Voted event raised by the BallotContract contract.
type BallotContractVoted struct {
	Ballot common.Address
	Voter  common.Address
	Vote   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x174ba19ba3c8bb5c679c87e51db79fff7c3f04bb84c1fd55b7dacb470b674aa6.
//
// Solidity: event Voted(address indexed ballot, address indexed voter, uint256 vote)
func (_BallotContract *BallotContractFilterer) FilterVoted(opts *bind.FilterOpts, ballot []common.Address, voter []common.Address) (*BallotContractVotedIterator, error) {

	var ballotRule []interface{}
	for _, ballotItem := range ballot {
		ballotRule = append(ballotRule, ballotItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _BallotContract.contract.FilterLogs(opts, "Voted", ballotRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &BallotContractVotedIterator{contract: _BallotContract.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x174ba19ba3c8bb5c679c87e51db79fff7c3f04bb84c1fd55b7dacb470b674aa6.
//
// Solidity: event Voted(address indexed ballot, address indexed voter, uint256 vote)
func (_BallotContract *BallotContractFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *BallotContractVoted, ballot []common.Address, voter []common.Address) (event.Subscription, error) {

	var ballotRule []interface{}
	for _, ballotItem := range ballot {
		ballotRule = append(ballotRule, ballotItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _BallotContract.contract.WatchLogs(opts, "Voted", ballotRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotContractVoted)
				if err := _BallotContract.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x174ba19ba3c8bb5c679c87e51db79fff7c3f04bb84c1fd55b7dacb470b674aa6.
//
// Solidity: event Voted(address indexed ballot, address indexed voter, uint256 vote)
func (_BallotContract *BallotContractFilterer) ParseVoted(log types.Log) (*BallotContractVoted, error) {
	event := new(BallotContractVoted)
	if err := _BallotContract.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	return event, nil
}

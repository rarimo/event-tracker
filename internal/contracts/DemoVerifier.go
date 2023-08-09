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
	_ = abi.ConvertType
)

// IDemoVerifierIdentityProofInfo is an auto generated low-level Go binding around an user-defined struct.
type IDemoVerifierIdentityProofInfo struct {
	SenderAddr common.Address
	IsProved   bool
}

// DemoVerifierMetaData contains all meta data concerning the DemoVerifier contract.
var DemoVerifierMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"identityId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"senderAddr\",\"type\":\"address\"}],\"name\":\"IdentityProved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIZKPQueriesStorage\",\"name\":\"zkpQueriesStorage_\",\"type\":\"address\"}],\"name\":\"__DemoVerifier_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressToIdentityId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"identityId_\",\"type\":\"uint256\"}],\"name\":\"getIdentityProofInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"senderAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isProved\",\"type\":\"bool\"}],\"internalType\":\"structIDemoVerifier.IdentityProofInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"identityId_\",\"type\":\"uint256\"}],\"name\":\"isIdentityProved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"inputs_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"a_\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b_\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c_\",\"type\":\"uint256[2]\"}],\"name\":\"proveIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIZKPQueriesStorage\",\"name\":\"newZKPQueriesStorage_\",\"type\":\"address\"}],\"name\":\"setZKPQueriesStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zkpQueriesStorage\",\"outputs\":[{\"internalType\":\"contractIZKPQueriesStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DemoVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use DemoVerifierMetaData.ABI instead.
var DemoVerifierABI = DemoVerifierMetaData.ABI

// DemoVerifier is an auto generated Go binding around an Ethereum contract.
type DemoVerifier struct {
	DemoVerifierCaller     // Read-only binding to the contract
	DemoVerifierTransactor // Write-only binding to the contract
	DemoVerifierFilterer   // Log filterer for contract events
}

// DemoVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type DemoVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DemoVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DemoVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DemoVerifierSession struct {
	Contract     *DemoVerifier     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DemoVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DemoVerifierCallerSession struct {
	Contract *DemoVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DemoVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DemoVerifierTransactorSession struct {
	Contract     *DemoVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DemoVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type DemoVerifierRaw struct {
	Contract *DemoVerifier // Generic contract binding to access the raw methods on
}

// DemoVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DemoVerifierCallerRaw struct {
	Contract *DemoVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// DemoVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DemoVerifierTransactorRaw struct {
	Contract *DemoVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDemoVerifier creates a new instance of DemoVerifier, bound to a specific deployed contract.
func NewDemoVerifier(address common.Address, backend bind.ContractBackend) (*DemoVerifier, error) {
	contract, err := bindDemoVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DemoVerifier{DemoVerifierCaller: DemoVerifierCaller{contract: contract}, DemoVerifierTransactor: DemoVerifierTransactor{contract: contract}, DemoVerifierFilterer: DemoVerifierFilterer{contract: contract}}, nil
}

// NewDemoVerifierCaller creates a new read-only instance of DemoVerifier, bound to a specific deployed contract.
func NewDemoVerifierCaller(address common.Address, caller bind.ContractCaller) (*DemoVerifierCaller, error) {
	contract, err := bindDemoVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DemoVerifierCaller{contract: contract}, nil
}

// NewDemoVerifierTransactor creates a new write-only instance of DemoVerifier, bound to a specific deployed contract.
func NewDemoVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*DemoVerifierTransactor, error) {
	contract, err := bindDemoVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DemoVerifierTransactor{contract: contract}, nil
}

// NewDemoVerifierFilterer creates a new log filterer instance of DemoVerifier, bound to a specific deployed contract.
func NewDemoVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*DemoVerifierFilterer, error) {
	contract, err := bindDemoVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DemoVerifierFilterer{contract: contract}, nil
}

// bindDemoVerifier binds a generic wrapper to an already deployed contract.
func bindDemoVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DemoVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DemoVerifier *DemoVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DemoVerifier.Contract.DemoVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DemoVerifier *DemoVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DemoVerifier.Contract.DemoVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DemoVerifier *DemoVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DemoVerifier.Contract.DemoVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DemoVerifier *DemoVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DemoVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DemoVerifier *DemoVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DemoVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DemoVerifier *DemoVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DemoVerifier.Contract.contract.Transact(opts, method, params...)
}

// AddressToIdentityId is a free data retrieval call binding the contract method 0xb4528ff9.
//
// Solidity: function addressToIdentityId(address ) view returns(uint256)
func (_DemoVerifier *DemoVerifierCaller) AddressToIdentityId(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DemoVerifier.contract.Call(opts, &out, "addressToIdentityId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToIdentityId is a free data retrieval call binding the contract method 0xb4528ff9.
//
// Solidity: function addressToIdentityId(address ) view returns(uint256)
func (_DemoVerifier *DemoVerifierSession) AddressToIdentityId(arg0 common.Address) (*big.Int, error) {
	return _DemoVerifier.Contract.AddressToIdentityId(&_DemoVerifier.CallOpts, arg0)
}

// AddressToIdentityId is a free data retrieval call binding the contract method 0xb4528ff9.
//
// Solidity: function addressToIdentityId(address ) view returns(uint256)
func (_DemoVerifier *DemoVerifierCallerSession) AddressToIdentityId(arg0 common.Address) (*big.Int, error) {
	return _DemoVerifier.Contract.AddressToIdentityId(&_DemoVerifier.CallOpts, arg0)
}

// GetIdentityProofInfo is a free data retrieval call binding the contract method 0x5332d5ec.
//
// Solidity: function getIdentityProofInfo(uint256 identityId_) view returns((address,bool))
func (_DemoVerifier *DemoVerifierCaller) GetIdentityProofInfo(opts *bind.CallOpts, identityId_ *big.Int) (IDemoVerifierIdentityProofInfo, error) {
	var out []interface{}
	err := _DemoVerifier.contract.Call(opts, &out, "getIdentityProofInfo", identityId_)

	if err != nil {
		return *new(IDemoVerifierIdentityProofInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IDemoVerifierIdentityProofInfo)).(*IDemoVerifierIdentityProofInfo)

	return out0, err

}

// GetIdentityProofInfo is a free data retrieval call binding the contract method 0x5332d5ec.
//
// Solidity: function getIdentityProofInfo(uint256 identityId_) view returns((address,bool))
func (_DemoVerifier *DemoVerifierSession) GetIdentityProofInfo(identityId_ *big.Int) (IDemoVerifierIdentityProofInfo, error) {
	return _DemoVerifier.Contract.GetIdentityProofInfo(&_DemoVerifier.CallOpts, identityId_)
}

// GetIdentityProofInfo is a free data retrieval call binding the contract method 0x5332d5ec.
//
// Solidity: function getIdentityProofInfo(uint256 identityId_) view returns((address,bool))
func (_DemoVerifier *DemoVerifierCallerSession) GetIdentityProofInfo(identityId_ *big.Int) (IDemoVerifierIdentityProofInfo, error) {
	return _DemoVerifier.Contract.GetIdentityProofInfo(&_DemoVerifier.CallOpts, identityId_)
}

// IsIdentityProved is a free data retrieval call binding the contract method 0x5428764d.
//
// Solidity: function isIdentityProved(uint256 identityId_) view returns(bool)
func (_DemoVerifier *DemoVerifierCaller) IsIdentityProved(opts *bind.CallOpts, identityId_ *big.Int) (bool, error) {
	var out []interface{}
	err := _DemoVerifier.contract.Call(opts, &out, "isIdentityProved", identityId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIdentityProved is a free data retrieval call binding the contract method 0x5428764d.
//
// Solidity: function isIdentityProved(uint256 identityId_) view returns(bool)
func (_DemoVerifier *DemoVerifierSession) IsIdentityProved(identityId_ *big.Int) (bool, error) {
	return _DemoVerifier.Contract.IsIdentityProved(&_DemoVerifier.CallOpts, identityId_)
}

// IsIdentityProved is a free data retrieval call binding the contract method 0x5428764d.
//
// Solidity: function isIdentityProved(uint256 identityId_) view returns(bool)
func (_DemoVerifier *DemoVerifierCallerSession) IsIdentityProved(identityId_ *big.Int) (bool, error) {
	return _DemoVerifier.Contract.IsIdentityProved(&_DemoVerifier.CallOpts, identityId_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DemoVerifier *DemoVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DemoVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DemoVerifier *DemoVerifierSession) Owner() (common.Address, error) {
	return _DemoVerifier.Contract.Owner(&_DemoVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DemoVerifier *DemoVerifierCallerSession) Owner() (common.Address, error) {
	return _DemoVerifier.Contract.Owner(&_DemoVerifier.CallOpts)
}

// ZkpQueriesStorage is a free data retrieval call binding the contract method 0xb4db08cc.
//
// Solidity: function zkpQueriesStorage() view returns(address)
func (_DemoVerifier *DemoVerifierCaller) ZkpQueriesStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DemoVerifier.contract.Call(opts, &out, "zkpQueriesStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkpQueriesStorage is a free data retrieval call binding the contract method 0xb4db08cc.
//
// Solidity: function zkpQueriesStorage() view returns(address)
func (_DemoVerifier *DemoVerifierSession) ZkpQueriesStorage() (common.Address, error) {
	return _DemoVerifier.Contract.ZkpQueriesStorage(&_DemoVerifier.CallOpts)
}

// ZkpQueriesStorage is a free data retrieval call binding the contract method 0xb4db08cc.
//
// Solidity: function zkpQueriesStorage() view returns(address)
func (_DemoVerifier *DemoVerifierCallerSession) ZkpQueriesStorage() (common.Address, error) {
	return _DemoVerifier.Contract.ZkpQueriesStorage(&_DemoVerifier.CallOpts)
}

// DemoVerifierInit is a paid mutator transaction binding the contract method 0xc18d9a8a.
//
// Solidity: function __DemoVerifier_init(address zkpQueriesStorage_) returns()
func (_DemoVerifier *DemoVerifierTransactor) DemoVerifierInit(opts *bind.TransactOpts, zkpQueriesStorage_ common.Address) (*types.Transaction, error) {
	return _DemoVerifier.contract.Transact(opts, "__DemoVerifier_init", zkpQueriesStorage_)
}

// DemoVerifierInit is a paid mutator transaction binding the contract method 0xc18d9a8a.
//
// Solidity: function __DemoVerifier_init(address zkpQueriesStorage_) returns()
func (_DemoVerifier *DemoVerifierSession) DemoVerifierInit(zkpQueriesStorage_ common.Address) (*types.Transaction, error) {
	return _DemoVerifier.Contract.DemoVerifierInit(&_DemoVerifier.TransactOpts, zkpQueriesStorage_)
}

// DemoVerifierInit is a paid mutator transaction binding the contract method 0xc18d9a8a.
//
// Solidity: function __DemoVerifier_init(address zkpQueriesStorage_) returns()
func (_DemoVerifier *DemoVerifierTransactorSession) DemoVerifierInit(zkpQueriesStorage_ common.Address) (*types.Transaction, error) {
	return _DemoVerifier.Contract.DemoVerifierInit(&_DemoVerifier.TransactOpts, zkpQueriesStorage_)
}

// ProveIdentity is a paid mutator transaction binding the contract method 0xa2f71251.
//
// Solidity: function proveIdentity(uint256[] inputs_, uint256[2] a_, uint256[2][2] b_, uint256[2] c_) returns()
func (_DemoVerifier *DemoVerifierTransactor) ProveIdentity(opts *bind.TransactOpts, inputs_ []*big.Int, a_ [2]*big.Int, b_ [2][2]*big.Int, c_ [2]*big.Int) (*types.Transaction, error) {
	return _DemoVerifier.contract.Transact(opts, "proveIdentity", inputs_, a_, b_, c_)
}

// ProveIdentity is a paid mutator transaction binding the contract method 0xa2f71251.
//
// Solidity: function proveIdentity(uint256[] inputs_, uint256[2] a_, uint256[2][2] b_, uint256[2] c_) returns()
func (_DemoVerifier *DemoVerifierSession) ProveIdentity(inputs_ []*big.Int, a_ [2]*big.Int, b_ [2][2]*big.Int, c_ [2]*big.Int) (*types.Transaction, error) {
	return _DemoVerifier.Contract.ProveIdentity(&_DemoVerifier.TransactOpts, inputs_, a_, b_, c_)
}

// ProveIdentity is a paid mutator transaction binding the contract method 0xa2f71251.
//
// Solidity: function proveIdentity(uint256[] inputs_, uint256[2] a_, uint256[2][2] b_, uint256[2] c_) returns()
func (_DemoVerifier *DemoVerifierTransactorSession) ProveIdentity(inputs_ []*big.Int, a_ [2]*big.Int, b_ [2][2]*big.Int, c_ [2]*big.Int) (*types.Transaction, error) {
	return _DemoVerifier.Contract.ProveIdentity(&_DemoVerifier.TransactOpts, inputs_, a_, b_, c_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DemoVerifier *DemoVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DemoVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DemoVerifier *DemoVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _DemoVerifier.Contract.RenounceOwnership(&_DemoVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DemoVerifier *DemoVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DemoVerifier.Contract.RenounceOwnership(&_DemoVerifier.TransactOpts)
}

// SetZKPQueriesStorage is a paid mutator transaction binding the contract method 0xddebe5c0.
//
// Solidity: function setZKPQueriesStorage(address newZKPQueriesStorage_) returns()
func (_DemoVerifier *DemoVerifierTransactor) SetZKPQueriesStorage(opts *bind.TransactOpts, newZKPQueriesStorage_ common.Address) (*types.Transaction, error) {
	return _DemoVerifier.contract.Transact(opts, "setZKPQueriesStorage", newZKPQueriesStorage_)
}

// SetZKPQueriesStorage is a paid mutator transaction binding the contract method 0xddebe5c0.
//
// Solidity: function setZKPQueriesStorage(address newZKPQueriesStorage_) returns()
func (_DemoVerifier *DemoVerifierSession) SetZKPQueriesStorage(newZKPQueriesStorage_ common.Address) (*types.Transaction, error) {
	return _DemoVerifier.Contract.SetZKPQueriesStorage(&_DemoVerifier.TransactOpts, newZKPQueriesStorage_)
}

// SetZKPQueriesStorage is a paid mutator transaction binding the contract method 0xddebe5c0.
//
// Solidity: function setZKPQueriesStorage(address newZKPQueriesStorage_) returns()
func (_DemoVerifier *DemoVerifierTransactorSession) SetZKPQueriesStorage(newZKPQueriesStorage_ common.Address) (*types.Transaction, error) {
	return _DemoVerifier.Contract.SetZKPQueriesStorage(&_DemoVerifier.TransactOpts, newZKPQueriesStorage_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DemoVerifier *DemoVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DemoVerifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DemoVerifier *DemoVerifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DemoVerifier.Contract.TransferOwnership(&_DemoVerifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DemoVerifier *DemoVerifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DemoVerifier.Contract.TransferOwnership(&_DemoVerifier.TransactOpts, newOwner)
}

// DemoVerifierIdentityProvedIterator is returned from FilterIdentityProved and is used to iterate over the raw logs and unpacked data for IdentityProved events raised by the DemoVerifier contract.
type DemoVerifierIdentityProvedIterator struct {
	Event *DemoVerifierIdentityProved // Event containing the contract specifics and raw log

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
func (it *DemoVerifierIdentityProvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DemoVerifierIdentityProved)
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
		it.Event = new(DemoVerifierIdentityProved)
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
func (it *DemoVerifierIdentityProvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DemoVerifierIdentityProvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DemoVerifierIdentityProved represents a IdentityProved event raised by the DemoVerifier contract.
type DemoVerifierIdentityProved struct {
	IdentityId *big.Int
	SenderAddr common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterIdentityProved is a free log retrieval operation binding the contract event 0x71e684bd20a8d6234a44a6af5b056eae6faa465c43327b30f9867d6ec0e327e6.
//
// Solidity: event IdentityProved(uint256 indexed identityId, address senderAddr)
func (_DemoVerifier *DemoVerifierFilterer) FilterIdentityProved(opts *bind.FilterOpts, identityId []*big.Int) (*DemoVerifierIdentityProvedIterator, error) {

	var identityIdRule []interface{}
	for _, identityIdItem := range identityId {
		identityIdRule = append(identityIdRule, identityIdItem)
	}

	logs, sub, err := _DemoVerifier.contract.FilterLogs(opts, "IdentityProved", identityIdRule)
	if err != nil {
		return nil, err
	}
	return &DemoVerifierIdentityProvedIterator{contract: _DemoVerifier.contract, event: "IdentityProved", logs: logs, sub: sub}, nil
}

// WatchIdentityProved is a free log subscription operation binding the contract event 0x71e684bd20a8d6234a44a6af5b056eae6faa465c43327b30f9867d6ec0e327e6.
//
// Solidity: event IdentityProved(uint256 indexed identityId, address senderAddr)
func (_DemoVerifier *DemoVerifierFilterer) WatchIdentityProved(opts *bind.WatchOpts, sink chan<- *DemoVerifierIdentityProved, identityId []*big.Int) (event.Subscription, error) {

	var identityIdRule []interface{}
	for _, identityIdItem := range identityId {
		identityIdRule = append(identityIdRule, identityIdItem)
	}

	logs, sub, err := _DemoVerifier.contract.WatchLogs(opts, "IdentityProved", identityIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DemoVerifierIdentityProved)
				if err := _DemoVerifier.contract.UnpackLog(event, "IdentityProved", log); err != nil {
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

// ParseIdentityProved is a log parse operation binding the contract event 0x71e684bd20a8d6234a44a6af5b056eae6faa465c43327b30f9867d6ec0e327e6.
//
// Solidity: event IdentityProved(uint256 indexed identityId, address senderAddr)
func (_DemoVerifier *DemoVerifierFilterer) ParseIdentityProved(log types.Log) (*DemoVerifierIdentityProved, error) {
	event := new(DemoVerifierIdentityProved)
	if err := _DemoVerifier.contract.UnpackLog(event, "IdentityProved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DemoVerifierInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DemoVerifier contract.
type DemoVerifierInitializedIterator struct {
	Event *DemoVerifierInitialized // Event containing the contract specifics and raw log

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
func (it *DemoVerifierInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DemoVerifierInitialized)
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
		it.Event = new(DemoVerifierInitialized)
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
func (it *DemoVerifierInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DemoVerifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DemoVerifierInitialized represents a Initialized event raised by the DemoVerifier contract.
type DemoVerifierInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DemoVerifier *DemoVerifierFilterer) FilterInitialized(opts *bind.FilterOpts) (*DemoVerifierInitializedIterator, error) {

	logs, sub, err := _DemoVerifier.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DemoVerifierInitializedIterator{contract: _DemoVerifier.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DemoVerifier *DemoVerifierFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DemoVerifierInitialized) (event.Subscription, error) {

	logs, sub, err := _DemoVerifier.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DemoVerifierInitialized)
				if err := _DemoVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DemoVerifier *DemoVerifierFilterer) ParseInitialized(log types.Log) (*DemoVerifierInitialized, error) {
	event := new(DemoVerifierInitialized)
	if err := _DemoVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DemoVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DemoVerifier contract.
type DemoVerifierOwnershipTransferredIterator struct {
	Event *DemoVerifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DemoVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DemoVerifierOwnershipTransferred)
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
		it.Event = new(DemoVerifierOwnershipTransferred)
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
func (it *DemoVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DemoVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DemoVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the DemoVerifier contract.
type DemoVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DemoVerifier *DemoVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DemoVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DemoVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DemoVerifierOwnershipTransferredIterator{contract: _DemoVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DemoVerifier *DemoVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DemoVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DemoVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DemoVerifierOwnershipTransferred)
				if err := _DemoVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DemoVerifier *DemoVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*DemoVerifierOwnershipTransferred, error) {
	event := new(DemoVerifierOwnershipTransferred)
	if err := _DemoVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

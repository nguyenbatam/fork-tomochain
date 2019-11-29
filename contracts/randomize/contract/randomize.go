// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"github.com/fns/fns/accounts/abi"
	"github.com/fns/fns/accounts/abi/bind"
	"github.com/fns/fns/common"
	"github.com/fns/fns/core/types"
	"strings"
)

// FNSRandomizeABI is the input ABI used to generate the binding from.
const FNSRandomizeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getSecret\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_secret\",\"type\":\"bytes32[]\"}],\"name\":\"setSecret\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getOpening\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_opening\",\"type\":\"bytes32\"}],\"name\":\"setOpening\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// FNSRandomizeBin is the compiled bytecode used for deploying new contracts.
const FNSRandomizeBin = `0x608060405234801561001057600080fd5b50610335806100206000396000f3006080604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663284180fc811461006657806334d38600146100e4578063d442d6cc1461013b578063e11f5ba21461017b575b600080fd5b34801561007257600080fd5b5061009473ffffffffffffffffffffffffffffffffffffffff60043516610193565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100d05781810151838201526020016100b8565b505050509050019250505060405180910390f35b3480156100f057600080fd5b50604080516020600480358082013583810280860185019096528085526101399536959394602494938501929182918501908490808284375094975061020b9650505050505050565b005b34801561014757600080fd5b5061016973ffffffffffffffffffffffffffffffffffffffff60043516610250565b60408051918252519081900360200190f35b34801561018757600080fd5b50610139600435610278565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260208181526040918290208054835181840281018401909452808452606093928301828280156101ff57602002820191906000526020600020905b815481526001909101906020018083116101ea575b50505050509050919050565b610384430661032081101561021f57600080fd5b610352811061022d57600080fd5b33600090815260208181526040909120835161024b9285019061029f565b505050565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b610384430661035281101561028c57600080fd5b5033600090815260016020526040902055565b8280548282559060005260206000209081019282156102dc579160200282015b828111156102dc57825182556020909201916001909101906102bf565b506102e89291506102ec565b5090565b61030691905b808211156102e857600081556001016102f2565b905600a165627a7a723058201f52ebc2713e1ca7eae309c3df69ec4b869733ffe5f987561ac8c2d5b5ee49a50029`

// DeployFNSRandomize deploys a new Ethereum contract, binding an instance of FNSRandomize to it.
func DeployFNSRandomize(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FNSRandomize, error) {
	parsed, err := abi.JSON(strings.NewReader(FNSRandomizeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FNSRandomizeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FNSRandomize{FNSRandomizeCaller: FNSRandomizeCaller{contract: contract}, FNSRandomizeTransactor: FNSRandomizeTransactor{contract: contract}, FNSRandomizeFilterer: FNSRandomizeFilterer{contract: contract}}, nil
}

// FNSRandomize is an auto generated Go binding around an Ethereum contract.
type FNSRandomize struct {
	FNSRandomizeCaller     // Read-only binding to the contract
	FNSRandomizeTransactor // Write-only binding to the contract
	FNSRandomizeFilterer   // Log filterer for contract events
}

// FNSRandomizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type FNSRandomizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FNSRandomizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FNSRandomizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FNSRandomizeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FNSRandomizeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FNSRandomizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FNSRandomizeSession struct {
	Contract     *FNSRandomize     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FNSRandomizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FNSRandomizeCallerSession struct {
	Contract *FNSRandomizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FNSRandomizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FNSRandomizeTransactorSession struct {
	Contract     *FNSRandomizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FNSRandomizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type FNSRandomizeRaw struct {
	Contract *FNSRandomize // Generic contract binding to access the raw methods on
}

// FNSRandomizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FNSRandomizeCallerRaw struct {
	Contract *FNSRandomizeCaller // Generic read-only contract binding to access the raw methods on
}

// FNSRandomizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FNSRandomizeTransactorRaw struct {
	Contract *FNSRandomizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFNSRandomize creates a new instance of FNSRandomize, bound to a specific deployed contract.
func NewFNSRandomize(address common.Address, backend bind.ContractBackend) (*FNSRandomize, error) {
	contract, err := bindFNSRandomize(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FNSRandomize{FNSRandomizeCaller: FNSRandomizeCaller{contract: contract}, FNSRandomizeTransactor: FNSRandomizeTransactor{contract: contract}, FNSRandomizeFilterer: FNSRandomizeFilterer{contract: contract}}, nil
}

// NewFNSRandomizeCaller creates a new read-only instance of FNSRandomize, bound to a specific deployed contract.
func NewFNSRandomizeCaller(address common.Address, caller bind.ContractCaller) (*FNSRandomizeCaller, error) {
	contract, err := bindFNSRandomize(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FNSRandomizeCaller{contract: contract}, nil
}

// NewFNSRandomizeTransactor creates a new write-only instance of FNSRandomize, bound to a specific deployed contract.
func NewFNSRandomizeTransactor(address common.Address, transactor bind.ContractTransactor) (*FNSRandomizeTransactor, error) {
	contract, err := bindFNSRandomize(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FNSRandomizeTransactor{contract: contract}, nil
}

// NewFNSRandomizeFilterer creates a new log filterer instance of FNSRandomize, bound to a specific deployed contract.
func NewFNSRandomizeFilterer(address common.Address, filterer bind.ContractFilterer) (*FNSRandomizeFilterer, error) {
	contract, err := bindFNSRandomize(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FNSRandomizeFilterer{contract: contract}, nil
}

// bindFNSRandomize binds a generic wrapper to an already deployed contract.
func bindFNSRandomize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FNSRandomizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FNSRandomize *FNSRandomizeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FNSRandomize.Contract.FNSRandomizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FNSRandomize *FNSRandomizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FNSRandomize.Contract.FNSRandomizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FNSRandomize *FNSRandomizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FNSRandomize.Contract.FNSRandomizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FNSRandomize *FNSRandomizeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FNSRandomize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FNSRandomize *FNSRandomizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FNSRandomize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FNSRandomize *FNSRandomizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FNSRandomize.Contract.contract.Transact(opts, method, params...)
}

// GetOpening is a free data retrieval call binding the contract method 0xd442d6cc.
//
// Solidity: function getOpening(_validator address) constant returns(bytes32)
func (_FNSRandomize *FNSRandomizeCaller) GetOpening(opts *bind.CallOpts, _validator common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _FNSRandomize.contract.Call(opts, out, "getOpening", _validator)
	return *ret0, err
}

// GetOpening is a free data retrieval call binding the contract method 0xd442d6cc.
//
// Solidity: function getOpening(_validator address) constant returns(bytes32)
func (_FNSRandomize *FNSRandomizeSession) GetOpening(_validator common.Address) ([32]byte, error) {
	return _FNSRandomize.Contract.GetOpening(&_FNSRandomize.CallOpts, _validator)
}

// GetOpening is a free data retrieval call binding the contract method 0xd442d6cc.
//
// Solidity: function getOpening(_validator address) constant returns(bytes32)
func (_FNSRandomize *FNSRandomizeCallerSession) GetOpening(_validator common.Address) ([32]byte, error) {
	return _FNSRandomize.Contract.GetOpening(&_FNSRandomize.CallOpts, _validator)
}

// GetSecret is a free data retrieval call binding the contract method 0x284180fc.
//
// Solidity: function getSecret(_validator address) constant returns(bytes32[])
func (_FNSRandomize *FNSRandomizeCaller) GetSecret(opts *bind.CallOpts, _validator common.Address) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _FNSRandomize.contract.Call(opts, out, "getSecret", _validator)
	return *ret0, err
}

// GetSecret is a free data retrieval call binding the contract method 0x284180fc.
//
// Solidity: function getSecret(_validator address) constant returns(bytes32[])
func (_FNSRandomize *FNSRandomizeSession) GetSecret(_validator common.Address) ([][32]byte, error) {
	return _FNSRandomize.Contract.GetSecret(&_FNSRandomize.CallOpts, _validator)
}

// GetSecret is a free data retrieval call binding the contract method 0x284180fc.
//
// Solidity: function getSecret(_validator address) constant returns(bytes32[])
func (_FNSRandomize *FNSRandomizeCallerSession) GetSecret(_validator common.Address) ([][32]byte, error) {
	return _FNSRandomize.Contract.GetSecret(&_FNSRandomize.CallOpts, _validator)
}

// SetOpening is a paid mutator transaction binding the contract method 0xe11f5ba2.
//
// Solidity: function setOpening(_opening bytes32) returns()
func (_FNSRandomize *FNSRandomizeTransactor) SetOpening(opts *bind.TransactOpts, _opening [32]byte) (*types.Transaction, error) {
	return _FNSRandomize.contract.Transact(opts, "setOpening", _opening)
}

// SetOpening is a paid mutator transaction binding the contract method 0xe11f5ba2.
//
// Solidity: function setOpening(_opening bytes32) returns()
func (_FNSRandomize *FNSRandomizeSession) SetOpening(_opening [32]byte) (*types.Transaction, error) {
	return _FNSRandomize.Contract.SetOpening(&_FNSRandomize.TransactOpts, _opening)
}

// SetOpening is a paid mutator transaction binding the contract method 0xe11f5ba2.
//
// Solidity: function setOpening(_opening bytes32) returns()
func (_FNSRandomize *FNSRandomizeTransactorSession) SetOpening(_opening [32]byte) (*types.Transaction, error) {
	return _FNSRandomize.Contract.SetOpening(&_FNSRandomize.TransactOpts, _opening)
}

// SetSecret is a paid mutator transaction binding the contract method 0x34d38600.
//
// Solidity: function setSecret(_secret bytes32[]) returns()
func (_FNSRandomize *FNSRandomizeTransactor) SetSecret(opts *bind.TransactOpts, _secret [][32]byte) (*types.Transaction, error) {
	return _FNSRandomize.contract.Transact(opts, "setSecret", _secret)
}

// SetSecret is a paid mutator transaction binding the contract method 0x34d38600.
//
// Solidity: function setSecret(_secret bytes32[]) returns()
func (_FNSRandomize *FNSRandomizeSession) SetSecret(_secret [][32]byte) (*types.Transaction, error) {
	return _FNSRandomize.Contract.SetSecret(&_FNSRandomize.TransactOpts, _secret)
}

// SetSecret is a paid mutator transaction binding the contract method 0x34d38600.
//
// Solidity: function setSecret(_secret bytes32[]) returns()
func (_FNSRandomize *FNSRandomizeTransactorSession) SetSecret(_secret [][32]byte) (*types.Transaction, error) {
	return _FNSRandomize.Contract.SetSecret(&_FNSRandomize.TransactOpts, _secret)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820c17235a98d4b09da2543b6484368e90c5c1b3bf8a8226a7be1eb6802372842100029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

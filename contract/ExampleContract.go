// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ExampleContractABI is the input ABI used to generate the binding from.
const ExampleContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_number\",\"type\":\"uint32\"}],\"name\":\"setNumber\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ExampleContractBin is the compiled bytecode used for deploying new contracts.
const ExampleContractBin = `0x608060405234801561001057600080fd5b5060f48061001f6000396000f3fe6080604052600436106042577c0100000000000000000000000000000000000000000000000000000000600035046325202bac811460475780638381f58a146075575b600080fd5b348015605257600080fd5b50607360048036036020811015606757600080fd5b503563ffffffff1660a0565b005b348015608057600080fd5b50608760bc565b6040805163ffffffff9092168252519081900360200190f35b6000805463ffffffff191663ffffffff92909216919091179055565b60005463ffffffff168156fea165627a7a723058204e10b1e35382594b30a2c20b20fa6a149769a1b7ca3bd60ca76602f462186f9a0029`

// DeployExampleContract deploys a new Ethereum contract, binding an instance of ExampleContract to it.
func DeployExampleContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExampleContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExampleContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExampleContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExampleContract{ExampleContractCaller: ExampleContractCaller{contract: contract}, ExampleContractTransactor: ExampleContractTransactor{contract: contract}, ExampleContractFilterer: ExampleContractFilterer{contract: contract}}, nil
}

// ExampleContract is an auto generated Go binding around an Ethereum contract.
type ExampleContract struct {
	ExampleContractCaller     // Read-only binding to the contract
	ExampleContractTransactor // Write-only binding to the contract
	ExampleContractFilterer   // Log filterer for contract events
}

// ExampleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleContractSession struct {
	Contract     *ExampleContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExampleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleContractCallerSession struct {
	Contract *ExampleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ExampleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleContractTransactorSession struct {
	Contract     *ExampleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ExampleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleContractRaw struct {
	Contract *ExampleContract // Generic contract binding to access the raw methods on
}

// ExampleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleContractCallerRaw struct {
	Contract *ExampleContractCaller // Generic read-only contract binding to access the raw methods on
}

// ExampleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleContractTransactorRaw struct {
	Contract *ExampleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleContract creates a new instance of ExampleContract, bound to a specific deployed contract.
func NewExampleContract(address common.Address, backend bind.ContractBackend) (*ExampleContract, error) {
	contract, err := bindExampleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleContract{ExampleContractCaller: ExampleContractCaller{contract: contract}, ExampleContractTransactor: ExampleContractTransactor{contract: contract}, ExampleContractFilterer: ExampleContractFilterer{contract: contract}}, nil
}

// NewExampleContractCaller creates a new read-only instance of ExampleContract, bound to a specific deployed contract.
func NewExampleContractCaller(address common.Address, caller bind.ContractCaller) (*ExampleContractCaller, error) {
	contract, err := bindExampleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleContractCaller{contract: contract}, nil
}

// NewExampleContractTransactor creates a new write-only instance of ExampleContract, bound to a specific deployed contract.
func NewExampleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ExampleContractTransactor, error) {
	contract, err := bindExampleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleContractTransactor{contract: contract}, nil
}

// NewExampleContractFilterer creates a new log filterer instance of ExampleContract, bound to a specific deployed contract.
func NewExampleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ExampleContractFilterer, error) {
	contract, err := bindExampleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleContractFilterer{contract: contract}, nil
}

// bindExampleContract binds a generic wrapper to an already deployed contract.
func bindExampleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExampleContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleContract *ExampleContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExampleContract.Contract.ExampleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleContract *ExampleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleContract.Contract.ExampleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleContract *ExampleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleContract.Contract.ExampleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleContract *ExampleContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExampleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleContract *ExampleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleContract *ExampleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleContract.Contract.contract.Transact(opts, method, params...)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() constant returns(uint32)
func (_ExampleContract *ExampleContractCaller) Number(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _ExampleContract.contract.Call(opts, out, "number")
	return *ret0, err
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() constant returns(uint32)
func (_ExampleContract *ExampleContractSession) Number() (uint32, error) {
	return _ExampleContract.Contract.Number(&_ExampleContract.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() constant returns(uint32)
func (_ExampleContract *ExampleContractCallerSession) Number() (uint32, error) {
	return _ExampleContract.Contract.Number(&_ExampleContract.CallOpts)
}

// SetNumber is a paid mutator transaction binding the contract method 0x25202bac.
//
// Solidity: function setNumber(_number uint32) returns()
func (_ExampleContract *ExampleContractTransactor) SetNumber(opts *bind.TransactOpts, _number uint32) (*types.Transaction, error) {
	return _ExampleContract.contract.Transact(opts, "setNumber", _number)
}

// SetNumber is a paid mutator transaction binding the contract method 0x25202bac.
//
// Solidity: function setNumber(_number uint32) returns()
func (_ExampleContract *ExampleContractSession) SetNumber(_number uint32) (*types.Transaction, error) {
	return _ExampleContract.Contract.SetNumber(&_ExampleContract.TransactOpts, _number)
}

// SetNumber is a paid mutator transaction binding the contract method 0x25202bac.
//
// Solidity: function setNumber(_number uint32) returns()
func (_ExampleContract *ExampleContractTransactorSession) SetNumber(_number uint32) (*types.Transaction, error) {
	return _ExampleContract.Contract.SetNumber(&_ExampleContract.TransactOpts, _number)
}

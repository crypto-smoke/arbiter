// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package utils

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

// ThingMetaData contains all meta data concerning the Thing contract.
var ThingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"truePriceTokenA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"truePriceTokenB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"computeProfitMaximizingTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"aToB\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ThingABI is the input ABI used to generate the binding from.
// Deprecated: Use ThingMetaData.ABI instead.
var ThingABI = ThingMetaData.ABI

// Thing is an auto generated Go binding around an Ethereum contract.
type Thing struct {
	ThingCaller     // Read-only binding to the contract
	ThingTransactor // Write-only binding to the contract
	ThingFilterer   // Log filterer for contract events
}

// ThingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ThingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ThingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ThingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ThingSession struct {
	Contract     *Thing            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ThingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ThingCallerSession struct {
	Contract *ThingCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ThingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ThingTransactorSession struct {
	Contract     *ThingTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ThingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ThingRaw struct {
	Contract *Thing // Generic contract binding to access the raw methods on
}

// ThingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ThingCallerRaw struct {
	Contract *ThingCaller // Generic read-only contract binding to access the raw methods on
}

// ThingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ThingTransactorRaw struct {
	Contract *ThingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewThing creates a new instance of Thing, bound to a specific deployed contract.
func NewThing(address common.Address, backend bind.ContractBackend) (*Thing, error) {
	contract, err := bindThing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Thing{ThingCaller: ThingCaller{contract: contract}, ThingTransactor: ThingTransactor{contract: contract}, ThingFilterer: ThingFilterer{contract: contract}}, nil
}

// NewThingCaller creates a new read-only instance of Thing, bound to a specific deployed contract.
func NewThingCaller(address common.Address, caller bind.ContractCaller) (*ThingCaller, error) {
	contract, err := bindThing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ThingCaller{contract: contract}, nil
}

// NewThingTransactor creates a new write-only instance of Thing, bound to a specific deployed contract.
func NewThingTransactor(address common.Address, transactor bind.ContractTransactor) (*ThingTransactor, error) {
	contract, err := bindThing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ThingTransactor{contract: contract}, nil
}

// NewThingFilterer creates a new log filterer instance of Thing, bound to a specific deployed contract.
func NewThingFilterer(address common.Address, filterer bind.ContractFilterer) (*ThingFilterer, error) {
	contract, err := bindThing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ThingFilterer{contract: contract}, nil
}

// bindThing binds a generic wrapper to an already deployed contract.
func bindThing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ThingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Thing *ThingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Thing.Contract.ThingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Thing *ThingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thing.Contract.ThingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Thing *ThingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Thing.Contract.ThingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Thing *ThingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Thing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Thing *ThingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Thing *ThingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Thing.Contract.contract.Transact(opts, method, params...)
}

// ComputeProfitMaximizingTrade is a free data retrieval call binding the contract method 0xfa653154.
//
// Solidity: function computeProfitMaximizingTrade(uint256 truePriceTokenA, uint256 truePriceTokenB, uint256 reserveA, uint256 reserveB) view returns(bool aToB, uint256 amountIn)
func (_Thing *ThingCaller) ComputeProfitMaximizingTrade(opts *bind.CallOpts, truePriceTokenA *big.Int, truePriceTokenB *big.Int, reserveA *big.Int, reserveB *big.Int) (struct {
	AToB     bool
	AmountIn *big.Int
}, error) {
	var out []interface{}
	err := _Thing.contract.Call(opts, &out, "computeProfitMaximizingTrade", truePriceTokenA, truePriceTokenB, reserveA, reserveB)

	outstruct := new(struct {
		AToB     bool
		AmountIn *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AToB = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.AmountIn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ComputeProfitMaximizingTrade is a free data retrieval call binding the contract method 0xfa653154.
//
// Solidity: function computeProfitMaximizingTrade(uint256 truePriceTokenA, uint256 truePriceTokenB, uint256 reserveA, uint256 reserveB) view returns(bool aToB, uint256 amountIn)
func (_Thing *ThingSession) ComputeProfitMaximizingTrade(truePriceTokenA *big.Int, truePriceTokenB *big.Int, reserveA *big.Int, reserveB *big.Int) (struct {
	AToB     bool
	AmountIn *big.Int
}, error) {
	return _Thing.Contract.ComputeProfitMaximizingTrade(&_Thing.CallOpts, truePriceTokenA, truePriceTokenB, reserveA, reserveB)
}

// ComputeProfitMaximizingTrade is a free data retrieval call binding the contract method 0xfa653154.
//
// Solidity: function computeProfitMaximizingTrade(uint256 truePriceTokenA, uint256 truePriceTokenB, uint256 reserveA, uint256 reserveB) view returns(bool aToB, uint256 amountIn)
func (_Thing *ThingCallerSession) ComputeProfitMaximizingTrade(truePriceTokenA *big.Int, truePriceTokenB *big.Int, reserveA *big.Int, reserveB *big.Int) (struct {
	AToB     bool
	AmountIn *big.Int
}, error) {
	return _Thing.Contract.ComputeProfitMaximizingTrade(&_Thing.CallOpts, truePriceTokenA, truePriceTokenB, reserveA, reserveB)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arb

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

// ArbABI is the input ABI used to generate the binding from.
const ArbABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"dualDexTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"estimateDualDexTrade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token3\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"estimateSingleDexTrade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router3\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token3\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"estimateTriDexTrade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getAmountOutMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContractAddress\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"recoverTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Arb is an auto generated Go binding around an Ethereum contract.
type Arb struct {
	ArbCaller     // Read-only binding to the contract
	ArbTransactor // Write-only binding to the contract
	ArbFilterer   // Log filterer for contract events
}

// ArbCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbSession struct {
	Contract     *Arb              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbCallerSession struct {
	Contract *ArbCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbTransactorSession struct {
	Contract     *ArbTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbRaw struct {
	Contract *Arb // Generic contract binding to access the raw methods on
}

// ArbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbCallerRaw struct {
	Contract *ArbCaller // Generic read-only contract binding to access the raw methods on
}

// ArbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbTransactorRaw struct {
	Contract *ArbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArb creates a new instance of Arb, bound to a specific deployed contract.
func NewArb(address common.Address, backend bind.ContractBackend) (*Arb, error) {
	contract, err := bindArb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arb{ArbCaller: ArbCaller{contract: contract}, ArbTransactor: ArbTransactor{contract: contract}, ArbFilterer: ArbFilterer{contract: contract}}, nil
}

// NewArbCaller creates a new read-only instance of Arb, bound to a specific deployed contract.
func NewArbCaller(address common.Address, caller bind.ContractCaller) (*ArbCaller, error) {
	contract, err := bindArb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbCaller{contract: contract}, nil
}

// NewArbTransactor creates a new write-only instance of Arb, bound to a specific deployed contract.
func NewArbTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbTransactor, error) {
	contract, err := bindArb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbTransactor{contract: contract}, nil
}

// NewArbFilterer creates a new log filterer instance of Arb, bound to a specific deployed contract.
func NewArbFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbFilterer, error) {
	contract, err := bindArb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbFilterer{contract: contract}, nil
}

// bindArb binds a generic wrapper to an already deployed contract.
func bindArb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arb *ArbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arb.Contract.ArbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arb *ArbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.Contract.ArbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arb *ArbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arb.Contract.ArbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arb *ArbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arb *ArbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arb *ArbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arb.Contract.contract.Transact(opts, method, params...)
}

// EstimateDualDexTrade is a free data retrieval call binding the contract method 0x068e7ca1.
//
// Solidity: function estimateDualDexTrade(address _router1, address _router2, address _token1, address _token2, uint256 _amount) view returns(uint256)
func (_Arb *ArbCaller) EstimateDualDexTrade(opts *bind.CallOpts, _router1 common.Address, _router2 common.Address, _token1 common.Address, _token2 common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "estimateDualDexTrade", _router1, _router2, _token1, _token2, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateDualDexTrade is a free data retrieval call binding the contract method 0x068e7ca1.
//
// Solidity: function estimateDualDexTrade(address _router1, address _router2, address _token1, address _token2, uint256 _amount) view returns(uint256)
func (_Arb *ArbSession) EstimateDualDexTrade(_router1 common.Address, _router2 common.Address, _token1 common.Address, _token2 common.Address, _amount *big.Int) (*big.Int, error) {
	return _Arb.Contract.EstimateDualDexTrade(&_Arb.CallOpts, _router1, _router2, _token1, _token2, _amount)
}

// EstimateDualDexTrade is a free data retrieval call binding the contract method 0x068e7ca1.
//
// Solidity: function estimateDualDexTrade(address _router1, address _router2, address _token1, address _token2, uint256 _amount) view returns(uint256)
func (_Arb *ArbCallerSession) EstimateDualDexTrade(_router1 common.Address, _router2 common.Address, _token1 common.Address, _token2 common.Address, _amount *big.Int) (*big.Int, error) {
	return _Arb.Contract.EstimateDualDexTrade(&_Arb.CallOpts, _router1, _router2, _token1, _token2, _amount)
}

// EstimateSingleDexTrade is a free data retrieval call binding the contract method 0xd8273ffe.
//
// Solidity: function estimateSingleDexTrade(address _router1, address _token1, address _token2, address _token3, uint256 _amount) view returns(uint256)
func (_Arb *ArbCaller) EstimateSingleDexTrade(opts *bind.CallOpts, _router1 common.Address, _token1 common.Address, _token2 common.Address, _token3 common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "estimateSingleDexTrade", _router1, _token1, _token2, _token3, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateSingleDexTrade is a free data retrieval call binding the contract method 0xd8273ffe.
//
// Solidity: function estimateSingleDexTrade(address _router1, address _token1, address _token2, address _token3, uint256 _amount) view returns(uint256)
func (_Arb *ArbSession) EstimateSingleDexTrade(_router1 common.Address, _token1 common.Address, _token2 common.Address, _token3 common.Address, _amount *big.Int) (*big.Int, error) {
	return _Arb.Contract.EstimateSingleDexTrade(&_Arb.CallOpts, _router1, _token1, _token2, _token3, _amount)
}

// EstimateSingleDexTrade is a free data retrieval call binding the contract method 0xd8273ffe.
//
// Solidity: function estimateSingleDexTrade(address _router1, address _token1, address _token2, address _token3, uint256 _amount) view returns(uint256)
func (_Arb *ArbCallerSession) EstimateSingleDexTrade(_router1 common.Address, _token1 common.Address, _token2 common.Address, _token3 common.Address, _amount *big.Int) (*big.Int, error) {
	return _Arb.Contract.EstimateSingleDexTrade(&_Arb.CallOpts, _router1, _token1, _token2, _token3, _amount)
}

// EstimateTriDexTrade is a free data retrieval call binding the contract method 0x26a215ec.
//
// Solidity: function estimateTriDexTrade(address _router1, address _router2, address _router3, address _token1, address _token2, address _token3, uint256 _amount) view returns(uint256)
func (_Arb *ArbCaller) EstimateTriDexTrade(opts *bind.CallOpts, _router1 common.Address, _router2 common.Address, _router3 common.Address, _token1 common.Address, _token2 common.Address, _token3 common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "estimateTriDexTrade", _router1, _router2, _router3, _token1, _token2, _token3, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTriDexTrade is a free data retrieval call binding the contract method 0x26a215ec.
//
// Solidity: function estimateTriDexTrade(address _router1, address _router2, address _router3, address _token1, address _token2, address _token3, uint256 _amount) view returns(uint256)
func (_Arb *ArbSession) EstimateTriDexTrade(_router1 common.Address, _router2 common.Address, _router3 common.Address, _token1 common.Address, _token2 common.Address, _token3 common.Address, _amount *big.Int) (*big.Int, error) {
	return _Arb.Contract.EstimateTriDexTrade(&_Arb.CallOpts, _router1, _router2, _router3, _token1, _token2, _token3, _amount)
}

// EstimateTriDexTrade is a free data retrieval call binding the contract method 0x26a215ec.
//
// Solidity: function estimateTriDexTrade(address _router1, address _router2, address _router3, address _token1, address _token2, address _token3, uint256 _amount) view returns(uint256)
func (_Arb *ArbCallerSession) EstimateTriDexTrade(_router1 common.Address, _router2 common.Address, _router3 common.Address, _token1 common.Address, _token2 common.Address, _token3 common.Address, _amount *big.Int) (*big.Int, error) {
	return _Arb.Contract.EstimateTriDexTrade(&_Arb.CallOpts, _router1, _router2, _router3, _token1, _token2, _token3, _amount)
}

// GetAmountOutMin is a free data retrieval call binding the contract method 0x43cf6f24.
//
// Solidity: function getAmountOutMin(address router, address _tokenIn, address _tokenOut, uint256 _amount) view returns(uint256)
func (_Arb *ArbCaller) GetAmountOutMin(opts *bind.CallOpts, router common.Address, _tokenIn common.Address, _tokenOut common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "getAmountOutMin", router, _tokenIn, _tokenOut, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOutMin is a free data retrieval call binding the contract method 0x43cf6f24.
//
// Solidity: function getAmountOutMin(address router, address _tokenIn, address _tokenOut, uint256 _amount) view returns(uint256)
func (_Arb *ArbSession) GetAmountOutMin(router common.Address, _tokenIn common.Address, _tokenOut common.Address, _amount *big.Int) (*big.Int, error) {
	return _Arb.Contract.GetAmountOutMin(&_Arb.CallOpts, router, _tokenIn, _tokenOut, _amount)
}

// GetAmountOutMin is a free data retrieval call binding the contract method 0x43cf6f24.
//
// Solidity: function getAmountOutMin(address router, address _tokenIn, address _tokenOut, uint256 _amount) view returns(uint256)
func (_Arb *ArbCallerSession) GetAmountOutMin(router common.Address, _tokenIn common.Address, _tokenOut common.Address, _amount *big.Int) (*big.Int, error) {
	return _Arb.Contract.GetAmountOutMin(&_Arb.CallOpts, router, _tokenIn, _tokenOut, _amount)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _tokenContractAddress) view returns(uint256)
func (_Arb *ArbCaller) GetBalance(opts *bind.CallOpts, _tokenContractAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "getBalance", _tokenContractAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _tokenContractAddress) view returns(uint256)
func (_Arb *ArbSession) GetBalance(_tokenContractAddress common.Address) (*big.Int, error) {
	return _Arb.Contract.GetBalance(&_Arb.CallOpts, _tokenContractAddress)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _tokenContractAddress) view returns(uint256)
func (_Arb *ArbCallerSession) GetBalance(_tokenContractAddress common.Address) (*big.Int, error) {
	return _Arb.Contract.GetBalance(&_Arb.CallOpts, _tokenContractAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arb *ArbCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arb *ArbSession) Owner() (common.Address, error) {
	return _Arb.Contract.Owner(&_Arb.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arb *ArbCallerSession) Owner() (common.Address, error) {
	return _Arb.Contract.Owner(&_Arb.CallOpts)
}

// DualDexTrade is a paid mutator transaction binding the contract method 0x1aa51318.
//
// Solidity: function dualDexTrade(address _router1, address _router2, address _token1, address _token2, uint256 _amount) returns()
func (_Arb *ArbTransactor) DualDexTrade(opts *bind.TransactOpts, _router1 common.Address, _router2 common.Address, _token1 common.Address, _token2 common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "dualDexTrade", _router1, _router2, _token1, _token2, _amount)
}

// DualDexTrade is a paid mutator transaction binding the contract method 0x1aa51318.
//
// Solidity: function dualDexTrade(address _router1, address _router2, address _token1, address _token2, uint256 _amount) returns()
func (_Arb *ArbSession) DualDexTrade(_router1 common.Address, _router2 common.Address, _token1 common.Address, _token2 common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.DualDexTrade(&_Arb.TransactOpts, _router1, _router2, _token1, _token2, _amount)
}

// DualDexTrade is a paid mutator transaction binding the contract method 0x1aa51318.
//
// Solidity: function dualDexTrade(address _router1, address _router2, address _token1, address _token2, uint256 _amount) returns()
func (_Arb *ArbTransactorSession) DualDexTrade(_router1 common.Address, _router2 common.Address, _token1 common.Address, _token2 common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.DualDexTrade(&_Arb.TransactOpts, _router1, _router2, _token1, _token2, _amount)
}

// RecoverEth is a paid mutator transaction binding the contract method 0xbcdb446b.
//
// Solidity: function recoverEth() returns()
func (_Arb *ArbTransactor) RecoverEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "recoverEth")
}

// RecoverEth is a paid mutator transaction binding the contract method 0xbcdb446b.
//
// Solidity: function recoverEth() returns()
func (_Arb *ArbSession) RecoverEth() (*types.Transaction, error) {
	return _Arb.Contract.RecoverEth(&_Arb.TransactOpts)
}

// RecoverEth is a paid mutator transaction binding the contract method 0xbcdb446b.
//
// Solidity: function recoverEth() returns()
func (_Arb *ArbTransactorSession) RecoverEth() (*types.Transaction, error) {
	return _Arb.Contract.RecoverEth(&_Arb.TransactOpts)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(address tokenAddress) returns()
func (_Arb *ArbTransactor) RecoverTokens(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "recoverTokens", tokenAddress)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(address tokenAddress) returns()
func (_Arb *ArbSession) RecoverTokens(tokenAddress common.Address) (*types.Transaction, error) {
	return _Arb.Contract.RecoverTokens(&_Arb.TransactOpts, tokenAddress)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(address tokenAddress) returns()
func (_Arb *ArbTransactorSession) RecoverTokens(tokenAddress common.Address) (*types.Transaction, error) {
	return _Arb.Contract.RecoverTokens(&_Arb.TransactOpts, tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arb *ArbTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arb *ArbSession) RenounceOwnership() (*types.Transaction, error) {
	return _Arb.Contract.RenounceOwnership(&_Arb.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arb *ArbTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Arb.Contract.RenounceOwnership(&_Arb.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arb *ArbTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arb *ArbSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Arb.Contract.TransferOwnership(&_Arb.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arb *ArbTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Arb.Contract.TransferOwnership(&_Arb.TransactOpts, newOwner)
}

// ArbOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Arb contract.
type ArbOwnershipTransferredIterator struct {
	Event *ArbOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ArbOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbOwnershipTransferred)
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
		it.Event = new(ArbOwnershipTransferred)
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
func (it *ArbOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbOwnershipTransferred represents a OwnershipTransferred event raised by the Arb contract.
type ArbOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Arb *ArbFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ArbOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArbOwnershipTransferredIterator{contract: _Arb.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Arb *ArbFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbOwnershipTransferred)
				if err := _Arb.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Arb *ArbFilterer) ParseOwnershipTransferred(log types.Log) (*ArbOwnershipTransferred, error) {
	event := new(ArbOwnershipTransferred)
	if err := _Arb.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

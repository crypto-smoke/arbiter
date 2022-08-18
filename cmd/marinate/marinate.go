// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// MarinateMetaData contains all meta data concerning the Marinate contract.
var MarinateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_UMAMI\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_depositLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"AddToContractWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"RemoveFromContractWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rps\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardCollection\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multipliedAmount\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multipliedAmount\",\"type\":\"uint256\"}],\"name\":\"StakeMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multipliedAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawMultiplier\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SCALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UMAMI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"name\":\"addApprovedMultiplierToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addApprovedRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"addToContractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getAvailableTokenRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedMultiplierNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedRewardToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isNFTStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marinatorInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multipliedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"migrateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"multiplierNFTs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierStakingEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierWithdrawEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nftMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paidTokenRewardsPerStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payRewardsEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeApprovedMultiplierToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeApprovedRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"removeFromContractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setDepositLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setMultiplierStakeEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setMultiplierWithdrawEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setPayRewardsEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_scale\",\"type\":\"uint256\"}],\"name\":\"setScale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setStakeEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setStakingWithdrawEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setTransferEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"stakeMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"toBePaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMultipliedStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalTokenRewardsPerStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedContracts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarinateABI is the input ABI used to generate the binding from.
// Deprecated: Use MarinateMetaData.ABI instead.
var MarinateABI = MarinateMetaData.ABI

// Marinate is an auto generated Go binding around an Ethereum contract.
type Marinate struct {
	MarinateCaller     // Read-only binding to the contract
	MarinateTransactor // Write-only binding to the contract
	MarinateFilterer   // Log filterer for contract events
}

// MarinateCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarinateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarinateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarinateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarinateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarinateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarinateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarinateSession struct {
	Contract     *Marinate         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarinateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarinateCallerSession struct {
	Contract *MarinateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MarinateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarinateTransactorSession struct {
	Contract     *MarinateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MarinateRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarinateRaw struct {
	Contract *Marinate // Generic contract binding to access the raw methods on
}

// MarinateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarinateCallerRaw struct {
	Contract *MarinateCaller // Generic read-only contract binding to access the raw methods on
}

// MarinateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarinateTransactorRaw struct {
	Contract *MarinateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarinate creates a new instance of Marinate, bound to a specific deployed contract.
func NewMarinate(address common.Address, backend bind.ContractBackend) (*Marinate, error) {
	contract, err := bindMarinate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marinate{MarinateCaller: MarinateCaller{contract: contract}, MarinateTransactor: MarinateTransactor{contract: contract}, MarinateFilterer: MarinateFilterer{contract: contract}}, nil
}

// NewMarinateCaller creates a new read-only instance of Marinate, bound to a specific deployed contract.
func NewMarinateCaller(address common.Address, caller bind.ContractCaller) (*MarinateCaller, error) {
	contract, err := bindMarinate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarinateCaller{contract: contract}, nil
}

// NewMarinateTransactor creates a new write-only instance of Marinate, bound to a specific deployed contract.
func NewMarinateTransactor(address common.Address, transactor bind.ContractTransactor) (*MarinateTransactor, error) {
	contract, err := bindMarinate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarinateTransactor{contract: contract}, nil
}

// NewMarinateFilterer creates a new log filterer instance of Marinate, bound to a specific deployed contract.
func NewMarinateFilterer(address common.Address, filterer bind.ContractFilterer) (*MarinateFilterer, error) {
	contract, err := bindMarinate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarinateFilterer{contract: contract}, nil
}

// bindMarinate binds a generic wrapper to an already deployed contract.
func bindMarinate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarinateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marinate *MarinateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marinate.Contract.MarinateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marinate *MarinateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marinate.Contract.MarinateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marinate *MarinateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marinate.Contract.MarinateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marinate *MarinateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marinate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marinate *MarinateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marinate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marinate *MarinateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marinate.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Marinate *MarinateCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Marinate *MarinateSession) ADMINROLE() ([32]byte, error) {
	return _Marinate.Contract.ADMINROLE(&_Marinate.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Marinate *MarinateCallerSession) ADMINROLE() ([32]byte, error) {
	return _Marinate.Contract.ADMINROLE(&_Marinate.CallOpts)
}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(uint256)
func (_Marinate *MarinateCaller) BASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(uint256)
func (_Marinate *MarinateSession) BASE() (*big.Int, error) {
	return _Marinate.Contract.BASE(&_Marinate.CallOpts)
}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(uint256)
func (_Marinate *MarinateCallerSession) BASE() (*big.Int, error) {
	return _Marinate.Contract.BASE(&_Marinate.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Marinate *MarinateCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Marinate *MarinateSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Marinate.Contract.DEFAULTADMINROLE(&_Marinate.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Marinate *MarinateCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Marinate.Contract.DEFAULTADMINROLE(&_Marinate.CallOpts)
}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint256)
func (_Marinate *MarinateCaller) SCALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "SCALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint256)
func (_Marinate *MarinateSession) SCALE() (*big.Int, error) {
	return _Marinate.Contract.SCALE(&_Marinate.CallOpts)
}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint256)
func (_Marinate *MarinateCallerSession) SCALE() (*big.Int, error) {
	return _Marinate.Contract.SCALE(&_Marinate.CallOpts)
}

// UMAMI is a free data retrieval call binding the contract method 0x1eee01b0.
//
// Solidity: function UMAMI() view returns(address)
func (_Marinate *MarinateCaller) UMAMI(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "UMAMI")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UMAMI is a free data retrieval call binding the contract method 0x1eee01b0.
//
// Solidity: function UMAMI() view returns(address)
func (_Marinate *MarinateSession) UMAMI() (common.Address, error) {
	return _Marinate.Contract.UMAMI(&_Marinate.CallOpts)
}

// UMAMI is a free data retrieval call binding the contract method 0x1eee01b0.
//
// Solidity: function UMAMI() view returns(address)
func (_Marinate *MarinateCallerSession) UMAMI() (common.Address, error) {
	return _Marinate.Contract.UMAMI(&_Marinate.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Marinate *MarinateCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Marinate *MarinateSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Marinate.Contract.Allowance(&_Marinate.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Marinate *MarinateCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Marinate.Contract.Allowance(&_Marinate.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Marinate *MarinateCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Marinate *MarinateSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Marinate.Contract.BalanceOf(&_Marinate.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Marinate *MarinateCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Marinate.Contract.BalanceOf(&_Marinate.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Marinate *MarinateCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Marinate *MarinateSession) Decimals() (uint8, error) {
	return _Marinate.Contract.Decimals(&_Marinate.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Marinate *MarinateCallerSession) Decimals() (uint8, error) {
	return _Marinate.Contract.Decimals(&_Marinate.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Marinate *MarinateCaller) DepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "depositLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Marinate *MarinateSession) DepositLimit() (*big.Int, error) {
	return _Marinate.Contract.DepositLimit(&_Marinate.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_Marinate *MarinateCallerSession) DepositLimit() (*big.Int, error) {
	return _Marinate.Contract.DepositLimit(&_Marinate.CallOpts)
}

// GetAvailableTokenRewards is a free data retrieval call binding the contract method 0xc5644c86.
//
// Solidity: function getAvailableTokenRewards(address staker, address token) view returns(uint256 totalRewards)
func (_Marinate *MarinateCaller) GetAvailableTokenRewards(opts *bind.CallOpts, staker common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "getAvailableTokenRewards", staker, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableTokenRewards is a free data retrieval call binding the contract method 0xc5644c86.
//
// Solidity: function getAvailableTokenRewards(address staker, address token) view returns(uint256 totalRewards)
func (_Marinate *MarinateSession) GetAvailableTokenRewards(staker common.Address, token common.Address) (*big.Int, error) {
	return _Marinate.Contract.GetAvailableTokenRewards(&_Marinate.CallOpts, staker, token)
}

// GetAvailableTokenRewards is a free data retrieval call binding the contract method 0xc5644c86.
//
// Solidity: function getAvailableTokenRewards(address staker, address token) view returns(uint256 totalRewards)
func (_Marinate *MarinateCallerSession) GetAvailableTokenRewards(staker common.Address, token common.Address) (*big.Int, error) {
	return _Marinate.Contract.GetAvailableTokenRewards(&_Marinate.CallOpts, staker, token)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Marinate *MarinateCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Marinate *MarinateSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Marinate.Contract.GetRoleAdmin(&_Marinate.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Marinate *MarinateCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Marinate.Contract.GetRoleAdmin(&_Marinate.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Marinate *MarinateCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Marinate *MarinateSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Marinate.Contract.HasRole(&_Marinate.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Marinate *MarinateCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Marinate.Contract.HasRole(&_Marinate.CallOpts, role, account)
}

// IsApprovedMultiplierNFT is a free data retrieval call binding the contract method 0x8c6fd36b.
//
// Solidity: function isApprovedMultiplierNFT(address ) view returns(bool)
func (_Marinate *MarinateCaller) IsApprovedMultiplierNFT(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "isApprovedMultiplierNFT", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedMultiplierNFT is a free data retrieval call binding the contract method 0x8c6fd36b.
//
// Solidity: function isApprovedMultiplierNFT(address ) view returns(bool)
func (_Marinate *MarinateSession) IsApprovedMultiplierNFT(arg0 common.Address) (bool, error) {
	return _Marinate.Contract.IsApprovedMultiplierNFT(&_Marinate.CallOpts, arg0)
}

// IsApprovedMultiplierNFT is a free data retrieval call binding the contract method 0x8c6fd36b.
//
// Solidity: function isApprovedMultiplierNFT(address ) view returns(bool)
func (_Marinate *MarinateCallerSession) IsApprovedMultiplierNFT(arg0 common.Address) (bool, error) {
	return _Marinate.Contract.IsApprovedMultiplierNFT(&_Marinate.CallOpts, arg0)
}

// IsApprovedRewardToken is a free data retrieval call binding the contract method 0x7095bffc.
//
// Solidity: function isApprovedRewardToken(address ) view returns(bool)
func (_Marinate *MarinateCaller) IsApprovedRewardToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "isApprovedRewardToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedRewardToken is a free data retrieval call binding the contract method 0x7095bffc.
//
// Solidity: function isApprovedRewardToken(address ) view returns(bool)
func (_Marinate *MarinateSession) IsApprovedRewardToken(arg0 common.Address) (bool, error) {
	return _Marinate.Contract.IsApprovedRewardToken(&_Marinate.CallOpts, arg0)
}

// IsApprovedRewardToken is a free data retrieval call binding the contract method 0x7095bffc.
//
// Solidity: function isApprovedRewardToken(address ) view returns(bool)
func (_Marinate *MarinateCallerSession) IsApprovedRewardToken(arg0 common.Address) (bool, error) {
	return _Marinate.Contract.IsApprovedRewardToken(&_Marinate.CallOpts, arg0)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) view returns(bool)
func (_Marinate *MarinateCaller) IsContract(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "isContract", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) view returns(bool)
func (_Marinate *MarinateSession) IsContract(addr common.Address) (bool, error) {
	return _Marinate.Contract.IsContract(&_Marinate.CallOpts, addr)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address addr) view returns(bool)
func (_Marinate *MarinateCallerSession) IsContract(addr common.Address) (bool, error) {
	return _Marinate.Contract.IsContract(&_Marinate.CallOpts, addr)
}

// IsNFTStaked is a free data retrieval call binding the contract method 0x5526b05f.
//
// Solidity: function isNFTStaked(address , address ) view returns(bool)
func (_Marinate *MarinateCaller) IsNFTStaked(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "isNFTStaked", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNFTStaked is a free data retrieval call binding the contract method 0x5526b05f.
//
// Solidity: function isNFTStaked(address , address ) view returns(bool)
func (_Marinate *MarinateSession) IsNFTStaked(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Marinate.Contract.IsNFTStaked(&_Marinate.CallOpts, arg0, arg1)
}

// IsNFTStaked is a free data retrieval call binding the contract method 0x5526b05f.
//
// Solidity: function isNFTStaked(address , address ) view returns(bool)
func (_Marinate *MarinateCallerSession) IsNFTStaked(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Marinate.Contract.IsNFTStaked(&_Marinate.CallOpts, arg0, arg1)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address addr) view returns(bool)
func (_Marinate *MarinateCaller) IsWhitelisted(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "isWhitelisted", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address addr) view returns(bool)
func (_Marinate *MarinateSession) IsWhitelisted(addr common.Address) (bool, error) {
	return _Marinate.Contract.IsWhitelisted(&_Marinate.CallOpts, addr)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address addr) view returns(bool)
func (_Marinate *MarinateCallerSession) IsWhitelisted(addr common.Address) (bool, error) {
	return _Marinate.Contract.IsWhitelisted(&_Marinate.CallOpts, addr)
}

// MarinatorInfo is a free data retrieval call binding the contract method 0x9b5f61a4.
//
// Solidity: function marinatorInfo(address ) view returns(uint256 amount, uint256 multipliedAmount)
func (_Marinate *MarinateCaller) MarinatorInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount           *big.Int
	MultipliedAmount *big.Int
}, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "marinatorInfo", arg0)

	outstruct := new(struct {
		Amount           *big.Int
		MultipliedAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MultipliedAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MarinatorInfo is a free data retrieval call binding the contract method 0x9b5f61a4.
//
// Solidity: function marinatorInfo(address ) view returns(uint256 amount, uint256 multipliedAmount)
func (_Marinate *MarinateSession) MarinatorInfo(arg0 common.Address) (struct {
	Amount           *big.Int
	MultipliedAmount *big.Int
}, error) {
	return _Marinate.Contract.MarinatorInfo(&_Marinate.CallOpts, arg0)
}

// MarinatorInfo is a free data retrieval call binding the contract method 0x9b5f61a4.
//
// Solidity: function marinatorInfo(address ) view returns(uint256 amount, uint256 multipliedAmount)
func (_Marinate *MarinateCallerSession) MarinatorInfo(arg0 common.Address) (struct {
	Amount           *big.Int
	MultipliedAmount *big.Int
}, error) {
	return _Marinate.Contract.MarinatorInfo(&_Marinate.CallOpts, arg0)
}

// MultiplierNFTs is a free data retrieval call binding the contract method 0x55d84c1c.
//
// Solidity: function multiplierNFTs(uint256 ) view returns(address)
func (_Marinate *MarinateCaller) MultiplierNFTs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "multiplierNFTs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MultiplierNFTs is a free data retrieval call binding the contract method 0x55d84c1c.
//
// Solidity: function multiplierNFTs(uint256 ) view returns(address)
func (_Marinate *MarinateSession) MultiplierNFTs(arg0 *big.Int) (common.Address, error) {
	return _Marinate.Contract.MultiplierNFTs(&_Marinate.CallOpts, arg0)
}

// MultiplierNFTs is a free data retrieval call binding the contract method 0x55d84c1c.
//
// Solidity: function multiplierNFTs(uint256 ) view returns(address)
func (_Marinate *MarinateCallerSession) MultiplierNFTs(arg0 *big.Int) (common.Address, error) {
	return _Marinate.Contract.MultiplierNFTs(&_Marinate.CallOpts, arg0)
}

// MultiplierStakingEnabled is a free data retrieval call binding the contract method 0x700d27cf.
//
// Solidity: function multiplierStakingEnabled() view returns(bool)
func (_Marinate *MarinateCaller) MultiplierStakingEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "multiplierStakingEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MultiplierStakingEnabled is a free data retrieval call binding the contract method 0x700d27cf.
//
// Solidity: function multiplierStakingEnabled() view returns(bool)
func (_Marinate *MarinateSession) MultiplierStakingEnabled() (bool, error) {
	return _Marinate.Contract.MultiplierStakingEnabled(&_Marinate.CallOpts)
}

// MultiplierStakingEnabled is a free data retrieval call binding the contract method 0x700d27cf.
//
// Solidity: function multiplierStakingEnabled() view returns(bool)
func (_Marinate *MarinateCallerSession) MultiplierStakingEnabled() (bool, error) {
	return _Marinate.Contract.MultiplierStakingEnabled(&_Marinate.CallOpts)
}

// MultiplierWithdrawEnabled is a free data retrieval call binding the contract method 0x311d8a51.
//
// Solidity: function multiplierWithdrawEnabled() view returns(bool)
func (_Marinate *MarinateCaller) MultiplierWithdrawEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "multiplierWithdrawEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MultiplierWithdrawEnabled is a free data retrieval call binding the contract method 0x311d8a51.
//
// Solidity: function multiplierWithdrawEnabled() view returns(bool)
func (_Marinate *MarinateSession) MultiplierWithdrawEnabled() (bool, error) {
	return _Marinate.Contract.MultiplierWithdrawEnabled(&_Marinate.CallOpts)
}

// MultiplierWithdrawEnabled is a free data retrieval call binding the contract method 0x311d8a51.
//
// Solidity: function multiplierWithdrawEnabled() view returns(bool)
func (_Marinate *MarinateCallerSession) MultiplierWithdrawEnabled() (bool, error) {
	return _Marinate.Contract.MultiplierWithdrawEnabled(&_Marinate.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Marinate *MarinateCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Marinate *MarinateSession) Name() (string, error) {
	return _Marinate.Contract.Name(&_Marinate.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Marinate *MarinateCallerSession) Name() (string, error) {
	return _Marinate.Contract.Name(&_Marinate.CallOpts)
}

// NftMultiplier is a free data retrieval call binding the contract method 0x0b09e5fb.
//
// Solidity: function nftMultiplier(address ) view returns(uint256)
func (_Marinate *MarinateCaller) NftMultiplier(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "nftMultiplier", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftMultiplier is a free data retrieval call binding the contract method 0x0b09e5fb.
//
// Solidity: function nftMultiplier(address ) view returns(uint256)
func (_Marinate *MarinateSession) NftMultiplier(arg0 common.Address) (*big.Int, error) {
	return _Marinate.Contract.NftMultiplier(&_Marinate.CallOpts, arg0)
}

// NftMultiplier is a free data retrieval call binding the contract method 0x0b09e5fb.
//
// Solidity: function nftMultiplier(address ) view returns(uint256)
func (_Marinate *MarinateCallerSession) NftMultiplier(arg0 common.Address) (*big.Int, error) {
	return _Marinate.Contract.NftMultiplier(&_Marinate.CallOpts, arg0)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Marinate *MarinateCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Marinate *MarinateSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Marinate.Contract.OnERC721Received(&_Marinate.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Marinate *MarinateCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Marinate.Contract.OnERC721Received(&_Marinate.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marinate *MarinateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marinate *MarinateSession) Owner() (common.Address, error) {
	return _Marinate.Contract.Owner(&_Marinate.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marinate *MarinateCallerSession) Owner() (common.Address, error) {
	return _Marinate.Contract.Owner(&_Marinate.CallOpts)
}

// PaidTokenRewardsPerStake is a free data retrieval call binding the contract method 0xbfe0bc98.
//
// Solidity: function paidTokenRewardsPerStake(address , address ) view returns(uint256)
func (_Marinate *MarinateCaller) PaidTokenRewardsPerStake(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "paidTokenRewardsPerStake", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PaidTokenRewardsPerStake is a free data retrieval call binding the contract method 0xbfe0bc98.
//
// Solidity: function paidTokenRewardsPerStake(address , address ) view returns(uint256)
func (_Marinate *MarinateSession) PaidTokenRewardsPerStake(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Marinate.Contract.PaidTokenRewardsPerStake(&_Marinate.CallOpts, arg0, arg1)
}

// PaidTokenRewardsPerStake is a free data retrieval call binding the contract method 0xbfe0bc98.
//
// Solidity: function paidTokenRewardsPerStake(address , address ) view returns(uint256)
func (_Marinate *MarinateCallerSession) PaidTokenRewardsPerStake(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Marinate.Contract.PaidTokenRewardsPerStake(&_Marinate.CallOpts, arg0, arg1)
}

// PayRewardsEnabled is a free data retrieval call binding the contract method 0xdb0e9d66.
//
// Solidity: function payRewardsEnabled() view returns(bool)
func (_Marinate *MarinateCaller) PayRewardsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "payRewardsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PayRewardsEnabled is a free data retrieval call binding the contract method 0xdb0e9d66.
//
// Solidity: function payRewardsEnabled() view returns(bool)
func (_Marinate *MarinateSession) PayRewardsEnabled() (bool, error) {
	return _Marinate.Contract.PayRewardsEnabled(&_Marinate.CallOpts)
}

// PayRewardsEnabled is a free data retrieval call binding the contract method 0xdb0e9d66.
//
// Solidity: function payRewardsEnabled() view returns(bool)
func (_Marinate *MarinateCallerSession) PayRewardsEnabled() (bool, error) {
	return _Marinate.Contract.PayRewardsEnabled(&_Marinate.CallOpts)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_Marinate *MarinateCaller) RewardTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "rewardTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_Marinate *MarinateSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _Marinate.Contract.RewardTokens(&_Marinate.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_Marinate *MarinateCallerSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _Marinate.Contract.RewardTokens(&_Marinate.CallOpts, arg0)
}

// StakeEnabled is a free data retrieval call binding the contract method 0x3b17c736.
//
// Solidity: function stakeEnabled() view returns(bool)
func (_Marinate *MarinateCaller) StakeEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "stakeEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StakeEnabled is a free data retrieval call binding the contract method 0x3b17c736.
//
// Solidity: function stakeEnabled() view returns(bool)
func (_Marinate *MarinateSession) StakeEnabled() (bool, error) {
	return _Marinate.Contract.StakeEnabled(&_Marinate.CallOpts)
}

// StakeEnabled is a free data retrieval call binding the contract method 0x3b17c736.
//
// Solidity: function stakeEnabled() view returns(bool)
func (_Marinate *MarinateCallerSession) StakeEnabled() (bool, error) {
	return _Marinate.Contract.StakeEnabled(&_Marinate.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Marinate *MarinateCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Marinate *MarinateSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Marinate.Contract.SupportsInterface(&_Marinate.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Marinate *MarinateCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Marinate.Contract.SupportsInterface(&_Marinate.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Marinate *MarinateCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Marinate *MarinateSession) Symbol() (string, error) {
	return _Marinate.Contract.Symbol(&_Marinate.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Marinate *MarinateCallerSession) Symbol() (string, error) {
	return _Marinate.Contract.Symbol(&_Marinate.CallOpts)
}

// ToBePaid is a free data retrieval call binding the contract method 0xbc35e3fa.
//
// Solidity: function toBePaid(address , address ) view returns(uint256)
func (_Marinate *MarinateCaller) ToBePaid(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "toBePaid", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToBePaid is a free data retrieval call binding the contract method 0xbc35e3fa.
//
// Solidity: function toBePaid(address , address ) view returns(uint256)
func (_Marinate *MarinateSession) ToBePaid(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Marinate.Contract.ToBePaid(&_Marinate.CallOpts, arg0, arg1)
}

// ToBePaid is a free data retrieval call binding the contract method 0xbc35e3fa.
//
// Solidity: function toBePaid(address , address ) view returns(uint256)
func (_Marinate *MarinateCallerSession) ToBePaid(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Marinate.Contract.ToBePaid(&_Marinate.CallOpts, arg0, arg1)
}

// TotalMultipliedStaked is a free data retrieval call binding the contract method 0x5f46c8cc.
//
// Solidity: function totalMultipliedStaked() view returns(uint256)
func (_Marinate *MarinateCaller) TotalMultipliedStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "totalMultipliedStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMultipliedStaked is a free data retrieval call binding the contract method 0x5f46c8cc.
//
// Solidity: function totalMultipliedStaked() view returns(uint256)
func (_Marinate *MarinateSession) TotalMultipliedStaked() (*big.Int, error) {
	return _Marinate.Contract.TotalMultipliedStaked(&_Marinate.CallOpts)
}

// TotalMultipliedStaked is a free data retrieval call binding the contract method 0x5f46c8cc.
//
// Solidity: function totalMultipliedStaked() view returns(uint256)
func (_Marinate *MarinateCallerSession) TotalMultipliedStaked() (*big.Int, error) {
	return _Marinate.Contract.TotalMultipliedStaked(&_Marinate.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Marinate *MarinateCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Marinate *MarinateSession) TotalStaked() (*big.Int, error) {
	return _Marinate.Contract.TotalStaked(&_Marinate.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Marinate *MarinateCallerSession) TotalStaked() (*big.Int, error) {
	return _Marinate.Contract.TotalStaked(&_Marinate.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Marinate *MarinateCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Marinate *MarinateSession) TotalSupply() (*big.Int, error) {
	return _Marinate.Contract.TotalSupply(&_Marinate.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Marinate *MarinateCallerSession) TotalSupply() (*big.Int, error) {
	return _Marinate.Contract.TotalSupply(&_Marinate.CallOpts)
}

// TotalTokenRewardsPerStake is a free data retrieval call binding the contract method 0x01b6c969.
//
// Solidity: function totalTokenRewardsPerStake(address ) view returns(uint256)
func (_Marinate *MarinateCaller) TotalTokenRewardsPerStake(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "totalTokenRewardsPerStake", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokenRewardsPerStake is a free data retrieval call binding the contract method 0x01b6c969.
//
// Solidity: function totalTokenRewardsPerStake(address ) view returns(uint256)
func (_Marinate *MarinateSession) TotalTokenRewardsPerStake(arg0 common.Address) (*big.Int, error) {
	return _Marinate.Contract.TotalTokenRewardsPerStake(&_Marinate.CallOpts, arg0)
}

// TotalTokenRewardsPerStake is a free data retrieval call binding the contract method 0x01b6c969.
//
// Solidity: function totalTokenRewardsPerStake(address ) view returns(uint256)
func (_Marinate *MarinateCallerSession) TotalTokenRewardsPerStake(arg0 common.Address) (*big.Int, error) {
	return _Marinate.Contract.TotalTokenRewardsPerStake(&_Marinate.CallOpts, arg0)
}

// TransferEnabled is a free data retrieval call binding the contract method 0x4cd412d5.
//
// Solidity: function transferEnabled() view returns(bool)
func (_Marinate *MarinateCaller) TransferEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "transferEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferEnabled is a free data retrieval call binding the contract method 0x4cd412d5.
//
// Solidity: function transferEnabled() view returns(bool)
func (_Marinate *MarinateSession) TransferEnabled() (bool, error) {
	return _Marinate.Contract.TransferEnabled(&_Marinate.CallOpts)
}

// TransferEnabled is a free data retrieval call binding the contract method 0x4cd412d5.
//
// Solidity: function transferEnabled() view returns(bool)
func (_Marinate *MarinateCallerSession) TransferEnabled() (bool, error) {
	return _Marinate.Contract.TransferEnabled(&_Marinate.CallOpts)
}

// WhitelistedContracts is a free data retrieval call binding the contract method 0x391feebb.
//
// Solidity: function whitelistedContracts(address ) view returns(bool)
func (_Marinate *MarinateCaller) WhitelistedContracts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "whitelistedContracts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedContracts is a free data retrieval call binding the contract method 0x391feebb.
//
// Solidity: function whitelistedContracts(address ) view returns(bool)
func (_Marinate *MarinateSession) WhitelistedContracts(arg0 common.Address) (bool, error) {
	return _Marinate.Contract.WhitelistedContracts(&_Marinate.CallOpts, arg0)
}

// WhitelistedContracts is a free data retrieval call binding the contract method 0x391feebb.
//
// Solidity: function whitelistedContracts(address ) view returns(bool)
func (_Marinate *MarinateCallerSession) WhitelistedContracts(arg0 common.Address) (bool, error) {
	return _Marinate.Contract.WhitelistedContracts(&_Marinate.CallOpts, arg0)
}

// WithdrawEnabled is a free data retrieval call binding the contract method 0x2287e96a.
//
// Solidity: function withdrawEnabled() view returns(bool)
func (_Marinate *MarinateCaller) WithdrawEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Marinate.contract.Call(opts, &out, "withdrawEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawEnabled is a free data retrieval call binding the contract method 0x2287e96a.
//
// Solidity: function withdrawEnabled() view returns(bool)
func (_Marinate *MarinateSession) WithdrawEnabled() (bool, error) {
	return _Marinate.Contract.WithdrawEnabled(&_Marinate.CallOpts)
}

// WithdrawEnabled is a free data retrieval call binding the contract method 0x2287e96a.
//
// Solidity: function withdrawEnabled() view returns(bool)
func (_Marinate *MarinateCallerSession) WithdrawEnabled() (bool, error) {
	return _Marinate.Contract.WithdrawEnabled(&_Marinate.CallOpts)
}

// AddApprovedMultiplierToken is a paid mutator transaction binding the contract method 0x08a8c0e0.
//
// Solidity: function addApprovedMultiplierToken(address token, uint256 multiplier) returns()
func (_Marinate *MarinateTransactor) AddApprovedMultiplierToken(opts *bind.TransactOpts, token common.Address, multiplier *big.Int) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "addApprovedMultiplierToken", token, multiplier)
}

// AddApprovedMultiplierToken is a paid mutator transaction binding the contract method 0x08a8c0e0.
//
// Solidity: function addApprovedMultiplierToken(address token, uint256 multiplier) returns()
func (_Marinate *MarinateSession) AddApprovedMultiplierToken(token common.Address, multiplier *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.AddApprovedMultiplierToken(&_Marinate.TransactOpts, token, multiplier)
}

// AddApprovedMultiplierToken is a paid mutator transaction binding the contract method 0x08a8c0e0.
//
// Solidity: function addApprovedMultiplierToken(address token, uint256 multiplier) returns()
func (_Marinate *MarinateTransactorSession) AddApprovedMultiplierToken(token common.Address, multiplier *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.AddApprovedMultiplierToken(&_Marinate.TransactOpts, token, multiplier)
}

// AddApprovedRewardToken is a paid mutator transaction binding the contract method 0xca49ff65.
//
// Solidity: function addApprovedRewardToken(address token) returns()
func (_Marinate *MarinateTransactor) AddApprovedRewardToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "addApprovedRewardToken", token)
}

// AddApprovedRewardToken is a paid mutator transaction binding the contract method 0xca49ff65.
//
// Solidity: function addApprovedRewardToken(address token) returns()
func (_Marinate *MarinateSession) AddApprovedRewardToken(token common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.AddApprovedRewardToken(&_Marinate.TransactOpts, token)
}

// AddApprovedRewardToken is a paid mutator transaction binding the contract method 0xca49ff65.
//
// Solidity: function addApprovedRewardToken(address token) returns()
func (_Marinate *MarinateTransactorSession) AddApprovedRewardToken(token common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.AddApprovedRewardToken(&_Marinate.TransactOpts, token)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address token, uint256 amount) returns()
func (_Marinate *MarinateTransactor) AddReward(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "addReward", token, amount)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address token, uint256 amount) returns()
func (_Marinate *MarinateSession) AddReward(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.AddReward(&_Marinate.TransactOpts, token, amount)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address token, uint256 amount) returns()
func (_Marinate *MarinateTransactorSession) AddReward(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.AddReward(&_Marinate.TransactOpts, token, amount)
}

// AddToContractWhitelist is a paid mutator transaction binding the contract method 0xacc3a006.
//
// Solidity: function addToContractWhitelist(address _contract) returns(bool)
func (_Marinate *MarinateTransactor) AddToContractWhitelist(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "addToContractWhitelist", _contract)
}

// AddToContractWhitelist is a paid mutator transaction binding the contract method 0xacc3a006.
//
// Solidity: function addToContractWhitelist(address _contract) returns(bool)
func (_Marinate *MarinateSession) AddToContractWhitelist(_contract common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.AddToContractWhitelist(&_Marinate.TransactOpts, _contract)
}

// AddToContractWhitelist is a paid mutator transaction binding the contract method 0xacc3a006.
//
// Solidity: function addToContractWhitelist(address _contract) returns(bool)
func (_Marinate *MarinateTransactorSession) AddToContractWhitelist(_contract common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.AddToContractWhitelist(&_Marinate.TransactOpts, _contract)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Marinate *MarinateTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Marinate *MarinateSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.Approve(&_Marinate.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Marinate *MarinateTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.Approve(&_Marinate.TransactOpts, spender, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_Marinate *MarinateTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "claimRewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_Marinate *MarinateSession) ClaimRewards() (*types.Transaction, error) {
	return _Marinate.Contract.ClaimRewards(&_Marinate.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_Marinate *MarinateTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _Marinate.Contract.ClaimRewards(&_Marinate.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Marinate *MarinateTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Marinate *MarinateSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.DecreaseAllowance(&_Marinate.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Marinate *MarinateTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.DecreaseAllowance(&_Marinate.TransactOpts, spender, subtractedValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Marinate *MarinateTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Marinate *MarinateSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.GrantRole(&_Marinate.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Marinate *MarinateTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.GrantRole(&_Marinate.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Marinate *MarinateTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Marinate *MarinateSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.IncreaseAllowance(&_Marinate.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Marinate *MarinateTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.IncreaseAllowance(&_Marinate.TransactOpts, spender, addedValue)
}

// MigrateToken is a paid mutator transaction binding the contract method 0x93bfa282.
//
// Solidity: function migrateToken(address token, address destination, uint256 amount) returns()
func (_Marinate *MarinateTransactor) MigrateToken(opts *bind.TransactOpts, token common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "migrateToken", token, destination, amount)
}

// MigrateToken is a paid mutator transaction binding the contract method 0x93bfa282.
//
// Solidity: function migrateToken(address token, address destination, uint256 amount) returns()
func (_Marinate *MarinateSession) MigrateToken(token common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.MigrateToken(&_Marinate.TransactOpts, token, destination, amount)
}

// MigrateToken is a paid mutator transaction binding the contract method 0x93bfa282.
//
// Solidity: function migrateToken(address token, address destination, uint256 amount) returns()
func (_Marinate *MarinateTransactorSession) MigrateToken(token common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.MigrateToken(&_Marinate.TransactOpts, token, destination, amount)
}

// RecoverEth is a paid mutator transaction binding the contract method 0xbcdb446b.
//
// Solidity: function recoverEth() returns()
func (_Marinate *MarinateTransactor) RecoverEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "recoverEth")
}

// RecoverEth is a paid mutator transaction binding the contract method 0xbcdb446b.
//
// Solidity: function recoverEth() returns()
func (_Marinate *MarinateSession) RecoverEth() (*types.Transaction, error) {
	return _Marinate.Contract.RecoverEth(&_Marinate.TransactOpts)
}

// RecoverEth is a paid mutator transaction binding the contract method 0xbcdb446b.
//
// Solidity: function recoverEth() returns()
func (_Marinate *MarinateTransactorSession) RecoverEth() (*types.Transaction, error) {
	return _Marinate.Contract.RecoverEth(&_Marinate.TransactOpts)
}

// RemoveApprovedMultiplierToken is a paid mutator transaction binding the contract method 0x3abfbfef.
//
// Solidity: function removeApprovedMultiplierToken(address token) returns()
func (_Marinate *MarinateTransactor) RemoveApprovedMultiplierToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "removeApprovedMultiplierToken", token)
}

// RemoveApprovedMultiplierToken is a paid mutator transaction binding the contract method 0x3abfbfef.
//
// Solidity: function removeApprovedMultiplierToken(address token) returns()
func (_Marinate *MarinateSession) RemoveApprovedMultiplierToken(token common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.RemoveApprovedMultiplierToken(&_Marinate.TransactOpts, token)
}

// RemoveApprovedMultiplierToken is a paid mutator transaction binding the contract method 0x3abfbfef.
//
// Solidity: function removeApprovedMultiplierToken(address token) returns()
func (_Marinate *MarinateTransactorSession) RemoveApprovedMultiplierToken(token common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.RemoveApprovedMultiplierToken(&_Marinate.TransactOpts, token)
}

// RemoveApprovedRewardToken is a paid mutator transaction binding the contract method 0x6f5d3a5a.
//
// Solidity: function removeApprovedRewardToken(address token) returns()
func (_Marinate *MarinateTransactor) RemoveApprovedRewardToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "removeApprovedRewardToken", token)
}

// RemoveApprovedRewardToken is a paid mutator transaction binding the contract method 0x6f5d3a5a.
//
// Solidity: function removeApprovedRewardToken(address token) returns()
func (_Marinate *MarinateSession) RemoveApprovedRewardToken(token common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.RemoveApprovedRewardToken(&_Marinate.TransactOpts, token)
}

// RemoveApprovedRewardToken is a paid mutator transaction binding the contract method 0x6f5d3a5a.
//
// Solidity: function removeApprovedRewardToken(address token) returns()
func (_Marinate *MarinateTransactorSession) RemoveApprovedRewardToken(token common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.RemoveApprovedRewardToken(&_Marinate.TransactOpts, token)
}

// RemoveFromContractWhitelist is a paid mutator transaction binding the contract method 0xc3d9ed39.
//
// Solidity: function removeFromContractWhitelist(address _contract) returns(bool)
func (_Marinate *MarinateTransactor) RemoveFromContractWhitelist(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "removeFromContractWhitelist", _contract)
}

// RemoveFromContractWhitelist is a paid mutator transaction binding the contract method 0xc3d9ed39.
//
// Solidity: function removeFromContractWhitelist(address _contract) returns(bool)
func (_Marinate *MarinateSession) RemoveFromContractWhitelist(_contract common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.RemoveFromContractWhitelist(&_Marinate.TransactOpts, _contract)
}

// RemoveFromContractWhitelist is a paid mutator transaction binding the contract method 0xc3d9ed39.
//
// Solidity: function removeFromContractWhitelist(address _contract) returns(bool)
func (_Marinate *MarinateTransactorSession) RemoveFromContractWhitelist(_contract common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.RemoveFromContractWhitelist(&_Marinate.TransactOpts, _contract)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marinate *MarinateTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marinate *MarinateSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marinate.Contract.RenounceOwnership(&_Marinate.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marinate *MarinateTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marinate.Contract.RenounceOwnership(&_Marinate.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Marinate *MarinateTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Marinate *MarinateSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.RenounceRole(&_Marinate.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Marinate *MarinateTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.RenounceRole(&_Marinate.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Marinate *MarinateTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Marinate *MarinateSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.RevokeRole(&_Marinate.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Marinate *MarinateTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.RevokeRole(&_Marinate.TransactOpts, role, account)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_Marinate *MarinateTransactor) SetDepositLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "setDepositLimit", limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_Marinate *MarinateSession) SetDepositLimit(limit *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.SetDepositLimit(&_Marinate.TransactOpts, limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_Marinate *MarinateTransactorSession) SetDepositLimit(limit *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.SetDepositLimit(&_Marinate.TransactOpts, limit)
}

// SetMultiplierStakeEnabled is a paid mutator transaction binding the contract method 0x5d80dbe9.
//
// Solidity: function setMultiplierStakeEnabled(bool enabled) returns()
func (_Marinate *MarinateTransactor) SetMultiplierStakeEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "setMultiplierStakeEnabled", enabled)
}

// SetMultiplierStakeEnabled is a paid mutator transaction binding the contract method 0x5d80dbe9.
//
// Solidity: function setMultiplierStakeEnabled(bool enabled) returns()
func (_Marinate *MarinateSession) SetMultiplierStakeEnabled(enabled bool) (*types.Transaction, error) {
	return _Marinate.Contract.SetMultiplierStakeEnabled(&_Marinate.TransactOpts, enabled)
}

// SetMultiplierStakeEnabled is a paid mutator transaction binding the contract method 0x5d80dbe9.
//
// Solidity: function setMultiplierStakeEnabled(bool enabled) returns()
func (_Marinate *MarinateTransactorSession) SetMultiplierStakeEnabled(enabled bool) (*types.Transaction, error) {
	return _Marinate.Contract.SetMultiplierStakeEnabled(&_Marinate.TransactOpts, enabled)
}

// SetMultiplierWithdrawEnabled is a paid mutator transaction binding the contract method 0x8fa9484d.
//
// Solidity: function setMultiplierWithdrawEnabled(bool enabled) returns()
func (_Marinate *MarinateTransactor) SetMultiplierWithdrawEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "setMultiplierWithdrawEnabled", enabled)
}

// SetMultiplierWithdrawEnabled is a paid mutator transaction binding the contract method 0x8fa9484d.
//
// Solidity: function setMultiplierWithdrawEnabled(bool enabled) returns()
func (_Marinate *MarinateSession) SetMultiplierWithdrawEnabled(enabled bool) (*types.Transaction, error) {
	return _Marinate.Contract.SetMultiplierWithdrawEnabled(&_Marinate.TransactOpts, enabled)
}

// SetMultiplierWithdrawEnabled is a paid mutator transaction binding the contract method 0x8fa9484d.
//
// Solidity: function setMultiplierWithdrawEnabled(bool enabled) returns()
func (_Marinate *MarinateTransactorSession) SetMultiplierWithdrawEnabled(enabled bool) (*types.Transaction, error) {
	return _Marinate.Contract.SetMultiplierWithdrawEnabled(&_Marinate.TransactOpts, enabled)
}

// SetPayRewardsEnabled is a paid mutator transaction binding the contract method 0xbfd994ce.
//
// Solidity: function setPayRewardsEnabled(bool enabled) returns()
func (_Marinate *MarinateTransactor) SetPayRewardsEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "setPayRewardsEnabled", enabled)
}

// SetPayRewardsEnabled is a paid mutator transaction binding the contract method 0xbfd994ce.
//
// Solidity: function setPayRewardsEnabled(bool enabled) returns()
func (_Marinate *MarinateSession) SetPayRewardsEnabled(enabled bool) (*types.Transaction, error) {
	return _Marinate.Contract.SetPayRewardsEnabled(&_Marinate.TransactOpts, enabled)
}

// SetPayRewardsEnabled is a paid mutator transaction binding the contract method 0xbfd994ce.
//
// Solidity: function setPayRewardsEnabled(bool enabled) returns()
func (_Marinate *MarinateTransactorSession) SetPayRewardsEnabled(enabled bool) (*types.Transaction, error) {
	return _Marinate.Contract.SetPayRewardsEnabled(&_Marinate.TransactOpts, enabled)
}

// SetScale is a paid mutator transaction binding the contract method 0x3edc3519.
//
// Solidity: function setScale(uint256 _scale) returns()
func (_Marinate *MarinateTransactor) SetScale(opts *bind.TransactOpts, _scale *big.Int) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "setScale", _scale)
}

// SetScale is a paid mutator transaction binding the contract method 0x3edc3519.
//
// Solidity: function setScale(uint256 _scale) returns()
func (_Marinate *MarinateSession) SetScale(_scale *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.SetScale(&_Marinate.TransactOpts, _scale)
}

// SetScale is a paid mutator transaction binding the contract method 0x3edc3519.
//
// Solidity: function setScale(uint256 _scale) returns()
func (_Marinate *MarinateTransactorSession) SetScale(_scale *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.SetScale(&_Marinate.TransactOpts, _scale)
}

// SetStakeEnabled is a paid mutator transaction binding the contract method 0x82461948.
//
// Solidity: function setStakeEnabled(bool enabled) returns()
func (_Marinate *MarinateTransactor) SetStakeEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "setStakeEnabled", enabled)
}

// SetStakeEnabled is a paid mutator transaction binding the contract method 0x82461948.
//
// Solidity: function setStakeEnabled(bool enabled) returns()
func (_Marinate *MarinateSession) SetStakeEnabled(enabled bool) (*types.Transaction, error) {
	return _Marinate.Contract.SetStakeEnabled(&_Marinate.TransactOpts, enabled)
}

// SetStakeEnabled is a paid mutator transaction binding the contract method 0x82461948.
//
// Solidity: function setStakeEnabled(bool enabled) returns()
func (_Marinate *MarinateTransactorSession) SetStakeEnabled(enabled bool) (*types.Transaction, error) {
	return _Marinate.Contract.SetStakeEnabled(&_Marinate.TransactOpts, enabled)
}

// SetStakingWithdrawEnabled is a paid mutator transaction binding the contract method 0x6580221f.
//
// Solidity: function setStakingWithdrawEnabled(bool enabled) returns()
func (_Marinate *MarinateTransactor) SetStakingWithdrawEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "setStakingWithdrawEnabled", enabled)
}

// SetStakingWithdrawEnabled is a paid mutator transaction binding the contract method 0x6580221f.
//
// Solidity: function setStakingWithdrawEnabled(bool enabled) returns()
func (_Marinate *MarinateSession) SetStakingWithdrawEnabled(enabled bool) (*types.Transaction, error) {
	return _Marinate.Contract.SetStakingWithdrawEnabled(&_Marinate.TransactOpts, enabled)
}

// SetStakingWithdrawEnabled is a paid mutator transaction binding the contract method 0x6580221f.
//
// Solidity: function setStakingWithdrawEnabled(bool enabled) returns()
func (_Marinate *MarinateTransactorSession) SetStakingWithdrawEnabled(enabled bool) (*types.Transaction, error) {
	return _Marinate.Contract.SetStakingWithdrawEnabled(&_Marinate.TransactOpts, enabled)
}

// SetTransferEnabled is a paid mutator transaction binding the contract method 0x9fe9f623.
//
// Solidity: function setTransferEnabled(bool enabled) returns()
func (_Marinate *MarinateTransactor) SetTransferEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "setTransferEnabled", enabled)
}

// SetTransferEnabled is a paid mutator transaction binding the contract method 0x9fe9f623.
//
// Solidity: function setTransferEnabled(bool enabled) returns()
func (_Marinate *MarinateSession) SetTransferEnabled(enabled bool) (*types.Transaction, error) {
	return _Marinate.Contract.SetTransferEnabled(&_Marinate.TransactOpts, enabled)
}

// SetTransferEnabled is a paid mutator transaction binding the contract method 0x9fe9f623.
//
// Solidity: function setTransferEnabled(bool enabled) returns()
func (_Marinate *MarinateTransactorSession) SetTransferEnabled(enabled bool) (*types.Transaction, error) {
	return _Marinate.Contract.SetTransferEnabled(&_Marinate.TransactOpts, enabled)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Marinate *MarinateTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Marinate *MarinateSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.Stake(&_Marinate.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Marinate *MarinateTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.Stake(&_Marinate.TransactOpts, amount)
}

// StakeMultiplier is a paid mutator transaction binding the contract method 0xafb27a75.
//
// Solidity: function stakeMultiplier(address nft, uint256 tokenId) returns()
func (_Marinate *MarinateTransactor) StakeMultiplier(opts *bind.TransactOpts, nft common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "stakeMultiplier", nft, tokenId)
}

// StakeMultiplier is a paid mutator transaction binding the contract method 0xafb27a75.
//
// Solidity: function stakeMultiplier(address nft, uint256 tokenId) returns()
func (_Marinate *MarinateSession) StakeMultiplier(nft common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.StakeMultiplier(&_Marinate.TransactOpts, nft, tokenId)
}

// StakeMultiplier is a paid mutator transaction binding the contract method 0xafb27a75.
//
// Solidity: function stakeMultiplier(address nft, uint256 tokenId) returns()
func (_Marinate *MarinateTransactorSession) StakeMultiplier(nft common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.StakeMultiplier(&_Marinate.TransactOpts, nft, tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Marinate *MarinateTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Marinate *MarinateSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.Transfer(&_Marinate.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Marinate *MarinateTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.Transfer(&_Marinate.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Marinate *MarinateTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Marinate *MarinateSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.TransferFrom(&_Marinate.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Marinate *MarinateTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.TransferFrom(&_Marinate.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marinate *MarinateTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marinate *MarinateSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.TransferOwnership(&_Marinate.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marinate *MarinateTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marinate.Contract.TransferOwnership(&_Marinate.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Marinate *MarinateTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Marinate *MarinateSession) Withdraw() (*types.Transaction, error) {
	return _Marinate.Contract.Withdraw(&_Marinate.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Marinate *MarinateTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Marinate.Contract.Withdraw(&_Marinate.TransactOpts)
}

// WithdrawMultiplier is a paid mutator transaction binding the contract method 0x72c21597.
//
// Solidity: function withdrawMultiplier(address nft, uint256 tokenId) returns()
func (_Marinate *MarinateTransactor) WithdrawMultiplier(opts *bind.TransactOpts, nft common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Marinate.contract.Transact(opts, "withdrawMultiplier", nft, tokenId)
}

// WithdrawMultiplier is a paid mutator transaction binding the contract method 0x72c21597.
//
// Solidity: function withdrawMultiplier(address nft, uint256 tokenId) returns()
func (_Marinate *MarinateSession) WithdrawMultiplier(nft common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.WithdrawMultiplier(&_Marinate.TransactOpts, nft, tokenId)
}

// WithdrawMultiplier is a paid mutator transaction binding the contract method 0x72c21597.
//
// Solidity: function withdrawMultiplier(address nft, uint256 tokenId) returns()
func (_Marinate *MarinateTransactorSession) WithdrawMultiplier(nft common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Marinate.Contract.WithdrawMultiplier(&_Marinate.TransactOpts, nft, tokenId)
}

// MarinateAddToContractWhitelistIterator is returned from FilterAddToContractWhitelist and is used to iterate over the raw logs and unpacked data for AddToContractWhitelist events raised by the Marinate contract.
type MarinateAddToContractWhitelistIterator struct {
	Event *MarinateAddToContractWhitelist // Event containing the contract specifics and raw log

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
func (it *MarinateAddToContractWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateAddToContractWhitelist)
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
		it.Event = new(MarinateAddToContractWhitelist)
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
func (it *MarinateAddToContractWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateAddToContractWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateAddToContractWhitelist represents a AddToContractWhitelist event raised by the Marinate contract.
type MarinateAddToContractWhitelist struct {
	Contract common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddToContractWhitelist is a free log retrieval operation binding the contract event 0xfbd3cde7ff522a917e485c8ed2a6e87590887ab399f5ac312307903f49854307.
//
// Solidity: event AddToContractWhitelist(address indexed _contract)
func (_Marinate *MarinateFilterer) FilterAddToContractWhitelist(opts *bind.FilterOpts, _contract []common.Address) (*MarinateAddToContractWhitelistIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "AddToContractWhitelist", _contractRule)
	if err != nil {
		return nil, err
	}
	return &MarinateAddToContractWhitelistIterator{contract: _Marinate.contract, event: "AddToContractWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddToContractWhitelist is a free log subscription operation binding the contract event 0xfbd3cde7ff522a917e485c8ed2a6e87590887ab399f5ac312307903f49854307.
//
// Solidity: event AddToContractWhitelist(address indexed _contract)
func (_Marinate *MarinateFilterer) WatchAddToContractWhitelist(opts *bind.WatchOpts, sink chan<- *MarinateAddToContractWhitelist, _contract []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "AddToContractWhitelist", _contractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateAddToContractWhitelist)
				if err := _Marinate.contract.UnpackLog(event, "AddToContractWhitelist", log); err != nil {
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

// ParseAddToContractWhitelist is a log parse operation binding the contract event 0xfbd3cde7ff522a917e485c8ed2a6e87590887ab399f5ac312307903f49854307.
//
// Solidity: event AddToContractWhitelist(address indexed _contract)
func (_Marinate *MarinateFilterer) ParseAddToContractWhitelist(log types.Log) (*MarinateAddToContractWhitelist, error) {
	event := new(MarinateAddToContractWhitelist)
	if err := _Marinate.contract.UnpackLog(event, "AddToContractWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarinateApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Marinate contract.
type MarinateApprovalIterator struct {
	Event *MarinateApproval // Event containing the contract specifics and raw log

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
func (it *MarinateApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateApproval)
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
		it.Event = new(MarinateApproval)
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
func (it *MarinateApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateApproval represents a Approval event raised by the Marinate contract.
type MarinateApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Marinate *MarinateFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MarinateApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MarinateApprovalIterator{contract: _Marinate.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Marinate *MarinateFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MarinateApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateApproval)
				if err := _Marinate.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Marinate *MarinateFilterer) ParseApproval(log types.Log) (*MarinateApproval, error) {
	event := new(MarinateApproval)
	if err := _Marinate.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarinateOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Marinate contract.
type MarinateOwnershipTransferredIterator struct {
	Event *MarinateOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarinateOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateOwnershipTransferred)
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
		it.Event = new(MarinateOwnershipTransferred)
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
func (it *MarinateOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateOwnershipTransferred represents a OwnershipTransferred event raised by the Marinate contract.
type MarinateOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marinate *MarinateFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarinateOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarinateOwnershipTransferredIterator{contract: _Marinate.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marinate *MarinateFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarinateOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateOwnershipTransferred)
				if err := _Marinate.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Marinate *MarinateFilterer) ParseOwnershipTransferred(log types.Log) (*MarinateOwnershipTransferred, error) {
	event := new(MarinateOwnershipTransferred)
	if err := _Marinate.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarinateRemoveFromContractWhitelistIterator is returned from FilterRemoveFromContractWhitelist and is used to iterate over the raw logs and unpacked data for RemoveFromContractWhitelist events raised by the Marinate contract.
type MarinateRemoveFromContractWhitelistIterator struct {
	Event *MarinateRemoveFromContractWhitelist // Event containing the contract specifics and raw log

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
func (it *MarinateRemoveFromContractWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateRemoveFromContractWhitelist)
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
		it.Event = new(MarinateRemoveFromContractWhitelist)
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
func (it *MarinateRemoveFromContractWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateRemoveFromContractWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateRemoveFromContractWhitelist represents a RemoveFromContractWhitelist event raised by the Marinate contract.
type MarinateRemoveFromContractWhitelist struct {
	Contract common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemoveFromContractWhitelist is a free log retrieval operation binding the contract event 0x8e81447740597754af5db3e176253a36f7981a9549f48ace3f0cb233913f9d85.
//
// Solidity: event RemoveFromContractWhitelist(address indexed _contract)
func (_Marinate *MarinateFilterer) FilterRemoveFromContractWhitelist(opts *bind.FilterOpts, _contract []common.Address) (*MarinateRemoveFromContractWhitelistIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "RemoveFromContractWhitelist", _contractRule)
	if err != nil {
		return nil, err
	}
	return &MarinateRemoveFromContractWhitelistIterator{contract: _Marinate.contract, event: "RemoveFromContractWhitelist", logs: logs, sub: sub}, nil
}

// WatchRemoveFromContractWhitelist is a free log subscription operation binding the contract event 0x8e81447740597754af5db3e176253a36f7981a9549f48ace3f0cb233913f9d85.
//
// Solidity: event RemoveFromContractWhitelist(address indexed _contract)
func (_Marinate *MarinateFilterer) WatchRemoveFromContractWhitelist(opts *bind.WatchOpts, sink chan<- *MarinateRemoveFromContractWhitelist, _contract []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "RemoveFromContractWhitelist", _contractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateRemoveFromContractWhitelist)
				if err := _Marinate.contract.UnpackLog(event, "RemoveFromContractWhitelist", log); err != nil {
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

// ParseRemoveFromContractWhitelist is a log parse operation binding the contract event 0x8e81447740597754af5db3e176253a36f7981a9549f48ace3f0cb233913f9d85.
//
// Solidity: event RemoveFromContractWhitelist(address indexed _contract)
func (_Marinate *MarinateFilterer) ParseRemoveFromContractWhitelist(log types.Log) (*MarinateRemoveFromContractWhitelist, error) {
	event := new(MarinateRemoveFromContractWhitelist)
	if err := _Marinate.contract.UnpackLog(event, "RemoveFromContractWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarinateRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the Marinate contract.
type MarinateRewardAddedIterator struct {
	Event *MarinateRewardAdded // Event containing the contract specifics and raw log

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
func (it *MarinateRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateRewardAdded)
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
		it.Event = new(MarinateRewardAdded)
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
func (it *MarinateRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateRewardAdded represents a RewardAdded event raised by the Marinate contract.
type MarinateRewardAdded struct {
	Token  common.Address
	Amount *big.Int
	Rps    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0x6a6f77044107a33658235d41bedbbaf2fe9ccdceb313143c947a5e76e1ec8474.
//
// Solidity: event RewardAdded(address token, uint256 amount, uint256 rps)
func (_Marinate *MarinateFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*MarinateRewardAddedIterator, error) {

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &MarinateRewardAddedIterator{contract: _Marinate.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0x6a6f77044107a33658235d41bedbbaf2fe9ccdceb313143c947a5e76e1ec8474.
//
// Solidity: event RewardAdded(address token, uint256 amount, uint256 rps)
func (_Marinate *MarinateFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *MarinateRewardAdded) (event.Subscription, error) {

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateRewardAdded)
				if err := _Marinate.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0x6a6f77044107a33658235d41bedbbaf2fe9ccdceb313143c947a5e76e1ec8474.
//
// Solidity: event RewardAdded(address token, uint256 amount, uint256 rps)
func (_Marinate *MarinateFilterer) ParseRewardAdded(log types.Log) (*MarinateRewardAdded, error) {
	event := new(MarinateRewardAdded)
	if err := _Marinate.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarinateRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Marinate contract.
type MarinateRewardClaimedIterator struct {
	Event *MarinateRewardClaimed // Event containing the contract specifics and raw log

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
func (it *MarinateRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateRewardClaimed)
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
		it.Event = new(MarinateRewardClaimed)
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
func (it *MarinateRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateRewardClaimed represents a RewardClaimed event raised by the Marinate contract.
type MarinateRewardClaimed struct {
	Token  common.Address
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address token, address staker, uint256 amount)
func (_Marinate *MarinateFilterer) FilterRewardClaimed(opts *bind.FilterOpts) (*MarinateRewardClaimedIterator, error) {

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "RewardClaimed")
	if err != nil {
		return nil, err
	}
	return &MarinateRewardClaimedIterator{contract: _Marinate.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address token, address staker, uint256 amount)
func (_Marinate *MarinateFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *MarinateRewardClaimed) (event.Subscription, error) {

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "RewardClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateRewardClaimed)
				if err := _Marinate.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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

// ParseRewardClaimed is a log parse operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address token, address staker, uint256 amount)
func (_Marinate *MarinateFilterer) ParseRewardClaimed(log types.Log) (*MarinateRewardClaimed, error) {
	event := new(MarinateRewardClaimed)
	if err := _Marinate.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarinateRewardCollectionIterator is returned from FilterRewardCollection and is used to iterate over the raw logs and unpacked data for RewardCollection events raised by the Marinate contract.
type MarinateRewardCollectionIterator struct {
	Event *MarinateRewardCollection // Event containing the contract specifics and raw log

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
func (it *MarinateRewardCollectionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateRewardCollection)
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
		it.Event = new(MarinateRewardCollection)
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
func (it *MarinateRewardCollectionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateRewardCollectionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateRewardCollection represents a RewardCollection event raised by the Marinate contract.
type MarinateRewardCollection struct {
	Token  common.Address
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardCollection is a free log retrieval operation binding the contract event 0x49d368020d08fde7587b5cf8cf012abc1e73214d401dfe803c531501d0bd4017.
//
// Solidity: event RewardCollection(address token, address addr, uint256 amount)
func (_Marinate *MarinateFilterer) FilterRewardCollection(opts *bind.FilterOpts) (*MarinateRewardCollectionIterator, error) {

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "RewardCollection")
	if err != nil {
		return nil, err
	}
	return &MarinateRewardCollectionIterator{contract: _Marinate.contract, event: "RewardCollection", logs: logs, sub: sub}, nil
}

// WatchRewardCollection is a free log subscription operation binding the contract event 0x49d368020d08fde7587b5cf8cf012abc1e73214d401dfe803c531501d0bd4017.
//
// Solidity: event RewardCollection(address token, address addr, uint256 amount)
func (_Marinate *MarinateFilterer) WatchRewardCollection(opts *bind.WatchOpts, sink chan<- *MarinateRewardCollection) (event.Subscription, error) {

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "RewardCollection")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateRewardCollection)
				if err := _Marinate.contract.UnpackLog(event, "RewardCollection", log); err != nil {
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

// ParseRewardCollection is a log parse operation binding the contract event 0x49d368020d08fde7587b5cf8cf012abc1e73214d401dfe803c531501d0bd4017.
//
// Solidity: event RewardCollection(address token, address addr, uint256 amount)
func (_Marinate *MarinateFilterer) ParseRewardCollection(log types.Log) (*MarinateRewardCollection, error) {
	event := new(MarinateRewardCollection)
	if err := _Marinate.contract.UnpackLog(event, "RewardCollection", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarinateRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Marinate contract.
type MarinateRoleAdminChangedIterator struct {
	Event *MarinateRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MarinateRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateRoleAdminChanged)
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
		it.Event = new(MarinateRoleAdminChanged)
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
func (it *MarinateRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateRoleAdminChanged represents a RoleAdminChanged event raised by the Marinate contract.
type MarinateRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Marinate *MarinateFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MarinateRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MarinateRoleAdminChangedIterator{contract: _Marinate.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Marinate *MarinateFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MarinateRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateRoleAdminChanged)
				if err := _Marinate.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Marinate *MarinateFilterer) ParseRoleAdminChanged(log types.Log) (*MarinateRoleAdminChanged, error) {
	event := new(MarinateRoleAdminChanged)
	if err := _Marinate.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarinateRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Marinate contract.
type MarinateRoleGrantedIterator struct {
	Event *MarinateRoleGranted // Event containing the contract specifics and raw log

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
func (it *MarinateRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateRoleGranted)
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
		it.Event = new(MarinateRoleGranted)
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
func (it *MarinateRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateRoleGranted represents a RoleGranted event raised by the Marinate contract.
type MarinateRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marinate *MarinateFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MarinateRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MarinateRoleGrantedIterator{contract: _Marinate.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marinate *MarinateFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MarinateRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateRoleGranted)
				if err := _Marinate.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marinate *MarinateFilterer) ParseRoleGranted(log types.Log) (*MarinateRoleGranted, error) {
	event := new(MarinateRoleGranted)
	if err := _Marinate.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarinateRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Marinate contract.
type MarinateRoleRevokedIterator struct {
	Event *MarinateRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MarinateRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateRoleRevoked)
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
		it.Event = new(MarinateRoleRevoked)
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
func (it *MarinateRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateRoleRevoked represents a RoleRevoked event raised by the Marinate contract.
type MarinateRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marinate *MarinateFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MarinateRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MarinateRoleRevokedIterator{contract: _Marinate.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marinate *MarinateFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MarinateRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateRoleRevoked)
				if err := _Marinate.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marinate *MarinateFilterer) ParseRoleRevoked(log types.Log) (*MarinateRoleRevoked, error) {
	event := new(MarinateRoleRevoked)
	if err := _Marinate.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarinateStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the Marinate contract.
type MarinateStakeIterator struct {
	Event *MarinateStake // Event containing the contract specifics and raw log

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
func (it *MarinateStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateStake)
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
		it.Event = new(MarinateStake)
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
func (it *MarinateStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateStake represents a Stake event raised by the Marinate contract.
type MarinateStake struct {
	Addr             common.Address
	Amount           *big.Int
	MultipliedAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0x5af417134f72a9d41143ace85b0a26dce6f550f894f2cbc1eeee8810603d91b6.
//
// Solidity: event Stake(address addr, uint256 amount, uint256 multipliedAmount)
func (_Marinate *MarinateFilterer) FilterStake(opts *bind.FilterOpts) (*MarinateStakeIterator, error) {

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "Stake")
	if err != nil {
		return nil, err
	}
	return &MarinateStakeIterator{contract: _Marinate.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0x5af417134f72a9d41143ace85b0a26dce6f550f894f2cbc1eeee8810603d91b6.
//
// Solidity: event Stake(address addr, uint256 amount, uint256 multipliedAmount)
func (_Marinate *MarinateFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *MarinateStake) (event.Subscription, error) {

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "Stake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateStake)
				if err := _Marinate.contract.UnpackLog(event, "Stake", log); err != nil {
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

// ParseStake is a log parse operation binding the contract event 0x5af417134f72a9d41143ace85b0a26dce6f550f894f2cbc1eeee8810603d91b6.
//
// Solidity: event Stake(address addr, uint256 amount, uint256 multipliedAmount)
func (_Marinate *MarinateFilterer) ParseStake(log types.Log) (*MarinateStake, error) {
	event := new(MarinateStake)
	if err := _Marinate.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarinateStakeMultiplierIterator is returned from FilterStakeMultiplier and is used to iterate over the raw logs and unpacked data for StakeMultiplier events raised by the Marinate contract.
type MarinateStakeMultiplierIterator struct {
	Event *MarinateStakeMultiplier // Event containing the contract specifics and raw log

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
func (it *MarinateStakeMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateStakeMultiplier)
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
		it.Event = new(MarinateStakeMultiplier)
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
func (it *MarinateStakeMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateStakeMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateStakeMultiplier represents a StakeMultiplier event raised by the Marinate contract.
type MarinateStakeMultiplier struct {
	Addr             common.Address
	Nft              common.Address
	TokenId          *big.Int
	MultipliedAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStakeMultiplier is a free log retrieval operation binding the contract event 0xc3daea287989ddc10f62ff010419ab2227ed3b8c7497506e97a02fa5184c3ec4.
//
// Solidity: event StakeMultiplier(address addr, address nft, uint256 tokenId, uint256 multipliedAmount)
func (_Marinate *MarinateFilterer) FilterStakeMultiplier(opts *bind.FilterOpts) (*MarinateStakeMultiplierIterator, error) {

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "StakeMultiplier")
	if err != nil {
		return nil, err
	}
	return &MarinateStakeMultiplierIterator{contract: _Marinate.contract, event: "StakeMultiplier", logs: logs, sub: sub}, nil
}

// WatchStakeMultiplier is a free log subscription operation binding the contract event 0xc3daea287989ddc10f62ff010419ab2227ed3b8c7497506e97a02fa5184c3ec4.
//
// Solidity: event StakeMultiplier(address addr, address nft, uint256 tokenId, uint256 multipliedAmount)
func (_Marinate *MarinateFilterer) WatchStakeMultiplier(opts *bind.WatchOpts, sink chan<- *MarinateStakeMultiplier) (event.Subscription, error) {

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "StakeMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateStakeMultiplier)
				if err := _Marinate.contract.UnpackLog(event, "StakeMultiplier", log); err != nil {
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

// ParseStakeMultiplier is a log parse operation binding the contract event 0xc3daea287989ddc10f62ff010419ab2227ed3b8c7497506e97a02fa5184c3ec4.
//
// Solidity: event StakeMultiplier(address addr, address nft, uint256 tokenId, uint256 multipliedAmount)
func (_Marinate *MarinateFilterer) ParseStakeMultiplier(log types.Log) (*MarinateStakeMultiplier, error) {
	event := new(MarinateStakeMultiplier)
	if err := _Marinate.contract.UnpackLog(event, "StakeMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarinateTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Marinate contract.
type MarinateTransferIterator struct {
	Event *MarinateTransfer // Event containing the contract specifics and raw log

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
func (it *MarinateTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateTransfer)
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
		it.Event = new(MarinateTransfer)
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
func (it *MarinateTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateTransfer represents a Transfer event raised by the Marinate contract.
type MarinateTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Marinate *MarinateFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MarinateTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MarinateTransferIterator{contract: _Marinate.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Marinate *MarinateFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MarinateTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateTransfer)
				if err := _Marinate.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Marinate *MarinateFilterer) ParseTransfer(log types.Log) (*MarinateTransfer, error) {
	event := new(MarinateTransfer)
	if err := _Marinate.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarinateWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Marinate contract.
type MarinateWithdrawIterator struct {
	Event *MarinateWithdraw // Event containing the contract specifics and raw log

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
func (it *MarinateWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateWithdraw)
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
		it.Event = new(MarinateWithdraw)
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
func (it *MarinateWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateWithdraw represents a Withdraw event raised by the Marinate contract.
type MarinateWithdraw struct {
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address addr, uint256 amount)
func (_Marinate *MarinateFilterer) FilterWithdraw(opts *bind.FilterOpts) (*MarinateWithdrawIterator, error) {

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &MarinateWithdrawIterator{contract: _Marinate.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address addr, uint256 amount)
func (_Marinate *MarinateFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MarinateWithdraw) (event.Subscription, error) {

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateWithdraw)
				if err := _Marinate.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address addr, uint256 amount)
func (_Marinate *MarinateFilterer) ParseWithdraw(log types.Log) (*MarinateWithdraw, error) {
	event := new(MarinateWithdraw)
	if err := _Marinate.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarinateWithdrawMultiplierIterator is returned from FilterWithdrawMultiplier and is used to iterate over the raw logs and unpacked data for WithdrawMultiplier events raised by the Marinate contract.
type MarinateWithdrawMultiplierIterator struct {
	Event *MarinateWithdrawMultiplier // Event containing the contract specifics and raw log

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
func (it *MarinateWithdrawMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarinateWithdrawMultiplier)
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
		it.Event = new(MarinateWithdrawMultiplier)
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
func (it *MarinateWithdrawMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarinateWithdrawMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarinateWithdrawMultiplier represents a WithdrawMultiplier event raised by the Marinate contract.
type MarinateWithdrawMultiplier struct {
	Addr             common.Address
	Nft              common.Address
	TokenId          *big.Int
	MultipliedAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawMultiplier is a free log retrieval operation binding the contract event 0x45d2a91600dc69305825e109ffd66b221ea47086a5ac4ed7ce4afde017f78bc6.
//
// Solidity: event WithdrawMultiplier(address addr, address nft, uint256 tokenId, uint256 multipliedAmount)
func (_Marinate *MarinateFilterer) FilterWithdrawMultiplier(opts *bind.FilterOpts) (*MarinateWithdrawMultiplierIterator, error) {

	logs, sub, err := _Marinate.contract.FilterLogs(opts, "WithdrawMultiplier")
	if err != nil {
		return nil, err
	}
	return &MarinateWithdrawMultiplierIterator{contract: _Marinate.contract, event: "WithdrawMultiplier", logs: logs, sub: sub}, nil
}

// WatchWithdrawMultiplier is a free log subscription operation binding the contract event 0x45d2a91600dc69305825e109ffd66b221ea47086a5ac4ed7ce4afde017f78bc6.
//
// Solidity: event WithdrawMultiplier(address addr, address nft, uint256 tokenId, uint256 multipliedAmount)
func (_Marinate *MarinateFilterer) WatchWithdrawMultiplier(opts *bind.WatchOpts, sink chan<- *MarinateWithdrawMultiplier) (event.Subscription, error) {

	logs, sub, err := _Marinate.contract.WatchLogs(opts, "WithdrawMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarinateWithdrawMultiplier)
				if err := _Marinate.contract.UnpackLog(event, "WithdrawMultiplier", log); err != nil {
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

// ParseWithdrawMultiplier is a log parse operation binding the contract event 0x45d2a91600dc69305825e109ffd66b221ea47086a5ac4ed7ce4afde017f78bc6.
//
// Solidity: event WithdrawMultiplier(address addr, address nft, uint256 tokenId, uint256 multipliedAmount)
func (_Marinate *MarinateFilterer) ParseWithdrawMultiplier(log types.Log) (*MarinateWithdrawMultiplier, error) {
	event := new(MarinateWithdrawMultiplier)
	if err := _Marinate.contract.UnpackLog(event, "WithdrawMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

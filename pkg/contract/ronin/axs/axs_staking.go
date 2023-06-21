// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package axs

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

// AxsMetaData contains all meta data concerning the Axs contract.
var AxsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_oldAdmin\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_newAdmin\",\"type\":\"address\",\"indexed\":true}],\"name\":\"AdminChanged\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_oldAdmin\",\"type\":\"address\",\"indexed\":true}],\"name\":\"AdminRemoved\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_token\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":true}],\"name\":\"EmergencyUnstaked\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_token\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":true}],\"name\":\"RewardClaimed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_token\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":true}],\"name\":\"Staked\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_token\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":true}],\"name\":\"Unstaked\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_newAdmin\",\"type\":\"address\",\"indexed\":false}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"claimPendingRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"emergencyUnstake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":true,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":false}],\"name\":\"getPendingRewards\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardToken\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":true,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":false}],\"name\":\"getStakingAmount\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":true,\"inputs\":[],\"name\":\"getStakingToken\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":true,\"inputs\":[],\"name\":\"getStakingTotal\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":true,\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bytes4\",\"indexed\":false}],\"name\":\"methodPaused\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_method\",\"type\":\"bytes4\",\"indexed\":false}],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":true,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_method\",\"type\":\"bytes4\",\"indexed\":false}],\"name\":\"paused\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"restakeRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_method\",\"type\":\"bytes4\",\"indexed\":false}],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"unpauseAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"unstake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"unstakeAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false}]",
}

// AxsABI is the input ABI used to generate the binding from.
// Deprecated: Use AxsMetaData.ABI instead.
var AxsABI = AxsMetaData.ABI

// Axs is an auto generated Go binding around an Ethereum contract.
type Axs struct {
	AxsCaller     // Read-only binding to the contract
	AxsTransactor // Write-only binding to the contract
	AxsFilterer   // Log filterer for contract events
}

// AxsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AxsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AxsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AxsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AxsSession struct {
	Contract     *Axs              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AxsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AxsCallerSession struct {
	Contract *AxsCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AxsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AxsTransactorSession struct {
	Contract     *AxsTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AxsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AxsRaw struct {
	Contract *Axs // Generic contract binding to access the raw methods on
}

// AxsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AxsCallerRaw struct {
	Contract *AxsCaller // Generic read-only contract binding to access the raw methods on
}

// AxsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AxsTransactorRaw struct {
	Contract *AxsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAxs creates a new instance of Axs, bound to a specific deployed contract.
func NewAxs(address common.Address, backend bind.ContractBackend) (*Axs, error) {
	contract, err := bindAxs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Axs{AxsCaller: AxsCaller{contract: contract}, AxsTransactor: AxsTransactor{contract: contract}, AxsFilterer: AxsFilterer{contract: contract}}, nil
}

// NewAxsCaller creates a new read-only instance of Axs, bound to a specific deployed contract.
func NewAxsCaller(address common.Address, caller bind.ContractCaller) (*AxsCaller, error) {
	contract, err := bindAxs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AxsCaller{contract: contract}, nil
}

// NewAxsTransactor creates a new write-only instance of Axs, bound to a specific deployed contract.
func NewAxsTransactor(address common.Address, transactor bind.ContractTransactor) (*AxsTransactor, error) {
	contract, err := bindAxs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AxsTransactor{contract: contract}, nil
}

// NewAxsFilterer creates a new log filterer instance of Axs, bound to a specific deployed contract.
func NewAxsFilterer(address common.Address, filterer bind.ContractFilterer) (*AxsFilterer, error) {
	contract, err := bindAxs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AxsFilterer{contract: contract}, nil
}

// bindAxs binds a generic wrapper to an already deployed contract.
func bindAxs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AxsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Axs *AxsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Axs.Contract.AxsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Axs *AxsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axs.Contract.AxsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Axs *AxsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Axs.Contract.AxsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Axs *AxsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Axs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Axs *AxsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Axs *AxsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Axs.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Axs *AxsCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Axs.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Axs *AxsSession) Admin() (common.Address, error) {
	return _Axs.Contract.Admin(&_Axs.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Axs *AxsCallerSession) Admin() (common.Address, error) {
	return _Axs.Contract.Admin(&_Axs.CallOpts)
}

// GetPendingRewards is a free data retrieval call binding the contract method 0xf6ed2017.
//
// Solidity: function getPendingRewards(address _user) view returns(uint256)
func (_Axs *AxsCaller) GetPendingRewards(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Axs.contract.Call(opts, &out, "getPendingRewards", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingRewards is a free data retrieval call binding the contract method 0xf6ed2017.
//
// Solidity: function getPendingRewards(address _user) view returns(uint256)
func (_Axs *AxsSession) GetPendingRewards(_user common.Address) (*big.Int, error) {
	return _Axs.Contract.GetPendingRewards(&_Axs.CallOpts, _user)
}

// GetPendingRewards is a free data retrieval call binding the contract method 0xf6ed2017.
//
// Solidity: function getPendingRewards(address _user) view returns(uint256)
func (_Axs *AxsCallerSession) GetPendingRewards(_user common.Address) (*big.Int, error) {
	return _Axs.Contract.GetPendingRewards(&_Axs.CallOpts, _user)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x69940d79.
//
// Solidity: function getRewardToken() view returns(address)
func (_Axs *AxsCaller) GetRewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Axs.contract.Call(opts, &out, "getRewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardToken is a free data retrieval call binding the contract method 0x69940d79.
//
// Solidity: function getRewardToken() view returns(address)
func (_Axs *AxsSession) GetRewardToken() (common.Address, error) {
	return _Axs.Contract.GetRewardToken(&_Axs.CallOpts)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x69940d79.
//
// Solidity: function getRewardToken() view returns(address)
func (_Axs *AxsCallerSession) GetRewardToken() (common.Address, error) {
	return _Axs.Contract.GetRewardToken(&_Axs.CallOpts)
}

// GetStakingAmount is a free data retrieval call binding the contract method 0x74363daa.
//
// Solidity: function getStakingAmount(address _user) view returns(uint256)
func (_Axs *AxsCaller) GetStakingAmount(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Axs.contract.Call(opts, &out, "getStakingAmount", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingAmount is a free data retrieval call binding the contract method 0x74363daa.
//
// Solidity: function getStakingAmount(address _user) view returns(uint256)
func (_Axs *AxsSession) GetStakingAmount(_user common.Address) (*big.Int, error) {
	return _Axs.Contract.GetStakingAmount(&_Axs.CallOpts, _user)
}

// GetStakingAmount is a free data retrieval call binding the contract method 0x74363daa.
//
// Solidity: function getStakingAmount(address _user) view returns(uint256)
func (_Axs *AxsCallerSession) GetStakingAmount(_user common.Address) (*big.Int, error) {
	return _Axs.Contract.GetStakingAmount(&_Axs.CallOpts, _user)
}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_Axs *AxsCaller) GetStakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Axs.contract.Call(opts, &out, "getStakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_Axs *AxsSession) GetStakingToken() (common.Address, error) {
	return _Axs.Contract.GetStakingToken(&_Axs.CallOpts)
}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_Axs *AxsCallerSession) GetStakingToken() (common.Address, error) {
	return _Axs.Contract.GetStakingToken(&_Axs.CallOpts)
}

// GetStakingTotal is a free data retrieval call binding the contract method 0x61ee98dc.
//
// Solidity: function getStakingTotal() view returns(uint256)
func (_Axs *AxsCaller) GetStakingTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Axs.contract.Call(opts, &out, "getStakingTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingTotal is a free data retrieval call binding the contract method 0x61ee98dc.
//
// Solidity: function getStakingTotal() view returns(uint256)
func (_Axs *AxsSession) GetStakingTotal() (*big.Int, error) {
	return _Axs.Contract.GetStakingTotal(&_Axs.CallOpts)
}

// GetStakingTotal is a free data retrieval call binding the contract method 0x61ee98dc.
//
// Solidity: function getStakingTotal() view returns(uint256)
func (_Axs *AxsCallerSession) GetStakingTotal() (*big.Int, error) {
	return _Axs.Contract.GetStakingTotal(&_Axs.CallOpts)
}

// MethodPaused is a free data retrieval call binding the contract method 0xfc9dd5c4.
//
// Solidity: function methodPaused(bytes4 ) view returns(bool)
func (_Axs *AxsCaller) MethodPaused(opts *bind.CallOpts, arg0 [4]byte) (bool, error) {
	var out []interface{}
	err := _Axs.contract.Call(opts, &out, "methodPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MethodPaused is a free data retrieval call binding the contract method 0xfc9dd5c4.
//
// Solidity: function methodPaused(bytes4 ) view returns(bool)
func (_Axs *AxsSession) MethodPaused(arg0 [4]byte) (bool, error) {
	return _Axs.Contract.MethodPaused(&_Axs.CallOpts, arg0)
}

// MethodPaused is a free data retrieval call binding the contract method 0xfc9dd5c4.
//
// Solidity: function methodPaused(bytes4 ) view returns(bool)
func (_Axs *AxsCallerSession) MethodPaused(arg0 [4]byte) (bool, error) {
	return _Axs.Contract.MethodPaused(&_Axs.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x59237eba.
//
// Solidity: function paused(bytes4 _method) view returns(bool)
func (_Axs *AxsCaller) Paused(opts *bind.CallOpts, _method [4]byte) (bool, error) {
	var out []interface{}
	err := _Axs.contract.Call(opts, &out, "paused", _method)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x59237eba.
//
// Solidity: function paused(bytes4 _method) view returns(bool)
func (_Axs *AxsSession) Paused(_method [4]byte) (bool, error) {
	return _Axs.Contract.Paused(&_Axs.CallOpts, _method)
}

// Paused is a free data retrieval call binding the contract method 0x59237eba.
//
// Solidity: function paused(bytes4 _method) view returns(bool)
func (_Axs *AxsCallerSession) Paused(_method [4]byte) (bool, error) {
	return _Axs.Contract.Paused(&_Axs.CallOpts, _method)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Axs *AxsTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Axs.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Axs *AxsSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Axs.Contract.ChangeAdmin(&_Axs.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Axs *AxsTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Axs.Contract.ChangeAdmin(&_Axs.TransactOpts, _newAdmin)
}

// ClaimPendingRewards is a paid mutator transaction binding the contract method 0x92bd7b2c.
//
// Solidity: function claimPendingRewards() returns()
func (_Axs *AxsTransactor) ClaimPendingRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axs.contract.Transact(opts, "claimPendingRewards")
}

// ClaimPendingRewards is a paid mutator transaction binding the contract method 0x92bd7b2c.
//
// Solidity: function claimPendingRewards() returns()
func (_Axs *AxsSession) ClaimPendingRewards() (*types.Transaction, error) {
	return _Axs.Contract.ClaimPendingRewards(&_Axs.TransactOpts)
}

// ClaimPendingRewards is a paid mutator transaction binding the contract method 0x92bd7b2c.
//
// Solidity: function claimPendingRewards() returns()
func (_Axs *AxsTransactorSession) ClaimPendingRewards() (*types.Transaction, error) {
	return _Axs.Contract.ClaimPendingRewards(&_Axs.TransactOpts)
}

// EmergencyUnstake is a paid mutator transaction binding the contract method 0x7589cf2f.
//
// Solidity: function emergencyUnstake() returns()
func (_Axs *AxsTransactor) EmergencyUnstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axs.contract.Transact(opts, "emergencyUnstake")
}

// EmergencyUnstake is a paid mutator transaction binding the contract method 0x7589cf2f.
//
// Solidity: function emergencyUnstake() returns()
func (_Axs *AxsSession) EmergencyUnstake() (*types.Transaction, error) {
	return _Axs.Contract.EmergencyUnstake(&_Axs.TransactOpts)
}

// EmergencyUnstake is a paid mutator transaction binding the contract method 0x7589cf2f.
//
// Solidity: function emergencyUnstake() returns()
func (_Axs *AxsTransactorSession) EmergencyUnstake() (*types.Transaction, error) {
	return _Axs.Contract.EmergencyUnstake(&_Axs.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _method) returns()
func (_Axs *AxsTransactor) Pause(opts *bind.TransactOpts, _method [4]byte) (*types.Transaction, error) {
	return _Axs.contract.Transact(opts, "pause", _method)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _method) returns()
func (_Axs *AxsSession) Pause(_method [4]byte) (*types.Transaction, error) {
	return _Axs.Contract.Pause(&_Axs.TransactOpts, _method)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _method) returns()
func (_Axs *AxsTransactorSession) Pause(_method [4]byte) (*types.Transaction, error) {
	return _Axs.Contract.Pause(&_Axs.TransactOpts, _method)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Axs *AxsTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axs.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Axs *AxsSession) PauseAll() (*types.Transaction, error) {
	return _Axs.Contract.PauseAll(&_Axs.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Axs *AxsTransactorSession) PauseAll() (*types.Transaction, error) {
	return _Axs.Contract.PauseAll(&_Axs.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Axs *AxsTransactor) RemoveAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axs.contract.Transact(opts, "removeAdmin")
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Axs *AxsSession) RemoveAdmin() (*types.Transaction, error) {
	return _Axs.Contract.RemoveAdmin(&_Axs.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Axs *AxsTransactorSession) RemoveAdmin() (*types.Transaction, error) {
	return _Axs.Contract.RemoveAdmin(&_Axs.TransactOpts)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x3d8527ba.
//
// Solidity: function restakeRewards() returns()
func (_Axs *AxsTransactor) RestakeRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axs.contract.Transact(opts, "restakeRewards")
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x3d8527ba.
//
// Solidity: function restakeRewards() returns()
func (_Axs *AxsSession) RestakeRewards() (*types.Transaction, error) {
	return _Axs.Contract.RestakeRewards(&_Axs.TransactOpts)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x3d8527ba.
//
// Solidity: function restakeRewards() returns()
func (_Axs *AxsTransactorSession) RestakeRewards() (*types.Transaction, error) {
	return _Axs.Contract.RestakeRewards(&_Axs.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_Axs *AxsTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Axs.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_Axs *AxsSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _Axs.Contract.Stake(&_Axs.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_Axs *AxsTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _Axs.Contract.Stake(&_Axs.TransactOpts, _amount)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _method) returns()
func (_Axs *AxsTransactor) Unpause(opts *bind.TransactOpts, _method [4]byte) (*types.Transaction, error) {
	return _Axs.contract.Transact(opts, "unpause", _method)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _method) returns()
func (_Axs *AxsSession) Unpause(_method [4]byte) (*types.Transaction, error) {
	return _Axs.Contract.Unpause(&_Axs.TransactOpts, _method)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _method) returns()
func (_Axs *AxsTransactorSession) Unpause(_method [4]byte) (*types.Transaction, error) {
	return _Axs.Contract.Unpause(&_Axs.TransactOpts, _method)
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_Axs *AxsTransactor) UnpauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axs.contract.Transact(opts, "unpauseAll")
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_Axs *AxsSession) UnpauseAll() (*types.Transaction, error) {
	return _Axs.Contract.UnpauseAll(&_Axs.TransactOpts)
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_Axs *AxsTransactorSession) UnpauseAll() (*types.Transaction, error) {
	return _Axs.Contract.UnpauseAll(&_Axs.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Axs *AxsTransactor) Unstake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Axs.contract.Transact(opts, "unstake", _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Axs *AxsSession) Unstake(_amount *big.Int) (*types.Transaction, error) {
	return _Axs.Contract.Unstake(&_Axs.TransactOpts, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Axs *AxsTransactorSession) Unstake(_amount *big.Int) (*types.Transaction, error) {
	return _Axs.Contract.Unstake(&_Axs.TransactOpts, _amount)
}

// UnstakeAll is a paid mutator transaction binding the contract method 0x35322f37.
//
// Solidity: function unstakeAll() returns()
func (_Axs *AxsTransactor) UnstakeAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axs.contract.Transact(opts, "unstakeAll")
}

// UnstakeAll is a paid mutator transaction binding the contract method 0x35322f37.
//
// Solidity: function unstakeAll() returns()
func (_Axs *AxsSession) UnstakeAll() (*types.Transaction, error) {
	return _Axs.Contract.UnstakeAll(&_Axs.TransactOpts)
}

// UnstakeAll is a paid mutator transaction binding the contract method 0x35322f37.
//
// Solidity: function unstakeAll() returns()
func (_Axs *AxsTransactorSession) UnstakeAll() (*types.Transaction, error) {
	return _Axs.Contract.UnstakeAll(&_Axs.TransactOpts)
}

// AxsAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Axs contract.
type AxsAdminChangedIterator struct {
	Event *AxsAdminChanged // Event containing the contract specifics and raw log

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
func (it *AxsAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AxsAdminChanged)
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
		it.Event = new(AxsAdminChanged)
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
func (it *AxsAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AxsAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AxsAdminChanged represents a AdminChanged event raised by the Axs contract.
type AxsAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Axs *AxsFilterer) FilterAdminChanged(opts *bind.FilterOpts, _oldAdmin []common.Address, _newAdmin []common.Address) (*AxsAdminChangedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Axs.contract.FilterLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return &AxsAdminChangedIterator{contract: _Axs.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Axs *AxsFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AxsAdminChanged, _oldAdmin []common.Address, _newAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Axs.contract.WatchLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AxsAdminChanged)
				if err := _Axs.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Axs *AxsFilterer) ParseAdminChanged(log types.Log) (*AxsAdminChanged, error) {
	event := new(AxsAdminChanged)
	if err := _Axs.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AxsAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the Axs contract.
type AxsAdminRemovedIterator struct {
	Event *AxsAdminRemoved // Event containing the contract specifics and raw log

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
func (it *AxsAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AxsAdminRemoved)
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
		it.Event = new(AxsAdminRemoved)
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
func (it *AxsAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AxsAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AxsAdminRemoved represents a AdminRemoved event raised by the Axs contract.
type AxsAdminRemoved struct {
	OldAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Axs *AxsFilterer) FilterAdminRemoved(opts *bind.FilterOpts, _oldAdmin []common.Address) (*AxsAdminRemovedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Axs.contract.FilterLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return &AxsAdminRemovedIterator{contract: _Axs.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Axs *AxsFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *AxsAdminRemoved, _oldAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Axs.contract.WatchLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AxsAdminRemoved)
				if err := _Axs.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Axs *AxsFilterer) ParseAdminRemoved(log types.Log) (*AxsAdminRemoved, error) {
	event := new(AxsAdminRemoved)
	if err := _Axs.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AxsEmergencyUnstakedIterator is returned from FilterEmergencyUnstaked and is used to iterate over the raw logs and unpacked data for EmergencyUnstaked events raised by the Axs contract.
type AxsEmergencyUnstakedIterator struct {
	Event *AxsEmergencyUnstaked // Event containing the contract specifics and raw log

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
func (it *AxsEmergencyUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AxsEmergencyUnstaked)
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
		it.Event = new(AxsEmergencyUnstaked)
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
func (it *AxsEmergencyUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AxsEmergencyUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AxsEmergencyUnstaked represents a EmergencyUnstaked event raised by the Axs contract.
type AxsEmergencyUnstaked struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyUnstaked is a free log retrieval operation binding the contract event 0x589526ce978dd18660ed3d203132d4f86762231a31e8fe21896e2ee637069551.
//
// Solidity: event EmergencyUnstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Axs *AxsFilterer) FilterEmergencyUnstaked(opts *bind.FilterOpts, _user []common.Address, _token []common.Address, _amount []*big.Int) (*AxsEmergencyUnstakedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Axs.contract.FilterLogs(opts, "EmergencyUnstaked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &AxsEmergencyUnstakedIterator{contract: _Axs.contract, event: "EmergencyUnstaked", logs: logs, sub: sub}, nil
}

// WatchEmergencyUnstaked is a free log subscription operation binding the contract event 0x589526ce978dd18660ed3d203132d4f86762231a31e8fe21896e2ee637069551.
//
// Solidity: event EmergencyUnstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Axs *AxsFilterer) WatchEmergencyUnstaked(opts *bind.WatchOpts, sink chan<- *AxsEmergencyUnstaked, _user []common.Address, _token []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Axs.contract.WatchLogs(opts, "EmergencyUnstaked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AxsEmergencyUnstaked)
				if err := _Axs.contract.UnpackLog(event, "EmergencyUnstaked", log); err != nil {
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

// ParseEmergencyUnstaked is a log parse operation binding the contract event 0x589526ce978dd18660ed3d203132d4f86762231a31e8fe21896e2ee637069551.
//
// Solidity: event EmergencyUnstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Axs *AxsFilterer) ParseEmergencyUnstaked(log types.Log) (*AxsEmergencyUnstaked, error) {
	event := new(AxsEmergencyUnstaked)
	if err := _Axs.contract.UnpackLog(event, "EmergencyUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AxsRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Axs contract.
type AxsRewardClaimedIterator struct {
	Event *AxsRewardClaimed // Event containing the contract specifics and raw log

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
func (it *AxsRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AxsRewardClaimed)
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
		it.Event = new(AxsRewardClaimed)
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
func (it *AxsRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AxsRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AxsRewardClaimed represents a RewardClaimed event raised by the Axs contract.
type AxsRewardClaimed struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Axs *AxsFilterer) FilterRewardClaimed(opts *bind.FilterOpts, _user []common.Address, _token []common.Address, _amount []*big.Int) (*AxsRewardClaimedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Axs.contract.FilterLogs(opts, "RewardClaimed", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &AxsRewardClaimedIterator{contract: _Axs.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Axs *AxsFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *AxsRewardClaimed, _user []common.Address, _token []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Axs.contract.WatchLogs(opts, "RewardClaimed", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AxsRewardClaimed)
				if err := _Axs.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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
// Solidity: event RewardClaimed(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Axs *AxsFilterer) ParseRewardClaimed(log types.Log) (*AxsRewardClaimed, error) {
	event := new(AxsRewardClaimed)
	if err := _Axs.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AxsStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Axs contract.
type AxsStakedIterator struct {
	Event *AxsStaked // Event containing the contract specifics and raw log

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
func (it *AxsStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AxsStaked)
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
		it.Event = new(AxsStaked)
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
func (it *AxsStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AxsStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AxsStaked represents a Staked event raised by the Axs contract.
type AxsStaked struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Axs *AxsFilterer) FilterStaked(opts *bind.FilterOpts, _user []common.Address, _token []common.Address, _amount []*big.Int) (*AxsStakedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Axs.contract.FilterLogs(opts, "Staked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &AxsStakedIterator{contract: _Axs.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Axs *AxsFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *AxsStaked, _user []common.Address, _token []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Axs.contract.WatchLogs(opts, "Staked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AxsStaked)
				if err := _Axs.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Axs *AxsFilterer) ParseStaked(log types.Log) (*AxsStaked, error) {
	event := new(AxsStaked)
	if err := _Axs.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AxsUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Axs contract.
type AxsUnstakedIterator struct {
	Event *AxsUnstaked // Event containing the contract specifics and raw log

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
func (it *AxsUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AxsUnstaked)
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
		it.Event = new(AxsUnstaked)
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
func (it *AxsUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AxsUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AxsUnstaked represents a Unstaked event raised by the Axs contract.
type AxsUnstaked struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Axs *AxsFilterer) FilterUnstaked(opts *bind.FilterOpts, _user []common.Address, _token []common.Address, _amount []*big.Int) (*AxsUnstakedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Axs.contract.FilterLogs(opts, "Unstaked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &AxsUnstakedIterator{contract: _Axs.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Axs *AxsFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *AxsUnstaked, _user []common.Address, _token []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Axs.contract.WatchLogs(opts, "Unstaked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AxsUnstaked)
				if err := _Axs.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Axs *AxsFilterer) ParseUnstaked(log types.Log) (*AxsUnstaked, error) {
	event := new(AxsUnstaked)
	if err := _Axs.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

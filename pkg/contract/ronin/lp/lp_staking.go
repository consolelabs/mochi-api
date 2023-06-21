// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lp

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

// LpMetaData contains all meta data concerning the Lp contract.
var LpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_oldAdmin\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_newAdmin\",\"type\":\"address\",\"indexed\":true}],\"name\":\"AdminChanged\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_oldAdmin\",\"type\":\"address\",\"indexed\":true}],\"name\":\"AdminRemoved\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_token\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":true}],\"name\":\"EmergencyUnstaked\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_token\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":true}],\"name\":\"RewardClaimed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_token\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":true}],\"name\":\"Staked\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_token\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":true}],\"name\":\"Unstaked\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_newAdmin\",\"type\":\"address\",\"indexed\":false}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"claimPendingRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"emergencyUnstake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":true,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":false}],\"name\":\"getPendingRewards\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardToken\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":true,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":false}],\"name\":\"getStakingAmount\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":true,\"inputs\":[],\"name\":\"getStakingToken\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":true,\"inputs\":[],\"name\":\"getStakingTotal\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"constant\":true,\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bytes4\",\"indexed\":false}],\"name\":\"methodPaused\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_method\",\"type\":\"bytes4\",\"indexed\":false}],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"constant\":true,\"inputs\":[{\"internal_type\":\"\",\"name\":\"_method\",\"type\":\"bytes4\",\"indexed\":false}],\"name\":\"paused\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"restakeRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_method\",\"type\":\"bytes4\",\"indexed\":false}],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"unpauseAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"unstake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"unstakeAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false}]",
}

// LpABI is the input ABI used to generate the binding from.
// Deprecated: Use LpMetaData.ABI instead.
var LpABI = LpMetaData.ABI

// Lp is an auto generated Go binding around an Ethereum contract.
type Lp struct {
	LpCaller     // Read-only binding to the contract
	LpTransactor // Write-only binding to the contract
	LpFilterer   // Log filterer for contract events
}

// LpCaller is an auto generated read-only Go binding around an Ethereum contract.
type LpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LpSession struct {
	Contract     *Lp               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LpCallerSession struct {
	Contract *LpCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LpTransactorSession struct {
	Contract     *LpTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LpRaw is an auto generated low-level Go binding around an Ethereum contract.
type LpRaw struct {
	Contract *Lp // Generic contract binding to access the raw methods on
}

// LpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LpCallerRaw struct {
	Contract *LpCaller // Generic read-only contract binding to access the raw methods on
}

// LpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LpTransactorRaw struct {
	Contract *LpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLp creates a new instance of Lp, bound to a specific deployed contract.
func NewLp(address common.Address, backend bind.ContractBackend) (*Lp, error) {
	contract, err := bindLp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lp{LpCaller: LpCaller{contract: contract}, LpTransactor: LpTransactor{contract: contract}, LpFilterer: LpFilterer{contract: contract}}, nil
}

// NewLpCaller creates a new read-only instance of Lp, bound to a specific deployed contract.
func NewLpCaller(address common.Address, caller bind.ContractCaller) (*LpCaller, error) {
	contract, err := bindLp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LpCaller{contract: contract}, nil
}

// NewLpTransactor creates a new write-only instance of Lp, bound to a specific deployed contract.
func NewLpTransactor(address common.Address, transactor bind.ContractTransactor) (*LpTransactor, error) {
	contract, err := bindLp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LpTransactor{contract: contract}, nil
}

// NewLpFilterer creates a new log filterer instance of Lp, bound to a specific deployed contract.
func NewLpFilterer(address common.Address, filterer bind.ContractFilterer) (*LpFilterer, error) {
	contract, err := bindLp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LpFilterer{contract: contract}, nil
}

// bindLp binds a generic wrapper to an already deployed contract.
func bindLp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lp *LpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lp.Contract.LpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lp *LpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.Contract.LpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lp *LpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lp.Contract.LpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lp *LpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lp *LpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lp *LpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lp.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Lp *LpCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Lp *LpSession) Admin() (common.Address, error) {
	return _Lp.Contract.Admin(&_Lp.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Lp *LpCallerSession) Admin() (common.Address, error) {
	return _Lp.Contract.Admin(&_Lp.CallOpts)
}

// GetPendingRewards is a free data retrieval call binding the contract method 0xf6ed2017.
//
// Solidity: function getPendingRewards(address _user) view returns(uint256)
func (_Lp *LpCaller) GetPendingRewards(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getPendingRewards", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingRewards is a free data retrieval call binding the contract method 0xf6ed2017.
//
// Solidity: function getPendingRewards(address _user) view returns(uint256)
func (_Lp *LpSession) GetPendingRewards(_user common.Address) (*big.Int, error) {
	return _Lp.Contract.GetPendingRewards(&_Lp.CallOpts, _user)
}

// GetPendingRewards is a free data retrieval call binding the contract method 0xf6ed2017.
//
// Solidity: function getPendingRewards(address _user) view returns(uint256)
func (_Lp *LpCallerSession) GetPendingRewards(_user common.Address) (*big.Int, error) {
	return _Lp.Contract.GetPendingRewards(&_Lp.CallOpts, _user)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x69940d79.
//
// Solidity: function getRewardToken() view returns(address)
func (_Lp *LpCaller) GetRewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getRewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardToken is a free data retrieval call binding the contract method 0x69940d79.
//
// Solidity: function getRewardToken() view returns(address)
func (_Lp *LpSession) GetRewardToken() (common.Address, error) {
	return _Lp.Contract.GetRewardToken(&_Lp.CallOpts)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x69940d79.
//
// Solidity: function getRewardToken() view returns(address)
func (_Lp *LpCallerSession) GetRewardToken() (common.Address, error) {
	return _Lp.Contract.GetRewardToken(&_Lp.CallOpts)
}

// GetStakingAmount is a free data retrieval call binding the contract method 0x74363daa.
//
// Solidity: function getStakingAmount(address _user) view returns(uint256)
func (_Lp *LpCaller) GetStakingAmount(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getStakingAmount", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingAmount is a free data retrieval call binding the contract method 0x74363daa.
//
// Solidity: function getStakingAmount(address _user) view returns(uint256)
func (_Lp *LpSession) GetStakingAmount(_user common.Address) (*big.Int, error) {
	return _Lp.Contract.GetStakingAmount(&_Lp.CallOpts, _user)
}

// GetStakingAmount is a free data retrieval call binding the contract method 0x74363daa.
//
// Solidity: function getStakingAmount(address _user) view returns(uint256)
func (_Lp *LpCallerSession) GetStakingAmount(_user common.Address) (*big.Int, error) {
	return _Lp.Contract.GetStakingAmount(&_Lp.CallOpts, _user)
}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_Lp *LpCaller) GetStakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getStakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_Lp *LpSession) GetStakingToken() (common.Address, error) {
	return _Lp.Contract.GetStakingToken(&_Lp.CallOpts)
}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_Lp *LpCallerSession) GetStakingToken() (common.Address, error) {
	return _Lp.Contract.GetStakingToken(&_Lp.CallOpts)
}

// GetStakingTotal is a free data retrieval call binding the contract method 0x61ee98dc.
//
// Solidity: function getStakingTotal() view returns(uint256)
func (_Lp *LpCaller) GetStakingTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getStakingTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingTotal is a free data retrieval call binding the contract method 0x61ee98dc.
//
// Solidity: function getStakingTotal() view returns(uint256)
func (_Lp *LpSession) GetStakingTotal() (*big.Int, error) {
	return _Lp.Contract.GetStakingTotal(&_Lp.CallOpts)
}

// GetStakingTotal is a free data retrieval call binding the contract method 0x61ee98dc.
//
// Solidity: function getStakingTotal() view returns(uint256)
func (_Lp *LpCallerSession) GetStakingTotal() (*big.Int, error) {
	return _Lp.Contract.GetStakingTotal(&_Lp.CallOpts)
}

// MethodPaused is a free data retrieval call binding the contract method 0xfc9dd5c4.
//
// Solidity: function methodPaused(bytes4 ) view returns(bool)
func (_Lp *LpCaller) MethodPaused(opts *bind.CallOpts, arg0 [4]byte) (bool, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "methodPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MethodPaused is a free data retrieval call binding the contract method 0xfc9dd5c4.
//
// Solidity: function methodPaused(bytes4 ) view returns(bool)
func (_Lp *LpSession) MethodPaused(arg0 [4]byte) (bool, error) {
	return _Lp.Contract.MethodPaused(&_Lp.CallOpts, arg0)
}

// MethodPaused is a free data retrieval call binding the contract method 0xfc9dd5c4.
//
// Solidity: function methodPaused(bytes4 ) view returns(bool)
func (_Lp *LpCallerSession) MethodPaused(arg0 [4]byte) (bool, error) {
	return _Lp.Contract.MethodPaused(&_Lp.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x59237eba.
//
// Solidity: function paused(bytes4 _method) view returns(bool)
func (_Lp *LpCaller) Paused(opts *bind.CallOpts, _method [4]byte) (bool, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "paused", _method)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x59237eba.
//
// Solidity: function paused(bytes4 _method) view returns(bool)
func (_Lp *LpSession) Paused(_method [4]byte) (bool, error) {
	return _Lp.Contract.Paused(&_Lp.CallOpts, _method)
}

// Paused is a free data retrieval call binding the contract method 0x59237eba.
//
// Solidity: function paused(bytes4 _method) view returns(bool)
func (_Lp *LpCallerSession) Paused(_method [4]byte) (bool, error) {
	return _Lp.Contract.Paused(&_Lp.CallOpts, _method)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Lp *LpTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Lp *LpSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Lp.Contract.ChangeAdmin(&_Lp.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Lp *LpTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Lp.Contract.ChangeAdmin(&_Lp.TransactOpts, _newAdmin)
}

// ClaimPendingRewards is a paid mutator transaction binding the contract method 0x92bd7b2c.
//
// Solidity: function claimPendingRewards() returns()
func (_Lp *LpTransactor) ClaimPendingRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "claimPendingRewards")
}

// ClaimPendingRewards is a paid mutator transaction binding the contract method 0x92bd7b2c.
//
// Solidity: function claimPendingRewards() returns()
func (_Lp *LpSession) ClaimPendingRewards() (*types.Transaction, error) {
	return _Lp.Contract.ClaimPendingRewards(&_Lp.TransactOpts)
}

// ClaimPendingRewards is a paid mutator transaction binding the contract method 0x92bd7b2c.
//
// Solidity: function claimPendingRewards() returns()
func (_Lp *LpTransactorSession) ClaimPendingRewards() (*types.Transaction, error) {
	return _Lp.Contract.ClaimPendingRewards(&_Lp.TransactOpts)
}

// EmergencyUnstake is a paid mutator transaction binding the contract method 0x7589cf2f.
//
// Solidity: function emergencyUnstake() returns()
func (_Lp *LpTransactor) EmergencyUnstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "emergencyUnstake")
}

// EmergencyUnstake is a paid mutator transaction binding the contract method 0x7589cf2f.
//
// Solidity: function emergencyUnstake() returns()
func (_Lp *LpSession) EmergencyUnstake() (*types.Transaction, error) {
	return _Lp.Contract.EmergencyUnstake(&_Lp.TransactOpts)
}

// EmergencyUnstake is a paid mutator transaction binding the contract method 0x7589cf2f.
//
// Solidity: function emergencyUnstake() returns()
func (_Lp *LpTransactorSession) EmergencyUnstake() (*types.Transaction, error) {
	return _Lp.Contract.EmergencyUnstake(&_Lp.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _method) returns()
func (_Lp *LpTransactor) Pause(opts *bind.TransactOpts, _method [4]byte) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "pause", _method)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _method) returns()
func (_Lp *LpSession) Pause(_method [4]byte) (*types.Transaction, error) {
	return _Lp.Contract.Pause(&_Lp.TransactOpts, _method)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _method) returns()
func (_Lp *LpTransactorSession) Pause(_method [4]byte) (*types.Transaction, error) {
	return _Lp.Contract.Pause(&_Lp.TransactOpts, _method)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Lp *LpTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Lp *LpSession) PauseAll() (*types.Transaction, error) {
	return _Lp.Contract.PauseAll(&_Lp.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Lp *LpTransactorSession) PauseAll() (*types.Transaction, error) {
	return _Lp.Contract.PauseAll(&_Lp.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Lp *LpTransactor) RemoveAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "removeAdmin")
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Lp *LpSession) RemoveAdmin() (*types.Transaction, error) {
	return _Lp.Contract.RemoveAdmin(&_Lp.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Lp *LpTransactorSession) RemoveAdmin() (*types.Transaction, error) {
	return _Lp.Contract.RemoveAdmin(&_Lp.TransactOpts)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x3d8527ba.
//
// Solidity: function restakeRewards() returns()
func (_Lp *LpTransactor) RestakeRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "restakeRewards")
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x3d8527ba.
//
// Solidity: function restakeRewards() returns()
func (_Lp *LpSession) RestakeRewards() (*types.Transaction, error) {
	return _Lp.Contract.RestakeRewards(&_Lp.TransactOpts)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x3d8527ba.
//
// Solidity: function restakeRewards() returns()
func (_Lp *LpTransactorSession) RestakeRewards() (*types.Transaction, error) {
	return _Lp.Contract.RestakeRewards(&_Lp.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_Lp *LpTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_Lp *LpSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Stake(&_Lp.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_Lp *LpTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Stake(&_Lp.TransactOpts, _amount)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _method) returns()
func (_Lp *LpTransactor) Unpause(opts *bind.TransactOpts, _method [4]byte) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "unpause", _method)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _method) returns()
func (_Lp *LpSession) Unpause(_method [4]byte) (*types.Transaction, error) {
	return _Lp.Contract.Unpause(&_Lp.TransactOpts, _method)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _method) returns()
func (_Lp *LpTransactorSession) Unpause(_method [4]byte) (*types.Transaction, error) {
	return _Lp.Contract.Unpause(&_Lp.TransactOpts, _method)
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_Lp *LpTransactor) UnpauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "unpauseAll")
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_Lp *LpSession) UnpauseAll() (*types.Transaction, error) {
	return _Lp.Contract.UnpauseAll(&_Lp.TransactOpts)
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_Lp *LpTransactorSession) UnpauseAll() (*types.Transaction, error) {
	return _Lp.Contract.UnpauseAll(&_Lp.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Lp *LpTransactor) Unstake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "unstake", _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Lp *LpSession) Unstake(_amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Unstake(&_Lp.TransactOpts, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Lp *LpTransactorSession) Unstake(_amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Unstake(&_Lp.TransactOpts, _amount)
}

// UnstakeAll is a paid mutator transaction binding the contract method 0x35322f37.
//
// Solidity: function unstakeAll() returns()
func (_Lp *LpTransactor) UnstakeAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "unstakeAll")
}

// UnstakeAll is a paid mutator transaction binding the contract method 0x35322f37.
//
// Solidity: function unstakeAll() returns()
func (_Lp *LpSession) UnstakeAll() (*types.Transaction, error) {
	return _Lp.Contract.UnstakeAll(&_Lp.TransactOpts)
}

// UnstakeAll is a paid mutator transaction binding the contract method 0x35322f37.
//
// Solidity: function unstakeAll() returns()
func (_Lp *LpTransactorSession) UnstakeAll() (*types.Transaction, error) {
	return _Lp.Contract.UnstakeAll(&_Lp.TransactOpts)
}

// LpAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Lp contract.
type LpAdminChangedIterator struct {
	Event *LpAdminChanged // Event containing the contract specifics and raw log

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
func (it *LpAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpAdminChanged)
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
		it.Event = new(LpAdminChanged)
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
func (it *LpAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpAdminChanged represents a AdminChanged event raised by the Lp contract.
type LpAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Lp *LpFilterer) FilterAdminChanged(opts *bind.FilterOpts, _oldAdmin []common.Address, _newAdmin []common.Address) (*LpAdminChangedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return &LpAdminChangedIterator{contract: _Lp.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Lp *LpFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LpAdminChanged, _oldAdmin []common.Address, _newAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpAdminChanged)
				if err := _Lp.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Lp *LpFilterer) ParseAdminChanged(log types.Log) (*LpAdminChanged, error) {
	event := new(LpAdminChanged)
	if err := _Lp.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the Lp contract.
type LpAdminRemovedIterator struct {
	Event *LpAdminRemoved // Event containing the contract specifics and raw log

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
func (it *LpAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpAdminRemoved)
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
		it.Event = new(LpAdminRemoved)
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
func (it *LpAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpAdminRemoved represents a AdminRemoved event raised by the Lp contract.
type LpAdminRemoved struct {
	OldAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Lp *LpFilterer) FilterAdminRemoved(opts *bind.FilterOpts, _oldAdmin []common.Address) (*LpAdminRemovedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return &LpAdminRemovedIterator{contract: _Lp.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Lp *LpFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *LpAdminRemoved, _oldAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpAdminRemoved)
				if err := _Lp.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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
func (_Lp *LpFilterer) ParseAdminRemoved(log types.Log) (*LpAdminRemoved, error) {
	event := new(LpAdminRemoved)
	if err := _Lp.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpEmergencyUnstakedIterator is returned from FilterEmergencyUnstaked and is used to iterate over the raw logs and unpacked data for EmergencyUnstaked events raised by the Lp contract.
type LpEmergencyUnstakedIterator struct {
	Event *LpEmergencyUnstaked // Event containing the contract specifics and raw log

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
func (it *LpEmergencyUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpEmergencyUnstaked)
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
		it.Event = new(LpEmergencyUnstaked)
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
func (it *LpEmergencyUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpEmergencyUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpEmergencyUnstaked represents a EmergencyUnstaked event raised by the Lp contract.
type LpEmergencyUnstaked struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyUnstaked is a free log retrieval operation binding the contract event 0x589526ce978dd18660ed3d203132d4f86762231a31e8fe21896e2ee637069551.
//
// Solidity: event EmergencyUnstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Lp *LpFilterer) FilterEmergencyUnstaked(opts *bind.FilterOpts, _user []common.Address, _token []common.Address, _amount []*big.Int) (*LpEmergencyUnstakedIterator, error) {

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

	logs, sub, err := _Lp.contract.FilterLogs(opts, "EmergencyUnstaked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &LpEmergencyUnstakedIterator{contract: _Lp.contract, event: "EmergencyUnstaked", logs: logs, sub: sub}, nil
}

// WatchEmergencyUnstaked is a free log subscription operation binding the contract event 0x589526ce978dd18660ed3d203132d4f86762231a31e8fe21896e2ee637069551.
//
// Solidity: event EmergencyUnstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Lp *LpFilterer) WatchEmergencyUnstaked(opts *bind.WatchOpts, sink chan<- *LpEmergencyUnstaked, _user []common.Address, _token []common.Address, _amount []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Lp.contract.WatchLogs(opts, "EmergencyUnstaked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpEmergencyUnstaked)
				if err := _Lp.contract.UnpackLog(event, "EmergencyUnstaked", log); err != nil {
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
func (_Lp *LpFilterer) ParseEmergencyUnstaked(log types.Log) (*LpEmergencyUnstaked, error) {
	event := new(LpEmergencyUnstaked)
	if err := _Lp.contract.UnpackLog(event, "EmergencyUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Lp contract.
type LpRewardClaimedIterator struct {
	Event *LpRewardClaimed // Event containing the contract specifics and raw log

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
func (it *LpRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpRewardClaimed)
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
		it.Event = new(LpRewardClaimed)
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
func (it *LpRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpRewardClaimed represents a RewardClaimed event raised by the Lp contract.
type LpRewardClaimed struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Lp *LpFilterer) FilterRewardClaimed(opts *bind.FilterOpts, _user []common.Address, _token []common.Address, _amount []*big.Int) (*LpRewardClaimedIterator, error) {

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

	logs, sub, err := _Lp.contract.FilterLogs(opts, "RewardClaimed", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &LpRewardClaimedIterator{contract: _Lp.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Lp *LpFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *LpRewardClaimed, _user []common.Address, _token []common.Address, _amount []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Lp.contract.WatchLogs(opts, "RewardClaimed", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpRewardClaimed)
				if err := _Lp.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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
func (_Lp *LpFilterer) ParseRewardClaimed(log types.Log) (*LpRewardClaimed, error) {
	event := new(LpRewardClaimed)
	if err := _Lp.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Lp contract.
type LpStakedIterator struct {
	Event *LpStaked // Event containing the contract specifics and raw log

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
func (it *LpStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpStaked)
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
		it.Event = new(LpStaked)
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
func (it *LpStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpStaked represents a Staked event raised by the Lp contract.
type LpStaked struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Lp *LpFilterer) FilterStaked(opts *bind.FilterOpts, _user []common.Address, _token []common.Address, _amount []*big.Int) (*LpStakedIterator, error) {

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

	logs, sub, err := _Lp.contract.FilterLogs(opts, "Staked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &LpStakedIterator{contract: _Lp.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Lp *LpFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *LpStaked, _user []common.Address, _token []common.Address, _amount []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Lp.contract.WatchLogs(opts, "Staked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpStaked)
				if err := _Lp.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_Lp *LpFilterer) ParseStaked(log types.Log) (*LpStaked, error) {
	event := new(LpStaked)
	if err := _Lp.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Lp contract.
type LpUnstakedIterator struct {
	Event *LpUnstaked // Event containing the contract specifics and raw log

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
func (it *LpUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpUnstaked)
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
		it.Event = new(LpUnstaked)
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
func (it *LpUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpUnstaked represents a Unstaked event raised by the Lp contract.
type LpUnstaked struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Lp *LpFilterer) FilterUnstaked(opts *bind.FilterOpts, _user []common.Address, _token []common.Address, _amount []*big.Int) (*LpUnstakedIterator, error) {

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

	logs, sub, err := _Lp.contract.FilterLogs(opts, "Unstaked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &LpUnstakedIterator{contract: _Lp.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Lp *LpFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *LpUnstaked, _user []common.Address, _token []common.Address, _amount []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Lp.contract.WatchLogs(opts, "Unstaked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpUnstaked)
				if err := _Lp.contract.UnpackLog(event, "Unstaked", log); err != nil {
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
func (_Lp *LpFilterer) ParseUnstaked(log types.Log) (*LpUnstaked, error) {
	event := new(LpUnstaked)
	if err := _Lp.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

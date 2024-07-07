// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nekostaking

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

// NekostakingMetaData contains all meta data concerning the Nekostaking contract.
var NekostakingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NekoContract\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"addRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableStaking\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUnstaking\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"_neko\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isPrizeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodFinish\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerTokenStored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardTokenLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"stake\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"unstake\",\"type\":\"bool\"}],\"name\":\"setDisable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDuration\",\"type\":\"uint256\"}],\"name\":\"setRewardDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakeNekosFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"unstakeNekos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NekostakingABI is the input ABI used to generate the binding from.
// Deprecated: Use NekostakingMetaData.ABI instead.
var NekostakingABI = NekostakingMetaData.ABI

// Nekostaking is an auto generated Go binding around an Ethereum contract.
type Nekostaking struct {
	NekostakingCaller     // Read-only binding to the contract
	NekostakingTransactor // Write-only binding to the contract
	NekostakingFilterer   // Log filterer for contract events
}

// NekostakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type NekostakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NekostakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NekostakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NekostakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NekostakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NekostakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NekostakingSession struct {
	Contract     *Nekostaking      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NekostakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NekostakingCallerSession struct {
	Contract *NekostakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NekostakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NekostakingTransactorSession struct {
	Contract     *NekostakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NekostakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type NekostakingRaw struct {
	Contract *Nekostaking // Generic contract binding to access the raw methods on
}

// NekostakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NekostakingCallerRaw struct {
	Contract *NekostakingCaller // Generic read-only contract binding to access the raw methods on
}

// NekostakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NekostakingTransactorRaw struct {
	Contract *NekostakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNekostaking creates a new instance of Nekostaking, bound to a specific deployed contract.
func NewNekostaking(address common.Address, backend bind.ContractBackend) (*Nekostaking, error) {
	contract, err := bindNekostaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nekostaking{NekostakingCaller: NekostakingCaller{contract: contract}, NekostakingTransactor: NekostakingTransactor{contract: contract}, NekostakingFilterer: NekostakingFilterer{contract: contract}}, nil
}

// NewNekostakingCaller creates a new read-only instance of Nekostaking, bound to a specific deployed contract.
func NewNekostakingCaller(address common.Address, caller bind.ContractCaller) (*NekostakingCaller, error) {
	contract, err := bindNekostaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NekostakingCaller{contract: contract}, nil
}

// NewNekostakingTransactor creates a new write-only instance of Nekostaking, bound to a specific deployed contract.
func NewNekostakingTransactor(address common.Address, transactor bind.ContractTransactor) (*NekostakingTransactor, error) {
	contract, err := bindNekostaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NekostakingTransactor{contract: contract}, nil
}

// NewNekostakingFilterer creates a new log filterer instance of Nekostaking, bound to a specific deployed contract.
func NewNekostakingFilterer(address common.Address, filterer bind.ContractFilterer) (*NekostakingFilterer, error) {
	contract, err := bindNekostaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NekostakingFilterer{contract: contract}, nil
}

// bindNekostaking binds a generic wrapper to an already deployed contract.
func bindNekostaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NekostakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nekostaking *NekostakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nekostaking.Contract.NekostakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nekostaking *NekostakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nekostaking.Contract.NekostakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nekostaking *NekostakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nekostaking.Contract.NekostakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nekostaking *NekostakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nekostaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nekostaking *NekostakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nekostaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nekostaking *NekostakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nekostaking.Contract.contract.Transact(opts, method, params...)
}

// NekoContract is a free data retrieval call binding the contract method 0x31046c1d.
//
// Solidity: function NekoContract() view returns(address)
func (_Nekostaking *NekostakingCaller) NekoContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "NekoContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NekoContract is a free data retrieval call binding the contract method 0x31046c1d.
//
// Solidity: function NekoContract() view returns(address)
func (_Nekostaking *NekostakingSession) NekoContract() (common.Address, error) {
	return _Nekostaking.Contract.NekoContract(&_Nekostaking.CallOpts)
}

// NekoContract is a free data retrieval call binding the contract method 0x31046c1d.
//
// Solidity: function NekoContract() view returns(address)
func (_Nekostaking *NekostakingCallerSession) NekoContract() (common.Address, error) {
	return _Nekostaking.Contract.NekoContract(&_Nekostaking.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Nekostaking *NekostakingCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Nekostaking *NekostakingSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Nekostaking.Contract.Balances(&_Nekostaking.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Nekostaking *NekostakingCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Nekostaking.Contract.Balances(&_Nekostaking.CallOpts, arg0)
}

// DisableStaking is a free data retrieval call binding the contract method 0x28696de2.
//
// Solidity: function disableStaking() view returns(bool)
func (_Nekostaking *NekostakingCaller) DisableStaking(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "disableStaking")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DisableStaking is a free data retrieval call binding the contract method 0x28696de2.
//
// Solidity: function disableStaking() view returns(bool)
func (_Nekostaking *NekostakingSession) DisableStaking() (bool, error) {
	return _Nekostaking.Contract.DisableStaking(&_Nekostaking.CallOpts)
}

// DisableStaking is a free data retrieval call binding the contract method 0x28696de2.
//
// Solidity: function disableStaking() view returns(bool)
func (_Nekostaking *NekostakingCallerSession) DisableStaking() (bool, error) {
	return _Nekostaking.Contract.DisableStaking(&_Nekostaking.CallOpts)
}

// DisableUnstaking is a free data retrieval call binding the contract method 0xb46bbf0e.
//
// Solidity: function disableUnstaking() view returns(bool)
func (_Nekostaking *NekostakingCaller) DisableUnstaking(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "disableUnstaking")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DisableUnstaking is a free data retrieval call binding the contract method 0xb46bbf0e.
//
// Solidity: function disableUnstaking() view returns(bool)
func (_Nekostaking *NekostakingSession) DisableUnstaking() (bool, error) {
	return _Nekostaking.Contract.DisableUnstaking(&_Nekostaking.CallOpts)
}

// DisableUnstaking is a free data retrieval call binding the contract method 0xb46bbf0e.
//
// Solidity: function disableUnstaking() view returns(bool)
func (_Nekostaking *NekostakingCallerSession) DisableUnstaking() (bool, error) {
	return _Nekostaking.Contract.DisableUnstaking(&_Nekostaking.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x3e491d47.
//
// Solidity: function earned(address account, uint256 id) view returns(uint256)
func (_Nekostaking *NekostakingCaller) Earned(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "earned", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x3e491d47.
//
// Solidity: function earned(address account, uint256 id) view returns(uint256)
func (_Nekostaking *NekostakingSession) Earned(account common.Address, id *big.Int) (*big.Int, error) {
	return _Nekostaking.Contract.Earned(&_Nekostaking.CallOpts, account, id)
}

// Earned is a free data retrieval call binding the contract method 0x3e491d47.
//
// Solidity: function earned(address account, uint256 id) view returns(uint256)
func (_Nekostaking *NekostakingCallerSession) Earned(account common.Address, id *big.Int) (*big.Int, error) {
	return _Nekostaking.Contract.Earned(&_Nekostaking.CallOpts, account, id)
}

// IsPrizeToken is a free data retrieval call binding the contract method 0xa1100fed.
//
// Solidity: function isPrizeToken(address ) view returns(bool)
func (_Nekostaking *NekostakingCaller) IsPrizeToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "isPrizeToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPrizeToken is a free data retrieval call binding the contract method 0xa1100fed.
//
// Solidity: function isPrizeToken(address ) view returns(bool)
func (_Nekostaking *NekostakingSession) IsPrizeToken(arg0 common.Address) (bool, error) {
	return _Nekostaking.Contract.IsPrizeToken(&_Nekostaking.CallOpts, arg0)
}

// IsPrizeToken is a free data retrieval call binding the contract method 0xa1100fed.
//
// Solidity: function isPrizeToken(address ) view returns(bool)
func (_Nekostaking *NekostakingCallerSession) IsPrizeToken(arg0 common.Address) (bool, error) {
	return _Nekostaking.Contract.IsPrizeToken(&_Nekostaking.CallOpts, arg0)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0xeeca1562.
//
// Solidity: function lastTimeRewardApplicable(uint256 id) view returns(uint256)
func (_Nekostaking *NekostakingCaller) LastTimeRewardApplicable(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "lastTimeRewardApplicable", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0xeeca1562.
//
// Solidity: function lastTimeRewardApplicable(uint256 id) view returns(uint256)
func (_Nekostaking *NekostakingSession) LastTimeRewardApplicable(id *big.Int) (*big.Int, error) {
	return _Nekostaking.Contract.LastTimeRewardApplicable(&_Nekostaking.CallOpts, id)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0xeeca1562.
//
// Solidity: function lastTimeRewardApplicable(uint256 id) view returns(uint256)
func (_Nekostaking *NekostakingCallerSession) LastTimeRewardApplicable(id *big.Int) (*big.Int, error) {
	return _Nekostaking.Contract.LastTimeRewardApplicable(&_Nekostaking.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nekostaking *NekostakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nekostaking *NekostakingSession) Owner() (common.Address, error) {
	return _Nekostaking.Contract.Owner(&_Nekostaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nekostaking *NekostakingCallerSession) Owner() (common.Address, error) {
	return _Nekostaking.Contract.Owner(&_Nekostaking.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Nekostaking *NekostakingCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Nekostaking *NekostakingSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Nekostaking.Contract.Owners(&_Nekostaking.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Nekostaking *NekostakingCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Nekostaking.Contract.Owners(&_Nekostaking.CallOpts, arg0)
}

// RewardInfo is a free data retrieval call binding the contract method 0x81a00f83.
//
// Solidity: function rewardInfo(uint256 ) view returns(address rewardToken, uint256 duration, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored, uint256 rewardBalance)
func (_Nekostaking *NekostakingCaller) RewardInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RewardToken          common.Address
	Duration             *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
	RewardBalance        *big.Int
}, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "rewardInfo", arg0)

	outstruct := new(struct {
		RewardToken          common.Address
		Duration             *big.Int
		PeriodFinish         *big.Int
		RewardRate           *big.Int
		LastUpdateTime       *big.Int
		RewardPerTokenStored *big.Int
		RewardBalance        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Duration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PeriodFinish = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RewardRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RewardPerTokenStored = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RewardBalance = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardInfo is a free data retrieval call binding the contract method 0x81a00f83.
//
// Solidity: function rewardInfo(uint256 ) view returns(address rewardToken, uint256 duration, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored, uint256 rewardBalance)
func (_Nekostaking *NekostakingSession) RewardInfo(arg0 *big.Int) (struct {
	RewardToken          common.Address
	Duration             *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
	RewardBalance        *big.Int
}, error) {
	return _Nekostaking.Contract.RewardInfo(&_Nekostaking.CallOpts, arg0)
}

// RewardInfo is a free data retrieval call binding the contract method 0x81a00f83.
//
// Solidity: function rewardInfo(uint256 ) view returns(address rewardToken, uint256 duration, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored, uint256 rewardBalance)
func (_Nekostaking *NekostakingCallerSession) RewardInfo(arg0 *big.Int) (struct {
	RewardToken          common.Address
	Duration             *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
	RewardBalance        *big.Int
}, error) {
	return _Nekostaking.Contract.RewardInfo(&_Nekostaking.CallOpts, arg0)
}

// RewardPerToken is a free data retrieval call binding the contract method 0x874c120b.
//
// Solidity: function rewardPerToken(uint256 id) view returns(uint256)
func (_Nekostaking *NekostakingCaller) RewardPerToken(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "rewardPerToken", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0x874c120b.
//
// Solidity: function rewardPerToken(uint256 id) view returns(uint256)
func (_Nekostaking *NekostakingSession) RewardPerToken(id *big.Int) (*big.Int, error) {
	return _Nekostaking.Contract.RewardPerToken(&_Nekostaking.CallOpts, id)
}

// RewardPerToken is a free data retrieval call binding the contract method 0x874c120b.
//
// Solidity: function rewardPerToken(uint256 id) view returns(uint256)
func (_Nekostaking *NekostakingCallerSession) RewardPerToken(id *big.Int) (*big.Int, error) {
	return _Nekostaking.Contract.RewardPerToken(&_Nekostaking.CallOpts, id)
}

// RewardTokenLength is a free data retrieval call binding the contract method 0x857cb94a.
//
// Solidity: function rewardTokenLength() view returns(uint256)
func (_Nekostaking *NekostakingCaller) RewardTokenLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "rewardTokenLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardTokenLength is a free data retrieval call binding the contract method 0x857cb94a.
//
// Solidity: function rewardTokenLength() view returns(uint256)
func (_Nekostaking *NekostakingSession) RewardTokenLength() (*big.Int, error) {
	return _Nekostaking.Contract.RewardTokenLength(&_Nekostaking.CallOpts)
}

// RewardTokenLength is a free data retrieval call binding the contract method 0x857cb94a.
//
// Solidity: function rewardTokenLength() view returns(uint256)
func (_Nekostaking *NekostakingCallerSession) RewardTokenLength() (*big.Int, error) {
	return _Nekostaking.Contract.RewardTokenLength(&_Nekostaking.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_Nekostaking *NekostakingCaller) Rewards(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "rewards", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_Nekostaking *NekostakingSession) Rewards(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Nekostaking.Contract.Rewards(&_Nekostaking.CallOpts, arg0, arg1)
}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_Nekostaking *NekostakingCallerSession) Rewards(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Nekostaking.Contract.Rewards(&_Nekostaking.CallOpts, arg0, arg1)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nekostaking *NekostakingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nekostaking *NekostakingSession) TotalSupply() (*big.Int, error) {
	return _Nekostaking.Contract.TotalSupply(&_Nekostaking.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nekostaking *NekostakingCallerSession) TotalSupply() (*big.Int, error) {
	return _Nekostaking.Contract.TotalSupply(&_Nekostaking.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_Nekostaking *NekostakingCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nekostaking.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_Nekostaking *NekostakingSession) UserRewardPerTokenPaid(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Nekostaking.Contract.UserRewardPerTokenPaid(&_Nekostaking.CallOpts, arg0, arg1)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_Nekostaking *NekostakingCallerSession) UserRewardPerTokenPaid(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Nekostaking.Contract.UserRewardPerTokenPaid(&_Nekostaking.CallOpts, arg0, arg1)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0x0e3802e9.
//
// Solidity: function addRewardToken(address _rewardToken, uint256 _duration) returns()
func (_Nekostaking *NekostakingTransactor) AddRewardToken(opts *bind.TransactOpts, _rewardToken common.Address, _duration *big.Int) (*types.Transaction, error) {
	return _Nekostaking.contract.Transact(opts, "addRewardToken", _rewardToken, _duration)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0x0e3802e9.
//
// Solidity: function addRewardToken(address _rewardToken, uint256 _duration) returns()
func (_Nekostaking *NekostakingSession) AddRewardToken(_rewardToken common.Address, _duration *big.Int) (*types.Transaction, error) {
	return _Nekostaking.Contract.AddRewardToken(&_Nekostaking.TransactOpts, _rewardToken, _duration)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0x0e3802e9.
//
// Solidity: function addRewardToken(address _rewardToken, uint256 _duration) returns()
func (_Nekostaking *NekostakingTransactorSession) AddRewardToken(_rewardToken common.Address, _duration *big.Int) (*types.Transaction, error) {
	return _Nekostaking.Contract.AddRewardToken(&_Nekostaking.TransactOpts, _rewardToken, _duration)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Nekostaking *NekostakingTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nekostaking.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Nekostaking *NekostakingSession) GetReward() (*types.Transaction, error) {
	return _Nekostaking.Contract.GetReward(&_Nekostaking.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Nekostaking *NekostakingTransactorSession) GetReward() (*types.Transaction, error) {
	return _Nekostaking.Contract.GetReward(&_Nekostaking.TransactOpts)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address user) returns()
func (_Nekostaking *NekostakingTransactor) GetReward0(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Nekostaking.contract.Transact(opts, "getReward0", user)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address user) returns()
func (_Nekostaking *NekostakingSession) GetReward0(user common.Address) (*types.Transaction, error) {
	return _Nekostaking.Contract.GetReward0(&_Nekostaking.TransactOpts, user)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address user) returns()
func (_Nekostaking *NekostakingTransactorSession) GetReward0(user common.Address) (*types.Transaction, error) {
	return _Nekostaking.Contract.GetReward0(&_Nekostaking.TransactOpts, user)
}

// GetReward1 is a paid mutator transaction binding the contract method 0xf474c8ce.
//
// Solidity: function getReward(address user, uint256 id) returns()
func (_Nekostaking *NekostakingTransactor) GetReward1(opts *bind.TransactOpts, user common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nekostaking.contract.Transact(opts, "getReward1", user, id)
}

// GetReward1 is a paid mutator transaction binding the contract method 0xf474c8ce.
//
// Solidity: function getReward(address user, uint256 id) returns()
func (_Nekostaking *NekostakingSession) GetReward1(user common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nekostaking.Contract.GetReward1(&_Nekostaking.TransactOpts, user, id)
}

// GetReward1 is a paid mutator transaction binding the contract method 0xf474c8ce.
//
// Solidity: function getReward(address user, uint256 id) returns()
func (_Nekostaking *NekostakingTransactorSession) GetReward1(user common.Address, id *big.Int) (*types.Transaction, error) {
	return _Nekostaking.Contract.GetReward1(&_Nekostaking.TransactOpts, user, id)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _neko) returns()
func (_Nekostaking *NekostakingTransactor) Initialize(opts *bind.TransactOpts, _neko common.Address) (*types.Transaction, error) {
	return _Nekostaking.contract.Transact(opts, "initialize", _neko)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _neko) returns()
func (_Nekostaking *NekostakingSession) Initialize(_neko common.Address) (*types.Transaction, error) {
	return _Nekostaking.Contract.Initialize(&_Nekostaking.TransactOpts, _neko)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _neko) returns()
func (_Nekostaking *NekostakingTransactorSession) Initialize(_neko common.Address) (*types.Transaction, error) {
	return _Nekostaking.Contract.Initialize(&_Nekostaking.TransactOpts, _neko)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x246132f9.
//
// Solidity: function notifyRewardAmount(uint256 id, uint256 reward) returns()
func (_Nekostaking *NekostakingTransactor) NotifyRewardAmount(opts *bind.TransactOpts, id *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _Nekostaking.contract.Transact(opts, "notifyRewardAmount", id, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x246132f9.
//
// Solidity: function notifyRewardAmount(uint256 id, uint256 reward) returns()
func (_Nekostaking *NekostakingSession) NotifyRewardAmount(id *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _Nekostaking.Contract.NotifyRewardAmount(&_Nekostaking.TransactOpts, id, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x246132f9.
//
// Solidity: function notifyRewardAmount(uint256 id, uint256 reward) returns()
func (_Nekostaking *NekostakingTransactorSession) NotifyRewardAmount(id *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _Nekostaking.Contract.NotifyRewardAmount(&_Nekostaking.TransactOpts, id, reward)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Nekostaking *NekostakingTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Nekostaking.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Nekostaking *NekostakingSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Nekostaking.Contract.OnERC721Received(&_Nekostaking.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Nekostaking *NekostakingTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Nekostaking.Contract.OnERC721Received(&_Nekostaking.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nekostaking *NekostakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nekostaking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nekostaking *NekostakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nekostaking.Contract.RenounceOwnership(&_Nekostaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nekostaking *NekostakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nekostaking.Contract.RenounceOwnership(&_Nekostaking.TransactOpts)
}

// SetDisable is a paid mutator transaction binding the contract method 0x070a7e5b.
//
// Solidity: function setDisable(bool stake, bool unstake) returns()
func (_Nekostaking *NekostakingTransactor) SetDisable(opts *bind.TransactOpts, stake bool, unstake bool) (*types.Transaction, error) {
	return _Nekostaking.contract.Transact(opts, "setDisable", stake, unstake)
}

// SetDisable is a paid mutator transaction binding the contract method 0x070a7e5b.
//
// Solidity: function setDisable(bool stake, bool unstake) returns()
func (_Nekostaking *NekostakingSession) SetDisable(stake bool, unstake bool) (*types.Transaction, error) {
	return _Nekostaking.Contract.SetDisable(&_Nekostaking.TransactOpts, stake, unstake)
}

// SetDisable is a paid mutator transaction binding the contract method 0x070a7e5b.
//
// Solidity: function setDisable(bool stake, bool unstake) returns()
func (_Nekostaking *NekostakingTransactorSession) SetDisable(stake bool, unstake bool) (*types.Transaction, error) {
	return _Nekostaking.Contract.SetDisable(&_Nekostaking.TransactOpts, stake, unstake)
}

// SetRewardDuration is a paid mutator transaction binding the contract method 0x5d835ae7.
//
// Solidity: function setRewardDuration(uint256 id, uint256 rewardDuration) returns()
func (_Nekostaking *NekostakingTransactor) SetRewardDuration(opts *bind.TransactOpts, id *big.Int, rewardDuration *big.Int) (*types.Transaction, error) {
	return _Nekostaking.contract.Transact(opts, "setRewardDuration", id, rewardDuration)
}

// SetRewardDuration is a paid mutator transaction binding the contract method 0x5d835ae7.
//
// Solidity: function setRewardDuration(uint256 id, uint256 rewardDuration) returns()
func (_Nekostaking *NekostakingSession) SetRewardDuration(id *big.Int, rewardDuration *big.Int) (*types.Transaction, error) {
	return _Nekostaking.Contract.SetRewardDuration(&_Nekostaking.TransactOpts, id, rewardDuration)
}

// SetRewardDuration is a paid mutator transaction binding the contract method 0x5d835ae7.
//
// Solidity: function setRewardDuration(uint256 id, uint256 rewardDuration) returns()
func (_Nekostaking *NekostakingTransactorSession) SetRewardDuration(id *big.Int, rewardDuration *big.Int) (*types.Transaction, error) {
	return _Nekostaking.Contract.SetRewardDuration(&_Nekostaking.TransactOpts, id, rewardDuration)
}

// StakeNekosFor is a paid mutator transaction binding the contract method 0x3aaed54e.
//
// Solidity: function stakeNekosFor(uint256[] tokenIds, address staker) returns()
func (_Nekostaking *NekostakingTransactor) StakeNekosFor(opts *bind.TransactOpts, tokenIds []*big.Int, staker common.Address) (*types.Transaction, error) {
	return _Nekostaking.contract.Transact(opts, "stakeNekosFor", tokenIds, staker)
}

// StakeNekosFor is a paid mutator transaction binding the contract method 0x3aaed54e.
//
// Solidity: function stakeNekosFor(uint256[] tokenIds, address staker) returns()
func (_Nekostaking *NekostakingSession) StakeNekosFor(tokenIds []*big.Int, staker common.Address) (*types.Transaction, error) {
	return _Nekostaking.Contract.StakeNekosFor(&_Nekostaking.TransactOpts, tokenIds, staker)
}

// StakeNekosFor is a paid mutator transaction binding the contract method 0x3aaed54e.
//
// Solidity: function stakeNekosFor(uint256[] tokenIds, address staker) returns()
func (_Nekostaking *NekostakingTransactorSession) StakeNekosFor(tokenIds []*big.Int, staker common.Address) (*types.Transaction, error) {
	return _Nekostaking.Contract.StakeNekosFor(&_Nekostaking.TransactOpts, tokenIds, staker)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nekostaking *NekostakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nekostaking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nekostaking *NekostakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nekostaking.Contract.TransferOwnership(&_Nekostaking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nekostaking *NekostakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nekostaking.Contract.TransferOwnership(&_Nekostaking.TransactOpts, newOwner)
}

// UnstakeNekos is a paid mutator transaction binding the contract method 0x0adef4e1.
//
// Solidity: function unstakeNekos(uint256[] tokenIds) returns()
func (_Nekostaking *NekostakingTransactor) UnstakeNekos(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Nekostaking.contract.Transact(opts, "unstakeNekos", tokenIds)
}

// UnstakeNekos is a paid mutator transaction binding the contract method 0x0adef4e1.
//
// Solidity: function unstakeNekos(uint256[] tokenIds) returns()
func (_Nekostaking *NekostakingSession) UnstakeNekos(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Nekostaking.Contract.UnstakeNekos(&_Nekostaking.TransactOpts, tokenIds)
}

// UnstakeNekos is a paid mutator transaction binding the contract method 0x0adef4e1.
//
// Solidity: function unstakeNekos(uint256[] tokenIds) returns()
func (_Nekostaking *NekostakingTransactorSession) UnstakeNekos(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Nekostaking.Contract.UnstakeNekos(&_Nekostaking.TransactOpts, tokenIds)
}

// NekostakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Nekostaking contract.
type NekostakingOwnershipTransferredIterator struct {
	Event *NekostakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NekostakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NekostakingOwnershipTransferred)
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
		it.Event = new(NekostakingOwnershipTransferred)
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
func (it *NekostakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NekostakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NekostakingOwnershipTransferred represents a OwnershipTransferred event raised by the Nekostaking contract.
type NekostakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nekostaking *NekostakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NekostakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nekostaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NekostakingOwnershipTransferredIterator{contract: _Nekostaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nekostaking *NekostakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NekostakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nekostaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NekostakingOwnershipTransferred)
				if err := _Nekostaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Nekostaking *NekostakingFilterer) ParseOwnershipTransferred(log types.Log) (*NekostakingOwnershipTransferred, error) {
	event := new(NekostakingOwnershipTransferred)
	if err := _Nekostaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NekostakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Nekostaking contract.
type NekostakingStakedIterator struct {
	Event *NekostakingStaked // Event containing the contract specifics and raw log

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
func (it *NekostakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NekostakingStaked)
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
		it.Event = new(NekostakingStaked)
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
func (it *NekostakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NekostakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NekostakingStaked represents a Staked event raised by the Nekostaking contract.
type NekostakingStaked struct {
	Staker common.Address
	Ids    []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x134b166c6094cc1ccbf1e3353ce5c3cd9fd29869051bdb999895854d77cc5ef6.
//
// Solidity: event Staked(address staker, uint256[] ids)
func (_Nekostaking *NekostakingFilterer) FilterStaked(opts *bind.FilterOpts) (*NekostakingStakedIterator, error) {

	logs, sub, err := _Nekostaking.contract.FilterLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return &NekostakingStakedIterator{contract: _Nekostaking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x134b166c6094cc1ccbf1e3353ce5c3cd9fd29869051bdb999895854d77cc5ef6.
//
// Solidity: event Staked(address staker, uint256[] ids)
func (_Nekostaking *NekostakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *NekostakingStaked) (event.Subscription, error) {

	logs, sub, err := _Nekostaking.contract.WatchLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NekostakingStaked)
				if err := _Nekostaking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x134b166c6094cc1ccbf1e3353ce5c3cd9fd29869051bdb999895854d77cc5ef6.
//
// Solidity: event Staked(address staker, uint256[] ids)
func (_Nekostaking *NekostakingFilterer) ParseStaked(log types.Log) (*NekostakingStaked, error) {
	event := new(NekostakingStaked)
	if err := _Nekostaking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NekostakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Nekostaking contract.
type NekostakingUnstakedIterator struct {
	Event *NekostakingUnstaked // Event containing the contract specifics and raw log

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
func (it *NekostakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NekostakingUnstaked)
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
		it.Event = new(NekostakingUnstaked)
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
func (it *NekostakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NekostakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NekostakingUnstaked represents a Unstaked event raised by the Nekostaking contract.
type NekostakingUnstaked struct {
	Staker common.Address
	Ids    []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x20748b935fd9f21155c2e98cb2bd5df6fe86f21b193cebaae8d9ad7db0ba5416.
//
// Solidity: event Unstaked(address staker, uint256[] ids)
func (_Nekostaking *NekostakingFilterer) FilterUnstaked(opts *bind.FilterOpts) (*NekostakingUnstakedIterator, error) {

	logs, sub, err := _Nekostaking.contract.FilterLogs(opts, "Unstaked")
	if err != nil {
		return nil, err
	}
	return &NekostakingUnstakedIterator{contract: _Nekostaking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x20748b935fd9f21155c2e98cb2bd5df6fe86f21b193cebaae8d9ad7db0ba5416.
//
// Solidity: event Unstaked(address staker, uint256[] ids)
func (_Nekostaking *NekostakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *NekostakingUnstaked) (event.Subscription, error) {

	logs, sub, err := _Nekostaking.contract.WatchLogs(opts, "Unstaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NekostakingUnstaked)
				if err := _Nekostaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x20748b935fd9f21155c2e98cb2bd5df6fe86f21b193cebaae8d9ad7db0ba5416.
//
// Solidity: event Unstaked(address staker, uint256[] ids)
func (_Nekostaking *NekostakingFilterer) ParseUnstaked(log types.Log) (*NekostakingUnstaked, error) {
	event := new(NekostakingUnstaked)
	if err := _Nekostaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

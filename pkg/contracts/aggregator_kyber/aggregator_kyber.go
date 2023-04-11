// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggregatorkyber

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

// MetaAggregationRouterV2SwapDescriptionV2 is an auto generated low-level Go binding around an user-defined struct.
type MetaAggregationRouterV2SwapDescriptionV2 struct {
	SrcToken        common.Address
	DstToken        common.Address
	SrcReceivers    []common.Address
	SrcAmounts      []*big.Int
	FeeReceivers    []common.Address
	FeeAmounts      []*big.Int
	DstReceiver     common.Address
	Amount          *big.Int
	MinReturnAmount *big.Int
	Flags           *big.Int
	Permit          []byte
}

// MetaAggregationRouterV2SwapExecutionParams is an auto generated low-level Go binding around an user-defined struct.
type MetaAggregationRouterV2SwapExecutionParams struct {
	CallTarget    common.Address
	ApproveTarget common.Address
	TargetData    []byte
	Desc          MetaAggregationRouterV2SwapDescriptionV2
	ClientData    []byte
}

// AggregatorkyberMetaData contains all meta data concerning the Aggregatorkyber contract.
var AggregatorkyberMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"ClientData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Error\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"}],\"name\":\"Exchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBps\",\"type\":\"bool\"}],\"name\":\"Fee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"callTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTarget\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"feeReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapDescriptionV2\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapExecutionParams\",\"name\":\"execution\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"callTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTarget\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"feeReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapDescriptionV2\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapExecutionParams\",\"name\":\"execution\",\"type\":\"tuple\"}],\"name\":\"swapGeneric\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"feeReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapDescriptionV2\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"executorData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swapSimpleMode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addr\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AggregatorkyberABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregatorkyberMetaData.ABI instead.
var AggregatorkyberABI = AggregatorkyberMetaData.ABI

// Aggregatorkyber is an auto generated Go binding around an Ethereum contract.
type Aggregatorkyber struct {
	AggregatorkyberCaller     // Read-only binding to the contract
	AggregatorkyberTransactor // Write-only binding to the contract
	AggregatorkyberFilterer   // Log filterer for contract events
}

// AggregatorkyberCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggregatorkyberCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorkyberTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregatorkyberTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorkyberFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregatorkyberFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorkyberSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregatorkyberSession struct {
	Contract     *Aggregatorkyber  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggregatorkyberCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregatorkyberCallerSession struct {
	Contract *AggregatorkyberCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AggregatorkyberTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregatorkyberTransactorSession struct {
	Contract     *AggregatorkyberTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AggregatorkyberRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggregatorkyberRaw struct {
	Contract *Aggregatorkyber // Generic contract binding to access the raw methods on
}

// AggregatorkyberCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregatorkyberCallerRaw struct {
	Contract *AggregatorkyberCaller // Generic read-only contract binding to access the raw methods on
}

// AggregatorkyberTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregatorkyberTransactorRaw struct {
	Contract *AggregatorkyberTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregatorkyber creates a new instance of Aggregatorkyber, bound to a specific deployed contract.
func NewAggregatorkyber(address common.Address, backend bind.ContractBackend) (*Aggregatorkyber, error) {
	contract, err := bindAggregatorkyber(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aggregatorkyber{AggregatorkyberCaller: AggregatorkyberCaller{contract: contract}, AggregatorkyberTransactor: AggregatorkyberTransactor{contract: contract}, AggregatorkyberFilterer: AggregatorkyberFilterer{contract: contract}}, nil
}

// NewAggregatorkyberCaller creates a new read-only instance of Aggregatorkyber, bound to a specific deployed contract.
func NewAggregatorkyberCaller(address common.Address, caller bind.ContractCaller) (*AggregatorkyberCaller, error) {
	contract, err := bindAggregatorkyber(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorkyberCaller{contract: contract}, nil
}

// NewAggregatorkyberTransactor creates a new write-only instance of Aggregatorkyber, bound to a specific deployed contract.
func NewAggregatorkyberTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregatorkyberTransactor, error) {
	contract, err := bindAggregatorkyber(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorkyberTransactor{contract: contract}, nil
}

// NewAggregatorkyberFilterer creates a new log filterer instance of Aggregatorkyber, bound to a specific deployed contract.
func NewAggregatorkyberFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregatorkyberFilterer, error) {
	contract, err := bindAggregatorkyber(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatorkyberFilterer{contract: contract}, nil
}

// bindAggregatorkyber binds a generic wrapper to an already deployed contract.
func bindAggregatorkyber(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AggregatorkyberABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggregatorkyber *AggregatorkyberRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggregatorkyber.Contract.AggregatorkyberCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggregatorkyber *AggregatorkyberRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.AggregatorkyberTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggregatorkyber *AggregatorkyberRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.AggregatorkyberTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggregatorkyber *AggregatorkyberCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggregatorkyber.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggregatorkyber *AggregatorkyberTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggregatorkyber *AggregatorkyberTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Aggregatorkyber *AggregatorkyberCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggregatorkyber.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Aggregatorkyber *AggregatorkyberSession) WETH() (common.Address, error) {
	return _Aggregatorkyber.Contract.WETH(&_Aggregatorkyber.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Aggregatorkyber *AggregatorkyberCallerSession) WETH() (common.Address, error) {
	return _Aggregatorkyber.Contract.WETH(&_Aggregatorkyber.CallOpts)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_Aggregatorkyber *AggregatorkyberCaller) IsWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Aggregatorkyber.contract.Call(opts, &out, "isWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_Aggregatorkyber *AggregatorkyberSession) IsWhitelist(arg0 common.Address) (bool, error) {
	return _Aggregatorkyber.Contract.IsWhitelist(&_Aggregatorkyber.CallOpts, arg0)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_Aggregatorkyber *AggregatorkyberCallerSession) IsWhitelist(arg0 common.Address) (bool, error) {
	return _Aggregatorkyber.Contract.IsWhitelist(&_Aggregatorkyber.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggregatorkyber *AggregatorkyberCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggregatorkyber.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggregatorkyber *AggregatorkyberSession) Owner() (common.Address, error) {
	return _Aggregatorkyber.Contract.Owner(&_Aggregatorkyber.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggregatorkyber *AggregatorkyberCallerSession) Owner() (common.Address, error) {
	return _Aggregatorkyber.Contract.Owner(&_Aggregatorkyber.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggregatorkyber *AggregatorkyberTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregatorkyber.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggregatorkyber *AggregatorkyberSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.RenounceOwnership(&_Aggregatorkyber.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggregatorkyber *AggregatorkyberTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.RenounceOwnership(&_Aggregatorkyber.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Aggregatorkyber *AggregatorkyberTransactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Aggregatorkyber.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Aggregatorkyber *AggregatorkyberSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.RescueFunds(&_Aggregatorkyber.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Aggregatorkyber *AggregatorkyberTransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.RescueFunds(&_Aggregatorkyber.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xe21fd0e9.
//
// Solidity: function swap((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Aggregatorkyber *AggregatorkyberTransactor) Swap(opts *bind.TransactOpts, execution MetaAggregationRouterV2SwapExecutionParams) (*types.Transaction, error) {
	return _Aggregatorkyber.contract.Transact(opts, "swap", execution)
}

// Swap is a paid mutator transaction binding the contract method 0xe21fd0e9.
//
// Solidity: function swap((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Aggregatorkyber *AggregatorkyberSession) Swap(execution MetaAggregationRouterV2SwapExecutionParams) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.Swap(&_Aggregatorkyber.TransactOpts, execution)
}

// Swap is a paid mutator transaction binding the contract method 0xe21fd0e9.
//
// Solidity: function swap((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Aggregatorkyber *AggregatorkyberTransactorSession) Swap(execution MetaAggregationRouterV2SwapExecutionParams) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.Swap(&_Aggregatorkyber.TransactOpts, execution)
}

// SwapGeneric is a paid mutator transaction binding the contract method 0x59e50fed.
//
// Solidity: function swapGeneric((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Aggregatorkyber *AggregatorkyberTransactor) SwapGeneric(opts *bind.TransactOpts, execution MetaAggregationRouterV2SwapExecutionParams) (*types.Transaction, error) {
	return _Aggregatorkyber.contract.Transact(opts, "swapGeneric", execution)
}

// SwapGeneric is a paid mutator transaction binding the contract method 0x59e50fed.
//
// Solidity: function swapGeneric((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Aggregatorkyber *AggregatorkyberSession) SwapGeneric(execution MetaAggregationRouterV2SwapExecutionParams) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.SwapGeneric(&_Aggregatorkyber.TransactOpts, execution)
}

// SwapGeneric is a paid mutator transaction binding the contract method 0x59e50fed.
//
// Solidity: function swapGeneric((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Aggregatorkyber *AggregatorkyberTransactorSession) SwapGeneric(execution MetaAggregationRouterV2SwapExecutionParams) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.SwapGeneric(&_Aggregatorkyber.TransactOpts, execution)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0x8af033fb.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_Aggregatorkyber *AggregatorkyberTransactor) SwapSimpleMode(opts *bind.TransactOpts, caller common.Address, desc MetaAggregationRouterV2SwapDescriptionV2, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Aggregatorkyber.contract.Transact(opts, "swapSimpleMode", caller, desc, executorData, clientData)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0x8af033fb.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_Aggregatorkyber *AggregatorkyberSession) SwapSimpleMode(caller common.Address, desc MetaAggregationRouterV2SwapDescriptionV2, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.SwapSimpleMode(&_Aggregatorkyber.TransactOpts, caller, desc, executorData, clientData)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0x8af033fb.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_Aggregatorkyber *AggregatorkyberTransactorSession) SwapSimpleMode(caller common.Address, desc MetaAggregationRouterV2SwapDescriptionV2, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.SwapSimpleMode(&_Aggregatorkyber.TransactOpts, caller, desc, executorData, clientData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggregatorkyber *AggregatorkyberTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Aggregatorkyber.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggregatorkyber *AggregatorkyberSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.TransferOwnership(&_Aggregatorkyber.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggregatorkyber *AggregatorkyberTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.TransferOwnership(&_Aggregatorkyber.TransactOpts, newOwner)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x33320de3.
//
// Solidity: function updateWhitelist(address[] addr, bool[] value) returns()
func (_Aggregatorkyber *AggregatorkyberTransactor) UpdateWhitelist(opts *bind.TransactOpts, addr []common.Address, value []bool) (*types.Transaction, error) {
	return _Aggregatorkyber.contract.Transact(opts, "updateWhitelist", addr, value)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x33320de3.
//
// Solidity: function updateWhitelist(address[] addr, bool[] value) returns()
func (_Aggregatorkyber *AggregatorkyberSession) UpdateWhitelist(addr []common.Address, value []bool) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.UpdateWhitelist(&_Aggregatorkyber.TransactOpts, addr, value)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x33320de3.
//
// Solidity: function updateWhitelist(address[] addr, bool[] value) returns()
func (_Aggregatorkyber *AggregatorkyberTransactorSession) UpdateWhitelist(addr []common.Address, value []bool) (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.UpdateWhitelist(&_Aggregatorkyber.TransactOpts, addr, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Aggregatorkyber *AggregatorkyberTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregatorkyber.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Aggregatorkyber *AggregatorkyberSession) Receive() (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.Receive(&_Aggregatorkyber.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Aggregatorkyber *AggregatorkyberTransactorSession) Receive() (*types.Transaction, error) {
	return _Aggregatorkyber.Contract.Receive(&_Aggregatorkyber.TransactOpts)
}

// AggregatorkyberClientDataIterator is returned from FilterClientData and is used to iterate over the raw logs and unpacked data for ClientData events raised by the Aggregatorkyber contract.
type AggregatorkyberClientDataIterator struct {
	Event *AggregatorkyberClientData // Event containing the contract specifics and raw log

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
func (it *AggregatorkyberClientDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorkyberClientData)
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
		it.Event = new(AggregatorkyberClientData)
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
func (it *AggregatorkyberClientDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorkyberClientDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorkyberClientData represents a ClientData event raised by the Aggregatorkyber contract.
type AggregatorkyberClientData struct {
	ClientData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClientData is a free log retrieval operation binding the contract event 0x095e66fa4dd6a6f7b43fb8444a7bd0edb870508c7abf639bc216efb0bcff9779.
//
// Solidity: event ClientData(bytes clientData)
func (_Aggregatorkyber *AggregatorkyberFilterer) FilterClientData(opts *bind.FilterOpts) (*AggregatorkyberClientDataIterator, error) {

	logs, sub, err := _Aggregatorkyber.contract.FilterLogs(opts, "ClientData")
	if err != nil {
		return nil, err
	}
	return &AggregatorkyberClientDataIterator{contract: _Aggregatorkyber.contract, event: "ClientData", logs: logs, sub: sub}, nil
}

// WatchClientData is a free log subscription operation binding the contract event 0x095e66fa4dd6a6f7b43fb8444a7bd0edb870508c7abf639bc216efb0bcff9779.
//
// Solidity: event ClientData(bytes clientData)
func (_Aggregatorkyber *AggregatorkyberFilterer) WatchClientData(opts *bind.WatchOpts, sink chan<- *AggregatorkyberClientData) (event.Subscription, error) {

	logs, sub, err := _Aggregatorkyber.contract.WatchLogs(opts, "ClientData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorkyberClientData)
				if err := _Aggregatorkyber.contract.UnpackLog(event, "ClientData", log); err != nil {
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

// ParseClientData is a log parse operation binding the contract event 0x095e66fa4dd6a6f7b43fb8444a7bd0edb870508c7abf639bc216efb0bcff9779.
//
// Solidity: event ClientData(bytes clientData)
func (_Aggregatorkyber *AggregatorkyberFilterer) ParseClientData(log types.Log) (*AggregatorkyberClientData, error) {
	event := new(AggregatorkyberClientData)
	if err := _Aggregatorkyber.contract.UnpackLog(event, "ClientData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorkyberErrorIterator is returned from FilterError and is used to iterate over the raw logs and unpacked data for Error events raised by the Aggregatorkyber contract.
type AggregatorkyberErrorIterator struct {
	Event *AggregatorkyberError // Event containing the contract specifics and raw log

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
func (it *AggregatorkyberErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorkyberError)
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
		it.Event = new(AggregatorkyberError)
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
func (it *AggregatorkyberErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorkyberErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorkyberError represents a Error event raised by the Aggregatorkyber contract.
type AggregatorkyberError struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterError is a free log retrieval operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Aggregatorkyber *AggregatorkyberFilterer) FilterError(opts *bind.FilterOpts) (*AggregatorkyberErrorIterator, error) {

	logs, sub, err := _Aggregatorkyber.contract.FilterLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return &AggregatorkyberErrorIterator{contract: _Aggregatorkyber.contract, event: "Error", logs: logs, sub: sub}, nil
}

// WatchError is a free log subscription operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Aggregatorkyber *AggregatorkyberFilterer) WatchError(opts *bind.WatchOpts, sink chan<- *AggregatorkyberError) (event.Subscription, error) {

	logs, sub, err := _Aggregatorkyber.contract.WatchLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorkyberError)
				if err := _Aggregatorkyber.contract.UnpackLog(event, "Error", log); err != nil {
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

// ParseError is a log parse operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Aggregatorkyber *AggregatorkyberFilterer) ParseError(log types.Log) (*AggregatorkyberError, error) {
	event := new(AggregatorkyberError)
	if err := _Aggregatorkyber.contract.UnpackLog(event, "Error", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorkyberExchangeIterator is returned from FilterExchange and is used to iterate over the raw logs and unpacked data for Exchange events raised by the Aggregatorkyber contract.
type AggregatorkyberExchangeIterator struct {
	Event *AggregatorkyberExchange // Event containing the contract specifics and raw log

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
func (it *AggregatorkyberExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorkyberExchange)
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
		it.Event = new(AggregatorkyberExchange)
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
func (it *AggregatorkyberExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorkyberExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorkyberExchange represents a Exchange event raised by the Aggregatorkyber contract.
type AggregatorkyberExchange struct {
	Pair      common.Address
	AmountOut *big.Int
	Output    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExchange is a free log retrieval operation binding the contract event 0xddac40937f35385a34f721af292e5a83fc5b840f722bff57c2fc71adba708c48.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (_Aggregatorkyber *AggregatorkyberFilterer) FilterExchange(opts *bind.FilterOpts) (*AggregatorkyberExchangeIterator, error) {

	logs, sub, err := _Aggregatorkyber.contract.FilterLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return &AggregatorkyberExchangeIterator{contract: _Aggregatorkyber.contract, event: "Exchange", logs: logs, sub: sub}, nil
}

// WatchExchange is a free log subscription operation binding the contract event 0xddac40937f35385a34f721af292e5a83fc5b840f722bff57c2fc71adba708c48.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (_Aggregatorkyber *AggregatorkyberFilterer) WatchExchange(opts *bind.WatchOpts, sink chan<- *AggregatorkyberExchange) (event.Subscription, error) {

	logs, sub, err := _Aggregatorkyber.contract.WatchLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorkyberExchange)
				if err := _Aggregatorkyber.contract.UnpackLog(event, "Exchange", log); err != nil {
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

// ParseExchange is a log parse operation binding the contract event 0xddac40937f35385a34f721af292e5a83fc5b840f722bff57c2fc71adba708c48.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (_Aggregatorkyber *AggregatorkyberFilterer) ParseExchange(log types.Log) (*AggregatorkyberExchange, error) {
	event := new(AggregatorkyberExchange)
	if err := _Aggregatorkyber.contract.UnpackLog(event, "Exchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorkyberFeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the Aggregatorkyber contract.
type AggregatorkyberFeeIterator struct {
	Event *AggregatorkyberFee // Event containing the contract specifics and raw log

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
func (it *AggregatorkyberFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorkyberFee)
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
		it.Event = new(AggregatorkyberFee)
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
func (it *AggregatorkyberFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorkyberFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorkyberFee represents a Fee event raised by the Aggregatorkyber contract.
type AggregatorkyberFee struct {
	Token       common.Address
	TotalAmount *big.Int
	TotalFee    *big.Int
	Recipients  []common.Address
	Amounts     []*big.Int
	IsBps       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0x4c39b7ce5f4f514f45cb6f82b171b8b0b7f2cbf488ad28e4eff451588e2f014b.
//
// Solidity: event Fee(address token, uint256 totalAmount, uint256 totalFee, address[] recipients, uint256[] amounts, bool isBps)
func (_Aggregatorkyber *AggregatorkyberFilterer) FilterFee(opts *bind.FilterOpts) (*AggregatorkyberFeeIterator, error) {

	logs, sub, err := _Aggregatorkyber.contract.FilterLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return &AggregatorkyberFeeIterator{contract: _Aggregatorkyber.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0x4c39b7ce5f4f514f45cb6f82b171b8b0b7f2cbf488ad28e4eff451588e2f014b.
//
// Solidity: event Fee(address token, uint256 totalAmount, uint256 totalFee, address[] recipients, uint256[] amounts, bool isBps)
func (_Aggregatorkyber *AggregatorkyberFilterer) WatchFee(opts *bind.WatchOpts, sink chan<- *AggregatorkyberFee) (event.Subscription, error) {

	logs, sub, err := _Aggregatorkyber.contract.WatchLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorkyberFee)
				if err := _Aggregatorkyber.contract.UnpackLog(event, "Fee", log); err != nil {
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

// ParseFee is a log parse operation binding the contract event 0x4c39b7ce5f4f514f45cb6f82b171b8b0b7f2cbf488ad28e4eff451588e2f014b.
//
// Solidity: event Fee(address token, uint256 totalAmount, uint256 totalFee, address[] recipients, uint256[] amounts, bool isBps)
func (_Aggregatorkyber *AggregatorkyberFilterer) ParseFee(log types.Log) (*AggregatorkyberFee, error) {
	event := new(AggregatorkyberFee)
	if err := _Aggregatorkyber.contract.UnpackLog(event, "Fee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorkyberOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Aggregatorkyber contract.
type AggregatorkyberOwnershipTransferredIterator struct {
	Event *AggregatorkyberOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AggregatorkyberOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorkyberOwnershipTransferred)
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
		it.Event = new(AggregatorkyberOwnershipTransferred)
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
func (it *AggregatorkyberOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorkyberOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorkyberOwnershipTransferred represents a OwnershipTransferred event raised by the Aggregatorkyber contract.
type AggregatorkyberOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggregatorkyber *AggregatorkyberFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AggregatorkyberOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aggregatorkyber.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorkyberOwnershipTransferredIterator{contract: _Aggregatorkyber.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggregatorkyber *AggregatorkyberFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AggregatorkyberOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aggregatorkyber.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorkyberOwnershipTransferred)
				if err := _Aggregatorkyber.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Aggregatorkyber *AggregatorkyberFilterer) ParseOwnershipTransferred(log types.Log) (*AggregatorkyberOwnershipTransferred, error) {
	event := new(AggregatorkyberOwnershipTransferred)
	if err := _Aggregatorkyber.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorkyberSwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the Aggregatorkyber contract.
type AggregatorkyberSwappedIterator struct {
	Event *AggregatorkyberSwapped // Event containing the contract specifics and raw log

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
func (it *AggregatorkyberSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorkyberSwapped)
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
		it.Event = new(AggregatorkyberSwapped)
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
func (it *AggregatorkyberSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorkyberSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorkyberSwapped represents a Swapped event raised by the Aggregatorkyber contract.
type AggregatorkyberSwapped struct {
	Sender       common.Address
	SrcToken     common.Address
	DstToken     common.Address
	DstReceiver  common.Address
	SpentAmount  *big.Int
	ReturnAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Aggregatorkyber *AggregatorkyberFilterer) FilterSwapped(opts *bind.FilterOpts) (*AggregatorkyberSwappedIterator, error) {

	logs, sub, err := _Aggregatorkyber.contract.FilterLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return &AggregatorkyberSwappedIterator{contract: _Aggregatorkyber.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Aggregatorkyber *AggregatorkyberFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *AggregatorkyberSwapped) (event.Subscription, error) {

	logs, sub, err := _Aggregatorkyber.contract.WatchLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorkyberSwapped)
				if err := _Aggregatorkyber.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Aggregatorkyber *AggregatorkyberFilterer) ParseSwapped(log types.Log) (*AggregatorkyberSwapped, error) {
	event := new(AggregatorkyberSwapped)
	if err := _Aggregatorkyber.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

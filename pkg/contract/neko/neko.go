// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package neko

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

// NekoMetaData contains all meta data concerning the Neko contract.
var NekoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_NEKO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NR_OF_SUPPORTED_TOKEN\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WHITELIST_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"addWhitelists\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numNeko\",\"type\":\"uint256\"}],\"name\":\"calculatePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasSaleStarted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasWhitelistStarted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numNeko\",\"type\":\"uint256\"}],\"name\":\"mintNeko\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numNeko\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenId\",\"type\":\"uint8\"}],\"name\":\"mintNekoByCustomToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintWhitelist\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenId\",\"type\":\"uint8\"}],\"name\":\"mintWhitelistByCustomToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseWhitelistSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"removeWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numNeko\",\"type\":\"uint256\"}],\"name\":\"reserveGiveaway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAddressRevealer\",\"name\":\"_revealer\",\"type\":\"address\"}],\"name\":\"setRevealer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenId\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"conversionRate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"setSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startWhitelistSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"supportedToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"conversionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenId\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenId\",\"type\":\"uint8\"}],\"name\":\"turnOffSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenId\",\"type\":\"uint8\"}],\"name\":\"turnOnSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelister\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NekoABI is the input ABI used to generate the binding from.
// Deprecated: Use NekoMetaData.ABI instead.
var NekoABI = NekoMetaData.ABI

// Neko is an auto generated Go binding around an Ethereum contract.
type Neko struct {
	NekoCaller     // Read-only binding to the contract
	NekoTransactor // Write-only binding to the contract
	NekoFilterer   // Log filterer for contract events
}

// NekoCaller is an auto generated read-only Go binding around an Ethereum contract.
type NekoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NekoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NekoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NekoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NekoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NekoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NekoSession struct {
	Contract     *Neko             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NekoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NekoCallerSession struct {
	Contract *NekoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NekoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NekoTransactorSession struct {
	Contract     *NekoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NekoRaw is an auto generated low-level Go binding around an Ethereum contract.
type NekoRaw struct {
	Contract *Neko // Generic contract binding to access the raw methods on
}

// NekoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NekoCallerRaw struct {
	Contract *NekoCaller // Generic read-only contract binding to access the raw methods on
}

// NekoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NekoTransactorRaw struct {
	Contract *NekoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNeko creates a new instance of Neko, bound to a specific deployed contract.
func NewNeko(address common.Address, backend bind.ContractBackend) (*Neko, error) {
	contract, err := bindNeko(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Neko{NekoCaller: NekoCaller{contract: contract}, NekoTransactor: NekoTransactor{contract: contract}, NekoFilterer: NekoFilterer{contract: contract}}, nil
}

// NewNekoCaller creates a new read-only instance of Neko, bound to a specific deployed contract.
func NewNekoCaller(address common.Address, caller bind.ContractCaller) (*NekoCaller, error) {
	contract, err := bindNeko(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NekoCaller{contract: contract}, nil
}

// NewNekoTransactor creates a new write-only instance of Neko, bound to a specific deployed contract.
func NewNekoTransactor(address common.Address, transactor bind.ContractTransactor) (*NekoTransactor, error) {
	contract, err := bindNeko(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NekoTransactor{contract: contract}, nil
}

// NewNekoFilterer creates a new log filterer instance of Neko, bound to a specific deployed contract.
func NewNekoFilterer(address common.Address, filterer bind.ContractFilterer) (*NekoFilterer, error) {
	contract, err := bindNeko(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NekoFilterer{contract: contract}, nil
}

// bindNeko binds a generic wrapper to an already deployed contract.
func bindNeko(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NekoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Neko *NekoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Neko.Contract.NekoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Neko *NekoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neko.Contract.NekoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Neko *NekoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Neko.Contract.NekoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Neko *NekoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Neko.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Neko *NekoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neko.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Neko *NekoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Neko.Contract.contract.Transact(opts, method, params...)
}

// MAXNEKO is a free data retrieval call binding the contract method 0x8219c09e.
//
// Solidity: function MAX_NEKO() view returns(uint256)
func (_Neko *NekoCaller) MAXNEKO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "MAX_NEKO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNEKO is a free data retrieval call binding the contract method 0x8219c09e.
//
// Solidity: function MAX_NEKO() view returns(uint256)
func (_Neko *NekoSession) MAXNEKO() (*big.Int, error) {
	return _Neko.Contract.MAXNEKO(&_Neko.CallOpts)
}

// MAXNEKO is a free data retrieval call binding the contract method 0x8219c09e.
//
// Solidity: function MAX_NEKO() view returns(uint256)
func (_Neko *NekoCallerSession) MAXNEKO() (*big.Int, error) {
	return _Neko.Contract.MAXNEKO(&_Neko.CallOpts)
}

// NROFSUPPORTEDTOKEN is a free data retrieval call binding the contract method 0xa8b76e51.
//
// Solidity: function NR_OF_SUPPORTED_TOKEN() view returns(uint16)
func (_Neko *NekoCaller) NROFSUPPORTEDTOKEN(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "NR_OF_SUPPORTED_TOKEN")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// NROFSUPPORTEDTOKEN is a free data retrieval call binding the contract method 0xa8b76e51.
//
// Solidity: function NR_OF_SUPPORTED_TOKEN() view returns(uint16)
func (_Neko *NekoSession) NROFSUPPORTEDTOKEN() (uint16, error) {
	return _Neko.Contract.NROFSUPPORTEDTOKEN(&_Neko.CallOpts)
}

// NROFSUPPORTEDTOKEN is a free data retrieval call binding the contract method 0xa8b76e51.
//
// Solidity: function NR_OF_SUPPORTED_TOKEN() view returns(uint16)
func (_Neko *NekoCallerSession) NROFSUPPORTEDTOKEN() (uint16, error) {
	return _Neko.Contract.NROFSUPPORTEDTOKEN(&_Neko.CallOpts)
}

// WHITELISTPRICE is a free data retrieval call binding the contract method 0x17e7f295.
//
// Solidity: function WHITELIST_PRICE() view returns(uint256)
func (_Neko *NekoCaller) WHITELISTPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "WHITELIST_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WHITELISTPRICE is a free data retrieval call binding the contract method 0x17e7f295.
//
// Solidity: function WHITELIST_PRICE() view returns(uint256)
func (_Neko *NekoSession) WHITELISTPRICE() (*big.Int, error) {
	return _Neko.Contract.WHITELISTPRICE(&_Neko.CallOpts)
}

// WHITELISTPRICE is a free data retrieval call binding the contract method 0x17e7f295.
//
// Solidity: function WHITELIST_PRICE() view returns(uint256)
func (_Neko *NekoCallerSession) WHITELISTPRICE() (*big.Int, error) {
	return _Neko.Contract.WHITELISTPRICE(&_Neko.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Neko *NekoCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Neko *NekoSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Neko.Contract.BalanceOf(&_Neko.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Neko *NekoCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Neko.Contract.BalanceOf(&_Neko.CallOpts, owner)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_Neko *NekoCaller) Base(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_Neko *NekoSession) Base() (*big.Int, error) {
	return _Neko.Contract.Base(&_Neko.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_Neko *NekoCallerSession) Base() (*big.Int, error) {
	return _Neko.Contract.Base(&_Neko.CallOpts)
}

// CalculatePrice is a free data retrieval call binding the contract method 0xae104265.
//
// Solidity: function calculatePrice(uint256 numNeko) view returns(uint256)
func (_Neko *NekoCaller) CalculatePrice(opts *bind.CallOpts, numNeko *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "calculatePrice", numNeko)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePrice is a free data retrieval call binding the contract method 0xae104265.
//
// Solidity: function calculatePrice(uint256 numNeko) view returns(uint256)
func (_Neko *NekoSession) CalculatePrice(numNeko *big.Int) (*big.Int, error) {
	return _Neko.Contract.CalculatePrice(&_Neko.CallOpts, numNeko)
}

// CalculatePrice is a free data retrieval call binding the contract method 0xae104265.
//
// Solidity: function calculatePrice(uint256 numNeko) view returns(uint256)
func (_Neko *NekoCallerSession) CalculatePrice(numNeko *big.Int) (*big.Int, error) {
	return _Neko.Contract.CalculatePrice(&_Neko.CallOpts, numNeko)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Neko *NekoCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Neko *NekoSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Neko.Contract.GetApproved(&_Neko.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Neko *NekoCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Neko.Contract.GetApproved(&_Neko.CallOpts, tokenId)
}

// HasSaleStarted is a free data retrieval call binding the contract method 0x1c8b232d.
//
// Solidity: function hasSaleStarted() view returns(bool)
func (_Neko *NekoCaller) HasSaleStarted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "hasSaleStarted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSaleStarted is a free data retrieval call binding the contract method 0x1c8b232d.
//
// Solidity: function hasSaleStarted() view returns(bool)
func (_Neko *NekoSession) HasSaleStarted() (bool, error) {
	return _Neko.Contract.HasSaleStarted(&_Neko.CallOpts)
}

// HasSaleStarted is a free data retrieval call binding the contract method 0x1c8b232d.
//
// Solidity: function hasSaleStarted() view returns(bool)
func (_Neko *NekoCallerSession) HasSaleStarted() (bool, error) {
	return _Neko.Contract.HasSaleStarted(&_Neko.CallOpts)
}

// HasWhitelistStarted is a free data retrieval call binding the contract method 0xb4f519f1.
//
// Solidity: function hasWhitelistStarted() view returns(bool)
func (_Neko *NekoCaller) HasWhitelistStarted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "hasWhitelistStarted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasWhitelistStarted is a free data retrieval call binding the contract method 0xb4f519f1.
//
// Solidity: function hasWhitelistStarted() view returns(bool)
func (_Neko *NekoSession) HasWhitelistStarted() (bool, error) {
	return _Neko.Contract.HasWhitelistStarted(&_Neko.CallOpts)
}

// HasWhitelistStarted is a free data retrieval call binding the contract method 0xb4f519f1.
//
// Solidity: function hasWhitelistStarted() view returns(bool)
func (_Neko *NekoCallerSession) HasWhitelistStarted() (bool, error) {
	return _Neko.Contract.HasWhitelistStarted(&_Neko.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Neko *NekoCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Neko *NekoSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Neko.Contract.IsApprovedForAll(&_Neko.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Neko *NekoCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Neko.Contract.IsApprovedForAll(&_Neko.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Neko *NekoCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Neko *NekoSession) Name() (string, error) {
	return _Neko.Contract.Name(&_Neko.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Neko *NekoCallerSession) Name() (string, error) {
	return _Neko.Contract.Name(&_Neko.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Neko *NekoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Neko *NekoSession) Owner() (common.Address, error) {
	return _Neko.Contract.Owner(&_Neko.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Neko *NekoCallerSession) Owner() (common.Address, error) {
	return _Neko.Contract.Owner(&_Neko.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Neko *NekoCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Neko *NekoSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Neko.Contract.OwnerOf(&_Neko.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Neko *NekoCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Neko.Contract.OwnerOf(&_Neko.CallOpts, tokenId)
}

// SupportedToken is a free data retrieval call binding the contract method 0x35bc5db3.
//
// Solidity: function supportedToken(uint8 ) view returns(address tokenAddress, uint256 conversionRate, uint8 tokenId, string symbol, bool initialized, bool supported)
func (_Neko *NekoCaller) SupportedToken(opts *bind.CallOpts, arg0 uint8) (struct {
	TokenAddress   common.Address
	ConversionRate *big.Int
	TokenId        uint8
	Symbol         string
	Initialized    bool
	Supported      bool
}, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "supportedToken", arg0)

	outstruct := new(struct {
		TokenAddress   common.Address
		ConversionRate *big.Int
		TokenId        uint8
		Symbol         string
		Initialized    bool
		Supported      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ConversionRate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TokenId = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Symbol = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Initialized = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Supported = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// SupportedToken is a free data retrieval call binding the contract method 0x35bc5db3.
//
// Solidity: function supportedToken(uint8 ) view returns(address tokenAddress, uint256 conversionRate, uint8 tokenId, string symbol, bool initialized, bool supported)
func (_Neko *NekoSession) SupportedToken(arg0 uint8) (struct {
	TokenAddress   common.Address
	ConversionRate *big.Int
	TokenId        uint8
	Symbol         string
	Initialized    bool
	Supported      bool
}, error) {
	return _Neko.Contract.SupportedToken(&_Neko.CallOpts, arg0)
}

// SupportedToken is a free data retrieval call binding the contract method 0x35bc5db3.
//
// Solidity: function supportedToken(uint8 ) view returns(address tokenAddress, uint256 conversionRate, uint8 tokenId, string symbol, bool initialized, bool supported)
func (_Neko *NekoCallerSession) SupportedToken(arg0 uint8) (struct {
	TokenAddress   common.Address
	ConversionRate *big.Int
	TokenId        uint8
	Symbol         string
	Initialized    bool
	Supported      bool
}, error) {
	return _Neko.Contract.SupportedToken(&_Neko.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Neko *NekoCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Neko *NekoSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Neko.Contract.SupportsInterface(&_Neko.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Neko *NekoCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Neko.Contract.SupportsInterface(&_Neko.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Neko *NekoCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Neko *NekoSession) Symbol() (string, error) {
	return _Neko.Contract.Symbol(&_Neko.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Neko *NekoCallerSession) Symbol() (string, error) {
	return _Neko.Contract.Symbol(&_Neko.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Neko *NekoCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Neko *NekoSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Neko.Contract.TokenByIndex(&_Neko.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Neko *NekoCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Neko.Contract.TokenByIndex(&_Neko.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Neko *NekoCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Neko *NekoSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Neko.Contract.TokenOfOwnerByIndex(&_Neko.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Neko *NekoCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Neko.Contract.TokenOfOwnerByIndex(&_Neko.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Neko *NekoCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Neko *NekoSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Neko.Contract.TokenURI(&_Neko.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Neko *NekoCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Neko.Contract.TokenURI(&_Neko.CallOpts, _tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_Neko *NekoCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "tokensOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_Neko *NekoSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Neko.Contract.TokensOfOwner(&_Neko.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_Neko *NekoCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Neko.Contract.TokensOfOwner(&_Neko.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Neko *NekoCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Neko *NekoSession) TotalSupply() (*big.Int, error) {
	return _Neko.Contract.TotalSupply(&_Neko.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Neko *NekoCallerSession) TotalSupply() (*big.Int, error) {
	return _Neko.Contract.TotalSupply(&_Neko.CallOpts)
}

// Whitelister is a free data retrieval call binding the contract method 0x0c7253cf.
//
// Solidity: function whitelister(address ) view returns(bool)
func (_Neko *NekoCaller) Whitelister(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Neko.contract.Call(opts, &out, "whitelister", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelister is a free data retrieval call binding the contract method 0x0c7253cf.
//
// Solidity: function whitelister(address ) view returns(bool)
func (_Neko *NekoSession) Whitelister(arg0 common.Address) (bool, error) {
	return _Neko.Contract.Whitelister(&_Neko.CallOpts, arg0)
}

// Whitelister is a free data retrieval call binding the contract method 0x0c7253cf.
//
// Solidity: function whitelister(address ) view returns(bool)
func (_Neko *NekoCallerSession) Whitelister(arg0 common.Address) (bool, error) {
	return _Neko.Contract.Whitelister(&_Neko.CallOpts, arg0)
}

// AddWhitelists is a paid mutator transaction binding the contract method 0xc8eaf28f.
//
// Solidity: function addWhitelists(address[] addresses) returns()
func (_Neko *NekoTransactor) AddWhitelists(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "addWhitelists", addresses)
}

// AddWhitelists is a paid mutator transaction binding the contract method 0xc8eaf28f.
//
// Solidity: function addWhitelists(address[] addresses) returns()
func (_Neko *NekoSession) AddWhitelists(addresses []common.Address) (*types.Transaction, error) {
	return _Neko.Contract.AddWhitelists(&_Neko.TransactOpts, addresses)
}

// AddWhitelists is a paid mutator transaction binding the contract method 0xc8eaf28f.
//
// Solidity: function addWhitelists(address[] addresses) returns()
func (_Neko *NekoTransactorSession) AddWhitelists(addresses []common.Address) (*types.Transaction, error) {
	return _Neko.Contract.AddWhitelists(&_Neko.TransactOpts, addresses)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Neko *NekoTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Neko *NekoSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Neko.Contract.Approve(&_Neko.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Neko *NekoTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Neko.Contract.Approve(&_Neko.TransactOpts, to, tokenId)
}

// MintNeko is a paid mutator transaction binding the contract method 0x19d9d446.
//
// Solidity: function mintNeko(uint256 numNeko) payable returns()
func (_Neko *NekoTransactor) MintNeko(opts *bind.TransactOpts, numNeko *big.Int) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "mintNeko", numNeko)
}

// MintNeko is a paid mutator transaction binding the contract method 0x19d9d446.
//
// Solidity: function mintNeko(uint256 numNeko) payable returns()
func (_Neko *NekoSession) MintNeko(numNeko *big.Int) (*types.Transaction, error) {
	return _Neko.Contract.MintNeko(&_Neko.TransactOpts, numNeko)
}

// MintNeko is a paid mutator transaction binding the contract method 0x19d9d446.
//
// Solidity: function mintNeko(uint256 numNeko) payable returns()
func (_Neko *NekoTransactorSession) MintNeko(numNeko *big.Int) (*types.Transaction, error) {
	return _Neko.Contract.MintNeko(&_Neko.TransactOpts, numNeko)
}

// MintNekoByCustomToken is a paid mutator transaction binding the contract method 0x5cedb52f.
//
// Solidity: function mintNekoByCustomToken(uint256 numNeko, uint256 amount, uint8 tokenId) returns()
func (_Neko *NekoTransactor) MintNekoByCustomToken(opts *bind.TransactOpts, numNeko *big.Int, amount *big.Int, tokenId uint8) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "mintNekoByCustomToken", numNeko, amount, tokenId)
}

// MintNekoByCustomToken is a paid mutator transaction binding the contract method 0x5cedb52f.
//
// Solidity: function mintNekoByCustomToken(uint256 numNeko, uint256 amount, uint8 tokenId) returns()
func (_Neko *NekoSession) MintNekoByCustomToken(numNeko *big.Int, amount *big.Int, tokenId uint8) (*types.Transaction, error) {
	return _Neko.Contract.MintNekoByCustomToken(&_Neko.TransactOpts, numNeko, amount, tokenId)
}

// MintNekoByCustomToken is a paid mutator transaction binding the contract method 0x5cedb52f.
//
// Solidity: function mintNekoByCustomToken(uint256 numNeko, uint256 amount, uint8 tokenId) returns()
func (_Neko *NekoTransactorSession) MintNekoByCustomToken(numNeko *big.Int, amount *big.Int, tokenId uint8) (*types.Transaction, error) {
	return _Neko.Contract.MintNekoByCustomToken(&_Neko.TransactOpts, numNeko, amount, tokenId)
}

// MintWhitelist is a paid mutator transaction binding the contract method 0x2d3df31f.
//
// Solidity: function mintWhitelist() payable returns()
func (_Neko *NekoTransactor) MintWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "mintWhitelist")
}

// MintWhitelist is a paid mutator transaction binding the contract method 0x2d3df31f.
//
// Solidity: function mintWhitelist() payable returns()
func (_Neko *NekoSession) MintWhitelist() (*types.Transaction, error) {
	return _Neko.Contract.MintWhitelist(&_Neko.TransactOpts)
}

// MintWhitelist is a paid mutator transaction binding the contract method 0x2d3df31f.
//
// Solidity: function mintWhitelist() payable returns()
func (_Neko *NekoTransactorSession) MintWhitelist() (*types.Transaction, error) {
	return _Neko.Contract.MintWhitelist(&_Neko.TransactOpts)
}

// MintWhitelistByCustomToken is a paid mutator transaction binding the contract method 0x4ac5f9b2.
//
// Solidity: function mintWhitelistByCustomToken(uint256 amount, uint8 tokenId) returns()
func (_Neko *NekoTransactor) MintWhitelistByCustomToken(opts *bind.TransactOpts, amount *big.Int, tokenId uint8) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "mintWhitelistByCustomToken", amount, tokenId)
}

// MintWhitelistByCustomToken is a paid mutator transaction binding the contract method 0x4ac5f9b2.
//
// Solidity: function mintWhitelistByCustomToken(uint256 amount, uint8 tokenId) returns()
func (_Neko *NekoSession) MintWhitelistByCustomToken(amount *big.Int, tokenId uint8) (*types.Transaction, error) {
	return _Neko.Contract.MintWhitelistByCustomToken(&_Neko.TransactOpts, amount, tokenId)
}

// MintWhitelistByCustomToken is a paid mutator transaction binding the contract method 0x4ac5f9b2.
//
// Solidity: function mintWhitelistByCustomToken(uint256 amount, uint8 tokenId) returns()
func (_Neko *NekoTransactorSession) MintWhitelistByCustomToken(amount *big.Int, tokenId uint8) (*types.Transaction, error) {
	return _Neko.Contract.MintWhitelistByCustomToken(&_Neko.TransactOpts, amount, tokenId)
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_Neko *NekoTransactor) PauseSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "pauseSale")
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_Neko *NekoSession) PauseSale() (*types.Transaction, error) {
	return _Neko.Contract.PauseSale(&_Neko.TransactOpts)
}

// PauseSale is a paid mutator transaction binding the contract method 0x55367ba9.
//
// Solidity: function pauseSale() returns()
func (_Neko *NekoTransactorSession) PauseSale() (*types.Transaction, error) {
	return _Neko.Contract.PauseSale(&_Neko.TransactOpts)
}

// PauseWhitelistSale is a paid mutator transaction binding the contract method 0x0d0dcda5.
//
// Solidity: function pauseWhitelistSale() returns()
func (_Neko *NekoTransactor) PauseWhitelistSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "pauseWhitelistSale")
}

// PauseWhitelistSale is a paid mutator transaction binding the contract method 0x0d0dcda5.
//
// Solidity: function pauseWhitelistSale() returns()
func (_Neko *NekoSession) PauseWhitelistSale() (*types.Transaction, error) {
	return _Neko.Contract.PauseWhitelistSale(&_Neko.TransactOpts)
}

// PauseWhitelistSale is a paid mutator transaction binding the contract method 0x0d0dcda5.
//
// Solidity: function pauseWhitelistSale() returns()
func (_Neko *NekoTransactorSession) PauseWhitelistSale() (*types.Transaction, error) {
	return _Neko.Contract.PauseWhitelistSale(&_Neko.TransactOpts)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x23245216.
//
// Solidity: function removeWhitelist(address[] addresses) returns()
func (_Neko *NekoTransactor) RemoveWhitelist(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "removeWhitelist", addresses)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x23245216.
//
// Solidity: function removeWhitelist(address[] addresses) returns()
func (_Neko *NekoSession) RemoveWhitelist(addresses []common.Address) (*types.Transaction, error) {
	return _Neko.Contract.RemoveWhitelist(&_Neko.TransactOpts, addresses)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x23245216.
//
// Solidity: function removeWhitelist(address[] addresses) returns()
func (_Neko *NekoTransactorSession) RemoveWhitelist(addresses []common.Address) (*types.Transaction, error) {
	return _Neko.Contract.RemoveWhitelist(&_Neko.TransactOpts, addresses)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Neko *NekoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Neko *NekoSession) RenounceOwnership() (*types.Transaction, error) {
	return _Neko.Contract.RenounceOwnership(&_Neko.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Neko *NekoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Neko.Contract.RenounceOwnership(&_Neko.TransactOpts)
}

// ReserveGiveaway is a paid mutator transaction binding the contract method 0xa40f1aa5.
//
// Solidity: function reserveGiveaway(uint256 numNeko) returns()
func (_Neko *NekoTransactor) ReserveGiveaway(opts *bind.TransactOpts, numNeko *big.Int) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "reserveGiveaway", numNeko)
}

// ReserveGiveaway is a paid mutator transaction binding the contract method 0xa40f1aa5.
//
// Solidity: function reserveGiveaway(uint256 numNeko) returns()
func (_Neko *NekoSession) ReserveGiveaway(numNeko *big.Int) (*types.Transaction, error) {
	return _Neko.Contract.ReserveGiveaway(&_Neko.TransactOpts, numNeko)
}

// ReserveGiveaway is a paid mutator transaction binding the contract method 0xa40f1aa5.
//
// Solidity: function reserveGiveaway(uint256 numNeko) returns()
func (_Neko *NekoTransactorSession) ReserveGiveaway(numNeko *big.Int) (*types.Transaction, error) {
	return _Neko.Contract.ReserveGiveaway(&_Neko.TransactOpts, numNeko)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Neko *NekoTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Neko *NekoSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Neko.Contract.SafeTransferFrom(&_Neko.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Neko *NekoTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Neko.Contract.SafeTransferFrom(&_Neko.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Neko *NekoTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Neko *NekoSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Neko.Contract.SafeTransferFrom0(&_Neko.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Neko *NekoTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Neko.Contract.SafeTransferFrom0(&_Neko.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Neko *NekoTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Neko *NekoSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Neko.Contract.SetApprovalForAll(&_Neko.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Neko *NekoTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Neko.Contract.SetApprovalForAll(&_Neko.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Neko *NekoTransactor) SetBaseURI(opts *bind.TransactOpts, _baseURI string) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "setBaseURI", _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Neko *NekoSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _Neko.Contract.SetBaseURI(&_Neko.TransactOpts, _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Neko *NekoTransactorSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _Neko.Contract.SetBaseURI(&_Neko.TransactOpts, _baseURI)
}

// SetRevealer is a paid mutator transaction binding the contract method 0xc17c3fe3.
//
// Solidity: function setRevealer(address _revealer) returns()
func (_Neko *NekoTransactor) SetRevealer(opts *bind.TransactOpts, _revealer common.Address) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "setRevealer", _revealer)
}

// SetRevealer is a paid mutator transaction binding the contract method 0xc17c3fe3.
//
// Solidity: function setRevealer(address _revealer) returns()
func (_Neko *NekoSession) SetRevealer(_revealer common.Address) (*types.Transaction, error) {
	return _Neko.Contract.SetRevealer(&_Neko.TransactOpts, _revealer)
}

// SetRevealer is a paid mutator transaction binding the contract method 0xc17c3fe3.
//
// Solidity: function setRevealer(address _revealer) returns()
func (_Neko *NekoTransactorSession) SetRevealer(_revealer common.Address) (*types.Transaction, error) {
	return _Neko.Contract.SetRevealer(&_Neko.TransactOpts, _revealer)
}

// SetSupportedToken is a paid mutator transaction binding the contract method 0x91bfc1e2.
//
// Solidity: function setSupportedToken(uint8 tokenId, address tokenAddress, uint256 conversionRate, string symbol) returns()
func (_Neko *NekoTransactor) SetSupportedToken(opts *bind.TransactOpts, tokenId uint8, tokenAddress common.Address, conversionRate *big.Int, symbol string) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "setSupportedToken", tokenId, tokenAddress, conversionRate, symbol)
}

// SetSupportedToken is a paid mutator transaction binding the contract method 0x91bfc1e2.
//
// Solidity: function setSupportedToken(uint8 tokenId, address tokenAddress, uint256 conversionRate, string symbol) returns()
func (_Neko *NekoSession) SetSupportedToken(tokenId uint8, tokenAddress common.Address, conversionRate *big.Int, symbol string) (*types.Transaction, error) {
	return _Neko.Contract.SetSupportedToken(&_Neko.TransactOpts, tokenId, tokenAddress, conversionRate, symbol)
}

// SetSupportedToken is a paid mutator transaction binding the contract method 0x91bfc1e2.
//
// Solidity: function setSupportedToken(uint8 tokenId, address tokenAddress, uint256 conversionRate, string symbol) returns()
func (_Neko *NekoTransactorSession) SetSupportedToken(tokenId uint8, tokenAddress common.Address, conversionRate *big.Int, symbol string) (*types.Transaction, error) {
	return _Neko.Contract.SetSupportedToken(&_Neko.TransactOpts, tokenId, tokenAddress, conversionRate, symbol)
}

// StartSale is a paid mutator transaction binding the contract method 0xb66a0e5d.
//
// Solidity: function startSale() returns()
func (_Neko *NekoTransactor) StartSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "startSale")
}

// StartSale is a paid mutator transaction binding the contract method 0xb66a0e5d.
//
// Solidity: function startSale() returns()
func (_Neko *NekoSession) StartSale() (*types.Transaction, error) {
	return _Neko.Contract.StartSale(&_Neko.TransactOpts)
}

// StartSale is a paid mutator transaction binding the contract method 0xb66a0e5d.
//
// Solidity: function startSale() returns()
func (_Neko *NekoTransactorSession) StartSale() (*types.Transaction, error) {
	return _Neko.Contract.StartSale(&_Neko.TransactOpts)
}

// StartWhitelistSale is a paid mutator transaction binding the contract method 0xff44e915.
//
// Solidity: function startWhitelistSale() returns()
func (_Neko *NekoTransactor) StartWhitelistSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "startWhitelistSale")
}

// StartWhitelistSale is a paid mutator transaction binding the contract method 0xff44e915.
//
// Solidity: function startWhitelistSale() returns()
func (_Neko *NekoSession) StartWhitelistSale() (*types.Transaction, error) {
	return _Neko.Contract.StartWhitelistSale(&_Neko.TransactOpts)
}

// StartWhitelistSale is a paid mutator transaction binding the contract method 0xff44e915.
//
// Solidity: function startWhitelistSale() returns()
func (_Neko *NekoTransactorSession) StartWhitelistSale() (*types.Transaction, error) {
	return _Neko.Contract.StartWhitelistSale(&_Neko.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Neko *NekoTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Neko *NekoSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Neko.Contract.TransferFrom(&_Neko.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Neko *NekoTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Neko.Contract.TransferFrom(&_Neko.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Neko *NekoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Neko *NekoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Neko.Contract.TransferOwnership(&_Neko.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Neko *NekoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Neko.Contract.TransferOwnership(&_Neko.TransactOpts, newOwner)
}

// TurnOffSupportedToken is a paid mutator transaction binding the contract method 0x4a219da1.
//
// Solidity: function turnOffSupportedToken(uint8 tokenId) returns()
func (_Neko *NekoTransactor) TurnOffSupportedToken(opts *bind.TransactOpts, tokenId uint8) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "turnOffSupportedToken", tokenId)
}

// TurnOffSupportedToken is a paid mutator transaction binding the contract method 0x4a219da1.
//
// Solidity: function turnOffSupportedToken(uint8 tokenId) returns()
func (_Neko *NekoSession) TurnOffSupportedToken(tokenId uint8) (*types.Transaction, error) {
	return _Neko.Contract.TurnOffSupportedToken(&_Neko.TransactOpts, tokenId)
}

// TurnOffSupportedToken is a paid mutator transaction binding the contract method 0x4a219da1.
//
// Solidity: function turnOffSupportedToken(uint8 tokenId) returns()
func (_Neko *NekoTransactorSession) TurnOffSupportedToken(tokenId uint8) (*types.Transaction, error) {
	return _Neko.Contract.TurnOffSupportedToken(&_Neko.TransactOpts, tokenId)
}

// TurnOnSupportedToken is a paid mutator transaction binding the contract method 0xb79d7ca8.
//
// Solidity: function turnOnSupportedToken(uint8 tokenId) returns()
func (_Neko *NekoTransactor) TurnOnSupportedToken(opts *bind.TransactOpts, tokenId uint8) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "turnOnSupportedToken", tokenId)
}

// TurnOnSupportedToken is a paid mutator transaction binding the contract method 0xb79d7ca8.
//
// Solidity: function turnOnSupportedToken(uint8 tokenId) returns()
func (_Neko *NekoSession) TurnOnSupportedToken(tokenId uint8) (*types.Transaction, error) {
	return _Neko.Contract.TurnOnSupportedToken(&_Neko.TransactOpts, tokenId)
}

// TurnOnSupportedToken is a paid mutator transaction binding the contract method 0xb79d7ca8.
//
// Solidity: function turnOnSupportedToken(uint8 tokenId) returns()
func (_Neko *NekoTransactorSession) TurnOnSupportedToken(tokenId uint8) (*types.Transaction, error) {
	return _Neko.Contract.TurnOnSupportedToken(&_Neko.TransactOpts, tokenId)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() payable returns()
func (_Neko *NekoTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Neko.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() payable returns()
func (_Neko *NekoSession) WithdrawAll() (*types.Transaction, error) {
	return _Neko.Contract.WithdrawAll(&_Neko.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() payable returns()
func (_Neko *NekoTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Neko.Contract.WithdrawAll(&_Neko.TransactOpts)
}

// NekoApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Neko contract.
type NekoApprovalIterator struct {
	Event *NekoApproval // Event containing the contract specifics and raw log

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
func (it *NekoApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NekoApproval)
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
		it.Event = new(NekoApproval)
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
func (it *NekoApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NekoApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NekoApproval represents a Approval event raised by the Neko contract.
type NekoApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Neko *NekoFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NekoApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Neko.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NekoApprovalIterator{contract: _Neko.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Neko *NekoFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NekoApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Neko.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NekoApproval)
				if err := _Neko.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Neko *NekoFilterer) ParseApproval(log types.Log) (*NekoApproval, error) {
	event := new(NekoApproval)
	if err := _Neko.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NekoApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Neko contract.
type NekoApprovalForAllIterator struct {
	Event *NekoApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NekoApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NekoApprovalForAll)
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
		it.Event = new(NekoApprovalForAll)
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
func (it *NekoApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NekoApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NekoApprovalForAll represents a ApprovalForAll event raised by the Neko contract.
type NekoApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Neko *NekoFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NekoApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Neko.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NekoApprovalForAllIterator{contract: _Neko.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Neko *NekoFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NekoApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Neko.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NekoApprovalForAll)
				if err := _Neko.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Neko *NekoFilterer) ParseApprovalForAll(log types.Log) (*NekoApprovalForAll, error) {
	event := new(NekoApprovalForAll)
	if err := _Neko.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NekoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Neko contract.
type NekoOwnershipTransferredIterator struct {
	Event *NekoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NekoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NekoOwnershipTransferred)
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
		it.Event = new(NekoOwnershipTransferred)
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
func (it *NekoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NekoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NekoOwnershipTransferred represents a OwnershipTransferred event raised by the Neko contract.
type NekoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Neko *NekoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NekoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Neko.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NekoOwnershipTransferredIterator{contract: _Neko.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Neko *NekoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NekoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Neko.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NekoOwnershipTransferred)
				if err := _Neko.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Neko *NekoFilterer) ParseOwnershipTransferred(log types.Log) (*NekoOwnershipTransferred, error) {
	event := new(NekoOwnershipTransferred)
	if err := _Neko.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NekoTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Neko contract.
type NekoTransferIterator struct {
	Event *NekoTransfer // Event containing the contract specifics and raw log

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
func (it *NekoTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NekoTransfer)
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
		it.Event = new(NekoTransfer)
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
func (it *NekoTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NekoTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NekoTransfer represents a Transfer event raised by the Neko contract.
type NekoTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Neko *NekoFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NekoTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Neko.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NekoTransferIterator{contract: _Neko.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Neko *NekoFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NekoTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Neko.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NekoTransfer)
				if err := _Neko.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Neko *NekoFilterer) ParseTransfer(log types.Log) (*NekoTransfer, error) {
	event := new(NekoTransfer)
	if err := _Neko.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

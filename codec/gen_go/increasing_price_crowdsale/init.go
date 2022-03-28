// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package increasing_price_crowdsale

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

// IncreasingPriceCrowdsaleMetaData contains all meta data concerning the IncreasingPriceCrowdsale contract.
var IncreasingPriceCrowdsaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevClosingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newClosingTime\",\"type\":\"uint256\"}],\"name\":\"TimedCrowdsaleExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IncreasingPriceCrowdsaleABI is the input ABI used to generate the binding from.
// Deprecated: Use IncreasingPriceCrowdsaleMetaData.ABI instead.
var IncreasingPriceCrowdsaleABI = IncreasingPriceCrowdsaleMetaData.ABI

// IncreasingPriceCrowdsale is an auto generated Go binding around an Ethereum contract.
type IncreasingPriceCrowdsale struct {
	IncreasingPriceCrowdsaleCaller     // Read-only binding to the contract
	IncreasingPriceCrowdsaleTransactor // Write-only binding to the contract
	IncreasingPriceCrowdsaleFilterer   // Log filterer for contract events
}

// IncreasingPriceCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncreasingPriceCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncreasingPriceCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncreasingPriceCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncreasingPriceCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncreasingPriceCrowdsaleSession struct {
	Contract     *IncreasingPriceCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IncreasingPriceCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncreasingPriceCrowdsaleCallerSession struct {
	Contract *IncreasingPriceCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IncreasingPriceCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncreasingPriceCrowdsaleTransactorSession struct {
	Contract     *IncreasingPriceCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IncreasingPriceCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleRaw struct {
	Contract *IncreasingPriceCrowdsale // Generic contract binding to access the raw methods on
}

// IncreasingPriceCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleCallerRaw struct {
	Contract *IncreasingPriceCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// IncreasingPriceCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncreasingPriceCrowdsaleTransactorRaw struct {
	Contract *IncreasingPriceCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncreasingPriceCrowdsale creates a new instance of IncreasingPriceCrowdsale, bound to a specific deployed contract.
func NewIncreasingPriceCrowdsale(address common.Address, backend bind.ContractBackend) (*IncreasingPriceCrowdsale, error) {
	contract, err := bindIncreasingPriceCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IncreasingPriceCrowdsale{IncreasingPriceCrowdsaleCaller: IncreasingPriceCrowdsaleCaller{contract: contract}, IncreasingPriceCrowdsaleTransactor: IncreasingPriceCrowdsaleTransactor{contract: contract}, IncreasingPriceCrowdsaleFilterer: IncreasingPriceCrowdsaleFilterer{contract: contract}}, nil
}

// NewIncreasingPriceCrowdsaleCaller creates a new read-only instance of IncreasingPriceCrowdsale, bound to a specific deployed contract.
func NewIncreasingPriceCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*IncreasingPriceCrowdsaleCaller, error) {
	contract, err := bindIncreasingPriceCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncreasingPriceCrowdsaleCaller{contract: contract}, nil
}

// NewIncreasingPriceCrowdsaleTransactor creates a new write-only instance of IncreasingPriceCrowdsale, bound to a specific deployed contract.
func NewIncreasingPriceCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*IncreasingPriceCrowdsaleTransactor, error) {
	contract, err := bindIncreasingPriceCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncreasingPriceCrowdsaleTransactor{contract: contract}, nil
}

// NewIncreasingPriceCrowdsaleFilterer creates a new log filterer instance of IncreasingPriceCrowdsale, bound to a specific deployed contract.
func NewIncreasingPriceCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*IncreasingPriceCrowdsaleFilterer, error) {
	contract, err := bindIncreasingPriceCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncreasingPriceCrowdsaleFilterer{contract: contract}, nil
}

// bindIncreasingPriceCrowdsale binds a generic wrapper to an already deployed contract.
func bindIncreasingPriceCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncreasingPriceCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IncreasingPriceCrowdsale.Contract.IncreasingPriceCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.Contract.IncreasingPriceCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.Contract.IncreasingPriceCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IncreasingPriceCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IncreasingPriceCrowdsale.contract.Call(opts, &out, "closingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.ClosingTime(&_IncreasingPriceCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.ClosingTime(&_IncreasingPriceCrowdsale.CallOpts)
}

// FinalRate is a free data retrieval call binding the contract method 0x21106109.
//
// Solidity: function finalRate() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) FinalRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IncreasingPriceCrowdsale.contract.Call(opts, &out, "finalRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FinalRate is a free data retrieval call binding the contract method 0x21106109.
//
// Solidity: function finalRate() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) FinalRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.FinalRate(&_IncreasingPriceCrowdsale.CallOpts)
}

// FinalRate is a free data retrieval call binding the contract method 0x21106109.
//
// Solidity: function finalRate() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) FinalRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.FinalRate(&_IncreasingPriceCrowdsale.CallOpts)
}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) GetCurrentRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IncreasingPriceCrowdsale.contract.Call(opts, &out, "getCurrentRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) GetCurrentRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.GetCurrentRate(&_IncreasingPriceCrowdsale.CallOpts)
}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) GetCurrentRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.GetCurrentRate(&_IncreasingPriceCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IncreasingPriceCrowdsale.contract.Call(opts, &out, "hasClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) HasClosed() (bool, error) {
	return _IncreasingPriceCrowdsale.Contract.HasClosed(&_IncreasingPriceCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _IncreasingPriceCrowdsale.Contract.HasClosed(&_IncreasingPriceCrowdsale.CallOpts)
}

// InitialRate is a free data retrieval call binding the contract method 0x9e51051f.
//
// Solidity: function initialRate() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) InitialRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IncreasingPriceCrowdsale.contract.Call(opts, &out, "initialRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialRate is a free data retrieval call binding the contract method 0x9e51051f.
//
// Solidity: function initialRate() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) InitialRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.InitialRate(&_IncreasingPriceCrowdsale.CallOpts)
}

// InitialRate is a free data retrieval call binding the contract method 0x9e51051f.
//
// Solidity: function initialRate() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) InitialRate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.InitialRate(&_IncreasingPriceCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IncreasingPriceCrowdsale.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) IsOpen() (bool, error) {
	return _IncreasingPriceCrowdsale.Contract.IsOpen(&_IncreasingPriceCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) IsOpen() (bool, error) {
	return _IncreasingPriceCrowdsale.Contract.IsOpen(&_IncreasingPriceCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IncreasingPriceCrowdsale.contract.Call(opts, &out, "openingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.OpeningTime(&_IncreasingPriceCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.OpeningTime(&_IncreasingPriceCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IncreasingPriceCrowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) Rate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.Rate(&_IncreasingPriceCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.Rate(&_IncreasingPriceCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IncreasingPriceCrowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) Token() (common.Address, error) {
	return _IncreasingPriceCrowdsale.Contract.Token(&_IncreasingPriceCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) Token() (common.Address, error) {
	return _IncreasingPriceCrowdsale.Contract.Token(&_IncreasingPriceCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IncreasingPriceCrowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) Wallet() (common.Address, error) {
	return _IncreasingPriceCrowdsale.Contract.Wallet(&_IncreasingPriceCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _IncreasingPriceCrowdsale.Contract.Wallet(&_IncreasingPriceCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IncreasingPriceCrowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.WeiRaised(&_IncreasingPriceCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _IncreasingPriceCrowdsale.Contract.WeiRaised(&_IncreasingPriceCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.Contract.BuyTokens(&_IncreasingPriceCrowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.Contract.BuyTokens(&_IncreasingPriceCrowdsale.TransactOpts, beneficiary)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.Contract.Fallback(&_IncreasingPriceCrowdsale.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _IncreasingPriceCrowdsale.Contract.Fallback(&_IncreasingPriceCrowdsale.TransactOpts, calldata)
}

// IncreasingPriceCrowdsaleTimedCrowdsaleExtendedIterator is returned from FilterTimedCrowdsaleExtended and is used to iterate over the raw logs and unpacked data for TimedCrowdsaleExtended events raised by the IncreasingPriceCrowdsale contract.
type IncreasingPriceCrowdsaleTimedCrowdsaleExtendedIterator struct {
	Event *IncreasingPriceCrowdsaleTimedCrowdsaleExtended // Event containing the contract specifics and raw log

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
func (it *IncreasingPriceCrowdsaleTimedCrowdsaleExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncreasingPriceCrowdsaleTimedCrowdsaleExtended)
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
		it.Event = new(IncreasingPriceCrowdsaleTimedCrowdsaleExtended)
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
func (it *IncreasingPriceCrowdsaleTimedCrowdsaleExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncreasingPriceCrowdsaleTimedCrowdsaleExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncreasingPriceCrowdsaleTimedCrowdsaleExtended represents a TimedCrowdsaleExtended event raised by the IncreasingPriceCrowdsale contract.
type IncreasingPriceCrowdsaleTimedCrowdsaleExtended struct {
	PrevClosingTime *big.Int
	NewClosingTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedCrowdsaleExtended is a free log retrieval operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleFilterer) FilterTimedCrowdsaleExtended(opts *bind.FilterOpts) (*IncreasingPriceCrowdsaleTimedCrowdsaleExtendedIterator, error) {

	logs, sub, err := _IncreasingPriceCrowdsale.contract.FilterLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return &IncreasingPriceCrowdsaleTimedCrowdsaleExtendedIterator{contract: _IncreasingPriceCrowdsale.contract, event: "TimedCrowdsaleExtended", logs: logs, sub: sub}, nil
}

// WatchTimedCrowdsaleExtended is a free log subscription operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleFilterer) WatchTimedCrowdsaleExtended(opts *bind.WatchOpts, sink chan<- *IncreasingPriceCrowdsaleTimedCrowdsaleExtended) (event.Subscription, error) {

	logs, sub, err := _IncreasingPriceCrowdsale.contract.WatchLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncreasingPriceCrowdsaleTimedCrowdsaleExtended)
				if err := _IncreasingPriceCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
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

// ParseTimedCrowdsaleExtended is a log parse operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleFilterer) ParseTimedCrowdsaleExtended(log types.Log) (*IncreasingPriceCrowdsaleTimedCrowdsaleExtended, error) {
	event := new(IncreasingPriceCrowdsaleTimedCrowdsaleExtended)
	if err := _IncreasingPriceCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncreasingPriceCrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the IncreasingPriceCrowdsale contract.
type IncreasingPriceCrowdsaleTokensPurchasedIterator struct {
	Event *IncreasingPriceCrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *IncreasingPriceCrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncreasingPriceCrowdsaleTokensPurchased)
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
		it.Event = new(IncreasingPriceCrowdsaleTokensPurchased)
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
func (it *IncreasingPriceCrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncreasingPriceCrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncreasingPriceCrowdsaleTokensPurchased represents a TokensPurchased event raised by the IncreasingPriceCrowdsale contract.
type IncreasingPriceCrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*IncreasingPriceCrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _IncreasingPriceCrowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &IncreasingPriceCrowdsaleTokensPurchasedIterator{contract: _IncreasingPriceCrowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *IncreasingPriceCrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _IncreasingPriceCrowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncreasingPriceCrowdsaleTokensPurchased)
				if err := _IncreasingPriceCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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

// ParseTokensPurchased is a log parse operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_IncreasingPriceCrowdsale *IncreasingPriceCrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*IncreasingPriceCrowdsaleTokensPurchased, error) {
	event := new(IncreasingPriceCrowdsaleTokensPurchased)
	if err := _IncreasingPriceCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

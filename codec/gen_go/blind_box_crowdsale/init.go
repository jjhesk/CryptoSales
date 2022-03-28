// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blind_box_crowdsale

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

// BlindBoxCrowdsaleMetaData contains all meta data concerning the BlindBoxCrowdsale contract.
var BlindBoxCrowdsaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevClosingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newClosingTime\",\"type\":\"uint256\"}],\"name\":\"TimedCrowdsaleExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ceilingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"floorRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxRandom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"privateAddress\",\"type\":\"address\"}],\"name\":\"getRandomNumber\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BlindBoxCrowdsaleABI is the input ABI used to generate the binding from.
// Deprecated: Use BlindBoxCrowdsaleMetaData.ABI instead.
var BlindBoxCrowdsaleABI = BlindBoxCrowdsaleMetaData.ABI

// BlindBoxCrowdsale is an auto generated Go binding around an Ethereum contract.
type BlindBoxCrowdsale struct {
	BlindBoxCrowdsaleCaller     // Read-only binding to the contract
	BlindBoxCrowdsaleTransactor // Write-only binding to the contract
	BlindBoxCrowdsaleFilterer   // Log filterer for contract events
}

// BlindBoxCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlindBoxCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindBoxCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlindBoxCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindBoxCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlindBoxCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindBoxCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlindBoxCrowdsaleSession struct {
	Contract     *BlindBoxCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BlindBoxCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlindBoxCrowdsaleCallerSession struct {
	Contract *BlindBoxCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BlindBoxCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlindBoxCrowdsaleTransactorSession struct {
	Contract     *BlindBoxCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BlindBoxCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlindBoxCrowdsaleRaw struct {
	Contract *BlindBoxCrowdsale // Generic contract binding to access the raw methods on
}

// BlindBoxCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlindBoxCrowdsaleCallerRaw struct {
	Contract *BlindBoxCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// BlindBoxCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlindBoxCrowdsaleTransactorRaw struct {
	Contract *BlindBoxCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlindBoxCrowdsale creates a new instance of BlindBoxCrowdsale, bound to a specific deployed contract.
func NewBlindBoxCrowdsale(address common.Address, backend bind.ContractBackend) (*BlindBoxCrowdsale, error) {
	contract, err := bindBlindBoxCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlindBoxCrowdsale{BlindBoxCrowdsaleCaller: BlindBoxCrowdsaleCaller{contract: contract}, BlindBoxCrowdsaleTransactor: BlindBoxCrowdsaleTransactor{contract: contract}, BlindBoxCrowdsaleFilterer: BlindBoxCrowdsaleFilterer{contract: contract}}, nil
}

// NewBlindBoxCrowdsaleCaller creates a new read-only instance of BlindBoxCrowdsale, bound to a specific deployed contract.
func NewBlindBoxCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*BlindBoxCrowdsaleCaller, error) {
	contract, err := bindBlindBoxCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlindBoxCrowdsaleCaller{contract: contract}, nil
}

// NewBlindBoxCrowdsaleTransactor creates a new write-only instance of BlindBoxCrowdsale, bound to a specific deployed contract.
func NewBlindBoxCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*BlindBoxCrowdsaleTransactor, error) {
	contract, err := bindBlindBoxCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlindBoxCrowdsaleTransactor{contract: contract}, nil
}

// NewBlindBoxCrowdsaleFilterer creates a new log filterer instance of BlindBoxCrowdsale, bound to a specific deployed contract.
func NewBlindBoxCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*BlindBoxCrowdsaleFilterer, error) {
	contract, err := bindBlindBoxCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlindBoxCrowdsaleFilterer{contract: contract}, nil
}

// bindBlindBoxCrowdsale binds a generic wrapper to an already deployed contract.
func bindBlindBoxCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlindBoxCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlindBoxCrowdsale.Contract.BlindBoxCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindBoxCrowdsale.Contract.BlindBoxCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlindBoxCrowdsale.Contract.BlindBoxCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlindBoxCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindBoxCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlindBoxCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// CeilingRate is a free data retrieval call binding the contract method 0xebd9fc4b.
//
// Solidity: function ceilingRate() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCaller) CeilingRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlindBoxCrowdsale.contract.Call(opts, &out, "ceilingRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CeilingRate is a free data retrieval call binding the contract method 0xebd9fc4b.
//
// Solidity: function ceilingRate() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleSession) CeilingRate() (*big.Int, error) {
	return _BlindBoxCrowdsale.Contract.CeilingRate(&_BlindBoxCrowdsale.CallOpts)
}

// CeilingRate is a free data retrieval call binding the contract method 0xebd9fc4b.
//
// Solidity: function ceilingRate() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCallerSession) CeilingRate() (*big.Int, error) {
	return _BlindBoxCrowdsale.Contract.CeilingRate(&_BlindBoxCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlindBoxCrowdsale.contract.Call(opts, &out, "closingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _BlindBoxCrowdsale.Contract.ClosingTime(&_BlindBoxCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _BlindBoxCrowdsale.Contract.ClosingTime(&_BlindBoxCrowdsale.CallOpts)
}

// FloorRate is a free data retrieval call binding the contract method 0xb371e4ce.
//
// Solidity: function floorRate() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCaller) FloorRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlindBoxCrowdsale.contract.Call(opts, &out, "floorRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FloorRate is a free data retrieval call binding the contract method 0xb371e4ce.
//
// Solidity: function floorRate() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleSession) FloorRate() (*big.Int, error) {
	return _BlindBoxCrowdsale.Contract.FloorRate(&_BlindBoxCrowdsale.CallOpts)
}

// FloorRate is a free data retrieval call binding the contract method 0xb371e4ce.
//
// Solidity: function floorRate() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCallerSession) FloorRate() (*big.Int, error) {
	return _BlindBoxCrowdsale.Contract.FloorRate(&_BlindBoxCrowdsale.CallOpts)
}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCaller) GetCurrentRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlindBoxCrowdsale.contract.Call(opts, &out, "getCurrentRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleSession) GetCurrentRate() (*big.Int, error) {
	return _BlindBoxCrowdsale.Contract.GetCurrentRate(&_BlindBoxCrowdsale.CallOpts)
}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCallerSession) GetCurrentRate() (*big.Int, error) {
	return _BlindBoxCrowdsale.Contract.GetCurrentRate(&_BlindBoxCrowdsale.CallOpts)
}

// GetRandomNumber is a free data retrieval call binding the contract method 0x00ef33c2.
//
// Solidity: function getRandomNumber(uint256 maxRandom, uint256 min, address privateAddress) view returns(uint8)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCaller) GetRandomNumber(opts *bind.CallOpts, maxRandom *big.Int, min *big.Int, privateAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _BlindBoxCrowdsale.contract.Call(opts, &out, "getRandomNumber", maxRandom, min, privateAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetRandomNumber is a free data retrieval call binding the contract method 0x00ef33c2.
//
// Solidity: function getRandomNumber(uint256 maxRandom, uint256 min, address privateAddress) view returns(uint8)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleSession) GetRandomNumber(maxRandom *big.Int, min *big.Int, privateAddress common.Address) (uint8, error) {
	return _BlindBoxCrowdsale.Contract.GetRandomNumber(&_BlindBoxCrowdsale.CallOpts, maxRandom, min, privateAddress)
}

// GetRandomNumber is a free data retrieval call binding the contract method 0x00ef33c2.
//
// Solidity: function getRandomNumber(uint256 maxRandom, uint256 min, address privateAddress) view returns(uint8)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCallerSession) GetRandomNumber(maxRandom *big.Int, min *big.Int, privateAddress common.Address) (uint8, error) {
	return _BlindBoxCrowdsale.Contract.GetRandomNumber(&_BlindBoxCrowdsale.CallOpts, maxRandom, min, privateAddress)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BlindBoxCrowdsale.contract.Call(opts, &out, "hasClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleSession) HasClosed() (bool, error) {
	return _BlindBoxCrowdsale.Contract.HasClosed(&_BlindBoxCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _BlindBoxCrowdsale.Contract.HasClosed(&_BlindBoxCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BlindBoxCrowdsale.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleSession) IsOpen() (bool, error) {
	return _BlindBoxCrowdsale.Contract.IsOpen(&_BlindBoxCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCallerSession) IsOpen() (bool, error) {
	return _BlindBoxCrowdsale.Contract.IsOpen(&_BlindBoxCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlindBoxCrowdsale.contract.Call(opts, &out, "openingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _BlindBoxCrowdsale.Contract.OpeningTime(&_BlindBoxCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _BlindBoxCrowdsale.Contract.OpeningTime(&_BlindBoxCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlindBoxCrowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleSession) Rate() (*big.Int, error) {
	return _BlindBoxCrowdsale.Contract.Rate(&_BlindBoxCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _BlindBoxCrowdsale.Contract.Rate(&_BlindBoxCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlindBoxCrowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleSession) Token() (common.Address, error) {
	return _BlindBoxCrowdsale.Contract.Token(&_BlindBoxCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCallerSession) Token() (common.Address, error) {
	return _BlindBoxCrowdsale.Contract.Token(&_BlindBoxCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlindBoxCrowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleSession) Wallet() (common.Address, error) {
	return _BlindBoxCrowdsale.Contract.Wallet(&_BlindBoxCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _BlindBoxCrowdsale.Contract.Wallet(&_BlindBoxCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlindBoxCrowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _BlindBoxCrowdsale.Contract.WeiRaised(&_BlindBoxCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _BlindBoxCrowdsale.Contract.WeiRaised(&_BlindBoxCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _BlindBoxCrowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _BlindBoxCrowdsale.Contract.BuyTokens(&_BlindBoxCrowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _BlindBoxCrowdsale.Contract.BuyTokens(&_BlindBoxCrowdsale.TransactOpts, beneficiary)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BlindBoxCrowdsale.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BlindBoxCrowdsale.Contract.Fallback(&_BlindBoxCrowdsale.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BlindBoxCrowdsale.Contract.Fallback(&_BlindBoxCrowdsale.TransactOpts, calldata)
}

// BlindBoxCrowdsaleTimedCrowdsaleExtendedIterator is returned from FilterTimedCrowdsaleExtended and is used to iterate over the raw logs and unpacked data for TimedCrowdsaleExtended events raised by the BlindBoxCrowdsale contract.
type BlindBoxCrowdsaleTimedCrowdsaleExtendedIterator struct {
	Event *BlindBoxCrowdsaleTimedCrowdsaleExtended // Event containing the contract specifics and raw log

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
func (it *BlindBoxCrowdsaleTimedCrowdsaleExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindBoxCrowdsaleTimedCrowdsaleExtended)
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
		it.Event = new(BlindBoxCrowdsaleTimedCrowdsaleExtended)
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
func (it *BlindBoxCrowdsaleTimedCrowdsaleExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindBoxCrowdsaleTimedCrowdsaleExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindBoxCrowdsaleTimedCrowdsaleExtended represents a TimedCrowdsaleExtended event raised by the BlindBoxCrowdsale contract.
type BlindBoxCrowdsaleTimedCrowdsaleExtended struct {
	PrevClosingTime *big.Int
	NewClosingTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedCrowdsaleExtended is a free log retrieval operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleFilterer) FilterTimedCrowdsaleExtended(opts *bind.FilterOpts) (*BlindBoxCrowdsaleTimedCrowdsaleExtendedIterator, error) {

	logs, sub, err := _BlindBoxCrowdsale.contract.FilterLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return &BlindBoxCrowdsaleTimedCrowdsaleExtendedIterator{contract: _BlindBoxCrowdsale.contract, event: "TimedCrowdsaleExtended", logs: logs, sub: sub}, nil
}

// WatchTimedCrowdsaleExtended is a free log subscription operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleFilterer) WatchTimedCrowdsaleExtended(opts *bind.WatchOpts, sink chan<- *BlindBoxCrowdsaleTimedCrowdsaleExtended) (event.Subscription, error) {

	logs, sub, err := _BlindBoxCrowdsale.contract.WatchLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindBoxCrowdsaleTimedCrowdsaleExtended)
				if err := _BlindBoxCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
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
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleFilterer) ParseTimedCrowdsaleExtended(log types.Log) (*BlindBoxCrowdsaleTimedCrowdsaleExtended, error) {
	event := new(BlindBoxCrowdsaleTimedCrowdsaleExtended)
	if err := _BlindBoxCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindBoxCrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the BlindBoxCrowdsale contract.
type BlindBoxCrowdsaleTokensPurchasedIterator struct {
	Event *BlindBoxCrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *BlindBoxCrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindBoxCrowdsaleTokensPurchased)
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
		it.Event = new(BlindBoxCrowdsaleTokensPurchased)
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
func (it *BlindBoxCrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindBoxCrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindBoxCrowdsaleTokensPurchased represents a TokensPurchased event raised by the BlindBoxCrowdsale contract.
type BlindBoxCrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*BlindBoxCrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _BlindBoxCrowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &BlindBoxCrowdsaleTokensPurchasedIterator{contract: _BlindBoxCrowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *BlindBoxCrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _BlindBoxCrowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindBoxCrowdsaleTokensPurchased)
				if err := _BlindBoxCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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
func (_BlindBoxCrowdsale *BlindBoxCrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*BlindBoxCrowdsaleTokensPurchased, error) {
	event := new(BlindBoxCrowdsaleTokensPurchased)
	if err := _BlindBoxCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

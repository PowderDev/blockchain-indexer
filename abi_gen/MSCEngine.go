// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi_gen

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
	_ = abi.ConvertType
)

// MSCEngineMetaData contains all meta data concerning the MSCEngine contract.
var MSCEngineMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"priceFeeds\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"mscAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CollateralDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CollateralRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DebtBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DebtMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Liquidated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"depositAndMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getAccountCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateralAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenAmountForUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getUSDValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getUserCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtToCover\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeemCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"redeemCollateralForMSC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MSCEngineABI is the input ABI used to generate the binding from.
// Deprecated: Use MSCEngineMetaData.ABI instead.
var MSCEngineABI = MSCEngineMetaData.ABI

// MSCEngine is an auto generated Go binding around an Ethereum contract.
type MSCEngine struct {
	MSCEngineCaller     // Read-only binding to the contract
	MSCEngineTransactor // Write-only binding to the contract
	MSCEngineFilterer   // Log filterer for contract events
}

// MSCEngineCaller is an auto generated read-only Go binding around an Ethereum contract.
type MSCEngineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MSCEngineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MSCEngineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MSCEngineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MSCEngineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MSCEngineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MSCEngineSession struct {
	Contract     *MSCEngine        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MSCEngineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MSCEngineCallerSession struct {
	Contract *MSCEngineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MSCEngineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MSCEngineTransactorSession struct {
	Contract     *MSCEngineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MSCEngineRaw is an auto generated low-level Go binding around an Ethereum contract.
type MSCEngineRaw struct {
	Contract *MSCEngine // Generic contract binding to access the raw methods on
}

// MSCEngineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MSCEngineCallerRaw struct {
	Contract *MSCEngineCaller // Generic read-only contract binding to access the raw methods on
}

// MSCEngineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MSCEngineTransactorRaw struct {
	Contract *MSCEngineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMSCEngine creates a new instance of MSCEngine, bound to a specific deployed contract.
func NewMSCEngine(address common.Address, backend bind.ContractBackend) (*MSCEngine, error) {
	contract, err := bindMSCEngine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MSCEngine{MSCEngineCaller: MSCEngineCaller{contract: contract}, MSCEngineTransactor: MSCEngineTransactor{contract: contract}, MSCEngineFilterer: MSCEngineFilterer{contract: contract}}, nil
}

// NewMSCEngineCaller creates a new read-only instance of MSCEngine, bound to a specific deployed contract.
func NewMSCEngineCaller(address common.Address, caller bind.ContractCaller) (*MSCEngineCaller, error) {
	contract, err := bindMSCEngine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MSCEngineCaller{contract: contract}, nil
}

// NewMSCEngineTransactor creates a new write-only instance of MSCEngine, bound to a specific deployed contract.
func NewMSCEngineTransactor(address common.Address, transactor bind.ContractTransactor) (*MSCEngineTransactor, error) {
	contract, err := bindMSCEngine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MSCEngineTransactor{contract: contract}, nil
}

// NewMSCEngineFilterer creates a new log filterer instance of MSCEngine, bound to a specific deployed contract.
func NewMSCEngineFilterer(address common.Address, filterer bind.ContractFilterer) (*MSCEngineFilterer, error) {
	contract, err := bindMSCEngine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MSCEngineFilterer{contract: contract}, nil
}

// bindMSCEngine binds a generic wrapper to an already deployed contract.
func bindMSCEngine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MSCEngineMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MSCEngine *MSCEngineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MSCEngine.Contract.MSCEngineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MSCEngine *MSCEngineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MSCEngine.Contract.MSCEngineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MSCEngine *MSCEngineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MSCEngine.Contract.MSCEngineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MSCEngine *MSCEngineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MSCEngine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MSCEngine *MSCEngineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MSCEngine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MSCEngine *MSCEngineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MSCEngine.Contract.contract.Transact(opts, method, params...)
}

// GetAccountCollateral is a free data retrieval call binding the contract method 0x5a1ddb78.
//
// Solidity: function getAccountCollateral(address user) view returns(uint256 totalValue)
func (_MSCEngine *MSCEngineCaller) GetAccountCollateral(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MSCEngine.contract.Call(opts, &out, "getAccountCollateral", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountCollateral is a free data retrieval call binding the contract method 0x5a1ddb78.
//
// Solidity: function getAccountCollateral(address user) view returns(uint256 totalValue)
func (_MSCEngine *MSCEngineSession) GetAccountCollateral(user common.Address) (*big.Int, error) {
	return _MSCEngine.Contract.GetAccountCollateral(&_MSCEngine.CallOpts, user)
}

// GetAccountCollateral is a free data retrieval call binding the contract method 0x5a1ddb78.
//
// Solidity: function getAccountCollateral(address user) view returns(uint256 totalValue)
func (_MSCEngine *MSCEngineCallerSession) GetAccountCollateral(user common.Address) (*big.Int, error) {
	return _MSCEngine.Contract.GetAccountCollateral(&_MSCEngine.CallOpts, user)
}

// GetCollateralAddresses is a free data retrieval call binding the contract method 0x1834f2a4.
//
// Solidity: function getCollateralAddresses() view returns(address[])
func (_MSCEngine *MSCEngineCaller) GetCollateralAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MSCEngine.contract.Call(opts, &out, "getCollateralAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCollateralAddresses is a free data retrieval call binding the contract method 0x1834f2a4.
//
// Solidity: function getCollateralAddresses() view returns(address[])
func (_MSCEngine *MSCEngineSession) GetCollateralAddresses() ([]common.Address, error) {
	return _MSCEngine.Contract.GetCollateralAddresses(&_MSCEngine.CallOpts)
}

// GetCollateralAddresses is a free data retrieval call binding the contract method 0x1834f2a4.
//
// Solidity: function getCollateralAddresses() view returns(address[])
func (_MSCEngine *MSCEngineCallerSession) GetCollateralAddresses() ([]common.Address, error) {
	return _MSCEngine.Contract.GetCollateralAddresses(&_MSCEngine.CallOpts)
}

// GetTokenAmountForUSD is a free data retrieval call binding the contract method 0x968534bf.
//
// Solidity: function getTokenAmountForUSD(address token, uint256 usdAmount) view returns(uint256)
func (_MSCEngine *MSCEngineCaller) GetTokenAmountForUSD(opts *bind.CallOpts, token common.Address, usdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MSCEngine.contract.Call(opts, &out, "getTokenAmountForUSD", token, usdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenAmountForUSD is a free data retrieval call binding the contract method 0x968534bf.
//
// Solidity: function getTokenAmountForUSD(address token, uint256 usdAmount) view returns(uint256)
func (_MSCEngine *MSCEngineSession) GetTokenAmountForUSD(token common.Address, usdAmount *big.Int) (*big.Int, error) {
	return _MSCEngine.Contract.GetTokenAmountForUSD(&_MSCEngine.CallOpts, token, usdAmount)
}

// GetTokenAmountForUSD is a free data retrieval call binding the contract method 0x968534bf.
//
// Solidity: function getTokenAmountForUSD(address token, uint256 usdAmount) view returns(uint256)
func (_MSCEngine *MSCEngineCallerSession) GetTokenAmountForUSD(token common.Address, usdAmount *big.Int) (*big.Int, error) {
	return _MSCEngine.Contract.GetTokenAmountForUSD(&_MSCEngine.CallOpts, token, usdAmount)
}

// GetUSDValue is a free data retrieval call binding the contract method 0xfa76dcf2.
//
// Solidity: function getUSDValue(address token, uint256 amount) view returns(uint256)
func (_MSCEngine *MSCEngineCaller) GetUSDValue(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MSCEngine.contract.Call(opts, &out, "getUSDValue", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUSDValue is a free data retrieval call binding the contract method 0xfa76dcf2.
//
// Solidity: function getUSDValue(address token, uint256 amount) view returns(uint256)
func (_MSCEngine *MSCEngineSession) GetUSDValue(token common.Address, amount *big.Int) (*big.Int, error) {
	return _MSCEngine.Contract.GetUSDValue(&_MSCEngine.CallOpts, token, amount)
}

// GetUSDValue is a free data retrieval call binding the contract method 0xfa76dcf2.
//
// Solidity: function getUSDValue(address token, uint256 amount) view returns(uint256)
func (_MSCEngine *MSCEngineCallerSession) GetUSDValue(token common.Address, amount *big.Int) (*big.Int, error) {
	return _MSCEngine.Contract.GetUSDValue(&_MSCEngine.CallOpts, token, amount)
}

// GetUserCollateral is a free data retrieval call binding the contract method 0xa55e8668.
//
// Solidity: function getUserCollateral(address user, address token) view returns(uint256)
func (_MSCEngine *MSCEngineCaller) GetUserCollateral(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MSCEngine.contract.Call(opts, &out, "getUserCollateral", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserCollateral is a free data retrieval call binding the contract method 0xa55e8668.
//
// Solidity: function getUserCollateral(address user, address token) view returns(uint256)
func (_MSCEngine *MSCEngineSession) GetUserCollateral(user common.Address, token common.Address) (*big.Int, error) {
	return _MSCEngine.Contract.GetUserCollateral(&_MSCEngine.CallOpts, user, token)
}

// GetUserCollateral is a free data retrieval call binding the contract method 0xa55e8668.
//
// Solidity: function getUserCollateral(address user, address token) view returns(uint256)
func (_MSCEngine *MSCEngineCallerSession) GetUserCollateral(user common.Address, token common.Address) (*big.Int, error) {
	return _MSCEngine.Contract.GetUserCollateral(&_MSCEngine.CallOpts, user, token)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address user) view returns(uint256 collateralValue, uint256 debtValue)
func (_MSCEngine *MSCEngineCaller) GetUserInfo(opts *bind.CallOpts, user common.Address) (struct {
	CollateralValue *big.Int
	DebtValue       *big.Int
}, error) {
	var out []interface{}
	err := _MSCEngine.contract.Call(opts, &out, "getUserInfo", user)

	outstruct := new(struct {
		CollateralValue *big.Int
		DebtValue       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CollateralValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DebtValue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address user) view returns(uint256 collateralValue, uint256 debtValue)
func (_MSCEngine *MSCEngineSession) GetUserInfo(user common.Address) (struct {
	CollateralValue *big.Int
	DebtValue       *big.Int
}, error) {
	return _MSCEngine.Contract.GetUserInfo(&_MSCEngine.CallOpts, user)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address user) view returns(uint256 collateralValue, uint256 debtValue)
func (_MSCEngine *MSCEngineCallerSession) GetUserInfo(user common.Address) (struct {
	CollateralValue *big.Int
	DebtValue       *big.Int
}, error) {
	return _MSCEngine.Contract.GetUserInfo(&_MSCEngine.CallOpts, user)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MSCEngine *MSCEngineTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MSCEngine *MSCEngineSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.Contract.Burn(&_MSCEngine.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MSCEngine *MSCEngineTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.Contract.Burn(&_MSCEngine.TransactOpts, amount)
}

// DepositAndMint is a paid mutator transaction binding the contract method 0x3ac6b7a3.
//
// Solidity: function depositAndMint(address collateralAddress, uint256 amount, uint256 mintAmount) returns()
func (_MSCEngine *MSCEngineTransactor) DepositAndMint(opts *bind.TransactOpts, collateralAddress common.Address, amount *big.Int, mintAmount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.contract.Transact(opts, "depositAndMint", collateralAddress, amount, mintAmount)
}

// DepositAndMint is a paid mutator transaction binding the contract method 0x3ac6b7a3.
//
// Solidity: function depositAndMint(address collateralAddress, uint256 amount, uint256 mintAmount) returns()
func (_MSCEngine *MSCEngineSession) DepositAndMint(collateralAddress common.Address, amount *big.Int, mintAmount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.Contract.DepositAndMint(&_MSCEngine.TransactOpts, collateralAddress, amount, mintAmount)
}

// DepositAndMint is a paid mutator transaction binding the contract method 0x3ac6b7a3.
//
// Solidity: function depositAndMint(address collateralAddress, uint256 amount, uint256 mintAmount) returns()
func (_MSCEngine *MSCEngineTransactorSession) DepositAndMint(collateralAddress common.Address, amount *big.Int, mintAmount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.Contract.DepositAndMint(&_MSCEngine.TransactOpts, collateralAddress, amount, mintAmount)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address collateralAddress, uint256 amount) returns()
func (_MSCEngine *MSCEngineTransactor) DepositCollateral(opts *bind.TransactOpts, collateralAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.contract.Transact(opts, "depositCollateral", collateralAddress, amount)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address collateralAddress, uint256 amount) returns()
func (_MSCEngine *MSCEngineSession) DepositCollateral(collateralAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.Contract.DepositCollateral(&_MSCEngine.TransactOpts, collateralAddress, amount)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address collateralAddress, uint256 amount) returns()
func (_MSCEngine *MSCEngineTransactorSession) DepositCollateral(collateralAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.Contract.DepositCollateral(&_MSCEngine.TransactOpts, collateralAddress, amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x26c01303.
//
// Solidity: function liquidate(address user, address collateralAddress, uint256 debtToCover) returns()
func (_MSCEngine *MSCEngineTransactor) Liquidate(opts *bind.TransactOpts, user common.Address, collateralAddress common.Address, debtToCover *big.Int) (*types.Transaction, error) {
	return _MSCEngine.contract.Transact(opts, "liquidate", user, collateralAddress, debtToCover)
}

// Liquidate is a paid mutator transaction binding the contract method 0x26c01303.
//
// Solidity: function liquidate(address user, address collateralAddress, uint256 debtToCover) returns()
func (_MSCEngine *MSCEngineSession) Liquidate(user common.Address, collateralAddress common.Address, debtToCover *big.Int) (*types.Transaction, error) {
	return _MSCEngine.Contract.Liquidate(&_MSCEngine.TransactOpts, user, collateralAddress, debtToCover)
}

// Liquidate is a paid mutator transaction binding the contract method 0x26c01303.
//
// Solidity: function liquidate(address user, address collateralAddress, uint256 debtToCover) returns()
func (_MSCEngine *MSCEngineTransactorSession) Liquidate(user common.Address, collateralAddress common.Address, debtToCover *big.Int) (*types.Transaction, error) {
	return _MSCEngine.Contract.Liquidate(&_MSCEngine.TransactOpts, user, collateralAddress, debtToCover)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MSCEngine *MSCEngineTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MSCEngine *MSCEngineSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.Contract.Mint(&_MSCEngine.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MSCEngine *MSCEngineTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.Contract.Mint(&_MSCEngine.TransactOpts, to, amount)
}

// RedeemCollateral is a paid mutator transaction binding the contract method 0x9acd81b3.
//
// Solidity: function redeemCollateral(address collateralAddress, uint256 amount) returns()
func (_MSCEngine *MSCEngineTransactor) RedeemCollateral(opts *bind.TransactOpts, collateralAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.contract.Transact(opts, "redeemCollateral", collateralAddress, amount)
}

// RedeemCollateral is a paid mutator transaction binding the contract method 0x9acd81b3.
//
// Solidity: function redeemCollateral(address collateralAddress, uint256 amount) returns()
func (_MSCEngine *MSCEngineSession) RedeemCollateral(collateralAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.Contract.RedeemCollateral(&_MSCEngine.TransactOpts, collateralAddress, amount)
}

// RedeemCollateral is a paid mutator transaction binding the contract method 0x9acd81b3.
//
// Solidity: function redeemCollateral(address collateralAddress, uint256 amount) returns()
func (_MSCEngine *MSCEngineTransactorSession) RedeemCollateral(collateralAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.Contract.RedeemCollateral(&_MSCEngine.TransactOpts, collateralAddress, amount)
}

// RedeemCollateralForMSC is a paid mutator transaction binding the contract method 0x8b2498a4.
//
// Solidity: function redeemCollateralForMSC(address collateralAddress, uint256 amount, uint256 burnAmount) returns()
func (_MSCEngine *MSCEngineTransactor) RedeemCollateralForMSC(opts *bind.TransactOpts, collateralAddress common.Address, amount *big.Int, burnAmount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.contract.Transact(opts, "redeemCollateralForMSC", collateralAddress, amount, burnAmount)
}

// RedeemCollateralForMSC is a paid mutator transaction binding the contract method 0x8b2498a4.
//
// Solidity: function redeemCollateralForMSC(address collateralAddress, uint256 amount, uint256 burnAmount) returns()
func (_MSCEngine *MSCEngineSession) RedeemCollateralForMSC(collateralAddress common.Address, amount *big.Int, burnAmount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.Contract.RedeemCollateralForMSC(&_MSCEngine.TransactOpts, collateralAddress, amount, burnAmount)
}

// RedeemCollateralForMSC is a paid mutator transaction binding the contract method 0x8b2498a4.
//
// Solidity: function redeemCollateralForMSC(address collateralAddress, uint256 amount, uint256 burnAmount) returns()
func (_MSCEngine *MSCEngineTransactorSession) RedeemCollateralForMSC(collateralAddress common.Address, amount *big.Int, burnAmount *big.Int) (*types.Transaction, error) {
	return _MSCEngine.Contract.RedeemCollateralForMSC(&_MSCEngine.TransactOpts, collateralAddress, amount, burnAmount)
}

// MSCEngineCollateralDepositedIterator is returned from FilterCollateralDeposited and is used to iterate over the raw logs and unpacked data for CollateralDeposited events raised by the MSCEngine contract.
type MSCEngineCollateralDepositedIterator struct {
	Event *MSCEngineCollateralDeposited // Event containing the contract specifics and raw log

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
func (it *MSCEngineCollateralDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MSCEngineCollateralDeposited)
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
		it.Event = new(MSCEngineCollateralDeposited)
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
func (it *MSCEngineCollateralDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MSCEngineCollateralDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MSCEngineCollateralDeposited represents a CollateralDeposited event raised by the MSCEngine contract.
type MSCEngineCollateralDeposited struct {
	Depositor common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollateralDeposited is a free log retrieval operation binding the contract event 0xf1c0dd7e9b98bbff859029005ef89b127af049cd18df1a8d79f0b7e019911e56.
//
// Solidity: event CollateralDeposited(address indexed depositor, address indexed token, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) FilterCollateralDeposited(opts *bind.FilterOpts, depositor []common.Address, token []common.Address) (*MSCEngineCollateralDepositedIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MSCEngine.contract.FilterLogs(opts, "CollateralDeposited", depositorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MSCEngineCollateralDepositedIterator{contract: _MSCEngine.contract, event: "CollateralDeposited", logs: logs, sub: sub}, nil
}

// WatchCollateralDeposited is a free log subscription operation binding the contract event 0xf1c0dd7e9b98bbff859029005ef89b127af049cd18df1a8d79f0b7e019911e56.
//
// Solidity: event CollateralDeposited(address indexed depositor, address indexed token, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) WatchCollateralDeposited(opts *bind.WatchOpts, sink chan<- *MSCEngineCollateralDeposited, depositor []common.Address, token []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MSCEngine.contract.WatchLogs(opts, "CollateralDeposited", depositorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MSCEngineCollateralDeposited)
				if err := _MSCEngine.contract.UnpackLog(event, "CollateralDeposited", log); err != nil {
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

// ParseCollateralDeposited is a log parse operation binding the contract event 0xf1c0dd7e9b98bbff859029005ef89b127af049cd18df1a8d79f0b7e019911e56.
//
// Solidity: event CollateralDeposited(address indexed depositor, address indexed token, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) ParseCollateralDeposited(log types.Log) (*MSCEngineCollateralDeposited, error) {
	event := new(MSCEngineCollateralDeposited)
	if err := _MSCEngine.contract.UnpackLog(event, "CollateralDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MSCEngineCollateralRedeemedIterator is returned from FilterCollateralRedeemed and is used to iterate over the raw logs and unpacked data for CollateralRedeemed events raised by the MSCEngine contract.
type MSCEngineCollateralRedeemedIterator struct {
	Event *MSCEngineCollateralRedeemed // Event containing the contract specifics and raw log

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
func (it *MSCEngineCollateralRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MSCEngineCollateralRedeemed)
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
		it.Event = new(MSCEngineCollateralRedeemed)
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
func (it *MSCEngineCollateralRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MSCEngineCollateralRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MSCEngineCollateralRedeemed represents a CollateralRedeemed event raised by the MSCEngine contract.
type MSCEngineCollateralRedeemed struct {
	Redeemer common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCollateralRedeemed is a free log retrieval operation binding the contract event 0xa5f9505801b85736b93411e5083d5a6003f3add45d82754efd49b4cca6b8e007.
//
// Solidity: event CollateralRedeemed(address indexed redeemer, address indexed token, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) FilterCollateralRedeemed(opts *bind.FilterOpts, redeemer []common.Address, token []common.Address) (*MSCEngineCollateralRedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MSCEngine.contract.FilterLogs(opts, "CollateralRedeemed", redeemerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MSCEngineCollateralRedeemedIterator{contract: _MSCEngine.contract, event: "CollateralRedeemed", logs: logs, sub: sub}, nil
}

// WatchCollateralRedeemed is a free log subscription operation binding the contract event 0xa5f9505801b85736b93411e5083d5a6003f3add45d82754efd49b4cca6b8e007.
//
// Solidity: event CollateralRedeemed(address indexed redeemer, address indexed token, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) WatchCollateralRedeemed(opts *bind.WatchOpts, sink chan<- *MSCEngineCollateralRedeemed, redeemer []common.Address, token []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MSCEngine.contract.WatchLogs(opts, "CollateralRedeemed", redeemerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MSCEngineCollateralRedeemed)
				if err := _MSCEngine.contract.UnpackLog(event, "CollateralRedeemed", log); err != nil {
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

// ParseCollateralRedeemed is a log parse operation binding the contract event 0xa5f9505801b85736b93411e5083d5a6003f3add45d82754efd49b4cca6b8e007.
//
// Solidity: event CollateralRedeemed(address indexed redeemer, address indexed token, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) ParseCollateralRedeemed(log types.Log) (*MSCEngineCollateralRedeemed, error) {
	event := new(MSCEngineCollateralRedeemed)
	if err := _MSCEngine.contract.UnpackLog(event, "CollateralRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MSCEngineDebtBurnedIterator is returned from FilterDebtBurned and is used to iterate over the raw logs and unpacked data for DebtBurned events raised by the MSCEngine contract.
type MSCEngineDebtBurnedIterator struct {
	Event *MSCEngineDebtBurned // Event containing the contract specifics and raw log

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
func (it *MSCEngineDebtBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MSCEngineDebtBurned)
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
		it.Event = new(MSCEngineDebtBurned)
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
func (it *MSCEngineDebtBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MSCEngineDebtBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MSCEngineDebtBurned represents a DebtBurned event raised by the MSCEngine contract.
type MSCEngineDebtBurned struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDebtBurned is a free log retrieval operation binding the contract event 0x97a507cdb275af68c600db8c00bdcae6ed23981cc9f27fa4f2b86a5bdc076eb8.
//
// Solidity: event DebtBurned(address indexed from, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) FilterDebtBurned(opts *bind.FilterOpts, from []common.Address) (*MSCEngineDebtBurnedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MSCEngine.contract.FilterLogs(opts, "DebtBurned", fromRule)
	if err != nil {
		return nil, err
	}
	return &MSCEngineDebtBurnedIterator{contract: _MSCEngine.contract, event: "DebtBurned", logs: logs, sub: sub}, nil
}

// WatchDebtBurned is a free log subscription operation binding the contract event 0x97a507cdb275af68c600db8c00bdcae6ed23981cc9f27fa4f2b86a5bdc076eb8.
//
// Solidity: event DebtBurned(address indexed from, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) WatchDebtBurned(opts *bind.WatchOpts, sink chan<- *MSCEngineDebtBurned, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MSCEngine.contract.WatchLogs(opts, "DebtBurned", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MSCEngineDebtBurned)
				if err := _MSCEngine.contract.UnpackLog(event, "DebtBurned", log); err != nil {
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

// ParseDebtBurned is a log parse operation binding the contract event 0x97a507cdb275af68c600db8c00bdcae6ed23981cc9f27fa4f2b86a5bdc076eb8.
//
// Solidity: event DebtBurned(address indexed from, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) ParseDebtBurned(log types.Log) (*MSCEngineDebtBurned, error) {
	event := new(MSCEngineDebtBurned)
	if err := _MSCEngine.contract.UnpackLog(event, "DebtBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MSCEngineDebtMintedIterator is returned from FilterDebtMinted and is used to iterate over the raw logs and unpacked data for DebtMinted events raised by the MSCEngine contract.
type MSCEngineDebtMintedIterator struct {
	Event *MSCEngineDebtMinted // Event containing the contract specifics and raw log

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
func (it *MSCEngineDebtMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MSCEngineDebtMinted)
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
		it.Event = new(MSCEngineDebtMinted)
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
func (it *MSCEngineDebtMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MSCEngineDebtMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MSCEngineDebtMinted represents a DebtMinted event raised by the MSCEngine contract.
type MSCEngineDebtMinted struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDebtMinted is a free log retrieval operation binding the contract event 0x9b147abdd144aea38aa2d6db5c7851352d6de64a7e633d19f87b03a79febf1c3.
//
// Solidity: event DebtMinted(address indexed to, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) FilterDebtMinted(opts *bind.FilterOpts, to []common.Address) (*MSCEngineDebtMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MSCEngine.contract.FilterLogs(opts, "DebtMinted", toRule)
	if err != nil {
		return nil, err
	}
	return &MSCEngineDebtMintedIterator{contract: _MSCEngine.contract, event: "DebtMinted", logs: logs, sub: sub}, nil
}

// WatchDebtMinted is a free log subscription operation binding the contract event 0x9b147abdd144aea38aa2d6db5c7851352d6de64a7e633d19f87b03a79febf1c3.
//
// Solidity: event DebtMinted(address indexed to, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) WatchDebtMinted(opts *bind.WatchOpts, sink chan<- *MSCEngineDebtMinted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MSCEngine.contract.WatchLogs(opts, "DebtMinted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MSCEngineDebtMinted)
				if err := _MSCEngine.contract.UnpackLog(event, "DebtMinted", log); err != nil {
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

// ParseDebtMinted is a log parse operation binding the contract event 0x9b147abdd144aea38aa2d6db5c7851352d6de64a7e633d19f87b03a79febf1c3.
//
// Solidity: event DebtMinted(address indexed to, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) ParseDebtMinted(log types.Log) (*MSCEngineDebtMinted, error) {
	event := new(MSCEngineDebtMinted)
	if err := _MSCEngine.contract.UnpackLog(event, "DebtMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MSCEngineLiquidatedIterator is returned from FilterLiquidated and is used to iterate over the raw logs and unpacked data for Liquidated events raised by the MSCEngine contract.
type MSCEngineLiquidatedIterator struct {
	Event *MSCEngineLiquidated // Event containing the contract specifics and raw log

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
func (it *MSCEngineLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MSCEngineLiquidated)
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
		it.Event = new(MSCEngineLiquidated)
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
func (it *MSCEngineLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MSCEngineLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MSCEngineLiquidated represents a Liquidated event raised by the MSCEngine contract.
type MSCEngineLiquidated struct {
	User              common.Address
	CollateralAddress common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLiquidated is a free log retrieval operation binding the contract event 0xde0aa27286f5cb3a4ed853dc4823ead62d63e92cdf13de09d6aece56970721a4.
//
// Solidity: event Liquidated(address indexed user, address indexed collateralAddress, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) FilterLiquidated(opts *bind.FilterOpts, user []common.Address, collateralAddress []common.Address) (*MSCEngineLiquidatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var collateralAddressRule []interface{}
	for _, collateralAddressItem := range collateralAddress {
		collateralAddressRule = append(collateralAddressRule, collateralAddressItem)
	}

	logs, sub, err := _MSCEngine.contract.FilterLogs(opts, "Liquidated", userRule, collateralAddressRule)
	if err != nil {
		return nil, err
	}
	return &MSCEngineLiquidatedIterator{contract: _MSCEngine.contract, event: "Liquidated", logs: logs, sub: sub}, nil
}

// WatchLiquidated is a free log subscription operation binding the contract event 0xde0aa27286f5cb3a4ed853dc4823ead62d63e92cdf13de09d6aece56970721a4.
//
// Solidity: event Liquidated(address indexed user, address indexed collateralAddress, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) WatchLiquidated(opts *bind.WatchOpts, sink chan<- *MSCEngineLiquidated, user []common.Address, collateralAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var collateralAddressRule []interface{}
	for _, collateralAddressItem := range collateralAddress {
		collateralAddressRule = append(collateralAddressRule, collateralAddressItem)
	}

	logs, sub, err := _MSCEngine.contract.WatchLogs(opts, "Liquidated", userRule, collateralAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MSCEngineLiquidated)
				if err := _MSCEngine.contract.UnpackLog(event, "Liquidated", log); err != nil {
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

// ParseLiquidated is a log parse operation binding the contract event 0xde0aa27286f5cb3a4ed853dc4823ead62d63e92cdf13de09d6aece56970721a4.
//
// Solidity: event Liquidated(address indexed user, address indexed collateralAddress, uint256 amount)
func (_MSCEngine *MSCEngineFilterer) ParseLiquidated(log types.Log) (*MSCEngineLiquidated, error) {
	event := new(MSCEngineLiquidated)
	if err := _MSCEngine.contract.UnpackLog(event, "Liquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// KeepBondingABI is the input ABI used to generate the binding from.
const KeepBondingABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenStakingAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenGrantAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sortitionPool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referenceID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"referenceID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newHolder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReferenceID\",\"type\":\"uint256\"}],\"name\":\"BondReassigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"referenceID\",\"type\":\"uint256\"}],\"name\":\"BondReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"referenceID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondSeized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnbondedValueDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnbondedValueWithdrawn\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"authorizeSortitionPoolContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"authorizerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bondCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"authorizedSortitionPool\",\"type\":\"address\"}],\"name\":\"availableUnbondedValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"beneficiaryOf\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"referenceID\",\"type\":\"uint256\"}],\"name\":\"bondAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"referenceID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"authorizedSortitionPool\",\"type\":\"address\"}],\"name\":\"createBond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"deauthorizeSortitionPoolContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"referenceID\",\"type\":\"uint256\"}],\"name\":\"freeBond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"hasSecondaryAuthorization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operatorContract\",\"type\":\"address\"}],\"name\":\"isAuthorizedForOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"referenceID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newHolder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newReferenceID\",\"type\":\"uint256\"}],\"name\":\"reassignBond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"referenceID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"seizeBond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unbondedValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"managedGrant\",\"type\":\"address\"}],\"name\":\"withdrawAsManagedGrantee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KeepBonding is an auto generated Go binding around an Ethereum contract.
type KeepBonding struct {
	KeepBondingCaller     // Read-only binding to the contract
	KeepBondingTransactor // Write-only binding to the contract
	KeepBondingFilterer   // Log filterer for contract events
}

// KeepBondingCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeepBondingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeepBondingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeepBondingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeepBondingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeepBondingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeepBondingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeepBondingSession struct {
	Contract     *KeepBonding      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeepBondingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeepBondingCallerSession struct {
	Contract *KeepBondingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// KeepBondingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeepBondingTransactorSession struct {
	Contract     *KeepBondingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// KeepBondingRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeepBondingRaw struct {
	Contract *KeepBonding // Generic contract binding to access the raw methods on
}

// KeepBondingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeepBondingCallerRaw struct {
	Contract *KeepBondingCaller // Generic read-only contract binding to access the raw methods on
}

// KeepBondingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeepBondingTransactorRaw struct {
	Contract *KeepBondingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeepBonding creates a new instance of KeepBonding, bound to a specific deployed contract.
func NewKeepBonding(address common.Address, backend bind.ContractBackend) (*KeepBonding, error) {
	contract, err := bindKeepBonding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeepBonding{KeepBondingCaller: KeepBondingCaller{contract: contract}, KeepBondingTransactor: KeepBondingTransactor{contract: contract}, KeepBondingFilterer: KeepBondingFilterer{contract: contract}}, nil
}

// NewKeepBondingCaller creates a new read-only instance of KeepBonding, bound to a specific deployed contract.
func NewKeepBondingCaller(address common.Address, caller bind.ContractCaller) (*KeepBondingCaller, error) {
	contract, err := bindKeepBonding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeepBondingCaller{contract: contract}, nil
}

// NewKeepBondingTransactor creates a new write-only instance of KeepBonding, bound to a specific deployed contract.
func NewKeepBondingTransactor(address common.Address, transactor bind.ContractTransactor) (*KeepBondingTransactor, error) {
	contract, err := bindKeepBonding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeepBondingTransactor{contract: contract}, nil
}

// NewKeepBondingFilterer creates a new log filterer instance of KeepBonding, bound to a specific deployed contract.
func NewKeepBondingFilterer(address common.Address, filterer bind.ContractFilterer) (*KeepBondingFilterer, error) {
	contract, err := bindKeepBonding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeepBondingFilterer{contract: contract}, nil
}

// bindKeepBonding binds a generic wrapper to an already deployed contract.
func bindKeepBonding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeepBondingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeepBonding *KeepBondingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KeepBonding.Contract.KeepBondingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeepBonding *KeepBondingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeepBonding.Contract.KeepBondingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeepBonding *KeepBondingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeepBonding.Contract.KeepBondingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeepBonding *KeepBondingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KeepBonding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeepBonding *KeepBondingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeepBonding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeepBonding *KeepBondingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeepBonding.Contract.contract.Transact(opts, method, params...)
}

// AuthorizerOf is a free data retrieval call binding the contract method 0xfb1677b1.
//
// Solidity: function authorizerOf(address _operator) constant returns(address)
func (_KeepBonding *KeepBondingCaller) AuthorizerOf(opts *bind.CallOpts, _operator common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KeepBonding.contract.Call(opts, out, "authorizerOf", _operator)
	return *ret0, err
}

// AuthorizerOf is a free data retrieval call binding the contract method 0xfb1677b1.
//
// Solidity: function authorizerOf(address _operator) constant returns(address)
func (_KeepBonding *KeepBondingSession) AuthorizerOf(_operator common.Address) (common.Address, error) {
	return _KeepBonding.Contract.AuthorizerOf(&_KeepBonding.CallOpts, _operator)
}

// AuthorizerOf is a free data retrieval call binding the contract method 0xfb1677b1.
//
// Solidity: function authorizerOf(address _operator) constant returns(address)
func (_KeepBonding *KeepBondingCallerSession) AuthorizerOf(_operator common.Address) (common.Address, error) {
	return _KeepBonding.Contract.AuthorizerOf(&_KeepBonding.CallOpts, _operator)
}

// AvailableUnbondedValue is a free data retrieval call binding the contract method 0x42bcb965.
//
// Solidity: function availableUnbondedValue(address operator, address bondCreator, address authorizedSortitionPool) constant returns(uint256)
func (_KeepBonding *KeepBondingCaller) AvailableUnbondedValue(opts *bind.CallOpts, operator common.Address, bondCreator common.Address, authorizedSortitionPool common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KeepBonding.contract.Call(opts, out, "availableUnbondedValue", operator, bondCreator, authorizedSortitionPool)
	return *ret0, err
}

// AvailableUnbondedValue is a free data retrieval call binding the contract method 0x42bcb965.
//
// Solidity: function availableUnbondedValue(address operator, address bondCreator, address authorizedSortitionPool) constant returns(uint256)
func (_KeepBonding *KeepBondingSession) AvailableUnbondedValue(operator common.Address, bondCreator common.Address, authorizedSortitionPool common.Address) (*big.Int, error) {
	return _KeepBonding.Contract.AvailableUnbondedValue(&_KeepBonding.CallOpts, operator, bondCreator, authorizedSortitionPool)
}

// AvailableUnbondedValue is a free data retrieval call binding the contract method 0x42bcb965.
//
// Solidity: function availableUnbondedValue(address operator, address bondCreator, address authorizedSortitionPool) constant returns(uint256)
func (_KeepBonding *KeepBondingCallerSession) AvailableUnbondedValue(operator common.Address, bondCreator common.Address, authorizedSortitionPool common.Address) (*big.Int, error) {
	return _KeepBonding.Contract.AvailableUnbondedValue(&_KeepBonding.CallOpts, operator, bondCreator, authorizedSortitionPool)
}

// BeneficiaryOf is a free data retrieval call binding the contract method 0xba7bffd3.
//
// Solidity: function beneficiaryOf(address _operator) constant returns(address)
func (_KeepBonding *KeepBondingCaller) BeneficiaryOf(opts *bind.CallOpts, _operator common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KeepBonding.contract.Call(opts, out, "beneficiaryOf", _operator)
	return *ret0, err
}

// BeneficiaryOf is a free data retrieval call binding the contract method 0xba7bffd3.
//
// Solidity: function beneficiaryOf(address _operator) constant returns(address)
func (_KeepBonding *KeepBondingSession) BeneficiaryOf(_operator common.Address) (common.Address, error) {
	return _KeepBonding.Contract.BeneficiaryOf(&_KeepBonding.CallOpts, _operator)
}

// BeneficiaryOf is a free data retrieval call binding the contract method 0xba7bffd3.
//
// Solidity: function beneficiaryOf(address _operator) constant returns(address)
func (_KeepBonding *KeepBondingCallerSession) BeneficiaryOf(_operator common.Address) (common.Address, error) {
	return _KeepBonding.Contract.BeneficiaryOf(&_KeepBonding.CallOpts, _operator)
}

// BondAmount is a free data retrieval call binding the contract method 0x446f0f9e.
//
// Solidity: function bondAmount(address operator, address holder, uint256 referenceID) constant returns(uint256)
func (_KeepBonding *KeepBondingCaller) BondAmount(opts *bind.CallOpts, operator common.Address, holder common.Address, referenceID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KeepBonding.contract.Call(opts, out, "bondAmount", operator, holder, referenceID)
	return *ret0, err
}

// BondAmount is a free data retrieval call binding the contract method 0x446f0f9e.
//
// Solidity: function bondAmount(address operator, address holder, uint256 referenceID) constant returns(uint256)
func (_KeepBonding *KeepBondingSession) BondAmount(operator common.Address, holder common.Address, referenceID *big.Int) (*big.Int, error) {
	return _KeepBonding.Contract.BondAmount(&_KeepBonding.CallOpts, operator, holder, referenceID)
}

// BondAmount is a free data retrieval call binding the contract method 0x446f0f9e.
//
// Solidity: function bondAmount(address operator, address holder, uint256 referenceID) constant returns(uint256)
func (_KeepBonding *KeepBondingCallerSession) BondAmount(operator common.Address, holder common.Address, referenceID *big.Int) (*big.Int, error) {
	return _KeepBonding.Contract.BondAmount(&_KeepBonding.CallOpts, operator, holder, referenceID)
}

// HasSecondaryAuthorization is a free data retrieval call binding the contract method 0x78f011c1.
//
// Solidity: function hasSecondaryAuthorization(address _operator, address _poolAddress) constant returns(bool)
func (_KeepBonding *KeepBondingCaller) HasSecondaryAuthorization(opts *bind.CallOpts, _operator common.Address, _poolAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KeepBonding.contract.Call(opts, out, "hasSecondaryAuthorization", _operator, _poolAddress)
	return *ret0, err
}

// HasSecondaryAuthorization is a free data retrieval call binding the contract method 0x78f011c1.
//
// Solidity: function hasSecondaryAuthorization(address _operator, address _poolAddress) constant returns(bool)
func (_KeepBonding *KeepBondingSession) HasSecondaryAuthorization(_operator common.Address, _poolAddress common.Address) (bool, error) {
	return _KeepBonding.Contract.HasSecondaryAuthorization(&_KeepBonding.CallOpts, _operator, _poolAddress)
}

// HasSecondaryAuthorization is a free data retrieval call binding the contract method 0x78f011c1.
//
// Solidity: function hasSecondaryAuthorization(address _operator, address _poolAddress) constant returns(bool)
func (_KeepBonding *KeepBondingCallerSession) HasSecondaryAuthorization(_operator common.Address, _poolAddress common.Address) (bool, error) {
	return _KeepBonding.Contract.HasSecondaryAuthorization(&_KeepBonding.CallOpts, _operator, _poolAddress)
}

// IsAuthorizedForOperator is a free data retrieval call binding the contract method 0xef1f9661.
//
// Solidity: function isAuthorizedForOperator(address _operator, address _operatorContract) constant returns(bool)
func (_KeepBonding *KeepBondingCaller) IsAuthorizedForOperator(opts *bind.CallOpts, _operator common.Address, _operatorContract common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KeepBonding.contract.Call(opts, out, "isAuthorizedForOperator", _operator, _operatorContract)
	return *ret0, err
}

// IsAuthorizedForOperator is a free data retrieval call binding the contract method 0xef1f9661.
//
// Solidity: function isAuthorizedForOperator(address _operator, address _operatorContract) constant returns(bool)
func (_KeepBonding *KeepBondingSession) IsAuthorizedForOperator(_operator common.Address, _operatorContract common.Address) (bool, error) {
	return _KeepBonding.Contract.IsAuthorizedForOperator(&_KeepBonding.CallOpts, _operator, _operatorContract)
}

// IsAuthorizedForOperator is a free data retrieval call binding the contract method 0xef1f9661.
//
// Solidity: function isAuthorizedForOperator(address _operator, address _operatorContract) constant returns(bool)
func (_KeepBonding *KeepBondingCallerSession) IsAuthorizedForOperator(_operator common.Address, _operatorContract common.Address) (bool, error) {
	return _KeepBonding.Contract.IsAuthorizedForOperator(&_KeepBonding.CallOpts, _operator, _operatorContract)
}

// UnbondedValue is a free data retrieval call binding the contract method 0x5823cfad.
//
// Solidity: function unbondedValue(address ) constant returns(uint256)
func (_KeepBonding *KeepBondingCaller) UnbondedValue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KeepBonding.contract.Call(opts, out, "unbondedValue", arg0)
	return *ret0, err
}

// UnbondedValue is a free data retrieval call binding the contract method 0x5823cfad.
//
// Solidity: function unbondedValue(address ) constant returns(uint256)
func (_KeepBonding *KeepBondingSession) UnbondedValue(arg0 common.Address) (*big.Int, error) {
	return _KeepBonding.Contract.UnbondedValue(&_KeepBonding.CallOpts, arg0)
}

// UnbondedValue is a free data retrieval call binding the contract method 0x5823cfad.
//
// Solidity: function unbondedValue(address ) constant returns(uint256)
func (_KeepBonding *KeepBondingCallerSession) UnbondedValue(arg0 common.Address) (*big.Int, error) {
	return _KeepBonding.Contract.UnbondedValue(&_KeepBonding.CallOpts, arg0)
}

// AuthorizeSortitionPoolContract is a paid mutator transaction binding the contract method 0xc5786174.
//
// Solidity: function authorizeSortitionPoolContract(address _operator, address _poolAddress) returns()
func (_KeepBonding *KeepBondingTransactor) AuthorizeSortitionPoolContract(opts *bind.TransactOpts, _operator common.Address, _poolAddress common.Address) (*types.Transaction, error) {
	return _KeepBonding.contract.Transact(opts, "authorizeSortitionPoolContract", _operator, _poolAddress)
}

// AuthorizeSortitionPoolContract is a paid mutator transaction binding the contract method 0xc5786174.
//
// Solidity: function authorizeSortitionPoolContract(address _operator, address _poolAddress) returns()
func (_KeepBonding *KeepBondingSession) AuthorizeSortitionPoolContract(_operator common.Address, _poolAddress common.Address) (*types.Transaction, error) {
	return _KeepBonding.Contract.AuthorizeSortitionPoolContract(&_KeepBonding.TransactOpts, _operator, _poolAddress)
}

// AuthorizeSortitionPoolContract is a paid mutator transaction binding the contract method 0xc5786174.
//
// Solidity: function authorizeSortitionPoolContract(address _operator, address _poolAddress) returns()
func (_KeepBonding *KeepBondingTransactorSession) AuthorizeSortitionPoolContract(_operator common.Address, _poolAddress common.Address) (*types.Transaction, error) {
	return _KeepBonding.Contract.AuthorizeSortitionPoolContract(&_KeepBonding.TransactOpts, _operator, _poolAddress)
}

// CreateBond is a paid mutator transaction binding the contract method 0xd20a62fc.
//
// Solidity: function createBond(address operator, address holder, uint256 referenceID, uint256 amount, address authorizedSortitionPool) returns()
func (_KeepBonding *KeepBondingTransactor) CreateBond(opts *bind.TransactOpts, operator common.Address, holder common.Address, referenceID *big.Int, amount *big.Int, authorizedSortitionPool common.Address) (*types.Transaction, error) {
	return _KeepBonding.contract.Transact(opts, "createBond", operator, holder, referenceID, amount, authorizedSortitionPool)
}

// CreateBond is a paid mutator transaction binding the contract method 0xd20a62fc.
//
// Solidity: function createBond(address operator, address holder, uint256 referenceID, uint256 amount, address authorizedSortitionPool) returns()
func (_KeepBonding *KeepBondingSession) CreateBond(operator common.Address, holder common.Address, referenceID *big.Int, amount *big.Int, authorizedSortitionPool common.Address) (*types.Transaction, error) {
	return _KeepBonding.Contract.CreateBond(&_KeepBonding.TransactOpts, operator, holder, referenceID, amount, authorizedSortitionPool)
}

// CreateBond is a paid mutator transaction binding the contract method 0xd20a62fc.
//
// Solidity: function createBond(address operator, address holder, uint256 referenceID, uint256 amount, address authorizedSortitionPool) returns()
func (_KeepBonding *KeepBondingTransactorSession) CreateBond(operator common.Address, holder common.Address, referenceID *big.Int, amount *big.Int, authorizedSortitionPool common.Address) (*types.Transaction, error) {
	return _KeepBonding.Contract.CreateBond(&_KeepBonding.TransactOpts, operator, holder, referenceID, amount, authorizedSortitionPool)
}

// DeauthorizeSortitionPoolContract is a paid mutator transaction binding the contract method 0x0b102471.
//
// Solidity: function deauthorizeSortitionPoolContract(address _operator, address _poolAddress) returns()
func (_KeepBonding *KeepBondingTransactor) DeauthorizeSortitionPoolContract(opts *bind.TransactOpts, _operator common.Address, _poolAddress common.Address) (*types.Transaction, error) {
	return _KeepBonding.contract.Transact(opts, "deauthorizeSortitionPoolContract", _operator, _poolAddress)
}

// DeauthorizeSortitionPoolContract is a paid mutator transaction binding the contract method 0x0b102471.
//
// Solidity: function deauthorizeSortitionPoolContract(address _operator, address _poolAddress) returns()
func (_KeepBonding *KeepBondingSession) DeauthorizeSortitionPoolContract(_operator common.Address, _poolAddress common.Address) (*types.Transaction, error) {
	return _KeepBonding.Contract.DeauthorizeSortitionPoolContract(&_KeepBonding.TransactOpts, _operator, _poolAddress)
}

// DeauthorizeSortitionPoolContract is a paid mutator transaction binding the contract method 0x0b102471.
//
// Solidity: function deauthorizeSortitionPoolContract(address _operator, address _poolAddress) returns()
func (_KeepBonding *KeepBondingTransactorSession) DeauthorizeSortitionPoolContract(_operator common.Address, _poolAddress common.Address) (*types.Transaction, error) {
	return _KeepBonding.Contract.DeauthorizeSortitionPoolContract(&_KeepBonding.TransactOpts, _operator, _poolAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address operator) returns()
func (_KeepBonding *KeepBondingTransactor) Deposit(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _KeepBonding.contract.Transact(opts, "deposit", operator)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address operator) returns()
func (_KeepBonding *KeepBondingSession) Deposit(operator common.Address) (*types.Transaction, error) {
	return _KeepBonding.Contract.Deposit(&_KeepBonding.TransactOpts, operator)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address operator) returns()
func (_KeepBonding *KeepBondingTransactorSession) Deposit(operator common.Address) (*types.Transaction, error) {
	return _KeepBonding.Contract.Deposit(&_KeepBonding.TransactOpts, operator)
}

// FreeBond is a paid mutator transaction binding the contract method 0x7ab3cf93.
//
// Solidity: function freeBond(address operator, uint256 referenceID) returns()
func (_KeepBonding *KeepBondingTransactor) FreeBond(opts *bind.TransactOpts, operator common.Address, referenceID *big.Int) (*types.Transaction, error) {
	return _KeepBonding.contract.Transact(opts, "freeBond", operator, referenceID)
}

// FreeBond is a paid mutator transaction binding the contract method 0x7ab3cf93.
//
// Solidity: function freeBond(address operator, uint256 referenceID) returns()
func (_KeepBonding *KeepBondingSession) FreeBond(operator common.Address, referenceID *big.Int) (*types.Transaction, error) {
	return _KeepBonding.Contract.FreeBond(&_KeepBonding.TransactOpts, operator, referenceID)
}

// FreeBond is a paid mutator transaction binding the contract method 0x7ab3cf93.
//
// Solidity: function freeBond(address operator, uint256 referenceID) returns()
func (_KeepBonding *KeepBondingTransactorSession) FreeBond(operator common.Address, referenceID *big.Int) (*types.Transaction, error) {
	return _KeepBonding.Contract.FreeBond(&_KeepBonding.TransactOpts, operator, referenceID)
}

// ReassignBond is a paid mutator transaction binding the contract method 0x972f2457.
//
// Solidity: function reassignBond(address operator, uint256 referenceID, address newHolder, uint256 newReferenceID) returns()
func (_KeepBonding *KeepBondingTransactor) ReassignBond(opts *bind.TransactOpts, operator common.Address, referenceID *big.Int, newHolder common.Address, newReferenceID *big.Int) (*types.Transaction, error) {
	return _KeepBonding.contract.Transact(opts, "reassignBond", operator, referenceID, newHolder, newReferenceID)
}

// ReassignBond is a paid mutator transaction binding the contract method 0x972f2457.
//
// Solidity: function reassignBond(address operator, uint256 referenceID, address newHolder, uint256 newReferenceID) returns()
func (_KeepBonding *KeepBondingSession) ReassignBond(operator common.Address, referenceID *big.Int, newHolder common.Address, newReferenceID *big.Int) (*types.Transaction, error) {
	return _KeepBonding.Contract.ReassignBond(&_KeepBonding.TransactOpts, operator, referenceID, newHolder, newReferenceID)
}

// ReassignBond is a paid mutator transaction binding the contract method 0x972f2457.
//
// Solidity: function reassignBond(address operator, uint256 referenceID, address newHolder, uint256 newReferenceID) returns()
func (_KeepBonding *KeepBondingTransactorSession) ReassignBond(operator common.Address, referenceID *big.Int, newHolder common.Address, newReferenceID *big.Int) (*types.Transaction, error) {
	return _KeepBonding.Contract.ReassignBond(&_KeepBonding.TransactOpts, operator, referenceID, newHolder, newReferenceID)
}

// SeizeBond is a paid mutator transaction binding the contract method 0x0cb0a677.
//
// Solidity: function seizeBond(address operator, uint256 referenceID, uint256 amount, address destination) returns()
func (_KeepBonding *KeepBondingTransactor) SeizeBond(opts *bind.TransactOpts, operator common.Address, referenceID *big.Int, amount *big.Int, destination common.Address) (*types.Transaction, error) {
	return _KeepBonding.contract.Transact(opts, "seizeBond", operator, referenceID, amount, destination)
}

// SeizeBond is a paid mutator transaction binding the contract method 0x0cb0a677.
//
// Solidity: function seizeBond(address operator, uint256 referenceID, uint256 amount, address destination) returns()
func (_KeepBonding *KeepBondingSession) SeizeBond(operator common.Address, referenceID *big.Int, amount *big.Int, destination common.Address) (*types.Transaction, error) {
	return _KeepBonding.Contract.SeizeBond(&_KeepBonding.TransactOpts, operator, referenceID, amount, destination)
}

// SeizeBond is a paid mutator transaction binding the contract method 0x0cb0a677.
//
// Solidity: function seizeBond(address operator, uint256 referenceID, uint256 amount, address destination) returns()
func (_KeepBonding *KeepBondingTransactorSession) SeizeBond(operator common.Address, referenceID *big.Int, amount *big.Int, destination common.Address) (*types.Transaction, error) {
	return _KeepBonding.Contract.SeizeBond(&_KeepBonding.TransactOpts, operator, referenceID, amount, destination)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address operator) returns()
func (_KeepBonding *KeepBondingTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, operator common.Address) (*types.Transaction, error) {
	return _KeepBonding.contract.Transact(opts, "withdraw", amount, operator)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address operator) returns()
func (_KeepBonding *KeepBondingSession) Withdraw(amount *big.Int, operator common.Address) (*types.Transaction, error) {
	return _KeepBonding.Contract.Withdraw(&_KeepBonding.TransactOpts, amount, operator)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address operator) returns()
func (_KeepBonding *KeepBondingTransactorSession) Withdraw(amount *big.Int, operator common.Address) (*types.Transaction, error) {
	return _KeepBonding.Contract.Withdraw(&_KeepBonding.TransactOpts, amount, operator)
}

// WithdrawAsManagedGrantee is a paid mutator transaction binding the contract method 0x5fcac8ff.
//
// Solidity: function withdrawAsManagedGrantee(uint256 amount, address operator, address managedGrant) returns()
func (_KeepBonding *KeepBondingTransactor) WithdrawAsManagedGrantee(opts *bind.TransactOpts, amount *big.Int, operator common.Address, managedGrant common.Address) (*types.Transaction, error) {
	return _KeepBonding.contract.Transact(opts, "withdrawAsManagedGrantee", amount, operator, managedGrant)
}

// WithdrawAsManagedGrantee is a paid mutator transaction binding the contract method 0x5fcac8ff.
//
// Solidity: function withdrawAsManagedGrantee(uint256 amount, address operator, address managedGrant) returns()
func (_KeepBonding *KeepBondingSession) WithdrawAsManagedGrantee(amount *big.Int, operator common.Address, managedGrant common.Address) (*types.Transaction, error) {
	return _KeepBonding.Contract.WithdrawAsManagedGrantee(&_KeepBonding.TransactOpts, amount, operator, managedGrant)
}

// WithdrawAsManagedGrantee is a paid mutator transaction binding the contract method 0x5fcac8ff.
//
// Solidity: function withdrawAsManagedGrantee(uint256 amount, address operator, address managedGrant) returns()
func (_KeepBonding *KeepBondingTransactorSession) WithdrawAsManagedGrantee(amount *big.Int, operator common.Address, managedGrant common.Address) (*types.Transaction, error) {
	return _KeepBonding.Contract.WithdrawAsManagedGrantee(&_KeepBonding.TransactOpts, amount, operator, managedGrant)
}

// KeepBondingBondCreatedIterator is returned from FilterBondCreated and is used to iterate over the raw logs and unpacked data for BondCreated events raised by the KeepBonding contract.
type KeepBondingBondCreatedIterator struct {
	Event *KeepBondingBondCreated // Event containing the contract specifics and raw log

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
func (it *KeepBondingBondCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeepBondingBondCreated)
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
		it.Event = new(KeepBondingBondCreated)
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
func (it *KeepBondingBondCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeepBondingBondCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeepBondingBondCreated represents a BondCreated event raised by the KeepBonding contract.
type KeepBondingBondCreated struct {
	Operator      common.Address
	Holder        common.Address
	SortitionPool common.Address
	ReferenceID   *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBondCreated is a free log retrieval operation binding the contract event 0xa5543d8e139d9ab4342d5c4f6ec1bff5a97f9a52d71f7ffe9845b94f1449fc91.
//
// Solidity: event BondCreated(address indexed operator, address indexed holder, address indexed sortitionPool, uint256 referenceID, uint256 amount)
func (_KeepBonding *KeepBondingFilterer) FilterBondCreated(opts *bind.FilterOpts, operator []common.Address, holder []common.Address, sortitionPool []common.Address) (*KeepBondingBondCreatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var sortitionPoolRule []interface{}
	for _, sortitionPoolItem := range sortitionPool {
		sortitionPoolRule = append(sortitionPoolRule, sortitionPoolItem)
	}

	logs, sub, err := _KeepBonding.contract.FilterLogs(opts, "BondCreated", operatorRule, holderRule, sortitionPoolRule)
	if err != nil {
		return nil, err
	}
	return &KeepBondingBondCreatedIterator{contract: _KeepBonding.contract, event: "BondCreated", logs: logs, sub: sub}, nil
}

// WatchBondCreated is a free log subscription operation binding the contract event 0xa5543d8e139d9ab4342d5c4f6ec1bff5a97f9a52d71f7ffe9845b94f1449fc91.
//
// Solidity: event BondCreated(address indexed operator, address indexed holder, address indexed sortitionPool, uint256 referenceID, uint256 amount)
func (_KeepBonding *KeepBondingFilterer) WatchBondCreated(opts *bind.WatchOpts, sink chan<- *KeepBondingBondCreated, operator []common.Address, holder []common.Address, sortitionPool []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var sortitionPoolRule []interface{}
	for _, sortitionPoolItem := range sortitionPool {
		sortitionPoolRule = append(sortitionPoolRule, sortitionPoolItem)
	}

	logs, sub, err := _KeepBonding.contract.WatchLogs(opts, "BondCreated", operatorRule, holderRule, sortitionPoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeepBondingBondCreated)
				if err := _KeepBonding.contract.UnpackLog(event, "BondCreated", log); err != nil {
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

// ParseBondCreated is a log parse operation binding the contract event 0xa5543d8e139d9ab4342d5c4f6ec1bff5a97f9a52d71f7ffe9845b94f1449fc91.
//
// Solidity: event BondCreated(address indexed operator, address indexed holder, address indexed sortitionPool, uint256 referenceID, uint256 amount)
func (_KeepBonding *KeepBondingFilterer) ParseBondCreated(log types.Log) (*KeepBondingBondCreated, error) {
	event := new(KeepBondingBondCreated)
	if err := _KeepBonding.contract.UnpackLog(event, "BondCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KeepBondingBondReassignedIterator is returned from FilterBondReassigned and is used to iterate over the raw logs and unpacked data for BondReassigned events raised by the KeepBonding contract.
type KeepBondingBondReassignedIterator struct {
	Event *KeepBondingBondReassigned // Event containing the contract specifics and raw log

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
func (it *KeepBondingBondReassignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeepBondingBondReassigned)
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
		it.Event = new(KeepBondingBondReassigned)
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
func (it *KeepBondingBondReassignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeepBondingBondReassignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeepBondingBondReassigned represents a BondReassigned event raised by the KeepBonding contract.
type KeepBondingBondReassigned struct {
	Operator       common.Address
	ReferenceID    *big.Int
	NewHolder      common.Address
	NewReferenceID *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondReassigned is a free log retrieval operation binding the contract event 0xb1d917176802bfbc813f2d82e745526029a4ccf0ea98d14e7a09a08703595b1e.
//
// Solidity: event BondReassigned(address indexed operator, uint256 indexed referenceID, address newHolder, uint256 newReferenceID)
func (_KeepBonding *KeepBondingFilterer) FilterBondReassigned(opts *bind.FilterOpts, operator []common.Address, referenceID []*big.Int) (*KeepBondingBondReassignedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var referenceIDRule []interface{}
	for _, referenceIDItem := range referenceID {
		referenceIDRule = append(referenceIDRule, referenceIDItem)
	}

	logs, sub, err := _KeepBonding.contract.FilterLogs(opts, "BondReassigned", operatorRule, referenceIDRule)
	if err != nil {
		return nil, err
	}
	return &KeepBondingBondReassignedIterator{contract: _KeepBonding.contract, event: "BondReassigned", logs: logs, sub: sub}, nil
}

// WatchBondReassigned is a free log subscription operation binding the contract event 0xb1d917176802bfbc813f2d82e745526029a4ccf0ea98d14e7a09a08703595b1e.
//
// Solidity: event BondReassigned(address indexed operator, uint256 indexed referenceID, address newHolder, uint256 newReferenceID)
func (_KeepBonding *KeepBondingFilterer) WatchBondReassigned(opts *bind.WatchOpts, sink chan<- *KeepBondingBondReassigned, operator []common.Address, referenceID []*big.Int) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var referenceIDRule []interface{}
	for _, referenceIDItem := range referenceID {
		referenceIDRule = append(referenceIDRule, referenceIDItem)
	}

	logs, sub, err := _KeepBonding.contract.WatchLogs(opts, "BondReassigned", operatorRule, referenceIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeepBondingBondReassigned)
				if err := _KeepBonding.contract.UnpackLog(event, "BondReassigned", log); err != nil {
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

// ParseBondReassigned is a log parse operation binding the contract event 0xb1d917176802bfbc813f2d82e745526029a4ccf0ea98d14e7a09a08703595b1e.
//
// Solidity: event BondReassigned(address indexed operator, uint256 indexed referenceID, address newHolder, uint256 newReferenceID)
func (_KeepBonding *KeepBondingFilterer) ParseBondReassigned(log types.Log) (*KeepBondingBondReassigned, error) {
	event := new(KeepBondingBondReassigned)
	if err := _KeepBonding.contract.UnpackLog(event, "BondReassigned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KeepBondingBondReleasedIterator is returned from FilterBondReleased and is used to iterate over the raw logs and unpacked data for BondReleased events raised by the KeepBonding contract.
type KeepBondingBondReleasedIterator struct {
	Event *KeepBondingBondReleased // Event containing the contract specifics and raw log

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
func (it *KeepBondingBondReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeepBondingBondReleased)
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
		it.Event = new(KeepBondingBondReleased)
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
func (it *KeepBondingBondReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeepBondingBondReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeepBondingBondReleased represents a BondReleased event raised by the KeepBonding contract.
type KeepBondingBondReleased struct {
	Operator    common.Address
	ReferenceID *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBondReleased is a free log retrieval operation binding the contract event 0x60b8ef4216791426b3d7acfb0b6d11a400872350afd70a3ce5ebf62bea7cb0d4.
//
// Solidity: event BondReleased(address indexed operator, uint256 indexed referenceID)
func (_KeepBonding *KeepBondingFilterer) FilterBondReleased(opts *bind.FilterOpts, operator []common.Address, referenceID []*big.Int) (*KeepBondingBondReleasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var referenceIDRule []interface{}
	for _, referenceIDItem := range referenceID {
		referenceIDRule = append(referenceIDRule, referenceIDItem)
	}

	logs, sub, err := _KeepBonding.contract.FilterLogs(opts, "BondReleased", operatorRule, referenceIDRule)
	if err != nil {
		return nil, err
	}
	return &KeepBondingBondReleasedIterator{contract: _KeepBonding.contract, event: "BondReleased", logs: logs, sub: sub}, nil
}

// WatchBondReleased is a free log subscription operation binding the contract event 0x60b8ef4216791426b3d7acfb0b6d11a400872350afd70a3ce5ebf62bea7cb0d4.
//
// Solidity: event BondReleased(address indexed operator, uint256 indexed referenceID)
func (_KeepBonding *KeepBondingFilterer) WatchBondReleased(opts *bind.WatchOpts, sink chan<- *KeepBondingBondReleased, operator []common.Address, referenceID []*big.Int) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var referenceIDRule []interface{}
	for _, referenceIDItem := range referenceID {
		referenceIDRule = append(referenceIDRule, referenceIDItem)
	}

	logs, sub, err := _KeepBonding.contract.WatchLogs(opts, "BondReleased", operatorRule, referenceIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeepBondingBondReleased)
				if err := _KeepBonding.contract.UnpackLog(event, "BondReleased", log); err != nil {
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

// ParseBondReleased is a log parse operation binding the contract event 0x60b8ef4216791426b3d7acfb0b6d11a400872350afd70a3ce5ebf62bea7cb0d4.
//
// Solidity: event BondReleased(address indexed operator, uint256 indexed referenceID)
func (_KeepBonding *KeepBondingFilterer) ParseBondReleased(log types.Log) (*KeepBondingBondReleased, error) {
	event := new(KeepBondingBondReleased)
	if err := _KeepBonding.contract.UnpackLog(event, "BondReleased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KeepBondingBondSeizedIterator is returned from FilterBondSeized and is used to iterate over the raw logs and unpacked data for BondSeized events raised by the KeepBonding contract.
type KeepBondingBondSeizedIterator struct {
	Event *KeepBondingBondSeized // Event containing the contract specifics and raw log

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
func (it *KeepBondingBondSeizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeepBondingBondSeized)
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
		it.Event = new(KeepBondingBondSeized)
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
func (it *KeepBondingBondSeizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeepBondingBondSeizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeepBondingBondSeized represents a BondSeized event raised by the KeepBonding contract.
type KeepBondingBondSeized struct {
	Operator    common.Address
	ReferenceID *big.Int
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBondSeized is a free log retrieval operation binding the contract event 0xf8e947b47b515d01aa96426822ddcf23a08f42d8c2dbfd65e674ba824f551382.
//
// Solidity: event BondSeized(address indexed operator, uint256 indexed referenceID, address destination, uint256 amount)
func (_KeepBonding *KeepBondingFilterer) FilterBondSeized(opts *bind.FilterOpts, operator []common.Address, referenceID []*big.Int) (*KeepBondingBondSeizedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var referenceIDRule []interface{}
	for _, referenceIDItem := range referenceID {
		referenceIDRule = append(referenceIDRule, referenceIDItem)
	}

	logs, sub, err := _KeepBonding.contract.FilterLogs(opts, "BondSeized", operatorRule, referenceIDRule)
	if err != nil {
		return nil, err
	}
	return &KeepBondingBondSeizedIterator{contract: _KeepBonding.contract, event: "BondSeized", logs: logs, sub: sub}, nil
}

// WatchBondSeized is a free log subscription operation binding the contract event 0xf8e947b47b515d01aa96426822ddcf23a08f42d8c2dbfd65e674ba824f551382.
//
// Solidity: event BondSeized(address indexed operator, uint256 indexed referenceID, address destination, uint256 amount)
func (_KeepBonding *KeepBondingFilterer) WatchBondSeized(opts *bind.WatchOpts, sink chan<- *KeepBondingBondSeized, operator []common.Address, referenceID []*big.Int) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var referenceIDRule []interface{}
	for _, referenceIDItem := range referenceID {
		referenceIDRule = append(referenceIDRule, referenceIDItem)
	}

	logs, sub, err := _KeepBonding.contract.WatchLogs(opts, "BondSeized", operatorRule, referenceIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeepBondingBondSeized)
				if err := _KeepBonding.contract.UnpackLog(event, "BondSeized", log); err != nil {
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

// ParseBondSeized is a log parse operation binding the contract event 0xf8e947b47b515d01aa96426822ddcf23a08f42d8c2dbfd65e674ba824f551382.
//
// Solidity: event BondSeized(address indexed operator, uint256 indexed referenceID, address destination, uint256 amount)
func (_KeepBonding *KeepBondingFilterer) ParseBondSeized(log types.Log) (*KeepBondingBondSeized, error) {
	event := new(KeepBondingBondSeized)
	if err := _KeepBonding.contract.UnpackLog(event, "BondSeized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KeepBondingUnbondedValueDepositedIterator is returned from FilterUnbondedValueDeposited and is used to iterate over the raw logs and unpacked data for UnbondedValueDeposited events raised by the KeepBonding contract.
type KeepBondingUnbondedValueDepositedIterator struct {
	Event *KeepBondingUnbondedValueDeposited // Event containing the contract specifics and raw log

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
func (it *KeepBondingUnbondedValueDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeepBondingUnbondedValueDeposited)
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
		it.Event = new(KeepBondingUnbondedValueDeposited)
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
func (it *KeepBondingUnbondedValueDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeepBondingUnbondedValueDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeepBondingUnbondedValueDeposited represents a UnbondedValueDeposited event raised by the KeepBonding contract.
type KeepBondingUnbondedValueDeposited struct {
	Operator    common.Address
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnbondedValueDeposited is a free log retrieval operation binding the contract event 0xfd586a32ad24d585b1f7b36ee48e66304ad7627b48b39a0ab1d8a3e84741ea2a.
//
// Solidity: event UnbondedValueDeposited(address indexed operator, address indexed beneficiary, uint256 amount)
func (_KeepBonding *KeepBondingFilterer) FilterUnbondedValueDeposited(opts *bind.FilterOpts, operator []common.Address, beneficiary []common.Address) (*KeepBondingUnbondedValueDepositedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _KeepBonding.contract.FilterLogs(opts, "UnbondedValueDeposited", operatorRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &KeepBondingUnbondedValueDepositedIterator{contract: _KeepBonding.contract, event: "UnbondedValueDeposited", logs: logs, sub: sub}, nil
}

// WatchUnbondedValueDeposited is a free log subscription operation binding the contract event 0xfd586a32ad24d585b1f7b36ee48e66304ad7627b48b39a0ab1d8a3e84741ea2a.
//
// Solidity: event UnbondedValueDeposited(address indexed operator, address indexed beneficiary, uint256 amount)
func (_KeepBonding *KeepBondingFilterer) WatchUnbondedValueDeposited(opts *bind.WatchOpts, sink chan<- *KeepBondingUnbondedValueDeposited, operator []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _KeepBonding.contract.WatchLogs(opts, "UnbondedValueDeposited", operatorRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeepBondingUnbondedValueDeposited)
				if err := _KeepBonding.contract.UnpackLog(event, "UnbondedValueDeposited", log); err != nil {
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

// ParseUnbondedValueDeposited is a log parse operation binding the contract event 0xfd586a32ad24d585b1f7b36ee48e66304ad7627b48b39a0ab1d8a3e84741ea2a.
//
// Solidity: event UnbondedValueDeposited(address indexed operator, address indexed beneficiary, uint256 amount)
func (_KeepBonding *KeepBondingFilterer) ParseUnbondedValueDeposited(log types.Log) (*KeepBondingUnbondedValueDeposited, error) {
	event := new(KeepBondingUnbondedValueDeposited)
	if err := _KeepBonding.contract.UnpackLog(event, "UnbondedValueDeposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KeepBondingUnbondedValueWithdrawnIterator is returned from FilterUnbondedValueWithdrawn and is used to iterate over the raw logs and unpacked data for UnbondedValueWithdrawn events raised by the KeepBonding contract.
type KeepBondingUnbondedValueWithdrawnIterator struct {
	Event *KeepBondingUnbondedValueWithdrawn // Event containing the contract specifics and raw log

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
func (it *KeepBondingUnbondedValueWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeepBondingUnbondedValueWithdrawn)
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
		it.Event = new(KeepBondingUnbondedValueWithdrawn)
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
func (it *KeepBondingUnbondedValueWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeepBondingUnbondedValueWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeepBondingUnbondedValueWithdrawn represents a UnbondedValueWithdrawn event raised by the KeepBonding contract.
type KeepBondingUnbondedValueWithdrawn struct {
	Operator    common.Address
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnbondedValueWithdrawn is a free log retrieval operation binding the contract event 0x5ebf1d16423ab39117c0ca9327215b5bcd423aaf7042044c87248a4423d252d9.
//
// Solidity: event UnbondedValueWithdrawn(address indexed operator, address indexed beneficiary, uint256 amount)
func (_KeepBonding *KeepBondingFilterer) FilterUnbondedValueWithdrawn(opts *bind.FilterOpts, operator []common.Address, beneficiary []common.Address) (*KeepBondingUnbondedValueWithdrawnIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _KeepBonding.contract.FilterLogs(opts, "UnbondedValueWithdrawn", operatorRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &KeepBondingUnbondedValueWithdrawnIterator{contract: _KeepBonding.contract, event: "UnbondedValueWithdrawn", logs: logs, sub: sub}, nil
}

// WatchUnbondedValueWithdrawn is a free log subscription operation binding the contract event 0x5ebf1d16423ab39117c0ca9327215b5bcd423aaf7042044c87248a4423d252d9.
//
// Solidity: event UnbondedValueWithdrawn(address indexed operator, address indexed beneficiary, uint256 amount)
func (_KeepBonding *KeepBondingFilterer) WatchUnbondedValueWithdrawn(opts *bind.WatchOpts, sink chan<- *KeepBondingUnbondedValueWithdrawn, operator []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _KeepBonding.contract.WatchLogs(opts, "UnbondedValueWithdrawn", operatorRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeepBondingUnbondedValueWithdrawn)
				if err := _KeepBonding.contract.UnpackLog(event, "UnbondedValueWithdrawn", log); err != nil {
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

// ParseUnbondedValueWithdrawn is a log parse operation binding the contract event 0x5ebf1d16423ab39117c0ca9327215b5bcd423aaf7042044c87248a4423d252d9.
//
// Solidity: event UnbondedValueWithdrawn(address indexed operator, address indexed beneficiary, uint256 amount)
func (_KeepBonding *KeepBondingFilterer) ParseUnbondedValueWithdrawn(log types.Log) (*KeepBondingUnbondedValueWithdrawn, error) {
	event := new(KeepBondingUnbondedValueWithdrawn)
	if err := _KeepBonding.contract.UnpackLog(event, "UnbondedValueWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

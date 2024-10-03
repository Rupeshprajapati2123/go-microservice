// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// PayableTransaction is an auto generated low-level Go binding around an user-defined struct.
type PayableTransaction struct {
	To               common.Address
	Amount           *big.Int
	Approvals        []common.Address
	Disapprovals     []common.Address
	Executed         bool
	Timestamp        *big.Int
	ApprovalCount    *big.Int
	DisapprovalCount *big.Int
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_approver\",\"type\":\"address\"}],\"name\":\"AppliedForApprovers\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"}],\"name\":\"ApproverApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"TransactionApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransactionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"disapprover\",\"type\":\"address\"}],\"name\":\"TransactionDisapproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransactionFailed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"addApprovers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"appliedApproversList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adress\",\"type\":\"address\"}],\"name\":\"applyForApprover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adres\",\"type\":\"address\"}],\"name\":\"approveApprover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"approveTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approversList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"createTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"disapproveTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTransactions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"approvals\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"disapprovals\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"approvalCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"disapprovalCount\",\"type\":\"uint256\"}],\"internalType\":\"structPayable.Transaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ad\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getTransactionDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"approvals\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"disapprovals\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"approvalCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"disapprovalCount\",\"type\":\"uint256\"}],\"internalType\":\"structPayable.Transaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receivEthere\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}
var StoreBin = StoreMetaData.Bin
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _version string) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// AppliedApproversList is a free data retrieval call binding the contract method 0x35b0d745.
//
// Solidity: function appliedApproversList() view returns(address[])
func (_Store *StoreCaller) AppliedApproversList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "appliedApproversList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AppliedApproversList is a free data retrieval call binding the contract method 0x35b0d745.
//
// Solidity: function appliedApproversList() view returns(address[])
func (_Store *StoreSession) AppliedApproversList() ([]common.Address, error) {
	return _Store.Contract.AppliedApproversList(&_Store.CallOpts)
}

// AppliedApproversList is a free data retrieval call binding the contract method 0x35b0d745.
//
// Solidity: function appliedApproversList() view returns(address[])
func (_Store *StoreCallerSession) AppliedApproversList() ([]common.Address, error) {
	return _Store.Contract.AppliedApproversList(&_Store.CallOpts)
}

// ApproversList is a free data retrieval call binding the contract method 0xbe2af9cb.
//
// Solidity: function approversList() view returns(address[])
func (_Store *StoreCaller) ApproversList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "approversList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ApproversList is a free data retrieval call binding the contract method 0xbe2af9cb.
//
// Solidity: function approversList() view returns(address[])
func (_Store *StoreSession) ApproversList() ([]common.Address, error) {
	return _Store.Contract.ApproversList(&_Store.CallOpts)
}

// ApproversList is a free data retrieval call binding the contract method 0xbe2af9cb.
//
// Solidity: function approversList() view returns(address[])
func (_Store *StoreCallerSession) ApproversList() ([]common.Address, error) {
	return _Store.Contract.ApproversList(&_Store.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Store *StoreCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Store *StoreSession) Balance() (*big.Int, error) {
	return _Store.Contract.Balance(&_Store.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Store *StoreCallerSession) Balance() (*big.Int, error) {
	return _Store.Contract.Balance(&_Store.CallOpts)
}

// GetAllTransactions is a free data retrieval call binding the contract method 0x27506f53.
//
// Solidity: function getAllTransactions() view returns((address,uint256,address[],address[],bool,uint256,uint256,uint256)[])
func (_Store *StoreCaller) GetAllTransactions(opts *bind.CallOpts) ([]PayableTransaction, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getAllTransactions")

	if err != nil {
		return *new([]PayableTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]PayableTransaction)).(*[]PayableTransaction)

	return out0, err

}

// GetAllTransactions is a free data retrieval call binding the contract method 0x27506f53.
//
// Solidity: function getAllTransactions() view returns((address,uint256,address[],address[],bool,uint256,uint256,uint256)[])
func (_Store *StoreSession) GetAllTransactions() ([]PayableTransaction, error) {
	return _Store.Contract.GetAllTransactions(&_Store.CallOpts)
}

// GetAllTransactions is a free data retrieval call binding the contract method 0x27506f53.
//
// Solidity: function getAllTransactions() view returns((address,uint256,address[],address[],bool,uint256,uint256,uint256)[])
func (_Store *StoreCallerSession) GetAllTransactions() ([]PayableTransaction, error) {
	return _Store.Contract.GetAllTransactions(&_Store.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address ad) view returns(uint256)
func (_Store *StoreCaller) GetBalance(opts *bind.CallOpts, ad common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getBalance", ad)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address ad) view returns(uint256)
func (_Store *StoreSession) GetBalance(ad common.Address) (*big.Int, error) {
	return _Store.Contract.GetBalance(&_Store.CallOpts, ad)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address ad) view returns(uint256)
func (_Store *StoreCallerSession) GetBalance(ad common.Address) (*big.Int, error) {
	return _Store.Contract.GetBalance(&_Store.CallOpts, ad)
}

// GetTransactionDetails is a free data retrieval call binding the contract method 0x0fa683d3.
//
// Solidity: function getTransactionDetails(uint256 transactionId) view returns((address,uint256,address[],address[],bool,uint256,uint256,uint256))
func (_Store *StoreCaller) GetTransactionDetails(opts *bind.CallOpts, transactionId *big.Int) (PayableTransaction, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getTransactionDetails", transactionId)

	if err != nil {
		return *new(PayableTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(PayableTransaction)).(*PayableTransaction)

	return out0, err

}

// GetTransactionDetails is a free data retrieval call binding the contract method 0x0fa683d3.
//
// Solidity: function getTransactionDetails(uint256 transactionId) view returns((address,uint256,address[],address[],bool,uint256,uint256,uint256))
func (_Store *StoreSession) GetTransactionDetails(transactionId *big.Int) (PayableTransaction, error) {
	return _Store.Contract.GetTransactionDetails(&_Store.CallOpts, transactionId)
}

// GetTransactionDetails is a free data retrieval call binding the contract method 0x0fa683d3.
//
// Solidity: function getTransactionDetails(uint256 transactionId) view returns((address,uint256,address[],address[],bool,uint256,uint256,uint256))
func (_Store *StoreCallerSession) GetTransactionDetails(transactionId *big.Int) (PayableTransaction, error) {
	return _Store.Contract.GetTransactionDetails(&_Store.CallOpts, transactionId)
}

// AddApprovers is a paid mutator transaction binding the contract method 0x6a882d51.
//
// Solidity: function addApprovers(address[] addresses) returns()
func (_Store *StoreTransactor) AddApprovers(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addApprovers", addresses)
}

// AddApprovers is a paid mutator transaction binding the contract method 0x6a882d51.
//
// Solidity: function addApprovers(address[] addresses) returns()
func (_Store *StoreSession) AddApprovers(addresses []common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddApprovers(&_Store.TransactOpts, addresses)
}

// AddApprovers is a paid mutator transaction binding the contract method 0x6a882d51.
//
// Solidity: function addApprovers(address[] addresses) returns()
func (_Store *StoreTransactorSession) AddApprovers(addresses []common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddApprovers(&_Store.TransactOpts, addresses)
}

// ApplyForApprover is a paid mutator transaction binding the contract method 0xc82830ea.
//
// Solidity: function applyForApprover(address adress) returns()
func (_Store *StoreTransactor) ApplyForApprover(opts *bind.TransactOpts, adress common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "applyForApprover", adress)
}

// ApplyForApprover is a paid mutator transaction binding the contract method 0xc82830ea.
//
// Solidity: function applyForApprover(address adress) returns()
func (_Store *StoreSession) ApplyForApprover(adress common.Address) (*types.Transaction, error) {
	return _Store.Contract.ApplyForApprover(&_Store.TransactOpts, adress)
}

// ApplyForApprover is a paid mutator transaction binding the contract method 0xc82830ea.
//
// Solidity: function applyForApprover(address adress) returns()
func (_Store *StoreTransactorSession) ApplyForApprover(adress common.Address) (*types.Transaction, error) {
	return _Store.Contract.ApplyForApprover(&_Store.TransactOpts, adress)
}

// ApproveApprover is a paid mutator transaction binding the contract method 0x838d4cd1.
//
// Solidity: function approveApprover(address adres) returns()
func (_Store *StoreTransactor) ApproveApprover(opts *bind.TransactOpts, adres common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "approveApprover", adres)
}

// ApproveApprover is a paid mutator transaction binding the contract method 0x838d4cd1.
//
// Solidity: function approveApprover(address adres) returns()
func (_Store *StoreSession) ApproveApprover(adres common.Address) (*types.Transaction, error) {
	return _Store.Contract.ApproveApprover(&_Store.TransactOpts, adres)
}

// ApproveApprover is a paid mutator transaction binding the contract method 0x838d4cd1.
//
// Solidity: function approveApprover(address adres) returns()
func (_Store *StoreTransactorSession) ApproveApprover(adres common.Address) (*types.Transaction, error) {
	return _Store.Contract.ApproveApprover(&_Store.TransactOpts, adres)
}

// ApproveTransaction is a paid mutator transaction binding the contract method 0x242232d1.
//
// Solidity: function approveTransaction(uint256 transactionId) returns()
func (_Store *StoreTransactor) ApproveTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "approveTransaction", transactionId)
}

// ApproveTransaction is a paid mutator transaction binding the contract method 0x242232d1.
//
// Solidity: function approveTransaction(uint256 transactionId) returns()
func (_Store *StoreSession) ApproveTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ApproveTransaction(&_Store.TransactOpts, transactionId)
}

// ApproveTransaction is a paid mutator transaction binding the contract method 0x242232d1.
//
// Solidity: function approveTransaction(uint256 transactionId) returns()
func (_Store *StoreTransactorSession) ApproveTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ApproveTransaction(&_Store.TransactOpts, transactionId)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0x1e7f8f0f.
//
// Solidity: function createTransaction(address to, uint256 _amount) returns(uint256)
func (_Store *StoreTransactor) CreateTransaction(opts *bind.TransactOpts, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "createTransaction", to, _amount)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0x1e7f8f0f.
//
// Solidity: function createTransaction(address to, uint256 _amount) returns(uint256)
func (_Store *StoreSession) CreateTransaction(to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.CreateTransaction(&_Store.TransactOpts, to, _amount)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0x1e7f8f0f.
//
// Solidity: function createTransaction(address to, uint256 _amount) returns(uint256)
func (_Store *StoreTransactorSession) CreateTransaction(to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.CreateTransaction(&_Store.TransactOpts, to, _amount)
}

// DisapproveTransaction is a paid mutator transaction binding the contract method 0x9911a149.
//
// Solidity: function disapproveTransaction(uint256 transactionId) returns()
func (_Store *StoreTransactor) DisapproveTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "disapproveTransaction", transactionId)
}

// DisapproveTransaction is a paid mutator transaction binding the contract method 0x9911a149.
//
// Solidity: function disapproveTransaction(uint256 transactionId) returns()
func (_Store *StoreSession) DisapproveTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DisapproveTransaction(&_Store.TransactOpts, transactionId)
}

// DisapproveTransaction is a paid mutator transaction binding the contract method 0x9911a149.
//
// Solidity: function disapproveTransaction(uint256 transactionId) returns()
func (_Store *StoreTransactorSession) DisapproveTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DisapproveTransaction(&_Store.TransactOpts, transactionId)
}

// ReceivEthere is a paid mutator transaction binding the contract method 0xf2c7aded.
//
// Solidity: function receivEthere() payable returns()
func (_Store *StoreTransactor) ReceivEthere(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "receivEthere")
}

// ReceivEthere is a paid mutator transaction binding the contract method 0xf2c7aded.
//
// Solidity: function receivEthere() payable returns()
func (_Store *StoreSession) ReceivEthere() (*types.Transaction, error) {
	return _Store.Contract.ReceivEthere(&_Store.TransactOpts)
}

// ReceivEthere is a paid mutator transaction binding the contract method 0xf2c7aded.
//
// Solidity: function receivEthere() payable returns()
func (_Store *StoreTransactorSession) ReceivEthere() (*types.Transaction, error) {
	return _Store.Contract.ReceivEthere(&_Store.TransactOpts)
}

// StoreAppliedForApproversIterator is returned from FilterAppliedForApprovers and is used to iterate over the raw logs and unpacked data for AppliedForApprovers events raised by the Store contract.
type StoreAppliedForApproversIterator struct {
	Event *StoreAppliedForApprovers // Event containing the contract specifics and raw log

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
func (it *StoreAppliedForApproversIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreAppliedForApprovers)
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
		it.Event = new(StoreAppliedForApprovers)
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
func (it *StoreAppliedForApproversIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreAppliedForApproversIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreAppliedForApprovers represents a AppliedForApprovers event raised by the Store contract.
type StoreAppliedForApprovers struct {
	Approver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAppliedForApprovers is a free log retrieval operation binding the contract event 0x626e4500941465f8684ab220d11d81f3a776c4ed7010715f25cd8e772dcbb484.
//
// Solidity: event AppliedForApprovers(address _approver)
func (_Store *StoreFilterer) FilterAppliedForApprovers(opts *bind.FilterOpts) (*StoreAppliedForApproversIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "AppliedForApprovers")
	if err != nil {
		return nil, err
	}
	return &StoreAppliedForApproversIterator{contract: _Store.contract, event: "AppliedForApprovers", logs: logs, sub: sub}, nil
}

// WatchAppliedForApprovers is a free log subscription operation binding the contract event 0x626e4500941465f8684ab220d11d81f3a776c4ed7010715f25cd8e772dcbb484.
//
// Solidity: event AppliedForApprovers(address _approver)
func (_Store *StoreFilterer) WatchAppliedForApprovers(opts *bind.WatchOpts, sink chan<- *StoreAppliedForApprovers) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "AppliedForApprovers")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreAppliedForApprovers)
				if err := _Store.contract.UnpackLog(event, "AppliedForApprovers", log); err != nil {
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

// ParseAppliedForApprovers is a log parse operation binding the contract event 0x626e4500941465f8684ab220d11d81f3a776c4ed7010715f25cd8e772dcbb484.
//
// Solidity: event AppliedForApprovers(address _approver)
func (_Store *StoreFilterer) ParseAppliedForApprovers(log types.Log) (*StoreAppliedForApprovers, error) {
	event := new(StoreAppliedForApprovers)
	if err := _Store.contract.UnpackLog(event, "AppliedForApprovers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreApproverApprovedIterator is returned from FilterApproverApproved and is used to iterate over the raw logs and unpacked data for ApproverApproved events raised by the Store contract.
type StoreApproverApprovedIterator struct {
	Event *StoreApproverApproved // Event containing the contract specifics and raw log

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
func (it *StoreApproverApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreApproverApproved)
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
		it.Event = new(StoreApproverApproved)
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
func (it *StoreApproverApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreApproverApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreApproverApproved represents a ApproverApproved event raised by the Store contract.
type StoreApproverApproved struct {
	Approved common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproverApproved is a free log retrieval operation binding the contract event 0x8ffbf603707f31ab466f6b711c89ae5521b6f93b25d3f86200054018476696d6.
//
// Solidity: event ApproverApproved(address approved)
func (_Store *StoreFilterer) FilterApproverApproved(opts *bind.FilterOpts) (*StoreApproverApprovedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ApproverApproved")
	if err != nil {
		return nil, err
	}
	return &StoreApproverApprovedIterator{contract: _Store.contract, event: "ApproverApproved", logs: logs, sub: sub}, nil
}

// WatchApproverApproved is a free log subscription operation binding the contract event 0x8ffbf603707f31ab466f6b711c89ae5521b6f93b25d3f86200054018476696d6.
//
// Solidity: event ApproverApproved(address approved)
func (_Store *StoreFilterer) WatchApproverApproved(opts *bind.WatchOpts, sink chan<- *StoreApproverApproved) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ApproverApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreApproverApproved)
				if err := _Store.contract.UnpackLog(event, "ApproverApproved", log); err != nil {
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

// ParseApproverApproved is a log parse operation binding the contract event 0x8ffbf603707f31ab466f6b711c89ae5521b6f93b25d3f86200054018476696d6.
//
// Solidity: event ApproverApproved(address approved)
func (_Store *StoreFilterer) ParseApproverApproved(log types.Log) (*StoreApproverApproved, error) {
	event := new(StoreApproverApproved)
	if err := _Store.contract.UnpackLog(event, "ApproverApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreFundsTransferredIterator is returned from FilterFundsTransferred and is used to iterate over the raw logs and unpacked data for FundsTransferred events raised by the Store contract.
type StoreFundsTransferredIterator struct {
	Event *StoreFundsTransferred // Event containing the contract specifics and raw log

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
func (it *StoreFundsTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreFundsTransferred)
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
		it.Event = new(StoreFundsTransferred)
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
func (it *StoreFundsTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreFundsTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreFundsTransferred represents a FundsTransferred event raised by the Store contract.
type StoreFundsTransferred struct {
	Id     *big.Int
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsTransferred is a free log retrieval operation binding the contract event 0xe90ae330e5aefc2a73ddff0dd88bcaf8f4a0629d6553a2a85b22cac3d5f4a19e.
//
// Solidity: event FundsTransferred(uint256 id, address to, uint256 amount)
func (_Store *StoreFilterer) FilterFundsTransferred(opts *bind.FilterOpts) (*StoreFundsTransferredIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "FundsTransferred")
	if err != nil {
		return nil, err
	}
	return &StoreFundsTransferredIterator{contract: _Store.contract, event: "FundsTransferred", logs: logs, sub: sub}, nil
}

// WatchFundsTransferred is a free log subscription operation binding the contract event 0xe90ae330e5aefc2a73ddff0dd88bcaf8f4a0629d6553a2a85b22cac3d5f4a19e.
//
// Solidity: event FundsTransferred(uint256 id, address to, uint256 amount)
func (_Store *StoreFilterer) WatchFundsTransferred(opts *bind.WatchOpts, sink chan<- *StoreFundsTransferred) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "FundsTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreFundsTransferred)
				if err := _Store.contract.UnpackLog(event, "FundsTransferred", log); err != nil {
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

// ParseFundsTransferred is a log parse operation binding the contract event 0xe90ae330e5aefc2a73ddff0dd88bcaf8f4a0629d6553a2a85b22cac3d5f4a19e.
//
// Solidity: event FundsTransferred(uint256 id, address to, uint256 amount)
func (_Store *StoreFilterer) ParseFundsTransferred(log types.Log) (*StoreFundsTransferred, error) {
	event := new(StoreFundsTransferred)
	if err := _Store.contract.UnpackLog(event, "FundsTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreTransactionApprovedIterator is returned from FilterTransactionApproved and is used to iterate over the raw logs and unpacked data for TransactionApproved events raised by the Store contract.
type StoreTransactionApprovedIterator struct {
	Event *StoreTransactionApproved // Event containing the contract specifics and raw log

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
func (it *StoreTransactionApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreTransactionApproved)
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
		it.Event = new(StoreTransactionApproved)
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
func (it *StoreTransactionApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreTransactionApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreTransactionApproved represents a TransactionApproved event raised by the Store contract.
type StoreTransactionApproved struct {
	Id       *big.Int
	Approver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransactionApproved is a free log retrieval operation binding the contract event 0x924813d717e221b5f46dcd8a56da1679e4612584ab3237d55e5faabf6f6a3079.
//
// Solidity: event TransactionApproved(uint256 id, address approver)
func (_Store *StoreFilterer) FilterTransactionApproved(opts *bind.FilterOpts) (*StoreTransactionApprovedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "TransactionApproved")
	if err != nil {
		return nil, err
	}
	return &StoreTransactionApprovedIterator{contract: _Store.contract, event: "TransactionApproved", logs: logs, sub: sub}, nil
}

// WatchTransactionApproved is a free log subscription operation binding the contract event 0x924813d717e221b5f46dcd8a56da1679e4612584ab3237d55e5faabf6f6a3079.
//
// Solidity: event TransactionApproved(uint256 id, address approver)
func (_Store *StoreFilterer) WatchTransactionApproved(opts *bind.WatchOpts, sink chan<- *StoreTransactionApproved) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "TransactionApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreTransactionApproved)
				if err := _Store.contract.UnpackLog(event, "TransactionApproved", log); err != nil {
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

// ParseTransactionApproved is a log parse operation binding the contract event 0x924813d717e221b5f46dcd8a56da1679e4612584ab3237d55e5faabf6f6a3079.
//
// Solidity: event TransactionApproved(uint256 id, address approver)
func (_Store *StoreFilterer) ParseTransactionApproved(log types.Log) (*StoreTransactionApproved, error) {
	event := new(StoreTransactionApproved)
	if err := _Store.contract.UnpackLog(event, "TransactionApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreTransactionCreatedIterator is returned from FilterTransactionCreated and is used to iterate over the raw logs and unpacked data for TransactionCreated events raised by the Store contract.
type StoreTransactionCreatedIterator struct {
	Event *StoreTransactionCreated // Event containing the contract specifics and raw log

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
func (it *StoreTransactionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreTransactionCreated)
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
		it.Event = new(StoreTransactionCreated)
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
func (it *StoreTransactionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreTransactionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreTransactionCreated represents a TransactionCreated event raised by the Store contract.
type StoreTransactionCreated struct {
	Id     *big.Int
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransactionCreated is a free log retrieval operation binding the contract event 0xfd27d176b6ebb21e0182c0e43df5818e0d632867e1eb47c383c58d221205fce3.
//
// Solidity: event TransactionCreated(uint256 id, address to, uint256 amount)
func (_Store *StoreFilterer) FilterTransactionCreated(opts *bind.FilterOpts) (*StoreTransactionCreatedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "TransactionCreated")
	if err != nil {
		return nil, err
	}
	return &StoreTransactionCreatedIterator{contract: _Store.contract, event: "TransactionCreated", logs: logs, sub: sub}, nil
}

// WatchTransactionCreated is a free log subscription operation binding the contract event 0xfd27d176b6ebb21e0182c0e43df5818e0d632867e1eb47c383c58d221205fce3.
//
// Solidity: event TransactionCreated(uint256 id, address to, uint256 amount)
func (_Store *StoreFilterer) WatchTransactionCreated(opts *bind.WatchOpts, sink chan<- *StoreTransactionCreated) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "TransactionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreTransactionCreated)
				if err := _Store.contract.UnpackLog(event, "TransactionCreated", log); err != nil {
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

// ParseTransactionCreated is a log parse operation binding the contract event 0xfd27d176b6ebb21e0182c0e43df5818e0d632867e1eb47c383c58d221205fce3.
//
// Solidity: event TransactionCreated(uint256 id, address to, uint256 amount)
func (_Store *StoreFilterer) ParseTransactionCreated(log types.Log) (*StoreTransactionCreated, error) {
	event := new(StoreTransactionCreated)
	if err := _Store.contract.UnpackLog(event, "TransactionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreTransactionDisapprovedIterator is returned from FilterTransactionDisapproved and is used to iterate over the raw logs and unpacked data for TransactionDisapproved events raised by the Store contract.
type StoreTransactionDisapprovedIterator struct {
	Event *StoreTransactionDisapproved // Event containing the contract specifics and raw log

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
func (it *StoreTransactionDisapprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreTransactionDisapproved)
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
		it.Event = new(StoreTransactionDisapproved)
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
func (it *StoreTransactionDisapprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreTransactionDisapprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreTransactionDisapproved represents a TransactionDisapproved event raised by the Store contract.
type StoreTransactionDisapproved struct {
	Id          *big.Int
	Disapprover common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransactionDisapproved is a free log retrieval operation binding the contract event 0x51e2a629616d21b16f9d1d2549b4b88aba4d3b132f12a47ef7696e4a4263d1e8.
//
// Solidity: event TransactionDisapproved(uint256 id, address disapprover)
func (_Store *StoreFilterer) FilterTransactionDisapproved(opts *bind.FilterOpts) (*StoreTransactionDisapprovedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "TransactionDisapproved")
	if err != nil {
		return nil, err
	}
	return &StoreTransactionDisapprovedIterator{contract: _Store.contract, event: "TransactionDisapproved", logs: logs, sub: sub}, nil
}

// WatchTransactionDisapproved is a free log subscription operation binding the contract event 0x51e2a629616d21b16f9d1d2549b4b88aba4d3b132f12a47ef7696e4a4263d1e8.
//
// Solidity: event TransactionDisapproved(uint256 id, address disapprover)
func (_Store *StoreFilterer) WatchTransactionDisapproved(opts *bind.WatchOpts, sink chan<- *StoreTransactionDisapproved) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "TransactionDisapproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreTransactionDisapproved)
				if err := _Store.contract.UnpackLog(event, "TransactionDisapproved", log); err != nil {
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

// ParseTransactionDisapproved is a log parse operation binding the contract event 0x51e2a629616d21b16f9d1d2549b4b88aba4d3b132f12a47ef7696e4a4263d1e8.
//
// Solidity: event TransactionDisapproved(uint256 id, address disapprover)
func (_Store *StoreFilterer) ParseTransactionDisapproved(log types.Log) (*StoreTransactionDisapproved, error) {
	event := new(StoreTransactionDisapproved)
	if err := _Store.contract.UnpackLog(event, "TransactionDisapproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreTransactionFailedIterator is returned from FilterTransactionFailed and is used to iterate over the raw logs and unpacked data for TransactionFailed events raised by the Store contract.
type StoreTransactionFailedIterator struct {
	Event *StoreTransactionFailed // Event containing the contract specifics and raw log

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
func (it *StoreTransactionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreTransactionFailed)
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
		it.Event = new(StoreTransactionFailed)
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
func (it *StoreTransactionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreTransactionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreTransactionFailed represents a TransactionFailed event raised by the Store contract.
type StoreTransactionFailed struct {
	Id     *big.Int
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransactionFailed is a free log retrieval operation binding the contract event 0xe015d34eef89d312e7f4ce9ac2b7d0a952a1a5329981aa762e2a64ccfd88dcc3.
//
// Solidity: event TransactionFailed(uint256 id, address to, uint256 amount)
func (_Store *StoreFilterer) FilterTransactionFailed(opts *bind.FilterOpts) (*StoreTransactionFailedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "TransactionFailed")
	if err != nil {
		return nil, err
	}
	return &StoreTransactionFailedIterator{contract: _Store.contract, event: "TransactionFailed", logs: logs, sub: sub}, nil
}

// WatchTransactionFailed is a free log subscription operation binding the contract event 0xe015d34eef89d312e7f4ce9ac2b7d0a952a1a5329981aa762e2a64ccfd88dcc3.
//
// Solidity: event TransactionFailed(uint256 id, address to, uint256 amount)
func (_Store *StoreFilterer) WatchTransactionFailed(opts *bind.WatchOpts, sink chan<- *StoreTransactionFailed) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "TransactionFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreTransactionFailed)
				if err := _Store.contract.UnpackLog(event, "TransactionFailed", log); err != nil {
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

// ParseTransactionFailed is a log parse operation binding the contract event 0xe015d34eef89d312e7f4ce9ac2b7d0a952a1a5329981aa762e2a64ccfd88dcc3.
//
// Solidity: event TransactionFailed(uint256 id, address to, uint256 amount)
func (_Store *StoreFilterer) ParseTransactionFailed(log types.Log) (*StoreTransactionFailed, error) {
	event := new(StoreTransactionFailed)
	if err := _Store.contract.UnpackLog(event, "TransactionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

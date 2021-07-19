// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_fis_drop

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FisDropREthABI is the input ABI used to generate the binding from.
const FisDropREthABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_FIS\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"initialDroppers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"initialThreshold\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FIS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_proposals\",\"outputs\":[{\"internalType\":\"enumFisDropREth.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"_yesVotes\",\"type\":\"uint40\"},{\"internalType\":\"uint8\",\"name\":\"_yesVotesTotal\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dateDrop\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dropper\",\"type\":\"address\"}],\"name\":\"addDropper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dropper\",\"type\":\"address\"}],\"name\":\"removeDropper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dropper\",\"type\":\"address\"}],\"name\":\"getDropperIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FisDropREth is an auto generated Go binding around an Ethereum contract.
type FisDropREth struct {
	FisDropREthCaller     // Read-only binding to the contract
	FisDropREthTransactor // Write-only binding to the contract
	FisDropREthFilterer   // Log filterer for contract events
}

// FisDropREthCaller is an auto generated read-only Go binding around an Ethereum contract.
type FisDropREthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FisDropREthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FisDropREthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FisDropREthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FisDropREthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FisDropREthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FisDropREthSession struct {
	Contract     *FisDropREth      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FisDropREthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FisDropREthCallerSession struct {
	Contract *FisDropREthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FisDropREthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FisDropREthTransactorSession struct {
	Contract     *FisDropREthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FisDropREthRaw is an auto generated low-level Go binding around an Ethereum contract.
type FisDropREthRaw struct {
	Contract *FisDropREth // Generic contract binding to access the raw methods on
}

// FisDropREthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FisDropREthCallerRaw struct {
	Contract *FisDropREthCaller // Generic read-only contract binding to access the raw methods on
}

// FisDropREthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FisDropREthTransactorRaw struct {
	Contract *FisDropREthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFisDropREth creates a new instance of FisDropREth, bound to a specific deployed contract.
func NewFisDropREth(address common.Address, backend bind.ContractBackend) (*FisDropREth, error) {
	contract, err := bindFisDropREth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FisDropREth{FisDropREthCaller: FisDropREthCaller{contract: contract}, FisDropREthTransactor: FisDropREthTransactor{contract: contract}, FisDropREthFilterer: FisDropREthFilterer{contract: contract}}, nil
}

// NewFisDropREthCaller creates a new read-only instance of FisDropREth, bound to a specific deployed contract.
func NewFisDropREthCaller(address common.Address, caller bind.ContractCaller) (*FisDropREthCaller, error) {
	contract, err := bindFisDropREth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FisDropREthCaller{contract: contract}, nil
}

// NewFisDropREthTransactor creates a new write-only instance of FisDropREth, bound to a specific deployed contract.
func NewFisDropREthTransactor(address common.Address, transactor bind.ContractTransactor) (*FisDropREthTransactor, error) {
	contract, err := bindFisDropREth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FisDropREthTransactor{contract: contract}, nil
}

// NewFisDropREthFilterer creates a new log filterer instance of FisDropREth, bound to a specific deployed contract.
func NewFisDropREthFilterer(address common.Address, filterer bind.ContractFilterer) (*FisDropREthFilterer, error) {
	contract, err := bindFisDropREth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FisDropREthFilterer{contract: contract}, nil
}

// bindFisDropREth binds a generic wrapper to an already deployed contract.
func bindFisDropREth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FisDropREthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FisDropREth *FisDropREthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FisDropREth.Contract.FisDropREthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FisDropREth *FisDropREthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FisDropREth.Contract.FisDropREthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FisDropREth *FisDropREthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FisDropREth.Contract.FisDropREthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FisDropREth *FisDropREthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FisDropREth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FisDropREth *FisDropREthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FisDropREth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FisDropREth *FisDropREthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FisDropREth.Contract.contract.Transact(opts, method, params...)
}

// FIS is a free data retrieval call binding the contract method 0xa4b8ab76.
//
// Solidity: function FIS() view returns(address)
func (_FisDropREth *FisDropREthCaller) FIS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FisDropREth.contract.Call(opts, &out, "FIS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FIS is a free data retrieval call binding the contract method 0xa4b8ab76.
//
// Solidity: function FIS() view returns(address)
func (_FisDropREth *FisDropREthSession) FIS() (common.Address, error) {
	return _FisDropREth.Contract.FIS(&_FisDropREth.CallOpts)
}

// FIS is a free data retrieval call binding the contract method 0xa4b8ab76.
//
// Solidity: function FIS() view returns(address)
func (_FisDropREth *FisDropREthCallerSession) FIS() (common.Address, error) {
	return _FisDropREth.Contract.FIS(&_FisDropREth.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0xf2e2af94.
//
// Solidity: function _proposals(bytes32 ) view returns(uint8 _status, uint40 _yesVotes, uint8 _yesVotesTotal)
func (_FisDropREth *FisDropREthCaller) Proposals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status        uint8
	YesVotes      *big.Int
	YesVotesTotal uint8
}, error) {
	var out []interface{}
	err := _FisDropREth.contract.Call(opts, &out, "_proposals", arg0)

	outstruct := new(struct {
		Status        uint8
		YesVotes      *big.Int
		YesVotesTotal uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.YesVotes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.YesVotesTotal = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0xf2e2af94.
//
// Solidity: function _proposals(bytes32 ) view returns(uint8 _status, uint40 _yesVotes, uint8 _yesVotesTotal)
func (_FisDropREth *FisDropREthSession) Proposals(arg0 [32]byte) (struct {
	Status        uint8
	YesVotes      *big.Int
	YesVotesTotal uint8
}, error) {
	return _FisDropREth.Contract.Proposals(&_FisDropREth.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0xf2e2af94.
//
// Solidity: function _proposals(bytes32 ) view returns(uint8 _status, uint40 _yesVotes, uint8 _yesVotesTotal)
func (_FisDropREth *FisDropREthCallerSession) Proposals(arg0 [32]byte) (struct {
	Status        uint8
	YesVotes      *big.Int
	YesVotesTotal uint8
}, error) {
	return _FisDropREth.Contract.Proposals(&_FisDropREth.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0x7f3c8160.
//
// Solidity: function _threshold() view returns(uint8)
func (_FisDropREth *FisDropREthCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FisDropREth.contract.Call(opts, &out, "_threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x7f3c8160.
//
// Solidity: function _threshold() view returns(uint8)
func (_FisDropREth *FisDropREthSession) Threshold() (uint8, error) {
	return _FisDropREth.Contract.Threshold(&_FisDropREth.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x7f3c8160.
//
// Solidity: function _threshold() view returns(uint8)
func (_FisDropREth *FisDropREthCallerSession) Threshold() (uint8, error) {
	return _FisDropREth.Contract.Threshold(&_FisDropREth.CallOpts)
}

// ClaimOpen is a free data retrieval call binding the contract method 0x4b8bcb58.
//
// Solidity: function claimOpen() view returns(bool)
func (_FisDropREth *FisDropREthCaller) ClaimOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FisDropREth.contract.Call(opts, &out, "claimOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClaimOpen is a free data retrieval call binding the contract method 0x4b8bcb58.
//
// Solidity: function claimOpen() view returns(bool)
func (_FisDropREth *FisDropREthSession) ClaimOpen() (bool, error) {
	return _FisDropREth.Contract.ClaimOpen(&_FisDropREth.CallOpts)
}

// ClaimOpen is a free data retrieval call binding the contract method 0x4b8bcb58.
//
// Solidity: function claimOpen() view returns(bool)
func (_FisDropREth *FisDropREthCallerSession) ClaimOpen() (bool, error) {
	return _FisDropREth.Contract.ClaimOpen(&_FisDropREth.CallOpts)
}

// ClaimRound is a free data retrieval call binding the contract method 0x9c8fe8b9.
//
// Solidity: function claimRound() view returns(uint256)
func (_FisDropREth *FisDropREthCaller) ClaimRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FisDropREth.contract.Call(opts, &out, "claimRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimRound is a free data retrieval call binding the contract method 0x9c8fe8b9.
//
// Solidity: function claimRound() view returns(uint256)
func (_FisDropREth *FisDropREthSession) ClaimRound() (*big.Int, error) {
	return _FisDropREth.Contract.ClaimRound(&_FisDropREth.CallOpts)
}

// ClaimRound is a free data retrieval call binding the contract method 0x9c8fe8b9.
//
// Solidity: function claimRound() view returns(uint256)
func (_FisDropREth *FisDropREthCallerSession) ClaimRound() (*big.Int, error) {
	return _FisDropREth.Contract.ClaimRound(&_FisDropREth.CallOpts)
}

// DateDrop is a free data retrieval call binding the contract method 0x9cedd341.
//
// Solidity: function dateDrop(bytes32 ) view returns(bool)
func (_FisDropREth *FisDropREthCaller) DateDrop(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _FisDropREth.contract.Call(opts, &out, "dateDrop", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DateDrop is a free data retrieval call binding the contract method 0x9cedd341.
//
// Solidity: function dateDrop(bytes32 ) view returns(bool)
func (_FisDropREth *FisDropREthSession) DateDrop(arg0 [32]byte) (bool, error) {
	return _FisDropREth.Contract.DateDrop(&_FisDropREth.CallOpts, arg0)
}

// DateDrop is a free data retrieval call binding the contract method 0x9cedd341.
//
// Solidity: function dateDrop(bytes32 ) view returns(bool)
func (_FisDropREth *FisDropREthCallerSession) DateDrop(arg0 [32]byte) (bool, error) {
	return _FisDropREth.Contract.DateDrop(&_FisDropREth.CallOpts, arg0)
}

// GetDropperIndex is a free data retrieval call binding the contract method 0xe6422c8c.
//
// Solidity: function getDropperIndex(address dropper) view returns(uint256)
func (_FisDropREth *FisDropREthCaller) GetDropperIndex(opts *bind.CallOpts, dropper common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FisDropREth.contract.Call(opts, &out, "getDropperIndex", dropper)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDropperIndex is a free data retrieval call binding the contract method 0xe6422c8c.
//
// Solidity: function getDropperIndex(address dropper) view returns(uint256)
func (_FisDropREth *FisDropREthSession) GetDropperIndex(dropper common.Address) (*big.Int, error) {
	return _FisDropREth.Contract.GetDropperIndex(&_FisDropREth.CallOpts, dropper)
}

// GetDropperIndex is a free data retrieval call binding the contract method 0xe6422c8c.
//
// Solidity: function getDropperIndex(address dropper) view returns(uint256)
func (_FisDropREth *FisDropREthCallerSession) GetDropperIndex(dropper common.Address) (*big.Int, error) {
	return _FisDropREth.Contract.GetDropperIndex(&_FisDropREth.CallOpts, dropper)
}

// IsClaimed is a free data retrieval call binding the contract method 0xf364c90c.
//
// Solidity: function isClaimed(uint256 round, uint256 index) view returns(bool)
func (_FisDropREth *FisDropREthCaller) IsClaimed(opts *bind.CallOpts, round *big.Int, index *big.Int) (bool, error) {
	var out []interface{}
	err := _FisDropREth.contract.Call(opts, &out, "isClaimed", round, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0xf364c90c.
//
// Solidity: function isClaimed(uint256 round, uint256 index) view returns(bool)
func (_FisDropREth *FisDropREthSession) IsClaimed(round *big.Int, index *big.Int) (bool, error) {
	return _FisDropREth.Contract.IsClaimed(&_FisDropREth.CallOpts, round, index)
}

// IsClaimed is a free data retrieval call binding the contract method 0xf364c90c.
//
// Solidity: function isClaimed(uint256 round, uint256 index) view returns(bool)
func (_FisDropREth *FisDropREthCallerSession) IsClaimed(round *big.Int, index *big.Int) (bool, error) {
	return _FisDropREth.Contract.IsClaimed(&_FisDropREth.CallOpts, round, index)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_FisDropREth *FisDropREthCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FisDropREth.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_FisDropREth *FisDropREthSession) MerkleRoot() ([32]byte, error) {
	return _FisDropREth.Contract.MerkleRoot(&_FisDropREth.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_FisDropREth *FisDropREthCallerSession) MerkleRoot() ([32]byte, error) {
	return _FisDropREth.Contract.MerkleRoot(&_FisDropREth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FisDropREth *FisDropREthCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FisDropREth.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FisDropREth *FisDropREthSession) Owner() (common.Address, error) {
	return _FisDropREth.Contract.Owner(&_FisDropREth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FisDropREth *FisDropREthCallerSession) Owner() (common.Address, error) {
	return _FisDropREth.Contract.Owner(&_FisDropREth.CallOpts)
}

// AddDropper is a paid mutator transaction binding the contract method 0x9f443aa0.
//
// Solidity: function addDropper(address dropper) returns()
func (_FisDropREth *FisDropREthTransactor) AddDropper(opts *bind.TransactOpts, dropper common.Address) (*types.Transaction, error) {
	return _FisDropREth.contract.Transact(opts, "addDropper", dropper)
}

// AddDropper is a paid mutator transaction binding the contract method 0x9f443aa0.
//
// Solidity: function addDropper(address dropper) returns()
func (_FisDropREth *FisDropREthSession) AddDropper(dropper common.Address) (*types.Transaction, error) {
	return _FisDropREth.Contract.AddDropper(&_FisDropREth.TransactOpts, dropper)
}

// AddDropper is a paid mutator transaction binding the contract method 0x9f443aa0.
//
// Solidity: function addDropper(address dropper) returns()
func (_FisDropREth *FisDropREthTransactorSession) AddDropper(dropper common.Address) (*types.Transaction, error) {
	return _FisDropREth.Contract.AddDropper(&_FisDropREth.TransactOpts, dropper)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 newThreshold) returns()
func (_FisDropREth *FisDropREthTransactor) ChangeThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _FisDropREth.contract.Transact(opts, "changeThreshold", newThreshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 newThreshold) returns()
func (_FisDropREth *FisDropREthSession) ChangeThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _FisDropREth.Contract.ChangeThreshold(&_FisDropREth.TransactOpts, newThreshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 newThreshold) returns()
func (_FisDropREth *FisDropREthTransactorSession) ChangeThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _FisDropREth.Contract.ChangeThreshold(&_FisDropREth.TransactOpts, newThreshold)
}

// Claim is a paid mutator transaction binding the contract method 0x2e7ba6ef.
//
// Solidity: function claim(uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns()
func (_FisDropREth *FisDropREthTransactor) Claim(opts *bind.TransactOpts, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _FisDropREth.contract.Transact(opts, "claim", index, account, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x2e7ba6ef.
//
// Solidity: function claim(uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns()
func (_FisDropREth *FisDropREthSession) Claim(index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _FisDropREth.Contract.Claim(&_FisDropREth.TransactOpts, index, account, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x2e7ba6ef.
//
// Solidity: function claim(uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns()
func (_FisDropREth *FisDropREthTransactorSession) Claim(index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _FisDropREth.Contract.Claim(&_FisDropREth.TransactOpts, index, account, amount, merkleProof)
}

// CloseClaim is a paid mutator transaction binding the contract method 0xc2fc94a1.
//
// Solidity: function closeClaim() returns()
func (_FisDropREth *FisDropREthTransactor) CloseClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FisDropREth.contract.Transact(opts, "closeClaim")
}

// CloseClaim is a paid mutator transaction binding the contract method 0xc2fc94a1.
//
// Solidity: function closeClaim() returns()
func (_FisDropREth *FisDropREthSession) CloseClaim() (*types.Transaction, error) {
	return _FisDropREth.Contract.CloseClaim(&_FisDropREth.TransactOpts)
}

// CloseClaim is a paid mutator transaction binding the contract method 0xc2fc94a1.
//
// Solidity: function closeClaim() returns()
func (_FisDropREth *FisDropREthTransactorSession) CloseClaim() (*types.Transaction, error) {
	return _FisDropREth.Contract.CloseClaim(&_FisDropREth.TransactOpts)
}

// OpenClaim is a paid mutator transaction binding the contract method 0x293cdbf1.
//
// Solidity: function openClaim() returns()
func (_FisDropREth *FisDropREthTransactor) OpenClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FisDropREth.contract.Transact(opts, "openClaim")
}

// OpenClaim is a paid mutator transaction binding the contract method 0x293cdbf1.
//
// Solidity: function openClaim() returns()
func (_FisDropREth *FisDropREthSession) OpenClaim() (*types.Transaction, error) {
	return _FisDropREth.Contract.OpenClaim(&_FisDropREth.TransactOpts)
}

// OpenClaim is a paid mutator transaction binding the contract method 0x293cdbf1.
//
// Solidity: function openClaim() returns()
func (_FisDropREth *FisDropREthTransactorSession) OpenClaim() (*types.Transaction, error) {
	return _FisDropREth.Contract.OpenClaim(&_FisDropREth.TransactOpts)
}

// RemoveDropper is a paid mutator transaction binding the contract method 0x2d772e53.
//
// Solidity: function removeDropper(address dropper) returns()
func (_FisDropREth *FisDropREthTransactor) RemoveDropper(opts *bind.TransactOpts, dropper common.Address) (*types.Transaction, error) {
	return _FisDropREth.contract.Transact(opts, "removeDropper", dropper)
}

// RemoveDropper is a paid mutator transaction binding the contract method 0x2d772e53.
//
// Solidity: function removeDropper(address dropper) returns()
func (_FisDropREth *FisDropREthSession) RemoveDropper(dropper common.Address) (*types.Transaction, error) {
	return _FisDropREth.Contract.RemoveDropper(&_FisDropREth.TransactOpts, dropper)
}

// RemoveDropper is a paid mutator transaction binding the contract method 0x2d772e53.
//
// Solidity: function removeDropper(address dropper) returns()
func (_FisDropREth *FisDropREthTransactorSession) RemoveDropper(dropper common.Address) (*types.Transaction, error) {
	return _FisDropREth.Contract.RemoveDropper(&_FisDropREth.TransactOpts, dropper)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FisDropREth *FisDropREthTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FisDropREth.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FisDropREth *FisDropREthSession) RenounceOwnership() (*types.Transaction, error) {
	return _FisDropREth.Contract.RenounceOwnership(&_FisDropREth.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FisDropREth *FisDropREthTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FisDropREth.Contract.RenounceOwnership(&_FisDropREth.TransactOpts)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x75edcbe0.
//
// Solidity: function setMerkleRoot(bytes32 dateHash, bytes32 _merkleRoot) returns()
func (_FisDropREth *FisDropREthTransactor) SetMerkleRoot(opts *bind.TransactOpts, dateHash [32]byte, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _FisDropREth.contract.Transact(opts, "setMerkleRoot", dateHash, _merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x75edcbe0.
//
// Solidity: function setMerkleRoot(bytes32 dateHash, bytes32 _merkleRoot) returns()
func (_FisDropREth *FisDropREthSession) SetMerkleRoot(dateHash [32]byte, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _FisDropREth.Contract.SetMerkleRoot(&_FisDropREth.TransactOpts, dateHash, _merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x75edcbe0.
//
// Solidity: function setMerkleRoot(bytes32 dateHash, bytes32 _merkleRoot) returns()
func (_FisDropREth *FisDropREthTransactorSession) SetMerkleRoot(dateHash [32]byte, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _FisDropREth.Contract.SetMerkleRoot(&_FisDropREth.TransactOpts, dateHash, _merkleRoot)
}

// SetMerkleRoot0 is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_FisDropREth *FisDropREthTransactor) SetMerkleRoot0(opts *bind.TransactOpts, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _FisDropREth.contract.Transact(opts, "setMerkleRoot0", _merkleRoot)
}

// SetMerkleRoot0 is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_FisDropREth *FisDropREthSession) SetMerkleRoot0(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _FisDropREth.Contract.SetMerkleRoot0(&_FisDropREth.TransactOpts, _merkleRoot)
}

// SetMerkleRoot0 is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_FisDropREth *FisDropREthTransactorSession) SetMerkleRoot0(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _FisDropREth.Contract.SetMerkleRoot0(&_FisDropREth.TransactOpts, _merkleRoot)
}

// SwitchClaim is a paid mutator transaction binding the contract method 0x4f5aa167.
//
// Solidity: function switchClaim() returns()
func (_FisDropREth *FisDropREthTransactor) SwitchClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FisDropREth.contract.Transact(opts, "switchClaim")
}

// SwitchClaim is a paid mutator transaction binding the contract method 0x4f5aa167.
//
// Solidity: function switchClaim() returns()
func (_FisDropREth *FisDropREthSession) SwitchClaim() (*types.Transaction, error) {
	return _FisDropREth.Contract.SwitchClaim(&_FisDropREth.TransactOpts)
}

// SwitchClaim is a paid mutator transaction binding the contract method 0x4f5aa167.
//
// Solidity: function switchClaim() returns()
func (_FisDropREth *FisDropREthTransactorSession) SwitchClaim() (*types.Transaction, error) {
	return _FisDropREth.Contract.SwitchClaim(&_FisDropREth.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FisDropREth *FisDropREthTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FisDropREth.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FisDropREth *FisDropREthSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FisDropREth.Contract.TransferOwnership(&_FisDropREth.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FisDropREth *FisDropREthTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FisDropREth.Contract.TransferOwnership(&_FisDropREth.TransactOpts, newOwner)
}

// FisDropREthClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the FisDropREth contract.
type FisDropREthClaimedIterator struct {
	Event *FisDropREthClaimed // Event containing the contract specifics and raw log

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
func (it *FisDropREthClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FisDropREthClaimed)
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
		it.Event = new(FisDropREthClaimed)
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
func (it *FisDropREthClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FisDropREthClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FisDropREthClaimed represents a Claimed event raised by the FisDropREth contract.
type FisDropREthClaimed struct {
	Round   *big.Int
	Index   *big.Int
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xb94bf7f9302edf52a596286915a69b4b0685574cffdedd0712e3c62f2550f0ba.
//
// Solidity: event Claimed(uint256 round, uint256 index, address account, uint256 amount)
func (_FisDropREth *FisDropREthFilterer) FilterClaimed(opts *bind.FilterOpts) (*FisDropREthClaimedIterator, error) {

	logs, sub, err := _FisDropREth.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &FisDropREthClaimedIterator{contract: _FisDropREth.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xb94bf7f9302edf52a596286915a69b4b0685574cffdedd0712e3c62f2550f0ba.
//
// Solidity: event Claimed(uint256 round, uint256 index, address account, uint256 amount)
func (_FisDropREth *FisDropREthFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *FisDropREthClaimed) (event.Subscription, error) {

	logs, sub, err := _FisDropREth.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FisDropREthClaimed)
				if err := _FisDropREth.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xb94bf7f9302edf52a596286915a69b4b0685574cffdedd0712e3c62f2550f0ba.
//
// Solidity: event Claimed(uint256 round, uint256 index, address account, uint256 amount)
func (_FisDropREth *FisDropREthFilterer) ParseClaimed(log types.Log) (*FisDropREthClaimed, error) {
	event := new(FisDropREthClaimed)
	if err := _FisDropREth.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FisDropREthOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FisDropREth contract.
type FisDropREthOwnershipTransferredIterator struct {
	Event *FisDropREthOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FisDropREthOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FisDropREthOwnershipTransferred)
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
		it.Event = new(FisDropREthOwnershipTransferred)
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
func (it *FisDropREthOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FisDropREthOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FisDropREthOwnershipTransferred represents a OwnershipTransferred event raised by the FisDropREth contract.
type FisDropREthOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FisDropREth *FisDropREthFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FisDropREthOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FisDropREth.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FisDropREthOwnershipTransferredIterator{contract: _FisDropREth.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FisDropREth *FisDropREthFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FisDropREthOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FisDropREth.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FisDropREthOwnershipTransferred)
				if err := _FisDropREth.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FisDropREth *FisDropREthFilterer) ParseOwnershipTransferred(log types.Log) (*FisDropREthOwnershipTransferred, error) {
	event := new(FisDropREthOwnershipTransferred)
	if err := _FisDropREth.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crabcaller

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

// CrabcallerMetaData contains all meta data concerning the Crabcaller contract.
var CrabcallerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getActiveGameInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"defBattlePoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defTimePoint\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"attackTeamId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"atkOwner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"atkBattlePoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"atkTimePoint\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getAttackTeam\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"attackTeamId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"teamOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId3\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"battlePoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"timePoint\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"currentGameId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"lockTo\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getDefenseTeam\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"teamOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId3\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"battlePoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"timePoint\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"currentGameId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"lockTo\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getGameDefTeamInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"battlePoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"timePoint\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"teamIds\",\"type\":\"uint256[]\"}],\"name\":\"getTeamInfos\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"teamOwners\",\"type\":\"address[]\"},{\"internalType\":\"uint16[]\",\"name\":\"battlePoints\",\"type\":\"uint16[]\"},{\"internalType\":\"uint16[]\",\"name\":\"timePoints\",\"type\":\"uint16[]\"},{\"internalType\":\"uint256[]\",\"name\":\"currentGameIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128[]\",\"name\":\"lockTos\",\"type\":\"uint128[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crabadaAddress\",\"type\":\"address\"}],\"name\":\"setCrabadaAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CrabcallerABI is the input ABI used to generate the binding from.
// Deprecated: Use CrabcallerMetaData.ABI instead.
var CrabcallerABI = CrabcallerMetaData.ABI

// Crabcaller is an auto generated Go binding around an Ethereum contract.
type Crabcaller struct {
	CrabcallerCaller     // Read-only binding to the contract
	CrabcallerTransactor // Write-only binding to the contract
	CrabcallerFilterer   // Log filterer for contract events
}

// CrabcallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrabcallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrabcallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrabcallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrabcallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrabcallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrabcallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrabcallerSession struct {
	Contract     *Crabcaller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrabcallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrabcallerCallerSession struct {
	Contract *CrabcallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CrabcallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrabcallerTransactorSession struct {
	Contract     *CrabcallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CrabcallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrabcallerRaw struct {
	Contract *Crabcaller // Generic contract binding to access the raw methods on
}

// CrabcallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrabcallerCallerRaw struct {
	Contract *CrabcallerCaller // Generic read-only contract binding to access the raw methods on
}

// CrabcallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrabcallerTransactorRaw struct {
	Contract *CrabcallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrabcaller creates a new instance of Crabcaller, bound to a specific deployed contract.
func NewCrabcaller(address common.Address, backend bind.ContractBackend) (*Crabcaller, error) {
	contract, err := bindCrabcaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crabcaller{CrabcallerCaller: CrabcallerCaller{contract: contract}, CrabcallerTransactor: CrabcallerTransactor{contract: contract}, CrabcallerFilterer: CrabcallerFilterer{contract: contract}}, nil
}

// NewCrabcallerCaller creates a new read-only instance of Crabcaller, bound to a specific deployed contract.
func NewCrabcallerCaller(address common.Address, caller bind.ContractCaller) (*CrabcallerCaller, error) {
	contract, err := bindCrabcaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrabcallerCaller{contract: contract}, nil
}

// NewCrabcallerTransactor creates a new write-only instance of Crabcaller, bound to a specific deployed contract.
func NewCrabcallerTransactor(address common.Address, transactor bind.ContractTransactor) (*CrabcallerTransactor, error) {
	contract, err := bindCrabcaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrabcallerTransactor{contract: contract}, nil
}

// NewCrabcallerFilterer creates a new log filterer instance of Crabcaller, bound to a specific deployed contract.
func NewCrabcallerFilterer(address common.Address, filterer bind.ContractFilterer) (*CrabcallerFilterer, error) {
	contract, err := bindCrabcaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrabcallerFilterer{contract: contract}, nil
}

// bindCrabcaller binds a generic wrapper to an already deployed contract.
func bindCrabcaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrabcallerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crabcaller *CrabcallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crabcaller.Contract.CrabcallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crabcaller *CrabcallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crabcaller.Contract.CrabcallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crabcaller *CrabcallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crabcaller.Contract.CrabcallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crabcaller *CrabcallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crabcaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crabcaller *CrabcallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crabcaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crabcaller *CrabcallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crabcaller.Contract.contract.Transact(opts, method, params...)
}

// GetActiveGameInfo is a free data retrieval call binding the contract method 0x9319189a.
//
// Solidity: function getActiveGameInfo(uint256 gameId) view returns(uint32 startTime, uint16 defBattlePoint, uint16 defTimePoint, uint256 attackTeamId, address atkOwner, uint16 atkBattlePoint, uint16 atkTimePoint)
func (_Crabcaller *CrabcallerCaller) GetActiveGameInfo(opts *bind.CallOpts, gameId *big.Int) (struct {
	StartTime      uint32
	DefBattlePoint uint16
	DefTimePoint   uint16
	AttackTeamId   *big.Int
	AtkOwner       common.Address
	AtkBattlePoint uint16
	AtkTimePoint   uint16
}, error) {
	var out []interface{}
	err := _Crabcaller.contract.Call(opts, &out, "getActiveGameInfo", gameId)

	outstruct := new(struct {
		StartTime      uint32
		DefBattlePoint uint16
		DefTimePoint   uint16
		AttackTeamId   *big.Int
		AtkOwner       common.Address
		AtkBattlePoint uint16
		AtkTimePoint   uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTime = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.DefBattlePoint = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.DefTimePoint = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.AttackTeamId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AtkOwner = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.AtkBattlePoint = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.AtkTimePoint = *abi.ConvertType(out[6], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetActiveGameInfo is a free data retrieval call binding the contract method 0x9319189a.
//
// Solidity: function getActiveGameInfo(uint256 gameId) view returns(uint32 startTime, uint16 defBattlePoint, uint16 defTimePoint, uint256 attackTeamId, address atkOwner, uint16 atkBattlePoint, uint16 atkTimePoint)
func (_Crabcaller *CrabcallerSession) GetActiveGameInfo(gameId *big.Int) (struct {
	StartTime      uint32
	DefBattlePoint uint16
	DefTimePoint   uint16
	AttackTeamId   *big.Int
	AtkOwner       common.Address
	AtkBattlePoint uint16
	AtkTimePoint   uint16
}, error) {
	return _Crabcaller.Contract.GetActiveGameInfo(&_Crabcaller.CallOpts, gameId)
}

// GetActiveGameInfo is a free data retrieval call binding the contract method 0x9319189a.
//
// Solidity: function getActiveGameInfo(uint256 gameId) view returns(uint32 startTime, uint16 defBattlePoint, uint16 defTimePoint, uint256 attackTeamId, address atkOwner, uint16 atkBattlePoint, uint16 atkTimePoint)
func (_Crabcaller *CrabcallerCallerSession) GetActiveGameInfo(gameId *big.Int) (struct {
	StartTime      uint32
	DefBattlePoint uint16
	DefTimePoint   uint16
	AttackTeamId   *big.Int
	AtkOwner       common.Address
	AtkBattlePoint uint16
	AtkTimePoint   uint16
}, error) {
	return _Crabcaller.Contract.GetActiveGameInfo(&_Crabcaller.CallOpts, gameId)
}

// GetAttackTeam is a free data retrieval call binding the contract method 0x201a5b38.
//
// Solidity: function getAttackTeam(uint256 gameId) view returns(uint256 attackTeamId, address teamOwner, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint16 battlePoint, uint16 timePoint, uint256 currentGameId, uint128 lockTo)
func (_Crabcaller *CrabcallerCaller) GetAttackTeam(opts *bind.CallOpts, gameId *big.Int) (struct {
	AttackTeamId  *big.Int
	TeamOwner     common.Address
	CrabadaId1    *big.Int
	CrabadaId2    *big.Int
	CrabadaId3    *big.Int
	BattlePoint   uint16
	TimePoint     uint16
	CurrentGameId *big.Int
	LockTo        *big.Int
}, error) {
	var out []interface{}
	err := _Crabcaller.contract.Call(opts, &out, "getAttackTeam", gameId)

	outstruct := new(struct {
		AttackTeamId  *big.Int
		TeamOwner     common.Address
		CrabadaId1    *big.Int
		CrabadaId2    *big.Int
		CrabadaId3    *big.Int
		BattlePoint   uint16
		TimePoint     uint16
		CurrentGameId *big.Int
		LockTo        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttackTeamId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TeamOwner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CrabadaId1 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CrabadaId2 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CrabadaId3 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.BattlePoint = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.TimePoint = *abi.ConvertType(out[6], new(uint16)).(*uint16)
	outstruct.CurrentGameId = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.LockTo = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAttackTeam is a free data retrieval call binding the contract method 0x201a5b38.
//
// Solidity: function getAttackTeam(uint256 gameId) view returns(uint256 attackTeamId, address teamOwner, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint16 battlePoint, uint16 timePoint, uint256 currentGameId, uint128 lockTo)
func (_Crabcaller *CrabcallerSession) GetAttackTeam(gameId *big.Int) (struct {
	AttackTeamId  *big.Int
	TeamOwner     common.Address
	CrabadaId1    *big.Int
	CrabadaId2    *big.Int
	CrabadaId3    *big.Int
	BattlePoint   uint16
	TimePoint     uint16
	CurrentGameId *big.Int
	LockTo        *big.Int
}, error) {
	return _Crabcaller.Contract.GetAttackTeam(&_Crabcaller.CallOpts, gameId)
}

// GetAttackTeam is a free data retrieval call binding the contract method 0x201a5b38.
//
// Solidity: function getAttackTeam(uint256 gameId) view returns(uint256 attackTeamId, address teamOwner, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint16 battlePoint, uint16 timePoint, uint256 currentGameId, uint128 lockTo)
func (_Crabcaller *CrabcallerCallerSession) GetAttackTeam(gameId *big.Int) (struct {
	AttackTeamId  *big.Int
	TeamOwner     common.Address
	CrabadaId1    *big.Int
	CrabadaId2    *big.Int
	CrabadaId3    *big.Int
	BattlePoint   uint16
	TimePoint     uint16
	CurrentGameId *big.Int
	LockTo        *big.Int
}, error) {
	return _Crabcaller.Contract.GetAttackTeam(&_Crabcaller.CallOpts, gameId)
}

// GetDefenseTeam is a free data retrieval call binding the contract method 0xb8088875.
//
// Solidity: function getDefenseTeam(uint256 gameId) view returns(uint256 teamId, address teamOwner, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint16 battlePoint, uint16 timePoint, uint256 currentGameId, uint128 lockTo)
func (_Crabcaller *CrabcallerCaller) GetDefenseTeam(opts *bind.CallOpts, gameId *big.Int) (struct {
	TeamId        *big.Int
	TeamOwner     common.Address
	CrabadaId1    *big.Int
	CrabadaId2    *big.Int
	CrabadaId3    *big.Int
	BattlePoint   uint16
	TimePoint     uint16
	CurrentGameId *big.Int
	LockTo        *big.Int
}, error) {
	var out []interface{}
	err := _Crabcaller.contract.Call(opts, &out, "getDefenseTeam", gameId)

	outstruct := new(struct {
		TeamId        *big.Int
		TeamOwner     common.Address
		CrabadaId1    *big.Int
		CrabadaId2    *big.Int
		CrabadaId3    *big.Int
		BattlePoint   uint16
		TimePoint     uint16
		CurrentGameId *big.Int
		LockTo        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TeamId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TeamOwner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CrabadaId1 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CrabadaId2 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CrabadaId3 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.BattlePoint = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.TimePoint = *abi.ConvertType(out[6], new(uint16)).(*uint16)
	outstruct.CurrentGameId = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.LockTo = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDefenseTeam is a free data retrieval call binding the contract method 0xb8088875.
//
// Solidity: function getDefenseTeam(uint256 gameId) view returns(uint256 teamId, address teamOwner, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint16 battlePoint, uint16 timePoint, uint256 currentGameId, uint128 lockTo)
func (_Crabcaller *CrabcallerSession) GetDefenseTeam(gameId *big.Int) (struct {
	TeamId        *big.Int
	TeamOwner     common.Address
	CrabadaId1    *big.Int
	CrabadaId2    *big.Int
	CrabadaId3    *big.Int
	BattlePoint   uint16
	TimePoint     uint16
	CurrentGameId *big.Int
	LockTo        *big.Int
}, error) {
	return _Crabcaller.Contract.GetDefenseTeam(&_Crabcaller.CallOpts, gameId)
}

// GetDefenseTeam is a free data retrieval call binding the contract method 0xb8088875.
//
// Solidity: function getDefenseTeam(uint256 gameId) view returns(uint256 teamId, address teamOwner, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint16 battlePoint, uint16 timePoint, uint256 currentGameId, uint128 lockTo)
func (_Crabcaller *CrabcallerCallerSession) GetDefenseTeam(gameId *big.Int) (struct {
	TeamId        *big.Int
	TeamOwner     common.Address
	CrabadaId1    *big.Int
	CrabadaId2    *big.Int
	CrabadaId3    *big.Int
	BattlePoint   uint16
	TimePoint     uint16
	CurrentGameId *big.Int
	LockTo        *big.Int
}, error) {
	return _Crabcaller.Contract.GetDefenseTeam(&_Crabcaller.CallOpts, gameId)
}

// GetGameDefTeamInfo is a free data retrieval call binding the contract method 0x6657193a.
//
// Solidity: function getGameDefTeamInfo(uint256 gameId) view returns(uint32 startTime, uint16 battlePoint, uint16 timePoint)
func (_Crabcaller *CrabcallerCaller) GetGameDefTeamInfo(opts *bind.CallOpts, gameId *big.Int) (struct {
	StartTime   uint32
	BattlePoint uint16
	TimePoint   uint16
}, error) {
	var out []interface{}
	err := _Crabcaller.contract.Call(opts, &out, "getGameDefTeamInfo", gameId)

	outstruct := new(struct {
		StartTime   uint32
		BattlePoint uint16
		TimePoint   uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTime = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BattlePoint = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.TimePoint = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetGameDefTeamInfo is a free data retrieval call binding the contract method 0x6657193a.
//
// Solidity: function getGameDefTeamInfo(uint256 gameId) view returns(uint32 startTime, uint16 battlePoint, uint16 timePoint)
func (_Crabcaller *CrabcallerSession) GetGameDefTeamInfo(gameId *big.Int) (struct {
	StartTime   uint32
	BattlePoint uint16
	TimePoint   uint16
}, error) {
	return _Crabcaller.Contract.GetGameDefTeamInfo(&_Crabcaller.CallOpts, gameId)
}

// GetGameDefTeamInfo is a free data retrieval call binding the contract method 0x6657193a.
//
// Solidity: function getGameDefTeamInfo(uint256 gameId) view returns(uint32 startTime, uint16 battlePoint, uint16 timePoint)
func (_Crabcaller *CrabcallerCallerSession) GetGameDefTeamInfo(gameId *big.Int) (struct {
	StartTime   uint32
	BattlePoint uint16
	TimePoint   uint16
}, error) {
	return _Crabcaller.Contract.GetGameDefTeamInfo(&_Crabcaller.CallOpts, gameId)
}

// GetTeamInfos is a free data retrieval call binding the contract method 0xed61b15f.
//
// Solidity: function getTeamInfos(uint256[] teamIds) view returns(address[] teamOwners, uint16[] battlePoints, uint16[] timePoints, uint256[] currentGameIds, uint128[] lockTos)
func (_Crabcaller *CrabcallerCaller) GetTeamInfos(opts *bind.CallOpts, teamIds []*big.Int) (struct {
	TeamOwners     []common.Address
	BattlePoints   []uint16
	TimePoints     []uint16
	CurrentGameIds []*big.Int
	LockTos        []*big.Int
}, error) {
	var out []interface{}
	err := _Crabcaller.contract.Call(opts, &out, "getTeamInfos", teamIds)

	outstruct := new(struct {
		TeamOwners     []common.Address
		BattlePoints   []uint16
		TimePoints     []uint16
		CurrentGameIds []*big.Int
		LockTos        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TeamOwners = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.BattlePoints = *abi.ConvertType(out[1], new([]uint16)).(*[]uint16)
	outstruct.TimePoints = *abi.ConvertType(out[2], new([]uint16)).(*[]uint16)
	outstruct.CurrentGameIds = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)
	outstruct.LockTos = *abi.ConvertType(out[4], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTeamInfos is a free data retrieval call binding the contract method 0xed61b15f.
//
// Solidity: function getTeamInfos(uint256[] teamIds) view returns(address[] teamOwners, uint16[] battlePoints, uint16[] timePoints, uint256[] currentGameIds, uint128[] lockTos)
func (_Crabcaller *CrabcallerSession) GetTeamInfos(teamIds []*big.Int) (struct {
	TeamOwners     []common.Address
	BattlePoints   []uint16
	TimePoints     []uint16
	CurrentGameIds []*big.Int
	LockTos        []*big.Int
}, error) {
	return _Crabcaller.Contract.GetTeamInfos(&_Crabcaller.CallOpts, teamIds)
}

// GetTeamInfos is a free data retrieval call binding the contract method 0xed61b15f.
//
// Solidity: function getTeamInfos(uint256[] teamIds) view returns(address[] teamOwners, uint16[] battlePoints, uint16[] timePoints, uint256[] currentGameIds, uint128[] lockTos)
func (_Crabcaller *CrabcallerCallerSession) GetTeamInfos(teamIds []*big.Int) (struct {
	TeamOwners     []common.Address
	BattlePoints   []uint16
	TimePoints     []uint16
	CurrentGameIds []*big.Int
	LockTos        []*big.Int
}, error) {
	return _Crabcaller.Contract.GetTeamInfos(&_Crabcaller.CallOpts, teamIds)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crabcaller *CrabcallerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crabcaller.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crabcaller *CrabcallerSession) Owner() (common.Address, error) {
	return _Crabcaller.Contract.Owner(&_Crabcaller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crabcaller *CrabcallerCallerSession) Owner() (common.Address, error) {
	return _Crabcaller.Contract.Owner(&_Crabcaller.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crabcaller *CrabcallerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crabcaller.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crabcaller *CrabcallerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crabcaller.Contract.RenounceOwnership(&_Crabcaller.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crabcaller *CrabcallerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crabcaller.Contract.RenounceOwnership(&_Crabcaller.TransactOpts)
}

// SetCrabadaAddress is a paid mutator transaction binding the contract method 0x971a7709.
//
// Solidity: function setCrabadaAddress(address crabadaAddress) returns()
func (_Crabcaller *CrabcallerTransactor) SetCrabadaAddress(opts *bind.TransactOpts, crabadaAddress common.Address) (*types.Transaction, error) {
	return _Crabcaller.contract.Transact(opts, "setCrabadaAddress", crabadaAddress)
}

// SetCrabadaAddress is a paid mutator transaction binding the contract method 0x971a7709.
//
// Solidity: function setCrabadaAddress(address crabadaAddress) returns()
func (_Crabcaller *CrabcallerSession) SetCrabadaAddress(crabadaAddress common.Address) (*types.Transaction, error) {
	return _Crabcaller.Contract.SetCrabadaAddress(&_Crabcaller.TransactOpts, crabadaAddress)
}

// SetCrabadaAddress is a paid mutator transaction binding the contract method 0x971a7709.
//
// Solidity: function setCrabadaAddress(address crabadaAddress) returns()
func (_Crabcaller *CrabcallerTransactorSession) SetCrabadaAddress(crabadaAddress common.Address) (*types.Transaction, error) {
	return _Crabcaller.Contract.SetCrabadaAddress(&_Crabcaller.TransactOpts, crabadaAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crabcaller *CrabcallerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Crabcaller.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crabcaller *CrabcallerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crabcaller.Contract.TransferOwnership(&_Crabcaller.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crabcaller *CrabcallerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crabcaller.Contract.TransferOwnership(&_Crabcaller.TransactOpts, newOwner)
}

// CrabcallerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Crabcaller contract.
type CrabcallerOwnershipTransferredIterator struct {
	Event *CrabcallerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CrabcallerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabcallerOwnershipTransferred)
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
		it.Event = new(CrabcallerOwnershipTransferred)
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
func (it *CrabcallerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabcallerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabcallerOwnershipTransferred represents a OwnershipTransferred event raised by the Crabcaller contract.
type CrabcallerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crabcaller *CrabcallerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrabcallerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crabcaller.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrabcallerOwnershipTransferredIterator{contract: _Crabcaller.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crabcaller *CrabcallerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrabcallerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crabcaller.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabcallerOwnershipTransferred)
				if err := _Crabcaller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Crabcaller *CrabcallerFilterer) ParseOwnershipTransferred(log types.Log) (*CrabcallerOwnershipTransferred, error) {
	event := new(CrabcallerOwnershipTransferred)
	if err := _Crabcaller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

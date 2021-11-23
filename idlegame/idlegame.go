// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package idlegame

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

// IdlegameMetaData contains all meta data concerning the Idlegame contract.
var IdlegameMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"battlePoint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timePoint\",\"type\":\"uint256\"}],\"name\":\"AddCrabada\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"CloseGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId3\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"battlePoint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timePoint\",\"type\":\"uint256\"}],\"name\":\"CreateTeam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"turn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"attackTeamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defenseTeamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"soldierId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"attackTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"attackPoint\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"defensePoint\",\"type\":\"uint16\"}],\"name\":\"Fight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Lend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"battlePoint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timePoint\",\"type\":\"uint256\"}],\"name\":\"RemoveCrabada\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SetLendingPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerTeamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerCraReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerTusReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loserCraReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loserTusReward\",\"type\":\"uint256\"}],\"name\":\"SettleGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"craReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tusReward\",\"type\":\"uint256\"}],\"name\":\"StartGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CRAB_STATUS_IN_TEAM\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CRAB_STATUS_LENDING\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAME_STATUS_CLAIMED\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAME_STATUS_SETTLE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STEP_DURATION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attackCooldownDuration\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseCraReward\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTusReward\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusCraClass\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusTusClass\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"craIncentiveEnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"craToken\",\"outputs\":[{\"internalType\":\"contractIToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"crabaInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"lockTo\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"status\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crabada\",\"outputs\":[{\"internalType\":\"contractICrabada\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameStatContract\",\"outputs\":[{\"internalType\":\"contractIGameStat\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendingFeeHolerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendingFeePercent\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendingLockDuration\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lootingCraReward\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lootingPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lootingTusReward\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miningDuration\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinforceLockDuration\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tusToken\",\"outputs\":[{\"internalType\":\"contractIToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrabada\",\"name\":\"_crabada\",\"type\":\"address\"},{\"internalType\":\"contractIToken\",\"name\":\"_cra\",\"type\":\"address\"},{\"internalType\":\"contractIToken\",\"name\":\"_tus\",\"type\":\"address\"},{\"internalType\":\"contractIGameStat\",\"name\":\"_gameStat\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"crabadaIds\",\"type\":\"uint256[]\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"crabadaIds\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setLendingPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crabadaId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId3\",\"type\":\"uint256\"}],\"name\":\"createTeam\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"}],\"name\":\"addCrabadaToTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"removeCrabadaFromTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"}],\"name\":\"startGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"attackTeamId\",\"type\":\"uint256\"}],\"name\":\"attack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowPrice\",\"type\":\"uint256\"}],\"name\":\"reinforceAttack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowPrice\",\"type\":\"uint256\"}],\"name\":\"reinforceDefense\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"settleGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"closeGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"}],\"name\":\"getStats\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"battlePoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"timePoint\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"}],\"name\":\"getTeamInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crabadaId3\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"battlePoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"timePoint\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"currentGameId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"lockTo\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getGameBasicInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"teamId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"craReward\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tusReward\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"status\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getGameBattleInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"attackTeamId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"attackTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastAttackTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastDefTime\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"attackId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"attackId2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defId2\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getLootingStatsInfo\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"attackPoint\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defensePoint\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crabadaId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"craRewardAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tusRewardAmount\",\"type\":\"uint128\"}],\"name\":\"setRewardMining\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"craRewardAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tusRewardAmount\",\"type\":\"uint128\"}],\"name\":\"setRewardLooting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"setLootingPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"duration\",\"type\":\"uint128\"}],\"name\":\"setStepDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"duration\",\"type\":\"uint128\"}],\"name\":\"setMiningDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"duration\",\"type\":\"uint128\"}],\"name\":\"setReinforceLockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"duration\",\"type\":\"uint128\"}],\"name\":\"setAttackCooldownDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"setLendingFeeHolerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isEnable\",\"type\":\"bool\"}],\"name\":\"setCRAIncentiveStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IdlegameABI is the input ABI used to generate the binding from.
// Deprecated: Use IdlegameMetaData.ABI instead.
var IdlegameABI = IdlegameMetaData.ABI

// Idlegame is an auto generated Go binding around an Ethereum contract.
type Idlegame struct {
	IdlegameCaller     // Read-only binding to the contract
	IdlegameTransactor // Write-only binding to the contract
	IdlegameFilterer   // Log filterer for contract events
}

// IdlegameCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdlegameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdlegameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdlegameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdlegameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdlegameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdlegameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdlegameSession struct {
	Contract     *Idlegame         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdlegameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdlegameCallerSession struct {
	Contract *IdlegameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IdlegameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdlegameTransactorSession struct {
	Contract     *IdlegameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IdlegameRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdlegameRaw struct {
	Contract *Idlegame // Generic contract binding to access the raw methods on
}

// IdlegameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdlegameCallerRaw struct {
	Contract *IdlegameCaller // Generic read-only contract binding to access the raw methods on
}

// IdlegameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdlegameTransactorRaw struct {
	Contract *IdlegameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdlegame creates a new instance of Idlegame, bound to a specific deployed contract.
func NewIdlegame(address common.Address, backend bind.ContractBackend) (*Idlegame, error) {
	contract, err := bindIdlegame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Idlegame{IdlegameCaller: IdlegameCaller{contract: contract}, IdlegameTransactor: IdlegameTransactor{contract: contract}, IdlegameFilterer: IdlegameFilterer{contract: contract}}, nil
}

// NewIdlegameCaller creates a new read-only instance of Idlegame, bound to a specific deployed contract.
func NewIdlegameCaller(address common.Address, caller bind.ContractCaller) (*IdlegameCaller, error) {
	contract, err := bindIdlegame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdlegameCaller{contract: contract}, nil
}

// NewIdlegameTransactor creates a new write-only instance of Idlegame, bound to a specific deployed contract.
func NewIdlegameTransactor(address common.Address, transactor bind.ContractTransactor) (*IdlegameTransactor, error) {
	contract, err := bindIdlegame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdlegameTransactor{contract: contract}, nil
}

// NewIdlegameFilterer creates a new log filterer instance of Idlegame, bound to a specific deployed contract.
func NewIdlegameFilterer(address common.Address, filterer bind.ContractFilterer) (*IdlegameFilterer, error) {
	contract, err := bindIdlegame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdlegameFilterer{contract: contract}, nil
}

// bindIdlegame binds a generic wrapper to an already deployed contract.
func bindIdlegame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdlegameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Idlegame *IdlegameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Idlegame.Contract.IdlegameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Idlegame *IdlegameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Idlegame.Contract.IdlegameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Idlegame *IdlegameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Idlegame.Contract.IdlegameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Idlegame *IdlegameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Idlegame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Idlegame *IdlegameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Idlegame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Idlegame *IdlegameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Idlegame.Contract.contract.Transact(opts, method, params...)
}

// CRABSTATUSINTEAM is a free data retrieval call binding the contract method 0xd937194e.
//
// Solidity: function CRAB_STATUS_IN_TEAM() view returns(uint128)
func (_Idlegame *IdlegameCaller) CRABSTATUSINTEAM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "CRAB_STATUS_IN_TEAM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CRABSTATUSINTEAM is a free data retrieval call binding the contract method 0xd937194e.
//
// Solidity: function CRAB_STATUS_IN_TEAM() view returns(uint128)
func (_Idlegame *IdlegameSession) CRABSTATUSINTEAM() (*big.Int, error) {
	return _Idlegame.Contract.CRABSTATUSINTEAM(&_Idlegame.CallOpts)
}

// CRABSTATUSINTEAM is a free data retrieval call binding the contract method 0xd937194e.
//
// Solidity: function CRAB_STATUS_IN_TEAM() view returns(uint128)
func (_Idlegame *IdlegameCallerSession) CRABSTATUSINTEAM() (*big.Int, error) {
	return _Idlegame.Contract.CRABSTATUSINTEAM(&_Idlegame.CallOpts)
}

// CRABSTATUSLENDING is a free data retrieval call binding the contract method 0xce8f1fbe.
//
// Solidity: function CRAB_STATUS_LENDING() view returns(uint128)
func (_Idlegame *IdlegameCaller) CRABSTATUSLENDING(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "CRAB_STATUS_LENDING")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CRABSTATUSLENDING is a free data retrieval call binding the contract method 0xce8f1fbe.
//
// Solidity: function CRAB_STATUS_LENDING() view returns(uint128)
func (_Idlegame *IdlegameSession) CRABSTATUSLENDING() (*big.Int, error) {
	return _Idlegame.Contract.CRABSTATUSLENDING(&_Idlegame.CallOpts)
}

// CRABSTATUSLENDING is a free data retrieval call binding the contract method 0xce8f1fbe.
//
// Solidity: function CRAB_STATUS_LENDING() view returns(uint128)
func (_Idlegame *IdlegameCallerSession) CRABSTATUSLENDING() (*big.Int, error) {
	return _Idlegame.Contract.CRABSTATUSLENDING(&_Idlegame.CallOpts)
}

// GAMESTATUSCLAIMED is a free data retrieval call binding the contract method 0x3f72ce85.
//
// Solidity: function GAME_STATUS_CLAIMED() view returns(uint32)
func (_Idlegame *IdlegameCaller) GAMESTATUSCLAIMED(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "GAME_STATUS_CLAIMED")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GAMESTATUSCLAIMED is a free data retrieval call binding the contract method 0x3f72ce85.
//
// Solidity: function GAME_STATUS_CLAIMED() view returns(uint32)
func (_Idlegame *IdlegameSession) GAMESTATUSCLAIMED() (uint32, error) {
	return _Idlegame.Contract.GAMESTATUSCLAIMED(&_Idlegame.CallOpts)
}

// GAMESTATUSCLAIMED is a free data retrieval call binding the contract method 0x3f72ce85.
//
// Solidity: function GAME_STATUS_CLAIMED() view returns(uint32)
func (_Idlegame *IdlegameCallerSession) GAMESTATUSCLAIMED() (uint32, error) {
	return _Idlegame.Contract.GAMESTATUSCLAIMED(&_Idlegame.CallOpts)
}

// GAMESTATUSSETTLE is a free data retrieval call binding the contract method 0x3c08b0ae.
//
// Solidity: function GAME_STATUS_SETTLE() view returns(uint32)
func (_Idlegame *IdlegameCaller) GAMESTATUSSETTLE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "GAME_STATUS_SETTLE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GAMESTATUSSETTLE is a free data retrieval call binding the contract method 0x3c08b0ae.
//
// Solidity: function GAME_STATUS_SETTLE() view returns(uint32)
func (_Idlegame *IdlegameSession) GAMESTATUSSETTLE() (uint32, error) {
	return _Idlegame.Contract.GAMESTATUSSETTLE(&_Idlegame.CallOpts)
}

// GAMESTATUSSETTLE is a free data retrieval call binding the contract method 0x3c08b0ae.
//
// Solidity: function GAME_STATUS_SETTLE() view returns(uint32)
func (_Idlegame *IdlegameCallerSession) GAMESTATUSSETTLE() (uint32, error) {
	return _Idlegame.Contract.GAMESTATUSSETTLE(&_Idlegame.CallOpts)
}

// STEPDURATION is a free data retrieval call binding the contract method 0xf67718cb.
//
// Solidity: function STEP_DURATION() view returns(uint128)
func (_Idlegame *IdlegameCaller) STEPDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "STEP_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STEPDURATION is a free data retrieval call binding the contract method 0xf67718cb.
//
// Solidity: function STEP_DURATION() view returns(uint128)
func (_Idlegame *IdlegameSession) STEPDURATION() (*big.Int, error) {
	return _Idlegame.Contract.STEPDURATION(&_Idlegame.CallOpts)
}

// STEPDURATION is a free data retrieval call binding the contract method 0xf67718cb.
//
// Solidity: function STEP_DURATION() view returns(uint128)
func (_Idlegame *IdlegameCallerSession) STEPDURATION() (*big.Int, error) {
	return _Idlegame.Contract.STEPDURATION(&_Idlegame.CallOpts)
}

// AttackCooldownDuration is a free data retrieval call binding the contract method 0x6c1fcc04.
//
// Solidity: function attackCooldownDuration() view returns(uint128)
func (_Idlegame *IdlegameCaller) AttackCooldownDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "attackCooldownDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttackCooldownDuration is a free data retrieval call binding the contract method 0x6c1fcc04.
//
// Solidity: function attackCooldownDuration() view returns(uint128)
func (_Idlegame *IdlegameSession) AttackCooldownDuration() (*big.Int, error) {
	return _Idlegame.Contract.AttackCooldownDuration(&_Idlegame.CallOpts)
}

// AttackCooldownDuration is a free data retrieval call binding the contract method 0x6c1fcc04.
//
// Solidity: function attackCooldownDuration() view returns(uint128)
func (_Idlegame *IdlegameCallerSession) AttackCooldownDuration() (*big.Int, error) {
	return _Idlegame.Contract.AttackCooldownDuration(&_Idlegame.CallOpts)
}

// BaseCraReward is a free data retrieval call binding the contract method 0x7d2694e4.
//
// Solidity: function baseCraReward() view returns(uint128)
func (_Idlegame *IdlegameCaller) BaseCraReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "baseCraReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseCraReward is a free data retrieval call binding the contract method 0x7d2694e4.
//
// Solidity: function baseCraReward() view returns(uint128)
func (_Idlegame *IdlegameSession) BaseCraReward() (*big.Int, error) {
	return _Idlegame.Contract.BaseCraReward(&_Idlegame.CallOpts)
}

// BaseCraReward is a free data retrieval call binding the contract method 0x7d2694e4.
//
// Solidity: function baseCraReward() view returns(uint128)
func (_Idlegame *IdlegameCallerSession) BaseCraReward() (*big.Int, error) {
	return _Idlegame.Contract.BaseCraReward(&_Idlegame.CallOpts)
}

// BaseTusReward is a free data retrieval call binding the contract method 0x580c484e.
//
// Solidity: function baseTusReward() view returns(uint128)
func (_Idlegame *IdlegameCaller) BaseTusReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "baseTusReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseTusReward is a free data retrieval call binding the contract method 0x580c484e.
//
// Solidity: function baseTusReward() view returns(uint128)
func (_Idlegame *IdlegameSession) BaseTusReward() (*big.Int, error) {
	return _Idlegame.Contract.BaseTusReward(&_Idlegame.CallOpts)
}

// BaseTusReward is a free data retrieval call binding the contract method 0x580c484e.
//
// Solidity: function baseTusReward() view returns(uint128)
func (_Idlegame *IdlegameCallerSession) BaseTusReward() (*big.Int, error) {
	return _Idlegame.Contract.BaseTusReward(&_Idlegame.CallOpts)
}

// BonusCraClass is a free data retrieval call binding the contract method 0xd0f0cd1f.
//
// Solidity: function bonusCraClass() view returns(uint8)
func (_Idlegame *IdlegameCaller) BonusCraClass(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "bonusCraClass")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BonusCraClass is a free data retrieval call binding the contract method 0xd0f0cd1f.
//
// Solidity: function bonusCraClass() view returns(uint8)
func (_Idlegame *IdlegameSession) BonusCraClass() (uint8, error) {
	return _Idlegame.Contract.BonusCraClass(&_Idlegame.CallOpts)
}

// BonusCraClass is a free data retrieval call binding the contract method 0xd0f0cd1f.
//
// Solidity: function bonusCraClass() view returns(uint8)
func (_Idlegame *IdlegameCallerSession) BonusCraClass() (uint8, error) {
	return _Idlegame.Contract.BonusCraClass(&_Idlegame.CallOpts)
}

// BonusTusClass is a free data retrieval call binding the contract method 0x4baa6a6b.
//
// Solidity: function bonusTusClass() view returns(uint8)
func (_Idlegame *IdlegameCaller) BonusTusClass(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "bonusTusClass")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BonusTusClass is a free data retrieval call binding the contract method 0x4baa6a6b.
//
// Solidity: function bonusTusClass() view returns(uint8)
func (_Idlegame *IdlegameSession) BonusTusClass() (uint8, error) {
	return _Idlegame.Contract.BonusTusClass(&_Idlegame.CallOpts)
}

// BonusTusClass is a free data retrieval call binding the contract method 0x4baa6a6b.
//
// Solidity: function bonusTusClass() view returns(uint8)
func (_Idlegame *IdlegameCallerSession) BonusTusClass() (uint8, error) {
	return _Idlegame.Contract.BonusTusClass(&_Idlegame.CallOpts)
}

// CraIncentiveEnable is a free data retrieval call binding the contract method 0x1ef17d36.
//
// Solidity: function craIncentiveEnable() view returns(bool)
func (_Idlegame *IdlegameCaller) CraIncentiveEnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "craIncentiveEnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CraIncentiveEnable is a free data retrieval call binding the contract method 0x1ef17d36.
//
// Solidity: function craIncentiveEnable() view returns(bool)
func (_Idlegame *IdlegameSession) CraIncentiveEnable() (bool, error) {
	return _Idlegame.Contract.CraIncentiveEnable(&_Idlegame.CallOpts)
}

// CraIncentiveEnable is a free data retrieval call binding the contract method 0x1ef17d36.
//
// Solidity: function craIncentiveEnable() view returns(bool)
func (_Idlegame *IdlegameCallerSession) CraIncentiveEnable() (bool, error) {
	return _Idlegame.Contract.CraIncentiveEnable(&_Idlegame.CallOpts)
}

// CraToken is a free data retrieval call binding the contract method 0xc9835144.
//
// Solidity: function craToken() view returns(address)
func (_Idlegame *IdlegameCaller) CraToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "craToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CraToken is a free data retrieval call binding the contract method 0xc9835144.
//
// Solidity: function craToken() view returns(address)
func (_Idlegame *IdlegameSession) CraToken() (common.Address, error) {
	return _Idlegame.Contract.CraToken(&_Idlegame.CallOpts)
}

// CraToken is a free data retrieval call binding the contract method 0xc9835144.
//
// Solidity: function craToken() view returns(address)
func (_Idlegame *IdlegameCallerSession) CraToken() (common.Address, error) {
	return _Idlegame.Contract.CraToken(&_Idlegame.CallOpts)
}

// CrabaInfos is a free data retrieval call binding the contract method 0x62428e4e.
//
// Solidity: function crabaInfos(uint256 ) view returns(address owner, uint128 lockTo, uint128 status)
func (_Idlegame *IdlegameCaller) CrabaInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner  common.Address
	LockTo *big.Int
	Status *big.Int
}, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "crabaInfos", arg0)

	outstruct := new(struct {
		Owner  common.Address
		LockTo *big.Int
		Status *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LockTo = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CrabaInfos is a free data retrieval call binding the contract method 0x62428e4e.
//
// Solidity: function crabaInfos(uint256 ) view returns(address owner, uint128 lockTo, uint128 status)
func (_Idlegame *IdlegameSession) CrabaInfos(arg0 *big.Int) (struct {
	Owner  common.Address
	LockTo *big.Int
	Status *big.Int
}, error) {
	return _Idlegame.Contract.CrabaInfos(&_Idlegame.CallOpts, arg0)
}

// CrabaInfos is a free data retrieval call binding the contract method 0x62428e4e.
//
// Solidity: function crabaInfos(uint256 ) view returns(address owner, uint128 lockTo, uint128 status)
func (_Idlegame *IdlegameCallerSession) CrabaInfos(arg0 *big.Int) (struct {
	Owner  common.Address
	LockTo *big.Int
	Status *big.Int
}, error) {
	return _Idlegame.Contract.CrabaInfos(&_Idlegame.CallOpts, arg0)
}

// Crabada is a free data retrieval call binding the contract method 0x9e19916d.
//
// Solidity: function crabada() view returns(address)
func (_Idlegame *IdlegameCaller) Crabada(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "crabada")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Crabada is a free data retrieval call binding the contract method 0x9e19916d.
//
// Solidity: function crabada() view returns(address)
func (_Idlegame *IdlegameSession) Crabada() (common.Address, error) {
	return _Idlegame.Contract.Crabada(&_Idlegame.CallOpts)
}

// Crabada is a free data retrieval call binding the contract method 0x9e19916d.
//
// Solidity: function crabada() view returns(address)
func (_Idlegame *IdlegameCallerSession) Crabada() (common.Address, error) {
	return _Idlegame.Contract.Crabada(&_Idlegame.CallOpts)
}

// GameStatContract is a free data retrieval call binding the contract method 0x27a4ff44.
//
// Solidity: function gameStatContract() view returns(address)
func (_Idlegame *IdlegameCaller) GameStatContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "gameStatContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GameStatContract is a free data retrieval call binding the contract method 0x27a4ff44.
//
// Solidity: function gameStatContract() view returns(address)
func (_Idlegame *IdlegameSession) GameStatContract() (common.Address, error) {
	return _Idlegame.Contract.GameStatContract(&_Idlegame.CallOpts)
}

// GameStatContract is a free data retrieval call binding the contract method 0x27a4ff44.
//
// Solidity: function gameStatContract() view returns(address)
func (_Idlegame *IdlegameCallerSession) GameStatContract() (common.Address, error) {
	return _Idlegame.Contract.GameStatContract(&_Idlegame.CallOpts)
}

// GetGameBasicInfo is a free data retrieval call binding the contract method 0xf0344e36.
//
// Solidity: function getGameBasicInfo(uint256 gameId) view returns(uint256 teamId, uint128 craReward, uint128 tusReward, uint32 startTime, uint32 duration, uint32 status)
func (_Idlegame *IdlegameCaller) GetGameBasicInfo(opts *bind.CallOpts, gameId *big.Int) (struct {
	TeamId    *big.Int
	CraReward *big.Int
	TusReward *big.Int
	StartTime uint32
	Duration  uint32
	Status    uint32
}, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "getGameBasicInfo", gameId)

	outstruct := new(struct {
		TeamId    *big.Int
		CraReward *big.Int
		TusReward *big.Int
		StartTime uint32
		Duration  uint32
		Status    uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TeamId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CraReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TusReward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Duration = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.Status = *abi.ConvertType(out[5], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetGameBasicInfo is a free data retrieval call binding the contract method 0xf0344e36.
//
// Solidity: function getGameBasicInfo(uint256 gameId) view returns(uint256 teamId, uint128 craReward, uint128 tusReward, uint32 startTime, uint32 duration, uint32 status)
func (_Idlegame *IdlegameSession) GetGameBasicInfo(gameId *big.Int) (struct {
	TeamId    *big.Int
	CraReward *big.Int
	TusReward *big.Int
	StartTime uint32
	Duration  uint32
	Status    uint32
}, error) {
	return _Idlegame.Contract.GetGameBasicInfo(&_Idlegame.CallOpts, gameId)
}

// GetGameBasicInfo is a free data retrieval call binding the contract method 0xf0344e36.
//
// Solidity: function getGameBasicInfo(uint256 gameId) view returns(uint256 teamId, uint128 craReward, uint128 tusReward, uint32 startTime, uint32 duration, uint32 status)
func (_Idlegame *IdlegameCallerSession) GetGameBasicInfo(gameId *big.Int) (struct {
	TeamId    *big.Int
	CraReward *big.Int
	TusReward *big.Int
	StartTime uint32
	Duration  uint32
	Status    uint32
}, error) {
	return _Idlegame.Contract.GetGameBasicInfo(&_Idlegame.CallOpts, gameId)
}

// GetGameBattleInfo is a free data retrieval call binding the contract method 0x183ce75d.
//
// Solidity: function getGameBattleInfo(uint256 gameId) view returns(uint256 attackTeamId, uint32 attackTime, uint32 lastAttackTime, uint32 lastDefTime, uint256 attackId1, uint256 attackId2, uint256 defId1, uint256 defId2)
func (_Idlegame *IdlegameCaller) GetGameBattleInfo(opts *bind.CallOpts, gameId *big.Int) (struct {
	AttackTeamId   *big.Int
	AttackTime     uint32
	LastAttackTime uint32
	LastDefTime    uint32
	AttackId1      *big.Int
	AttackId2      *big.Int
	DefId1         *big.Int
	DefId2         *big.Int
}, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "getGameBattleInfo", gameId)

	outstruct := new(struct {
		AttackTeamId   *big.Int
		AttackTime     uint32
		LastAttackTime uint32
		LastDefTime    uint32
		AttackId1      *big.Int
		AttackId2      *big.Int
		DefId1         *big.Int
		DefId2         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttackTeamId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AttackTime = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.LastAttackTime = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.LastDefTime = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.AttackId1 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AttackId2 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DefId1 = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DefId2 = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGameBattleInfo is a free data retrieval call binding the contract method 0x183ce75d.
//
// Solidity: function getGameBattleInfo(uint256 gameId) view returns(uint256 attackTeamId, uint32 attackTime, uint32 lastAttackTime, uint32 lastDefTime, uint256 attackId1, uint256 attackId2, uint256 defId1, uint256 defId2)
func (_Idlegame *IdlegameSession) GetGameBattleInfo(gameId *big.Int) (struct {
	AttackTeamId   *big.Int
	AttackTime     uint32
	LastAttackTime uint32
	LastDefTime    uint32
	AttackId1      *big.Int
	AttackId2      *big.Int
	DefId1         *big.Int
	DefId2         *big.Int
}, error) {
	return _Idlegame.Contract.GetGameBattleInfo(&_Idlegame.CallOpts, gameId)
}

// GetGameBattleInfo is a free data retrieval call binding the contract method 0x183ce75d.
//
// Solidity: function getGameBattleInfo(uint256 gameId) view returns(uint256 attackTeamId, uint32 attackTime, uint32 lastAttackTime, uint32 lastDefTime, uint256 attackId1, uint256 attackId2, uint256 defId1, uint256 defId2)
func (_Idlegame *IdlegameCallerSession) GetGameBattleInfo(gameId *big.Int) (struct {
	AttackTeamId   *big.Int
	AttackTime     uint32
	LastAttackTime uint32
	LastDefTime    uint32
	AttackId1      *big.Int
	AttackId2      *big.Int
	DefId1         *big.Int
	DefId2         *big.Int
}, error) {
	return _Idlegame.Contract.GetGameBattleInfo(&_Idlegame.CallOpts, gameId)
}

// GetLootingStatsInfo is a free data retrieval call binding the contract method 0xab0c8f8d.
//
// Solidity: function getLootingStatsInfo(uint256 gameId) view returns(uint16 attackPoint, uint16 defensePoint)
func (_Idlegame *IdlegameCaller) GetLootingStatsInfo(opts *bind.CallOpts, gameId *big.Int) (struct {
	AttackPoint  uint16
	DefensePoint uint16
}, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "getLootingStatsInfo", gameId)

	outstruct := new(struct {
		AttackPoint  uint16
		DefensePoint uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttackPoint = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.DefensePoint = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetLootingStatsInfo is a free data retrieval call binding the contract method 0xab0c8f8d.
//
// Solidity: function getLootingStatsInfo(uint256 gameId) view returns(uint16 attackPoint, uint16 defensePoint)
func (_Idlegame *IdlegameSession) GetLootingStatsInfo(gameId *big.Int) (struct {
	AttackPoint  uint16
	DefensePoint uint16
}, error) {
	return _Idlegame.Contract.GetLootingStatsInfo(&_Idlegame.CallOpts, gameId)
}

// GetLootingStatsInfo is a free data retrieval call binding the contract method 0xab0c8f8d.
//
// Solidity: function getLootingStatsInfo(uint256 gameId) view returns(uint16 attackPoint, uint16 defensePoint)
func (_Idlegame *IdlegameCallerSession) GetLootingStatsInfo(gameId *big.Int) (struct {
	AttackPoint  uint16
	DefensePoint uint16
}, error) {
	return _Idlegame.Contract.GetLootingStatsInfo(&_Idlegame.CallOpts, gameId)
}

// GetStats is a free data retrieval call binding the contract method 0x7b303965.
//
// Solidity: function getStats(uint256 crabadaId) view returns(uint16 battlePoint, uint16 timePoint)
func (_Idlegame *IdlegameCaller) GetStats(opts *bind.CallOpts, crabadaId *big.Int) (struct {
	BattlePoint uint16
	TimePoint   uint16
}, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "getStats", crabadaId)

	outstruct := new(struct {
		BattlePoint uint16
		TimePoint   uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BattlePoint = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.TimePoint = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetStats is a free data retrieval call binding the contract method 0x7b303965.
//
// Solidity: function getStats(uint256 crabadaId) view returns(uint16 battlePoint, uint16 timePoint)
func (_Idlegame *IdlegameSession) GetStats(crabadaId *big.Int) (struct {
	BattlePoint uint16
	TimePoint   uint16
}, error) {
	return _Idlegame.Contract.GetStats(&_Idlegame.CallOpts, crabadaId)
}

// GetStats is a free data retrieval call binding the contract method 0x7b303965.
//
// Solidity: function getStats(uint256 crabadaId) view returns(uint16 battlePoint, uint16 timePoint)
func (_Idlegame *IdlegameCallerSession) GetStats(crabadaId *big.Int) (struct {
	BattlePoint uint16
	TimePoint   uint16
}, error) {
	return _Idlegame.Contract.GetStats(&_Idlegame.CallOpts, crabadaId)
}

// GetTeamInfo is a free data retrieval call binding the contract method 0x969215ba.
//
// Solidity: function getTeamInfo(uint256 teamId) view returns(address owner, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint16 battlePoint, uint16 timePoint, uint256 currentGameId, uint128 lockTo)
func (_Idlegame *IdlegameCaller) GetTeamInfo(opts *bind.CallOpts, teamId *big.Int) (struct {
	Owner         common.Address
	CrabadaId1    *big.Int
	CrabadaId2    *big.Int
	CrabadaId3    *big.Int
	BattlePoint   uint16
	TimePoint     uint16
	CurrentGameId *big.Int
	LockTo        *big.Int
}, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "getTeamInfo", teamId)

	outstruct := new(struct {
		Owner         common.Address
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

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CrabadaId1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CrabadaId2 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CrabadaId3 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BattlePoint = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.TimePoint = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.CurrentGameId = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LockTo = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTeamInfo is a free data retrieval call binding the contract method 0x969215ba.
//
// Solidity: function getTeamInfo(uint256 teamId) view returns(address owner, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint16 battlePoint, uint16 timePoint, uint256 currentGameId, uint128 lockTo)
func (_Idlegame *IdlegameSession) GetTeamInfo(teamId *big.Int) (struct {
	Owner         common.Address
	CrabadaId1    *big.Int
	CrabadaId2    *big.Int
	CrabadaId3    *big.Int
	BattlePoint   uint16
	TimePoint     uint16
	CurrentGameId *big.Int
	LockTo        *big.Int
}, error) {
	return _Idlegame.Contract.GetTeamInfo(&_Idlegame.CallOpts, teamId)
}

// GetTeamInfo is a free data retrieval call binding the contract method 0x969215ba.
//
// Solidity: function getTeamInfo(uint256 teamId) view returns(address owner, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint16 battlePoint, uint16 timePoint, uint256 currentGameId, uint128 lockTo)
func (_Idlegame *IdlegameCallerSession) GetTeamInfo(teamId *big.Int) (struct {
	Owner         common.Address
	CrabadaId1    *big.Int
	CrabadaId2    *big.Int
	CrabadaId3    *big.Int
	BattlePoint   uint16
	TimePoint     uint16
	CurrentGameId *big.Int
	LockTo        *big.Int
}, error) {
	return _Idlegame.Contract.GetTeamInfo(&_Idlegame.CallOpts, teamId)
}

// LendingFeeHolerAddress is a free data retrieval call binding the contract method 0x6dad6b4b.
//
// Solidity: function lendingFeeHolerAddress() view returns(address)
func (_Idlegame *IdlegameCaller) LendingFeeHolerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "lendingFeeHolerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LendingFeeHolerAddress is a free data retrieval call binding the contract method 0x6dad6b4b.
//
// Solidity: function lendingFeeHolerAddress() view returns(address)
func (_Idlegame *IdlegameSession) LendingFeeHolerAddress() (common.Address, error) {
	return _Idlegame.Contract.LendingFeeHolerAddress(&_Idlegame.CallOpts)
}

// LendingFeeHolerAddress is a free data retrieval call binding the contract method 0x6dad6b4b.
//
// Solidity: function lendingFeeHolerAddress() view returns(address)
func (_Idlegame *IdlegameCallerSession) LendingFeeHolerAddress() (common.Address, error) {
	return _Idlegame.Contract.LendingFeeHolerAddress(&_Idlegame.CallOpts)
}

// LendingFeePercent is a free data retrieval call binding the contract method 0x4699f846.
//
// Solidity: function lendingFeePercent() view returns(uint32)
func (_Idlegame *IdlegameCaller) LendingFeePercent(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "lendingFeePercent")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LendingFeePercent is a free data retrieval call binding the contract method 0x4699f846.
//
// Solidity: function lendingFeePercent() view returns(uint32)
func (_Idlegame *IdlegameSession) LendingFeePercent() (uint32, error) {
	return _Idlegame.Contract.LendingFeePercent(&_Idlegame.CallOpts)
}

// LendingFeePercent is a free data retrieval call binding the contract method 0x4699f846.
//
// Solidity: function lendingFeePercent() view returns(uint32)
func (_Idlegame *IdlegameCallerSession) LendingFeePercent() (uint32, error) {
	return _Idlegame.Contract.LendingFeePercent(&_Idlegame.CallOpts)
}

// LendingLockDuration is a free data retrieval call binding the contract method 0x24a9c445.
//
// Solidity: function lendingLockDuration() view returns(uint32)
func (_Idlegame *IdlegameCaller) LendingLockDuration(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "lendingLockDuration")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LendingLockDuration is a free data retrieval call binding the contract method 0x24a9c445.
//
// Solidity: function lendingLockDuration() view returns(uint32)
func (_Idlegame *IdlegameSession) LendingLockDuration() (uint32, error) {
	return _Idlegame.Contract.LendingLockDuration(&_Idlegame.CallOpts)
}

// LendingLockDuration is a free data retrieval call binding the contract method 0x24a9c445.
//
// Solidity: function lendingLockDuration() view returns(uint32)
func (_Idlegame *IdlegameCallerSession) LendingLockDuration() (uint32, error) {
	return _Idlegame.Contract.LendingLockDuration(&_Idlegame.CallOpts)
}

// LootingCraReward is a free data retrieval call binding the contract method 0xc68616b1.
//
// Solidity: function lootingCraReward() view returns(uint128)
func (_Idlegame *IdlegameCaller) LootingCraReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "lootingCraReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LootingCraReward is a free data retrieval call binding the contract method 0xc68616b1.
//
// Solidity: function lootingCraReward() view returns(uint128)
func (_Idlegame *IdlegameSession) LootingCraReward() (*big.Int, error) {
	return _Idlegame.Contract.LootingCraReward(&_Idlegame.CallOpts)
}

// LootingCraReward is a free data retrieval call binding the contract method 0xc68616b1.
//
// Solidity: function lootingCraReward() view returns(uint128)
func (_Idlegame *IdlegameCallerSession) LootingCraReward() (*big.Int, error) {
	return _Idlegame.Contract.LootingCraReward(&_Idlegame.CallOpts)
}

// LootingPercent is a free data retrieval call binding the contract method 0xfec8d7d1.
//
// Solidity: function lootingPercent() view returns(uint256)
func (_Idlegame *IdlegameCaller) LootingPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "lootingPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LootingPercent is a free data retrieval call binding the contract method 0xfec8d7d1.
//
// Solidity: function lootingPercent() view returns(uint256)
func (_Idlegame *IdlegameSession) LootingPercent() (*big.Int, error) {
	return _Idlegame.Contract.LootingPercent(&_Idlegame.CallOpts)
}

// LootingPercent is a free data retrieval call binding the contract method 0xfec8d7d1.
//
// Solidity: function lootingPercent() view returns(uint256)
func (_Idlegame *IdlegameCallerSession) LootingPercent() (*big.Int, error) {
	return _Idlegame.Contract.LootingPercent(&_Idlegame.CallOpts)
}

// LootingTusReward is a free data retrieval call binding the contract method 0xa68f5b6a.
//
// Solidity: function lootingTusReward() view returns(uint128)
func (_Idlegame *IdlegameCaller) LootingTusReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "lootingTusReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LootingTusReward is a free data retrieval call binding the contract method 0xa68f5b6a.
//
// Solidity: function lootingTusReward() view returns(uint128)
func (_Idlegame *IdlegameSession) LootingTusReward() (*big.Int, error) {
	return _Idlegame.Contract.LootingTusReward(&_Idlegame.CallOpts)
}

// LootingTusReward is a free data retrieval call binding the contract method 0xa68f5b6a.
//
// Solidity: function lootingTusReward() view returns(uint128)
func (_Idlegame *IdlegameCallerSession) LootingTusReward() (*big.Int, error) {
	return _Idlegame.Contract.LootingTusReward(&_Idlegame.CallOpts)
}

// MiningDuration is a free data retrieval call binding the contract method 0x7dd8a48f.
//
// Solidity: function miningDuration() view returns(uint128)
func (_Idlegame *IdlegameCaller) MiningDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "miningDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningDuration is a free data retrieval call binding the contract method 0x7dd8a48f.
//
// Solidity: function miningDuration() view returns(uint128)
func (_Idlegame *IdlegameSession) MiningDuration() (*big.Int, error) {
	return _Idlegame.Contract.MiningDuration(&_Idlegame.CallOpts)
}

// MiningDuration is a free data retrieval call binding the contract method 0x7dd8a48f.
//
// Solidity: function miningDuration() view returns(uint128)
func (_Idlegame *IdlegameCallerSession) MiningDuration() (*big.Int, error) {
	return _Idlegame.Contract.MiningDuration(&_Idlegame.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Idlegame *IdlegameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Idlegame *IdlegameSession) Owner() (common.Address, error) {
	return _Idlegame.Contract.Owner(&_Idlegame.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Idlegame *IdlegameCallerSession) Owner() (common.Address, error) {
	return _Idlegame.Contract.Owner(&_Idlegame.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 crabadaId) view returns(address)
func (_Idlegame *IdlegameCaller) OwnerOf(opts *bind.CallOpts, crabadaId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "ownerOf", crabadaId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 crabadaId) view returns(address)
func (_Idlegame *IdlegameSession) OwnerOf(crabadaId *big.Int) (common.Address, error) {
	return _Idlegame.Contract.OwnerOf(&_Idlegame.CallOpts, crabadaId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 crabadaId) view returns(address)
func (_Idlegame *IdlegameCallerSession) OwnerOf(crabadaId *big.Int) (common.Address, error) {
	return _Idlegame.Contract.OwnerOf(&_Idlegame.CallOpts, crabadaId)
}

// ReinforceLockDuration is a free data retrieval call binding the contract method 0xe55cebde.
//
// Solidity: function reinforceLockDuration() view returns(uint128)
func (_Idlegame *IdlegameCaller) ReinforceLockDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "reinforceLockDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReinforceLockDuration is a free data retrieval call binding the contract method 0xe55cebde.
//
// Solidity: function reinforceLockDuration() view returns(uint128)
func (_Idlegame *IdlegameSession) ReinforceLockDuration() (*big.Int, error) {
	return _Idlegame.Contract.ReinforceLockDuration(&_Idlegame.CallOpts)
}

// ReinforceLockDuration is a free data retrieval call binding the contract method 0xe55cebde.
//
// Solidity: function reinforceLockDuration() view returns(uint128)
func (_Idlegame *IdlegameCallerSession) ReinforceLockDuration() (*big.Int, error) {
	return _Idlegame.Contract.ReinforceLockDuration(&_Idlegame.CallOpts)
}

// TusToken is a free data retrieval call binding the contract method 0xc33cacf4.
//
// Solidity: function tusToken() view returns(address)
func (_Idlegame *IdlegameCaller) TusToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Idlegame.contract.Call(opts, &out, "tusToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TusToken is a free data retrieval call binding the contract method 0xc33cacf4.
//
// Solidity: function tusToken() view returns(address)
func (_Idlegame *IdlegameSession) TusToken() (common.Address, error) {
	return _Idlegame.Contract.TusToken(&_Idlegame.CallOpts)
}

// TusToken is a free data retrieval call binding the contract method 0xc33cacf4.
//
// Solidity: function tusToken() view returns(address)
func (_Idlegame *IdlegameCallerSession) TusToken() (common.Address, error) {
	return _Idlegame.Contract.TusToken(&_Idlegame.CallOpts)
}

// AddCrabadaToTeam is a paid mutator transaction binding the contract method 0xc0d8080c.
//
// Solidity: function addCrabadaToTeam(uint256 teamId, uint256 position, uint256 crabadaId) returns()
func (_Idlegame *IdlegameTransactor) AddCrabadaToTeam(opts *bind.TransactOpts, teamId *big.Int, position *big.Int, crabadaId *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "addCrabadaToTeam", teamId, position, crabadaId)
}

// AddCrabadaToTeam is a paid mutator transaction binding the contract method 0xc0d8080c.
//
// Solidity: function addCrabadaToTeam(uint256 teamId, uint256 position, uint256 crabadaId) returns()
func (_Idlegame *IdlegameSession) AddCrabadaToTeam(teamId *big.Int, position *big.Int, crabadaId *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.AddCrabadaToTeam(&_Idlegame.TransactOpts, teamId, position, crabadaId)
}

// AddCrabadaToTeam is a paid mutator transaction binding the contract method 0xc0d8080c.
//
// Solidity: function addCrabadaToTeam(uint256 teamId, uint256 position, uint256 crabadaId) returns()
func (_Idlegame *IdlegameTransactorSession) AddCrabadaToTeam(teamId *big.Int, position *big.Int, crabadaId *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.AddCrabadaToTeam(&_Idlegame.TransactOpts, teamId, position, crabadaId)
}

// Attack is a paid mutator transaction binding the contract method 0xe1fa7638.
//
// Solidity: function attack(uint256 gameId, uint256 attackTeamId) returns()
func (_Idlegame *IdlegameTransactor) Attack(opts *bind.TransactOpts, gameId *big.Int, attackTeamId *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "attack", gameId, attackTeamId)
}

// Attack is a paid mutator transaction binding the contract method 0xe1fa7638.
//
// Solidity: function attack(uint256 gameId, uint256 attackTeamId) returns()
func (_Idlegame *IdlegameSession) Attack(gameId *big.Int, attackTeamId *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.Attack(&_Idlegame.TransactOpts, gameId, attackTeamId)
}

// Attack is a paid mutator transaction binding the contract method 0xe1fa7638.
//
// Solidity: function attack(uint256 gameId, uint256 attackTeamId) returns()
func (_Idlegame *IdlegameTransactorSession) Attack(gameId *big.Int, attackTeamId *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.Attack(&_Idlegame.TransactOpts, gameId, attackTeamId)
}

// CloseGame is a paid mutator transaction binding the contract method 0x2d6ef310.
//
// Solidity: function closeGame(uint256 gameId) returns()
func (_Idlegame *IdlegameTransactor) CloseGame(opts *bind.TransactOpts, gameId *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "closeGame", gameId)
}

// CloseGame is a paid mutator transaction binding the contract method 0x2d6ef310.
//
// Solidity: function closeGame(uint256 gameId) returns()
func (_Idlegame *IdlegameSession) CloseGame(gameId *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.CloseGame(&_Idlegame.TransactOpts, gameId)
}

// CloseGame is a paid mutator transaction binding the contract method 0x2d6ef310.
//
// Solidity: function closeGame(uint256 gameId) returns()
func (_Idlegame *IdlegameTransactorSession) CloseGame(gameId *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.CloseGame(&_Idlegame.TransactOpts, gameId)
}

// CreateTeam is a paid mutator transaction binding the contract method 0xcf034493.
//
// Solidity: function createTeam(uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3) returns(uint256 teamId)
func (_Idlegame *IdlegameTransactor) CreateTeam(opts *bind.TransactOpts, crabadaId1 *big.Int, crabadaId2 *big.Int, crabadaId3 *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "createTeam", crabadaId1, crabadaId2, crabadaId3)
}

// CreateTeam is a paid mutator transaction binding the contract method 0xcf034493.
//
// Solidity: function createTeam(uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3) returns(uint256 teamId)
func (_Idlegame *IdlegameSession) CreateTeam(crabadaId1 *big.Int, crabadaId2 *big.Int, crabadaId3 *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.CreateTeam(&_Idlegame.TransactOpts, crabadaId1, crabadaId2, crabadaId3)
}

// CreateTeam is a paid mutator transaction binding the contract method 0xcf034493.
//
// Solidity: function createTeam(uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3) returns(uint256 teamId)
func (_Idlegame *IdlegameTransactorSession) CreateTeam(crabadaId1 *big.Int, crabadaId2 *big.Int, crabadaId3 *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.CreateTeam(&_Idlegame.TransactOpts, crabadaId1, crabadaId2, crabadaId3)
}

// Deposit is a paid mutator transaction binding the contract method 0x598b8e71.
//
// Solidity: function deposit(uint256[] crabadaIds) returns()
func (_Idlegame *IdlegameTransactor) Deposit(opts *bind.TransactOpts, crabadaIds []*big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "deposit", crabadaIds)
}

// Deposit is a paid mutator transaction binding the contract method 0x598b8e71.
//
// Solidity: function deposit(uint256[] crabadaIds) returns()
func (_Idlegame *IdlegameSession) Deposit(crabadaIds []*big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.Deposit(&_Idlegame.TransactOpts, crabadaIds)
}

// Deposit is a paid mutator transaction binding the contract method 0x598b8e71.
//
// Solidity: function deposit(uint256[] crabadaIds) returns()
func (_Idlegame *IdlegameTransactorSession) Deposit(crabadaIds []*big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.Deposit(&_Idlegame.TransactOpts, crabadaIds)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _crabada, address _cra, address _tus, address _gameStat, address owner) returns()
func (_Idlegame *IdlegameTransactor) Initialize(opts *bind.TransactOpts, _crabada common.Address, _cra common.Address, _tus common.Address, _gameStat common.Address, owner common.Address) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "initialize", _crabada, _cra, _tus, _gameStat, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _crabada, address _cra, address _tus, address _gameStat, address owner) returns()
func (_Idlegame *IdlegameSession) Initialize(_crabada common.Address, _cra common.Address, _tus common.Address, _gameStat common.Address, owner common.Address) (*types.Transaction, error) {
	return _Idlegame.Contract.Initialize(&_Idlegame.TransactOpts, _crabada, _cra, _tus, _gameStat, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _crabada, address _cra, address _tus, address _gameStat, address owner) returns()
func (_Idlegame *IdlegameTransactorSession) Initialize(_crabada common.Address, _cra common.Address, _tus common.Address, _gameStat common.Address, owner common.Address) (*types.Transaction, error) {
	return _Idlegame.Contract.Initialize(&_Idlegame.TransactOpts, _crabada, _cra, _tus, _gameStat, owner)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Idlegame *IdlegameTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Idlegame *IdlegameSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Idlegame.Contract.OnERC721Received(&_Idlegame.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Idlegame *IdlegameTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Idlegame.Contract.OnERC721Received(&_Idlegame.TransactOpts, operator, from, tokenId, data)
}

// ReinforceAttack is a paid mutator transaction binding the contract method 0x3dc8d5ce.
//
// Solidity: function reinforceAttack(uint256 gameId, uint256 crabadaId, uint256 borrowPrice) returns()
func (_Idlegame *IdlegameTransactor) ReinforceAttack(opts *bind.TransactOpts, gameId *big.Int, crabadaId *big.Int, borrowPrice *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "reinforceAttack", gameId, crabadaId, borrowPrice)
}

// ReinforceAttack is a paid mutator transaction binding the contract method 0x3dc8d5ce.
//
// Solidity: function reinforceAttack(uint256 gameId, uint256 crabadaId, uint256 borrowPrice) returns()
func (_Idlegame *IdlegameSession) ReinforceAttack(gameId *big.Int, crabadaId *big.Int, borrowPrice *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.ReinforceAttack(&_Idlegame.TransactOpts, gameId, crabadaId, borrowPrice)
}

// ReinforceAttack is a paid mutator transaction binding the contract method 0x3dc8d5ce.
//
// Solidity: function reinforceAttack(uint256 gameId, uint256 crabadaId, uint256 borrowPrice) returns()
func (_Idlegame *IdlegameTransactorSession) ReinforceAttack(gameId *big.Int, crabadaId *big.Int, borrowPrice *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.ReinforceAttack(&_Idlegame.TransactOpts, gameId, crabadaId, borrowPrice)
}

// ReinforceDefense is a paid mutator transaction binding the contract method 0x08873bfb.
//
// Solidity: function reinforceDefense(uint256 gameId, uint256 crabadaId, uint256 borrowPrice) returns()
func (_Idlegame *IdlegameTransactor) ReinforceDefense(opts *bind.TransactOpts, gameId *big.Int, crabadaId *big.Int, borrowPrice *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "reinforceDefense", gameId, crabadaId, borrowPrice)
}

// ReinforceDefense is a paid mutator transaction binding the contract method 0x08873bfb.
//
// Solidity: function reinforceDefense(uint256 gameId, uint256 crabadaId, uint256 borrowPrice) returns()
func (_Idlegame *IdlegameSession) ReinforceDefense(gameId *big.Int, crabadaId *big.Int, borrowPrice *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.ReinforceDefense(&_Idlegame.TransactOpts, gameId, crabadaId, borrowPrice)
}

// ReinforceDefense is a paid mutator transaction binding the contract method 0x08873bfb.
//
// Solidity: function reinforceDefense(uint256 gameId, uint256 crabadaId, uint256 borrowPrice) returns()
func (_Idlegame *IdlegameTransactorSession) ReinforceDefense(gameId *big.Int, crabadaId *big.Int, borrowPrice *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.ReinforceDefense(&_Idlegame.TransactOpts, gameId, crabadaId, borrowPrice)
}

// RemoveCrabadaFromTeam is a paid mutator transaction binding the contract method 0x31c1bf82.
//
// Solidity: function removeCrabadaFromTeam(uint256 teamId, uint256 position) returns()
func (_Idlegame *IdlegameTransactor) RemoveCrabadaFromTeam(opts *bind.TransactOpts, teamId *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "removeCrabadaFromTeam", teamId, position)
}

// RemoveCrabadaFromTeam is a paid mutator transaction binding the contract method 0x31c1bf82.
//
// Solidity: function removeCrabadaFromTeam(uint256 teamId, uint256 position) returns()
func (_Idlegame *IdlegameSession) RemoveCrabadaFromTeam(teamId *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.RemoveCrabadaFromTeam(&_Idlegame.TransactOpts, teamId, position)
}

// RemoveCrabadaFromTeam is a paid mutator transaction binding the contract method 0x31c1bf82.
//
// Solidity: function removeCrabadaFromTeam(uint256 teamId, uint256 position) returns()
func (_Idlegame *IdlegameTransactorSession) RemoveCrabadaFromTeam(teamId *big.Int, position *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.RemoveCrabadaFromTeam(&_Idlegame.TransactOpts, teamId, position)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Idlegame *IdlegameTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Idlegame *IdlegameSession) RenounceOwnership() (*types.Transaction, error) {
	return _Idlegame.Contract.RenounceOwnership(&_Idlegame.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Idlegame *IdlegameTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Idlegame.Contract.RenounceOwnership(&_Idlegame.TransactOpts)
}

// SetAttackCooldownDuration is a paid mutator transaction binding the contract method 0xbb11c5cb.
//
// Solidity: function setAttackCooldownDuration(uint128 duration) returns()
func (_Idlegame *IdlegameTransactor) SetAttackCooldownDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "setAttackCooldownDuration", duration)
}

// SetAttackCooldownDuration is a paid mutator transaction binding the contract method 0xbb11c5cb.
//
// Solidity: function setAttackCooldownDuration(uint128 duration) returns()
func (_Idlegame *IdlegameSession) SetAttackCooldownDuration(duration *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetAttackCooldownDuration(&_Idlegame.TransactOpts, duration)
}

// SetAttackCooldownDuration is a paid mutator transaction binding the contract method 0xbb11c5cb.
//
// Solidity: function setAttackCooldownDuration(uint128 duration) returns()
func (_Idlegame *IdlegameTransactorSession) SetAttackCooldownDuration(duration *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetAttackCooldownDuration(&_Idlegame.TransactOpts, duration)
}

// SetCRAIncentiveStatus is a paid mutator transaction binding the contract method 0xe2c41cee.
//
// Solidity: function setCRAIncentiveStatus(bool isEnable) returns()
func (_Idlegame *IdlegameTransactor) SetCRAIncentiveStatus(opts *bind.TransactOpts, isEnable bool) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "setCRAIncentiveStatus", isEnable)
}

// SetCRAIncentiveStatus is a paid mutator transaction binding the contract method 0xe2c41cee.
//
// Solidity: function setCRAIncentiveStatus(bool isEnable) returns()
func (_Idlegame *IdlegameSession) SetCRAIncentiveStatus(isEnable bool) (*types.Transaction, error) {
	return _Idlegame.Contract.SetCRAIncentiveStatus(&_Idlegame.TransactOpts, isEnable)
}

// SetCRAIncentiveStatus is a paid mutator transaction binding the contract method 0xe2c41cee.
//
// Solidity: function setCRAIncentiveStatus(bool isEnable) returns()
func (_Idlegame *IdlegameTransactorSession) SetCRAIncentiveStatus(isEnable bool) (*types.Transaction, error) {
	return _Idlegame.Contract.SetCRAIncentiveStatus(&_Idlegame.TransactOpts, isEnable)
}

// SetLendingFeeHolerAddress is a paid mutator transaction binding the contract method 0x8ea53ee4.
//
// Solidity: function setLendingFeeHolerAddress(address newAddr) returns()
func (_Idlegame *IdlegameTransactor) SetLendingFeeHolerAddress(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "setLendingFeeHolerAddress", newAddr)
}

// SetLendingFeeHolerAddress is a paid mutator transaction binding the contract method 0x8ea53ee4.
//
// Solidity: function setLendingFeeHolerAddress(address newAddr) returns()
func (_Idlegame *IdlegameSession) SetLendingFeeHolerAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Idlegame.Contract.SetLendingFeeHolerAddress(&_Idlegame.TransactOpts, newAddr)
}

// SetLendingFeeHolerAddress is a paid mutator transaction binding the contract method 0x8ea53ee4.
//
// Solidity: function setLendingFeeHolerAddress(address newAddr) returns()
func (_Idlegame *IdlegameTransactorSession) SetLendingFeeHolerAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Idlegame.Contract.SetLendingFeeHolerAddress(&_Idlegame.TransactOpts, newAddr)
}

// SetLendingPrice is a paid mutator transaction binding the contract method 0xa9686101.
//
// Solidity: function setLendingPrice(uint256 crabadaId, uint256 price) returns()
func (_Idlegame *IdlegameTransactor) SetLendingPrice(opts *bind.TransactOpts, crabadaId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "setLendingPrice", crabadaId, price)
}

// SetLendingPrice is a paid mutator transaction binding the contract method 0xa9686101.
//
// Solidity: function setLendingPrice(uint256 crabadaId, uint256 price) returns()
func (_Idlegame *IdlegameSession) SetLendingPrice(crabadaId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetLendingPrice(&_Idlegame.TransactOpts, crabadaId, price)
}

// SetLendingPrice is a paid mutator transaction binding the contract method 0xa9686101.
//
// Solidity: function setLendingPrice(uint256 crabadaId, uint256 price) returns()
func (_Idlegame *IdlegameTransactorSession) SetLendingPrice(crabadaId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetLendingPrice(&_Idlegame.TransactOpts, crabadaId, price)
}

// SetLootingPercent is a paid mutator transaction binding the contract method 0x7d9aee2a.
//
// Solidity: function setLootingPercent(uint256 percent) returns()
func (_Idlegame *IdlegameTransactor) SetLootingPercent(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "setLootingPercent", percent)
}

// SetLootingPercent is a paid mutator transaction binding the contract method 0x7d9aee2a.
//
// Solidity: function setLootingPercent(uint256 percent) returns()
func (_Idlegame *IdlegameSession) SetLootingPercent(percent *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetLootingPercent(&_Idlegame.TransactOpts, percent)
}

// SetLootingPercent is a paid mutator transaction binding the contract method 0x7d9aee2a.
//
// Solidity: function setLootingPercent(uint256 percent) returns()
func (_Idlegame *IdlegameTransactorSession) SetLootingPercent(percent *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetLootingPercent(&_Idlegame.TransactOpts, percent)
}

// SetMiningDuration is a paid mutator transaction binding the contract method 0x4b18970a.
//
// Solidity: function setMiningDuration(uint128 duration) returns()
func (_Idlegame *IdlegameTransactor) SetMiningDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "setMiningDuration", duration)
}

// SetMiningDuration is a paid mutator transaction binding the contract method 0x4b18970a.
//
// Solidity: function setMiningDuration(uint128 duration) returns()
func (_Idlegame *IdlegameSession) SetMiningDuration(duration *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetMiningDuration(&_Idlegame.TransactOpts, duration)
}

// SetMiningDuration is a paid mutator transaction binding the contract method 0x4b18970a.
//
// Solidity: function setMiningDuration(uint128 duration) returns()
func (_Idlegame *IdlegameTransactorSession) SetMiningDuration(duration *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetMiningDuration(&_Idlegame.TransactOpts, duration)
}

// SetReinforceLockDuration is a paid mutator transaction binding the contract method 0x2938f7c7.
//
// Solidity: function setReinforceLockDuration(uint128 duration) returns()
func (_Idlegame *IdlegameTransactor) SetReinforceLockDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "setReinforceLockDuration", duration)
}

// SetReinforceLockDuration is a paid mutator transaction binding the contract method 0x2938f7c7.
//
// Solidity: function setReinforceLockDuration(uint128 duration) returns()
func (_Idlegame *IdlegameSession) SetReinforceLockDuration(duration *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetReinforceLockDuration(&_Idlegame.TransactOpts, duration)
}

// SetReinforceLockDuration is a paid mutator transaction binding the contract method 0x2938f7c7.
//
// Solidity: function setReinforceLockDuration(uint128 duration) returns()
func (_Idlegame *IdlegameTransactorSession) SetReinforceLockDuration(duration *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetReinforceLockDuration(&_Idlegame.TransactOpts, duration)
}

// SetRewardLooting is a paid mutator transaction binding the contract method 0x5ba4e7c8.
//
// Solidity: function setRewardLooting(uint128 craRewardAmount, uint128 tusRewardAmount) returns()
func (_Idlegame *IdlegameTransactor) SetRewardLooting(opts *bind.TransactOpts, craRewardAmount *big.Int, tusRewardAmount *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "setRewardLooting", craRewardAmount, tusRewardAmount)
}

// SetRewardLooting is a paid mutator transaction binding the contract method 0x5ba4e7c8.
//
// Solidity: function setRewardLooting(uint128 craRewardAmount, uint128 tusRewardAmount) returns()
func (_Idlegame *IdlegameSession) SetRewardLooting(craRewardAmount *big.Int, tusRewardAmount *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetRewardLooting(&_Idlegame.TransactOpts, craRewardAmount, tusRewardAmount)
}

// SetRewardLooting is a paid mutator transaction binding the contract method 0x5ba4e7c8.
//
// Solidity: function setRewardLooting(uint128 craRewardAmount, uint128 tusRewardAmount) returns()
func (_Idlegame *IdlegameTransactorSession) SetRewardLooting(craRewardAmount *big.Int, tusRewardAmount *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetRewardLooting(&_Idlegame.TransactOpts, craRewardAmount, tusRewardAmount)
}

// SetRewardMining is a paid mutator transaction binding the contract method 0x68dedbac.
//
// Solidity: function setRewardMining(uint128 craRewardAmount, uint128 tusRewardAmount) returns()
func (_Idlegame *IdlegameTransactor) SetRewardMining(opts *bind.TransactOpts, craRewardAmount *big.Int, tusRewardAmount *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "setRewardMining", craRewardAmount, tusRewardAmount)
}

// SetRewardMining is a paid mutator transaction binding the contract method 0x68dedbac.
//
// Solidity: function setRewardMining(uint128 craRewardAmount, uint128 tusRewardAmount) returns()
func (_Idlegame *IdlegameSession) SetRewardMining(craRewardAmount *big.Int, tusRewardAmount *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetRewardMining(&_Idlegame.TransactOpts, craRewardAmount, tusRewardAmount)
}

// SetRewardMining is a paid mutator transaction binding the contract method 0x68dedbac.
//
// Solidity: function setRewardMining(uint128 craRewardAmount, uint128 tusRewardAmount) returns()
func (_Idlegame *IdlegameTransactorSession) SetRewardMining(craRewardAmount *big.Int, tusRewardAmount *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetRewardMining(&_Idlegame.TransactOpts, craRewardAmount, tusRewardAmount)
}

// SetStepDuration is a paid mutator transaction binding the contract method 0x45228060.
//
// Solidity: function setStepDuration(uint128 duration) returns()
func (_Idlegame *IdlegameTransactor) SetStepDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "setStepDuration", duration)
}

// SetStepDuration is a paid mutator transaction binding the contract method 0x45228060.
//
// Solidity: function setStepDuration(uint128 duration) returns()
func (_Idlegame *IdlegameSession) SetStepDuration(duration *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetStepDuration(&_Idlegame.TransactOpts, duration)
}

// SetStepDuration is a paid mutator transaction binding the contract method 0x45228060.
//
// Solidity: function setStepDuration(uint128 duration) returns()
func (_Idlegame *IdlegameTransactorSession) SetStepDuration(duration *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SetStepDuration(&_Idlegame.TransactOpts, duration)
}

// SettleGame is a paid mutator transaction binding the contract method 0x312d7bbc.
//
// Solidity: function settleGame(uint256 gameId) returns()
func (_Idlegame *IdlegameTransactor) SettleGame(opts *bind.TransactOpts, gameId *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "settleGame", gameId)
}

// SettleGame is a paid mutator transaction binding the contract method 0x312d7bbc.
//
// Solidity: function settleGame(uint256 gameId) returns()
func (_Idlegame *IdlegameSession) SettleGame(gameId *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SettleGame(&_Idlegame.TransactOpts, gameId)
}

// SettleGame is a paid mutator transaction binding the contract method 0x312d7bbc.
//
// Solidity: function settleGame(uint256 gameId) returns()
func (_Idlegame *IdlegameTransactorSession) SettleGame(gameId *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.SettleGame(&_Idlegame.TransactOpts, gameId)
}

// StartGame is a paid mutator transaction binding the contract method 0xe5ed1d59.
//
// Solidity: function startGame(uint256 teamId) returns()
func (_Idlegame *IdlegameTransactor) StartGame(opts *bind.TransactOpts, teamId *big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "startGame", teamId)
}

// StartGame is a paid mutator transaction binding the contract method 0xe5ed1d59.
//
// Solidity: function startGame(uint256 teamId) returns()
func (_Idlegame *IdlegameSession) StartGame(teamId *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.StartGame(&_Idlegame.TransactOpts, teamId)
}

// StartGame is a paid mutator transaction binding the contract method 0xe5ed1d59.
//
// Solidity: function startGame(uint256 teamId) returns()
func (_Idlegame *IdlegameTransactorSession) StartGame(teamId *big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.StartGame(&_Idlegame.TransactOpts, teamId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Idlegame *IdlegameTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Idlegame *IdlegameSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Idlegame.Contract.TransferOwnership(&_Idlegame.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Idlegame *IdlegameTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Idlegame.Contract.TransferOwnership(&_Idlegame.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8293744b.
//
// Solidity: function withdraw(address to, uint256[] crabadaIds) returns()
func (_Idlegame *IdlegameTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, crabadaIds []*big.Int) (*types.Transaction, error) {
	return _Idlegame.contract.Transact(opts, "withdraw", to, crabadaIds)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8293744b.
//
// Solidity: function withdraw(address to, uint256[] crabadaIds) returns()
func (_Idlegame *IdlegameSession) Withdraw(to common.Address, crabadaIds []*big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.Withdraw(&_Idlegame.TransactOpts, to, crabadaIds)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8293744b.
//
// Solidity: function withdraw(address to, uint256[] crabadaIds) returns()
func (_Idlegame *IdlegameTransactorSession) Withdraw(to common.Address, crabadaIds []*big.Int) (*types.Transaction, error) {
	return _Idlegame.Contract.Withdraw(&_Idlegame.TransactOpts, to, crabadaIds)
}

// IdlegameAddCrabadaIterator is returned from FilterAddCrabada and is used to iterate over the raw logs and unpacked data for AddCrabada events raised by the Idlegame contract.
type IdlegameAddCrabadaIterator struct {
	Event *IdlegameAddCrabada // Event containing the contract specifics and raw log

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
func (it *IdlegameAddCrabadaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdlegameAddCrabada)
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
		it.Event = new(IdlegameAddCrabada)
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
func (it *IdlegameAddCrabadaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdlegameAddCrabadaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdlegameAddCrabada represents a AddCrabada event raised by the Idlegame contract.
type IdlegameAddCrabada struct {
	TeamId      *big.Int
	Position    *big.Int
	CrabadaId   *big.Int
	BattlePoint *big.Int
	TimePoint   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddCrabada is a free log retrieval operation binding the contract event 0x95f06f23548778c52447db2dca1c565debe55244f3904032ac45c61df764c9a2.
//
// Solidity: event AddCrabada(uint256 teamId, uint256 position, uint256 crabadaId, uint256 battlePoint, uint256 timePoint)
func (_Idlegame *IdlegameFilterer) FilterAddCrabada(opts *bind.FilterOpts) (*IdlegameAddCrabadaIterator, error) {

	logs, sub, err := _Idlegame.contract.FilterLogs(opts, "AddCrabada")
	if err != nil {
		return nil, err
	}
	return &IdlegameAddCrabadaIterator{contract: _Idlegame.contract, event: "AddCrabada", logs: logs, sub: sub}, nil
}

// WatchAddCrabada is a free log subscription operation binding the contract event 0x95f06f23548778c52447db2dca1c565debe55244f3904032ac45c61df764c9a2.
//
// Solidity: event AddCrabada(uint256 teamId, uint256 position, uint256 crabadaId, uint256 battlePoint, uint256 timePoint)
func (_Idlegame *IdlegameFilterer) WatchAddCrabada(opts *bind.WatchOpts, sink chan<- *IdlegameAddCrabada) (event.Subscription, error) {

	logs, sub, err := _Idlegame.contract.WatchLogs(opts, "AddCrabada")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdlegameAddCrabada)
				if err := _Idlegame.contract.UnpackLog(event, "AddCrabada", log); err != nil {
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

// ParseAddCrabada is a log parse operation binding the contract event 0x95f06f23548778c52447db2dca1c565debe55244f3904032ac45c61df764c9a2.
//
// Solidity: event AddCrabada(uint256 teamId, uint256 position, uint256 crabadaId, uint256 battlePoint, uint256 timePoint)
func (_Idlegame *IdlegameFilterer) ParseAddCrabada(log types.Log) (*IdlegameAddCrabada, error) {
	event := new(IdlegameAddCrabada)
	if err := _Idlegame.contract.UnpackLog(event, "AddCrabada", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdlegameCloseGameIterator is returned from FilterCloseGame and is used to iterate over the raw logs and unpacked data for CloseGame events raised by the Idlegame contract.
type IdlegameCloseGameIterator struct {
	Event *IdlegameCloseGame // Event containing the contract specifics and raw log

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
func (it *IdlegameCloseGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdlegameCloseGame)
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
		it.Event = new(IdlegameCloseGame)
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
func (it *IdlegameCloseGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdlegameCloseGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdlegameCloseGame represents a CloseGame event raised by the Idlegame contract.
type IdlegameCloseGame struct {
	GameId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCloseGame is a free log retrieval operation binding the contract event 0x827e8bff9f46048f8351964d2c871d09e4f4231513cf2fcb786649c68732e24f.
//
// Solidity: event CloseGame(uint256 gameId)
func (_Idlegame *IdlegameFilterer) FilterCloseGame(opts *bind.FilterOpts) (*IdlegameCloseGameIterator, error) {

	logs, sub, err := _Idlegame.contract.FilterLogs(opts, "CloseGame")
	if err != nil {
		return nil, err
	}
	return &IdlegameCloseGameIterator{contract: _Idlegame.contract, event: "CloseGame", logs: logs, sub: sub}, nil
}

// WatchCloseGame is a free log subscription operation binding the contract event 0x827e8bff9f46048f8351964d2c871d09e4f4231513cf2fcb786649c68732e24f.
//
// Solidity: event CloseGame(uint256 gameId)
func (_Idlegame *IdlegameFilterer) WatchCloseGame(opts *bind.WatchOpts, sink chan<- *IdlegameCloseGame) (event.Subscription, error) {

	logs, sub, err := _Idlegame.contract.WatchLogs(opts, "CloseGame")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdlegameCloseGame)
				if err := _Idlegame.contract.UnpackLog(event, "CloseGame", log); err != nil {
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

// ParseCloseGame is a log parse operation binding the contract event 0x827e8bff9f46048f8351964d2c871d09e4f4231513cf2fcb786649c68732e24f.
//
// Solidity: event CloseGame(uint256 gameId)
func (_Idlegame *IdlegameFilterer) ParseCloseGame(log types.Log) (*IdlegameCloseGame, error) {
	event := new(IdlegameCloseGame)
	if err := _Idlegame.contract.UnpackLog(event, "CloseGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdlegameCreateTeamIterator is returned from FilterCreateTeam and is used to iterate over the raw logs and unpacked data for CreateTeam events raised by the Idlegame contract.
type IdlegameCreateTeamIterator struct {
	Event *IdlegameCreateTeam // Event containing the contract specifics and raw log

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
func (it *IdlegameCreateTeamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdlegameCreateTeam)
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
		it.Event = new(IdlegameCreateTeam)
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
func (it *IdlegameCreateTeamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdlegameCreateTeamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdlegameCreateTeam represents a CreateTeam event raised by the Idlegame contract.
type IdlegameCreateTeam struct {
	Owner       common.Address
	TeamId      *big.Int
	CrabadaId1  *big.Int
	CrabadaId2  *big.Int
	CrabadaId3  *big.Int
	BattlePoint *big.Int
	TimePoint   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreateTeam is a free log retrieval operation binding the contract event 0xe2f710515de5fe7937eb9d7ff8093cb7cd011b496de6574d80e51440784e2a71.
//
// Solidity: event CreateTeam(address owner, uint256 teamId, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint256 battlePoint, uint256 timePoint)
func (_Idlegame *IdlegameFilterer) FilterCreateTeam(opts *bind.FilterOpts) (*IdlegameCreateTeamIterator, error) {

	logs, sub, err := _Idlegame.contract.FilterLogs(opts, "CreateTeam")
	if err != nil {
		return nil, err
	}
	return &IdlegameCreateTeamIterator{contract: _Idlegame.contract, event: "CreateTeam", logs: logs, sub: sub}, nil
}

// WatchCreateTeam is a free log subscription operation binding the contract event 0xe2f710515de5fe7937eb9d7ff8093cb7cd011b496de6574d80e51440784e2a71.
//
// Solidity: event CreateTeam(address owner, uint256 teamId, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint256 battlePoint, uint256 timePoint)
func (_Idlegame *IdlegameFilterer) WatchCreateTeam(opts *bind.WatchOpts, sink chan<- *IdlegameCreateTeam) (event.Subscription, error) {

	logs, sub, err := _Idlegame.contract.WatchLogs(opts, "CreateTeam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdlegameCreateTeam)
				if err := _Idlegame.contract.UnpackLog(event, "CreateTeam", log); err != nil {
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

// ParseCreateTeam is a log parse operation binding the contract event 0xe2f710515de5fe7937eb9d7ff8093cb7cd011b496de6574d80e51440784e2a71.
//
// Solidity: event CreateTeam(address owner, uint256 teamId, uint256 crabadaId1, uint256 crabadaId2, uint256 crabadaId3, uint256 battlePoint, uint256 timePoint)
func (_Idlegame *IdlegameFilterer) ParseCreateTeam(log types.Log) (*IdlegameCreateTeam, error) {
	event := new(IdlegameCreateTeam)
	if err := _Idlegame.contract.UnpackLog(event, "CreateTeam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdlegameDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Idlegame contract.
type IdlegameDepositIterator struct {
	Event *IdlegameDeposit // Event containing the contract specifics and raw log

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
func (it *IdlegameDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdlegameDeposit)
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
		it.Event = new(IdlegameDeposit)
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
func (it *IdlegameDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdlegameDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdlegameDeposit represents a Deposit event raised by the Idlegame contract.
type IdlegameDeposit struct {
	CrabadaId *big.Int
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4bcc17093cdf51079c755de089be5a85e70fa374ec656c194480fbdcda224a53.
//
// Solidity: event Deposit(uint256 crabadaId, address owner)
func (_Idlegame *IdlegameFilterer) FilterDeposit(opts *bind.FilterOpts) (*IdlegameDepositIterator, error) {

	logs, sub, err := _Idlegame.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &IdlegameDepositIterator{contract: _Idlegame.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4bcc17093cdf51079c755de089be5a85e70fa374ec656c194480fbdcda224a53.
//
// Solidity: event Deposit(uint256 crabadaId, address owner)
func (_Idlegame *IdlegameFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *IdlegameDeposit) (event.Subscription, error) {

	logs, sub, err := _Idlegame.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdlegameDeposit)
				if err := _Idlegame.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x4bcc17093cdf51079c755de089be5a85e70fa374ec656c194480fbdcda224a53.
//
// Solidity: event Deposit(uint256 crabadaId, address owner)
func (_Idlegame *IdlegameFilterer) ParseDeposit(log types.Log) (*IdlegameDeposit, error) {
	event := new(IdlegameDeposit)
	if err := _Idlegame.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdlegameFightIterator is returned from FilterFight and is used to iterate over the raw logs and unpacked data for Fight events raised by the Idlegame contract.
type IdlegameFightIterator struct {
	Event *IdlegameFight // Event containing the contract specifics and raw log

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
func (it *IdlegameFightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdlegameFight)
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
		it.Event = new(IdlegameFight)
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
func (it *IdlegameFightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdlegameFightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdlegameFight represents a Fight event raised by the Idlegame contract.
type IdlegameFight struct {
	GameId        *big.Int
	Turn          *big.Int
	AttackTeamId  *big.Int
	DefenseTeamId *big.Int
	SoldierId     *big.Int
	AttackTime    *big.Int
	AttackPoint   uint16
	DefensePoint  uint16
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFight is a free log retrieval operation binding the contract event 0xc9116f8ab51c97fa0f2ebd6cc9f75395464ae81bebc69d2d468b508a460a7136.
//
// Solidity: event Fight(uint256 gameId, uint256 turn, uint256 attackTeamId, uint256 defenseTeamId, uint256 soldierId, uint256 attackTime, uint16 attackPoint, uint16 defensePoint)
func (_Idlegame *IdlegameFilterer) FilterFight(opts *bind.FilterOpts) (*IdlegameFightIterator, error) {

	logs, sub, err := _Idlegame.contract.FilterLogs(opts, "Fight")
	if err != nil {
		return nil, err
	}
	return &IdlegameFightIterator{contract: _Idlegame.contract, event: "Fight", logs: logs, sub: sub}, nil
}

// WatchFight is a free log subscription operation binding the contract event 0xc9116f8ab51c97fa0f2ebd6cc9f75395464ae81bebc69d2d468b508a460a7136.
//
// Solidity: event Fight(uint256 gameId, uint256 turn, uint256 attackTeamId, uint256 defenseTeamId, uint256 soldierId, uint256 attackTime, uint16 attackPoint, uint16 defensePoint)
func (_Idlegame *IdlegameFilterer) WatchFight(opts *bind.WatchOpts, sink chan<- *IdlegameFight) (event.Subscription, error) {

	logs, sub, err := _Idlegame.contract.WatchLogs(opts, "Fight")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdlegameFight)
				if err := _Idlegame.contract.UnpackLog(event, "Fight", log); err != nil {
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

// ParseFight is a log parse operation binding the contract event 0xc9116f8ab51c97fa0f2ebd6cc9f75395464ae81bebc69d2d468b508a460a7136.
//
// Solidity: event Fight(uint256 gameId, uint256 turn, uint256 attackTeamId, uint256 defenseTeamId, uint256 soldierId, uint256 attackTime, uint16 attackPoint, uint16 defensePoint)
func (_Idlegame *IdlegameFilterer) ParseFight(log types.Log) (*IdlegameFight, error) {
	event := new(IdlegameFight)
	if err := _Idlegame.contract.UnpackLog(event, "Fight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdlegameLendIterator is returned from FilterLend and is used to iterate over the raw logs and unpacked data for Lend events raised by the Idlegame contract.
type IdlegameLendIterator struct {
	Event *IdlegameLend // Event containing the contract specifics and raw log

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
func (it *IdlegameLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdlegameLend)
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
		it.Event = new(IdlegameLend)
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
func (it *IdlegameLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdlegameLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdlegameLend represents a Lend event raised by the Idlegame contract.
type IdlegameLend struct {
	Borrower  common.Address
	CrabadaId *big.Int
	GameId    *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLend is a free log retrieval operation binding the contract event 0xaf77103a98a12032881b6214985e03ff85a9c1ad6fed28f4b0f3c52e01bcc58b.
//
// Solidity: event Lend(address borrower, uint256 crabadaId, uint256 gameId, uint256 price)
func (_Idlegame *IdlegameFilterer) FilterLend(opts *bind.FilterOpts) (*IdlegameLendIterator, error) {

	logs, sub, err := _Idlegame.contract.FilterLogs(opts, "Lend")
	if err != nil {
		return nil, err
	}
	return &IdlegameLendIterator{contract: _Idlegame.contract, event: "Lend", logs: logs, sub: sub}, nil
}

// WatchLend is a free log subscription operation binding the contract event 0xaf77103a98a12032881b6214985e03ff85a9c1ad6fed28f4b0f3c52e01bcc58b.
//
// Solidity: event Lend(address borrower, uint256 crabadaId, uint256 gameId, uint256 price)
func (_Idlegame *IdlegameFilterer) WatchLend(opts *bind.WatchOpts, sink chan<- *IdlegameLend) (event.Subscription, error) {

	logs, sub, err := _Idlegame.contract.WatchLogs(opts, "Lend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdlegameLend)
				if err := _Idlegame.contract.UnpackLog(event, "Lend", log); err != nil {
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

// ParseLend is a log parse operation binding the contract event 0xaf77103a98a12032881b6214985e03ff85a9c1ad6fed28f4b0f3c52e01bcc58b.
//
// Solidity: event Lend(address borrower, uint256 crabadaId, uint256 gameId, uint256 price)
func (_Idlegame *IdlegameFilterer) ParseLend(log types.Log) (*IdlegameLend, error) {
	event := new(IdlegameLend)
	if err := _Idlegame.contract.UnpackLog(event, "Lend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdlegameOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Idlegame contract.
type IdlegameOwnershipTransferredIterator struct {
	Event *IdlegameOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdlegameOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdlegameOwnershipTransferred)
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
		it.Event = new(IdlegameOwnershipTransferred)
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
func (it *IdlegameOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdlegameOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdlegameOwnershipTransferred represents a OwnershipTransferred event raised by the Idlegame contract.
type IdlegameOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Idlegame *IdlegameFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdlegameOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Idlegame.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdlegameOwnershipTransferredIterator{contract: _Idlegame.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Idlegame *IdlegameFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdlegameOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Idlegame.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdlegameOwnershipTransferred)
				if err := _Idlegame.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Idlegame *IdlegameFilterer) ParseOwnershipTransferred(log types.Log) (*IdlegameOwnershipTransferred, error) {
	event := new(IdlegameOwnershipTransferred)
	if err := _Idlegame.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdlegameRemoveCrabadaIterator is returned from FilterRemoveCrabada and is used to iterate over the raw logs and unpacked data for RemoveCrabada events raised by the Idlegame contract.
type IdlegameRemoveCrabadaIterator struct {
	Event *IdlegameRemoveCrabada // Event containing the contract specifics and raw log

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
func (it *IdlegameRemoveCrabadaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdlegameRemoveCrabada)
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
		it.Event = new(IdlegameRemoveCrabada)
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
func (it *IdlegameRemoveCrabadaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdlegameRemoveCrabadaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdlegameRemoveCrabada represents a RemoveCrabada event raised by the Idlegame contract.
type IdlegameRemoveCrabada struct {
	TeamId      *big.Int
	Position    *big.Int
	CrabadaId   *big.Int
	BattlePoint *big.Int
	TimePoint   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveCrabada is a free log retrieval operation binding the contract event 0x164649e040e4f90ed2500b40e25f8e4c1e7991798b4359bf2a16ea564f8ded60.
//
// Solidity: event RemoveCrabada(uint256 teamId, uint256 position, uint256 crabadaId, uint256 battlePoint, uint256 timePoint)
func (_Idlegame *IdlegameFilterer) FilterRemoveCrabada(opts *bind.FilterOpts) (*IdlegameRemoveCrabadaIterator, error) {

	logs, sub, err := _Idlegame.contract.FilterLogs(opts, "RemoveCrabada")
	if err != nil {
		return nil, err
	}
	return &IdlegameRemoveCrabadaIterator{contract: _Idlegame.contract, event: "RemoveCrabada", logs: logs, sub: sub}, nil
}

// WatchRemoveCrabada is a free log subscription operation binding the contract event 0x164649e040e4f90ed2500b40e25f8e4c1e7991798b4359bf2a16ea564f8ded60.
//
// Solidity: event RemoveCrabada(uint256 teamId, uint256 position, uint256 crabadaId, uint256 battlePoint, uint256 timePoint)
func (_Idlegame *IdlegameFilterer) WatchRemoveCrabada(opts *bind.WatchOpts, sink chan<- *IdlegameRemoveCrabada) (event.Subscription, error) {

	logs, sub, err := _Idlegame.contract.WatchLogs(opts, "RemoveCrabada")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdlegameRemoveCrabada)
				if err := _Idlegame.contract.UnpackLog(event, "RemoveCrabada", log); err != nil {
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

// ParseRemoveCrabada is a log parse operation binding the contract event 0x164649e040e4f90ed2500b40e25f8e4c1e7991798b4359bf2a16ea564f8ded60.
//
// Solidity: event RemoveCrabada(uint256 teamId, uint256 position, uint256 crabadaId, uint256 battlePoint, uint256 timePoint)
func (_Idlegame *IdlegameFilterer) ParseRemoveCrabada(log types.Log) (*IdlegameRemoveCrabada, error) {
	event := new(IdlegameRemoveCrabada)
	if err := _Idlegame.contract.UnpackLog(event, "RemoveCrabada", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdlegameSetLendingPriceIterator is returned from FilterSetLendingPrice and is used to iterate over the raw logs and unpacked data for SetLendingPrice events raised by the Idlegame contract.
type IdlegameSetLendingPriceIterator struct {
	Event *IdlegameSetLendingPrice // Event containing the contract specifics and raw log

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
func (it *IdlegameSetLendingPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdlegameSetLendingPrice)
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
		it.Event = new(IdlegameSetLendingPrice)
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
func (it *IdlegameSetLendingPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdlegameSetLendingPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdlegameSetLendingPrice represents a SetLendingPrice event raised by the Idlegame contract.
type IdlegameSetLendingPrice struct {
	CrabadaId *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetLendingPrice is a free log retrieval operation binding the contract event 0x85937e28b96ea9433b7abff68e1485fb68b1a6a61edc4fe945676f48f12b56b9.
//
// Solidity: event SetLendingPrice(uint256 crabadaId, uint256 price)
func (_Idlegame *IdlegameFilterer) FilterSetLendingPrice(opts *bind.FilterOpts) (*IdlegameSetLendingPriceIterator, error) {

	logs, sub, err := _Idlegame.contract.FilterLogs(opts, "SetLendingPrice")
	if err != nil {
		return nil, err
	}
	return &IdlegameSetLendingPriceIterator{contract: _Idlegame.contract, event: "SetLendingPrice", logs: logs, sub: sub}, nil
}

// WatchSetLendingPrice is a free log subscription operation binding the contract event 0x85937e28b96ea9433b7abff68e1485fb68b1a6a61edc4fe945676f48f12b56b9.
//
// Solidity: event SetLendingPrice(uint256 crabadaId, uint256 price)
func (_Idlegame *IdlegameFilterer) WatchSetLendingPrice(opts *bind.WatchOpts, sink chan<- *IdlegameSetLendingPrice) (event.Subscription, error) {

	logs, sub, err := _Idlegame.contract.WatchLogs(opts, "SetLendingPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdlegameSetLendingPrice)
				if err := _Idlegame.contract.UnpackLog(event, "SetLendingPrice", log); err != nil {
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

// ParseSetLendingPrice is a log parse operation binding the contract event 0x85937e28b96ea9433b7abff68e1485fb68b1a6a61edc4fe945676f48f12b56b9.
//
// Solidity: event SetLendingPrice(uint256 crabadaId, uint256 price)
func (_Idlegame *IdlegameFilterer) ParseSetLendingPrice(log types.Log) (*IdlegameSetLendingPrice, error) {
	event := new(IdlegameSetLendingPrice)
	if err := _Idlegame.contract.UnpackLog(event, "SetLendingPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdlegameSettleGameIterator is returned from FilterSettleGame and is used to iterate over the raw logs and unpacked data for SettleGame events raised by the Idlegame contract.
type IdlegameSettleGameIterator struct {
	Event *IdlegameSettleGame // Event containing the contract specifics and raw log

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
func (it *IdlegameSettleGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdlegameSettleGame)
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
		it.Event = new(IdlegameSettleGame)
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
func (it *IdlegameSettleGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdlegameSettleGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdlegameSettleGame represents a SettleGame event raised by the Idlegame contract.
type IdlegameSettleGame struct {
	GameId          *big.Int
	WinnerTeamId    *big.Int
	WinnerCraReward *big.Int
	WinnerTusReward *big.Int
	LoserCraReward  *big.Int
	LoserTusReward  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSettleGame is a free log retrieval operation binding the contract event 0xf8abd91a64a8ff03c6198f7f7625d9c2b26e964546d6624e997e2d9ae84107b3.
//
// Solidity: event SettleGame(uint256 gameId, uint256 winnerTeamId, uint256 winnerCraReward, uint256 winnerTusReward, uint256 loserCraReward, uint256 loserTusReward)
func (_Idlegame *IdlegameFilterer) FilterSettleGame(opts *bind.FilterOpts) (*IdlegameSettleGameIterator, error) {

	logs, sub, err := _Idlegame.contract.FilterLogs(opts, "SettleGame")
	if err != nil {
		return nil, err
	}
	return &IdlegameSettleGameIterator{contract: _Idlegame.contract, event: "SettleGame", logs: logs, sub: sub}, nil
}

// WatchSettleGame is a free log subscription operation binding the contract event 0xf8abd91a64a8ff03c6198f7f7625d9c2b26e964546d6624e997e2d9ae84107b3.
//
// Solidity: event SettleGame(uint256 gameId, uint256 winnerTeamId, uint256 winnerCraReward, uint256 winnerTusReward, uint256 loserCraReward, uint256 loserTusReward)
func (_Idlegame *IdlegameFilterer) WatchSettleGame(opts *bind.WatchOpts, sink chan<- *IdlegameSettleGame) (event.Subscription, error) {

	logs, sub, err := _Idlegame.contract.WatchLogs(opts, "SettleGame")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdlegameSettleGame)
				if err := _Idlegame.contract.UnpackLog(event, "SettleGame", log); err != nil {
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

// ParseSettleGame is a log parse operation binding the contract event 0xf8abd91a64a8ff03c6198f7f7625d9c2b26e964546d6624e997e2d9ae84107b3.
//
// Solidity: event SettleGame(uint256 gameId, uint256 winnerTeamId, uint256 winnerCraReward, uint256 winnerTusReward, uint256 loserCraReward, uint256 loserTusReward)
func (_Idlegame *IdlegameFilterer) ParseSettleGame(log types.Log) (*IdlegameSettleGame, error) {
	event := new(IdlegameSettleGame)
	if err := _Idlegame.contract.UnpackLog(event, "SettleGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdlegameStartGameIterator is returned from FilterStartGame and is used to iterate over the raw logs and unpacked data for StartGame events raised by the Idlegame contract.
type IdlegameStartGameIterator struct {
	Event *IdlegameStartGame // Event containing the contract specifics and raw log

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
func (it *IdlegameStartGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdlegameStartGame)
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
		it.Event = new(IdlegameStartGame)
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
func (it *IdlegameStartGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdlegameStartGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdlegameStartGame represents a StartGame event raised by the Idlegame contract.
type IdlegameStartGame struct {
	GameId    *big.Int
	TeamId    *big.Int
	Duration  *big.Int
	CraReward *big.Int
	TusReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStartGame is a free log retrieval operation binding the contract event 0x0eef6f7452b7d2ee11184579c086fb47626e796a83df2b2e16254df60ab761eb.
//
// Solidity: event StartGame(uint256 gameId, uint256 teamId, uint256 duration, uint256 craReward, uint256 tusReward)
func (_Idlegame *IdlegameFilterer) FilterStartGame(opts *bind.FilterOpts) (*IdlegameStartGameIterator, error) {

	logs, sub, err := _Idlegame.contract.FilterLogs(opts, "StartGame")
	if err != nil {
		return nil, err
	}
	return &IdlegameStartGameIterator{contract: _Idlegame.contract, event: "StartGame", logs: logs, sub: sub}, nil
}

// WatchStartGame is a free log subscription operation binding the contract event 0x0eef6f7452b7d2ee11184579c086fb47626e796a83df2b2e16254df60ab761eb.
//
// Solidity: event StartGame(uint256 gameId, uint256 teamId, uint256 duration, uint256 craReward, uint256 tusReward)
func (_Idlegame *IdlegameFilterer) WatchStartGame(opts *bind.WatchOpts, sink chan<- *IdlegameStartGame) (event.Subscription, error) {

	logs, sub, err := _Idlegame.contract.WatchLogs(opts, "StartGame")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdlegameStartGame)
				if err := _Idlegame.contract.UnpackLog(event, "StartGame", log); err != nil {
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

// ParseStartGame is a log parse operation binding the contract event 0x0eef6f7452b7d2ee11184579c086fb47626e796a83df2b2e16254df60ab761eb.
//
// Solidity: event StartGame(uint256 gameId, uint256 teamId, uint256 duration, uint256 craReward, uint256 tusReward)
func (_Idlegame *IdlegameFilterer) ParseStartGame(log types.Log) (*IdlegameStartGame, error) {
	event := new(IdlegameStartGame)
	if err := _Idlegame.contract.UnpackLog(event, "StartGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdlegameWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Idlegame contract.
type IdlegameWithdrawIterator struct {
	Event *IdlegameWithdraw // Event containing the contract specifics and raw log

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
func (it *IdlegameWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdlegameWithdraw)
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
		it.Event = new(IdlegameWithdraw)
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
func (it *IdlegameWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdlegameWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdlegameWithdraw represents a Withdraw event raised by the Idlegame contract.
type IdlegameWithdraw struct {
	CrabadaId *big.Int
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x8353ffcac0876ad14e226d9783c04540bfebf13871e868157d2a391cad98e918.
//
// Solidity: event Withdraw(uint256 crabadaId, address owner)
func (_Idlegame *IdlegameFilterer) FilterWithdraw(opts *bind.FilterOpts) (*IdlegameWithdrawIterator, error) {

	logs, sub, err := _Idlegame.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &IdlegameWithdrawIterator{contract: _Idlegame.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x8353ffcac0876ad14e226d9783c04540bfebf13871e868157d2a391cad98e918.
//
// Solidity: event Withdraw(uint256 crabadaId, address owner)
func (_Idlegame *IdlegameFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IdlegameWithdraw) (event.Subscription, error) {

	logs, sub, err := _Idlegame.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdlegameWithdraw)
				if err := _Idlegame.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x8353ffcac0876ad14e226d9783c04540bfebf13871e868157d2a391cad98e918.
//
// Solidity: event Withdraw(uint256 crabadaId, address owner)
func (_Idlegame *IdlegameFilterer) ParseWithdraw(log types.Log) (*IdlegameWithdraw, error) {
	event := new(IdlegameWithdraw)
	if err := _Idlegame.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

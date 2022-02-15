// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package traderjoe

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

// TraderjoeMetaData contains all meta data concerning the Traderjoe contract.
var TraderjoeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WAVAX\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WAVAX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityAVAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityAVAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityAVAXSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityAVAXWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapAVAXForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactAVAXForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactAVAXForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForAVAX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForAVAXSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactAVAX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TraderjoeABI is the input ABI used to generate the binding from.
// Deprecated: Use TraderjoeMetaData.ABI instead.
var TraderjoeABI = TraderjoeMetaData.ABI

// Traderjoe is an auto generated Go binding around an Ethereum contract.
type Traderjoe struct {
	TraderjoeCaller     // Read-only binding to the contract
	TraderjoeTransactor // Write-only binding to the contract
	TraderjoeFilterer   // Log filterer for contract events
}

// TraderjoeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraderjoeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraderjoeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraderjoeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraderjoeSession struct {
	Contract     *Traderjoe        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraderjoeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraderjoeCallerSession struct {
	Contract *TraderjoeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TraderjoeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraderjoeTransactorSession struct {
	Contract     *TraderjoeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TraderjoeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraderjoeRaw struct {
	Contract *Traderjoe // Generic contract binding to access the raw methods on
}

// TraderjoeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraderjoeCallerRaw struct {
	Contract *TraderjoeCaller // Generic read-only contract binding to access the raw methods on
}

// TraderjoeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraderjoeTransactorRaw struct {
	Contract *TraderjoeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTraderjoe creates a new instance of Traderjoe, bound to a specific deployed contract.
func NewTraderjoe(address common.Address, backend bind.ContractBackend) (*Traderjoe, error) {
	contract, err := bindTraderjoe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Traderjoe{TraderjoeCaller: TraderjoeCaller{contract: contract}, TraderjoeTransactor: TraderjoeTransactor{contract: contract}, TraderjoeFilterer: TraderjoeFilterer{contract: contract}}, nil
}

// NewTraderjoeCaller creates a new read-only instance of Traderjoe, bound to a specific deployed contract.
func NewTraderjoeCaller(address common.Address, caller bind.ContractCaller) (*TraderjoeCaller, error) {
	contract, err := bindTraderjoe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeCaller{contract: contract}, nil
}

// NewTraderjoeTransactor creates a new write-only instance of Traderjoe, bound to a specific deployed contract.
func NewTraderjoeTransactor(address common.Address, transactor bind.ContractTransactor) (*TraderjoeTransactor, error) {
	contract, err := bindTraderjoe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeTransactor{contract: contract}, nil
}

// NewTraderjoeFilterer creates a new log filterer instance of Traderjoe, bound to a specific deployed contract.
func NewTraderjoeFilterer(address common.Address, filterer bind.ContractFilterer) (*TraderjoeFilterer, error) {
	contract, err := bindTraderjoe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraderjoeFilterer{contract: contract}, nil
}

// bindTraderjoe binds a generic wrapper to an already deployed contract.
func bindTraderjoe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TraderjoeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Traderjoe *TraderjoeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Traderjoe.Contract.TraderjoeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Traderjoe *TraderjoeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Traderjoe.Contract.TraderjoeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Traderjoe *TraderjoeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Traderjoe.Contract.TraderjoeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Traderjoe *TraderjoeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Traderjoe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Traderjoe *TraderjoeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Traderjoe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Traderjoe *TraderjoeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Traderjoe.Contract.contract.Transact(opts, method, params...)
}

// WAVAX is a free data retrieval call binding the contract method 0x73b295c2.
//
// Solidity: function WAVAX() view returns(address)
func (_Traderjoe *TraderjoeCaller) WAVAX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "WAVAX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WAVAX is a free data retrieval call binding the contract method 0x73b295c2.
//
// Solidity: function WAVAX() view returns(address)
func (_Traderjoe *TraderjoeSession) WAVAX() (common.Address, error) {
	return _Traderjoe.Contract.WAVAX(&_Traderjoe.CallOpts)
}

// WAVAX is a free data retrieval call binding the contract method 0x73b295c2.
//
// Solidity: function WAVAX() view returns(address)
func (_Traderjoe *TraderjoeCallerSession) WAVAX() (common.Address, error) {
	return _Traderjoe.Contract.WAVAX(&_Traderjoe.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Traderjoe *TraderjoeCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Traderjoe *TraderjoeSession) Factory() (common.Address, error) {
	return _Traderjoe.Contract.Factory(&_Traderjoe.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Traderjoe *TraderjoeCallerSession) Factory() (common.Address, error) {
	return _Traderjoe.Contract.Factory(&_Traderjoe.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Traderjoe *TraderjoeCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Traderjoe *TraderjoeSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Traderjoe.Contract.GetAmountIn(&_Traderjoe.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Traderjoe *TraderjoeCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Traderjoe.Contract.GetAmountIn(&_Traderjoe.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Traderjoe *TraderjoeCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Traderjoe *TraderjoeSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Traderjoe.Contract.GetAmountOut(&_Traderjoe.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Traderjoe *TraderjoeCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Traderjoe.Contract.GetAmountOut(&_Traderjoe.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Traderjoe *TraderjoeCaller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Traderjoe *TraderjoeSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Traderjoe.Contract.GetAmountsIn(&_Traderjoe.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Traderjoe *TraderjoeCallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Traderjoe.Contract.GetAmountsIn(&_Traderjoe.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Traderjoe *TraderjoeCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Traderjoe *TraderjoeSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Traderjoe.Contract.GetAmountsOut(&_Traderjoe.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Traderjoe *TraderjoeCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Traderjoe.Contract.GetAmountsOut(&_Traderjoe.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Traderjoe *TraderjoeCaller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Traderjoe *TraderjoeSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _Traderjoe.Contract.Quote(&_Traderjoe.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Traderjoe *TraderjoeCallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _Traderjoe.Contract.Quote(&_Traderjoe.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Traderjoe *TraderjoeTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Traderjoe *TraderjoeSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.AddLiquidity(&_Traderjoe.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Traderjoe *TraderjoeTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.AddLiquidity(&_Traderjoe.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xf91b3f72.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (_Traderjoe *TraderjoeTransactor) AddLiquidityAVAX(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "addLiquidityAVAX", token, amountTokenDesired, amountTokenMin, amountAVAXMin, to, deadline)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xf91b3f72.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (_Traderjoe *TraderjoeSession) AddLiquidityAVAX(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.AddLiquidityAVAX(&_Traderjoe.TransactOpts, token, amountTokenDesired, amountTokenMin, amountAVAXMin, to, deadline)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xf91b3f72.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (_Traderjoe *TraderjoeTransactorSession) AddLiquidityAVAX(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.AddLiquidityAVAX(&_Traderjoe.TransactOpts, token, amountTokenDesired, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Traderjoe *TraderjoeTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Traderjoe *TraderjoeSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemoveLiquidity(&_Traderjoe.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Traderjoe *TraderjoeTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemoveLiquidity(&_Traderjoe.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0x33c6b725.
//
// Solidity: function removeLiquidityAVAX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_Traderjoe *TraderjoeTransactor) RemoveLiquidityAVAX(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "removeLiquidityAVAX", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0x33c6b725.
//
// Solidity: function removeLiquidityAVAX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_Traderjoe *TraderjoeSession) RemoveLiquidityAVAX(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemoveLiquidityAVAX(&_Traderjoe.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0x33c6b725.
//
// Solidity: function removeLiquidityAVAX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_Traderjoe *TraderjoeTransactorSession) RemoveLiquidityAVAX(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemoveLiquidityAVAX(&_Traderjoe.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x73bc79cf.
//
// Solidity: function removeLiquidityAVAXSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountAVAX)
func (_Traderjoe *TraderjoeTransactor) RemoveLiquidityAVAXSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "removeLiquidityAVAXSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x73bc79cf.
//
// Solidity: function removeLiquidityAVAXSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountAVAX)
func (_Traderjoe *TraderjoeSession) RemoveLiquidityAVAXSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemoveLiquidityAVAXSupportingFeeOnTransferTokens(&_Traderjoe.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x73bc79cf.
//
// Solidity: function removeLiquidityAVAXSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountAVAX)
func (_Traderjoe *TraderjoeTransactorSession) RemoveLiquidityAVAXSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemoveLiquidityAVAXSupportingFeeOnTransferTokens(&_Traderjoe.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAXWithPermit is a paid mutator transaction binding the contract method 0x2c407024.
//
// Solidity: function removeLiquidityAVAXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountAVAX)
func (_Traderjoe *TraderjoeTransactor) RemoveLiquidityAVAXWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "removeLiquidityAVAXWithPermit", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermit is a paid mutator transaction binding the contract method 0x2c407024.
//
// Solidity: function removeLiquidityAVAXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountAVAX)
func (_Traderjoe *TraderjoeSession) RemoveLiquidityAVAXWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemoveLiquidityAVAXWithPermit(&_Traderjoe.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermit is a paid mutator transaction binding the contract method 0x2c407024.
//
// Solidity: function removeLiquidityAVAXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountAVAX)
func (_Traderjoe *TraderjoeTransactorSession) RemoveLiquidityAVAXWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemoveLiquidityAVAXWithPermit(&_Traderjoe.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9fc27226.
//
// Solidity: function removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountAVAX)
func (_Traderjoe *TraderjoeTransactor) RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9fc27226.
//
// Solidity: function removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountAVAX)
func (_Traderjoe *TraderjoeSession) RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(&_Traderjoe.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9fc27226.
//
// Solidity: function removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountAVAX)
func (_Traderjoe *TraderjoeTransactorSession) RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(&_Traderjoe.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Traderjoe *TraderjoeTransactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Traderjoe *TraderjoeSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemoveLiquidityWithPermit(&_Traderjoe.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Traderjoe *TraderjoeTransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemoveLiquidityWithPermit(&_Traderjoe.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x8a657e67.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Traderjoe *TraderjoeTransactor) SwapAVAXForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "swapAVAXForExactTokens", amountOut, path, to, deadline)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x8a657e67.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Traderjoe *TraderjoeSession) SwapAVAXForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapAVAXForExactTokens(&_Traderjoe.TransactOpts, amountOut, path, to, deadline)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x8a657e67.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Traderjoe *TraderjoeTransactorSession) SwapAVAXForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapAVAXForExactTokens(&_Traderjoe.TransactOpts, amountOut, path, to, deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0xa2a1623d.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Traderjoe *TraderjoeTransactor) SwapExactAVAXForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "swapExactAVAXForTokens", amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0xa2a1623d.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Traderjoe *TraderjoeSession) SwapExactAVAXForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapExactAVAXForTokens(&_Traderjoe.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0xa2a1623d.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Traderjoe *TraderjoeTransactorSession) SwapExactAVAXForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapExactAVAXForTokens(&_Traderjoe.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc57559dd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Traderjoe *TraderjoeTransactor) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "swapExactAVAXForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc57559dd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Traderjoe *TraderjoeSession) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapExactAVAXForTokensSupportingFeeOnTransferTokens(&_Traderjoe.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc57559dd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Traderjoe *TraderjoeTransactorSession) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapExactAVAXForTokensSupportingFeeOnTransferTokens(&_Traderjoe.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0x676528d1.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Traderjoe *TraderjoeTransactor) SwapExactTokensForAVAX(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "swapExactTokensForAVAX", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0x676528d1.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Traderjoe *TraderjoeSession) SwapExactTokensForAVAX(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapExactTokensForAVAX(&_Traderjoe.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0x676528d1.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Traderjoe *TraderjoeTransactorSession) SwapExactTokensForAVAX(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapExactTokensForAVAX(&_Traderjoe.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x762b1562.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Traderjoe *TraderjoeTransactor) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "swapExactTokensForAVAXSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x762b1562.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Traderjoe *TraderjoeSession) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapExactTokensForAVAXSupportingFeeOnTransferTokens(&_Traderjoe.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x762b1562.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Traderjoe *TraderjoeTransactorSession) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapExactTokensForAVAXSupportingFeeOnTransferTokens(&_Traderjoe.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Traderjoe *TraderjoeTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Traderjoe *TraderjoeSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapExactTokensForTokens(&_Traderjoe.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Traderjoe *TraderjoeTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapExactTokensForTokens(&_Traderjoe.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Traderjoe *TraderjoeTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Traderjoe *TraderjoeSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Traderjoe.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Traderjoe *TraderjoeTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Traderjoe.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x7a42416a.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Traderjoe *TraderjoeTransactor) SwapTokensForExactAVAX(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "swapTokensForExactAVAX", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x7a42416a.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Traderjoe *TraderjoeSession) SwapTokensForExactAVAX(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapTokensForExactAVAX(&_Traderjoe.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x7a42416a.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Traderjoe *TraderjoeTransactorSession) SwapTokensForExactAVAX(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapTokensForExactAVAX(&_Traderjoe.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Traderjoe *TraderjoeTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Traderjoe *TraderjoeSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapTokensForExactTokens(&_Traderjoe.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Traderjoe *TraderjoeTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SwapTokensForExactTokens(&_Traderjoe.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Traderjoe *TraderjoeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Traderjoe.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Traderjoe *TraderjoeSession) Receive() (*types.Transaction, error) {
	return _Traderjoe.Contract.Receive(&_Traderjoe.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Traderjoe *TraderjoeTransactorSession) Receive() (*types.Transaction, error) {
	return _Traderjoe.Contract.Receive(&_Traderjoe.TransactOpts)
}

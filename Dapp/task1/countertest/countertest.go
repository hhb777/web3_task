// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package countertest

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

// CountertestMetaData contains all meta data concerning the Countertest contract.
var CountertestMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"CounterAddEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CountAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506101868061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c806306661abd14610038578063e567df5014610056575b5f80fd5b610040610060565b60405161004d91906100c3565b60405180910390f35b61005e610065565b005b5f5481565b5f8081548092919061007690610109565b91905055505f547f925e485aa80d3d9d79e7d63560e9a5f35538f3c3e5e823075ac2c1c8bbadddd060405160405180910390a2565b5f819050919050565b6100bd816100ab565b82525050565b5f6020820190506100d65f8301846100b4565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610113826100ab565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610145576101446100dc565b5b60018201905091905056fea2646970667358221220316efb97d00d6878702b36e544c2db69ba8b45d359696606ab4ecd41f8d4fef664736f6c634300081a0033",
}

// CountertestABI is the input ABI used to generate the binding from.
// Deprecated: Use CountertestMetaData.ABI instead.
var CountertestABI = CountertestMetaData.ABI

// CountertestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CountertestMetaData.Bin instead.
var CountertestBin = CountertestMetaData.Bin

// DeployCountertest deploys a new Ethereum contract, binding an instance of Countertest to it.
func DeployCountertest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Countertest, error) {
	parsed, err := CountertestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CountertestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Countertest{CountertestCaller: CountertestCaller{contract: contract}, CountertestTransactor: CountertestTransactor{contract: contract}, CountertestFilterer: CountertestFilterer{contract: contract}}, nil
}

// Countertest is an auto generated Go binding around an Ethereum contract.
type Countertest struct {
	CountertestCaller     // Read-only binding to the contract
	CountertestTransactor // Write-only binding to the contract
	CountertestFilterer   // Log filterer for contract events
}

// CountertestCaller is an auto generated read-only Go binding around an Ethereum contract.
type CountertestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountertestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CountertestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountertestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CountertestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountertestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CountertestSession struct {
	Contract     *Countertest      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CountertestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CountertestCallerSession struct {
	Contract *CountertestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CountertestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CountertestTransactorSession struct {
	Contract     *CountertestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CountertestRaw is an auto generated low-level Go binding around an Ethereum contract.
type CountertestRaw struct {
	Contract *Countertest // Generic contract binding to access the raw methods on
}

// CountertestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CountertestCallerRaw struct {
	Contract *CountertestCaller // Generic read-only contract binding to access the raw methods on
}

// CountertestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CountertestTransactorRaw struct {
	Contract *CountertestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCountertest creates a new instance of Countertest, bound to a specific deployed contract.
func NewCountertest(address common.Address, backend bind.ContractBackend) (*Countertest, error) {
	contract, err := bindCountertest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Countertest{CountertestCaller: CountertestCaller{contract: contract}, CountertestTransactor: CountertestTransactor{contract: contract}, CountertestFilterer: CountertestFilterer{contract: contract}}, nil
}

// NewCountertestCaller creates a new read-only instance of Countertest, bound to a specific deployed contract.
func NewCountertestCaller(address common.Address, caller bind.ContractCaller) (*CountertestCaller, error) {
	contract, err := bindCountertest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CountertestCaller{contract: contract}, nil
}

// NewCountertestTransactor creates a new write-only instance of Countertest, bound to a specific deployed contract.
func NewCountertestTransactor(address common.Address, transactor bind.ContractTransactor) (*CountertestTransactor, error) {
	contract, err := bindCountertest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CountertestTransactor{contract: contract}, nil
}

// NewCountertestFilterer creates a new log filterer instance of Countertest, bound to a specific deployed contract.
func NewCountertestFilterer(address common.Address, filterer bind.ContractFilterer) (*CountertestFilterer, error) {
	contract, err := bindCountertest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CountertestFilterer{contract: contract}, nil
}

// bindCountertest binds a generic wrapper to an already deployed contract.
func bindCountertest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CountertestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Countertest *CountertestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Countertest.Contract.CountertestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Countertest *CountertestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Countertest.Contract.CountertestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Countertest *CountertestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Countertest.Contract.CountertestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Countertest *CountertestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Countertest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Countertest *CountertestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Countertest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Countertest *CountertestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Countertest.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Countertest *CountertestCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Countertest.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Countertest *CountertestSession) Count() (*big.Int, error) {
	return _Countertest.Contract.Count(&_Countertest.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Countertest *CountertestCallerSession) Count() (*big.Int, error) {
	return _Countertest.Contract.Count(&_Countertest.CallOpts)
}

// CountAdd is a paid mutator transaction binding the contract method 0xe567df50.
//
// Solidity: function CountAdd() returns()
func (_Countertest *CountertestTransactor) CountAdd(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Countertest.contract.Transact(opts, "CountAdd")
}

// CountAdd is a paid mutator transaction binding the contract method 0xe567df50.
//
// Solidity: function CountAdd() returns()
func (_Countertest *CountertestSession) CountAdd() (*types.Transaction, error) {
	return _Countertest.Contract.CountAdd(&_Countertest.TransactOpts)
}

// CountAdd is a paid mutator transaction binding the contract method 0xe567df50.
//
// Solidity: function CountAdd() returns()
func (_Countertest *CountertestTransactorSession) CountAdd() (*types.Transaction, error) {
	return _Countertest.Contract.CountAdd(&_Countertest.TransactOpts)
}

// CountertestCounterAddEventIterator is returned from FilterCounterAddEvent and is used to iterate over the raw logs and unpacked data for CounterAddEvent events raised by the Countertest contract.
type CountertestCounterAddEventIterator struct {
	Event *CountertestCounterAddEvent // Event containing the contract specifics and raw log

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
func (it *CountertestCounterAddEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CountertestCounterAddEvent)
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
		it.Event = new(CountertestCounterAddEvent)
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
func (it *CountertestCounterAddEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CountertestCounterAddEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CountertestCounterAddEvent represents a CounterAddEvent event raised by the Countertest contract.
type CountertestCounterAddEvent struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCounterAddEvent is a free log retrieval operation binding the contract event 0x925e485aa80d3d9d79e7d63560e9a5f35538f3c3e5e823075ac2c1c8bbadddd0.
//
// Solidity: event CounterAddEvent(uint256 indexed count)
func (_Countertest *CountertestFilterer) FilterCounterAddEvent(opts *bind.FilterOpts, count []*big.Int) (*CountertestCounterAddEventIterator, error) {

	var countRule []interface{}
	for _, countItem := range count {
		countRule = append(countRule, countItem)
	}

	logs, sub, err := _Countertest.contract.FilterLogs(opts, "CounterAddEvent", countRule)
	if err != nil {
		return nil, err
	}
	return &CountertestCounterAddEventIterator{contract: _Countertest.contract, event: "CounterAddEvent", logs: logs, sub: sub}, nil
}

// WatchCounterAddEvent is a free log subscription operation binding the contract event 0x925e485aa80d3d9d79e7d63560e9a5f35538f3c3e5e823075ac2c1c8bbadddd0.
//
// Solidity: event CounterAddEvent(uint256 indexed count)
func (_Countertest *CountertestFilterer) WatchCounterAddEvent(opts *bind.WatchOpts, sink chan<- *CountertestCounterAddEvent, count []*big.Int) (event.Subscription, error) {

	var countRule []interface{}
	for _, countItem := range count {
		countRule = append(countRule, countItem)
	}

	logs, sub, err := _Countertest.contract.WatchLogs(opts, "CounterAddEvent", countRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CountertestCounterAddEvent)
				if err := _Countertest.contract.UnpackLog(event, "CounterAddEvent", log); err != nil {
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

// ParseCounterAddEvent is a log parse operation binding the contract event 0x925e485aa80d3d9d79e7d63560e9a5f35538f3c3e5e823075ac2c1c8bbadddd0.
//
// Solidity: event CounterAddEvent(uint256 indexed count)
func (_Countertest *CountertestFilterer) ParseCounterAddEvent(log types.Log) (*CountertestCounterAddEvent, error) {
	event := new(CountertestCounterAddEvent)
	if err := _Countertest.contract.UnpackLog(event, "CounterAddEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

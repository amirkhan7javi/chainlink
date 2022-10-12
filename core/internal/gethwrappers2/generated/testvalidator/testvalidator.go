// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testvalidator

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated"
)

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

var TestValidatorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousRoundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"previousAnswer\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentRoundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"currentAnswer\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialGas\",\"type\":\"uint256\"}],\"name\":\"Validated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"latestRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"minGasUse\",\"type\":\"uint32\"}],\"name\":\"setMinGasUse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"previousRoundId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"previousAnswer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"currentRoundId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"currentAnswer\",\"type\":\"int256\"}],\"name\":\"validate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061020f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806311a8f41314610046578063beed9b5114610064578063c4792df414610084575b600080fd5b61004e610099565b60405161005b919061019a565b60405180910390f35b610077610072366004610133565b61009f565b60405161005b919061018f565b610097610092366004610164565b610117565b005b60015490565b6000805a90507fdb623f4f39d41e75ae1cbe50460c3d1496b6cf9a0db391b7197f82cab2744d2186868686856040516100dc9594939291906101a3565b60405180910390a1600184905560005463ffffffff165b805a6100ff90846101c6565b101561010a576100f3565b5060019695505050505050565b6000805463ffffffff191663ffffffff92909216919091179055565b60008060008060808587031215610148578384fd5b5050823594602084013594506040840135936060013592509050565b600060208284031215610175578081fd5b813563ffffffff81168114610188578182fd5b9392505050565b901515815260200190565b90815260200190565b948552602085019390935260408401919091526060830152608082015260a00190565b6000828210156101fd577f4e487b710000000000000000000000000000000000000000000000000000000081526011600452602481fd5b50039056fea164736f6c6343000800000a",
}

var TestValidatorABI = TestValidatorMetaData.ABI

var TestValidatorBin = TestValidatorMetaData.Bin

func DeployTestValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestValidator, error) {
	parsed, err := TestValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestValidator{TestValidatorCaller: TestValidatorCaller{contract: contract}, TestValidatorTransactor: TestValidatorTransactor{contract: contract}, TestValidatorFilterer: TestValidatorFilterer{contract: contract}}, nil
}

type TestValidator struct {
	address common.Address
	abi     abi.ABI
	TestValidatorCaller
	TestValidatorTransactor
	TestValidatorFilterer
}

type TestValidatorCaller struct {
	contract *bind.BoundContract
}

type TestValidatorTransactor struct {
	contract *bind.BoundContract
}

type TestValidatorFilterer struct {
	contract *bind.BoundContract
}

type TestValidatorSession struct {
	Contract     *TestValidator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TestValidatorCallerSession struct {
	Contract *TestValidatorCaller
	CallOpts bind.CallOpts
}

type TestValidatorTransactorSession struct {
	Contract     *TestValidatorTransactor
	TransactOpts bind.TransactOpts
}

type TestValidatorRaw struct {
	Contract *TestValidator
}

type TestValidatorCallerRaw struct {
	Contract *TestValidatorCaller
}

type TestValidatorTransactorRaw struct {
	Contract *TestValidatorTransactor
}

func NewTestValidator(address common.Address, backend bind.ContractBackend) (*TestValidator, error) {
	abi, err := abi.JSON(strings.NewReader(TestValidatorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindTestValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestValidator{address: address, abi: abi, TestValidatorCaller: TestValidatorCaller{contract: contract}, TestValidatorTransactor: TestValidatorTransactor{contract: contract}, TestValidatorFilterer: TestValidatorFilterer{contract: contract}}, nil
}

func NewTestValidatorCaller(address common.Address, caller bind.ContractCaller) (*TestValidatorCaller, error) {
	contract, err := bindTestValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestValidatorCaller{contract: contract}, nil
}

func NewTestValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*TestValidatorTransactor, error) {
	contract, err := bindTestValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestValidatorTransactor{contract: contract}, nil
}

func NewTestValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*TestValidatorFilterer, error) {
	contract, err := bindTestValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestValidatorFilterer{contract: contract}, nil
}

func bindTestValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_TestValidator *TestValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestValidator.Contract.TestValidatorCaller.contract.Call(opts, result, method, params...)
}

func (_TestValidator *TestValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestValidator.Contract.TestValidatorTransactor.contract.Transfer(opts)
}

func (_TestValidator *TestValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestValidator.Contract.TestValidatorTransactor.contract.Transact(opts, method, params...)
}

func (_TestValidator *TestValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestValidator.Contract.contract.Call(opts, result, method, params...)
}

func (_TestValidator *TestValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestValidator.Contract.contract.Transfer(opts)
}

func (_TestValidator *TestValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestValidator.Contract.contract.Transact(opts, method, params...)
}

func (_TestValidator *TestValidatorCaller) LatestRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestValidator.contract.Call(opts, &out, "latestRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestValidator *TestValidatorSession) LatestRoundId() (*big.Int, error) {
	return _TestValidator.Contract.LatestRoundId(&_TestValidator.CallOpts)
}

func (_TestValidator *TestValidatorCallerSession) LatestRoundId() (*big.Int, error) {
	return _TestValidator.Contract.LatestRoundId(&_TestValidator.CallOpts)
}

func (_TestValidator *TestValidatorTransactor) SetMinGasUse(opts *bind.TransactOpts, minGasUse uint32) (*types.Transaction, error) {
	return _TestValidator.contract.Transact(opts, "setMinGasUse", minGasUse)
}

func (_TestValidator *TestValidatorSession) SetMinGasUse(minGasUse uint32) (*types.Transaction, error) {
	return _TestValidator.Contract.SetMinGasUse(&_TestValidator.TransactOpts, minGasUse)
}

func (_TestValidator *TestValidatorTransactorSession) SetMinGasUse(minGasUse uint32) (*types.Transaction, error) {
	return _TestValidator.Contract.SetMinGasUse(&_TestValidator.TransactOpts, minGasUse)
}

func (_TestValidator *TestValidatorTransactor) Validate(opts *bind.TransactOpts, previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error) {
	return _TestValidator.contract.Transact(opts, "validate", previousRoundId, previousAnswer, currentRoundId, currentAnswer)
}

func (_TestValidator *TestValidatorSession) Validate(previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error) {
	return _TestValidator.Contract.Validate(&_TestValidator.TransactOpts, previousRoundId, previousAnswer, currentRoundId, currentAnswer)
}

func (_TestValidator *TestValidatorTransactorSession) Validate(previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error) {
	return _TestValidator.Contract.Validate(&_TestValidator.TransactOpts, previousRoundId, previousAnswer, currentRoundId, currentAnswer)
}

type TestValidatorValidatedIterator struct {
	Event *TestValidatorValidated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestValidatorValidatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestValidatorValidated)
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

	select {
	case log := <-it.logs:
		it.Event = new(TestValidatorValidated)
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

func (it *TestValidatorValidatedIterator) Error() error {
	return it.fail
}

func (it *TestValidatorValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestValidatorValidated struct {
	PreviousRoundId *big.Int
	PreviousAnswer  *big.Int
	CurrentRoundId  *big.Int
	CurrentAnswer   *big.Int
	InitialGas      *big.Int
	Raw             types.Log
}

func (_TestValidator *TestValidatorFilterer) FilterValidated(opts *bind.FilterOpts) (*TestValidatorValidatedIterator, error) {

	logs, sub, err := _TestValidator.contract.FilterLogs(opts, "Validated")
	if err != nil {
		return nil, err
	}
	return &TestValidatorValidatedIterator{contract: _TestValidator.contract, event: "Validated", logs: logs, sub: sub}, nil
}

func (_TestValidator *TestValidatorFilterer) WatchValidated(opts *bind.WatchOpts, sink chan<- *TestValidatorValidated) (event.Subscription, error) {

	logs, sub, err := _TestValidator.contract.WatchLogs(opts, "Validated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestValidatorValidated)
				if err := _TestValidator.contract.UnpackLog(event, "Validated", log); err != nil {
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

func (_TestValidator *TestValidatorFilterer) ParseValidated(log types.Log) (*TestValidatorValidated, error) {
	event := new(TestValidatorValidated)
	if err := _TestValidator.contract.UnpackLog(event, "Validated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_TestValidator *TestValidator) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _TestValidator.abi.Events["Validated"].ID:
		return _TestValidator.ParseValidated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (TestValidatorValidated) Topic() common.Hash {
	return common.HexToHash("0xdb623f4f39d41e75ae1cbe50460c3d1496b6cf9a0db391b7197f82cab2744d21")
}

func (_TestValidator *TestValidator) Address() common.Address {
	return _TestValidator.address
}

type TestValidatorInterface interface {
	LatestRoundId(opts *bind.CallOpts) (*big.Int, error)

	SetMinGasUse(opts *bind.TransactOpts, minGasUse uint32) (*types.Transaction, error)

	Validate(opts *bind.TransactOpts, previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error)

	FilterValidated(opts *bind.FilterOpts) (*TestValidatorValidatedIterator, error)

	WatchValidated(opts *bind.WatchOpts, sink chan<- *TestValidatorValidated) (event.Subscription, error)

	ParseValidated(log types.Log) (*TestValidatorValidated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accesscontroltesthelper

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

var AccessControlTestHelperMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[],\"name\":\"Dummy\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_roundID\",\"type\":\"uint256\"}],\"name\":\"readGetAnswer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"_roundID\",\"type\":\"uint80\"}],\"name\":\"readGetRoundData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_roundID\",\"type\":\"uint256\"}],\"name\":\"readGetTimestamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"readLatestAnswer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"readLatestRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"readLatestRoundData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"readLatestTimestamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"testLatestTransmissionDetails\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506106f1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063c0c9c7db1161005b578063c0c9c7db146100db578063c9592ab9146100ee578063d2f79c4714610101578063eea2913a1461011457610088565b806304cefda51461008d57806320f2c97c146100a257806395319deb146100b5578063bf5fc18b146100c8575b600080fd5b6100a061009b36600461053f565b610127565b005b6100a06100b036600461053f565b6101a0565b6100a06100c3366004610589565b610243565b6100a06100d6366004610560565b61030b565b6100a06100e936600461053f565b6103ce565b6100a06100fc366004610560565b61046c565b6100a061010f36600461053f565b6104b1565b6100a061012236600461053f565b6104ea565b806001600160a01b031663e5fe45776040518163ffffffff1660e01b815260040160a06040518083038186803b15801561016057600080fd5b505afa158015610174573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061019891906105bf565b505050505050565b806001600160a01b031663feaf968c6040518163ffffffff1660e01b815260040160a06040518083038186803b1580156101d957600080fd5b505afa1580156101ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102119190610660565b50506040517f10e4ab9f2ce395bf5539d7c60c9bfeef0b416602954734c5bb8bfd9433c9ff689350600092509050a150565b6040517f9a6fc8f50000000000000000000000000000000000000000000000000000000081526001600160a01b03831690639a6fc8f5906102889084906004016106b2565b60a06040518083038186803b1580156102a057600080fd5b505afa1580156102b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102d89190610660565b50506040517f10e4ab9f2ce395bf5539d7c60c9bfeef0b416602954734c5bb8bfd9433c9ff689350600092509050a15050565b6040517fb5ab58dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063b5ab58dc906103509084906004016106a9565b60206040518083038186803b15801561036857600080fd5b505afa15801561037c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103a09190610648565b506040517f10e4ab9f2ce395bf5539d7c60c9bfeef0b416602954734c5bb8bfd9433c9ff6890600090a15050565b806001600160a01b03166350d25bcd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561040757600080fd5b505afa15801561041b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043f9190610648565b506040517f10e4ab9f2ce395bf5539d7c60c9bfeef0b416602954734c5bb8bfd9433c9ff6890600090a150565b6040517fb633620c0000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063b633620c906103509084906004016106a9565b806001600160a01b0316638205bf6a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561040757600080fd5b806001600160a01b031663668a0f026040518163ffffffff1660e01b815260040160206040518083038186803b15801561040757600080fd5b80356001600160a01b038116811461053a57600080fd5b919050565b600060208284031215610550578081fd5b61055982610523565b9392505050565b60008060408385031215610572578081fd5b61057b83610523565b946020939093013593505050565b6000806040838503121561059b578182fd5b6105a483610523565b915060208301356105b4816106c9565b809150509250929050565b600080600080600060a086880312156105d6578081fd5b85519450602086015163ffffffff811681146105f0578182fd5b604087015190945060ff81168114610606578182fd5b8093505060608601518060170b811461061d578182fd5b608087015190925067ffffffffffffffff8116811461063a578182fd5b809150509295509295909350565b600060208284031215610659578081fd5b5051919050565b600080600080600060a08688031215610677578081fd5b8551610682816106c9565b80955050602086015193506040860151925060608601519150608086015161063a816106c9565b90815260200190565b69ffffffffffffffffffff91909116815260200190565b69ffffffffffffffffffff811681146106e157600080fd5b5056fea164736f6c6343000800000a",
}

var AccessControlTestHelperABI = AccessControlTestHelperMetaData.ABI

var AccessControlTestHelperBin = AccessControlTestHelperMetaData.Bin

func DeployAccessControlTestHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccessControlTestHelper, error) {
	parsed, err := AccessControlTestHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccessControlTestHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessControlTestHelper{AccessControlTestHelperCaller: AccessControlTestHelperCaller{contract: contract}, AccessControlTestHelperTransactor: AccessControlTestHelperTransactor{contract: contract}, AccessControlTestHelperFilterer: AccessControlTestHelperFilterer{contract: contract}}, nil
}

type AccessControlTestHelper struct {
	address common.Address
	abi     abi.ABI
	AccessControlTestHelperCaller
	AccessControlTestHelperTransactor
	AccessControlTestHelperFilterer
}

type AccessControlTestHelperCaller struct {
	contract *bind.BoundContract
}

type AccessControlTestHelperTransactor struct {
	contract *bind.BoundContract
}

type AccessControlTestHelperFilterer struct {
	contract *bind.BoundContract
}

type AccessControlTestHelperSession struct {
	Contract     *AccessControlTestHelper
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type AccessControlTestHelperCallerSession struct {
	Contract *AccessControlTestHelperCaller
	CallOpts bind.CallOpts
}

type AccessControlTestHelperTransactorSession struct {
	Contract     *AccessControlTestHelperTransactor
	TransactOpts bind.TransactOpts
}

type AccessControlTestHelperRaw struct {
	Contract *AccessControlTestHelper
}

type AccessControlTestHelperCallerRaw struct {
	Contract *AccessControlTestHelperCaller
}

type AccessControlTestHelperTransactorRaw struct {
	Contract *AccessControlTestHelperTransactor
}

func NewAccessControlTestHelper(address common.Address, backend bind.ContractBackend) (*AccessControlTestHelper, error) {
	abi, err := abi.JSON(strings.NewReader(AccessControlTestHelperABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindAccessControlTestHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlTestHelper{address: address, abi: abi, AccessControlTestHelperCaller: AccessControlTestHelperCaller{contract: contract}, AccessControlTestHelperTransactor: AccessControlTestHelperTransactor{contract: contract}, AccessControlTestHelperFilterer: AccessControlTestHelperFilterer{contract: contract}}, nil
}

func NewAccessControlTestHelperCaller(address common.Address, caller bind.ContractCaller) (*AccessControlTestHelperCaller, error) {
	contract, err := bindAccessControlTestHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTestHelperCaller{contract: contract}, nil
}

func NewAccessControlTestHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlTestHelperTransactor, error) {
	contract, err := bindAccessControlTestHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTestHelperTransactor{contract: contract}, nil
}

func NewAccessControlTestHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlTestHelperFilterer, error) {
	contract, err := bindAccessControlTestHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlTestHelperFilterer{contract: contract}, nil
}

func bindAccessControlTestHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlTestHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_AccessControlTestHelper *AccessControlTestHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlTestHelper.Contract.AccessControlTestHelperCaller.contract.Call(opts, result, method, params...)
}

func (_AccessControlTestHelper *AccessControlTestHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.AccessControlTestHelperTransactor.contract.Transfer(opts)
}

func (_AccessControlTestHelper *AccessControlTestHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.AccessControlTestHelperTransactor.contract.Transact(opts, method, params...)
}

func (_AccessControlTestHelper *AccessControlTestHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlTestHelper.Contract.contract.Call(opts, result, method, params...)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.contract.Transfer(opts)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.contract.Transact(opts, method, params...)
}

func (_AccessControlTestHelper *AccessControlTestHelperCaller) TestLatestTransmissionDetails(opts *bind.CallOpts, _aggregator common.Address) error {
	var out []interface{}
	err := _AccessControlTestHelper.contract.Call(opts, &out, "testLatestTransmissionDetails", _aggregator)

	if err != nil {
		return err
	}

	return err

}

func (_AccessControlTestHelper *AccessControlTestHelperSession) TestLatestTransmissionDetails(_aggregator common.Address) error {
	return _AccessControlTestHelper.Contract.TestLatestTransmissionDetails(&_AccessControlTestHelper.CallOpts, _aggregator)
}

func (_AccessControlTestHelper *AccessControlTestHelperCallerSession) TestLatestTransmissionDetails(_aggregator common.Address) error {
	return _AccessControlTestHelper.Contract.TestLatestTransmissionDetails(&_AccessControlTestHelper.CallOpts, _aggregator)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactor) ReadGetAnswer(opts *bind.TransactOpts, _aggregator common.Address, _roundID *big.Int) (*types.Transaction, error) {
	return _AccessControlTestHelper.contract.Transact(opts, "readGetAnswer", _aggregator, _roundID)
}

func (_AccessControlTestHelper *AccessControlTestHelperSession) ReadGetAnswer(_aggregator common.Address, _roundID *big.Int) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.ReadGetAnswer(&_AccessControlTestHelper.TransactOpts, _aggregator, _roundID)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactorSession) ReadGetAnswer(_aggregator common.Address, _roundID *big.Int) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.ReadGetAnswer(&_AccessControlTestHelper.TransactOpts, _aggregator, _roundID)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactor) ReadGetRoundData(opts *bind.TransactOpts, _aggregator common.Address, _roundID *big.Int) (*types.Transaction, error) {
	return _AccessControlTestHelper.contract.Transact(opts, "readGetRoundData", _aggregator, _roundID)
}

func (_AccessControlTestHelper *AccessControlTestHelperSession) ReadGetRoundData(_aggregator common.Address, _roundID *big.Int) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.ReadGetRoundData(&_AccessControlTestHelper.TransactOpts, _aggregator, _roundID)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactorSession) ReadGetRoundData(_aggregator common.Address, _roundID *big.Int) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.ReadGetRoundData(&_AccessControlTestHelper.TransactOpts, _aggregator, _roundID)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactor) ReadGetTimestamp(opts *bind.TransactOpts, _aggregator common.Address, _roundID *big.Int) (*types.Transaction, error) {
	return _AccessControlTestHelper.contract.Transact(opts, "readGetTimestamp", _aggregator, _roundID)
}

func (_AccessControlTestHelper *AccessControlTestHelperSession) ReadGetTimestamp(_aggregator common.Address, _roundID *big.Int) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.ReadGetTimestamp(&_AccessControlTestHelper.TransactOpts, _aggregator, _roundID)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactorSession) ReadGetTimestamp(_aggregator common.Address, _roundID *big.Int) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.ReadGetTimestamp(&_AccessControlTestHelper.TransactOpts, _aggregator, _roundID)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactor) ReadLatestAnswer(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _AccessControlTestHelper.contract.Transact(opts, "readLatestAnswer", _aggregator)
}

func (_AccessControlTestHelper *AccessControlTestHelperSession) ReadLatestAnswer(_aggregator common.Address) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.ReadLatestAnswer(&_AccessControlTestHelper.TransactOpts, _aggregator)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactorSession) ReadLatestAnswer(_aggregator common.Address) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.ReadLatestAnswer(&_AccessControlTestHelper.TransactOpts, _aggregator)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactor) ReadLatestRound(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _AccessControlTestHelper.contract.Transact(opts, "readLatestRound", _aggregator)
}

func (_AccessControlTestHelper *AccessControlTestHelperSession) ReadLatestRound(_aggregator common.Address) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.ReadLatestRound(&_AccessControlTestHelper.TransactOpts, _aggregator)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactorSession) ReadLatestRound(_aggregator common.Address) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.ReadLatestRound(&_AccessControlTestHelper.TransactOpts, _aggregator)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactor) ReadLatestRoundData(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _AccessControlTestHelper.contract.Transact(opts, "readLatestRoundData", _aggregator)
}

func (_AccessControlTestHelper *AccessControlTestHelperSession) ReadLatestRoundData(_aggregator common.Address) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.ReadLatestRoundData(&_AccessControlTestHelper.TransactOpts, _aggregator)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactorSession) ReadLatestRoundData(_aggregator common.Address) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.ReadLatestRoundData(&_AccessControlTestHelper.TransactOpts, _aggregator)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactor) ReadLatestTimestamp(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _AccessControlTestHelper.contract.Transact(opts, "readLatestTimestamp", _aggregator)
}

func (_AccessControlTestHelper *AccessControlTestHelperSession) ReadLatestTimestamp(_aggregator common.Address) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.ReadLatestTimestamp(&_AccessControlTestHelper.TransactOpts, _aggregator)
}

func (_AccessControlTestHelper *AccessControlTestHelperTransactorSession) ReadLatestTimestamp(_aggregator common.Address) (*types.Transaction, error) {
	return _AccessControlTestHelper.Contract.ReadLatestTimestamp(&_AccessControlTestHelper.TransactOpts, _aggregator)
}

type AccessControlTestHelperDummyIterator struct {
	Event *AccessControlTestHelperDummy

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlTestHelperDummyIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlTestHelperDummy)
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
		it.Event = new(AccessControlTestHelperDummy)
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

func (it *AccessControlTestHelperDummyIterator) Error() error {
	return it.fail
}

func (it *AccessControlTestHelperDummyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlTestHelperDummy struct {
	Raw types.Log
}

func (_AccessControlTestHelper *AccessControlTestHelperFilterer) FilterDummy(opts *bind.FilterOpts) (*AccessControlTestHelperDummyIterator, error) {

	logs, sub, err := _AccessControlTestHelper.contract.FilterLogs(opts, "Dummy")
	if err != nil {
		return nil, err
	}
	return &AccessControlTestHelperDummyIterator{contract: _AccessControlTestHelper.contract, event: "Dummy", logs: logs, sub: sub}, nil
}

func (_AccessControlTestHelper *AccessControlTestHelperFilterer) WatchDummy(opts *bind.WatchOpts, sink chan<- *AccessControlTestHelperDummy) (event.Subscription, error) {

	logs, sub, err := _AccessControlTestHelper.contract.WatchLogs(opts, "Dummy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlTestHelperDummy)
				if err := _AccessControlTestHelper.contract.UnpackLog(event, "Dummy", log); err != nil {
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

func (_AccessControlTestHelper *AccessControlTestHelperFilterer) ParseDummy(log types.Log) (*AccessControlTestHelperDummy, error) {
	event := new(AccessControlTestHelperDummy)
	if err := _AccessControlTestHelper.contract.UnpackLog(event, "Dummy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_AccessControlTestHelper *AccessControlTestHelper) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _AccessControlTestHelper.abi.Events["Dummy"].ID:
		return _AccessControlTestHelper.ParseDummy(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (AccessControlTestHelperDummy) Topic() common.Hash {
	return common.HexToHash("0x10e4ab9f2ce395bf5539d7c60c9bfeef0b416602954734c5bb8bfd9433c9ff68")
}

func (_AccessControlTestHelper *AccessControlTestHelper) Address() common.Address {
	return _AccessControlTestHelper.address
}

type AccessControlTestHelperInterface interface {
	TestLatestTransmissionDetails(opts *bind.CallOpts, _aggregator common.Address) error

	ReadGetAnswer(opts *bind.TransactOpts, _aggregator common.Address, _roundID *big.Int) (*types.Transaction, error)

	ReadGetRoundData(opts *bind.TransactOpts, _aggregator common.Address, _roundID *big.Int) (*types.Transaction, error)

	ReadGetTimestamp(opts *bind.TransactOpts, _aggregator common.Address, _roundID *big.Int) (*types.Transaction, error)

	ReadLatestAnswer(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error)

	ReadLatestRound(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error)

	ReadLatestRoundData(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error)

	ReadLatestTimestamp(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error)

	FilterDummy(opts *bind.FilterOpts) (*AccessControlTestHelperDummyIterator, error)

	WatchDummy(opts *bind.WatchOpts, sink chan<- *AccessControlTestHelperDummy) (event.Subscription, error)

	ParseDummy(log types.Log) (*AccessControlTestHelperDummy, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}

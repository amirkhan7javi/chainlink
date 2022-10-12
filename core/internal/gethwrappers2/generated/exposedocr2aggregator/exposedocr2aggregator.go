// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exposedocr2aggregator

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

var ExposedOCR2AggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"BillingAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"BillingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"oldLinkToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"newLinkToken\",\"type\":\"address\"}],\"name\":\"LinkTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int192\",\"name\":\"answer\",\"type\":\"int192\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"observationsTimestamp\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int192[]\",\"name\":\"observations\",\"type\":\"int192[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"observers\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"int192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"int192\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"RequesterAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"}],\"name\":\"RoundRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"previousValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousGasLimit\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"currentValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"currentGasLimit\",\"type\":\"uint32\"}],\"name\":\"ValidatorConfigSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_configCount\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_encodedConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_encodedConfig\",\"type\":\"bytes\"}],\"name\":\"exposedConfigDigestFromConfigData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBillingAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinkToken\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequesterAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId_\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorConfig\",\"outputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTransmissionDetails\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"},{\"internalType\":\"int192\",\"name\":\"latestAnswer_\",\"type\":\"int192\"},{\"internalType\":\"uint64\",\"name\":\"latestTimestamp_\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAnswer\",\"outputs\":[{\"internalType\":\"int192\",\"name\":\"\",\"type\":\"int192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAnswer\",\"outputs\":[{\"internalType\":\"int192\",\"name\":\"\",\"type\":\"int192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"oracleObservationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"owedPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestNewRound\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"setBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"}],\"name\":\"setBillingAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"setLinkToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"requesterAccessController\",\"type\":\"address\"}],\"name\":\"setRequesterAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"newValidator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newGasLimit\",\"type\":\"uint32\"}],\"name\":\"setValidatorConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506000806000806000806040518060200160405280600081525033806000806001600160a01b0316826001600160a01b031614156200006d5760405162461bcd60e51b81526004016200006490620004f7565b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000a057620000a08162000176565b5050601180546001600160a01b0319166001600160a01b038a169081179091556040519091506000907f4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a908290a3620000f984620001f3565b7fff0000000000000000000000000000000000000000000000000000000000000060f883901b1660c05280516200013890601090602084019062000400565b50620001448362000267565b62000151600080620002d8565b50505050601791820b820b604090811b60805290820b90910b901b60a05250620005b9565b6001600160a01b038116331415620001a25760405162461bcd60e51b815260040162000064906200052e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6012546001600160a01b0390811690821681146200026357601280546001600160a01b0319166001600160a01b0384161790556040517f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912906200025a9083908590620004a6565b60405180910390a15b5050565b62000271620003d1565b600f546001600160a01b0390811690821681146200026357600f80546001600160a01b0319166001600160a01b0384161790556040517f27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634906200025a9083908590620004a6565b620002e2620003d1565b60408051808201909152600e546001600160a01b03808216808452600160a01b90920463ffffffff16602084015284161415806200033057508163ffffffff16816020015163ffffffff1614155b15620003cc576040805180820182526001600160a01b0385811680835263ffffffff86166020938401819052600e80546001600160a01b031916831763ffffffff60a01b1916600160a01b9092029190911790558451928501519351909392909116917fb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca154191620003c39190879062000565565b60405180910390a35b505050565b6000546001600160a01b03163314620003fe5760405162461bcd60e51b81526004016200006490620004c0565b565b8280546200040e906200057c565b90600052602060002090601f0160209004810192826200043257600085556200047d565b82601f106200044d57805160ff19168380011785556200047d565b828001600101855582156200047d579182015b828111156200047d57825182559160200191906001019062000460565b506200048b9291506200048f565b5090565b5b808211156200048b576000815560010162000490565b6001600160a01b0392831681529116602082015260400190565b60208082526016908201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604082015260600190565b60208082526018908201527f43616e6e6f7420736574206f776e657220746f207a65726f0000000000000000604082015260600190565b60208082526017908201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604082015260600190565b63ffffffff92831681529116602082015260400190565b6002810460018216806200059157607f821691505b60208210811415620005b357634e487b7160e01b600052602260045260246000fd5b50919050565b60805160401c60a05160401c60c05160f81c6154f56200060e60003960006107f0015260008181610d2b01528181611dd6015261316101526000818161074101528181611db5015261313401526154f56000f3fe608060405234801561001057600080fd5b50600436106102de5760003560e01c80639bd2c0b111610186578063d09dc339116100e3578063e76d516811610097578063f2fde38b11610071578063f2fde38b146105a8578063fbffd2c1146105bb578063feaf968c146105ce576102de565b8063e76d51681461057a578063eb45716314610582578063eb5dcd6c14610595576102de565b8063e3d0e712116100c8578063e3d0e7121461052e578063e4902f8214610541578063e5fe457714610561576102de565b8063d09dc3391461051e578063daffc4b514610526576102de565b8063b1dc65a41161013a578063b633620c1161011f578063b633620c146104f0578063c107532914610503578063c4c92b3714610516576102de565b8063b1dc65a4146104ca578063b5ab58dc146104dd576102de565b80639e3ceeab1161016b5780639e3ceeab1461048d578063afcb95d7146104a0578063b121e147146104b7576102de565b80639bd2c0b1146104645780639c849b301461047a576102de565b8063668a0f021161023f57806381ff7048116101f35780638da5cb5b116101cd5780638da5cb5b1461041657806398e5b12a1461042b5780639a6fc8f514610440576102de565b806381ff7048146103e45780638205bf6a146103fb5780638ac28d5a14610403576102de565b80637284e416116102245780637284e416146103c1578063739376c1146103c957806379ba5097146103dc576102de565b8063668a0f02146103b157806370da2f67146103b9576102de565b80634fb174701161029657806354fd4d501161027b57806354fd4d5014610381578063643dc10514610389578063666cab8d1461039c576102de565b80634fb174701461036457806350d25bcd14610379576102de565b806322adbc78116102c757806322adbc78146103215780632993726814610336578063313ce5671461034f576102de565b80630eafb25b146102e3578063181f5a771461030c575b600080fd5b6102f66102f1366004613f95565b6105d6565b60405161030391906145ab565b60405180910390f35b610314610708565b604051610303919061483d565b61032961073f565b6040516103039190614790565b61033e610763565b604051610303959493929190615188565b6103576107ee565b6040516103039190615206565b610377610372366004614244565b610812565b005b6102f6610a9b565b6102f6610ac7565b61037761039736600461444c565b610acc565b6103a4610cb1565b6040516103039190614696565b6102f6610d13565b610329610d29565b610314610d4d565b6102f66103d7366004614286565b610dd6565b610377610e30565b6103ec610ec9565b60405161030393929190615167565b6102f6610ee5565b610377610411366004613f95565b610f2d565b61041e610f72565b604051610303919061460c565b610433610f81565b60405161030391906151bc565b61045361044e3660046144c3565b6110e0565b6040516103039594939291906151d3565b61046c6111a8565b604051610303929190614771565b610377610488366004614014565b6111ec565b61037761049b366004613f95565b61139e565b6104a861142e565b604051610303939291906146a9565b6103776104c5366004613f95565b611449565b6103776104d8366004614145565b61150d565b6102f66104eb366004614256565b61192c565b6102f66104fe366004614256565b611962565b610377610511366004613fe9565b6119b4565b61041e611c11565b6102f6611c20565b61041e611cdd565b61037761053c36600461407d565b611cec565b61055461054f366004613f95565b61248d565b60405161030391906150b9565b61056961255a565b604051610303959493929190614700565b61041e6125ef565b610377610590366004614217565b6125fe565b6103776105a3366004613fb1565b612743565b6103776105b6366004613f95565b612834565b6103776105c9366004613f95565b612845565b610453612856565b6001600160a01b03811660009081526002602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046bffffffffffffffffffffffff16918101919091529061063e576000915050610703565b600b5460208201516000917201000000000000000000000000000000000000900463ffffffff169060069060ff16601f811061068a57634e487b7160e01b600052603260045260246000fd5b600881049190910154600b546106c0926007166004026101000a90910463ffffffff9081169166010000000000009004166153d8565b63ffffffff166106d091906152e7565b6106de90633b9aca006152e7565b905081604001516bffffffffffffffffffffffff16816106fe9190615262565b925050505b919050565b60408051808201909152601a81527f4f43523241676772656761746f7220312e302e302d616c706861000000000000602082015290565b7f000000000000000000000000000000000000000000000000000000000000000081565b600b546a0100000000000000000000810463ffffffff908116926e010000000000000000000000000000830482169272010000000000000000000000000000000000008104831692760100000000000000000000000000000000000000000000820416917a01000000000000000000000000000000000000000000000000000090910462ffffff1690565b7f000000000000000000000000000000000000000000000000000000000000000081565b61081a6128ef565b6011546001600160a01b039081169083168114156108385750610a97565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b038416906370a082319061087d90309060040161460c565b60206040518083038186803b15801561089557600080fd5b505afa1580156108a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108cd919061426e565b506108d661291b565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526000906001600160a01b038316906370a082319061091e90309060040161460c565b60206040518083038186803b15801561093657600080fd5b505afa15801561094a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096e919061426e565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000081529091506001600160a01b0383169063a9059cbb906109b8908690859060040161467d565b602060405180830381600087803b1580156109d257600080fd5b505af11580156109e6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a0a91906141f7565b610a2f5760405162461bcd60e51b8152600401610a2690614aab565b60405180910390fd5b601180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0386811691821790925560405190918416907f4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a90600090a350505b5050565b600b546601000000000000900463ffffffff166000908152600c6020526040902054601790810b900b90565b600681565b6012546001600160a01b0316610ae0610f72565b6001600160a01b0316336001600160a01b03161480610b9457506040517f6b14daf80000000000000000000000000000000000000000000000000000000081526001600160a01b03821690636b14daf890610b449033906000903690600401614620565b60206040518083038186803b158015610b5c57600080fd5b505afa158015610b70573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9491906141f7565b610bb05760405162461bcd60e51b8152600401610a2690614d08565b610bb861291b565b85600b600001600a6101000a81548163ffffffff021916908363ffffffff16021790555084600b600001600e6101000a81548163ffffffff021916908363ffffffff16021790555083600b60000160126101000a81548163ffffffff021916908363ffffffff16021790555082600b60000160166101000a81548163ffffffff021916908363ffffffff16021790555081600b600001601a6101000a81548162ffffff021916908362ffffff1602179055507f0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f8686868686604051610ca1959493929190615188565b60405180910390a1505050505050565b60606005805480602002602001604051908101604052809291908181526020018280548015610d0957602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610ceb575b5050505050905090565b600b546601000000000000900463ffffffff1690565b7f000000000000000000000000000000000000000000000000000000000000000081565b606060108054610d5c906153fd565b80601f0160208091040260200160405190810160405280929190818152602001828054610d88906153fd565b8015610d095780601f10610daa57610100808354040283529160200191610d09565b820191906000526020600020905b815481529060010190602001808311610db857509395945050505050565b6000610e218b8b8b8b8b8b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d92508c9150612d2e9050565b9b9a5050505050505050505050565b6001546001600160a01b03163314610e5a5760405162461bcd60e51b8152600401610a2690614887565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b600d54600a5463ffffffff808316926401000000009004169192565b600b5463ffffffff660100000000000090910481166000908152600c60205260409020547c010000000000000000000000000000000000000000000000000000000090041690565b6001600160a01b03818116600090815260136020526040902054163314610f665760405162461bcd60e51b8152600401610a2690614dab565b610f6f81612dda565b50565b6000546001600160a01b031690565b6000610f8b610f72565b6001600160a01b0316336001600160a01b031614806110435750600f546040517f6b14daf80000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636b14daf890610ff39033906000903690600401614620565b60206040518083038186803b15801561100b57600080fd5b505afa15801561101f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061104391906141f7565b61105f5760405162461bcd60e51b8152600401610a2690614998565b600b54600a54604051610100830464ffffffffff8116936601000000000000900463ffffffff9081169333937f41e3990591fd372502daa15842da15bc7f41c75309ab3ff4f56f1848c178825c936110c09360089190911c169087906146e1565b60405180910390a26110d381600161527a565b63ffffffff169250505090565b60008080808063ffffffff69ffffffffffffffffffff871611156111125750600093508392508291508190508061119f565b50505063ffffffff8084166000908152600c602090815260409182902082516060810184529054601781810b810b810b8084527801000000000000000000000000000000000000000000000000830487169484018590527c0100000000000000000000000000000000000000000000000000000000909204909516919093018190528695509190920b9250835b91939590929450565b60408051808201909152600e546001600160a01b0381168083527401000000000000000000000000000000000000000090910463ffffffff16602090920182905291565b6111f46128ef565b8281146112135760405162461bcd60e51b8152600401610a2690614ef5565b60005b8381101561139757600085858381811061124057634e487b7160e01b600052603260045260246000fd5b90506020020160208101906112559190613f95565b9050600084848481811061127957634e487b7160e01b600052603260045260246000fd5b905060200201602081019061128e9190613f95565b6001600160a01b0380841660009081526013602052604090205491925016801580806112cb5750826001600160a01b0316826001600160a01b0316145b6112e75760405162461bcd60e51b8152600401610a2690614cd1565b6001600160a01b03848116600090815260136020526040902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000168583169081179091559083161461138057826001600160a01b0316826001600160a01b0316856001600160a01b03167f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b360405160405180910390a45b50505050808061138f90615438565b915050611216565b5050505050565b6113a66128ef565b600f546001600160a01b039081169082168114610a9757600f80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384161790556040517f27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634906114229083908590614757565b60405180910390a15050565b600a54600b5460009261010090910460081c63ffffffff1690565b6001600160a01b038181166000908152601460205260409020541633146114825760405162461bcd60e51b8152600401610a2690614850565b6001600160a01b0381811660008181526013602090815260408083208054337fffffffffffffffffffffffff000000000000000000000000000000000000000080831682179093556014909452828520805490921690915590519416939092849290917f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b39190a45050565b60005a604080516101008082018352600b5460ff8116835290810464ffffffffff90811660208085018290526601000000000000840463ffffffff908116968601969096526a01000000000000000000008404861660608601526e01000000000000000000000000000084048616608086015272010000000000000000000000000000000000008404861660a0860152760100000000000000000000000000000000000000000000840490951660c08501527a01000000000000000000000000000000000000000000000000000090920462ffffff1660e08401529394509092918c0135918216116116115760405162461bcd60e51b8152600401610a2690614961565b3360009081526002602052604090205460ff166116405760405162461bcd60e51b8152600401610a2690614d3d565b600a548b35146116625760405162461bcd60e51b8152600401610a2690614e87565b6116708a8a8a8a8a8a613019565b815161167d9060016152a2565b60ff16871461169e5760405162461bcd60e51b8152600401610a2690614d74565b8685146116bd5760405162461bcd60e51b8152600401610a2690614fcf565b60008a8a6040516116cf9291906145b4565b6040519081900381206116e6918e906020016146c7565b6040516020818303038152906040528051906020012090506000611708613c39565b60005b8a8110156118805760006001858a846020811061173857634e487b7160e01b600052603260045260246000fd5b61174591901a601b6152a2565b8f8f8681811061176557634e487b7160e01b600052603260045260246000fd5b905060200201358e8e8781811061178c57634e487b7160e01b600052603260045260246000fd5b90506020020135604051600081526020016040526040516117b09493929190614739565b6020604051602081039080840390855afa1580156117d2573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101516001600160a01b03811660009081526003602090815290849020838501909452925460ff80821615158085526101009092041693830193909352909550925090506118595760405162461bcd60e51b8152600401610a2690614e19565b826020015160080260ff166001901b8401935050808061187890615438565b91505061170b565b5081827e0101010101010101010101010101010101010101010101010101010101010116146118c15760405162461bcd60e51b8152600401610a2690614b19565b50600091506119109050838d836020020135848e8e8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061308692505050565b905061191e83828633613508565b505050505050505050505050565b600063ffffffff82111561194257506000610703565b5063ffffffff166000908152600c6020526040902054601790810b900b90565b600063ffffffff82111561197857506000610703565b5063ffffffff9081166000908152600c60205260409020547c010000000000000000000000000000000000000000000000000000000090041690565b6119bc610f72565b6001600160a01b0316336001600160a01b03161480611a7457506012546040517f6b14daf80000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636b14daf890611a249033906000903690600401614620565b60206040518083038186803b158015611a3c57600080fd5b505afa158015611a50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a7491906141f7565b611a905760405162461bcd60e51b8152600401610a2690614d08565b6000611a9a613659565b6011546040517f70a082310000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b03909116906370a0823190611ae990309060040161460c565b60206040518083038186803b158015611b0157600080fd5b505afa158015611b15573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b39919061426e565b905081811015611b5b5760405162461bcd60e51b8152600401610a2690614de2565b6011546001600160a01b031663a9059cbb85611b80611b7a86866153c1565b87613856565b6040518363ffffffff1660e01b8152600401611b9d92919061467d565b602060405180830381600087803b158015611bb757600080fd5b505af1158015611bcb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bef91906141f7565b611c0b5760405162461bcd60e51b8152600401610a2690614f61565b50505050565b6012546001600160a01b031690565b6011546040517f70a0823100000000000000000000000000000000000000000000000000000000815260009182916001600160a01b03909116906370a0823190611c6e90309060040161460c565b60206040518083038186803b158015611c8657600080fd5b505afa158015611c9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cbe919061426e565b90506000611cca613659565b9050611cd6818361534d565b9250505090565b600f546001600160a01b031690565b611cf46128ef565b601f86511115611d165760405162461bcd60e51b8152600401610a2690614a06565b8451865114611d375760405162461bcd60e51b8152600401610a2690614a3d565b8551611d44856003615324565b60ff1610611d645760405162461bcd60e51b8152600401610a2690614f2a565b611d708460ff16613870565b825115611d8f5760405162461bcd60e51b8152600401610a2690614c9a565b60006040518060c001604052808881526020018781526020018660ff16815260200160017f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000604051602001611e07939291906145c4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815290825267ffffffffffffffff8616602083015201839052600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000ff1690559050611e7e61291b565b60045460005b81811015611f7957600060048281548110611eaf57634e487b7160e01b600052603260045260246000fd5b6000918252602082200154600580546001600160a01b0390921693509084908110611eea57634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b03948516835260038252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905594168252600290529190912080547fffffffffffffffffffffffffffffffffffff00000000000000000000000000001690555080611f7181615438565b915050611e84565b50611f8660046000613c50565b611f9260056000613c50565b60005b825151811015612276576003600084600001518381518110611fc757634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff161561200b5760405162461bcd60e51b8152600401610a2690614bbe565b604080518082019091526001815260ff82166020820152835180516003916000918590811061204a57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252818101929092526040016000908120835181549484015160ff16610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff9115157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0090961695909517169390931790925584015180516002929190849081106120fb57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff161561213f5760405162461bcd60e51b8152600401610a26906148f5565b60405180606001604052806001151581526020018260ff16815260200160006bffffffffffffffffffffffff16815250600260008560200151848151811061219757634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252818101929092526040908101600020835181549385015194909201517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909316911515919091177fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1661010060ff90941693909302929092177fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff16620100006bffffffffffffffffffffffff909216919091021790558061226e81615438565b915050611f95565b508151805161228d91600491602090910190613c6e565b5060208083015180516122a4926005920190613c6e565b506040820151600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600d80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292909217808555920481169260019290916000916123329185911661527a565b92506101000a81548163ffffffff021916908363ffffffff1602179055506123914630600d60009054906101000a900463ffffffff1663ffffffff1686600001518760200151886040015189606001518a608001518b60a00151612d2e565b600a819055600d5484516020860151604080880151606089015160808a015160a08b015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05986123fa988b98919763ffffffff9091169691959094909390929091906150ca565b60405180910390a1600b546601000000000000900463ffffffff1660005b8451518110156124805781600682601f811061244457634e487b7160e01b600052603260045260246000fd5b600891828204019190066004026101000a81548163ffffffff021916908363ffffffff160217905550808061247890615438565b915050612418565b5050505050505050505050565b6001600160a01b03811660009081526002602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046bffffffffffffffffffffffff1691810191909152906124f5576000915050610703565b6006816020015160ff16601f811061251d57634e487b7160e01b600052603260045260246000fd5b600881049190910154600b54612553926007166004026101000a90910463ffffffff9081169166010000000000009004166153d8565b9392505050565b60008080808033321461257f5760405162461bcd60e51b8152600401610a2690614ae2565b5050600a54600b5463ffffffff6601000000000000820481166000908152600c60205260409020549296610100909204600881901c8216965064ffffffffff169450601783900b93507c010000000000000000000000000000000000000000000000000000000090920490911690565b6011546001600160a01b031690565b6126066128ef565b60408051808201909152600e546001600160a01b038082168084527401000000000000000000000000000000000000000090920463ffffffff166020840152841614158061266457508163ffffffff16816020015163ffffffff1614155b1561273e576040805180820182526001600160a01b0385811680835263ffffffff86166020938401819052600e80547fffffffffffffffffffffffff00000000000000000000000000000000000000001683177fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000009092029190911790558451928501519351909392909116917fb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca15419161273591908790615150565b60405180910390a35b505050565b6001600160a01b0382811660009081526013602052604090205416331461277c5760405162461bcd60e51b8152600401610a26906148be565b336001600160a01b03821614156127a55760405162461bcd60e51b8152600401610a2690614ebe565b6001600160a01b03808316600090815260146020526040902080548383167fffffffffffffffffffffffff00000000000000000000000000000000000000008216811790925590911690811461273e576040516001600160a01b038084169133918616907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e3836790600090a4505050565b61283c6128ef565b610f6f81613890565b61284d6128ef565b610f6f81613922565b600b546601000000000000900463ffffffff9081166000818152600c602090815260409182902082516060810184529054601781810b810b810b8084527801000000000000000000000000000000000000000000000000830488169484018590527c01000000000000000000000000000000000000000000000000000000009092049096169190930181905292949190930b9291908490565b6000546001600160a01b031633146129195760405162461bcd60e51b8152600401610a26906149cf565b565b601154600b54604080516103e08101918290526001600160a01b0390931692660100000000000090920463ffffffff1691600091600690601f908285855b82829054906101000a900463ffffffff1663ffffffff168152602001906004019060208260030104928301926001038202915080841161295957905050505050509050600060058054806020026020016040519081016040528092919081815260200182805480156129f457602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116129d6575b5050505050905060005b8151811015612d2057600060026000848481518110612a2d57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160029054906101000a90046bffffffffffffffffffffffff166bffffffffffffffffffffffff169050600060026000858581518110612aa757634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160026101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff16021790555060008483601f8110612b2257634e487b7160e01b600052603260045260246000fd5b6020020151600b5490870363ffffffff90811692507201000000000000000000000000000000000000909104168102633b9aca000282018015612d1557600060136000878781518110612b8557634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b03169050886001600160a01b031663a9059cbb82846040518363ffffffff1660e01b8152600401612bf192919061467d565b602060405180830381600087803b158015612c0b57600080fd5b505af1158015612c1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c4391906141f7565b612c5f5760405162461bcd60e51b8152600401610a2690614f61565b878786601f8110612c8057634e487b7160e01b600052603260045260246000fd5b602002019063ffffffff16908163ffffffff1681525050886001600160a01b0316816001600160a01b0316878781518110612ccb57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c85604051612d0b91906145ab565b60405180910390a4505b5050506001016129fe565b50611397600683601f613ceb565b6000808a8a8a8a8a8a8a8a8a604051602001612d5299989796959493929190615006565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150505b9998505050505050505050565b6001600160a01b0381166000908152600260209081526040918290208251606081018452905460ff80821615158084526101008304909116938301939093526201000090046bffffffffffffffffffffffff1692810192909252612e3e5750610f6f565b6000612e49836105d6565b9050801561273e576001600160a01b03808416600090815260136020526040908190205460115491517fa9059cbb00000000000000000000000000000000000000000000000000000000815290831692919091169063a9059cbb90612eb4908490869060040161467d565b602060405180830381600087803b158015612ece57600080fd5b505af1158015612ee2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f0691906141f7565b612f225760405162461bcd60e51b8152600401610a2690614f61565b600b60000160069054906101000a900463ffffffff166006846020015160ff16601f8110612f6057634e487b7160e01b600052603260045260246000fd5b6008810491909101805460079092166004026101000a63ffffffff8181021990931693909216919091029190911790556001600160a01b038481166000818152600260205260409081902080547fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff169055601154905190831692841691907fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c9061300b9087906145ab565b60405180910390a450505050565b60006130268260206152e7565b6130318560206152e7565b61303d88610144615262565b6130479190615262565b6130519190615262565b61305c906000615262565b905036811461307d5760405162461bcd60e51b8152600401610a2690614a74565b50505050505050565b6000806130928361399e565b9050601f81604001515111156130ba5760405162461bcd60e51b8152600401610a2690614c2c565b604081015151865160ff16106130e25760405162461bcd60e51b8152600401610a2690614b50565b64ffffffffff841660208701526040810151805160009190613106906002906152c7565b8151811061312457634e487b7160e01b600052603260045260246000fd5b602002602001015190508060170b7f000000000000000000000000000000000000000000000000000000000000000060170b1315801561318a57507f000000000000000000000000000000000000000000000000000000000000000060170b8160170b13155b6131a65760405162461bcd60e51b8152600401610a2690614c63565b604087018051906131b682615471565b63ffffffff1663ffffffff168152505060405180606001604052808260170b8152602001836000015163ffffffff1681526020014263ffffffff16815250600c6000896040015163ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a81548177ffffffffffffffffffffffffffffffffffffffffffffffff021916908360170b77ffffffffffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160186101000a81548163ffffffff021916908363ffffffff160217905550604082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555090505086600b60008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548164ffffffffff021916908364ffffffffff16021790555060408201518160000160066101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001600a6101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600e6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160000160126101000a81548163ffffffff021916908363ffffffff16021790555060c08201518160000160166101000a81548163ffffffff021916908363ffffffff16021790555060e082015181600001601a6101000a81548162ffffff021916908362ffffff160217905550905050866040015163ffffffff167fc797025feeeaf2cd924c99e9205acb8ec04d5cad21c41ce637a38fb6dee6016a823385600001518660400151876020015188606001518d8d60405161344a98979695949392919061479e565b60405180910390a260006001600160a01b0316876040015163ffffffff167f0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271846000015160405161349b91906150b9565b60405180910390a3866040015163ffffffff168160170b7f0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f426040516134e191906145ab565b60405180910390a36134fa87604001518260170b613a23565b506060015195945050505050565b60008360170b121561351957611c0b565b6000613540633b9aca003a04866080015163ffffffff16876060015163ffffffff16613b5b565b90506010360260005a905060006135698663ffffffff1685858b60e0015162ffffff1686613b81565b90506000670de0b6b3a764000077ffffffffffffffffffffffffffffffffffffffffffffffff891683026001600160a01b03881660009081526002602052604090205460c08c01519290910492506201000090046bffffffffffffffffffffffff9081169163ffffffff16633b9aca0002828401019081168211156135f45750505050505050611c0b565b6001600160a01b038816600090815260026020526040902080546bffffffffffffffffffffffff90921662010000027fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff90921691909117905550505050505050505050565b60008060058054806020026020016040519081016040528092919081815260200182805480156136b257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613694575b50508351600b54604080516103e08101918290529697509195660100000000000090910463ffffffff169450600093509150600690601f908285855b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116136ee5790505050505050905060005b8381101561378f578181601f811061375c57634e487b7160e01b600052603260045260246000fd5b602002015161376b90846153d8565b61377b9063ffffffff1687615262565b95508061378781615438565b915050613734565b50600b546137bd907201000000000000000000000000000000000000900463ffffffff16633b9aca006152e7565b6137c790866152e7565b945060005b8381101561384e57600260008683815181106137f857634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b031682528101919091526040016000205461383a906201000090046bffffffffffffffffffffffff1687615262565b95508061384681615438565b9150506137cc565b505050505090565b60008183101561386757508161386a565b50805b92915050565b80600010610f6f5760405162461bcd60e51b8152600401610a2690614e50565b6001600160a01b0381163314156138b95760405162461bcd60e51b8152600401610a2690614f98565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6012546001600160a01b039081169082168114610a9757601280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384161790556040517f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912906114229083908590614757565b6139a6613d82565b60008060606000858060200190518101906139c19190614384565b929650909450925090506139d58683613bb5565b81516040516000906139eb9086906020016145ab565b60408051918152928152825160808101845263ffffffff909716875260208701525084019190915260170b6060830152509050919050565b60408051808201909152600e546001600160a01b0381168083527401000000000000000000000000000000000000000090910463ffffffff166020830152613a6b5750610a97565b6000613a786001856153d8565b63ffffffff8082166000818152600c6020908152604091829020549087015187519251959650601791820b90910b94613b3f9491821693613ac492909187918c16908b9060240161509e565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fbeed9b5100000000000000000000000000000000000000000000000000000000179052613bfd565b6113975760405162461bcd60e51b8152600401610a2690614bf5565b60008383811015613b6e57600285850304015b613b788184613856565b95945050505050565b600081861015613ba35760405162461bcd60e51b8152600401610a269061492c565b50633b9aca0094039190910101020290565b600081516020613bc591906152e7565b613bd09060a0615262565b613bdb906000615262565b90508083511461273e5760405162461bcd60e51b8152600401610a2690614b87565b60005a6113888110613c315761138881039050846040820482031115613c31576000808451602086016000888af150600191505b509392505050565b604080518082019091526000808252602082015290565b5080546000825590600052602060002090810190610f6f9190613db3565b828054828255906000526020600020908101928215613cdb579160200282015b82811115613cdb57825182547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03909116178255602090920191600190910190613c8e565b50613ce7929150613db3565b5090565b600483019183908215613cdb5791602002820160005b83821115613d4557835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302613d01565b8015613d755782816101000a81549063ffffffff0219169055600401602081600301049283019260010302613d45565b5050613ce7929150613db3565b6040518060800160405280600063ffffffff1681526020016060815260200160608152602001600060170b81525090565b5b80821115613ce75760008155600101613db4565b8035610703816154c1565b60008083601f840112613de4578182fd5b50813567ffffffffffffffff811115613dfb578182fd5b6020830191508360208083028501011115613e1557600080fd5b9250929050565b600082601f830112613e2c578081fd5b81356020613e41613e3c8361523e565b615214565b8281528181019085830183850287018401881015613e5d578586fd5b855b85811015613e84578135613e72816154c1565b84529284019290840190600101613e5f565b5090979650505050505050565b60008083601f840112613ea2578182fd5b50813567ffffffffffffffff811115613eb9578182fd5b602083019150836020828501011115613e1557600080fd5b600082601f830112613ee1578081fd5b813567ffffffffffffffff811115613efb57613efb6154ab565b613f2c60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601615214565b818152846020838601011115613f40578283fd5b816020850160208301379081016020019190915292915050565b8051601781900b811461070357600080fd5b803567ffffffffffffffff8116811461070357600080fd5b803560ff8116811461070357600080fd5b600060208284031215613fa6578081fd5b8135612553816154c1565b60008060408385031215613fc3578081fd5b8235613fce816154c1565b91506020830135613fde816154c1565b809150509250929050565b60008060408385031215613ffb578182fd5b8235614006816154c1565b946020939093013593505050565b60008060008060408587031215614029578182fd5b843567ffffffffffffffff80821115614040578384fd5b61404c88838901613dd3565b90965094506020870135915080821115614064578384fd5b5061407187828801613dd3565b95989497509550505050565b60008060008060008060c08789031215614095578384fd5b863567ffffffffffffffff808211156140ac578586fd5b6140b88a838b01613e1c565b975060208901359150808211156140cd578586fd5b6140d98a838b01613e1c565b96506140e760408a01613f84565b955060608901359150808211156140fc578384fd5b6141088a838b01613ed1565b945061411660808a01613f6c565b935060a089013591508082111561412b578283fd5b5061413889828a01613ed1565b9150509295509295509295565b60008060008060008060008060e0898b031215614160578586fd5b606089018a811115614170578687fd5b8998503567ffffffffffffffff80821115614189578788fd5b6141958c838d01613e91565b909950975060808b01359150808211156141ad578384fd5b6141b98c838d01613dd3565b909750955060a08b01359150808211156141d1578384fd5b506141de8b828c01613dd3565b999c989b50969995989497949560c00135949350505050565b600060208284031215614208578081fd5b81518015158114612553578182fd5b60008060408385031215614229578182fd5b8235614234816154c1565b91506020830135613fde816154d6565b60008060408385031215613fc3578182fd5b600060208284031215614267578081fd5b5035919050565b60006020828403121561427f578081fd5b5051919050565b6000806000806000806000806000806101208b8d0312156142a5578384fd5b8a3599506142b560208c01613dc8565b98506142c360408c01613f6c565b975060608b013567ffffffffffffffff808211156142df578586fd5b6142eb8e838f01613e1c565b985060808d0135915080821115614300578586fd5b61430c8e838f01613e1c565b975061431a60a08e01613f84565b965060c08d013591508082111561432f578586fd5b61433b8e838f01613e91565b909650945084915061434f60e08e01613f6c565b93506101008d0135915080821115614365578283fd5b506143728d828e01613ed1565b9150509295989b9194979a5092959850565b60008060008060808587031215614399578182fd5b84516143a4816154d6565b809450506020808601519350604086015167ffffffffffffffff8111156143c9578384fd5b8601601f810188136143d9578384fd5b80516143e7613e3c8261523e565b81815283810190838501858402850186018c1015614403578788fd5b8794505b8385101561442c5761441881613f5a565b835260019490940193918501918501614407565b50809650505050505061444160608601613f5a565b905092959194509250565b600080600080600060a08688031215614463578283fd5b853561446e816154d6565b9450602086013561447e816154d6565b9350604086013561448e816154d6565b9250606086013561449e816154d6565b9150608086013562ffffff811681146144b5578182fd5b809150509295509295909350565b6000602082840312156144d4578081fd5b813569ffffffffffffffffffff81168114612553578182fd5b6000815180845260208085019450808401835b838110156145255781516001600160a01b031687529582019590820190600101614500565b509495945050505050565b60008151808452815b8181101561455557602081850181015186830182015201614539565b818111156145665782602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60170b9052565b64ffffffffff169052565b90815260200190565b6000828483379101908152919050565b60f89390931b7fff00000000000000000000000000000000000000000000000000000000000000168352601791820b604090811b6001850152910b901b601982015260310190565b6001600160a01b0391909116815260200190565b60006001600160a01b03851682526040602083015282604083015282846060840137818301606090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010192915050565b6001600160a01b03929092168252602082015260400190565b60006020825261255360208301846144ed565b9215158352602083019190915263ffffffff16604082015260600190565b828152608081016060836020840137600081529392505050565b92835263ffffffff91909116602083015260ff16604082015260600190565b94855263ffffffff93909316602085015260ff91909116604084015260170b606083015267ffffffffffffffff16608082015260a00190565b93845260ff9290921660208401526040830152606082015260800190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b0392909216825263ffffffff16602082015260400190565b60179190910b815260200190565b600061010080830160178c810b855260206001600160a01b038d168187015263ffffffff8c1660408701528360608701528293508a5180845261012087019450818c019350855b81811015614803578451840b865294820194938201936001016147e5565b5050505050828103608084015261481a8188614530565b91505061482a60a0830186614599565b8360c0830152612dcd60e08301846145a0565b6000602082526125536020830184614530565b6020808252601f908201527f6f6e6c792070726f706f736564207061796565732063616e2061636365707400604082015260600190565b60208082526016908201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604082015260600190565b6020808252601d908201527f6f6e6c792063757272656e742070617965652063616e20757064617465000000604082015260600190565b6020808252601c908201527f7265706561746564207472616e736d6974746572206164647265737300000000604082015260600190565b6020808252818101527f6c6566744761732063616e6e6f742065786365656420696e697469616c476173604082015260600190565b6020808252600c908201527f7374616c65207265706f72740000000000000000000000000000000000000000604082015260600190565b6020808252601d908201527f4f6e6c79206f776e6572267265717565737465722063616e2063616c6c000000604082015260600190565b60208082526016908201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604082015260600190565b60208082526010908201527f746f6f206d616e79206f7261636c657300000000000000000000000000000000604082015260600190565b60208082526016908201527f6f7261636c65206c656e677468206d69736d6174636800000000000000000000604082015260600190565b60208082526018908201527f63616c6c64617461206c656e677468206d69736d617463680000000000000000604082015260600190565b6020808252601f908201527f7472616e736665722072656d61696e696e672066756e6473206661696c656400604082015260600190565b60208082526014908201527f4f6e6c792063616c6c61626c6520627920454f41000000000000000000000000604082015260600190565b60208082526010908201527f6475706c6963617465207369676e657200000000000000000000000000000000604082015260600190565b6020808252601e908201527f746f6f206665772076616c75657320746f207472757374206d656469616e0000604082015260600190565b60208082526016908201527f7265706f7274206c656e677468206d69736d6174636800000000000000000000604082015260600190565b60208082526017908201527f7265706561746564207369676e65722061646472657373000000000000000000604082015260600190565b60208082526010908201527f696e73756666696369656e742067617300000000000000000000000000000000604082015260600190565b6020808252601e908201527f6e756d206f62736572766174696f6e73206f7574206f6620626f756e64730000604082015260600190565b6020808252601e908201527f6d656469616e206973206f7574206f66206d696e2d6d61782072616e67650000604082015260600190565b6020808252601b908201527f6f6e636861696e436f6e666967206d75737420626520656d7074790000000000604082015260600190565b60208082526011908201527f706179656520616c726561647920736574000000000000000000000000000000604082015260600190565b6020808252818101527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c604082015260600190565b60208082526018908201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604082015260600190565b6020808252601a908201527f77726f6e67206e756d626572206f66207369676e617475726573000000000000604082015260600190565b60208082526017908201527f4f6e6c792070617965652063616e207769746864726177000000000000000000604082015260600190565b60208082526014908201527f696e73756666696369656e742062616c616e6365000000000000000000000000604082015260600190565b6020808252600f908201527f7369676e6174757265206572726f720000000000000000000000000000000000604082015260600190565b60208082526012908201527f66206d75737420626520706f7369746976650000000000000000000000000000604082015260600190565b60208082526015908201527f636f6e666967446967657374206d69736d617463680000000000000000000000604082015260600190565b60208082526017908201527f63616e6e6f74207472616e7366657220746f2073656c66000000000000000000604082015260600190565b6020808252818101527f7472616e736d6974746572732e73697a6520213d207061796565732e73697a65604082015260600190565b60208082526018908201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604082015260600190565b60208082526012908201527f696e73756666696369656e742066756e64730000000000000000000000000000604082015260600190565b60208082526017908201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604082015260600190565b6020808252601e908201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604082015260600190565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526150408285018b6144ed565b91508382036080850152615054828a6144ed565b915060ff881660a085015283820360c08501526150718288614530565b90861660e0850152838103610100850152905061508e8185614530565b9c9b505050505050505050505050565b93845260208401929092526040830152606082015260800190565b63ffffffff91909116815260200190565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526150fa8184018a6144ed565b9050828103608084015261510e81896144ed565b905060ff871660a084015282810360c084015261512b8187614530565b905067ffffffffffffffff851660e084015282810361010084015261508e8185614530565b63ffffffff92831681529116602082015260400190565b63ffffffff9384168152919092166020820152604081019190915260600190565b63ffffffff958616815293851660208501529184166040840152909216606082015262ffffff909116608082015260a00190565b69ffffffffffffffffffff91909116815260200190565b69ffffffffffffffffffff9586168152602081019490945260408401929092526060830152909116608082015260a00190565b60ff91909116815260200190565b60405181810167ffffffffffffffff81118282101715615236576152366154ab565b604052919050565b600067ffffffffffffffff821115615258576152586154ab565b5060209081020190565b6000821982111561527557615275615495565b500190565b600063ffffffff80831681851680830382111561529957615299615495565b01949350505050565b600060ff821660ff84168060ff038211156152bf576152bf615495565b019392505050565b6000826152e257634e487b7160e01b81526012600452602481fd5b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561531f5761531f615495565b500290565b600060ff821660ff84168160ff048111821515161561534557615345615495565b029392505050565b6000808312837f80000000000000000000000000000000000000000000000000000000000000000183128115161561538757615387615495565b837f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0183138116156153bb576153bb615495565b50500390565b6000828210156153d3576153d3615495565b500390565b600063ffffffff838116908316818110156153f5576153f5615495565b039392505050565b60028104600182168061541157607f821691505b6020821081141561543257634e487b7160e01b600052602260045260246000fd5b50919050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561546a5761546a615495565b5060010190565b600063ffffffff8083168181141561548b5761548b615495565b6001019392505050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114610f6f57600080fd5b63ffffffff81168114610f6f57600080fdfea164736f6c6343000800000a",
}

var ExposedOCR2AggregatorABI = ExposedOCR2AggregatorMetaData.ABI

var ExposedOCR2AggregatorBin = ExposedOCR2AggregatorMetaData.Bin

func DeployExposedOCR2Aggregator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExposedOCR2Aggregator, error) {
	parsed, err := ExposedOCR2AggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExposedOCR2AggregatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExposedOCR2Aggregator{ExposedOCR2AggregatorCaller: ExposedOCR2AggregatorCaller{contract: contract}, ExposedOCR2AggregatorTransactor: ExposedOCR2AggregatorTransactor{contract: contract}, ExposedOCR2AggregatorFilterer: ExposedOCR2AggregatorFilterer{contract: contract}}, nil
}

type ExposedOCR2Aggregator struct {
	address common.Address
	abi     abi.ABI
	ExposedOCR2AggregatorCaller
	ExposedOCR2AggregatorTransactor
	ExposedOCR2AggregatorFilterer
}

type ExposedOCR2AggregatorCaller struct {
	contract *bind.BoundContract
}

type ExposedOCR2AggregatorTransactor struct {
	contract *bind.BoundContract
}

type ExposedOCR2AggregatorFilterer struct {
	contract *bind.BoundContract
}

type ExposedOCR2AggregatorSession struct {
	Contract     *ExposedOCR2Aggregator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ExposedOCR2AggregatorCallerSession struct {
	Contract *ExposedOCR2AggregatorCaller
	CallOpts bind.CallOpts
}

type ExposedOCR2AggregatorTransactorSession struct {
	Contract     *ExposedOCR2AggregatorTransactor
	TransactOpts bind.TransactOpts
}

type ExposedOCR2AggregatorRaw struct {
	Contract *ExposedOCR2Aggregator
}

type ExposedOCR2AggregatorCallerRaw struct {
	Contract *ExposedOCR2AggregatorCaller
}

type ExposedOCR2AggregatorTransactorRaw struct {
	Contract *ExposedOCR2AggregatorTransactor
}

func NewExposedOCR2Aggregator(address common.Address, backend bind.ContractBackend) (*ExposedOCR2Aggregator, error) {
	abi, err := abi.JSON(strings.NewReader(ExposedOCR2AggregatorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindExposedOCR2Aggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2Aggregator{address: address, abi: abi, ExposedOCR2AggregatorCaller: ExposedOCR2AggregatorCaller{contract: contract}, ExposedOCR2AggregatorTransactor: ExposedOCR2AggregatorTransactor{contract: contract}, ExposedOCR2AggregatorFilterer: ExposedOCR2AggregatorFilterer{contract: contract}}, nil
}

func NewExposedOCR2AggregatorCaller(address common.Address, caller bind.ContractCaller) (*ExposedOCR2AggregatorCaller, error) {
	contract, err := bindExposedOCR2Aggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorCaller{contract: contract}, nil
}

func NewExposedOCR2AggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ExposedOCR2AggregatorTransactor, error) {
	contract, err := bindExposedOCR2Aggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorTransactor{contract: contract}, nil
}

func NewExposedOCR2AggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ExposedOCR2AggregatorFilterer, error) {
	contract, err := bindExposedOCR2Aggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorFilterer{contract: contract}, nil
}

func bindExposedOCR2Aggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExposedOCR2AggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExposedOCR2Aggregator.Contract.ExposedOCR2AggregatorCaller.contract.Call(opts, result, method, params...)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.ExposedOCR2AggregatorTransactor.contract.Transfer(opts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.ExposedOCR2AggregatorTransactor.contract.Transact(opts, method, params...)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExposedOCR2Aggregator.Contract.contract.Call(opts, result, method, params...)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.contract.Transfer(opts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.contract.Transact(opts, method, params...)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) Decimals() (uint8, error) {
	return _ExposedOCR2Aggregator.Contract.Decimals(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) Decimals() (uint8, error) {
	return _ExposedOCR2Aggregator.Contract.Decimals(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) Description() (string, error) {
	return _ExposedOCR2Aggregator.Contract.Description(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) Description() (string, error) {
	return _ExposedOCR2Aggregator.Contract.Description(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) ExposedConfigDigestFromConfigData(opts *bind.CallOpts, _chainId *big.Int, _contractAddress common.Address, _configCount uint64, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _encodedConfigVersion uint64, _encodedConfig []byte) ([32]byte, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "exposedConfigDigestFromConfigData", _chainId, _contractAddress, _configCount, _signers, _transmitters, _f, _onchainConfig, _encodedConfigVersion, _encodedConfig)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) ExposedConfigDigestFromConfigData(_chainId *big.Int, _contractAddress common.Address, _configCount uint64, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _encodedConfigVersion uint64, _encodedConfig []byte) ([32]byte, error) {
	return _ExposedOCR2Aggregator.Contract.ExposedConfigDigestFromConfigData(&_ExposedOCR2Aggregator.CallOpts, _chainId, _contractAddress, _configCount, _signers, _transmitters, _f, _onchainConfig, _encodedConfigVersion, _encodedConfig)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) ExposedConfigDigestFromConfigData(_chainId *big.Int, _contractAddress common.Address, _configCount uint64, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _encodedConfigVersion uint64, _encodedConfig []byte) ([32]byte, error) {
	return _ExposedOCR2Aggregator.Contract.ExposedConfigDigestFromConfigData(&_ExposedOCR2Aggregator.CallOpts, _chainId, _contractAddress, _configCount, _signers, _transmitters, _f, _onchainConfig, _encodedConfigVersion, _encodedConfig)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) GetAnswer(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "getAnswer", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.GetAnswer(&_ExposedOCR2Aggregator.CallOpts, roundId)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.GetAnswer(&_ExposedOCR2Aggregator.CallOpts, roundId)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) GetBilling(opts *bind.CallOpts) (GetBilling,

	error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "getBilling")

	outstruct := new(GetBilling)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaximumGasPriceGwei = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.ReasonableGasPriceGwei = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ObservationPaymentGjuels = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.TransmissionPaymentGjuels = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.AccountingGas = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) GetBilling() (GetBilling,

	error) {
	return _ExposedOCR2Aggregator.Contract.GetBilling(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) GetBilling() (GetBilling,

	error) {
	return _ExposedOCR2Aggregator.Contract.GetBilling(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) GetBillingAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "getBillingAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) GetBillingAccessController() (common.Address, error) {
	return _ExposedOCR2Aggregator.Contract.GetBillingAccessController(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) GetBillingAccessController() (common.Address, error) {
	return _ExposedOCR2Aggregator.Contract.GetBillingAccessController(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) GetLinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "getLinkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) GetLinkToken() (common.Address, error) {
	return _ExposedOCR2Aggregator.Contract.GetLinkToken(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) GetLinkToken() (common.Address, error) {
	return _ExposedOCR2Aggregator.Contract.GetLinkToken(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) GetRequesterAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "getRequesterAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) GetRequesterAccessController() (common.Address, error) {
	return _ExposedOCR2Aggregator.Contract.GetRequesterAccessController(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) GetRequesterAccessController() (common.Address, error) {
	return _ExposedOCR2Aggregator.Contract.GetRequesterAccessController(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) GetRoundData(opts *bind.CallOpts, roundId *big.Int) (GetRoundData,

	error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "getRoundData", roundId)

	outstruct := new(GetRoundData)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) GetRoundData(roundId *big.Int) (GetRoundData,

	error) {
	return _ExposedOCR2Aggregator.Contract.GetRoundData(&_ExposedOCR2Aggregator.CallOpts, roundId)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) GetRoundData(roundId *big.Int) (GetRoundData,

	error) {
	return _ExposedOCR2Aggregator.Contract.GetRoundData(&_ExposedOCR2Aggregator.CallOpts, roundId)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) GetTimestamp(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "getTimestamp", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.GetTimestamp(&_ExposedOCR2Aggregator.CallOpts, roundId)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.GetTimestamp(&_ExposedOCR2Aggregator.CallOpts, roundId)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) GetTransmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "getTransmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) GetTransmitters() ([]common.Address, error) {
	return _ExposedOCR2Aggregator.Contract.GetTransmitters(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) GetTransmitters() ([]common.Address, error) {
	return _ExposedOCR2Aggregator.Contract.GetTransmitters(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) GetValidatorConfig(opts *bind.CallOpts) (GetValidatorConfig,

	error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "getValidatorConfig")

	outstruct := new(GetValidatorConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.GasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) GetValidatorConfig() (GetValidatorConfig,

	error) {
	return _ExposedOCR2Aggregator.Contract.GetValidatorConfig(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) GetValidatorConfig() (GetValidatorConfig,

	error) {
	return _ExposedOCR2Aggregator.Contract.GetValidatorConfig(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) LatestAnswer() (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.LatestAnswer(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) LatestAnswer() (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.LatestAnswer(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _ExposedOCR2Aggregator.Contract.LatestConfigDetails(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _ExposedOCR2Aggregator.Contract.LatestConfigDetails(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _ExposedOCR2Aggregator.Contract.LatestConfigDigestAndEpoch(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _ExposedOCR2Aggregator.Contract.LatestConfigDigestAndEpoch(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) LatestRound() (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.LatestRound(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) LatestRound() (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.LatestRound(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) LatestRoundData(opts *bind.CallOpts) (LatestRoundData,

	error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(LatestRoundData)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) LatestRoundData() (LatestRoundData,

	error) {
	return _ExposedOCR2Aggregator.Contract.LatestRoundData(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) LatestRoundData() (LatestRoundData,

	error) {
	return _ExposedOCR2Aggregator.Contract.LatestRoundData(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) LatestTimestamp() (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.LatestTimestamp(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) LatestTimestamp() (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.LatestTimestamp(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) LatestTransmissionDetails(opts *bind.CallOpts) (LatestTransmissionDetails,

	error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "latestTransmissionDetails")

	outstruct := new(LatestTransmissionDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigDigest = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Round = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.LatestAnswer = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LatestTimestamp = *abi.ConvertType(out[4], new(uint64)).(*uint64)

	return *outstruct, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) LatestTransmissionDetails() (LatestTransmissionDetails,

	error) {
	return _ExposedOCR2Aggregator.Contract.LatestTransmissionDetails(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) LatestTransmissionDetails() (LatestTransmissionDetails,

	error) {
	return _ExposedOCR2Aggregator.Contract.LatestTransmissionDetails(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) LinkAvailableForPayment() (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.LinkAvailableForPayment(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.LinkAvailableForPayment(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) MaxAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "maxAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) MaxAnswer() (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.MaxAnswer(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) MaxAnswer() (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.MaxAnswer(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) MinAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "minAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) MinAnswer() (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.MinAnswer(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) MinAnswer() (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.MinAnswer(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) OracleObservationCount(opts *bind.CallOpts, transmitterAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "oracleObservationCount", transmitterAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _ExposedOCR2Aggregator.Contract.OracleObservationCount(&_ExposedOCR2Aggregator.CallOpts, transmitterAddress)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _ExposedOCR2Aggregator.Contract.OracleObservationCount(&_ExposedOCR2Aggregator.CallOpts, transmitterAddress)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) OwedPayment(opts *bind.CallOpts, transmitterAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "owedPayment", transmitterAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.OwedPayment(&_ExposedOCR2Aggregator.CallOpts, transmitterAddress)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.OwedPayment(&_ExposedOCR2Aggregator.CallOpts, transmitterAddress)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) Owner() (common.Address, error) {
	return _ExposedOCR2Aggregator.Contract.Owner(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) Owner() (common.Address, error) {
	return _ExposedOCR2Aggregator.Contract.Owner(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) TypeAndVersion() (string, error) {
	return _ExposedOCR2Aggregator.Contract.TypeAndVersion(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) TypeAndVersion() (string, error) {
	return _ExposedOCR2Aggregator.Contract.TypeAndVersion(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExposedOCR2Aggregator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) Version() (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.Version(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorCallerSession) Version() (*big.Int, error) {
	return _ExposedOCR2Aggregator.Contract.Version(&_ExposedOCR2Aggregator.CallOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "acceptOwnership")
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.AcceptOwnership(&_ExposedOCR2Aggregator.TransactOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.AcceptOwnership(&_ExposedOCR2Aggregator.TransactOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) AcceptPayeeship(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "acceptPayeeship", transmitter)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.AcceptPayeeship(&_ExposedOCR2Aggregator.TransactOpts, transmitter)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.AcceptPayeeship(&_ExposedOCR2Aggregator.TransactOpts, transmitter)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) RequestNewRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "requestNewRound")
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) RequestNewRound() (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.RequestNewRound(&_ExposedOCR2Aggregator.TransactOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) RequestNewRound() (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.RequestNewRound(&_ExposedOCR2Aggregator.TransactOpts)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) SetBilling(opts *bind.TransactOpts, maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "setBilling", maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.SetBilling(&_ExposedOCR2Aggregator.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.SetBilling(&_ExposedOCR2Aggregator.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) SetBillingAccessController(opts *bind.TransactOpts, _billingAccessController common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "setBillingAccessController", _billingAccessController)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.SetBillingAccessController(&_ExposedOCR2Aggregator.TransactOpts, _billingAccessController)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.SetBillingAccessController(&_ExposedOCR2Aggregator.TransactOpts, _billingAccessController)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "setConfig", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.SetConfig(&_ExposedOCR2Aggregator.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.SetConfig(&_ExposedOCR2Aggregator.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) SetLinkToken(opts *bind.TransactOpts, linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "setLinkToken", linkToken, recipient)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) SetLinkToken(linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.SetLinkToken(&_ExposedOCR2Aggregator.TransactOpts, linkToken, recipient)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) SetLinkToken(linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.SetLinkToken(&_ExposedOCR2Aggregator.TransactOpts, linkToken, recipient)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) SetPayees(opts *bind.TransactOpts, transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "setPayees", transmitters, payees)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.SetPayees(&_ExposedOCR2Aggregator.TransactOpts, transmitters, payees)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.SetPayees(&_ExposedOCR2Aggregator.TransactOpts, transmitters, payees)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) SetRequesterAccessController(opts *bind.TransactOpts, requesterAccessController common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "setRequesterAccessController", requesterAccessController)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) SetRequesterAccessController(requesterAccessController common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.SetRequesterAccessController(&_ExposedOCR2Aggregator.TransactOpts, requesterAccessController)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) SetRequesterAccessController(requesterAccessController common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.SetRequesterAccessController(&_ExposedOCR2Aggregator.TransactOpts, requesterAccessController)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) SetValidatorConfig(opts *bind.TransactOpts, newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "setValidatorConfig", newValidator, newGasLimit)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) SetValidatorConfig(newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.SetValidatorConfig(&_ExposedOCR2Aggregator.TransactOpts, newValidator, newGasLimit)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) SetValidatorConfig(newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.SetValidatorConfig(&_ExposedOCR2Aggregator.TransactOpts, newValidator, newGasLimit)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "transferOwnership", to)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.TransferOwnership(&_ExposedOCR2Aggregator.TransactOpts, to)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.TransferOwnership(&_ExposedOCR2Aggregator.TransactOpts, to)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) TransferPayeeship(opts *bind.TransactOpts, transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "transferPayeeship", transmitter, proposed)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.TransferPayeeship(&_ExposedOCR2Aggregator.TransactOpts, transmitter, proposed)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.TransferPayeeship(&_ExposedOCR2Aggregator.TransactOpts, transmitter, proposed)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.Transmit(&_ExposedOCR2Aggregator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.Transmit(&_ExposedOCR2Aggregator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) WithdrawFunds(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "withdrawFunds", recipient, amount)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.WithdrawFunds(&_ExposedOCR2Aggregator.TransactOpts, recipient, amount)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.WithdrawFunds(&_ExposedOCR2Aggregator.TransactOpts, recipient, amount)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactor) WithdrawPayment(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.contract.Transact(opts, "withdrawPayment", transmitter)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.WithdrawPayment(&_ExposedOCR2Aggregator.TransactOpts, transmitter)
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorTransactorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _ExposedOCR2Aggregator.Contract.WithdrawPayment(&_ExposedOCR2Aggregator.TransactOpts, transmitter)
}

type ExposedOCR2AggregatorAnswerUpdatedIterator struct {
	Event *ExposedOCR2AggregatorAnswerUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorAnswerUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorAnswerUpdated)
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
		it.Event = new(ExposedOCR2AggregatorAnswerUpdated)
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

func (it *ExposedOCR2AggregatorAnswerUpdatedIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*ExposedOCR2AggregatorAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorAnswerUpdatedIterator{contract: _ExposedOCR2Aggregator.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorAnswerUpdated)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParseAnswerUpdated(log types.Log) (*ExposedOCR2AggregatorAnswerUpdated, error) {
	event := new(ExposedOCR2AggregatorAnswerUpdated)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorBillingAccessControllerSetIterator struct {
	Event *ExposedOCR2AggregatorBillingAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorBillingAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorBillingAccessControllerSet)
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
		it.Event = new(ExposedOCR2AggregatorBillingAccessControllerSet)
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

func (it *ExposedOCR2AggregatorBillingAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorBillingAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorBillingAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*ExposedOCR2AggregatorBillingAccessControllerSetIterator, error) {

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorBillingAccessControllerSetIterator{contract: _ExposedOCR2Aggregator.contract, event: "BillingAccessControllerSet", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorBillingAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorBillingAccessControllerSet)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParseBillingAccessControllerSet(log types.Log) (*ExposedOCR2AggregatorBillingAccessControllerSet, error) {
	event := new(ExposedOCR2AggregatorBillingAccessControllerSet)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorBillingSetIterator struct {
	Event *ExposedOCR2AggregatorBillingSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorBillingSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorBillingSet)
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
		it.Event = new(ExposedOCR2AggregatorBillingSet)
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

func (it *ExposedOCR2AggregatorBillingSetIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorBillingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorBillingSet struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
	Raw                       types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterBillingSet(opts *bind.FilterOpts) (*ExposedOCR2AggregatorBillingSetIterator, error) {

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorBillingSetIterator{contract: _ExposedOCR2Aggregator.contract, event: "BillingSet", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchBillingSet(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorBillingSet) (event.Subscription, error) {

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorBillingSet)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "BillingSet", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParseBillingSet(log types.Log) (*ExposedOCR2AggregatorBillingSet, error) {
	event := new(ExposedOCR2AggregatorBillingSet)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "BillingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorConfigSetIterator struct {
	Event *ExposedOCR2AggregatorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorConfigSet)
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
		it.Event = new(ExposedOCR2AggregatorConfigSet)
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

func (it *ExposedOCR2AggregatorConfigSetIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterConfigSet(opts *bind.FilterOpts) (*ExposedOCR2AggregatorConfigSetIterator, error) {

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorConfigSetIterator{contract: _ExposedOCR2Aggregator.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorConfigSet) (event.Subscription, error) {

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorConfigSet)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParseConfigSet(log types.Log) (*ExposedOCR2AggregatorConfigSet, error) {
	event := new(ExposedOCR2AggregatorConfigSet)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorLinkTokenSetIterator struct {
	Event *ExposedOCR2AggregatorLinkTokenSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorLinkTokenSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorLinkTokenSet)
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
		it.Event = new(ExposedOCR2AggregatorLinkTokenSet)
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

func (it *ExposedOCR2AggregatorLinkTokenSetIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorLinkTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorLinkTokenSet struct {
	OldLinkToken common.Address
	NewLinkToken common.Address
	Raw          types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterLinkTokenSet(opts *bind.FilterOpts, oldLinkToken []common.Address, newLinkToken []common.Address) (*ExposedOCR2AggregatorLinkTokenSetIterator, error) {

	var oldLinkTokenRule []interface{}
	for _, oldLinkTokenItem := range oldLinkToken {
		oldLinkTokenRule = append(oldLinkTokenRule, oldLinkTokenItem)
	}
	var newLinkTokenRule []interface{}
	for _, newLinkTokenItem := range newLinkToken {
		newLinkTokenRule = append(newLinkTokenRule, newLinkTokenItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "LinkTokenSet", oldLinkTokenRule, newLinkTokenRule)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorLinkTokenSetIterator{contract: _ExposedOCR2Aggregator.contract, event: "LinkTokenSet", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchLinkTokenSet(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorLinkTokenSet, oldLinkToken []common.Address, newLinkToken []common.Address) (event.Subscription, error) {

	var oldLinkTokenRule []interface{}
	for _, oldLinkTokenItem := range oldLinkToken {
		oldLinkTokenRule = append(oldLinkTokenRule, oldLinkTokenItem)
	}
	var newLinkTokenRule []interface{}
	for _, newLinkTokenItem := range newLinkToken {
		newLinkTokenRule = append(newLinkTokenRule, newLinkTokenItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "LinkTokenSet", oldLinkTokenRule, newLinkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorLinkTokenSet)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "LinkTokenSet", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParseLinkTokenSet(log types.Log) (*ExposedOCR2AggregatorLinkTokenSet, error) {
	event := new(ExposedOCR2AggregatorLinkTokenSet)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "LinkTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorNewRoundIterator struct {
	Event *ExposedOCR2AggregatorNewRound

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorNewRoundIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorNewRound)
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
		it.Event = new(ExposedOCR2AggregatorNewRound)
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

func (it *ExposedOCR2AggregatorNewRoundIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*ExposedOCR2AggregatorNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorNewRoundIterator{contract: _ExposedOCR2Aggregator.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorNewRound)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParseNewRound(log types.Log) (*ExposedOCR2AggregatorNewRound, error) {
	event := new(ExposedOCR2AggregatorNewRound)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorNewTransmissionIterator struct {
	Event *ExposedOCR2AggregatorNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorNewTransmission)
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
		it.Event = new(ExposedOCR2AggregatorNewTransmission)
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

func (it *ExposedOCR2AggregatorNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorNewTransmission struct {
	AggregatorRoundId     uint32
	Answer                *big.Int
	Transmitter           common.Address
	ObservationsTimestamp uint32
	Observations          []*big.Int
	Observers             []byte
	JuelsPerFeeCoin       *big.Int
	ConfigDigest          [32]byte
	EpochAndRound         *big.Int
	Raw                   types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*ExposedOCR2AggregatorNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorNewTransmissionIterator{contract: _ExposedOCR2Aggregator.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorNewTransmission)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParseNewTransmission(log types.Log) (*ExposedOCR2AggregatorNewTransmission, error) {
	event := new(ExposedOCR2AggregatorNewTransmission)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorOraclePaidIterator struct {
	Event *ExposedOCR2AggregatorOraclePaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorOraclePaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorOraclePaid)
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
		it.Event = new(ExposedOCR2AggregatorOraclePaid)
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

func (it *ExposedOCR2AggregatorOraclePaidIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorOraclePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorOraclePaid struct {
	Transmitter common.Address
	Payee       common.Address
	Amount      *big.Int
	LinkToken   common.Address
	Raw         types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*ExposedOCR2AggregatorOraclePaidIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var linkTokenRule []interface{}
	for _, linkTokenItem := range linkToken {
		linkTokenRule = append(linkTokenRule, linkTokenItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorOraclePaidIterator{contract: _ExposedOCR2Aggregator.contract, event: "OraclePaid", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var linkTokenRule []interface{}
	for _, linkTokenItem := range linkToken {
		linkTokenRule = append(linkTokenRule, linkTokenItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorOraclePaid)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParseOraclePaid(log types.Log) (*ExposedOCR2AggregatorOraclePaid, error) {
	event := new(ExposedOCR2AggregatorOraclePaid)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorOwnershipTransferRequestedIterator struct {
	Event *ExposedOCR2AggregatorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorOwnershipTransferRequested)
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
		it.Event = new(ExposedOCR2AggregatorOwnershipTransferRequested)
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

func (it *ExposedOCR2AggregatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExposedOCR2AggregatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorOwnershipTransferRequestedIterator{contract: _ExposedOCR2Aggregator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorOwnershipTransferRequested)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*ExposedOCR2AggregatorOwnershipTransferRequested, error) {
	event := new(ExposedOCR2AggregatorOwnershipTransferRequested)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorOwnershipTransferredIterator struct {
	Event *ExposedOCR2AggregatorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorOwnershipTransferred)
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
		it.Event = new(ExposedOCR2AggregatorOwnershipTransferred)
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

func (it *ExposedOCR2AggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExposedOCR2AggregatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorOwnershipTransferredIterator{contract: _ExposedOCR2Aggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorOwnershipTransferred)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*ExposedOCR2AggregatorOwnershipTransferred, error) {
	event := new(ExposedOCR2AggregatorOwnershipTransferred)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorPayeeshipTransferRequestedIterator struct {
	Event *ExposedOCR2AggregatorPayeeshipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorPayeeshipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorPayeeshipTransferRequested)
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
		it.Event = new(ExposedOCR2AggregatorPayeeshipTransferRequested)
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

func (it *ExposedOCR2AggregatorPayeeshipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorPayeeshipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorPayeeshipTransferRequested struct {
	Transmitter common.Address
	Current     common.Address
	Proposed    common.Address
	Raw         types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*ExposedOCR2AggregatorPayeeshipTransferRequestedIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var proposedRule []interface{}
	for _, proposedItem := range proposed {
		proposedRule = append(proposedRule, proposedItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorPayeeshipTransferRequestedIterator{contract: _ExposedOCR2Aggregator.contract, event: "PayeeshipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var proposedRule []interface{}
	for _, proposedItem := range proposed {
		proposedRule = append(proposedRule, proposedItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorPayeeshipTransferRequested)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParsePayeeshipTransferRequested(log types.Log) (*ExposedOCR2AggregatorPayeeshipTransferRequested, error) {
	event := new(ExposedOCR2AggregatorPayeeshipTransferRequested)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorPayeeshipTransferredIterator struct {
	Event *ExposedOCR2AggregatorPayeeshipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorPayeeshipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorPayeeshipTransferred)
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
		it.Event = new(ExposedOCR2AggregatorPayeeshipTransferred)
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

func (it *ExposedOCR2AggregatorPayeeshipTransferredIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorPayeeshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorPayeeshipTransferred struct {
	Transmitter common.Address
	Previous    common.Address
	Current     common.Address
	Raw         types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*ExposedOCR2AggregatorPayeeshipTransferredIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorPayeeshipTransferredIterator{contract: _ExposedOCR2Aggregator.contract, event: "PayeeshipTransferred", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorPayeeshipTransferred)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParsePayeeshipTransferred(log types.Log) (*ExposedOCR2AggregatorPayeeshipTransferred, error) {
	event := new(ExposedOCR2AggregatorPayeeshipTransferred)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorRequesterAccessControllerSetIterator struct {
	Event *ExposedOCR2AggregatorRequesterAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorRequesterAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorRequesterAccessControllerSet)
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
		it.Event = new(ExposedOCR2AggregatorRequesterAccessControllerSet)
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

func (it *ExposedOCR2AggregatorRequesterAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorRequesterAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorRequesterAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterRequesterAccessControllerSet(opts *bind.FilterOpts) (*ExposedOCR2AggregatorRequesterAccessControllerSetIterator, error) {

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "RequesterAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorRequesterAccessControllerSetIterator{contract: _ExposedOCR2Aggregator.contract, event: "RequesterAccessControllerSet", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchRequesterAccessControllerSet(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorRequesterAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "RequesterAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorRequesterAccessControllerSet)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "RequesterAccessControllerSet", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParseRequesterAccessControllerSet(log types.Log) (*ExposedOCR2AggregatorRequesterAccessControllerSet, error) {
	event := new(ExposedOCR2AggregatorRequesterAccessControllerSet)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "RequesterAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorRoundRequestedIterator struct {
	Event *ExposedOCR2AggregatorRoundRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorRoundRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorRoundRequested)
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
		it.Event = new(ExposedOCR2AggregatorRoundRequested)
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

func (it *ExposedOCR2AggregatorRoundRequestedIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorRoundRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorRoundRequested struct {
	Requester    common.Address
	ConfigDigest [32]byte
	Epoch        uint32
	Round        uint8
	Raw          types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterRoundRequested(opts *bind.FilterOpts, requester []common.Address) (*ExposedOCR2AggregatorRoundRequestedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "RoundRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorRoundRequestedIterator{contract: _ExposedOCR2Aggregator.contract, event: "RoundRequested", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchRoundRequested(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorRoundRequested, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "RoundRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorRoundRequested)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "RoundRequested", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParseRoundRequested(log types.Log) (*ExposedOCR2AggregatorRoundRequested, error) {
	event := new(ExposedOCR2AggregatorRoundRequested)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "RoundRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorTransmittedIterator struct {
	Event *ExposedOCR2AggregatorTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorTransmitted)
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
		it.Event = new(ExposedOCR2AggregatorTransmitted)
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

func (it *ExposedOCR2AggregatorTransmittedIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterTransmitted(opts *bind.FilterOpts) (*ExposedOCR2AggregatorTransmittedIterator, error) {

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorTransmittedIterator{contract: _ExposedOCR2Aggregator.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorTransmitted) (event.Subscription, error) {

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorTransmitted)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParseTransmitted(log types.Log) (*ExposedOCR2AggregatorTransmitted, error) {
	event := new(ExposedOCR2AggregatorTransmitted)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedOCR2AggregatorValidatorConfigSetIterator struct {
	Event *ExposedOCR2AggregatorValidatorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedOCR2AggregatorValidatorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedOCR2AggregatorValidatorConfigSet)
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
		it.Event = new(ExposedOCR2AggregatorValidatorConfigSet)
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

func (it *ExposedOCR2AggregatorValidatorConfigSetIterator) Error() error {
	return it.fail
}

func (it *ExposedOCR2AggregatorValidatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedOCR2AggregatorValidatorConfigSet struct {
	PreviousValidator common.Address
	PreviousGasLimit  uint32
	CurrentValidator  common.Address
	CurrentGasLimit   uint32
	Raw               types.Log
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) FilterValidatorConfigSet(opts *bind.FilterOpts, previousValidator []common.Address, currentValidator []common.Address) (*ExposedOCR2AggregatorValidatorConfigSetIterator, error) {

	var previousValidatorRule []interface{}
	for _, previousValidatorItem := range previousValidator {
		previousValidatorRule = append(previousValidatorRule, previousValidatorItem)
	}

	var currentValidatorRule []interface{}
	for _, currentValidatorItem := range currentValidator {
		currentValidatorRule = append(currentValidatorRule, currentValidatorItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.FilterLogs(opts, "ValidatorConfigSet", previousValidatorRule, currentValidatorRule)
	if err != nil {
		return nil, err
	}
	return &ExposedOCR2AggregatorValidatorConfigSetIterator{contract: _ExposedOCR2Aggregator.contract, event: "ValidatorConfigSet", logs: logs, sub: sub}, nil
}

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) WatchValidatorConfigSet(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorValidatorConfigSet, previousValidator []common.Address, currentValidator []common.Address) (event.Subscription, error) {

	var previousValidatorRule []interface{}
	for _, previousValidatorItem := range previousValidator {
		previousValidatorRule = append(previousValidatorRule, previousValidatorItem)
	}

	var currentValidatorRule []interface{}
	for _, currentValidatorItem := range currentValidator {
		currentValidatorRule = append(currentValidatorRule, currentValidatorItem)
	}

	logs, sub, err := _ExposedOCR2Aggregator.contract.WatchLogs(opts, "ValidatorConfigSet", previousValidatorRule, currentValidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedOCR2AggregatorValidatorConfigSet)
				if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "ValidatorConfigSet", log); err != nil {
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

func (_ExposedOCR2Aggregator *ExposedOCR2AggregatorFilterer) ParseValidatorConfigSet(log types.Log) (*ExposedOCR2AggregatorValidatorConfigSet, error) {
	event := new(ExposedOCR2AggregatorValidatorConfigSet)
	if err := _ExposedOCR2Aggregator.contract.UnpackLog(event, "ValidatorConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetBilling struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
}
type GetRoundData struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}
type GetValidatorConfig struct {
	Validator common.Address
	GasLimit  uint32
}
type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}
type LatestConfigDigestAndEpoch struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}
type LatestRoundData struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}
type LatestTransmissionDetails struct {
	ConfigDigest    [32]byte
	Epoch           uint32
	Round           uint8
	LatestAnswer    *big.Int
	LatestTimestamp uint64
}

func (_ExposedOCR2Aggregator *ExposedOCR2Aggregator) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _ExposedOCR2Aggregator.abi.Events["AnswerUpdated"].ID:
		return _ExposedOCR2Aggregator.ParseAnswerUpdated(log)
	case _ExposedOCR2Aggregator.abi.Events["BillingAccessControllerSet"].ID:
		return _ExposedOCR2Aggregator.ParseBillingAccessControllerSet(log)
	case _ExposedOCR2Aggregator.abi.Events["BillingSet"].ID:
		return _ExposedOCR2Aggregator.ParseBillingSet(log)
	case _ExposedOCR2Aggregator.abi.Events["ConfigSet"].ID:
		return _ExposedOCR2Aggregator.ParseConfigSet(log)
	case _ExposedOCR2Aggregator.abi.Events["LinkTokenSet"].ID:
		return _ExposedOCR2Aggregator.ParseLinkTokenSet(log)
	case _ExposedOCR2Aggregator.abi.Events["NewRound"].ID:
		return _ExposedOCR2Aggregator.ParseNewRound(log)
	case _ExposedOCR2Aggregator.abi.Events["NewTransmission"].ID:
		return _ExposedOCR2Aggregator.ParseNewTransmission(log)
	case _ExposedOCR2Aggregator.abi.Events["OraclePaid"].ID:
		return _ExposedOCR2Aggregator.ParseOraclePaid(log)
	case _ExposedOCR2Aggregator.abi.Events["OwnershipTransferRequested"].ID:
		return _ExposedOCR2Aggregator.ParseOwnershipTransferRequested(log)
	case _ExposedOCR2Aggregator.abi.Events["OwnershipTransferred"].ID:
		return _ExposedOCR2Aggregator.ParseOwnershipTransferred(log)
	case _ExposedOCR2Aggregator.abi.Events["PayeeshipTransferRequested"].ID:
		return _ExposedOCR2Aggregator.ParsePayeeshipTransferRequested(log)
	case _ExposedOCR2Aggregator.abi.Events["PayeeshipTransferred"].ID:
		return _ExposedOCR2Aggregator.ParsePayeeshipTransferred(log)
	case _ExposedOCR2Aggregator.abi.Events["RequesterAccessControllerSet"].ID:
		return _ExposedOCR2Aggregator.ParseRequesterAccessControllerSet(log)
	case _ExposedOCR2Aggregator.abi.Events["RoundRequested"].ID:
		return _ExposedOCR2Aggregator.ParseRoundRequested(log)
	case _ExposedOCR2Aggregator.abi.Events["Transmitted"].ID:
		return _ExposedOCR2Aggregator.ParseTransmitted(log)
	case _ExposedOCR2Aggregator.abi.Events["ValidatorConfigSet"].ID:
		return _ExposedOCR2Aggregator.ParseValidatorConfigSet(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ExposedOCR2AggregatorAnswerUpdated) Topic() common.Hash {
	return common.HexToHash("0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f")
}

func (ExposedOCR2AggregatorBillingAccessControllerSet) Topic() common.Hash {
	return common.HexToHash("0x793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912")
}

func (ExposedOCR2AggregatorBillingSet) Topic() common.Hash {
	return common.HexToHash("0x0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f")
}

func (ExposedOCR2AggregatorConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (ExposedOCR2AggregatorLinkTokenSet) Topic() common.Hash {
	return common.HexToHash("0x4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a")
}

func (ExposedOCR2AggregatorNewRound) Topic() common.Hash {
	return common.HexToHash("0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271")
}

func (ExposedOCR2AggregatorNewTransmission) Topic() common.Hash {
	return common.HexToHash("0xc797025feeeaf2cd924c99e9205acb8ec04d5cad21c41ce637a38fb6dee6016a")
}

func (ExposedOCR2AggregatorOraclePaid) Topic() common.Hash {
	return common.HexToHash("0xd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c")
}

func (ExposedOCR2AggregatorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ExposedOCR2AggregatorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (ExposedOCR2AggregatorPayeeshipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367")
}

func (ExposedOCR2AggregatorPayeeshipTransferred) Topic() common.Hash {
	return common.HexToHash("0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3")
}

func (ExposedOCR2AggregatorRequesterAccessControllerSet) Topic() common.Hash {
	return common.HexToHash("0x27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634")
}

func (ExposedOCR2AggregatorRoundRequested) Topic() common.Hash {
	return common.HexToHash("0x41e3990591fd372502daa15842da15bc7f41c75309ab3ff4f56f1848c178825c")
}

func (ExposedOCR2AggregatorTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (ExposedOCR2AggregatorValidatorConfigSet) Topic() common.Hash {
	return common.HexToHash("0xb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541")
}

func (_ExposedOCR2Aggregator *ExposedOCR2Aggregator) Address() common.Address {
	return _ExposedOCR2Aggregator.address
}

type ExposedOCR2AggregatorInterface interface {
	Decimals(opts *bind.CallOpts) (uint8, error)

	Description(opts *bind.CallOpts) (string, error)

	ExposedConfigDigestFromConfigData(opts *bind.CallOpts, _chainId *big.Int, _contractAddress common.Address, _configCount uint64, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _encodedConfigVersion uint64, _encodedConfig []byte) ([32]byte, error)

	GetAnswer(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error)

	GetBilling(opts *bind.CallOpts) (GetBilling,

		error)

	GetBillingAccessController(opts *bind.CallOpts) (common.Address, error)

	GetLinkToken(opts *bind.CallOpts) (common.Address, error)

	GetRequesterAccessController(opts *bind.CallOpts) (common.Address, error)

	GetRoundData(opts *bind.CallOpts, roundId *big.Int) (GetRoundData,

		error)

	GetTimestamp(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error)

	GetTransmitters(opts *bind.CallOpts) ([]common.Address, error)

	GetValidatorConfig(opts *bind.CallOpts) (GetValidatorConfig,

		error)

	LatestAnswer(opts *bind.CallOpts) (*big.Int, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	LatestRound(opts *bind.CallOpts) (*big.Int, error)

	LatestRoundData(opts *bind.CallOpts) (LatestRoundData,

		error)

	LatestTimestamp(opts *bind.CallOpts) (*big.Int, error)

	LatestTransmissionDetails(opts *bind.CallOpts) (LatestTransmissionDetails,

		error)

	LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error)

	MaxAnswer(opts *bind.CallOpts) (*big.Int, error)

	MinAnswer(opts *bind.CallOpts) (*big.Int, error)

	OracleObservationCount(opts *bind.CallOpts, transmitterAddress common.Address) (uint32, error)

	OwedPayment(opts *bind.CallOpts, transmitterAddress common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	Version(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptPayeeship(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error)

	RequestNewRound(opts *bind.TransactOpts) (*types.Transaction, error)

	SetBilling(opts *bind.TransactOpts, maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error)

	SetBillingAccessController(opts *bind.TransactOpts, _billingAccessController common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	SetLinkToken(opts *bind.TransactOpts, linkToken common.Address, recipient common.Address) (*types.Transaction, error)

	SetPayees(opts *bind.TransactOpts, transmitters []common.Address, payees []common.Address) (*types.Transaction, error)

	SetRequesterAccessController(opts *bind.TransactOpts, requesterAccessController common.Address) (*types.Transaction, error)

	SetValidatorConfig(opts *bind.TransactOpts, newValidator common.Address, newGasLimit uint32) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	TransferPayeeship(opts *bind.TransactOpts, transmitter common.Address, proposed common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	WithdrawFunds(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	WithdrawPayment(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error)

	FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*ExposedOCR2AggregatorAnswerUpdatedIterator, error)

	WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error)

	ParseAnswerUpdated(log types.Log) (*ExposedOCR2AggregatorAnswerUpdated, error)

	FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*ExposedOCR2AggregatorBillingAccessControllerSetIterator, error)

	WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorBillingAccessControllerSet) (event.Subscription, error)

	ParseBillingAccessControllerSet(log types.Log) (*ExposedOCR2AggregatorBillingAccessControllerSet, error)

	FilterBillingSet(opts *bind.FilterOpts) (*ExposedOCR2AggregatorBillingSetIterator, error)

	WatchBillingSet(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorBillingSet) (event.Subscription, error)

	ParseBillingSet(log types.Log) (*ExposedOCR2AggregatorBillingSet, error)

	FilterConfigSet(opts *bind.FilterOpts) (*ExposedOCR2AggregatorConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*ExposedOCR2AggregatorConfigSet, error)

	FilterLinkTokenSet(opts *bind.FilterOpts, oldLinkToken []common.Address, newLinkToken []common.Address) (*ExposedOCR2AggregatorLinkTokenSetIterator, error)

	WatchLinkTokenSet(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorLinkTokenSet, oldLinkToken []common.Address, newLinkToken []common.Address) (event.Subscription, error)

	ParseLinkTokenSet(log types.Log) (*ExposedOCR2AggregatorLinkTokenSet, error)

	FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*ExposedOCR2AggregatorNewRoundIterator, error)

	WatchNewRound(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error)

	ParseNewRound(log types.Log) (*ExposedOCR2AggregatorNewRound, error)

	FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*ExposedOCR2AggregatorNewTransmissionIterator, error)

	WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error)

	ParseNewTransmission(log types.Log) (*ExposedOCR2AggregatorNewTransmission, error)

	FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*ExposedOCR2AggregatorOraclePaidIterator, error)

	WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error)

	ParseOraclePaid(log types.Log) (*ExposedOCR2AggregatorOraclePaid, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExposedOCR2AggregatorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ExposedOCR2AggregatorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExposedOCR2AggregatorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ExposedOCR2AggregatorOwnershipTransferred, error)

	FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*ExposedOCR2AggregatorPayeeshipTransferRequestedIterator, error)

	WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error)

	ParsePayeeshipTransferRequested(log types.Log) (*ExposedOCR2AggregatorPayeeshipTransferRequested, error)

	FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*ExposedOCR2AggregatorPayeeshipTransferredIterator, error)

	WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error)

	ParsePayeeshipTransferred(log types.Log) (*ExposedOCR2AggregatorPayeeshipTransferred, error)

	FilterRequesterAccessControllerSet(opts *bind.FilterOpts) (*ExposedOCR2AggregatorRequesterAccessControllerSetIterator, error)

	WatchRequesterAccessControllerSet(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorRequesterAccessControllerSet) (event.Subscription, error)

	ParseRequesterAccessControllerSet(log types.Log) (*ExposedOCR2AggregatorRequesterAccessControllerSet, error)

	FilterRoundRequested(opts *bind.FilterOpts, requester []common.Address) (*ExposedOCR2AggregatorRoundRequestedIterator, error)

	WatchRoundRequested(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorRoundRequested, requester []common.Address) (event.Subscription, error)

	ParseRoundRequested(log types.Log) (*ExposedOCR2AggregatorRoundRequested, error)

	FilterTransmitted(opts *bind.FilterOpts) (*ExposedOCR2AggregatorTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*ExposedOCR2AggregatorTransmitted, error)

	FilterValidatorConfigSet(opts *bind.FilterOpts, previousValidator []common.Address, currentValidator []common.Address) (*ExposedOCR2AggregatorValidatorConfigSetIterator, error)

	WatchValidatorConfigSet(opts *bind.WatchOpts, sink chan<- *ExposedOCR2AggregatorValidatorConfigSet, previousValidator []common.Address, currentValidator []common.Address) (event.Subscription, error)

	ParseValidatorConfigSet(log types.Log) (*ExposedOCR2AggregatorValidatorConfigSet, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}

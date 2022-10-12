// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ocr2aggregator

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

var OCR2AggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"int192\",\"name\":\"minAnswer_\",\"type\":\"int192\"},{\"internalType\":\"int192\",\"name\":\"maxAnswer_\",\"type\":\"int192\"},{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"billingAccessController\",\"type\":\"address\"},{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"requesterAccessController\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"description_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"BillingAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"BillingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"oldLinkToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"newLinkToken\",\"type\":\"address\"}],\"name\":\"LinkTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int192\",\"name\":\"answer\",\"type\":\"int192\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"observationsTimestamp\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int192[]\",\"name\":\"observations\",\"type\":\"int192[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"observers\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"int192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"int192\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"RequesterAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"}],\"name\":\"RoundRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"previousValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousGasLimit\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"currentValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"currentGasLimit\",\"type\":\"uint32\"}],\"name\":\"ValidatorConfigSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBillingAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinkToken\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequesterAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId_\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorConfig\",\"outputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTransmissionDetails\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"},{\"internalType\":\"int192\",\"name\":\"latestAnswer_\",\"type\":\"int192\"},{\"internalType\":\"uint64\",\"name\":\"latestTimestamp_\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAnswer\",\"outputs\":[{\"internalType\":\"int192\",\"name\":\"\",\"type\":\"int192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAnswer\",\"outputs\":[{\"internalType\":\"int192\",\"name\":\"\",\"type\":\"int192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"oracleObservationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"owedPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestNewRound\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"setBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"}],\"name\":\"setBillingAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"setLinkToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"requesterAccessController\",\"type\":\"address\"}],\"name\":\"setRequesterAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"newValidator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newGasLimit\",\"type\":\"uint32\"}],\"name\":\"setValidatorConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005af838038062005af88339810160408190526200003491620004b2565b3380600081620000615760405162461bcd60e51b8152600401620000589062000653565b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000945762000094816200016a565b5050601180546001600160a01b0319166001600160a01b038a169081179091556040519091506000907f4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a908290a3620000ed84620001e7565b7fff0000000000000000000000000000000000000000000000000000000000000060f883901b1660c05280516200012c906010906020840190620003f4565b5062000138836200025b565b62000145600080620002cc565b50505050601791820b820b604090811b60805290820b90910b901b60a0525062000744565b6001600160a01b038116331415620001965760405162461bcd60e51b815260040162000058906200068a565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6012546001600160a01b0390811690821681146200025757601280546001600160a01b0319166001600160a01b0384161790556040517f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912906200024e908390859062000602565b60405180910390a15b5050565b62000265620003c5565b600f546001600160a01b0390811690821681146200025757600f80546001600160a01b0319166001600160a01b0384161790556040517f27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634906200024e908390859062000602565b620002d6620003c5565b60408051808201909152600e546001600160a01b03808216808452600160a01b90920463ffffffff16602084015284161415806200032457508163ffffffff16816020015163ffffffff1614155b15620003c0576040805180820182526001600160a01b0385811680835263ffffffff86166020938401819052600e80546001600160a01b031916831763ffffffff60a01b1916600160a01b9092029190911790558451928501519351909392909116917fb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca154191620003b791908790620006c1565b60405180910390a35b505050565b6000546001600160a01b03163314620003f25760405162461bcd60e51b815260040162000058906200061c565b565b8280546200040290620006d8565b90600052602060002090601f01602090048101928262000426576000855562000471565b82601f106200044157805160ff191683800117855562000471565b8280016001018555821562000471579182015b828111156200047157825182559160200191906001019062000454565b506200047f92915062000483565b5090565b5b808211156200047f576000815560010162000484565b8051601781900b8114620004ad57600080fd5b919050565b600080600080600080600060e0888a031215620004cd578283fd5b8751620004da816200072b565b96506020620004eb8982016200049a565b9650620004fb60408a016200049a565b955060608901516200050d816200072b565b60808a015190955062000520816200072b565b60a08a015190945060ff8116811462000537578384fd5b60c08a01519093506001600160401b038082111562000554578384fd5b818b0191508b601f83011262000568578384fd5b8151818111156200057d576200057d62000715565b604051601f8201601f1916810185018381118282101715620005a357620005a362000715565b60405281815283820185018e1015620005ba578586fd5b8592505b81831015620005dd5783830185015181840186015291840191620005be565b81831115620005ee57858583830101525b809550505050505092959891949750929550565b6001600160a01b0392831681529116602082015260400190565b60208082526016908201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604082015260600190565b60208082526018908201527f43616e6e6f7420736574206f776e657220746f207a65726f0000000000000000604082015260600190565b60208082526017908201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604082015260600190565b63ffffffff92831681529116602082015260400190565b600281046001821680620006ed57607f821691505b602082108114156200070f57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146200074157600080fd5b50565b60805160401c60a05160401c60c05160f81c61535f6200079960003960006107d2015260008181610d0d01528181611d5e015261303d01526000818161072301528181611d3d0152613010015261535f6000f3fe608060405234801561001057600080fd5b50600436106102d35760003560e01c80639bd2c0b111610186578063d09dc339116100e3578063e76d516811610097578063f2fde38b11610071578063f2fde38b1461058a578063fbffd2c11461059d578063feaf968c146105b0576102d3565b8063e76d51681461055c578063eb45716314610564578063eb5dcd6c14610577576102d3565b8063e3d0e712116100c8578063e3d0e71214610510578063e4902f8214610523578063e5fe457714610543576102d3565b8063d09dc33914610500578063daffc4b514610508576102d3565b8063b1dc65a41161013a578063b633620c1161011f578063b633620c146104d2578063c1075329146104e5578063c4c92b37146104f8576102d3565b8063b1dc65a4146104ac578063b5ab58dc146104bf576102d3565b80639e3ceeab1161016b5780639e3ceeab1461046f578063afcb95d714610482578063b121e14714610499576102d3565b80639bd2c0b1146104465780639c849b301461045c576102d3565b8063668a0f02116102345780638205bf6a116101e85780638da5cb5b116101cd5780638da5cb5b146103f857806398e5b12a1461040d5780639a6fc8f514610422576102d3565b80638205bf6a146103dd5780638ac28d5a146103e5576102d3565b80637284e416116102195780637284e416146103b657806379ba5097146103be57806381ff7048146103c6576102d3565b8063668a0f02146103a657806370da2f67146103ae576102d3565b80634fb174701161028b57806354fd4d501161027057806354fd4d5014610376578063643dc1051461037e578063666cab8d14610391576102d3565b80634fb174701461035957806350d25bcd1461036e576102d3565b806322adbc78116102bc57806322adbc7814610316578063299372681461032b578063313ce56714610344576102d3565b80630eafb25b146102d8578063181f5a7714610301575b600080fd5b6102eb6102e6366004613ed2565b6105b8565b6040516102f89190614415565b60405180910390f35b6103096106ea565b6040516102f891906146a7565b61031e610721565b6040516102f891906145fa565b610333610745565b6040516102f8959493929190614ff2565b61034c6107d0565b6040516102f89190615070565b61036c6103673660046141ac565b6107f4565b005b6102eb610a7d565b6102eb610aa9565b61036c61038c3660046142b6565b610aae565b610399610c93565b6040516102f89190614500565b6102eb610cf5565b61031e610d0b565b610309610d2f565b61036c610db8565b6103ce610e51565b6040516102f893929190614fd1565b6102eb610e6d565b61036c6103f3366004613ed2565b610eb5565b610400610efa565b6040516102f89190614476565b610415610f09565b6040516102f89190615026565b61043561043036600461432d565b611068565b6040516102f895949392919061503d565b61044e611130565b6040516102f89291906145db565b61036c61046a366004613f51565b611174565b61036c61047d366004613ed2565b611326565b61048a6113b6565b6040516102f893929190614513565b61036c6104a7366004613ed2565b6113d1565b61036c6104ba366004614082565b611495565b6102eb6104cd3660046141be565b6118b4565b6102eb6104e03660046141be565b6118ea565b61036c6104f3366004613f26565b61193c565b610400611b99565b6102eb611ba8565b610400611c65565b61036c61051e366004613fba565b611c74565b610536610531366004613ed2565b612415565b6040516102f89190614f23565b61054b6124e2565b6040516102f895949392919061456a565b610400612577565b61036c61057236600461417f565b612586565b61036c610585366004613eee565b6126cb565b61036c610598366004613ed2565b6127bc565b61036c6105ab366004613ed2565b6127cd565b6104356127de565b6001600160a01b03811660009081526002602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046bffffffffffffffffffffffff1691810191909152906106205760009150506106e5565b600b5460208201516000917201000000000000000000000000000000000000900463ffffffff169060069060ff16601f811061066c57634e487b7160e01b600052603260045260246000fd5b600881049190910154600b546106a2926007166004026101000a90910463ffffffff908116916601000000000000900416615242565b63ffffffff166106b29190615151565b6106c090633b9aca00615151565b905081604001516bffffffffffffffffffffffff16816106e091906150cc565b925050505b919050565b60408051808201909152601a81527f4f43523241676772656761746f7220312e302e302d616c706861000000000000602082015290565b7f000000000000000000000000000000000000000000000000000000000000000081565b600b546a0100000000000000000000810463ffffffff908116926e010000000000000000000000000000830482169272010000000000000000000000000000000000008104831692760100000000000000000000000000000000000000000000820416917a01000000000000000000000000000000000000000000000000000090910462ffffff1690565b7f000000000000000000000000000000000000000000000000000000000000000081565b6107fc612877565b6011546001600160a01b0390811690831681141561081a5750610a79565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b038416906370a082319061085f903090600401614476565b60206040518083038186803b15801561087757600080fd5b505afa15801561088b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108af91906141d6565b506108b86128a3565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526000906001600160a01b038316906370a0823190610900903090600401614476565b60206040518083038186803b15801561091857600080fd5b505afa15801561092c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095091906141d6565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000081529091506001600160a01b0383169063a9059cbb9061099a90869085906004016144e7565b602060405180830381600087803b1580156109b457600080fd5b505af11580156109c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ec919061415f565b610a115760405162461bcd60e51b8152600401610a0890614915565b60405180910390fd5b601180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0386811691821790925560405190918416907f4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a90600090a350505b5050565b600b546601000000000000900463ffffffff166000908152600c6020526040902054601790810b900b90565b600681565b6012546001600160a01b0316610ac2610efa565b6001600160a01b0316336001600160a01b03161480610b7657506040517f6b14daf80000000000000000000000000000000000000000000000000000000081526001600160a01b03821690636b14daf890610b26903390600090369060040161448a565b60206040518083038186803b158015610b3e57600080fd5b505afa158015610b52573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b76919061415f565b610b925760405162461bcd60e51b8152600401610a0890614b72565b610b9a6128a3565b85600b600001600a6101000a81548163ffffffff021916908363ffffffff16021790555084600b600001600e6101000a81548163ffffffff021916908363ffffffff16021790555083600b60000160126101000a81548163ffffffff021916908363ffffffff16021790555082600b60000160166101000a81548163ffffffff021916908363ffffffff16021790555081600b600001601a6101000a81548162ffffff021916908362ffffff1602179055507f0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f8686868686604051610c83959493929190614ff2565b60405180910390a1505050505050565b60606005805480602002602001604051908101604052809291908181526020018280548015610ceb57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610ccd575b5050505050905090565b600b546601000000000000900463ffffffff1690565b7f000000000000000000000000000000000000000000000000000000000000000081565b606060108054610d3e90615267565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6a90615267565b8015610ceb5780601f10610d8c57610100808354040283529160200191610ceb565b820191906000526020600020905b815481529060010190602001808311610d9a57509395945050505050565b6001546001600160a01b03163314610de25760405162461bcd60e51b8152600401610a08906146f1565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b600d54600a5463ffffffff808316926401000000009004169192565b600b5463ffffffff660100000000000090910481166000908152600c60205260409020547c010000000000000000000000000000000000000000000000000000000090041690565b6001600160a01b03818116600090815260136020526040902054163314610eee5760405162461bcd60e51b8152600401610a0890614c15565b610ef781612cb6565b50565b6000546001600160a01b031690565b6000610f13610efa565b6001600160a01b0316336001600160a01b03161480610fcb5750600f546040517f6b14daf80000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636b14daf890610f7b903390600090369060040161448a565b60206040518083038186803b158015610f9357600080fd5b505afa158015610fa7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fcb919061415f565b610fe75760405162461bcd60e51b8152600401610a0890614802565b600b54600a54604051610100830464ffffffffff8116936601000000000000900463ffffffff9081169333937f41e3990591fd372502daa15842da15bc7f41c75309ab3ff4f56f1848c178825c936110489360089190911c1690879061454b565b60405180910390a261105b8160016150e4565b63ffffffff169250505090565b60008080808063ffffffff69ffffffffffffffffffff8716111561109a57506000935083925082915081905080611127565b50505063ffffffff8084166000908152600c602090815260409182902082516060810184529054601781810b810b810b8084527801000000000000000000000000000000000000000000000000830487169484018590527c0100000000000000000000000000000000000000000000000000000000909204909516919093018190528695509190920b9250835b91939590929450565b60408051808201909152600e546001600160a01b0381168083527401000000000000000000000000000000000000000090910463ffffffff16602090920182905291565b61117c612877565b82811461119b5760405162461bcd60e51b8152600401610a0890614d5f565b60005b8381101561131f5760008585838181106111c857634e487b7160e01b600052603260045260246000fd5b90506020020160208101906111dd9190613ed2565b9050600084848481811061120157634e487b7160e01b600052603260045260246000fd5b90506020020160208101906112169190613ed2565b6001600160a01b0380841660009081526013602052604090205491925016801580806112535750826001600160a01b0316826001600160a01b0316145b61126f5760405162461bcd60e51b8152600401610a0890614b3b565b6001600160a01b03848116600090815260136020526040902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000168583169081179091559083161461130857826001600160a01b0316826001600160a01b0316856001600160a01b03167f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b360405160405180910390a45b505050508080611317906152a2565b91505061119e565b5050505050565b61132e612877565b600f546001600160a01b039081169082168114610a7957600f80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384161790556040517f27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634906113aa90839085906145c1565b60405180910390a15050565b600a54600b5460009261010090910460081c63ffffffff1690565b6001600160a01b0381811660009081526014602052604090205416331461140a5760405162461bcd60e51b8152600401610a08906146ba565b6001600160a01b0381811660008181526013602090815260408083208054337fffffffffffffffffffffffff000000000000000000000000000000000000000080831682179093556014909452828520805490921690915590519416939092849290917f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b39190a45050565b60005a604080516101008082018352600b5460ff8116835290810464ffffffffff90811660208085018290526601000000000000840463ffffffff908116968601969096526a01000000000000000000008404861660608601526e01000000000000000000000000000084048616608086015272010000000000000000000000000000000000008404861660a0860152760100000000000000000000000000000000000000000000840490951660c08501527a01000000000000000000000000000000000000000000000000000090920462ffffff1660e08401529394509092918c0135918216116115995760405162461bcd60e51b8152600401610a08906147cb565b3360009081526002602052604090205460ff166115c85760405162461bcd60e51b8152600401610a0890614ba7565b600a548b35146115ea5760405162461bcd60e51b8152600401610a0890614cf1565b6115f88a8a8a8a8a8a612ef5565b815161160590600161510c565b60ff1687146116265760405162461bcd60e51b8152600401610a0890614bde565b8685146116455760405162461bcd60e51b8152600401610a0890614e39565b60008a8a60405161165792919061441e565b60405190819003812061166e918e90602001614531565b6040516020818303038152906040528051906020012090506000611690613bc1565b60005b8a8110156118085760006001858a84602081106116c057634e487b7160e01b600052603260045260246000fd5b6116cd91901a601b61510c565b8f8f868181106116ed57634e487b7160e01b600052603260045260246000fd5b905060200201358e8e8781811061171457634e487b7160e01b600052603260045260246000fd5b905060200201356040516000815260200160405260405161173894939291906145a3565b6020604051602081039080840390855afa15801561175a573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101516001600160a01b03811660009081526003602090815290849020838501909452925460ff80821615158085526101009092041693830193909352909550925090506117e15760405162461bcd60e51b8152600401610a0890614c83565b826020015160080260ff166001901b84019350508080611800906152a2565b915050611693565b5081827e0101010101010101010101010101010101010101010101010101010101010116146118495760405162461bcd60e51b8152600401610a0890614983565b50600091506118989050838d836020020135848e8e8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612f6292505050565b90506118a6838286336133e4565b505050505050505050505050565b600063ffffffff8211156118ca575060006106e5565b5063ffffffff166000908152600c6020526040902054601790810b900b90565b600063ffffffff821115611900575060006106e5565b5063ffffffff9081166000908152600c60205260409020547c010000000000000000000000000000000000000000000000000000000090041690565b611944610efa565b6001600160a01b0316336001600160a01b031614806119fc57506012546040517f6b14daf80000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636b14daf8906119ac903390600090369060040161448a565b60206040518083038186803b1580156119c457600080fd5b505afa1580156119d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119fc919061415f565b611a185760405162461bcd60e51b8152600401610a0890614b72565b6000611a22613535565b6011546040517f70a082310000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b03909116906370a0823190611a71903090600401614476565b60206040518083038186803b158015611a8957600080fd5b505afa158015611a9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ac191906141d6565b905081811015611ae35760405162461bcd60e51b8152600401610a0890614c4c565b6011546001600160a01b031663a9059cbb85611b08611b02868661522b565b87613732565b6040518363ffffffff1660e01b8152600401611b259291906144e7565b602060405180830381600087803b158015611b3f57600080fd5b505af1158015611b53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b77919061415f565b611b935760405162461bcd60e51b8152600401610a0890614dcb565b50505050565b6012546001600160a01b031690565b6011546040517f70a0823100000000000000000000000000000000000000000000000000000000815260009182916001600160a01b03909116906370a0823190611bf6903090600401614476565b60206040518083038186803b158015611c0e57600080fd5b505afa158015611c22573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c4691906141d6565b90506000611c52613535565b9050611c5e81836151b7565b9250505090565b600f546001600160a01b031690565b611c7c612877565b601f86511115611c9e5760405162461bcd60e51b8152600401610a0890614870565b8451865114611cbf5760405162461bcd60e51b8152600401610a08906148a7565b8551611ccc85600361518e565b60ff1610611cec5760405162461bcd60e51b8152600401610a0890614d94565b611cf88460ff1661374c565b825115611d175760405162461bcd60e51b8152600401610a0890614b04565b60006040518060c001604052808881526020018781526020018660ff16815260200160017f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000604051602001611d8f9392919061442e565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815290825267ffffffffffffffff8616602083015201839052600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000ff1690559050611e066128a3565b60045460005b81811015611f0157600060048281548110611e3757634e487b7160e01b600052603260045260246000fd5b6000918252602082200154600580546001600160a01b0390921693509084908110611e7257634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b03948516835260038252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905594168252600290529190912080547fffffffffffffffffffffffffffffffffffff00000000000000000000000000001690555080611ef9816152a2565b915050611e0c565b50611f0e60046000613bd8565b611f1a60056000613bd8565b60005b8251518110156121fe576003600084600001518381518110611f4f57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615611f935760405162461bcd60e51b8152600401610a0890614a28565b604080518082019091526001815260ff821660208201528351805160039160009185908110611fd257634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252818101929092526040016000908120835181549484015160ff16610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff9115157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909616959095171693909317909255840151805160029291908490811061208357634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff16156120c75760405162461bcd60e51b8152600401610a089061475f565b60405180606001604052806001151581526020018260ff16815260200160006bffffffffffffffffffffffff16815250600260008560200151848151811061211f57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252818101929092526040908101600020835181549385015194909201517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909316911515919091177fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1661010060ff90941693909302929092177fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff16620100006bffffffffffffffffffffffff90921691909102179055806121f6816152a2565b915050611f1d565b508151805161221591600491602090910190613bf6565b50602080830151805161222c926005920190613bf6565b506040820151600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600d80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292909217808555920481169260019290916000916122ba918591166150e4565b92506101000a81548163ffffffff021916908363ffffffff1602179055506123194630600d60009054906101000a900463ffffffff1663ffffffff1686600001518760200151886040015189606001518a608001518b60a0015161376c565b600a819055600d5484516020860151604080880151606089015160808a015160a08b015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612382988b98919763ffffffff909116969195909490939092909190614f34565b60405180910390a1600b546601000000000000900463ffffffff1660005b8451518110156124085781600682601f81106123cc57634e487b7160e01b600052603260045260246000fd5b600891828204019190066004026101000a81548163ffffffff021916908363ffffffff1602179055508080612400906152a2565b9150506123a0565b5050505050505050505050565b6001600160a01b03811660009081526002602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046bffffffffffffffffffffffff16918101919091529061247d5760009150506106e5565b6006816020015160ff16601f81106124a557634e487b7160e01b600052603260045260246000fd5b600881049190910154600b546124db926007166004026101000a90910463ffffffff908116916601000000000000900416615242565b9392505050565b6000808080803332146125075760405162461bcd60e51b8152600401610a089061494c565b5050600a54600b5463ffffffff6601000000000000820481166000908152600c60205260409020549296610100909204600881901c8216965064ffffffffff169450601783900b93507c010000000000000000000000000000000000000000000000000000000090920490911690565b6011546001600160a01b031690565b61258e612877565b60408051808201909152600e546001600160a01b038082168084527401000000000000000000000000000000000000000090920463ffffffff16602084015284161415806125ec57508163ffffffff16816020015163ffffffff1614155b156126c6576040805180820182526001600160a01b0385811680835263ffffffff86166020938401819052600e80547fffffffffffffffffffffffff00000000000000000000000000000000000000001683177fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000009092029190911790558451928501519351909392909116917fb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541916126bd91908790614fba565b60405180910390a35b505050565b6001600160a01b038281166000908152601360205260409020541633146127045760405162461bcd60e51b8152600401610a0890614728565b336001600160a01b038216141561272d5760405162461bcd60e51b8152600401610a0890614d28565b6001600160a01b03808316600090815260146020526040902080548383167fffffffffffffffffffffffff0000000000000000000000000000000000000000821681179092559091169081146126c6576040516001600160a01b038084169133918616907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e3836790600090a4505050565b6127c4612877565b610ef781613818565b6127d5612877565b610ef7816138aa565b600b546601000000000000900463ffffffff9081166000818152600c602090815260409182902082516060810184529054601781810b810b810b8084527801000000000000000000000000000000000000000000000000830488169484018590527c01000000000000000000000000000000000000000000000000000000009092049096169190930181905292949190930b9291908490565b6000546001600160a01b031633146128a15760405162461bcd60e51b8152600401610a0890614839565b565b601154600b54604080516103e08101918290526001600160a01b0390931692660100000000000090920463ffffffff1691600091600690601f908285855b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116128e1579050505050505090506000600580548060200260200160405190810160405280929190818152602001828054801561297c57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161295e575b5050505050905060005b8151811015612ca8576000600260008484815181106129b557634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160029054906101000a90046bffffffffffffffffffffffff166bffffffffffffffffffffffff169050600060026000858581518110612a2f57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160026101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff16021790555060008483601f8110612aaa57634e487b7160e01b600052603260045260246000fd5b6020020151600b5490870363ffffffff90811692507201000000000000000000000000000000000000909104168102633b9aca000282018015612c9d57600060136000878781518110612b0d57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b03169050886001600160a01b031663a9059cbb82846040518363ffffffff1660e01b8152600401612b799291906144e7565b602060405180830381600087803b158015612b9357600080fd5b505af1158015612ba7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bcb919061415f565b612be75760405162461bcd60e51b8152600401610a0890614dcb565b878786601f8110612c0857634e487b7160e01b600052603260045260246000fd5b602002019063ffffffff16908163ffffffff1681525050886001600160a01b0316816001600160a01b0316878781518110612c5357634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c85604051612c939190614415565b60405180910390a4505b505050600101612986565b5061131f600683601f613c73565b6001600160a01b0381166000908152600260209081526040918290208251606081018452905460ff80821615158084526101008304909116938301939093526201000090046bffffffffffffffffffffffff1692810192909252612d1a5750610ef7565b6000612d25836105b8565b905080156126c6576001600160a01b03808416600090815260136020526040908190205460115491517fa9059cbb00000000000000000000000000000000000000000000000000000000815290831692919091169063a9059cbb90612d9090849086906004016144e7565b602060405180830381600087803b158015612daa57600080fd5b505af1158015612dbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612de2919061415f565b612dfe5760405162461bcd60e51b8152600401610a0890614dcb565b600b60000160069054906101000a900463ffffffff166006846020015160ff16601f8110612e3c57634e487b7160e01b600052603260045260246000fd5b6008810491909101805460079092166004026101000a63ffffffff8181021990931693909216919091029190911790556001600160a01b038481166000818152600260205260409081902080547fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff169055601154905190831692841691907fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c90612ee7908790614415565b60405180910390a450505050565b6000612f02826020615151565b612f0d856020615151565b612f19886101446150cc565b612f2391906150cc565b612f2d91906150cc565b612f389060006150cc565b9050368114612f595760405162461bcd60e51b8152600401610a08906148de565b50505050505050565b600080612f6e83613926565b9050601f8160400151511115612f965760405162461bcd60e51b8152600401610a0890614a96565b604081015151865160ff1610612fbe5760405162461bcd60e51b8152600401610a08906149ba565b64ffffffffff841660208701526040810151805160009190612fe290600290615131565b8151811061300057634e487b7160e01b600052603260045260246000fd5b602002602001015190508060170b7f000000000000000000000000000000000000000000000000000000000000000060170b1315801561306657507f000000000000000000000000000000000000000000000000000000000000000060170b8160170b13155b6130825760405162461bcd60e51b8152600401610a0890614acd565b60408701805190613092826152db565b63ffffffff1663ffffffff168152505060405180606001604052808260170b8152602001836000015163ffffffff1681526020014263ffffffff16815250600c6000896040015163ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a81548177ffffffffffffffffffffffffffffffffffffffffffffffff021916908360170b77ffffffffffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160186101000a81548163ffffffff021916908363ffffffff160217905550604082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555090505086600b60008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548164ffffffffff021916908364ffffffffff16021790555060408201518160000160066101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001600a6101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600e6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160000160126101000a81548163ffffffff021916908363ffffffff16021790555060c08201518160000160166101000a81548163ffffffff021916908363ffffffff16021790555060e082015181600001601a6101000a81548162ffffff021916908362ffffff160217905550905050866040015163ffffffff167fc797025feeeaf2cd924c99e9205acb8ec04d5cad21c41ce637a38fb6dee6016a823385600001518660400151876020015188606001518d8d604051613326989796959493929190614608565b60405180910390a260006001600160a01b0316876040015163ffffffff167f0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac6027184600001516040516133779190614f23565b60405180910390a3866040015163ffffffff168160170b7f0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f426040516133bd9190614415565b60405180910390a36133d687604001518260170b6139ab565b506060015195945050505050565b60008360170b12156133f557611b93565b600061341c633b9aca003a04866080015163ffffffff16876060015163ffffffff16613ae3565b90506010360260005a905060006134458663ffffffff1685858b60e0015162ffffff1686613b09565b90506000670de0b6b3a764000077ffffffffffffffffffffffffffffffffffffffffffffffff891683026001600160a01b03881660009081526002602052604090205460c08c01519290910492506201000090046bffffffffffffffffffffffff9081169163ffffffff16633b9aca0002828401019081168211156134d05750505050505050611b93565b6001600160a01b038816600090815260026020526040902080546bffffffffffffffffffffffff90921662010000027fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff90921691909117905550505050505050505050565b600080600580548060200260200160405190810160405280929190818152602001828054801561358e57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613570575b50508351600b54604080516103e08101918290529697509195660100000000000090910463ffffffff169450600093509150600690601f908285855b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116135ca5790505050505050905060005b8381101561366b578181601f811061363857634e487b7160e01b600052603260045260246000fd5b60200201516136479084615242565b6136579063ffffffff16876150cc565b955080613663816152a2565b915050613610565b50600b54613699907201000000000000000000000000000000000000900463ffffffff16633b9aca00615151565b6136a39086615151565b945060005b8381101561372a57600260008683815181106136d457634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0316825281019190915260400160002054613716906201000090046bffffffffffffffffffffffff16876150cc565b955080613722816152a2565b9150506136a8565b505050505090565b600081831015613743575081613746565b50805b92915050565b80600010610ef75760405162461bcd60e51b8152600401610a0890614cba565b6000808a8a8a8a8a8a8a8a8a60405160200161379099989796959493929190614e70565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150505b9998505050505050505050565b6001600160a01b0381163314156138415760405162461bcd60e51b8152600401610a0890614e02565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6012546001600160a01b039081169082168114610a7957601280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384161790556040517f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912906113aa90839085906145c1565b61392e613d0a565b600080606060008580602001905181019061394991906141ee565b9296509094509250905061395d8683613b3d565b8151604051600090613973908690602001614415565b60408051918152928152825160808101845263ffffffff909716875260208701525084019190915260170b6060830152509050919050565b60408051808201909152600e546001600160a01b0381168083527401000000000000000000000000000000000000000090910463ffffffff1660208301526139f35750610a79565b6000613a00600185615242565b63ffffffff8082166000818152600c6020908152604091829020549087015187519251959650601791820b90910b94613ac79491821693613a4c92909187918c16908b90602401614f08565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fbeed9b5100000000000000000000000000000000000000000000000000000000179052613b85565b61131f5760405162461bcd60e51b8152600401610a0890614a5f565b60008383811015613af657600285850304015b613b008184613732565b95945050505050565b600081861015613b2b5760405162461bcd60e51b8152600401610a0890614796565b50633b9aca0094039190910101020290565b600081516020613b4d9190615151565b613b589060a06150cc565b613b639060006150cc565b9050808351146126c65760405162461bcd60e51b8152600401610a08906149f1565b60005a6113888110613bb95761138881039050846040820482031115613bb9576000808451602086016000888af150600191505b509392505050565b604080518082019091526000808252602082015290565b5080546000825590600052602060002090810190610ef79190613d3b565b828054828255906000526020600020908101928215613c63579160200282015b82811115613c6357825182547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03909116178255602090920191600190910190613c16565b50613c6f929150613d3b565b5090565b600483019183908215613c635791602002820160005b83821115613ccd57835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302613c89565b8015613cfd5782816101000a81549063ffffffff0219169055600401602081600301049283019260010302613ccd565b5050613c6f929150613d3b565b6040518060800160405280600063ffffffff1681526020016060815260200160608152602001600060170b81525090565b5b80821115613c6f5760008155600101613d3c565b60008083601f840112613d61578182fd5b50813567ffffffffffffffff811115613d78578182fd5b6020830191508360208083028501011115613d9257600080fd5b9250929050565b600082601f830112613da9578081fd5b81356020613dbe613db9836150a8565b61507e565b8281528181019085830183850287018401881015613dda578586fd5b855b85811015613e01578135613def8161532b565b84529284019290840190600101613ddc565b5090979650505050505050565b600082601f830112613e1e578081fd5b813567ffffffffffffffff811115613e3857613e38615315565b613e6960207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161507e565b818152846020838601011115613e7d578283fd5b816020850160208301379081016020019190915292915050565b8051601781900b81146106e557600080fd5b803567ffffffffffffffff811681146106e557600080fd5b803560ff811681146106e557600080fd5b600060208284031215613ee3578081fd5b81356124db8161532b565b60008060408385031215613f00578081fd5b8235613f0b8161532b565b91506020830135613f1b8161532b565b809150509250929050565b60008060408385031215613f38578182fd5b8235613f438161532b565b946020939093013593505050565b60008060008060408587031215613f66578182fd5b843567ffffffffffffffff80821115613f7d578384fd5b613f8988838901613d50565b90965094506020870135915080821115613fa1578384fd5b50613fae87828801613d50565b95989497509550505050565b60008060008060008060c08789031215613fd2578384fd5b863567ffffffffffffffff80821115613fe9578586fd5b613ff58a838b01613d99565b9750602089013591508082111561400a578586fd5b6140168a838b01613d99565b965061402460408a01613ec1565b95506060890135915080821115614039578384fd5b6140458a838b01613e0e565b945061405360808a01613ea9565b935060a0890135915080821115614068578283fd5b5061407589828a01613e0e565b9150509295509295509295565b60008060008060008060008060e0898b03121561409d578586fd5b606089018a8111156140ad578687fd5b8998503567ffffffffffffffff808211156140c6578788fd5b818b0191508b601f8301126140d9578788fd5b8135818111156140e7578889fd5b8c60208285010111156140f8578889fd5b6020830199508098505060808b0135915080821115614115578384fd5b6141218c838d01613d50565b909750955060a08b0135915080821115614139578384fd5b506141468b828c01613d50565b999c989b50969995989497949560c00135949350505050565b600060208284031215614170578081fd5b815180151581146124db578182fd5b60008060408385031215614191578182fd5b823561419c8161532b565b91506020830135613f1b81615340565b60008060408385031215613f00578182fd5b6000602082840312156141cf578081fd5b5035919050565b6000602082840312156141e7578081fd5b5051919050565b60008060008060808587031215614203578182fd5b845161420e81615340565b809450506020808601519350604086015167ffffffffffffffff811115614233578384fd5b8601601f81018813614243578384fd5b8051614251613db9826150a8565b81815283810190838501858402850186018c101561426d578788fd5b8794505b838510156142965761428281613e97565b835260019490940193918501918501614271565b5080965050505050506142ab60608601613e97565b905092959194509250565b600080600080600060a086880312156142cd578283fd5b85356142d881615340565b945060208601356142e881615340565b935060408601356142f881615340565b9250606086013561430881615340565b9150608086013562ffffff8116811461431f578182fd5b809150509295509295909350565b60006020828403121561433e578081fd5b813569ffffffffffffffffffff811681146124db578182fd5b6000815180845260208085019450808401835b8381101561438f5781516001600160a01b03168752958201959082019060010161436a565b509495945050505050565b60008151808452815b818110156143bf576020818501810151868301820152016143a3565b818111156143d05782602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60170b9052565b64ffffffffff169052565b90815260200190565b6000828483379101908152919050565b60f89390931b7fff00000000000000000000000000000000000000000000000000000000000000168352601791820b604090811b6001850152910b901b601982015260310190565b6001600160a01b0391909116815260200190565b60006001600160a01b03851682526040602083015282604083015282846060840137818301606090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010192915050565b6001600160a01b03929092168252602082015260400190565b6000602082526124db6020830184614357565b9215158352602083019190915263ffffffff16604082015260600190565b828152608081016060836020840137600081529392505050565b92835263ffffffff91909116602083015260ff16604082015260600190565b94855263ffffffff93909316602085015260ff91909116604084015260170b606083015267ffffffffffffffff16608082015260a00190565b93845260ff9290921660208401526040830152606082015260800190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b0392909216825263ffffffff16602082015260400190565b60179190910b815260200190565b600061010080830160178c810b855260206001600160a01b038d168187015263ffffffff8c1660408701528360608701528293508a5180845261012087019450818c019350855b8181101561466d578451840b8652948201949382019360010161464f565b50505050508281036080840152614684818861439a565b91505061469460a0830186614403565b8360c083015261380b60e083018461440a565b6000602082526124db602083018461439a565b6020808252601f908201527f6f6e6c792070726f706f736564207061796565732063616e2061636365707400604082015260600190565b60208082526016908201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604082015260600190565b6020808252601d908201527f6f6e6c792063757272656e742070617965652063616e20757064617465000000604082015260600190565b6020808252601c908201527f7265706561746564207472616e736d6974746572206164647265737300000000604082015260600190565b6020808252818101527f6c6566744761732063616e6e6f742065786365656420696e697469616c476173604082015260600190565b6020808252600c908201527f7374616c65207265706f72740000000000000000000000000000000000000000604082015260600190565b6020808252601d908201527f4f6e6c79206f776e6572267265717565737465722063616e2063616c6c000000604082015260600190565b60208082526016908201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604082015260600190565b60208082526010908201527f746f6f206d616e79206f7261636c657300000000000000000000000000000000604082015260600190565b60208082526016908201527f6f7261636c65206c656e677468206d69736d6174636800000000000000000000604082015260600190565b60208082526018908201527f63616c6c64617461206c656e677468206d69736d617463680000000000000000604082015260600190565b6020808252601f908201527f7472616e736665722072656d61696e696e672066756e6473206661696c656400604082015260600190565b60208082526014908201527f4f6e6c792063616c6c61626c6520627920454f41000000000000000000000000604082015260600190565b60208082526010908201527f6475706c6963617465207369676e657200000000000000000000000000000000604082015260600190565b6020808252601e908201527f746f6f206665772076616c75657320746f207472757374206d656469616e0000604082015260600190565b60208082526016908201527f7265706f7274206c656e677468206d69736d6174636800000000000000000000604082015260600190565b60208082526017908201527f7265706561746564207369676e65722061646472657373000000000000000000604082015260600190565b60208082526010908201527f696e73756666696369656e742067617300000000000000000000000000000000604082015260600190565b6020808252601e908201527f6e756d206f62736572766174696f6e73206f7574206f6620626f756e64730000604082015260600190565b6020808252601e908201527f6d656469616e206973206f7574206f66206d696e2d6d61782072616e67650000604082015260600190565b6020808252601b908201527f6f6e636861696e436f6e666967206d75737420626520656d7074790000000000604082015260600190565b60208082526011908201527f706179656520616c726561647920736574000000000000000000000000000000604082015260600190565b6020808252818101527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c604082015260600190565b60208082526018908201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604082015260600190565b6020808252601a908201527f77726f6e67206e756d626572206f66207369676e617475726573000000000000604082015260600190565b60208082526017908201527f4f6e6c792070617965652063616e207769746864726177000000000000000000604082015260600190565b60208082526014908201527f696e73756666696369656e742062616c616e6365000000000000000000000000604082015260600190565b6020808252600f908201527f7369676e6174757265206572726f720000000000000000000000000000000000604082015260600190565b60208082526012908201527f66206d75737420626520706f7369746976650000000000000000000000000000604082015260600190565b60208082526015908201527f636f6e666967446967657374206d69736d617463680000000000000000000000604082015260600190565b60208082526017908201527f63616e6e6f74207472616e7366657220746f2073656c66000000000000000000604082015260600190565b6020808252818101527f7472616e736d6974746572732e73697a6520213d207061796565732e73697a65604082015260600190565b60208082526018908201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604082015260600190565b60208082526012908201527f696e73756666696369656e742066756e64730000000000000000000000000000604082015260600190565b60208082526017908201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604082015260600190565b6020808252601e908201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604082015260600190565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b166040850152816060850152614eaa8285018b614357565b91508382036080850152614ebe828a614357565b915060ff881660a085015283820360c0850152614edb828861439a565b90861660e08501528381036101008501529050614ef8818561439a565b9c9b505050505050505050505050565b93845260208401929092526040830152606082015260800190565b63ffffffff91909116815260200190565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614f648184018a614357565b90508281036080840152614f788189614357565b905060ff871660a084015282810360c0840152614f95818761439a565b905067ffffffffffffffff851660e0840152828103610100840152614ef8818561439a565b63ffffffff92831681529116602082015260400190565b63ffffffff9384168152919092166020820152604081019190915260600190565b63ffffffff958616815293851660208501529184166040840152909216606082015262ffffff909116608082015260a00190565b69ffffffffffffffffffff91909116815260200190565b69ffffffffffffffffffff9586168152602081019490945260408401929092526060830152909116608082015260a00190565b60ff91909116815260200190565b60405181810167ffffffffffffffff811182821017156150a0576150a0615315565b604052919050565b600067ffffffffffffffff8211156150c2576150c2615315565b5060209081020190565b600082198211156150df576150df6152ff565b500190565b600063ffffffff808316818516808303821115615103576151036152ff565b01949350505050565b600060ff821660ff84168060ff03821115615129576151296152ff565b019392505050565b60008261514c57634e487b7160e01b81526012600452602481fd5b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615615189576151896152ff565b500290565b600060ff821660ff84168160ff04811182151516156151af576151af6152ff565b029392505050565b6000808312837f8000000000000000000000000000000000000000000000000000000000000000018312811516156151f1576151f16152ff565b837f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018313811615615225576152256152ff565b50500390565b60008282101561523d5761523d6152ff565b500390565b600063ffffffff8381169083168181101561525f5761525f6152ff565b039392505050565b60028104600182168061527b57607f821691505b6020821081141561529c57634e487b7160e01b600052602260045260246000fd5b50919050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156152d4576152d46152ff565b5060010190565b600063ffffffff808316818114156152f5576152f56152ff565b6001019392505050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114610ef757600080fd5b63ffffffff81168114610ef757600080fdfea164736f6c6343000800000a",
}

var OCR2AggregatorABI = OCR2AggregatorMetaData.ABI

var OCR2AggregatorBin = OCR2AggregatorMetaData.Bin

func DeployOCR2Aggregator(auth *bind.TransactOpts, backend bind.ContractBackend, link common.Address, minAnswer_ *big.Int, maxAnswer_ *big.Int, billingAccessController common.Address, requesterAccessController common.Address, decimals_ uint8, description_ string) (common.Address, *types.Transaction, *OCR2Aggregator, error) {
	parsed, err := OCR2AggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OCR2AggregatorBin), backend, link, minAnswer_, maxAnswer_, billingAccessController, requesterAccessController, decimals_, description_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OCR2Aggregator{OCR2AggregatorCaller: OCR2AggregatorCaller{contract: contract}, OCR2AggregatorTransactor: OCR2AggregatorTransactor{contract: contract}, OCR2AggregatorFilterer: OCR2AggregatorFilterer{contract: contract}}, nil
}

type OCR2Aggregator struct {
	address common.Address
	abi     abi.ABI
	OCR2AggregatorCaller
	OCR2AggregatorTransactor
	OCR2AggregatorFilterer
}

type OCR2AggregatorCaller struct {
	contract *bind.BoundContract
}

type OCR2AggregatorTransactor struct {
	contract *bind.BoundContract
}

type OCR2AggregatorFilterer struct {
	contract *bind.BoundContract
}

type OCR2AggregatorSession struct {
	Contract     *OCR2Aggregator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OCR2AggregatorCallerSession struct {
	Contract *OCR2AggregatorCaller
	CallOpts bind.CallOpts
}

type OCR2AggregatorTransactorSession struct {
	Contract     *OCR2AggregatorTransactor
	TransactOpts bind.TransactOpts
}

type OCR2AggregatorRaw struct {
	Contract *OCR2Aggregator
}

type OCR2AggregatorCallerRaw struct {
	Contract *OCR2AggregatorCaller
}

type OCR2AggregatorTransactorRaw struct {
	Contract *OCR2AggregatorTransactor
}

func NewOCR2Aggregator(address common.Address, backend bind.ContractBackend) (*OCR2Aggregator, error) {
	abi, err := abi.JSON(strings.NewReader(OCR2AggregatorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOCR2Aggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OCR2Aggregator{address: address, abi: abi, OCR2AggregatorCaller: OCR2AggregatorCaller{contract: contract}, OCR2AggregatorTransactor: OCR2AggregatorTransactor{contract: contract}, OCR2AggregatorFilterer: OCR2AggregatorFilterer{contract: contract}}, nil
}

func NewOCR2AggregatorCaller(address common.Address, caller bind.ContractCaller) (*OCR2AggregatorCaller, error) {
	contract, err := bindOCR2Aggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorCaller{contract: contract}, nil
}

func NewOCR2AggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*OCR2AggregatorTransactor, error) {
	contract, err := bindOCR2Aggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorTransactor{contract: contract}, nil
}

func NewOCR2AggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*OCR2AggregatorFilterer, error) {
	contract, err := bindOCR2Aggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorFilterer{contract: contract}, nil
}

func bindOCR2Aggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OCR2AggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_OCR2Aggregator *OCR2AggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OCR2Aggregator.Contract.OCR2AggregatorCaller.contract.Call(opts, result, method, params...)
}

func (_OCR2Aggregator *OCR2AggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.OCR2AggregatorTransactor.contract.Transfer(opts)
}

func (_OCR2Aggregator *OCR2AggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.OCR2AggregatorTransactor.contract.Transact(opts, method, params...)
}

func (_OCR2Aggregator *OCR2AggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OCR2Aggregator.Contract.contract.Call(opts, result, method, params...)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.contract.Transfer(opts)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.contract.Transact(opts, method, params...)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) Decimals() (uint8, error) {
	return _OCR2Aggregator.Contract.Decimals(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) Decimals() (uint8, error) {
	return _OCR2Aggregator.Contract.Decimals(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) Description() (string, error) {
	return _OCR2Aggregator.Contract.Description(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) Description() (string, error) {
	return _OCR2Aggregator.Contract.Description(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) GetAnswer(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "getAnswer", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _OCR2Aggregator.Contract.GetAnswer(&_OCR2Aggregator.CallOpts, roundId)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _OCR2Aggregator.Contract.GetAnswer(&_OCR2Aggregator.CallOpts, roundId)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) GetBilling(opts *bind.CallOpts) (GetBilling,

	error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "getBilling")

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

func (_OCR2Aggregator *OCR2AggregatorSession) GetBilling() (GetBilling,

	error) {
	return _OCR2Aggregator.Contract.GetBilling(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) GetBilling() (GetBilling,

	error) {
	return _OCR2Aggregator.Contract.GetBilling(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) GetBillingAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "getBillingAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) GetBillingAccessController() (common.Address, error) {
	return _OCR2Aggregator.Contract.GetBillingAccessController(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) GetBillingAccessController() (common.Address, error) {
	return _OCR2Aggregator.Contract.GetBillingAccessController(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) GetLinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "getLinkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) GetLinkToken() (common.Address, error) {
	return _OCR2Aggregator.Contract.GetLinkToken(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) GetLinkToken() (common.Address, error) {
	return _OCR2Aggregator.Contract.GetLinkToken(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) GetRequesterAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "getRequesterAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) GetRequesterAccessController() (common.Address, error) {
	return _OCR2Aggregator.Contract.GetRequesterAccessController(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) GetRequesterAccessController() (common.Address, error) {
	return _OCR2Aggregator.Contract.GetRequesterAccessController(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) GetRoundData(opts *bind.CallOpts, roundId *big.Int) (GetRoundData,

	error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "getRoundData", roundId)

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

func (_OCR2Aggregator *OCR2AggregatorSession) GetRoundData(roundId *big.Int) (GetRoundData,

	error) {
	return _OCR2Aggregator.Contract.GetRoundData(&_OCR2Aggregator.CallOpts, roundId)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) GetRoundData(roundId *big.Int) (GetRoundData,

	error) {
	return _OCR2Aggregator.Contract.GetRoundData(&_OCR2Aggregator.CallOpts, roundId)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) GetTimestamp(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "getTimestamp", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _OCR2Aggregator.Contract.GetTimestamp(&_OCR2Aggregator.CallOpts, roundId)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _OCR2Aggregator.Contract.GetTimestamp(&_OCR2Aggregator.CallOpts, roundId)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) GetTransmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "getTransmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) GetTransmitters() ([]common.Address, error) {
	return _OCR2Aggregator.Contract.GetTransmitters(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) GetTransmitters() ([]common.Address, error) {
	return _OCR2Aggregator.Contract.GetTransmitters(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) GetValidatorConfig(opts *bind.CallOpts) (GetValidatorConfig,

	error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "getValidatorConfig")

	outstruct := new(GetValidatorConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.GasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) GetValidatorConfig() (GetValidatorConfig,

	error) {
	return _OCR2Aggregator.Contract.GetValidatorConfig(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) GetValidatorConfig() (GetValidatorConfig,

	error) {
	return _OCR2Aggregator.Contract.GetValidatorConfig(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) LatestAnswer() (*big.Int, error) {
	return _OCR2Aggregator.Contract.LatestAnswer(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) LatestAnswer() (*big.Int, error) {
	return _OCR2Aggregator.Contract.LatestAnswer(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _OCR2Aggregator.Contract.LatestConfigDetails(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _OCR2Aggregator.Contract.LatestConfigDetails(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _OCR2Aggregator.Contract.LatestConfigDigestAndEpoch(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _OCR2Aggregator.Contract.LatestConfigDigestAndEpoch(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) LatestRound() (*big.Int, error) {
	return _OCR2Aggregator.Contract.LatestRound(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) LatestRound() (*big.Int, error) {
	return _OCR2Aggregator.Contract.LatestRound(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) LatestRoundData(opts *bind.CallOpts) (LatestRoundData,

	error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "latestRoundData")

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

func (_OCR2Aggregator *OCR2AggregatorSession) LatestRoundData() (LatestRoundData,

	error) {
	return _OCR2Aggregator.Contract.LatestRoundData(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) LatestRoundData() (LatestRoundData,

	error) {
	return _OCR2Aggregator.Contract.LatestRoundData(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) LatestTimestamp() (*big.Int, error) {
	return _OCR2Aggregator.Contract.LatestTimestamp(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) LatestTimestamp() (*big.Int, error) {
	return _OCR2Aggregator.Contract.LatestTimestamp(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) LatestTransmissionDetails(opts *bind.CallOpts) (LatestTransmissionDetails,

	error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "latestTransmissionDetails")

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

func (_OCR2Aggregator *OCR2AggregatorSession) LatestTransmissionDetails() (LatestTransmissionDetails,

	error) {
	return _OCR2Aggregator.Contract.LatestTransmissionDetails(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) LatestTransmissionDetails() (LatestTransmissionDetails,

	error) {
	return _OCR2Aggregator.Contract.LatestTransmissionDetails(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) LinkAvailableForPayment() (*big.Int, error) {
	return _OCR2Aggregator.Contract.LinkAvailableForPayment(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _OCR2Aggregator.Contract.LinkAvailableForPayment(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) MaxAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "maxAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) MaxAnswer() (*big.Int, error) {
	return _OCR2Aggregator.Contract.MaxAnswer(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) MaxAnswer() (*big.Int, error) {
	return _OCR2Aggregator.Contract.MaxAnswer(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) MinAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "minAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) MinAnswer() (*big.Int, error) {
	return _OCR2Aggregator.Contract.MinAnswer(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) MinAnswer() (*big.Int, error) {
	return _OCR2Aggregator.Contract.MinAnswer(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) OracleObservationCount(opts *bind.CallOpts, transmitterAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "oracleObservationCount", transmitterAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _OCR2Aggregator.Contract.OracleObservationCount(&_OCR2Aggregator.CallOpts, transmitterAddress)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _OCR2Aggregator.Contract.OracleObservationCount(&_OCR2Aggregator.CallOpts, transmitterAddress)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) OwedPayment(opts *bind.CallOpts, transmitterAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "owedPayment", transmitterAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _OCR2Aggregator.Contract.OwedPayment(&_OCR2Aggregator.CallOpts, transmitterAddress)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _OCR2Aggregator.Contract.OwedPayment(&_OCR2Aggregator.CallOpts, transmitterAddress)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) Owner() (common.Address, error) {
	return _OCR2Aggregator.Contract.Owner(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) Owner() (common.Address, error) {
	return _OCR2Aggregator.Contract.Owner(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) TypeAndVersion() (string, error) {
	return _OCR2Aggregator.Contract.TypeAndVersion(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) TypeAndVersion() (string, error) {
	return _OCR2Aggregator.Contract.TypeAndVersion(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OCR2Aggregator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OCR2Aggregator *OCR2AggregatorSession) Version() (*big.Int, error) {
	return _OCR2Aggregator.Contract.Version(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorCallerSession) Version() (*big.Int, error) {
	return _OCR2Aggregator.Contract.Version(&_OCR2Aggregator.CallOpts)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "acceptOwnership")
}

func (_OCR2Aggregator *OCR2AggregatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.AcceptOwnership(&_OCR2Aggregator.TransactOpts)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.AcceptOwnership(&_OCR2Aggregator.TransactOpts)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) AcceptPayeeship(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "acceptPayeeship", transmitter)
}

func (_OCR2Aggregator *OCR2AggregatorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.AcceptPayeeship(&_OCR2Aggregator.TransactOpts, transmitter)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.AcceptPayeeship(&_OCR2Aggregator.TransactOpts, transmitter)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) RequestNewRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "requestNewRound")
}

func (_OCR2Aggregator *OCR2AggregatorSession) RequestNewRound() (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.RequestNewRound(&_OCR2Aggregator.TransactOpts)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) RequestNewRound() (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.RequestNewRound(&_OCR2Aggregator.TransactOpts)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) SetBilling(opts *bind.TransactOpts, maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "setBilling", maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_OCR2Aggregator *OCR2AggregatorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.SetBilling(&_OCR2Aggregator.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.SetBilling(&_OCR2Aggregator.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) SetBillingAccessController(opts *bind.TransactOpts, _billingAccessController common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "setBillingAccessController", _billingAccessController)
}

func (_OCR2Aggregator *OCR2AggregatorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.SetBillingAccessController(&_OCR2Aggregator.TransactOpts, _billingAccessController)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.SetBillingAccessController(&_OCR2Aggregator.TransactOpts, _billingAccessController)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "setConfig", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_OCR2Aggregator *OCR2AggregatorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.SetConfig(&_OCR2Aggregator.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.SetConfig(&_OCR2Aggregator.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) SetLinkToken(opts *bind.TransactOpts, linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "setLinkToken", linkToken, recipient)
}

func (_OCR2Aggregator *OCR2AggregatorSession) SetLinkToken(linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.SetLinkToken(&_OCR2Aggregator.TransactOpts, linkToken, recipient)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) SetLinkToken(linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.SetLinkToken(&_OCR2Aggregator.TransactOpts, linkToken, recipient)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) SetPayees(opts *bind.TransactOpts, transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "setPayees", transmitters, payees)
}

func (_OCR2Aggregator *OCR2AggregatorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.SetPayees(&_OCR2Aggregator.TransactOpts, transmitters, payees)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.SetPayees(&_OCR2Aggregator.TransactOpts, transmitters, payees)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) SetRequesterAccessController(opts *bind.TransactOpts, requesterAccessController common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "setRequesterAccessController", requesterAccessController)
}

func (_OCR2Aggregator *OCR2AggregatorSession) SetRequesterAccessController(requesterAccessController common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.SetRequesterAccessController(&_OCR2Aggregator.TransactOpts, requesterAccessController)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) SetRequesterAccessController(requesterAccessController common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.SetRequesterAccessController(&_OCR2Aggregator.TransactOpts, requesterAccessController)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) SetValidatorConfig(opts *bind.TransactOpts, newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "setValidatorConfig", newValidator, newGasLimit)
}

func (_OCR2Aggregator *OCR2AggregatorSession) SetValidatorConfig(newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.SetValidatorConfig(&_OCR2Aggregator.TransactOpts, newValidator, newGasLimit)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) SetValidatorConfig(newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.SetValidatorConfig(&_OCR2Aggregator.TransactOpts, newValidator, newGasLimit)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "transferOwnership", to)
}

func (_OCR2Aggregator *OCR2AggregatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.TransferOwnership(&_OCR2Aggregator.TransactOpts, to)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.TransferOwnership(&_OCR2Aggregator.TransactOpts, to)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) TransferPayeeship(opts *bind.TransactOpts, transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "transferPayeeship", transmitter, proposed)
}

func (_OCR2Aggregator *OCR2AggregatorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.TransferPayeeship(&_OCR2Aggregator.TransactOpts, transmitter, proposed)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.TransferPayeeship(&_OCR2Aggregator.TransactOpts, transmitter, proposed)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_OCR2Aggregator *OCR2AggregatorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.Transmit(&_OCR2Aggregator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.Transmit(&_OCR2Aggregator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) WithdrawFunds(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "withdrawFunds", recipient, amount)
}

func (_OCR2Aggregator *OCR2AggregatorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.WithdrawFunds(&_OCR2Aggregator.TransactOpts, recipient, amount)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.WithdrawFunds(&_OCR2Aggregator.TransactOpts, recipient, amount)
}

func (_OCR2Aggregator *OCR2AggregatorTransactor) WithdrawPayment(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.contract.Transact(opts, "withdrawPayment", transmitter)
}

func (_OCR2Aggregator *OCR2AggregatorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.WithdrawPayment(&_OCR2Aggregator.TransactOpts, transmitter)
}

func (_OCR2Aggregator *OCR2AggregatorTransactorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _OCR2Aggregator.Contract.WithdrawPayment(&_OCR2Aggregator.TransactOpts, transmitter)
}

type OCR2AggregatorAnswerUpdatedIterator struct {
	Event *OCR2AggregatorAnswerUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorAnswerUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorAnswerUpdated)
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
		it.Event = new(OCR2AggregatorAnswerUpdated)
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

func (it *OCR2AggregatorAnswerUpdatedIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*OCR2AggregatorAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorAnswerUpdatedIterator{contract: _OCR2Aggregator.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorAnswerUpdated)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParseAnswerUpdated(log types.Log) (*OCR2AggregatorAnswerUpdated, error) {
	event := new(OCR2AggregatorAnswerUpdated)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorBillingAccessControllerSetIterator struct {
	Event *OCR2AggregatorBillingAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorBillingAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorBillingAccessControllerSet)
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
		it.Event = new(OCR2AggregatorBillingAccessControllerSet)
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

func (it *OCR2AggregatorBillingAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorBillingAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorBillingAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*OCR2AggregatorBillingAccessControllerSetIterator, error) {

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorBillingAccessControllerSetIterator{contract: _OCR2Aggregator.contract, event: "BillingAccessControllerSet", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorBillingAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorBillingAccessControllerSet)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParseBillingAccessControllerSet(log types.Log) (*OCR2AggregatorBillingAccessControllerSet, error) {
	event := new(OCR2AggregatorBillingAccessControllerSet)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorBillingSetIterator struct {
	Event *OCR2AggregatorBillingSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorBillingSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorBillingSet)
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
		it.Event = new(OCR2AggregatorBillingSet)
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

func (it *OCR2AggregatorBillingSetIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorBillingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorBillingSet struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
	Raw                       types.Log
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterBillingSet(opts *bind.FilterOpts) (*OCR2AggregatorBillingSetIterator, error) {

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorBillingSetIterator{contract: _OCR2Aggregator.contract, event: "BillingSet", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchBillingSet(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorBillingSet) (event.Subscription, error) {

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorBillingSet)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "BillingSet", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParseBillingSet(log types.Log) (*OCR2AggregatorBillingSet, error) {
	event := new(OCR2AggregatorBillingSet)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "BillingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorConfigSetIterator struct {
	Event *OCR2AggregatorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorConfigSet)
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
		it.Event = new(OCR2AggregatorConfigSet)
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

func (it *OCR2AggregatorConfigSetIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorConfigSet struct {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OCR2AggregatorConfigSetIterator, error) {

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorConfigSetIterator{contract: _OCR2Aggregator.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorConfigSet) (event.Subscription, error) {

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorConfigSet)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParseConfigSet(log types.Log) (*OCR2AggregatorConfigSet, error) {
	event := new(OCR2AggregatorConfigSet)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorLinkTokenSetIterator struct {
	Event *OCR2AggregatorLinkTokenSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorLinkTokenSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorLinkTokenSet)
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
		it.Event = new(OCR2AggregatorLinkTokenSet)
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

func (it *OCR2AggregatorLinkTokenSetIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorLinkTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorLinkTokenSet struct {
	OldLinkToken common.Address
	NewLinkToken common.Address
	Raw          types.Log
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterLinkTokenSet(opts *bind.FilterOpts, oldLinkToken []common.Address, newLinkToken []common.Address) (*OCR2AggregatorLinkTokenSetIterator, error) {

	var oldLinkTokenRule []interface{}
	for _, oldLinkTokenItem := range oldLinkToken {
		oldLinkTokenRule = append(oldLinkTokenRule, oldLinkTokenItem)
	}
	var newLinkTokenRule []interface{}
	for _, newLinkTokenItem := range newLinkToken {
		newLinkTokenRule = append(newLinkTokenRule, newLinkTokenItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "LinkTokenSet", oldLinkTokenRule, newLinkTokenRule)
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorLinkTokenSetIterator{contract: _OCR2Aggregator.contract, event: "LinkTokenSet", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchLinkTokenSet(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorLinkTokenSet, oldLinkToken []common.Address, newLinkToken []common.Address) (event.Subscription, error) {

	var oldLinkTokenRule []interface{}
	for _, oldLinkTokenItem := range oldLinkToken {
		oldLinkTokenRule = append(oldLinkTokenRule, oldLinkTokenItem)
	}
	var newLinkTokenRule []interface{}
	for _, newLinkTokenItem := range newLinkToken {
		newLinkTokenRule = append(newLinkTokenRule, newLinkTokenItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "LinkTokenSet", oldLinkTokenRule, newLinkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorLinkTokenSet)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "LinkTokenSet", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParseLinkTokenSet(log types.Log) (*OCR2AggregatorLinkTokenSet, error) {
	event := new(OCR2AggregatorLinkTokenSet)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "LinkTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorNewRoundIterator struct {
	Event *OCR2AggregatorNewRound

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorNewRoundIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorNewRound)
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
		it.Event = new(OCR2AggregatorNewRound)
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

func (it *OCR2AggregatorNewRoundIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*OCR2AggregatorNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorNewRoundIterator{contract: _OCR2Aggregator.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorNewRound)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParseNewRound(log types.Log) (*OCR2AggregatorNewRound, error) {
	event := new(OCR2AggregatorNewRound)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorNewTransmissionIterator struct {
	Event *OCR2AggregatorNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorNewTransmission)
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
		it.Event = new(OCR2AggregatorNewTransmission)
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

func (it *OCR2AggregatorNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorNewTransmission struct {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*OCR2AggregatorNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorNewTransmissionIterator{contract: _OCR2Aggregator.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorNewTransmission)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParseNewTransmission(log types.Log) (*OCR2AggregatorNewTransmission, error) {
	event := new(OCR2AggregatorNewTransmission)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorOraclePaidIterator struct {
	Event *OCR2AggregatorOraclePaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorOraclePaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorOraclePaid)
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
		it.Event = new(OCR2AggregatorOraclePaid)
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

func (it *OCR2AggregatorOraclePaidIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorOraclePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorOraclePaid struct {
	Transmitter common.Address
	Payee       common.Address
	Amount      *big.Int
	LinkToken   common.Address
	Raw         types.Log
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*OCR2AggregatorOraclePaidIterator, error) {

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

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorOraclePaidIterator{contract: _OCR2Aggregator.contract, event: "OraclePaid", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorOraclePaid)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParseOraclePaid(log types.Log) (*OCR2AggregatorOraclePaid, error) {
	event := new(OCR2AggregatorOraclePaid)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorOwnershipTransferRequestedIterator struct {
	Event *OCR2AggregatorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorOwnershipTransferRequested)
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
		it.Event = new(OCR2AggregatorOwnershipTransferRequested)
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

func (it *OCR2AggregatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OCR2AggregatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorOwnershipTransferRequestedIterator{contract: _OCR2Aggregator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorOwnershipTransferRequested)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*OCR2AggregatorOwnershipTransferRequested, error) {
	event := new(OCR2AggregatorOwnershipTransferRequested)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorOwnershipTransferredIterator struct {
	Event *OCR2AggregatorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorOwnershipTransferred)
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
		it.Event = new(OCR2AggregatorOwnershipTransferred)
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

func (it *OCR2AggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OCR2AggregatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorOwnershipTransferredIterator{contract: _OCR2Aggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorOwnershipTransferred)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*OCR2AggregatorOwnershipTransferred, error) {
	event := new(OCR2AggregatorOwnershipTransferred)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorPayeeshipTransferRequestedIterator struct {
	Event *OCR2AggregatorPayeeshipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorPayeeshipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorPayeeshipTransferRequested)
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
		it.Event = new(OCR2AggregatorPayeeshipTransferRequested)
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

func (it *OCR2AggregatorPayeeshipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorPayeeshipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorPayeeshipTransferRequested struct {
	Transmitter common.Address
	Current     common.Address
	Proposed    common.Address
	Raw         types.Log
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*OCR2AggregatorPayeeshipTransferRequestedIterator, error) {

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

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorPayeeshipTransferRequestedIterator{contract: _OCR2Aggregator.contract, event: "PayeeshipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorPayeeshipTransferRequested)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParsePayeeshipTransferRequested(log types.Log) (*OCR2AggregatorPayeeshipTransferRequested, error) {
	event := new(OCR2AggregatorPayeeshipTransferRequested)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorPayeeshipTransferredIterator struct {
	Event *OCR2AggregatorPayeeshipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorPayeeshipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorPayeeshipTransferred)
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
		it.Event = new(OCR2AggregatorPayeeshipTransferred)
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

func (it *OCR2AggregatorPayeeshipTransferredIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorPayeeshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorPayeeshipTransferred struct {
	Transmitter common.Address
	Previous    common.Address
	Current     common.Address
	Raw         types.Log
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*OCR2AggregatorPayeeshipTransferredIterator, error) {

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

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorPayeeshipTransferredIterator{contract: _OCR2Aggregator.contract, event: "PayeeshipTransferred", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorPayeeshipTransferred)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParsePayeeshipTransferred(log types.Log) (*OCR2AggregatorPayeeshipTransferred, error) {
	event := new(OCR2AggregatorPayeeshipTransferred)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorRequesterAccessControllerSetIterator struct {
	Event *OCR2AggregatorRequesterAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorRequesterAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorRequesterAccessControllerSet)
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
		it.Event = new(OCR2AggregatorRequesterAccessControllerSet)
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

func (it *OCR2AggregatorRequesterAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorRequesterAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorRequesterAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterRequesterAccessControllerSet(opts *bind.FilterOpts) (*OCR2AggregatorRequesterAccessControllerSetIterator, error) {

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "RequesterAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorRequesterAccessControllerSetIterator{contract: _OCR2Aggregator.contract, event: "RequesterAccessControllerSet", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchRequesterAccessControllerSet(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorRequesterAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "RequesterAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorRequesterAccessControllerSet)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "RequesterAccessControllerSet", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParseRequesterAccessControllerSet(log types.Log) (*OCR2AggregatorRequesterAccessControllerSet, error) {
	event := new(OCR2AggregatorRequesterAccessControllerSet)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "RequesterAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorRoundRequestedIterator struct {
	Event *OCR2AggregatorRoundRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorRoundRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorRoundRequested)
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
		it.Event = new(OCR2AggregatorRoundRequested)
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

func (it *OCR2AggregatorRoundRequestedIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorRoundRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorRoundRequested struct {
	Requester    common.Address
	ConfigDigest [32]byte
	Epoch        uint32
	Round        uint8
	Raw          types.Log
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterRoundRequested(opts *bind.FilterOpts, requester []common.Address) (*OCR2AggregatorRoundRequestedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "RoundRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorRoundRequestedIterator{contract: _OCR2Aggregator.contract, event: "RoundRequested", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchRoundRequested(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorRoundRequested, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "RoundRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorRoundRequested)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "RoundRequested", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParseRoundRequested(log types.Log) (*OCR2AggregatorRoundRequested, error) {
	event := new(OCR2AggregatorRoundRequested)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "RoundRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorTransmittedIterator struct {
	Event *OCR2AggregatorTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorTransmitted)
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
		it.Event = new(OCR2AggregatorTransmitted)
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

func (it *OCR2AggregatorTransmittedIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterTransmitted(opts *bind.FilterOpts) (*OCR2AggregatorTransmittedIterator, error) {

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorTransmittedIterator{contract: _OCR2Aggregator.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorTransmitted) (event.Subscription, error) {

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorTransmitted)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParseTransmitted(log types.Log) (*OCR2AggregatorTransmitted, error) {
	event := new(OCR2AggregatorTransmitted)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AggregatorValidatorConfigSetIterator struct {
	Event *OCR2AggregatorValidatorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AggregatorValidatorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AggregatorValidatorConfigSet)
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
		it.Event = new(OCR2AggregatorValidatorConfigSet)
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

func (it *OCR2AggregatorValidatorConfigSetIterator) Error() error {
	return it.fail
}

func (it *OCR2AggregatorValidatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AggregatorValidatorConfigSet struct {
	PreviousValidator common.Address
	PreviousGasLimit  uint32
	CurrentValidator  common.Address
	CurrentGasLimit   uint32
	Raw               types.Log
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) FilterValidatorConfigSet(opts *bind.FilterOpts, previousValidator []common.Address, currentValidator []common.Address) (*OCR2AggregatorValidatorConfigSetIterator, error) {

	var previousValidatorRule []interface{}
	for _, previousValidatorItem := range previousValidator {
		previousValidatorRule = append(previousValidatorRule, previousValidatorItem)
	}

	var currentValidatorRule []interface{}
	for _, currentValidatorItem := range currentValidator {
		currentValidatorRule = append(currentValidatorRule, currentValidatorItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.FilterLogs(opts, "ValidatorConfigSet", previousValidatorRule, currentValidatorRule)
	if err != nil {
		return nil, err
	}
	return &OCR2AggregatorValidatorConfigSetIterator{contract: _OCR2Aggregator.contract, event: "ValidatorConfigSet", logs: logs, sub: sub}, nil
}

func (_OCR2Aggregator *OCR2AggregatorFilterer) WatchValidatorConfigSet(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorValidatorConfigSet, previousValidator []common.Address, currentValidator []common.Address) (event.Subscription, error) {

	var previousValidatorRule []interface{}
	for _, previousValidatorItem := range previousValidator {
		previousValidatorRule = append(previousValidatorRule, previousValidatorItem)
	}

	var currentValidatorRule []interface{}
	for _, currentValidatorItem := range currentValidator {
		currentValidatorRule = append(currentValidatorRule, currentValidatorItem)
	}

	logs, sub, err := _OCR2Aggregator.contract.WatchLogs(opts, "ValidatorConfigSet", previousValidatorRule, currentValidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AggregatorValidatorConfigSet)
				if err := _OCR2Aggregator.contract.UnpackLog(event, "ValidatorConfigSet", log); err != nil {
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

func (_OCR2Aggregator *OCR2AggregatorFilterer) ParseValidatorConfigSet(log types.Log) (*OCR2AggregatorValidatorConfigSet, error) {
	event := new(OCR2AggregatorValidatorConfigSet)
	if err := _OCR2Aggregator.contract.UnpackLog(event, "ValidatorConfigSet", log); err != nil {
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

func (_OCR2Aggregator *OCR2Aggregator) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OCR2Aggregator.abi.Events["AnswerUpdated"].ID:
		return _OCR2Aggregator.ParseAnswerUpdated(log)
	case _OCR2Aggregator.abi.Events["BillingAccessControllerSet"].ID:
		return _OCR2Aggregator.ParseBillingAccessControllerSet(log)
	case _OCR2Aggregator.abi.Events["BillingSet"].ID:
		return _OCR2Aggregator.ParseBillingSet(log)
	case _OCR2Aggregator.abi.Events["ConfigSet"].ID:
		return _OCR2Aggregator.ParseConfigSet(log)
	case _OCR2Aggregator.abi.Events["LinkTokenSet"].ID:
		return _OCR2Aggregator.ParseLinkTokenSet(log)
	case _OCR2Aggregator.abi.Events["NewRound"].ID:
		return _OCR2Aggregator.ParseNewRound(log)
	case _OCR2Aggregator.abi.Events["NewTransmission"].ID:
		return _OCR2Aggregator.ParseNewTransmission(log)
	case _OCR2Aggregator.abi.Events["OraclePaid"].ID:
		return _OCR2Aggregator.ParseOraclePaid(log)
	case _OCR2Aggregator.abi.Events["OwnershipTransferRequested"].ID:
		return _OCR2Aggregator.ParseOwnershipTransferRequested(log)
	case _OCR2Aggregator.abi.Events["OwnershipTransferred"].ID:
		return _OCR2Aggregator.ParseOwnershipTransferred(log)
	case _OCR2Aggregator.abi.Events["PayeeshipTransferRequested"].ID:
		return _OCR2Aggregator.ParsePayeeshipTransferRequested(log)
	case _OCR2Aggregator.abi.Events["PayeeshipTransferred"].ID:
		return _OCR2Aggregator.ParsePayeeshipTransferred(log)
	case _OCR2Aggregator.abi.Events["RequesterAccessControllerSet"].ID:
		return _OCR2Aggregator.ParseRequesterAccessControllerSet(log)
	case _OCR2Aggregator.abi.Events["RoundRequested"].ID:
		return _OCR2Aggregator.ParseRoundRequested(log)
	case _OCR2Aggregator.abi.Events["Transmitted"].ID:
		return _OCR2Aggregator.ParseTransmitted(log)
	case _OCR2Aggregator.abi.Events["ValidatorConfigSet"].ID:
		return _OCR2Aggregator.ParseValidatorConfigSet(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (OCR2AggregatorAnswerUpdated) Topic() common.Hash {
	return common.HexToHash("0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f")
}

func (OCR2AggregatorBillingAccessControllerSet) Topic() common.Hash {
	return common.HexToHash("0x793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912")
}

func (OCR2AggregatorBillingSet) Topic() common.Hash {
	return common.HexToHash("0x0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f")
}

func (OCR2AggregatorConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (OCR2AggregatorLinkTokenSet) Topic() common.Hash {
	return common.HexToHash("0x4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a")
}

func (OCR2AggregatorNewRound) Topic() common.Hash {
	return common.HexToHash("0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271")
}

func (OCR2AggregatorNewTransmission) Topic() common.Hash {
	return common.HexToHash("0xc797025feeeaf2cd924c99e9205acb8ec04d5cad21c41ce637a38fb6dee6016a")
}

func (OCR2AggregatorOraclePaid) Topic() common.Hash {
	return common.HexToHash("0xd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c")
}

func (OCR2AggregatorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OCR2AggregatorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (OCR2AggregatorPayeeshipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367")
}

func (OCR2AggregatorPayeeshipTransferred) Topic() common.Hash {
	return common.HexToHash("0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3")
}

func (OCR2AggregatorRequesterAccessControllerSet) Topic() common.Hash {
	return common.HexToHash("0x27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634")
}

func (OCR2AggregatorRoundRequested) Topic() common.Hash {
	return common.HexToHash("0x41e3990591fd372502daa15842da15bc7f41c75309ab3ff4f56f1848c178825c")
}

func (OCR2AggregatorTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (OCR2AggregatorValidatorConfigSet) Topic() common.Hash {
	return common.HexToHash("0xb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541")
}

func (_OCR2Aggregator *OCR2Aggregator) Address() common.Address {
	return _OCR2Aggregator.address
}

type OCR2AggregatorInterface interface {
	Decimals(opts *bind.CallOpts) (uint8, error)

	Description(opts *bind.CallOpts) (string, error)

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

	FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*OCR2AggregatorAnswerUpdatedIterator, error)

	WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error)

	ParseAnswerUpdated(log types.Log) (*OCR2AggregatorAnswerUpdated, error)

	FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*OCR2AggregatorBillingAccessControllerSetIterator, error)

	WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorBillingAccessControllerSet) (event.Subscription, error)

	ParseBillingAccessControllerSet(log types.Log) (*OCR2AggregatorBillingAccessControllerSet, error)

	FilterBillingSet(opts *bind.FilterOpts) (*OCR2AggregatorBillingSetIterator, error)

	WatchBillingSet(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorBillingSet) (event.Subscription, error)

	ParseBillingSet(log types.Log) (*OCR2AggregatorBillingSet, error)

	FilterConfigSet(opts *bind.FilterOpts) (*OCR2AggregatorConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*OCR2AggregatorConfigSet, error)

	FilterLinkTokenSet(opts *bind.FilterOpts, oldLinkToken []common.Address, newLinkToken []common.Address) (*OCR2AggregatorLinkTokenSetIterator, error)

	WatchLinkTokenSet(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorLinkTokenSet, oldLinkToken []common.Address, newLinkToken []common.Address) (event.Subscription, error)

	ParseLinkTokenSet(log types.Log) (*OCR2AggregatorLinkTokenSet, error)

	FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*OCR2AggregatorNewRoundIterator, error)

	WatchNewRound(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error)

	ParseNewRound(log types.Log) (*OCR2AggregatorNewRound, error)

	FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*OCR2AggregatorNewTransmissionIterator, error)

	WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error)

	ParseNewTransmission(log types.Log) (*OCR2AggregatorNewTransmission, error)

	FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*OCR2AggregatorOraclePaidIterator, error)

	WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error)

	ParseOraclePaid(log types.Log) (*OCR2AggregatorOraclePaid, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OCR2AggregatorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OCR2AggregatorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OCR2AggregatorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OCR2AggregatorOwnershipTransferred, error)

	FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*OCR2AggregatorPayeeshipTransferRequestedIterator, error)

	WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error)

	ParsePayeeshipTransferRequested(log types.Log) (*OCR2AggregatorPayeeshipTransferRequested, error)

	FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*OCR2AggregatorPayeeshipTransferredIterator, error)

	WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error)

	ParsePayeeshipTransferred(log types.Log) (*OCR2AggregatorPayeeshipTransferred, error)

	FilterRequesterAccessControllerSet(opts *bind.FilterOpts) (*OCR2AggregatorRequesterAccessControllerSetIterator, error)

	WatchRequesterAccessControllerSet(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorRequesterAccessControllerSet) (event.Subscription, error)

	ParseRequesterAccessControllerSet(log types.Log) (*OCR2AggregatorRequesterAccessControllerSet, error)

	FilterRoundRequested(opts *bind.FilterOpts, requester []common.Address) (*OCR2AggregatorRoundRequestedIterator, error)

	WatchRoundRequested(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorRoundRequested, requester []common.Address) (event.Subscription, error)

	ParseRoundRequested(log types.Log) (*OCR2AggregatorRoundRequested, error)

	FilterTransmitted(opts *bind.FilterOpts) (*OCR2AggregatorTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*OCR2AggregatorTransmitted, error)

	FilterValidatorConfigSet(opts *bind.FilterOpts, previousValidator []common.Address, currentValidator []common.Address) (*OCR2AggregatorValidatorConfigSetIterator, error)

	WatchValidatorConfigSet(opts *bind.WatchOpts, sink chan<- *OCR2AggregatorValidatorConfigSet, previousValidator []common.Address, currentValidator []common.Address) (event.Subscription, error)

	ParseValidatorConfigSet(log types.Log) (*OCR2AggregatorValidatorConfigSet, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accesscontrolledocr2aggregator

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

var AccessControlledOCR2AggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"int192\",\"name\":\"_minAnswer\",\"type\":\"int192\"},{\"internalType\":\"int192\",\"name\":\"_maxAnswer\",\"type\":\"int192\"},{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"},{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_requesterAccessController\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"AddedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"BillingAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"BillingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CheckAccessDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CheckAccessEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"oldLinkToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"newLinkToken\",\"type\":\"address\"}],\"name\":\"LinkTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int192\",\"name\":\"answer\",\"type\":\"int192\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"observationsTimestamp\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int192[]\",\"name\":\"observations\",\"type\":\"int192[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"observers\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"int192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"int192\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RemovedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"RequesterAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"}],\"name\":\"RoundRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"previousValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousGasLimit\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"currentValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"currentGasLimit\",\"type\":\"uint32\"}],\"name\":\"ValidatorConfigSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"addAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableAccessCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableAccessCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBillingAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinkToken\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequesterAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorConfig\",\"outputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"hasAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTransmissionDetails\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"},{\"internalType\":\"int192\",\"name\":\"latestAnswer_\",\"type\":\"int192\"},{\"internalType\":\"uint64\",\"name\":\"latestTimestamp_\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAnswer\",\"outputs\":[{\"internalType\":\"int192\",\"name\":\"\",\"type\":\"int192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAnswer\",\"outputs\":[{\"internalType\":\"int192\",\"name\":\"\",\"type\":\"int192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"oracleObservationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"owedPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"removeAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestNewRound\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"setBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"}],\"name\":\"setBillingAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"setLinkToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"requesterAccessController\",\"type\":\"address\"}],\"name\":\"setRequesterAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"newValidator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newGasLimit\",\"type\":\"uint32\"}],\"name\":\"setValidatorConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162006221380380620062218339810160408190526200003491620004ce565b868686868686863380600081620000685760405162461bcd60e51b81526004016200005f906200066f565b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156200009b576200009b8162000186565b5050601180546001600160a01b0319166001600160a01b038a169081179091556040519091506000907f4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a908290a3620000f48462000203565b7fff0000000000000000000000000000000000000000000000000000000000000060f883901b1660c05280516200013390601090602084019062000410565b506200013f8362000277565b6200014c600080620002e8565b50505050601791820b820b604090811b60805290820b90910b901b60a05250506015805460ff191660011790555062000760945050505050565b6001600160a01b038116331415620001b25760405162461bcd60e51b81526004016200005f90620006a6565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6012546001600160a01b0390811690821681146200027357601280546001600160a01b0319166001600160a01b0384161790556040517f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912906200026a90839085906200061e565b60405180910390a15b5050565b62000281620003e1565b600f546001600160a01b0390811690821681146200027357600f80546001600160a01b0319166001600160a01b0384161790556040517f27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634906200026a90839085906200061e565b620002f2620003e1565b60408051808201909152600e546001600160a01b03808216808452600160a01b90920463ffffffff16602084015284161415806200034057508163ffffffff16816020015163ffffffff1614155b15620003dc576040805180820182526001600160a01b0385811680835263ffffffff86166020938401819052600e80546001600160a01b031916831763ffffffff60a01b1916600160a01b9092029190911790558451928501519351909392909116917fb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca154191620003d391908790620006dd565b60405180910390a35b505050565b6000546001600160a01b031633146200040e5760405162461bcd60e51b81526004016200005f9062000638565b565b8280546200041e90620006f4565b90600052602060002090601f0160209004810192826200044257600085556200048d565b82601f106200045d57805160ff19168380011785556200048d565b828001600101855582156200048d579182015b828111156200048d57825182559160200191906001019062000470565b506200049b9291506200049f565b5090565b5b808211156200049b5760008155600101620004a0565b8051601781900b8114620004c957600080fd5b919050565b600080600080600080600060e0888a031215620004e9578283fd5b8751620004f68162000747565b9650602062000507898201620004b6565b96506200051760408a01620004b6565b95506060890151620005298162000747565b60808a01519095506200053c8162000747565b60a08a015190945060ff8116811462000553578384fd5b60c08a01519093506001600160401b038082111562000570578384fd5b818b0191508b601f83011262000584578384fd5b81518181111562000599576200059962000731565b604051601f8201601f1916810185018381118282101715620005bf57620005bf62000731565b60405281815283820185018e1015620005d6578586fd5b8592505b81831015620005f95783830185015181840186015291840191620005da565b818311156200060a57858583830101525b809550505050505092959891949750929550565b6001600160a01b0392831681529116602082015260400190565b60208082526016908201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604082015260600190565b60208082526018908201527f43616e6e6f7420736574206f776e657220746f207a65726f0000000000000000604082015260600190565b60208082526017908201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604082015260600190565b63ffffffff92831681529116602082015260400190565b6002810460018216806200070957607f821691505b602082108114156200072b57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146200075d57600080fd5b50565b60805160401c60a05160401c60c05160f81c615a6c620007b560003960006108c2015260008181610eb6015281816120b001526135780152600081816108130152818161208f015261354b0152615a6c6000f3fe608060405234801561001057600080fd5b50600436106103155760003560e01c80639a6fc8f5116101a7578063d09dc339116100ee578063e76d516811610097578063f2fde38b11610071578063f2fde38b1461062a578063fbffd2c11461063d578063feaf968c1461065057610315565b8063e76d5168146105fc578063eb45716314610604578063eb5dcd6c1461061757610315565b8063e3d0e712116100c8578063e3d0e712146105b0578063e4902f82146105c3578063e5fe4577146105e357610315565b8063d09dc33914610598578063daffc4b5146105a0578063dc7f0124146105a857610315565b8063b121e14711610150578063b633620c1161012a578063b633620c1461056a578063c10753291461057d578063c4c92b371461059057610315565b8063b121e14714610531578063b1dc65a414610544578063b5ab58dc1461055757610315565b80639e3ceeab116101815780639e3ceeab146104f4578063a118f24914610507578063afcb95d71461051a57610315565b80639a6fc8f5146104a75780639bd2c0b1146104cb5780639c849b30146104e157610315565b8063668a0f021161026b57806381ff7048116102145780638ac28d5a116101ee5780638ac28d5a1461046a5780638da5cb5b1461047d57806398e5b12a1461049257610315565b806381ff7048146104385780638205bf6a1461044f5780638823da6c1461045757610315565b80637284e416116102455780637284e4161461042057806379ba5097146104285780638038e4a11461043057610315565b8063668a0f02146103f05780636b14daf8146103f857806370da2f671461041857610315565b8063313ce567116102cd57806354fd4d50116102a757806354fd4d50146103c0578063643dc105146103c8578063666cab8d146103db57610315565b8063313ce567146103905780634fb17470146103a557806350d25bcd146103b857610315565b8063181f5a77116102fe578063181f5a771461034d57806322adbc7814610362578063299372681461037757610315565b80630a7569831461031a5780630eafb25b14610324575b600080fd5b610322610658565b005b61033761033236600461452b565b6106bf565b6040516103449190614abc565b60405180910390f35b6103556107f1565b6040516103449190614d59565b61036a610811565b6040516103449190614cac565b61037f610835565b6040516103449594939291906156db565b6103986108c0565b6040516103449190615759565b6103226103b3366004614853565b6108e4565b610337610b6d565b610337610bd9565b6103226103d636600461495d565b610bde565b6103e3610dc3565b6040516103449190614ba7565b610337610e25565b61040b61040636600461457f565b610e8c565b6040516103449190614bba565b61036a610eb4565b610355610ed8565b610322610f3f565b610322610fd8565b610440611040565b604051610344939291906156ba565b61033761105c565b61032261046536600461452b565b6110c3565b61032261047836600461452b565b611166565b6104856111a8565b6040516103449190614b1d565b61049a6111b7565b604051610344919061570f565b6104ba6104b53660046149d4565b611316565b604051610344959493929190615726565b6104d3611398565b604051610344929190614c8d565b6103226104ef3660046145f8565b6113dc565b61032261050236600461452b565b61158e565b61032261051536600461452b565b61161e565b6105226116b7565b60405161034493929190614bc5565b61032261053f36600461452b565b6116d2565b610322610552366004614729565b611796565b610337610565366004614865565b611bb5565b610337610578366004614865565b611c1d565b61032261058b3660046145cd565b611c85565b610485611ee2565b610337611ef1565b610485611fae565b61040b611fbd565b6103226105be366004614661565b611fc6565b6105d66105d136600461452b565b612767565b604051610344919061560c565b6105eb612834565b604051610344959493929190614c1c565b6104856128c9565b610322610612366004614826565b6128d8565b610322610625366004614547565b612a1d565b61032261063836600461452b565b612b0e565b61032261064b36600461452b565b612b1f565b6104ba612b30565b610660612bae565b60155460ff16156106bd57601580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f53963890600090a15b565b6001600160a01b03811660009081526002602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046bffffffffffffffffffffffff1691810191909152906107275760009150506107ec565b600b5460208201516000917201000000000000000000000000000000000000900463ffffffff169060069060ff16601f811061077357634e487b7160e01b600052603260045260246000fd5b600881049190910154600b546107a9926007166004026101000a90910463ffffffff90811691660100000000000090041661592b565b63ffffffff166107b9919061583a565b6107c790633b9aca0061583a565b905081604001516bffffffffffffffffffffffff16816107e791906157b5565b925050505b919050565b60606040518060600160405280602a8152602001615a36602a9139905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b600b546a0100000000000000000000810463ffffffff908116926e010000000000000000000000000000830482169272010000000000000000000000000000000000008104831692760100000000000000000000000000000000000000000000820416917a01000000000000000000000000000000000000000000000000000090910462ffffff1690565b7f000000000000000000000000000000000000000000000000000000000000000081565b6108ec612bae565b6011546001600160a01b0390811690831681141561090a5750610b69565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b038416906370a082319061094f903090600401614b1d565b60206040518083038186803b15801561096757600080fd5b505afa15801561097b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099f919061487d565b506109a8612bd8565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526000906001600160a01b038316906370a08231906109f0903090600401614b1d565b60206040518083038186803b158015610a0857600080fd5b505afa158015610a1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a40919061487d565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000081529091506001600160a01b0383169063a9059cbb90610a8a9086908590600401614b8e565b602060405180830381600087803b158015610aa457600080fd5b505af1158015610ab8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610adc9190614806565b610b015760405162461bcd60e51b8152600401610af890614fc7565b60405180910390fd5b601180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0386811691821790925560405190918416907f4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a90600090a350505b5050565b6000610bb0336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e8c92505050565b610bcc5760405162461bcd60e51b8152600401610af8906151b6565b610bd4612feb565b905090565b600681565b6012546001600160a01b0316610bf26111a8565b6001600160a01b0316336001600160a01b03161480610ca657506040517f6b14daf80000000000000000000000000000000000000000000000000000000081526001600160a01b03821690636b14daf890610c569033906000903690600401614b31565b60206040518083038186803b158015610c6e57600080fd5b505afa158015610c82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ca69190614806565b610cc25760405162461bcd60e51b8152600401610af89061525b565b610cca612bd8565b85600b600001600a6101000a81548163ffffffff021916908363ffffffff16021790555084600b600001600e6101000a81548163ffffffff021916908363ffffffff16021790555083600b60000160126101000a81548163ffffffff021916908363ffffffff16021790555082600b60000160166101000a81548163ffffffff021916908363ffffffff16021790555081600b600001601a6101000a81548162ffffff021916908362ffffff1602179055507f0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f8686868686604051610db39594939291906156db565b60405180910390a1505050505050565b60606005805480602002602001604051908101604052809291908181526020018280548015610e1b57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610dfd575b5050505050905090565b6000610e68336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e8c92505050565b610e845760405162461bcd60e51b8152600401610af8906151b6565b610bd4613017565b6000610e98838361302d565b80610eab57506001600160a01b03831632145b90505b92915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6060610f1b336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e8c92505050565b610f375760405162461bcd60e51b8152600401610af8906151b6565b610bd461305d565b6001546001600160a01b03163314610f695760405162461bcd60e51b8152600401610af890614da3565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610fe0612bae565b60155460ff166106bd57601580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517faebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c348090600090a1565b600d54600a5463ffffffff808316926401000000009004169192565b600061109f336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e8c92505050565b6110bb5760405162461bcd60e51b8152600401610af8906151b6565b610bd46130e6565b6110cb612bae565b6001600160a01b03811660009081526016602052604090205460ff1615611163576001600160a01b0381166000908152601660205260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517f3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d19061115a908390614b1d565b60405180910390a15b50565b6001600160a01b0381811660009081526013602052604090205416331461119f5760405162461bcd60e51b8152600401610af8906152fe565b6111638161312e565b6000546001600160a01b031690565b60006111c16111a8565b6001600160a01b0316336001600160a01b031614806112795750600f546040517f6b14daf80000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636b14daf8906112299033906000903690600401614b31565b60206040518083038186803b15801561124157600080fd5b505afa158015611255573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112799190614806565b6112955760405162461bcd60e51b8152600401610af890614eb4565b600b54600a54604051610100830464ffffffffff8116936601000000000000900463ffffffff9081169333937f41e3990591fd372502daa15842da15bc7f41c75309ab3ff4f56f1848c178825c936112f69360089190911c16908790614bfd565b60405180910390a26113098160016157cd565b63ffffffff169250505090565b600080600080600061135f336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e8c92505050565b61137b5760405162461bcd60e51b8152600401610af8906151b6565b6113848661336d565b945094509450945094505b91939590929450565b60408051808201909152600e546001600160a01b0381168083527401000000000000000000000000000000000000000090910463ffffffff16602090920182905291565b6113e4612bae565b8281146114035760405162461bcd60e51b8152600401610af890615448565b60005b8381101561158757600085858381811061143057634e487b7160e01b600052603260045260246000fd5b9050602002016020810190611445919061452b565b9050600084848481811061146957634e487b7160e01b600052603260045260246000fd5b905060200201602081019061147e919061452b565b6001600160a01b0380841660009081526013602052604090205491925016801580806114bb5750826001600160a01b0316826001600160a01b0316145b6114d75760405162461bcd60e51b8152600401610af890615224565b6001600160a01b03848116600090815260136020526040902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000168583169081179091559083161461157057826001600160a01b0316826001600160a01b0316856001600160a01b03167f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b360405160405180910390a45b50505050808061157f90615985565b915050611406565b5050505050565b611596612bae565b600f546001600160a01b039081169082168114610b6957600f80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384161790556040517f27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634906116129083908590614c73565b60405180910390a15050565b611626612bae565b6001600160a01b03811660009081526016602052604090205460ff16611163576001600160a01b0381166000908152601660205260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055517f87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db49061115a908390614b1d565b600a54600b5460009261010090910460081c63ffffffff1690565b6001600160a01b0381811660009081526014602052604090205416331461170b5760405162461bcd60e51b8152600401610af890614d6c565b6001600160a01b0381811660008181526013602090815260408083208054337fffffffffffffffffffffffff000000000000000000000000000000000000000080831682179093556014909452828520805490921690915590519416939092849290917f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b39190a45050565b60005a604080516101008082018352600b5460ff8116835290810464ffffffffff90811660208085018290526601000000000000840463ffffffff908116968601969096526a01000000000000000000008404861660608601526e01000000000000000000000000000084048616608086015272010000000000000000000000000000000000008404861660a0860152760100000000000000000000000000000000000000000000840490951660c08501527a01000000000000000000000000000000000000000000000000000090920462ffffff1660e08401529394509092918c01359182161161189a5760405162461bcd60e51b8152600401610af890614e7d565b3360009081526002602052604090205460ff166118c95760405162461bcd60e51b8152600401610af890615290565b600a548b35146118eb5760405162461bcd60e51b8152600401610af8906153da565b6118f98a8a8a8a8a8a613430565b81516119069060016157f5565b60ff1687146119275760405162461bcd60e51b8152600401610af8906152c7565b8685146119465760405162461bcd60e51b8152600401610af890615522565b60008a8a604051611958929190614ac5565b60405190819003812061196f918e90602001614be3565b604051602081830303815290604052805190602001209050600061199161421a565b60005b8a811015611b095760006001858a84602081106119c157634e487b7160e01b600052603260045260246000fd5b6119ce91901a601b6157f5565b8f8f868181106119ee57634e487b7160e01b600052603260045260246000fd5b905060200201358e8e87818110611a1557634e487b7160e01b600052603260045260246000fd5b9050602002013560405160008152602001604052604051611a399493929190614c55565b6020604051602081039080840390855afa158015611a5b573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101516001600160a01b03811660009081526003602090815290849020838501909452925460ff8082161515808552610100909204169383019390935290955092509050611ae25760405162461bcd60e51b8152600401610af89061536c565b826020015160080260ff166001901b84019350508080611b0190615985565b915050611994565b5081827e010101010101010101010101010101010101010101010101010101010101011614611b4a5760405162461bcd60e51b8152600401610af890615035565b5060009150611b999050838d836020020135848e8e8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061349d92505050565b9050611ba78382863361391f565b505050505050505050505050565b6000611bf8336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e8c92505050565b611c145760405162461bcd60e51b8152600401610af8906151b6565b610eae82613a70565b6000611c60336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e8c92505050565b611c7c5760405162461bcd60e51b8152600401610af8906151b6565b610eae82613aa6565b611c8d6111a8565b6001600160a01b0316336001600160a01b03161480611d4557506012546040517f6b14daf80000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636b14daf890611cf59033906000903690600401614b31565b60206040518083038186803b158015611d0d57600080fd5b505afa158015611d21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d459190614806565b611d615760405162461bcd60e51b8152600401610af89061525b565b6000611d6b613af8565b6011546040517f70a082310000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b03909116906370a0823190611dba903090600401614b1d565b60206040518083038186803b158015611dd257600080fd5b505afa158015611de6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e0a919061487d565b905081811015611e2c5760405162461bcd60e51b8152600401610af890615335565b6011546001600160a01b031663a9059cbb85611e51611e4b8686615914565b87613cf5565b6040518363ffffffff1660e01b8152600401611e6e929190614b8e565b602060405180830381600087803b158015611e8857600080fd5b505af1158015611e9c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ec09190614806565b611edc5760405162461bcd60e51b8152600401610af8906154b4565b50505050565b6012546001600160a01b031690565b6011546040517f70a0823100000000000000000000000000000000000000000000000000000000815260009182916001600160a01b03909116906370a0823190611f3f903090600401614b1d565b60206040518083038186803b158015611f5757600080fd5b505afa158015611f6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f8f919061487d565b90506000611f9b613af8565b9050611fa781836158a0565b9250505090565b600f546001600160a01b031690565b60155460ff1681565b611fce612bae565b601f86511115611ff05760405162461bcd60e51b8152600401610af890614f22565b84518651146120115760405162461bcd60e51b8152600401610af890614f59565b855161201e856003615877565b60ff161061203e5760405162461bcd60e51b8152600401610af89061547d565b61204a8460ff16613d0c565b8251156120695760405162461bcd60e51b8152600401610af8906151ed565b60006040518060c001604052808881526020018781526020018660ff16815260200160017f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006040516020016120e193929190614ad5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815290825267ffffffffffffffff8616602083015201839052600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000ff1690559050612158612bd8565b60045460005b818110156122535760006004828154811061218957634e487b7160e01b600052603260045260246000fd5b6000918252602082200154600580546001600160a01b03909216935090849081106121c457634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b03948516835260038252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905594168252600290529190912080547fffffffffffffffffffffffffffffffffffff0000000000000000000000000000169055508061224b81615985565b91505061215e565b5061226060046000614231565b61226c60056000614231565b60005b8251518110156125505760036000846000015183815181106122a157634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff16156122e55760405162461bcd60e51b8152600401610af8906150da565b604080518082019091526001815260ff82166020820152835180516003916000918590811061232457634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252818101929092526040016000908120835181549484015160ff16610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff9115157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0090961695909517169390931790925584015180516002929190849081106123d557634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff16156124195760405162461bcd60e51b8152600401610af890614e11565b60405180606001604052806001151581526020018260ff16815260200160006bffffffffffffffffffffffff16815250600260008560200151848151811061247157634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252818101929092526040908101600020835181549385015194909201517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909316911515919091177fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1661010060ff90941693909302929092177fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff16620100006bffffffffffffffffffffffff909216919091021790558061254881615985565b91505061226f565b50815180516125679160049160209091019061424f565b50602080830151805161257e92600592019061424f565b506040820151600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600d80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff4381168202929092178085559204811692600192909160009161260c918591166157cd565b92506101000a81548163ffffffff021916908363ffffffff16021790555061266b4630600d60009054906101000a900463ffffffff1663ffffffff1686600001518760200151886040015189606001518a608001518b60a00151613d2c565b600a819055600d5484516020860151604080880151606089015160808a015160a08b015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05986126d4988b98919763ffffffff90911696919590949093909290919061561d565b60405180910390a1600b546601000000000000900463ffffffff1660005b84515181101561275a5781600682601f811061271e57634e487b7160e01b600052603260045260246000fd5b600891828204019190066004026101000a81548163ffffffff021916908363ffffffff160217905550808061275290615985565b9150506126f2565b5050505050505050505050565b6001600160a01b03811660009081526002602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046bffffffffffffffffffffffff1691810191909152906127cf5760009150506107ec565b6006816020015160ff16601f81106127f757634e487b7160e01b600052603260045260246000fd5b600881049190910154600b5461282d926007166004026101000a90910463ffffffff90811691660100000000000090041661592b565b9392505050565b6000808080803332146128595760405162461bcd60e51b8152600401610af890614ffe565b5050600a54600b5463ffffffff6601000000000000820481166000908152600c60205260409020549296610100909204600881901c8216965064ffffffffff169450601783900b93507c010000000000000000000000000000000000000000000000000000000090920490911690565b6011546001600160a01b031690565b6128e0612bae565b60408051808201909152600e546001600160a01b038082168084527401000000000000000000000000000000000000000090920463ffffffff166020840152841614158061293e57508163ffffffff16816020015163ffffffff1614155b15612a18576040805180820182526001600160a01b0385811680835263ffffffff86166020938401819052600e80547fffffffffffffffffffffffff00000000000000000000000000000000000000001683177fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000009092029190911790558451928501519351909392909116917fb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca154191612a0f919087906156a3565b60405180910390a35b505050565b6001600160a01b03828116600090815260136020526040902054163314612a565760405162461bcd60e51b8152600401610af890614dda565b336001600160a01b0382161415612a7f5760405162461bcd60e51b8152600401610af890615411565b6001600160a01b03808316600090815260146020526040902080548383167fffffffffffffffffffffffff000000000000000000000000000000000000000082168117909255909116908114612a18576040516001600160a01b038084169133918616907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e3836790600090a4505050565b612b16612bae565b61116381613dd8565b612b27612bae565b61116381613e6a565b6000806000806000612b79336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e8c92505050565b612b955760405162461bcd60e51b8152600401610af8906151b6565b612b9d613ee6565b945094509450945094509091929394565b6000546001600160a01b031633146106bd5760405162461bcd60e51b8152600401610af890614eeb565b601154600b54604080516103e08101918290526001600160a01b0390931692660100000000000090920463ffffffff1691600091600690601f908285855b82829054906101000a900463ffffffff1663ffffffff1681526020019060040190602082600301049283019260010382029150808411612c165790505050505050905060006005805480602002602001604051908101604052809291908181526020018280548015612cb157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612c93575b5050505050905060005b8151811015612fdd57600060026000848481518110612cea57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160029054906101000a90046bffffffffffffffffffffffff166bffffffffffffffffffffffff169050600060026000858581518110612d6457634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160026101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff16021790555060008483601f8110612ddf57634e487b7160e01b600052603260045260246000fd5b6020020151600b5490870363ffffffff90811692507201000000000000000000000000000000000000909104168102633b9aca000282018015612fd257600060136000878781518110612e4257634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b03169050886001600160a01b031663a9059cbb82846040518363ffffffff1660e01b8152600401612eae929190614b8e565b602060405180830381600087803b158015612ec857600080fd5b505af1158015612edc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f009190614806565b612f1c5760405162461bcd60e51b8152600401610af8906154b4565b878786601f8110612f3d57634e487b7160e01b600052603260045260246000fd5b602002019063ffffffff16908163ffffffff1681525050886001600160a01b0316816001600160a01b0316878781518110612f8857634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c85604051612fc89190614abc565b60405180910390a4505b505050600101612cbb565b50611587600683601f6142cc565b600b546601000000000000900463ffffffff166000908152600c6020526040902054601790810b900b90565b600b546601000000000000900463ffffffff1690565b6001600160a01b03821660009081526016602052604081205460ff1680610eab57505060155460ff161592915050565b60606010805461306c90615950565b80601f016020809104026020016040519081016040528092919081815260200182805461309890615950565b8015610e1b5780601f106130ba57610100808354040283529160200191610e1b565b820191906000526020600020905b8154815290600101906020018083116130c857509395945050505050565b600b5463ffffffff660100000000000090910481166000908152600c60205260409020547c010000000000000000000000000000000000000000000000000000000090041690565b6001600160a01b0381166000908152600260209081526040918290208251606081018452905460ff80821615158084526101008304909116938301939093526201000090046bffffffffffffffffffffffff16928101929092526131925750611163565b600061319d836106bf565b90508015612a18576001600160a01b03808416600090815260136020526040908190205460115491517fa9059cbb00000000000000000000000000000000000000000000000000000000815290831692919091169063a9059cbb906132089084908690600401614b8e565b602060405180830381600087803b15801561322257600080fd5b505af1158015613236573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061325a9190614806565b6132765760405162461bcd60e51b8152600401610af8906154b4565b600b60000160069054906101000a900463ffffffff166006846020015160ff16601f81106132b457634e487b7160e01b600052603260045260246000fd5b6008810491909101805460079092166004026101000a63ffffffff8181021990931693909216919091029190911790556001600160a01b038481166000818152600260205260409081902080547fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff169055601154905190831692841691907fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c9061335f908790614abc565b60405180910390a450505050565b60008080808063ffffffff69ffffffffffffffffffff8716111561339f5750600093508392508291508190508061138f565b5050505063ffffffff8281166000908152600c602090815260409182902082516060810184529054601781810b810b810b8084527801000000000000000000000000000000000000000000000000830487169484018590527c01000000000000000000000000000000000000000000000000000000009092049095169190930181905294959190920b939192508490565b600061343d82602061583a565b61344885602061583a565b613454886101446157b5565b61345e91906157b5565b61346891906157b5565b6134739060006157b5565b90503681146134945760405162461bcd60e51b8152600401610af890614f90565b50505050505050565b6000806134a983613f7f565b9050601f81604001515111156134d15760405162461bcd60e51b8152600401610af890615148565b604081015151865160ff16106134f95760405162461bcd60e51b8152600401610af89061506c565b64ffffffffff84166020870152604081015180516000919061351d9060029061581a565b8151811061353b57634e487b7160e01b600052603260045260246000fd5b602002602001015190508060170b7f000000000000000000000000000000000000000000000000000000000000000060170b131580156135a157507f000000000000000000000000000000000000000000000000000000000000000060170b8160170b13155b6135bd5760405162461bcd60e51b8152600401610af89061517f565b604087018051906135cd826159be565b63ffffffff1663ffffffff168152505060405180606001604052808260170b8152602001836000015163ffffffff1681526020014263ffffffff16815250600c6000896040015163ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a81548177ffffffffffffffffffffffffffffffffffffffffffffffff021916908360170b77ffffffffffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160186101000a81548163ffffffff021916908363ffffffff160217905550604082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555090505086600b60008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548164ffffffffff021916908364ffffffffff16021790555060408201518160000160066101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001600a6101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600e6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160000160126101000a81548163ffffffff021916908363ffffffff16021790555060c08201518160000160166101000a81548163ffffffff021916908363ffffffff16021790555060e082015181600001601a6101000a81548162ffffff021916908362ffffff160217905550905050866040015163ffffffff167fc797025feeeaf2cd924c99e9205acb8ec04d5cad21c41ce637a38fb6dee6016a823385600001518660400151876020015188606001518d8d604051613861989796959493929190614cba565b60405180910390a260006001600160a01b0316876040015163ffffffff167f0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac6027184600001516040516138b2919061560c565b60405180910390a3866040015163ffffffff168160170b7f0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f426040516138f89190614abc565b60405180910390a361391187604001518260170b614004565b506060015195945050505050565b60008360170b121561393057611edc565b6000613957633b9aca003a04866080015163ffffffff16876060015163ffffffff1661413c565b90506010360260005a905060006139808663ffffffff1685858b60e0015162ffffff1686614162565b90506000670de0b6b3a764000077ffffffffffffffffffffffffffffffffffffffffffffffff891683026001600160a01b03881660009081526002602052604090205460c08c01519290910492506201000090046bffffffffffffffffffffffff9081169163ffffffff16633b9aca000282840101908116821115613a0b5750505050505050611edc565b6001600160a01b038816600090815260026020526040902080546bffffffffffffffffffffffff90921662010000027fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff90921691909117905550505050505050505050565b600063ffffffff821115613a86575060006107ec565b5063ffffffff166000908152600c6020526040902054601790810b900b90565b600063ffffffff821115613abc575060006107ec565b5063ffffffff9081166000908152600c60205260409020547c010000000000000000000000000000000000000000000000000000000090041690565b6000806005805480602002602001604051908101604052809291908181526020018280548015613b5157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613b33575b50508351600b54604080516103e08101918290529697509195660100000000000090910463ffffffff169450600093509150600690601f908285855b82829054906101000a900463ffffffff1663ffffffff1681526020019060040190602082600301049283019260010382029150808411613b8d5790505050505050905060005b83811015613c2e578181601f8110613bfb57634e487b7160e01b600052603260045260246000fd5b6020020151613c0a908461592b565b613c1a9063ffffffff16876157b5565b955080613c2681615985565b915050613bd3565b50600b54613c5c907201000000000000000000000000000000000000900463ffffffff16633b9aca0061583a565b613c66908661583a565b945060005b83811015613ced5760026000868381518110613c9757634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0316825281019190915260400160002054613cd9906201000090046bffffffffffffffffffffffff16876157b5565b955080613ce581615985565b915050613c6b565b505050505090565b600081831015613d06575081610eae565b50919050565b806000106111635760405162461bcd60e51b8152600401610af8906153a3565b6000808a8a8a8a8a8a8a8a8a604051602001613d5099989796959493929190615559565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150505b9998505050505050505050565b6001600160a01b038116331415613e015760405162461bcd60e51b8152600401610af8906154eb565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6012546001600160a01b039081169082168114610b6957601280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384161790556040517f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912906116129083908590614c73565b600b546601000000000000900463ffffffff9081166000818152600c602090815260409182902082516060810184529054601781810b810b810b8084527801000000000000000000000000000000000000000000000000830488169484018590527c01000000000000000000000000000000000000000000000000000000009092049096169190930181905292949190930b9291908490565b613f87614363565b6000806060600085806020019051810190613fa29190614895565b92965090945092509050613fb68683614196565b8151604051600090613fcc908690602001614abc565b60408051918152928152825160808101845263ffffffff909716875260208701525084019190915260170b6060830152509050919050565b60408051808201909152600e546001600160a01b0381168083527401000000000000000000000000000000000000000090910463ffffffff16602083015261404c5750610b69565b600061405960018561592b565b63ffffffff8082166000818152600c6020908152604091829020549087015187519251959650601791820b90910b9461412094918216936140a592909187918c16908b906024016155f1565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fbeed9b51000000000000000000000000000000000000000000000000000000001790526141de565b6115875760405162461bcd60e51b8152600401610af890615111565b6000838381101561414f57600285850304015b6141598184613cf5565b95945050505050565b6000818610156141845760405162461bcd60e51b8152600401610af890614e48565b50633b9aca0094039190910101020290565b6000815160206141a6919061583a565b6141b19060a06157b5565b6141bc9060006157b5565b905080835114612a185760405162461bcd60e51b8152600401610af8906150a3565b60005a61138881106142125761138881039050846040820482031115614212576000808451602086016000888af150600191505b509392505050565b604080518082019091526000808252602082015290565b50805460008255906000526020600020908101906111639190614394565b8280548282559060005260206000209081019282156142bc579160200282015b828111156142bc57825182547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0390911617825560209092019160019091019061426f565b506142c8929150614394565b5090565b6004830191839082156142bc5791602002820160005b8382111561432657835183826101000a81548163ffffffff021916908363ffffffff16021790555092602001926004016020816003010492830192600103026142e2565b80156143565782816101000a81549063ffffffff0219169055600401602081600301049283019260010302614326565b50506142c8929150614394565b6040518060800160405280600063ffffffff1681526020016060815260200160608152602001600060170b81525090565b5b808211156142c85760008155600101614395565b60008083601f8401126143ba578182fd5b50813567ffffffffffffffff8111156143d1578182fd5b60208301915083602080830285010111156143eb57600080fd5b9250929050565b600082601f830112614402578081fd5b8135602061441761441283615791565b615767565b8281528181019085830183850287018401881015614433578586fd5b855b8581101561445a57813561444881615a0e565b84529284019290840190600101614435565b5090979650505050505050565b600082601f830112614477578081fd5b813567ffffffffffffffff811115614491576144916159f8565b6144c260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601615767565b8181528460208386010111156144d6578283fd5b816020850160208301379081016020019190915292915050565b8051601781900b81146107ec57600080fd5b803567ffffffffffffffff811681146107ec57600080fd5b803560ff811681146107ec57600080fd5b60006020828403121561453c578081fd5b813561282d81615a0e565b60008060408385031215614559578081fd5b823561456481615a0e565b9150602083013561457481615a0e565b809150509250929050565b60008060408385031215614591578182fd5b823561459c81615a0e565b9150602083013567ffffffffffffffff8111156145b7578182fd5b6145c385828601614467565b9150509250929050565b600080604083850312156145df578182fd5b82356145ea81615a0e565b946020939093013593505050565b6000806000806040858703121561460d578182fd5b843567ffffffffffffffff80821115614624578384fd5b614630888389016143a9565b90965094506020870135915080821115614648578384fd5b50614655878288016143a9565b95989497509550505050565b60008060008060008060c08789031215614679578384fd5b863567ffffffffffffffff80821115614690578586fd5b61469c8a838b016143f2565b975060208901359150808211156146b1578586fd5b6146bd8a838b016143f2565b96506146cb60408a0161451a565b955060608901359150808211156146e0578384fd5b6146ec8a838b01614467565b94506146fa60808a01614502565b935060a089013591508082111561470f578283fd5b5061471c89828a01614467565b9150509295509295509295565b60008060008060008060008060e0898b031215614744578586fd5b606089018a811115614754578687fd5b8998503567ffffffffffffffff8082111561476d578788fd5b818b0191508b601f830112614780578788fd5b81358181111561478e578889fd5b8c602082850101111561479f578889fd5b6020830199508098505060808b01359150808211156147bc578384fd5b6147c88c838d016143a9565b909750955060a08b01359150808211156147e0578384fd5b506147ed8b828c016143a9565b999c989b50969995989497949560c00135949350505050565b600060208284031215614817578081fd5b8151801515811461282d578182fd5b60008060408385031215614838578182fd5b823561484381615a0e565b9150602083013561457481615a23565b60008060408385031215614559578182fd5b600060208284031215614876578081fd5b5035919050565b60006020828403121561488e578081fd5b5051919050565b600080600080608085870312156148aa578182fd5b84516148b581615a23565b809450506020808601519350604086015167ffffffffffffffff8111156148da578384fd5b8601601f810188136148ea578384fd5b80516148f861441282615791565b81815283810190838501858402850186018c1015614914578788fd5b8794505b8385101561493d57614929816144f0565b835260019490940193918501918501614918565b508096505050505050614952606086016144f0565b905092959194509250565b600080600080600060a08688031215614974578283fd5b853561497f81615a23565b9450602086013561498f81615a23565b9350604086013561499f81615a23565b925060608601356149af81615a23565b9150608086013562ffffff811681146149c6578182fd5b809150509295509295909350565b6000602082840312156149e5578081fd5b813569ffffffffffffffffffff8116811461282d578182fd5b6000815180845260208085019450808401835b83811015614a365781516001600160a01b031687529582019590820190600101614a11565b509495945050505050565b60008151808452815b81811015614a6657602081850181015186830182015201614a4a565b81811115614a775782602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60170b9052565b64ffffffffff169052565b90815260200190565b6000828483379101908152919050565b60f89390931b7fff00000000000000000000000000000000000000000000000000000000000000168352601791820b604090811b6001850152910b901b601982015260310190565b6001600160a01b0391909116815260200190565b60006001600160a01b03851682526040602083015282604083015282846060840137818301606090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010192915050565b6001600160a01b03929092168252602082015260400190565b600060208252610eab60208301846149fe565b901515815260200190565b9215158352602083019190915263ffffffff16604082015260600190565b828152608081016060836020840137600081529392505050565b92835263ffffffff91909116602083015260ff16604082015260600190565b94855263ffffffff93909316602085015260ff91909116604084015260170b606083015267ffffffffffffffff16608082015260a00190565b93845260ff9290921660208401526040830152606082015260800190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b0392909216825263ffffffff16602082015260400190565b60179190910b815260200190565b600061010080830160178c810b855260206001600160a01b038d168187015263ffffffff8c1660408701528360608701528293508a5180845261012087019450818c019350855b81811015614d1f578451840b86529482019493820193600101614d01565b50505050508281036080840152614d368188614a41565b915050614d4660a0830186614aaa565b8360c0830152613dcb60e0830184614ab1565b600060208252610eab6020830184614a41565b6020808252601f908201527f6f6e6c792070726f706f736564207061796565732063616e2061636365707400604082015260600190565b60208082526016908201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604082015260600190565b6020808252601d908201527f6f6e6c792063757272656e742070617965652063616e20757064617465000000604082015260600190565b6020808252601c908201527f7265706561746564207472616e736d6974746572206164647265737300000000604082015260600190565b6020808252818101527f6c6566744761732063616e6e6f742065786365656420696e697469616c476173604082015260600190565b6020808252600c908201527f7374616c65207265706f72740000000000000000000000000000000000000000604082015260600190565b6020808252601d908201527f4f6e6c79206f776e6572267265717565737465722063616e2063616c6c000000604082015260600190565b60208082526016908201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604082015260600190565b60208082526010908201527f746f6f206d616e79206f7261636c657300000000000000000000000000000000604082015260600190565b60208082526016908201527f6f7261636c65206c656e677468206d69736d6174636800000000000000000000604082015260600190565b60208082526018908201527f63616c6c64617461206c656e677468206d69736d617463680000000000000000604082015260600190565b6020808252601f908201527f7472616e736665722072656d61696e696e672066756e6473206661696c656400604082015260600190565b60208082526014908201527f4f6e6c792063616c6c61626c6520627920454f41000000000000000000000000604082015260600190565b60208082526010908201527f6475706c6963617465207369676e657200000000000000000000000000000000604082015260600190565b6020808252601e908201527f746f6f206665772076616c75657320746f207472757374206d656469616e0000604082015260600190565b60208082526016908201527f7265706f7274206c656e677468206d69736d6174636800000000000000000000604082015260600190565b60208082526017908201527f7265706561746564207369676e65722061646472657373000000000000000000604082015260600190565b60208082526010908201527f696e73756666696369656e742067617300000000000000000000000000000000604082015260600190565b6020808252601e908201527f6e756d206f62736572766174696f6e73206f7574206f6620626f756e64730000604082015260600190565b6020808252601e908201527f6d656469616e206973206f7574206f66206d696e2d6d61782072616e67650000604082015260600190565b60208082526009908201527f4e6f206163636573730000000000000000000000000000000000000000000000604082015260600190565b6020808252601b908201527f6f6e636861696e436f6e666967206d75737420626520656d7074790000000000604082015260600190565b60208082526011908201527f706179656520616c726561647920736574000000000000000000000000000000604082015260600190565b6020808252818101527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c604082015260600190565b60208082526018908201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604082015260600190565b6020808252601a908201527f77726f6e67206e756d626572206f66207369676e617475726573000000000000604082015260600190565b60208082526017908201527f4f6e6c792070617965652063616e207769746864726177000000000000000000604082015260600190565b60208082526014908201527f696e73756666696369656e742062616c616e6365000000000000000000000000604082015260600190565b6020808252600f908201527f7369676e6174757265206572726f720000000000000000000000000000000000604082015260600190565b60208082526012908201527f66206d75737420626520706f7369746976650000000000000000000000000000604082015260600190565b60208082526015908201527f636f6e666967446967657374206d69736d617463680000000000000000000000604082015260600190565b60208082526017908201527f63616e6e6f74207472616e7366657220746f2073656c66000000000000000000604082015260600190565b6020808252818101527f7472616e736d6974746572732e73697a6520213d207061796565732e73697a65604082015260600190565b60208082526018908201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604082015260600190565b60208082526012908201527f696e73756666696369656e742066756e64730000000000000000000000000000604082015260600190565b60208082526017908201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604082015260600190565b6020808252601e908201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604082015260600190565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526155938285018b6149fe565b915083820360808501526155a7828a6149fe565b915060ff881660a085015283820360c08501526155c48288614a41565b90861660e085015283810361010085015290506155e18185614a41565b9c9b505050505050505050505050565b93845260208401929092526040830152606082015260800190565b63ffffffff91909116815260200190565b600061012063ffffffff808d1684528b6020850152808b1660408501525080606084015261564d8184018a6149fe565b9050828103608084015261566181896149fe565b905060ff871660a084015282810360c084015261567e8187614a41565b905067ffffffffffffffff851660e08401528281036101008401526155e18185614a41565b63ffffffff92831681529116602082015260400190565b63ffffffff9384168152919092166020820152604081019190915260600190565b63ffffffff958616815293851660208501529184166040840152909216606082015262ffffff909116608082015260a00190565b69ffffffffffffffffffff91909116815260200190565b69ffffffffffffffffffff9586168152602081019490945260408401929092526060830152909116608082015260a00190565b60ff91909116815260200190565b60405181810167ffffffffffffffff81118282101715615789576157896159f8565b604052919050565b600067ffffffffffffffff8211156157ab576157ab6159f8565b5060209081020190565b600082198211156157c8576157c86159e2565b500190565b600063ffffffff8083168185168083038211156157ec576157ec6159e2565b01949350505050565b600060ff821660ff84168060ff03821115615812576158126159e2565b019392505050565b60008261583557634e487b7160e01b81526012600452602481fd5b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615615872576158726159e2565b500290565b600060ff821660ff84168160ff0481118215151615615898576158986159e2565b029392505050565b6000808312837f8000000000000000000000000000000000000000000000000000000000000000018312811516156158da576158da6159e2565b837f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01831381161561590e5761590e6159e2565b50500390565b600082821015615926576159266159e2565b500390565b600063ffffffff83811690831681811015615948576159486159e2565b039392505050565b60028104600182168061596457607f821691505b60208210811415613d0657634e487b7160e01b600052602260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156159b7576159b76159e2565b5060010190565b600063ffffffff808316818114156159d8576159d86159e2565b6001019392505050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461116357600080fd5b63ffffffff8116811461116357600080fdfe416363657373436f6e74726f6c6c65644f43523241676772656761746f7220312e302e302d616c706861a164736f6c6343000800000a",
}

var AccessControlledOCR2AggregatorABI = AccessControlledOCR2AggregatorMetaData.ABI

var AccessControlledOCR2AggregatorBin = AccessControlledOCR2AggregatorMetaData.Bin

func DeployAccessControlledOCR2Aggregator(auth *bind.TransactOpts, backend bind.ContractBackend, _link common.Address, _minAnswer *big.Int, _maxAnswer *big.Int, _billingAccessController common.Address, _requesterAccessController common.Address, _decimals uint8, description string) (common.Address, *types.Transaction, *AccessControlledOCR2Aggregator, error) {
	parsed, err := AccessControlledOCR2AggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccessControlledOCR2AggregatorBin), backend, _link, _minAnswer, _maxAnswer, _billingAccessController, _requesterAccessController, _decimals, description)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessControlledOCR2Aggregator{AccessControlledOCR2AggregatorCaller: AccessControlledOCR2AggregatorCaller{contract: contract}, AccessControlledOCR2AggregatorTransactor: AccessControlledOCR2AggregatorTransactor{contract: contract}, AccessControlledOCR2AggregatorFilterer: AccessControlledOCR2AggregatorFilterer{contract: contract}}, nil
}

type AccessControlledOCR2Aggregator struct {
	address common.Address
	abi     abi.ABI
	AccessControlledOCR2AggregatorCaller
	AccessControlledOCR2AggregatorTransactor
	AccessControlledOCR2AggregatorFilterer
}

type AccessControlledOCR2AggregatorCaller struct {
	contract *bind.BoundContract
}

type AccessControlledOCR2AggregatorTransactor struct {
	contract *bind.BoundContract
}

type AccessControlledOCR2AggregatorFilterer struct {
	contract *bind.BoundContract
}

type AccessControlledOCR2AggregatorSession struct {
	Contract     *AccessControlledOCR2Aggregator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type AccessControlledOCR2AggregatorCallerSession struct {
	Contract *AccessControlledOCR2AggregatorCaller
	CallOpts bind.CallOpts
}

type AccessControlledOCR2AggregatorTransactorSession struct {
	Contract     *AccessControlledOCR2AggregatorTransactor
	TransactOpts bind.TransactOpts
}

type AccessControlledOCR2AggregatorRaw struct {
	Contract *AccessControlledOCR2Aggregator
}

type AccessControlledOCR2AggregatorCallerRaw struct {
	Contract *AccessControlledOCR2AggregatorCaller
}

type AccessControlledOCR2AggregatorTransactorRaw struct {
	Contract *AccessControlledOCR2AggregatorTransactor
}

func NewAccessControlledOCR2Aggregator(address common.Address, backend bind.ContractBackend) (*AccessControlledOCR2Aggregator, error) {
	abi, err := abi.JSON(strings.NewReader(AccessControlledOCR2AggregatorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindAccessControlledOCR2Aggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2Aggregator{address: address, abi: abi, AccessControlledOCR2AggregatorCaller: AccessControlledOCR2AggregatorCaller{contract: contract}, AccessControlledOCR2AggregatorTransactor: AccessControlledOCR2AggregatorTransactor{contract: contract}, AccessControlledOCR2AggregatorFilterer: AccessControlledOCR2AggregatorFilterer{contract: contract}}, nil
}

func NewAccessControlledOCR2AggregatorCaller(address common.Address, caller bind.ContractCaller) (*AccessControlledOCR2AggregatorCaller, error) {
	contract, err := bindAccessControlledOCR2Aggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorCaller{contract: contract}, nil
}

func NewAccessControlledOCR2AggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlledOCR2AggregatorTransactor, error) {
	contract, err := bindAccessControlledOCR2Aggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorTransactor{contract: contract}, nil
}

func NewAccessControlledOCR2AggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlledOCR2AggregatorFilterer, error) {
	contract, err := bindAccessControlledOCR2Aggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorFilterer{contract: contract}, nil
}

func bindAccessControlledOCR2Aggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlledOCR2AggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlledOCR2Aggregator.Contract.AccessControlledOCR2AggregatorCaller.contract.Call(opts, result, method, params...)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.AccessControlledOCR2AggregatorTransactor.contract.Transfer(opts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.AccessControlledOCR2AggregatorTransactor.contract.Transact(opts, method, params...)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlledOCR2Aggregator.Contract.contract.Call(opts, result, method, params...)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.contract.Transfer(opts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.contract.Transact(opts, method, params...)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) CheckEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "checkEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) CheckEnabled() (bool, error) {
	return _AccessControlledOCR2Aggregator.Contract.CheckEnabled(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) CheckEnabled() (bool, error) {
	return _AccessControlledOCR2Aggregator.Contract.CheckEnabled(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) Decimals() (uint8, error) {
	return _AccessControlledOCR2Aggregator.Contract.Decimals(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) Decimals() (uint8, error) {
	return _AccessControlledOCR2Aggregator.Contract.Decimals(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) Description() (string, error) {
	return _AccessControlledOCR2Aggregator.Contract.Description(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) Description() (string, error) {
	return _AccessControlledOCR2Aggregator.Contract.Description(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "getAnswer", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.GetAnswer(&_AccessControlledOCR2Aggregator.CallOpts, _roundId)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.GetAnswer(&_AccessControlledOCR2Aggregator.CallOpts, _roundId)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) GetBilling(opts *bind.CallOpts) (GetBilling,

	error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "getBilling")

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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) GetBilling() (GetBilling,

	error) {
	return _AccessControlledOCR2Aggregator.Contract.GetBilling(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) GetBilling() (GetBilling,

	error) {
	return _AccessControlledOCR2Aggregator.Contract.GetBilling(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) GetBillingAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "getBillingAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) GetBillingAccessController() (common.Address, error) {
	return _AccessControlledOCR2Aggregator.Contract.GetBillingAccessController(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) GetBillingAccessController() (common.Address, error) {
	return _AccessControlledOCR2Aggregator.Contract.GetBillingAccessController(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) GetLinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "getLinkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) GetLinkToken() (common.Address, error) {
	return _AccessControlledOCR2Aggregator.Contract.GetLinkToken(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) GetLinkToken() (common.Address, error) {
	return _AccessControlledOCR2Aggregator.Contract.GetLinkToken(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) GetRequesterAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "getRequesterAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) GetRequesterAccessController() (common.Address, error) {
	return _AccessControlledOCR2Aggregator.Contract.GetRequesterAccessController(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) GetRequesterAccessController() (common.Address, error) {
	return _AccessControlledOCR2Aggregator.Contract.GetRequesterAccessController(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (GetRoundData,

	error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "getRoundData", _roundId)

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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) GetRoundData(_roundId *big.Int) (GetRoundData,

	error) {
	return _AccessControlledOCR2Aggregator.Contract.GetRoundData(&_AccessControlledOCR2Aggregator.CallOpts, _roundId)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) GetRoundData(_roundId *big.Int) (GetRoundData,

	error) {
	return _AccessControlledOCR2Aggregator.Contract.GetRoundData(&_AccessControlledOCR2Aggregator.CallOpts, _roundId)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) GetTimestamp(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "getTimestamp", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.GetTimestamp(&_AccessControlledOCR2Aggregator.CallOpts, _roundId)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.GetTimestamp(&_AccessControlledOCR2Aggregator.CallOpts, _roundId)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) GetTransmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "getTransmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) GetTransmitters() ([]common.Address, error) {
	return _AccessControlledOCR2Aggregator.Contract.GetTransmitters(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) GetTransmitters() ([]common.Address, error) {
	return _AccessControlledOCR2Aggregator.Contract.GetTransmitters(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) GetValidatorConfig(opts *bind.CallOpts) (GetValidatorConfig,

	error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "getValidatorConfig")

	outstruct := new(GetValidatorConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.GasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) GetValidatorConfig() (GetValidatorConfig,

	error) {
	return _AccessControlledOCR2Aggregator.Contract.GetValidatorConfig(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) GetValidatorConfig() (GetValidatorConfig,

	error) {
	return _AccessControlledOCR2Aggregator.Contract.GetValidatorConfig(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) HasAccess(opts *bind.CallOpts, _user common.Address, _calldata []byte) (bool, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "hasAccess", _user, _calldata)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) HasAccess(_user common.Address, _calldata []byte) (bool, error) {
	return _AccessControlledOCR2Aggregator.Contract.HasAccess(&_AccessControlledOCR2Aggregator.CallOpts, _user, _calldata)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) HasAccess(_user common.Address, _calldata []byte) (bool, error) {
	return _AccessControlledOCR2Aggregator.Contract.HasAccess(&_AccessControlledOCR2Aggregator.CallOpts, _user, _calldata)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) LatestAnswer() (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.LatestAnswer(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) LatestAnswer() (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.LatestAnswer(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _AccessControlledOCR2Aggregator.Contract.LatestConfigDetails(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _AccessControlledOCR2Aggregator.Contract.LatestConfigDetails(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _AccessControlledOCR2Aggregator.Contract.LatestConfigDigestAndEpoch(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _AccessControlledOCR2Aggregator.Contract.LatestConfigDigestAndEpoch(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) LatestRound() (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.LatestRound(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) LatestRound() (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.LatestRound(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) LatestRoundData(opts *bind.CallOpts) (LatestRoundData,

	error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "latestRoundData")

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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) LatestRoundData() (LatestRoundData,

	error) {
	return _AccessControlledOCR2Aggregator.Contract.LatestRoundData(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) LatestRoundData() (LatestRoundData,

	error) {
	return _AccessControlledOCR2Aggregator.Contract.LatestRoundData(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) LatestTimestamp() (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.LatestTimestamp(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) LatestTimestamp() (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.LatestTimestamp(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) LatestTransmissionDetails(opts *bind.CallOpts) (LatestTransmissionDetails,

	error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "latestTransmissionDetails")

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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) LatestTransmissionDetails() (LatestTransmissionDetails,

	error) {
	return _AccessControlledOCR2Aggregator.Contract.LatestTransmissionDetails(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) LatestTransmissionDetails() (LatestTransmissionDetails,

	error) {
	return _AccessControlledOCR2Aggregator.Contract.LatestTransmissionDetails(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) LinkAvailableForPayment() (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.LinkAvailableForPayment(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.LinkAvailableForPayment(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) MaxAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "maxAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) MaxAnswer() (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.MaxAnswer(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) MaxAnswer() (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.MaxAnswer(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) MinAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "minAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) MinAnswer() (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.MinAnswer(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) MinAnswer() (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.MinAnswer(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) OracleObservationCount(opts *bind.CallOpts, transmitterAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "oracleObservationCount", transmitterAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _AccessControlledOCR2Aggregator.Contract.OracleObservationCount(&_AccessControlledOCR2Aggregator.CallOpts, transmitterAddress)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _AccessControlledOCR2Aggregator.Contract.OracleObservationCount(&_AccessControlledOCR2Aggregator.CallOpts, transmitterAddress)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) OwedPayment(opts *bind.CallOpts, transmitterAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "owedPayment", transmitterAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.OwedPayment(&_AccessControlledOCR2Aggregator.CallOpts, transmitterAddress)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.OwedPayment(&_AccessControlledOCR2Aggregator.CallOpts, transmitterAddress)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) Owner() (common.Address, error) {
	return _AccessControlledOCR2Aggregator.Contract.Owner(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) Owner() (common.Address, error) {
	return _AccessControlledOCR2Aggregator.Contract.Owner(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) TypeAndVersion() (string, error) {
	return _AccessControlledOCR2Aggregator.Contract.TypeAndVersion(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) TypeAndVersion() (string, error) {
	return _AccessControlledOCR2Aggregator.Contract.TypeAndVersion(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOCR2Aggregator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) Version() (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.Version(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorCallerSession) Version() (*big.Int, error) {
	return _AccessControlledOCR2Aggregator.Contract.Version(&_AccessControlledOCR2Aggregator.CallOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "acceptOwnership")
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.AcceptOwnership(&_AccessControlledOCR2Aggregator.TransactOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.AcceptOwnership(&_AccessControlledOCR2Aggregator.TransactOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) AcceptPayeeship(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "acceptPayeeship", transmitter)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.AcceptPayeeship(&_AccessControlledOCR2Aggregator.TransactOpts, transmitter)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.AcceptPayeeship(&_AccessControlledOCR2Aggregator.TransactOpts, transmitter)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "addAccess", _user)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.AddAccess(&_AccessControlledOCR2Aggregator.TransactOpts, _user)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.AddAccess(&_AccessControlledOCR2Aggregator.TransactOpts, _user)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "disableAccessCheck")
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.DisableAccessCheck(&_AccessControlledOCR2Aggregator.TransactOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.DisableAccessCheck(&_AccessControlledOCR2Aggregator.TransactOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "enableAccessCheck")
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.EnableAccessCheck(&_AccessControlledOCR2Aggregator.TransactOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.EnableAccessCheck(&_AccessControlledOCR2Aggregator.TransactOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "removeAccess", _user)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.RemoveAccess(&_AccessControlledOCR2Aggregator.TransactOpts, _user)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.RemoveAccess(&_AccessControlledOCR2Aggregator.TransactOpts, _user)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) RequestNewRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "requestNewRound")
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) RequestNewRound() (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.RequestNewRound(&_AccessControlledOCR2Aggregator.TransactOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) RequestNewRound() (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.RequestNewRound(&_AccessControlledOCR2Aggregator.TransactOpts)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) SetBilling(opts *bind.TransactOpts, maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "setBilling", maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.SetBilling(&_AccessControlledOCR2Aggregator.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.SetBilling(&_AccessControlledOCR2Aggregator.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) SetBillingAccessController(opts *bind.TransactOpts, _billingAccessController common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "setBillingAccessController", _billingAccessController)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.SetBillingAccessController(&_AccessControlledOCR2Aggregator.TransactOpts, _billingAccessController)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.SetBillingAccessController(&_AccessControlledOCR2Aggregator.TransactOpts, _billingAccessController)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "setConfig", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.SetConfig(&_AccessControlledOCR2Aggregator.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.SetConfig(&_AccessControlledOCR2Aggregator.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) SetLinkToken(opts *bind.TransactOpts, linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "setLinkToken", linkToken, recipient)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) SetLinkToken(linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.SetLinkToken(&_AccessControlledOCR2Aggregator.TransactOpts, linkToken, recipient)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) SetLinkToken(linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.SetLinkToken(&_AccessControlledOCR2Aggregator.TransactOpts, linkToken, recipient)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) SetPayees(opts *bind.TransactOpts, transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "setPayees", transmitters, payees)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.SetPayees(&_AccessControlledOCR2Aggregator.TransactOpts, transmitters, payees)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.SetPayees(&_AccessControlledOCR2Aggregator.TransactOpts, transmitters, payees)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) SetRequesterAccessController(opts *bind.TransactOpts, requesterAccessController common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "setRequesterAccessController", requesterAccessController)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) SetRequesterAccessController(requesterAccessController common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.SetRequesterAccessController(&_AccessControlledOCR2Aggregator.TransactOpts, requesterAccessController)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) SetRequesterAccessController(requesterAccessController common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.SetRequesterAccessController(&_AccessControlledOCR2Aggregator.TransactOpts, requesterAccessController)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) SetValidatorConfig(opts *bind.TransactOpts, newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "setValidatorConfig", newValidator, newGasLimit)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) SetValidatorConfig(newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.SetValidatorConfig(&_AccessControlledOCR2Aggregator.TransactOpts, newValidator, newGasLimit)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) SetValidatorConfig(newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.SetValidatorConfig(&_AccessControlledOCR2Aggregator.TransactOpts, newValidator, newGasLimit)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "transferOwnership", to)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.TransferOwnership(&_AccessControlledOCR2Aggregator.TransactOpts, to)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.TransferOwnership(&_AccessControlledOCR2Aggregator.TransactOpts, to)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) TransferPayeeship(opts *bind.TransactOpts, transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "transferPayeeship", transmitter, proposed)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.TransferPayeeship(&_AccessControlledOCR2Aggregator.TransactOpts, transmitter, proposed)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.TransferPayeeship(&_AccessControlledOCR2Aggregator.TransactOpts, transmitter, proposed)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.Transmit(&_AccessControlledOCR2Aggregator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.Transmit(&_AccessControlledOCR2Aggregator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) WithdrawFunds(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "withdrawFunds", recipient, amount)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.WithdrawFunds(&_AccessControlledOCR2Aggregator.TransactOpts, recipient, amount)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.WithdrawFunds(&_AccessControlledOCR2Aggregator.TransactOpts, recipient, amount)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactor) WithdrawPayment(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.contract.Transact(opts, "withdrawPayment", transmitter)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.WithdrawPayment(&_AccessControlledOCR2Aggregator.TransactOpts, transmitter)
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorTransactorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _AccessControlledOCR2Aggregator.Contract.WithdrawPayment(&_AccessControlledOCR2Aggregator.TransactOpts, transmitter)
}

type AccessControlledOCR2AggregatorAddedAccessIterator struct {
	Event *AccessControlledOCR2AggregatorAddedAccess

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorAddedAccessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorAddedAccess)
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
		it.Event = new(AccessControlledOCR2AggregatorAddedAccess)
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

func (it *AccessControlledOCR2AggregatorAddedAccessIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorAddedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorAddedAccess struct {
	User common.Address
	Raw  types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterAddedAccess(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorAddedAccessIterator, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorAddedAccessIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "AddedAccess", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorAddedAccess) (event.Subscription, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorAddedAccess)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "AddedAccess", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseAddedAccess(log types.Log) (*AccessControlledOCR2AggregatorAddedAccess, error) {
	event := new(AccessControlledOCR2AggregatorAddedAccess)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "AddedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorAnswerUpdatedIterator struct {
	Event *AccessControlledOCR2AggregatorAnswerUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorAnswerUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorAnswerUpdated)
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
		it.Event = new(AccessControlledOCR2AggregatorAnswerUpdated)
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

func (it *AccessControlledOCR2AggregatorAnswerUpdatedIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*AccessControlledOCR2AggregatorAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorAnswerUpdatedIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorAnswerUpdated)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseAnswerUpdated(log types.Log) (*AccessControlledOCR2AggregatorAnswerUpdated, error) {
	event := new(AccessControlledOCR2AggregatorAnswerUpdated)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorBillingAccessControllerSetIterator struct {
	Event *AccessControlledOCR2AggregatorBillingAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorBillingAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorBillingAccessControllerSet)
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
		it.Event = new(AccessControlledOCR2AggregatorBillingAccessControllerSet)
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

func (it *AccessControlledOCR2AggregatorBillingAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorBillingAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorBillingAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorBillingAccessControllerSetIterator, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorBillingAccessControllerSetIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "BillingAccessControllerSet", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorBillingAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorBillingAccessControllerSet)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseBillingAccessControllerSet(log types.Log) (*AccessControlledOCR2AggregatorBillingAccessControllerSet, error) {
	event := new(AccessControlledOCR2AggregatorBillingAccessControllerSet)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorBillingSetIterator struct {
	Event *AccessControlledOCR2AggregatorBillingSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorBillingSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorBillingSet)
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
		it.Event = new(AccessControlledOCR2AggregatorBillingSet)
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

func (it *AccessControlledOCR2AggregatorBillingSetIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorBillingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorBillingSet struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
	Raw                       types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterBillingSet(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorBillingSetIterator, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorBillingSetIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "BillingSet", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchBillingSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorBillingSet) (event.Subscription, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorBillingSet)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "BillingSet", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseBillingSet(log types.Log) (*AccessControlledOCR2AggregatorBillingSet, error) {
	event := new(AccessControlledOCR2AggregatorBillingSet)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "BillingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorCheckAccessDisabledIterator struct {
	Event *AccessControlledOCR2AggregatorCheckAccessDisabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorCheckAccessDisabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorCheckAccessDisabled)
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
		it.Event = new(AccessControlledOCR2AggregatorCheckAccessDisabled)
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

func (it *AccessControlledOCR2AggregatorCheckAccessDisabledIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorCheckAccessDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorCheckAccessDisabled struct {
	Raw types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterCheckAccessDisabled(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorCheckAccessDisabledIterator, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorCheckAccessDisabledIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "CheckAccessDisabled", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorCheckAccessDisabled) (event.Subscription, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorCheckAccessDisabled)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseCheckAccessDisabled(log types.Log) (*AccessControlledOCR2AggregatorCheckAccessDisabled, error) {
	event := new(AccessControlledOCR2AggregatorCheckAccessDisabled)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorCheckAccessEnabledIterator struct {
	Event *AccessControlledOCR2AggregatorCheckAccessEnabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorCheckAccessEnabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorCheckAccessEnabled)
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
		it.Event = new(AccessControlledOCR2AggregatorCheckAccessEnabled)
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

func (it *AccessControlledOCR2AggregatorCheckAccessEnabledIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorCheckAccessEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorCheckAccessEnabled struct {
	Raw types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterCheckAccessEnabled(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorCheckAccessEnabledIterator, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorCheckAccessEnabledIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "CheckAccessEnabled", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorCheckAccessEnabled) (event.Subscription, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorCheckAccessEnabled)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseCheckAccessEnabled(log types.Log) (*AccessControlledOCR2AggregatorCheckAccessEnabled, error) {
	event := new(AccessControlledOCR2AggregatorCheckAccessEnabled)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorConfigSetIterator struct {
	Event *AccessControlledOCR2AggregatorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorConfigSet)
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
		it.Event = new(AccessControlledOCR2AggregatorConfigSet)
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

func (it *AccessControlledOCR2AggregatorConfigSetIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorConfigSet struct {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterConfigSet(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorConfigSetIterator, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorConfigSetIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorConfigSet) (event.Subscription, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorConfigSet)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseConfigSet(log types.Log) (*AccessControlledOCR2AggregatorConfigSet, error) {
	event := new(AccessControlledOCR2AggregatorConfigSet)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorLinkTokenSetIterator struct {
	Event *AccessControlledOCR2AggregatorLinkTokenSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorLinkTokenSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorLinkTokenSet)
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
		it.Event = new(AccessControlledOCR2AggregatorLinkTokenSet)
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

func (it *AccessControlledOCR2AggregatorLinkTokenSetIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorLinkTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorLinkTokenSet struct {
	OldLinkToken common.Address
	NewLinkToken common.Address
	Raw          types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterLinkTokenSet(opts *bind.FilterOpts, oldLinkToken []common.Address, newLinkToken []common.Address) (*AccessControlledOCR2AggregatorLinkTokenSetIterator, error) {

	var oldLinkTokenRule []interface{}
	for _, oldLinkTokenItem := range oldLinkToken {
		oldLinkTokenRule = append(oldLinkTokenRule, oldLinkTokenItem)
	}
	var newLinkTokenRule []interface{}
	for _, newLinkTokenItem := range newLinkToken {
		newLinkTokenRule = append(newLinkTokenRule, newLinkTokenItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "LinkTokenSet", oldLinkTokenRule, newLinkTokenRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorLinkTokenSetIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "LinkTokenSet", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchLinkTokenSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorLinkTokenSet, oldLinkToken []common.Address, newLinkToken []common.Address) (event.Subscription, error) {

	var oldLinkTokenRule []interface{}
	for _, oldLinkTokenItem := range oldLinkToken {
		oldLinkTokenRule = append(oldLinkTokenRule, oldLinkTokenItem)
	}
	var newLinkTokenRule []interface{}
	for _, newLinkTokenItem := range newLinkToken {
		newLinkTokenRule = append(newLinkTokenRule, newLinkTokenItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "LinkTokenSet", oldLinkTokenRule, newLinkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorLinkTokenSet)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "LinkTokenSet", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseLinkTokenSet(log types.Log) (*AccessControlledOCR2AggregatorLinkTokenSet, error) {
	event := new(AccessControlledOCR2AggregatorLinkTokenSet)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "LinkTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorNewRoundIterator struct {
	Event *AccessControlledOCR2AggregatorNewRound

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorNewRoundIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorNewRound)
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
		it.Event = new(AccessControlledOCR2AggregatorNewRound)
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

func (it *AccessControlledOCR2AggregatorNewRoundIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*AccessControlledOCR2AggregatorNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorNewRoundIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorNewRound)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseNewRound(log types.Log) (*AccessControlledOCR2AggregatorNewRound, error) {
	event := new(AccessControlledOCR2AggregatorNewRound)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorNewTransmissionIterator struct {
	Event *AccessControlledOCR2AggregatorNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorNewTransmission)
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
		it.Event = new(AccessControlledOCR2AggregatorNewTransmission)
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

func (it *AccessControlledOCR2AggregatorNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorNewTransmission struct {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*AccessControlledOCR2AggregatorNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorNewTransmissionIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorNewTransmission)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseNewTransmission(log types.Log) (*AccessControlledOCR2AggregatorNewTransmission, error) {
	event := new(AccessControlledOCR2AggregatorNewTransmission)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorOraclePaidIterator struct {
	Event *AccessControlledOCR2AggregatorOraclePaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorOraclePaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorOraclePaid)
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
		it.Event = new(AccessControlledOCR2AggregatorOraclePaid)
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

func (it *AccessControlledOCR2AggregatorOraclePaidIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorOraclePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorOraclePaid struct {
	Transmitter common.Address
	Payee       common.Address
	Amount      *big.Int
	LinkToken   common.Address
	Raw         types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*AccessControlledOCR2AggregatorOraclePaidIterator, error) {

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

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorOraclePaidIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "OraclePaid", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorOraclePaid)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseOraclePaid(log types.Log) (*AccessControlledOCR2AggregatorOraclePaid, error) {
	event := new(AccessControlledOCR2AggregatorOraclePaid)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorOwnershipTransferRequestedIterator struct {
	Event *AccessControlledOCR2AggregatorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorOwnershipTransferRequested)
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
		it.Event = new(AccessControlledOCR2AggregatorOwnershipTransferRequested)
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

func (it *AccessControlledOCR2AggregatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AccessControlledOCR2AggregatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorOwnershipTransferRequestedIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorOwnershipTransferRequested)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*AccessControlledOCR2AggregatorOwnershipTransferRequested, error) {
	event := new(AccessControlledOCR2AggregatorOwnershipTransferRequested)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorOwnershipTransferredIterator struct {
	Event *AccessControlledOCR2AggregatorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorOwnershipTransferred)
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
		it.Event = new(AccessControlledOCR2AggregatorOwnershipTransferred)
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

func (it *AccessControlledOCR2AggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AccessControlledOCR2AggregatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorOwnershipTransferredIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorOwnershipTransferred)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*AccessControlledOCR2AggregatorOwnershipTransferred, error) {
	event := new(AccessControlledOCR2AggregatorOwnershipTransferred)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorPayeeshipTransferRequestedIterator struct {
	Event *AccessControlledOCR2AggregatorPayeeshipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorPayeeshipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorPayeeshipTransferRequested)
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
		it.Event = new(AccessControlledOCR2AggregatorPayeeshipTransferRequested)
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

func (it *AccessControlledOCR2AggregatorPayeeshipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorPayeeshipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorPayeeshipTransferRequested struct {
	Transmitter common.Address
	Current     common.Address
	Proposed    common.Address
	Raw         types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*AccessControlledOCR2AggregatorPayeeshipTransferRequestedIterator, error) {

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

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorPayeeshipTransferRequestedIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "PayeeshipTransferRequested", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorPayeeshipTransferRequested)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParsePayeeshipTransferRequested(log types.Log) (*AccessControlledOCR2AggregatorPayeeshipTransferRequested, error) {
	event := new(AccessControlledOCR2AggregatorPayeeshipTransferRequested)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorPayeeshipTransferredIterator struct {
	Event *AccessControlledOCR2AggregatorPayeeshipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorPayeeshipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorPayeeshipTransferred)
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
		it.Event = new(AccessControlledOCR2AggregatorPayeeshipTransferred)
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

func (it *AccessControlledOCR2AggregatorPayeeshipTransferredIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorPayeeshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorPayeeshipTransferred struct {
	Transmitter common.Address
	Previous    common.Address
	Current     common.Address
	Raw         types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*AccessControlledOCR2AggregatorPayeeshipTransferredIterator, error) {

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

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorPayeeshipTransferredIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "PayeeshipTransferred", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorPayeeshipTransferred)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParsePayeeshipTransferred(log types.Log) (*AccessControlledOCR2AggregatorPayeeshipTransferred, error) {
	event := new(AccessControlledOCR2AggregatorPayeeshipTransferred)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorRemovedAccessIterator struct {
	Event *AccessControlledOCR2AggregatorRemovedAccess

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorRemovedAccessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorRemovedAccess)
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
		it.Event = new(AccessControlledOCR2AggregatorRemovedAccess)
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

func (it *AccessControlledOCR2AggregatorRemovedAccessIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorRemovedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorRemovedAccess struct {
	User common.Address
	Raw  types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterRemovedAccess(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorRemovedAccessIterator, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorRemovedAccessIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "RemovedAccess", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorRemovedAccess) (event.Subscription, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorRemovedAccess)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseRemovedAccess(log types.Log) (*AccessControlledOCR2AggregatorRemovedAccess, error) {
	event := new(AccessControlledOCR2AggregatorRemovedAccess)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorRequesterAccessControllerSetIterator struct {
	Event *AccessControlledOCR2AggregatorRequesterAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorRequesterAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorRequesterAccessControllerSet)
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
		it.Event = new(AccessControlledOCR2AggregatorRequesterAccessControllerSet)
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

func (it *AccessControlledOCR2AggregatorRequesterAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorRequesterAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorRequesterAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterRequesterAccessControllerSet(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorRequesterAccessControllerSetIterator, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "RequesterAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorRequesterAccessControllerSetIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "RequesterAccessControllerSet", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchRequesterAccessControllerSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorRequesterAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "RequesterAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorRequesterAccessControllerSet)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "RequesterAccessControllerSet", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseRequesterAccessControllerSet(log types.Log) (*AccessControlledOCR2AggregatorRequesterAccessControllerSet, error) {
	event := new(AccessControlledOCR2AggregatorRequesterAccessControllerSet)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "RequesterAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorRoundRequestedIterator struct {
	Event *AccessControlledOCR2AggregatorRoundRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorRoundRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorRoundRequested)
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
		it.Event = new(AccessControlledOCR2AggregatorRoundRequested)
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

func (it *AccessControlledOCR2AggregatorRoundRequestedIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorRoundRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorRoundRequested struct {
	Requester    common.Address
	ConfigDigest [32]byte
	Epoch        uint32
	Round        uint8
	Raw          types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterRoundRequested(opts *bind.FilterOpts, requester []common.Address) (*AccessControlledOCR2AggregatorRoundRequestedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "RoundRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorRoundRequestedIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "RoundRequested", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchRoundRequested(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorRoundRequested, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "RoundRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorRoundRequested)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "RoundRequested", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseRoundRequested(log types.Log) (*AccessControlledOCR2AggregatorRoundRequested, error) {
	event := new(AccessControlledOCR2AggregatorRoundRequested)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "RoundRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorTransmittedIterator struct {
	Event *AccessControlledOCR2AggregatorTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorTransmitted)
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
		it.Event = new(AccessControlledOCR2AggregatorTransmitted)
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

func (it *AccessControlledOCR2AggregatorTransmittedIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterTransmitted(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorTransmittedIterator, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorTransmittedIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorTransmitted) (event.Subscription, error) {

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorTransmitted)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseTransmitted(log types.Log) (*AccessControlledOCR2AggregatorTransmitted, error) {
	event := new(AccessControlledOCR2AggregatorTransmitted)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AccessControlledOCR2AggregatorValidatorConfigSetIterator struct {
	Event *AccessControlledOCR2AggregatorValidatorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AccessControlledOCR2AggregatorValidatorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOCR2AggregatorValidatorConfigSet)
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
		it.Event = new(AccessControlledOCR2AggregatorValidatorConfigSet)
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

func (it *AccessControlledOCR2AggregatorValidatorConfigSetIterator) Error() error {
	return it.fail
}

func (it *AccessControlledOCR2AggregatorValidatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AccessControlledOCR2AggregatorValidatorConfigSet struct {
	PreviousValidator common.Address
	PreviousGasLimit  uint32
	CurrentValidator  common.Address
	CurrentGasLimit   uint32
	Raw               types.Log
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) FilterValidatorConfigSet(opts *bind.FilterOpts, previousValidator []common.Address, currentValidator []common.Address) (*AccessControlledOCR2AggregatorValidatorConfigSetIterator, error) {

	var previousValidatorRule []interface{}
	for _, previousValidatorItem := range previousValidator {
		previousValidatorRule = append(previousValidatorRule, previousValidatorItem)
	}

	var currentValidatorRule []interface{}
	for _, currentValidatorItem := range currentValidator {
		currentValidatorRule = append(currentValidatorRule, currentValidatorItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.FilterLogs(opts, "ValidatorConfigSet", previousValidatorRule, currentValidatorRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOCR2AggregatorValidatorConfigSetIterator{contract: _AccessControlledOCR2Aggregator.contract, event: "ValidatorConfigSet", logs: logs, sub: sub}, nil
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) WatchValidatorConfigSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorValidatorConfigSet, previousValidator []common.Address, currentValidator []common.Address) (event.Subscription, error) {

	var previousValidatorRule []interface{}
	for _, previousValidatorItem := range previousValidator {
		previousValidatorRule = append(previousValidatorRule, previousValidatorItem)
	}

	var currentValidatorRule []interface{}
	for _, currentValidatorItem := range currentValidator {
		currentValidatorRule = append(currentValidatorRule, currentValidatorItem)
	}

	logs, sub, err := _AccessControlledOCR2Aggregator.contract.WatchLogs(opts, "ValidatorConfigSet", previousValidatorRule, currentValidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AccessControlledOCR2AggregatorValidatorConfigSet)
				if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "ValidatorConfigSet", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2AggregatorFilterer) ParseValidatorConfigSet(log types.Log) (*AccessControlledOCR2AggregatorValidatorConfigSet, error) {
	event := new(AccessControlledOCR2AggregatorValidatorConfigSet)
	if err := _AccessControlledOCR2Aggregator.contract.UnpackLog(event, "ValidatorConfigSet", log); err != nil {
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

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2Aggregator) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _AccessControlledOCR2Aggregator.abi.Events["AddedAccess"].ID:
		return _AccessControlledOCR2Aggregator.ParseAddedAccess(log)
	case _AccessControlledOCR2Aggregator.abi.Events["AnswerUpdated"].ID:
		return _AccessControlledOCR2Aggregator.ParseAnswerUpdated(log)
	case _AccessControlledOCR2Aggregator.abi.Events["BillingAccessControllerSet"].ID:
		return _AccessControlledOCR2Aggregator.ParseBillingAccessControllerSet(log)
	case _AccessControlledOCR2Aggregator.abi.Events["BillingSet"].ID:
		return _AccessControlledOCR2Aggregator.ParseBillingSet(log)
	case _AccessControlledOCR2Aggregator.abi.Events["CheckAccessDisabled"].ID:
		return _AccessControlledOCR2Aggregator.ParseCheckAccessDisabled(log)
	case _AccessControlledOCR2Aggregator.abi.Events["CheckAccessEnabled"].ID:
		return _AccessControlledOCR2Aggregator.ParseCheckAccessEnabled(log)
	case _AccessControlledOCR2Aggregator.abi.Events["ConfigSet"].ID:
		return _AccessControlledOCR2Aggregator.ParseConfigSet(log)
	case _AccessControlledOCR2Aggregator.abi.Events["LinkTokenSet"].ID:
		return _AccessControlledOCR2Aggregator.ParseLinkTokenSet(log)
	case _AccessControlledOCR2Aggregator.abi.Events["NewRound"].ID:
		return _AccessControlledOCR2Aggregator.ParseNewRound(log)
	case _AccessControlledOCR2Aggregator.abi.Events["NewTransmission"].ID:
		return _AccessControlledOCR2Aggregator.ParseNewTransmission(log)
	case _AccessControlledOCR2Aggregator.abi.Events["OraclePaid"].ID:
		return _AccessControlledOCR2Aggregator.ParseOraclePaid(log)
	case _AccessControlledOCR2Aggregator.abi.Events["OwnershipTransferRequested"].ID:
		return _AccessControlledOCR2Aggregator.ParseOwnershipTransferRequested(log)
	case _AccessControlledOCR2Aggregator.abi.Events["OwnershipTransferred"].ID:
		return _AccessControlledOCR2Aggregator.ParseOwnershipTransferred(log)
	case _AccessControlledOCR2Aggregator.abi.Events["PayeeshipTransferRequested"].ID:
		return _AccessControlledOCR2Aggregator.ParsePayeeshipTransferRequested(log)
	case _AccessControlledOCR2Aggregator.abi.Events["PayeeshipTransferred"].ID:
		return _AccessControlledOCR2Aggregator.ParsePayeeshipTransferred(log)
	case _AccessControlledOCR2Aggregator.abi.Events["RemovedAccess"].ID:
		return _AccessControlledOCR2Aggregator.ParseRemovedAccess(log)
	case _AccessControlledOCR2Aggregator.abi.Events["RequesterAccessControllerSet"].ID:
		return _AccessControlledOCR2Aggregator.ParseRequesterAccessControllerSet(log)
	case _AccessControlledOCR2Aggregator.abi.Events["RoundRequested"].ID:
		return _AccessControlledOCR2Aggregator.ParseRoundRequested(log)
	case _AccessControlledOCR2Aggregator.abi.Events["Transmitted"].ID:
		return _AccessControlledOCR2Aggregator.ParseTransmitted(log)
	case _AccessControlledOCR2Aggregator.abi.Events["ValidatorConfigSet"].ID:
		return _AccessControlledOCR2Aggregator.ParseValidatorConfigSet(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (AccessControlledOCR2AggregatorAddedAccess) Topic() common.Hash {
	return common.HexToHash("0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4")
}

func (AccessControlledOCR2AggregatorAnswerUpdated) Topic() common.Hash {
	return common.HexToHash("0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f")
}

func (AccessControlledOCR2AggregatorBillingAccessControllerSet) Topic() common.Hash {
	return common.HexToHash("0x793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912")
}

func (AccessControlledOCR2AggregatorBillingSet) Topic() common.Hash {
	return common.HexToHash("0x0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f")
}

func (AccessControlledOCR2AggregatorCheckAccessDisabled) Topic() common.Hash {
	return common.HexToHash("0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638")
}

func (AccessControlledOCR2AggregatorCheckAccessEnabled) Topic() common.Hash {
	return common.HexToHash("0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480")
}

func (AccessControlledOCR2AggregatorConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (AccessControlledOCR2AggregatorLinkTokenSet) Topic() common.Hash {
	return common.HexToHash("0x4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a")
}

func (AccessControlledOCR2AggregatorNewRound) Topic() common.Hash {
	return common.HexToHash("0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271")
}

func (AccessControlledOCR2AggregatorNewTransmission) Topic() common.Hash {
	return common.HexToHash("0xc797025feeeaf2cd924c99e9205acb8ec04d5cad21c41ce637a38fb6dee6016a")
}

func (AccessControlledOCR2AggregatorOraclePaid) Topic() common.Hash {
	return common.HexToHash("0xd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c")
}

func (AccessControlledOCR2AggregatorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (AccessControlledOCR2AggregatorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (AccessControlledOCR2AggregatorPayeeshipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367")
}

func (AccessControlledOCR2AggregatorPayeeshipTransferred) Topic() common.Hash {
	return common.HexToHash("0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3")
}

func (AccessControlledOCR2AggregatorRemovedAccess) Topic() common.Hash {
	return common.HexToHash("0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1")
}

func (AccessControlledOCR2AggregatorRequesterAccessControllerSet) Topic() common.Hash {
	return common.HexToHash("0x27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634")
}

func (AccessControlledOCR2AggregatorRoundRequested) Topic() common.Hash {
	return common.HexToHash("0x41e3990591fd372502daa15842da15bc7f41c75309ab3ff4f56f1848c178825c")
}

func (AccessControlledOCR2AggregatorTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (AccessControlledOCR2AggregatorValidatorConfigSet) Topic() common.Hash {
	return common.HexToHash("0xb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541")
}

func (_AccessControlledOCR2Aggregator *AccessControlledOCR2Aggregator) Address() common.Address {
	return _AccessControlledOCR2Aggregator.address
}

type AccessControlledOCR2AggregatorInterface interface {
	CheckEnabled(opts *bind.CallOpts) (bool, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	Description(opts *bind.CallOpts) (string, error)

	GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error)

	GetBilling(opts *bind.CallOpts) (GetBilling,

		error)

	GetBillingAccessController(opts *bind.CallOpts) (common.Address, error)

	GetLinkToken(opts *bind.CallOpts) (common.Address, error)

	GetRequesterAccessController(opts *bind.CallOpts) (common.Address, error)

	GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (GetRoundData,

		error)

	GetTimestamp(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error)

	GetTransmitters(opts *bind.CallOpts) ([]common.Address, error)

	GetValidatorConfig(opts *bind.CallOpts) (GetValidatorConfig,

		error)

	HasAccess(opts *bind.CallOpts, _user common.Address, _calldata []byte) (bool, error)

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

	AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error)

	DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error)

	EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error)

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

	FilterAddedAccess(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorAddedAccessIterator, error)

	WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorAddedAccess) (event.Subscription, error)

	ParseAddedAccess(log types.Log) (*AccessControlledOCR2AggregatorAddedAccess, error)

	FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*AccessControlledOCR2AggregatorAnswerUpdatedIterator, error)

	WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error)

	ParseAnswerUpdated(log types.Log) (*AccessControlledOCR2AggregatorAnswerUpdated, error)

	FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorBillingAccessControllerSetIterator, error)

	WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorBillingAccessControllerSet) (event.Subscription, error)

	ParseBillingAccessControllerSet(log types.Log) (*AccessControlledOCR2AggregatorBillingAccessControllerSet, error)

	FilterBillingSet(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorBillingSetIterator, error)

	WatchBillingSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorBillingSet) (event.Subscription, error)

	ParseBillingSet(log types.Log) (*AccessControlledOCR2AggregatorBillingSet, error)

	FilterCheckAccessDisabled(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorCheckAccessDisabledIterator, error)

	WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorCheckAccessDisabled) (event.Subscription, error)

	ParseCheckAccessDisabled(log types.Log) (*AccessControlledOCR2AggregatorCheckAccessDisabled, error)

	FilterCheckAccessEnabled(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorCheckAccessEnabledIterator, error)

	WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorCheckAccessEnabled) (event.Subscription, error)

	ParseCheckAccessEnabled(log types.Log) (*AccessControlledOCR2AggregatorCheckAccessEnabled, error)

	FilterConfigSet(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*AccessControlledOCR2AggregatorConfigSet, error)

	FilterLinkTokenSet(opts *bind.FilterOpts, oldLinkToken []common.Address, newLinkToken []common.Address) (*AccessControlledOCR2AggregatorLinkTokenSetIterator, error)

	WatchLinkTokenSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorLinkTokenSet, oldLinkToken []common.Address, newLinkToken []common.Address) (event.Subscription, error)

	ParseLinkTokenSet(log types.Log) (*AccessControlledOCR2AggregatorLinkTokenSet, error)

	FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*AccessControlledOCR2AggregatorNewRoundIterator, error)

	WatchNewRound(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error)

	ParseNewRound(log types.Log) (*AccessControlledOCR2AggregatorNewRound, error)

	FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*AccessControlledOCR2AggregatorNewTransmissionIterator, error)

	WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error)

	ParseNewTransmission(log types.Log) (*AccessControlledOCR2AggregatorNewTransmission, error)

	FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*AccessControlledOCR2AggregatorOraclePaidIterator, error)

	WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error)

	ParseOraclePaid(log types.Log) (*AccessControlledOCR2AggregatorOraclePaid, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AccessControlledOCR2AggregatorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*AccessControlledOCR2AggregatorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AccessControlledOCR2AggregatorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*AccessControlledOCR2AggregatorOwnershipTransferred, error)

	FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*AccessControlledOCR2AggregatorPayeeshipTransferRequestedIterator, error)

	WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error)

	ParsePayeeshipTransferRequested(log types.Log) (*AccessControlledOCR2AggregatorPayeeshipTransferRequested, error)

	FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*AccessControlledOCR2AggregatorPayeeshipTransferredIterator, error)

	WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error)

	ParsePayeeshipTransferred(log types.Log) (*AccessControlledOCR2AggregatorPayeeshipTransferred, error)

	FilterRemovedAccess(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorRemovedAccessIterator, error)

	WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorRemovedAccess) (event.Subscription, error)

	ParseRemovedAccess(log types.Log) (*AccessControlledOCR2AggregatorRemovedAccess, error)

	FilterRequesterAccessControllerSet(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorRequesterAccessControllerSetIterator, error)

	WatchRequesterAccessControllerSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorRequesterAccessControllerSet) (event.Subscription, error)

	ParseRequesterAccessControllerSet(log types.Log) (*AccessControlledOCR2AggregatorRequesterAccessControllerSet, error)

	FilterRoundRequested(opts *bind.FilterOpts, requester []common.Address) (*AccessControlledOCR2AggregatorRoundRequestedIterator, error)

	WatchRoundRequested(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorRoundRequested, requester []common.Address) (event.Subscription, error)

	ParseRoundRequested(log types.Log) (*AccessControlledOCR2AggregatorRoundRequested, error)

	FilterTransmitted(opts *bind.FilterOpts) (*AccessControlledOCR2AggregatorTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*AccessControlledOCR2AggregatorTransmitted, error)

	FilterValidatorConfigSet(opts *bind.FilterOpts, previousValidator []common.Address, currentValidator []common.Address) (*AccessControlledOCR2AggregatorValidatorConfigSetIterator, error)

	WatchValidatorConfigSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOCR2AggregatorValidatorConfigSet, previousValidator []common.Address, currentValidator []common.Address) (event.Subscription, error)

	ParseValidatorConfigSet(log types.Log) (*AccessControlledOCR2AggregatorValidatorConfigSet, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}

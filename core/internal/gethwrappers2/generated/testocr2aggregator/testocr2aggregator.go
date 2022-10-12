// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testocr2aggregator

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

var TestOCR2AggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"int192\",\"name\":\"_minAnswer\",\"type\":\"int192\"},{\"internalType\":\"int192\",\"name\":\"_maxAnswer\",\"type\":\"int192\"},{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"},{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_requesterAdminAccessController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"AddedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"BillingAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"BillingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CheckAccessDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CheckAccessEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"oldLinkToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"newLinkToken\",\"type\":\"address\"}],\"name\":\"LinkTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int192\",\"name\":\"answer\",\"type\":\"int192\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"observationsTimestamp\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int192[]\",\"name\":\"observations\",\"type\":\"int192[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"observers\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"int192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"int192\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RemovedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"RequesterAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"}],\"name\":\"RoundRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"previousValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousGasLimit\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"currentValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"currentGasLimit\",\"type\":\"uint32\"}],\"name\":\"ValidatorConfigSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"addAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableAccessCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableAccessCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBillingAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfigDigest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinkToken\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequesterAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorConfig\",\"outputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"hasAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTransmissionDetails\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"},{\"internalType\":\"int192\",\"name\":\"latestAnswer_\",\"type\":\"int192\"},{\"internalType\":\"uint64\",\"name\":\"latestTimestamp_\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAnswer\",\"outputs\":[{\"internalType\":\"int192\",\"name\":\"\",\"type\":\"int192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAnswer\",\"outputs\":[{\"internalType\":\"int192\",\"name\":\"\",\"type\":\"int192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"oracleObservationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"owedPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"removeAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestNewRound\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"setBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"}],\"name\":\"setBillingAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"setLinkToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"requesterAccessController\",\"type\":\"address\"}],\"name\":\"setRequesterAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"newValidator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newGasLimit\",\"type\":\"uint32\"}],\"name\":\"setValidatorConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"testBurnLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"testDecodeReport\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"int192[]\",\"name\":\"\",\"type\":\"int192[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transmitter\",\"type\":\"address\"}],\"name\":\"testPayee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"txGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reasonableGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumGasPrice\",\"type\":\"uint256\"}],\"name\":\"testReimbursementGasPriceGwei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"paymentJuels\",\"type\":\"uint96\"}],\"name\":\"testSetTransmitterPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testTotalLinkDue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"linkDue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceGwei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callDataGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountingGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftGas\",\"type\":\"uint256\"}],\"name\":\"testTransmitterGasCostWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005ef838038062005ef883398101604081905262000034916200050c565b8484848484600060405180604001604052806004815260200163151154d560e21b8152508686868686868633806000806001600160a01b0316826001600160a01b03161415620000a15760405162461bcd60e51b81526004016200009890620005d5565b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000d457620000d481620001c4565b5050601180546001600160a01b0319166001600160a01b038a169081179091556040519091506000907f4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a908290a36200012d8462000241565b7fff0000000000000000000000000000000000000000000000000000000000000060f883901b1660c05280516200016c9060109060208401906200044e565b506200017883620002b5565b6200018560008062000326565b50505050601791820b820b604090811b60805290820b90910b901b60a05250506015805460ff1916600117905550620006b09950505050505050505050565b6001600160a01b038116331415620001f05760405162461bcd60e51b815260040162000098906200060c565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6012546001600160a01b039081169082168114620002b157601280546001600160a01b0319166001600160a01b0384161790556040517f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d4891290620002a8908390859062000584565b60405180910390a15b5050565b620002bf6200041f565b600f546001600160a01b039081169082168114620002b157600f80546001600160a01b0319166001600160a01b0384161790556040517f27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae63490620002a8908390859062000584565b620003306200041f565b60408051808201909152600e546001600160a01b03808216808452600160a01b90920463ffffffff16602084015284161415806200037e57508163ffffffff16816020015163ffffffff1614155b156200041a576040805180820182526001600160a01b0385811680835263ffffffff86166020938401819052600e80546001600160a01b031916831763ffffffff60a01b1916600160a01b9092029190911790558451928501519351909392909116917fb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca154191620004119190879062000643565b60405180910390a35b505050565b6000546001600160a01b031633146200044c5760405162461bcd60e51b815260040162000098906200059e565b565b8280546200045c906200065a565b90600052602060002090601f016020900481019282620004805760008555620004cb565b82601f106200049b57805160ff1916838001178555620004cb565b82800160010185558215620004cb579182015b82811115620004cb578251825591602001919060010190620004ae565b50620004d9929150620004dd565b5090565b5b80821115620004d95760008155600101620004de565b8051601781900b81146200050757600080fd5b919050565b600080600080600060a0868803121562000524578081fd5b8551620005318162000697565b94506200054160208701620004f4565b93506200055160408701620004f4565b92506060860151620005638162000697565b6080870151909250620005768162000697565b809150509295509295909350565b6001600160a01b0392831681529116602082015260400190565b60208082526016908201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604082015260600190565b60208082526018908201527f43616e6e6f7420736574206f776e657220746f207a65726f0000000000000000604082015260600190565b60208082526017908201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604082015260600190565b63ffffffff92831681529116602082015260400190565b6002810460018216806200066f57607f821691505b602082108114156200069157634e487b7160e01b600052602260045260246000fd5b50919050565b6001600160a01b0381168114620006ad57600080fd5b50565b60805160401c60a05160401c60c05160f81c6157f3620007056000396000610947015260008181610fbc015281816120c501526136090152600081816108d1015281816120a401526135dc01526157f36000f3fe608060405234801561001057600080fd5b506004361061036d5760003560e01c80638da5cb5b116101d3578063c4c92b3711610104578063e5fe4577116100a2578063eb5dcd6c1161007c578063eb5dcd6c14610700578063f2fde38b14610713578063fbffd2c114610726578063feaf968c146107395761036d565b8063e5fe4577146106cc578063e76d5168146106e5578063eb457163146106ed5761036d565b8063dc7f0124116100de578063dc7f01241461067e578063e259c1f614610686578063e3d0e71214610699578063e4902f82146106ac5761036d565b8063c4c92b3714610666578063d09dc3391461066e578063daffc4b5146106765761036d565b8063a118f24911610171578063b1dc65a41161014b578063b1dc65a41461061a578063b5ab58dc1461062d578063b633620c14610640578063c1075329146106535761036d565b8063a118f249146105dd578063afcb95d7146105f0578063b121e147146106075761036d565b80639bd2c0b1116101ad5780639bd2c0b11461058e5780639c849b30146105a45780639e3ceeab146105b75780639eb6e060146105ca5761036d565b80638da5cb5b1461054057806398e5b12a146105555780639a6fc8f51461056a5761036d565b8063643dc105116102ad5780637284e4161161024b57806381ff70481161022557806381ff7048146104fb5780638205bf6a146105125780638823da6c1461051a5780638ac28d5a1461052d5761036d565b80637284e416146104e357806379ba5097146104eb5780638038e4a1146104f35761036d565b806366cfeaf11161028757806366cfeaf1146104a05780636b14daf8146104a857806370031468146104c857806370da2f67146104db5761036d565b8063643dc10514610470578063666cab8d14610483578063668a0f02146104985761036d565b8063313ce5671161031a5780634fb17470116102f45780634fb174701461043a57806350d25bcd1461044d57806354fd4d50146104555780635765b76a1461045d5761036d565b8063313ce567146103f057806333e8f9f0146104055780633b5cdfa2146104185761036d565b8063181f5a771161034b578063181f5a77146103ad57806322adbc78146103c257806329937268146103d75761036d565b80630a756983146103725780630eafb25b1461037c578063102a474b146103a5575b600080fd5b61037a610741565b005b61038f61038a3660046141c3565b61078a565b60405161039c9190614826565b60405180910390f35b61038f6108a0565b6103b56108af565b60405161039c9190614a7d565b6103ca6108cf565b60405161039c91906149f8565b6103df6108f3565b60405161039c95949392919061549e565b6103f8610945565b60405161039c919061551c565b61038f6104133660046145c3565b610969565b61042b6104263660046144f6565b610982565b60405161039c93929190615435565b61037a610448366004614556565b6109ac565b61038f610bd2565b61038f610c39565b61038f61046b366004614598565b610c3e565b61037a61047e3660046146c5565b610c53565b61048b610e1f565b60405161039c91906148f3565b61038f610e81565b61038f610ee8565b6104bb6104b6366004614217565b610eee565b60405161039c9190614906565b61037a6104d6366004614568565b610f16565b6103ca610fba565b6103b5610fde565b61037a611045565b61037a6110c6565b610503611110565b60405161039c9392919061547d565b61038f61112c565b61037a6105283660046141c3565b611193565b61037a61053b3660046141c3565b611217565b610548611259565b60405161039c9190614887565b61055d611268565b60405161039c91906154d2565b61057d61057836600461473c565b6113ab565b60405161039c9594939291906154e9565b61059661142d565b60405161039c9291906149d9565b61037a6105b23660046142c8565b611460565b61037a6105c53660046141c3565b6115fa565b6105486105d83660046141c3565b611672565b61037a6105eb3660046141c3565b611690565b6105f861170b565b60405161039c93929190614911565b61037a6106153660046141c3565b611726565b61037a6106283660046143f9565b6117d2565b61038f61063b366004614568565b611b97565b61038f61064e366004614568565b611bff565b61037a610661366004614265565b611c67565b610548611e92565b61038f611ea1565b610548611f45565b6104bb611f54565b61037a610694366004614290565b611f5d565b61037a6106a7366004614331565b611fdb565b6106bf6106ba3660046141c3565b61264e565b60405161039c919061539e565b6106d4612713565b60405161039c959493929190614968565b61054861278c565b61037a6106fb366004614529565b61279b565b61037a61070e3660046141df565b6128a6565b61037a6107213660046141c3565b61297f565b61037a6107343660046141c3565b612990565b61057d6129a1565b610749612a1f565b60155460ff1615610788576015805460ff191690556040517f3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f53963890600090a15b565b6001600160a01b03811660009081526002602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b031691810191909152906107ed57600091505061089b565b600b546020820151600091600160901b900463ffffffff169060069060ff16601f811061082a57634e487b7160e01b600052603260045260246000fd5b600881049190910154600b5461085d926007166004026101000a90910463ffffffff90811691600160301b9004166156d0565b63ffffffff1661086d91906155fd565b61087b90633b9aca006155fd565b905081604001516001600160601b0316816108969190615578565b925050505b919050565b60006108aa612a49565b905090565b60606040518060600160405280602a81526020016157bd602a9139905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b600b546a0100000000000000000000810463ffffffff908116926e0100000000000000000000000000008304821692600160901b8104831692600160b01b82041691600160d01b90910462ffffff1690565b7f000000000000000000000000000000000000000000000000000000000000000081565b60006109788686868686612c2f565b9695505050505050565b6000606080600061099285612c63565b805160208201516040909201519097919650945092505050565b6109b4612a1f565b6011546001600160a01b039081169083168114156109d25750610bce565b6040516370a0823160e01b81526001600160a01b038416906370a08231906109fe903090600401614887565b60206040518083038186803b158015610a1657600080fd5b505afa158015610a2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4e9190614580565b50610a57612ce8565b6040516370a0823160e01b81526000906001600160a01b038316906370a0823190610a86903090600401614887565b60206040518083038186803b158015610a9e57600080fd5b505afa158015610ab2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad69190614580565b60405163a9059cbb60e01b81529091506001600160a01b0383169063a9059cbb90610b0790869085906004016148da565b602060405180830381600087803b158015610b2157600080fd5b505af1158015610b35573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5991906144d6565b610b7e5760405162461bcd60e51b8152600401610b7590614ceb565b60405180910390fd5b601180546001600160a01b0319166001600160a01b0386811691821790925560405190918416907f4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a90600090a350505b5050565b6000610c15336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610eee92505050565b610c315760405162461bcd60e51b8152600401610b7590614eda565b6108aa6130d5565b600681565b6000610c4b8484846130fe565b949350505050565b6012546001600160a01b0316610c67611259565b6001600160a01b0316336001600160a01b03161480610d025750604051630d629b5f60e31b81526001600160a01b03821690636b14daf890610cb2903390600090369060040161489b565b60206040518083038186803b158015610cca57600080fd5b505afa158015610cde573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d0291906144d6565b610d1e5760405162461bcd60e51b8152600401610b7590614fb6565b610d26612ce8565b85600b600001600a6101000a81548163ffffffff021916908363ffffffff16021790555084600b600001600e6101000a81548163ffffffff021916908363ffffffff16021790555083600b60000160126101000a81548163ffffffff021916908363ffffffff16021790555082600b60000160166101000a81548163ffffffff021916908363ffffffff16021790555081600b600001601a6101000a81548162ffffff021916908362ffffff1602179055507f0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f8686868686604051610e0f95949392919061549e565b60405180910390a1505050505050565b60606005805480602002602001604051908101604052809291908181526020018280548015610e7757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e59575b5050505050905090565b6000610ec4336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610eee92505050565b610ee05760405162461bcd60e51b8152600401610b7590614eda565b6108aa613124565b600a5490565b6000610efa8383613137565b80610f0d57506001600160a01b03831632145b90505b92915050565b60115460405163a9059cbb60e01b81526001600160a01b039091169063a9059cbb90610f499060019085906004016148da565b602060405180830381600087803b158015610f6357600080fd5b505af1158015610f77573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f9b91906144d6565b610fb75760405162461bcd60e51b8152600401610b759061527d565b50565b7f000000000000000000000000000000000000000000000000000000000000000081565b6060611021336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610eee92505050565b61103d5760405162461bcd60e51b8152600401610b7590614eda565b6108aa613167565b6001546001600160a01b0316331461106f5760405162461bcd60e51b8152600401610b7590614ac7565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6110ce612a1f565b60155460ff16610788576015805460ff191660011790556040517faebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c348090600090a1565b600d54600a5463ffffffff808316926401000000009004169192565b600061116f336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610eee92505050565b61118b5760405162461bcd60e51b8152600401610b7590614eda565b6108aa6131f0565b61119b612a1f565b6001600160a01b03811660009081526016602052604090205460ff1615610fb7576001600160a01b03811660009081526016602052604090819020805460ff19169055517f3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d19061120c908390614887565b60405180910390a150565b6001600160a01b038181166000908152601360205260409020541633146112505760405162461bcd60e51b8152600401610b7590615059565b610fb78161321c565b6000546001600160a01b031690565b6000611272611259565b6001600160a01b0316336001600160a01b031614806113115750600f54604051630d629b5f60e31b81526001600160a01b0390911690636b14daf8906112c1903390600090369060040161489b565b60206040518083038186803b1580156112d957600080fd5b505afa1580156112ed573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061131191906144d6565b61132d5760405162461bcd60e51b8152600401610b7590614bd8565b600b54600a54604051610100830464ffffffffff811693600160301b900463ffffffff9081169333937f41e3990591fd372502daa15842da15bc7f41c75309ab3ff4f56f1848c178825c9361138b9360089190911c16908790614949565b60405180910390a261139e816001615590565b63ffffffff169250505090565b60008060008060006113f4336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610eee92505050565b6114105760405162461bcd60e51b8152600401610b7590614eda565b6114198661342c565b945094509450945094505b91939590929450565b60408051808201909152600e546001600160a01b038116808352600160a01b90910463ffffffff16602090920182905291565b611468612a1f565b8281146114875760405162461bcd60e51b8152600401610b75906151a3565b60005b838110156115f35760008585838181106114b457634e487b7160e01b600052603260045260246000fd5b90506020020160208101906114c991906141c3565b905060008484848181106114ed57634e487b7160e01b600052603260045260246000fd5b905060200201602081019061150291906141c3565b6001600160a01b03808416600090815260136020526040902054919250168015808061153f5750826001600160a01b0316826001600160a01b0316145b61155b5760405162461bcd60e51b8152600401610b7590614f7f565b6001600160a01b03848116600090815260136020526040902080546001600160a01b031916858316908117909155908316146115dc57826001600160a01b0316826001600160a01b0316856001600160a01b03167f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b360405160405180910390a45b5050505080806115eb9061572a565b91505061148a565b5050505050565b611602612a1f565b600f546001600160a01b039081169082168114610bce57600f80546001600160a01b0319166001600160a01b0384161790556040517f27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae6349061166690839085906149bf565b60405180910390a15050565b6001600160a01b039081166000908152601360205260409020541690565b611698612a1f565b6001600160a01b03811660009081526016602052604090205460ff16610fb7576001600160a01b03811660009081526016602052604090819020805460ff19166001179055517f87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db49061120c908390614887565b600a54600b5460009261010090910460081c63ffffffff1690565b6001600160a01b0381811660009081526014602052604090205416331461175f5760405162461bcd60e51b8152600401610b7590614a90565b6001600160a01b0381811660008181526013602090815260408083208054336001600160a01b031980831682179093556014909452828520805490921690915590519416939092849290917f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b39190a45050565b60005a604080516101008082018352600b5460ff8116835290810464ffffffffff9081166020808501829052600160301b840463ffffffff908116968601969096526a01000000000000000000008404861660608601526e010000000000000000000000000000840486166080860152600160901b8404861660a0860152600160b01b840490951660c0850152600160d01b90920462ffffff1660e08401529394509092918c01359182161161189a5760405162461bcd60e51b8152600401610b7590614ba1565b3360009081526002602052604090205460ff166118c95760405162461bcd60e51b8152600401610b7590614feb565b600a548b35146118eb5760405162461bcd60e51b8152600401610b7590615135565b6118f98a8a8a8a8a8a6134c1565b81516119069060016155b8565b60ff1687146119275760405162461bcd60e51b8152600401610b7590615022565b8685146119465760405162461bcd60e51b8152600401610b75906152b4565b60008a8a60405161195892919061482f565b60405190819003812061196f918e9060200161492f565b6040516020818303038152906040528051906020012090506000611991613eec565b60005b8a811015611aeb5760006001858a84602081106119c157634e487b7160e01b600052603260045260246000fd5b6119ce91901a601b6155b8565b8f8f868181106119ee57634e487b7160e01b600052603260045260246000fd5b905060200201358e8e87818110611a1557634e487b7160e01b600052603260045260246000fd5b9050602002013560405160008152602001604052604051611a3994939291906149a1565b6020604051602081039080840390855afa158015611a5b573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526003602090815290849020838501909452925460ff8082161515808552610100909204169383019390935290955092509050611ac45760405162461bcd60e51b8152600401610b75906150c7565b826020015160080260ff166001901b84019350508080611ae39061572a565b915050611994565b5081827e010101010101010101010101010101010101010101010101010101010101011614611b2c5760405162461bcd60e51b8152600401610b7590614d59565b5060009150611b7b9050838d836020020135848e8e8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061352e92505050565b9050611b89838286336139b0565b505050505050505050505050565b6000611bda336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610eee92505050565b611bf65760405162461bcd60e51b8152600401610b7590614eda565b610f1082613ae6565b6000611c42336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610eee92505050565b611c5e5760405162461bcd60e51b8152600401610b7590614eda565b610f1082613b1c565b611c6f611259565b6001600160a01b0316336001600160a01b03161480611d0e5750601254604051630d629b5f60e31b81526001600160a01b0390911690636b14daf890611cbe903390600090369060040161489b565b60206040518083038186803b158015611cd657600080fd5b505afa158015611cea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d0e91906144d6565b611d2a5760405162461bcd60e51b8152600401610b7590614fb6565b6000611d34612a49565b6011546040516370a0823160e01b81529192506000916001600160a01b03909116906370a0823190611d6a903090600401614887565b60206040518083038186803b158015611d8257600080fd5b505afa158015611d96573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dba9190614580565b905081811015611ddc5760405162461bcd60e51b8152600401610b7590615090565b6011546001600160a01b031663a9059cbb85611e01611dfb86866156b9565b87613b55565b6040518363ffffffff1660e01b8152600401611e1e9291906148da565b602060405180830381600087803b158015611e3857600080fd5b505af1158015611e4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e7091906144d6565b611e8c5760405162461bcd60e51b8152600401610b759061520f565b50505050565b6012546001600160a01b031690565b6011546040516370a0823160e01b815260009182916001600160a01b03909116906370a0823190611ed6903090600401614887565b60206040518083038186803b158015611eee57600080fd5b505afa158015611f02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f269190614580565b90506000611f32612a49565b9050611f3e8183615645565b9250505090565b600f546001600160a01b031690565b60155460ff1681565b6001600160a01b03821660009081526002602052604090205460ff16611f955760405162461bcd60e51b8152600401610b7590614f11565b6001600160a01b03909116600090815260026020526040902080546001600160601b0390921662010000026dffffffffffffffffffffffff000019909216919091179055565b611fe3612a1f565b601f865111156120055760405162461bcd60e51b8152600401610b7590614c46565b84518651146120265760405162461bcd60e51b8152600401610b7590614c7d565b855161203385600361561c565b60ff16106120535760405162461bcd60e51b8152600401610b75906151d8565b61205f8460ff16613b6c565b82511561207e5760405162461bcd60e51b8152600401610b7590614f48565b60006040518060c001604052808881526020018781526020018660ff16815260200160017f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006040516020016120f69392919061483f565b60408051601f1981840301815291815290825267ffffffffffffffff8616602083015201839052600b805465ffffffffff00191690559050612136612ce8565b60045460005b818110156122035760006004828154811061216757634e487b7160e01b600052603260045260246000fd5b6000918252602082200154600580546001600160a01b03909216935090849081106121a257634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b039485168352600382526040808420805461ffff1916905594168252600290529190912080546dffffffffffffffffffffffffffff1916905550806121fb8161572a565b91505061213c565b5061221060046000613f03565b61221c60056000613f03565b60005b82515181101561246f57600360008460000151838151811061225157634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff16156122955760405162461bcd60e51b8152600401610b7590614dfe565b604080518082019091526001815260ff8216602082015283518051600391600091859081106122d457634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252818101929092526040016000908120835181549484015160ff166101000261ff001991151560ff19909616959095171693909317909255840151805160029291908490811061234a57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff161561238e5760405162461bcd60e51b8152600401610b7590614b35565b60405180606001604052806001151581526020018260ff16815260200160006001600160601b031681525060026000856020015184815181106123e157634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b031682528181019290925260409081016000208351815493850151949092015160ff199093169115159190911761ff00191661010060ff90941693909302929092176dffffffffffffffffffffffff00001916620100006001600160601b0390921691909102179055806124678161572a565b91505061221f565b508151805161248691600491602090910190613f21565b50602080830151805161249d926005920190613f21565b506040820151600b805460ff191660ff909216919091179055600d805467ffffffff0000000019811664010000000063ffffffff438116820292909217808555920481169260019290916000916124f691859116615590565b92506101000a81548163ffffffff021916908363ffffffff1602179055506125554630600d60009054906101000a900463ffffffff1663ffffffff1686600001518760200151886040015189606001518a608001518b60a00151613b8c565b600a819055600d5484516020860151604080880151606089015160808a015160a08b015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05986125be988b98919763ffffffff9091169691959094909390929091906153af565b60405180910390a1600b54600160301b900463ffffffff1660005b8451518110156126415781600682601f811061260557634e487b7160e01b600052603260045260246000fd5b600891828204019190066004026101000a81548163ffffffff021916908363ffffffff16021790555080806126399061572a565b9150506125d9565b5050505050505050505050565b6001600160a01b03811660009081526002602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b031691810191909152906126b157600091505061089b565b6006816020015160ff16601f81106126d957634e487b7160e01b600052603260045260246000fd5b600881049190910154600b5461270c926007166004026101000a90910463ffffffff90811691600160301b9004166156d0565b9392505050565b6000808080803332146127385760405162461bcd60e51b8152600401610b7590614d22565b5050600a54600b5463ffffffff600160301b820481166000908152600c60205260409020549296610100909204600881901c8216965064ffffffffff169450601783900b9350600160e01b90920490911690565b6011546001600160a01b031690565b6127a3612a1f565b60408051808201909152600e546001600160a01b03808216808452600160a01b90920463ffffffff16602084015284161415806127f057508163ffffffff16816020015163ffffffff1614155b156128a1576040805180820182526001600160a01b0385811680835263ffffffff86166020938401819052600e80546001600160a01b03191683177fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff16600160a01b9092029190911790558451928501519351909392909116917fb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca15419161289891908790615466565b60405180910390a35b505050565b6001600160a01b038281166000908152601360205260409020541633146128df5760405162461bcd60e51b8152600401610b7590614afe565b336001600160a01b03821614156129085760405162461bcd60e51b8152600401610b759061516c565b6001600160a01b03808316600090815260146020526040902080548383166001600160a01b0319821681179092559091169081146128a1576040516001600160a01b038084169133918616907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e3836790600090a4505050565b612987612a1f565b610fb781613c19565b612998612a1f565b610fb781613c93565b60008060008060006129ea336000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610eee92505050565b612a065760405162461bcd60e51b8152600401610b7590614eda565b612a0e613cf7565b945094509450945094509091929394565b6000546001600160a01b031633146107885760405162461bcd60e51b8152600401610b7590614c0f565b6000806005805480602002602001604051908101604052809291908181526020018280548015612aa257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612a84575b50508351600b54604080516103e08101918290529697509195600160301b90910463ffffffff169450600093509150600690601f908285855b82829054906101000a900463ffffffff1663ffffffff1681526020019060040190602082600301049283019260010382029150808411612adb5790505050505050905060005b83811015612b7c578181601f8110612b4957634e487b7160e01b600052603260045260246000fd5b6020020151612b5890846156d0565b612b689063ffffffff1687615578565b955080612b748161572a565b915050612b21565b50600b54612b9b90600160901b900463ffffffff16633b9aca006155fd565b612ba590866155fd565b945060005b83811015612c275760026000868381518110612bd657634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0316825281019190915260400160002054612c13906201000090046001600160601b031687615578565b955080612c1f8161572a565b915050612baa565b505050505090565b600081861015612c515760405162461bcd60e51b8152600401610b7590614b6c565b50633b9aca0094039190910101020290565b612c6b613f86565b6000806060600085806020019051810190612c8691906145fd565b92965090945092509050612c9a8683613d5f565b8151604051600090612cb0908690602001614826565b60408051918152928152825160808101845263ffffffff909716875260208701525084019190915260170b6060830152509050919050565b601154600b54604080516103e08101918290526001600160a01b0390931692600160301b90920463ffffffff1691600091600690601f908285855b82829054906101000a900463ffffffff1663ffffffff1681526020019060040190602082600301049283019260010382029150808411612d235790505050505050905060006005805480602002602001604051908101604052809291908181526020018280548015612dbe57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612da0575b5050505050905060005b81518110156130c757600060026000848481518110612df757634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160029054906101000a90046001600160601b03166001600160601b03169050600060026000858581518110612e6757634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160026101000a8154816001600160601b0302191690836001600160601b0316021790555060008483601f8110612ed857634e487b7160e01b600052603260045260246000fd5b6020020151600b5490870363ffffffff9081169250600160901b909104168102633b9aca0002820180156130bc57600060136000878781518110612f2c57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b03169050886001600160a01b031663a9059cbb82846040518363ffffffff1660e01b8152600401612f989291906148da565b602060405180830381600087803b158015612fb257600080fd5b505af1158015612fc6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fea91906144d6565b6130065760405162461bcd60e51b8152600401610b759061520f565b878786601f811061302757634e487b7160e01b600052603260045260246000fd5b602002019063ffffffff16908163ffffffff1681525050886001600160a01b0316816001600160a01b031687878151811061307257634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c856040516130b29190614826565b60405180910390a4505b505050600101612dc8565b506115f3600683601f613fb7565b600b54600160301b900463ffffffff166000908152600c6020526040902054601790810b900b90565b6000838381101561311157600285850304015b61311b8184613b55565b95945050505050565b600b54600160301b900463ffffffff1690565b6001600160a01b03821660009081526016602052604081205460ff1680610f0d57505060155460ff161592915050565b606060108054613176906156f5565b80601f01602080910402602001604051908101604052809291908181526020018280546131a2906156f5565b8015610e775780601f106131c457610100808354040283529160200191610e77565b820191906000526020600020905b8154815290600101906020018083116131d257509395945050505050565b600b5463ffffffff600160301b90910481166000908152600c6020526040902054600160e01b90041690565b6001600160a01b0381166000908152600260209081526040918290208251606081018452905460ff80821615158084526101008304909116938301939093526201000090046001600160601b03169281019290925261327b5750610fb7565b60006132868361078a565b905080156128a1576001600160a01b038084166000908152601360205260409081902054601154915163a9059cbb60e01b815290831692919091169063a9059cbb906132d890849086906004016148da565b602060405180830381600087803b1580156132f257600080fd5b505af1158015613306573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061332a91906144d6565b6133465760405162461bcd60e51b8152600401610b759061520f565b600b60000160069054906101000a900463ffffffff166006846020015160ff16601f811061338457634e487b7160e01b600052603260045260246000fd5b6008810491909101805460079092166004026101000a63ffffffff8181021990931693909216919091029190911790556001600160a01b038481166000818152600260205260409081902080546dffffffffffffffffffffffff000019169055601154905190831692841691907fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c9061341e908790614826565b60405180910390a450505050565b60008080808063ffffffff69ffffffffffffffffffff8716111561345e57506000935083925082915081905080611424565b5050505063ffffffff8281166000908152600c602090815260409182902082516060810184529054601781810b810b810b808452600160c01b83048716948401859052600160e01b9092049095169190930181905294959190920b939192508490565b60006134ce8260206155fd565b6134d98560206155fd565b6134e588610144615578565b6134ef9190615578565b6134f99190615578565b613504906000615578565b90503681146135255760405162461bcd60e51b8152600401610b7590614cb4565b50505050505050565b60008061353a83612c63565b9050601f81604001515111156135625760405162461bcd60e51b8152600401610b7590614e6c565b604081015151865160ff161061358a5760405162461bcd60e51b8152600401610b7590614d90565b64ffffffffff8416602087015260408101518051600091906135ae906002906155dd565b815181106135cc57634e487b7160e01b600052603260045260246000fd5b602002602001015190508060170b7f000000000000000000000000000000000000000000000000000000000000000060170b1315801561363257507f000000000000000000000000000000000000000000000000000000000000000060170b8160170b13155b61364e5760405162461bcd60e51b8152600401610b7590614ea3565b6040870180519061365e82615745565b63ffffffff1663ffffffff168152505060405180606001604052808260170b8152602001836000015163ffffffff1681526020014263ffffffff16815250600c6000896040015163ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a81548177ffffffffffffffffffffffffffffffffffffffffffffffff021916908360170b77ffffffffffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160186101000a81548163ffffffff021916908363ffffffff160217905550604082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555090505086600b60008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548164ffffffffff021916908364ffffffffff16021790555060408201518160000160066101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001600a6101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600e6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160000160126101000a81548163ffffffff021916908363ffffffff16021790555060c08201518160000160166101000a81548163ffffffff021916908363ffffffff16021790555060e082015181600001601a6101000a81548162ffffff021916908362ffffff160217905550905050866040015163ffffffff167fc797025feeeaf2cd924c99e9205acb8ec04d5cad21c41ce637a38fb6dee6016a823385600001518660400151876020015188606001518d8d6040516138f2989796959493929190614a06565b60405180910390a260006001600160a01b0316876040015163ffffffff167f0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac602718460000151604051613943919061539e565b60405180910390a3866040015163ffffffff168160170b7f0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f426040516139899190614826565b60405180910390a36139a287604001518260170b613da7565b506060015195945050505050565b60008360170b12156139c157611e8c565b60006139e8633b9aca003a04866080015163ffffffff16876060015163ffffffff166130fe565b90506010360260005a90506000613a118663ffffffff1685858b60e0015162ffffff1686612c2f565b90506000670de0b6b3a764000077ffffffffffffffffffffffffffffffffffffffffffffffff891683026001600160a01b03881660009081526002602052604090205460c08c01519290910492506201000090046001600160601b039081169163ffffffff16633b9aca000282840101908116821115613a975750505050505050611e8c565b6001600160a01b038816600090815260026020526040902080546001600160601b0390921662010000026dffffffffffffffffffffffff00001990921691909117905550505050505050505050565b600063ffffffff821115613afc5750600061089b565b5063ffffffff166000908152600c6020526040902054601790810b900b90565b600063ffffffff821115613b325750600061089b565b5063ffffffff9081166000908152600c6020526040902054600160e01b90041690565b600081831015613b66575081610f10565b50919050565b80600010610fb75760405162461bcd60e51b8152600401610b75906150fe565b6000808a8a8a8a8a8a8a8a8a604051602001613bb0999897969594939291906152eb565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b6001600160a01b038116331415613c425760405162461bcd60e51b8152600401610b7590615246565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6012546001600160a01b039081169082168114610bce57601280546001600160a01b0319166001600160a01b0384161790556040517f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d489129061166690839085906149bf565b600b54600160301b900463ffffffff9081166000818152600c602090815260409182902082516060810184529054601781810b810b810b808452600160c01b83048816948401859052600160e01b9092049096169190930181905292949190930b9291908490565b600081516020613d6f91906155fd565b613d7a9060a0615578565b613d85906000615578565b9050808351146128a15760405162461bcd60e51b8152600401610b7590614dc7565b60408051808201909152600e546001600160a01b038116808352600160a01b90910463ffffffff166020830152613dde5750610bce565b6000613deb6001856156d0565b63ffffffff8082166000818152600c6020908152604091829020549087015187519251959650601791820b90910b94613e949491821693613e3792909187918c16908b90602401615383565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fbeed9b5100000000000000000000000000000000000000000000000000000000179052613eb0565b6115f35760405162461bcd60e51b8152600401610b7590614e35565b60005a6113888110613ee45761138881039050846040820482031115613ee4576000808451602086016000888af150600191505b509392505050565b604080518082019091526000808252602082015290565b5080546000825590600052602060002090810190610fb7919061404a565b828054828255906000526020600020908101928215613f76579160200282015b82811115613f7657825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613f41565b50613f8292915061404a565b5090565b6040518060800160405280600063ffffffff1681526020016060815260200160608152602001600060170b81525090565b600483019183908215613f765791602002820160005b8382111561401157835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302613fcd565b80156140415782816101000a81549063ffffffff0219169055600401602081600301049283019260010302614011565b5050613f829291505b5b80821115613f82576000815560010161404b565b60008083601f840112614070578182fd5b50813567ffffffffffffffff811115614087578182fd5b60208301915083602080830285010111156140a157600080fd5b9250929050565b600082601f8301126140b8578081fd5b813560206140cd6140c883615554565b61552a565b82815281810190858301838502870184018810156140e9578586fd5b855b858110156141105781356140fe81615795565b845292840192908401906001016140eb565b5090979650505050505050565b600082601f83011261412d578081fd5b813567ffffffffffffffff8111156141475761414761577f565b61415a601f8201601f191660200161552a565b81815284602083860101111561416e578283fd5b816020850160208301379081016020019190915292915050565b8051601781900b811461089b57600080fd5b803567ffffffffffffffff8116811461089b57600080fd5b803560ff8116811461089b57600080fd5b6000602082840312156141d4578081fd5b813561270c81615795565b600080604083850312156141f1578081fd5b82356141fc81615795565b9150602083013561420c81615795565b809150509250929050565b60008060408385031215614229578182fd5b823561423481615795565b9150602083013567ffffffffffffffff81111561424f578182fd5b61425b8582860161411d565b9150509250929050565b60008060408385031215614277578182fd5b823561428281615795565b946020939093013593505050565b600080604083850312156142a2578182fd5b82356142ad81615795565b915060208301356001600160601b038116811461420c578182fd5b600080600080604085870312156142dd578182fd5b843567ffffffffffffffff808211156142f4578384fd5b6143008883890161405f565b90965094506020870135915080821115614318578384fd5b506143258782880161405f565b95989497509550505050565b60008060008060008060c08789031215614349578384fd5b863567ffffffffffffffff80821115614360578586fd5b61436c8a838b016140a8565b97506020890135915080821115614381578586fd5b61438d8a838b016140a8565b965061439b60408a016141b2565b955060608901359150808211156143b0578384fd5b6143bc8a838b0161411d565b94506143ca60808a0161419a565b935060a08901359150808211156143df578283fd5b506143ec89828a0161411d565b9150509295509295509295565b60008060008060008060008060e0898b031215614414578586fd5b606089018a811115614424578687fd5b8998503567ffffffffffffffff8082111561443d578788fd5b818b0191508b601f830112614450578788fd5b81358181111561445e578889fd5b8c602082850101111561446f578889fd5b6020830199508098505060808b013591508082111561448c578384fd5b6144988c838d0161405f565b909750955060a08b01359150808211156144b0578384fd5b506144bd8b828c0161405f565b999c989b50969995989497949560c00135949350505050565b6000602082840312156144e7578081fd5b8151801515811461270c578182fd5b600060208284031215614507578081fd5b813567ffffffffffffffff81111561451d578182fd5b610c4b8482850161411d565b6000806040838503121561453b578182fd5b823561454681615795565b9150602083013561420c816157aa565b600080604083850312156141f1578182fd5b600060208284031215614579578081fd5b5035919050565b600060208284031215614591578081fd5b5051919050565b6000806000606084860312156145ac578081fd5b505081359360208301359350604090920135919050565b600080600080600060a086880312156145da578283fd5b505083359560208501359550604085013594606081013594506080013592509050565b60008060008060808587031215614612578182fd5b845161461d816157aa565b809450506020808601519350604086015167ffffffffffffffff811115614642578384fd5b8601601f81018813614652578384fd5b80516146606140c882615554565b81815283810190838501858402850186018c101561467c578788fd5b8794505b838510156146a55761469181614188565b835260019490940193918501918501614680565b5080965050505050506146ba60608601614188565b905092959194509250565b600080600080600060a086880312156146dc578283fd5b85356146e7816157aa565b945060208601356146f7816157aa565b93506040860135614707816157aa565b92506060860135614717816157aa565b9150608086013562ffffff8116811461472e578182fd5b809150509295509295909350565b60006020828403121561474d578081fd5b813569ffffffffffffffffffff8116811461270c578182fd5b6000815180845260208085019450808401835b8381101561479e5781516001600160a01b031687529582019590820190600101614779565b509495945050505050565b6000815180845260208085019450808401835b8381101561479e57815160170b875295820195908201906001016147bc565b60008151808452815b81811015614800576020818501810151868301820152016147e4565b818111156148115782602083870101525b50601f01601f19169290920160200192915050565b90815260200190565b6000828483379101908152919050565b60f89390931b7fff00000000000000000000000000000000000000000000000000000000000000168352601791820b604090811b6001850152910b901b601982015260310190565b6001600160a01b0391909116815260200190565b60006001600160a01b03851682526040602083015282604083015282846060840137818301606090810191909152601f909201601f1916010192915050565b6001600160a01b03929092168252602082015260400190565b600060208252610f0d6020830184614766565b901515815260200190565b9215158352602083019190915263ffffffff16604082015260600190565b828152608081016060836020840137600081529392505050565b92835263ffffffff91909116602083015260ff16604082015260600190565b94855263ffffffff93909316602085015260ff91909116604084015260170b606083015267ffffffffffffffff16608082015260a00190565b93845260ff9290921660208401526040830152606082015260800190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b0392909216825263ffffffff16602082015260400190565b60179190910b815260200190565b60006101008a60170b83526001600160a01b038a16602084015263ffffffff89166040840152806060840152614a3e818401896147a9565b90508281036080840152614a5281886147db565b60179690960b60a0840152505060c081019290925264ffffffffff1660e09091015295945050505050565b600060208252610f0d60208301846147db565b6020808252601f908201527f6f6e6c792070726f706f736564207061796565732063616e2061636365707400604082015260600190565b60208082526016908201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604082015260600190565b6020808252601d908201527f6f6e6c792063757272656e742070617965652063616e20757064617465000000604082015260600190565b6020808252601c908201527f7265706561746564207472616e736d6974746572206164647265737300000000604082015260600190565b6020808252818101527f6c6566744761732063616e6e6f742065786365656420696e697469616c476173604082015260600190565b6020808252600c908201527f7374616c65207265706f72740000000000000000000000000000000000000000604082015260600190565b6020808252601d908201527f4f6e6c79206f776e6572267265717565737465722063616e2063616c6c000000604082015260600190565b60208082526016908201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604082015260600190565b60208082526010908201527f746f6f206d616e79206f7261636c657300000000000000000000000000000000604082015260600190565b60208082526016908201527f6f7261636c65206c656e677468206d69736d6174636800000000000000000000604082015260600190565b60208082526018908201527f63616c6c64617461206c656e677468206d69736d617463680000000000000000604082015260600190565b6020808252601f908201527f7472616e736665722072656d61696e696e672066756e6473206661696c656400604082015260600190565b60208082526014908201527f4f6e6c792063616c6c61626c6520627920454f41000000000000000000000000604082015260600190565b60208082526010908201527f6475706c6963617465207369676e657200000000000000000000000000000000604082015260600190565b6020808252601e908201527f746f6f206665772076616c75657320746f207472757374206d656469616e0000604082015260600190565b60208082526016908201527f7265706f7274206c656e677468206d69736d6174636800000000000000000000604082015260600190565b60208082526017908201527f7265706561746564207369676e65722061646472657373000000000000000000604082015260600190565b60208082526010908201527f696e73756666696369656e742067617300000000000000000000000000000000604082015260600190565b6020808252601e908201527f6e756d206f62736572766174696f6e73206f7574206f6620626f756e64730000604082015260600190565b6020808252601e908201527f6d656469616e206973206f7574206f66206d696e2d6d61782072616e67650000604082015260600190565b60208082526009908201527f4e6f206163636573730000000000000000000000000000000000000000000000604082015260600190565b60208082526013908201527f696e76616c6964207472616e736d697474657200000000000000000000000000604082015260600190565b6020808252601b908201527f6f6e636861696e436f6e666967206d75737420626520656d7074790000000000604082015260600190565b60208082526011908201527f706179656520616c726561647920736574000000000000000000000000000000604082015260600190565b6020808252818101527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c604082015260600190565b60208082526018908201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604082015260600190565b6020808252601a908201527f77726f6e67206e756d626572206f66207369676e617475726573000000000000604082015260600190565b60208082526017908201527f4f6e6c792070617965652063616e207769746864726177000000000000000000604082015260600190565b60208082526014908201527f696e73756666696369656e742062616c616e6365000000000000000000000000604082015260600190565b6020808252600f908201527f7369676e6174757265206572726f720000000000000000000000000000000000604082015260600190565b60208082526012908201527f66206d75737420626520706f7369746976650000000000000000000000000000604082015260600190565b60208082526015908201527f636f6e666967446967657374206d69736d617463680000000000000000000000604082015260600190565b60208082526017908201527f63616e6e6f74207472616e7366657220746f2073656c66000000000000000000604082015260600190565b6020808252818101527f7472616e736d6974746572732e73697a6520213d207061796565732e73697a65604082015260600190565b60208082526018908201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604082015260600190565b60208082526012908201527f696e73756666696369656e742066756e64730000000000000000000000000000604082015260600190565b60208082526017908201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604082015260600190565b6020808252600f908201527f7472616e73666572206661696c65640000000000000000000000000000000000604082015260600190565b6020808252601e908201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604082015260600190565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526153258285018b614766565b91508382036080850152615339828a614766565b915060ff881660a085015283820360c085015261535682886147db565b90861660e0850152838103610100850152905061537381856147db565b9c9b505050505050505050505050565b93845260208401929092526040830152606082015260800190565b63ffffffff91909116815260200190565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526153df8184018a614766565b905082810360808401526153f38189614766565b905060ff871660a084015282810360c084015261541081876147db565b905067ffffffffffffffff851660e084015282810361010084015261537381856147db565b600063ffffffff851682526060602083015261545460608301856147db565b828103604084015261097881856147a9565b63ffffffff92831681529116602082015260400190565b63ffffffff9384168152919092166020820152604081019190915260600190565b63ffffffff958616815293851660208501529184166040840152909216606082015262ffffff909116608082015260a00190565b69ffffffffffffffffffff91909116815260200190565b69ffffffffffffffffffff9586168152602081019490945260408401929092526060830152909116608082015260a00190565b60ff91909116815260200190565b60405181810167ffffffffffffffff8111828210171561554c5761554c61577f565b604052919050565b600067ffffffffffffffff82111561556e5761556e61577f565b5060209081020190565b6000821982111561558b5761558b615769565b500190565b600063ffffffff8083168185168083038211156155af576155af615769565b01949350505050565b600060ff821660ff84168060ff038211156155d5576155d5615769565b019392505050565b6000826155f857634e487b7160e01b81526012600452602481fd5b500490565b600081600019048311821515161561561757615617615769565b500290565b600060ff821660ff84168160ff048111821515161561563d5761563d615769565b029392505050565b6000808312837f80000000000000000000000000000000000000000000000000000000000000000183128115161561567f5761567f615769565b837f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0183138116156156b3576156b3615769565b50500390565b6000828210156156cb576156cb615769565b500390565b600063ffffffff838116908316818110156156ed576156ed615769565b039392505050565b60028104600182168061570957607f821691505b60208210811415613b6657634e487b7160e01b600052602260045260246000fd5b600060001982141561573e5761573e615769565b5060010190565b600063ffffffff8083168181141561575f5761575f615769565b6001019392505050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114610fb757600080fd5b63ffffffff81168114610fb757600080fdfe416363657373436f6e74726f6c6c65644f43523241676772656761746f7220312e302e302d616c706861a164736f6c6343000800000a",
}

var TestOCR2AggregatorABI = TestOCR2AggregatorMetaData.ABI

var TestOCR2AggregatorBin = TestOCR2AggregatorMetaData.Bin

func DeployTestOCR2Aggregator(auth *bind.TransactOpts, backend bind.ContractBackend, _link common.Address, _minAnswer *big.Int, _maxAnswer *big.Int, _billingAccessController common.Address, _requesterAdminAccessController common.Address) (common.Address, *types.Transaction, *TestOCR2Aggregator, error) {
	parsed, err := TestOCR2AggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestOCR2AggregatorBin), backend, _link, _minAnswer, _maxAnswer, _billingAccessController, _requesterAdminAccessController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestOCR2Aggregator{TestOCR2AggregatorCaller: TestOCR2AggregatorCaller{contract: contract}, TestOCR2AggregatorTransactor: TestOCR2AggregatorTransactor{contract: contract}, TestOCR2AggregatorFilterer: TestOCR2AggregatorFilterer{contract: contract}}, nil
}

type TestOCR2Aggregator struct {
	address common.Address
	abi     abi.ABI
	TestOCR2AggregatorCaller
	TestOCR2AggregatorTransactor
	TestOCR2AggregatorFilterer
}

type TestOCR2AggregatorCaller struct {
	contract *bind.BoundContract
}

type TestOCR2AggregatorTransactor struct {
	contract *bind.BoundContract
}

type TestOCR2AggregatorFilterer struct {
	contract *bind.BoundContract
}

type TestOCR2AggregatorSession struct {
	Contract     *TestOCR2Aggregator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TestOCR2AggregatorCallerSession struct {
	Contract *TestOCR2AggregatorCaller
	CallOpts bind.CallOpts
}

type TestOCR2AggregatorTransactorSession struct {
	Contract     *TestOCR2AggregatorTransactor
	TransactOpts bind.TransactOpts
}

type TestOCR2AggregatorRaw struct {
	Contract *TestOCR2Aggregator
}

type TestOCR2AggregatorCallerRaw struct {
	Contract *TestOCR2AggregatorCaller
}

type TestOCR2AggregatorTransactorRaw struct {
	Contract *TestOCR2AggregatorTransactor
}

func NewTestOCR2Aggregator(address common.Address, backend bind.ContractBackend) (*TestOCR2Aggregator, error) {
	abi, err := abi.JSON(strings.NewReader(TestOCR2AggregatorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindTestOCR2Aggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestOCR2Aggregator{address: address, abi: abi, TestOCR2AggregatorCaller: TestOCR2AggregatorCaller{contract: contract}, TestOCR2AggregatorTransactor: TestOCR2AggregatorTransactor{contract: contract}, TestOCR2AggregatorFilterer: TestOCR2AggregatorFilterer{contract: contract}}, nil
}

func NewTestOCR2AggregatorCaller(address common.Address, caller bind.ContractCaller) (*TestOCR2AggregatorCaller, error) {
	contract, err := bindTestOCR2Aggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorCaller{contract: contract}, nil
}

func NewTestOCR2AggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*TestOCR2AggregatorTransactor, error) {
	contract, err := bindTestOCR2Aggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorTransactor{contract: contract}, nil
}

func NewTestOCR2AggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*TestOCR2AggregatorFilterer, error) {
	contract, err := bindTestOCR2Aggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorFilterer{contract: contract}, nil
}

func bindTestOCR2Aggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestOCR2AggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestOCR2Aggregator.Contract.TestOCR2AggregatorCaller.contract.Call(opts, result, method, params...)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.TestOCR2AggregatorTransactor.contract.Transfer(opts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.TestOCR2AggregatorTransactor.contract.Transact(opts, method, params...)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestOCR2Aggregator.Contract.contract.Call(opts, result, method, params...)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.contract.Transfer(opts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.contract.Transact(opts, method, params...)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) CheckEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "checkEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) CheckEnabled() (bool, error) {
	return _TestOCR2Aggregator.Contract.CheckEnabled(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) CheckEnabled() (bool, error) {
	return _TestOCR2Aggregator.Contract.CheckEnabled(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) Decimals() (uint8, error) {
	return _TestOCR2Aggregator.Contract.Decimals(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) Decimals() (uint8, error) {
	return _TestOCR2Aggregator.Contract.Decimals(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) Description() (string, error) {
	return _TestOCR2Aggregator.Contract.Description(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) Description() (string, error) {
	return _TestOCR2Aggregator.Contract.Description(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "getAnswer", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.GetAnswer(&_TestOCR2Aggregator.CallOpts, _roundId)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.GetAnswer(&_TestOCR2Aggregator.CallOpts, _roundId)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) GetBilling(opts *bind.CallOpts) (GetBilling,

	error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "getBilling")

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

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) GetBilling() (GetBilling,

	error) {
	return _TestOCR2Aggregator.Contract.GetBilling(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) GetBilling() (GetBilling,

	error) {
	return _TestOCR2Aggregator.Contract.GetBilling(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) GetBillingAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "getBillingAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) GetBillingAccessController() (common.Address, error) {
	return _TestOCR2Aggregator.Contract.GetBillingAccessController(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) GetBillingAccessController() (common.Address, error) {
	return _TestOCR2Aggregator.Contract.GetBillingAccessController(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) GetConfigDigest(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "getConfigDigest")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) GetConfigDigest() ([32]byte, error) {
	return _TestOCR2Aggregator.Contract.GetConfigDigest(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) GetConfigDigest() ([32]byte, error) {
	return _TestOCR2Aggregator.Contract.GetConfigDigest(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) GetLinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "getLinkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) GetLinkToken() (common.Address, error) {
	return _TestOCR2Aggregator.Contract.GetLinkToken(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) GetLinkToken() (common.Address, error) {
	return _TestOCR2Aggregator.Contract.GetLinkToken(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) GetRequesterAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "getRequesterAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) GetRequesterAccessController() (common.Address, error) {
	return _TestOCR2Aggregator.Contract.GetRequesterAccessController(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) GetRequesterAccessController() (common.Address, error) {
	return _TestOCR2Aggregator.Contract.GetRequesterAccessController(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (GetRoundData,

	error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "getRoundData", _roundId)

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

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) GetRoundData(_roundId *big.Int) (GetRoundData,

	error) {
	return _TestOCR2Aggregator.Contract.GetRoundData(&_TestOCR2Aggregator.CallOpts, _roundId)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) GetRoundData(_roundId *big.Int) (GetRoundData,

	error) {
	return _TestOCR2Aggregator.Contract.GetRoundData(&_TestOCR2Aggregator.CallOpts, _roundId)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) GetTimestamp(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "getTimestamp", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.GetTimestamp(&_TestOCR2Aggregator.CallOpts, _roundId)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.GetTimestamp(&_TestOCR2Aggregator.CallOpts, _roundId)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) GetTransmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "getTransmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) GetTransmitters() ([]common.Address, error) {
	return _TestOCR2Aggregator.Contract.GetTransmitters(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) GetTransmitters() ([]common.Address, error) {
	return _TestOCR2Aggregator.Contract.GetTransmitters(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) GetValidatorConfig(opts *bind.CallOpts) (GetValidatorConfig,

	error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "getValidatorConfig")

	outstruct := new(GetValidatorConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.GasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) GetValidatorConfig() (GetValidatorConfig,

	error) {
	return _TestOCR2Aggregator.Contract.GetValidatorConfig(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) GetValidatorConfig() (GetValidatorConfig,

	error) {
	return _TestOCR2Aggregator.Contract.GetValidatorConfig(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) HasAccess(opts *bind.CallOpts, _user common.Address, _calldata []byte) (bool, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "hasAccess", _user, _calldata)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) HasAccess(_user common.Address, _calldata []byte) (bool, error) {
	return _TestOCR2Aggregator.Contract.HasAccess(&_TestOCR2Aggregator.CallOpts, _user, _calldata)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) HasAccess(_user common.Address, _calldata []byte) (bool, error) {
	return _TestOCR2Aggregator.Contract.HasAccess(&_TestOCR2Aggregator.CallOpts, _user, _calldata)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) LatestAnswer() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.LatestAnswer(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) LatestAnswer() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.LatestAnswer(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _TestOCR2Aggregator.Contract.LatestConfigDetails(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _TestOCR2Aggregator.Contract.LatestConfigDetails(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _TestOCR2Aggregator.Contract.LatestConfigDigestAndEpoch(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _TestOCR2Aggregator.Contract.LatestConfigDigestAndEpoch(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) LatestRound() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.LatestRound(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) LatestRound() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.LatestRound(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) LatestRoundData(opts *bind.CallOpts) (LatestRoundData,

	error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "latestRoundData")

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

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) LatestRoundData() (LatestRoundData,

	error) {
	return _TestOCR2Aggregator.Contract.LatestRoundData(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) LatestRoundData() (LatestRoundData,

	error) {
	return _TestOCR2Aggregator.Contract.LatestRoundData(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) LatestTimestamp() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.LatestTimestamp(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) LatestTimestamp() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.LatestTimestamp(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) LatestTransmissionDetails(opts *bind.CallOpts) (LatestTransmissionDetails,

	error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "latestTransmissionDetails")

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

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) LatestTransmissionDetails() (LatestTransmissionDetails,

	error) {
	return _TestOCR2Aggregator.Contract.LatestTransmissionDetails(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) LatestTransmissionDetails() (LatestTransmissionDetails,

	error) {
	return _TestOCR2Aggregator.Contract.LatestTransmissionDetails(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) LinkAvailableForPayment() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.LinkAvailableForPayment(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.LinkAvailableForPayment(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) MaxAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "maxAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) MaxAnswer() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.MaxAnswer(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) MaxAnswer() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.MaxAnswer(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) MinAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "minAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) MinAnswer() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.MinAnswer(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) MinAnswer() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.MinAnswer(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) OracleObservationCount(opts *bind.CallOpts, transmitterAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "oracleObservationCount", transmitterAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _TestOCR2Aggregator.Contract.OracleObservationCount(&_TestOCR2Aggregator.CallOpts, transmitterAddress)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _TestOCR2Aggregator.Contract.OracleObservationCount(&_TestOCR2Aggregator.CallOpts, transmitterAddress)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) OwedPayment(opts *bind.CallOpts, transmitterAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "owedPayment", transmitterAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.OwedPayment(&_TestOCR2Aggregator.CallOpts, transmitterAddress)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.OwedPayment(&_TestOCR2Aggregator.CallOpts, transmitterAddress)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) Owner() (common.Address, error) {
	return _TestOCR2Aggregator.Contract.Owner(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) Owner() (common.Address, error) {
	return _TestOCR2Aggregator.Contract.Owner(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) TestDecodeReport(opts *bind.CallOpts, report []byte) (uint32, []byte, []*big.Int, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "testDecodeReport", report)

	if err != nil {
		return *new(uint32), *new([]byte), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	out2 := *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, out2, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) TestDecodeReport(report []byte) (uint32, []byte, []*big.Int, error) {
	return _TestOCR2Aggregator.Contract.TestDecodeReport(&_TestOCR2Aggregator.CallOpts, report)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) TestDecodeReport(report []byte) (uint32, []byte, []*big.Int, error) {
	return _TestOCR2Aggregator.Contract.TestDecodeReport(&_TestOCR2Aggregator.CallOpts, report)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) TestPayee(opts *bind.CallOpts, _transmitter common.Address) (common.Address, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "testPayee", _transmitter)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) TestPayee(_transmitter common.Address) (common.Address, error) {
	return _TestOCR2Aggregator.Contract.TestPayee(&_TestOCR2Aggregator.CallOpts, _transmitter)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) TestPayee(_transmitter common.Address) (common.Address, error) {
	return _TestOCR2Aggregator.Contract.TestPayee(&_TestOCR2Aggregator.CallOpts, _transmitter)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) TestReimbursementGasPriceGwei(opts *bind.CallOpts, txGasPrice *big.Int, reasonableGasPrice *big.Int, maximumGasPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "testReimbursementGasPriceGwei", txGasPrice, reasonableGasPrice, maximumGasPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) TestReimbursementGasPriceGwei(txGasPrice *big.Int, reasonableGasPrice *big.Int, maximumGasPrice *big.Int) (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.TestReimbursementGasPriceGwei(&_TestOCR2Aggregator.CallOpts, txGasPrice, reasonableGasPrice, maximumGasPrice)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) TestReimbursementGasPriceGwei(txGasPrice *big.Int, reasonableGasPrice *big.Int, maximumGasPrice *big.Int) (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.TestReimbursementGasPriceGwei(&_TestOCR2Aggregator.CallOpts, txGasPrice, reasonableGasPrice, maximumGasPrice)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) TestTotalLinkDue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "testTotalLinkDue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) TestTotalLinkDue() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.TestTotalLinkDue(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) TestTotalLinkDue() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.TestTotalLinkDue(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) TestTransmitterGasCostWei(opts *bind.CallOpts, initialGas *big.Int, gasPriceGwei *big.Int, callDataGas *big.Int, accountingGas *big.Int, leftGas *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "testTransmitterGasCostWei", initialGas, gasPriceGwei, callDataGas, accountingGas, leftGas)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) TestTransmitterGasCostWei(initialGas *big.Int, gasPriceGwei *big.Int, callDataGas *big.Int, accountingGas *big.Int, leftGas *big.Int) (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.TestTransmitterGasCostWei(&_TestOCR2Aggregator.CallOpts, initialGas, gasPriceGwei, callDataGas, accountingGas, leftGas)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) TestTransmitterGasCostWei(initialGas *big.Int, gasPriceGwei *big.Int, callDataGas *big.Int, accountingGas *big.Int, leftGas *big.Int) (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.TestTransmitterGasCostWei(&_TestOCR2Aggregator.CallOpts, initialGas, gasPriceGwei, callDataGas, accountingGas, leftGas)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) TypeAndVersion() (string, error) {
	return _TestOCR2Aggregator.Contract.TypeAndVersion(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) TypeAndVersion() (string, error) {
	return _TestOCR2Aggregator.Contract.TypeAndVersion(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestOCR2Aggregator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) Version() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.Version(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorCallerSession) Version() (*big.Int, error) {
	return _TestOCR2Aggregator.Contract.Version(&_TestOCR2Aggregator.CallOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "acceptOwnership")
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.AcceptOwnership(&_TestOCR2Aggregator.TransactOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.AcceptOwnership(&_TestOCR2Aggregator.TransactOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) AcceptPayeeship(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "acceptPayeeship", transmitter)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.AcceptPayeeship(&_TestOCR2Aggregator.TransactOpts, transmitter)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.AcceptPayeeship(&_TestOCR2Aggregator.TransactOpts, transmitter)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "addAccess", _user)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.AddAccess(&_TestOCR2Aggregator.TransactOpts, _user)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.AddAccess(&_TestOCR2Aggregator.TransactOpts, _user)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "disableAccessCheck")
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.DisableAccessCheck(&_TestOCR2Aggregator.TransactOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.DisableAccessCheck(&_TestOCR2Aggregator.TransactOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "enableAccessCheck")
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.EnableAccessCheck(&_TestOCR2Aggregator.TransactOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.EnableAccessCheck(&_TestOCR2Aggregator.TransactOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "removeAccess", _user)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.RemoveAccess(&_TestOCR2Aggregator.TransactOpts, _user)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.RemoveAccess(&_TestOCR2Aggregator.TransactOpts, _user)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) RequestNewRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "requestNewRound")
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) RequestNewRound() (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.RequestNewRound(&_TestOCR2Aggregator.TransactOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) RequestNewRound() (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.RequestNewRound(&_TestOCR2Aggregator.TransactOpts)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) SetBilling(opts *bind.TransactOpts, maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "setBilling", maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.SetBilling(&_TestOCR2Aggregator.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.SetBilling(&_TestOCR2Aggregator.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) SetBillingAccessController(opts *bind.TransactOpts, _billingAccessController common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "setBillingAccessController", _billingAccessController)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.SetBillingAccessController(&_TestOCR2Aggregator.TransactOpts, _billingAccessController)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.SetBillingAccessController(&_TestOCR2Aggregator.TransactOpts, _billingAccessController)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "setConfig", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.SetConfig(&_TestOCR2Aggregator.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.SetConfig(&_TestOCR2Aggregator.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) SetLinkToken(opts *bind.TransactOpts, linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "setLinkToken", linkToken, recipient)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) SetLinkToken(linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.SetLinkToken(&_TestOCR2Aggregator.TransactOpts, linkToken, recipient)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) SetLinkToken(linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.SetLinkToken(&_TestOCR2Aggregator.TransactOpts, linkToken, recipient)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) SetPayees(opts *bind.TransactOpts, transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "setPayees", transmitters, payees)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.SetPayees(&_TestOCR2Aggregator.TransactOpts, transmitters, payees)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.SetPayees(&_TestOCR2Aggregator.TransactOpts, transmitters, payees)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) SetRequesterAccessController(opts *bind.TransactOpts, requesterAccessController common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "setRequesterAccessController", requesterAccessController)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) SetRequesterAccessController(requesterAccessController common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.SetRequesterAccessController(&_TestOCR2Aggregator.TransactOpts, requesterAccessController)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) SetRequesterAccessController(requesterAccessController common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.SetRequesterAccessController(&_TestOCR2Aggregator.TransactOpts, requesterAccessController)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) SetValidatorConfig(opts *bind.TransactOpts, newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "setValidatorConfig", newValidator, newGasLimit)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) SetValidatorConfig(newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.SetValidatorConfig(&_TestOCR2Aggregator.TransactOpts, newValidator, newGasLimit)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) SetValidatorConfig(newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.SetValidatorConfig(&_TestOCR2Aggregator.TransactOpts, newValidator, newGasLimit)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) TestBurnLink(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "testBurnLink", amount)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) TestBurnLink(amount *big.Int) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.TestBurnLink(&_TestOCR2Aggregator.TransactOpts, amount)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) TestBurnLink(amount *big.Int) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.TestBurnLink(&_TestOCR2Aggregator.TransactOpts, amount)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) TestSetTransmitterPayment(opts *bind.TransactOpts, transmitter common.Address, paymentJuels *big.Int) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "testSetTransmitterPayment", transmitter, paymentJuels)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) TestSetTransmitterPayment(transmitter common.Address, paymentJuels *big.Int) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.TestSetTransmitterPayment(&_TestOCR2Aggregator.TransactOpts, transmitter, paymentJuels)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) TestSetTransmitterPayment(transmitter common.Address, paymentJuels *big.Int) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.TestSetTransmitterPayment(&_TestOCR2Aggregator.TransactOpts, transmitter, paymentJuels)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "transferOwnership", to)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.TransferOwnership(&_TestOCR2Aggregator.TransactOpts, to)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.TransferOwnership(&_TestOCR2Aggregator.TransactOpts, to)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) TransferPayeeship(opts *bind.TransactOpts, transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "transferPayeeship", transmitter, proposed)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.TransferPayeeship(&_TestOCR2Aggregator.TransactOpts, transmitter, proposed)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.TransferPayeeship(&_TestOCR2Aggregator.TransactOpts, transmitter, proposed)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.Transmit(&_TestOCR2Aggregator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.Transmit(&_TestOCR2Aggregator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) WithdrawFunds(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "withdrawFunds", recipient, amount)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.WithdrawFunds(&_TestOCR2Aggregator.TransactOpts, recipient, amount)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.WithdrawFunds(&_TestOCR2Aggregator.TransactOpts, recipient, amount)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactor) WithdrawPayment(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.contract.Transact(opts, "withdrawPayment", transmitter)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.WithdrawPayment(&_TestOCR2Aggregator.TransactOpts, transmitter)
}

func (_TestOCR2Aggregator *TestOCR2AggregatorTransactorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _TestOCR2Aggregator.Contract.WithdrawPayment(&_TestOCR2Aggregator.TransactOpts, transmitter)
}

type TestOCR2AggregatorAddedAccessIterator struct {
	Event *TestOCR2AggregatorAddedAccess

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorAddedAccessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorAddedAccess)
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
		it.Event = new(TestOCR2AggregatorAddedAccess)
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

func (it *TestOCR2AggregatorAddedAccessIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorAddedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorAddedAccess struct {
	User common.Address
	Raw  types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterAddedAccess(opts *bind.FilterOpts) (*TestOCR2AggregatorAddedAccessIterator, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorAddedAccessIterator{contract: _TestOCR2Aggregator.contract, event: "AddedAccess", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorAddedAccess) (event.Subscription, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorAddedAccess)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "AddedAccess", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseAddedAccess(log types.Log) (*TestOCR2AggregatorAddedAccess, error) {
	event := new(TestOCR2AggregatorAddedAccess)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "AddedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorAnswerUpdatedIterator struct {
	Event *TestOCR2AggregatorAnswerUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorAnswerUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorAnswerUpdated)
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
		it.Event = new(TestOCR2AggregatorAnswerUpdated)
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

func (it *TestOCR2AggregatorAnswerUpdatedIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*TestOCR2AggregatorAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorAnswerUpdatedIterator{contract: _TestOCR2Aggregator.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorAnswerUpdated)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseAnswerUpdated(log types.Log) (*TestOCR2AggregatorAnswerUpdated, error) {
	event := new(TestOCR2AggregatorAnswerUpdated)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorBillingAccessControllerSetIterator struct {
	Event *TestOCR2AggregatorBillingAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorBillingAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorBillingAccessControllerSet)
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
		it.Event = new(TestOCR2AggregatorBillingAccessControllerSet)
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

func (it *TestOCR2AggregatorBillingAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorBillingAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorBillingAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*TestOCR2AggregatorBillingAccessControllerSetIterator, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorBillingAccessControllerSetIterator{contract: _TestOCR2Aggregator.contract, event: "BillingAccessControllerSet", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorBillingAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorBillingAccessControllerSet)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseBillingAccessControllerSet(log types.Log) (*TestOCR2AggregatorBillingAccessControllerSet, error) {
	event := new(TestOCR2AggregatorBillingAccessControllerSet)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorBillingSetIterator struct {
	Event *TestOCR2AggregatorBillingSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorBillingSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorBillingSet)
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
		it.Event = new(TestOCR2AggregatorBillingSet)
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

func (it *TestOCR2AggregatorBillingSetIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorBillingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorBillingSet struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
	Raw                       types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterBillingSet(opts *bind.FilterOpts) (*TestOCR2AggregatorBillingSetIterator, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorBillingSetIterator{contract: _TestOCR2Aggregator.contract, event: "BillingSet", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchBillingSet(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorBillingSet) (event.Subscription, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorBillingSet)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "BillingSet", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseBillingSet(log types.Log) (*TestOCR2AggregatorBillingSet, error) {
	event := new(TestOCR2AggregatorBillingSet)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "BillingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorCheckAccessDisabledIterator struct {
	Event *TestOCR2AggregatorCheckAccessDisabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorCheckAccessDisabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorCheckAccessDisabled)
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
		it.Event = new(TestOCR2AggregatorCheckAccessDisabled)
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

func (it *TestOCR2AggregatorCheckAccessDisabledIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorCheckAccessDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorCheckAccessDisabled struct {
	Raw types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterCheckAccessDisabled(opts *bind.FilterOpts) (*TestOCR2AggregatorCheckAccessDisabledIterator, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorCheckAccessDisabledIterator{contract: _TestOCR2Aggregator.contract, event: "CheckAccessDisabled", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorCheckAccessDisabled) (event.Subscription, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorCheckAccessDisabled)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseCheckAccessDisabled(log types.Log) (*TestOCR2AggregatorCheckAccessDisabled, error) {
	event := new(TestOCR2AggregatorCheckAccessDisabled)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorCheckAccessEnabledIterator struct {
	Event *TestOCR2AggregatorCheckAccessEnabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorCheckAccessEnabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorCheckAccessEnabled)
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
		it.Event = new(TestOCR2AggregatorCheckAccessEnabled)
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

func (it *TestOCR2AggregatorCheckAccessEnabledIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorCheckAccessEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorCheckAccessEnabled struct {
	Raw types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterCheckAccessEnabled(opts *bind.FilterOpts) (*TestOCR2AggregatorCheckAccessEnabledIterator, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorCheckAccessEnabledIterator{contract: _TestOCR2Aggregator.contract, event: "CheckAccessEnabled", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorCheckAccessEnabled) (event.Subscription, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorCheckAccessEnabled)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseCheckAccessEnabled(log types.Log) (*TestOCR2AggregatorCheckAccessEnabled, error) {
	event := new(TestOCR2AggregatorCheckAccessEnabled)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorConfigSetIterator struct {
	Event *TestOCR2AggregatorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorConfigSet)
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
		it.Event = new(TestOCR2AggregatorConfigSet)
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

func (it *TestOCR2AggregatorConfigSetIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorConfigSet struct {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterConfigSet(opts *bind.FilterOpts) (*TestOCR2AggregatorConfigSetIterator, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorConfigSetIterator{contract: _TestOCR2Aggregator.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorConfigSet) (event.Subscription, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorConfigSet)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseConfigSet(log types.Log) (*TestOCR2AggregatorConfigSet, error) {
	event := new(TestOCR2AggregatorConfigSet)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorLinkTokenSetIterator struct {
	Event *TestOCR2AggregatorLinkTokenSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorLinkTokenSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorLinkTokenSet)
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
		it.Event = new(TestOCR2AggregatorLinkTokenSet)
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

func (it *TestOCR2AggregatorLinkTokenSetIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorLinkTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorLinkTokenSet struct {
	OldLinkToken common.Address
	NewLinkToken common.Address
	Raw          types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterLinkTokenSet(opts *bind.FilterOpts, oldLinkToken []common.Address, newLinkToken []common.Address) (*TestOCR2AggregatorLinkTokenSetIterator, error) {

	var oldLinkTokenRule []interface{}
	for _, oldLinkTokenItem := range oldLinkToken {
		oldLinkTokenRule = append(oldLinkTokenRule, oldLinkTokenItem)
	}
	var newLinkTokenRule []interface{}
	for _, newLinkTokenItem := range newLinkToken {
		newLinkTokenRule = append(newLinkTokenRule, newLinkTokenItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "LinkTokenSet", oldLinkTokenRule, newLinkTokenRule)
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorLinkTokenSetIterator{contract: _TestOCR2Aggregator.contract, event: "LinkTokenSet", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchLinkTokenSet(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorLinkTokenSet, oldLinkToken []common.Address, newLinkToken []common.Address) (event.Subscription, error) {

	var oldLinkTokenRule []interface{}
	for _, oldLinkTokenItem := range oldLinkToken {
		oldLinkTokenRule = append(oldLinkTokenRule, oldLinkTokenItem)
	}
	var newLinkTokenRule []interface{}
	for _, newLinkTokenItem := range newLinkToken {
		newLinkTokenRule = append(newLinkTokenRule, newLinkTokenItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "LinkTokenSet", oldLinkTokenRule, newLinkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorLinkTokenSet)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "LinkTokenSet", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseLinkTokenSet(log types.Log) (*TestOCR2AggregatorLinkTokenSet, error) {
	event := new(TestOCR2AggregatorLinkTokenSet)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "LinkTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorNewRoundIterator struct {
	Event *TestOCR2AggregatorNewRound

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorNewRoundIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorNewRound)
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
		it.Event = new(TestOCR2AggregatorNewRound)
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

func (it *TestOCR2AggregatorNewRoundIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*TestOCR2AggregatorNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorNewRoundIterator{contract: _TestOCR2Aggregator.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorNewRound)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseNewRound(log types.Log) (*TestOCR2AggregatorNewRound, error) {
	event := new(TestOCR2AggregatorNewRound)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorNewTransmissionIterator struct {
	Event *TestOCR2AggregatorNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorNewTransmission)
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
		it.Event = new(TestOCR2AggregatorNewTransmission)
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

func (it *TestOCR2AggregatorNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorNewTransmission struct {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*TestOCR2AggregatorNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorNewTransmissionIterator{contract: _TestOCR2Aggregator.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorNewTransmission)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseNewTransmission(log types.Log) (*TestOCR2AggregatorNewTransmission, error) {
	event := new(TestOCR2AggregatorNewTransmission)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorOraclePaidIterator struct {
	Event *TestOCR2AggregatorOraclePaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorOraclePaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorOraclePaid)
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
		it.Event = new(TestOCR2AggregatorOraclePaid)
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

func (it *TestOCR2AggregatorOraclePaidIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorOraclePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorOraclePaid struct {
	Transmitter common.Address
	Payee       common.Address
	Amount      *big.Int
	LinkToken   common.Address
	Raw         types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*TestOCR2AggregatorOraclePaidIterator, error) {

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

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorOraclePaidIterator{contract: _TestOCR2Aggregator.contract, event: "OraclePaid", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorOraclePaid)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseOraclePaid(log types.Log) (*TestOCR2AggregatorOraclePaid, error) {
	event := new(TestOCR2AggregatorOraclePaid)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorOwnershipTransferRequestedIterator struct {
	Event *TestOCR2AggregatorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorOwnershipTransferRequested)
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
		it.Event = new(TestOCR2AggregatorOwnershipTransferRequested)
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

func (it *TestOCR2AggregatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestOCR2AggregatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorOwnershipTransferRequestedIterator{contract: _TestOCR2Aggregator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorOwnershipTransferRequested)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*TestOCR2AggregatorOwnershipTransferRequested, error) {
	event := new(TestOCR2AggregatorOwnershipTransferRequested)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorOwnershipTransferredIterator struct {
	Event *TestOCR2AggregatorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorOwnershipTransferred)
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
		it.Event = new(TestOCR2AggregatorOwnershipTransferred)
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

func (it *TestOCR2AggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestOCR2AggregatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorOwnershipTransferredIterator{contract: _TestOCR2Aggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorOwnershipTransferred)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*TestOCR2AggregatorOwnershipTransferred, error) {
	event := new(TestOCR2AggregatorOwnershipTransferred)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorPayeeshipTransferRequestedIterator struct {
	Event *TestOCR2AggregatorPayeeshipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorPayeeshipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorPayeeshipTransferRequested)
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
		it.Event = new(TestOCR2AggregatorPayeeshipTransferRequested)
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

func (it *TestOCR2AggregatorPayeeshipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorPayeeshipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorPayeeshipTransferRequested struct {
	Transmitter common.Address
	Current     common.Address
	Proposed    common.Address
	Raw         types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*TestOCR2AggregatorPayeeshipTransferRequestedIterator, error) {

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

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorPayeeshipTransferRequestedIterator{contract: _TestOCR2Aggregator.contract, event: "PayeeshipTransferRequested", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorPayeeshipTransferRequested)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParsePayeeshipTransferRequested(log types.Log) (*TestOCR2AggregatorPayeeshipTransferRequested, error) {
	event := new(TestOCR2AggregatorPayeeshipTransferRequested)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorPayeeshipTransferredIterator struct {
	Event *TestOCR2AggregatorPayeeshipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorPayeeshipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorPayeeshipTransferred)
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
		it.Event = new(TestOCR2AggregatorPayeeshipTransferred)
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

func (it *TestOCR2AggregatorPayeeshipTransferredIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorPayeeshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorPayeeshipTransferred struct {
	Transmitter common.Address
	Previous    common.Address
	Current     common.Address
	Raw         types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*TestOCR2AggregatorPayeeshipTransferredIterator, error) {

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

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorPayeeshipTransferredIterator{contract: _TestOCR2Aggregator.contract, event: "PayeeshipTransferred", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorPayeeshipTransferred)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParsePayeeshipTransferred(log types.Log) (*TestOCR2AggregatorPayeeshipTransferred, error) {
	event := new(TestOCR2AggregatorPayeeshipTransferred)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorRemovedAccessIterator struct {
	Event *TestOCR2AggregatorRemovedAccess

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorRemovedAccessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorRemovedAccess)
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
		it.Event = new(TestOCR2AggregatorRemovedAccess)
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

func (it *TestOCR2AggregatorRemovedAccessIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorRemovedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorRemovedAccess struct {
	User common.Address
	Raw  types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterRemovedAccess(opts *bind.FilterOpts) (*TestOCR2AggregatorRemovedAccessIterator, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorRemovedAccessIterator{contract: _TestOCR2Aggregator.contract, event: "RemovedAccess", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorRemovedAccess) (event.Subscription, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorRemovedAccess)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseRemovedAccess(log types.Log) (*TestOCR2AggregatorRemovedAccess, error) {
	event := new(TestOCR2AggregatorRemovedAccess)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorRequesterAccessControllerSetIterator struct {
	Event *TestOCR2AggregatorRequesterAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorRequesterAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorRequesterAccessControllerSet)
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
		it.Event = new(TestOCR2AggregatorRequesterAccessControllerSet)
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

func (it *TestOCR2AggregatorRequesterAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorRequesterAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorRequesterAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterRequesterAccessControllerSet(opts *bind.FilterOpts) (*TestOCR2AggregatorRequesterAccessControllerSetIterator, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "RequesterAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorRequesterAccessControllerSetIterator{contract: _TestOCR2Aggregator.contract, event: "RequesterAccessControllerSet", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchRequesterAccessControllerSet(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorRequesterAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "RequesterAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorRequesterAccessControllerSet)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "RequesterAccessControllerSet", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseRequesterAccessControllerSet(log types.Log) (*TestOCR2AggregatorRequesterAccessControllerSet, error) {
	event := new(TestOCR2AggregatorRequesterAccessControllerSet)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "RequesterAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorRoundRequestedIterator struct {
	Event *TestOCR2AggregatorRoundRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorRoundRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorRoundRequested)
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
		it.Event = new(TestOCR2AggregatorRoundRequested)
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

func (it *TestOCR2AggregatorRoundRequestedIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorRoundRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorRoundRequested struct {
	Requester    common.Address
	ConfigDigest [32]byte
	Epoch        uint32
	Round        uint8
	Raw          types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterRoundRequested(opts *bind.FilterOpts, requester []common.Address) (*TestOCR2AggregatorRoundRequestedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "RoundRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorRoundRequestedIterator{contract: _TestOCR2Aggregator.contract, event: "RoundRequested", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchRoundRequested(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorRoundRequested, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "RoundRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorRoundRequested)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "RoundRequested", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseRoundRequested(log types.Log) (*TestOCR2AggregatorRoundRequested, error) {
	event := new(TestOCR2AggregatorRoundRequested)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "RoundRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorTransmittedIterator struct {
	Event *TestOCR2AggregatorTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorTransmitted)
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
		it.Event = new(TestOCR2AggregatorTransmitted)
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

func (it *TestOCR2AggregatorTransmittedIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterTransmitted(opts *bind.FilterOpts) (*TestOCR2AggregatorTransmittedIterator, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorTransmittedIterator{contract: _TestOCR2Aggregator.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorTransmitted) (event.Subscription, error) {

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorTransmitted)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseTransmitted(log types.Log) (*TestOCR2AggregatorTransmitted, error) {
	event := new(TestOCR2AggregatorTransmitted)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TestOCR2AggregatorValidatorConfigSetIterator struct {
	Event *TestOCR2AggregatorValidatorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TestOCR2AggregatorValidatorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestOCR2AggregatorValidatorConfigSet)
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
		it.Event = new(TestOCR2AggregatorValidatorConfigSet)
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

func (it *TestOCR2AggregatorValidatorConfigSetIterator) Error() error {
	return it.fail
}

func (it *TestOCR2AggregatorValidatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TestOCR2AggregatorValidatorConfigSet struct {
	PreviousValidator common.Address
	PreviousGasLimit  uint32
	CurrentValidator  common.Address
	CurrentGasLimit   uint32
	Raw               types.Log
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) FilterValidatorConfigSet(opts *bind.FilterOpts, previousValidator []common.Address, currentValidator []common.Address) (*TestOCR2AggregatorValidatorConfigSetIterator, error) {

	var previousValidatorRule []interface{}
	for _, previousValidatorItem := range previousValidator {
		previousValidatorRule = append(previousValidatorRule, previousValidatorItem)
	}

	var currentValidatorRule []interface{}
	for _, currentValidatorItem := range currentValidator {
		currentValidatorRule = append(currentValidatorRule, currentValidatorItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.FilterLogs(opts, "ValidatorConfigSet", previousValidatorRule, currentValidatorRule)
	if err != nil {
		return nil, err
	}
	return &TestOCR2AggregatorValidatorConfigSetIterator{contract: _TestOCR2Aggregator.contract, event: "ValidatorConfigSet", logs: logs, sub: sub}, nil
}

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) WatchValidatorConfigSet(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorValidatorConfigSet, previousValidator []common.Address, currentValidator []common.Address) (event.Subscription, error) {

	var previousValidatorRule []interface{}
	for _, previousValidatorItem := range previousValidator {
		previousValidatorRule = append(previousValidatorRule, previousValidatorItem)
	}

	var currentValidatorRule []interface{}
	for _, currentValidatorItem := range currentValidator {
		currentValidatorRule = append(currentValidatorRule, currentValidatorItem)
	}

	logs, sub, err := _TestOCR2Aggregator.contract.WatchLogs(opts, "ValidatorConfigSet", previousValidatorRule, currentValidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TestOCR2AggregatorValidatorConfigSet)
				if err := _TestOCR2Aggregator.contract.UnpackLog(event, "ValidatorConfigSet", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2AggregatorFilterer) ParseValidatorConfigSet(log types.Log) (*TestOCR2AggregatorValidatorConfigSet, error) {
	event := new(TestOCR2AggregatorValidatorConfigSet)
	if err := _TestOCR2Aggregator.contract.UnpackLog(event, "ValidatorConfigSet", log); err != nil {
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

func (_TestOCR2Aggregator *TestOCR2Aggregator) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _TestOCR2Aggregator.abi.Events["AddedAccess"].ID:
		return _TestOCR2Aggregator.ParseAddedAccess(log)
	case _TestOCR2Aggregator.abi.Events["AnswerUpdated"].ID:
		return _TestOCR2Aggregator.ParseAnswerUpdated(log)
	case _TestOCR2Aggregator.abi.Events["BillingAccessControllerSet"].ID:
		return _TestOCR2Aggregator.ParseBillingAccessControllerSet(log)
	case _TestOCR2Aggregator.abi.Events["BillingSet"].ID:
		return _TestOCR2Aggregator.ParseBillingSet(log)
	case _TestOCR2Aggregator.abi.Events["CheckAccessDisabled"].ID:
		return _TestOCR2Aggregator.ParseCheckAccessDisabled(log)
	case _TestOCR2Aggregator.abi.Events["CheckAccessEnabled"].ID:
		return _TestOCR2Aggregator.ParseCheckAccessEnabled(log)
	case _TestOCR2Aggregator.abi.Events["ConfigSet"].ID:
		return _TestOCR2Aggregator.ParseConfigSet(log)
	case _TestOCR2Aggregator.abi.Events["LinkTokenSet"].ID:
		return _TestOCR2Aggregator.ParseLinkTokenSet(log)
	case _TestOCR2Aggregator.abi.Events["NewRound"].ID:
		return _TestOCR2Aggregator.ParseNewRound(log)
	case _TestOCR2Aggregator.abi.Events["NewTransmission"].ID:
		return _TestOCR2Aggregator.ParseNewTransmission(log)
	case _TestOCR2Aggregator.abi.Events["OraclePaid"].ID:
		return _TestOCR2Aggregator.ParseOraclePaid(log)
	case _TestOCR2Aggregator.abi.Events["OwnershipTransferRequested"].ID:
		return _TestOCR2Aggregator.ParseOwnershipTransferRequested(log)
	case _TestOCR2Aggregator.abi.Events["OwnershipTransferred"].ID:
		return _TestOCR2Aggregator.ParseOwnershipTransferred(log)
	case _TestOCR2Aggregator.abi.Events["PayeeshipTransferRequested"].ID:
		return _TestOCR2Aggregator.ParsePayeeshipTransferRequested(log)
	case _TestOCR2Aggregator.abi.Events["PayeeshipTransferred"].ID:
		return _TestOCR2Aggregator.ParsePayeeshipTransferred(log)
	case _TestOCR2Aggregator.abi.Events["RemovedAccess"].ID:
		return _TestOCR2Aggregator.ParseRemovedAccess(log)
	case _TestOCR2Aggregator.abi.Events["RequesterAccessControllerSet"].ID:
		return _TestOCR2Aggregator.ParseRequesterAccessControllerSet(log)
	case _TestOCR2Aggregator.abi.Events["RoundRequested"].ID:
		return _TestOCR2Aggregator.ParseRoundRequested(log)
	case _TestOCR2Aggregator.abi.Events["Transmitted"].ID:
		return _TestOCR2Aggregator.ParseTransmitted(log)
	case _TestOCR2Aggregator.abi.Events["ValidatorConfigSet"].ID:
		return _TestOCR2Aggregator.ParseValidatorConfigSet(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (TestOCR2AggregatorAddedAccess) Topic() common.Hash {
	return common.HexToHash("0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4")
}

func (TestOCR2AggregatorAnswerUpdated) Topic() common.Hash {
	return common.HexToHash("0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f")
}

func (TestOCR2AggregatorBillingAccessControllerSet) Topic() common.Hash {
	return common.HexToHash("0x793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912")
}

func (TestOCR2AggregatorBillingSet) Topic() common.Hash {
	return common.HexToHash("0x0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f")
}

func (TestOCR2AggregatorCheckAccessDisabled) Topic() common.Hash {
	return common.HexToHash("0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638")
}

func (TestOCR2AggregatorCheckAccessEnabled) Topic() common.Hash {
	return common.HexToHash("0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480")
}

func (TestOCR2AggregatorConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (TestOCR2AggregatorLinkTokenSet) Topic() common.Hash {
	return common.HexToHash("0x4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a")
}

func (TestOCR2AggregatorNewRound) Topic() common.Hash {
	return common.HexToHash("0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271")
}

func (TestOCR2AggregatorNewTransmission) Topic() common.Hash {
	return common.HexToHash("0xc797025feeeaf2cd924c99e9205acb8ec04d5cad21c41ce637a38fb6dee6016a")
}

func (TestOCR2AggregatorOraclePaid) Topic() common.Hash {
	return common.HexToHash("0xd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c")
}

func (TestOCR2AggregatorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (TestOCR2AggregatorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (TestOCR2AggregatorPayeeshipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367")
}

func (TestOCR2AggregatorPayeeshipTransferred) Topic() common.Hash {
	return common.HexToHash("0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3")
}

func (TestOCR2AggregatorRemovedAccess) Topic() common.Hash {
	return common.HexToHash("0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1")
}

func (TestOCR2AggregatorRequesterAccessControllerSet) Topic() common.Hash {
	return common.HexToHash("0x27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634")
}

func (TestOCR2AggregatorRoundRequested) Topic() common.Hash {
	return common.HexToHash("0x41e3990591fd372502daa15842da15bc7f41c75309ab3ff4f56f1848c178825c")
}

func (TestOCR2AggregatorTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (TestOCR2AggregatorValidatorConfigSet) Topic() common.Hash {
	return common.HexToHash("0xb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541")
}

func (_TestOCR2Aggregator *TestOCR2Aggregator) Address() common.Address {
	return _TestOCR2Aggregator.address
}

type TestOCR2AggregatorInterface interface {
	CheckEnabled(opts *bind.CallOpts) (bool, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	Description(opts *bind.CallOpts) (string, error)

	GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error)

	GetBilling(opts *bind.CallOpts) (GetBilling,

		error)

	GetBillingAccessController(opts *bind.CallOpts) (common.Address, error)

	GetConfigDigest(opts *bind.CallOpts) ([32]byte, error)

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

	TestDecodeReport(opts *bind.CallOpts, report []byte) (uint32, []byte, []*big.Int, error)

	TestPayee(opts *bind.CallOpts, _transmitter common.Address) (common.Address, error)

	TestReimbursementGasPriceGwei(opts *bind.CallOpts, txGasPrice *big.Int, reasonableGasPrice *big.Int, maximumGasPrice *big.Int) (*big.Int, error)

	TestTotalLinkDue(opts *bind.CallOpts) (*big.Int, error)

	TestTransmitterGasCostWei(opts *bind.CallOpts, initialGas *big.Int, gasPriceGwei *big.Int, callDataGas *big.Int, accountingGas *big.Int, leftGas *big.Int) (*big.Int, error)

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

	TestBurnLink(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	TestSetTransmitterPayment(opts *bind.TransactOpts, transmitter common.Address, paymentJuels *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	TransferPayeeship(opts *bind.TransactOpts, transmitter common.Address, proposed common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	WithdrawFunds(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	WithdrawPayment(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error)

	FilterAddedAccess(opts *bind.FilterOpts) (*TestOCR2AggregatorAddedAccessIterator, error)

	WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorAddedAccess) (event.Subscription, error)

	ParseAddedAccess(log types.Log) (*TestOCR2AggregatorAddedAccess, error)

	FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*TestOCR2AggregatorAnswerUpdatedIterator, error)

	WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error)

	ParseAnswerUpdated(log types.Log) (*TestOCR2AggregatorAnswerUpdated, error)

	FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*TestOCR2AggregatorBillingAccessControllerSetIterator, error)

	WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorBillingAccessControllerSet) (event.Subscription, error)

	ParseBillingAccessControllerSet(log types.Log) (*TestOCR2AggregatorBillingAccessControllerSet, error)

	FilterBillingSet(opts *bind.FilterOpts) (*TestOCR2AggregatorBillingSetIterator, error)

	WatchBillingSet(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorBillingSet) (event.Subscription, error)

	ParseBillingSet(log types.Log) (*TestOCR2AggregatorBillingSet, error)

	FilterCheckAccessDisabled(opts *bind.FilterOpts) (*TestOCR2AggregatorCheckAccessDisabledIterator, error)

	WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorCheckAccessDisabled) (event.Subscription, error)

	ParseCheckAccessDisabled(log types.Log) (*TestOCR2AggregatorCheckAccessDisabled, error)

	FilterCheckAccessEnabled(opts *bind.FilterOpts) (*TestOCR2AggregatorCheckAccessEnabledIterator, error)

	WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorCheckAccessEnabled) (event.Subscription, error)

	ParseCheckAccessEnabled(log types.Log) (*TestOCR2AggregatorCheckAccessEnabled, error)

	FilterConfigSet(opts *bind.FilterOpts) (*TestOCR2AggregatorConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*TestOCR2AggregatorConfigSet, error)

	FilterLinkTokenSet(opts *bind.FilterOpts, oldLinkToken []common.Address, newLinkToken []common.Address) (*TestOCR2AggregatorLinkTokenSetIterator, error)

	WatchLinkTokenSet(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorLinkTokenSet, oldLinkToken []common.Address, newLinkToken []common.Address) (event.Subscription, error)

	ParseLinkTokenSet(log types.Log) (*TestOCR2AggregatorLinkTokenSet, error)

	FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*TestOCR2AggregatorNewRoundIterator, error)

	WatchNewRound(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error)

	ParseNewRound(log types.Log) (*TestOCR2AggregatorNewRound, error)

	FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*TestOCR2AggregatorNewTransmissionIterator, error)

	WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error)

	ParseNewTransmission(log types.Log) (*TestOCR2AggregatorNewTransmission, error)

	FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*TestOCR2AggregatorOraclePaidIterator, error)

	WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error)

	ParseOraclePaid(log types.Log) (*TestOCR2AggregatorOraclePaid, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestOCR2AggregatorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*TestOCR2AggregatorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestOCR2AggregatorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*TestOCR2AggregatorOwnershipTransferred, error)

	FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*TestOCR2AggregatorPayeeshipTransferRequestedIterator, error)

	WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error)

	ParsePayeeshipTransferRequested(log types.Log) (*TestOCR2AggregatorPayeeshipTransferRequested, error)

	FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*TestOCR2AggregatorPayeeshipTransferredIterator, error)

	WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error)

	ParsePayeeshipTransferred(log types.Log) (*TestOCR2AggregatorPayeeshipTransferred, error)

	FilterRemovedAccess(opts *bind.FilterOpts) (*TestOCR2AggregatorRemovedAccessIterator, error)

	WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorRemovedAccess) (event.Subscription, error)

	ParseRemovedAccess(log types.Log) (*TestOCR2AggregatorRemovedAccess, error)

	FilterRequesterAccessControllerSet(opts *bind.FilterOpts) (*TestOCR2AggregatorRequesterAccessControllerSetIterator, error)

	WatchRequesterAccessControllerSet(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorRequesterAccessControllerSet) (event.Subscription, error)

	ParseRequesterAccessControllerSet(log types.Log) (*TestOCR2AggregatorRequesterAccessControllerSet, error)

	FilterRoundRequested(opts *bind.FilterOpts, requester []common.Address) (*TestOCR2AggregatorRoundRequestedIterator, error)

	WatchRoundRequested(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorRoundRequested, requester []common.Address) (event.Subscription, error)

	ParseRoundRequested(log types.Log) (*TestOCR2AggregatorRoundRequested, error)

	FilterTransmitted(opts *bind.FilterOpts) (*TestOCR2AggregatorTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*TestOCR2AggregatorTransmitted, error)

	FilterValidatorConfigSet(opts *bind.FilterOpts, previousValidator []common.Address, currentValidator []common.Address) (*TestOCR2AggregatorValidatorConfigSetIterator, error)

	WatchValidatorConfigSet(opts *bind.WatchOpts, sink chan<- *TestOCR2AggregatorValidatorConfigSet, previousValidator []common.Address, currentValidator []common.Address) (event.Subscription, error)

	ParseValidatorConfigSet(log types.Log) (*TestOCR2AggregatorValidatorConfigSet, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}

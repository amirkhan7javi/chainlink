package contracts

import (
	"errors"
	"github.com/smartcontractkit/chainlink-testing-framework/contracts/ethereum"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
)

// ContractLoader is an interface for abstracting the contract loading methods across network implementations
type ContractLoader interface {
	LoadKeeperConsumerBenchmark(address common.Address) (KeeperConsumerBenchmark, error)
}

// NewContractLoader returns an instance of a contract Loader based on the client type
func NewContractLoader(bcClient blockchain.EVMClient) (ContractLoader, error) {
	switch clientImpl := bcClient.Get().(type) {
	case *blockchain.EthereumClient:
		return NewEthereumContractLoader(clientImpl), nil
	case *blockchain.KlaytnClient:
		return &KlaytnContractLoader{NewEthereumContractLoader(clientImpl)}, nil
	case *blockchain.MetisClient:
		return &MetisContractLoader{NewEthereumContractLoader(clientImpl)}, nil
	case *blockchain.ArbitrumClient:
		return &MetisContractLoader{NewEthereumContractLoader(clientImpl)}, nil

	}
	return nil, errors.New("unknown blockchain client implementation for contract Loader, register blockchain client in NewContractLoader")
}

// EthereumContractLoader provides the implementations for deploying ETH (EVM) based contracts
type EthereumContractLoader struct {
	client blockchain.EVMClient
}

// KlaytnContractLoader wraps ethereum contract deployments for Klaytn
type KlaytnContractLoader struct {
	*EthereumContractLoader
}

// MetisContractLoader wraps ethereum contract deployments for Metis
type MetisContractLoader struct {
	*EthereumContractLoader
}

// ArbitrumContractLoader wraps for Arbitrum
type ArbitrumContractLoader struct {
	*EthereumContractLoader
}

// NewEthereumContractLoader returns an instantiated instance of the ETH contract Loader
func NewEthereumContractLoader(ethClient blockchain.EVMClient) *EthereumContractLoader {
	return &EthereumContractLoader{
		client: ethClient,
	}
}

// LoadKeeperConsumerBenchmark returns deployed on given address Keeper Consumer Contract
func (e *EthereumContractLoader) LoadKeeperConsumerBenchmark(address common.Address) (KeeperConsumerBenchmark, error) {
	instance, err := e.client.LoadContract("KeeperConsumerBenchmark", address, func(
		address common.Address,
		backend bind.ContractBackend,
	) (interface{}, error) {
		return ethereum.NewKeeperConsumerBenchmark(address, backend)
	})
	if err != nil {
		return nil, err
	}
	return &EthereumKeeperConsumerBenchmark{
		address:  &address,
		client:   e.client,
		consumer: instance.(*ethereum.KeeperConsumerBenchmark),
	}, err
}

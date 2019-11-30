package ethereum

import (
	ethclient "github.com/ethereum/go-ethereum/ethclient"
)

// CreateEthClient creates an ethereum client
func CreateEthClient(network string) (*ethclient.Client, error) {
	conn, err := ethclient.Dial(network)
	return conn, err
}

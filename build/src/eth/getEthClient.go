package eth

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetEthClient() (*ethclient.Client, error) {
	gethRpc := "wss://mainnet.infura.io/ws/v3/e6c920580178424bbdf6dde266bfb5bd"
	client, err := ethclient.Dial(gethRpc)
	if err != nil {
		err := fmt.Errorf("unable to connect to %s. %s", gethRpc, err)
		return nil, err
	}

	return client, nil
}
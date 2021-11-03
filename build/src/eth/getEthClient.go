package eth

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetEthClient(rpc string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	return client, nil
}
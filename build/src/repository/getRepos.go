package repository

import (
	"announcements-bot/params"
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetRepos(ethClient *ethclient.Client) ([]params.NewRepoEvent, error) {
	contractAddress := common.HexToAddress("0x266bfdb2124a68beb6769dc887bd655f78778923")
	firstBlock := big.NewInt(10153269)
	latestBlock, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	query := ethereum.FilterQuery{
        FromBlock:  firstBlock,
        ToBlock:    big.NewInt(int64(latestBlock)),
        Addresses: []common.Address{
            contractAddress,
        },
    }

    logs, err := ethClient.FilterLogs(context.Background(), query)
    if err != nil {
        return nil, err
    }

	contractAbi, err := abi.JSON(strings.NewReader(params.RepositoryAbi))
    if err != nil {
        return nil, err
    }

	var repos []params.NewRepoEvent

    for _, vLog := range logs {
		// This smart contract has multiple events
		// Continue until found "NewRepo"
        event, err := contractAbi.Unpack("NewRepo", vLog.Data)
        if err != nil {
            continue
		}

		// Parse event
		id := event[0].([32]uint8)
		name := event[1].(string)
		address := event[2].(common.Address)

		eventParsed := params.NewRepoEvent{Id: common.BytesToAddress(id[:]), Name: name, Address: address}
		repos = append(repos, eventParsed)
    }

	return repos, nil
}



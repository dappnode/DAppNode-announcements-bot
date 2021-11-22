package repository

import (
	"announcements-bot/params"
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetRepos(ethClient *ethclient.Client) ([]params.NewRepoEvent, error) {
	contractAddress := common.HexToAddress("0x266bfdb2124a68beb6769dc887bd655f78778923")
	firstBlock := big.NewInt(0)
	latestBlock, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	fmt.Printf(params.InfoLog + "Getting repositories of address %s from block %d to %d\n", contractAddress.Hex(), firstBlock.Int64(), latestBlock)

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
			fmt.Printf(params.WarnLog + "Error unpacking NewRepo event: %w\n", err)
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

// Utils

func ParseVersionEvent(event []interface{}) params.NewVersionEvent {
    versionId := event[0].(*big.Int)
    semanticVersion := event[1].([3]uint16)

    return params.NewVersionEvent{VersionId: versionId, SemanticVersion: semanticVersion}
}

func GetAddresses(repos []params.NewRepoEvent) (addresses []common.Address)  {
	for _, r := range repos {
		addresses = append(addresses, r.Address)	
	}
	return addresses
}

func GetNames(repos []params.NewRepoEvent) (names []string) {
	for _, r := range repos {
		names = append(names, r.Name)	
	}
	return names
}

func GetIds(repos []params.NewRepoEvent) (ids []common.Address) {
	for _, r := range repos {
		ids = append(ids, r.Id)	
	}
	return ids
}

package eth

import (
	"announcements-bot/discord"
	"announcements-bot/params"
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Subscribe to "NewVersion" events
func SubscribeNewVersion(ethClient *ethclient.Client, dc *discordgo.Session, discordChannel string, repos []params.NewRepoEvent) {
    // Create query with all repos addresses
    contractAbi, err := abi.JSON(strings.NewReader(params.RegistryAbi))
    if err != nil {
        log.Fatal(err)
    }

	query := ethereum.FilterQuery{
        Addresses:  GetAddresses(repos),
        Topics:     [][]common.Hash{{contractAbi.Events["NewVersion"].ID}},
    }

    logs := make(chan types.Log)
    sub, err := ethClient.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }

    for {
        select {
            case err := <-sub.Err():
                log.Fatal(err)
            case vLog := <-logs:       
                fmt.Println(vLog) // pointer to event log
                event, err := contractAbi.Unpack("NewVersion", vLog.Data)
                if err != nil {
                    continue
                }

                // Parse event
                versionId := event[0].(*big.Int)
                semanticVersion := event[1].([3]uint16)
                eventParsed := params.NewVersionEvent{VersionId: versionId, SemanticVersion: semanticVersion}

                // Look for vLog.Address in repos
                for _, r := range repos {
                    if r.Address == vLog.Address {
                        // Write New version message
                        discord.WriteNewVersionMessage(dc, discordChannel, &eventParsed, r.Name)
                        break
                    }
                }
            }
    }
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
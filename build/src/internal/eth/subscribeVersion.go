package eth

import (
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
func SubscribeNewVersion(ethClient *ethclient.Client, discord *discordgo.Session, discordChannel string, repos []NewRepoEvent) {
    // Create query with all repos addresses
    contractAbi, err := abi.JSON(strings.NewReader(registryAbi))
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
                eventParsed := NewVersionEvent{versionId: versionId, semanticVersion: semanticVersion}

                // Write New version message
                WriteNewVersionMessage(discord, discordChannel, &eventParsed)
            }
    }
}

// Utils

func ParseVersionEvent(event []interface{}) NewVersionEvent {
    versionId := event[0].(*big.Int)
    semanticVersion := event[1].([3]uint16)

    return NewVersionEvent{versionId: versionId, semanticVersion: semanticVersion}
}
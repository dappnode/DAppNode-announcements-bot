package eth

import (
	"announcements-bot/discord"
	"announcements-bot/params"
	"announcements-bot/repository"
	"context"
	"fmt"
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
        err := fmt.Errorf(params.ErrorLog + "error parsing registry abi. %w", err)
        panic(err)
    }

	query := ethereum.FilterQuery{
        Addresses:  repository.GetAddresses(repos),
        Topics:     [][]common.Hash{{contractAbi.Events["NewVersion"].ID}},
    }

    logs := make(chan types.Log)
    sub, err := ethClient.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        err := fmt.Errorf(params.ErrorLog + "error subscribing to registry logs. %w", err)
        panic(err)
    }

    for {
        select {
            case err := <-sub.Err():
                fmt.Println(err.Error())
            case vLog := <-logs:       
                fmt.Printf(params.InfoLog + "new version released: %s\n", vLog)
                event, err := contractAbi.Unpack("NewVersion", vLog.Data)
                if err != nil {
                    fmt.Printf(params.WarnLog + "error unpacking NewVersion event: %w\n", err)
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


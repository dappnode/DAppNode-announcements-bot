package eth

import (
	"announcements-bot/discord"
	"announcements-bot/params"
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Subscribe to "NewRepo" events
func SubscribeNewRepo(client *ethclient.Client, dc *discordgo.Session, discordChannel string) {
	contractAddress := common.HexToAddress("0x266bfdb2124a68beb6769dc887bd655f78778923")

    contractAbi, err := abi.JSON(strings.NewReader(params.RepositoryAbi))
    if err != nil {
        err := fmt.Errorf(params.ErrorLog + "error parsing directory abi. %w", err)
        panic(err.Error())
    }

    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
        Topics:    [][]common.Hash{{contractAbi.Events["NewRepo"].ID}},
    }

    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        err := fmt.Errorf(params.ErrorLog + "error subscribing to logs. %w", err)
        panic(err.Error())
    }

    for {
        select {
            case err := <-sub.Err():
                fmt.Println(err.Error())
            case vLog := <-logs:
                fmt.Printf(params.InfoLog + "new package released: %s\n", vLog)
                event, err := contractAbi.Unpack("NewRepo", vLog.Data)
                if err != nil {
                    fmt.Printf(params.WarnLog + "error unpacking NewRepo event: %w\n", err.Error())
                    continue
                }

                // Write New version message
                eventParsed := ParseRepoEvent(event)
                discord.WriteNewRepoMessage(dc, discordChannel, &eventParsed)
        }
    }
}

// Utils

func ParseRepoEvent(event []interface{}) params.NewRepoEvent{
    id := event[0].([32]uint8)
    name := event[1].(string)
    address := event[2].(common.Address)

    return params.NewRepoEvent{Id: common.BytesToAddress(id[:]), Name: name, Address: address}
}
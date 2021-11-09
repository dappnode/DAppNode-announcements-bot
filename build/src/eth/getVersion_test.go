package eth

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func TestGetVersion(t *testing.T) {
	t.Log("heyyy")
	ethClient, err := GetEthClient("wss://mainnet.infura.io/ws/v3/e6c920580178424bbdf6dde266bfb5bd")
	if(err != nil){
		t.Error(err)
	}
	contractAddress := common.HexToAddress("0xda33dA12E19a3579f1b522D0D215D2333431173e")
	firstBlock := big.NewInt(13581149)
	latestBlock := big.NewInt(13581151)

	query := ethereum.FilterQuery{
		FromBlock: firstBlock,
		ToBlock:   latestBlock,
		Addresses: []common.Address{
			contractAddress,
		},
	}


    logs, err := ethClient.FilterLogs(context.Background(), query)
    if err != nil {
        t.Error(err)
    }

	contractAbi, err := abi.JSON(strings.NewReader(registryAbi))
    if err != nil {
        t.Error(err)
    }

	for _, vLog := range logs {
		// This smart contract has multiple events
		// Continue until found "NewRepo"
        event, err := contractAbi.Unpack("NewVersion", vLog.Data)
        if err != nil {
            t.Error(err)
		}
		
		// Parse event
		versionId := event[0].(*big.Int)
		semanticVersion := event[1].([3]uint16)
		eventParsed := NewVersionEvent{versionId: versionId, semanticVersion: semanticVersion}
		fmt.Println(eventParsed)
    }

}
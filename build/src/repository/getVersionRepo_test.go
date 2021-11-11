package repository

import (
	"announcements-bot/env"
	"announcements-bot/eth"
	"announcements-bot/params"
	"context"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var gethRpc string

func init() {
	env.LoadEnv()
	gethRpc = os.Getenv("GETH_RPC")

	if gethRpc == "" {
		panic("gethRpc not set")
	}
}

func TestGetVersion(t *testing.T) {
	gethClient, err := ethclient.Dial(gethRpc)
	if err != nil {
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

    logs, err := gethClient.FilterLogs(context.Background(), query)
    if err != nil {
        t.Error(err)
    }

	contractAbi, err := abi.JSON(strings.NewReader(params.RegistryAbi))
    if err != nil {
        t.Error(err)
    }

	for _, vLog := range logs {
        event, err := contractAbi.Unpack("NewVersion", vLog.Data)
        if err != nil {
            t.Error(err)
		}

		eventParsed := eth.ParseVersionEvent(event)
		
		if eventParsed.VersionId.Cmp(big.NewInt(16)) != 0 {
			t.Error("Version ID is not 16")
		}
		if eventParsed.SemanticVersion != [3]uint16{0, 1, 15} {
			t.Error("Semantic version is not [0 1 15]")
		}
    }
}

func TestGetRepos(t *testing.T) {
	gethClient, err := ethclient.Dial(gethRpc)
	if err != nil {
		t.Error(err)
	}

	repos, err := GetRepos(gethClient)
	if(err != nil){
		t.Error(err)
	}

	rotkiEvent := repos[20]
	expectedId := common.HexToAddress("0x8B7a2eD2997A9a0cD635ba6AC74FC58b2a38aca1")
	expectedAddress := common.HexToAddress("0x8730413f2d7aF5a0cF63a988a0F6417fec05F328")
	expectedName := "rotki"

	if rotkiEvent.Name != expectedName {
		t.Error("rotki event name is not rotki")
	}
	if rotkiEvent.Id != expectedId {
		t.Errorf("event id expected %s but received %s", rotkiEvent.Id.Hex(), expectedId.Hex())
	}
	if rotkiEvent.Address != expectedAddress {
		t.Errorf("event id expected %s but received %s", rotkiEvent.Address.Hex(), expectedAddress.Hex())
	}
}
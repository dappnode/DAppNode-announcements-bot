package eth

import (
	"announcements-bot/env"
	"context"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	env.LoadEnv()
}

func TestGetVersion(t *testing.T) {
	ethClient, err := GetEthClient(os.Getenv("GETH_RPC"))
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
        event, err := contractAbi.Unpack("NewVersion", vLog.Data)
        if err != nil {
            t.Error(err)
		}

		eventParsed := ParseVersionEvent(event)
		
		if eventParsed.versionId.Cmp(big.NewInt(16)) != 0 {
			t.Error("Version ID is not 16")
		}
		if eventParsed.semanticVersion != [3]uint16{0, 1, 15} {
			t.Error("Semantic version is not [0 1 15]")
		}
    }
}
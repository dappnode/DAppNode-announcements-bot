package eth

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var ethClient, _ = GetEthClient()

func TestNewVersionSub(t *testing.T) {
	contractAddress := common.HexToAddress("0x448BfB454718f20941FE8a1bfA63a0024F21Ba50")

	query := ethereum.FilterQuery{
        FromBlock:  big.NewInt(13481339),
        ToBlock:    big.NewInt(13481341),
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

	// Expected values
	expectedBlockHash := "0xf1abf58dfd3ce477c34a287a8d45c671250716d449bb3c13d7c10476f2b9b5a9"
	expectedBlockNumber := uint64(13481340)
	expectedTxHash := "0xe9aa0e89ea3b470fbb8cdf2c200a65c34391762d9b3dbd727828bf797b0200e9"
	expectedSemanticVersion := [3]uint16{0,1,4}
	expectedVersionId := big.NewInt(5)

    for _, vLog := range logs {
		// Check transaction is expected
		if vLog.BlockHash.Hex() != expectedBlockHash ||  vLog.BlockNumber != expectedBlockNumber || vLog.TxHash.Hex() != expectedTxHash {
			t.Errorf("Unexpected transaction. blockHash: %s. blockNumber: %d. txHash: %s", vLog.BlockHash.Hex(), vLog.BlockNumber, vLog.TxHash.Hex())
		}  

        event,err := contractAbi.Unpack("NewVersion", vLog.Data)
        if err != nil {
            t.Error(err)
        }

		// Check event is expected
        eventParsed := NewVersionEvent{versionId: event[0].(*big.Int), semanticVersion: event[1].([3]uint16)}
		if eventParsed.semanticVersion != expectedSemanticVersion || eventParsed.versionId.Cmp(expectedVersionId) != 0 {
			t.Errorf("Unexpected version")
		}
    }
}

func TestNewRepoSub(t *testing.T) {
	contractAddress := common.HexToAddress("0x266bfdb2124a68beb6769dc887bd655f78778923")

	query := ethereum.FilterQuery{
        FromBlock:  big.NewInt(13144555),
        ToBlock:    big.NewInt(13144557),
        Addresses: []common.Address{
            contractAddress,
        },
    }

    logs, err := ethClient.FilterLogs(context.Background(), query)
    if err != nil {
        t.Error(err)
    }

	contractAbi, err := abi.JSON(strings.NewReader(repositoryAbi))
    if err != nil {
        t.Error(err)
    }

	// Expected values
	expectedBlockHash := "0x2d8f1ba90bf6214f8d55f195ae9969e17ddd760d300c8f1c98837370524ca926"
	expectedBlockNumber := uint64(13144556)
	expectedTxHash := "0x76393f8858866426c9a996ac00e3d184d68dc3e95efebca2b4883a383c57492c"
	expectedName := "ssv-prater"
	expectedScAddress := common.HexToAddress("0xEa6D433366C42faECbd4C604E81A138d28666c59")

    for _, vLog := range logs {
		// Check transaction is expected
        if vLog.BlockHash.Hex() != expectedBlockHash ||  vLog.BlockNumber != expectedBlockNumber || vLog.TxHash.Hex() != expectedTxHash {
			t.Errorf("Unexpected transaction. blockHash: %s. blockNumber: %d. txHash: %s", vLog.BlockHash.Hex(), vLog.BlockNumber, vLog.TxHash.Hex())
		}

		// This smart contract has multiple events
		// Continue until found "NewRepo"
        event, err := contractAbi.Unpack("NewRepo", vLog.Data)
        if err != nil {
            continue
		}

		// Check event is expected
		eventParsed := NewRepoEvent{id: event[0].([32]byte), name: event[1].(string), repo: event[2].(common.Address)}
		if eventParsed.name != expectedName || eventParsed.repo != expectedScAddress {
			t.Errorf("Unexpected log. Name: %s. ScAddress: %s", event[1], event[2])
		}  
    }
}
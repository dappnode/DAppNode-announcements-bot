package eth

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestGetRepos(t *testing.T) {
	ethClient, err := GetEthClient(os.Getenv("GETH_RPC"))
	if(err != nil){
		t.Error(err)
	}

	repos, err := GetRepos(ethClient)
	if(err != nil){
		t.Error(err)
	}

	rotkiEvent := repos[20]
	expectedId := common.HexToAddress("0x8B7a2eD2997A9a0cD635ba6AC74FC58b2a38aca1")
	expectedAddress := common.HexToAddress("0x8730413f2d7aF5a0cF63a988a0F6417fec05F328")
	expectedName := "rotki"

	if rotkiEvent.name != expectedName {
		t.Error("rotki event name is not rotki")
	}
	if rotkiEvent.id != expectedId {
		t.Errorf("event id expected %s but received %s", rotkiEvent.id.Hex(), expectedId.Hex())
	}
	if rotkiEvent.address != expectedAddress {
		t.Errorf("event id expected %s but received %s", rotkiEvent.address.Hex(), expectedAddress.Hex())
	}
}
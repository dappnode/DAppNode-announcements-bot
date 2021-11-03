package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type NewVersionEvent struct {
    versionId *big.Int
    semanticVersion [3]uint16
}

type NewRepoEvent struct {
	id   common.Address
	name string
	address common.Address
} 
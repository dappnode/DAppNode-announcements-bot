package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type NewVersionEvent struct {
    VersionId *big.Int
    SemanticVersion [3]uint16
}

type NewRepoEvent struct {
	Id   common.Address
	Name string
	Address common.Address
} 
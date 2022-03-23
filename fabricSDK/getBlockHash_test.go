package fabricSDK

import (
	"fmt"
	"sidechain/fabricSDK"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Println(fabricSDK.GetBlockNumber())
}

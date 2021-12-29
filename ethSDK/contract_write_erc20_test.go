package ethSDK

import (
	"fmt"
	"sidechain/ethSDK"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Println(ethSDK.Transfer("/home/eth-poa/signer1/data/geth.ipc", "0xD78d66C33933a05c57c503d61667918f95cee351", "f48dec702794f00ad8c7af81dfd1220721c4d38d97c87a9b2d3cb8663f80a1f4", "0xf745069D290dE951508CA088D198678758DcA46c", "1"))
}

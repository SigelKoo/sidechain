package ethSDK

import (
	"fmt"
	"sidechain/ethSDK"
	"testing"
)

func TestHello(t *testing.T) {
	//ethSDK.TransferTokens("/home/eth-poa/signer1/data/geth.ipc", "0xD78d66C33933a05c57c503d61667918f95cee351", "8c7ee582167250ee80c52d813f1747592e78c6c311d3576fa15570662b63dd74", "0x416b1e5329Bd97BB704866bD489747b26848fA42", "100")
	fmt.Println(ethSDK.Transfer("/home/eth-poa/signer1/data/geth.ipc", "0xD78d66C33933a05c57c503d61667918f95cee351", "f48dec702794f00ad8c7af81dfd1220721c4d38d97c87a9b2d3cb8663f80a1f4", "0xf745069D290dE951508CA088D198678758DcA46c", "1"))
}

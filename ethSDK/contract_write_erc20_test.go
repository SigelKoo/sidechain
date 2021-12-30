package ethSDK

import (
	"fmt"
	"sidechain/ethSDK"
	"testing"
)

func TestHello(t *testing.T) {
	//ethSDK.TransferTokens("/home/eth-poa/signer1/data/geth.ipc", "0xD78d66C33933a05c57c503d61667918f95cee351", "8c7ee582167250ee80c52d813f1747592e78c6c311d3576fa15570662b63dd74", "0x416b1e5329Bd97BB704866bD489747b26848fA42", "100")
	fmt.Println(ethSDK.Transfer("/home/eth-poa/signer1/data/geth.ipc", "0xD78d66C33933a05c57c503d61667918f95cee351", "3e0aae2e1274a2ce81838d07c1c4a2fb840d691de35f9a0f641d459ac159e7f8", "0xf745069D290dE951508CA088D198678758DcA46c", "10"))
}

package main

import (
	"sidechain/ethfabricListen"
)

func main() {
	// fmt.Println(ethSDK.ContractDeploy("HTTP://222.201.187.76:8501", "8c7ee582167250ee80c52d813f1747592e78c6c311d3576fa15570662b63dd74"))
	// 0xD78d66C33933a05c57c503d61667918f95cee351 0xa9aec05c96c8f822bbf4d6dcccc67630835b04c7ca6443ca5686fc8ff5610db1
	// fmt.Println(ethSDK.GetUserBalance("HTTP://222.201.187.76:8501", "0xD78d66C33933a05c57c503d61667918f95cee351", "0x60BD95E835ADe2552545DfC21ADB23069A0A7aD4"))
	// fmt.Println(ethSDK.GetTokenInfo("HTTP://222.201.187.76:8501", "0xD78d66C33933a05c57c503d61667918f95cee351"))
	// fmt.Println(ethSDK.Transfer("HTTP://222.201.187.76:8501", "0xD78d66C33933a05c57c503d61667918f95cee351", "8c7ee582167250ee80c52d813f1747592e78c6c311d3576fa15570662b63dd74", "0x416b1e5329Bd97BB704866bD489747b26848fA42", "100"))
	go ethfabricListen.Eth_listen_erc20_transfer("/home/eth-poa/signer1/data/geth.ipc", "0xD78d66C33933a05c57c503d61667918f95cee351")
	// fabricSDK.GetBlockNumber()
}

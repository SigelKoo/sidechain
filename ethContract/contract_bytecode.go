package ethContract

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func ContractBytecode(url string, address string) string {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(address)
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytecode)
}

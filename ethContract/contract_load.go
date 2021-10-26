package ethContract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	store "sidechain/ethContract/contracts"
)

func ContractLoad(url string, address string) *store.Store {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(address)
	instance, err := store.NewStore(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance
}

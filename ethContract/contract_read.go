package ethContract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	store "sidechain/ethContract/contracts"
)

func ContractRead(url string, address string, BlockNumber *big.Int) string {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(address)
	instance, err := store.NewStore(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	blockHash, err := instance.Get(nil, BlockNumber)
	if err != nil {
		log.Fatal(err)
	}
	return blockHash
}

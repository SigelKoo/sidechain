package ethSDK

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"sidechain/ethContract/openzeppelin-contracts/contracts/token/ERC20"
)

func ContractDeploy(url string, privateString string) string {
	// store "2e8749fd1ba7a42586d2bb38c10fab2e8845abd7733378a95a03fdcdbd1b854e"
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(privateString)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetUint64(uint64(57825)))
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	// address, tx, _, err := store.DeployStore(auth, client)
	address, tx, _, err := ERC20.DeployERC20(auth, client, "SidechainCoin", "sc")
	if err != nil {
		log.Fatal(err)
	}
	return address.Hex() + " " + tx.Hash().Hex()
}

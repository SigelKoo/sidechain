package ethSDK

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	token_erc20 "sidechain/ethContract/openzeppelin-contracts/contracts/token/ERC20"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Transfer(url string, contractAddress string, privateString string, to string, value string) string {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress(contractAddress)
	instance, err := token_erc20.NewTokenErc20(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privateString)
	if err != nil {
		log.Fatal(err)
	}
	opts := bind.NewKeyedTransactor(privateKey)

	publicKey := privateKey.Public()
	publicKeyECSDA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECSDA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	toAddress := common.HexToAddress(to)

	parsed, err := abi.JSON(strings.NewReader(token_erc20.TokenErc20ABI))
	input, err := parsed.Pack("transfer", toAddress, new(big.Int).SetString(value, 10))
	if err != nil {
		log.Fatal(err)
	}

	amount := opts.opts.Value
	if amount == nil {
		amount = new(big.Int)
	}

	tx, _ := erc20.BuildTransfer(opts, address, web3go.NewBigInt(10))
	fmt.Println(tx.GetHash().GetHex())
	tx, _ = erc20.Transfer(opts, address, web3go.NewBigInt(10))
	fmt.Println(tx.GetHash().GetHex())

	//打印账户token余额
	address, _ = web3go.NewAddressFromHex("0xfe04cb1d7d6715169edc07c8e3c2fdba3a0854af")
	balance, err = erc20.BalanceOf(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance.String())
}

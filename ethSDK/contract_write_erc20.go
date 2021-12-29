package ethSDK

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	token_erc20 "sidechain/ethContract/openzeppelin-contracts/contracts/token/ERC20"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetUint64(uint64(686868)))
	if err != nil {
		log.Fatal(err)
	}

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

	opts.Value = big.NewInt(0)
	opts.GasLimit = uint64(300000)
	opts.GasPrice = gasPrice

	toAddress := common.HexToAddress(to)

	parsed, err := abi.JSON(strings.NewReader(token_erc20.TokenErc20MetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	bigintValue, b := new(big.Int).SetString(value, 10)
	if !b {
		log.Fatal(b)
	}

	input, err := parsed.Pack("transfer", toAddress, bigintValue)
	if err != nil {
		log.Fatal(err)
	}

	rawTx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &tokenAddress,
		Value:    opts.Value,
		Gas:      opts.GasLimit,
		GasPrice: opts.GasPrice,
		Data:     input,
	})
	signedTx, err := opts.Signer(fromAddress, rawTx)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.Transfer(opts, toAddress, bigintValue)
	if err != nil {
		log.Fatal(err)
	}

	return signedTx.Hash().String() + "," + tx.Hash().String()
}

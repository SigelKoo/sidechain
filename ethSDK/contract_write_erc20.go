package ethSDK

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	token_erc20 "sidechain/ethContract/openzeppelin-contracts/contracts/token/ERC20"

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

func TransferTokens(url string, contractAddress string, privateString string, to string, amountStr string) {
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
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(to)
	tokenAddress := common.HexToAddress(contractAddress)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Printf("Method ID: %s\n", hexutil.Encode(methodID))

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Printf("To address: %s\n", hexutil.Encode(paddedAddress))

	amount := new(big.Int)
	amount.SetString(amountStr, 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Printf("Token amount: %s", hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Gas limit: %d", gasLimit)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tokens sent at TX: %s", signedTx.Hash().Hex())
}

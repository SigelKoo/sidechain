package ethSDK

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	store "sidechain/ethContract/store-contracts"
)

type StoreSetup struct {
	client         *ethclient.Client
	privateKey     *ecdsa.PrivateKey
	publicKeyECSDA *ecdsa.PublicKey
	auth           *bind.TransactOpts
	instance       *store.Store
}

func InitStoreSetup(url string, privateString string, address string) (*StoreSetup, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateString)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECSDA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECSDA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetUint64(uint64(686868)))
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	contractAddress := common.HexToAddress(address)
	instance, err := store.NewStore(contractAddress, client)
	if err != nil {
		return nil, err
	}

	return &StoreSetup{
		client:         client,
		privateKey:     privateKey,
		publicKeyECSDA: publicKeyECSDA,
		auth:           auth,
		instance:       instance,
	}, nil
}

func (t *StoreSetup) Set(blockNumber string, hash string) (string, error) {
	num := new(big.Int)
	num, ok := num.SetString(blockNumber, 10)
	if !ok {
		return "", fmt.Errorf("SetString: error")
	}
	tx, err := t.instance.Set(t.auth, num, hash)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (t *StoreSetup) Get(blockNumber string) (string, error) {
	num := new(big.Int)
	num, ok := num.SetString(blockNumber, 10)
	if !ok {
		return "", fmt.Errorf("SetString: error")
	}
	res, err := t.instance.Get(nil, num)
	if err != nil {
		return "", err
	}
	return res, nil
}

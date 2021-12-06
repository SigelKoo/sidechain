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
	token_erc20 "sidechain/ethContract/openzeppelin-contracts/contracts/token/ERC20"
)

type ServiceSetup struct {
	client         *ethclient.Client
	privateKey     *ecdsa.PrivateKey
	publicKeyECSDA *ecdsa.PublicKey
	auth           *bind.TransactOpts
	instance       *token_erc20.TokenErc20
}

func InitServiceSetup(url string, privateString string, address string) (*ServiceSetup, error) {
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
	instance, err := token_erc20.NewTokenErc20(contractAddress, client)
	if err != nil {
		return nil, err
	}

	return &ServiceSetup{
		client:         client,
		privateKey:     privateKey,
		publicKeyECSDA: publicKeyECSDA,
		auth:           auth,
		instance:       instance,
	}, nil
}

func (t *ServiceSetup) Transfer(recipient string, amount string) (string, error) {
	num := new(big.Int)
	num, ok := num.SetString(amount, 10)
	if !ok {
		return "", fmt.Errorf("SetString: error")
	}
	tx, err := t.instance.Transfer(t.auth, common.HexToAddress(recipient), num)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (t *ServiceSetup) BalanceOf(account string) (string, error) {
	res, err := t.instance.BalanceOf(nil, common.HexToAddress(account))
	if err != nil {
		return "", err
	}
	return (*res).String(), nil
}

func (t *ServiceSetup) Burn(amount string) (string, error) {
	num := new(big.Int)
	num, ok := num.SetString(amount, 10)
	if !ok {
		return "", fmt.Errorf("SetString: error")
	}
	tx, err := t.instance.Transfer(t.auth, common.HexToAddress("0x0"), num)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

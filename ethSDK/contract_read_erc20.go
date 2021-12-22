package ethSDK

import (
	"fmt"
	"log"
	"math"
	"math/big"

	token_erc20 "sidechain/ethContract/openzeppelin-contracts/contracts/token/ERC20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	// 首先创建一个ERC20智能合约interface。 这只是与您可以调用的函数的函数定义的契约。
)

// 查询ERC20代币智能合约
func GetUserBalance(url string, contractAddress string, userAddress string) string {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	// 假设我们已经像往常一样设置了以太坊客户端，我们现在可以将新的token包导入我们的应用程序并实例化它。
	// 这个例子里我们用BNB代币的地址. https://etherscan.io/token/0xB8c77482e45F1F44dE1745F52C74426C631bDD52
	tokenAddress := common.HexToAddress(contractAddress)
	instance, err := token_erc20.NewTokenErc20Caller(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// 我们现在可以调用任何ERC20的方法。 例如，我们可以查询用户的代币余额。
	address := common.HexToAddress(userAddress)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("wei: %s\n", bal)

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// 我们可以做一些简单的数学运算将余额转换为可读的十进制格式。
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	return value.String()
}

func GetTokenInfo(url string, contractAddress string) string {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress(contractAddress)
	instance, err := token_erc20.NewTokenErc20Caller(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)
	fmt.Printf("symbol: %s\n", symbol)
	fmt.Printf("decimals: %v\n", decimals)

	return name + "," + symbol + "," + string(decimals)
}

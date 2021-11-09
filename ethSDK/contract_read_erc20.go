package ethSDK

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	// 首先创建一个ERC20智能合约interface。 这只是与您可以调用的函数的函数定义的契约。
	token "sidechain/ethContract/openzeppelin-contracts/contracts/token/ERC20"
	"log"
	"math"
	"math/big"
)

// 查询ERC20代币智能合约
func main() {
	client, err := ethclient.Dial("HTTP://192.168.132.80:8501")
	if err != nil {
		log.Fatal(err)
	}

	// 假设我们已经像往常一样设置了以太坊客户端，我们现在可以将新的token包导入我们的应用程序并实例化它。
	// 这个例子里我们用BNB代币的地址. https://etherscan.io/token/0xB8c77482e45F1F44dE1745F52C74426C631bDD52
	tokenAddress := common.HexToAddress("0xB8c77482e45F1F44dE1745F52C74426C631bDD52")
	instance, err := token.NewERC20Caller(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// 我们现在可以调用任何ERC20的方法。 例如，我们可以查询用户的代币余额。
	address := common.HexToAddress("0xd62933233c74ada45ff997baf625267f29b2d6ad")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	// 我们还可以读ERC20智能合约的公共变量。
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

	fmt.Printf("wei: %s\n", bal)

	// 我们可以做一些简单的数学运算将余额转换为可读的十进制格式。
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f", value)
}
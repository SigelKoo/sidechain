package ethSDK

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	token "sidechain/ethContract/openzeppelin-contracts/contracts/token/ERC20"
	"strings"
)

// 现在在我们的Go应用程序中，让我们创建与ERC-20事件日志签名类型相匹配的结构类型：
type LogTransfer struct {
	From common.Address
	To common.Address
	Value *big.Int
}

type LogApproval struct {
	TokenOwner common.Address
	Spender common.Address
	Value *big.Int
}

// 读取ERC-20代币的事件日志
func eventListenERC20(url string, address string, start, end *big.Int) {
	// 初始化以太坊客户端
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	// 按照ERC-20智能合约地址和所需的块范围创建一个“FilterQuery”:
	contractAddress := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		FromBlock: start,
		ToBlock: end,
		//FromBlock: big.NewInt(6383820),
		//ToBlock:   big.NewInt(6383840),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	// 用FilterLogs来过滤日志：
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// 接下来我们将解析JSON abi，稍后我们将使用解压缩原始日志数据：
	contractAbi, err := abi.JSON(strings.NewReader(string(token.ERC20MetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}

	// 为了按某种日志类型进行过滤，我们需要弄清楚每个事件日志函数签名的keccak256哈希值。
	// 事件日志函数签名哈希始终是topic [0]，我们很快就会看到。
	// 以下是使用go-ethereumcrypto包计算keccak256哈希的方法：
	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	// 现在我们将遍历所有日志并设置switch语句以按事件日志类型进行过滤：
	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		// 现在要解析Transfer事件日志，我们将使用abi.Unpack将原始日志数据解析为我们的日志类型结构。
		// 解包不会解析indexed事件类型，因为它们存储在topics下，所以对于那些我们必须单独解析，如下例所示：
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent LogTransfer

			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Value.String())
		// Approval 日志也是类似的方法：
		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

			var approvalEvent LogApproval

			err := contractAbi.UnpackIntoInterface(&approvalEvent, "Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Tokens: %s\n", approvalEvent.Value.String())
		}

		fmt.Printf("\n\n")
	}
}

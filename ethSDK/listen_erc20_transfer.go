package ethSDK

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	token_erc20 "sidechain/ethContract/openzeppelin-contracts/contracts/token/ERC20"
	"sidechain/fabricSDK"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type LogTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

// 监听事件日志
func Eth_listen_erc20_transfer(url string, address string) {
	// 为了订阅事件日志，我们需要做的第一件事就是拨打启用websocket的以太坊客户端。
	// client, err := ethclient.Dial("HTTP://192.168.132.80:8501")
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	// 下一步是创建筛选查询。 在这个例子中，我们将阅读来自我们在之前课程中创建的示例合约中的所有事件。
	contractAddress := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	// 我们接收事件的方式是通过Go channel。 让我们从go-ethereumcore/types包创建一个类型为Log的channel。
	logs := make(chan types.Log)

	contractAbi, err := abi.JSON(strings.NewReader(string(token_erc20.TokenErc20MetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	// 现在我们所要做的就是通过从客户端调用SubscribeFilterLogs来订阅，它接收查询选项和输出通道。
	// 这将返回包含unsubscribe和error方法的订阅结构。
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	sdk, err := fabricSDK.SetupSDK("./config/crypto-config.yaml", false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer sdk.Close()

	// 最后，我们要做的就是使用select语句设置一个连续循环来读入新的日志事件或订阅错误。
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
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

				fmt.Printf("\n\n")
				if transferEvent.To == common.HexToAddress("0xf745069D290dE951508CA088D198678758DcA46c") {
					Org1User1Info := fabricSDK.InitInfo{
						ChannelID:     "mychannel",
						ChannelConfig: "/home/fabric-samples/test-network/channel-artifacts/mychannel.block",

						OrgAdmin:       "Admin",
						OrgName:        "Org1",
						OrdererOrgName: "orderer.example.com",

						ChaincodeID:     "token_erc20",
						ChaincodeGoPath: os.Getenv("GOPATH"),
						ChaincodePath:   "/home/fabric-samples/token-erc-20/chaincode-go",
						UserName:        "minter",
					}

					Org2UserInfo := fabricSDK.InitInfo{
						ChannelID:     "mychannel",
						ChannelConfig: "/home/fabric-samples/test-network/channel-artifacts/mychannel.block",

						OrgAdmin:       "Admin",
						OrgName:        "Org2",
						OrdererOrgName: "orderer.example.com",

						ChaincodeID:     "token_erc20",
						ChaincodeGoPath: os.Getenv("GOPATH"),
						ChaincodePath:   "/home/fabric-samples/token-erc-20/chaincode-go",
						UserName:        transferEvent.From.Hex(),
					}

					minterChannelContext := sdk.ChannelContext(Org1User1Info.ChannelID, fabsdk.WithUser(Org1User1Info.UserName), fabsdk.WithOrg(Org1User1Info.OrgName))
					clientChannelContext := sdk.ChannelContext(Org2UserInfo.ChannelID, fabsdk.WithUser(Org2UserInfo.UserName), fabsdk.WithOrg(Org2UserInfo.OrgName))
					channelMinter, err := channel.New(minterChannelContext)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
					channelClient, err := channel.New(clientChannelContext)
					if err != nil {
						fmt.Println(err.Error())
						return
					}

					minterServiceSetup := fabricSDK.ServiceSetup{
						ChaincodeID: Org1User1Info.ChaincodeID,
						Client:      channelMinter,
					}
					client1ServiceSetup := fabricSDK.ServiceSetup{
						ChaincodeID: Org2UserInfo.ChaincodeID,
						Client:      channelClient,
					}

					clientID, err := client1ServiceSetup.ClientAccountID()
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println("client，交易id为：" + clientID)
					}

					msg, err := minterServiceSetup.Transfer(clientID, transferEvent.Value.String())
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println("转账成功，交易编号为：" + msg)
					}
				}
			}
		}
	}
}

/*

Log Block Number: 5573
Log Index: 0
Log Name: Transfer
From: 0x60BD95E835ADe2552545DfC21ADB23069A0A7aD4
To: 0xb1ABb29CC3CD7b6c8D028866c370f92A2D1c870c
Tokens: 100


Log Block Number: 5600
Log Index: 0
Log Name: Transfer
From: 0x60BD95E835ADe2552545DfC21ADB23069A0A7aD4
To: 0x416b1e5329Bd97BB704866bD489747b26848fA42
Tokens: 100

*/

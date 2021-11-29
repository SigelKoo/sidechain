package main

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"os"
	"sidechain/fabricSDK"
)

func main() {
	Org2User1Info := fabricSDK.InitInfo{
		ChannelID:     "mychannel",
		ChannelConfig: "/home/fabric-samples/test-network/channel-artifacts/mychannel.block",

		OrgAdmin:       "Admin",
		OrgName:        "Org2",
		OrdererOrgName: "orderer.example.com",

		ChaincodeID:     "token_erc20",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "/home/fabric-samples/token-erc-20/chaincode-go",
		UserName:        "0xb1ABb29CC3CD7b6c8D028866c370f92A2D1c870c",
	}

	Org2User2Info := fabricSDK.InitInfo{
		ChannelID:     "mychannel",
		ChannelConfig: "/home/fabric-samples/test-network/channel-artifacts/mychannel.block",

		OrgAdmin:       "Admin",
		OrgName:        "Org2",
		OrdererOrgName: "orderer.example.com",

		ChaincodeID:     "token_erc20",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "/home/fabric-samples/token-erc-20/chaincode-go",
		UserName:        "0x416b1e5329Bd97BB704866bD489747b26848fA42",
	}

	sdk, err := fabricSDK.SetupSDK("./config/crypto-config.yaml", false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer sdk.Close()
	client1ChannelContext := sdk.ChannelContext(Org2User1Info.ChannelID, fabsdk.WithUser(Org2User1Info.UserName), fabsdk.WithOrg(Org2User1Info.OrgName))
	client2ChannelContext := sdk.ChannelContext(Org2User2Info.ChannelID, fabsdk.WithUser(Org2User2Info.UserName), fabsdk.WithOrg(Org2User2Info.OrgName))
	channelClient1, err := channel.New(client1ChannelContext)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	channelClient2, err := channel.New(client2ChannelContext)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	client1ServiceSetup := fabricSDK.ServiceSetup{
		ChaincodeID: Org2User1Info.ChaincodeID,
		Client:      channelClient1,
	}
	client2ServiceSetup := fabricSDK.ServiceSetup{
		ChaincodeID: Org2User2Info.ChaincodeID,
		Client:      channelClient2,
	}

	client1Balance, err := client1ServiceSetup.ClientAccountBalance()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("client1，账户余额为：" + client1Balance)
	}

	client1ID, err := client1ServiceSetup.ClientAccountID()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("client1，交易id为：" + client1ID)
	}

	client2ID, err := client2ServiceSetup.ClientAccountID()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("client2，交易id为：" + client2ID)
	}

	msg, err := client1ServiceSetup.Transfer(client2ID, "1")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("转账成功，交易编号为：" + msg)
	}

	client1Balance, err = client2ServiceSetup.BalanceOf(client1ID)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("client2查到的client1余额为：" + client1Balance)
	}
}

package main

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"os"
	"sidechain/fabricSDK"
)

func main() {
	initInfo := fabricSDK.InitInfo {
		ChannelID:     "mychannel",
		ChannelConfig: "/home/fabric-samples/test-network/channel-artifacts/mychannel.block",

		OrgAdmin:       "Admin",
		OrgName:        "org2.example.com",
		OrdererOrgName: "orderer.example.com",

		ChaincodeID:     "token_erc20",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "/home/fabric-samples/token-erc-20/chaincode-go",
		UserName:        "0xb1ABb29CC3CD7b6c8D028866c370f92A2D1c870c@org2.example.com",
	}

	sdk, err := fabricSDK.SetupSDK("./config/crypto-config.yaml", false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer sdk.Close()
	clientChannelContext := sdk.ChannelContext(initInfo.ChannelID, fabsdk.WithUser(initInfo.UserName), fabsdk.WithOrg(initInfo.OrgName))
	channelClient, err := channel.New(clientChannelContext)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	serviceSetup := fabricSDK.ServiceSetup{
		ChaincodeID: initInfo.ChaincodeID,
		Client:      channelClient,
	}
	recipient := "eDUwOTo6Q049MHg0MTZiMWU1MzI5QmQ5N0JCNzA0ODY2YkQ0ODk3NDdiMjY4NDhmQTQyLE9VPWNsaWVudCxPPUh5cGVybGVkZ2VyLFNUPU5vcnRoIENhcm9saW5hLEM9VVM6OkNOPWNhLm9yZzIuZXhhbXBsZS5jb20sTz1vcmcyLmV4YW1wbGUuY29tLEw9SHVyc2xleSxTVD1IYW1wc2hpcmUsQz1VSw=="
	msg, err := serviceSetup.Transfer(recipient,"1")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("转账成功, 交易编号为: " + msg)
	}
}
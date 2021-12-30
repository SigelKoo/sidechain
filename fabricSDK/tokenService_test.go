package fabricSDK

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/events/deliverclient/seek"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"os"
	"sidechain/fabricSDK"
	"testing"
)

func TestHello(t *testing.T) {
	sdk, err := fabricSDK.SetupSDK("./../config/crypto-config.yaml", false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer sdk.Close()

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
		UserName:        "0xb1ABb29CC3CD7b6c8D028866c370f92A2D1c870c",
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

	eventMinter, err := event.New(minterChannelContext, event.WithBlockEvents(), event.WithSeekType(seek.Newest))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	eventClient, err := event.New(clientChannelContext, event.WithBlockEvents(), event.WithSeekType(seek.Newest))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	minterServiceSetup := fabricSDK.ServiceSetup{
		ChaincodeID:   Org1User1Info.ChaincodeID,
		ChannelClient: channelMinter,
		EventClient:   eventMinter,
	}
	client1ServiceSetup := fabricSDK.ServiceSetup{
		ChaincodeID:   Org2UserInfo.ChaincodeID,
		ChannelClient: channelClient,
		EventClient:   eventClient,
	}

	minterID, err := minterServiceSetup.ClientAccountID()
	if err != nil {
		fmt.Println(err.Error())
	}

	value := "1"

	msg, err := client1ServiceSetup.Transfer(minterID, value)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("转账成功，交易编号为：" + msg)
	}
}

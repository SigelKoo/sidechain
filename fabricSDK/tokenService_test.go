package fabricSDK

import (
	"fmt"
	"sidechain/fabricSDK"
	"testing"
)

func TestHello(t *testing.T) {
	sdk, err := fabricSDK.SetupSDK("./config/crypto-config.yaml", false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer sdk.Close()

	Org2UserInfo := fabricSDK.InitInfo{
		ChannelID:     "mychannel",
		ChannelConfig: "/home/fabric-samples/test-network/channel-artifacts/mychannel.block",

		OrgAdmin:       "Admin",
		OrgName:        "Org2",
		OrdererOrgName: "orderer.example.com",

		ChaincodeID:     "token_erc20",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "/home/fabric-samples/token-erc-20/chaincode-go",
		UserName:        "",
	}

	clientChannelContext := sdk.ChannelContext(Org2UserInfo.ChannelID, fabsdk.WithUser(Org2UserInfo.UserName), fabsdk.WithOrg(Org2UserInfo.OrgName))

	channelClient, err := channel.New(clientChannelContext)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	eventClient, err := event.New(clientChannelContext, event.WithBlockEvents(), event.WithSeekType(seek.Newest))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client1ServiceSetup := fabricSDK.ServiceSetup{
		ChaincodeID:   Org2UserInfo.ChaincodeID,
		ChannelClient: channelClient,
		EventClient:   eventClient,
	}

	value := "1"

	msg, err := client1ServiceSetup.Burn(value)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("毁币成功，交易编号为：" + msg)
	}

}

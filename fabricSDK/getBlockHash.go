package fabricSDK

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"os"
)

func getBlockNumber() {
	sdk, err := SetupSDK("./config/crypto-config.yaml", false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer sdk.Close()
	Org1UserInfo := InitInfo{
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
	channelContext := sdk.ChannelContext(Org1UserInfo.ChannelID, fabsdk.WithUser(Org1UserInfo.UserName), fabsdk.WithOrg(Org1UserInfo.OrgName))
	client, err := ledger.New(channelContext)
	block, err := client.QueryBlock(788)
	fmt.Println(block)
}

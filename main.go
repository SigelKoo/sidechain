package main

import (
	"fmt"
	"os"
	sdkInit "sidechain/sdkinit"
	sidechain "sidechain/src"
)

var App sdkInit.Application


func main() {
	orgs := []*sdkInit.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    2,
		},
	}

	// init sdk env info
	info := sdkInit.SdkEnvInfo{
		ChannelID:        "mychannel",
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
	}

	// sdk setup
	sdk, err := sdkInit.Setup("./config/crypto-config.yaml", &info)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}
	fmt.Println(sdk)
	sidechain.InitChainBrowserService()
	sidechain.QueryLatestBlocksInfo()
}

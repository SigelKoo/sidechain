package ethfabricListen

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sidechain/ethSDK"
	"sidechain/fabricSDK"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/events/deliverclient/seek"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type thisEvent struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value int    `json:"value"`
}

func fabric_listen_erc20_transfer() {
	sdk, err := fabricSDK.SetupSDK("./config/crypto-config.yaml", false)
	if err != nil {
		log.Fatal(err)
	}
	defer sdk.Close()

	Org1UserInfo := fabricSDK.InitInfo{
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

	channelProvider := sdk.ChannelContext(Org1UserInfo.ChannelID, fabsdk.WithUser(Org1UserInfo.UserName), fabsdk.WithOrg(Org1UserInfo.OrgName))
	ec, err := event.New(channelProvider, event.WithBlockEvents(), event.WithSeekType(seek.Newest))

	if err != nil {
		log.Fatal("failed to create fabcli, error: %v", err)
	}

	registration, notifier, err := ec.RegisterChaincodeEvent(Org1UserInfo.ChaincodeID, "eventBurn")
	if err != nil {
		log.Fatal("failed to register chaincode event, error: %v", err)
	}

	defer ec.Unregister(registration)

	// todo: add context
	go func() {
		for {
			select {
			case ccEvent := <-notifier:
				if ccEvent != nil {
					ent := thisEvent{}
					json.Unmarshal(ccEvent.Payload, &ent)
					if fabricSDK.GetX509UserName(ent.To) == "minter" {
						ethSDK.Transfer("HTTP://127.0.0.1:8501", "8c7ee582167250ee80c52d813f1747592e78c6c311d3576fa15570662b63dd74", fabricSDK.GetX509UserName(ent.From), strconv.Itoa(ent.Value))
					}
				}
			case <-time.After(time.Second * 5):
				fmt.Println("timeout while waiting for chaincode event")
			}
		}
	}()
}

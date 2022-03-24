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

func (te *thisEvent) string() string {
	return te.From + "," + te.To + "," + strconv.Itoa(te.Value)
}

func Fabric_listen_erc20_transfer(done func()) {
	defer done()
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

	registration, notifier, err := ec.RegisterChaincodeEvent(Org1UserInfo.ChaincodeID, "Burn")

	if err != nil {
		log.Fatal("failed to register chaincode event, error: %v", err)
	}

	defer ec.Unregister(registration)

	for {
		select {
		case ccEvent := <-notifier:
			if ccEvent != nil {
				ent := thisEvent{}
				json.Unmarshal(ccEvent.Payload, &ent)
				fmt.Println("---------------------------------------------------------")
				fmt.Println("Fabric 交易事件：")
				fmt.Printf("Chaincode ID: %s\n", ccEvent.ChaincodeID)
				fmt.Printf("Block Number: %d\n", ccEvent.BlockNumber)
				fmt.Printf("Transaction Hash: %s\n", ccEvent.TxID)
				fmt.Printf("Event Name: %s\n", ccEvent.EventName)
				sigandver(ent.string())
				fmt.Println("转账成功，交易编号为：", ethSDK.Transfer("HTTP://127.0.0.1:8501", "0xD78d66C33933a05c57c503d61667918f95cee351", "8c7ee582167250ee80c52d813f1747592e78c6c311d3576fa15570662b63dd74", fabricSDK.GetX509UserName(ent.From), strconv.Itoa(ent.Value)))
			}
		case <-time.After(time.Second * 5):
			fmt.Println("timeout while waiting for chaincode event")
		}
	}
}

package fabricSDK

/*import (
	"encoding/json"
	"fmt"
	"os"
	"sidechain/ethSDK"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type thisEvent struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value int    `json:"value"`
}

func fabric_listen_erc20_transfer() error {
	sdk, err := SetupSDK("./config/crypto-config.yaml", false)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer sdk.Close()

	Org2UserInfo := InitInfo{
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

	channelProvider := sdk.ChannelContext(Org2UserInfo.ChannelID, fabsdk.WithUser(Org2UserInfo.UserName), fabsdk.WithOrg(Org2UserInfo.OrgName))
	ec, err := event.New(channelProvider, event.WithBlockEvents())

	if err != nil {
		return fmt.Errorf("failed to create fabcli, error: %v", err)
	}

	registration, notifier, err := ec.RegisterChaincodeEvent(Org2UserInfo.ChaincodeID, "eventBurn")
	if err != nil {
		return fmt.Errorf("failed to register chaincode event, error: %v", err)
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
					if ent.To == "0x0" {
						// Problem：此处缺智能合约的地址
<<<<<<< Updated upstream
						ethSDK.ContractWrite("HTTP://222.201.187.76:8501", "8c7ee582167250ee80c52d813f1747592e78c6c311d3576fa15570662b63dd74", "")
=======
						ss, err := ethSDK.InitServiceSetup("HTTP://192.168.132.80:8501", "8c7ee582167250ee80c52d813f1747592e78c6c311d3576fa15570662b63dd74", "")
						if err != nil {
							fmt.Println("failed to init service setup, error: %v", err)
						}
						ss.Transfer(ent.From, string(ent.Value))
					}
				}
			case <-time.After(time.Second * 5):
				fmt.Println("timeout while waiting for chaincode event")
			}
		}
	}()
	return nil
}

func fabric_loop_erc20_push_eth() error {
	sdk, err := SetupSDK("./config/crypto-config.yaml", false)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer sdk.Close()

	Org2UserInfo := InitInfo{
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

	channelProvider := sdk.ChannelContext(Org2UserInfo.ChannelID, fabsdk.WithUser(Org2UserInfo.UserName), fabsdk.WithOrg(Org2UserInfo.OrgName))
	ec, err := event.New(channelProvider, event.WithBlockEvents())

	if err != nil {
		return fmt.Errorf("failed to create fabcli, error: %v", err)
	}

	registration, notifier, err := ec.RegisterBlockEvent()
	if err != nil {
		return fmt.Errorf("failed to register block event, error: %v", err)
	}

	defer ec.Unregister(registration)

	// todo: add context
	go func() {
		for {
			select {
			case bEvent := <-notifier:
				if bEvent != nil {
					ent := bEvent.Block.Header
					if ent.Number % 10 == 0 {
						// Problem：此处缺智能合约的地址
						ss, err := ethSDK.InitStoreSetup("HTTP://192.168.132.80:8501", "8c7ee582167250ee80c52d813f1747592e78c6c311d3576fa15570662b63dd74", "")
						if err != nil {
							fmt.Println("failed to init store setup, error: %v", err)
						}
						ss.Set(string(ent.Number), string(ent.DataHash))
						// Problem：交易也传到区块链上，可能修改智能合约为merkle结构
>>>>>>> Stashed changes
					}
				}
			case <-time.After(time.Second * 5):
				fmt.Println("timeout while waiting for chaincode event")
			}
		}
	}()
	return nil
}*/

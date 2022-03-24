package main

import (
	"fmt"
	"net/http"
	"os"
	"sidechain/ethSDK"
	"sidechain/ethfabricListen"
	"sidechain/fabricSDK"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/events/deliverclient/seek"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go ethfabricListen.Eth_listen_erc20_transfer("/home/eth-poa/signer1/data/geth.ipc", "0xD78d66C33933a05c57c503d61667918f95cee351", wg.Done)
	go ethfabricListen.Fabric_listen_erc20_transfer(wg.Done)
	go startService(wg.Done)
	wg.Wait()
}

func startService(done func()) {
	defer done()
	r := gin.Default()
	r.LoadHTMLGlob("./web/**/*")
	r.StaticFS("/css", http.Dir("./web/css"))
	r.StaticFS("/js", http.Dir("./web/js"))
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET("/transfer", func(c *gin.Context) {
		Org2UserInfo := fabricSDK.InitInfo{
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
		balance, err := client1ServiceSetup.ClientAccountBalance()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.HTML(http.StatusOK, "transfer.html", gin.H{
			"token1": ethSDK.GetUserBalance("/home/eth-poa/signer1/data/geth.ipc", "0xD78d66C33933a05c57c503d61667918f95cee351", "0x416b1e5329Bd97BB704866bD489747b26848fA42"),
			"token2": balance,
		})
	})
	r.POST("/login", func(c *gin.Context) {
		username := c.DefaultPostForm("username", "somebody")
		password := c.DefaultPostForm("password", "***")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
	r.POST("/submit", func(c *gin.Context) {
		to := c.PostForm("to")
		value := c.PostForm("value")
		Org2User1Info := fabricSDK.InitInfo{
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
		channelClient1, err := channel.New(client1ChannelContext)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		eventClient1, err := event.New(client1ChannelContext, event.WithBlockEvents(), event.WithSeekType(seek.Newest))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client1ServiceSetup := fabricSDK.ServiceSetup{
			ChaincodeID:   Org2User1Info.ChaincodeID,
			ChannelClient: channelClient1,
			EventClient:   eventClient1,
		}

		tranMsg, err := client1ServiceSetup.Transfer(to, value)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("转账成功，交易编号为：" + tranMsg)
		}

		balance, err := client1ServiceSetup.ClientAccountBalance()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.HTML(http.StatusOK, "transfer.html", gin.H{
			"token1": ethSDK.GetUserBalance("/home/eth-poa/signer1/data/geth.ipc", "0xD78d66C33933a05c57c503d61667918f95cee351", "0x416b1e5329Bd97BB704866bD489747b26848fA42"),
			"token2": balance,
		})
	})
	r.Run(":9090")
}

package ethfabricListen

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	token_erc20 "sidechain/ethContract/openzeppelin-contracts/contracts/token/ERC20"
	"sidechain/fabricSDK"
	"sidechain/musig"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/events/deliverclient/seek"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type LogTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

func (lt *LogTransfer) string() string {
	return lt.From.Hex() + "," + lt.To.Hex() + "," + lt.Value.String()
}

// 监听事件日志
func Eth_listen_erc20_transfer(url string, address string, done func()) {
	defer done()
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	contractAbi, err := abi.JSON(strings.NewReader(string(token_erc20.TokenErc20MetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	sdk, err := fabricSDK.SetupSDK("./config/crypto-config.yaml", false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer sdk.Close()

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
			fmt.Printf("Log Index: %d\n", vLog.Index)

			switch vLog.Topics[0].Hex() {
			case logTransferSigHash.Hex():
				fmt.Printf("Log Name: Transfer\n")

				var transferEvent LogTransfer

				err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

				fmt.Printf("From: %s\n", transferEvent.From.Hex())
				fmt.Printf("To: %s\n", transferEvent.To.Hex())
				fmt.Printf("Tokens: %s\n", transferEvent.Value.String())

				fmt.Println()
				if transferEvent.To == common.HexToAddress("0xf745069D290dE951508CA088D198678758DcA46c") {
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
						UserName:        transferEvent.From.Hex(),
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

					sigandver(transferEvent.string())

					clientID, err := client1ServiceSetup.ClientAccountID()
					if err != nil {
						fmt.Println(err.Error())
					}

					mintMsg, err := minterServiceSetup.Mint(transferEvent.Value.String())
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println("铸币成功，交易编号为：" + mintMsg)
					}

					tranMsg, err := minterServiceSetup.Transfer(clientID, transferEvent.Value.String())
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println("转账成功，交易编号为：" + tranMsg)
					}

					accbal, err := client1ServiceSetup.ClientAccountBalance()
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println(accbal)
					}
				}
			}
		}
	}
}

func sigandver(tx string) {
	// Alice
	// private/public keys
	x1, X1 := musig.KeyGen()
	// random value
	r1, R1 := musig.KeyGen()

	// Bob
	// private/public keys
	x2, X2 := musig.KeyGen()
	// random value
	r2, R2 := musig.KeyGen()

	// Carol
	// private/public keys
	x3, X3 := musig.KeyGen()
	// random value
	r3, R3 := musig.KeyGen()

	// L is a multiset of public keys.
	var L []*btcec.PublicKey
	L = append(L, X1)
	L = append(L, X2)
	L = append(L, X3)
	fmt.Println("Public keys")
	for i, l := range L {
		fmt.Printf("L%d:%x\n", i, l.SerializeCompressed())
	}

	// R is a part of signature.
	R := musig.AddPubs(R1, R2, R3)

	// message
	m := []byte(tx)
	fmt.Printf("tx:%x\n", tx)

	//Signing
	// Alice signs.
	s1 := musig.Sign(L, R, m, x1, r1)
	// Bob signs.
	s2 := musig.Sign(L, R, m, x2, r2)
	// Bob signs.
	s3 := musig.Sign(L, R, m, x3, r3)
	// s is a part of signature.
	s := musig.AddSigs(s1, s2, s3)
	// signature
	fmt.Println("Signature")
	fmt.Printf("R:%x\n", R.SerializeCompressed())
	fmt.Printf("s:%v\n", s)

	// Verification
	v := musig.Ver(L, m, R, s)
	fmt.Printf("Result:%v\n", v)

	if v != 1 {
		fmt.Println("Fail")
	}
}

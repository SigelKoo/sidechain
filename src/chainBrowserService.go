package sidechain

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	providersFab "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
	"math/big"
	ethContract "sidechain/ethContract"
	"strings"
)

//区块链浏览器服务
var mainSDK *fabsdk.FabricSDK
var ledgerClient *ledger.Client

const (
	org1Name      = "Org1"
	org1Peer0 	  = "peer0.org1.example.com"
	org1AdminUser = "Admin"
	org1User      = "User1"
	channelID     = "mychannel"
	configPath = "/home/gopath/sidechain/config/crypto-config.yaml"
)
var chainBrowserConfigPath = configPath

//初始化区块浏览器SDK
func InitChainBrowserService() {
	log.Println("============ 初始化区块浏览器服务 ============")
	//获取fabsdk
	var err error
	ConfigBackend := config.FromFile(chainBrowserConfigPath)
	mainSDK, err = fabsdk.New(ConfigBackend)
	if err != nil {
		panic(fmt.Sprintf("Failed to create new SDK: %s", err))
	}
	//获取context
	org1AdminChannelContext := mainSDK.ChannelContext(channelID, fabsdk.WithUser(org1User), fabsdk.WithOrg(org1Name))
	// org1AdminChannelContext := mainSDK.ChannelContext(channelID, fabsdk.WithUser(org1AdminUser), fabsdk.WithOrg(org1Name))
	//Ledger client
	ledgerClient, err = ledger.New(org1AdminChannelContext)
	if err != nil {
		fmt.Printf("Failed to create new resource management client: %s", err)
	}
}

//查询账本信息
func QueryLedgerInfo() (*providersFab.BlockchainInfoResponse, error) {
	ledgerInfo, err := ledgerClient.QueryInfo()
	if err != nil {
		fmt.Printf("QueryInfo return error: %s", err)
		return nil, err
	}
	QueryPeerConfig(ledgerInfo)
	return ledgerInfo, nil
}

//查询节点信息
func QueryPeerConfig(ledgerInfo *providersFab.BlockchainInfoResponse) (*providersFab.EndpointConfig, error) {
	sdk := mainSDK
	configBackend, err := sdk.Config()
	if err != nil {
		fmt.Println("failed to get config backend, error: %s", err)
	}

	endpointConfig, err := fab.ConfigFromBackend(configBackend)
	if err != nil {
		fmt.Println("failed to get endpoint config, error: %s", err)
	}

	expectedPeerConfig1, _ := endpointConfig.PeerConfig("peer0.org1.example.com")
	fmt.Println("Unable to fetch Peer config for %s", "peer0.org1.example.com")
	expectedPeerConfig2, _ := endpointConfig.PeerConfig("peer1.org1.example.com")
	fmt.Println("Unable to fetch Peer config for %s", "peer1.org1.example.com")

	if !strings.Contains(ledgerInfo.Endorser, expectedPeerConfig1.URL) && !strings.Contains(ledgerInfo.Endorser, expectedPeerConfig2.URL) {
		fmt.Println("Expecting %s or %s, got %s", expectedPeerConfig1.URL, expectedPeerConfig2.URL, ledgerInfo.Endorser)
	}

	return &endpointConfig, nil
}

//查询最新10个区块信息
func QueryLatestBlocksInfo() ([]*Block, error) {
	ledgerInfo, err := ledgerClient.QueryInfo()
	if err != nil {
		fmt.Printf("QueryLatestBlocksInfo return error: %s\n", err)
		return nil, err
	}
	latestBlockList := []*Block{}
	lastetBlockNum := ledgerInfo.BCI.Height - 1
	minBlockNum := 1
	if lastetBlockNum > 10 {
		minBlockNum = int(lastetBlockNum - 10)
	}
	for i := lastetBlockNum; i > 0 && int(i) > minBlockNum; i-- {
		block,err := QueryBlockByBlockNumber(int64(i))
		if err != nil {
			fmt.Printf("QueryLatestBlocksInfo return error: %s", err)
			return latestBlockList, err
		}
		latestBlockList = append(latestBlockList, block)
	}
	return latestBlockList, nil
}

func QueryLatestBlocksInfoJsonStr() (string, error) {
	blockList, err := QueryLatestBlocksInfo()
	jsonStr, err := json.Marshal(blockList)
	return string(jsonStr), err
}

// 查询指定区块信息
func QueryBlockByBlockNumber(num int64) (*Block, error) {
	rawBlock,err := ledgerClient.QueryBlock(uint64(num))
	if err != nil {
		fmt.Printf("QueryBlock return error: %s", err)
		return nil, err
	}

	//解析区块体
	txList := []*Transaction{}
	for i :=range rawBlock.Data.Data {
		rawEnvelope, err := GetEnvelopeFromBlock(rawBlock.Data.Data[i])
		if err != nil {
			fmt.Printf("QueryBlock return error: %s", err)
			return nil, err
		}
		transaction, err := GetTransactionFromEnvelopeDeep(rawEnvelope)
		if err != nil {
			fmt.Printf("QueryBlock return error: %s", err)
			return nil, err
		}
		for i := range transaction.TransactionActionList {
			transaction.TransactionActionList[i].BlockNum = rawBlock.Header.Number
		}
		txList = append(txList, transaction)
	}

	block := Block {
		Number:       rawBlock.Header.Number,
		PreviousHash: rawBlock.Header.PreviousHash,
		DataHash:     rawBlock.Header.DataHash,
		BlockHash:    rawBlock.Header.DataHash, //需要计算
		TxNum:        len(rawBlock.Data.Data),
		TransactionList:     txList,
		CreateTime:   txList[0].TransactionActionList[0].Timestamp,
	}

	return &block, nil
}

//查询交易信息
func QueryTransactionByTxId(txId string) (*Transaction, error) {
	rawTx,err := ledgerClient.QueryTransaction(providersFab.TransactionID(txId))
	if err != nil {
		fmt.Printf("QueryBlock return error: %s", err)
		return nil, err
	}

	transaction, err := GetTransactionFromEnvelopeDeep(rawTx.TransactionEnvelope)
	if err != nil {
		fmt.Printf("QueryBlock return error: %s", err)
		return nil, err
	}
	block,err := ledgerClient.QueryBlockByTxID(providersFab.TransactionID(txId))
	if err != nil {
		fmt.Printf("QueryBlock return error: %s", err)
		return nil, err
	}
	for i := range transaction.TransactionActionList {
		transaction.TransactionActionList[i].BlockNum = block.Header.Number
	}
	return transaction, nil
}

func QueryTransactionByTxIdJsonStr(txId string) (string, error) {
	transaction,err := QueryTransactionByTxId(txId)
	if err!=nil {
		return "", err
	}
	jsonStr,err := json.Marshal(transaction)
	return string(jsonStr), err
}

// 获取当前最新区块号并记录
func latestBlockNum() (int, error) {
	ledgerInfo, err := ledgerClient.QueryInfo()
	if err != nil {
		fmt.Printf("QueryLatestBlocksInfo return error: %s\n", err)
		return -1, err
	}
	lastetBlockNum := ledgerInfo.BCI.Height - 1
	return int(lastetBlockNum), err
}

// 判断是否出了8个区块
func QueryIsEightBlock(rawBlockNum int) int {
	ledgerInfo, err := ledgerClient.QueryInfo()
	if err != nil {
		fmt.Printf("QueryIsEightBlock return error: %s\n", err)
	}
	latestBlockNum := ledgerInfo.BCI.Height - 1
	flag := int(latestBlockNum) - rawBlockNum
	return flag
}

// 将区块哈希取出，发送给以太坊客户端
func QueryBlockHash(rawBlockNum *int) string {
	for {
		temp := QueryIsEightBlock(*rawBlockNum)
		if temp < 7 {
			continue
		} else {
			arr := [][]byte{}
			for i := 0; i < 8; i++ {
				blockInfo, err := QueryBlockByBlockNumber(int64(*rawBlockNum + i))
				if err != nil {
					fmt.Printf("QueryBlockHash return error: %s\n", err)
				}
				arr = append(arr, blockInfo.BlockHash)
				*rawBlockNum += 8
			}
			ethContract.ContractWrite("HTTP://192.168.132.80:8545", "2e8749fd1ba7a42586d2bb38c10fab2e8845abd7733378a95a03fdcdbd1b854e", "", new(big.Int).SetUint64(uint64(*rawBlockNum - 7)), string(ethHash(arr)))
		}
	}
}

// 模拟merkle树根
func ethHash(arr [][]byte) []byte {
	return crypto.Keccak256(
		crypto.Keccak256(
			crypto.Keccak256(
				crypto.Keccak256(arr[0]),
				crypto.Keccak256(arr[1]),
			),
			crypto.Keccak256(
				crypto.Keccak256(arr[2]),
				crypto.Keccak256(arr[3]),
			),
		),
		crypto.Keccak256(
			crypto.Keccak256(
				crypto.Keccak256(arr[4]),
				crypto.Keccak256(arr[5]),
			),
			crypto.Keccak256(
				crypto.Keccak256(arr[6]),
				crypto.Keccak256(arr[7]),
			),
		),
	)
}
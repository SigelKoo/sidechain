package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func main() {
	log.Println("============ application-golang starts ============")

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	userName := "0x416b1e5329Bd97BB704866bD489747b26848fA42"

	if !wallet.Exists(userName) {
		err = populateWallet(userName, wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"..",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org2.example.com",
		"connection-org2.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, userName),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	contract := network.GetContract("token_erc20")

	log.Println("--> Submit Transaction: Transfer, transfers amount of token_erc20 to recipient")
	result, err := contract.SubmitTransaction("Transfer", "eDUwOTo6Q049MHg0MTZiMWU1MzI5QmQ5N0JCNzA0ODY2YkQ0ODk3NDdiMjY4NDhmQTQyLE9VPWNsaWVudCxPPUh5cGVybGVkZ2VyLFNUPU5vcnRoIENhcm9saW5hLEM9VVM6OkNOPWNhLm9yZzIuZXhhbXBsZS5jb20sTz1vcmcyLmV4YW1wbGUuY29tLEw9SHVyc2xleSxTVD1IYW1wc2hpcmUsQz1VSw==", "1")
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}
	log.Println(string(result))

	//log.Println("--> Evaluate Transaction: ReadAsset, function returns an asset with a given assetID")
	//result, err = contract.EvaluateTransaction("ReadAsset", "asset13")
	//if err != nil {
	//	log.Fatalf("Failed to evaluate transaction: %v\n", err)
	//}
	//log.Println(string(result))
	//
	//log.Println("--> Evaluate Transaction: AssetExists, function returns 'true' if an asset with given assetID exist")
	//result, err = contract.EvaluateTransaction("AssetExists", "asset1")
	//if err != nil {
	//	log.Fatalf("Failed to evaluate transaction: %v\n", err)
	//}
	//log.Println(string(result))
	//
	//log.Println("--> Submit Transaction: TransferAsset asset1, transfer to new owner of Tom")
	//_, err = contract.SubmitTransaction("TransferAsset", "asset1", "Tom")
	//if err != nil {
	//	log.Fatalf("Failed to Submit transaction: %v", err)
	//}
	//
	//log.Println("--> Evaluate Transaction: ReadAsset, function returns 'asset1' attributes")
	//result, err = contract.EvaluateTransaction("ReadAsset", "asset1")
	//if err != nil {
	//	log.Fatalf("Failed to evaluate transaction: %v", err)
	//}
	//log.Println(string(result))
	log.Println("============ application-golang ends ============")
}

func populateWallet(userName string, wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")
	credPath := filepath.Join(
		"..",
		"..",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org2.example.com",
		"users",
		userName + "@org2.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org2MSP", string(cert), string(key))

	return wallet.Put(userName, identity)
}
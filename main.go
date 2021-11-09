package main

import (
	// ethsdk "sidechain/ethSDK"
	sc "sidechain/src"
)

func main() {
	// fmt.Println(ethsdk.ContractDeploy("HTTP://192.168.132.80:8501", "2e8749fd1ba7a42586d2bb38c10fab2e8845abd7733378a95a03fdcdbd1b854e"))
	sc.RegisterUser("123", "Org1", "123")
}

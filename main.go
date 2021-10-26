package main

import (
	"fmt"
	"sidechain/ethContract"
)

func main() {
	fmt.Println(ethContract.ContractDeploy("HTTP://192.168.132.80:8501", "2e8749fd1ba7a42586d2bb38c10fab2e8845abd7733378a95a03fdcdbd1b854e"))
}

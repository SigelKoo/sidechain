package main

import sidechain "sidechain/src"

func main() {
	sidechain.InitChainBrowserService()
	sidechain.QueryLatestBlocksInfo()
}

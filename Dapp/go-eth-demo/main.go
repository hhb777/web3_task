package main

import "github.com/local/go-eth-demo/goethereumbase"

func main() {
	var Pkey string = "xxx"

	//goethereumbase.SelectBlock(Pkey)
	//goethereumbase.SelectTrans(Pkey)
	//goethereumbase.SelectRecv(Pkey)
	//goethereumbase.CreateWallet()
	//goethereumbase.ETHTrans(Pkey)
	//goethereumbase.Erc20Trans(Pkey)
	//goethereumbase.SelectBalance(Pkey)
	//goethereumbase.SelectERC20Balance(Pkey)
	//goethereumbase.SubscribeBlock(Pkey)
	//goethereumbase.DeployContract(Pkey)
	goethereumbase.LoadContract(Pkey)
}

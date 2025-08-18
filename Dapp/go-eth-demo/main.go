package main

import "github.com/local/go-eth-demo/goethereumbase"

func main() {
	var Pkey string = "xxx"

	//goethereumbase.SelectBlock(Pkey)
	//goethereumbase.SelectTrans(Pkey)
	//goethereumbase.SelectRecv(Pkey)
	//goethereumbase.CreateWallet()
	goethereumbase.ETHTrans(Pkey)
}

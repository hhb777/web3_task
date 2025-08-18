package goethereumbase

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	token "github.com/local/go-eth-demo/erc20"
	"log"
	"math"
	"math/big"
)

func SelectERC20Balance(Pkey string) {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/" + Pkey)
	if err != nil {
		log.Fatal(err)
	}
	tokenAddress := common.HexToAddress("0x77Eb631EeC72A666BD9a7B0Ad3c99a715Fa42eE9")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0x18540EF36c5aD9Ab916d378e31Db7d3a836A0A44")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("wei:", bal, "name:", name, "symbol:", symbol, "decimals:", decimals)
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(18)))
	fmt.Println("balance:", value)
}

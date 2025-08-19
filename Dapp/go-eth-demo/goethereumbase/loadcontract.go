package goethereumbase

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/local/go-eth-demo/store"
	"log"
)

const (
	contractAddr = "0x568ce5829a2475c5b4a22B853c567166717758dE"
)

func LoadContract(Pkey string) {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/" + Pkey)
	if err != nil {
		log.Fatal(err)
	}

	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}
	_ = storeContract

}

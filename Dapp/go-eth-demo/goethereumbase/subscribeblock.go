package goethereumbase

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func SubscribeBlock(Pkey string) {
	//client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/" + Pkey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//headers := make(chan *types.Header)
	//sub, err := client.SubscribeNewHead(context.Background(), headers)
	//if err != nil {
	//
	//	log.Fatal(err)
	//}
	//for {
	//	select {
	//	case err := <-sub.Err():
	//		log.Fatal(err)
	//	case header := <-headers:
	//		fmt.Println(header.Hash().Hex())
	//		block, err := client.BlockByHash(context.Background(), header.Hash())
	//		if err != nil {
	//			fmt.Println("ERROR:", err.Error())
	//			log.Fatal(err)
	//		}
	//		fmt.Println(header.Hash().Hex())
	//		fmt.Println(header.Number.Uint64())
	//		fmt.Println(header.Time)
	//		fmt.Println(header.Nonce.Uint64())
	//
	//		fmt.Println(block.Hash().Hex())
	//		fmt.Println(block.Number().Uint64())
	//		fmt.Println(block.Time())
	//		fmt.Println(block.Nonce())
	//		fmt.Println(len(block.Transactions()))
	//
	//	}
	//}
	client, err := ethclient.Dial("wss://go.getblock.io/xxx")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}

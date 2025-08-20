package goethereumbase

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
)

var StoreABI = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

func EventContract(Pkey string) {
	//client, err := ethclient.Dial("https://sepolia.infura.io/v3/" + Pkey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//contractAddress := common.HexToAddress("0x171bCD02237133E7D215A7177A5CD08557161615")
	//query := ethereum.FilterQuery{
	//	FromBlock: big.NewInt(9024199),
	//	// ToBlock:   big.NewInt(2394201),
	//	Addresses: []common.Address{
	//		contractAddress,
	//	},
	//	// Topics: [][]common.Hash{
	//	//  {},
	//	//  {},
	//	// },
	//}
	//
	//logs, err := client.FilterLogs(context.Background(), query)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//contractAbi, err := abi.JSON(strings.NewReader(StoreABI))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, vLog := range logs {
	//	fmt.Println(vLog.BlockHash.Hex())
	//	fmt.Println(vLog.BlockNumber)
	//	fmt.Println(vLog.TxHash.Hex())
	//	event := struct {
	//		Key   [32]byte
	//		Value [32]byte
	//	}{}
	//	err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(common.Bytes2Hex(event.Key[:]))
	//	fmt.Println(common.Bytes2Hex(event.Value[:]))
	//	var topics []string
	//	for i := range vLog.Topics {
	//		topics = append(topics, vLog.Topics[i].Hex())
	//	}
	//
	//	fmt.Println("topics[0]=", topics[0])
	//	if len(topics) > 1 {
	//		fmt.Println("indexed topics:", topics[1:])
	//	}
	//}
	//
	//eventSignature := []byte("ItemSet(bytes32,bytes32)")
	//hash := crypto.Keccak256Hash(eventSignature)
	//fmt.Println("signature topics=", hash.Hex())

	// =====================================

	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/" + Pkey)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x171bCD02237133E7D215A7177A5CD08557161615")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(StoreABI)))
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)
			fmt.Println(vLog.TxHash.Hex())
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}
			err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(common.Bytes2Hex(event.Key[:]))
			fmt.Println(common.Bytes2Hex(event.Value[:]))
			var topics []string
			for i := range vLog.Topics {
				topics = append(topics, vLog.Topics[i].Hex())
			}
			fmt.Println("topics[0]=", topics[0])
			if len(topics) > 1 {
				fmt.Println("index topic:", topics[1:])
			}
		}
	}
}

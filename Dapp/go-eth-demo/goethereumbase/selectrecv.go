package goethereumbase

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
	"time"
)

func SelectRecv(Pkey string) {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/" + Pkey)
	if err != nil {
		log.Fatal(err)
	}
	var receiptByHash []*types.Receipt
	ch := make(chan int)

	go func(ch chan int) {
		blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
		receiptByHash, err = client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
		ch <- 1
		if err != nil {
			log.Fatal(err)
		}
		close(ch)

	}(ch)
	for i := 0; i < 10; i++ {
		select {
		case <-ch:
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	blockNumber := big.NewInt(5671744)
	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receiptsByNum)
	fmt.Println(receiptByHash)
	fmt.Println("receiptByHash[0] == receiptsByNum[0]:", receiptByHash[0] == receiptsByNum[0])

	for _, receipt := range receiptsByNum {
		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
		fmt.Println(receipt.TxHash.Hex())
		fmt.Println(receipt.TransactionIndex)
		fmt.Println(receipt.ContractAddress.Hex())
		break
	}
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.Status)
	fmt.Println(receipt.Logs)
	fmt.Println(receipt.TxHash.Hex())
	fmt.Println(receipt.TransactionIndex)
	fmt.Println(receipt.ContractAddress.Hex())

}

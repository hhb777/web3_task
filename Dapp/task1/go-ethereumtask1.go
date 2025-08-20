package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"task1/countertest"
	"time"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/xxx")
	if err != nil {
		fmt.Println(err)
	}
	////区块号信息
	//blockNumber := big.NewInt(int64(9024199))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//block, err := client.BlockByNumber(context.Background(), blockNumber)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("block hash:", block.Hash()) // 0xcb3cb8469acbab1c8dc39ba4039586d875bb4310321e40e64dc90ff5259e658b
	//fmt.Println("block time:", block.Time()) //1755678564
	//count, err := client.TransactionCount(context.Background(), block.Hash())
	//fmt.Println("block trans count:", count) // 182

	//发送交易
	//privateKey, err := crypto.HexToECDSA("xxx")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//publickey := privateKey.Public()
	//publicECDSA, ok := publickey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal(err)
	//}
	//fromAddress := crypto.PubkeyToAddress(*publicECDSA)
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//gasLimit := uint64(23000)
	//toAddress := common.HexToAddress("0x87410827BCE8E8BF9F98c587f4Fd482dDcf5E804")
	//
	//value := big.NewInt(1000000000000000)
	//var data []byte
	//tx := types.NewTx(&types.LegacyTx{Nonce: nonce, Value: value, GasPrice: gasPrice, Gas: gasLimit, Data: data, To: &toAddress})
	//chainID, err := client.NetworkID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = client.SendTransaction(context.Background(), signedTx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

	// deploy contract
	privateKey, err := crypto.HexToECDSA("xxx")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasLimit := uint64(3000000)
	chainID, err := client.ChainID(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(gasLimit)
	auth.GasPrice = gasPrice
	address, tx, instance, err := countertest.DeployCountertest(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	_ = instance

	_, err = waitcontract(tx.Hash(), client)
	if err != nil {
		log.Fatal(err)
	}
	// load contract
	//address := common.HexToAddress("0x8b0AA10ef404ea10A5CcC337F294DB38dBa9A1c6")
	countercontract, err := countertest.NewCountertest(address, client)
	if err != nil {
		log.Fatal(err)
	}
	// call contract
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	tx1, err := countercontract.CountAdd(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx1.Hash().Hex())
	_, err = waitcontract(tx1.Hash(), client)
	if err != nil {
		log.Fatal(err)
	}
	callopt := &bind.CallOpts{Context: context.Background()}
	value, err := countercontract.Count(callopt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count value:", value)

}

func waitcontract(txHash common.Hash, client *ethclient.Client) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, err
		}
		// 等待一段时间后再次查询
		time.Sleep(1 * time.Second)
	}
}

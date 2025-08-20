package goethereumbase

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"time"

	//"crypto/ecdsa"
	"fmt"
	_ "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/local/go-eth-demo/store"
	"log"
	"math/big"
)

const (
	contractAddre = "0x568ce5829a2475c5b4a22B853c567166717758dE"
)

func ExcuteContract(Pkey string) {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/" + Pkey)
	if err != nil {
		log.Fatal(err)
	}
	//abi生成go代码的方式执行合约
	//storeContract, err := store.NewStore(common.HexToAddress(contractAddre), client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//privateKey, err := crypto.HexToECDSA("xxx")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var key [32]byte
	//var value [32]byte
	//
	//copy(key[:], []byte("demo_save_key"))
	//copy(value[:], []byte("demo_save_valuetest"))
	//
	//opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//tx, err := storeContract.SetItem(opt, key, value)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("tx hash: %x\n", tx.Hash().Hex())
	//
	//callOpt := &bind.CallOpts{Context: context.Background()}
	//valueInContract, err := storeContract.Items(callOpt, key)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("valueInContract:", valueInContract)
	//fmt.Println("value: ", value)
	//fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)

	//==========================================

	//ethclient的abi方式执行合约
	//privateKey, err := crypto.HexToECDSA("xxx")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//publickey := privateKey.Public()
	//publicKeyECDSA, ok := publickey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//contractABIe, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//methodName := "setItem"
	//var key [32]byte
	//var value [32]byte
	//
	//copy(key[:], []byte("demo_save_key"))
	//copy(value[:], []byte("demo_save_valuetest"))
	//
	//input, err := contractABIe.Pack(methodName, key, value)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//toaddress := common.HexToAddress(contractAddre)
	//chainID := big.NewInt(int64(11155111))
	//tx := types.NewTx(&types.LegacyTx{Nonce: nonce, To: &toaddress, Value: big.NewInt(0), Gas: 300000, GasPrice: gasPrice, Data: input})
	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = client.SendTransaction(context.Background(), signedTx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Transaction sent: ", signedTx.Hash().Hex())
	//_, err = waitForReceipte(client, signedTx.Hash())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//callInput, err := contractABIe.Pack("items", key)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//callMsg := ethereum.CallMsg{
	//	To:   &toaddress,
	//	Data: callInput,
	//}
	//result, err := client.CallContract(context.Background(), callMsg, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var unpacked [32]byte
	//err = contractABIe.UnpackIntoInterface(&unpacked, "items", result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("unpacked:", unpacked)
	//fmt.Println("value: ", value)
	//fmt.Println("is value saving in contract equals to origin value:", unpacked == value)

	//===============================================

	//ethclient不使用abi方式执行合约
	privateKey, err := crypto.HexToECDSA("xxx")
	if err != nil {
		log.Fatal(err)
	}
	publickey := privateKey.Public()
	publicKeyECDSA, ok := publickey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
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
	methodSignature := []byte("setItem(bytes32,bytes32)")
	methodSelector := crypto.Keccak256(methodSignature)[:4]

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("demo_save_key_no_use_abi"))
	copy(value[:], []byte("demo_save_value_no_use_abi_11111"))

	// 组合调用数据
	var input []byte
	input = append(input, methodSelector...)
	input = append(input, key[:]...)
	input = append(input, value[:]...)
	toaddress := common.HexToAddress(contractAddre)
	chainID := big.NewInt(int64(11155111))
	tx := types.NewTx(&types.LegacyTx{Nonce: nonce, To: &toaddress, Value: big.NewInt(0), Gas: 300000, GasPrice: gasPrice, Data: input})
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	_, err = waitForReceipte(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	itemsSignature := []byte("items(bytes32)")
	itemsSelector := crypto.Keccak256(itemsSignature)[:4]

	var callInput []byte
	callInput = append(callInput, itemsSelector...)
	callInput = append(callInput, key[:]...)
	callMsg := ethereum.CallMsg{
		To:   &toaddress,
		Data: callInput,
	}
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result: ", result)
	var unpacked [32]byte
	copy(unpacked[:], result)

	fmt.Println("unpacked3:", unpacked)
	fmt.Println("value: ", value)
	fmt.Println("is value saving in contract equals to origin value:", unpacked == value)
}

func waitForReceipte(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != ethereum.NotFound {
			return nil, err
		} else {
			return receipt, nil
		}
		time.Sleep(1 * time.Second)

	}
}

package goethereumbase

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	_ "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/local/go-eth-demo/store"
	"log"
	"math/big"
	_ "time"
)

const (
	contractBytecode = "608060405234801561000f575f80fd5b5060405161084d38038061084d83398181016040528101906100319190610193565b805f908161003f91906103e7565b50506104b6565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6100a58261005f565b810181811067ffffffffffffffff821117156100c4576100c361006f565b5b80604052505050565b5f6100d6610046565b90506100e2828261009c565b919050565b5f67ffffffffffffffff8211156101015761010061006f565b5b61010a8261005f565b9050602081019050919050565b8281835e5f83830152505050565b5f610137610132846100e7565b6100cd565b9050828152602081018484840111156101535761015261005b565b5b61015e848285610117565b509392505050565b5f82601f83011261017a57610179610057565b5b815161018a848260208601610125565b91505092915050565b5f602082840312156101a8576101a761004f565b5b5f82015167ffffffffffffffff8111156101c5576101c4610053565b5b6101d184828501610166565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061022857607f821691505b60208210810361023b5761023a6101e4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261029d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610262565b6102a78683610262565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102eb6102e66102e1846102bf565b6102c8565b6102bf565b9050919050565b5f819050919050565b610304836102d1565b610318610310826102f2565b84845461026e565b825550505050565b5f90565b61032c610320565b6103378184846102fb565b505050565b5b8181101561035a5761034f5f82610324565b60018101905061033d565b5050565b601f82111561039f5761037081610241565b61037984610253565b81016020851015610388578190505b61039c61039485610253565b83018261033c565b50505b505050565b5f82821c905092915050565b5f6103bf5f19846008026103a4565b1980831691505092915050565b5f6103d783836103b0565b9150826002028217905092915050565b6103f0826101da565b67ffffffffffffffff8111156104095761040861006f565b5b6104138254610211565b61041e82828561035e565b5f60209050601f83116001811461044f575f841561043d578287015190505b61044785826103cc565b8655506104ae565b601f19841661045d86610241565b5f5b828110156104845784890151825560018201915060208501945060208101905061045f565b868310156104a1578489015161049d601f8916826103b0565b8355505b6001600288020188555050505b505050505050565b61038a806104c35f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c806348f343f31461004357806354fd4d5014610073578063f56256c714610091575b5f80fd5b61005d600480360381019061005891906101d6565b6100ad565b60405161006a9190610210565b60405180910390f35b61007b6100c2565b6040516100889190610299565b60405180910390f35b6100ab60048036038101906100a691906102b9565b61014d565b005b6001602052805f5260405f205f915090505481565b5f80546100ce90610324565b80601f01602080910402602001604051908101604052809291908181526020018280546100fa90610324565b80156101455780601f1061011c57610100808354040283529160200191610145565b820191905f5260205f20905b81548152906001019060200180831161012857829003601f168201915b505050505081565b8060015f8481526020019081526020015f2081905550817fe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4826040516101939190610210565b60405180910390a25050565b5f80fd5b5f819050919050565b6101b5816101a3565b81146101bf575f80fd5b50565b5f813590506101d0816101ac565b92915050565b5f602082840312156101eb576101ea61019f565b5b5f6101f8848285016101c2565b91505092915050565b61020a816101a3565b82525050565b5f6020820190506102235f830184610201565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61026b82610229565b6102758185610233565b9350610285818560208601610243565b61028e81610251565b840191505092915050565b5f6020820190508181035f8301526102b18184610261565b905092915050565b5f80604083850312156102cf576102ce61019f565b5b5f6102dc858286016101c2565b92505060206102ed858286016101c2565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061033b57607f821691505b60208210810361034e5761034d6102f7565b5b5091905056fea26469706673582212209281a0f4aaf3481ca222acf17d6de8e763c53b9c686e08808dc2a62f8cc552fe64736f6c634300081a0033"
	contractABI      = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
)

func DeployContract(Pkey string) {
	////abigen方式部署合约
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/" + Pkey)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("xxx")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
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
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	_ = instance

	//==========================================

	//ethclient 工具部署合约
	//client, err := ethclient.Dial("https://sepolia.infura.io/v3/" + Pkey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//privateKey, err := crypto.HexToECDSA("xxx")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("error casting public key to ECDSA")
	//}
	//
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//
	//// 获取nonce
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 获取建议的gas价格
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//input := "1.0"
	//parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//constructorparam, err := parsedABI.Pack("", input)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//data := append(common.FromHex(contractBytecode), constructorparam...)
	//
	//tx := types.NewTx(&types.LegacyTx{Nonce: nonce, GasPrice: big.NewInt(0).Add(gasPrice, big.NewInt(int64(10000000000))), Data: data, Gas: uint64(300000), Value: big.NewInt(0)})
	//
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
	//receipt, err := waitForReceipt(client, signedTx.Hash())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Contract deployed at: %s\n", receipt.ContractAddress.Hex())
}

//func waitForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
//	for {
//		receipt, err := client.TransactionReceipt(context.Background(), txHash)
//		if err == nil {
//			return receipt, nil
//		}
//		if err != ethereum.NotFound {
//			return nil, err
//		}
//		// 等待一段时间后再次查询
//		time.Sleep(1 * time.Second)
//	}
//}

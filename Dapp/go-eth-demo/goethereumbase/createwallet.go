package goethereumbase

import (
	"crypto/ecdsa"
	_ "encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

func CreateWallet() {
	//privateKey, err := crypto.GenerateKey()
	//if err != nil {
	//	log.Fatal(err)
	//}

	privateKey, err := crypto.HexToECDSA("ed67084e2f5717fa9b98250f40b93c3d998115168f782ad676a71e30f2ee974c")
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("from publickey:", hexutil.Encode(publicKeyBytes))
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address:", address)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}

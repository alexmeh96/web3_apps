package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKey := crypto.FromECDSA(pvk)
	fmt.Println(hexutil.Encode(privateKey))

	publicKey := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(publicKey))

	fmt.Println(crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}

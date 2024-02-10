package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
)

func main() {
	password := "password"

	// создание стора для хранения аккаунтов(кошельков)
	keyStore := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	// создание аккаунта(кошелька) и сохранение в дирректорию ./keystore файл,
	// который содержит закрытый ключ, зашированный паролем
	a, err := keyStore.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Address", a.Address)
	fmt.Println("File", a.URL.Path)
	fmt.Println()

	b, err := ioutil.ReadFile(a.URL.Path)
	if err != nil {
		log.Fatal(err)
	}

	// получение и расшифровка паролем информации об аккаунте(кошельке) из файла
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Private", hexutil.Encode(pData))

	pData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Public", hexutil.Encode(pData))

	fmt.Println("Address", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())
}

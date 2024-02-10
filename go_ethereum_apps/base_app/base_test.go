package base_app

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"testing"
)

const sepoliaUrl = ""
const mainnetUrl = ""
const localUrl = "http://127.0.0.1:8545"

func TestBalance(t *testing.T) {
	client, err := ethclient.DialContext(context.Background(), mainnetUrl)
	if err != nil {
		log.Fatalf("Error to create a ether client:%v", err)
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get block:%v", err)
	}

	fmt.Println("The block last number: ", block.Number())

	addr := "0x4299885E97668415CD67484d4a2c5214480Ff76d"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance:%v", err)
	}
	fmt.Println("The balance:", balance)
	// 1 ether = 10^18 wei
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fmt.Println(fBalance)
	balanceEther := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(balanceEther)
}

func TestWallet(t *testing.T) {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKey := crypto.FromECDSA(pvk)
	fmt.Println("privateKey:", hexutil.Encode(privateKey))

	publicKey := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println("publicKey:", hexutil.Encode(publicKey))

	fmt.Println("address:", crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}

func TestKeystore(t *testing.T) {
	password := "password"

	// создание стора для хранения аккаунтов(кошельков)
	keyStore := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)

	// создание аккаунта(кошелька) и сохранение в дирректорию ./wallets файл,
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

func TestBalanceDifferentNetworks(t *testing.T) {
	client1, err := ethclient.Dial(mainnetUrl)
	if err != nil {
		log.Fatal(err)
	}
	client2, err := ethclient.Dial(sepoliaUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client1.Close()
	defer client2.Close()

	address := common.HexToAddress("cd00ac6c1d67ebd195bc296b65a29e8196df2dab")

	b1, err := client1.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := client2.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mainnet balance: ", b1)
	fmt.Println("sepolia balance: ", b2)
}

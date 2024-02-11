package base_app

import (
	todo "base_app/gen"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"testing"
)

const mainnetUrl = ""
const sepoliaUrl = ""
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

func TestMakeKeystore(t *testing.T) {
	password := "password"

	// создание стора для хранения аккаунтов(кошельков)
	keyStore := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	// создание аккаунта(кошелька) и сохранение в дирректорию ./wallet файл,
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

func TestMakeAndSendTransaction(t *testing.T) {
	password := "password"
	client, err := ethclient.Dial(sepoliaUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	address1 := common.HexToAddress("236a0e676bf8ae80fdf731de7cbec06c27b1cd5c")
	address2 := common.HexToAddress("4d2a47eeab5caf04ff4bf5d3a13bf82f72f69595")

	// получение порядкового номера нашего кошелька из сети
	nonce, err := client.PendingNonceAt(context.Background(), address1)
	if err != nil {
		log.Fatal(err)
	}

	// получаем из сети текущую стоимость газа
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// колличество wei, которые мы переводим
	// 1 ether = 1000000000000000000 wei
	amount := big.NewInt(100000000000000000)

	// создаём транзакцию (комиссия = 21000 * gasPrice)
	// вся сумма которая будет списана с нашего кошелька равна: amount + 21000 * gasPrice
	tx := types.NewTransaction(nonce, address2, amount, 21000, gasPrice, nil)

	// получаем идентификатор сети, в которой мы выполняем транзакцию
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// файл-кошелька
	b, err := ioutil.ReadFile("wallet/UTC--2024-02-10T19-50-15.564706810Z--236a0e676bf8ae80fdf731de7cbec06c27b1cd5c")
	if err != nil {
		log.Fatal(err)
	}

	// получаем и расшифровываем ключ из файл-кошелька
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	// подписываем транзакцию приватным ключём
	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	// отправляем транзакцию в сеть на выполнение
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	// hash отправленной транзакции
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func TestMakeAndDeployContract(t *testing.T) {
	password := "password"

	// файл-кошелька
	b, err := ioutil.ReadFile("wallet/UTC--2024-02-10T19-50-15.564706810Z--236a0e676bf8ae80fdf731de7cbec06c27b1cd5c")
	if err != nil {
		log.Fatal(err)
	}

	// получаем и расшифровываем ключ из файл-кошелька
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(sepoliaUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	// получение порядкового номера нашего кошелька из сети
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}

	// получаем из сети текущую стоимость газа
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// получаем идентификатор сети, в которой мы выполняем транзакцию
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	txOpts.GasPrice = gasPrice
	txOpts.GasLimit = uint64(3000000)
	txOpts.Nonce = big.NewInt(int64(nonce))

	// отправка контракта в сеть
	addr, tx, _, err := todo.DeployTodo(txOpts, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(addr.Hex())
	fmt.Println(tx.Hash().Hex())
}

package base_app

import (
	todo "base_app/gen"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"testing"
)

const contractAddress = ""

func TestRunAddFunc(t *testing.T) {
	password := "password"

	b, err := ioutil.ReadFile("wallet/UTC--2024-02-10T19-50-15.564706810Z--236a0e676bf8ae80fdf731de7cbec06c27b1cd5c")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(sepoliaUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	contractAddr := common.HexToAddress(contractAddress)
	td, err := todo.NewTodo(contractAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	txOpt, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	txOpt.GasLimit = 3000000
	txOpt.GasPrice = gasPrice

	// вызов функции Add у контракта
	tx, err := td.Add(txOpt, "First Task")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash())
}

func TestRunListFunc(t *testing.T) {
	password := "password"

	b, err := ioutil.ReadFile("wallet/UTC--2024-02-10T19-50-15.564706810Z--236a0e676bf8ae80fdf731de7cbec06c27b1cd5c")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(sepoliaUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	cAdd := common.HexToAddress(contractAddress)
	td, err := todo.NewTodo(cAdd, client)
	if err != nil {
		log.Fatal(err)
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	txOpts.GasLimit = 3000000
	txOpts.GasPrice = gasPrice

	addr := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	// вызвов функции List у контракта
	tasks, err := td.List(&bind.CallOpts{
		From: addr,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tasks)
}

func TestRunUpdateFunc(t *testing.T) {
	password := "password"

	b, err := ioutil.ReadFile("wallet/UTC--2024-02-10T19-50-15.564706810Z--236a0e676bf8ae80fdf731de7cbec06c27b1cd5c")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(sepoliaUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	cAdd := common.HexToAddress(contractAddress)
	td, err := todo.NewTodo(cAdd, client)
	if err != nil {
		log.Fatal(err)
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	txOpts.GasLimit = 3000000
	txOpts.GasPrice = gasPrice

	tx, err := td.Update(txOpts, big.NewInt(0), "update task content")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Toggle tx", tx.Hash())
}

func TestRunToggleFunc(t *testing.T) {
	password := "password"

	b, err := ioutil.ReadFile("wallet/UTC--2024-02-10T19-50-15.564706810Z--236a0e676bf8ae80fdf731de7cbec06c27b1cd5c")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(sepoliaUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	cAdd := common.HexToAddress(contractAddress)
	td, err := todo.NewTodo(cAdd, client)
	if err != nil {
		log.Fatal(err)
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	txOpts.GasLimit = 3000000
	txOpts.GasPrice = gasPrice

	tx, err := td.Toggle(txOpts, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Toggle tx", tx.Hash())
}

func TestRunRemoveFunc(t *testing.T) {
	password := "password"

	b, err := ioutil.ReadFile("wallet/UTC--2024-02-10T19-50-15.564706810Z--236a0e676bf8ae80fdf731de7cbec06c27b1cd5c")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(sepoliaUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	cAdd := common.HexToAddress(contractAddress)
	td, err := todo.NewTodo(cAdd, client)
	if err != nil {
		log.Fatal(err)
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	txOpts.GasLimit = 3000000
	txOpts.GasPrice = gasPrice

	tx, err := td.Remove(txOpts, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Toggle tx", tx.Hash())
}

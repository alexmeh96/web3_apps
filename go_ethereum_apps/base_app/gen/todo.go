// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	Content string
	Status  bool
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"toggle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506111868061005c5f395ff3fe608060405234801561000f575f80fd5b506004361061007b575f3560e01c80638da5cb5b116100595780638da5cb5b146100d55780639507d39a146100f3578063b0c8f9dc14610123578063f745630f1461013f5761007b565b80630f560cd71461007f5780634cc822151461009d578063751ef753146100b9575b5f80fd5b61008761015b565b6040516100949190610992565b60405180910390f35b6100b760048036038101906100b291906109f6565b6102b6565b005b6100d360048036038101906100ce91906109f6565b610410565b005b6100dd6104d3565b6040516100ea9190610a60565b60405180910390f35b61010d600480360381019061010891906109f6565b6104f6565b60405161011a9190610ab3565b60405180910390f35b61013d60048036038101906101389190610bff565b61062f565b005b61015960048036038101906101549190610c46565b6106fc565b005b60603373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101b3575f80fd5b6001805480602002602001604051908101604052809291908181526020015f905b828210156102ad578382905f5260205f2090600202016040518060400160405290815f8201805461020490610ccd565b80601f016020809104026020016040519081016040528092919081815260200182805461023090610ccd565b801561027b5780601f106102525761010080835404028352916020019161027b565b820191905f5260205f20905b81548152906001019060200180831161025e57829003601f168201915b50505050508152602001600182015f9054906101000a900460ff161515151581525050815260200190600101906101d4565b50505050905090565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461030c575f80fd5b5f8190505b600180805490506103229190610d2a565b8110156103c257600180826103379190610d5d565b8154811061034857610347610d90565b5b905f5260205f2090600202016001828154811061036857610367610d90565b5b905f5260205f2090600202015f8201815f0190816103869190610f6f565b50600182015f9054906101000a900460ff16816001015f6101000a81548160ff0219169083151502179055509050508080600101915050610311565b5060018054806103d5576103d4611054565b5b600190038181905f5260205f2090600202015f8082015f6103f69190610786565b600182015f6101000a81549060ff02191690555050905550565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610466575f80fd5b6001818154811061047a57610479610d90565b5b905f5260205f2090600202016001015f9054906101000a900460ff1615600182815481106104ab576104aa610d90565b5b905f5260205f2090600202016001015f6101000a81548160ff02191690831515021790555050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6104fe6107c3565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610554575f80fd5b6001828154811061056857610567610d90565b5b905f5260205f2090600202016040518060400160405290815f8201805461058e90610ccd565b80601f01602080910402602001604051908101604052809291908181526020018280546105ba90610ccd565b80156106055780601f106105dc57610100808354040283529160200191610605565b820191905f5260205f20905b8154815290600101906020018083116105e857829003601f168201915b50505050508152602001600182015f9054906101000a900460ff1615151515815250509050919050565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610685575f80fd5b600160405180604001604052808381526020015f1515815250908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f0190816106d79190611081565b506020820151816001015f6101000a81548160ff021916908315150217905550505050565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610752575f80fd5b806001838154811061076757610766610d90565b5b905f5260205f2090600202015f0190816107819190611081565b505050565b50805461079290610ccd565b5f825580601f106107a357506107c0565b601f0160209004905f5260205f20908101906107bf91906107de565b5b50565b6040518060400160405280606081526020015f151581525090565b5b808211156107f5575f815f9055506001016107df565b5090565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561085957808201518184015260208101905061083e565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61087e82610822565b610888818561082c565b935061089881856020860161083c565b6108a181610864565b840191505092915050565b5f8115159050919050565b6108c0816108ac565b82525050565b5f604083015f8301518482035f8601526108e08282610874565b91505060208301516108f560208601826108b7565b508091505092915050565b5f61090b83836108c6565b905092915050565b5f602082019050919050565b5f610929826107f9565b6109338185610803565b93508360208202850161094585610813565b805f5b8581101561098057848403895281516109618582610900565b945061096c83610913565b925060208a01995050600181019050610948565b50829750879550505050505092915050565b5f6020820190508181035f8301526109aa818461091f565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b6109d5816109c3565b81146109df575f80fd5b50565b5f813590506109f0816109cc565b92915050565b5f60208284031215610a0b57610a0a6109bb565b5b5f610a18848285016109e2565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610a4a82610a21565b9050919050565b610a5a81610a40565b82525050565b5f602082019050610a735f830184610a51565b92915050565b5f604083015f8301518482035f860152610a938282610874565b9150506020830151610aa860208601826108b7565b508091505092915050565b5f6020820190508181035f830152610acb8184610a79565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610b1182610864565b810181811067ffffffffffffffff82111715610b3057610b2f610adb565b5b80604052505050565b5f610b426109b2565b9050610b4e8282610b08565b919050565b5f67ffffffffffffffff821115610b6d57610b6c610adb565b5b610b7682610864565b9050602081019050919050565b828183375f83830152505050565b5f610ba3610b9e84610b53565b610b39565b905082815260208101848484011115610bbf57610bbe610ad7565b5b610bca848285610b83565b509392505050565b5f82601f830112610be657610be5610ad3565b5b8135610bf6848260208601610b91565b91505092915050565b5f60208284031215610c1457610c136109bb565b5b5f82013567ffffffffffffffff811115610c3157610c306109bf565b5b610c3d84828501610bd2565b91505092915050565b5f8060408385031215610c5c57610c5b6109bb565b5b5f610c69858286016109e2565b925050602083013567ffffffffffffffff811115610c8a57610c896109bf565b5b610c9685828601610bd2565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610ce457607f821691505b602082108103610cf757610cf6610ca0565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610d34826109c3565b9150610d3f836109c3565b9250828203905081811115610d5757610d56610cfd565b5b92915050565b5f610d67826109c3565b9150610d72836109c3565b9250828201905080821115610d8a57610d89610cfd565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81549050610dcb81610ccd565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610e2e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610df3565b610e388683610df3565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610e73610e6e610e69846109c3565b610e50565b6109c3565b9050919050565b5f819050919050565b610e8c83610e59565b610ea0610e9882610e7a565b848454610dff565b825550505050565b5f90565b610eb4610ea8565b610ebf818484610e83565b505050565b5b81811015610ee257610ed75f82610eac565b600181019050610ec5565b5050565b601f821115610f2757610ef881610dd2565b610f0184610de4565b81016020851015610f10578190505b610f24610f1c85610de4565b830182610ec4565b50505b505050565b5f82821c905092915050565b5f610f475f1984600802610f2c565b1980831691505092915050565b5f610f5f8383610f38565b9150826002028217905092915050565b818103610f7d575050611052565b610f8682610dbd565b67ffffffffffffffff811115610f9f57610f9e610adb565b5b610fa98254610ccd565b610fb4828285610ee6565b5f601f831160018114610fe1575f8415610fcf578287015490505b610fd98582610f54565b86555061104b565b601f198416610fef87610dd2565b9650610ffa86610dd2565b5f5b8281101561102157848901548255600182019150600185019450602081019050610ffc565b8683101561103e578489015461103a601f891682610f38565b8355505b6001600288020188555050505b5050505050505b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b61108a82610822565b67ffffffffffffffff8111156110a3576110a2610adb565b5b6110ad8254610ccd565b6110b8828285610ee6565b5f60209050601f8311600181146110e9575f84156110d7578287015190505b6110e18582610f54565b865550611148565b601f1984166110f786610dd2565b5f5b8281101561111e578489015182556001820191506020850194506020810190506110f9565b8683101561113b5784890151611137601f891682610f38565b8355505b6001600288020188555050505b50505050505056fea264697066735822122023b4cb6746fcc1244fd25f76ae3841837b1b69a33f9ec95b29672fd61021277e64736f6c63430008180033",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoCaller) Get(opts *bind.CallOpts, _id *big.Int) (TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "get", _id)

	if err != nil {
		return *new(TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTask)).(*TodoTask)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoSession) Get(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoCallerSession) Get(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _id)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCaller) List(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCallerSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCallerSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoTransactor) Add(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "add", _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoSession) Add(_content string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoTransactorSession) Add(_content string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactor) Remove(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "remove", _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactorSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Todo *TodoTransactor) Toggle(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "toggle", _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Todo *TodoSession) Toggle(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Toggle(&_Todo.TransactOpts, _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Todo *TodoTransactorSession) Toggle(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Toggle(&_Todo.TransactOpts, _id)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactor) Update(opts *bind.TransactOpts, _id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "update", _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactorSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}

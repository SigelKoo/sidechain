// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token_erc20

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
)

// TokenErc20MetaData contains all meta data concerning the TokenErc20 contract.
var TokenErc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200186f3803806200186f83398181016040528101906200003791906200033a565b6040518060400160405280600b81526020017f746f6b656e5f65726332300000000000000000000000000000000000000000008152506040518060400160405280600b81526020017f746f6b656e5f65726332300000000000000000000000000000000000000000008152508160039080519060200190620000bb92919062000273565b508060049080519060200190620000d492919062000273565b505050620000e93382620000f060201b60201c565b506200051e565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000163576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200015a90620003b9565b60405180910390fd5b62000177600083836200026960201b60201c565b80600260008282546200018b919062000409565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254620001e2919062000409565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051620002499190620003db565b60405180910390a362000265600083836200026e60201b60201c565b5050565b505050565b505050565b828054620002819062000470565b90600052602060002090601f016020900481019282620002a55760008555620002f1565b82601f10620002c057805160ff1916838001178555620002f1565b82800160010185558215620002f1579182015b82811115620002f0578251825591602001919060010190620002d3565b5b50905062000300919062000304565b5090565b5b808211156200031f57600081600090555060010162000305565b5090565b600081519050620003348162000504565b92915050565b6000602082840312156200034d57600080fd5b60006200035d8482850162000323565b91505092915050565b600062000375601f83620003f8565b91507f45524332303a206d696e7420746f20746865207a65726f2061646472657373006000830152602082019050919050565b620003b38162000466565b82525050565b60006020820190508181036000830152620003d48162000366565b9050919050565b6000602082019050620003f26000830184620003a8565b92915050565b600082825260208201905092915050565b6000620004168262000466565b9150620004238362000466565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156200045b576200045a620004a6565b5b828201905092915050565b6000819050919050565b600060028204905060018216806200048957607f821691505b60208210811415620004a0576200049f620004d5565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6200050f8162000466565b81146200051b57600080fd5b50565b611341806200052e6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461016857806370a082311461019857806395d89b41146101c8578063a457c2d7146101e6578063a9059cbb14610216578063dd62ed3e14610246576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100fc57806323b872dd1461011a578063313ce5671461014a575b600080fd5b6100b6610276565b6040516100c3919061100a565b60405180910390f35b6100e660048036038101906100e19190610c83565b610308565b6040516100f39190610fef565b60405180910390f35b610104610326565b604051610111919061110c565b60405180910390f35b610134600480360381019061012f9190610c34565b610330565b6040516101419190610fef565b60405180910390f35b610152610428565b60405161015f9190611127565b60405180910390f35b610182600480360381019061017d9190610c83565b610431565b60405161018f9190610fef565b60405180910390f35b6101b260048036038101906101ad9190610bcf565b6104dd565b6040516101bf919061110c565b60405180910390f35b6101d0610525565b6040516101dd919061100a565b60405180910390f35b61020060048036038101906101fb9190610c83565b6105b7565b60405161020d9190610fef565b60405180910390f35b610230600480360381019061022b9190610c83565b6106a2565b60405161023d9190610fef565b60405180910390f35b610260600480360381019061025b9190610bf8565b6106c0565b60405161026d919061110c565b60405180910390f35b6060600380546102859061123c565b80601f01602080910402602001604051908101604052809291908181526020018280546102b19061123c565b80156102fe5780601f106102d3576101008083540402835291602001916102fe565b820191906000526020600020905b8154815290600101906020018083116102e157829003601f168201915b5050505050905090565b600061031c610315610747565b848461074f565b6001905092915050565b6000600254905090565b600061033d84848461091a565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610388610747565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610408576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ff9061108c565b60405180910390fd5b61041c85610414610747565b85840361074f565b60019150509392505050565b60006012905090565b60006104d361043e610747565b84846001600061044c610747565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546104ce919061115e565b61074f565b6001905092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600480546105349061123c565b80601f01602080910402602001604051908101604052809291908181526020018280546105609061123c565b80156105ad5780601f10610582576101008083540402835291602001916105ad565b820191906000526020600020905b81548152906001019060200180831161059057829003601f168201915b5050505050905090565b600080600160006105c6610747565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610683576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067a906110ec565b60405180910390fd5b61069761068e610747565b8585840361074f565b600191505092915050565b60006106b66106af610747565b848461091a565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156107bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b6906110cc565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561082f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108269061104c565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258360405161090d919061110c565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561098a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610981906110ac565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156109fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109f19061102c565b60405180910390fd5b610a05838383610b9b565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610a8b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a829061106c565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610b1e919061115e565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610b82919061110c565b60405180910390a3610b95848484610ba0565b50505050565b505050565b505050565b600081359050610bb4816112dd565b92915050565b600081359050610bc9816112f4565b92915050565b600060208284031215610be157600080fd5b6000610bef84828501610ba5565b91505092915050565b60008060408385031215610c0b57600080fd5b6000610c1985828601610ba5565b9250506020610c2a85828601610ba5565b9150509250929050565b600080600060608486031215610c4957600080fd5b6000610c5786828701610ba5565b9350506020610c6886828701610ba5565b9250506040610c7986828701610bba565b9150509250925092565b60008060408385031215610c9657600080fd5b6000610ca485828601610ba5565b9250506020610cb585828601610bba565b9150509250929050565b610cc8816111c6565b82525050565b6000610cd982611142565b610ce3818561114d565b9350610cf3818560208601611209565b610cfc816112cc565b840191505092915050565b6000610d1460238361114d565b91507f45524332303a207472616e7366657220746f20746865207a65726f206164647260008301527f65737300000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610d7a60228361114d565b91507f45524332303a20617070726f766520746f20746865207a65726f20616464726560008301527f73730000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610de060268361114d565b91507f45524332303a207472616e7366657220616d6f756e742065786365656473206260008301527f616c616e636500000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610e4660288361114d565b91507f45524332303a207472616e7366657220616d6f756e742065786365656473206160008301527f6c6c6f77616e63650000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610eac60258361114d565b91507f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008301527f64726573730000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610f1260248361114d565b91507f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008301527f72657373000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610f7860258361114d565b91507f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008301527f207a65726f0000000000000000000000000000000000000000000000000000006020830152604082019050919050565b610fda816111f2565b82525050565b610fe9816111fc565b82525050565b60006020820190506110046000830184610cbf565b92915050565b600060208201905081810360008301526110248184610cce565b905092915050565b6000602082019050818103600083015261104581610d07565b9050919050565b6000602082019050818103600083015261106581610d6d565b9050919050565b6000602082019050818103600083015261108581610dd3565b9050919050565b600060208201905081810360008301526110a581610e39565b9050919050565b600060208201905081810360008301526110c581610e9f565b9050919050565b600060208201905081810360008301526110e581610f05565b9050919050565b6000602082019050818103600083015261110581610f6b565b9050919050565b60006020820190506111216000830184610fd1565b92915050565b600060208201905061113c6000830184610fe0565b92915050565b600081519050919050565b600082825260208201905092915050565b6000611169826111f2565b9150611174836111f2565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156111a9576111a861126e565b5b828201905092915050565b60006111bf826111d2565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b8381101561122757808201518184015260208101905061120c565b83811115611236576000848401525b50505050565b6000600282049050600182168061125457607f821691505b602082108114156112685761126761129d565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b6112e6816111b4565b81146112f157600080fd5b50565b6112fd816111f2565b811461130857600080fd5b5056fea264697066735822122004c6866ed6ad4f01af6bd1fa5efbe8a61f56b281e9e62f5643cd85d44838ea7264736f6c63430008000033",
}

// TokenErc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenErc20MetaData.ABI instead.
var TokenErc20ABI = TokenErc20MetaData.ABI

// TokenErc20Bin is the compiled bytecode used for deploying new store-contracts.
// Deprecated: Use TokenErc20MetaData.Bin instead.
var TokenErc20Bin = TokenErc20MetaData.Bin

// DeployTokenErc20 deploys a new Ethereum contract, binding an instance of TokenErc20 to it.
func DeployTokenErc20(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *TokenErc20, error) {
	parsed, err := TokenErc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenErc20Bin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenErc20{TokenErc20Caller: TokenErc20Caller{contract: contract}, TokenErc20Transactor: TokenErc20Transactor{contract: contract}, TokenErc20Filterer: TokenErc20Filterer{contract: contract}}, nil
}

// TokenErc20 is an auto generated Go binding around an Ethereum contract.
type TokenErc20 struct {
	TokenErc20Caller     // Read-only binding to the contract
	TokenErc20Transactor // Write-only binding to the contract
	TokenErc20Filterer   // Log filterer for contract events
}

// TokenErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type TokenErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenErc20Session struct {
	Contract     *TokenErc20       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenErc20CallerSession struct {
	Contract *TokenErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenErc20TransactorSession struct {
	Contract     *TokenErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type TokenErc20Raw struct {
	Contract *TokenErc20 // Generic contract binding to access the raw methods on
}

// TokenErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenErc20CallerRaw struct {
	Contract *TokenErc20Caller // Generic read-only contract binding to access the raw methods on
}

// TokenErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenErc20TransactorRaw struct {
	Contract *TokenErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenErc20 creates a new instance of TokenErc20, bound to a specific deployed contract.
func NewTokenErc20(address common.Address, backend bind.ContractBackend) (*TokenErc20, error) {
	contract, err := bindTokenErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenErc20{TokenErc20Caller: TokenErc20Caller{contract: contract}, TokenErc20Transactor: TokenErc20Transactor{contract: contract}, TokenErc20Filterer: TokenErc20Filterer{contract: contract}}, nil
}

// NewTokenErc20Caller creates a new read-only instance of TokenErc20, bound to a specific deployed contract.
func NewTokenErc20Caller(address common.Address, caller bind.ContractCaller) (*TokenErc20Caller, error) {
	contract, err := bindTokenErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenErc20Caller{contract: contract}, nil
}

// NewTokenErc20Transactor creates a new write-only instance of TokenErc20, bound to a specific deployed contract.
func NewTokenErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*TokenErc20Transactor, error) {
	contract, err := bindTokenErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenErc20Transactor{contract: contract}, nil
}

// NewTokenErc20Filterer creates a new log filterer instance of TokenErc20, bound to a specific deployed contract.
func NewTokenErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*TokenErc20Filterer, error) {
	contract, err := bindTokenErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenErc20Filterer{contract: contract}, nil
}

// bindTokenErc20 binds a generic wrapper to an already deployed contract.
func bindTokenErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenErc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenErc20 *TokenErc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenErc20.Contract.TokenErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenErc20 *TokenErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenErc20.Contract.TokenErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenErc20 *TokenErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenErc20.Contract.TokenErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenErc20 *TokenErc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenErc20 *TokenErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenErc20 *TokenErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenErc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenErc20 *TokenErc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenErc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenErc20 *TokenErc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TokenErc20.Contract.Allowance(&_TokenErc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenErc20 *TokenErc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TokenErc20.Contract.Allowance(&_TokenErc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenErc20 *TokenErc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenErc20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenErc20 *TokenErc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _TokenErc20.Contract.BalanceOf(&_TokenErc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenErc20 *TokenErc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TokenErc20.Contract.BalanceOf(&_TokenErc20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenErc20 *TokenErc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenErc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenErc20 *TokenErc20Session) Decimals() (uint8, error) {
	return _TokenErc20.Contract.Decimals(&_TokenErc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenErc20 *TokenErc20CallerSession) Decimals() (uint8, error) {
	return _TokenErc20.Contract.Decimals(&_TokenErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenErc20 *TokenErc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenErc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenErc20 *TokenErc20Session) Name() (string, error) {
	return _TokenErc20.Contract.Name(&_TokenErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenErc20 *TokenErc20CallerSession) Name() (string, error) {
	return _TokenErc20.Contract.Name(&_TokenErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenErc20 *TokenErc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenErc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenErc20 *TokenErc20Session) Symbol() (string, error) {
	return _TokenErc20.Contract.Symbol(&_TokenErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenErc20 *TokenErc20CallerSession) Symbol() (string, error) {
	return _TokenErc20.Contract.Symbol(&_TokenErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenErc20 *TokenErc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenErc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenErc20 *TokenErc20Session) TotalSupply() (*big.Int, error) {
	return _TokenErc20.Contract.TotalSupply(&_TokenErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenErc20 *TokenErc20CallerSession) TotalSupply() (*big.Int, error) {
	return _TokenErc20.Contract.TotalSupply(&_TokenErc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenErc20 *TokenErc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenErc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenErc20 *TokenErc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenErc20.Contract.Approve(&_TokenErc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenErc20 *TokenErc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenErc20.Contract.Approve(&_TokenErc20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenErc20 *TokenErc20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenErc20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenErc20 *TokenErc20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenErc20.Contract.DecreaseAllowance(&_TokenErc20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenErc20 *TokenErc20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenErc20.Contract.DecreaseAllowance(&_TokenErc20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenErc20 *TokenErc20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenErc20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenErc20 *TokenErc20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenErc20.Contract.IncreaseAllowance(&_TokenErc20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenErc20 *TokenErc20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenErc20.Contract.IncreaseAllowance(&_TokenErc20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenErc20 *TokenErc20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenErc20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenErc20 *TokenErc20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenErc20.Contract.Transfer(&_TokenErc20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenErc20 *TokenErc20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenErc20.Contract.Transfer(&_TokenErc20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenErc20 *TokenErc20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenErc20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenErc20 *TokenErc20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenErc20.Contract.TransferFrom(&_TokenErc20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenErc20 *TokenErc20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenErc20.Contract.TransferFrom(&_TokenErc20.TransactOpts, sender, recipient, amount)
}

// TokenErc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TokenErc20 contract.
type TokenErc20ApprovalIterator struct {
	Event *TokenErc20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenErc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenErc20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenErc20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenErc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenErc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenErc20Approval represents a Approval event raised by the TokenErc20 contract.
type TokenErc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenErc20 *TokenErc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenErc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenErc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenErc20ApprovalIterator{contract: _TokenErc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenErc20 *TokenErc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenErc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenErc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenErc20Approval)
				if err := _TokenErc20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenErc20 *TokenErc20Filterer) ParseApproval(log types.Log) (*TokenErc20Approval, error) {
	event := new(TokenErc20Approval)
	if err := _TokenErc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenErc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenErc20 contract.
type TokenErc20TransferIterator struct {
	Event *TokenErc20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenErc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenErc20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenErc20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenErc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenErc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenErc20Transfer represents a Transfer event raised by the TokenErc20 contract.
type TokenErc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TokenErc20 *TokenErc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenErc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenErc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenErc20TransferIterator{contract: _TokenErc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TokenErc20 *TokenErc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenErc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenErc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenErc20Transfer)
				if err := _TokenErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TokenErc20 *TokenErc20Filterer) ParseTransfer(log types.Log) (*TokenErc20Transfer, error) {
	event := new(TokenErc20Transfer)
	if err := _TokenErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

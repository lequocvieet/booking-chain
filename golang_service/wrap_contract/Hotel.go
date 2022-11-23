// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrap_contract

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

// HotelMetaData contains all meta data concerning the Hotel contract.
var HotelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_landLord\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roomId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numberOfdates\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"booker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BookRoom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CancelBookRoom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"checker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CheckIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CheckOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listRoomId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CreateListRoom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deleter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DeleteListRoom\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roomId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numberOfdates\",\"type\":\"uint256\"}],\"name\":\"bookRoom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"cancelBookRoom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"checkIn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"checkOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createListRoom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listRoomId\",\"type\":\"uint256\"}],\"name\":\"deleteListRoom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getOwnerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"landLord\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listRoomCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listRoomIdToListRoom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listRoomId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listRooms\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listRoomId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractRoomNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalPrice\",\"type\":\"uint256\"}],\"name\":\"requestPayment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIdToRoomId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610f76380380610f7683398101604081905261002f9161007d565b6001600055600480546001600160a01b039384166001600160a01b031991821617909155600280549290931691161790556100b7565b6001600160a01b038116811461007a57600080fd5b50565b6000806040838503121561009057600080fd5b825161009b81610065565b60208401519092506100ac81610065565b809150509250929050565b610eb0806100c66000396000f3fe6080604052600436106100dd5760003560e01c80637c4e87fd1161007f5780639c79e301116100595780639c79e30114610252578063a8339a591461028f578063f38aac05146102a2578063fd56f5b5146102b757600080fd5b80637c4e87fd146101d55780637e6e2d66146101f5578063836387101461023257600080fd5b80631bd2eeee116100bb5780631bd2eeee1461014a57806347ccca021461016a5780635de988ab146101a25780635fe27291146101b557600080fd5b80630a523740146100e25780631358646e1461012257806315663ef614610137575b600080fd5b3480156100ee57600080fd5b5061010f6100fd366004610bdd565b60056020526000908152604090205481565b6040519081526020015b60405180910390f35b610135610130366004610bf6565b6102cd565b005b610135610145366004610cd3565b610410565b34801561015657600080fd5b50610135610165366004610bdd565b6105f9565b34801561017657600080fd5b5060045461018a906001600160a01b031681565b6040516001600160a01b039091168152602001610119565b6101356101b0366004610bdd565b6106f3565b3480156101c157600080fd5b5060025461018a906001600160a01b031681565b3480156101e157600080fd5b506101356101f0366004610bdd565b61078f565b34801561020157600080fd5b50610215610210366004610bdd565b610889565b604080519283526001600160a01b03909116602083015201610119565b34801561023e57600080fd5b5061018a61024d366004610bdd565b6108c0565b34801561025e57600080fd5b5061021561026d366004610bdd565b600660205260009081526040902080546001909101546001600160a01b031682565b61013561029d366004610d10565b610936565b3480156102ae57600080fd5b50610135610a44565b3480156102c357600080fd5b5061010f60015481565b6102d5610b84565b813410156103355760405162461bcd60e51b815260206004820152602260248201527f6e6f7420656e6f75676820657468657220746f20626f6f6b207468697320726f6044820152616f6d60f01b60648201526084015b60405180910390fd5b600480546040516309792c4960e31b81523392810192909252602482018390526000916001600160a01b0390911690634bc96248906044016020604051808303816000875af115801561038c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b09190610d55565b60408051868152602081018590529081018290523360608201524260808201529091507f324f54df56e9321588859c151dc1d48a827444c19bfdc05908760ec366f343049060a00160405180910390a15061040b6001600055565b505050565b60005b81518110156105ba5760045482516000916001600160a01b031690636352211e9085908590811061044657610446610d6e565b60200260200101516040518263ffffffff1660e01b815260040161046c91815260200190565b602060405180830381865afa158015610489573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ad9190610d84565b90506001600160a01b03811633146105075760405162461bcd60e51b815260206004820152601d60248201527f596f7520617265206e6f74206f776e6572206f662074686973206e6674000000604482015260640161032c565b60045483516001600160a01b03909116906323b872dd903390309087908790811061053457610534610d6e565b60209081029190910101516040516001600160e01b031960e086901b1681526001600160a01b0393841660048201529290911660248301526044820152606401600060405180830381600087803b15801561058e57600080fd5b505af11580156105a2573d6000803e3d6000fd5b505050505080806105b290610db4565b915050610413565b507fb6aa69ca1581faaa664c9aa421dea39b8a13e9ed194214b8538fe4fc3f5e46f18133426040516105ee93929190610ddb565b60405180910390a150565b600480546040516331a9108f60e11b815291820183905233916001600160a01b0390911690636352211e90602401602060405180830381865afa158015610644573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106689190610d84565b6001600160a01b0316146106be5760405162461bcd60e51b815260206004820152601d60248201527f596f7520617265206e6f74206f776e6572206f662074686973206e6674000000604482015260640161032c565b604080518281524260208201527f2c0711cfda20e0b1afe5fb3c3ea5a875558abe4b154988930eaa92b42f12e42e91016105ee565b6002546001600160a01b0316331461075e5760405162461bcd60e51b815260206004820152602860248201527f596f75206d757374206265204c616e64204c6f726420746f20526571756573746044820152670814185e5b595b9d60c21b606482015260840161032c565b604051339082156108fc029083906000818181858888f1935050505015801561078b573d6000803e3d6000fd5b5050565b610797610b84565b6002546001600160a01b031633146107c15760405162461bcd60e51b815260040161032c90610e31565b60005b6003548110156108465781600382815481106107e2576107e2610d6e565b90600052602060002090600202016000015403610834576003818154811061080c5761080c610d6e565b60009182526020822060029091020190815560010180546001600160a01b0319169055610846565b8061083e81610db4565b9150506107c4565b5060405142815233907fce2b0d64af608a7801d45bbe06df40258994f355b230e9b051daf5c18e5a08459060200160405180910390a26108866001600055565b50565b6003818154811061089957600080fd5b6000918252602090912060029091020180546001909101549091506001600160a01b031682565b600480546040516331a9108f60e11b81529182018390526000916001600160a01b0390911690636352211e90602401602060405180830381865afa15801561090c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109309190610d84565b92915050565b60005b82518110156109d65760045483516001600160a01b03909116906342966c689085908490811061096b5761096b610d6e565b60200260200101516040518263ffffffff1660e01b815260040161099191815260200190565b600060405180830381600087803b1580156109ab57600080fd5b505af11580156109bf573d6000803e3d6000fd5b5050505080806109ce90610db4565b915050610939565b50604051339082156108fc029083906000818181858888f19350505050158015610a04573d6000803e3d6000fd5b507f73446524eea227fd1c194dd0b9014ca0e1a8a828aca89faf082679cb67eeefe6823342604051610a3893929190610ddb565b60405180910390a15050565b610a4c610b84565b6002546001600160a01b03163314610a765760405162461bcd60e51b815260040161032c90610e31565b604080518082019091526000808252602082015260018054906000610a9a83610db4565b9091555050600180548252600280546001600160a01b0390811660208501908152600380548086018255600091909152855193027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b810193909355517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c9092018054929091166001600160a01b03199092169190911790555460405133917f41bb3013ce0a5e645449c98720f443086dd75d13e6b583535e44cfe2041dc28591610b6f91904290918252602082015260400190565b60405180910390a250610b826001600055565b565b600260005403610bd65760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161032c565b6002600055565b600060208284031215610bef57600080fd5b5035919050565b600080600060608486031215610c0b57600080fd5b505081359360208301359350604090920135919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610c4957600080fd5b8135602067ffffffffffffffff80831115610c6657610c66610c22565b8260051b604051601f19603f83011681018181108482111715610c8b57610c8b610c22565b604052938452858101830193838101925087851115610ca957600080fd5b83870191505b84821015610cc857813583529183019190830190610caf565b979650505050505050565b600060208284031215610ce557600080fd5b813567ffffffffffffffff811115610cfc57600080fd5b610d0884828501610c38565b949350505050565b60008060408385031215610d2357600080fd5b823567ffffffffffffffff811115610d3a57600080fd5b610d4685828601610c38565b95602094909401359450505050565b600060208284031215610d6757600080fd5b5051919050565b634e487b7160e01b600052603260045260246000fd5b600060208284031215610d9657600080fd5b81516001600160a01b0381168114610dad57600080fd5b9392505050565b600060018201610dd457634e487b7160e01b600052601160045260246000fd5b5060010190565b606080825284519082018190526000906020906080840190828801845b82811015610e1457815184529284019290840190600101610df8565b5050506001600160a01b0395909516908301525060400152919050565b60208082526029908201527f596f75206d757374206265204c616e64204c6f726420746f20437265617465206040820152684c69737420526f6f6d60b81b60608201526080019056fea26469706673582212205de36a2979eb4a80f454b2ca8be2e95a9636082809c3592ac73cf64b6da9c59164736f6c63430008110033",
}

// HotelABI is the input ABI used to generate the binding from.
// Deprecated: Use HotelMetaData.ABI instead.
var HotelABI = HotelMetaData.ABI

// HotelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HotelMetaData.Bin instead.
var HotelBin = HotelMetaData.Bin

// DeployHotel deploys a new Ethereum contract, binding an instance of Hotel to it.
func DeployHotel(auth *bind.TransactOpts, backend bind.ContractBackend, _nftAddress common.Address, _landLord common.Address) (common.Address, *types.Transaction, *Hotel, error) {
	parsed, err := HotelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HotelBin), backend, _nftAddress, _landLord)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hotel{HotelCaller: HotelCaller{contract: contract}, HotelTransactor: HotelTransactor{contract: contract}, HotelFilterer: HotelFilterer{contract: contract}}, nil
}

// Hotel is an auto generated Go binding around an Ethereum contract.
type Hotel struct {
	HotelCaller     // Read-only binding to the contract
	HotelTransactor // Write-only binding to the contract
	HotelFilterer   // Log filterer for contract events
}

// HotelCaller is an auto generated read-only Go binding around an Ethereum contract.
type HotelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HotelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HotelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HotelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HotelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HotelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HotelSession struct {
	Contract     *Hotel            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HotelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HotelCallerSession struct {
	Contract *HotelCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HotelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HotelTransactorSession struct {
	Contract     *HotelTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HotelRaw is an auto generated low-level Go binding around an Ethereum contract.
type HotelRaw struct {
	Contract *Hotel // Generic contract binding to access the raw methods on
}

// HotelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HotelCallerRaw struct {
	Contract *HotelCaller // Generic read-only contract binding to access the raw methods on
}

// HotelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HotelTransactorRaw struct {
	Contract *HotelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHotel creates a new instance of Hotel, bound to a specific deployed contract.
func NewHotel(address common.Address, backend bind.ContractBackend) (*Hotel, error) {
	contract, err := bindHotel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hotel{HotelCaller: HotelCaller{contract: contract}, HotelTransactor: HotelTransactor{contract: contract}, HotelFilterer: HotelFilterer{contract: contract}}, nil
}

// NewHotelCaller creates a new read-only instance of Hotel, bound to a specific deployed contract.
func NewHotelCaller(address common.Address, caller bind.ContractCaller) (*HotelCaller, error) {
	contract, err := bindHotel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HotelCaller{contract: contract}, nil
}

// NewHotelTransactor creates a new write-only instance of Hotel, bound to a specific deployed contract.
func NewHotelTransactor(address common.Address, transactor bind.ContractTransactor) (*HotelTransactor, error) {
	contract, err := bindHotel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HotelTransactor{contract: contract}, nil
}

// NewHotelFilterer creates a new log filterer instance of Hotel, bound to a specific deployed contract.
func NewHotelFilterer(address common.Address, filterer bind.ContractFilterer) (*HotelFilterer, error) {
	contract, err := bindHotel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HotelFilterer{contract: contract}, nil
}

// bindHotel binds a generic wrapper to an already deployed contract.
func bindHotel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HotelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hotel *HotelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hotel.Contract.HotelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hotel *HotelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hotel.Contract.HotelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hotel *HotelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hotel.Contract.HotelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hotel *HotelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hotel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hotel *HotelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hotel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hotel *HotelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hotel.Contract.contract.Transact(opts, method, params...)
}

// GetOwnerOf is a free data retrieval call binding the contract method 0x83638710.
//
// Solidity: function getOwnerOf(uint256 _tokenId) view returns(address)
func (_Hotel *HotelCaller) GetOwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hotel.contract.Call(opts, &out, "getOwnerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwnerOf is a free data retrieval call binding the contract method 0x83638710.
//
// Solidity: function getOwnerOf(uint256 _tokenId) view returns(address)
func (_Hotel *HotelSession) GetOwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Hotel.Contract.GetOwnerOf(&_Hotel.CallOpts, _tokenId)
}

// GetOwnerOf is a free data retrieval call binding the contract method 0x83638710.
//
// Solidity: function getOwnerOf(uint256 _tokenId) view returns(address)
func (_Hotel *HotelCallerSession) GetOwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Hotel.Contract.GetOwnerOf(&_Hotel.CallOpts, _tokenId)
}

// LandLord is a free data retrieval call binding the contract method 0x5fe27291.
//
// Solidity: function landLord() view returns(address)
func (_Hotel *HotelCaller) LandLord(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hotel.contract.Call(opts, &out, "landLord")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LandLord is a free data retrieval call binding the contract method 0x5fe27291.
//
// Solidity: function landLord() view returns(address)
func (_Hotel *HotelSession) LandLord() (common.Address, error) {
	return _Hotel.Contract.LandLord(&_Hotel.CallOpts)
}

// LandLord is a free data retrieval call binding the contract method 0x5fe27291.
//
// Solidity: function landLord() view returns(address)
func (_Hotel *HotelCallerSession) LandLord() (common.Address, error) {
	return _Hotel.Contract.LandLord(&_Hotel.CallOpts)
}

// ListRoomCount is a free data retrieval call binding the contract method 0xfd56f5b5.
//
// Solidity: function listRoomCount() view returns(uint256)
func (_Hotel *HotelCaller) ListRoomCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hotel.contract.Call(opts, &out, "listRoomCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ListRoomCount is a free data retrieval call binding the contract method 0xfd56f5b5.
//
// Solidity: function listRoomCount() view returns(uint256)
func (_Hotel *HotelSession) ListRoomCount() (*big.Int, error) {
	return _Hotel.Contract.ListRoomCount(&_Hotel.CallOpts)
}

// ListRoomCount is a free data retrieval call binding the contract method 0xfd56f5b5.
//
// Solidity: function listRoomCount() view returns(uint256)
func (_Hotel *HotelCallerSession) ListRoomCount() (*big.Int, error) {
	return _Hotel.Contract.ListRoomCount(&_Hotel.CallOpts)
}

// ListRoomIdToListRoom is a free data retrieval call binding the contract method 0x9c79e301.
//
// Solidity: function listRoomIdToListRoom(uint256 ) view returns(uint256 listRoomId, address owner)
func (_Hotel *HotelCaller) ListRoomIdToListRoom(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ListRoomId *big.Int
	Owner      common.Address
}, error) {
	var out []interface{}
	err := _Hotel.contract.Call(opts, &out, "listRoomIdToListRoom", arg0)

	outstruct := new(struct {
		ListRoomId *big.Int
		Owner      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListRoomId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ListRoomIdToListRoom is a free data retrieval call binding the contract method 0x9c79e301.
//
// Solidity: function listRoomIdToListRoom(uint256 ) view returns(uint256 listRoomId, address owner)
func (_Hotel *HotelSession) ListRoomIdToListRoom(arg0 *big.Int) (struct {
	ListRoomId *big.Int
	Owner      common.Address
}, error) {
	return _Hotel.Contract.ListRoomIdToListRoom(&_Hotel.CallOpts, arg0)
}

// ListRoomIdToListRoom is a free data retrieval call binding the contract method 0x9c79e301.
//
// Solidity: function listRoomIdToListRoom(uint256 ) view returns(uint256 listRoomId, address owner)
func (_Hotel *HotelCallerSession) ListRoomIdToListRoom(arg0 *big.Int) (struct {
	ListRoomId *big.Int
	Owner      common.Address
}, error) {
	return _Hotel.Contract.ListRoomIdToListRoom(&_Hotel.CallOpts, arg0)
}

// ListRooms is a free data retrieval call binding the contract method 0x7e6e2d66.
//
// Solidity: function listRooms(uint256 ) view returns(uint256 listRoomId, address owner)
func (_Hotel *HotelCaller) ListRooms(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ListRoomId *big.Int
	Owner      common.Address
}, error) {
	var out []interface{}
	err := _Hotel.contract.Call(opts, &out, "listRooms", arg0)

	outstruct := new(struct {
		ListRoomId *big.Int
		Owner      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListRoomId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ListRooms is a free data retrieval call binding the contract method 0x7e6e2d66.
//
// Solidity: function listRooms(uint256 ) view returns(uint256 listRoomId, address owner)
func (_Hotel *HotelSession) ListRooms(arg0 *big.Int) (struct {
	ListRoomId *big.Int
	Owner      common.Address
}, error) {
	return _Hotel.Contract.ListRooms(&_Hotel.CallOpts, arg0)
}

// ListRooms is a free data retrieval call binding the contract method 0x7e6e2d66.
//
// Solidity: function listRooms(uint256 ) view returns(uint256 listRoomId, address owner)
func (_Hotel *HotelCallerSession) ListRooms(arg0 *big.Int) (struct {
	ListRoomId *big.Int
	Owner      common.Address
}, error) {
	return _Hotel.Contract.ListRooms(&_Hotel.CallOpts, arg0)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Hotel *HotelCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hotel.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Hotel *HotelSession) Nft() (common.Address, error) {
	return _Hotel.Contract.Nft(&_Hotel.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Hotel *HotelCallerSession) Nft() (common.Address, error) {
	return _Hotel.Contract.Nft(&_Hotel.CallOpts)
}

// TokenIdToRoomId is a free data retrieval call binding the contract method 0x0a523740.
//
// Solidity: function tokenIdToRoomId(uint256 ) view returns(uint256)
func (_Hotel *HotelCaller) TokenIdToRoomId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hotel.contract.Call(opts, &out, "tokenIdToRoomId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToRoomId is a free data retrieval call binding the contract method 0x0a523740.
//
// Solidity: function tokenIdToRoomId(uint256 ) view returns(uint256)
func (_Hotel *HotelSession) TokenIdToRoomId(arg0 *big.Int) (*big.Int, error) {
	return _Hotel.Contract.TokenIdToRoomId(&_Hotel.CallOpts, arg0)
}

// TokenIdToRoomId is a free data retrieval call binding the contract method 0x0a523740.
//
// Solidity: function tokenIdToRoomId(uint256 ) view returns(uint256)
func (_Hotel *HotelCallerSession) TokenIdToRoomId(arg0 *big.Int) (*big.Int, error) {
	return _Hotel.Contract.TokenIdToRoomId(&_Hotel.CallOpts, arg0)
}

// BookRoom is a paid mutator transaction binding the contract method 0x1358646e.
//
// Solidity: function bookRoom(uint256 _roomId, uint256 _totalPrice, uint256 _numberOfdates) payable returns()
func (_Hotel *HotelTransactor) BookRoom(opts *bind.TransactOpts, _roomId *big.Int, _totalPrice *big.Int, _numberOfdates *big.Int) (*types.Transaction, error) {
	return _Hotel.contract.Transact(opts, "bookRoom", _roomId, _totalPrice, _numberOfdates)
}

// BookRoom is a paid mutator transaction binding the contract method 0x1358646e.
//
// Solidity: function bookRoom(uint256 _roomId, uint256 _totalPrice, uint256 _numberOfdates) payable returns()
func (_Hotel *HotelSession) BookRoom(_roomId *big.Int, _totalPrice *big.Int, _numberOfdates *big.Int) (*types.Transaction, error) {
	return _Hotel.Contract.BookRoom(&_Hotel.TransactOpts, _roomId, _totalPrice, _numberOfdates)
}

// BookRoom is a paid mutator transaction binding the contract method 0x1358646e.
//
// Solidity: function bookRoom(uint256 _roomId, uint256 _totalPrice, uint256 _numberOfdates) payable returns()
func (_Hotel *HotelTransactorSession) BookRoom(_roomId *big.Int, _totalPrice *big.Int, _numberOfdates *big.Int) (*types.Transaction, error) {
	return _Hotel.Contract.BookRoom(&_Hotel.TransactOpts, _roomId, _totalPrice, _numberOfdates)
}

// CancelBookRoom is a paid mutator transaction binding the contract method 0xa8339a59.
//
// Solidity: function cancelBookRoom(uint256[] _tokenIds, uint256 _price) payable returns()
func (_Hotel *HotelTransactor) CancelBookRoom(opts *bind.TransactOpts, _tokenIds []*big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Hotel.contract.Transact(opts, "cancelBookRoom", _tokenIds, _price)
}

// CancelBookRoom is a paid mutator transaction binding the contract method 0xa8339a59.
//
// Solidity: function cancelBookRoom(uint256[] _tokenIds, uint256 _price) payable returns()
func (_Hotel *HotelSession) CancelBookRoom(_tokenIds []*big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Hotel.Contract.CancelBookRoom(&_Hotel.TransactOpts, _tokenIds, _price)
}

// CancelBookRoom is a paid mutator transaction binding the contract method 0xa8339a59.
//
// Solidity: function cancelBookRoom(uint256[] _tokenIds, uint256 _price) payable returns()
func (_Hotel *HotelTransactorSession) CancelBookRoom(_tokenIds []*big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Hotel.Contract.CancelBookRoom(&_Hotel.TransactOpts, _tokenIds, _price)
}

// CheckIn is a paid mutator transaction binding the contract method 0x15663ef6.
//
// Solidity: function checkIn(uint256[] _tokenIds) payable returns()
func (_Hotel *HotelTransactor) CheckIn(opts *bind.TransactOpts, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Hotel.contract.Transact(opts, "checkIn", _tokenIds)
}

// CheckIn is a paid mutator transaction binding the contract method 0x15663ef6.
//
// Solidity: function checkIn(uint256[] _tokenIds) payable returns()
func (_Hotel *HotelSession) CheckIn(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _Hotel.Contract.CheckIn(&_Hotel.TransactOpts, _tokenIds)
}

// CheckIn is a paid mutator transaction binding the contract method 0x15663ef6.
//
// Solidity: function checkIn(uint256[] _tokenIds) payable returns()
func (_Hotel *HotelTransactorSession) CheckIn(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _Hotel.Contract.CheckIn(&_Hotel.TransactOpts, _tokenIds)
}

// CheckOut is a paid mutator transaction binding the contract method 0x1bd2eeee.
//
// Solidity: function checkOut(uint256 _tokenId) returns()
func (_Hotel *HotelTransactor) CheckOut(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Hotel.contract.Transact(opts, "checkOut", _tokenId)
}

// CheckOut is a paid mutator transaction binding the contract method 0x1bd2eeee.
//
// Solidity: function checkOut(uint256 _tokenId) returns()
func (_Hotel *HotelSession) CheckOut(_tokenId *big.Int) (*types.Transaction, error) {
	return _Hotel.Contract.CheckOut(&_Hotel.TransactOpts, _tokenId)
}

// CheckOut is a paid mutator transaction binding the contract method 0x1bd2eeee.
//
// Solidity: function checkOut(uint256 _tokenId) returns()
func (_Hotel *HotelTransactorSession) CheckOut(_tokenId *big.Int) (*types.Transaction, error) {
	return _Hotel.Contract.CheckOut(&_Hotel.TransactOpts, _tokenId)
}

// CreateListRoom is a paid mutator transaction binding the contract method 0xf38aac05.
//
// Solidity: function createListRoom() returns()
func (_Hotel *HotelTransactor) CreateListRoom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hotel.contract.Transact(opts, "createListRoom")
}

// CreateListRoom is a paid mutator transaction binding the contract method 0xf38aac05.
//
// Solidity: function createListRoom() returns()
func (_Hotel *HotelSession) CreateListRoom() (*types.Transaction, error) {
	return _Hotel.Contract.CreateListRoom(&_Hotel.TransactOpts)
}

// CreateListRoom is a paid mutator transaction binding the contract method 0xf38aac05.
//
// Solidity: function createListRoom() returns()
func (_Hotel *HotelTransactorSession) CreateListRoom() (*types.Transaction, error) {
	return _Hotel.Contract.CreateListRoom(&_Hotel.TransactOpts)
}

// DeleteListRoom is a paid mutator transaction binding the contract method 0x7c4e87fd.
//
// Solidity: function deleteListRoom(uint256 _listRoomId) returns()
func (_Hotel *HotelTransactor) DeleteListRoom(opts *bind.TransactOpts, _listRoomId *big.Int) (*types.Transaction, error) {
	return _Hotel.contract.Transact(opts, "deleteListRoom", _listRoomId)
}

// DeleteListRoom is a paid mutator transaction binding the contract method 0x7c4e87fd.
//
// Solidity: function deleteListRoom(uint256 _listRoomId) returns()
func (_Hotel *HotelSession) DeleteListRoom(_listRoomId *big.Int) (*types.Transaction, error) {
	return _Hotel.Contract.DeleteListRoom(&_Hotel.TransactOpts, _listRoomId)
}

// DeleteListRoom is a paid mutator transaction binding the contract method 0x7c4e87fd.
//
// Solidity: function deleteListRoom(uint256 _listRoomId) returns()
func (_Hotel *HotelTransactorSession) DeleteListRoom(_listRoomId *big.Int) (*types.Transaction, error) {
	return _Hotel.Contract.DeleteListRoom(&_Hotel.TransactOpts, _listRoomId)
}

// RequestPayment is a paid mutator transaction binding the contract method 0x5de988ab.
//
// Solidity: function requestPayment(uint256 _totalPrice) payable returns()
func (_Hotel *HotelTransactor) RequestPayment(opts *bind.TransactOpts, _totalPrice *big.Int) (*types.Transaction, error) {
	return _Hotel.contract.Transact(opts, "requestPayment", _totalPrice)
}

// RequestPayment is a paid mutator transaction binding the contract method 0x5de988ab.
//
// Solidity: function requestPayment(uint256 _totalPrice) payable returns()
func (_Hotel *HotelSession) RequestPayment(_totalPrice *big.Int) (*types.Transaction, error) {
	return _Hotel.Contract.RequestPayment(&_Hotel.TransactOpts, _totalPrice)
}

// RequestPayment is a paid mutator transaction binding the contract method 0x5de988ab.
//
// Solidity: function requestPayment(uint256 _totalPrice) payable returns()
func (_Hotel *HotelTransactorSession) RequestPayment(_totalPrice *big.Int) (*types.Transaction, error) {
	return _Hotel.Contract.RequestPayment(&_Hotel.TransactOpts, _totalPrice)
}

// HotelBookRoomIterator is returned from FilterBookRoom and is used to iterate over the raw logs and unpacked data for BookRoom events raised by the Hotel contract.
type HotelBookRoomIterator struct {
	Event *HotelBookRoom // Event containing the contract specifics and raw log

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
func (it *HotelBookRoomIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotelBookRoom)
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
		it.Event = new(HotelBookRoom)
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
func (it *HotelBookRoomIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotelBookRoomIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotelBookRoom represents a BookRoom event raised by the Hotel contract.
type HotelBookRoom struct {
	RoomId        *big.Int
	NumberOfdates *big.Int
	StartTokenId  *big.Int
	Booker        common.Address
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBookRoom is a free log retrieval operation binding the contract event 0x324f54df56e9321588859c151dc1d48a827444c19bfdc05908760ec366f34304.
//
// Solidity: event BookRoom(uint256 roomId, uint256 numberOfdates, uint256 startTokenId, address booker, uint256 timestamp)
func (_Hotel *HotelFilterer) FilterBookRoom(opts *bind.FilterOpts) (*HotelBookRoomIterator, error) {

	logs, sub, err := _Hotel.contract.FilterLogs(opts, "BookRoom")
	if err != nil {
		return nil, err
	}
	return &HotelBookRoomIterator{contract: _Hotel.contract, event: "BookRoom", logs: logs, sub: sub}, nil
}

// WatchBookRoom is a free log subscription operation binding the contract event 0x324f54df56e9321588859c151dc1d48a827444c19bfdc05908760ec366f34304.
//
// Solidity: event BookRoom(uint256 roomId, uint256 numberOfdates, uint256 startTokenId, address booker, uint256 timestamp)
func (_Hotel *HotelFilterer) WatchBookRoom(opts *bind.WatchOpts, sink chan<- *HotelBookRoom) (event.Subscription, error) {

	logs, sub, err := _Hotel.contract.WatchLogs(opts, "BookRoom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotelBookRoom)
				if err := _Hotel.contract.UnpackLog(event, "BookRoom", log); err != nil {
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

// ParseBookRoom is a log parse operation binding the contract event 0x324f54df56e9321588859c151dc1d48a827444c19bfdc05908760ec366f34304.
//
// Solidity: event BookRoom(uint256 roomId, uint256 numberOfdates, uint256 startTokenId, address booker, uint256 timestamp)
func (_Hotel *HotelFilterer) ParseBookRoom(log types.Log) (*HotelBookRoom, error) {
	event := new(HotelBookRoom)
	if err := _Hotel.contract.UnpackLog(event, "BookRoom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotelCancelBookRoomIterator is returned from FilterCancelBookRoom and is used to iterate over the raw logs and unpacked data for CancelBookRoom events raised by the Hotel contract.
type HotelCancelBookRoomIterator struct {
	Event *HotelCancelBookRoom // Event containing the contract specifics and raw log

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
func (it *HotelCancelBookRoomIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotelCancelBookRoom)
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
		it.Event = new(HotelCancelBookRoom)
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
func (it *HotelCancelBookRoomIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotelCancelBookRoomIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotelCancelBookRoom represents a CancelBookRoom event raised by the Hotel contract.
type HotelCancelBookRoom struct {
	TokenIds  []*big.Int
	Canceler  common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCancelBookRoom is a free log retrieval operation binding the contract event 0x73446524eea227fd1c194dd0b9014ca0e1a8a828aca89faf082679cb67eeefe6.
//
// Solidity: event CancelBookRoom(uint256[] tokenIds, address canceler, uint256 timestamp)
func (_Hotel *HotelFilterer) FilterCancelBookRoom(opts *bind.FilterOpts) (*HotelCancelBookRoomIterator, error) {

	logs, sub, err := _Hotel.contract.FilterLogs(opts, "CancelBookRoom")
	if err != nil {
		return nil, err
	}
	return &HotelCancelBookRoomIterator{contract: _Hotel.contract, event: "CancelBookRoom", logs: logs, sub: sub}, nil
}

// WatchCancelBookRoom is a free log subscription operation binding the contract event 0x73446524eea227fd1c194dd0b9014ca0e1a8a828aca89faf082679cb67eeefe6.
//
// Solidity: event CancelBookRoom(uint256[] tokenIds, address canceler, uint256 timestamp)
func (_Hotel *HotelFilterer) WatchCancelBookRoom(opts *bind.WatchOpts, sink chan<- *HotelCancelBookRoom) (event.Subscription, error) {

	logs, sub, err := _Hotel.contract.WatchLogs(opts, "CancelBookRoom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotelCancelBookRoom)
				if err := _Hotel.contract.UnpackLog(event, "CancelBookRoom", log); err != nil {
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

// ParseCancelBookRoom is a log parse operation binding the contract event 0x73446524eea227fd1c194dd0b9014ca0e1a8a828aca89faf082679cb67eeefe6.
//
// Solidity: event CancelBookRoom(uint256[] tokenIds, address canceler, uint256 timestamp)
func (_Hotel *HotelFilterer) ParseCancelBookRoom(log types.Log) (*HotelCancelBookRoom, error) {
	event := new(HotelCancelBookRoom)
	if err := _Hotel.contract.UnpackLog(event, "CancelBookRoom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotelCheckInIterator is returned from FilterCheckIn and is used to iterate over the raw logs and unpacked data for CheckIn events raised by the Hotel contract.
type HotelCheckInIterator struct {
	Event *HotelCheckIn // Event containing the contract specifics and raw log

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
func (it *HotelCheckInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotelCheckIn)
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
		it.Event = new(HotelCheckIn)
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
func (it *HotelCheckInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotelCheckInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotelCheckIn represents a CheckIn event raised by the Hotel contract.
type HotelCheckIn struct {
	TokenIds  []*big.Int
	Checker   common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCheckIn is a free log retrieval operation binding the contract event 0xb6aa69ca1581faaa664c9aa421dea39b8a13e9ed194214b8538fe4fc3f5e46f1.
//
// Solidity: event CheckIn(uint256[] tokenIds, address checker, uint256 timestamp)
func (_Hotel *HotelFilterer) FilterCheckIn(opts *bind.FilterOpts) (*HotelCheckInIterator, error) {

	logs, sub, err := _Hotel.contract.FilterLogs(opts, "CheckIn")
	if err != nil {
		return nil, err
	}
	return &HotelCheckInIterator{contract: _Hotel.contract, event: "CheckIn", logs: logs, sub: sub}, nil
}

// WatchCheckIn is a free log subscription operation binding the contract event 0xb6aa69ca1581faaa664c9aa421dea39b8a13e9ed194214b8538fe4fc3f5e46f1.
//
// Solidity: event CheckIn(uint256[] tokenIds, address checker, uint256 timestamp)
func (_Hotel *HotelFilterer) WatchCheckIn(opts *bind.WatchOpts, sink chan<- *HotelCheckIn) (event.Subscription, error) {

	logs, sub, err := _Hotel.contract.WatchLogs(opts, "CheckIn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotelCheckIn)
				if err := _Hotel.contract.UnpackLog(event, "CheckIn", log); err != nil {
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

// ParseCheckIn is a log parse operation binding the contract event 0xb6aa69ca1581faaa664c9aa421dea39b8a13e9ed194214b8538fe4fc3f5e46f1.
//
// Solidity: event CheckIn(uint256[] tokenIds, address checker, uint256 timestamp)
func (_Hotel *HotelFilterer) ParseCheckIn(log types.Log) (*HotelCheckIn, error) {
	event := new(HotelCheckIn)
	if err := _Hotel.contract.UnpackLog(event, "CheckIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotelCheckOutIterator is returned from FilterCheckOut and is used to iterate over the raw logs and unpacked data for CheckOut events raised by the Hotel contract.
type HotelCheckOutIterator struct {
	Event *HotelCheckOut // Event containing the contract specifics and raw log

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
func (it *HotelCheckOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotelCheckOut)
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
		it.Event = new(HotelCheckOut)
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
func (it *HotelCheckOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotelCheckOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotelCheckOut represents a CheckOut event raised by the Hotel contract.
type HotelCheckOut struct {
	TokenId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCheckOut is a free log retrieval operation binding the contract event 0x2c0711cfda20e0b1afe5fb3c3ea5a875558abe4b154988930eaa92b42f12e42e.
//
// Solidity: event CheckOut(uint256 tokenId, uint256 timestamp)
func (_Hotel *HotelFilterer) FilterCheckOut(opts *bind.FilterOpts) (*HotelCheckOutIterator, error) {

	logs, sub, err := _Hotel.contract.FilterLogs(opts, "CheckOut")
	if err != nil {
		return nil, err
	}
	return &HotelCheckOutIterator{contract: _Hotel.contract, event: "CheckOut", logs: logs, sub: sub}, nil
}

// WatchCheckOut is a free log subscription operation binding the contract event 0x2c0711cfda20e0b1afe5fb3c3ea5a875558abe4b154988930eaa92b42f12e42e.
//
// Solidity: event CheckOut(uint256 tokenId, uint256 timestamp)
func (_Hotel *HotelFilterer) WatchCheckOut(opts *bind.WatchOpts, sink chan<- *HotelCheckOut) (event.Subscription, error) {

	logs, sub, err := _Hotel.contract.WatchLogs(opts, "CheckOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotelCheckOut)
				if err := _Hotel.contract.UnpackLog(event, "CheckOut", log); err != nil {
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

// ParseCheckOut is a log parse operation binding the contract event 0x2c0711cfda20e0b1afe5fb3c3ea5a875558abe4b154988930eaa92b42f12e42e.
//
// Solidity: event CheckOut(uint256 tokenId, uint256 timestamp)
func (_Hotel *HotelFilterer) ParseCheckOut(log types.Log) (*HotelCheckOut, error) {
	event := new(HotelCheckOut)
	if err := _Hotel.contract.UnpackLog(event, "CheckOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotelCreateListRoomIterator is returned from FilterCreateListRoom and is used to iterate over the raw logs and unpacked data for CreateListRoom events raised by the Hotel contract.
type HotelCreateListRoomIterator struct {
	Event *HotelCreateListRoom // Event containing the contract specifics and raw log

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
func (it *HotelCreateListRoomIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotelCreateListRoom)
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
		it.Event = new(HotelCreateListRoom)
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
func (it *HotelCreateListRoomIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotelCreateListRoomIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotelCreateListRoom represents a CreateListRoom event raised by the Hotel contract.
type HotelCreateListRoom struct {
	ListRoomId *big.Int
	Owner      common.Address
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreateListRoom is a free log retrieval operation binding the contract event 0x41bb3013ce0a5e645449c98720f443086dd75d13e6b583535e44cfe2041dc285.
//
// Solidity: event CreateListRoom(uint256 listRoomId, address indexed owner, uint256 timestamp)
func (_Hotel *HotelFilterer) FilterCreateListRoom(opts *bind.FilterOpts, owner []common.Address) (*HotelCreateListRoomIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Hotel.contract.FilterLogs(opts, "CreateListRoom", ownerRule)
	if err != nil {
		return nil, err
	}
	return &HotelCreateListRoomIterator{contract: _Hotel.contract, event: "CreateListRoom", logs: logs, sub: sub}, nil
}

// WatchCreateListRoom is a free log subscription operation binding the contract event 0x41bb3013ce0a5e645449c98720f443086dd75d13e6b583535e44cfe2041dc285.
//
// Solidity: event CreateListRoom(uint256 listRoomId, address indexed owner, uint256 timestamp)
func (_Hotel *HotelFilterer) WatchCreateListRoom(opts *bind.WatchOpts, sink chan<- *HotelCreateListRoom, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Hotel.contract.WatchLogs(opts, "CreateListRoom", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotelCreateListRoom)
				if err := _Hotel.contract.UnpackLog(event, "CreateListRoom", log); err != nil {
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

// ParseCreateListRoom is a log parse operation binding the contract event 0x41bb3013ce0a5e645449c98720f443086dd75d13e6b583535e44cfe2041dc285.
//
// Solidity: event CreateListRoom(uint256 listRoomId, address indexed owner, uint256 timestamp)
func (_Hotel *HotelFilterer) ParseCreateListRoom(log types.Log) (*HotelCreateListRoom, error) {
	event := new(HotelCreateListRoom)
	if err := _Hotel.contract.UnpackLog(event, "CreateListRoom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotelDeleteListRoomIterator is returned from FilterDeleteListRoom and is used to iterate over the raw logs and unpacked data for DeleteListRoom events raised by the Hotel contract.
type HotelDeleteListRoomIterator struct {
	Event *HotelDeleteListRoom // Event containing the contract specifics and raw log

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
func (it *HotelDeleteListRoomIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotelDeleteListRoom)
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
		it.Event = new(HotelDeleteListRoom)
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
func (it *HotelDeleteListRoomIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotelDeleteListRoomIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotelDeleteListRoom represents a DeleteListRoom event raised by the Hotel contract.
type HotelDeleteListRoom struct {
	Deleter   common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeleteListRoom is a free log retrieval operation binding the contract event 0xce2b0d64af608a7801d45bbe06df40258994f355b230e9b051daf5c18e5a0845.
//
// Solidity: event DeleteListRoom(address indexed deleter, uint256 timestamp)
func (_Hotel *HotelFilterer) FilterDeleteListRoom(opts *bind.FilterOpts, deleter []common.Address) (*HotelDeleteListRoomIterator, error) {

	var deleterRule []interface{}
	for _, deleterItem := range deleter {
		deleterRule = append(deleterRule, deleterItem)
	}

	logs, sub, err := _Hotel.contract.FilterLogs(opts, "DeleteListRoom", deleterRule)
	if err != nil {
		return nil, err
	}
	return &HotelDeleteListRoomIterator{contract: _Hotel.contract, event: "DeleteListRoom", logs: logs, sub: sub}, nil
}

// WatchDeleteListRoom is a free log subscription operation binding the contract event 0xce2b0d64af608a7801d45bbe06df40258994f355b230e9b051daf5c18e5a0845.
//
// Solidity: event DeleteListRoom(address indexed deleter, uint256 timestamp)
func (_Hotel *HotelFilterer) WatchDeleteListRoom(opts *bind.WatchOpts, sink chan<- *HotelDeleteListRoom, deleter []common.Address) (event.Subscription, error) {

	var deleterRule []interface{}
	for _, deleterItem := range deleter {
		deleterRule = append(deleterRule, deleterItem)
	}

	logs, sub, err := _Hotel.contract.WatchLogs(opts, "DeleteListRoom", deleterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotelDeleteListRoom)
				if err := _Hotel.contract.UnpackLog(event, "DeleteListRoom", log); err != nil {
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

// ParseDeleteListRoom is a log parse operation binding the contract event 0xce2b0d64af608a7801d45bbe06df40258994f355b230e9b051daf5c18e5a0845.
//
// Solidity: event DeleteListRoom(address indexed deleter, uint256 timestamp)
func (_Hotel *HotelFilterer) ParseDeleteListRoom(log types.Log) (*HotelDeleteListRoom, error) {
	event := new(HotelDeleteListRoom)
	if err := _Hotel.contract.UnpackLog(event, "DeleteListRoom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

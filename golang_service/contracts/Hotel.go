// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// HotelListRoom is an auto generated low-level Go binding around an user-defined struct.
type HotelListRoom struct {
	ListRoomId *big.Int
	Owner      common.Address
}

// HotelMetaData contains all meta data concerning the Hotel contract.
var HotelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numberOfdates\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"booker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BookRoom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"canceler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CancelBookRoom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"checker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CheckIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CheckOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listRoomId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CreateListRoom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deleter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DeleteListRoom\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numberOfdates\",\"type\":\"uint256\"}],\"name\":\"bookRoom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"cancelBookRoom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"checkIn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createListRoom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listRoomId\",\"type\":\"uint256\"}],\"name\":\"deleteListRoom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getListRooms\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listRoomId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structHotel.ListRoom[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getOwnerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listRoomCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listRoomIdToListRoom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listRoomId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listRooms\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listRoomId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractRoomNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalPrice\",\"type\":\"uint256\"}],\"name\":\"requestPayment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIdToRoomId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIdToRoomNFTState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"tranferRoomNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161105f38038061105f83398101604081905261002f91610059565b6001600055600380546001600160a01b0319166001600160a01b0392909216919091179055610089565b60006020828403121561006b57600080fd5b81516001600160a01b038116811461008257600080fd5b9392505050565b610fc7806100986000396000f3fe6080604052600436106101095760003560e01c80638a3a1a1811610095578063a8339a5911610064578063a8339a5914610302578063cba5751c14610315578063e8d565201461032a578063f38aac051461034c578063fd56f5b51461036157600080fd5b80638a3a1a181461025257806394656214146102655780639c79e301146102a5578063a03d5ecc146102e257600080fd5b80636896fabf116100dc5780636896fabf146101ae5780636f9fb98a146101c25780637c4e87fd146101d55780637e6e2d66146101f5578063836387101461023257600080fd5b80630a5237401461010e57806315663ef61461014e57806347ccca02146101635780635de988ab1461019b575b600080fd5b34801561011a57600080fd5b5061013b610129366004610ca2565b60046020526000908152604090205481565b6040519081526020015b60405180910390f35b61016161015c366004610d6c565b610377565b005b34801561016f57600080fd5b50600354610183906001600160a01b031681565b6040516001600160a01b039091168152602001610145565b6101616101a9366004610ca2565b61059d565b3480156101ba57600080fd5b50333161013b565b3480156101ce57600080fd5b504761013b565b3480156101e157600080fd5b506101616101f0366004610ca2565b6105ce565b34801561020157600080fd5b50610215610210366004610ca2565b61069e565b604080519283526001600160a01b03909116602083015201610145565b34801561023e57600080fd5b5061018361024d366004610ca2565b6106d5565b610161610260366004610da9565b610749565b34801561027157600080fd5b50610295610280366004610ca2565b60066020526000908152604090205460ff1681565b6040519015158152602001610145565b3480156102b157600080fd5b506102156102c0366004610ca2565b600560205260009081526040902080546001909101546001600160a01b031682565b3480156102ee57600080fd5b506101616102fd366004610de0565b6108b8565b610161610310366004610e32565b61097b565b34801561032157600080fd5b50610161610a91565b34801561033657600080fd5b5061033f610ac6565b6040516101459190610e77565b34801561035857600080fd5b50610161610b3b565b34801561036d57600080fd5b5061013b60015481565b60005b81518110156105565760035482516000916001600160a01b031690636352211e908590859081106103ad576103ad610ecf565b60200260200101516040518263ffffffff1660e01b81526004016103d391815260200190565b602060405180830381865afa1580156103f0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104149190610ee5565b90506001600160a01b03811633146104735760405162461bcd60e51b815260206004820152601d60248201527f596f7520617265206e6f74206f776e6572206f662074686973206e667400000060448201526064015b60405180910390fd5b6006600084848151811061048957610489610ecf565b60209081029190910181015182528101919091526040016000205460ff166104fd5760405162461bcd60e51b815260206004820152602160248201527f596f7572204e4654206e6f7420617661696c61626c6520746f20636865636b496044820152603760f91b606482015260840161046a565b60006006600085858151811061051557610515610ecf565b6020026020010151815260200190815260200160002060006101000a81548160ff02191690831515021790555050808061054e90610f09565b91505061037a565b50336001600160a01b03167fb6aa69ca1581faaa664c9aa421dea39b8a13e9ed194214b8538fe4fc3f5e46f18242604051610592929190610f30565b60405180910390a250565b604051339082156108fc029083906000818181858888f193505050501580156105ca573d6000803e3d6000fd5b5050565b6105d6610c49565b60005b60025481101561065b5781600282815481106105f7576105f7610ecf565b90600052602060002090600202016000015403610649576002818154811061062157610621610ecf565b60009182526020822060029091020190815560010180546001600160a01b031916905561065b565b8061065381610f09565b9150506105d9565b5060405142815233907fce2b0d64af608a7801d45bbe06df40258994f355b230e9b051daf5c18e5a08459060200160405180910390a261069b6001600055565b50565b600281815481106106ae57600080fd5b6000918252602090912060029091020180546001909101549091506001600160a01b031682565b6003546040516331a9108f60e11b8152600481018390526000916001600160a01b031690636352211e90602401602060405180830381865afa15801561071f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107439190610ee5565b92915050565b610751610c49565b813410156107ac5760405162461bcd60e51b815260206004820152602260248201527f6e6f7420656e6f75676820657468657220746f20626f6f6b207468697320726f6044820152616f6d60f01b606482015260840161046a565b6003546040516309792c4960e31b8152336004820152602481018390526000916001600160a01b031690634bc96248906044016020604051808303816000875af11580156107fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108229190610f78565b90508060005b83811015610869576000828152600660205260409020805460ff191660011790558161085381610f09565b925050808061086190610f09565b915050610828565b5060408051848152602081018490524281830152905133917f133b114160665140104c84b124c4f21992bbdb4dc0ece92ddcfaa9b4016cfc28919081900360600190a250506105ca6001600055565b60005b82518110156109765760035483516001600160a01b03909116906323b872dd90339085908790869081106108f1576108f1610ecf565b60209081029190910101516040516001600160e01b031960e086901b1681526001600160a01b0393841660048201529290911660248301526044820152606401600060405180830381600087803b15801561094b57600080fd5b505af115801561095f573d6000803e3d6000fd5b50505050808061096e90610f09565b9150506108bb565b505050565b60005b8251811015610a1b5760035483516001600160a01b03909116906342966c68908590849081106109b0576109b0610ecf565b60200260200101516040518263ffffffff1660e01b81526004016109d691815260200190565b600060405180830381600087803b1580156109f057600080fd5b505af1158015610a04573d6000803e3d6000fd5b505050508080610a1390610f09565b91505061097e565b50604051339082156108fc029083906000818181858888f19350505050158015610a49573d6000803e3d6000fd5b50336001600160a01b03167f73446524eea227fd1c194dd0b9014ca0e1a8a828aca89faf082679cb67eeefe68342604051610a85929190610f30565b60405180910390a25050565b6040514281527f2c276699fefbad45f18509233bc03bb9cfdf41b18f32756933188bdef35fef559060200160405180910390a1565b60606002805480602002602001604051908101604052809291908181526020016000905b82821015610b325760008481526020908190206040805180820190915260028502909101805482526001908101546001600160a01b0316828401529083529092019101610aea565b50505050905090565b610b43610c49565b604080518082019091526000808252602082015260018054906000610b6783610f09565b909155505060018054825233602083018181526002805480850182556000829052855191027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace81019190915590517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf90910180546001600160a01b03929092166001600160a01b031990921691909117905590546040517f41bb3013ce0a5e645449c98720f443086dd75d13e6b583535e44cfe2041dc28591610c34914290918252602082015260400190565b60405180910390a250610c476001600055565b565b600260005403610c9b5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161046a565b6002600055565b600060208284031215610cb457600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610ce257600080fd5b8135602067ffffffffffffffff80831115610cff57610cff610cbb565b8260051b604051601f19603f83011681018181108482111715610d2457610d24610cbb565b604052938452858101830193838101925087851115610d4257600080fd5b83870191505b84821015610d6157813583529183019190830190610d48565b979650505050505050565b600060208284031215610d7e57600080fd5b813567ffffffffffffffff811115610d9557600080fd5b610da184828501610cd1565b949350505050565b60008060408385031215610dbc57600080fd5b50508035926020909101359150565b6001600160a01b038116811461069b57600080fd5b60008060408385031215610df357600080fd5b823567ffffffffffffffff811115610e0a57600080fd5b610e1685828601610cd1565b9250506020830135610e2781610dcb565b809150509250929050565b60008060408385031215610e4557600080fd5b823567ffffffffffffffff811115610e5c57600080fd5b610e6885828601610cd1565b95602094909401359450505050565b602080825282518282018190526000919060409081850190868401855b82811015610ec2578151805185528601516001600160a01b0316868501529284019290850190600101610e94565b5091979650505050505050565b634e487b7160e01b600052603260045260246000fd5b600060208284031215610ef757600080fd5b8151610f0281610dcb565b9392505050565b600060018201610f2957634e487b7160e01b600052601160045260246000fd5b5060010190565b604080825283519082018190526000906020906060840190828701845b82811015610f6957815184529284019290840190600101610f4d565b50505092019290925292915050565b600060208284031215610f8a57600080fd5b505191905056fea26469706673582212206b83e44505cb49b3c10068e589a5218c26a6ba380284667d4e6e100ea6e6ba7d64736f6c63430008110033",
}

// HotelABI is the input ABI used to generate the binding from.
// Deprecated: Use HotelMetaData.ABI instead.
var HotelABI = HotelMetaData.ABI

// HotelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HotelMetaData.Bin instead.
var HotelBin = HotelMetaData.Bin

// DeployHotel deploys a new Ethereum contract, binding an instance of Hotel to it.
func DeployHotel(auth *bind.TransactOpts, backend bind.ContractBackend, _nftAddress common.Address) (common.Address, *types.Transaction, *Hotel, error) {
	parsed, err := HotelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HotelBin), backend, _nftAddress)
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

// GetAccountBalance is a free data retrieval call binding the contract method 0x6896fabf.
//
// Solidity: function getAccountBalance() view returns(uint256)
func (_Hotel *HotelCaller) GetAccountBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hotel.contract.Call(opts, &out, "getAccountBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountBalance is a free data retrieval call binding the contract method 0x6896fabf.
//
// Solidity: function getAccountBalance() view returns(uint256)
func (_Hotel *HotelSession) GetAccountBalance() (*big.Int, error) {
	return _Hotel.Contract.GetAccountBalance(&_Hotel.CallOpts)
}

// GetAccountBalance is a free data retrieval call binding the contract method 0x6896fabf.
//
// Solidity: function getAccountBalance() view returns(uint256)
func (_Hotel *HotelCallerSession) GetAccountBalance() (*big.Int, error) {
	return _Hotel.Contract.GetAccountBalance(&_Hotel.CallOpts)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Hotel *HotelCaller) GetContractBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hotel.contract.Call(opts, &out, "getContractBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Hotel *HotelSession) GetContractBalance() (*big.Int, error) {
	return _Hotel.Contract.GetContractBalance(&_Hotel.CallOpts)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Hotel *HotelCallerSession) GetContractBalance() (*big.Int, error) {
	return _Hotel.Contract.GetContractBalance(&_Hotel.CallOpts)
}

// GetListRooms is a free data retrieval call binding the contract method 0xe8d56520.
//
// Solidity: function getListRooms() view returns((uint256,address)[])
func (_Hotel *HotelCaller) GetListRooms(opts *bind.CallOpts) ([]HotelListRoom, error) {
	var out []interface{}
	err := _Hotel.contract.Call(opts, &out, "getListRooms")

	if err != nil {
		return *new([]HotelListRoom), err
	}

	out0 := *abi.ConvertType(out[0], new([]HotelListRoom)).(*[]HotelListRoom)

	return out0, err

}

// GetListRooms is a free data retrieval call binding the contract method 0xe8d56520.
//
// Solidity: function getListRooms() view returns((uint256,address)[])
func (_Hotel *HotelSession) GetListRooms() ([]HotelListRoom, error) {
	return _Hotel.Contract.GetListRooms(&_Hotel.CallOpts)
}

// GetListRooms is a free data retrieval call binding the contract method 0xe8d56520.
//
// Solidity: function getListRooms() view returns((uint256,address)[])
func (_Hotel *HotelCallerSession) GetListRooms() ([]HotelListRoom, error) {
	return _Hotel.Contract.GetListRooms(&_Hotel.CallOpts)
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

// TokenIdToRoomNFTState is a free data retrieval call binding the contract method 0x94656214.
//
// Solidity: function tokenIdToRoomNFTState(uint256 ) view returns(bool)
func (_Hotel *HotelCaller) TokenIdToRoomNFTState(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Hotel.contract.Call(opts, &out, "tokenIdToRoomNFTState", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenIdToRoomNFTState is a free data retrieval call binding the contract method 0x94656214.
//
// Solidity: function tokenIdToRoomNFTState(uint256 ) view returns(bool)
func (_Hotel *HotelSession) TokenIdToRoomNFTState(arg0 *big.Int) (bool, error) {
	return _Hotel.Contract.TokenIdToRoomNFTState(&_Hotel.CallOpts, arg0)
}

// TokenIdToRoomNFTState is a free data retrieval call binding the contract method 0x94656214.
//
// Solidity: function tokenIdToRoomNFTState(uint256 ) view returns(bool)
func (_Hotel *HotelCallerSession) TokenIdToRoomNFTState(arg0 *big.Int) (bool, error) {
	return _Hotel.Contract.TokenIdToRoomNFTState(&_Hotel.CallOpts, arg0)
}

// BookRoom is a paid mutator transaction binding the contract method 0x8a3a1a18.
//
// Solidity: function bookRoom(uint256 _totalPrice, uint256 _numberOfdates) payable returns()
func (_Hotel *HotelTransactor) BookRoom(opts *bind.TransactOpts, _totalPrice *big.Int, _numberOfdates *big.Int) (*types.Transaction, error) {
	return _Hotel.contract.Transact(opts, "bookRoom", _totalPrice, _numberOfdates)
}

// BookRoom is a paid mutator transaction binding the contract method 0x8a3a1a18.
//
// Solidity: function bookRoom(uint256 _totalPrice, uint256 _numberOfdates) payable returns()
func (_Hotel *HotelSession) BookRoom(_totalPrice *big.Int, _numberOfdates *big.Int) (*types.Transaction, error) {
	return _Hotel.Contract.BookRoom(&_Hotel.TransactOpts, _totalPrice, _numberOfdates)
}

// BookRoom is a paid mutator transaction binding the contract method 0x8a3a1a18.
//
// Solidity: function bookRoom(uint256 _totalPrice, uint256 _numberOfdates) payable returns()
func (_Hotel *HotelTransactorSession) BookRoom(_totalPrice *big.Int, _numberOfdates *big.Int) (*types.Transaction, error) {
	return _Hotel.Contract.BookRoom(&_Hotel.TransactOpts, _totalPrice, _numberOfdates)
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

// CheckOut is a paid mutator transaction binding the contract method 0xcba5751c.
//
// Solidity: function checkOut() returns()
func (_Hotel *HotelTransactor) CheckOut(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hotel.contract.Transact(opts, "checkOut")
}

// CheckOut is a paid mutator transaction binding the contract method 0xcba5751c.
//
// Solidity: function checkOut() returns()
func (_Hotel *HotelSession) CheckOut() (*types.Transaction, error) {
	return _Hotel.Contract.CheckOut(&_Hotel.TransactOpts)
}

// CheckOut is a paid mutator transaction binding the contract method 0xcba5751c.
//
// Solidity: function checkOut() returns()
func (_Hotel *HotelTransactorSession) CheckOut() (*types.Transaction, error) {
	return _Hotel.Contract.CheckOut(&_Hotel.TransactOpts)
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

// TranferRoomNFT is a paid mutator transaction binding the contract method 0xa03d5ecc.
//
// Solidity: function tranferRoomNFT(uint256[] _tokenIds, address receiver) returns()
func (_Hotel *HotelTransactor) TranferRoomNFT(opts *bind.TransactOpts, _tokenIds []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Hotel.contract.Transact(opts, "tranferRoomNFT", _tokenIds, receiver)
}

// TranferRoomNFT is a paid mutator transaction binding the contract method 0xa03d5ecc.
//
// Solidity: function tranferRoomNFT(uint256[] _tokenIds, address receiver) returns()
func (_Hotel *HotelSession) TranferRoomNFT(_tokenIds []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Hotel.Contract.TranferRoomNFT(&_Hotel.TransactOpts, _tokenIds, receiver)
}

// TranferRoomNFT is a paid mutator transaction binding the contract method 0xa03d5ecc.
//
// Solidity: function tranferRoomNFT(uint256[] _tokenIds, address receiver) returns()
func (_Hotel *HotelTransactorSession) TranferRoomNFT(_tokenIds []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Hotel.Contract.TranferRoomNFT(&_Hotel.TransactOpts, _tokenIds, receiver)
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
	NumberOfdates *big.Int
	StartTokenId  *big.Int
	Booker        common.Address
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBookRoom is a free log retrieval operation binding the contract event 0x133b114160665140104c84b124c4f21992bbdb4dc0ece92ddcfaa9b4016cfc28.
//
// Solidity: event BookRoom(uint256 numberOfdates, uint256 startTokenId, address indexed booker, uint256 timestamp)
func (_Hotel *HotelFilterer) FilterBookRoom(opts *bind.FilterOpts, booker []common.Address) (*HotelBookRoomIterator, error) {

	var bookerRule []interface{}
	for _, bookerItem := range booker {
		bookerRule = append(bookerRule, bookerItem)
	}

	logs, sub, err := _Hotel.contract.FilterLogs(opts, "BookRoom", bookerRule)
	if err != nil {
		return nil, err
	}
	return &HotelBookRoomIterator{contract: _Hotel.contract, event: "BookRoom", logs: logs, sub: sub}, nil
}

// WatchBookRoom is a free log subscription operation binding the contract event 0x133b114160665140104c84b124c4f21992bbdb4dc0ece92ddcfaa9b4016cfc28.
//
// Solidity: event BookRoom(uint256 numberOfdates, uint256 startTokenId, address indexed booker, uint256 timestamp)
func (_Hotel *HotelFilterer) WatchBookRoom(opts *bind.WatchOpts, sink chan<- *HotelBookRoom, booker []common.Address) (event.Subscription, error) {

	var bookerRule []interface{}
	for _, bookerItem := range booker {
		bookerRule = append(bookerRule, bookerItem)
	}

	logs, sub, err := _Hotel.contract.WatchLogs(opts, "BookRoom", bookerRule)
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

// ParseBookRoom is a log parse operation binding the contract event 0x133b114160665140104c84b124c4f21992bbdb4dc0ece92ddcfaa9b4016cfc28.
//
// Solidity: event BookRoom(uint256 numberOfdates, uint256 startTokenId, address indexed booker, uint256 timestamp)
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
// Solidity: event CancelBookRoom(uint256[] tokenIds, address indexed canceler, uint256 timestamp)
func (_Hotel *HotelFilterer) FilterCancelBookRoom(opts *bind.FilterOpts, canceler []common.Address) (*HotelCancelBookRoomIterator, error) {

	var cancelerRule []interface{}
	for _, cancelerItem := range canceler {
		cancelerRule = append(cancelerRule, cancelerItem)
	}

	logs, sub, err := _Hotel.contract.FilterLogs(opts, "CancelBookRoom", cancelerRule)
	if err != nil {
		return nil, err
	}
	return &HotelCancelBookRoomIterator{contract: _Hotel.contract, event: "CancelBookRoom", logs: logs, sub: sub}, nil
}

// WatchCancelBookRoom is a free log subscription operation binding the contract event 0x73446524eea227fd1c194dd0b9014ca0e1a8a828aca89faf082679cb67eeefe6.
//
// Solidity: event CancelBookRoom(uint256[] tokenIds, address indexed canceler, uint256 timestamp)
func (_Hotel *HotelFilterer) WatchCancelBookRoom(opts *bind.WatchOpts, sink chan<- *HotelCancelBookRoom, canceler []common.Address) (event.Subscription, error) {

	var cancelerRule []interface{}
	for _, cancelerItem := range canceler {
		cancelerRule = append(cancelerRule, cancelerItem)
	}

	logs, sub, err := _Hotel.contract.WatchLogs(opts, "CancelBookRoom", cancelerRule)
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
// Solidity: event CancelBookRoom(uint256[] tokenIds, address indexed canceler, uint256 timestamp)
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
// Solidity: event CheckIn(uint256[] tokenIds, address indexed checker, uint256 timestamp)
func (_Hotel *HotelFilterer) FilterCheckIn(opts *bind.FilterOpts, checker []common.Address) (*HotelCheckInIterator, error) {

	var checkerRule []interface{}
	for _, checkerItem := range checker {
		checkerRule = append(checkerRule, checkerItem)
	}

	logs, sub, err := _Hotel.contract.FilterLogs(opts, "CheckIn", checkerRule)
	if err != nil {
		return nil, err
	}
	return &HotelCheckInIterator{contract: _Hotel.contract, event: "CheckIn", logs: logs, sub: sub}, nil
}

// WatchCheckIn is a free log subscription operation binding the contract event 0xb6aa69ca1581faaa664c9aa421dea39b8a13e9ed194214b8538fe4fc3f5e46f1.
//
// Solidity: event CheckIn(uint256[] tokenIds, address indexed checker, uint256 timestamp)
func (_Hotel *HotelFilterer) WatchCheckIn(opts *bind.WatchOpts, sink chan<- *HotelCheckIn, checker []common.Address) (event.Subscription, error) {

	var checkerRule []interface{}
	for _, checkerItem := range checker {
		checkerRule = append(checkerRule, checkerItem)
	}

	logs, sub, err := _Hotel.contract.WatchLogs(opts, "CheckIn", checkerRule)
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
// Solidity: event CheckIn(uint256[] tokenIds, address indexed checker, uint256 timestamp)
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
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCheckOut is a free log retrieval operation binding the contract event 0x2c276699fefbad45f18509233bc03bb9cfdf41b18f32756933188bdef35fef55.
//
// Solidity: event CheckOut(uint256 timestamp)
func (_Hotel *HotelFilterer) FilterCheckOut(opts *bind.FilterOpts) (*HotelCheckOutIterator, error) {

	logs, sub, err := _Hotel.contract.FilterLogs(opts, "CheckOut")
	if err != nil {
		return nil, err
	}
	return &HotelCheckOutIterator{contract: _Hotel.contract, event: "CheckOut", logs: logs, sub: sub}, nil
}

// WatchCheckOut is a free log subscription operation binding the contract event 0x2c276699fefbad45f18509233bc03bb9cfdf41b18f32756933188bdef35fef55.
//
// Solidity: event CheckOut(uint256 timestamp)
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

// ParseCheckOut is a log parse operation binding the contract event 0x2c276699fefbad45f18509233bc03bb9cfdf41b18f32756933188bdef35fef55.
//
// Solidity: event CheckOut(uint256 timestamp)
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

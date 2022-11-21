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

// RoomNFTMetaData contains all meta data concerning the RoomNFT contract.
var RoomNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"safeMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeMintBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405180604001604052806007815260200166149bdbdb53919560ca1b81525060405180604001604052806004815260200163524f4f4d60e01b81525081600090816200006091906200018d565b5060016200006f82826200018d565b5050506200008c620000866200009260201b60201c565b62000096565b62000259565b3390565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200011357607f821691505b6020821081036200013457634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200018857600081815260208120601f850160051c81016020861015620001635750805b601f850160051c820191505b8181101562000184578281556001016200016f565b5050505b505050565b81516001600160401b03811115620001a957620001a9620000e8565b620001c181620001ba8454620000fe565b846200013a565b602080601f831160018114620001f95760008415620001e05750858301515b600019600386901b1c1916600185901b17855562000184565b600085815260208120601f198616915b828110156200022a5788860151825594840194600190910190840162000209565b5085821015620002495787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61171680620002696000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80636352211e116100ad578063a22cb46511610071578063a22cb46514610257578063b88d4fde1461026a578063c87b56dd1461027d578063e985e9c514610290578063f2fde38b146102cc57600080fd5b80636352211e1461021057806370a0823114610223578063715018a6146102365780638da5cb5b1461023e57806395d89b411461024f57600080fd5b806323b872dd116100f457806323b872dd146101a357806340d097c3146101b657806342842e0e146101d757806342966c68146101ea5780634bc96248146101fd57600080fd5b806301ffc9a71461012657806306fdde031461014e578063081812fc14610163578063095ea7b31461018e575b600080fd5b610139610134366004611252565b6102df565b60405190151581526020015b60405180910390f35b610156610331565b60405161014591906112bf565b6101766101713660046112d2565b6103c3565b6040516001600160a01b039091168152602001610145565b6101a161019c366004611302565b6103ea565b005b6101a16101b136600461132c565b610504565b6101c96101c4366004611368565b610536565b604051908152602001610145565b6101a16101e536600461132c565b61056d565b6101a16101f83660046112d2565b610588565b6101c961020b366004611302565b6105b9565b61017661021e3660046112d2565b61061b565b6101c9610231366004611368565b61067b565b6101a1610701565b6006546001600160a01b0316610176565b610156610715565b6101a1610265366004611383565b610724565b6101a16102783660046113d5565b610733565b61015661028b3660046112d2565b61076b565b61013961029e3660046114b1565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6101a16102da366004611368565b6107df565b60006001600160e01b031982166380ac58cd60e01b148061031057506001600160e01b03198216635b5e139f60e01b145b8061032b57506301ffc9a760e01b6001600160e01b03198316145b92915050565b606060008054610340906114e4565b80601f016020809104026020016040519081016040528092919081815260200182805461036c906114e4565b80156103b95780601f1061038e576101008083540402835291602001916103b9565b820191906000526020600020905b81548152906001019060200180831161039c57829003601f168201915b5050505050905090565b60006103ce82610855565b506000908152600460205260409020546001600160a01b031690565b60006103f58261061b565b9050806001600160a01b0316836001600160a01b0316036104675760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b03821614806104835750610483813361029e565b6104f55760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c000000606482015260840161045e565b6104ff83836108b4565b505050565b61050f335b82610922565b61052b5760405162461bcd60e51b815260040161045e9061151e565b6104ff8383836109a1565b6000610540610b12565b600061054b60075490565b905061055b600780546001019055565b6105658382610b6c565b90505b919050565b6104ff83838360405180602001604052806000815250610733565b61059133610509565b6105ad5760405162461bcd60e51b815260040161045e9061151e565b6105b681610b86565b50565b60006105c3610b12565b60006105ce60075490565b905060005b838110156106135760006105e660075490565b90506105f6600780546001019055565b6106008682610b6c565b508061060b81611581565b9150506105d3565b509392505050565b6000818152600260205260408120546001600160a01b0316806105655760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604482015260640161045e565b60006001600160a01b0382166106e55760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b606482015260840161045e565b506001600160a01b031660009081526003602052604090205490565b610709610b12565b6107136000610c29565b565b606060018054610340906114e4565b61072f338383610c7b565b5050565b61073d3383610922565b6107595760405162461bcd60e51b815260040161045e9061151e565b61076584848484610d49565b50505050565b606061077682610855565b600061078d60408051602081019091526000815290565b905060008151116107ad57604051806020016040528060008152506107d8565b806107b784610d7c565b6040516020016107c892919061159a565b6040516020818303038152906040525b9392505050565b6107e7610b12565b6001600160a01b03811661084c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161045e565b6105b681610c29565b6000818152600260205260409020546001600160a01b03166105b65760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604482015260640161045e565b600081815260046020526040902080546001600160a01b0319166001600160a01b03841690811790915581906108e98261061b565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b60008061092e8361061b565b9050806001600160a01b0316846001600160a01b0316148061097557506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b806109995750836001600160a01b031661098e846103c3565b6001600160a01b0316145b949350505050565b826001600160a01b03166109b48261061b565b6001600160a01b0316146109da5760405162461bcd60e51b815260040161045e906115c9565b6001600160a01b038216610a3c5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b606482015260840161045e565b610a498383836001610e0f565b826001600160a01b0316610a5c8261061b565b6001600160a01b031614610a825760405162461bcd60e51b815260040161045e906115c9565b600081815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260038552838620805460001901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6006546001600160a01b031633146107135760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161045e565b61072f828260405180602001604052806000815250610e97565b6000610b918261061b565b9050610ba1816000846001610e0f565b610baa8261061b565b600083815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0385168085526003845282852080546000190190558785526002909352818420805490911690555192935084927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b816001600160a01b0316836001600160a01b031603610cdc5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015260640161045e565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b610d548484846109a1565b610d6084848484610eca565b6107655760405162461bcd60e51b815260040161045e9061160e565b60606000610d8983610fcb565b600101905060008167ffffffffffffffff811115610da957610da96113bf565b6040519080825280601f01601f191660200182016040528015610dd3576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a8504945084610ddd57509392505050565b6001811115610765576001600160a01b03841615610e55576001600160a01b03841660009081526003602052604081208054839290610e4f908490611660565b90915550505b6001600160a01b03831615610765576001600160a01b03831660009081526003602052604081208054839290610e8c908490611673565b909155505050505050565b610ea183836110a3565b610eae6000848484610eca565b6104ff5760405162461bcd60e51b815260040161045e9061160e565b60006001600160a01b0384163b15610fc057604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290610f0e903390899088908890600401611686565b6020604051808303816000875af1925050508015610f49575060408051601f3d908101601f19168201909252610f46918101906116c3565b60015b610fa6573d808015610f77576040519150601f19603f3d011682016040523d82523d6000602084013e610f7c565b606091505b508051600003610f9e5760405162461bcd60e51b815260040161045e9061160e565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610999565b506001949350505050565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b831061100a5772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611036576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061105457662386f26fc10000830492506010015b6305f5e100831061106c576305f5e100830492506008015b612710831061108057612710830492506004015b60648310611092576064830492506002015b600a83106105655760010192915050565b6001600160a01b0382166110f95760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604482015260640161045e565b6000818152600260205260409020546001600160a01b03161561115e5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161045e565b61116c600083836001610e0f565b6000818152600260205260409020546001600160a01b0316156111d15760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161045e565b6001600160a01b038216600081815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6001600160e01b0319811681146105b657600080fd5b60006020828403121561126457600080fd5b81356107d88161123c565b60005b8381101561128a578181015183820152602001611272565b50506000910152565b600081518084526112ab81602086016020860161126f565b601f01601f19169290920160200192915050565b6020815260006107d86020830184611293565b6000602082840312156112e457600080fd5b5035919050565b80356001600160a01b038116811461056857600080fd5b6000806040838503121561131557600080fd5b61131e836112eb565b946020939093013593505050565b60008060006060848603121561134157600080fd5b61134a846112eb565b9250611358602085016112eb565b9150604084013590509250925092565b60006020828403121561137a57600080fd5b6107d8826112eb565b6000806040838503121561139657600080fd5b61139f836112eb565b9150602083013580151581146113b457600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600080600080608085870312156113eb57600080fd5b6113f4856112eb565b9350611402602086016112eb565b925060408501359150606085013567ffffffffffffffff8082111561142657600080fd5b818701915087601f83011261143a57600080fd5b81358181111561144c5761144c6113bf565b604051601f8201601f19908116603f01168101908382118183101715611474576114746113bf565b816040528281528a602084870101111561148d57600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b600080604083850312156114c457600080fd5b6114cd836112eb565b91506114db602084016112eb565b90509250929050565b600181811c908216806114f857607f821691505b60208210810361151857634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b6000600182016115935761159361156b565b5060010190565b600083516115ac81846020880161126f565b8351908301906115c081836020880161126f565b01949350505050565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b8181038181111561032b5761032b61156b565b8082018082111561032b5761032b61156b565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906116b990830184611293565b9695505050505050565b6000602082840312156116d557600080fd5b81516107d88161123c56fea26469706673582212203d402de792a5d67ca1133a5c16a19e1f51b0821c6eebc7f040672a1776e9a20664736f6c63430008110033",
}

// RoomNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use RoomNFTMetaData.ABI instead.
var RoomNFTABI = RoomNFTMetaData.ABI

// RoomNFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoomNFTMetaData.Bin instead.
var RoomNFTBin = RoomNFTMetaData.Bin

// DeployRoomNFT deploys a new Ethereum contract, binding an instance of RoomNFT to it.
func DeployRoomNFT(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RoomNFT, error) {
	parsed, err := RoomNFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoomNFTBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RoomNFT{RoomNFTCaller: RoomNFTCaller{contract: contract}, RoomNFTTransactor: RoomNFTTransactor{contract: contract}, RoomNFTFilterer: RoomNFTFilterer{contract: contract}}, nil
}

// RoomNFT is an auto generated Go binding around an Ethereum contract.
type RoomNFT struct {
	RoomNFTCaller     // Read-only binding to the contract
	RoomNFTTransactor // Write-only binding to the contract
	RoomNFTFilterer   // Log filterer for contract events
}

// RoomNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoomNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoomNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoomNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoomNFTSession struct {
	Contract     *RoomNFT          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoomNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoomNFTCallerSession struct {
	Contract *RoomNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RoomNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoomNFTTransactorSession struct {
	Contract     *RoomNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RoomNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoomNFTRaw struct {
	Contract *RoomNFT // Generic contract binding to access the raw methods on
}

// RoomNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoomNFTCallerRaw struct {
	Contract *RoomNFTCaller // Generic read-only contract binding to access the raw methods on
}

// RoomNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoomNFTTransactorRaw struct {
	Contract *RoomNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoomNFT creates a new instance of RoomNFT, bound to a specific deployed contract.
func NewRoomNFT(address common.Address, backend bind.ContractBackend) (*RoomNFT, error) {
	contract, err := bindRoomNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoomNFT{RoomNFTCaller: RoomNFTCaller{contract: contract}, RoomNFTTransactor: RoomNFTTransactor{contract: contract}, RoomNFTFilterer: RoomNFTFilterer{contract: contract}}, nil
}

// NewRoomNFTCaller creates a new read-only instance of RoomNFT, bound to a specific deployed contract.
func NewRoomNFTCaller(address common.Address, caller bind.ContractCaller) (*RoomNFTCaller, error) {
	contract, err := bindRoomNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoomNFTCaller{contract: contract}, nil
}

// NewRoomNFTTransactor creates a new write-only instance of RoomNFT, bound to a specific deployed contract.
func NewRoomNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*RoomNFTTransactor, error) {
	contract, err := bindRoomNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoomNFTTransactor{contract: contract}, nil
}

// NewRoomNFTFilterer creates a new log filterer instance of RoomNFT, bound to a specific deployed contract.
func NewRoomNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*RoomNFTFilterer, error) {
	contract, err := bindRoomNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoomNFTFilterer{contract: contract}, nil
}

// bindRoomNFT binds a generic wrapper to an already deployed contract.
func bindRoomNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoomNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoomNFT *RoomNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoomNFT.Contract.RoomNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoomNFT *RoomNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomNFT.Contract.RoomNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoomNFT *RoomNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoomNFT.Contract.RoomNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoomNFT *RoomNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoomNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoomNFT *RoomNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoomNFT *RoomNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoomNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RoomNFT *RoomNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RoomNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RoomNFT *RoomNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RoomNFT.Contract.BalanceOf(&_RoomNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RoomNFT *RoomNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RoomNFT.Contract.BalanceOf(&_RoomNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RoomNFT *RoomNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RoomNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RoomNFT *RoomNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _RoomNFT.Contract.GetApproved(&_RoomNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RoomNFT *RoomNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _RoomNFT.Contract.GetApproved(&_RoomNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RoomNFT *RoomNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _RoomNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RoomNFT *RoomNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _RoomNFT.Contract.IsApprovedForAll(&_RoomNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RoomNFT *RoomNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _RoomNFT.Contract.IsApprovedForAll(&_RoomNFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RoomNFT *RoomNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RoomNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RoomNFT *RoomNFTSession) Name() (string, error) {
	return _RoomNFT.Contract.Name(&_RoomNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RoomNFT *RoomNFTCallerSession) Name() (string, error) {
	return _RoomNFT.Contract.Name(&_RoomNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoomNFT *RoomNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoomNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoomNFT *RoomNFTSession) Owner() (common.Address, error) {
	return _RoomNFT.Contract.Owner(&_RoomNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoomNFT *RoomNFTCallerSession) Owner() (common.Address, error) {
	return _RoomNFT.Contract.Owner(&_RoomNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RoomNFT *RoomNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RoomNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RoomNFT *RoomNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _RoomNFT.Contract.OwnerOf(&_RoomNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RoomNFT *RoomNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _RoomNFT.Contract.OwnerOf(&_RoomNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RoomNFT *RoomNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RoomNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RoomNFT *RoomNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RoomNFT.Contract.SupportsInterface(&_RoomNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RoomNFT *RoomNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RoomNFT.Contract.SupportsInterface(&_RoomNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RoomNFT *RoomNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RoomNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RoomNFT *RoomNFTSession) Symbol() (string, error) {
	return _RoomNFT.Contract.Symbol(&_RoomNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RoomNFT *RoomNFTCallerSession) Symbol() (string, error) {
	return _RoomNFT.Contract.Symbol(&_RoomNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RoomNFT *RoomNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _RoomNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RoomNFT *RoomNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _RoomNFT.Contract.TokenURI(&_RoomNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RoomNFT *RoomNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _RoomNFT.Contract.TokenURI(&_RoomNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RoomNFT *RoomNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RoomNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RoomNFT *RoomNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RoomNFT.Contract.Approve(&_RoomNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RoomNFT *RoomNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RoomNFT.Contract.Approve(&_RoomNFT.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_RoomNFT *RoomNFTTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _RoomNFT.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_RoomNFT *RoomNFTSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _RoomNFT.Contract.Burn(&_RoomNFT.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_RoomNFT *RoomNFTTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _RoomNFT.Contract.Burn(&_RoomNFT.TransactOpts, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoomNFT *RoomNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoomNFT *RoomNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _RoomNFT.Contract.RenounceOwnership(&_RoomNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoomNFT *RoomNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RoomNFT.Contract.RenounceOwnership(&_RoomNFT.TransactOpts)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns(uint256)
func (_RoomNFT *RoomNFTTransactor) SafeMint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RoomNFT.contract.Transact(opts, "safeMint", to)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns(uint256)
func (_RoomNFT *RoomNFTSession) SafeMint(to common.Address) (*types.Transaction, error) {
	return _RoomNFT.Contract.SafeMint(&_RoomNFT.TransactOpts, to)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns(uint256)
func (_RoomNFT *RoomNFTTransactorSession) SafeMint(to common.Address) (*types.Transaction, error) {
	return _RoomNFT.Contract.SafeMint(&_RoomNFT.TransactOpts, to)
}

// SafeMintBatch is a paid mutator transaction binding the contract method 0x4bc96248.
//
// Solidity: function safeMintBatch(address to, uint256 amount) returns(uint256)
func (_RoomNFT *RoomNFTTransactor) SafeMintBatch(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RoomNFT.contract.Transact(opts, "safeMintBatch", to, amount)
}

// SafeMintBatch is a paid mutator transaction binding the contract method 0x4bc96248.
//
// Solidity: function safeMintBatch(address to, uint256 amount) returns(uint256)
func (_RoomNFT *RoomNFTSession) SafeMintBatch(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RoomNFT.Contract.SafeMintBatch(&_RoomNFT.TransactOpts, to, amount)
}

// SafeMintBatch is a paid mutator transaction binding the contract method 0x4bc96248.
//
// Solidity: function safeMintBatch(address to, uint256 amount) returns(uint256)
func (_RoomNFT *RoomNFTTransactorSession) SafeMintBatch(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RoomNFT.Contract.SafeMintBatch(&_RoomNFT.TransactOpts, to, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RoomNFT *RoomNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RoomNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RoomNFT *RoomNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RoomNFT.Contract.SafeTransferFrom(&_RoomNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RoomNFT *RoomNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RoomNFT.Contract.SafeTransferFrom(&_RoomNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_RoomNFT *RoomNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _RoomNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_RoomNFT *RoomNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _RoomNFT.Contract.SafeTransferFrom0(&_RoomNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_RoomNFT *RoomNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _RoomNFT.Contract.SafeTransferFrom0(&_RoomNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RoomNFT *RoomNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _RoomNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RoomNFT *RoomNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _RoomNFT.Contract.SetApprovalForAll(&_RoomNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RoomNFT *RoomNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _RoomNFT.Contract.SetApprovalForAll(&_RoomNFT.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RoomNFT *RoomNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RoomNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RoomNFT *RoomNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RoomNFT.Contract.TransferFrom(&_RoomNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RoomNFT *RoomNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RoomNFT.Contract.TransferFrom(&_RoomNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RoomNFT *RoomNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RoomNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RoomNFT *RoomNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RoomNFT.Contract.TransferOwnership(&_RoomNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RoomNFT *RoomNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RoomNFT.Contract.TransferOwnership(&_RoomNFT.TransactOpts, newOwner)
}

// RoomNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RoomNFT contract.
type RoomNFTApprovalIterator struct {
	Event *RoomNFTApproval // Event containing the contract specifics and raw log

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
func (it *RoomNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomNFTApproval)
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
		it.Event = new(RoomNFTApproval)
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
func (it *RoomNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomNFTApproval represents a Approval event raised by the RoomNFT contract.
type RoomNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RoomNFT *RoomNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*RoomNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RoomNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RoomNFTApprovalIterator{contract: _RoomNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RoomNFT *RoomNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RoomNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RoomNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomNFTApproval)
				if err := _RoomNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RoomNFT *RoomNFTFilterer) ParseApproval(log types.Log) (*RoomNFTApproval, error) {
	event := new(RoomNFTApproval)
	if err := _RoomNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the RoomNFT contract.
type RoomNFTApprovalForAllIterator struct {
	Event *RoomNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *RoomNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomNFTApprovalForAll)
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
		it.Event = new(RoomNFTApprovalForAll)
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
func (it *RoomNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomNFTApprovalForAll represents a ApprovalForAll event raised by the RoomNFT contract.
type RoomNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RoomNFT *RoomNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*RoomNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RoomNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RoomNFTApprovalForAllIterator{contract: _RoomNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RoomNFT *RoomNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *RoomNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RoomNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomNFTApprovalForAll)
				if err := _RoomNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RoomNFT *RoomNFTFilterer) ParseApprovalForAll(log types.Log) (*RoomNFTApprovalForAll, error) {
	event := new(RoomNFTApprovalForAll)
	if err := _RoomNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RoomNFT contract.
type RoomNFTOwnershipTransferredIterator struct {
	Event *RoomNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RoomNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomNFTOwnershipTransferred)
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
		it.Event = new(RoomNFTOwnershipTransferred)
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
func (it *RoomNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomNFTOwnershipTransferred represents a OwnershipTransferred event raised by the RoomNFT contract.
type RoomNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoomNFT *RoomNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RoomNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RoomNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RoomNFTOwnershipTransferredIterator{contract: _RoomNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoomNFT *RoomNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RoomNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RoomNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomNFTOwnershipTransferred)
				if err := _RoomNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoomNFT *RoomNFTFilterer) ParseOwnershipTransferred(log types.Log) (*RoomNFTOwnershipTransferred, error) {
	event := new(RoomNFTOwnershipTransferred)
	if err := _RoomNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RoomNFT contract.
type RoomNFTTransferIterator struct {
	Event *RoomNFTTransfer // Event containing the contract specifics and raw log

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
func (it *RoomNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomNFTTransfer)
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
		it.Event = new(RoomNFTTransfer)
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
func (it *RoomNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomNFTTransfer represents a Transfer event raised by the RoomNFT contract.
type RoomNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RoomNFT *RoomNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*RoomNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RoomNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RoomNFTTransferIterator{contract: _RoomNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RoomNFT *RoomNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RoomNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RoomNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomNFTTransfer)
				if err := _RoomNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RoomNFT *RoomNFTFilterer) ParseTransfer(log types.Log) (*RoomNFTTransfer, error) {
	event := new(RoomNFTTransfer)
	if err := _RoomNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

package res

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func AddressToCommonAddress(address string) common.Address {
	return common.HexToAddress(address)
}

func StringToBigInt(number string) *big.Int {
	n := new(big.Int)
	bigIntNumber, _ := n.SetString(number, 10)
	return bigIntNumber
}

func Float32ToBigInt(number float32) *big.Int {
	n := new(big.Int)
	bigIntNumber := n.SetInt64(int64(number))
	return bigIntNumber
}

func IntToBigInt(number int) *big.Int {
	n := new(big.Int)
	bigIntNumber := n.SetInt64(int64(number))
	return bigIntNumber

}

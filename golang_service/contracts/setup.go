package contracts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnectHardHatNode() *ethclient.Client {
	client, err := ethclient.Dial(os.Getenv("LOCAL_HARDHAT_URL"))
	if err != nil {
		panic(err)
	}
	return client
}

func GetHotelContract() *Hotel {

	conn, err := NewHotel(common.HexToAddress(os.Getenv("DEPLOYED_HOTEL_CONTRACT_ADDRESS")), ConnectHardHatNode())
	if err != nil {
		panic(err)
	}
	return conn
}

func GetRoomNFTContract() *RoomNFT {
	conn, err := NewRoomNFT(common.HexToAddress(os.Getenv("DEPLOYED_ROOMNFT_CONTRACT_ADDRESS")), ConnectHardHatNode())
	if err != nil {
		panic(err)
	}
	return conn
}

// Only used for consumed gas function
// view function do not consume gas so no need to sign by PrivateKey
func SignTransaction(privateKeyAddress string, value uint64) (*bind.TransactOpts, error) {

	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println(fromAddress)
	nonce, err := ConnectHardHatNode().PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("endingNonce error: %v", err)
	}
	fmt.Println("nonce=", nonce)
	chainID, err := ConnectHardHatNode().ChainID(context.Background())
	if err != nil {
		return nil, errors.New("chainID error")
	}
	fmt.Println("ChainID", chainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, errors.New("newKeyedTransactorWithChainID error")
	}
	auth.Nonce = big.NewInt(int64(nonce))
	if value != 0 {
		auth.Value = big.NewInt(int64(value)) // in wei
	} else {
		auth.Value = big.NewInt(0) // in wei
	}
	auth.GasLimit = uint64(0) // in units
	auth.GasPrice, err = ConnectHardHatNode().SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.New("gas price setup error")
	}
	return auth, nil
}

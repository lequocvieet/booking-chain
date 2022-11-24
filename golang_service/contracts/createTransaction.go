package contracts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"golang_service/wrap_contract"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetHotelContract(privateKeyAddress string, value uint64) (*wrap_contract.Hotel, *bind.TransactOpts) {
	// address of etherum testnet node
	local_hardhat := "ws://127.0.0.1:8545"
	client, err := ethclient.Dial(local_hardhat)
	if err != nil {
		panic(err)
	}
	deployed_hotel_contract_address := os.Getenv("DEPLOYED_HOTEL_CONTRACT_ADDRESS")
	conn, err := wrap_contract.NewHotel(common.HexToAddress(deployed_hotel_contract_address), client)
	if err != nil {
		panic(err)
	}

	var auth *bind.TransactOpts
	if privateKeyAddress != "" {
		auth, err = GetAccountAuth(client, privateKeyAddress, value)
		if err != nil {
			panic(err)
		}
	}

	return conn, auth
}

func GetRoomNFTContract(privateKeyAddress string, value uint64) (*wrap_contract.RoomNFT, *bind.TransactOpts) {
	// address of etherum testnet node
	local_hardhat := "ws://127.0.0.1:8545"
	client, err := ethclient.Dial(local_hardhat)
	if err != nil {
		panic(err)
	}
	deployed_room_nft_contract_address := os.Getenv("DEPLOYED_ROOMNFT_CONTRACT_ADDRESS")
	conn, err := wrap_contract.NewRoomNFT(common.HexToAddress(deployed_room_nft_contract_address), client)
	if err != nil {
		panic(err)
	}

	var auth *bind.TransactOpts
	if privateKeyAddress != "" {
		auth, err = GetAccountAuth(client, privateKeyAddress, value)
		if err != nil {
			panic(err)
		}
	}

	return conn, auth
}

func GetAccountAuth(client *ethclient.Client, privateKeyAddress string, value uint64) (*bind.TransactOpts, error) {

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
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("endingNonce error: %v", err)
	}
	fmt.Println("nonce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, errors.New("ChainID error!")
	}
	fmt.Println("ChainID", chainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, errors.New("NewKeyedTransactorWithChainID error!")
	}
	auth.Nonce = big.NewInt(int64(nonce))
	if value != 0 {
		auth.Value = big.NewInt(int64(value)) // in wei
	} else {
		auth.Value = big.NewInt(0) // in wei
	}
	auth.GasLimit = uint64(0) // in units
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	return auth, nil
}

func ListenCreateListRoomEvent(landLordAddress common.Address) *wrap_contract.HotelCreateListRoom {

	// address of etherum local node
	local_hardhat := "ws://127.0.0.1:8545"
	client, err := ethclient.Dial(local_hardhat)
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	deployed_hotel_contract_address := os.Getenv("DEPLOYED_HOTEL_CONTRACT_ADDRESS")
	conn, err := wrap_contract.NewHotel(common.HexToAddress(deployed_hotel_contract_address), client)
	if err != nil {
		panic(err)
	}
	var landLordAddresses []common.Address
	landLordAddresses = append(landLordAddresses, landLordAddress)

	// Filter event
	filterOpts := &bind.FilterOpts{Context: context.Background(), Start: 0, End: nil}
	itr, err := conn.FilterCreateListRoom(filterOpts, landLordAddresses)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var events []*wrap_contract.HotelCreateListRoom
	//Loop over all found events
	for itr.Next() {
		event := itr.Event
		events = append(events, event)
	}
	return events[len(events)-1]
}

func ListenBookRoomEvent(userAddress common.Address) *wrap_contract.HotelBookRoom {

	// address of etherum local node
	local_hardhat := "ws://127.0.0.1:8545"
	client, err := ethclient.Dial(local_hardhat)
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	deployed_hotel_contract_address := os.Getenv("DEPLOYED_HOTEL_CONTRACT_ADDRESS")
	conn, err := wrap_contract.NewHotel(common.HexToAddress(deployed_hotel_contract_address), client)
	if err != nil {
		panic(err)
	}

	var userAddresses []common.Address
	userAddresses = append(userAddresses, userAddress)
	// Filter event
	filterOpts := &bind.FilterOpts{Context: context.Background(), Start: 0, End: nil}
	itr, err := conn.FilterBookRoom(filterOpts, userAddresses)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var events []*wrap_contract.HotelBookRoom
	//Loop over all found events
	for itr.Next() {
		event := itr.Event
		events = append(events, event)
	}
	return events[len(events)-1]
}

func ListenCancelBookRoomEvent(userAddress common.Address) *wrap_contract.HotelCancelBookRoom {

	// address of etherum local node
	local_hardhat := "ws://127.0.0.1:8545"
	client, err := ethclient.Dial(local_hardhat)
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	deployed_hotel_contract_address := os.Getenv("DEPLOYED_HOTEL_CONTRACT_ADDRESS")
	conn, err := wrap_contract.NewHotel(common.HexToAddress(deployed_hotel_contract_address), client)
	if err != nil {
		panic(err)
	}
	var userAddresses []common.Address
	userAddresses = append(userAddresses, userAddress)
	// Filter event
	filterOpts := &bind.FilterOpts{Context: context.Background(), Start: 0, End: nil}
	itr, err := conn.FilterCancelBookRoom(filterOpts, userAddresses)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var events []*wrap_contract.HotelCancelBookRoom
	//Loop over all found events
	for itr.Next() {
		event := itr.Event
		events = append(events, event)
	}
	return events[len(events)-1]
}

func ListenCheckInEvent(checker common.Address) *wrap_contract.HotelCheckIn {

	// address of etherum local node
	local_hardhat := "ws://127.0.0.1:8545"
	client, err := ethclient.Dial(local_hardhat)
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	deployed_hotel_contract_address := os.Getenv("DEPLOYED_HOTEL_CONTRACT_ADDRESS")
	conn, err := wrap_contract.NewHotel(common.HexToAddress(deployed_hotel_contract_address), client)
	if err != nil {
		panic(err)
	}
	var checkerAddresses []common.Address
	checkerAddresses = append(checkerAddresses, checker)
	// Filter event
	filterOpts := &bind.FilterOpts{Context: context.Background(), Start: 0, End: nil}
	itr, err := conn.FilterCheckIn(filterOpts, checkerAddresses)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var events []*wrap_contract.HotelCheckIn
	//Loop over all found events
	for itr.Next() {
		event := itr.Event
		events = append(events, event)
	}
	return events[len(events)-1]
}

func CaculateAllBookRoomEventByTokenID(tokenId int) int {
	// address of etherum local node
	local_hardhat := "ws://127.0.0.1:8545"
	client, err := ethclient.Dial(local_hardhat)
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	deployed_hotel_contract_address := os.Getenv("DEPLOYED_HOTEL_CONTRACT_ADDRESS")
	conn, err := wrap_contract.NewHotel(common.HexToAddress(deployed_hotel_contract_address), client)
	if err != nil {
		panic(err)
	}
	var userAddresses []common.Address
	// Filter event
	filterOpts := &bind.FilterOpts{Context: context.Background(), Start: 0, End: nil}
	itr, err := conn.FilterBookRoom(filterOpts, userAddresses)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//Loop over all found events
	var count int = 0
	for itr.Next() {
		event := itr.Event
		numberOfdates := event.NumberOfdates
		startTokenId := int(event.StartTokenId.Uint64())
		for i := 0; i < int(numberOfdates.Uint64()); i++ {
			if startTokenId == tokenId {
				count++
			}
			startTokenId++
		}

	}
	return count
}

func CaculateAllCancelBookRoomEventByTokenID(tokenId int) int {
	// address of etherum local node
	local_hardhat := "ws://127.0.0.1:8545"
	client, err := ethclient.Dial(local_hardhat)
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	deployed_hotel_contract_address := os.Getenv("DEPLOYED_HOTEL_CONTRACT_ADDRESS")
	conn, err := wrap_contract.NewHotel(common.HexToAddress(deployed_hotel_contract_address), client)
	if err != nil {
		panic(err)
	}
	var userAddresses []common.Address
	// Filter event
	filterOpts := &bind.FilterOpts{Context: context.Background(), Start: 0, End: nil}
	itr, err := conn.FilterCancelBookRoom(filterOpts, userAddresses)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//Loop over all found events
	var count int = 0
	for itr.Next() {
		event := itr.Event
		for i := 0; i < len(event.TokenIds); i++ {
			if int(event.TokenIds[i].Uint64()) == tokenId {
				count++
			}
		}

	}
	return count
}

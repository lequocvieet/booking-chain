package contracts

import (
	"context"
	"fmt"
	"golang_service/wrap_contract"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

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

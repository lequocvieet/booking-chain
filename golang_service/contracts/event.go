package contracts

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var filterOpts = &bind.FilterOpts{Context: context.Background(), Start: 0, End: nil}

func ListenEvent(addressIndexed common.Address, eventType string) interface{} {
	conn := GetHotelContract()
	var addressIndexeds []common.Address
	addressIndexeds = append(addressIndexeds, addressIndexed)
	switch eventType {
	case "CreateListRoom":
		itr, err := conn.FilterCreateListRoom(filterOpts, addressIndexeds)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		var events []*HotelCreateListRoom
		//Loop over all found events
		for itr.Next() {
			event := itr.Event
			events = append(events, event)
		}
		return events[len(events)-1]

	case "BookRoom":
		itr, err := conn.FilterBookRoom(filterOpts, addressIndexeds)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		var events []*HotelBookRoom
		//Loop over all found events
		for itr.Next() {
			event := itr.Event
			events = append(events, event)
		}
		return events[len(events)-1]

	case "CancelBookRoom":
		itr, err := conn.FilterCancelBookRoom(filterOpts, addressIndexeds)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		var events []*HotelCancelBookRoom
		//Loop over all found events
		for itr.Next() {
			event := itr.Event
			events = append(events, event)
		}
		return events[len(events)-1]

	case "CheckIn":
		itr, err := conn.FilterCheckIn(filterOpts, addressIndexeds)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		var events []*HotelCheckIn
		//Loop over all found events
		for itr.Next() {
			event := itr.Event
			events = append(events, event)
		}
		return events[len(events)-1]
	}
	return nil

}

func CaculateAllBookRoomEventByTokenID(tokenId int) int {
	// address of etherum local node
	conn := GetHotelContract()
	var userAddresses []common.Address
	// Filter event
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
	conn := GetHotelContract()
	var userAddresses []common.Address
	// Filter event
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

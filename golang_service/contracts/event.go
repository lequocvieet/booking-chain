package contracts

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func ListenEvent(addressIndexed common.Address, eventType string) interface{} {
	conn := GetHotelContract()
	var addressIndexeds []common.Address
	addressIndexeds = append(addressIndexeds, addressIndexed)
	filterOpts := &bind.FilterOpts{Context: context.Background(), Start: 0, End: nil}

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



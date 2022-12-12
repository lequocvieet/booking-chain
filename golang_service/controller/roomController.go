package controller

import (
	"encoding/json"
	"fmt"
	jwt "golang_service/auth"
	"golang_service/contracts"
	"golang_service/models"
	res "golang_service/utils"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type roomController struct {
	model *models.Model
}

var roomIdCount int = 1

const dateLayout string = "2006-01-02"

func NewRoomController(model *models.Model) roomController {
	return roomController{model}
}

func (c roomController) CreateListRoom(w http.ResponseWriter, r *http.Request) {
	landLordName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	landLord, err := c.model.FindUserByUserName(landLordName)
	if landLord.Role != "landlord" {
		res.JSON(w, 400, "you must be landlord to execute this function")
		return
	}
	//create transaction on blockchain
	auth, err := contracts.SignTransaction(landLord.PrivateKey, 0)
	if err != nil {
		res.JSON(w, 201, err)
		return
	}
	_, err = hotelContract.CreateListRoom(auth)
	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "create list room has fail")
		return
	}

	//Save listroom to backend after listen createListRoomEvent
	createListRoomEvent := contracts.ListenEvent(common.HexToAddress(landLord.Address), "CreateListRoom").(*contracts.HotelCreateListRoom)
	listRoom := models.NewListRoom(int(createListRoomEvent.ListRoomId.Uint64()), landLord.ID, res.EpochTimeToDate(createListRoomEvent.Timestamp))
	c.model.SaveListRoom(listRoom)
	res.JSON(w, 201, "create list room success")
	return

}

func (c roomController) DeleteListRoom(w http.ResponseWriter, r *http.Request) {
	landLordName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	landLord, err := c.model.FindUserByUserName(landLordName)
	if err != nil {
		res.JSON(w, 400, err)
		return
	}
	if landLord.Role != "landlord" {
		res.JSON(w, 400, "you must be landlord to execute this function")
		return
	}

	listRoomId := r.PostFormValue("listRoomId")
	//create transaction on blockchain
	auth, err := contracts.SignTransaction(landLord.PrivateKey, 0)
	if err != nil {
		res.JSON(w, 201, err)
		return
	}
	n := new(big.Int)
	listRoomIds, _ := n.SetString(listRoomId, 10)
	_, err = hotelContract.DeleteListRoom(auth, listRoomIds)
	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "delete list room has fail")
		return
	}
	// Delete list room off chain
	id, _ := strconv.Atoi(listRoomId)
	errDeleteListRoom := c.model.DeleteListRoom(id)
	if errDeleteListRoom != nil {
		res.JSON(w, 500, errDeleteListRoom)
		return
	}
	res.JSON(w, 201, "delete list room success")
	return

}

func (c roomController) CreateRoom(w http.ResponseWriter, r *http.Request) {
	landLordName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	landLord, err := c.model.FindUserByUserName(landLordName)
	if err != nil {
		res.JSON(w, 400, err)
		return
	}
	if landLord.Role != "landlord" {
		res.JSON(w, 400, "You must be landlord to execute this function!")
		return
	}
	listRoomId := r.PostFormValue("listRoomId")
	pricePerDay := r.PostFormValue("pricePerDay")
	id, _ := strconv.Atoi(listRoomId)
	_, errFindListRoom := c.model.FindListRoom(id)
	if errFindListRoom != nil {
		res.JSON(w, 400, errFindListRoom)
		return
	}
	value, _ := strconv.ParseFloat(pricePerDay, 32)
	room := models.NewRoom(roomIdCount, float32(value), id)
	c.model.SaveRoom(room) // save room off-chain

	//create default state
	id, _ = strconv.Atoi(listRoomId)
	stateRoom := models.NewStateRoom(roomIdCount, time.Time{}, time.Time{}, id, "Available")
	roomIdCount++
	c.model.SaveStateRoom(stateRoom) // save state of new room created
	res.JSON(w, 201, "create room successful")
	return

}

func (c roomController) UpdateRoomState(w http.ResponseWriter, r *http.Request) {
	landLordName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	landLord, err := c.model.FindUserByUserName(landLordName)
	if err != nil {
		res.JSON(w, 400, err)
		return
	}
	if landLord.Role != "landlord" {
		res.JSON(w, 400, "You must be landlord to execute this function!")
		return
	}
	stateRoomId, _ := strconv.Atoi(r.PostFormValue("stateRoomId"))
	listRoomId, _ := strconv.Atoi(r.PostFormValue("listRoomId"))
	roomId, _ := strconv.Atoi(r.PostFormValue("listRoomId"))
	startUpdate := res.FormatTime(r.PostFormValue("startUpdate"))
	endUpdate := res.FormatTime(r.PostFormValue("endUpdate"))
	state := r.PostFormValue("state")

	//verify list room id exist
	_, errFindListRoom := c.model.FindListRoom(listRoomId)
	if errFindListRoom != nil {
		res.JSON(w, 400, errFindListRoom)
		return
	}
	_, errFindRoom := c.model.FindRoomByID(roomId)
	if errFindRoom != nil {
		res.JSON(w, 400, errFindRoom)
		return
	}
	//verify update room state in the past
	if endUpdate.Before(time.Now()) {
		res.JSON(w, 400, "can not update state room in the past")
		return
	}

	//Check at that time room is already used
	var isUsed bool = true
	roomNFTs := c.model.FindRoomNFTByRoomIdAndDateValid(roomId, startUpdate, endUpdate)
	if len(roomNFTs) == 0 {
		isUsed = false
	}
	//Check at that time room available or not
	var available bool = true
	stateRooms := c.model.FindStateRoom(roomId, listRoomId, startUpdate, endUpdate)
	if len(stateRooms) != 0 {
		available = false
	}
	fmt.Print("available ", available)

	if isUsed == false && available == true {
		//ready to update state
		stateRoom := c.model.FindStateRoomById(stateRoomId)
		stateRoom.State = state
		stateRoom.StartUpdate = startUpdate
		stateRoom.EndUpdate = endUpdate
		c.model.SaveStateRoom(stateRoom) // save state room off-chain
		res.JSON(w, 201, "update state room successful")
		return
	}
	res.JSON(w, 400, "room already used or not available yet")
	return
}

func (c roomController) BookRoom(w http.ResponseWriter, r *http.Request) {
	userName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	user, err := c.model.FindUserByUserName(userName)
	if err != nil {
		res.JSON(w, 400, err)
		return
	}
	listRoomId, _ := strconv.Atoi(r.PostFormValue("listRoomId"))
	roomId, _ := strconv.Atoi(r.PostFormValue("listRoomId"))
	startDay := res.FormatTime(r.PostFormValue("startDay"))
	startDay = time.Date(startDay.Year(), startDay.Month(), startDay.Day(), 7, 0, 0, 0, startDay.Location())
	endDay := res.FormatTime(r.PostFormValue("endDay"))
	endDay = time.Date(endDay.Year(), endDay.Month(), endDay.Day(), 7, 0, 0, 0, endDay.Location())

	//verify list room id exist
	_, errFindListRoom := c.model.FindListRoom(listRoomId)
	if errFindListRoom != nil {
		res.JSON(w, 400, errFindListRoom)
		return
	}
	//verify that room id  exist
	room, errFindRoom := c.model.FindRoomByID(roomId)
	if errFindRoom != nil {
		res.JSON(w, 400, errFindRoom)
		return
	}

	//Todo check day format

	//verify book room in the past
	if startDay.Before(time.Now()) || endDay.Before(time.Now()) {
		res.JSON(w, 400, "can not book room in the past")
		return
	}

	//verify all room state in range book is availale
	var available bool = true
	stateRooms := c.model.FindStateRoomByDay(roomId, listRoomId, startDay, endDay)
	if len(stateRooms) != 0 {
		available = false
	}
	fmt.Print("available ", available)
	//verify all room in range book is not booked before
	var isUsed bool = true
	roomNFTs := c.model.FindRoomNFTByRoomIdAndDateValid(roomId, startDay, endDay)
	if len(roomNFTs) == 0 {
		isUsed = false
	}
	fmt.Print(" isUsed ", isUsed)
	if isUsed == false && available == true {
		//ready to book
		//Caculate total price

		numberDaysBook := res.SubtractTime(endDay, startDay)
		totalPrice := room.PricePerDay * float32(numberDaysBook)

		//create transaction on blockchain
		auth, err := contracts.SignTransaction(user.PrivateKey, uint64(totalPrice))
		if err != nil {
			res.JSON(w, 201, err)
			return
		}
		_, errBookRoom := hotelContract.BookRoom(auth, res.Float32ToBigInt(totalPrice), res.IntToBigInt(numberDaysBook))
		if errBookRoom != nil {
			res.JSON(w, 400, errBookRoom)
			return
		}

		// Listen bookroom Event save to RoomNFT table
		bookRoomEvent := contracts.ListenEvent(common.HexToAddress(user.Address), "BookRoom").(*contracts.HotelBookRoom)
		var tokenID int = int(bookRoomEvent.StartTokenId.Uint64())
		for i := 0; i < int(bookRoomEvent.NumberOfdates.Uint64()); i++ {
			roomNFT := models.NewRoomNFT(roomId, tokenID, bookRoomEvent.Booker.String(), startDay)
			startDay = startDay.AddDate(0, 0, 1) // increase 1 days for later RoomNFT in for loop
			tokenID++
			c.model.SaveRoomNFT(roomNFT)
		}
		res.JSON(w, 201, "book room success")
		return
	}
	res.JSON(w, 400, "room already booked or not available yet")
	return

}

func (c roomController) CancelBookRoom(w http.ResponseWriter, r *http.Request) {
	userName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	user, err := c.model.FindUserByUserName(userName)
	if err != nil {
		res.JSON(w, 400, err)
		return
	}
	var roomNFTId models.RoomNFTIds
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res.JSON(w, 400, "Read Body error!")
		return
	}
	err = json.Unmarshal(body, &roomNFTId)
	if err != nil {
		fmt.Print(err)
		res.JSON(w, 400, "Unmarshal error!")
		return
	}
	var roomNFTs []models.RoomNFT
	//verify that roomNFT id  exist
	roomNFTIds := roomNFTId.RoomNFTIds
	for i := 0; i < len(roomNFTIds); i++ {
		roomNFT, errFindRoomNFT := c.model.FindRoomNFTByID(roomNFTIds[i])
		if errFindRoomNFT != nil {
			res.JSON(w, 400, errFindRoomNFT)
			return
		}
		roomNFTs = append(roomNFTs, roomNFT)
	}

	//Check valid cancel happens before check in time
	//check in time hard code is 14pm - 15pm
	currentTime := time.Now()
	currentTime.Format(dateLayout)
	for i := 0; i < len(roomNFTs); i++ {
		if currentTime.After(roomNFTs[i].DateValid) {
			res.JSON(w, 400, "Cannot cancel after checkin!")
			return
		}

	}

	//Call ownerOf view function to verify that caller owns that tokenids
	var opts *bind.CallOpts
	var tokenIds []*big.Int
	for i := 0; i < len(roomNFTs); i++ {
		ownerAddress, err := hotelContract.GetOwnerOf(opts, res.IntToBigInt(roomNFTs[i].TokenId))
		if err != nil {
			fmt.Println(err)
			return
		}
		if user.Address != ownerAddress.Hex() {
			res.JSON(w, 400, "you not own these NFT to cancel")
			return
		}
		tokenIds = append(tokenIds, res.IntToBigInt(roomNFTs[i].TokenId))

	}
	//Find room by room id Caculate repay price
	var rooms []models.Room
	var price float32
	for i := 0; i < len(roomNFTs); i++ {
		var room models.Room
		room, errFindRoom := c.model.FindRoomByID(roomNFTs[i].RoomId)
		if errFindRoom != nil {
			res.JSON(w, 400, errFindRoom)
			return
		}
		rooms = append(rooms, room)
		price += room.PricePerDay

	}
	//fee cancel hard save in .env
	feeCancel, err := strconv.ParseFloat(os.Getenv("FEE_CANCEL"), 32)
	totalRepayPrice := price * float32(feeCancel)

	// nft.approve to approve burn permission to hotel contract
	authNFT, err := contracts.SignTransaction(user.PrivateKey, 0)
	if err != nil {
		res.JSON(w, 201, err)
		return
	}
	contractAddress := common.HexToAddress(os.Getenv("DEPLOYED_HOTEL_CONTRACT_ADDRESS"))
	_, errSetApproval := roomNFTCotract.SetApprovalForAll(authNFT, contractAddress, true)
	if errSetApproval != nil {
		res.JSON(w, 201, errSetApproval)
		return

	}

	//create transaction on blockchain
	auth, err := contracts.SignTransaction(user.PrivateKey, 0)
	if err != nil {
		res.JSON(w, 201, err)
		return
	}
	_, errCancelBook := hotelContract.CancelBookRoom(auth, tokenIds, res.Float32ToBigInt(totalRepayPrice))
	if errCancelBook != nil {
		res.JSON(w, 201, errCancelBook)
		return
	}
	//Delete in RoomNFT table
	for i := 0; i < len(roomNFTIds); i++ {
		err := c.model.DeleteRoomNFTByID(roomNFTIds[i])
		if err != nil {
			fmt.Print(err)
			res.JSON(w, 500, "cancel Book room has failed")
			return

		}

	}
	res.JSON(w, 201, "cancel Book room success")
	return
}

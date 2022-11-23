package handler

import (
	"fmt"
	jwt "golang_service/auth"
	"golang_service/contracts"
	"golang_service/models"
	res "golang_service/utils"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

var roomIdCount int = 0

const (
	checkInTimeHourBottom int = 14
	checkInTimeHourTop    int = 15
)

func New(db *gorm.DB) handler {
	return handler{db}
}

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	if govalidator.IsNull(username) || govalidator.IsNull(password) {
		res.JSON(w, 400, "Data can not empty")
		return
	}

	username = models.Santize(username)
	password = models.Santize(password)

	var user models.User
	result := h.DB.Where("user_name = ?", username).First(&user)
	// find username in user table

	if result.Error != nil {
		fmt.Println("err", result.Error)
		res.JSON(w, 400, "Username or Password incorrect")
		return
	}

	// convert interface to string
	hashedPassword := fmt.Sprintf("%v", user.Password)
	fmt.Println("hash", hashedPassword)

	err := models.CheckPasswordHash(hashedPassword, password)

	if err != nil {
		fmt.Println("err", err)
		res.JSON(w, 401, "Username or Password incorrect")
		return
	}

	token, errCreate := jwt.Create(username)

	if errCreate != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}

	res.JSON(w, 200, token)
	return
}

func (h handler) Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	fmt.Println("----------", username, email, password)

	if govalidator.IsNull(username) || govalidator.IsNull(email) || govalidator.IsNull(password) {
		res.JSON(w, 400, "Data can not empty")
		return
	}

	if !govalidator.IsEmail(email) {
		res.JSON(w, 400, "Email is invalid")
		return
	}
	var user models.User
	username = models.Santize(username)
	email = models.Santize(email)
	password = models.Santize(password)
	err := h.DB.
		Where("user_name = ? and email=?", username, email).
		First(&user).Error

	if err.Error() != "record not found" {
		res.JSON(w, 409, "Bad query")
		return
	}
	if user.UserName != "" {
		res.JSON(w, 409, "User exists")
		return
	}

	password, err = models.Hash(password)

	if err != nil {
		res.JSON(w, 500, "Register has failed")
		return
	}
	user.Email = email
	user.Password = password
	user.UserName = username
	h.DB.Save(&user)
	res.JSON(w, 201, "Register Succesfully")
	return
}

func (h handler) CreateListRoom(w http.ResponseWriter, r *http.Request) {
	landLordName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var landLord models.User
	errFindLandLord := h.DB.Where("user_name = ?", landLordName).First(&landLord)
	if errFindLandLord.Error != nil {
		res.JSON(w, 400, "Cannot find landLord!")
		return
	}
	//create transaction on blockchain
	conn, auth := contracts.GetHotelContract(landLord.PrivateKey, 0)
	_, err = conn.CreateListRoom(auth)
	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "Create list room has failed")
		return
	}

	//Save listroom to backend after listen createListRoomEvent
	event := contracts.ListenCreateListRoomEvent(common.HexToAddress(landLord.Address))
	var listRoom models.ListRoom
	listRoom.ID = int(event.ListRoomId.Uint64())
	listRoom.OwnerId = landLord.ID
	listRoom.CreateDay = res.EpochTimeToDate(event.Timestamp)
	h.DB.Save(&listRoom)
	res.JSON(w, 201, "Create list room success!")
	return

}

func (h handler) DeleteListRoom(w http.ResponseWriter, r *http.Request) {
	landLordName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var landLord models.User
	errFindLandLord := h.DB.Where("user_name = ?", landLordName).First(&landLord)
	if errFindLandLord.Error != nil {
		res.JSON(w, 400, "Cannot find landLord!")
		return
	}

	listRoomId := r.PostFormValue("listRoomId")
	//create transaction on blockchain
	conn, auth := contracts.GetHotelContract(landLord.PrivateKey, 0)
	n := new(big.Int)
	listRoomIds, _ := n.SetString(listRoomId, 10)
	_, err = conn.DeleteListRoom(auth, listRoomIds)
	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "Delete list room has failed")
		return
	}
	// Delete list room off chain
	var listRoom models.ListRoom
	h.DB.Where("list_room_id = ?", listRoomId).Delete(&listRoom)
	res.JSON(w, 201, "Delete list room success!")
	return

}

func (h handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	landLordName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var landLord models.User
	errFindLandLord := h.DB.Where("user_name = ?", landLordName).First(&landLord)
	if errFindLandLord.Error != nil {
		res.JSON(w, 400, "Cannot find landLord!")
		return
	}

	listRoomId := r.PostFormValue("listRoomId")
	pricePerDay := r.PostFormValue("pricePerDay")
	var room models.Room
	errFindRoom := h.DB.Where("list_room_id = ?", listRoomId).First(&room)
	if errFindRoom.Error != nil {
		res.JSON(w, 400, "Cannot find list room by list room id!")
		return
	}
	room.ListRoomId, _ = strconv.Atoi(listRoomId)
	room.ID = roomIdCount
	value, _ := strconv.ParseFloat(pricePerDay, 32)
	room.PricePerDay = float32(value)
	h.DB.Save(&room) // save room off-chain

	//create default state
	var stateRoom models.StateRoom
	stateRoom.RoomId = roomIdCount
	roomIdCount++
	stateRoom.ListRoomId, _ = strconv.Atoi(listRoomId)
	stateRoom.State = "Available"
	h.DB.Save(&stateRoom) // save state of new room created
	res.JSON(w, 201, "Create room successful")
	return

}

func (h handler) UpdateRoomState(w http.ResponseWriter, r *http.Request) {
	landLordName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var landLord models.User
	errFindLandLord := h.DB.Where("user_name = ?", landLordName).First(&landLord)
	if errFindLandLord.Error != nil {
		res.JSON(w, 400, "Cannot find landLord!")
		return
	}
	stateRoomId := r.PostFormValue("stateRoomId")
	listRoomId := r.PostFormValue("listRoomId")
	roomId := r.PostFormValue("roomId")
	startUpdate := res.FormatTime(r.PostFormValue("timeStartUpdate"))
	endUpdate := res.FormatTime(r.PostFormValue("endUpdate"))
	state := r.PostFormValue("state")

	var room models.Room
	errFindListRoom := h.DB.Where("list_room_id = ?", listRoomId).First(&room)
	if errFindListRoom.Error != nil {
		res.JSON(w, 400, "Cannot find list room!")
		return
	}

	errFindRoom := h.DB.Where("room_id = ?", roomId).First(&room)
	if errFindRoom.Error != nil {
		res.JSON(w, 400, "Cannot find room!")
		return
	}

	//Check at that time room is already used
	var roomNFTs []models.RoomNFT
	var isUsed bool = true

	sql := `room_id=? AND date_valid >= ? AND date_valid <= ? `
	errFindRoomNFT := h.DB.Where(sql, roomId, listRoomId, startUpdate, endUpdate).Scan(&roomNFTs)
	if errFindRoomNFT.Error != nil {
		res.JSON(w, 400, "Cannot find RoomNFT!")
		return
	}
	if len(roomNFTs) == 0 {
		isUsed = false
	}

	//Check at that time room available or not
	var stateRooms []models.StateRoom
	var available bool = true
	sql = `room_id=? AND list_room_id=? AND start_update >= ? AND end_update <= ?`
	errFindStateRooms := h.DB.Where(sql, roomId, listRoomId, startUpdate, endUpdate).Scan(&stateRooms)
	if errFindStateRooms.Error != nil {
		res.JSON(w, 400, "Cannot find State room!")
		return
	}
	if len(stateRooms) != 0 {
		available = false
	}

	if isUsed == false && available == true {
		//ready to update state
		var stateRoom models.StateRoom
		sql := "state_room_id=?"
		errFindStateRoom := h.DB.Where(sql, stateRoomId).Scan(&stateRoom)
		if errFindStateRoom.Error != nil {
			res.JSON(w, 400, "Cannot find state room!")
			return
		}
		stateRoom.State = state
		stateRoom.StartUpdate = startUpdate
		stateRoom.EndUpdate = endUpdate
		h.DB.Save(&stateRoom) // save state room off-chain
		res.JSON(w, 201, "Update state room successful")
		return
	} else {
		res.JSON(w, 400, "Room already used or not available yet!")
		return
	}
}

func (h handler) BookRoom(w http.ResponseWriter, r *http.Request) {
	userName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var user models.User
	errFindUser := h.DB.Where("user_name = ?", userName).First(&user)
	if errFindUser.Error != nil {
		res.JSON(w, 400, "Cannot find user!")
		return
	}
	roomId := r.PostFormValue("roomId")
	listRoomId := r.PostFormValue("listRoomId")
	startDay := res.FormatTime(r.PostFormValue("startDay"))
	endDay := res.FormatTime(r.PostFormValue("endDay"))

	//verify that list room id  exist
	var listRoom models.ListRoom
	errFindListRoom := h.DB.Where("list_room_id = ?", listRoomId).First(&listRoom)
	if errFindListRoom.Error != nil {
		res.JSON(w, 400, "List room is not exist !")
		return
	}
	//verify that room id  exist
	var room models.Room
	errFindRoom := h.DB.Where("list_room_id = ?", roomId).First(&room)
	if errFindRoom.Error != nil {
		res.JSON(w, 400, "Cannot find room !")
		return
	}

	//Todo check day format

	//verify all room state in range book is availale
	var stateRooms []models.StateRoom
	var available bool = true
	sql := `room_id=? AND list_room_id=? AND start_update >= ? AND end_update <= ?`
	errFindStateRooms := h.DB.Where(sql, roomId, listRoomId, startDay, endDay).Scan(&stateRooms)
	if errFindStateRooms.Error != nil {
		res.JSON(w, 400, "State Room is not exist !")
		return
	}
	if len(stateRooms) != 0 {
		available = false
	}
	//verify all room in range book is not booked before
	var roomNFTs []models.RoomNFT
	var isUsed bool = true

	sql = `room_id=? AND date_valid >= ? AND date_valid <= ? `
	errFindRoomNFT := h.DB.Where(sql, roomId, listRoomId, startDay, endDay).Scan(&roomNFTs)
	if errFindRoomNFT.Error != nil {
		res.JSON(w, 400, " Room NFT is not exist !")
		return
	}
	if len(roomNFTs) == 0 {
		isUsed = false
	}
	if isUsed == false && available == true {
		//ready to book
		//Caculate total price

		numberDaysBook := res.SubtractTime(endDay, startDay)
		totalPrice := room.PricePerDay * float32(numberDaysBook)

		//create transaction on blockchain
		conn, auth := contracts.GetHotelContract(user.PrivateKey, uint64(totalPrice))
		result, err := conn.BookRoom(auth, res.StringToBigInt(roomId), res.Float32ToBigInt(totalPrice), res.IntToBigInt(numberDaysBook))
		if err != nil {
			fmt.Println(err)
			res.JSON(w, 500, "Book room success!")
		}
		res.JSON(w, 201, result)

		// Listen bookroom Event save to RoomNFT table
		event := contracts.ListenBookRoomEvent()
		var tokenID int = int(event.StartTokenId.Uint64())
		for i := 0; i < int(event.NumberOfdates.Uint64()); i++ {
			roomNFTs[i].Booker = event.Booker.String()
			roomNFTs[i].RoomId = int(event.RoomId.Uint64())
			roomNFTs[i].TokenId = tokenID
			roomNFTs[i].DateValid = startDay
			startDay.AddDate(0, 0, 1) // increase 1 days for later RoomNFT
			tokenID++
			h.DB.Save(&roomNFTs[i])
		}
		return

	} else {
		res.JSON(w, 400, "Room already booked or not available yet!")
		return
	}

}

func (h handler) CancelBookRoom(w http.ResponseWriter, r *http.Request) {
	userName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var user models.User
	errFindUser := h.DB.Where("user_name = ?", userName).First(&user)
	if errFindUser.Error != nil {
		res.JSON(w, 400, "Cannot find user!")
		return
	}
	roomNFTIds := r.PostFormValue("roomNFTIds")
	var roomNFTs []models.RoomNFT
	//verify that roomNFT id  exist

	for i := 0; i < len(roomNFTIds); i++ {
		var roomNFT models.RoomNFT
		errFindRoomNFT := h.DB.Where("room_nft_id = ?", roomNFTIds[i]).First(&roomNFT)
		if errFindRoomNFT.Error != nil {
			res.JSON(w, 400, "Cannot find roomNFT !")
			return
		}
		roomNFTs = append(roomNFTs, roomNFT)
	}

	//Check valid cancel happens before check in time
	//check in time hard code is 14pm - 15pm
	currentTime := time.Now()
	currentTime.Format("2006-01-02")
	for i := 0; i < len(roomNFTs); i++ {
		if currentTime.After(roomNFTs[i].DateValid) {
			res.JSON(w, 400, "Cannot cancel after checkin!")
			return
		}

	}

	//Call ownerOf view function to verify that caller owner that tokenids
	conn, auth := contracts.GetHotelContract(user.PrivateKey, 0)
	var opts *bind.CallOpts
	var tokenIds []*big.Int
	for i := 0; i < len(roomNFTs); i++ {
		ownerAddress, err := conn.GetOwnerOf(opts, res.IntToBigInt(roomNFTs[i].TokenId))
		if err != nil {
			fmt.Println(err)
			return
		}
		if user.Address != ownerAddress.Hex() {
			res.JSON(w, 400, "You not own these NFT to cancel")
			return
		}
		tokenIds = append(tokenIds, res.IntToBigInt(roomNFTs[i].TokenId))

	}
	//Todo find room by roomNFT id Caculate repay price
	var rooms []models.Room
	var price float32
	for i := 0; i < len(roomNFTIds); i++ {
		var room models.Room
		errFindRoom := h.DB.Where("room_nft_id = ?", roomNFTIds[i]).First(&room)
		if errFindRoom.Error != nil {
			res.JSON(w, 400, "Cannot find roomNFT !")
			return
		}
		rooms = append(rooms, room)
		price += room.PricePerDay

	}
	//fee cancel hard save in .env
	feeCancel, _ := strconv.Atoi(os.Getenv("Fee_CANCEL"))
	totalRepayPrice := price * float32(feeCancel)

	//Todo nft.approve to approve burn permission to hotel contract
	roomnft_conn, auth_nft := contracts.GetRoomNFTContract(user.PrivateKey, 0)
	contractAddress := common.HexToAddress(os.Getenv("DEPLOYED_HOTEL_CONTRACT_ADDRESS"))
	roomnft_conn.SetApprovalForAll(auth_nft, contractAddress, true)

	//create transaction on blockchain
	_, err = conn.CancelBookRoom(auth, tokenIds, res.Float32ToBigInt(totalRepayPrice))
	if err != nil {
		fmt.Println(err)
		return
	}
	//Delete in RoomNFT table
	var roomNFT models.RoomNFT
	for i := 0; i < len(roomNFTIds); i++ {
		//Todo check error
		h.DB.Where("room_nft_id= ?", roomNFTIds[i]).Delete(&roomNFT)
	}
	res.JSON(w, 500, "Cancel Book room success!")
	return
}

func (h handler) CheckIn(w http.ResponseWriter, r *http.Request) {
	userName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var user models.User
	errFindUser := h.DB.Where("user_name = ?", userName).First(&user)
	if errFindUser.Error != nil {
		res.JSON(w, 400, "Cannot find user!")
		return
	}
	roomNFTId, _ := strconv.Atoi(r.PostFormValue("roomNFTid"))
	//verify that roomNFT id  exist
	var roomNFT models.RoomNFT
	errFindRoomNFT := h.DB.Where("room_nft_id = ?", roomNFTId).First(&roomNFT)
	if errFindRoomNFT.Error != nil {
		res.JSON(w, 400, "Cannot find roomNFT !")
		return
	}

	//Call ownerOf view function to verify that caller owner that tokenids
	conn, auth := contracts.GetHotelContract(user.PrivateKey, 0)
	var opts *bind.CallOpts
	ownerAddress, err := conn.GetOwnerOf(opts, res.IntToBigInt(roomNFT.TokenId))
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Address != ownerAddress.Hex() {
		res.JSON(w, 400, "You not own this NFT to checkin")
		return
	}

	//check time
	currentTime := time.Now()
	currentTime.Format("2006-01-02")
	if roomNFT.DateValid == currentTime && currentTime.Hour() > 14 && currentTime.Hour() < 15 {

		//Call check in on-chain
		conn, auth := contracts.GetHotelContract(user.PrivateKey, 0)
		_, err := conn.CheckIn(auth, roomNFT.TokenId)
		if err != nil {
			fmt.Println(err)
			return
		}
		res.JSON(w, 201, "Check in success!")
		return

	}
	res.JSON(w, 400, "Cannot checkin yet !")
	return
}

func (h handler) CheckOut(w http.ResponseWriter, r *http.Request) {
	userName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var user models.User
	errFindUser := h.DB.Where("user_name = ?", userName).First(&user)
	if errFindUser.Error != nil {
		res.JSON(w, 400, "Cannot find user!")
		return
	}
	//Listen checkin event to know that user already checkin can checkout
	event = contracts.ListenCheckInEvent(user.Address)
	if event != nil {
		//check time call checkout is 12pm
		currentTime := time.Now()
		currentTime.Format("2006-01-02")
		if currentTime.Hour() == 12 {

			//create transaction on blockchain
			conn, auth := contracts.GetHotelContract(user.PrivateKey, 0)
			result, err := conn.CheckOut(auth)
			if err != nil {
				fmt.Println(err)
				res.JSON(w, 201, "Check out success!")
				return
			}
			res.JSON(w, 201, result)
			return
		}
		res.JSON(w, 400, "Cannot checkout yet !")
		return
	}
	res.JSON(w, 400, "You do not have permission to checkout !")
	return

}

func (h handler) RequestPayment(w http.ResponseWriter, r *http.Request){
	landLordName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var landLord models.User
	errFindLandLord := h.DB.Where("user_name = ?", landLordName).First(&landLord)
	if errFindLandLord.Error != nil {
		res.JSON(w, 400, "Cannot find landLord!")
		return
	}
   
	roomNFTIds := r.PostFormValue("roomNFTids")
	//convert roomNFTIds to array
	//verify that roomNFT id  exist
	var roomNFTs []models.RoomNFT
	var roomIds []int
	for i := 0; i < len(roomNFTIds); i++ {
		var roomNFT models.RoomNFT
		errFindRoomNFT := h.DB.Where("room_nft_id = ?", roomNFTIds[i]).First(&roomNFT)
		if errFindRoomNFT.Error != nil {
			res.JSON(w, 400, "Cannot find roomNFT !")
			return
		}
        roomIds = append(roomIds,roomNFT.RoomId )
		roomNFTs=append(roomNFTs,roomNFT)

	}
	//Todo checktime of the last room requested is 15 days after
		currentTime := time.Now()
		currentTime.Format("2006-01-02")
		for i := 0; i < len(roomNFTs); i++ {

			if !currentTime.After(roomNFTs[i].DateValid.AddDate(0,0,15)){
				//Requested only affect after 15 days
                res.JSON(w, 400, "Not enough 15 days!")
				return
			}	
		}

	//Todo listen all bookEvent at that room
    //caculate number book request


	for _, v := range roomIds {
		skip := false
		for _, u := range roomIds {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			roomIds = append(roomIds, v)
		}
	}
	
	for i := 0; i < len(roomIds); i++ {
	numberOfBookRequest:=contracts.CaculateAllBookRoomEventByRoomID(roomIds[i])
		
	}
    



	//Todo listen all cancelEvent at that room
	//caculate number cancel request



	//=> caculate total price
	//TotalPrice=PricePerDay*bookNumber-FeeCancel*PricePerday*CancelNumber

	//create transaction on blockchain
	conn, auth := contracts.GetHotelContract(user.PrivateKey, 0)
	result, err := conn.RequestPayment(totalPrice)
	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "Request payment success!")
		return
	}
	else{
		res.JSON(w, 400, "Request payment Faild !")
		return
		}

	}

	else{
		res.JSON(w, 400, "Cannot request payment yet!")
			return
	}

}

package handler

import (
	"fmt"
	jwt "golang_service/auth"
	"golang_service/contracts"
	"golang_service/models"
	res "golang_service/utils"
	"math/big"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

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
	result, err := conn.CreateListRoom(auth)
	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "Create list room has failed")
		return
	}
	res.JSON(w, 201, result)

	//Todo save listroom to backend after listen createListRoomEvent

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
	result, err := conn.DeleteListRoom(auth, listRoomIds)
	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "Delete list room has failed")
		return
	}
	res.JSON(w, 201, result)

	//Todo Delete list room off chain
	return

}

func (h handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	roomIdCount := 0
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

	//Todo check at that time room is already used

	stateRoomId := r.PostFormValue("stateRoomId")
	listRoomId := r.PostFormValue("listRoomId")
	roomId := r.PostFormValue("roomId")
	startUpdate := r.PostFormValue("timeStartUpdate")
	endUpdate := r.PostFormValue("endUpdate")
	state := r.PostFormValue("state")

	var room models.Room

	errFindListRoom := h.DB.Where("list_room_id = ?", listRoomId).First(&room)
	if errFindListRoom.Error != nil {
		res.JSON(w, 400, "Cannot find list room by list room id!")
		return
	}

	errFindRoom := h.DB.Where("room_id = ?", roomId).First(&room)
	if errFindRoom.Error != nil {
		res.JSON(w, 400, "Cannot find room !")
		return
	}

	var stateRoom models.StateRoom

	errFindStateRoom := h.DB.Where("state_room_id = ?", stateRoomId).First(&stateRoom)
	if errFindStateRoom.Error != nil {
		res.JSON(w, 400, "Cannot find state room!")
		return
	}

	stateRoom.RoomId, _ = strconv.Atoi(roomId)
	stateRoom.ListRoomId, _ = strconv.Atoi(listRoomId)
	stateRoom.StartUpdate = startUpdate
	stateRoom.EndUpdate = endUpdate
	stateRoom.State = state

	h.DB.Save(&stateRoom) // save state room off-chain
	res.JSON(w, 201, "Update state room successful")
	return

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
	startDay := r.PostFormValue("startDay")
	endDay := r.PostFormValue("endDay")

	//verify that room id  exist
	var room models.Room
	errFindRoom := h.DB.Where("list_room_id = ?", roomId).First(&room)
	if errFindRoom.Error != nil {
		res.JSON(w, 400, "Cannot find room !")
		return
	}

	//Todo check day format

	//Todo verify all room state in range book is availale

	//Todo Caculate total price

	numberDaysBook := endDay - startDay
	totalPrice := room.PricePerDay * numberDaysBook

	//create transaction on blockchain
	conn, auth := contracts.GetHotelContract(user.PrivateKey, 0)
	n := new(big.Int)
	roomIds, _ := n.SetString(roomId, 10)
	result, err := conn.BookRoom(auth, roomId, totalPrice, numberDaysBook)
	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "Book room success!")
		return
	}
	res.JSON(w, 201, result)

	//Todo Listen bookroom Event save to RoomNFT table
	return

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
	roomNFTIds := r.PostFormValue("roomNFTids")
	//verify that roomNFT id  exist

	for i := 0; i < roomNFTIds.lengh(); i++ {
		var roomNFT models.RoomNFT
		errFindRoomNFT := h.DB.Where("room_nft_id = ?", roomNFTIds[i]).First(&roomNFT)
		roomNFTs.append(roomNFT)
		if errFindRoomNFT.Error != nil {
			res.JSON(w, 400, "Cannot find roomNFT !")
			return
		}
	}

	//Todo check valid cancel before checkout time
	//check in time hard code is 14pm

	for i := 0; i < roomNFTs.lengh(); i++ {
		if currentTime >= checkInTime {
			res.JSON(w, 400, "Out of cancel time!")
			return
		}
	}
	tokenID := roomNFT.TokenId

	//Todo listen ownerOf view function to verify that caller owner that tokenids

	//Todo Caculate repay price

	numberDaysCancelBook := roomNFTs.lengh()
	//Todo find room by roomNFT id
	//fee cancel hard save in .env
	payPrice := room.PricePerDay * numberDaysBook * feeCancel

	//Todo nft.approve to approve burn permission to hotel

	//create transaction on blockchain
	conn, auth := contracts.GetHotelContract(user.PrivateKey, 0)
	n := new(big.Int)
	roomIds, _ := n.SetString(roomId, 10)
	result, err := conn.CancelBookRoom(auth,_tokenIds,payPrice)
	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "Cancel Book room success!")
		return
	}
	res.JSON(w, 201, result)

	//Todo Listen cancel bookroom Event and delete in RoomNFT table
	return

}

func (h handler) CheckIn(w http.ResponseWriter, r *http.Request){
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
	roomNFTIds := r.PostFormValue("roomNFTids")
	//verify that roomNFT id  exist

	for i := 0; i < roomNFTIds.lengh(); i++ {
		var roomNFT models.RoomNFT
		errFindRoomNFT := h.DB.Where("room_nft_id = ?", roomNFTIds[i]).First(&roomNFT)
		roomNFTs.append(roomNFT)
		if errFindRoomNFT.Error != nil {
			res.JSON(w, 400, "Cannot find roomNFT !")
			return
		}
	}

	//check time call checkin hard checkin time is 14pm

	if(roomNFTS.get(1).DateValid==currentTime && currenTime.hour==14){
    
		//create transaction on blockchain
		conn, auth := contracts.GetHotelContract(user.PrivateKey, 0)
		result, err := conn.CheckIn(auth,tokenIds)
		if err != nil {
			fmt.Println(err)
			res.JSON(w, 500, "Check in success!")
			return
		}
		res.JSON(w, 201, result)
		return


	}
	else{
		res.JSON(w, 400, "Cannot checkin yet !")
			return
	}

}

func (h handler) CheckOut(w http.ResponseWriter, r *http.Request){
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
	roomNFTId := r.PostFormValue("roomNFTid")
	//verify that roomNFT id  exist

		var roomNFT models.RoomNFT
		errFindRoomNFT := h.DB.Where("room_nft_id = ?", roomNFTId).First(&roomNFT)
		if errFindRoomNFT.Error != nil {
			res.JSON(w, 400, "Cannot find roomNFT !")
			return
		}
	
	//check time call checkout hard checkin time is 12pm

	if(roomNFT.DateValid==currentTime.Date && currenTime.hour==12){
    
		//create transaction on blockchain
		conn, auth := contracts.GetHotelContract(user.PrivateKey, 0)
		result, err := conn.CheckOut(auth,tokenId)
		if err != nil {
			fmt.Println(err)
			res.JSON(w, 500, "Check out success!")
			return
		}
		res.JSON(w, 201, result)
		return
	}
	else{
		res.JSON(w, 400, "Cannot checkout yet !")
			return
	}

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
	//verify that roomNFT id  exist

	for i := 0; i < roomNFTIds.lengh(); i++ {
		var roomNFT models.RoomNFT
		errFindRoomNFT := h.DB.Where("room_nft_id = ?", roomNFTIds[i]).First(&roomNFT)
		roomNFTs.append(roomNFT)
		if errFindRoomNFT.Error != nil {
			res.JSON(w, 400, "Cannot find roomNFT !")
			return
		}
	}
	//Todo checktime of the last room requested is 15 days after

	if(currenTime.date >= roomNFTs.get(lengh())+15){
	//Todo listen all bookEvent at that room
    //caculate number book request

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





package handler

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
	"sort"
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

var roomIdCount int = 1

const (
	checkInTimeHourBottom int    = 14
	checkInTimeHourTop    int    = 15
	dateLayout            string = "2006-01-02"
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
	h.DB.Where("user_name = ?", landLordName).First(&landLord)
	if landLord.Role != "landlord" {
		res.JSON(w, 400, "You must be landlord to execute this function!")
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
	h.DB.Where("user_name = ?", landLordName).First(&landLord)
	if landLord.Role != "landlord" {
		res.JSON(w, 400, "You must be landlord to execute this function!")
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
	errDeleteListRoom := h.DB.Where("id = ?", listRoomId).Delete(&listRoom)
	if errDeleteListRoom.Error != nil {
		res.JSON(w, 500, "Delete list room has failed!")
		return
	}
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
	h.DB.Where("user_name = ?", landLordName).First(&landLord)
	if landLord.Role != "landlord" {
		res.JSON(w, 400, "You must be landlord to execute this function!")
		return
	}

	listRoomId := r.PostFormValue("listRoomId")
	pricePerDay := r.PostFormValue("pricePerDay")
	var room models.Room
	var listRoom models.ListRoom
	errFindListRoom := h.DB.Where("id = ?", listRoomId).First(&listRoom)
	if errFindListRoom.Error != nil {
		res.JSON(w, 400, "List room id does not exist!")
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
	h.DB.Where("user_name = ?", landLordName).First(&landLord)
	if landLord.Role != "landlord" {
		res.JSON(w, 400, "You must be landlord to execute this function!")
		return
	}
	stateRoomId := r.PostFormValue("stateRoomId")
	listRoomId := r.PostFormValue("listRoomId")
	roomId := r.PostFormValue("roomId")
	startUpdate := res.FormatTime(r.PostFormValue("startUpdate"))
	endUpdate := res.FormatTime(r.PostFormValue("endUpdate"))
	state := r.PostFormValue("state")

	//verify list room id exist
	var listRoom models.ListRoom
	errFindListRoom := h.DB.Where("id = ?", listRoomId).First(&listRoom)
	if errFindListRoom.Error != nil {
		res.JSON(w, 400, "List room id does not exist!")
		return
	}
	var room models.Room
	errFindRoom := h.DB.Where("id= ?", roomId).First(&room)
	if errFindRoom.Error != nil {
		res.JSON(w, 400, "Room id does not exist!")
		return
	}

	//verify update room state in the past
	if endUpdate.Before(time.Now()) {
		res.JSON(w, 400, "Can not update state room in the past!")
		return
	}

	//Check at that time room is already used
	var roomNFTs []models.RoomNFT
	var isUsed bool = true

	sql := `room_id=? AND date_valid >= ? AND date_valid <= ? `
	h.DB.Where(sql, roomId, startUpdate, endUpdate).Find(&roomNFTs)
	if len(roomNFTs) == 0 {
		isUsed = false
	}

	//Check at that time room available or not
	var stateRooms []models.StateRoom
	var available bool = true
	sql = `room_id=? AND list_room_id=? AND start_update >= ? AND end_update <= ?`
	h.DB.Where(sql, roomId, listRoomId, startUpdate, endUpdate).Find(&stateRooms)
	if len(stateRooms) != 0 {
		available = false
	}
	fmt.Print("available ", available)

	if isUsed == false && available == true {
		//ready to update state
		var stateRoom models.StateRoom
		errFindStateRoom := h.DB.Where("id = ? ", stateRoomId).First(&stateRoom)
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
	}
	res.JSON(w, 400, "Room already used or not available yet!")
	return
}

func (h handler) BookRoom(w http.ResponseWriter, r *http.Request) {
	userName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var user models.User
	h.DB.Where("user_name = ?", userName).First(&user)

	roomId := r.PostFormValue("roomId")
	listRoomId := r.PostFormValue("listRoomId")
	startDay := res.FormatTime(r.PostFormValue("startDay"))
	startDay = time.Date(startDay.Year(), startDay.Month(), startDay.Day(), 7, 0, 0, 0, startDay.Location())
	endDay := res.FormatTime(r.PostFormValue("endDay"))
	endDay = time.Date(endDay.Year(), endDay.Month(), endDay.Day(), 7, 0, 0, 0, endDay.Location())

	//verify that list room id  exist
	var listRoom models.ListRoom
	errFindListRoom := h.DB.Where("id = ?", listRoomId).First(&listRoom)
	if errFindListRoom.Error != nil {
		res.JSON(w, 400, "List room id is not exist !")
		return
	}
	//verify that room id  exist
	var room models.Room
	errFindRoom := h.DB.Where("id = ?", roomId).First(&room)
	if errFindRoom.Error != nil {
		res.JSON(w, 400, "Room id is not exist !")
		return
	}

	//Todo check day format

	//verify book room in the past
	if startDay.Before(time.Now()) || endDay.Before(time.Now()) {
		res.JSON(w, 400, "Can not book room in the past!")
		return
	}

	//verify all room state in range book is availale
	var stateRooms []models.StateRoom
	var available bool = true
	sql := `room_id=? AND list_room_id=? AND start_update >= ? AND end_update <= ?`
	h.DB.Where(sql, roomId, listRoomId, startDay, endDay).Find(&stateRooms)
	if len(stateRooms) != 0 {
		available = false
	}
	fmt.Print("available ", available)
	//verify all room in range book is not booked before
	var roomNFTs []models.RoomNFT
	var isUsed bool = true

	sql = `room_id=? AND date_valid >= ? AND date_valid <= ? `
	h.DB.Where(sql, roomId, startDay, endDay).Find(&roomNFTs)
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
		conn, auth := contracts.GetHotelContract(user.PrivateKey, uint64(totalPrice))
		_, err := conn.BookRoom(auth, res.StringToBigInt(roomId), res.Float32ToBigInt(totalPrice), res.IntToBigInt(numberDaysBook))
		if err != nil {
			fmt.Println(err)
			return
		}

		// Listen bookroom Event save to RoomNFT table
		event := contracts.ListenBookRoomEvent(common.HexToAddress(user.Address))
		var tokenID int = int(event.StartTokenId.Uint64())
		for i := 0; i < int(event.NumberOfdates.Uint64()); i++ {
			var roomNFT models.RoomNFT
			roomNFT.Booker = event.Booker.String()
			roomNFT.RoomId = int(event.RoomId.Uint64())
			roomNFT.TokenId = tokenID
			roomNFT.DateValid = startDay
			startDay = startDay.AddDate(0, 0, 1) // increase 1 days for later RoomNFT in for loop
			tokenID++
			h.DB.Save(&roomNFT)
		}
		res.JSON(w, 201, "Book room success!")
		return
	}
	res.JSON(w, 400, "Room already booked or not available yet!")
	return

}

func (h handler) CancelBookRoom(w http.ResponseWriter, r *http.Request) {
	userName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var user models.User
	h.DB.Where("user_name = ?", userName).First(&user)

	type RoomNFTIds struct {
		RoomNFTIds []int `json:"roomNFTIds"`
	}
	var roomNFTId RoomNFTIds
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
		var roomNFT models.RoomNFT
		errFindRoomNFT := h.DB.Where("id = ?", roomNFTIds[i]).First(&roomNFT)
		if errFindRoomNFT.Error != nil {
			res.JSON(w, 400, "Cannot find roomNFT !")
			return
		}
		roomNFTs = append(roomNFTs, roomNFT)
	}
	//fmt.Print(" roomNFTS ", roomNFTs)

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
	conn, auth := contracts.GetHotelContract(user.PrivateKey, 0)
	var opts *bind.CallOpts
	var tokenIds []*big.Int
	for i := 0; i < len(roomNFTs); i++ {
		ownerAddress, err := conn.GetOwnerOf(opts, res.IntToBigInt(roomNFTs[i].TokenId))
		//fmt.Print("ownerAddress ", ownerAddress)
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
	//fmt.Print("tokenIds ", tokenIds)
	//Find room by room id Caculate repay price
	var rooms []models.Room
	var price float32
	for i := 0; i < len(roomNFTs); i++ {
		var room models.Room
		errFindRoom := h.DB.Where("id = ?", roomNFTs[i].RoomId).First(&room)
		if errFindRoom.Error != nil {
			res.JSON(w, 400, "RoomNFT id does not exist !")
			return
		}
		rooms = append(rooms, room)
		price += room.PricePerDay

	}
	//fee cancel hard save in .env
	feeCancel, err := strconv.ParseFloat(os.Getenv("FEE_CANCEL"), 32)
	totalRepayPrice := price * float32(feeCancel)

	//Todo nft.approve to approve burn permission to hotel contract
	roomnft_conn, auth_nft := contracts.GetRoomNFTContract(user.PrivateKey, 0)
	contractAddress := common.HexToAddress(os.Getenv("DEPLOYED_HOTEL_CONTRACT_ADDRESS"))
	_, errSetApproval := roomnft_conn.SetApprovalForAll(auth_nft, contractAddress, true)
	if errSetApproval != nil {
		fmt.Println("error ", err)
		return

	}

	//create transaction on blockchain
	conn, auth = contracts.GetHotelContract(user.PrivateKey, 0)
	_, err = conn.CancelBookRoom(auth, tokenIds, res.Float32ToBigInt(totalRepayPrice))
	if err != nil {
		fmt.Println(err)
		return
	}
	//Delete in RoomNFT table
	var roomNFT models.RoomNFT
	for i := 0; i < len(roomNFTIds); i++ {
		err = h.DB.Where("id = ?", roomNFTIds[i]).Delete(&roomNFT).Error
		if err != nil {
			fmt.Print(err)
			res.JSON(w, 500, "Cancel Book room has failed!")
			return

		}

	}
	res.JSON(w, 201, "Cancel Book room success!")
	return
}

func (h handler) CheckIn(w http.ResponseWriter, r *http.Request) {
	userName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var user models.User
	h.DB.Where("user_name = ?", userName).First(&user)

	type RoomNFTIds struct {
		RoomNFTIds []int `json:"roomNFTIds"`
	}
	var roomNFTId RoomNFTIds
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
	//verify that roomNFT id  exist
	var roomNFTs []models.RoomNFT
	var tokenIds []*big.Int

	for i := 0; i < len(roomNFTId.RoomNFTIds); i++ {
		var roomNFT models.RoomNFT
		errFindRoomNFT := h.DB.Where("id = ?", roomNFTId.RoomNFTIds[i]).First(&roomNFT)
		if errFindRoomNFT.Error != nil {
			res.JSON(w, 400, " roomNFT does not exist !")
			return
		}
		roomNFTs = append(roomNFTs, roomNFT)
		tokenIds = append(tokenIds, res.IntToBigInt(roomNFT.TokenId))
	}

	//Call ownerOf view function to verify that caller owner that tokenids
	conn, _ := contracts.GetHotelContract(user.PrivateKey, 0)
	var opts *bind.CallOpts
	for i := 0; i < len(roomNFTs); i++ {
		ownerAddress, err := conn.GetOwnerOf(opts, res.IntToBigInt(roomNFTs[i].TokenId))
		if err != nil {
			fmt.Println(err)
			return
		}
		if user.Address != ownerAddress.Hex() {
			res.JSON(w, 400, "You not own this NFT to checkin")
			return
		}
		//get earliest nft minted in batch
		sort.SliceStable(roomNFTs, func(i, j int) bool {
			return roomNFTs[i].DateValid.Unix() < roomNFTs[j].DateValid.Unix()
		})

	}

	//check time
	currentTime := time.Now() //14pm-15pm
	//currentTime.After(roomNFTs[0].DateValid) &&
	if currentTime.After(roomNFTs[0].DateValid) && currentTime.Before(roomNFTs[0].DateValid.Add(time.Hour)) {

		//approve permission for hotel contract
		roomnft_conn, auth_nft := contracts.GetRoomNFTContract(user.PrivateKey, 0)
		contractAddress := common.HexToAddress(os.Getenv("DEPLOYED_HOTEL_CONTRACT_ADDRESS"))
		_, errSetApproval := roomnft_conn.SetApprovalForAll(auth_nft, contractAddress, true)
		if errSetApproval != nil {
			fmt.Println("error ", err)
			return

		}

		//Call check in on-chain
		conn, auth := contracts.GetHotelContract(user.PrivateKey, 0)
		fmt.Print("tokenIds ", tokenIds)
		_, err := conn.CheckIn(auth, tokenIds)
		if err != nil {
			fmt.Println("contract error", err)
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
	h.DB.Where("user_name = ?", userName).First(&user)

	//Listen checkin event to know that user already checkin can checkout
	event := contracts.ListenCheckInEvent(common.HexToAddress(user.Address))
	if event != nil {
		//check time call checkout is 12pm
		currentTime := time.Now()
		currentTime.Format(dateLayout)
		//test: currentTime.Hour() != 0   replace currentTime.Hour() == 12
		if currentTime.Hour() == 12 {

			//create transaction on blockchain
			conn, auth := contracts.GetHotelContract(user.PrivateKey, 0)
			_, err := conn.CheckOut(auth)
			if err != nil {
				fmt.Println(err)
				res.JSON(w, 400, "Check out has fail!")
				return
			}
			res.JSON(w, 201, "Check out success!")
			return
		}
		res.JSON(w, 400, "Cannot checkout yet !")
		return
	}
	res.JSON(w, 400, "You do not have permission to checkout !")
	return

}

func (h handler) RequestPayment(w http.ResponseWriter, r *http.Request) {
	landLordName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	var landLord models.User
	h.DB.Where("user_name = ?", landLordName).First(&landLord)
	if landLord.Role != "landlord" {
		res.JSON(w, 400, "You must be landlord to execute this function!")
		return
	}
	type RoomNFTIds struct {
		RoomNFTIds []int `json:"roomNFTIds"`
	}
	var roomNFTId RoomNFTIds
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
	//verify that roomNFT id  exist
	var roomNFTs []models.RoomNFT
	var totalPrice float32 = 0
	for i := 0; i < len(roomNFTId.RoomNFTIds); i++ {
		var roomNFT models.RoomNFT
		var room models.Room
		errFindRoomNFT := h.DB.Where("id = ?", roomNFTId.RoomNFTIds[i]).First(&roomNFT)
		if errFindRoomNFT.Error != nil {
			res.JSON(w, 400, "RoomNFT does not exist !")
			return
		}
		errFindRoom := h.DB.Where("id = ?", roomNFT.RoomId).First(&room)
		if errFindRoom.Error != nil {
			res.JSON(w, 400, "Cannot find room !")
			return
		}
		roomNFTs = append(roomNFTs, roomNFT)
		numberOfBookRequest := contracts.CaculateAllBookRoomEventByTokenID(roomNFT.TokenId)
		numberOfCancelBookRequest := contracts.CaculateAllCancelBookRoomEventByTokenID(roomNFT.TokenId)
		feeCancel, _ := strconv.Atoi(os.Getenv("FEE_CANCEL"))
		price := room.PricePerDay * (float32(numberOfBookRequest) - float32(feeCancel*numberOfCancelBookRequest))
		totalPrice += price
	}
	fmt.Print("TotalPrice ", totalPrice)

	//checktime is 15 days after DateValid
	currentTime := time.Now()
	for i := 0; i < len(roomNFTs); i++ {
		//test: comment all for loop to by pass this case
		if !currentTime.After(roomNFTs[i].DateValid.AddDate(0, 0, 15)) {
			//Requested only affect after 15 days
			res.JSON(w, 400, "Not enough 15 days!")
			return
		}
	}

	//create transaction on blockchain
	conn, auth := contracts.GetHotelContract(landLord.PrivateKey, 0)
	_, err = conn.RequestPayment(auth, res.Float32ToBigInt(totalPrice))
	if err != nil {
		fmt.Println(err)
		res.JSON(w, 500, "Request payment has faild!")
		return
	}
	res.JSON(w, 201, "Request payment success!")
	return
}

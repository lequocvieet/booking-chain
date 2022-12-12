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
	"sort"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var roomNFTCotract = &contracts.RoomNFT{}
var hotelContract = &contracts.Hotel{}

type userController struct {
	model *models.Model
}

func NewUserController(model *models.Model) userController {
	roomNFTCotract = contracts.GetRoomNFTContract()
	hotelContract = contracts.GetHotelContract()
	return userController{model}
}

func (c userController) Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	if govalidator.IsNull(username) || govalidator.IsNull(password) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	username = models.Santize(username)
	password = models.Santize(password)

	// find username in user table
	user, err := c.model.FindUserByUserName(username)
	if err != nil {
		res.JSON(w, 400, err)
		return
	}
	// convert interface to string
	hashedPassword := fmt.Sprintf("%v", user.Password)
	fmt.Println("hash", hashedPassword)
	err = models.CheckPasswordHash(hashedPassword, password)

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

func (c userController) Register(w http.ResponseWriter, r *http.Request) {
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
	username = models.Santize(username)
	email = models.Santize(email)
	password = models.Santize(password)
	user, err := c.model.FindUserByUserNameAndEmail(username, email)
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
	user = models.NewUser(username, email, password)
	c.model.SaveUser(user)
	res.JSON(w, 201, "Register Succesfully")
	return
}

func (c userController) CheckIn(w http.ResponseWriter, r *http.Request) {
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
	//verify that roomNFT id  exist
	var roomNFTs []models.RoomNFT
	var tokenIds []*big.Int

	for i := 0; i < len(roomNFTId.RoomNFTIds); i++ {
		roomNFT, err := c.model.FindRoomNFTByID(roomNFTId.RoomNFTIds[i])
		if err != nil {
			res.JSON(w, 400, err)
			return
		}
		roomNFTs = append(roomNFTs, roomNFT)
		tokenIds = append(tokenIds, res.IntToBigInt(roomNFT.TokenId))
	}

	//Call ownerOf view function to verify that caller owner that tokenids
	var opts *bind.CallOpts
	for i := 0; i < len(roomNFTs); i++ {
		ownerAddress, err := hotelContract.GetOwnerOf(opts, res.IntToBigInt(roomNFTs[i].TokenId))
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
		contractAddress := common.HexToAddress(os.Getenv("DEPLOYED_HOTEL_CONTRACT_ADDRESS"))
		auth, err := contracts.SignTransaction(user.PrivateKey, 0)
		if err != nil {
			res.JSON(w, 201, err)
			return
		}
		_, errSetApproval := roomNFTCotract.SetApprovalForAll(auth, contractAddress, true)
		if errSetApproval != nil {
			res.JSON(w, 201, errSetApproval)
			return

		}

		//Call check in on-chain
		auth, err = contracts.SignTransaction(user.PrivateKey, 0)
		if err != nil {
			res.JSON(w, 201, err)
			return
		}
		fmt.Print("tokenIds ", tokenIds)
		_, errCheckIn := hotelContract.CheckIn(auth, tokenIds)
		if errCheckIn != nil {
			res.JSON(w, 201, errCheckIn)
			return
		}
		res.JSON(w, 201, "check in success")
		return

	}
	res.JSON(w, 400, "cannot checkin yet")
	return
}

func (c userController) CheckOut(w http.ResponseWriter, r *http.Request) {
	userName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	user, err := c.model.FindUserByUserName(userName)
	if err != nil {
		res.JSON(w, 500, err)
		return
	}

	//Listen checkin event to know that user already checkin can checkout
	checkInEvent := contracts.ListenEvent(common.HexToAddress(user.Address), "CheckIn").(*contracts.HotelCheckIn)
	if checkInEvent != nil {
		//check time call checkout is 12pm
		currentTime := time.Now()
		currentTime.Format(dateLayout)
		//test: currentTime.Hour() != 0   replace currentTime.Hour() == 12
		if currentTime.Hour() == 12 {

			//create transaction on blockchain
			auth, err := contracts.SignTransaction(user.PrivateKey, 0)
			if err != nil {
				res.JSON(w, 201, err)
				return
			}
			_, errCheckOut := hotelContract.CheckOut(auth)
			if errCheckOut != nil {
				res.JSON(w, 400, "check out has fail")
				return
			}
			res.JSON(w, 201, "check out success")
			return
		}
		res.JSON(w, 400, "cannot checkout yet ")
		return
	}
	res.JSON(w, 400, "you do not have permission to checkout ")
	return

}

func (c userController) RequestPayment(w http.ResponseWriter, r *http.Request) {
	landLordName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	landLord, err := c.model.FindUserByUserName(landLordName)
	if err != nil {
		res.JSON(w, 500, err)
		return
	}
	if landLord.Role != "landlord" {
		res.JSON(w, 400, "you must be landlord to execute this function")
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
	//verify that roomNFT id  exist
	var roomNFTs []models.RoomNFT
	var totalPrice float32 = 0
	for i := 0; i < len(roomNFTId.RoomNFTIds); i++ {
		var room models.Room
		roomNFT, err := c.model.FindRoomNFTByID(roomNFTId.RoomNFTIds[i])
		if err != nil {
			res.JSON(w, 400, err)
			return
		}
		room, errFindRoom := c.model.FindRoomByID(roomNFT.RoomId)
		if errFindRoom != nil {
			res.JSON(w, 400, errFindRoom)
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
	auth, err := contracts.SignTransaction(landLord.PrivateKey, 0)
	if err != nil {
		res.JSON(w, 201, err)
		return
	}
	_, err = hotelContract.RequestPayment(auth, res.Float32ToBigInt(totalPrice))
	if err != nil {
		res.JSON(w, 500, "request payment has fail")
		return
	}
	res.JSON(w, 201, "request payment success")
	return
}

func (c userController) TransferRoomNFT(w http.ResponseWriter, r *http.Request) {
	userName, err := jwt.ExtractUsernameFromToken(r)
	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}
	user, err := c.model.FindUserByUserName(userName)
	if err != nil {
		res.JSON(w, 500, err)
		return
	}
	var tranferRequest models.TransferRequest
	var tokenIds []*big.Int
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res.JSON(w, 400, "Read Body error!")
		return
	}
	err = json.Unmarshal(body, &tranferRequest)
	if err != nil {
		fmt.Print(err)
		res.JSON(w, 400, "Unmarshal error!")
		return
	}

	//verify that roomNFT id  exist
	for i := 0; i < len(tranferRequest.RoomNFTIds); i++ {
		roomNFT, err := c.model.FindRoomNFTByID(tranferRequest.RoomNFTIds[i])
		if err != nil {
			res.JSON(w, 400, err)
			return
		}
		tokenIds = append(tokenIds, res.IntToBigInt(roomNFT.TokenId))

	}

	//transfer on blockchain
	auth, err := contracts.SignTransaction(user.PrivateKey, 0)
	if err != nil {
		res.JSON(w, 201, err)
		return
	}
	_, err = hotelContract.TranferRoomNFT(auth, tokenIds, common.HexToAddress(tranferRequest.Receiver))
	if err != nil {
		res.JSON(w, 500, "transfer RoomNFT has fail")
		return
	}
	res.JSON(w, 201, "transfer RoomNFT success")
	return

}

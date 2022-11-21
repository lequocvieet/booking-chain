package handler

import (
	"fmt"
	jwt "golang_service/auth"
	"golang_service/contracts"
	"golang_service/models"
	res "golang_service/utils"
	"net/http"

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
	return

}

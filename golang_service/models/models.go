package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Model {
	return Model{db}
}

type ListRoom struct {
	ID        int       `json:"id" gorm:"column:id"`
	OwnerId   int       `json:"owner_id" gorm:"column:owner_id"`
	CreateDay time.Time `json:"create_day" gorm:"column:create_day"`
}

type Room struct {
	ID          int     `json:"id" gorm:"column:id"`
	PricePerDay float32 `json:"price_per_day" gorm:"column:price_per_day"`
	ListRoomId  int     `json:"list_room_id" gorm:"column:list_room_id"`
}

type RoomNFT struct {
	ID        int       `json:"id" gorm:"column:id;autoIncrement:true"`
	RoomId    int       `json:"room_id" gorm:"column:room_id"`
	TokenId   int       `json:"token_id" gorm:"column:token_id"`
	Booker    string    `json:"booker" gorm:"column:booker"`
	DateValid time.Time `json:"date_valid" gorm:"column:date_valid"`
}

type StateRoom struct {
	ID          int       `json:"id" gorm:"column:id;autoIncrement:true"`
	RoomId      int       `json:"room_id" gorm:"column:room_id"`
	EndUpdate   time.Time `json:"end_update" gorm:"column:end_update"`
	StartUpdate time.Time `json:"start_update" gorm:"column:start_update"`
	ListRoomId  int       `json:"list_room_id" gorm:"column:list_room_id"`
	State       string    `json:"state" gorm:"column:state"`
}

type User struct {
	ID         int    `json:"id" gorm:"column:id;autoIncrement:true"`
	UserName   string `json:"user_name" gorm:"column:user_name"`
	Email      string `json:"email" gorm:"column:email"`
	Password   string `json:"password" gorm:"column:password;size:1000"`
	Address    string `json:"address" gorm:"column:address"`
	PrivateKey string `json:"private_key" gorm:"column:private_key"`
	Role       string `json:"role" gorm:"column:role"`
}

type RoomNFTIds struct {
	RoomNFTIds []int `json:"roomNFTIds"`
}

type TransferRequest struct {
	RoomNFTIds []int  `json:"roomNFTIds"`
	Receiver   string `json:"receiver"`
}

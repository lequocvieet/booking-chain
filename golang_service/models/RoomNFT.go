package models

import "time"

type RoomNFT struct {
	ID        int       `json:"id" gorm:"column:id"`
	RoomId    int       `json:"room_id" gorm:"column:room_id"`
	TokenId   int       `json:"token_id" gorm:"column:token_id"`
	Booker    string    `json:"booker" gorm:"column:booker"`
	DateValid time.Time `json:"date_valid" gorm:"column:date_valid"`
}

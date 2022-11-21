package models

import "time"

type RoomState struct {
	ID          int       `json:"id" gorm:"column:id"`
	RoomId      int       `json:"room_id" gorm:"column:room_id"`
	EndUpdate   time.Time `json:"end_update" gorm:"column:end_update"`
	StartUpdate time.Time `json:"start_update" gorm:"column:start_update"`
}

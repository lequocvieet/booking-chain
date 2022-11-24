package models

import "time"

type StateRoom struct {
	ID          int       `json:"id" gorm:"column:id;autoIncrement:true"`
	RoomId      int       `json:"room_id" gorm:"column:room_id"`
	EndUpdate   time.Time `json:"end_update" gorm:"column:end_update"`
	StartUpdate time.Time `json:"start_update" gorm:"column:start_update"`
	ListRoomId  int       `json:"list_room_id" gorm:"column:list_room_id"`
	State       string    `json:"state" gorm:"column:state"`
}

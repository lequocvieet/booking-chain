package models

type Room struct {
	ID          int     `json:"id" gorm:"column:id"`
	PricePerDay float32 `json:"price_per_day" gorm:"column:price_per_day"`
	ListRoomId  int     `json:"list_room_id" gorm:"column:list_room_id"`
}

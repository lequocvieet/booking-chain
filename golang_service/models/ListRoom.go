package models

import (
	"time"
)

type ListRoom struct {
	ID        int       `json:"id" gorm:"column:id"`
	OwnerId   string    `json:"owner_id" gorm:"column:owner_id"`
	CreateDay time.Time `json:"create_day" gorm:"column:create_day"`
}

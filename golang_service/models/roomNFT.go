package models

import (
	"fmt"
	"time"
)

func NewRoomNFT(roomId int, tokenId int, booker string, dateValid time.Time) RoomNFT {
	u := RoomNFT{
		RoomId:    roomId,
		TokenId:   tokenId,
		Booker:    booker,
		DateValid: dateValid,
	}
	return u
}

func (m *Model) FindRoomNFTByRoomIdAndDateValid(roomId int, start time.Time, end time.Time) []RoomNFT {
	var roomNFTs []RoomNFT
	sql := `room_id=? AND date_valid >= ? AND date_valid <= ? `
	m.DB.Where(sql, roomId, start, end).Find(&roomNFTs)
	return roomNFTs
}
func (m *Model) SaveRoomNFT(roomNFT RoomNFT) error {
	m.DB.Save(&roomNFT)
	return fmt.Errorf("save roomNFT success")
}
func (m *Model) FindRoomNFTByID(ID int) (RoomNFT, error) {
	var roomNFT RoomNFT
	result := m.DB.Where("id = ?", ID).First(&roomNFT)
	if result.Error != nil {
		return roomNFT, fmt.Errorf("roomNFT ID not exist")
	}
	return roomNFT, nil
}

func (m *Model) DeleteRoomNFTByID(ID int) error {
	var roomNFT RoomNFT
	result := m.DB.Where("id = ?", ID).Delete(&roomNFT)
	if result.Error != nil {
		return fmt.Errorf("delete roomNFT has failed")
	}
	return fmt.Errorf("delete roomNFT success")
}

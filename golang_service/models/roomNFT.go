package models

import (
	"fmt"
	"time"
)

func (m *Model) FindRoomNFTByRoomIdAndDateValid(roomId int, start time.Time, end time.Time) ([]RoomNFT, error) {
	var roomNFTs []RoomNFT
	sql := `room_id=? AND date_valid >= ? AND date_valid <= ? `
	m.DB.Where(sql, roomId, start, end).Find(&roomNFTs)
	return roomNFTs, nil
}
func (m *Model) SaveRoomNFT(roomNFT RoomNFT) error {
	m.DB.Save(&roomNFT)
	return fmt.Errorf("Save roomNFT success!")
}
func (m *Model) FindByID(ID int) (RoomNFT, error) {
	var roomNFT RoomNFT
	result := m.DB.Where("id = ?", ID).First(&roomNFT)
	if result.Error != nil {
		return roomNFT, fmt.Errorf("RoomNFT ID not exist")
	}
	return roomNFT, nil
}

func (m *Model) DeleteRoomNFTByID(ID int) error {
	var roomNFT RoomNFT
	result := m.DB.Where("id = ?", ID).Delete(&roomNFT)
	if result.Error != nil {
		return fmt.Errorf("Delete roomNFT has failed!")
	}
	return fmt.Errorf("Delete roomNFT success!")
}

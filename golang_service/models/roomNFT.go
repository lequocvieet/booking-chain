package models

import (
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

func (m *Model) FindRoomNFTByRoomIdAndDateValid(roomId int, start time.Time, end time.Time) ([]RoomNFT, error) {
	var roomNFTs []RoomNFT
	sql := `room_id=? AND date_valid >= ? AND date_valid <= ? `
	err := m.DB.Where(sql, roomId, start, end).Find(&roomNFTs).Error
	if err != nil {
		return []RoomNFT{}, err
	}
	return roomNFTs, nil
}
func (m *Model) SaveRoomNFT(roomNFT RoomNFT) error {
	err := m.DB.Save(&roomNFT).Error
	if err != nil {
		return err
	}
	return nil
}
func (m *Model) FindRoomNFTByID(ID int) (RoomNFT, error) {
	var roomNFT RoomNFT
	err := m.DB.Where("id = ?", ID).First(&roomNFT).Error
	if err != nil {
		return RoomNFT{}, err
	}
	return roomNFT, nil
}

func (m *Model) DeleteRoomNFTByID(ID int) error {
	var roomNFT RoomNFT
	err := m.DB.Where("id = ?", ID).Delete(&roomNFT).Error
	if err != nil {
		return err
	}
	return nil
}

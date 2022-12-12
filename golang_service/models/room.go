package models

import "fmt"

func NewRoom(ID int, pricePerDay float32, listRoomId int) Room {
	u := Room{
		ID:          ID,
		PricePerDay: pricePerDay,
		ListRoomId:  listRoomId,
	}
	return u
}

func (m *Model) FindRoomByID(ID int) (Room, error) {
	var room Room
	errFindRoom := m.DB.Where("id= ?", ID).First(&room)
	if errFindRoom.Error != nil {
		return room, fmt.Errorf("room id does not exist")
	}
	return room, nil
}

func (m *Model) SaveRoom(room Room) error {
	m.DB.Save(&room)
	return fmt.Errorf("Save Room success!")
}

func (m *Model) DeleteRoom(roomId int) error {
	var room Room
	errDeleteRoom := m.DB.Where("id = ?", roomId).Delete(&room)
	if errDeleteRoom.Error != nil {
		return fmt.Errorf("Delete room has failed!")
	}
	return fmt.Errorf("Delete room success!")
}

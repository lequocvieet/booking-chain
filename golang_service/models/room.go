package models

import "fmt"

func NewRoom(pricePerDay float32, listRoomId int) *Room {
	return &Room{
		PricePerDay: pricePerDay,
		ListRoomId:  listRoomId,
	}

}

func (m *Model) FindRoomByID(ID int) (Room, error) {
	var room Room
	err := m.DB.Where("id= ?", ID).First(&room).Error
	if err != nil {
		return Room{}, err
	}
	return room, nil
}

func (m *Model) SaveRoom(room *Room) (int, error) {
	fmt.Println(room)
	err := m.DB.Create(room).Error
	if err != nil {
		return 0, err
	}

	return room.ID, nil
}

func (m *Model) DeleteRoom(roomId int) error {
	var room Room
	err := m.DB.Where("id = ?", roomId).Delete(&room).Error
	if err != nil {
		return err
	}
	return nil
}

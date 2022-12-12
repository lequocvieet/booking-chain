package models

import (
	"fmt"
	"time"
)

func NewListRoom(ID int, ownerID int, createDay time.Time) ListRoom {
	u := ListRoom{
		ID:        ID,
		OwnerId:   ownerID,
		CreateDay: createDay,
	}
	return u
}
func (m *Model) FindListRoom(ID int) (ListRoom, error) {
	var listRoom ListRoom
	errFindListRoom := m.DB.Where("id = ?", ID).First(&listRoom)
	if errFindListRoom.Error != nil {
		return listRoom, fmt.Errorf("list room id does not exist")
	}
	return listRoom, nil
}
func (m *Model) SaveListRoom(listRoom ListRoom) error {
	m.DB.Save(&listRoom)
	return fmt.Errorf("save list room success")
}

func (m *Model) DeleteListRoom(listRoomId int) error {
	var listRoom ListRoom
	errDeleteListRoom := m.DB.Where("id = ?", listRoomId).Delete(&listRoom)
	if errDeleteListRoom.Error != nil {
		return fmt.Errorf("delete list room has failed")
	}
	return fmt.Errorf("delete list room success")
}

package models

import (
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
	err := m.DB.Where("id = ?", ID).First(&listRoom).Error
	if err != nil {
		return ListRoom{}, err
	}
	return listRoom, nil
}
func (m *Model) SaveListRoom(listRoom ListRoom) error {
	err := m.DB.Save(&listRoom).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) DeleteListRoom(listRoomId int) error {
	var listRoom ListRoom
	err := m.DB.Where("id = ?", listRoomId).Delete(&listRoom).Error
	if err != nil {
		return err
	}
	return nil
}

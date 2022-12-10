package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Model struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Model {
	return Model{db}
}

func (m *Model) FindListRoom(ID int) (ListRoom, error) {
	var listRoom ListRoom
	errFindListRoom := m.DB.Where("id = ?", ID).First(&listRoom)
	if errFindListRoom.Error != nil {
		return listRoom, fmt.Errorf("List room id does not exist!")
	}
	return listRoom, nil
}
func (m *Model) SaveListRoom(listRoom ListRoom) error {
	m.DB.Save(&listRoom)
	return fmt.Errorf("Save list room success!")
}

func (m *Model) DeleteListRoom(listRoomId int) error {
	var listRoom ListRoom
	errDeleteListRoom := m.DB.Where("id = ?", listRoomId).Delete(&listRoom)
	if errDeleteListRoom.Error != nil {
		return fmt.Errorf("Delete list room has failed!")
	}
	return fmt.Errorf("Delete list room success!")
}

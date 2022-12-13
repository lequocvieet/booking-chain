package models

import (
	"time"
)

func NewStateRoom(roomId int, endUpdate time.Time, startUpdate time.Time, listRoomId int, state string) StateRoom {
	u := StateRoom{
		RoomId:      roomId,
		EndUpdate:   endUpdate,
		StartUpdate: startUpdate,
		ListRoomId:  listRoomId,
		State:       state,
	}
	return u
}

func (m *Model) FindStateRoomById(ID int) (StateRoom, error) {
	var stateRoom StateRoom
	err := m.DB.Where("id = ? ", ID).First(&stateRoom).Error
	if err != nil {
		return StateRoom{}, err
	}
	return stateRoom, nil
}

func (m *Model) FindStateRoomByDay(roomId int, listRoomId int, startDay time.Time, endDay time.Time) []StateRoom {

	//find state room that has end day || start day in range
	var stateRooms []StateRoom
	sql := `room_id=? AND list_room_id=? AND start_update <= ? AND end_update >= ?`
	m.DB.Where(sql, roomId, listRoomId, startDay, startDay).Find(&stateRooms)
	m.DB.Where(sql, roomId, listRoomId, endDay, endDay).Find(&stateRooms)
	return stateRooms

}

func (m *Model) FindStateRoom(roomId int, listRoomId int, startUpdate time.Time, endUpdate time.Time) ([]StateRoom, error) {

	//find state room in range
	var stateRooms []StateRoom
	sql := `room_id=? AND list_room_id=? AND start_update >= ? AND end_update <= ?`
	err := m.DB.Where(sql, roomId, listRoomId, startUpdate, endUpdate).Find(&stateRooms).Error
	if err != nil {
		return []StateRoom{}, err
	}
	return stateRooms, nil

}

func (m *Model) SaveStateRoom(stateRoom StateRoom) error {
	err := m.DB.Save(&stateRoom).Error
	if err != nil {
		return err
	}
	return nil
}

package models

import (
	"fmt"
	"time"
)

func (m *Model) FindStateRoomByDay(roomId int, listRoomId int, startDay time.Time, endDay time.Time) ([]StateRoom, error) {

	//find state room that has end day || start day in range
	var stateRooms []StateRoom
	sql := `room_id=? AND list_room_id=? AND start_update <= ? AND end_update >= ?`
	m.DB.Where(sql, roomId, listRoomId, startDay, startDay).Find(&stateRooms)
	m.DB.Where(sql, roomId, listRoomId, endDay, endDay).Find(&stateRooms)

	return stateRooms, nil

}

func (m *Model) FindStateRoom(roomId int, listRoomId int, startUpdate time.Time, endUpdate time.Time) ([]StateRoom, error) {

	//find state room in range
	var stateRooms []StateRoom
	sql := `room_id=? AND list_room_id=? AND start_update >= ? AND end_update <= ?`
	m.DB.Where(sql, roomId, listRoomId, startUpdate, endUpdate).Find(&stateRooms)
	return stateRooms, nil

}

func (m *Model) SaveStateRoom(stateRoom StateRoom) error {
	m.DB.Save(&stateRoom)
	return fmt.Errorf("Save user success!")
}

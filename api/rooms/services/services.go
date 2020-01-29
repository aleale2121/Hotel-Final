package services

import (
	"fmt"
  "github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/rooms"
)

// RoomServiceImpl implements rooms.RoomService interface
type RoomServices struct {
	roomsRepo rooms.RoomRepository
}
// NewRoomServiceImpl will create new RoomService object
func NewRoomServiceImpl(RoomRepo rooms.RoomRepository) *RoomServices {
	return &RoomServices{roomsRepo: RoomRepo}
}

// Rooms returns list of all rooms
func (rs *RoomServices) Rooms() ([]entity.Room, []error) {
	fmt.Println("Retrieving Data From Room  Services...")
	roomsList, err := rs.roomsRepo.Rooms()

	if len( err) >0 {
		fmt.Println("Error Room  Services...")
		return nil, err
	}

	return roomsList, nil
}

// StoreRoom persists new room information
func (rs *RoomServices) StoreRoom(room entity.Room) (*entity.Room, []error) {
	roomstored,err:= rs.roomsRepo.StoreRoom(room)

	if len( err) >0  {
		return nil,err
	}

	return roomstored,nil
}

// Room returns a room object with a given id
func (rs *RoomServices) Room(id int) (*entity.Room, []error){

	r, err := rs.roomsRepo.Room(id)

	if len( err) >0 {
		return r, err
	}

	return r, nil
}

// UpdateRoom updates a cateogory with new data
func (rs *RoomServices) UpdateRoom(room *entity.Room) (*entity.Room, []error) {

	uprm,err := rs.roomsRepo.UpdateRoom(room)

	if len( err) >0  {
		return nil, err
	}

	return uprm,nil
}
func (rs *RoomServices) DeleteRoom(id int) (*entity.Room, []error) {
	room,err := rs.roomsRepo.DeleteRoom(id)
	if len( err) >0 {
		return nil,err
	}
	return room,nil
}
func (rs *RoomServices) RoomTypes() ([]entity.Type, []error){
	fmt.Println("Retrieving Data From Room Types...")
	roomtypes, err := rs.roomsRepo.RoomTypes()

	if len( err) >0 {
		fmt.Println("Error Room Type Services...")
		return nil, err
	}

	return roomtypes, nil
}
func (rs *RoomServices) RoomType(id int) (*entity.Type, []error){
	r, errs := rs.roomsRepo.RoomType(id)

	if len(errs) >0 {
		for i := 0; i < len(errs); i++ {
			fmt.Println(errs[i])
		}
		return r, errs
	}
	return r, nil
}
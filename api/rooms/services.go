package rooms

import "github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"

// RoomServices specifies room menu services
type RoomServices interface {
	Rooms() ([]entity.Room, []error)
	Room(id int) (*entity.Room, []error)
	UpdateRoom(room *entity.Room) (*entity.Room, []error)
	DeleteRoom(id int) (*entity.Room, []error)
	StoreRoom(room entity.Room) (*entity.Room, []error)
	RoomTypes() ([]entity.Type, []error)
	RoomType(id int) (*entity.Type, []error)
}
package rooms

import "github.com/aleale2121/Hotel-Final/api/entity"

// RoomRepository specifies room related database operations
type RoomRepository interface {
	Rooms() ([]entity.Room, []error)
	Room(id int) (*entity.Room, []error)
	UpdateRoom(room *entity.Room) (*entity.Room, []error)
	DeleteRoom(id int) (*entity.Room, []error)
	StoreRoom(category entity.Room) (*entity.Room, []error)
	RoomTypes() ([]entity.Type, []error)
	RoomType(id int) (*entity.Type, []error)
}
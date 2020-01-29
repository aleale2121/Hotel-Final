package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/aleale2121/Hotel-Final/api/entity"
)

type RoomGormRepo struct {
 conn *gorm.DB
}

func NewRoomGormRepo(db *gorm.DB) *RoomGormRepo {
	return &RoomGormRepo{conn: db}
}
func (r *RoomGormRepo) Rooms() ([]entity.Room, []error) {
	fmt.Println("Retrieving Data From Rooms")
	roomlist := []entity.Room{}
	errs := r.conn.Preload("Type").Preload("Orders").Find(&roomlist).GetErrors()
	if len(errs) > 0 {
		fmt.Println(len(errs))
		fmt.Println("From Gorm Room List")
		for i := 0; i < len(errs); i++ {
			fmt.Println(errs[i])
		}
		return nil, errs
	}
	return roomlist, errs
}

func (r *RoomGormRepo) Room(id int) (*entity.Room, []error) {
	room := entity.Room{}
	errs := r.conn.Preload("Type").Preload("Orders").First(&room, id).GetErrors()
	if len(errs) > 0 {

		return nil, errs
	}
	return &room, errs
}

func (r *RoomGormRepo) UpdateRoom(room *entity.Room) (*entity.Room, []error) {
	roomUpdated := room
	fmt.Println(roomUpdated.Id)
	//errs := r.conn.Save(&roomUpdated).GetErrors()
	rs := r.conn.Model(roomUpdated).Where("id = ?", roomUpdated.Id).UpdateColumns(
		map[string]interface{}{
			"room_number":  roomUpdated.RoomNumber,
			"price":  roomUpdated.Price,
			"description": roomUpdated.Description,
			"image":roomUpdated.Image,
		},
	)
	if rs.Error!=nil {
		err2:= []error{rs.Error}
		return nil, err2
	}
	return roomUpdated, nil
}

func (r *RoomGormRepo) DeleteRoom(id int) (*entity.Room, []error) {
	cat, errs := r.Room(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = r.conn.Delete(cat, cat.Id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cat, errs
}

func (r *RoomGormRepo) StoreRoom(room entity.Room)(*entity.Room, []error) {
	storeroom := room
	errs := r.conn.Create(&storeroom).GetErrors()
	if len(errs) > 0 {
		fmt.Println("Store Room Error")
		r.conn.Debug().Create(storeroom)
		for _, err := range errs {
			fmt.Println(err)
		}
		return nil, errs
	}
	return &storeroom, errs
}
func (r *RoomGormRepo) RoomTypes() ([]entity.Type, []error){
	fmt.Println("Retrieving Data From Room Types...")
	var roomtypes []entity.Type
	errs := r.conn.Find(&roomtypes).GetErrors()
	if len(errs) > 0 {
		fmt.Println(len(errs))
		fmt.Println("From Gorm Room Types Rooms...")
		for i := 0; i < len(errs); i++ {
			fmt.Println(errs[i])
		}
		return nil, errs
	}

	return roomtypes, errs
}
func (r *RoomGormRepo) RoomType(id int) (*entity.Type, []error){
	roomcat := entity.Type{}
	errs := r.conn.Preload("Room").First(&roomcat, id).GetErrors()
	if len(errs) > 0 {
		fmt.Println("Form GORM ROOM TYPE Error  ")
		return nil, errs
	}

	return &roomcat, errs
}
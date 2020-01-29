package repository
//
//import (
//	"errors"
//
//	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"
//	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/rooms"
//	"github.com/jinzhu/gorm"
//)
//
//// MockEventsRepo implements the menu.CategoryRepository interface
//type MockRoomsRepo struct {
//	conn *gorm.DB
//}
//// NewMockEventsRepo will create a new object of MockEventsRepo
//func NewMockRoomsRepo(db *gorm.DB) rooms.RoomRepository {
//	return &MockRoomsRepo{conn: db}
//}
//func (mCatRepo *MockRoomsRepo) Rooms() ([]entity.Room, []error) {
//	ctgs := []entity.Room{entity.RoomsMock}
//	return ctgs, nil
//}
//
//func (mCatRepo *MockRoomsRepo) Room(id int) (*entity.Room, []error) {
//	ctg := entity.RoomsMock
//	if id == 1 {
//		return &ctg, nil
//	}
//	return nil, []error{errors.New("Not found")}
//}
//
//func (mCatRepo *MockRoomsRepo) UpdateRoom(room *entity.Room) (*entity.Room, []error) {
//	cat := entity.RoomsMock
//	return &cat, nil
//}
//
//func (mCatRepo *MockRoomsRepo) DeleteRoom(id int) (*entity.Room, []error) {
//	cat := entity.RoomsMock
//	if id != 1 {
//		return nil, []error{errors.New("Not found")}
//	}
//	return &cat, nil
//}
//
//func (mCatRepo *MockRoomsRepo) StoreRoom(rooms entity.Room) (*entity.Room, []error) {
//	cat := rooms
//	return &cat, nil
//}
//
//func (mCatRepo *MockRoomsRepo) RoomTypes() ([]entity.Type, []error) {
//	panic("implement me")
//}
//
//func (mCatRepo *MockRoomsRepo) RoomType(id int) (*entity.Type, []error) {
//	panic("implement me")
//}
//
//

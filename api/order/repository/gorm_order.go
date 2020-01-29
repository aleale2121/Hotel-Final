package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/aleale2121/Hotel-Final/api/entity"
	"github.com/aleale2121/Hotel-Final/api/order"
	"github.com/aleale2121/Hotel-Final/api/utils"
)

// OrderGormRepo implements the menu.repository interface
type OrderGormRepo struct {
	conn *gorm.DB
}

// NewOrderGormRepo returns new object of OrderGormRepo
func NewOrderGormRepo(db *gorm.DB) order.OrderRepository {
	return &OrderGormRepo{conn: db}
}

// Orders returns all customer orders stored in the database
func (orderRepo *OrderGormRepo) Orders() ([]entity.Order, error) {
	orders := []entity.Order{}
	errs := orderRepo.conn.Preload("User").Preload("Room").Find(&orders).Error
	if errs!=nil {
		return orders, utils.ErrInternalServerError
	}
	return orders, errs
}

// Order retrieve customer order by order id
func (orderRepo *OrderGormRepo) Order(id uint32) (*entity.Order, error) {
	order := entity.Order{}
	errs := orderRepo.conn.Preload("User").Preload("Room").First(&order, id).Error
	if errs!=nil{
		return &order, utils.ErrInternalServerError
	}
	return &order, errs
}

// UpdateOrder updates a given customer order in the database
func (orderRepo *OrderGormRepo) UpdateOrder(order *entity.Order) (*entity.Order, error) {
	ordr := order
	errs := orderRepo.conn.Save(ordr).Error
	if errs!=nil {
		return ordr, utils.ErrInternalServerError
	}
	return ordr, errs
}

// DeleteOrder deletes a given order from the database
func (orderRepo *OrderGormRepo) DeleteOrder(id uint32) (*entity.Order, error) {
	ordr, errs := orderRepo.Order(id)

	if errs!=nil{
		return nil, utils.ErrInternalServerError
	}
	errs = orderRepo.conn.Delete(ordr, id).Error
	if errs!=nil {
		return nil, utils.ErrInternalServerError
	}
	return ordr, errs
}

// StoreOrder stores a given order in the database
func (orderRepo *OrderGormRepo) StoreOrder(order *entity.Order) (*entity.Order, error) {
	ordr := order

	errs := orderRepo.conn.Create(&ordr).Error
	if errs!=nil{
		println("Store Order Gorm Exception")
		return nil, utils.ErrInternalServerError
	}
	fmt.Println("Successfully Stored  --- From Store order repository")
	return ordr, nil
}

// CustomerOrders returns list of orders from the database for a given customer
func (orderRepo *OrderGormRepo) CustomerOrders(customerid int32) ([]entity.Order, error) {
	var custOrders []entity.Order
	errs := orderRepo.conn.Preload("User").Preload("Room").Where("user_id = ?",customerid).Find(&custOrders).Error
	if errs!=nil {
		return custOrders, utils.ErrInternalServerError
	}
	return custOrders, errs
}
//orders by the room id
func (orderRepo *OrderGormRepo)  RoomOrder(id uint32)([]entity.Order, error){
	var roomOrders []entity.Order
	errs := orderRepo.conn.Preload("User").Preload("Room").Where("room_id = ?",id).Find(&roomOrders).Error
	if errs!=nil {
		return roomOrders, utils.ErrInternalServerError
	}
	return roomOrders, errs
}

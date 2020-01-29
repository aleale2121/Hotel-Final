package services
import (
	"fmt"
 "github.com/aleale2121/Hotel-Final/api/entity"
 "github.com/aleale2121/Hotel-Final/api/order"
)

// OrderService implements menu.OrderService interface
type OrderService struct {
	orderRepo order.OrderRepository
}

// NewOrderService returns new OrderService object
func NewOrderService(orderRepository order.OrderRepository) order.OrderService {
	return &OrderService{orderRepo: orderRepository}
}

// Orders returns all stored food orders
func (os *OrderService) Orders() ([]entity.Order, error) {
	ords, errs := os.orderRepo.Orders()
	if errs!=nil{
		return ords, errs
	}
	return ords, nil
}

// Order retrieves an order by its id
func (os *OrderService) Order(id uint32) (*entity.Order,error) {
	ord, errs := os.orderRepo.Order(id)
	if errs!=nil {
		return ord, errs
	}
	return ord, nil
}

// CustomerOrders returns all orders of a given customer
func (os *OrderService) CustomerOrders(customerid int32) ([]entity.Order, error) {
	ords, errs := os.orderRepo.CustomerOrders(customerid)
	if errs!=nil {
		return ords, errs
	}
	return ords, nil
}

// UpdateOrder updates a given order
func (os *OrderService) UpdateOrder(order *entity.Order) (*entity.Order, error) {
	ord, errs := os.orderRepo.UpdateOrder(order)
	if errs!=nil {
		return ord, errs
	}
	return ord, nil
}

// DeleteOrder deletes a given order
func (os *OrderService) DeleteOrder(id uint32) (*entity.Order, error) {
	ord, errs := os.orderRepo.DeleteOrder(id)
	if errs!=nil {
		fmt.Println("Delete Room Services")
		return ord, errs
	}
	return ord, nil
}

// StoreOrder stores a given order
func (os *OrderService) StoreOrder(order *entity.Order) (*entity.Order, error) {
	ord, errs := os.orderRepo.StoreOrder(order)
	if errs!=nil {
		println("Store Order Service Exception")
		return ord, errs
	}
	fmt.Println("Successfully Stored  --- From Store order services")
	return ord, nil
}

func (os *OrderService)RoomOrder(id uint32)([]entity.Order, error) {
	ords, errs := os.orderRepo.RoomOrder(id)
	if errs!=nil {
		return ords, errs
	}
	return ords, nil
}

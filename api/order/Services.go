package order
import "github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"
// repository specifies customer menu order related database operations
type OrderService interface {
	Orders() ([]entity.Order, error)
	Order(id uint32) (*entity.Order, error)
	CustomerOrders(customerid int32) ([]entity.Order, error)
	UpdateOrder(order *entity.Order) (*entity.Order, error)
	DeleteOrder(id uint32) (*entity.Order, error)
	StoreOrder(order *entity.Order) (*entity.Order, error)
	RoomOrder(id uint32)([]entity.Order, error)
}

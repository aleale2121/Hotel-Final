package order
import "github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"
// repository specifies customer menu order related database operations
type UserRepository interface {
	Users() ([]entity.User, error)
	User(id uint32) (*entity.User, error)
	StoreUser(user *entity.User) (*entity.User, error)
	UpdateUser(order *entity.User) (*entity.User, error)
	DeleteUser(id uint32) (*entity.User, error)
	UserByUserName(user entity.User)(*entity.User, error)
	PhoneExists(phone string) bool
	EmailExists(email string) bool
	UserRoles(*entity.User) ([]entity.Role, []error)
}

// RoleRepository speifies application user role related database operations
type RoleRepository interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	RoleByName(name string) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}

// SessionRepository specifies logged in user session related database operations
type SessionRepository interface {
	Session(sessionID string) (*entity.Session, []error)
	StoreSession(session *entity.Session) (*entity.Session, []error)
	DeleteSession(sessionID string) (*entity.Session, []error)
}

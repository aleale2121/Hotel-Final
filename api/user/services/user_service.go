package services
import (
	"fmt"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"
	user"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/user"
)

type UserServices struct {
	userRepo user.UserRepository
}

func NewUserService(userRepository user.UserRepository) user.UserService {
	return &UserServices{userRepo: userRepository}
}
func (u *UserServices) Users() ([]entity.User, error) {
	users, errs := u.userRepo.Users()
	if errs!=nil{
		return users, errs
	}
	return users, nil
}

func (u *UserServices) User(id uint32) (*entity.User, error) {
	users, errs := u.userRepo.User(id)
	if errs!=nil{
		return users, errs
	}
	return users, nil
}

func (u *UserServices) UpdateUser(user *entity.User) (*entity.User, error) {
	user1, errs := u.userRepo.UpdateUser(user)
	if errs!=nil {
		return user1, errs
	}
	return user1, nil
}

func (u *UserServices) DeleteUser(id uint32) (*entity.User, error) {
	user1, errs := u.userRepo.DeleteUser(id)
	if errs!=nil {
		fmt.Println("Delete Room Services")
		return user1, errs
	}
	return user1, nil
}

func (u *UserServices) StoreUser(user *entity.User) (*entity.User, error) {
	us:=user
	user1, errs := u.userRepo.StoreUser(us)
	if errs!=nil {
		return  nil,errs
	}
	return user1, nil
}
func (u *UserServices) UserByUserName(user entity.User)(*entity.User, error){
	users, errs := u.userRepo.UserByUserName(user)
	if errs!=nil{
		return users, errs
	}
	fmt.Println("services--- ",*users)
	return users, nil
}


// PhoneExists check if there is a user with a given phone number
func (us *UserServices) PhoneExists(phone string) bool {
	exists := us.userRepo.PhoneExists(phone)
	return exists
}

// EmailExists checks if there exist a user with a given email address
func (us *UserServices) EmailExists(email string) bool {
	exists := us.userRepo.EmailExists(email)
	return exists
}

// UserRoles returns list of roles a user has
func (us *UserServices) UserRoles(user *entity.User) ([]entity.Role, []error) {
	userRoles, errs := us.userRepo.UserRoles(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}

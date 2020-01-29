package rate

import "github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"

type RateService interface {
	//rate methods
	AddRate(updaterate *entity.Rating) (*entity.Rating, []error)
	RetrieveHotelRateValue(id uint) (*entity.Rating, []error)
	GetAllRatings(updaterate *entity.Rating) ([]entity.Rating, []error)
	GetUserRateValue(user_id uint) (*entity.Rating, []error)
	UpdateUserRateValue(user_id uint, u int) ([]error)


}
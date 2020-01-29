package rate

import "github.com/aleale2121/Hotel-Final/entity"

type RateRepository interface {

	//rate methods

	AddRate(updaterate *entity.Rating) (*entity.Rating, []error)
	RetrieveHotelRateValue(id uint) (*entity.Rating, []error)
	GetAllRatings(updaterate *entity.Rating) ([]entity.Rating, []error)
	GetUserRateValue(user_id uint) (*entity.Rating, []error)
	UpdateUserRateValue(user_id uint, u int) ([]error)
	
}
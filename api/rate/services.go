package rate

import "github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"

type RateService interface {
	AddRate(updaterate *entity.Rating) (*entity.Rating, []error)
	//api
	Rate() ([]entity.Rating, []error)
	Rates(id uint) (*entity.Rating, []error)
	UpdateRates(rpay *entity.Rating) (*entity.Rating, []error)
	DeleteRates(id uint) (*entity.Rating, []error)
}
package services
import (
	"fmt"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/rate"
)


type RateService struct {
	rRepo rate.RateRepository
}

// NewRateService will create new RateService object
func NewRateService(rate rate.RateRepository) rate.RateService {
	return &RateService{rRepo: rate}
}
func (rservice *RateService) AddRate(updaterate *entity.Rating) (*entity.Rating, []error)  {
	fmt.Println("services class ")
	data,err := rservice.rRepo.AddRate(updaterate)
	if err != nil {
		return data,err
	}
	return data, nil
}



func (rservice *RateService) Rate() ([]entity.Rating, []error) {
	data := []entity.Rating{}
	data, errs:= rservice.rRepo.Rate()
	if len(errs) > 0 {
		return nil, errs
	}
	return data, errs
}
// Rates retrieves stored Rate by its id
func (rservice *RateService) Rates(id uint) (*entity.Rating, []error) {
	cmnt, errs := rservice.rRepo.Rates(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// UpdateRates updates a given Rates
func (rservice *RateService) UpdateRates(pay *entity.Rating) (*entity.Rating, []error) {
	cmnt, errs := rservice.rRepo.UpdateRates(pay)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
// DeleteRates deletes a given Rates
func (rservice *RateService) DeleteRates(id uint) (*entity.Rating, []error) {
	cmnt, errs := rservice.rRepo.DeleteRates(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

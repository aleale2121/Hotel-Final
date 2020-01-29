package services
import (
"fmt"
"github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"
"github.com/yuidegm/Hotel-Rental-Managemnet-System/rate"
)


type RateService struct {
	rRepo rate.RateRepository
}

// NewRateService will create new RateService object
func NewRateService(rate rate.RateRepository) rate.RateService {
	return &RateService{rRepo: rate}
}

func (rservice *RateService) GetAllRatings(updaterate *entity.Rating) ([]entity.Rating,[]error) {
	data,err := rservice.rRepo.GetAllRatings(updaterate)
	if len(err)>0{return nil,err}
	return data,nil
}

func (rservice *RateService) GetUserRateValue(user_id uint) (*entity.Rating,[]error) {
	data,err := rservice.rRepo.GetUserRateValue(user_id)
	return data,err
}

func (rservice *RateService) RetrieveHotelRateValue(id uint) (*entity.Rating, []error){
	val,err := rservice.rRepo.RetrieveHotelRateValue(id)
	if len(err) > 0 {
		return val, err
	}
	return val, nil
}
func (rservice *RateService) AddRate(updaterate *entity.Rating) (*entity.Rating, []error)  {
	fmt.Println("services class ")
	data,err := rservice.rRepo.AddRate(updaterate)
	if err != nil {
		return data,err
	}
	return data, nil
}


func (rservice *RateService) UpdateUserRateValue(user_id uint,u int) ([]error){
	errs := rservice.rRepo.UpdateUserRateValue(user_id,u)
	if len(errs) > 0 {
		return errs
	}
	return nil
}




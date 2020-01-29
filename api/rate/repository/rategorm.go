package repository

import (
"fmt"
"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"
"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/rate"
"github.com/jinzhu/gorm"
)

// PymentsGormRepo implements the menu.CategoryRepository interface
type RateGormRepo struct {
	Conn *gorm.DB
}
var dbconnect *gorm.DB

// NewCategoryGormRepo will create a new object of CategoryGormRepo
func NewRateGormRepo(db *gorm.DB) rate.RateRepository {
	return &RateGormRepo{Conn: db}
}
// AddRate updates a given  in the database
func (rRepo *RateGormRepo) AddRate(updaterate *entity.Rating) (*entity.Rating, []error) {
	ratevalue := updaterate
	fmt.Println("repository start")
	errs := rRepo.Conn.Create(ratevalue).GetErrors()
	//fmt.Println(errs,1000101001)
	if len(errs) > 0 {
		fmt.Println(errs)
	}
	fmt.Println("repository end")

	return ratevalue, errs
}



//Rate returns all rate values from dtatabases
func (rRepo *RateGormRepo) Rate() ([]entity.Rating, []error) {
	cmnts := []entity.Rating{}
	errs := rRepo.Conn.Find(&cmnts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

// Rates retrieves a customer rate values from the database by its id
func (rRepo *RateGormRepo) Rates(id uint) (*entity.Rating, []error) {
	cmnt := entity.Rating{}
	errs := rRepo.Conn.First(&cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cmnt, errs
}



// Updatepayments updates a given customer payments in the database
func (rRepo *RateGormRepo) UpdateRates(rpay *entity.Rating) (*entity.Rating, []error) {
	cmnt := rpay
	errs := rRepo.Conn.Save(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
// DeleteRates deletes a given customer rate value from the database
func (rRepo *RateGormRepo) DeleteRates(id uint) (*entity.Rating, []error) {
	cmnt, errs := rRepo.Rates(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = rRepo.Conn.Delete(cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}





















































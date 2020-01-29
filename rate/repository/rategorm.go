package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/aleale2121/Hotel-Final/entity"
	"github.com/aleale2121/Hotel-Final/rate"
)

// RateGormRepo implements rate.RateRepository  interface
type RateGormRepo struct {
	Conn *gorm.DB
}
//var dbconnect *gorm.DB

// NewRateGormRepo will create a new object of RateGormRepo
func NewRateGormRepo(db *gorm.DB) rate.RateRepository {
	return &RateGormRepo{Conn: db}
}
// RetrieveHotelRateValue retrieve a rate value from the database by its id
func (rRepo *RateGormRepo) RetrieveHotelRateValue(id uint) (*entity.Rating, []error) {
	ctg := entity.Rating{}
	errs := rRepo.Conn.First(&ctg, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &ctg, errs
}

// AddRate updates a given  in the database
func (rRepo *RateGormRepo) AddRate(rating *entity.Rating) (*entity.Rating, []error) {
	fmt.Println("repository start")
	errs := rRepo.Conn.Create(rating).GetErrors()
	//fmt.Println(errs,1000101001)
	if len(errs) > 0 {
		fmt.Println(errs)
	}
	fmt.Println("repository end")

	return rating, errs
}

func (rRepo *RateGormRepo) GetAllRatings(rating *entity.Rating) ([]entity.Rating,[]error){
	var ratings []entity.Rating
	rows,err:=rRepo.Conn.Model(&rating).Find(&ratings).Rows()
	defer rows.Close()
	if err!=nil{ panic(err)}
	return ratings,nil

}
func (rRepo *RateGormRepo) GetUserRateValue(user_id uint) (*entity.Rating,[]error){
	data:= entity.Rating{}
	err:=rRepo.Conn.Model(&data).Where("user_id=$1",user_id).Find(&data).GetErrors()
	if len(err)>0{
		return  nil,err
	}
	return  &data,nil
}

func (rRepo *RateGormRepo) UpdateUserRateValue(user_id uint,u int) ([]error) {
	user:=&entity.Rating{UserId: user_id,RateValue:u}
	errs:=rRepo.Conn.Model(user).Where("user_id=?",user.UserId).UpdateColumns(
		map[string]interface{}{
			"UserId": user.UserId,
			"RateValue": user.RateValue,
		})
	if errs.Error!=nil{
		panic(errs)
	}
	return nil
}



























































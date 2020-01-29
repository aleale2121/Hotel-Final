package repository

import (
	"fmt"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api_bank/bank"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api_bank/entity"
	"github.com/jinzhu/gorm"
)

// BankGormRepo implements the BankGormRepo interface
type BankGormRepo struct {
	Conn *gorm.DB
}

// NewBankGormRepo will create a new object of NewBankGormRepo
func NewBankGormRepo(db *gorm.DB) bank.Repository {
	return &BankGormRepo{Conn: db}
}


func  (bRepo *BankGormRepo) RetrieveAccountFromBank(id int64) (*entity.Customer, []error) {
	data:= &entity.Customer{}
	 fmt.Println("gorm id",id)
	fmt.Println("first step gorm",*data)
	errs := bRepo.Conn.Debug().Model(data).Where("account_number=$1",id).First(data).GetErrors()
	fmt.Println(errs)
	if len(errs)>0{
		//panic(errs)
		return nil,errs
	}
	fmt.Println("get",*data)


	return data, nil
}


func  (bRepo *BankGormRepo) UpdateUserAccount(userAccountNumber int64, userBalance float32) []error {
	user:=entity.Customer{AccountNumber: userAccountNumber,AccountBalance: userBalance}
	fmt.Println("gorm bank",user.AccountBalance)
	errs:=bRepo.Conn.Debug().Model(&user).UpdateColumn("account_balance", user.AccountBalance).GetErrors()
	if len(errs)>0{
		panic(errs)
	}
	return nil
}

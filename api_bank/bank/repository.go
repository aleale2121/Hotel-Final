package bank

import "github.com/yuidegm/Hotel-Rental-Managemnet-System/api_bank/entity"

type Repository interface {
	RetrieveAccountFromBank(id int64) (*entity.Customer, []error)
	UpdateUserAccount(userAccountNumber int64, userBalance float32) []error
}
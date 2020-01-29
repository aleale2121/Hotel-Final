package bank

import "github.com/aleale2121/Hotel-Final/api_bank/entity"

type Services interface {
	RetrieveAccountFromBank(id int64) (*entity.Customer, []error)
	UpdateUserAccount(userAccountNumber int64, userBalance float32) []error
}
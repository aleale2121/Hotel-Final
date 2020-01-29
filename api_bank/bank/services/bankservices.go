package services

import (
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api_bank/bank"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api_bank/entity"
)

// RoomServiceImpl implements rooms.RoomService interface
type BankService struct {
	bRepos bank.Repository
}

// NewBankService will create new NewBankService object
func NewBankService(bRepo bank.Repository) bank.Services {
	return &BankService{bRepos: bRepo}
}


// RetrieveAccountFromBank  retrieve account  FromBank
func (service *BankService) RetrieveAccountFromBank(id int64) (*entity.Customer, []error) {
	payments,err := service.bRepos.RetrieveAccountFromBank(id)
	if err != nil {
		return nil,err
	}
	return payments,nil
}



// UpdateUserAccount  updates User Account  account using their account number
func (service *BankService)  UpdateUserAccount(userAccountNumber int64, userBalance float32) []error {
	errs := service.bRepos.UpdateUserAccount(userAccountNumber, userBalance)
	if len(errs) > 0 {
		return errs
	}
	return nil
}


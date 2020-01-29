package entity

type Customer struct {
	AccountNumber   int64 `gorm:"primary_key" json:"account_number"`
	AccountBalance  float32 `json:"account_balance"`
}

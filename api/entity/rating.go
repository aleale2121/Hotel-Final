package entity

type Rating struct {
	UserId    uint `json:"user_id" gorm:"not null; unique;primary_key"`
	RateValue int   `json:"rate_value" gorm:"not null"`
}

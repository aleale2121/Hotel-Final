package entity

import "time"

type Comments struct {
	Id          int `json:"id"`
	Name      string `gorm:"type:varchar(255);not null"json:"name"`
	Email string `gorm:"type:varchar(255);not null" json:"email"`
	Comment       string `gorm:"type:varchar(255);not null" json:"Comments"`
	Publish  time.Time `gorm:"default:current_timestamp" json:"created_at"`
}
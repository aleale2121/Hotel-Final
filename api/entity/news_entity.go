package entity

import "time"

type News struct {
	Id          int `json:"id"`
	Header      string `gorm:"type:varchar(255);not null"json:"header"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	Image       string `gorm:"type:varchar(255);not null" json:"image"`
	Publish  time.Time `gorm:"default:current_timestamp" json:"created_at"`
}


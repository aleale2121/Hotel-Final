package entity

import "time"

type Events struct {
	Id          int `json:"id"`
	Header      string `gorm:"type:varchar(255);not null"json:"header"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	Image       string `gorm:"type:varchar(255);not null" json:"image"`
	Publish  time.Time `gorm:"default:current_timestamp" json:"created_at"`
}

//type EventsApi struct {
//	Id          int `json:"id"`
//	Header      string `json:"header"`
//	Description string `json:"description"`
//	Image       string `json:"image"`
//	//Publish     time.Time
//}
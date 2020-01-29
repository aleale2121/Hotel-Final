package entity

import "time"

type Order struct {
	Id          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	ArrivalDate time.Time `gorm:"default:current_timestamp" json:"arrival_date"`
	DepartureDate time.Time `gorm:"default:current_timestamp" json:"departure_date"`
	UserId     uint32 `gorm:"not null;" json:"user_id"`
	User User `json:"user"`
	RoomId      uint32 `gorm:"not null;" json:"room_id"`
	Room  Room `json:"room;auto_preload"`
	Adults      uint `json:"adults"`
	Child       uint `json:"child"`
}

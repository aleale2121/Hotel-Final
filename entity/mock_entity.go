package entity

import "time"

var EventsMock =Events{
	Id :       1,
	Header:    "Mock events 01",
	Description:  "two",
	Image :      "tutu.png",
	Publish:      time.Time{},
}
var NewsMock =News{
	Id :       1,
	Header:    "Mock News 01",
	Description:  "two",
	Image :      "tutu.png",
	Publish:      time.Time{},
}
var CommentsMock =Comments {
	Id   :1,
	Name    :"gech",
	Email :"getme@gmail.com",
	Comment  :"first comment",
	Publish  :time.Time{},
}

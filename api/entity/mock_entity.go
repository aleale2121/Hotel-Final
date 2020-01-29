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
	Header:    "Mock newss 01",
	Description:  "two newss",
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
//
//var MockType=Type {
//	Id :1,
//	RoomType :"one bed",
//	Rooms :[]Room{} ,
//}
//var RoomsMock =Room {
//	Id    :     1,
//	RoomNumber :1,
//	Price     :23,
//	TypeId    :1,
//	Type      :nil,
//	Description:"this is room one",
//	Image   :"room.png",
//	Orders   :[]Order{},
//}
//
//var OrderMock= Order {
//	Id   :1,
//	//ArrivalDate timestamp.Timestamp `json:"arrival_date"`
//	//DepartureDate timestamp.Timestamp `json:"departure_date"`
//	UserId  :1,
//	User   : nil,
//	RoomId  :1,
//	Room    :nil,
//	Adults  :1,
//	Child     :1,
//}
//
//var MockUser=User {
//	Id   :1,
//	UserName : "gech",
//	FullName :"getahun honelet",
//	Email    :"getbet04@gmail.com",
//	Phone    :"0949922604",
//	Password  :"0077",
//	Roles    :[]Role{},
//	Orders  :[]Order{},
//}
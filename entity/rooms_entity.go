package entity

type Type struct {
	Id uint32  `gorm:"primary_key;auto_increment" json:"id"`
	RoomType string `gorm:"type:varchar(255);not null" json:"room_type"`
	Rooms []Room `gorm:"ForeignKey:TypeId;auto_preload" json:"rooms"`
}
type Room struct {
	Id          uint32  `gorm:"primary_key;auto_increment" json:"id"`
	RoomNumber  int     `json:"room_number"`
	Price       float64  `json:"price"`
	TypeId  uint32 `gorm:"not null;auto_preload" json:"type_id"`
	Type    Type   `gorm:"auto_preload" json:"type"`
	Description string  `gorm:"type:text" json:"description"`
	Image       string  `gorm:"type:varchar(255)" json:"image"`
	Orders  []Order `gorm:"ForeignKey:RoomId;auto_preload" json:"orders"`
}




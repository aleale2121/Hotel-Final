package entity

type User struct {
	Id       uint32  `gorm:"primary_key;auto_increment js" json:"id"`
	FullName string `gorm:"type:varchar(255);not null" json:"full_name"`
	Email    string `gorm:"type:varchar(255);not null; unique" json:"email"`
	Phone    string `gorm:"type:varchar(100);not null; unique" json:"phone"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	RoleID   int    `json:"role_id"`
	Orders  []Order `gorm:"ForeignKey:UserId;auto_preload" json:"orders"`
}
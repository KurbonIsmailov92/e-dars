package models

type Student struct {
	ID      uint  `gorm:"primary_key" json:"id"`
	UserID  uint  `gorm:"not null" json:"user_id"`
	User    User  `gorm:"foreignKey:UserID;references:ID" json:"user"`
	RoleID  uint  `gorm:"not null" json:"role_id"`
	Role    Role  `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	GroupID uint  `gorm:"not null" json:"group_id"`
	Group   Group `gorm:"foreignKey:GroupID;references:ID" json:"group"`
}

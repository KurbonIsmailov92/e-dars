package models

type Teacher struct {
	ID     uint   `gorm:"primary_key;auto_increment" json:"id"`
	UserID uint   `gorm:"user_id" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
	RoleID uint64 `gorm:"role_id" json:"role_id"`
	Role   Role   `gorm:"foreignKey:RoleID;references:ID" json:"role"`
}

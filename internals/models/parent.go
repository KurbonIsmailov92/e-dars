package models

type Parent struct {
	ID      uint      `gorm:"primaryKey;auto_increment" json:"id"`
	UserID  uint      `gorm:"not null" json:"user_id"`
	User    User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	RoleID  uint      `gorm:"not null" json:"role_id"`
	Role    Role      `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	Student []Student `gorm:"many2many: parents_students" json:"student"`
}

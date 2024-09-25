package models

import "time"

type User struct {
	ID            uint      `gorm:"primary_key;auto_increment" json:"id"`
	FullName      string    `gorm:"size: 40;" json:"full_name"`
	Username      string    `gorm:"size:20;unique;" json:"username"`
	Password      string    `gorm:"size:255;" json:"password"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:NULL" json:"updated_at"`
	DeletedAt     time.Time `gorm:"default:NULL" json:"deleted_at"`
	DeactivatedAt time.Time `gorm:"default:NULL" json:"deactivated_at"`
	IsActive      bool      `gorm:"default:true" json:"is_active"`
	IsDeleted     bool      `gorm:"default:false" json:"is_deleted"`
	Email         string    `gorm:"size:255;unique;" json:"email"`
	Phone         *string   `gorm:"size:20;default null;" json:"phone"`
	RoleCode      string    `gorm:"role_code" json:"role_code"`
	Role          Role      `gorm:"foreignKey:RoleCode" json:"role"`
	GroupID       *uint     `gorm:"default null" json:"group_id"`
	Group         Group     `gorm:"foreignKey:GroupID" json:"group"`
	ParentID      *uint     `gorm:"default:NULL" json:"parent_id"`
}

type SwagUser struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type SwagUserForShow struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type SignInInput struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
}

type UserPassword struct {
	OldPassword string `json:"old_password"`
	Password    string `json:"password" gorm:"not null"`
}

type SwagUserForUpdateByAdmin struct {
	FullName  string `json:"full_name"`
	Username  string `json:"username" gorm:"unique"`
	RoleCode  string `json:"role_code" gorm:"not null"`
	Email     string `json:"email" gorm:"size:255;unique;"`
	Phone     string `json:"phone" gorm:"size:20;"`
	IsActive  bool   `json:"is_active" gorm:"default:true"`
	IsDeleted bool   `json:"is_deleted" gorm:"default:false"`
	GroupID   *uint  `json:"group_id"`
}

package models

import "time"

type Class struct {
	ID              uint      `gorm:"primary_key; auto_increment" json:"id"`
	Name            string    `gorm:"size:30; not null" json:"name"`
	Description     string    `gorm:"size:255; not null" json:"desc"`
	ClassroomNumber int       `gorm:"not null" json:"classroom_number"`
	IsDeleted       bool      `gorm:"default:false" json:"is_deleted"`
	DeletedAt       time.Time `gorm:"default:null" json:"deleted_at"`
	Teacher         []User    `gorm:"many2many:class_users;" json:"teacher"`
}

type SwagClass struct {
	Name            string `gorm:"size:30; not null" json:"name"`
	Description     string `gorm:"size:255; not null" json:"desc"`
	ClassroomNumber int    `gorm:"not null" json:"classroom_number"`
}

type SwagClassInfo struct {
	Name            string `gorm:"size:30; not null" json:"name"`
	Description     string `gorm:"size:255; not null" json:"desc"`
	ClassroomNumber int    `gorm:"not null" json:"classroom_number"`
	Teacher         []User `gorm:"many2many:class_users;" json:"teacher"`
}

type ClassUser struct {
	ClassID uint `json:"class_id"`
	UserID  uint `json:"user_id"`
}

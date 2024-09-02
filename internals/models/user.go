package models

type User struct {
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	FullName string `gorm:"full_name" json:"full_name"`
}

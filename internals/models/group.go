package models

type Group struct {
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	CodeName string `gorm:"size:3;not null" json:"code_name"`
}

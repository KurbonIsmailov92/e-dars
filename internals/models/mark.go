package models

type Mark struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Code string `gorm:"size:1;" json:"code"`
}

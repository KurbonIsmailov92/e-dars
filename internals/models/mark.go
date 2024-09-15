package models

type Mark struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Code string `gorm:"not null;" json:"code"`
}

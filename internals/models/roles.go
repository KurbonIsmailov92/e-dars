package models

type Role struct {
	Code string `gorm:"primaryKey;unique" json:"code"`
	Name string `gorm:"size:20;" json:"name"`
}

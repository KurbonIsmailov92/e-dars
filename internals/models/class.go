package models

type Class struct {
	ID              uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name            string `gorm:"size:255;not null" json:"name"`
	Description     string `gorm:"size:255;not null" json:"desc"`
	ClassroomNumber int    `json:"classroomNumber"`
}

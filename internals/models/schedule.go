package models

import "time"

type Schedule struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	GroupID     uint      `gorm:"not null" json:"group_id"`
	Group       Group     `gorm:"foreignKey:GroupID;references:ID" json:"group"`
	ClassID     uint      `gorm:"not null" json:"class_id"`
	Class       Class     `gorm:"foreignKey:ClassID;references:ID" json:"class"`
	PlannedDate time.Time `gorm:"not null" json:"planned_date"`
	IsExam      bool      `gorm:"not null;default:false" json:"is_exam"`
}

type SwagSchedule struct {
	GroupID     uint      `gorm:"not null" json:"group_id"`
	ClassID     uint      `gorm:"not null" json:"class_id"`
	PlannedDate time.Time `gorm:"not null" json:"planned_date"`
	IsExam      bool      `gorm:"not null;default:false" json:"is_exam"`
}

type SwagScheduleToShow struct {
	ID          uint      `gorm:"nut null" json:"id"`
	Group       string    `gorm:"not null" json:"group"`
	Class       string    `gorm:"not null" json:"class"`
	PlannedDate time.Time `gorm:"not null" json:"planned_date"`
	IsExam      bool      `gorm:"not null;default:false" json:"is_exam"`
}

type SwagScheduleForUsers struct {
	Date      time.Time `gorm:"not null" json:"date"`
	Class     string    `gorm:"not null" json:"class"`
	ClassRoom string    `gorm:"not null" json:"class_room"`
	Group     string    `gorm:"not null" json:"group"`
	Teacher   string    `gorm:"not null" json:"teacher"`
	IsExam    bool      `gorm:"not null;default:false" json:"is_exam"`
}

type ScheduleDates struct {
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
}

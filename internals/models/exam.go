package models

import "time"

type Exam struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	Date         time.Time `gorm:"not_null" json:"date"`
	MarkID       uint      `gorm:"not_null;column:mark_id" json:"mark_id"`
	Mark         Mark      `gorm:"foreignKey:MarkID;references:ID" json:"mark"`
	ScheduleID   uint      `gorm:"not_null;column:schedule_note_id" json:"schedule_id"`
	ScheduleNote Schedule  `gorm:"foreignKey:ScheduleID;references:ID" json:"schedule_note"`
	Student      []User    `gorm:"many2many:exam_users" json:"student"`
}

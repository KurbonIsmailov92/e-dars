package models

import "time"

type JournalNotes struct {
	ID           uint      `gorm:"primary_key;auto_increment" json:"id"`
	Date         time.Time `gorm:"not null" json:"date"`
	MarkID       uint      `gorm:"default:null" json:"mark_id"`
	Mark         Mark      `gorm:"foreignKey:MarkID;references:ID" json:"mark"`
	StudentID    uint      `gorm:"not null" json:"student_id"`
	Student      Student   `gorm:"foreignKey:StudentID;references:ID" json:"student"`
	ScheduleID   uint      `gorm:"not null" json:"schedule_id"`
	ScheduleNote Schedule  `gorm:"foreignKey:ScheduleID;references:ID" json:"schedule_note"`
}

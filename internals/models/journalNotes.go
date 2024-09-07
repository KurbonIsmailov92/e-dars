package models

import "time"

type JournalNotes struct {
	ID           uint      `gorm:"primary_key;auto_increment" json:"id"`
	Date         time.Time `gorm:"not null" json:"date"`
	MarkID       uint      `gorm:"default:null" json:"mark_id"`
	Mark         Mark      `gorm:"foreignKey:MarkID;references:ID" json:"mark"`
	UserID       uint      `gorm:"not null" json:"user_id"`
	Student      User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	ScheduleID   uint      `gorm:"not null" json:"schedule_id"`
	ScheduleNote Schedule  `gorm:"foreignKey:ScheduleID;references:ID" json:"schedule_note"`
}

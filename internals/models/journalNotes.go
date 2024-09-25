package models

import "time"

type JournalNote struct {
	ID           uint      `gorm:"primary_key;auto_increment" json:"id"`
	Date         time.Time `gorm:"not null" json:"date"`
	MarkID       *uint     `gorm:"default:null" json:"mark_id"`
	Mark         Mark      `gorm:"foreignKey:MarkID;references:ID" json:"mark"`
	MarkedAt     time.Time `gorm:"default:null" json:"marked_at"`
	UserID       *uint     `gorm:"default null" json:"user_id"`
	Student      User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	ScheduleID   *uint     `gorm:"not null" json:"schedule_id"`
	ScheduleNote Schedule  `gorm:"foreignKey:ScheduleID;references:ID" json:"schedule_note"`
}

type SwagJournalNote struct {
	ScheduleID *uint `gorm:"default:null" json:"schedule_id"`
}

type SwagJournalNotesOfChildren struct {
	Date        time.Time `gorm:"not null" json:"date"`
	Group       string    `gorm:"not null" json:"group"`
	StudentName string    `gorm:"not null" json:"student_name"`
	Class       string    `gorm:"not null" json:"class"`
	ClassRoom   string    `gorm:"not null" json:"class_room"`
	Mark        string    `gorm:"default null" json:"mark"`
	TeacherName string    `gorm:"not null" json:"teacher_name"`
}

type JournalDates struct {
	DateFrom string `gorm:"not null" json:"date_from"`
	DateTo   string `gorm:"not null" json:"date_to"`
}

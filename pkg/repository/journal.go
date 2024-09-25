package repository

import (
	"e-dars/internals/db"
	"e-dars/internals/models"
	"e-dars/logger"
)

func CreateJournalNote(note models.JournalNote) error {
	if err := db.GetDBConnection().
		Model(models.JournalNote{}).
		Save(&note).Error; err != nil {
		logger.Error.Printf("Error creating journal note: %v", err)
		return err
	}
	return nil
}

func GetAllJournalNotes() (notes []models.JournalNote, err error) {
	err = db.GetDBConnection().
		Preload("Mark").
		Preload("Student").
		Preload("ScheduleNote").
		Find(&notes).
		Error
	if err != nil {
		logger.Error.Printf("Failed to get all journal notes: %v", err)
		return notes, translateError(err)
	}
	return notes, nil
}

func GetJournalNoteByID(id uint) (note models.SwagJournalNotes, err error) {
	if err = db.GetDBConnection().
		Raw(db.DBGetJournalNotesByID, id).
		Scan(&note).
		Error; err != nil {
		logger.Error.Printf("Failed to get journal note: %v", err)
		return note, translateError(err)
	}

	return note, nil
}

func GetJournalNotesByParentIDAndDate(id uint, dates models.JournalDates) (notes []models.SwagJournalNotes, err error) {
	err = db.GetDBConnection().
		Raw(db.DBGetChildJournalNotesByDates, id, dates.DateFrom, dates.DateTo).
		Scan(&notes).
		Error

	if err != nil {
		logger.Error.Printf("Failed to get journal notes for parent: %v", err)
		return notes, translateError(err)
	}
	return notes, nil
}

func GetJournalNotesByStudent(studentID uint, dates models.JournalDates) (notes []models.SwagJournalNotes, err error) {
	err = db.GetDBConnection().
		Raw(db.DBGetOwnJournalNotesByDates, studentID, dates.DateFrom, dates.DateTo).
		Scan(&notes).
		Error

	if err != nil {
		logger.Error.Printf("Failed to get journal notes for student: %v", err)
		return notes, translateError(err)
	}
	return notes, nil
}

func GetJournalNotesByTeacher(teacherID uint, dates models.JournalDates) (notes []models.SwagJournalNotes, err error) {
	err = db.GetDBConnection().
		Raw(db.DBGetJournalNotesByTeacherAndDates, teacherID, dates.DateFrom, dates.DateTo).
		Scan(&notes).
		Error

	if err != nil {
		logger.Error.Printf("Failed to get journal notes for teacher: %v", err)
		return notes, translateError(err)
	}
	return notes, nil
}

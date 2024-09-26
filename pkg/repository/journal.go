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
		logger.Error.Printf("[repository.CreateJournalNote] Error: creating journal note: %v", err)
		return translateError(err)
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
		logger.Error.Printf(" [repository.GetAllJournalNotes] Failed to get all journal notes: %v", err)
		return notes, translateError(err)
	}
	return notes, nil
}

func GetJournalNoteByID(id uint) (note models.SwagJournalNotes, err error) {
	if err = db.GetDBConnection().
		Raw(db.GetJournalNotesByIDDB, id).
		Scan(&note).
		Error; err != nil {
		logger.Error.Printf("[repository.GetJournalNoteByID] Failed to get journal note: %v", err)
		return note, translateError(err)
	}

	return note, nil
}

func GetJournalNotesByParentIDAndDate(id uint, dates models.JournalDates) (notes []models.SwagJournalNotes, err error) {
	err = db.GetDBConnection().
		Raw(db.GetChildJournalNotesByDatesDB, id, dates.DateFrom, dates.DateTo).
		Scan(&notes).
		Error

	if err != nil {
		logger.Error.Printf("[repository.GetJournalNotesByParentIDAndDate] Failed to get journal notes for parent: %v", err)
		return notes, translateError(err)
	}
	return notes, nil
}

func GetJournalNotesByStudent(studentID uint, dates models.JournalDates) (notes []models.SwagJournalNotes, err error) {
	err = db.GetDBConnection().
		Raw(db.GetOwnJournalNotesByDatesDB, studentID, dates.DateFrom, dates.DateTo).
		Scan(&notes).
		Error

	if err != nil {
		logger.Error.Printf("[repository.GetJournalNotesByStudent] Failed to get journal notes for student: %v", err)
		return notes, translateError(err)
	}
	return notes, nil
}

func GetJournalNotesByTeacher(teacherID uint, dates models.JournalDates) (notes []models.SwagJournalNotes, err error) {
	err = db.GetDBConnection().
		Raw(db.GetJournalNotesByTeacherAndDatesDB, teacherID, dates.DateFrom, dates.DateTo).
		Scan(&notes).
		Error

	if err != nil {
		logger.Error.Printf("[repository.GetJournalNotesByTeacher] Failed to get journal notes for teacher: %v", err)
		return notes, translateError(err)
	}
	return notes, nil
}

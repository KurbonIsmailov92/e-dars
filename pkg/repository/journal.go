package repository

import (
	"e-dars/internals/db"
	"e-dars/internals/models"
	"e-dars/logger"
	"time"
)

func CreateJournalNote(scheduleID uint, scheduleDate time.Time) error {
	if err := db.GetDBConnection().
		Exec(db.CreateNewJournalNoteDB, scheduleID, scheduleDate).
		Error; err != nil {
		logger.Error.Printf("Failed to create journal note: %v", err)
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
		logger.Error.Printf("Failed to get all journal notes: %v", err)
		return notes, translateError(err)
	}
	return notes, nil
}

func GetJournalNoteByID(id uint) (note models.SwagJournalNotesOfChildren, err error) {
	if err = db.GetDBConnection().
		Raw(db.GetJournalNotesByID, id).
		Scan(&note).
		Error; err != nil {
		logger.Error.Printf("Failed to get journal note: %v", err)
		return note, translateError(err)
	}

	return note, nil
}

func SetMark() {

}

func GetJournalNotesByParentIDAndDate(id uint, dateFrom, dateTo string) (notes []models.SwagJournalNotesOfChildren, err error) {
	err = db.GetDBConnection().
		Raw(db.GetChildJournalNotesByDates, id, dateFrom, dateTo).
		Scan(&notes).
		Error

	if err != nil {
		logger.Error.Printf("Failed to get journal notes for parent: %v", err)
		return notes, translateError(err)
	}
	return notes, nil
}

package service

import (
	"e-dars/internals/models"
	"e-dars/pkg/repository"
)

func CreateJournalNote(note *models.JournalNote) error {

	scheduleNote, err := repository.GetScheduleNoteByID(*note.ScheduleID)
	if err != nil {
		return err
	}

	if err = repository.CreateJournalNote(*note.ScheduleID, scheduleNote.PlannedDate); err != nil {
		return err
	}
	return nil
}

func GetAllJournalNotes() ([]models.JournalNote, error) {
	var notes []models.JournalNote

	notes, err := repository.GetAllJournalNotes()
	if err != nil {
		return notes, err
	}

	return notes, nil
}

func GetJournalNoteByID(id uint) (note models.SwagJournalNotesOfChildren, err error) {
	if note, err = repository.GetJournalNoteByID(id); err != nil {
		return note, err
	}
	return note, nil
}

func GetJournalNotesByParentIDAndDate(id uint, dateFrom, dateTo string) (notes []models.SwagJournalNotesOfChildren, err error) {
	notes, err = repository.GetJournalNotesByParentIDAndDate(id, dateFrom, dateTo)
	if err != nil {
		return notes, err
	}
	return notes, nil
}

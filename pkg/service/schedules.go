package service

import (
	"e-dars/internals/models"
	"e-dars/pkg/repository"
)

func CreateNewScheduleNote(scheduleNote *models.Schedule) error {
	if err := repository.CreateNewScheduleNote(scheduleNote); err != nil {
		return err
	}

	scheduleNoteForJournal, err := repository.GetLastCreatedScheduleNote()
	if err != nil {
		return err
	}

	if err = repository.CreateJournalNote(scheduleNoteForJournal.ID,
		scheduleNoteForJournal.PlannedDate); err != nil {
		return err
	}
	return nil
}

func GetAllScheduleNotes() (notes []models.Schedule, err error) {
	if notes, err = repository.GetAllScheduleNotes(); err != nil {
		return nil, err
	}
	return notes, nil
}

func GetScheduleNoteByID(id uint) (note models.Schedule, err error) {
	note, err = repository.GetScheduleNoteByID(id)
	if err != nil {
		return note, err
	}
	return note, nil
}

func UpdateScheduleNoteByID(id uint, scheduleNote models.Schedule) (err error) {
	scheduleNoteFromDB, err := repository.GetScheduleNoteByID(id)
	if err = repository.UpdateScheduleNoteByID(id, scheduleNote, scheduleNoteFromDB); err != nil {
		return err
	}
	return nil
}

func DeleteScheduleNoteByID(id uint) (err error) {

	if _, err = repository.GetScheduleNoteByID(id); err != nil {
		return err
	}

	if err = repository.DeleteScheduleNoteByID(id); err != nil {
		return err
	}
	return nil
}

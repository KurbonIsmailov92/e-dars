package service

import (
	"e-dars/errs"
	"e-dars/internals/models"
	"e-dars/pkg/repository"
	"errors"
	"time"
)

func CreateNewScheduleNote(scheduleNote *models.Schedule) error {
	if err := repository.CreateNewScheduleNote(scheduleNote); err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	}

	return nil
}

func GetAllScheduleNotes() (notes []models.Schedule, err error) {
	if notes, err = repository.GetAllScheduleNotes(); err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return nil, err
	}
	return notes, nil
}

func GetScheduleNoteByID(id uint) (note models.Schedule, err error) {
	note, err = repository.GetScheduleNoteByID(id)
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return note, err
	}
	return note, nil
}

func UpdateScheduleNoteByID(id uint, scheduleNote models.Schedule) (err error) {
	scheduleNoteFromDB, err := repository.GetScheduleNoteByID(id)
	if err = repository.UpdateScheduleNoteByID(id, scheduleNote, scheduleNoteFromDB); err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	}
	return nil
}

func DeleteScheduleNoteByID(id uint) (err error) {

	var scheduleNote models.Schedule

	if scheduleNote, err = repository.GetScheduleNoteByID(id); err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	}

	if !scheduleNote.PlannedDate.Before(time.Now()) {
		err = errs.ErrDateIsPast
		return err
	}

	if err = repository.DeleteScheduleNoteByID(id); err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	}
	return nil
}

func GetTeacherScheduleByDates(id uint, dates models.ScheduleDates) (notes []models.SwagScheduleForUsers, err error) {
	if notes, err = repository.GetTeacherScheduleByDates(id, dates); err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return nil, err
	}
	return notes, nil
}
func GetStudentScheduleByDates(id uint, dates models.ScheduleDates) (notes []models.SwagScheduleForUsers, err error) {
	if notes, err = repository.GetStudentScheduleByDates(id, dates); err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return nil, err
	}
	return notes, nil
}

func GetParentScheduleByDates(id uint, dates models.ScheduleDates) (notes []models.SwagScheduleForUsers, err error) {
	if notes, err = repository.GetParentScheduleByDates(id, dates); err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return nil, err
	}
	return notes, nil
}

package repository

import (
	"e-dars/internals/db"
	"e-dars/internals/models"
	"e-dars/logger"
)

func CreateNewScheduleNote(scheduleNote *models.Schedule) (err error) {
	if err = db.GetDBConnection().
		Create(scheduleNote).Error; err != nil {
		logger.Error.Printf("[repository CreateNewScheduleNote] Failed to create new schedule note: %v", err)
	}
	return nil
}

func GetAllScheduleNotes() (schedules []models.Schedule, err error) {
	schedules = []models.Schedule{}
	err = db.GetDBConnection().
		Preload("Class").
		Preload("Group").
		Find(&schedules).Error
	if err != nil {
		logger.Error.Printf("[repository GetAllScheduleNotes] Error getting all notes: %v", err)
		return nil, err
	}
	return schedules, nil
}

func GetScheduleNoteByID(noteID uint) (note models.Schedule, err error) {
	err = db.GetDBConnection().
		Where("id = ?", noteID).
		Preload("Class").
		Preload("Group").
		First(&note).Error
	if err != nil {
		logger.Error.Printf("[repository.GetScheduleNoteByID] Error getting Schedule Note by ID: %v\n", err)
		return note, translateError(err)
	}
	return note, nil
}

func UpdateScheduleNoteByID(id uint, note, noteFromDB models.Schedule) (err error) {
	if err := db.GetDBConnection().
		Model(&noteFromDB).
		Updates(note).
		Where("id = ?", id).Error; err != nil {
		logger.Error.Printf("[repository UpdateScheduleNoteByID] Error updating Schedule Note: %v", err)
		return err
	}

	return nil
}

func DeleteScheduleNoteByID(id uint) error {
	err := db.GetDBConnection().
		Model(&models.Schedule{}).
		Where("id = ?", id).
		Delete(&models.Schedule{}).Error
	if err != nil {
		logger.Error.Printf("[repository DeleteScheduleNoteByID] Error deleting schedule note: %v", err)
		return translateError(err)
	}
	return nil
}

func GetLastCreatedScheduleNote() (schedule models.Schedule, err error) {
	err = db.GetDBConnection().
		Order("id DESC").
		First(&schedule).Error

	if err != nil {
		return models.Schedule{}, translateError(err)
	}
	return schedule, nil
}

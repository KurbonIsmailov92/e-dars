package service

import (
	"e-dars/internals/models"
	"e-dars/logger"
	"e-dars/pkg/repository"
	"errors"
	"time"
)

func CreateJournalNote(note *models.MarkSetter, id uint) error {
	var journalNote models.JournalNote
	var teacherID uint

	scheduleNote, err := repository.GetScheduleNoteByID(note.ScheduleNoteID)
	if err != nil {
		return err
	}
	student, err := repository.GetUserByID(note.StudentID)
	if err != nil {
		return err
	}
	if student.RoleCode != "student" {
		logger.Error.Printf("User %d is not a student", note.StudentID)
		return err
	}
	if teacherID, err = repository.GetTeacherIDFromDB(scheduleNote.ClassID); err != nil {
		return err
	}

	if id != teacherID {
		logger.Error.Printf("Teacher ID is not the same as active user ID")
		return errors.New("teacher is not the owner of the class")
	}

	journalNote = models.JournalNote{
		Date:       scheduleNote.PlannedDate,
		MarkID:     note.MarkID,
		UserID:     student.ID,
		ScheduleID: scheduleNote.ID,
		MarkedAt:   time.Now(),
	}

	if err = repository.CreateJournalNote(journalNote); err != nil {
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

func GetJournalNoteByID(id uint) (note models.SwagJournalNotes, err error) {
	if note, err = repository.GetJournalNoteByID(id); err != nil {
		return note, err
	}
	return note, nil
}

func GetJournalNotesByParentIDAndDate(id uint, dates models.JournalDates) (notes []models.SwagJournalNotes, err error) {
	notes, err = repository.GetJournalNotesByParentIDAndDate(id, dates)
	if err != nil {
		return notes, err
	}
	return notes, nil
}

func GetJournalNotesByStudent(studentID uint, dates models.JournalDates) (notes []models.SwagJournalNotes, err error) {
	notes, err = repository.GetJournalNotesByStudent(studentID, dates)
	if err != nil {
		return notes, err
	}
	return notes, nil
}

func GetJournalNotesByTeacher(studentID uint, dates models.JournalDates) (notes []models.SwagJournalNotes, err error) {
	notes, err = repository.GetJournalNotesByTeacher(studentID, dates)
	if err != nil {
		return notes, err
	}
	return notes, nil
}

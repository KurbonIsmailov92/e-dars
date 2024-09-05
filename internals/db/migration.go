package db

import "e-dars/internals/models"

func MigrateTables() error {
	err := session.AutoMigrate(models.Class{},
		models.Exam{},
		models.Group{},
		models.JournalNotes{},
		models.Mark{},
		models.Parent{},
		models.Schedule{},
		models.Student{},
		models.Teacher{},
		models.User{},
	)

	if err != nil {
		return err
	}
	return nil
}

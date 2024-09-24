package db

import "e-dars/internals/models"

func MigrateTables() error {
	err := dbSession.AutoMigrate(models.Class{},
		models.Group{},
		models.JournalNotes{},
		models.Mark{},
		models.Schedule{},
		models.User{},
	)

	if err != nil {
		return err
	}
	return nil
}

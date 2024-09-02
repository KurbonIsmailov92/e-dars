package db

import "e-dars/internals/models"

func MigrateTables() error {
	err := session.AutoMigrate(models.User{}, models.Role{}, models.Class{},
		models.Mark{})
	if err != nil {
		return err
	}
	return nil
}

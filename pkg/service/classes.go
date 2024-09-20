package service

import (
	"e-dars/errs"
	"e-dars/internals/models"
	"e-dars/pkg/repository"
	"errors"
)

func CreateNewClass(c *models.Class) error {
	classFromDB, err := repository.GetClassByName(c.Name)
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	}

	if classFromDB.ID > 0 {
		return errs.ErrUsernameUniquenessFailed
	}

	err = repository.CreateNewClass(c)
	if err != nil {
		return err
	}

	return nil
}

func GetAllClasses() (classes []models.Class, err error) {
	classes, err = repository.GetAllClasses()
	if err != nil {
		return nil, err
	}
	return classes, nil
}

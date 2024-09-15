package service

import (
	"e-dars/errs"
	"e-dars/internals/models"
	"e-dars/pkg/repository"
	"e-dars/utils"
	"errors"
)

func CreateNewUser(u *models.User) error {
	userFromDB, err := repository.GetUserByUsername(u.Username)
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	}

	if userFromDB.ID > 0 {
		return errs.ErrUsernameUniquenessFailed
	}

	u.Password = utils.GenerateHash(u.Password)
	err = repository.CreateNewUser(u)
	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id int) (user models.User, err error) {
	user, err = repository.GetUserByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

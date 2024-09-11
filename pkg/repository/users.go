package repository

import (
	"e-dars/errs"
	"e-dars/internals/db"
	"e-dars/internals/models"
	"e-dars/logger"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDbConnection().
		Where("username = ? AND password = ?", username, password).
		Preload("Role").
		First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, translateError(err)

	}
	fmt.Println("Role is:", user.Roles)
	return user, nil
}

func translateError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrRecordNotFound
	}

	return err
}

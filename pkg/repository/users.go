package repository

import (
	"e-dars/errs"
	"e-dars/internals/db"
	"e-dars/internals/models"
	"e-dars/logger"
	"errors"
	"gorm.io/gorm"
	"time"
)

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConnection().
		Where("username = ? AND password = ?", username, password).
		Preload("Role").
		First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, translateError(err)

	}
	return user, nil
}

func translateError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrRecordNotFound
	}

	return err
}

func GetUserByUsername(userName string) (user models.User, err error) {
	err = db.GetDBConnection().
		Where("username = ?", userName).
		First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
		return user, translateError(err)
	}
	return user, nil
}

func CreateNewUser(u *models.User) (err error) {
	if err = db.GetDBConnection().Create(&u).Error; err != nil {
		logger.Error.Println("[repository.CreateUser] cannot create user. Error is:", err.Error())
		return translateError(err)
	}

	return nil
}

func GetAllUsers() (users []models.User, err error) {
	users = []models.User{}
	err = db.GetDBConnection().
		Order("id").
		Preload("Role").
		Preload("Group").
		Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository] Error getting all users: %v", err)
		return nil, err
	}
	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	user = models.User{}
	err = db.GetDBConnection().
		Where("id = ?", id).
		Preload("Role").
		Preload("Group").
		First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository] Error getting user: %v", err)
		return user, err
	}
	return user, nil
}

func UpdateUser(id uint, user, existUser models.User) error {
	if err := db.GetDBConnection().
		Model(&existUser).
		Updates(user).
		Where("id = ?", id).Error; err != nil {
		logger.Error.Printf("[repository UpdateUser] Error updating user: %v", err)
		return err
	}

	return nil
}

func DeActiveUserByID(id int) error {
	if err := db.GetDBConnection().
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_active":      false,
			"deactivated_at": time.Now(),
		}).Error; err != nil {
		logger.Error.Printf("[repository] Error deactivating user: %v", err)
		return err
	}
	return nil
}

func ActiveUserByID(id int) error {
	if err := db.GetDBConnection().
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_active":      true,
			"deactivated_at": nil, //"0001-01-01 00:00:00",
		}).Error; err != nil {
		logger.Error.Printf("[repository] Error activating user: %v", err)
		return err
	}
	return nil
}

package service

import (
	"e-dars/errs"
	"e-dars/internals/models"
	"e-dars/logger"
	"e-dars/pkg/repository"
	"e-dars/utils"
	"errors"
	"os"
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

func GetUserByID(id uint) (user models.User, err error) {
	user, err = repository.GetUserByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(id uint, user models.User) error {
	existUser, err := repository.GetUserByID(id)
	if err != nil {
		return err
	}
	user.Username = existUser.Username
	user.Password = existUser.Password

	if err = repository.UpdateUser(id, user, existUser); err != nil {
		return err
	}
	return nil
}

func DeActiveUser(id uint) error {
	if err := repository.DeActiveUserByID(id); err != nil {
		return err
	}
	return nil
}

func ActivateUser(id uint) error {
	if err := repository.ActiveUserByID(id); err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint) error {
	if err := repository.DeleteUserByID(id); err != nil {
		return err
	}
	return nil
}

func ReturnUser(id uint) error {
	if err := repository.ReturnUserByID(id); err != nil {
		return err
	}
	return nil
}

func ResetUserPassToDefault(id uint) error {
	newPassword := utils.GenerateHash(os.Getenv("DEFAULT_USER_PASSWORD"))
	if err := repository.ResetUserPasswordToDefault(id, newPassword); err != nil {
		return err
	}
	return nil
}

func ChangeOwnPasswordByUser(id uint, newPassword, oldPassword string) error {
	user, err := GetUserByID(id)
	if err != nil {
		return err
	}

	oldPassword = utils.GenerateHash(oldPassword)

	if oldPassword != user.Password {
		logger.Error.Printf("[service ChangeOwnPasswordByUser] Incorrect old password")
		err = errors.New("incorrect old password")
		return err
	}

	newPassword = utils.GenerateHash(newPassword)

	if err = repository.ChangeOwnPasswordByUser(id, newPassword); err != nil {
		return err
	}

	return nil
}

func SetAdminRoleToUser(id uint) error {
	if err := repository.SetAdminRoleToUser(id); err != nil {
		return err
	}
	return nil
}

func SetParentToUser(userid, parenID uint) error {

	parent, err := repository.GetUserByID(parenID)
	if err != nil {
		return err
	}

	if parent.RoleCode != "parent" {
		return errors.New("user is not a parent")
	}

	if err := repository.SetParentToUser(userid, parenID); err != nil {
		return err
	}
	return nil
}

func SetRoleToUser(userid uint, roleCode string) error {
	if err := repository.SetRoleToUser(userid, roleCode); err != nil {
		return err
	}
	return nil
}

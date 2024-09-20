package service

import (
	"e-dars/errs"
	"e-dars/internals/models"
	"e-dars/logger"
	"e-dars/pkg/repository"
	"errors"
	"fmt"
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

func SetClassTeacher(classID, teacherID uint) error {
	teacher, err := repository.GetUserByID(teacherID)
	fmt.Println(teacher.Role.Name)
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	}
	if teacher.Role.Code != "teacher" {
		logger.Error.Printf("[service SetClassTeacher] User is not teacher")
		err = errs.ErrUserIsNotTeacher
		return err
	}

	class, err := repository.GetClassByID(classID)
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		logger.Error.Printf("[service SetClassTeacher] There is no such class")
		err = errs.ErrClassNotFound
		return err
	}

	if err := repository.SetClassTeacher(class.ID, teacher.ID); err != nil {
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

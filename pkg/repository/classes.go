package repository

import (
	"e-dars/internals/db"
	"e-dars/internals/models"
	"e-dars/logger"
)

func SetClassTeacher(classID, teacherID uint) error {
	if err := db.GetDBConnection().
		Table("class_users").
		Create(&models.ClassUser{ClassID: classID,
			UserID: teacherID}).Error; err != nil {
		logger.Error.Printf("[repository SetClassTeacher] Failed to set teacher for class: %v", err)
		return err
	}
	return nil
}

func CreateNewClass(class *models.Class) (err error) {
	tx := db.GetDBConnection().Begin()

	if err = tx.Create(&class).Error; err != nil {
		logger.Error.Println("[repository.CreateClass] Cannot create class. Error is:", err.Error())
		return translateError(err)
	}

	if len(class.Teacher) > 0 {
		if err = tx.Model(&class).Association("Teacher").Replace(class.Teacher); err != nil {
			tx.Rollback()
			logger.Error.Printf("[repository.CreateNewClass] Cannot associate teachers with Error: %v", err.Error())
			return translateError(err)
		}
	}

	if err = tx.Commit().Error; err != nil {
		logger.Error.Printf("[repository.CreateNewClass] Transaction commit failed: %v", err)
		return translateError(err)
	}

	logger.Info.Printf("[repository.CreateNewClass] Successfully created class with teachers")
	return nil
}

func GetClassByName(className string) (class models.Class, err error) {
	err = db.GetDBConnection().
		Where("name = ?", className).
		First(&class).Error
	if err != nil {
		logger.Error.Printf("[repository.GetClassByName] error getting class by name: %v\n", err)
		return class, translateError(err)
	}
	return class, nil
}

func GetAllClasses() (classes []models.Class, err error) {
	classes = []models.Class{}
	err = db.GetDBConnection().
		Order("id").
		Preload("Teacher").
		Find(&classes).Error
	if err != nil {
		logger.Error.Printf("[repository GetAllClasses] Error getting all classes: %v", err)
		return nil, err
	}
	return classes, nil
}

func GetClassByID(classID uint) (class models.Class, err error) {
	err = db.GetDBConnection().
		Where("id = ?", classID).
		First(&class).Error
	if err != nil {
		logger.Error.Printf("[repository.GetClassByID] Error getting class by name: %v\n", err)
		return class, translateError(err)
	}
	return class, nil
}

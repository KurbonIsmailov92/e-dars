package repository

import (
	"e-dars/internals/db"
	"e-dars/internals/models"
	"e-dars/logger"
)

func SetClassTeacher(classID uint, teacher []models.User) error {
	if err := db.GetDBConnection().
		Set("class_id = ?", classID).
		Set("user_id = ?", teacher).
		Error; err != nil {
		logger.Error.Printf("Error setting class teacher id to: %v", err)
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

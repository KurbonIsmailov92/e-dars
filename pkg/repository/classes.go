package repository

import (
	"e-dars/internals/db"
	"e-dars/internals/models"
	"e-dars/logger"
)

func CreateNewClass(c *models.Class) (err error) {
	if err = db.GetDBConnection().Create(&c).Error; err != nil {
		logger.Error.Println("[repository.CreateClass] cannot create class. Error is:", err.Error())
		return translateError(err)
	}

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

package db

import (
	"e-dars/internals/models"
	"fmt"
)

func InsertSeeds() error {
	var (
		roles  []models.Role
		marks  []models.Mark
		groups []models.Group
	)

	roles = append(roles, models.Role{Code: "admin", Name: "Администратор"})
	roles = append(roles, models.Role{Code: "teacher", Name: "Преподаватель"})
	roles = append(roles, models.Role{Code: "parent", Name: "Родитель"})
	roles = append(roles, models.Role{Code: "student", Name: "Ученик"})

	marks = append(marks, models.Mark{Code: "1"})
	marks = append(marks, models.Mark{Code: "2"})
	marks = append(marks, models.Mark{Code: "3"})
	marks = append(marks, models.Mark{Code: "4"})
	marks = append(marks, models.Mark{Code: "5"})
	marks = append(marks, models.Mark{Code: "N/A"})

	for i := 1; i <= 11; i++ {
		groups = append(groups, models.Group{CodeName: fmt.Sprintf("%dА", i)})
		groups = append(groups, models.Group{CodeName: fmt.Sprintf("%dБ", i)})
		groups = append(groups, models.Group{CodeName: fmt.Sprintf("%dВ", i)})
		groups = append(groups, models.Group{CodeName: fmt.Sprintf("%dГ", i)})
		groups = append(groups, models.Group{CodeName: fmt.Sprintf("%dД", i)})
	}

	if err := GetDBConnection().Save(&roles).Error; err != nil {
		return err
	}
	if err := GetDBConnection().Save(&marks).Error; err != nil {
		return err
	}
	if err := GetDBConnection().Save(&groups).Error; err != nil {
		return err
	}

	return nil
}

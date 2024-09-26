package controllers

import (
	"e-dars/errs"
	"e-dars/internals/models"
	"e-dars/logger"
	"e-dars/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateNewClass
// @Summary Create Class
// @Security ApiKeyAuth
// @Tags classes
// @Description Create new class
// @ID create-class
// @Accept json
// @Produce json
// @Param input body models.SwagClass true "New Class info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /classes/api/v1/ [post]
func CreateNewClass(c *gin.Context) {
	var class models.Class
	if err := c.BindJSON(&class); err != nil {
		handleError(c, err)
		return
	}

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)

		return
	}

	err := service.CreateNewClass(&class)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controller.CreateNewClass] Created New Class success]")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Class created successfully",
	})

}

// GetAllClasses
// @Summary Get All Classes
// @Security ApiKeyAuth
// @Tags classes
// @Description get list of all classes
// @ID get-all-classes
// @Produce json
// @Success 200 {array} models.SwagClassInfo
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /classes/api/v1/ [get]
func GetAllClasses(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	classes, err := service.GetAllClasses()
	if err != nil {
		handleError(c, err)
	}

	logger.Info.Printf("[controllers.GetAllClasses] Successfully got all classes")
	c.JSON(http.StatusOK, gin.H{"classes": classes})
}

// GetClassByID
// @Summary Get Class
// @Security ApiKeyAuth
// @Tags classes
// @Description get class by ID
// @ID get-class
// @Produce json
// @Param id path integer true "id of the class"
// @Success 200 {array} models.SwagClassInfo
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /classes/api/v1/{id} [get]
func GetClassByID(c *gin.Context) {
	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}

	class, err := service.GetClassByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetClassByID] Successfully got class by id %v", id)
	c.JSON(http.StatusOK, gin.H{"class": class})
}

// SetClassTeacher
// @Summary Set Class to Teacher
// @Security ApiKeyAuth
// @Tags classes
// @Description Set Class to Teacher
// @ID set-class-to-teacher
// @Accept json
// @Produce json
// @Param input body models.ClassUser true "Info for setting"
// @Success 200 {array} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /classes/api/v1/set/ [post]
func SetClassTeacher(c *gin.Context) {
	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	var set models.ClassUser

	if err := c.BindJSON(&set); err != nil {
		handleError(c, err)
		return
	}

	if err := service.SetClassTeacher(set.ClassID, set.UserID); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.SetClassTeacher] Successfully set class to teacher")
	c.JSON(http.StatusCreated, gin.H{"message": "Class set successfully to teacher"})

}

// UpdateClass
// @Summary Update Class
// @Security ApiKeyAuth
// @Tags classes
// @Description Update class by ID
// @ID update-class
// @Accept json
// @Produce json
// @Param id path integer true "id of the class"
// @Param input body models.SwagClass true "Update class data"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /classes/api/v1/update/{id} [put]
func UpdateClass(c *gin.Context) {
	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	var class models.Class
	if err := c.BindJSON(&class); err != nil {
		handleError(c, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.UpdateClass] Invalid class ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}

	class.ID = uint(id)

	if err = service.UpdateClass(uint(id), class); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.UpdateClass] Successfully updated class to %v", class)
	c.JSON(http.StatusOK, gin.H{"message": "Class updated successfully"})

}

// DeleteClass
// @Summary Delete class by ID
// @Security ApiKeyAuth
// @Tags classes
// @Description Delete class by ID
// @ID delete-class
// @Accept json
// @Produce json
// @Param id path integer true "id of the class"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /classes/api/v1/delete/{id} [delete]
func DeleteClass(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid class ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}
	if err = service.DeleteClass(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.DeleteClass] Successfully deleted class: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "Class Deleted!"})
}

// ReturnClass
// @Summary Return class by ID
// @Security ApiKeyAuth
// @Tags classes
// @Description Return class by ID
// @ID return-class
// @Accept json
// @Produce json
// @Param id path integer true "id of the class"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /classes/api/v1/return/{id} [delete]
func ReturnClass(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.ReturnClass] Invalid class ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}
	if err = service.ReturnClass(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.ReturnClass] Successfully returned class: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "Class Returned successfully!"})
}

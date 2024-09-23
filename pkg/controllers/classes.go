package controllers

import (
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
// @Description create new class
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if c.GetString(userRoleCtx) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to create a new class",
		})

		return
	}

	err := service.CreateNewClass(&class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

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
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to see all classes",
		})
		return
	}

	classes, err := service.GetAllClasses()
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"massage": "No classes found"})
	}

	logger.Info.Printf("[controllers] Successfully got all classes")
	c.JSON(http.StatusOK, gin.H{"classes": classes})
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
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to set classes to teachers",
		})
		return
	}

	var set models.ClassUser

	if err := c.BindJSON(&set); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	if err := service.SetClassTeacher(set.ClassID, set.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"massage": err.Error()})
		return
	}
	logger.Info.Printf("[controllers] Successfully set class to teacher")
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
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to update user",
		})
		return
	}

	var class models.Class
	if err := c.BindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid class ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	class.ID = uint(id)

	if err := service.UpdateClass(uint(id), class); err != nil {
		logger.Error.Printf("[controllers] Failed to update class %v: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	logger.Info.Printf("[controllers UpdateClass] Successfully updated class to %v", class)
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
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to delete class",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid class ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
	}
	if err = service.DeleteClass(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Class not found"})
		return
	}
	logger.Info.Printf("[controllers] Successfully deleted class: %v", id)
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
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to return class",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid class ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
	}
	if err = service.ReturnClass(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Class not found"})
		return
	}
	logger.Info.Printf("[controllers] Successfully returned class: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "Class Returned successfully!"})
}

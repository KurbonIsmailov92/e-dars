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
// @Router /classes [post]
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
// @Success 200 {array} models.Class
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /classes [get]
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

	logger.Info.Printf("[controllers] Successfully got all classes: %v", classes)
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
// @Success 200 {array} models.Class
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /classes/set [post]
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

	}

}

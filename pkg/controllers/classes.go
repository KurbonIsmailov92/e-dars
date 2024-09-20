package controllers

import (
	"e-dars/internals/models"
	"e-dars/logger"
	"e-dars/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
			"error": "You do not have permission to see all users",
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

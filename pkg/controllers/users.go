package controllers

import (
	"e-dars/internals/models"
	"e-dars/logger"
	"e-dars/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateNewUser
// @Summary Create User
// @Security ApiKeyAuth
// @Tags users
// @Description create new user
// @ID create-user
// @Accept json
// @Produce json
// @Param input body models.SwagUser true "new operation info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users [post]
func CreateNewUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if c.GetString(userRoleCtx) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to create a new user",
		})

		return
	}

	err := service.CreateNewUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})

}

// GetAllUsers
// @Summary Get All Users
// @Security ApiKeyAuth
// @Tags users
// @Description get list of all users
// @ID get-all-users
// @Produce json
// @Success 200 {array} models.User
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"massage": "No users found"})
	}
	logger.Info.Printf("[controllers] Successfully got all users: %v", users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUserByID
// @Summary Get User By ID
// @Security ApiKeyAuth
// @Tags users
// @Description get user by ID
// @ID get-user-by-id
// @Produce json
// @Param id path integer true "id of the operation"
// @Success 200 {object} models.User
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	var user models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Entered wrong id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if user, err = service.GetUserByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"massage": "User not found"})
		return
	}
	logger.Info.Printf("[controllers] Successfully got user: %v", user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

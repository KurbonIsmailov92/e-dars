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
// @Param input body models.SwagUser true "New User info"
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
		"message": "User created successfully",
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

	if c.GetString(userRoleCtx) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to see all users",
		})
		return
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
// @Param id path integer true "id of the user"
// @Success 200 {object} models.User
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	var user models.User

	if c.GetString(userRoleCtx) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to see user",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Entered wrong id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if user, err = service.GetUserByID(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"massage": "User not found"})
		return
	}
	logger.Info.Printf("[controllers] Successfully got user: %v", user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUser
// @Summary Update user by ID
// @Security ApiKeyAuth
// @Tags users
// @Description Update user by ID
// @ID update-user
// @Accept json
// @Produce json
// @Param id path integer true "id of the user"
// @Param input body models.SwagUser true "User info"
// @Success 200 {object} models.User
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	var user models.User

	if c.GetString(userRoleCtx) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to update user",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err = c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	user.ID = uint(id)

	if err = service.UpdateUser(uint(id), user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"massage": "User not found"})
		return
	}
	logger.Info.Printf("[controllers] Successfully updated user: %v", user)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})
}

// DeActivateUser
// @Summary Deactivate user by ID
// @Security ApiKeyAuth
// @Tags users
// @Description Deactivate user by ID
// @ID deactivate-user
// @Accept json
// @Produce json
// @Param id path integer true "id of the user"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/deactivate/{id} [patch]
func DeActivateUser(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to deactivate user",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}
	if err = service.DeActiveUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	logger.Info.Printf("[controllers] Successfully deactivated user: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "User deactivated!"})
}

// ActivateUser
// @Summary Activate user by ID
// @Security ApiKeyAuth
// @Tags users
// @Description Activate user by ID
// @ID activate-user
// @Accept json
// @Produce json
// @Param id path integer true "id of the user"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/activate/{id} [patch]
func ActivateUser(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to activate user",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}
	if err = service.ActivateUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	logger.Info.Printf("[controllers] Successfully activated user: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "User activated!"})
}

// DeleteUser
// @Summary Delete user by ID
// @Security ApiKeyAuth
// @Tags users
// @Description Delete user by ID
// @ID delete-user
// @Accept json
// @Produce json
// @Param id path integer true "id of the user"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/delete/{id} [delete]
func DeleteUser(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to activate user",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}
	if err = service.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	logger.Info.Printf("[controllers] Successfully deleted user: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "User Deleted!"})
}

// ReturnUser
// @Summary Return user by ID
// @Security ApiKeyAuth
// @Tags users
// @Description Return user by ID
// @ID return-user
// @Accept json
// @Produce json
// @Param id path integer true "id of the user"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/return/{id} [delete]
func ReturnUser(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to activate user",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}
	if err = service.ReturnUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	logger.Info.Printf("[controllers] Successfully returned user: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "User Returned successfully!"})
}

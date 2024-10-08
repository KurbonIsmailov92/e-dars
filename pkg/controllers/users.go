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
// @Router /users/api/v1/ [post]
func CreateNewUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handleError(c, err)
	}

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)

		return
	}

	err := service.CreateNewUser(&user)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.CreateNewUser] User created successfully")
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
// @Success 200 {array} models.SwagUserForShow
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/api/v1/ [get]
func GetAllUsers(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	users, err := service.GetAllUsers()
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetAllUsers] Successfully got all users: %v", users)
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
// @Success 200 {object} models.SwagUserForUpdateByAdmin
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/api/v1/{id} [get]
func GetUserByID(c *gin.Context) {
	var user models.User

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] Entered wrong id: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}

	if user, err = service.GetUserByID(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.GetUserByID] Successfully got user: %v", user)
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
// @Param input body models.SwagUserForUpdateByAdmin true "User info"
// @Success 200 {object} models.SwagUserForUpdateByAdmin
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/api/v1/{id} [put]
func UpdateUser(c *gin.Context) {
	var user models.User

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.UpdateUser] Invalid user ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}

	if err = c.BindJSON(&user); err != nil {
		handleError(c, err)
		return
	}

	user.ID = uint(id)

	if err = service.UpdateUser(user.ID, user); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.UpdateUser] Successfully updated user: %v", user)
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
// @Router /users/api/v1/deactivate/{id} [patch]
func DeActivateUser(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeActivateUser] Invalid user ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}
	if err = service.DeActiveUser(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.DeActivateUser] Successfully deactivated user: %v", id)
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
// @Router/users/api/v1/activate/{id} [patch]
func ActivateUser(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.ActivateUser] Invalid user ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}
	if err = service.ActivateUser(uint(id)); err != nil {
		handleError(c, errs.ErrRecordNotFound)
		return
	}
	logger.Info.Printf("[controllers.ActivateUser] Successfully activated user: %v", id)
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
// @Router /users/api/v1/delete/{id} [delete]
func DeleteUser(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteUser] Invalid user ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}
	if err = service.DeleteUser(uint(id)); err != nil {
		handleError(c, errs.ErrRecordNotFound)
		return
	}
	logger.Info.Printf("[controllers.DeleteUser] Successfully deleted user: %v", id)
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
// @Router /users/api/v1/return/{id} [delete]
func ReturnUser(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.ReturnUser] Invalid user ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}
	if err = service.ReturnUser(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.ReturnUser] Successfully returned user: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "User Returned successfully!"})
}

// ResetUserPasswordByAdmin
// @Summary Reset user`s password to default by ID
// @Security ApiKeyAuth
// @Tags users
// @Description Reset user`s password to default by ID
// @ID reset-user-password
// @Produce json
// @Param id path integer true "id of the user"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/api/v1/reset-password/{id} [patch]
func ResetUserPasswordByAdmin(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.ResetUserPasswordByAdmin] Invalid user ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}
	if err = service.ResetUserPassToDefault(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.ResetUserPasswordByAdmin] Successfully reseted user`s password to default: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "Password reseted successfully!"})
}

// ChangeOwnPasswordByUser
// @Summary Change user`s password to new by User
// @Security ApiKeyAuth
// @Tags users
// @Description Change user`s password to new by User
// @ID change-password
// @Accept json
// @Produce json
// @Param input body models.UserPassword true "User password"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/api/v1/change-password [patch]
func ChangeOwnPasswordByUser(c *gin.Context) {

	var userPasswords models.UserPassword

	id, err := strconv.Atoi(c.GetString(userIDCtx))
	if err != nil {
		logger.Error.Printf("[controllers.ChangeOwnPasswordByUser] Invalid user ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}

	if err = c.BindJSON(&userPasswords); err != nil {
		logger.Error.Printf("[controllers.ChangeOwnPasswordByUser] Invalid JSON: %v", err)
		handleError(c, err)
		return
	}

	if err = service.ChangeOwnPasswordByUser(uint(id), userPasswords.Password,
		userPasswords.OldPassword); err != nil {
		handleError(c, errs.ErrIncorrectOldPassword)
		return
	}
	logger.Info.Printf("[controllers.ChangeOwnPasswordByUser] Successfully changed own password: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

// SetAdminRoleToUser
// @Summary Set Admin Role To User
// @Security ApiKeyAuth
// @Tags users
// @Description Set Admin Role To User by ID
// @ID admin-user
// @Accept json
// @Produce json
// @Param id path integer true "id of the user"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/api/v1/set-admin/{id} [patch]
func SetAdminRoleToUser(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.SetAdminRoleToUser] Invalid user ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}
	if err = service.SetAdminRoleToUser(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.SetAdminRoleToUser] Successfully turned user`s role to Admin")
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})
}

// SetParentToUser
// @Summary Set Parent To User
// @Security ApiKeyAuth
// @Tags users
// @Description Set Parent To User by ID
// @ID parent-user
// @Accept json
// @Produce json
// @Param id path integer true "id of the user"
// @Param input body models.SwagUserForParentSetting true " Parent ID"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/api/v1/set-parent/{id} [patch]
func SetParentToUser(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handleError(c, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.SetParentToUser] Invalid user ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}
	if err = service.SetParentToUser(uint(id), *user.ParentID); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.SetParentToUser] Successfully set user`s parent")
	c.JSON(http.StatusOK, gin.H{"message": "Parent set successfully!"})
}

// SetRoleToUser
// @Summary Set Role To User
// @Security ApiKeyAuth
// @Tags users
// @Description Set Role To User by ID
// @ID role-user
// @Accept json
// @Produce json
// @Param id path integer true "id of the user"
// @Param input body models.SwagUserForRoleSetting true "Role code (teacher, student, parent)"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /users/api/v1/set-role/{id} [patch]
func SetRoleToUser(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handleError(c, err)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.SetRoleToUser] Invalid user ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
	}
	if err = service.SetRoleToUser(uint(id), user.RoleCode); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.SetRoleToUser] Successfully set user`s role")
	c.JSON(http.StatusOK, gin.H{"message": "Role set successfully!"})
}

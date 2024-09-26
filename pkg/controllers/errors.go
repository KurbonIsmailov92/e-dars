package controllers

import (
	"e-dars/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func newErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}

func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUsernameUniquenessFailed),
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword),
		errors.Is(err, errs.ErrIncorrectOldPassword),
		errors.Is(err, errs.ErrFailedValidation):
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrPermissionDenied):
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrClassNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrUserDeactivatedOrDeleted):
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrFailedSetTeacherToClass):
		c.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))

	default:
		c.JSON(http.StatusInternalServerError, newErrorResponse(errs.ErrSomethingWentWrong.Error()))
	}
}

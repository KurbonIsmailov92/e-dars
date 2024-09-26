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

// CreateJournalNote
// @Summary Create Journal Note
// @Security ApiKeyAuth
// @Tags journal
// @Description Create new Journal Note. Time must be in this format "2024-09-28T08:30:00Z"
// @ID create-journal-note
// @Accept json
// @Produce json
// @Param input body models.MarkSetter true "New Journal Note info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /journal/api/v1/ [post]
func CreateJournalNote(c *gin.Context) {
	var note models.MarkSetter

	if c.GetString(userRoleCtx) != "teacher" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.GetString(userIDCtx))
	if err != nil {
		handleError(c, errs.ErrFailedValidation)
		return
	}

	if err = c.BindJSON(&note); err != nil {
		handleError(c, err)
		return
	}

	if err = service.CreateJournalNote(&note, uint(id)); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.CreateJournalNote] Created new journal note")
	newDefaultResponse("Created new journal note")
}

// GetAllJournalNotes
// @Summary Get All Journal Notes
// @Security ApiKeyAuth
// @Tags journal
// @Description Get list of all Journal Notes
// @ID get-all-journal-notes
// @Produce json
// @Success 200 {array} models.JournalNote
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /journal/api/v1/ [get]
func GetAllJournalNotes(c *gin.Context) {
	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	notes, err := service.GetAllJournalNotes()
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetAllJournalNotes] Successfully got all schedule notes")
	c.JSON(http.StatusOK, gin.H{"Journal notes": notes})
}

// GetJournalNoteByID
// @Summary Get Journal Note By ID
// @Security ApiKeyAuth
// @Tags journal
// @Description get journal note by ID
// @ID get-journal-note-by-id
// @Produce json
// @Param id path integer true "id of the journal note"
// @Success 200 {object} models.SwagJournalNotes
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /journal/api/v1/{id} [get]
func GetJournalNoteByID(c *gin.Context) {
	if c.GetString(userRoleCtx) != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	var note models.SwagJournalNotes

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrFailedValidation)
		return
	}

	if note, err = service.GetJournalNoteByID(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.GetJournalNoteByID] Successfully got journal note")
	c.JSON(http.StatusOK, gin.H{"journal note": note})
}

// GetJournalNotesByParentIDAndDate
// @Summary Get Journal Note Of Child
// @Security ApiKeyAuth
// @Tags journal
// @Description  Get Journal Note Of Child. Time must be in this format "2024-09-28"
// @ID get-journal-note-of-child
// @Accept json
// @Produce json
// @Param input body models.JournalDates true "Example: 2024-09-28"
// @Success 200 {array} models.SwagJournalNotes
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /journal/api/v1/notes [post]
func GetJournalNotesByParentIDAndDate(c *gin.Context) {
	if c.GetString(userRoleCtx) != "parent" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	var dates models.JournalDates

	id, err := strconv.Atoi(c.GetString(userIDCtx))
	if err != nil {
		logger.Error.Printf("[controllers.GetJournalNotesByParentIDAndDate] Invalid user ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}

	if err = c.ShouldBindJSON(&dates); err != nil {
		handleError(c, err)
		return
	}

	notes, err := service.GetJournalNotesByParentIDAndDate(uint(id), dates)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetJournalNotesByParentIDAndDate] Successfully got journal notes of child / children")
	c.JSON(http.StatusOK, gin.H{"Journal notes": notes})

}

// GetJournalNotesByStudent
// @Summary Get Own Journal Notes
// @Security ApiKeyAuth
// @Tags journal
// @Description Get Own Journal Notes. Time must be in this format "2024-09-28"
// @ID get-journal-note-of-student
// @Accept json
// @Produce json
// @Param input body models.JournalDates true "Example: 2024-09-28"
// @Success 200 {array} models.SwagJournalNotes
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /journal/api/v1/my-notes [post]
func GetJournalNotesByStudent(c *gin.Context) {
	if c.GetString(userRoleCtx) != "student" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	var dates models.JournalDates

	id, err := strconv.Atoi(c.GetString(userIDCtx))
	if err != nil {
		logger.Error.Printf("[controllers.GetJournalNotesByStudent] Invalid user ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}

	if err = c.ShouldBindJSON(&dates); err != nil {
		handleError(c, err)
		return
	}

	notes, err := service.GetJournalNotesByStudent(uint(id), dates)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetJournalNotesByStudent] Successfully got own journal notes ")
	c.JSON(http.StatusOK, gin.H{"Journal notes": notes})

}

// GetJournalNotesByTeacher
// @Summary Get Journal Notes by teacher
// @Security ApiKeyAuth
// @Tags journal
// @Description Get Journal Notes by teacher. Time must be in this format "2024-09-28"
// @ID get-journal-note-of-teacher
// @Accept json
// @Produce json
// @Param input body models.JournalDates true "Example: 2024-09-28"
// @Success 200 {array} models.SwagJournalNotes
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /journal/api/v1/teacher-notes [post]
func GetJournalNotesByTeacher(c *gin.Context) {

	if c.GetString(userRoleCtx) != "teacher" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	var dates models.JournalDates

	id, err := strconv.Atoi(c.GetString(userIDCtx))
	if err != nil {
		logger.Error.Printf("[controllers.GetJournalNotesByTeacher] Invalid user ID: %v", err)
		handleError(c, errs.ErrFailedValidation)
		return
	}

	if err = c.BindJSON(&dates); err != nil {
		handleError(c, err)
		return
	}

	notes, err := service.GetJournalNotesByTeacher(uint(id), dates)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.GetJournalNotesByStudent] Successfully got own journal notes ")
	c.JSON(http.StatusOK, gin.H{"Journal notes": notes})

}

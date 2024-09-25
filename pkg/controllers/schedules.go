package controllers

import (
	"e-dars/internals/models"
	"e-dars/logger"
	"e-dars/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateNewScheduleNote
// @Summary Create Schedule Note
// @Security ApiKeyAuth
// @Tags schedules
// @Description Create new Schedule Note. Time must be in this format "2024-09-28T08:30:00Z"
// @ID create-schedule-note
// @Accept json
// @Produce json
// @Param input body models.SwagSchedule true "New Schedule Note info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /schedules/api/v1/ [post]
func CreateNewScheduleNote(c *gin.Context) {

	var note models.Schedule

	if c.GetString(userRoleCtx) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to create a new schedule note",
		})

		return
	}

	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := service.CreateNewScheduleNote(&note); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	logger.Info.Printf("[controllers CreateNewScheduleNote] Note created successfully")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Note created successfully",
	})

}

// GetAllScheduleNotes
// @Summary Get All Schedule Notes
// @Security ApiKeyAuth
// @Tags schedules
// @Description Get list of all Schedule Notes
// @ID get-all-schedule-notes
// @Produce json
// @Success 200 {array} models.SwagScheduleForUsers
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /schedules/api/v1/ [get]
func GetAllScheduleNotes(c *gin.Context) {

	if c.GetString(userRoleCtx) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to see all users",
		})
		return
	}

	notes, err := service.GetAllScheduleNotes()

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"massage": "No notes found"})
	}

	logger.Info.Printf("[controllers] Successfully got all schedule notes")
	c.JSON(http.StatusOK, gin.H{"Schedule Notes": notes})
}

// GetScheduleNoteByID
// @Summary Get Schedule Note By ID
// @Security ApiKeyAuth
// @Tags schedules
// @Description get schedule note by ID
// @ID get-schedule-note-by-id
// @Produce json
// @Param id path integer true "id of the user"
// @Success 200 {object} models.SwagSchedule
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /schedules/api/v1/{id} [get]
func GetScheduleNoteByID(c *gin.Context) {
	scheduleNote := models.Schedule{}

	if c.GetString(userRoleCtx) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to see a single schedule note",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Entered wrong id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule note ID"})
		return
	}

	if scheduleNote, err = service.GetScheduleNoteByID(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"massage": "Schedule note not found"})
		return
	}
	logger.Info.Printf("[controllers] Successfully got schedule note")
	c.JSON(http.StatusOK, gin.H{"Schedule note": scheduleNote})
}

// UpdateScheduleNote
// @Summary Update Schedule Note
// @Security ApiKeyAuth
// @Tags schedules
// @Description Update Schedule Note by ID Time must be in this format "2024-09-28T08:30:00Z"
// @ID update-schedule-note
// @Accept json
// @Produce json
// @Param id path integer true "id of the class"
// @Param input body models.SwagSchedule true "Update Schedule Note data "
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /schedules/api/v1/update/{id} [put]
func UpdateScheduleNote(c *gin.Context) {
	if c.GetString(userRoleCtx) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to update Schedule Note",
		})
		return
	}

	var scheduleNote models.Schedule
	if err := c.BindJSON(&scheduleNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid schedule note ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule note ID"})
		return
	}

	scheduleNote.ID = uint(id)

	if err := service.UpdateScheduleNoteByID(uint(id), scheduleNote); err != nil {
		logger.Error.Printf("[controllers] Failed to update schedule note %v: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	logger.Info.Printf("[controllers UpdateScheduleNote] Successfully updated schedule note to %v", scheduleNote.ID)
	c.JSON(http.StatusOK, gin.H{"message": "Schedule note updated successfully"})

}

// DeleteScheduleNote
// @Summary Delete Schedule Note
// @Security ApiKeyAuth
// @Tags schedules
// @Description Delete Schedule Note by ID
// @ID delete-schedule-note
// @Accept json
// @Produce json
// @Param id path integer true "id of the class"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /schedules/api/v1/delete/{id} [delete]
func DeleteScheduleNote(c *gin.Context) {
	if c.GetString(userRoleCtx) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to delete schedule note",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers] Invalid schedule note ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule note ID"})
		return
	}

	if err = service.DeleteScheduleNoteByID(uint(id)); err != nil {
		logger.Error.Printf("[controllers] Failed to delete schedule note %v: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("[controllers DeleteScheduleNote] Successfully deleted schedule note to %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "Schedule note deleted successfully"})
}

// GetTeacherScheduleByDates
// @Summary Get Teachers Schedule Notes
// @Security ApiKeyAuth
// @Tags schedules
// @Description Get Teachers Schedule Notes. Date format must be 2024-09-28
// @ID get-teacher-schedule-notes
// @Produce json
// @Param input body models.ScheduleDates true "Example: 2024-09-28"
// @Success 200 {array} models.SwagScheduleForUsers
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /schedules/api/v1/teacher [post]
func GetTeacherScheduleByDates(c *gin.Context) {
	if c.GetString(userRoleCtx) != "teacher" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to see teachers schedule note",
		})
		return
	}

	id, err := strconv.Atoi(c.GetString(userIDCtx))
	if err != nil {
		logger.Error.Printf("[controllers.GetTeacherScheduleByDates] Invalid ID: %v", err)
		return
	}

	var dates models.ScheduleDates

	if err := c.BindJSON(&dates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	notes, err := service.GetTeacherScheduleByDates(uint(id), dates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	logger.Info.Printf("[controller.GetTeacherScheduleByDates] Successfully got teacher`s schedule notes")
	c.JSON(http.StatusOK, gin.H{"Schedule Notes": notes})

}

// GetStudentScheduleByDates
// @Summary Get Student Schedule Notes
// @Security ApiKeyAuth
// @Tags schedules
// @Description Get Student Schedule Notes. Date format must be 2024-09-28
// @ID get-student-schedule-notes
// @Produce json
// @Param input body models.ScheduleDates true "Example: 2024-09-28"
// @Success 200 {array} models.SwagScheduleForUsers
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /schedules/api/v1/student [post]
func GetStudentScheduleByDates(c *gin.Context) {
	if c.GetString(userRoleCtx) != "student" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to see student schedule note",
		})
		return
	}

	id, err := strconv.Atoi(c.GetString(userIDCtx))
	if err != nil {
		logger.Error.Printf("[controllers.GetStudentScheduleByDates] Invalid ID: %v", err)
		return
	}

	var dates models.ScheduleDates

	if err := c.BindJSON(&dates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	notes, err := service.GetStudentScheduleByDates(uint(id), dates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	logger.Info.Printf("[controller.GetStudentScheduleByDates] Successfully got student`s schedule notes")
	c.JSON(http.StatusOK, gin.H{"Schedule Notes": notes})

}

// GetParentScheduleByDates
// @Summary Get Parent Schedule Notes
// @Security ApiKeyAuth
// @Tags schedules
// @Description Get Parent Schedule Notes. Date format must be 2024-09-28
// @ID get-parent-schedule-notes
// @Produce json
// @Param input body models.ScheduleDates true "Example: 2024-09-28"
// @Success 200 {array} models.SwagScheduleForUsers
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /schedules/api/v1/parent [post]
func GetParentScheduleByDates(c *gin.Context) {
	if c.GetString(userRoleCtx) != "parent" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You do not have permission to see parents schedule note",
		})
		return
	}

	id, err := strconv.Atoi(c.GetString(userIDCtx))
	if err != nil {
		logger.Error.Printf("[controllers.GetParentScheduleByDates] Invalid ID: %v", err)
		return
	}

	var dates models.ScheduleDates

	if err := c.BindJSON(&dates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	notes, err := service.GetParentScheduleByDates(uint(id), dates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	logger.Info.Printf("[controller.GetParentScheduleByDates] Successfully got parent`s schedule notes")
	c.JSON(http.StatusOK, gin.H{"Schedule Notes": notes})

}

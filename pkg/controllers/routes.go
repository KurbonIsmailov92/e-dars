package controllers

import (
	"e-dars/configs"
	_ "e-dars/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRoutes() *gin.Engine {

	router := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", PingPong)

	auth := router.Group("/auth/api/v1")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	UsersG := router.Group("users/api/v1", checkUserAuthentication)
	{
		UsersG.GET("/", GetAllUsers)
		UsersG.GET("/:id", GetUserByID)
		UsersG.POST("/", CreateNewUser)
		UsersG.PUT("/:id", UpdateUser)
		UsersG.PATCH("/deactivate/:id", DeActivateUser)
		UsersG.PATCH("/activate/:id", ActivateUser)
		UsersG.DELETE("/delete/:id", DeleteUser)
		UsersG.DELETE("/return/:id", ReturnUser)
		UsersG.PATCH("/reset-password/:id", ResetUserPasswordByAdmin)
		UsersG.PATCH("/change-password", ChangeOwnPasswordByUser)
		UsersG.PATCH("/set-admin/:id", SetAdminRoleToUser)
		UsersG.PATCH("/set-parent/:id", SetParentToUser)
		UsersG.PATCH("/set-role/:id", SetRoleToUser)
	}

	classesG := router.Group("classes/api/v1", checkUserAuthentication)

	{
		classesG.POST("/", CreateNewClass)
		classesG.GET("/", GetAllClasses)
		classesG.GET("/:id", GetClassByID)
		classesG.POST("/set", SetClassTeacher)
		classesG.PUT("/update/:id", UpdateClass)
		classesG.DELETE("/delete/:id", DeleteClass)
		classesG.DELETE("/return/:id", ReturnClass)
	}

	schedulesG := router.Group("schedules/api/v1", checkUserAuthentication)
	{
		schedulesG.POST("/", CreateNewScheduleNote)
		schedulesG.GET("/", GetAllScheduleNotes)
		schedulesG.GET("/:id", GetScheduleNoteByID)
		schedulesG.PUT("/update/:id", UpdateScheduleNote)
		schedulesG.DELETE("/delete/:id", DeleteScheduleNote)
		schedulesG.POST("/teacher", GetTeacherScheduleByDates)
		schedulesG.POST("/student", GetStudentScheduleByDates)
		schedulesG.POST("/parent", GetParentScheduleByDates)
	}

	journalG := router.Group("journal/api/v1", checkUserAuthentication)
	{
		journalG.POST("/", CreateJournalNote)
		journalG.GET("/", GetAllJournalNotes)
		journalG.GET("/:id", GetJournalNoteByID)
		journalG.POST("/notes", GetJournalNotesByParentIDAndDate)
		journalG.POST("/my-notes", GetJournalNotesByStudent)
		journalG.POST("/teacher-notes", GetJournalNotesByTeacher)
	}

	return router
}

// PingPong
// @Summary Check Connection
// @Tags check
// @Description Show logo if check is OK
// @ID ping
// @Accept json
// @Produce json
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /ping [get]
func PingPong(c *gin.Context) {

	filename := "logo.jpg"
	filepath := "./uploads/" + filename
	c.File(filepath)

	/*c.JSON(http.StatusOK, gin.H{
		"answer": []string{
			" ███████      ██████   █████  ██████   ██████ ",
			" ██           ██   ██ ██   ██ ██   ██ ██      ",
			" █████   ████ ██   ██ ███████ ██████  ███████ ",
			" ██           ██   ██ ██   ██ ██   ██      ██ ",
			" ███████      ██████  ██   ██ ██   ██ ██████  ",
		},
	})*/
}

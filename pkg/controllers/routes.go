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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	usersG := router.Group("/users", checkUserAuthentication)
	{
		usersG.GET("", GetAllUsers)
		usersG.GET(":id", GetUserByID)
		usersG.POST("", CreateNewUser)
		usersG.PUT(":id", UpdateUser)
		usersG.PATCH(":id", DeActivateUser)
		usersG.PATCH("active:id", ActivateUser)

	}

	classesG := router.Group("/classes", checkUserAuthentication)

	{
		classesG.POST("", CreateNewClass)
		classesG.GET("", GetAllClasses)

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

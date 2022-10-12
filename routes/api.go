package routes

import (
	"github.com/gin-gonic/gin"
	controllers "golangapi/controllers"
	services "golangapi/services"
)

var (
	userService    services.UserService       = services.New()
	loginController controllers.LoginController = controllers.New(userService)
)

func CreateRouter() *gin.Engine {
	routes := gin.Default()
	v1 := routes.Group("/api/v1")
	{
		v1.POST("login", loginController.Login)
		v1.GET("/healthcheck", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "OK",
			})
		})

	}

	return routes
}
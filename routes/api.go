package routes

import (
	"github.com/gin-gonic/gin"
	"golangapi/middleware"
	controllers "golangapi/controllers"
	services "golangapi/services"
)

var (
	userService     services.UserService         = services.NewUserService()
	hotelService    services.HotelService       = services.NewHotelService()
	loginController controllers.LoginController = controllers.NewLoginController(userService)
	hotelController controllers.HotelController = controllers.NewHotelController(hotelService)
)

func CreateRouter() *gin.Engine {
	routes := gin.Default()
	v1 := routes.Group("/api/v1")
	{
		v1.POST("login", loginController.Login)
		authorized := v1.Group("/")
		authorized.Use(middleware.Jwt())
		authorized.GET("hotels", hotelController.GetHotels)
		v1.GET("/healthcheck", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "OK",
			})
		})
	}

	return routes
}
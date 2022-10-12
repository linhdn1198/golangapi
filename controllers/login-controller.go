package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golangapi/services"
	"golangapi/util"
)

type LoginController interface {
	Login(ctx *gin.Context)
}

type controller struct {
	service services.UserService
}

func (c *controller) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	user := c.service.GetUser(username)
	if user.Id == 0 {
		ctx.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}
	password := ctx.PostForm("password")
	match := util.CheckPasswordHash(password, user.Password)
	if !match {
		ctx.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}
	accessToken := util.Generate(user.Username)
	ctx.JSON(http.StatusOK, accessToken)
}

func New(service services.UserService) LoginController {
	return &controller{
		service: service,
	}
}
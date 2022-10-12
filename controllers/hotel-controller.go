package controllers

import (
	"golangapi/services"
	"golangapi/util"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"strconv"
	// "golangapi/util"
)

type HotelController interface {
	GetHotels(ctx *gin.Context)
}

type hotelController struct {
	service services.HotelService
}

func (c *hotelController) GetHotels(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	accessToken := strings.Replace(token, "Bearer ", "", 1)
	ok := util.Parse(accessToken)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}
	offset, errOffset := strconv.Atoi(ctx.Request.URL.Query().Get("offset"))
	if errOffset != nil {
		offset = 0
	}
	limit, errLimit := strconv.Atoi(ctx.Request.URL.Query().Get("limit"))
	if errLimit != nil {
		limit = 10
	}

	hotels := c.service.GetHotels(offset, limit)

	ctx.JSON(http.StatusOK, hotels)
}

func NewHotelController(service services.HotelService) HotelController {
	return &hotelController{
		service: service,
	}
}
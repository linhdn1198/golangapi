package controllers

import (
	"golangapi/services"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type HotelController interface {
	GetHotels(ctx *gin.Context)
}

type hotelController struct {
	service services.HotelService
}

func (c *hotelController) GetHotels(ctx *gin.Context) {
	offset, errOffset := strconv.Atoi(ctx.Request.URL.Query().Get("offset"))
	limit, errLimit := strconv.Atoi(ctx.Request.URL.Query().Get("limit"))
	if errOffset != nil {
		offset = 0
	}
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
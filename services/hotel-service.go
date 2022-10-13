package services

import (
	"golangapi/database"
	"golangapi/models"
)

type HotelService interface {
	GetHotels(offset, limit int) []models.Hotel
}

type hotelService struct {
	hotels []models.Hotel
}

func (service *hotelService) GetHotels(offset, limit int) []models.Hotel {
	db := database.GetConnection()
	db.Limit(limit).Offset(offset).Find(&service.hotels)

	return service.hotels
}

func NewHotelService() HotelService {
	return &hotelService{}
}
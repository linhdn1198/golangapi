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
	var Hotels []models.Hotel
	db.Limit(limit).Offset(offset).Find(&Hotels)

	return Hotels
}

func NewHotelService() HotelService {
	return &hotelService{}
}
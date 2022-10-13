package services

import (
	"golangapi/database"
	"golangapi/models"
)

type UserService interface {
	GetUser(username string) models.User
}

type userService struct {
	user models.User
}

func (service *userService) GetUser(username string) models.User {
	db := database.GetConnection()
	db.Where("username = ? ", username).Find(&service.user);

	return service.user
}

func NewUserService() UserService {
	return &userService{}
}
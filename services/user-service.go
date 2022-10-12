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
	var User models.User
	db.Where("username = ? ", username).Find(&User);
	// match := util.CheckPasswordHash(password, User.Password)
	// if match {
	// 	return User
	// }

	return User
}

func New() UserService {
	return &userService{}
}
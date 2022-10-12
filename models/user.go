package models

import (
	"gorm.io/gorm"
)

type User struct {
	Id int `gorm:"primary_key;auto_increment" json:"id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:100;not null;" json:"password"`
}

func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error

	if err != nil {
	   return err
	}

	return nil
 }
package database

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"golangapi/models"
)
 
const DB_USERNAME = "root"
const DB_PASSWORD = "123456"
const DB_NAME = "golang"
const DB_HOST = "localhost"
const DB_PORT = "3306"
 
func GetConnection() (*gorm.DB) {
	dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
 
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("DB Connection established...")

	db.AutoMigrate(&models.User{})
 
	return db
}
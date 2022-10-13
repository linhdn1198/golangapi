package database

import (
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"golangapi/models"
)
 
func GetConnection() (*gorm.DB) {
	dsn := os.Getenv("DB_USERNAME") +":"+ os.Getenv("DB_PASSWORD") +"@tcp"+ "(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") +")/" + os.Getenv("DB_NAME") + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
 
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("DB Connection established...")

	db.AutoMigrate(&models.User{}, &models.Hotel{})
 
	return db
}
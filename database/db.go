package database

import (
	"fmt"
	"go-restapi/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB instance
var DB *gorm.DB

// Connect Database
func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Connection opened to database")

	//Migrate Schema
	err = db.AutoMigrate(&models.User{}, &models.Book{})
	if err != nil {
		panic("Failed to migrate database")
	}
	fmt.Println("Database migrated")

	DB = db
}

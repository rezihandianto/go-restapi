package main

import (
	"log"
	"test-go/database"
	"test-go/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	database.Connect()

	app := fiber.New()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello guys")
	})

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

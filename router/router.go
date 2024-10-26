package router

import (
	"test-go/controllers"
	"test-go/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	//Book
	book := api.Group("/books")
	book.Use(middleware.JWTProtected)
	book.Get("/", controllers.GetBooks)
	book.Get("/:id", controllers.GetBook)
	book.Post("/", controllers.CreateBook)
	book.Put("/:id", controllers.UpdateBooks)
	book.Delete("/:id", controllers.DeleteBooks)

	//Auth
	auth := api.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
}

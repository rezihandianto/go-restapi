package controllers

import (
	"go-restapi/database"
	"go-restapi/models"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Find(&books)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": books,
	})
}
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if err := database.DB.Find(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": book,
	})

}
func CreateBook(c *fiber.Ctx) error {
	var book bookRequest
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	bookModel := models.Book{
		Title:  book.Title,
		Author: book.Author,
	}

	resp := database.DB.Create(&bookModel)

	if resp.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": resp.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Book created successfully",
	})
}
func UpdateBooks(c *fiber.Ctx) error {
	id := c.Params("id")

	var book bookRequest

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var oldBook models.Book
	if err := database.DB.Find(&oldBook, id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
		})
	}

	database.DB.Model(&oldBook).Updates(book)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book updated successfully",
	})

}
func DeleteBooks(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	if err := database.DB.Find(&book, id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
		})
	}

	database.DB.Delete(&book)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})

}

type bookRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}

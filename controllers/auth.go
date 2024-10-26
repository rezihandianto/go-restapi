package controllers

import (
	"test-go/database"
	"test-go/models"
	"test-go/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var req authRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := models.User{
		Email:    req.Email,
		Password: utils.GeneratePassword(req.Password),
	}

	res := database.DB.Create(&user)
	if err := res.Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
func Login(c *fiber.Ctx) error {
	var req authRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if !utils.ComparePassword(user.Password, req.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user": userDataShow{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

type authRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type userDataShow struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package handlers

import (
	"TransactionSystem/internal/connection"
	"TransactionSystem/internal/model/user"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateHandler(c *fiber.Ctx) error {
	newUser := new(user.User)
	newUser.Money = 0
	if err := connection.DB.Create(newUser).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(newUser)
}

package utils

import (
	"github.com/gofiber/fiber/v3"
)

func ResponseJSON(c fiber.Ctx, statusCode int, message string, err string, data any) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"message": message,
		"error":   err,
		"data":    data,
	})
}

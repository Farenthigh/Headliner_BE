package ContactAdapter

import (
	Entities "headliner-be/entities"
	ContactUsecase "headliner-be/usecase/contact"

	"github.com/gofiber/fiber/v3"
)

type ContactHandler struct {
	service ContactUsecase.ContactUsecase
}

func NewContactHandler(s ContactUsecase.ContactUsecase) *ContactHandler {
	return &ContactHandler{service: s}
}

func (h *ContactHandler) CreateContact(c fiber.Ctx) error {
	var contact Entities.Contact

	if err := c.Bind().Body(&contact); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if contact.Subject == "" || contact.Description == "" {
		return c.Status(400).JSON(fiber.Map{"error": "All fields required"})
	}

	err := h.service.CreateContact(&contact)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save"})
	}

	return c.JSON(fiber.Map{"message": "Success"})
}
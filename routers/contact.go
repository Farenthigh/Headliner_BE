package routers

import (
	ContactAdapter "headliner-be/adapter/contact"
	ContactUsecase "headliner-be/usecase/contact"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func InitContactRoute(app *fiber.App, db *gorm.DB) {
	contactRepo := ContactAdapter.NewContactGorm(db)
	contactUsecase := ContactUsecase.NewContactService(contactRepo)
	contactHandler := ContactAdapter.NewContactHandler(contactUsecase)

	contacts := app.Group("/contact")
	contacts.Post("/", contactHandler.CreateContact)
}
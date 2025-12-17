package Entities

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, app *fiber.App) {
	db.AutoMigrate(&Users{})
}

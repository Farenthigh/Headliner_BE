package routers

import (
	UsersAdapter "headliner-be/adapter/users"
	UsersUsecase "headliner-be/usecase/users"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func InitUsersRoute(app *fiber.App, db *gorm.DB) {
	usersRepo := UsersAdapter.NewUsersGorm(db)
	usersUsecase := UsersUsecase.NewUsersService(usersRepo)
	usersHandler := UsersAdapter.NewUsersAdapter(usersUsecase)

	users := app.Group("/users")
	users.Post("/register", usersHandler.Register)
}

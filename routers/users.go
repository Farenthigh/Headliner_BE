package routers

import (
	UsersAdapter "headliner-be/adapter/users"
	UsersUsecase "headliner-be/usecase/users"
	"headliner-be/utils"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func InitUsersRoute(app *fiber.App, db *gorm.DB) {
	usersRepo := UsersAdapter.NewUsersGorm(db)
	usersUsecase := UsersUsecase.NewUsersService(usersRepo)
	usersHandler := UsersAdapter.NewUsersAdapter(usersUsecase)

	users := app.Group("/users")
	users.Post("/register", usersHandler.Register)
	users.Post("/login", usersHandler.Login)
	users.Post("/createcharacter", utils.IsExist, usersHandler.CreateCharacter)
	users.Get("/data", utils.IsExist, usersHandler.GetUserData)
}

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
	users.Post("/register-with-google",utils.FirebaseAuth, usersHandler.RegisterWithGoogle)
	users.Post("/login-with-google", utils.FirebaseAuth, usersHandler.LoginWithGoogle)
	users.Post("/createcharacter", utils.IsExist, usersHandler.CreateCharacter)
	users.Get("/data", utils.IsExist, usersHandler.GetUserData)

	users.Put("/update-username", utils.IsExist, usersHandler.UpdateUsername)
	users.Put("/update-password", utils.IsExist, usersHandler.UpdatePassword)
	users.Post("/update-chatbot", utils.IsExist, usersHandler.UpdateChatbotName)
}

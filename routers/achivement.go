package routers

import (
	AchivementAdapter "headliner-be/adapter/achivement"
	AchivementUsecase "headliner-be/usecase/achivement"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func InitAchivementRoute(app *fiber.App, db *gorm.DB) {
	achivementRepo := AchivementAdapter.NewAchivementGorm(db)
	achivementUsecase := AchivementUsecase.NewAchivementService(achivementRepo)
	achivementHandler := AchivementAdapter.NewAchivementAdapter(achivementUsecase)

	achivements := app.Group("/achivements")
	achivements.Post("/", achivementHandler.CreateAchivement)
	achivements.Get("/", achivementHandler.GetAllAchivements)
	achivements.Get("/:id", achivementHandler.GetAchivementByID)
}
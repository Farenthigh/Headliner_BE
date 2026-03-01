package routers

import (
	AchievementAdapter "headliner-be/adapter/achievement"
	AchievementUsecase "headliner-be/usecase/achievement"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func InitAchievementRoute(app *fiber.App, db *gorm.DB) {
	achievementRepo := AchievementAdapter.NewAchievementGorm(db)
	achievementUsecase := AchievementUsecase.NewAchievementService(achievementRepo)
	achievementHandler := AchievementAdapter.NewAchievementHandler(achievementUsecase)

	achievements := app.Group("/achievements")
	achievements.Post("/", achievementHandler.CreateAchievement)
	achievements.Get("/", achievementHandler.GetAllAchievements)
	achievements.Get("/:id", achievementHandler.GetAchievementByID)
}
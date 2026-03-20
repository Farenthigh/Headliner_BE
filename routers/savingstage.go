package routers

import (
	"headliner-be/adapter/savingstage"
	savingstageUsecase "headliner-be/usecase/savingstage"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func InitSavingStageRoute(app *fiber.App, db *gorm.DB) {

	savingstageUsecase := savingstageUsecase.NewSavingStageUsecase(db)
	savingstageHandler := savingstage.NewSavingStageHandler(savingstageUsecase)

	stageGroup := app.Group("/saving-stage")

	stageGroup.Post("/save", savingstageHandler.SaveStage)
	stageGroup.Get("/leaderboard", savingstageHandler.GetLeaderboard)
	stageGroup.Get("/leaderboard/me", savingstageHandler.GetMyRank)
	stageGroup.Get("/progress", savingstageHandler.GetUserStages)
	stageGroup.Get("/unlock", savingstageHandler.GetUnlockStage)
	stageGroup.Get("/stars", savingstageHandler.GetStageStars)
}

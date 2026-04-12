package routers

import (
	"headliner-be/adapter/stage"
	stageUsecase "headliner-be/usecase/stage"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func InitStageRoute(app *fiber.App, db *gorm.DB) {

	stageUsecase := stageUsecase.NewStageUsecase(db)
	stageHandler := stage.NewStageHandler(stageUsecase)

	stageGroup := app.Group("/stage")

	stageGroup.Post("/save", stageHandler.SaveStage)
	stageGroup.Get("/leaderboard", stageHandler.GetLeaderboard)
	stageGroup.Get("/leaderboard/me", stageHandler.GetMyRank)
	stageGroup.Get("/progress", stageHandler.GetUserStages)
	stageGroup.Get("/unlock", stageHandler.GetUnlockStage)
	stageGroup.Get("/stars", stageHandler.GetStageStars)
}

package routers

import (
	TaxGameScoreAdapter "headliner-be/adapter/taxgamescore"
	TaxGameScoreUsecase "headliner-be/usecase/taxgamescore"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func InitTaxGameScoreRoute(app *fiber.App, db *gorm.DB) {
	taxGameScoreRepo := TaxGameScoreAdapter.NewTaxGameScoreGorm(db)
	taxGameScoreUsecase := TaxGameScoreUsecase.NewTaxGameScoreService(taxGameScoreRepo)
	taxGameScoreHandler := TaxGameScoreAdapter.NewTaxGameScoreHandler(taxGameScoreUsecase)

	taxGameScores := app.Group("/taxgamescores")
	taxGameScores.Post("/", taxGameScoreHandler.CreateTaxGameScore)
	taxGameScores.Get("/", taxGameScoreHandler.GetAllTaxGameScores)
	taxGameScores.Get("/user/:id", taxGameScoreHandler.GetTaxGameScoreByUserID)
	taxGameScores.Put("/", taxGameScoreHandler.UpdateTaxGameScore)
}
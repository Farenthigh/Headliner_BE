package TaxGameScoreAdapter

import (
	TaxGameScoreUsecase "headliner-be/usecase/taxgamescore"

	"github.com/gofiber/fiber/v3"
)

type TaxGameScoreHandler struct {
	taxgamescoreusecase TaxGameScoreUsecase.TaxGameScoreUsecase
}

func NewTaxGameScoreHandler(taxgamescoreusecase TaxGameScoreUsecase.TaxGameScoreUsecase) *TaxGameScoreHandler {
	return &TaxGameScoreHandler{
		taxgamescoreusecase: taxgamescoreusecase,
	}
}

func (a *TaxGameScoreHandler) createTaxGameScore(c fiber.Ctx) error {
	var taxgamescore TaxGameScoreUsecase.
}
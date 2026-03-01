package TaxGameScoreAdapter

import (
	TaxGameScoreModels "headliner-be/model/taxgamescore"
	TaxGameScoreUsecase "headliner-be/usecase/taxgamescore"
	"strconv"

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

func (a *TaxGameScoreHandler) CreateTaxGameScore(c fiber.Ctx) error {
	var createTaxGameScoreInput TaxGameScoreModels.CreateTaxGameScoreInput
	if err := c.Bind().Body(&createTaxGameScoreInput); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	err := a.taxgamescoreusecase.CreateTaxGameScore(&createTaxGameScoreInput)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"message": "Tax game score created successfully",
	})
}

func (a *TaxGameScoreHandler) GetAllTaxGameScores(c fiber.Ctx) error {
	scores, err := a.taxgamescoreusecase.GetAllTaxGameScores()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(scores)
}

func (a *TaxGameScoreHandler) GetTaxGameScoreByUserID(c fiber.Ctx) error {
	userIDParam := c.Params("id")
	userIDValue, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid userID format")
	}
	userID := uint(userIDValue)
	score, err := a.taxgamescoreusecase.GetTaxGameScoreByUserID(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(score)
}

func (a *TaxGameScoreHandler) UpdateTaxGameScore(c fiber.Ctx) error {
	var updateTaxGameScoreInput TaxGameScoreModels.UpdateTaxGameScoreInput
	if err := c.Bind().Body(&updateTaxGameScoreInput); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	err := a.taxgamescoreusecase.UpdateTaxGameScore(&updateTaxGameScoreInput)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"message": "Tax game score updated successfully",
	})
}
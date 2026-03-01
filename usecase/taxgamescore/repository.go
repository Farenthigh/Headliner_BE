package TaxGameScoreUsecase

import (
	Entities "headliner-be/entites"
)

type TaxGameScoreRepository interface {
	CreateTaxGameScore(*Entities.TaxGameScore) error
	GetAllTaxGameScores() ([]*Entities.TaxGameScore, error)
	GetTaxGameScoreByUserID(userID uint) (*Entities.TaxGameScore, error)
	UpdateTaxGameScore(score *Entities.TaxGameScore) error
}
package TaxGameScoreUsecase

import Entities "headliner-be/entites"

type TaxGameScoreRepository interface {
	createTaxGameScore(Entities.TaxGameScore) error
	getAllTaxGameScores() ([]*Entities.TaxGameScore, error)
	getTaxGameScoreByUserID(userID uint) (*Entities.TaxGameScore, error)
	updateTaxGameScore(score *Entities.TaxGameScore) error
}
package TaxGameScoreUsecase

import Entities "headliner-be/entites"

type TaxGameScoreUsecase interface {
	createTaxGameScore(Entities.TaxGameScore) error
	getAllTaxGameScores() ([]*Entities.TaxGameScore, error)
	getTaxGameScoreByUserID(userID uint) (*Entities.TaxGameScore, error)
	updateTaxGameScore(score *Entities.TaxGameScore) error
}

type TaxGameScoreService struct {
	taxgamescorerepo TaxGameScoreRepository
}

func NewTaxGameScoreService(taxgamescorerepo TaxGameScoreRepository) TaxGameScoreUsecase {
	return &TaxGameScoreService{
		taxgamescorerepo: taxgamescorerepo,
	}
}

func (service *TaxGameScoreService) createTaxGameScore(score Entities.TaxGameScore) error {
	return service.taxgamescorerepo.createTaxGameScore(score)
}

func (service *TaxGameScoreService) getAllTaxGameScores() ([]*Entities.TaxGameScore, error) {
	return service.taxgamescorerepo.getAllTaxGameScores()
}

func (service *TaxGameScoreService) getTaxGameScoreByUserID(userID uint) (*Entities.TaxGameScore, error) {
	return service.taxgamescorerepo.getTaxGameScoreByUserID(userID)
}

func (service *TaxGameScoreService) updateTaxGameScore(score *Entities.TaxGameScore) error {
	return service.taxgamescorerepo.updateTaxGameScore(score)
}
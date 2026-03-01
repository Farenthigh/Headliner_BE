package TaxGameScoreUsecase

import (
	Entities "headliner-be/entites"
	TaxGameScoreModels "headliner-be/model/taxgamescore"
)

type TaxGameScoreUsecase interface {
	CreateTaxGameScore(*TaxGameScoreModels.CreateTaxGameScoreInput) error
	GetAllTaxGameScores() ([]*Entities.TaxGameScore, error)
	GetTaxGameScoreByUserID(userID uint) (*Entities.TaxGameScore, error)
	UpdateTaxGameScore(score *TaxGameScoreModels.UpdateTaxGameScoreInput) error
}

type TaxGameScoreService struct {
	taxgamescorerepo TaxGameScoreRepository
}

func NewTaxGameScoreService(taxgamescorerepo TaxGameScoreRepository) TaxGameScoreUsecase {
	return &TaxGameScoreService{
		taxgamescorerepo: taxgamescorerepo,
	}
}

func (service *TaxGameScoreService) CreateTaxGameScore(input *TaxGameScoreModels.CreateTaxGameScoreInput) error {

	return service.taxgamescorerepo.CreateTaxGameScore(&Entities.TaxGameScore{
		UserID: input.UserID,
		Score:  input.Score,
	})
}

func (service *TaxGameScoreService) GetAllTaxGameScores() ([]*Entities.TaxGameScore, error) {
	return service.taxgamescorerepo.GetAllTaxGameScores()
}

func (service *TaxGameScoreService) GetTaxGameScoreByUserID(userID uint) (*Entities.TaxGameScore, error) {
	return service.taxgamescorerepo.GetTaxGameScoreByUserID(userID)
}

func (service *TaxGameScoreService) UpdateTaxGameScore(score *TaxGameScoreModels.UpdateTaxGameScoreInput) error {
	 scoreEntity := Entities.TaxGameScore{}
	 scoreEntity.ID = score.Id
	 scoreEntity.UserID = score.UserID
	 scoreEntity.Score = score.Score
	return service.taxgamescorerepo.UpdateTaxGameScore(&scoreEntity)
}
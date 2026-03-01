package TaxGameScoreAdapter

import (
	Entities "headliner-be/entites"

	"gorm.io/gorm"
)

type TaxGameScoreGorm struct {
	db *gorm.DB
}

func NewTaxGameScoreGorm(db *gorm.DB) *TaxGameScoreGorm {
	return &TaxGameScoreGorm{
		db: db,
	}
}

func (g *TaxGameScoreGorm) CreateTaxGameScore(score *Entities.TaxGameScore) error {
	if err := g.db.Create(&score).Error; err != nil {
		return err
	}
	return nil
}

func (g *TaxGameScoreGorm) GetAllTaxGameScores() ([]*Entities.TaxGameScore, error) {
	var scores []*Entities.TaxGameScore
	if err := g.db.Find(&scores).Error; err != nil {
		return nil, err
	}
	return scores, nil
}

func (g *TaxGameScoreGorm) GetTaxGameScoreByUserID(userID uint) (*Entities.TaxGameScore, error) {
	var score Entities.TaxGameScore
	if err := g.db.Where("user_id = ?", userID).First(&score).Error; err != nil {
		return nil, err
	}
	return &score, nil
}

func (g *TaxGameScoreGorm) UpdateTaxGameScore(score *Entities.TaxGameScore) error {
	if err := g.db.Model(&Entities.TaxGameScore{}).Where("id = ?", score.ID).Updates(score).Error; err != nil {
		return err
	}
	return nil
}
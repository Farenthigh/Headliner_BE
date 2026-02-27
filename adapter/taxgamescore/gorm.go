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

func (g *TaxGameScoreGorm) createTaxGameScore(score *Entities.TaxGameScore) error {
	if err := g.db.Create(&score).Error; err != nil {
		return err
	}
	return nil
}

func (g *TaxGameScoreGorm) getAllTaxGameScores() ([]*Entities.TaxGameScore, error) {
	var scores []*Entities.TaxGameScore
	if err := g.db.Find(&scores).Error; err != nil {
		return nil, err
	}
	return scores, nil
}

func (g *TaxGameScoreGorm) getTaxGameScoreByUserID(userID uint) (*Entities.TaxGameScore, error) {
	var score Entities.TaxGameScore
	if err := g.db.Where("user_id = ?", userID).First(&score).Error; err != nil {
		return nil, err
	}
	return &score, nil
}

func (g *TaxGameScoreGorm) updateTaxGameScore(score *Entities.TaxGameScore) error {
	if err := g.db.Save(&score).Error; err != nil {
		return err
	}
	return nil
}
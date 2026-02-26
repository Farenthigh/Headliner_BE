package AchivementAdapter

import (
	Entities "headliner-be/entites"
	AchivementUsecase "headliner-be/usecase/achivement"

	"gorm.io/gorm"
)

type AchivementGorm struct {
	db *gorm.DB
}

func NewAchivementGorm(db *gorm.DB) AchivementUsecase.AchivementRepository {
	return &AchivementGorm{
		db: db,
	}
}

func (g *AchivementGorm) CreateAchivement(achivement *Entities.Achivement) error {
	if err := g.db.Create(&achivement).Error; err != nil {
		return err
	}
	return nil
}

func (g *AchivementGorm) GetAllAchivements() ([]*Entities.Achivement, error) {
	var achivements []*Entities.Achivement
	if err := g.db.Find(&achivements).Error; err != nil {
		return nil, err
	}
	return achivements, nil
}

func (g *AchivementGorm) GetAchivementByID(achivementID uint) (*Entities.Achivement, error) {
	var achivement Entities.Achivement
	if err := g.db.Where("id = ?", achivementID).First(&achivement).Error; err != nil {
		return nil, err
	}
	return &achivement, nil
}

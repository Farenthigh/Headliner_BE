package AchievementAdapter

import (
	Entities "headliner-be/entities"
	AchievementUsecase "headliner-be/usecase/achievement"

	"gorm.io/gorm"
)

type AchievementGorm struct {
	db *gorm.DB
}

func NewAchievementGorm(db *gorm.DB) AchievementUsecase.AchievementRepository {
	return &AchievementGorm{
		db: db,
	}
}

func (g *AchievementGorm) CreateAchievement(achievement *Entities.Achievement) error {
	if err := g.db.Create(achievement).Error; err != nil {
		return err
	}
	return nil
}

func (g *AchievementGorm) GetAllAchievements() ([]*Entities.Achievement, error) {
	var achievements []*Entities.Achievement
	if err := g.db.Find(achievements).Error; err != nil {
		return nil, err
	}
	return achievements, nil
}

func (g *AchievementGorm) GetAchievementByID(achievementID uint) (*Entities.Achievement, error) {
	var achievement Entities.Achievement
	if err := g.db.Where("id = ?", achievementID).First(&achievement).Error; err != nil {
		return nil, err
	}
	return &achievement, nil
}

func (g *AchievementGorm) SaveUserAchievement(userAch *Entities.UserAchievement) error {
	if err := g.db.FirstOrCreate(userAch, Entities.UserAchievement{
		UserID: userAch.UserID, 
		AchievementID: userAch.AchievementID,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (g *AchievementGorm) GetUserAchievements(userID uint) ([]*Entities.UserAchievement, error) {
	var userAchievements []*Entities.UserAchievement
	if err := g.db.Preload("Achievement").Where("user_id = ?", userID).Find(&userAchievements).Error; err != nil {
		return nil, err
	}
	return userAchievements, nil
}
package AchievementUsecase

import Entities "headliner-be/entites"

type AchievementUsecase interface {
	CreateAchievement(*Entities.Achievement) error
	GetAllAchievements() ([]*Entities.Achievement, error)
	GetAchievementByID(achievementID uint) (*Entities.Achievement, error)
}

type achievementService struct {
	repo AchievementRepository
}

func NewAchievementService(achievementRepo AchievementRepository) AchievementUsecase {
	return &achievementService{
		repo: achievementRepo,
	}
}

func (service *achievementService) CreateAchievement(achievement *Entities.Achievement) error {
	return service.repo.CreateAchievement(achievement)
}

func (service *achievementService) GetAllAchievements() ([]*Entities.Achievement, error) {
	return service.repo.GetAllAchievements()
}

func (service *achievementService) GetAchievementByID(achievementID uint) (*Entities.Achievement, error) {
	return service.repo.GetAchievementByID(achievementID)
}
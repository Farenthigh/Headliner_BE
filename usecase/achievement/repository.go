package AchievementUsecase

import Entities "headliner-be/entites"

type AchievementRepository interface {
	CreateAchievement(*Entities.Achievement) error
	GetAllAchievements() ([]*Entities.Achievement, error)
	GetAchievementByID(achievementID uint) (*Entities.Achievement, error)
}
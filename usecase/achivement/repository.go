package AchivementUsecase

import Entities "headliner-be/entites"

type AchivementRepository interface {
	CreateAchivement(*Entities.Achivement) error
	GetAllAchivements() ([]*Entities.Achivement, error)
	GetAchivementByID(achivementID uint) (*Entities.Achivement, error)
}
package AchivementUsecase

import Entities "headliner-be/entites"

type AchivementUsecase interface {
	CreateAchivement(*Entities.Achivement) error
	GetAllAchivements() ([]*Entities.Achivement, error)
	GetAchivementByID(achivementID uint) (*Entities.Achivement, error)
}

type achivementService struct {
	repo AchivementRepository
}

func NewAchivementService(achivementrepo AchivementRepository) AchivementUsecase {
	return &achivementService{
		repo: achivementrepo,
	}
}

func (service *achivementService) CreateAchivement(achivement *Entities.Achivement) error {
	return service.repo.CreateAchivement(achivement)
}

func (service *achivementService) GetAllAchivements() ([]*Entities.Achivement, error) {
	return service.repo.GetAllAchivements()
}

func (service *achivementService) GetAchivementByID(achivementID uint) (*Entities.Achivement, error) {
	return service.repo.GetAchivementByID(achivementID)
}
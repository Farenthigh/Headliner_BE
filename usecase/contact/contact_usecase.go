package ContactUsecase

import Entities "headliner-be/entities"

type ContactUsecase interface {
	CreateContact(contact *Entities.Contact) error
}

type contactService struct {
	repo ContactRepository
}

func NewContactService(contactRepo ContactRepository) ContactUsecase {
	return &contactService{
		repo: contactRepo,
	}
}

func (service *contactService) CreateContact(contact *Entities.Contact) error {
	return service.repo.CreateContact(contact)
}
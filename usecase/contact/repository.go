package ContactUsecase

import Entities "headliner-be/entities"

type ContactRepository interface {
	CreateContact(contact *Entities.Contact) error
}
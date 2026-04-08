package ContactAdapter

import (
	Entities "headliner-be/entities"
	ContactUsecase "headliner-be/usecase/contact"

	"gorm.io/gorm"
)

type ContactGorm struct {
	db *gorm.DB
}

func NewContactGorm(db *gorm.DB) ContactUsecase.ContactRepository {
	return &ContactGorm{db: db}
}

func (g *ContactGorm) CreateContact(contact *Entities.Contact) error {
	return g.db.Create(contact).Error
}
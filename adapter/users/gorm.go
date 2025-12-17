package UsersAdapter

import (
	"errors"
	"fmt"
	Entities "headliner-be/entites"
	UsersUsecase "headliner-be/usecase/users"

	"gorm.io/gorm"
)

type UsersGorm struct {
	db *gorm.DB
}

func NewUsersGorm(db *gorm.DB) UsersUsecase.UsersRepository {
	return &UsersGorm{db: db}
}

func (g *UsersGorm) Register(Users *Entities.Users) error {
	if err := g.db.Create(&Users).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("email %s already exists", Users.Email)
		}
		return err
	}
	return nil
}

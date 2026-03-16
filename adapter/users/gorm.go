package UsersAdapter

import (
	"errors"
	"fmt"
	Entities "headliner-be/entities"
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

func (g *UsersGorm) GetUserByEmail(email string) (*Entities.Users, error) {
	var user Entities.Users
	if err := g.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *UsersGorm) CreateCharacter(Users *Entities.Users) error {
	if err := g.db.Model(&Entities.Users{}).Where("id = ?", Users.ID).Updates(map[string]interface{}{
		"username":  Users.Username,
		"character": Users.Character,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (g *UsersGorm) GetUserByUsername(username string) (*Entities.Users, error) {
	var user Entities.Users
	if err := g.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (g *UsersGorm) GetUserByID(userID uint) (*Entities.Users, error) {
	var user Entities.Users
	if err := g.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (g *UsersGorm) UpdateUsername(userID uint, username string) error {

	if err := g.db.Model(&Entities.Users{}).
		Where("id = ?", userID).
		Update("username", username).Error; err != nil {

		return err
	}

	return nil
}
func (g *UsersGorm) UpdatePassword(userID uint, password string) error {

	if err := g.db.Model(&Entities.Users{}).
		Where("id = ?", userID).
		Update("password", password).Error; err != nil {

		return err
	}

	return nil
}
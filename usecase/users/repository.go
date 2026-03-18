package UsersUsecase

import Entities "headliner-be/entities"



type UsersRepository interface {
	Register(*Entities.Users) error
	GetUserByEmail(email string) (*Entities.Users, error)
	CreateCharacter(*Entities.Users) error
	GetUserByUsername(username string) (*Entities.Users, error)
	GetUserByID(userID uint) (*Entities.Users, error)
	UpdateUsername(userID uint, username string) error
	UpdatePassword(userID uint, password string) error
}

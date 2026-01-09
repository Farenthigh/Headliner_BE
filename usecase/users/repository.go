package UsersUsecase

import Entities "headliner-be/entites"

type UsersRepository interface {
	Register(*Entities.Users) error
	GetUserByEmail(email string) (*Entities.Users, error)
	CreateCharacter(*Entities.Users) error
	GetUserByUsername(username string) (*Entities.Users, error)
	GetUserByID(userID uint) (*Entities.Users, error)
}

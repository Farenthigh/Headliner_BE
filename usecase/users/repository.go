package UsersUsecase

import Entities "headliner-be/entites"

type UsersRepository interface {
	Register(*Entities.Users) error
	GetUserByEmail(email string) (*Entities.Users, error)
}

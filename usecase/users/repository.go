package UsersUsecase

import Entities "headliner-be/entites"

type UsersRepository interface {
	Register(*Entities.Users) error
}

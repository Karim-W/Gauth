package services

import (
	"log"
	"gauth.com/Dbmanager/models"
	repos "gauth.com/Repos"
)

type UserService interface {
	GetUser(id string) (user models.User, err error)
	CreateNewUser(user models.User) (models.User, error)
}
type userService struct {
	userRepository repos.UserRepository
}

func NewUserService(r repos.UserRepository) UserService {
	return userService{
		userRepository: r,
	}
}

func (u userService) GetUser(id string) (user models.User, err error) {
	log.Print("[UserService]...Get User")
	userRes, err := u.userRepository.GetUser(id)
	return userRes, err
}

func (u userService) CreateNewUser(user models.User) (models.User, error) {
	log.Print("[UserService]...Create User")
	user, err := u.userRepository.CreateUser(user)
	return user, err
}

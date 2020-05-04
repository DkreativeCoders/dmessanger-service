package service

import (
	"github.com/danieloluwadare/dmessanger/interfaces/irepository"
	"github.com/danieloluwadare/dmessanger/interfaces/iservice"
	"github.com/danieloluwadare/dmessanger/models"
)

func NewService(repository irepository.IUserRepository) iservice.IUserService {
	return service{repository}
}

type service struct {
	repository irepository.IUserRepository
}

func (s service) CreateUser(user models.User) (*models.User, error) {

	return &user, nil
}

func (s service) GetUser(id int) (*models.User, error) {
	var user models.User
	return &user, nil

}

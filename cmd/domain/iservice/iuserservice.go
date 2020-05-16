package iservice

import (
	"github.com/danieloluwadare/dmessanger/cmd/domain"
)

type IUserService interface {
	GetUser(id int) (*domain.User, error)
	GetAllUser()map[string]interface{}
	CreateUser(user domain.User) map[string]interface{}
}

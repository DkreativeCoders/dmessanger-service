package iservice

import "github.com/danieloluwadare/dmessanger/models"

type IUserService interface {
	GetUser(id int) (*models.User, error)
	GetAllUser()map[string]interface{}
	CreateUser(user models.User) map[string]interface{}
}

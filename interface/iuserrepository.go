package interface

import (
	"github.com/danieloluwadare/dmessanger/models"
)

type IUserRepository interface {
	GetUser(id int) (*models.User, error)
	CreateUser(user models.User) error
}
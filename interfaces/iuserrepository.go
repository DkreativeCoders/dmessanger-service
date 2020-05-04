package interfaces

import "github.com/danieloluwadare/dmessanger/models"

type IUserRepository interface {
	GetUser(id int) *models.User
	CreateUser(user models.User) error
}

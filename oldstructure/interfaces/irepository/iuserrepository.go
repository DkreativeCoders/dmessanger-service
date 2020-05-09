package irepository

import "github.com/danieloluwadare/dmessanger/oldstructure/models"

// IUserRepository
// Define your basic crud operation in this interface

type IUserRepository interface {
	FindByID(id int) *models.User
	FindAll() []models.User
	Save(user models.User) (*models.User, error)
}

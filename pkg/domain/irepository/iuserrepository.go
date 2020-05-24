package irepository

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
)

// IUserRepository
// Define your basic crud operation in this interface

type IUserRepository interface {
	FindByID(id int) (*domain.User, error)
	FindAll() []domain.User
	Save(user domain.User) (*domain.User, error)
	Update(user domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindUserExist(email string) bool
}

package irepository

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
)

// IUserRepository
// Define your basic crud operation in this interface

type IUserRepository interface {
	FindByID(id int) *domain.User
	FindAll() []domain.User
	Save(user domain.User) (*domain.User, error)
}

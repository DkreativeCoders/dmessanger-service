package irepository

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
)

// IUserRepository
// Define your basic crud operation in this interface

type ITokenRepository interface {
	Create(token domain.Token) (*domain.Token, error)
	FindByUserId(userId int) (*domain.Customer, error)
}
